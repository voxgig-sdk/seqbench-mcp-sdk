# SeqbenchMcp PHP SDK



The PHP SDK for the SeqbenchMcp API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->AlphafoldLookup()` — with named operations (`load`/`create`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/seqbench-mcp-sdk/releases](https://github.com/voxgig-sdk/seqbench-mcp-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'seqbenchmcp_sdk.php';

$client = new SeqbenchMcpSDK([
    "apikey" => getenv("SEQBENCH_MCP_APIKEY"),
]);
```

### 4. Create, update, and remove

```php
// create() returns the bare created AlphafoldLookup record.
$created = $client->AlphafoldLookup()->create(["accession" => "example_accession", "ok" => "example_ok", "provenance" => [], "result" => [], "tool" => "example_tool"]);

```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $batch = $client->Batch()->load();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required:

```php
$client = SeqbenchMcpSDK::test();

// Entity ops return the bare mock record (throws on error).
$batch = $client->Batch()->load();
print_r($batch);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new SeqbenchMcpSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
SEQBENCH_MCP_TEST_LIVE=TRUE
SEQBENCH_MCP_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### SeqbenchMcpSDK

```php
require_once 'seqbenchmcp_sdk.php';
$client = new SeqbenchMcpSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = SeqbenchMcpSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### SeqbenchMcpSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `AlphafoldLookup` | `($data): AlphafoldLookupEntity` | Create an AlphafoldLookup entity instance. |
| `AsoDesign` | `($data): AsoDesignEntity` | Create an AsoDesign entity instance. |
| `BaseEditingDesign` | `($data): BaseEditingDesignEntity` | Create a BaseEditingDesign entity instance. |
| `Batch` | `($data): BatchEntity` | Create a Batch entity instance. |
| `BatchWorkflow` | `($data): BatchWorkflowEntity` | Create a BatchWorkflow entity instance. |
| `CharacterizeSequence` | `($data): CharacterizeSequenceEntity` | Create a CharacterizeSequence entity instance. |
| `CloningSimulate` | `($data): CloningSimulateEntity` | Create a CloningSimulate entity instance. |
| `CodonAdaptationIndex` | `($data): CodonAdaptationIndexEntity` | Create a CodonAdaptationIndex entity instance. |
| `CodonOptimize` | `($data): CodonOptimizeEntity` | Create a CodonOptimize entity instance. |
| `ConstructAutofix` | `($data): ConstructAutofixEntity` | Create a ConstructAutofix entity instance. |
| `ConstructQc` | `($data): ConstructQcEntity` | Create a ConstructQc entity instance. |
| `CrisprGrnaDesign` | `($data): CrisprGrnaDesignEntity` | Create a CrisprGrnaDesign entity instance. |
| `CrisprHdrDonor` | `($data): CrisprHdrDonorEntity` | Create a CrisprHdrDonor entity instance. |
| `CrisprOfftargetCheck` | `($data): CrisprOfftargetCheckEntity` | Create a CrisprOfftargetCheck entity instance. |
| `CrossDimer` | `($data): CrossDimerEntity` | Create a CrossDimer entity instance. |
| `DnaMolarity` | `($data): DnaMolarityEntity` | Create a DnaMolarity entity instance. |
| `DoubleDigest` | `($data): DoubleDigestEntity` | Create a DoubleDigest entity instance. |
| `ExportEchoPicklist` | `($data): ExportEchoPicklistEntity` | Create an ExportEchoPicklist entity instance. |
| `ExportOpentronsProtocol` | `($data): ExportOpentronsProtocolEntity` | Create an ExportOpentronsProtocol entity instance. |
| `ExportPlateLayout` | `($data): ExportPlateLayoutEntity` | Create an ExportPlateLayout entity instance. |
| `ExpressionHeatmapCluster` | `($data): ExpressionHeatmapClusterEntity` | Create an ExpressionHeatmapCluster entity instance. |
| `FastqQcReport` | `($data): FastqQcReportEntity` | Create a FastqQcReport entity instance. |
| `FastqTrim` | `($data): FastqTrimEntity` | Create a FastqTrim entity instance. |
| `FindOrf` | `($data): FindOrfEntity` | Create a FindOrf entity instance. |
| `FormatSequence` | `($data): FormatSequenceEntity` | Create a FormatSequence entity instance. |
| `FunctionalEnrichment` | `($data): FunctionalEnrichmentEntity` | Create a FunctionalEnrichment entity instance. |
| `GcContent` | `($data): GcContentEntity` | Create a GcContent entity instance. |
| `GeneDossier` | `($data): GeneDossierEntity` | Create a GeneDossier entity instance. |
| `GeneExpression` | `($data): GeneExpressionEntity` | Create a GeneExpression entity instance. |
| `GeneModel` | `($data): GeneModelEntity` | Create a GeneModel entity instance. |
| `GoldenGateFidelity` | `($data): GoldenGateFidelityEntity` | Create a GoldenGateFidelity entity instance. |
| `HgvsConvert` | `($data): HgvsConvertEntity` | Create a HgvsConvert entity instance. |
| `IdMapPoll` | `($data): IdMapPollEntity` | Create an IdMapPoll entity instance. |
| `IdMapSubmit` | `($data): IdMapSubmitEntity` | Create an IdMapSubmit entity instance. |
| `InSilicoPcr` | `($data): InSilicoPcrEntity` | Create an InSilicoPcr entity instance. |
| `KaspPrimerDesign` | `($data): KaspPrimerDesignEntity` | Create a KaspPrimerDesign entity instance. |
| `ListTool` | `($data): ListToolEntity` | Create a ListTool entity instance. |
| `MeltingTemperature` | `($data): MeltingTemperatureEntity` | Create a MeltingTemperature entity instance. |
| `MotifFinder` | `($data): MotifFinderEntity` | Create a MotifFinder entity instance. |
| `MultipleSequenceAlignment` | `($data): MultipleSequenceAlignmentEntity` | Create a MultipleSequenceAlignment entity instance. |
| `OligoAnalysi` | `($data): OligoAnalysiEntity` | Create an OligoAnalysi entity instance. |
| `OrthologMap` | `($data): OrthologMapEntity` | Create an OrthologMap entity instance. |
| `PairwiseAlignment` | `($data): PairwiseAlignmentEntity` | Create a PairwiseAlignment entity instance. |
| `ParseGenbank` | `($data): ParseGenbankEntity` | Create a ParseGenbank entity instance. |
| `ParseSangerTrace` | `($data): ParseSangerTraceEntity` | Create a ParseSangerTrace entity instance. |
| `PlasmidAnnotate` | `($data): PlasmidAnnotateEntity` | Create a PlasmidAnnotate entity instance. |
| `PlasmidDeepAnnotate` | `($data): PlasmidDeepAnnotateEntity` | Create a PlasmidDeepAnnotate entity instance. |
| `PlasmidFullReport` | `($data): PlasmidFullReportEntity` | Create a PlasmidFullReport entity instance. |
| `PlasmidIdentify` | `($data): PlasmidIdentifyEntity` | Create a PlasmidIdentify entity instance. |
| `PrimeEditingDesign` | `($data): PrimeEditingDesignEntity` | Create a PrimeEditingDesign entity instance. |
| `PrimeEditingTwinDesign` | `($data): PrimeEditingTwinDesignEntity` | Create a PrimeEditingTwinDesign entity instance. |
| `PrimerDesign` | `($data): PrimerDesignEntity` | Create a PrimerDesign entity instance. |
| `PrimerSpecificity` | `($data): PrimerSpecificityEntity` | Create a PrimerSpecificity entity instance. |
| `ProteaseDigestion` | `($data): ProteaseDigestionEntity` | Create a ProteaseDigestion entity instance. |
| `ProteinAnnotatePoll` | `($data): ProteinAnnotatePollEntity` | Create a ProteinAnnotatePoll entity instance. |
| `ProteinAnnotateSubmit` | `($data): ProteinAnnotateSubmitEntity` | Create a ProteinAnnotateSubmit entity instance. |
| `ProteinHydrophobicity` | `($data): ProteinHydrophobicityEntity` | Create a ProteinHydrophobicity entity instance. |
| `ProteinProperty` | `($data): ProteinPropertyEntity` | Create a ProteinProperty entity instance. |
| `RandomSequence` | `($data): RandomSequenceEntity` | Create a RandomSequence entity instance. |
| `RestrictionSite` | `($data): RestrictionSiteEntity` | Create a RestrictionSite entity instance. |
| `ReverseComplement` | `($data): ReverseComplementEntity` | Create a ReverseComplement entity instance. |
| `ReverseTranslate` | `($data): ReverseTranslateEntity` | Create a ReverseTranslate entity instance. |
| `RnaFold` | `($data): RnaFoldEntity` | Create a RnaFold entity instance. |
| `SangerVsReference` | `($data): SangerVsReferenceEntity` | Create a SangerVsReference entity instance. |
| `SavePermalink` | `($data): SavePermalinkEntity` | Create a SavePermalink entity instance. |
| `SeqfileStat` | `($data): SeqfileStatEntity` | Create a SeqfileStat entity instance. |
| `SequenceFetch` | `($data): SequenceFetchEntity` | Create a SequenceFetch entity instance. |
| `SequenceFormatConvert` | `($data): SequenceFormatConvertEntity` | Create a SequenceFormatConvert entity instance. |
| `SequenceReport` | `($data): SequenceReportEntity` | Create a SequenceReport entity instance. |
| `SequenceSearch` | `($data): SequenceSearchEntity` | Create a SequenceSearch entity instance. |
| `SequencingReadbackVerify` | `($data): SequencingReadbackVerifyEntity` | Create a SequencingReadbackVerify entity instance. |
| `SessionCreate` | `($data): SessionCreateEntity` | Create a SessionCreate entity instance. |
| `SessionGet` | `($data): SessionGetEntity` | Create a SessionGet entity instance. |
| `SessionRun` | `($data): SessionRunEntity` | Create a SessionRun entity instance. |
| `SessionSet` | `($data): SessionSetEntity` | Create a SessionSet entity instance. |
| `SirnaDesign` | `($data): SirnaDesignEntity` | Create a SirnaDesign entity instance. |
| `SiteDirectedMutagenesi` | `($data): SiteDirectedMutagenesiEntity` | Create a SiteDirectedMutagenesi entity instance. |
| `Translate` | `($data): TranslateEntity` | Create a Translate entity instance. |
| `VariantAnnotate` | `($data): VariantAnnotateEntity` | Create a VariantAnnotate entity instance. |
| `VariantComparator` | `($data): VariantComparatorEntity` | Create a VariantComparator entity instance. |
| `VerifyAssembly` | `($data): VerifyAssemblyEntity` | Create a VerifyAssembly entity instance. |
| `VerifyConstruct` | `($data): VerifyConstructEntity` | Create a VerifyConstruct entity instance. |
| `VirtualGel` | `($data): VirtualGelEntity` | Create a VirtualGel entity instance. |
| `VolcanoPlotData` | `($data): VolcanoPlotDataEntity` | Create a VolcanoPlotData entity instance. |
| `WebSearch` | `($data): WebSearchEntity` | Create a WebSearch entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the bare result data (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

### Entities

#### AlphafoldLookup

| Field | Description |
| --- | --- |
| `accession` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/alphafold_lookup`

#### AsoDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `length` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `target` |  |
| `tool` |  |
| `wing` |  |

Operations: Create.

API path: `/aso_design`

#### BaseEditingDesign

| Field | Description |
| --- | --- |
| `editor` |  |
| `frame_start` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `target` |  |
| `target_position` |  |
| `tool` |  |

Operations: Create.

API path: `/base_editing_design`

#### Batch

| Field | Description |
| --- | --- |
| `arg` |  |
| `input` |  |
| `ok` |  |
| `result` |  |
| `tool` |  |

Operations: Create, Load.

API path: `/batch`

#### BatchWorkflow

| Field | Description |
| --- | --- |
| `input` |  |
| `ok` |  |
| `result` |  |
| `step` |  |

Operations: Create, Load.

API path: `/workflow`

#### CharacterizeSequence

| Field | Description |
| --- | --- |
| `end_primer_length` |  |
| `gate` |  |
| `max_orf` |  |
| `min_orf_aa` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/characterize_sequence`

#### CloningSimulate

| Field | Description |
| --- | --- |
| `arm_tm_target` |  |
| `circular` |  |
| `enzyme` |  |
| `enzyme3` |  |
| `enzyme5` |  |
| `fragment` |  |
| `gate` |  |
| `insert` |  |
| `method` |  |
| `name` |  |
| `ok` |  |
| `overlap_len` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `vector` |  |

Operations: Create.

API path: `/cloning_simulate`

#### CodonAdaptationIndex

| Field | Description |
| --- | --- |
| `frame_start` |  |
| `gate` |  |
| `ok` |  |
| `organism` |  |
| `provenance` |  |
| `rare_threshold` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/codon_adaptation_index`

#### CodonOptimize

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `organism` |  |
| `protein` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/codon_optimize`

#### ConstructAutofix

| Field | Description |
| --- | --- |
| `avoid_enzyme` |  |
| `cryptic_orf_min_aa` |  |
| `frame_start` |  |
| `gate` |  |
| `gc_high` |  |
| `gc_low` |  |
| `gc_window` |  |
| `homopolymer_min` |  |
| `max_pass` |  |
| `ok` |  |
| `organism` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/construct_autofix`

#### ConstructQc

| Field | Description |
| --- | --- |
| `avoid_enzyme` |  |
| `cryptic_orf_min_aa` |  |
| `frame_start` |  |
| `gate` |  |
| `gc_high` |  |
| `gc_low` |  |
| `gc_window` |  |
| `homopolymer_min` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/construct_qc`

#### CrisprGrnaDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `min_score` |  |
| `nuclease` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `search_reverse_strand` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/crispr_grna_design`

#### CrisprHdrDonor

| Field | Description |
| --- | --- |
| `arm_length` |  |
| `block_pam` |  |
| `design_genotyping_primer` |  |
| `edit_end` |  |
| `edit_start` |  |
| `frame_start` |  |
| `gate` |  |
| `guide_end` |  |
| `guide_start` |  |
| `guide_strand` |  |
| `nuclease` |  |
| `ok` |  |
| `provenance` |  |
| `replacement` |  |
| `result` |  |
| `target_sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/crispr_hdr_donor`

#### CrisprOfftargetCheck

| Field | Description |
| --- | --- |
| `gate` |  |
| `max_mismatch` |  |
| `nuclease` |  |
| `ok` |  |
| `protospacer` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/crispr_offtarget_check`

#### CrossDimer

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence_a` |  |
| `sequence_b` |  |
| `tool` |  |

Operations: Create.

API path: `/cross_dimer`

#### DnaMolarity

| Field | Description |
| --- | --- |
| `gate` |  |
| `length` |  |
| `mass_ng` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |
| `type` |  |
| `volume_ul` |  |

Operations: Create.

API path: `/dna_molarity`

#### DoubleDigest

| Field | Description |
| --- | --- |
| `enzyme_a` |  |
| `enzyme_b` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/double_digest`

#### ExportEchoPicklist

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `reaction` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/export_echo_picklist`

#### ExportOpentronsProtocol

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `protocol_name` |  |
| `provenance` |  |
| `reaction` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/export_opentrons_protocol`

#### ExportPlateLayout

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `reaction` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/export_plate_layout`

#### ExpressionHeatmapCluster

| Field | Description |
| --- | --- |
| `cluster_col` |  |
| `cluster_row` |  |
| `distance_metric` |  |
| `gate` |  |
| `gene` |  |
| `linkage` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sample` |  |
| `tool` |  |
| `value` |  |
| `z_score_row` |  |

Operations: Create.

API path: `/expression_heatmap_cluster`

#### FastqQcReport

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `ok` |  |
| `provenance` |  |
| `quality_offset` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/fastq_qc_report`

#### FastqTrim

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `min_length` |  |
| `ok` |  |
| `provenance` |  |
| `quality_offset` |  |
| `quality_threshold` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/fastq_trim`

#### FindOrf

| Field | Description |
| --- | --- |
| `gate` |  |
| `min_aa_length` |  |
| `ok` |  |
| `provenance` |  |
| `require_stop` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/find_orfs`

#### FormatSequence

| Field | Description |
| --- | --- |
| `case_mode` |  |
| `convert` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reverse` |  |
| `sequence` |  |
| `strip_non_letter` |  |
| `tool` |  |
| `width` |  |

Operations: Create.

API path: `/format_sequence`

#### FunctionalEnrichment

| Field | Description |
| --- | --- |
| `background` |  |
| `collection` |  |
| `gate` |  |
| `gene` |  |
| `max_term_size` |  |
| `min_term_size` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/functional_enrichment`

#### GcContent

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/gc_content`

#### GeneDossier

| Field | Description |
| --- | --- |
| `gate` |  |
| `gene` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/gene_dossier`

#### GeneExpression

| Field | Description |
| --- | --- |
| `gate` |  |
| `gene` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/gene_expression`

#### GeneModel

| Field | Description |
| --- | --- |
| `gate` |  |
| `gene` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/gene_model`

#### GoldenGateFidelity

| Field | Description |
| --- | --- |
| `compare_to_named_set` |  |
| `dataset` |  |
| `gate` |  |
| `ok` |  |
| `overhang` |  |
| `provenance` |  |
| `result` |  |
| `risk_threshold` |  |
| `tool` |  |

Operations: Create.

API path: `/golden_gate_fidelity`

#### HgvsConvert

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `variant` |  |

Operations: Create.

API path: `/hgvs_convert`

#### IdMapPoll

| Field | Description |
| --- | --- |
| `gate` |  |
| `job_id` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/id_map_poll`

#### IdMapSubmit

| Field | Description |
| --- | --- |
| `from` |  |
| `gate` |  |
| `ids` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tax_id` |  |
| `to` |  |
| `tool` |  |

Operations: Create.

API path: `/id_map_submit`

#### InSilicoPcr

| Field | Description |
| --- | --- |
| `circular` |  |
| `forward_primer` |  |
| `gate` |  |
| `max_mismatch` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reverse_primer` |  |
| `template` |  |
| `tool` |  |

Operations: Create.

API path: `/in_silico_pcr`

#### KaspPrimerDesign

| Field | Description |
| --- | --- |
| `add_secondary_mismatch` |  |
| `allele_a` |  |
| `allele_b` |  |
| `gate` |  |
| `max_amplicon` |  |
| `min_amplicon` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `snp_position` |  |
| `target` |  |
| `target_core_tm` |  |
| `tool` |  |

Operations: Create.

API path: `/kasp_primer_design`

#### ListTool

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/`

#### MeltingTemperature

| Field | Description |
| --- | --- |
| `dntp_mm` |  |
| `gate` |  |
| `mg_mm` |  |
| `na_mm` |  |
| `ok` |  |
| `oligo_nm` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `target_tm` |  |
| `tm_tolerance` |  |
| `tool` |  |

Operations: Create.

API path: `/melting_temperature`

#### MotifFinder

| Field | Description |
| --- | --- |
| `gate` |  |
| `max_mismatch` |  |
| `motif` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `search_reverse_strand` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/motif_finder`

#### MultipleSequenceAlignment

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/multiple_sequence_alignment`

#### OligoAnalysi

| Field | Description |
| --- | --- |
| `dntp_mm` |  |
| `gate` |  |
| `mg_mm` |  |
| `na_mm` |  |
| `ok` |  |
| `oligo_nm` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/oligo_analysis`

#### OrthologMap

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `source_species` |  |
| `symbol` |  |
| `target_species` |  |
| `tool` |  |
| `type` |  |

Operations: Create.

API path: `/ortholog_map`

#### PairwiseAlignment

| Field | Description |
| --- | --- |
| `gap` |  |
| `gate` |  |
| `match` |  |
| `mismatch` |  |
| `mode` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `seq_a` |  |
| `seq_b` |  |
| `tool` |  |

Operations: Create.

API path: `/pairwise_alignment`

#### ParseGenbank

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `text` |  |
| `tool` |  |

Operations: Create.

API path: `/parse_genbank`

#### ParseSangerTrace

| Field | Description |
| --- | --- |
| `file_base64` |  |
| `file_name` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/parse_sanger_trace`

#### PlasmidAnnotate

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/plasmid_annotate`

#### PlasmidDeepAnnotate

| Field | Description |
| --- | --- |
| `circular` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/plasmid_deep_annotate`

#### PlasmidFullReport

| Field | Description |
| --- | --- |
| `circular` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |
| `top_n` |  |

Operations: Create.

API path: `/plasmid_full_report`

#### PlasmidIdentify

| Field | Description |
| --- | --- |
| `circular` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |
| `top_n` |  |

Operations: Create.

API path: `/plasmid_identify`

#### PrimeEditingDesign

| Field | Description |
| --- | --- |
| `edit_end` |  |
| `edit_start` |  |
| `frame_start` |  |
| `gate` |  |
| `inserted_seq` |  |
| `ok` |  |
| `pbs_length` |  |
| `provenance` |  |
| `result` |  |
| `rtt_homology` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/prime_editing_design`

#### PrimeEditingTwinDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `new_sequence` |  |
| `ok` |  |
| `overlap_length` |  |
| `pbs_length` |  |
| `provenance` |  |
| `replace_end` |  |
| `replace_start` |  |
| `result` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/prime_editing_twin_design`

#### PrimerDesign

| Field | Description |
| --- | --- |
| `amplicon_max` |  |
| `amplicon_min` |  |
| `dntp_mm` |  |
| `gate` |  |
| `gc_max` |  |
| `gc_min` |  |
| `len_max` |  |
| `len_min` |  |
| `len_opt` |  |
| `max_return` |  |
| `mg_mm` |  |
| `na_mm` |  |
| `ok` |  |
| `oligo_nm` |  |
| `provenance` |  |
| `result` |  |
| `target_end` |  |
| `target_start` |  |
| `template` |  |
| `tm_max` |  |
| `tm_max_diff` |  |
| `tm_min` |  |
| `tm_opt` |  |
| `tool` |  |

Operations: Create.

API path: `/primer_design`

#### PrimerSpecificity

| Field | Description |
| --- | --- |
| `forward_primer` |  |
| `gate` |  |
| `max_mismatch` |  |
| `max_product_length` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reverse_primer` |  |
| `tool` |  |

Operations: Create.

API path: `/primer_specificity`

#### ProteaseDigestion

| Field | Description |
| --- | --- |
| `gate` |  |
| `max_mass` |  |
| `max_peptide` |  |
| `min_mass` |  |
| `missed_cleavage` |  |
| `ok` |  |
| `protease` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/protease_digestion`

#### ProteinAnnotatePoll

| Field | Description |
| --- | --- |
| `gate` |  |
| `job_id` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/protein_annotate_poll`

#### ProteinAnnotateSubmit

| Field | Description |
| --- | --- |
| `appl` |  |
| `gate` |  |
| `goterm` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/protein_annotate_submit`

#### ProteinHydrophobicity

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `scale` |  |
| `sequence` |  |
| `tool` |  |
| `window` |  |

Operations: Create.

API path: `/protein_hydrophobicity`

#### ProteinProperty

| Field | Description |
| --- | --- |
| `charge_step` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/protein_properties`

#### RandomSequence

| Field | Description |
| --- | --- |
| `gate` |  |
| `gc_content` |  |
| `kind` |  |
| `length` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/random_sequence`

#### RestrictionSite

| Field | Description |
| --- | --- |
| `enzyme` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/restriction_sites`

#### ReverseComplement

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |
| `type` |  |

Operations: Create.

API path: `/reverse_complement`

#### ReverseTranslate

| Field | Description |
| --- | --- |
| `gate` |  |
| `mode` |  |
| `ok` |  |
| `organism` |  |
| `protein` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/reverse_translate`

#### RnaFold

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/rna_fold`

#### SangerVsReference

| Field | Description |
| --- | --- |
| `file_base64` |  |
| `file_name` |  |
| `gate` |  |
| `min_coverage` |  |
| `ok` |  |
| `provenance` |  |
| `read` |  |
| `reference` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/sanger_vs_reference`

#### SavePermalink

| Field | Description |
| --- | --- |
| `arg` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/save_permalink`

#### SeqfileStat

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `ok` |  |
| `provenance` |  |
| `quality_offset` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/seqfile_stats`

#### SequenceFetch

| Field | Description |
| --- | --- |
| `accession` |  |
| `db` |  |
| `format` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/sequence_fetch`

#### SequenceFormatConvert

| Field | Description |
| --- | --- |
| `from` |  |
| `gate` |  |
| `input` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `to` |  |
| `tool` |  |

Operations: Create.

API path: `/sequence_format_convert`

#### SequenceReport

| Field | Description |
| --- | --- |
| `end_primer_length` |  |
| `gate` |  |
| `max_orf` |  |
| `min_orf_aa` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/sequence_report`

#### SequenceSearch

| Field | Description |
| --- | --- |
| `db` |  |
| `gate` |  |
| `gene` |  |
| `max_result` |  |
| `ok` |  |
| `organism` |  |
| `provenance` |  |
| `result` |  |
| `term` |  |
| `tool` |  |

Operations: Create.

API path: `/sequence_search`

#### SequencingReadbackVerify

| Field | Description |
| --- | --- |
| `gate` |  |
| `min_supporting_read` |  |
| `ok` |  |
| `provenance` |  |
| `read` |  |
| `reference` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/sequencing_readback_verify`

#### SessionCreate

| Field | Description |
| --- | --- |
| `entry` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/session_create`

#### SessionGet

| Field | Description |
| --- | --- |
| `gate` |  |
| `name` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `session_id` |  |
| `tool` |  |

Operations: Create.

API path: `/session_get`

#### SessionRun

| Field | Description |
| --- | --- |
| `arg` |  |
| `from_session` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `session_id` |  |
| `tool` |  |
| `write_back` |  |

Operations: Create.

API path: `/session_run`

#### SessionSet

| Field | Description |
| --- | --- |
| `entry` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `session_id` |  |
| `tool` |  |

Operations: Create.

API path: `/session_set`

#### SirnaDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `min_reynold` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sh_rna_loop` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/sirna_design`

#### SiteDirectedMutagenesi

| Field | Description |
| --- | --- |
| `arm_tm_target` |  |
| `dntp_mm` |  |
| `edit_kind` |  |
| `frame_start` |  |
| `gate` |  |
| `mg_mm` |  |
| `na_mm` |  |
| `new_base` |  |
| `ok` |  |
| `oligo_nm` |  |
| `organism` |  |
| `position` |  |
| `provenance` |  |
| `residue` |  |
| `result` |  |
| `style` |  |
| `target_aa` |  |
| `template` |  |
| `tool` |  |

Operations: Create.

API path: `/site_directed_mutagenesis`

#### Translate

| Field | Description |
| --- | --- |
| `frame` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `to_stop` |  |
| `tool` |  |

Operations: Create.

API path: `/translate`

#### VariantAnnotate

| Field | Description |
| --- | --- |
| `assembly` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `variant` |  |

Operations: Create.

API path: `/variant_annotate`

#### VariantComparator

| Field | Description |
| --- | --- |
| `coding` |  |
| `frame_start` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `query` |  |
| `reference` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/variant_comparator`

#### VerifyAssembly

| Field | Description |
| --- | --- |
| `arm_tm_target` |  |
| `circular` |  |
| `claimed_construct` |  |
| `coding` |  |
| `enzyme` |  |
| `enzyme3` |  |
| `enzyme5` |  |
| `fragment` |  |
| `fragment_pcr` |  |
| `frame_start` |  |
| `gate` |  |
| `insert` |  |
| `insert_pcr` |  |
| `method` |  |
| `name` |  |
| `ok` |  |
| `overlap_len` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `vector` |  |
| `vector_pcr` |  |

Operations: Create.

API path: `/verify_assembly`

#### VerifyConstruct

| Field | Description |
| --- | --- |
| `claimed_construct` |  |
| `expected_frame_start` |  |
| `gate` |  |
| `insert_forward_primer` |  |
| `insert_reverse_primer` |  |
| `insert_template` |  |
| `max_primer_mismatch` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `template_circular` |  |
| `tool` |  |

Operations: Create.

API path: `/verify_construct`

#### VirtualGel

| Field | Description |
| --- | --- |
| `circular` |  |
| `enzyme` |  |
| `gate` |  |
| `ladder` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/virtual_gel`

#### VolcanoPlotData

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `row` |  |
| `tool` |  |

Operations: Create.

API path: `/volcano_plot_data`

#### WebSearch

| Field | Description |
| --- | --- |
| `gate` |  |
| `max_result` |  |
| `ok` |  |
| `provenance` |  |
| `query` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/web_search`



## Entities


### AlphafoldLookup

Create an instance: `$alphafold_lookup = $client->AlphafoldLookup();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `string` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$alphafold_lookup = $client->AlphafoldLookup()->create([
    "accession" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### AsoDesign

Create an instance: `$aso_design = $client->AsoDesign();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `length` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `target` | `string` |  |
| `tool` | `string` |  |
| `wing` | `int` |  |

#### Example: Create

```php
$aso_design = $client->AsoDesign()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "target" => null, // string
    "tool" => null, // string
]);
```


### BaseEditingDesign

Create an instance: `$base_editing_design = $client->BaseEditingDesign();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `editor` | `string` |  |
| `frame_start` | `int` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `target` | `string` |  |
| `target_position` | `int` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$base_editing_design = $client->BaseEditingDesign()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "target" => null, // string
    "tool" => null, // string
]);
```


### Batch

Create an instance: `$batch = $client->Batch();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `array` |  |
| `input` | `string` |  |
| `ok` | `mixed` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Load

```php
// load() returns the bare Batch record (throws on error).
$batch = $client->Batch()->load();
```

#### Example: Create

```php
$batch = $client->Batch()->create([
    "input" => null, // string
    "ok" => null, // mixed
    "result" => null, // array
    "tool" => null, // string
]);
```


### BatchWorkflow

Create an instance: `$batch__workflow = $client->BatchWorkflow();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `input` | `string` |  |
| `ok` | `mixed` |  |
| `result` | `array` |  |
| `step` | `array` |  |

#### Example: Load

```php
// load() returns the bare BatchWorkflow record (throws on error).
$batch__workflow = $client->BatchWorkflow()->load();
```

#### Example: Create

```php
$batch__workflow = $client->BatchWorkflow()->create([
    "input" => null, // string
    "ok" => null, // mixed
    "result" => null, // array
    "step" => null, // array
]);
```


### CharacterizeSequence

Create an instance: `$characterize_sequence = $client->CharacterizeSequence();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `end_primer_length` | `int` |  |
| `gate` | `mixed` |  |
| `max_orf` | `int` |  |
| `min_orf_aa` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$characterize_sequence = $client->CharacterizeSequence()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### CloningSimulate

Create an instance: `$cloning_simulate = $client->CloningSimulate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_tm_target` | `float` |  |
| `circular` | `bool` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragment` | `array` |  |
| `gate` | `mixed` |  |
| `insert` | `string` |  |
| `method` | `string` |  |
| `name` | `array` |  |
| `ok` | `mixed` |  |
| `overlap_len` | `int` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |
| `vector` | `string` |  |

#### Example: Create

```php
$cloning_simulate = $client->CloningSimulate()->create([
    "method" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### CodonAdaptationIndex

Create an instance: `$codon_adaptation_index = $client->CodonAdaptationIndex();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `frame_start` | `int` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `organism` | `string` |  |
| `provenance` | `array` |  |
| `rare_threshold` | `float` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$codon_adaptation_index = $client->CodonAdaptationIndex()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### CodonOptimize

Create an instance: `$codon_optimize = $client->CodonOptimize();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `organism` | `string` |  |
| `protein` | `string` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$codon_optimize = $client->CodonOptimize()->create([
    "ok" => null, // mixed
    "protein" => null, // string
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### ConstructAutofix

Create an instance: `$construct_autofix = $client->ConstructAutofix();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoid_enzyme` | `array` |  |
| `cryptic_orf_min_aa` | `int` |  |
| `frame_start` | `int` |  |
| `gate` | `mixed` |  |
| `gc_high` | `float` |  |
| `gc_low` | `float` |  |
| `gc_window` | `int` |  |
| `homopolymer_min` | `int` |  |
| `max_pass` | `int` |  |
| `ok` | `mixed` |  |
| `organism` | `string` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$construct_autofix = $client->ConstructAutofix()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### ConstructQc

Create an instance: `$construct_qc = $client->ConstructQc();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoid_enzyme` | `array` |  |
| `cryptic_orf_min_aa` | `int` |  |
| `frame_start` | `int` |  |
| `gate` | `mixed` |  |
| `gc_high` | `float` |  |
| `gc_low` | `float` |  |
| `gc_window` | `int` |  |
| `homopolymer_min` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$construct_qc = $client->ConstructQc()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### CrisprGrnaDesign

Create an instance: `$crispr_grna_design = $client->CrisprGrnaDesign();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `min_score` | `float` |  |
| `nuclease` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `search_reverse_strand` | `bool` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$crispr_grna_design = $client->CrisprGrnaDesign()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### CrisprHdrDonor

Create an instance: `$crispr_hdr_donor = $client->CrisprHdrDonor();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_length` | `int` |  |
| `block_pam` | `bool` |  |
| `design_genotyping_primer` | `bool` |  |
| `edit_end` | `int` |  |
| `edit_start` | `int` |  |
| `frame_start` | `int` |  |
| `gate` | `mixed` |  |
| `guide_end` | `int` |  |
| `guide_start` | `int` |  |
| `guide_strand` | `string` |  |
| `nuclease` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `replacement` | `string` |  |
| `result` | `array` |  |
| `target_sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$crispr_hdr_donor = $client->CrisprHdrDonor()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "replacement" => null, // string
    "result" => null, // array
    "target_sequence" => null, // string
    "tool" => null, // string
]);
```


### CrisprOfftargetCheck

Create an instance: `$crispr_offtarget_check = $client->CrisprOfftargetCheck();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `max_mismatch` | `int` |  |
| `nuclease` | `string` |  |
| `ok` | `mixed` |  |
| `protospacer` | `string` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$crispr_offtarget_check = $client->CrisprOfftargetCheck()->create([
    "ok" => null, // mixed
    "protospacer" => null, // string
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### CrossDimer

Create an instance: `$cross_dimer = $client->CrossDimer();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence_a` | `string` |  |
| `sequence_b` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$cross_dimer = $client->CrossDimer()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence_a" => null, // string
    "sequence_b" => null, // string
    "tool" => null, // string
]);
```


### DnaMolarity

Create an instance: `$dna_molarity = $client->DnaMolarity();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `length` | `int` |  |
| `mass_ng` | `float` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |
| `volume_ul` | `float` |  |

#### Example: Create

```php
$dna_molarity = $client->DnaMolarity()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### DoubleDigest

Create an instance: `$double_digest = $client->DoubleDigest();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enzyme_a` | `string` |  |
| `enzyme_b` | `string` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$double_digest = $client->DoubleDigest()->create([
    "enzyme_a" => null, // string
    "enzyme_b" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### ExportEchoPicklist

Create an instance: `$export_echo_picklist = $client->ExportEchoPicklist();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `reaction` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$export_echo_picklist = $client->ExportEchoPicklist()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "reaction" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### ExportOpentronsProtocol

Create an instance: `$export_opentrons_protocol = $client->ExportOpentronsProtocol();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `protocol_name` | `string` |  |
| `provenance` | `array` |  |
| `reaction` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$export_opentrons_protocol = $client->ExportOpentronsProtocol()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "reaction" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### ExportPlateLayout

Create an instance: `$export_plate_layout = $client->ExportPlateLayout();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `reaction` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$export_plate_layout = $client->ExportPlateLayout()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "reaction" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### ExpressionHeatmapCluster

Create an instance: `$expression_heatmap_cluster = $client->ExpressionHeatmapCluster();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cluster_col` | `bool` |  |
| `cluster_row` | `bool` |  |
| `distance_metric` | `string` |  |
| `gate` | `mixed` |  |
| `gene` | `array` |  |
| `linkage` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sample` | `array` |  |
| `tool` | `string` |  |
| `value` | `array` |  |
| `z_score_row` | `bool` |  |

#### Example: Create

```php
$expression_heatmap_cluster = $client->ExpressionHeatmapCluster()->create([
    "gene" => null, // array
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sample" => null, // array
    "tool" => null, // string
    "value" => null, // array
]);
```


### FastqQcReport

Create an instance: `$fastq_qc_report = $client->FastqQcReport();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `input` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `quality_offset` | `int` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$fastq_qc_report = $client->FastqQcReport()->create([
    "input" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### FastqTrim

Create an instance: `$fastq_trim = $client->FastqTrim();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `input` | `string` |  |
| `min_length` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `quality_offset` | `int` |  |
| `quality_threshold` | `int` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$fastq_trim = $client->FastqTrim()->create([
    "input" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### FindOrf

Create an instance: `$find_orf = $client->FindOrf();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `min_aa_length` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `require_stop` | `bool` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$find_orf = $client->FindOrf()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### FormatSequence

Create an instance: `$format_sequence = $client->FormatSequence();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `case_mode` | `string` |  |
| `convert` | `string` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `reverse` | `bool` |  |
| `sequence` | `string` |  |
| `strip_non_letter` | `bool` |  |
| `tool` | `string` |  |
| `width` | `int` |  |

#### Example: Create

```php
$format_sequence = $client->FormatSequence()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### FunctionalEnrichment

Create an instance: `$functional_enrichment = $client->FunctionalEnrichment();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `background` | `array` |  |
| `collection` | `array` |  |
| `gate` | `mixed` |  |
| `gene` | `array` |  |
| `max_term_size` | `int` |  |
| `min_term_size` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$functional_enrichment = $client->FunctionalEnrichment()->create([
    "gene" => null, // array
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### GcContent

Create an instance: `$gc_content = $client->GcContent();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$gc_content = $client->GcContent()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### GeneDossier

Create an instance: `$gene_dossier = $client->GeneDossier();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `gene` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$gene_dossier = $client->GeneDossier()->create([
    "gene" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### GeneExpression

Create an instance: `$gene_expression = $client->GeneExpression();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `gene` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$gene_expression = $client->GeneExpression()->create([
    "gene" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### GeneModel

Create an instance: `$gene_model = $client->GeneModel();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `gene` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$gene_model = $client->GeneModel()->create([
    "gene" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### GoldenGateFidelity

Create an instance: `$golden_gate_fidelity = $client->GoldenGateFidelity();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `compare_to_named_set` | `string` |  |
| `dataset` | `string` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `overhang` | `array` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `risk_threshold` | `float` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$golden_gate_fidelity = $client->GoldenGateFidelity()->create([
    "ok" => null, // mixed
    "overhang" => null, // array
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### HgvsConvert

Create an instance: `$hgvs_convert = $client->HgvsConvert();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |
| `variant` | `string` |  |

#### Example: Create

```php
$hgvs_convert = $client->HgvsConvert()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
    "variant" => null, // string
]);
```


### IdMapPoll

Create an instance: `$id_map_poll = $client->IdMapPoll();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `job_id` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$id_map_poll = $client->IdMapPoll()->create([
    "job_id" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### IdMapSubmit

Create an instance: `$id_map_submit = $client->IdMapSubmit();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from` | `string` |  |
| `gate` | `mixed` |  |
| `ids` | `array` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tax_id` | `string` |  |
| `to` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$id_map_submit = $client->IdMapSubmit()->create([
    "from" => null, // string
    "ids" => null, // array
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "to" => null, // string
    "tool" => null, // string
]);
```


### InSilicoPcr

Create an instance: `$in_silico_pcr = $client->InSilicoPcr();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `forward_primer` | `string` |  |
| `gate` | `mixed` |  |
| `max_mismatch` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `reverse_primer` | `string` |  |
| `template` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$in_silico_pcr = $client->InSilicoPcr()->create([
    "forward_primer" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "reverse_primer" => null, // string
    "template" => null, // string
    "tool" => null, // string
]);
```


### KaspPrimerDesign

Create an instance: `$kasp_primer_design = $client->KaspPrimerDesign();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `add_secondary_mismatch` | `bool` |  |
| `allele_a` | `string` |  |
| `allele_b` | `string` |  |
| `gate` | `mixed` |  |
| `max_amplicon` | `int` |  |
| `min_amplicon` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `snp_position` | `int` |  |
| `target` | `string` |  |
| `target_core_tm` | `float` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$kasp_primer_design = $client->KaspPrimerDesign()->create([
    "allele_a" => null, // string
    "allele_b" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "snp_position" => null, // int
    "target" => null, // string
    "tool" => null, // string
]);
```


### ListTool

Create an instance: `$list_tool = $client->ListTool();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```php
// load() returns the bare ListTool record (throws on error).
$list_tool = $client->ListTool()->load();
```


### MeltingTemperature

Create an instance: `$melting_temperature = $client->MeltingTemperature();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dntp_mm` | `float` |  |
| `gate` | `mixed` |  |
| `mg_mm` | `float` |  |
| `na_mm` | `float` |  |
| `ok` | `mixed` |  |
| `oligo_nm` | `float` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `target_tm` | `float` |  |
| `tm_tolerance` | `float` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$melting_temperature = $client->MeltingTemperature()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### MotifFinder

Create an instance: `$motif_finder = $client->MotifFinder();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `max_mismatch` | `int` |  |
| `motif` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `search_reverse_strand` | `bool` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$motif_finder = $client->MotifFinder()->create([
    "motif" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### MultipleSequenceAlignment

Create an instance: `$multiple_sequence_alignment = $client->MultipleSequenceAlignment();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `input` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$multiple_sequence_alignment = $client->MultipleSequenceAlignment()->create([
    "input" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### OligoAnalysi

Create an instance: `$oligo_analysi = $client->OligoAnalysi();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dntp_mm` | `float` |  |
| `gate` | `mixed` |  |
| `mg_mm` | `float` |  |
| `na_mm` | `float` |  |
| `ok` | `mixed` |  |
| `oligo_nm` | `float` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$oligo_analysi = $client->OligoAnalysi()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### OrthologMap

Create an instance: `$ortholog_map = $client->OrthologMap();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `source_species` | `string` |  |
| `symbol` | `array` |  |
| `target_species` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |

#### Example: Create

```php
$ortholog_map = $client->OrthologMap()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "symbol" => null, // array
    "target_species" => null, // string
    "tool" => null, // string
]);
```


### PairwiseAlignment

Create an instance: `$pairwise_alignment = $client->PairwiseAlignment();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gap` | `float` |  |
| `gate` | `mixed` |  |
| `match` | `float` |  |
| `mismatch` | `float` |  |
| `mode` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `seq_a` | `string` |  |
| `seq_b` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$pairwise_alignment = $client->PairwiseAlignment()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "seq_a" => null, // string
    "seq_b" => null, // string
    "tool" => null, // string
]);
```


### ParseGenbank

Create an instance: `$parse_genbank = $client->ParseGenbank();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `text` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$parse_genbank = $client->ParseGenbank()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "text" => null, // string
    "tool" => null, // string
]);
```


### ParseSangerTrace

Create an instance: `$parse_sanger_trace = $client->ParseSangerTrace();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file_base64` | `string` |  |
| `file_name` | `string` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$parse_sanger_trace = $client->ParseSangerTrace()->create([
    "file_base64" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### PlasmidAnnotate

Create an instance: `$plasmid_annotate = $client->PlasmidAnnotate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$plasmid_annotate = $client->PlasmidAnnotate()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### PlasmidDeepAnnotate

Create an instance: `$plasmid_deep_annotate = $client->PlasmidDeepAnnotate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$plasmid_deep_annotate = $client->PlasmidDeepAnnotate()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### PlasmidFullReport

Create an instance: `$plasmid_full_report = $client->PlasmidFullReport();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `top_n` | `int` |  |

#### Example: Create

```php
$plasmid_full_report = $client->PlasmidFullReport()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### PlasmidIdentify

Create an instance: `$plasmid_identify = $client->PlasmidIdentify();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `top_n` | `int` |  |

#### Example: Create

```php
$plasmid_identify = $client->PlasmidIdentify()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### PrimeEditingDesign

Create an instance: `$prime_editing_design = $client->PrimeEditingDesign();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `edit_end` | `int` |  |
| `edit_start` | `int` |  |
| `frame_start` | `int` |  |
| `gate` | `mixed` |  |
| `inserted_seq` | `string` |  |
| `ok` | `mixed` |  |
| `pbs_length` | `int` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `rtt_homology` | `int` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$prime_editing_design = $client->PrimeEditingDesign()->create([
    "edit_end" => null, // int
    "edit_start" => null, // int
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "target" => null, // string
    "tool" => null, // string
]);
```


### PrimeEditingTwinDesign

Create an instance: `$prime_editing_twin_design = $client->PrimeEditingTwinDesign();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `new_sequence` | `string` |  |
| `ok` | `mixed` |  |
| `overlap_length` | `int` |  |
| `pbs_length` | `int` |  |
| `provenance` | `array` |  |
| `replace_end` | `int` |  |
| `replace_start` | `int` |  |
| `result` | `array` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$prime_editing_twin_design = $client->PrimeEditingTwinDesign()->create([
    "new_sequence" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "replace_end" => null, // int
    "replace_start" => null, // int
    "result" => null, // array
    "target" => null, // string
    "tool" => null, // string
]);
```


### PrimerDesign

Create an instance: `$primer_design = $client->PrimerDesign();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amplicon_max` | `int` |  |
| `amplicon_min` | `int` |  |
| `dntp_mm` | `float` |  |
| `gate` | `mixed` |  |
| `gc_max` | `float` |  |
| `gc_min` | `float` |  |
| `len_max` | `int` |  |
| `len_min` | `int` |  |
| `len_opt` | `int` |  |
| `max_return` | `int` |  |
| `mg_mm` | `float` |  |
| `na_mm` | `float` |  |
| `ok` | `mixed` |  |
| `oligo_nm` | `float` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `target_end` | `int` |  |
| `target_start` | `int` |  |
| `template` | `string` |  |
| `tm_max` | `float` |  |
| `tm_max_diff` | `float` |  |
| `tm_min` | `float` |  |
| `tm_opt` | `float` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$primer_design = $client->PrimerDesign()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "template" => null, // string
    "tool" => null, // string
]);
```


### PrimerSpecificity

Create an instance: `$primer_specificity = $client->PrimerSpecificity();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `forward_primer` | `string` |  |
| `gate` | `mixed` |  |
| `max_mismatch` | `int` |  |
| `max_product_length` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `reverse_primer` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$primer_specificity = $client->PrimerSpecificity()->create([
    "forward_primer" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "reverse_primer" => null, // string
    "tool" => null, // string
]);
```


### ProteaseDigestion

Create an instance: `$protease_digestion = $client->ProteaseDigestion();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `max_mass` | `float` |  |
| `max_peptide` | `int` |  |
| `min_mass` | `float` |  |
| `missed_cleavage` | `int` |  |
| `ok` | `mixed` |  |
| `protease` | `string` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$protease_digestion = $client->ProteaseDigestion()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### ProteinAnnotatePoll

Create an instance: `$protein_annotate_poll = $client->ProteinAnnotatePoll();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `job_id` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$protein_annotate_poll = $client->ProteinAnnotatePoll()->create([
    "job_id" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### ProteinAnnotateSubmit

Create an instance: `$protein_annotate_submit = $client->ProteinAnnotateSubmit();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `appl` | `string` |  |
| `gate` | `mixed` |  |
| `goterm` | `bool` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$protein_annotate_submit = $client->ProteinAnnotateSubmit()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### ProteinHydrophobicity

Create an instance: `$protein_hydrophobicity = $client->ProteinHydrophobicity();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `scale` | `string` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `window` | `int` |  |

#### Example: Create

```php
$protein_hydrophobicity = $client->ProteinHydrophobicity()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### ProteinProperty

Create an instance: `$protein_property = $client->ProteinProperty();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `charge_step` | `float` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$protein_property = $client->ProteinProperty()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### RandomSequence

Create an instance: `$random_sequence = $client->RandomSequence();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `gc_content` | `float` |  |
| `kind` | `string` |  |
| `length` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$random_sequence = $client->RandomSequence()->create([
    "length" => null, // int
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### RestrictionSite

Create an instance: `$restriction_site = $client->RestrictionSite();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enzyme` | `array` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$restriction_site = $client->RestrictionSite()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### ReverseComplement

Create an instance: `$reverse_complement = $client->ReverseComplement();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |

#### Example: Create

```php
$reverse_complement = $client->ReverseComplement()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### ReverseTranslate

Create an instance: `$reverse_translate = $client->ReverseTranslate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `mode` | `string` |  |
| `ok` | `mixed` |  |
| `organism` | `string` |  |
| `protein` | `string` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$reverse_translate = $client->ReverseTranslate()->create([
    "ok" => null, // mixed
    "protein" => null, // string
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### RnaFold

Create an instance: `$rna_fold = $client->RnaFold();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$rna_fold = $client->RnaFold()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### SangerVsReference

Create an instance: `$sanger_vs_reference = $client->SangerVsReference();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file_base64` | `string` |  |
| `file_name` | `string` |  |
| `gate` | `mixed` |  |
| `min_coverage` | `float` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `read` | `string` |  |
| `reference` | `string` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$sanger_vs_reference = $client->SangerVsReference()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "reference" => null, // string
    "result" => null, // array
    "tool" => null, // string
]);
```


### SavePermalink

Create an instance: `$save_permalink = $client->SavePermalink();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `array` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$save_permalink = $client->SavePermalink()->create([
    "arg" => null, // array
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### SeqfileStat

Create an instance: `$seqfile_stat = $client->SeqfileStat();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `input` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `quality_offset` | `int` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$seqfile_stat = $client->SeqfileStat()->create([
    "input" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### SequenceFetch

Create an instance: `$sequence_fetch = $client->SequenceFetch();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `string` |  |
| `db` | `string` |  |
| `format` | `string` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$sequence_fetch = $client->SequenceFetch()->create([
    "accession" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### SequenceFormatConvert

Create an instance: `$sequence_format_convert = $client->SequenceFormatConvert();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from` | `string` |  |
| `gate` | `mixed` |  |
| `input` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `to` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$sequence_format_convert = $client->SequenceFormatConvert()->create([
    "input" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### SequenceReport

Create an instance: `$sequence_report = $client->SequenceReport();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `end_primer_length` | `int` |  |
| `gate` | `mixed` |  |
| `max_orf` | `int` |  |
| `min_orf_aa` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$sequence_report = $client->SequenceReport()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### SequenceSearch

Create an instance: `$sequence_search = $client->SequenceSearch();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `db` | `string` |  |
| `gate` | `mixed` |  |
| `gene` | `string` |  |
| `max_result` | `int` |  |
| `ok` | `mixed` |  |
| `organism` | `string` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `term` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$sequence_search = $client->SequenceSearch()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### SequencingReadbackVerify

Create an instance: `$sequencing_readback_verify = $client->SequencingReadbackVerify();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `min_supporting_read` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `read` | `string` |  |
| `reference` | `string` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$sequencing_readback_verify = $client->SequencingReadbackVerify()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "read" => null, // string
    "reference" => null, // string
    "result" => null, // array
    "tool" => null, // string
]);
```


### SessionCreate

Create an instance: `$session_create = $client->SessionCreate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entry` | `array` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$session_create = $client->SessionCreate()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### SessionGet

Create an instance: `$session_get = $client->SessionGet();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `name` | `array` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$session_get = $client->SessionGet()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "session_id" => null, // string
    "tool" => null, // string
]);
```


### SessionRun

Create an instance: `$session_run = $client->SessionRun();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `array` |  |
| `from_session` | `array` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |
| `write_back` | `array` |  |

#### Example: Create

```php
$session_run = $client->SessionRun()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "session_id" => null, // string
    "tool" => null, // string
]);
```


### SessionSet

Create an instance: `$session_set = $client->SessionSet();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entry` | `array` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$session_set = $client->SessionSet()->create([
    "entry" => null, // array
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "session_id" => null, // string
    "tool" => null, // string
]);
```


### SirnaDesign

Create an instance: `$sirna_design = $client->SirnaDesign();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `min_reynold` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sh_rna_loop` | `string` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$sirna_design = $client->SirnaDesign()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "target" => null, // string
    "tool" => null, // string
]);
```


### SiteDirectedMutagenesi

Create an instance: `$site_directed_mutagenesi = $client->SiteDirectedMutagenesi();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_tm_target` | `float` |  |
| `dntp_mm` | `float` |  |
| `edit_kind` | `string` |  |
| `frame_start` | `int` |  |
| `gate` | `mixed` |  |
| `mg_mm` | `float` |  |
| `na_mm` | `float` |  |
| `new_base` | `string` |  |
| `ok` | `mixed` |  |
| `oligo_nm` | `float` |  |
| `organism` | `string` |  |
| `position` | `int` |  |
| `provenance` | `array` |  |
| `residue` | `int` |  |
| `result` | `array` |  |
| `style` | `string` |  |
| `target_aa` | `string` |  |
| `template` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$site_directed_mutagenesi = $client->SiteDirectedMutagenesi()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "template" => null, // string
    "tool" => null, // string
]);
```


### Translate

Create an instance: `$translate = $client->Translate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `frame` | `int` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `to_stop` | `bool` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$translate = $client->Translate()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### VariantAnnotate

Create an instance: `$variant_annotate = $client->VariantAnnotate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `assembly` | `string` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |
| `variant` | `string` |  |

#### Example: Create

```php
$variant_annotate = $client->VariantAnnotate()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
    "variant" => null, // string
]);
```


### VariantComparator

Create an instance: `$variant_comparator = $client->VariantComparator();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `coding` | `bool` |  |
| `frame_start` | `int` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `query` | `string` |  |
| `reference` | `string` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$variant_comparator = $client->VariantComparator()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "query" => null, // string
    "reference" => null, // string
    "result" => null, // array
    "tool" => null, // string
]);
```


### VerifyAssembly

Create an instance: `$verify_assembly = $client->VerifyAssembly();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_tm_target` | `float` |  |
| `circular` | `bool` |  |
| `claimed_construct` | `string` |  |
| `coding` | `bool` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragment` | `array` |  |
| `fragment_pcr` | `array` |  |
| `frame_start` | `int` |  |
| `gate` | `mixed` |  |
| `insert` | `string` |  |
| `insert_pcr` | `array` |  |
| `method` | `string` |  |
| `name` | `array` |  |
| `ok` | `mixed` |  |
| `overlap_len` | `int` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |
| `vector` | `string` |  |
| `vector_pcr` | `array` |  |

#### Example: Create

```php
$verify_assembly = $client->VerifyAssembly()->create([
    "claimed_construct" => null, // string
    "method" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### VerifyConstruct

Create an instance: `$verify_construct = $client->VerifyConstruct();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `claimed_construct` | `string` |  |
| `expected_frame_start` | `int` |  |
| `gate` | `mixed` |  |
| `insert_forward_primer` | `string` |  |
| `insert_reverse_primer` | `string` |  |
| `insert_template` | `string` |  |
| `max_primer_mismatch` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `template_circular` | `bool` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$verify_construct = $client->VerifyConstruct()->create([
    "claimed_construct" => null, // string
    "insert_forward_primer" => null, // string
    "insert_reverse_primer" => null, // string
    "insert_template" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "tool" => null, // string
]);
```


### VirtualGel

Create an instance: `$virtual_gel = $client->VirtualGel();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `enzyme` | `array` |  |
| `gate` | `mixed` |  |
| `ladder` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$virtual_gel = $client->VirtualGel()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequence" => null, // string
    "tool" => null, // string
]);
```


### VolcanoPlotData

Create an instance: `$volcano_plot_data = $client->VolcanoPlotData();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `row` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$volcano_plot_data = $client->VolcanoPlotData()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "row" => null, // array
    "tool" => null, // string
]);
```


### WebSearch

Create an instance: `$web_search = $client->WebSearch();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `mixed` |  |
| `max_result` | `float` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `query` | `string` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$web_search = $client->WebSearch()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "query" => null, // string
    "result" => null, // array
    "tool" => null, // string
]);
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── seqbenchmcp_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`seqbenchmcp_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```php
$batch = $client->Batch();
$batch->load();

// $batch->data_get() now returns the batch data from the last load
// $batch->match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
