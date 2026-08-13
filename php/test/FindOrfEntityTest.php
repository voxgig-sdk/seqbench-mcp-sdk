<?php
declare(strict_types=1);

// FindOrf entity test

require_once __DIR__ . '/../seqbenchmcp_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class FindOrfEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = SeqbenchMcpSDK::test(null, null);
        $ent = $testsdk->FindOrf(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = find_orf_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "find_orf." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set SEQBENCH_MCP_TEST_FIND_ORF_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $find_orf_ref01_ent = $client->FindOrf(null);
        $find_orf_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.find_orf"), "find_orf_ref01"));

        $find_orf_ref01_data_result = $find_orf_ref01_ent->create($find_orf_ref01_data, null);
        $find_orf_ref01_data = Helpers::to_map(is_object($find_orf_ref01_data_result) && method_exists($find_orf_ref01_data_result, 'data_get') ? $find_orf_ref01_data_result->data_get() : $find_orf_ref01_data_result);
        $this->assertNotNull($find_orf_ref01_data);

    }
}

function find_orf_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/find_orf/FindOrfTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = SeqbenchMcpSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["find_orf01", "find_orf02", "find_orf03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("SEQBENCH_MCP_TEST_FIND_ORF_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "SEQBENCH_MCP_TEST_FIND_ORF_ENTID" => $idmap,
        "SEQBENCH_MCP_TEST_LIVE" => "FALSE",
        "SEQBENCH_MCP_TEST_EXPLAIN" => "FALSE",
        "SEQBENCH_MCP_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["SEQBENCH_MCP_TEST_FIND_ORF_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["SEQBENCH_MCP_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["SEQBENCH_MCP_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new SeqbenchMcpSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["SEQBENCH_MCP_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["SEQBENCH_MCP_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
