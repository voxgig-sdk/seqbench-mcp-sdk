<?php
declare(strict_types=1);

// CharacterizeSequence entity test

require_once __DIR__ . '/../seqbenchmcp_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class CharacterizeSequenceEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = SeqbenchMcpSDK::test(null, null);
        $ent = $testsdk->CharacterizeSequence(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = characterize_sequence_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "characterize_sequence." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set SEQBENCHMCP_TEST_CHARACTERIZE_SEQUENCE_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $characterize_sequence_ref01_ent = $client->CharacterizeSequence(null);
        $characterize_sequence_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.characterize_sequence"), "characterize_sequence_ref01"));

        $characterize_sequence_ref01_data_result = $characterize_sequence_ref01_ent->create($characterize_sequence_ref01_data, null);
        $characterize_sequence_ref01_data = Helpers::to_map($characterize_sequence_ref01_data_result);
        $this->assertNotNull($characterize_sequence_ref01_data);

    }
}

function characterize_sequence_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/characterize_sequence/CharacterizeSequenceTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = SeqbenchMcpSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["characterize_sequence01", "characterize_sequence02", "characterize_sequence03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("SEQBENCHMCP_TEST_CHARACTERIZE_SEQUENCE_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "SEQBENCHMCP_TEST_CHARACTERIZE_SEQUENCE_ENTID" => $idmap,
        "SEQBENCHMCP_TEST_LIVE" => "FALSE",
        "SEQBENCHMCP_TEST_EXPLAIN" => "FALSE",
        "SEQBENCHMCP_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["SEQBENCHMCP_TEST_CHARACTERIZE_SEQUENCE_ENTID"]);
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
