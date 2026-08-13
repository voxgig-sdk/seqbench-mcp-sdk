<?php
declare(strict_types=1);

// CodonAdaptationIndex entity test

require_once __DIR__ . '/../seqbenchmcp_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class CodonAdaptationIndexEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = SeqbenchMcpSDK::test(null, null);
        $ent = $testsdk->CodonAdaptationIndex(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = codon_adaptation_index_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "codon_adaptation_index." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set SEQBENCH_MCP_TEST_CODON_ADAPTATION_INDEX_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $codon_adaptation_index_ref01_ent = $client->CodonAdaptationIndex(null);
        $codon_adaptation_index_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.codon_adaptation_index"), "codon_adaptation_index_ref01"));

        $codon_adaptation_index_ref01_data_result = $codon_adaptation_index_ref01_ent->create($codon_adaptation_index_ref01_data, null);
        $codon_adaptation_index_ref01_data = Helpers::to_map(is_object($codon_adaptation_index_ref01_data_result) && method_exists($codon_adaptation_index_ref01_data_result, 'data_get') ? $codon_adaptation_index_ref01_data_result->data_get() : $codon_adaptation_index_ref01_data_result);
        $this->assertNotNull($codon_adaptation_index_ref01_data);

    }
}

function codon_adaptation_index_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/codon_adaptation_index/CodonAdaptationIndexTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = SeqbenchMcpSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["codon_adaptation_index01", "codon_adaptation_index02", "codon_adaptation_index03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("SEQBENCH_MCP_TEST_CODON_ADAPTATION_INDEX_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "SEQBENCH_MCP_TEST_CODON_ADAPTATION_INDEX_ENTID" => $idmap,
        "SEQBENCH_MCP_TEST_LIVE" => "FALSE",
        "SEQBENCH_MCP_TEST_EXPLAIN" => "FALSE",
        "SEQBENCH_MCP_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["SEQBENCH_MCP_TEST_CODON_ADAPTATION_INDEX_ENTID"]);
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
