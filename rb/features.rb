# SeqbenchMcp SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module SeqbenchMcpFeatures
  def self.make_feature(name)
    case name
    when "base"
      SeqbenchMcpBaseFeature.new
    when "ratelimit"
      SeqbenchMcpRatelimitFeature.new
    when "retry"
      SeqbenchMcpRetryFeature.new
    when "test"
      SeqbenchMcpTestFeature.new
    when "timeout"
      SeqbenchMcpTimeoutFeature.new
    else
      SeqbenchMcpBaseFeature.new
    end
  end
end
