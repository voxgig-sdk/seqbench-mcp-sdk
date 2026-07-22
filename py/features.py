# SeqbenchMcp SDK feature factory

from feature.base_feature import SeqbenchMcpBaseFeature
from feature.test_feature import SeqbenchMcpTestFeature


def _make_feature(name):
    features = {
        "base": lambda: SeqbenchMcpBaseFeature(),
        "test": lambda: SeqbenchMcpTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
