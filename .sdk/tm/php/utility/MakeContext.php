<?php
declare(strict_types=1);

// SeqbenchMcp SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class SeqbenchMcpMakeContext
{
    public static function call(array $ctxmap, ?SeqbenchMcpContext $basectx): SeqbenchMcpContext
    {
        return new SeqbenchMcpContext($ctxmap, $basectx);
    }
}
