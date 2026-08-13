# ReverseComplement entity test

require "minitest/autorun"
require "json"
require_relative "../SeqbenchMcp_sdk"
require_relative "runner"

class ReverseComplementEntityTest < Minitest::Test
  def test_create_instance
    testsdk = SeqbenchMcpSDK.test(nil, nil)
    ent = testsdk.ReverseComplement(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = reverse_complement_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "reverse_complement." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set SEQBENCH_MCP_TEST_REVERSE_COMPLEMENT_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    reverse_complement_ref01_ent = client.ReverseComplement(nil)
    reverse_complement_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.reverse_complement"), "reverse_complement_ref01"))

    reverse_complement_ref01_data_result = reverse_complement_ref01_ent.create(reverse_complement_ref01_data, nil)
    reverse_complement_ref01_data = Helpers.to_map(reverse_complement_ref01_data_result.respond_to?(:data_get) ? reverse_complement_ref01_data_result.data_get : reverse_complement_ref01_data_result)
    assert !reverse_complement_ref01_data.nil?

  end
end

def reverse_complement_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "reverse_complement", "ReverseComplementTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = SeqbenchMcpSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["reverse_complement01", "reverse_complement02", "reverse_complement03"],
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
  entid_env_raw = ENV["SEQBENCH_MCP_TEST_REVERSE_COMPLEMENT_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "SEQBENCH_MCP_TEST_REVERSE_COMPLEMENT_ENTID" => idmap,
    "SEQBENCH_MCP_TEST_LIVE" => "FALSE",
    "SEQBENCH_MCP_TEST_EXPLAIN" => "FALSE",
    "SEQBENCH_MCP_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["SEQBENCH_MCP_TEST_REVERSE_COMPLEMENT_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["SEQBENCH_MCP_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
        "apikey" => env["SEQBENCH_MCP_APIKEY"],
      },
      extra || {},
    ])
    client = SeqbenchMcpSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["SEQBENCH_MCP_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["SEQBENCH_MCP_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
