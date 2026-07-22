# SangerVsReference entity test

require "minitest/autorun"
require "json"
require_relative "../SeqbenchMcp_sdk"
require_relative "runner"

class SangerVsReferenceEntityTest < Minitest::Test
  def test_create_instance
    testsdk = SeqbenchMcpSDK.test(nil, nil)
    ent = testsdk.SangerVsReference(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = sanger_vs_reference_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "sanger_vs_reference." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set SEQBENCHMCP_TEST_SANGER_VS_REFERENCE_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    sanger_vs_reference_ref01_ent = client.SangerVsReference(nil)
    sanger_vs_reference_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.sanger_vs_reference"), "sanger_vs_reference_ref01"))

    sanger_vs_reference_ref01_data_result = sanger_vs_reference_ref01_ent.create(sanger_vs_reference_ref01_data, nil)
    sanger_vs_reference_ref01_data = Helpers.to_map(sanger_vs_reference_ref01_data_result)
    assert !sanger_vs_reference_ref01_data.nil?

  end
end

def sanger_vs_reference_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "sanger_vs_reference", "SangerVsReferenceTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = SeqbenchMcpSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["sanger_vs_reference01", "sanger_vs_reference02", "sanger_vs_reference03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["SEQBENCHMCP_TEST_SANGER_VS_REFERENCE_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "SEQBENCHMCP_TEST_SANGER_VS_REFERENCE_ENTID" => idmap,
    "SEQBENCHMCP_TEST_LIVE" => "FALSE",
    "SEQBENCHMCP_TEST_EXPLAIN" => "FALSE",
    "SEQBENCHMCP_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["SEQBENCHMCP_TEST_SANGER_VS_REFERENCE_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["SEQBENCHMCP_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
        "apikey" => env["SEQBENCHMCP_APIKEY"],
      },
      extra || {},
    ])
    client = SeqbenchMcpSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["SEQBENCHMCP_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["SEQBENCHMCP_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
