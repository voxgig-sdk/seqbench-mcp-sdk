<?php
declare(strict_types=1);

// SeqbenchMcp SDK utility: result_headers

class SeqbenchMcpResultHeaders
{
    public static function call(SeqbenchMcpContext $ctx): ?SeqbenchMcpResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
