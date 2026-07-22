# SeqbenchMcp SDK exists test

require "minitest/autorun"
require_relative "../SeqbenchMcp_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = SeqbenchMcpSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
