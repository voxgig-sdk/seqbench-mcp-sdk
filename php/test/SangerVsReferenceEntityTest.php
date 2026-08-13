<?php
declare(strict_types=1);

// SangerVsReference entity test

require_once __DIR__ . '/../seqbenchmcp_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class SangerVsReferenceEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = SeqbenchMcpSDK::test(null, null);
        $ent = $testsdk->SangerVsReference(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = sanger_vs_reference_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "sanger_vs_reference." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set SEQBENCH_MCP_TEST_SANGER_VS_REFERENCE_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $sanger_vs_reference_ref01_ent = $client->SangerVsReference(null);
        $sanger_vs_reference_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.sanger_vs_reference"), "sanger_vs_reference_ref01"));

        $sanger_vs_reference_ref01_data_result = $sanger_vs_reference_ref01_ent->create($sanger_vs_reference_ref01_data, null);
        $sanger_vs_reference_ref01_data = Helpers::to_map(is_object($sanger_vs_reference_ref01_data_result) && method_exists($sanger_vs_reference_ref01_data_result, 'data_get') ? $sanger_vs_reference_ref01_data_result->data_get() : $sanger_vs_reference_ref01_data_result);
        $this->assertNotNull($sanger_vs_reference_ref01_data);

    }
}

function sanger_vs_reference_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/sanger_vs_reference/SangerVsReferenceTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = SeqbenchMcpSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["sanger_vs_reference01", "sanger_vs_reference02", "sanger_vs_reference03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("SEQBENCH_MCP_TEST_SANGER_VS_REFERENCE_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "SEQBENCH_MCP_TEST_SANGER_VS_REFERENCE_ENTID" => $idmap,
        "SEQBENCH_MCP_TEST_LIVE" => "FALSE",
        "SEQBENCH_MCP_TEST_EXPLAIN" => "FALSE",
        "SEQBENCH_MCP_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["SEQBENCH_MCP_TEST_SANGER_VS_REFERENCE_ENTID"]);
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
