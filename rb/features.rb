# SeqbenchMcp SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module SeqbenchMcpFeatures
  def self.make_feature(name)
    case name
    when "base"
      SeqbenchMcpBaseFeature.new
    when "test"
      SeqbenchMcpTestFeature.new
    else
      SeqbenchMcpBaseFeature.new
    end
  end
end
