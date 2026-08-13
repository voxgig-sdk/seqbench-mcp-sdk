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
// create() returns the ENTITY — call data_get() for the created AlphafoldLookup record.
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

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
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

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
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
| `frameStart` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `target` |  |
| `targetPosition` |  |
| `tool` |  |

Operations: Create.

API path: `/base_editing_design`

#### Batch

| Field | Description |
| --- | --- |
| `args` |  |
| `capped` |  |
| `columns` |  |
| `count` |  |
| `errors` |  |
| `input` |  |
| `limit` |  |
| `provenance` |  |
| `rows` |  |
| `tool` |  |

Operations: Create, Load.

API path: `/batch`

#### BatchWorkflow

| Field | Description |
| --- | --- |
| `capped` |  |
| `columns` |  |
| `count` |  |
| `errors` |  |
| `input` |  |
| `limit` |  |
| `provenance` |  |
| `rows` |  |
| `steps` |  |

Operations: Create, Load.

API path: `/workflow`

#### CharacterizeSequence

| Field | Description |
| --- | --- |
| `endPrimerLength` |  |
| `gate` |  |
| `maxOrfs` |  |
| `minOrfAa` |  |
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
| `armTmTarget` |  |
| `circular` |  |
| `enzyme` |  |
| `enzyme3` |  |
| `enzyme5` |  |
| `fragments` |  |
| `gate` |  |
| `insert` |  |
| `method` |  |
| `names` |  |
| `ok` |  |
| `overlapLen` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `vector` |  |

Operations: Create.

API path: `/cloning_simulate`

#### CodonAdaptationIndex

| Field | Description |
| --- | --- |
| `frameStart` |  |
| `gate` |  |
| `ok` |  |
| `organism` |  |
| `provenance` |  |
| `rareThreshold` |  |
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
| `avoidEnzymes` |  |
| `crypticOrfMinAa` |  |
| `frameStart` |  |
| `gate` |  |
| `gcHigh` |  |
| `gcLow` |  |
| `gcWindow` |  |
| `homopolymerMin` |  |
| `maxPasses` |  |
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
| `avoidEnzymes` |  |
| `crypticOrfMinAa` |  |
| `frameStart` |  |
| `gate` |  |
| `gcHigh` |  |
| `gcLow` |  |
| `gcWindow` |  |
| `homopolymerMin` |  |
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
| `minScore` |  |
| `nuclease` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `searchReverseStrand` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/crispr_grna_design`

#### CrisprHdrDonor

| Field | Description |
| --- | --- |
| `armLength` |  |
| `blockPam` |  |
| `designGenotypingPrimers` |  |
| `editEnd` |  |
| `editStart` |  |
| `frameStart` |  |
| `gate` |  |
| `guideEnd` |  |
| `guideStart` |  |
| `guideStrand` |  |
| `nuclease` |  |
| `ok` |  |
| `provenance` |  |
| `replacement` |  |
| `result` |  |
| `targetSequence` |  |
| `tool` |  |

Operations: Create.

API path: `/crispr_hdr_donor`

#### CrisprOfftargetCheck

| Field | Description |
| --- | --- |
| `gate` |  |
| `maxMismatches` |  |
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
| `sequenceA` |  |
| `sequenceB` |  |
| `tool` |  |

Operations: Create.

API path: `/cross_dimer`

#### DnaMolarity

| Field | Description |
| --- | --- |
| `gate` |  |
| `length` |  |
| `massNg` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |
| `type` |  |
| `volumeUl` |  |

Operations: Create.

API path: `/dna_molarity`

#### DoubleDigest

| Field | Description |
| --- | --- |
| `enzymeA` |  |
| `enzymeB` |  |
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
| `reactions` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/export_echo_picklist`

#### ExportOpentronsProtocol

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `protocolName` |  |
| `provenance` |  |
| `reactions` |  |
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
| `reactions` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/export_plate_layout`

#### ExpressionHeatmapCluster

| Field | Description |
| --- | --- |
| `clusterCols` |  |
| `clusterRows` |  |
| `distanceMetric` |  |
| `gate` |  |
| `genes` |  |
| `linkage` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `samples` |  |
| `tool` |  |
| `values` |  |
| `zScoreRows` |  |

Operations: Create.

API path: `/expression_heatmap_cluster`

#### FastqQcReport

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `ok` |  |
| `provenance` |  |
| `qualityOffset` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/fastq_qc_report`

