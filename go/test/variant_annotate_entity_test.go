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

func TestVariantAnnotateEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.VariantAnnotate(nil)
		if ent == nil {
			t.Fatal("expected non-nil VariantAnnotateEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := variant_annotateBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "variant_annotate." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set SEQBENCHMCP_TEST_VARIANT_ANNOTATE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		variantAnnotateRef01Ent := client.VariantAnnotate(nil)
		variantAnnotateRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "variant_annotate"}, setup.data), "variant_annotate_ref01"))

		variantAnnotateRef01DataResult, err := variantAnnotateRef01Ent.Create(variantAnnotateRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		variantAnnotateRef01Data = core.ToMapAny(variantAnnotateRef01DataResult)
		if variantAnnotateRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func variant_annotateBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "variant_annotate", "VariantAnnotateTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read variant_annotate test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse variant_annotate test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"variant_annotate01", "variant_annotate02", "variant_annotate03"},
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
	entidEnvRaw := os.Getenv("SEQBENCHMCP_TEST_VARIANT_ANNOTATE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"SEQBENCHMCP_TEST_VARIANT_ANNOTATE_ENTID": idmap,
		"SEQBENCHMCP_TEST_LIVE":      "FALSE",
		"SEQBENCHMCP_TEST_EXPLAIN":   "FALSE",
		"SEQBENCHMCP_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["SEQBENCHMCP_TEST_VARIANT_ANNOTATE_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["SEQBENCHMCP_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["SEQBENCHMCP_APIKEY"],
			},
			extra,
		})
		client = sdk.NewSeqbenchMcpSDK(core.ToMapAny(mergedOpts))
	}

	live := env["SEQBENCHMCP_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["SEQBENCHMCP_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
