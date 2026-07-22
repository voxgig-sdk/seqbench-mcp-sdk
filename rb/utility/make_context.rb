# SeqbenchMcp SDK utility: make_context
require_relative '../core/context'
module SeqbenchMcpUtilities
  MakeContext = ->(ctxmap, basectx) {
    SeqbenchMcpContext.new(ctxmap, basectx)
  }
end
