# BaseEditingDesign entity test

require "minitest/autorun"
require "json"
require_relative "../SeqbenchMcp_sdk"
require_relative "runner"

class BaseEditingDesignEntityTest < Minitest::Test
  def test_create_instance
    testsdk = SeqbenchMcpSDK.test(nil, nil)
    ent = testsdk.BaseEditingDesign(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = base_editing_design_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "base_editing_design." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set SEQBENCHMCP_TEST_BASE_EDITING_DESIGN_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    base_editing_design_ref01_ent = client.BaseEditingDesign(nil)
    base_editing_design_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.base_editing_design"), "base_editing_design_ref01"))

    base_editing_design_ref01_data_result = base_editing_design_ref01_ent.create(base_editing_design_ref01_data, nil)
    base_editing_design_ref01_data = Helpers.to_map(base_editing_design_ref01_data_result)
    assert !base_editing_design_ref01_data.nil?

  end
end

def base_editing_design_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "base_editing_design", "BaseEditingDesignTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = SeqbenchMcpSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["base_editing_design01", "base_editing_design02", "base_editing_design03"],
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
  entid_env_raw = ENV["SEQBENCHMCP_TEST_BASE_EDITING_DESIGN_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "SEQBENCHMCP_TEST_BASE_EDITING_DESIGN_ENTID" => idmap,
    "SEQBENCHMCP_TEST_LIVE" => "FALSE",
    "SEQBENCHMCP_TEST_EXPLAIN" => "FALSE",
    "SEQBENCHMCP_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["SEQBENCHMCP_TEST_BASE_EDITING_DESIGN_ENTID"])
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
