package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/seqbench-mcp-sdk/go"
	"github.com/voxgig-sdk/seqbench-mcp-sdk/go/core"

	vs "github.com/voxgig-sdk/seqbench-mcp-sdk/go/utility/struct"
)

func TestSequenceFetchEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.SequenceFetch(nil)
		if ent == nil {
			t.Fatal("expected non-nil SequenceFetchEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := sequence_fetchBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "sequence_fetch." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set SEQBENCH_MCP_TEST_SEQUENCE_FETCH_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		sequenceFetchRef01Ent := client.SequenceFetch(nil)
		sequenceFetchRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "sequence_fetch"}, setup.data), "sequence_fetch_ref01"))

		sequenceFetchRef01DataResult, err := sequenceFetchRef01Ent.Create(sequenceFetchRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		sequenceFetchRef01Data = core.ToMapAny(entityData(sequenceFetchRef01DataResult))
		if sequenceFetchRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func sequence_fetchBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "sequence_fetch", "SequenceFetchTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read sequence_fetch test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse sequence_fetch test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"sequence_fetch01", "sequence_fetch02", "sequence_fetch03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("SEQBENCH_MCP_TEST_SEQUENCE_FETCH_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"SEQBENCH_MCP_TEST_SEQUENCE_FETCH_ENTID": idmap,
		"SEQBENCH_MCP_TEST_LIVE":      "FALSE",
		"SEQBENCH_MCP_TEST_EXPLAIN":   "FALSE",
		"SEQBENCH_MCP_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["SEQBENCH_MCP_TEST_SEQUENCE_FETCH_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["SEQBENCH_MCP_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["SEQBENCH_MCP_APIKEY"],
			},
			extra,
		})
		client = sdk.NewSeqbenchMcpSDK(core.ToMapAny(mergedOpts))
	}

	live := env["SEQBENCH_MCP_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["SEQBENCH_MCP_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
