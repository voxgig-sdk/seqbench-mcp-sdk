<?php
declare(strict_types=1);

// SeqbenchMcp SDK utility: prepare_path

class SeqbenchMcpPreparePath
{
    public static function call(SeqbenchMcpContext $ctx): string
    {
        $point = $ctx->point;
        $parts = [];
        if ($point) {
            $p = \Voxgig\Struct\Struct::getprop($point, 'parts');
            if (is_array($p)) {
                $parts = $p;
            }
        }
        return \Voxgig\Struct\Struct::join($parts, '/', true);
    }
}
