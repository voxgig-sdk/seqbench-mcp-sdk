<?php
declare(strict_types=1);

// SeqbenchMcp SDK exists test

require_once __DIR__ . '/../seqbenchmcp_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = SeqbenchMcpSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