#### FastqTrim

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `minLength` |  |
| `ok` |  |
| `provenance` |  |
| `qualityOffset` |  |
| `qualityThreshold` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/fastq_trim`

#### FindOrf

| Field | Description |
| --- | --- |
| `gate` |  |
| `minAaLength` |  |
| `ok` |  |
| `provenance` |  |
| `requireStop` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/find_orfs`

#### FormatSequence

| Field | Description |
| --- | --- |
| `caseMode` |  |
| `convert` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reverse` |  |
| `sequence` |  |
| `stripNonLetters` |  |
| `tool` |  |
| `width` |  |

Operations: Create.

API path: `/format_sequence`

#### FunctionalEnrichment

| Field | Description |
| --- | --- |
| `background` |  |
| `collections` |  |
| `gate` |  |
| `genes` |  |
| `maxTermSize` |  |
| `minTermSize` |  |
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
| `compareToNamedSet` |  |
| `dataset` |  |
| `gate` |  |
| `ok` |  |
| `overhangs` |  |
| `provenance` |  |
| `result` |  |
| `riskThreshold` |  |
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
| `jobId` |  |
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
| `taxId` |  |
| `to` |  |
| `tool` |  |

Operations: Create.

API path: `/id_map_submit`

#### InSilicoPcr

| Field | Description |
| --- | --- |
| `circular` |  |
| `forwardPrimer` |  |
| `gate` |  |
| `maxMismatches` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reversePrimer` |  |
| `template` |  |
| `tool` |  |

Operations: Create.

API path: `/in_silico_pcr`

#### KaspPrimerDesign

| Field | Description |
| --- | --- |
| `addSecondaryMismatch` |  |
| `alleleA` |  |
| `alleleB` |  |
| `gate` |  |
| `maxAmplicon` |  |
| `minAmplicon` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `snpPosition` |  |
| `target` |  |
| `targetCoreTm` |  |
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
| `dntpMM` |  |
| `gate` |  |
| `mgMM` |  |
| `naMM` |  |
| `ok` |  |
| `oligoNM` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `targetTm` |  |
| `tmTolerance` |  |
| `tool` |  |

Operations: Create.

API path: `/melting_temperature`

#### MotifFinder

| Field | Description |
| --- | --- |
| `gate` |  |
| `maxMismatches` |  |
| `motif` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `searchReverseStrand` |  |
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
| `dntpMM` |  |
| `gate` |  |
| `mgMM` |  |
| `naMM` |  |
| `ok` |  |
| `oligoNM` |  |
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
| `sourceSpecies` |  |
| `symbols` |  |
| `targetSpecies` |  |
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
| `seqA` |  |
| `seqB` |  |
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
| `fileBase64` |  |
| `fileName` |  |
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
| `topN` |  |

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
| `topN` |  |

Operations: Create.

API path: `/plasmid_identify`

#### PrimeEditingDesign

| Field | Description |
| --- | --- |
| `editEnd` |  |
| `editStart` |  |
| `frameStart` |  |
| `gate` |  |
| `insertedSeq` |  |
| `ok` |  |
| `pbsLength` |  |
| `provenance` |  |
| `result` |  |
| `rttHomology` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/prime_editing_design`

#### PrimeEditingTwinDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `newSequence` |  |
| `ok` |  |
| `overlapLength` |  |
| `pbsLength` |  |
| `provenance` |  |
| `replaceEnd` |  |
| `replaceStart` |  |
| `result` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/prime_editing_twin_design`

#### PrimerDesign

| Field | Description |
| --- | --- |
| `ampliconMax` |  |
| `ampliconMin` |  |
| `dntpMM` |  |
| `gate` |  |
| `gcMax` |  |
| `gcMin` |  |
| `lenMax` |  |
| `lenMin` |  |
| `lenOpt` |  |
| `maxReturn` |  |
| `mgMM` |  |
| `naMM` |  |
| `ok` |  |
| `oligoNM` |  |
| `provenance` |  |
| `result` |  |
| `targetEnd` |  |
| `targetStart` |  |
| `template` |  |
| `tmMax` |  |
| `tmMaxDiff` |  |
| `tmMin` |  |
| `tmOpt` |  |
| `tool` |  |

Operations: Create.

API path: `/primer_design`

#### PrimerSpecificity

