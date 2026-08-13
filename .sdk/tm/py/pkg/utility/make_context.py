# SeqbenchMcp SDK utility: make_context

from projectname_sdk.core.context import SeqbenchMcpContext


def make_context_util(ctxmap, basectx):
    return SeqbenchMcpContext(ctxmap, basectx)
