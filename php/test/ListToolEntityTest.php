<?php
declare(strict_types=1);

// ListTool entity test

require_once __DIR__ . '/../seqbenchmcp_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class ListToolEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = SeqbenchMcpSDK::test(null, null);
        $ent = $testsdk->ListTool(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = list_tool_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "list_tool." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set SEQBENCHMCP_TEST_LIST_TOOL_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $list_tool_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.list_tool")));
        $list_tool_ref01_data = null;
        if (count($list_tool_ref01_data_raw) > 0) {
            $list_tool_ref01_data = Helpers::to_map($list_tool_ref01_data_raw[0][1]);
        }

        // LOAD
        $list_tool_ref01_ent = $client->ListTool(null);
        $list_tool_ref01_match_dt0 = [];
        $list_tool_ref01_data_dt0_loaded = $list_tool_ref01_ent->load($list_tool_ref01_match_dt0, null);
        $this->assertNotNull($list_tool_ref01_data_dt0_loaded);

    }
}

function list_tool_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/list_tool/ListToolTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = SeqbenchMcpSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["list_tool01", "list_tool02", "list_tool03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("SEQBENCHMCP_TEST_LIST_TOOL_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "SEQBENCHMCP_TEST_LIST_TOOL_ENTID" => $idmap,
        "SEQBENCHMCP_TEST_LIVE" => "FALSE",
        "SEQBENCHMCP_TEST_EXPLAIN" => "FALSE",
        "SEQBENCHMCP_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["SEQBENCHMCP_TEST_LIST_TOOL_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["SEQBENCHMCP_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["SEQBENCHMCP_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new SeqbenchMcpSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["SEQBENCHMCP_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["SEQBENCHMCP_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
