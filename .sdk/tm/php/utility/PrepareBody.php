<?php
declare(strict_types=1);

// SeqbenchMcp SDK utility: prepare_body

class SeqbenchMcpPrepareBody
{
    public static function call(SeqbenchMcpContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
