-- SeqbenchMcp SDK exists test

local sdk = require("seqbench-mcp_sdk")

describe("SeqbenchMcpSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