| Field | Description |
| --- | --- |
| `forwardPrimer` |  |
| `gate` |  |
| `maxMismatches` |  |
| `maxProductLength` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reversePrimer` |  |
| `tool` |  |

Operations: Create.

API path: `/primer_specificity`

#### ProteaseDigestion

| Field | Description |
| --- | --- |
| `gate` |  |
| `maxMass` |  |
| `maxPeptides` |  |
| `minMass` |  |
| `missedCleavages` |  |
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
| `jobId` |  |
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
| `goterms` |  |
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
| `chargeStep` |  |
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
| `gcContent` |  |
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
| `enzymes` |  |
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
| `fileBase64` |  |
| `fileName` |  |
| `gate` |  |
| `minCoverage` |  |
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
| `args` |  |
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
| `qualityOffset` |  |
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
| `endPrimerLength` |  |
| `gate` |  |
| `maxOrfs` |  |
| `minOrfAa` |  |
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
| `maxResults` |  |
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
| `minSupportingReads` |  |
| `ok` |  |
| `provenance` |  |
| `reads` |  |
| `reference` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/sequencing_readback_verify`

#### SessionCreate

| Field | Description |
| --- | --- |
| `entries` |  |
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
| `names` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sessionId` |  |
| `tool` |  |

Operations: Create.

API path: `/session_get`

#### SessionRun

| Field | Description |
| --- | --- |
| `args` |  |
| `fromSession` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sessionId` |  |
| `tool` |  |
| `writeBack` |  |

Operations: Create.

API path: `/session_run`

#### SessionSet

| Field | Description |
| --- | --- |
| `entries` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sessionId` |  |
| `tool` |  |

Operations: Create.

API path: `/session_set`

#### SirnaDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `minReynolds` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `shRnaLoop` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/sirna_design`

#### SiteDirectedMutagenesi

| Field | Description |
| --- | --- |
| `armTmTarget` |  |
| `dntpMM` |  |
| `editKind` |  |
| `frameStart` |  |
| `gate` |  |
| `mgMM` |  |
| `naMM` |  |
| `newBase` |  |
| `ok` |  |
| `oligoNM` |  |
| `organism` |  |
| `position` |  |
| `provenance` |  |
| `residue` |  |
| `result` |  |
| `style` |  |
| `targetAa` |  |
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
| `toStop` |  |
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
| `frameStart` |  |
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
| `armTmTarget` |  |
| `circular` |  |
| `claimedConstruct` |  |
| `coding` |  |
| `enzyme` |  |
| `enzyme3` |  |
| `enzyme5` |  |
| `fragmentPcrs` |  |
| `fragments` |  |
| `frameStart` |  |
| `gate` |  |
| `insert` |  |
| `insertPcr` |  |
| `method` |  |
| `names` |  |
| `ok` |  |
| `overlapLen` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `vector` |  |
| `vectorPcr` |  |

Operations: Create.

API path: `/verify_assembly`

#### VerifyConstruct

| Field | Description |
| --- | --- |
| `claimedConstruct` |  |
| `expectedFrameStart` |  |
| `gate` |  |
| `insertForwardPrimer` |  |
| `insertReversePrimer` |  |
| `insertTemplate` |  |
| `maxPrimerMismatches` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `templateCircular` |  |
| `tool` |  |

Operations: Create.

API path: `/verify_construct`

#### VirtualGel

| Field | Description |
| --- | --- |
| `circular` |  |
| `enzymes` |  |
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
| `rows` |  |
| `tool` |  |

Operations: Create.

API path: `/volcano_plot_data`

#### WebSearch

