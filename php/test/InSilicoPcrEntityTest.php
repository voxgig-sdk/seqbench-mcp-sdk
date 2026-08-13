<?php
declare(strict_types=1);

// InSilicoPcr entity test

require_once __DIR__ . '/../seqbenchmcp_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class InSilicoPcrEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = SeqbenchMcpSDK::test(null, null);
        $ent = $testsdk->InSilicoPcr(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = in_silico_pcr_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "in_silico_pcr." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set SEQBENCH_MCP_TEST_IN_SILICO_PCR_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $in_silico_pcr_ref01_ent = $client->InSilicoPcr(null);
        $in_silico_pcr_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.in_silico_pcr"), "in_silico_pcr_ref01"));

        $in_silico_pcr_ref01_data_result = $in_silico_pcr_ref01_ent->create($in_silico_pcr_ref01_data, null);
        $in_silico_pcr_ref01_data = Helpers::to_map(is_object($in_silico_pcr_ref01_data_result) && method_exists($in_silico_pcr_ref01_data_result, 'data_get') ? $in_silico_pcr_ref01_data_result->data_get() : $in_silico_pcr_ref01_data_result);
        $this->assertNotNull($in_silico_pcr_ref01_data);

    }
}

function in_silico_pcr_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/in_silico_pcr/InSilicoPcrTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = SeqbenchMcpSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["in_silico_pcr01", "in_silico_pcr02", "in_silico_pcr03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("SEQBENCH_MCP_TEST_IN_SILICO_PCR_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "SEQBENCH_MCP_TEST_IN_SILICO_PCR_ENTID" => $idmap,
        "SEQBENCH_MCP_TEST_LIVE" => "FALSE",
        "SEQBENCH_MCP_TEST_EXPLAIN" => "FALSE",
        "SEQBENCH_MCP_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["SEQBENCH_MCP_TEST_IN_SILICO_PCR_ENTID"]);
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
