# SeqbenchMcp SDK utility: make_context

from seqbenchmcp_sdk.core.context import SeqbenchMcpContext


def make_context_util(ctxmap, basectx):
    return SeqbenchMcpContext(ctxmap, basectx)
