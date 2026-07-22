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

func TestExportPlateLayoutEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.ExportPlateLayout(nil)
		if ent == nil {
			t.Fatal("expected non-nil ExportPlateLayoutEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := export_plate_layoutBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "export_plate_layout." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set SEQBENCHMCP_TEST_EXPORT_PLATE_LAYOUT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		exportPlateLayoutRef01Ent := client.ExportPlateLayout(nil)
		exportPlateLayoutRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "export_plate_layout"}, setup.data), "export_plate_layout_ref01"))

		exportPlateLayoutRef01DataResult, err := exportPlateLayoutRef01Ent.Create(exportPlateLayoutRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		exportPlateLayoutRef01Data = core.ToMapAny(exportPlateLayoutRef01DataResult)
		if exportPlateLayoutRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func export_plate_layoutBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "export_plate_layout", "ExportPlateLayoutTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read export_plate_layout test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse export_plate_layout test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"export_plate_layout01", "export_plate_layout02", "export_plate_layout03"},
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
	entidEnvRaw := os.Getenv("SEQBENCHMCP_TEST_EXPORT_PLATE_LAYOUT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"SEQBENCHMCP_TEST_EXPORT_PLATE_LAYOUT_ENTID": idmap,
		"SEQBENCHMCP_TEST_LIVE":      "FALSE",
		"SEQBENCHMCP_TEST_EXPLAIN":   "FALSE",
		"SEQBENCHMCP_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["SEQBENCHMCP_TEST_EXPORT_PLATE_LAYOUT_ENTID"])
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
