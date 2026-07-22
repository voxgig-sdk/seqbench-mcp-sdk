<?php
declare(strict_types=1);

// SeqbenchMcp SDK base feature

class SeqbenchMcpBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(SeqbenchMcpContext $ctx, array $options): void {}
    public function PostConstruct(SeqbenchMcpContext $ctx): void {}
    public function PostConstructEntity(SeqbenchMcpContext $ctx): void {}
    public function SetData(SeqbenchMcpContext $ctx): void {}
    public function GetData(SeqbenchMcpContext $ctx): void {}
    public function GetMatch(SeqbenchMcpContext $ctx): void {}
    public function SetMatch(SeqbenchMcpContext $ctx): void {}
    public function PrePoint(SeqbenchMcpContext $ctx): void {}
    public function PreSpec(SeqbenchMcpContext $ctx): void {}
    public function PreRequest(SeqbenchMcpContext $ctx): void {}
    public function PreResponse(SeqbenchMcpContext $ctx): void {}
    public function PreResult(SeqbenchMcpContext $ctx): void {}
    public function PreDone(SeqbenchMcpContext $ctx): void {}
    public function PreUnexpected(SeqbenchMcpContext $ctx): void {}
}
