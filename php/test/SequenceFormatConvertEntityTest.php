<?php
declare(strict_types=1);

// SequenceFormatConvert entity test

require_once __DIR__ . '/../seqbenchmcp_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class SequenceFormatConvertEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = SeqbenchMcpSDK::test(null, null);
        $ent = $testsdk->SequenceFormatConvert(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = sequence_format_convert_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "sequence_format_convert." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set SEQBENCH_MCP_TEST_SEQUENCE_FORMAT_CONVERT_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $sequence_format_convert_ref01_ent = $client->SequenceFormatConvert(null);
        $sequence_format_convert_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.sequence_format_convert"), "sequence_format_convert_ref01"));

        $sequence_format_convert_ref01_data_result = $sequence_format_convert_ref01_ent->create($sequence_format_convert_ref01_data, null);
        $sequence_format_convert_ref01_data = Helpers::to_map(is_object($sequence_format_convert_ref01_data_result) && method_exists($sequence_format_convert_ref01_data_result, 'data_get') ? $sequence_format_convert_ref01_data_result->data_get() : $sequence_format_convert_ref01_data_result);
        $this->assertNotNull($sequence_format_convert_ref01_data);

    }
}

function sequence_format_convert_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/sequence_format_convert/SequenceFormatConvertTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = SeqbenchMcpSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["sequence_format_convert01", "sequence_format_convert02", "sequence_format_convert03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("SEQBENCH_MCP_TEST_SEQUENCE_FORMAT_CONVERT_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "SEQBENCH_MCP_TEST_SEQUENCE_FORMAT_CONVERT_ENTID" => $idmap,
        "SEQBENCH_MCP_TEST_LIVE" => "FALSE",
        "SEQBENCH_MCP_TEST_EXPLAIN" => "FALSE",
        "SEQBENCH_MCP_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["SEQBENCH_MCP_TEST_SEQUENCE_FORMAT_CONVERT_ENTID"]);
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
