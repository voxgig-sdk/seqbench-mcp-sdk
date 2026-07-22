<?php
declare(strict_types=1);

// SeqbenchMcp SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class SeqbenchMcpFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new SeqbenchMcpBaseFeature();
            case "test":
                return new SeqbenchMcpTestFeature();
            default:
                return new SeqbenchMcpBaseFeature();
        }
    }
}
