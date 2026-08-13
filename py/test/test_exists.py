# SeqbenchMcp SDK exists test

import pytest
from seqbenchmcp_sdk import SeqbenchMcpSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = SeqbenchMcpSDK.test(None, None)
        assert testsdk is not None