| Field | Description |
| --- | --- |
| `gate` |  |
| `max_results` |  |
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
| `frameStart` | `int` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `target` | `string` |  |
| `targetPosition` | `int` |  |
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
| `args` | `array` |  |
| `capped` | `bool` |  |
| `columns` | `array` |  |
| `count` | `int` |  |
| `errors` | `int` |  |
| `input` | `string` |  |
| `limit` | `int` |  |
| `provenance` | `array` |  |
| `rows` | `array` |  |
| `tool` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Batch record (throws on error).
$batch = $client->Batch()->load();
```

#### Example: Create

```php
$batch = $client->Batch()->create([
    "capped" => null, // bool
    "columns" => null, // array
    "count" => null, // int
    "errors" => null, // int
    "input" => null, // string
    "limit" => null, // int
    "provenance" => null, // array
    "rows" => null, // array
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
| `capped` | `bool` |  |
| `columns` | `array` |  |
| `count` | `int` |  |
| `errors` | `int` |  |
| `input` | `string` |  |
| `limit` | `int` |  |
| `provenance` | `array` |  |
| `rows` | `array` |  |
| `steps` | `array` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the BatchWorkflow record (throws on error).
$batch__workflow = $client->BatchWorkflow()->load();
```

#### Example: Create

```php
$batch__workflow = $client->BatchWorkflow()->create([
    "capped" => null, // bool
    "columns" => null, // array
    "count" => null, // int
    "errors" => null, // int
    "input" => null, // string
    "limit" => null, // int
    "provenance" => null, // array
    "rows" => null, // array
    "steps" => null, // array
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
| `endPrimerLength` | `int` |  |
| `gate` | `mixed` |  |
| `maxOrfs` | `int` |  |
| `minOrfAa` | `int` |  |
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
| `armTmTarget` | `float` |  |
| `circular` | `bool` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragments` | `array` |  |
| `gate` | `mixed` |  |
| `insert` | `string` |  |
| `method` | `string` |  |
| `names` | `array` |  |
| `ok` | `mixed` |  |
| `overlapLen` | `int` |  |
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
| `frameStart` | `int` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `organism` | `string` |  |
| `provenance` | `array` |  |
| `rareThreshold` | `float` |  |
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
| `avoidEnzymes` | `array` |  |
| `crypticOrfMinAa` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `mixed` |  |
| `gcHigh` | `float` |  |
| `gcLow` | `float` |  |
| `gcWindow` | `int` |  |
| `homopolymerMin` | `int` |  |
| `maxPasses` | `int` |  |
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
| `avoidEnzymes` | `array` |  |
| `crypticOrfMinAa` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `mixed` |  |
| `gcHigh` | `float` |  |
| `gcLow` | `float` |  |
| `gcWindow` | `int` |  |
| `homopolymerMin` | `int` |  |
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
| `minScore` | `float` |  |
| `nuclease` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `searchReverseStrand` | `bool` |  |
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
| `armLength` | `int` |  |
| `blockPam` | `bool` |  |
| `designGenotypingPrimers` | `bool` |  |
| `editEnd` | `int` |  |
| `editStart` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `mixed` |  |
| `guideEnd` | `int` |  |
| `guideStart` | `int` |  |
| `guideStrand` | `string` |  |
| `nuclease` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `replacement` | `string` |  |
| `result` | `array` |  |
| `targetSequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$crispr_hdr_donor = $client->CrisprHdrDonor()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "replacement" => null, // string
    "result" => null, // array
    "targetSequence" => null, // string
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
| `maxMismatches` | `int` |  |
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
| `sequenceA` | `string` |  |
| `sequenceB` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$cross_dimer = $client->CrossDimer()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sequenceA" => null, // string
    "sequenceB" => null, // string
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
| `massNg` | `float` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |
| `volumeUl` | `float` |  |

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
| `enzymeA` | `string` |  |
| `enzymeB` | `string` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$double_digest = $client->DoubleDigest()->create([
    "enzymeA" => null, // string
    "enzymeB" => null, // string
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
| `reactions` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$export_echo_picklist = $client->ExportEchoPicklist()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "reactions" => null, // array
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
| `protocolName` | `string` |  |
| `provenance` | `array` |  |
| `reactions` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$export_opentrons_protocol = $client->ExportOpentronsProtocol()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "reactions" => null, // array
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
| `reactions` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$export_plate_layout = $client->ExportPlateLayout()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "reactions" => null, // array
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
| `clusterCols` | `bool` |  |
| `clusterRows` | `bool` |  |
| `distanceMetric` | `string` |  |
| `gate` | `mixed` |  |
| `genes` | `array` |  |
| `linkage` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `samples` | `array` |  |
| `tool` | `string` |  |
| `values` | `array` |  |
| `zScoreRows` | `bool` |  |

#### Example: Create

```php
$expression_heatmap_cluster = $client->ExpressionHeatmapCluster()->create([
    "genes" => null, // array
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "samples" => null, // array
    "tool" => null, // string
    "values" => null, // array
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
| `qualityOffset` | `int` |  |
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
| `minLength` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `qualityOffset` | `int` |  |
| `qualityThreshold` | `int` |  |
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
| `minAaLength` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `requireStop` | `bool` |  |
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
| `caseMode` | `string` |  |
| `convert` | `string` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `reverse` | `bool` |  |
| `sequence` | `string` |  |
| `stripNonLetters` | `bool` |  |
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
| `collections` | `array` |  |
| `gate` | `mixed` |  |
| `genes` | `array` |  |
| `maxTermSize` | `int` |  |
| `minTermSize` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$functional_enrichment = $client->FunctionalEnrichment()->create([
    "genes" => null, // array
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
| `compareToNamedSet` | `string` |  |
| `dataset` | `string` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `overhangs` | `array` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `riskThreshold` | `float` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$golden_gate_fidelity = $client->GoldenGateFidelity()->create([
    "ok" => null, // mixed
    "overhangs" => null, // array
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
| `jobId` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$id_map_poll = $client->IdMapPoll()->create([
    "jobId" => null, // string
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
| `taxId` | `string` |  |
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
| `forwardPrimer` | `string` |  |
| `gate` | `mixed` |  |
| `maxMismatches` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `reversePrimer` | `string` |  |
| `template` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$in_silico_pcr = $client->InSilicoPcr()->create([
    "forwardPrimer" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "reversePrimer" => null, // string
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
| `addSecondaryMismatch` | `bool` |  |
| `alleleA` | `string` |  |
| `alleleB` | `string` |  |
| `gate` | `mixed` |  |
| `maxAmplicon` | `int` |  |
| `minAmplicon` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `snpPosition` | `int` |  |
| `target` | `string` |  |
| `targetCoreTm` | `float` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$kasp_primer_design = $client->KaspPrimerDesign()->create([
    "alleleA" => null, // string
    "alleleB" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "snpPosition" => null, // int
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
// load() returns the ENTITY — call data_get() for the ListTool record (throws on error).
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
| `dntpMM` | `float` |  |
| `gate` | `mixed` |  |
| `mgMM` | `float` |  |
| `naMM` | `float` |  |
| `ok` | `mixed` |  |
| `oligoNM` | `float` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sequence` | `string` |  |
| `targetTm` | `float` |  |
| `tmTolerance` | `float` |  |
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
| `maxMismatches` | `int` |  |
| `motif` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `searchReverseStrand` | `bool` |  |
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
| `dntpMM` | `float` |  |
| `gate` | `mixed` |  |
| `mgMM` | `float` |  |
| `naMM` | `float` |  |
| `ok` | `mixed` |  |
| `oligoNM` | `float` |  |
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
| `sourceSpecies` | `string` |  |
| `symbols` | `array` |  |
| `targetSpecies` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |

#### Example: Create

```php
$ortholog_map = $client->OrthologMap()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "symbols" => null, // array
    "targetSpecies" => null, // string
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
| `seqA` | `string` |  |
| `seqB` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$pairwise_alignment = $client->PairwiseAlignment()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "seqA" => null, // string
    "seqB" => null, // string
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
| `fileBase64` | `string` |  |
| `fileName` | `string` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$parse_sanger_trace = $client->ParseSangerTrace()->create([
    "fileBase64" => null, // string
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
| `topN` | `int` |  |

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
| `topN` | `int` |  |

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
| `editEnd` | `int` |  |
| `editStart` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `mixed` |  |
| `insertedSeq` | `string` |  |
| `ok` | `mixed` |  |
| `pbsLength` | `int` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `rttHomology` | `int` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$prime_editing_design = $client->PrimeEditingDesign()->create([
    "editEnd" => null, // int
    "editStart" => null, // int
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
| `newSequence` | `string` |  |
| `ok` | `mixed` |  |
| `overlapLength` | `int` |  |
| `pbsLength` | `int` |  |
| `provenance` | `array` |  |
| `replaceEnd` | `int` |  |
| `replaceStart` | `int` |  |
| `result` | `array` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$prime_editing_twin_design = $client->PrimeEditingTwinDesign()->create([
    "newSequence" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "replaceEnd" => null, // int
    "replaceStart" => null, // int
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
| `ampliconMax` | `int` |  |
| `ampliconMin` | `int` |  |
| `dntpMM` | `float` |  |
| `gate` | `mixed` |  |
| `gcMax` | `float` |  |
| `gcMin` | `float` |  |
| `lenMax` | `int` |  |
| `lenMin` | `int` |  |
| `lenOpt` | `int` |  |
| `maxReturn` | `int` |  |
| `mgMM` | `float` |  |
| `naMM` | `float` |  |
| `ok` | `mixed` |  |
| `oligoNM` | `float` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `targetEnd` | `int` |  |
| `targetStart` | `int` |  |
| `template` | `string` |  |
| `tmMax` | `float` |  |
| `tmMaxDiff` | `float` |  |
| `tmMin` | `float` |  |
| `tmOpt` | `float` |  |
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
| `forwardPrimer` | `string` |  |
| `gate` | `mixed` |  |
| `maxMismatches` | `int` |  |
| `maxProductLength` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `reversePrimer` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$primer_specificity = $client->PrimerSpecificity()->create([
    "forwardPrimer" => null, // string
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "reversePrimer" => null, // string
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
| `maxMass` | `float` |  |
| `maxPeptides` | `int` |  |
| `minMass` | `float` |  |
| `missedCleavages` | `int` |  |
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
| `jobId` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$protein_annotate_poll = $client->ProteinAnnotatePoll()->create([
    "jobId" => null, // string
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
| `goterms` | `bool` |  |
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
| `chargeStep` | `float` |  |
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
| `gcContent` | `float` |  |
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
| `enzymes` | `array` |  |
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
| `fileBase64` | `string` |  |
| `fileName` | `string` |  |
| `gate` | `mixed` |  |
| `minCoverage` | `float` |  |
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
| `args` | `array` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$save_permalink = $client->SavePermalink()->create([
    "args" => null, // array
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
| `qualityOffset` | `int` |  |
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
| `endPrimerLength` | `int` |  |
| `gate` | `mixed` |  |
| `maxOrfs` | `int` |  |
| `minOrfAa` | `int` |  |
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
| `maxResults` | `int` |  |
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
| `minSupportingReads` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `reads` | `string` |  |
| `reference` | `string` |  |
| `result` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$sequencing_readback_verify = $client->SequencingReadbackVerify()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "reads" => null, // string
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
| `entries` | `array` |  |
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
| `names` | `array` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sessionId` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$session_get = $client->SessionGet()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sessionId" => null, // string
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
| `args` | `array` |  |
| `fromSession` | `array` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sessionId` | `string` |  |
| `tool` | `string` |  |
| `writeBack` | `array` |  |

#### Example: Create

```php
$session_run = $client->SessionRun()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sessionId" => null, // string
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
| `entries` | `array` |  |
| `gate` | `mixed` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `sessionId` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$session_set = $client->SessionSet()->create([
    "entries" => null, // array
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "sessionId" => null, // string
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
| `minReynolds` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `shRnaLoop` | `string` |  |
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
| `armTmTarget` | `float` |  |
| `dntpMM` | `float` |  |
| `editKind` | `string` |  |
| `frameStart` | `int` |  |
| `gate` | `mixed` |  |
| `mgMM` | `float` |  |
| `naMM` | `float` |  |
| `newBase` | `string` |  |
| `ok` | `mixed` |  |
| `oligoNM` | `float` |  |
| `organism` | `string` |  |
| `position` | `int` |  |
| `provenance` | `array` |  |
| `residue` | `int` |  |
| `result` | `array` |  |
| `style` | `string` |  |
| `targetAa` | `string` |  |
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
| `toStop` | `bool` |  |
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
| `frameStart` | `int` |  |
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
| `armTmTarget` | `float` |  |
| `circular` | `bool` |  |
| `claimedConstruct` | `string` |  |
| `coding` | `bool` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragmentPcrs` | `array` |  |
| `fragments` | `array` |  |
| `frameStart` | `int` |  |
| `gate` | `mixed` |  |
| `insert` | `string` |  |
| `insertPcr` | `array` |  |
| `method` | `string` |  |
| `names` | `array` |  |
| `ok` | `mixed` |  |
| `overlapLen` | `int` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `tool` | `string` |  |
| `vector` | `string` |  |
| `vectorPcr` | `array` |  |

#### Example: Create

```php
$verify_assembly = $client->VerifyAssembly()->create([
    "claimedConstruct" => null, // string
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
| `claimedConstruct` | `string` |  |
| `expectedFrameStart` | `int` |  |
| `gate` | `mixed` |  |
| `insertForwardPrimer` | `string` |  |
| `insertReversePrimer` | `string` |  |
| `insertTemplate` | `string` |  |
| `maxPrimerMismatches` | `int` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` |  |
| `templateCircular` | `bool` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$verify_construct = $client->VerifyConstruct()->create([
    "claimedConstruct" => null, // string
    "insertForwardPrimer" => null, // string
    "insertReversePrimer" => null, // string
    "insertTemplate" => null, // string
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
| `enzymes` | `array` |  |
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
| `rows` | `array` |  |
| `tool` | `string` |  |

#### Example: Create

```php
$volcano_plot_data = $client->VolcanoPlotData()->create([
    "ok" => null, // mixed
    "provenance" => null, // array
    "result" => null, // array
    "rows" => null, // array
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
| `max_results` | `float` |  |
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
