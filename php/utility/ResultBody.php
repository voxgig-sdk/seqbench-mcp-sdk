<?php
declare(strict_types=1);

// SeqbenchMcp SDK utility: result_body

class SeqbenchMcpResultBody
{
    public static function call(SeqbenchMcpContext $ctx): ?SeqbenchMcpResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
