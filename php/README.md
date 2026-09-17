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
print_r($batch->data_get());
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
| `accession` | UniProt accession, e.g. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/alphafold_lookup`

#### AsoDesign

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | Total gapmer length (nt). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `target` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |
| `wing` | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

Operations: Create.

API path: `/aso_design`

#### BaseEditingDesign

| Field | Description |
| --- | --- |
| `editor` | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `target` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/base_editing_design`

#### Batch

| Field | Description |
| --- | --- |
| `args` | Shared tool arguments applied to every record. |
| `capped` | True if input exceeded the record limit. |
| `columns` |  |
| `count` |  |
| `errors` |  |
| `input` | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | Maximum records per call (500). |
| `provenance` |  |
| `rows` |  |
| `tool` | A batchable tool slug (see `GET /batch`). |

Operations: Create, Load.

API path: `/batch`

#### BatchWorkflow

| Field | Description |
| --- | --- |
| `capped` |  |
| `columns` | Flattened "<step>·<tool>·<key>" column headers. |
| `count` |  |
| `errors` |  |
| `input` | Multi-FASTA text or one sequence per line. |
| `limit` | Maximum records per call (200). |
| `provenance` |  |
| `rows` |  |
| `steps` |  |

Operations: Create, Load.

API path: `/workflow`

#### CharacterizeSequence

| Field | Description |
| --- | --- |
| `endPrimerLength` | Length of the naive end primers taken from each end. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/characterize_sequence`

#### CloningSimulate

| Field | Description |
| --- | --- |
| `armTmTarget` | Target annealing Tm (°C) for primer arms. |
| `circular` | Produce a circular product. |
| `enzyme` | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | 3′ enzyme (restriction method). |
| `enzyme5` | 5′ enzyme (restriction method). |
| `fragments` | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | Insert sequence (restriction method). |
| `method` | Assembly method. |
| `names` | Optional labels for each fragment. |
| `ok` |  |
| `overlapLen` | Gibson homology-arm length (bp). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |
| `vector` | Vector sequence (restriction method). |

Operations: Create.

API path: `/cloning_simulate`

#### CodonAdaptationIndex

| Field | Description |
| --- | --- |
| `frameStart` | 1-based position to start reading codons. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `organism` |  |
| `provenance` |  |
| `rareThreshold` | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | Tool-specific output object. |
| `sequence` | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/codon_adaptation_index`

#### CodonOptimize

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `organism` |  |
| `protein` | Protein sequence (one-letter codes). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/codon_optimize`

#### ConstructAutofix

| Field | Description |
| --- | --- |
| `avoidEnzymes` | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | 1-based nucleotide where the reading frame begins. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` |  |
| `gcLow` |  |
| `gcWindow` |  |
| `homopolymerMin` |  |
| `maxPasses` | Repeat full passes until clean or no further progress. |
| `ok` |  |
| `organism` | Codon-usage table to prefer among synonymous options. |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/construct_autofix`

#### ConstructQc

| Field | Description |
| --- | --- |
| `avoidEnzymes` | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | 1-based nucleotide where the reading frame begins. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | GC% above this flags a GC-rich window. |
| `gcLow` | GC% below this flags an AT-rich window. |
| `gcWindow` | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | Minimum run length to flag a homopolymer. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/construct_qc`

#### CrisprGrnaDesign

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | Nuclease id. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `searchReverseStrand` | Also scan the reverse strand for guides. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/crispr_grna_design`

#### CrisprHdrDonor

| Field | Description |
| --- | --- |
| `armLength` | Homology arm length (bp) on each side. |
| `blockPam` | When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut. |
| `designGenotypingPrimers` | Also design a primer pair (on the original targetSequence) whose product spans the edit site. |
| `editEnd` | 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. |
| `editStart` | 1-based start of the region being replaced. |
| `frameStart` | Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | Strand the guide's protospacer is on. |
| `nuclease` | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` |  |
| `provenance` |  |
| `replacement` | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | Tool-specific output object. |
| `targetSequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/crispr_hdr_donor`

#### CrisprOfftargetCheck

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` |  |
| `protospacer` | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/crispr_offtarget_check`

#### CrossDimer

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequenceA` | First oligo (5'→3'). |
| `sequenceB` | Second oligo (5'→3'). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/cross_dimer`

#### DnaMolarity

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | Mass in nanograms. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | The tool slug that ran. |
| `type` | Molecule type. |
| `volumeUl` | Volume in microlitres (0 = unknown; needed for concentration). |

Operations: Create.

API path: `/dna_molarity`

#### DoubleDigest

| Field | Description |
| --- | --- |
| `enzymeA` | First enzyme name (e.g. |
| `enzymeB` | Second enzyme name (e.g. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/double_digest`

#### ExportEchoPicklist

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `reactions` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/export_echo_picklist`

#### ExportOpentronsProtocol

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `protocolName` | Optional protocol name (used in the script's metadata). |
| `provenance` |  |
| `reactions` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/export_opentrons_protocol`

#### ExportPlateLayout

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `reactions` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/export_plate_layout`

#### ExpressionHeatmapCluster

| Field | Description |
| --- | --- |
| `clusterCols` | Cluster (reorder) samples. |
| `clusterRows` | Cluster (reorder) genes. |
| `distanceMetric` | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | Row (gene) labels. |
| `linkage` | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `samples` | Column (sample) labels. |
| `tool` | The tool slug that ran. |
| `values` | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

Operations: Create.

API path: `/expression_heatmap_cluster`

#### FastqQcReport

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` |  |
| `provenance` |  |
| `qualityOffset` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/fastq_qc_report`

#### FastqTrim

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | Reads shorter than this after trimming are dropped. |
| `ok` |  |
| `provenance` |  |
| `qualityOffset` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | 3' quality-trim threshold (Phred score). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/fastq_trim`

#### FindOrf

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | Minimum protein length (aa) to report. |
| `ok` |  |
| `provenance` |  |
| `requireStop` | Only report ORFs terminated by a stop codon. |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/find_orfs`

#### FormatSequence

| Field | Description |
| --- | --- |
| `caseMode` |  |
| `convert` | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `reverse` | Reverse the sequence (no complement). |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | Remove digits, spaces and gaps (keep letters only). |
| `tool` | The tool slug that ran. |
| `width` | Line-wrap width; 0 = single line. |

Operations: Create.

API path: `/format_sequence`

#### FunctionalEnrichment

| Field | Description |
| --- | --- |
| `background` | Custom background/universe gene symbols. |
| `collections` | Which term collections to test. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | Query gene symbols (human, e.g. |
| `maxTermSize` | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | Skip terms/pathways with fewer than this many background genes. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/functional_enrichment`

#### GcContent

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/gc_content`

#### GeneDossier

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/gene_dossier`

#### GeneExpression

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/gene_expression`

#### GeneModel

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/gene_model`

#### GoldenGateFidelity

| Field | Description |
| --- | --- |
| `compareToNamedSet` | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `overhangs` | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `riskThreshold` | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/golden_gate_fidelity`

#### HgvsConvert

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |
| `variant` | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

Operations: Create.

API path: `/hgvs_convert`

#### IdMapPoll

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` |  |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/id_map_poll`

#### IdMapSubmit

| Field | Description |
| --- | --- |
| `from` | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | The ids to map, up to 1000 (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `taxId` | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | Target id type. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/id_map_submit`

#### InSilicoPcr

| Field | Description |
| --- | --- |
| `circular` | Treat the template as circular (plasmid). |
| `forwardPrimer` | Primer 1, 5'→3'. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | Mismatches tolerated per primer. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `reversePrimer` | Primer 2, 5'→3' (order does not matter). |
| `template` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/in_silico_pcr`

#### KaspPrimerDesign

| Field | Description |
| --- | --- |
| `addSecondaryMismatch` | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | First allele (single base) — gets the FAM tail. |
| `alleleB` | Second allele (single base) — gets the HEX tail. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | Minimum amplicon length for the common reverse primer. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `snpPosition` | 1-based position of the SNP on the forward strand. |
| `target` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | The tool slug that ran. |

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
| `dntpMM` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | Divalent cation [Mg2+] (mM). |
| `naMM` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` |  |
| `oligoNM` | Total strand concentration (nM). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | Optional target Tm (°C). |
| `tmTolerance` | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/melting_temperature`

#### MotifFinder

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | Maximum allowed mismatches per match. |
| `motif` | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `searchReverseStrand` | Also search the reverse strand. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/motif_finder`

#### MultipleSequenceAlignment

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/multiple_sequence_alignment`

#### OligoAnalysi

| Field | Description |
| --- | --- |
| `dntpMM` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | Divalent cation [Mg2+] (mM). |
| `naMM` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` |  |
| `oligoNM` | Total strand concentration (nM). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/oligo_analysis`

#### OrthologMap

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sourceSpecies` | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | Ensembl species slug to find homologs in (e.g. |
| `tool` | The tool slug that ran. |
| `type` | Homology type to return. |

Operations: Create.

API path: `/ortholog_map`

#### PairwiseAlignment

| Field | Description |
| --- | --- |
| `gap` | Linear gap penalty (per gap position). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | Match score. |
| `mismatch` | Mismatch penalty. |
| `mode` |  |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `seqA` | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/pairwise_alignment`

#### ParseGenbank

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `text` | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/parse_genbank`

#### ParseSangerTrace

| Field | Description |
| --- | --- |
| `fileBase64` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | Optional original file name (echoed back). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/parse_sanger_trace`

#### PlasmidAnnotate

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/plasmid_annotate`

#### PlasmidDeepAnnotate

| Field | Description |
| --- | --- |
| `circular` | Treat the sequence as a circular plasmid (vs. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/plasmid_deep_annotate`

#### PlasmidFullReport

| Field | Description |
| --- | --- |
| `circular` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |
| `topN` | How many top-ranked backbone candidates to report. |

Operations: Create.

API path: `/plasmid_full_report`

#### PlasmidIdentify

| Field | Description |
| --- | --- |
| `circular` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |
| `topN` | How many top-ranked backbone candidates to report. |

Operations: Create.

API path: `/plasmid_identify`

#### PrimeEditingDesign

| Field | Description |
| --- | --- |
| `editEnd` | 1-based inclusive end of the region being changed. |
| `editStart` | 1-based inclusive start of the region being changed. |
| `frameStart` | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | Replacement bases (forward strand). |
| `ok` |  |
| `pbsLength` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `rttHomology` | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/prime_editing_design`

#### PrimeEditingTwinDesign

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` |  |
| `overlapLength` | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` |  |
| `replaceEnd` | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | 1-based inclusive start of the region being replaced/deleted. |
| `result` | Tool-specific output object. |
| `target` | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/prime_editing_twin_design`

#### PrimerDesign

| Field | Description |
| --- | --- |
| `ampliconMax` |  |
| `ampliconMin` |  |
| `dntpMM` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` |  |
| `gcMin` |  |
| `lenMax` |  |
| `lenMin` |  |
| `lenOpt` |  |
| `maxReturn` | Number of best pairs to return. |
| `mgMM` | Divalent cation [Mg2+] (mM). |
| `naMM` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` |  |
| `oligoNM` | Total strand concentration (nM). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `targetEnd` | 1-based inclusive end of the target region (optional). |
| `targetStart` | 1-based inclusive start of a region the product must span (optional). |
| `template` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` |  |
| `tmMaxDiff` | Max Tm difference within a pair (°C). |
| `tmMin` |  |
| `tmOpt` |  |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/primer_design`

#### PrimerSpecificity

| Field | Description |
| --- | --- |
| `forwardPrimer` | Forward primer, 5'→3'. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `reversePrimer` | Reverse primer, 5'→3'. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/primer_specificity`

#### ProteaseDigestion

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | Cap on the number of returned peptides. |
| `minMass` | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | Allowed missed internal cleavages (0–2). |
| `ok` |  |
| `protease` | Protease or chemical cleavage agent. |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/protease_digestion`

#### ProteinAnnotatePoll

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` |  |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/protein_annotate_poll`

#### ProteinAnnotateSubmit

| Field | Description |
| --- | --- |
| `appl` | Restrict to one member database (e.g. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | Include GO-term cross-references. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/protein_annotate_submit`

#### ProteinHydrophobicity

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `scale` | Amino-acid scale. |
| `sequence` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | The tool slug that ran. |
| `window` | Sliding-window size (clamped to an odd number ≥ 1). |

Operations: Create.

API path: `/protein_hydrophobicity`

#### ProteinProperty

| Field | Description |
| --- | --- |
| `chargeStep` | pH step for the net-charge titration curve (0–14). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/protein_properties`

#### RandomSequence

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` |  |
| `length` | Number of residues to generate. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/random_sequence`

#### RestrictionSite

| Field | Description |
| --- | --- |
| `enzymes` | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/restriction_sites`

#### ReverseComplement

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |
| `type` |  |

Operations: Create.

API path: `/reverse_complement`

#### ReverseTranslate

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` |  |
| `ok` |  |
| `organism` | Codon-usage host (ignored in degenerate mode). |
| `protein` | Protein sequence (one-letter codes; * for stop). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/reverse_translate`

#### RnaFold

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/rna_fold`

#### SangerVsReference

| Field | Description |
| --- | --- |
| `fileBase64` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | Optional original file name (echoed back). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` |  |
| `provenance` |  |
| `read` | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | Expected reference sequence (FASTA or raw). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sanger_vs_reference`

#### SavePermalink

| Field | Description |
| --- | --- |
| `args` | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/save_permalink`

#### SeqfileStat

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` |  |
| `provenance` |  |
| `qualityOffset` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/seqfile_stats`

#### SequenceFetch

| Field | Description |
| --- | --- |
| `accession` | GenBank/RefSeq accession (e.g. |
| `db` | Database to query; auto-detects from the accession format. |
| `format` | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sequence_fetch`

#### SequenceFormatConvert

| Field | Description |
| --- | --- |
| `from` | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | A FASTA or GenBank record to convert. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `to` | Output format. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sequence_format_convert`

#### SequenceReport

| Field | Description |
| --- | --- |
| `endPrimerLength` | Length of the naive end primers taken from each end. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | Minimum ORF length in amino acids. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sequence_report`

#### SequenceSearch

| Field | Description |
| --- | --- |
| `db` |  |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | Gene symbol/name, e.g. |
| `maxResults` | Up to 20. |
| `ok` |  |
| `organism` | Organism name, e.g. |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `term` | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sequence_search`

#### SequencingReadbackVerify

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` |  |
| `provenance` |  |
| `reads` | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | The claimed/expected reference sequence. |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sequencing_readback_verify`

#### SessionCreate

| Field | Description |
| --- | --- |
| `entries` | Initial named entries, e.g. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/session_create`

#### SessionGet

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | Only return these entries; omit to return all of them. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sessionId` |  |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/session_get`

#### SessionRun

| Field | Description |
| --- | --- |
| `args` | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sessionId` |  |
| `tool` | The tool slug that ran. |
| `writeBack` | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

Operations: Create.

API path: `/session_run`

#### SessionSet

| Field | Description |
| --- | --- |
| `entries` | Named entries to add/overwrite, e.g. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sessionId` |  |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/session_set`

#### SirnaDesign

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `shRnaLoop` | Loop sequence used when assembling the shRNA cassette. |
| `target` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sirna_design`

#### SiteDirectedMutagenesi

| Field | Description |
| --- | --- |
| `armTmTarget` | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | Edit at the nucleotide or amino-acid level. |
| `frameStart` | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | Divalent cation [Mg2+] (mM). |
| `naMM` | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | Replacement base (editKind='nt'). |
| `ok` |  |
| `oligoNM` | Total strand concentration (nM). |
| `organism` | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | 1-based position to substitute (editKind='nt'). |
| `provenance` |  |
| `residue` | 1-based residue number to change (editKind='aa'). |
| `result` | Tool-specific output object. |
| `style` | Mutagenic primer style. |
| `targetAa` | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/site_directed_mutagenesis`

#### Translate

| Field | Description |
| --- | --- |
| `frame` |  |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | Stop at the first stop codon. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/translate`

#### VariantAnnotate

| Field | Description |
| --- | --- |
| `assembly` | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |
| `variant` | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

Operations: Create.

API path: `/variant_annotate`

#### VariantComparator

| Field | Description |
| --- | --- |
| `coding` | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | 1-based reading-frame start (used when coding is true). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `query` | Query / variant sequence (raw or FASTA). |
| `reference` | Reference / wild-type sequence (raw or FASTA). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/variant_comparator`

#### VerifyAssembly

| Field | Description |
| --- | --- |
| `armTmTarget` | Target annealing Tm (°C) for primer arms. |
| `circular` | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | The sequence you claim you ended up with. |
| `coding` | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | Type IIS enzyme for Golden Gate. |
| `enzyme3` | 3′ enzyme (restriction method). |
| `enzyme5` | 5′ enzyme (restriction method). |
| `fragmentPcrs` | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | Insert sequence (restriction method). |
| `insertPcr` | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | Assembly method used. |
| `names` | Optional labels for each fragment. |
| `ok` |  |
| `overlapLen` | Gibson homology-arm length (bp). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |
| `vector` | Vector sequence (restriction method). |
| `vectorPcr` | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

Operations: Create.

API path: `/verify_assembly`

#### VerifyConstruct

| Field | Description |
| --- | --- |
| `claimedConstruct` | The final sequence claimed to have been built. |
| `expectedFrameStart` | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | Mismatches tolerated per primer during PCR prediction. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `templateCircular` | Treat insertTemplate as circular (e.g. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/verify_construct`

#### VirtualGel

| Field | Description |
| --- | --- |
| `circular` | Treat the sequence as circular (plasmid). |
| `enzymes` | Enzyme names to digest with. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | DNA ladder to plot alongside the sample lane. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/virtual_gel`

#### VolcanoPlotData

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `rows` | Differential expression rows, one per gene. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/volcano_plot_data`

#### WebSearch

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | Maximum number of results to return (default 5, max 10). |
| `ok` |  |
| `provenance` |  |
| `query` | The search query. |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

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
| `accession` | `string` | UniProt accession, e.g. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `int` | Total gapmer length (nt). |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
| `wing` | `int` | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

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
| `editor` | `string` | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | `int` | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | `int` | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | `string` | The tool slug that ran. |

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
| `args` | `array` | Shared tool arguments applied to every record. |
| `capped` | `bool` | True if input exceeded the record limit. |
| `columns` | `array` |  |
| `count` | `int` |  |
| `errors` | `int` |  |
| `input` | `string` | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | `int` | Maximum records per call (500). |
| `provenance` | `array` |  |
| `rows` | `array` |  |
| `tool` | `string` | A batchable tool slug (see `GET /batch`). |

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
| `columns` | `array` | Flattened "<step>·<tool>·<key>" column headers. |
| `count` | `int` |  |
| `errors` | `int` |  |
| `input` | `string` | Multi-FASTA text or one sequence per line. |
| `limit` | `int` | Maximum records per call (200). |
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
| `endPrimerLength` | `int` | Length of the naive end primers taken from each end. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `int` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `int` | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `armTmTarget` | `float` | Target annealing Tm (°C) for primer arms. |
| `circular` | `bool` | Produce a circular product. |
| `enzyme` | `string` | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | `string` | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | 5′ enzyme (restriction method). |
| `fragments` | `array` | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | Insert sequence (restriction method). |
| `method` | `string` | Assembly method. |
| `names` | `array` | Optional labels for each fragment. |
| `ok` | `mixed` |  |
| `overlapLen` | `int` | Gibson homology-arm length (bp). |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `vector` | `string` | Vector sequence (restriction method). |

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
| `frameStart` | `int` | 1-based position to start reading codons. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `organism` | `string` |  |
| `provenance` | `array` |  |
| `rareThreshold` | `float` | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `organism` | `string` |  |
| `protein` | `string` | Protein sequence (one-letter codes). |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `avoidEnzymes` | `array` | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | `int` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `int` | 1-based nucleotide where the reading frame begins. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `float` |  |
| `gcLow` | `float` |  |
| `gcWindow` | `int` |  |
| `homopolymerMin` | `int` |  |
| `maxPasses` | `int` | Repeat full passes until clean or no further progress. |
| `ok` | `mixed` |  |
| `organism` | `string` | Codon-usage table to prefer among synonymous options. |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `avoidEnzymes` | `array` | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | `int` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `int` | 1-based nucleotide where the reading frame begins. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `float` | GC% above this flags a GC-rich window. |
| `gcLow` | `float` | GC% below this flags an AT-rich window. |
| `gcWindow` | `int` | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | `int` | Minimum run length to flag a homopolymer. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | `float` | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | `string` | Nuclease id. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `searchReverseStrand` | `bool` | Also scan the reverse strand for guides. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `armLength` | `int` | Homology arm length (bp) on each side. |
| `blockPam` | `bool` | When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut. |
| `designGenotypingPrimers` | `bool` | Also design a primer pair (on the original targetSequence) whose product spans the edit site. |
| `editEnd` | `int` | 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. |
| `editStart` | `int` | 1-based start of the region being replaced. |
| `frameStart` | `int` | Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | `int` | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | `int` | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | `string` | Strand the guide's protospacer is on. |
| `nuclease` | `string` | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `replacement` | `string` | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | `array` | Tool-specific output object. |
| `targetSequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | `string` | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` | `mixed` |  |
| `protospacer` | `string` | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequenceA` | `string` | First oligo (5'→3'). |
| `sequenceB` | `string` | Second oligo (5'→3'). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `int` | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | `float` | Mass in nanograms. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | `string` | The tool slug that ran. |
| `type` | `string` | Molecule type. |
| `volumeUl` | `float` | Volume in microlitres (0 = unknown; needed for concentration). |

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
| `enzymeA` | `string` | First enzyme name (e.g. |
| `enzymeB` | `string` | Second enzyme name (e.g. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `reactions` | `array` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `protocolName` | `string` | Optional protocol name (used in the script's metadata). |
| `provenance` | `array` |  |
| `reactions` | `array` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `reactions` | `array` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `clusterCols` | `bool` | Cluster (reorder) samples. |
| `clusterRows` | `bool` | Cluster (reorder) genes. |
| `distanceMetric` | `string` | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `array` | Row (gene) labels. |
| `linkage` | `string` | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `samples` | `array` | Column (sample) labels. |
| `tool` | `string` | The tool slug that ran. |
| `values` | `array` | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | `bool` | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `qualityOffset` | `int` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | `int` | Reads shorter than this after trimming are dropped. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `qualityOffset` | `int` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | `int` | 3' quality-trim threshold (Phred score). |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | `int` | Minimum protein length (aa) to report. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `requireStop` | `bool` | Only report ORFs terminated by a stop codon. |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `convert` | `string` | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `reverse` | `bool` | Reverse the sequence (no complement). |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | `bool` | Remove digits, spaces and gaps (keep letters only). |
| `tool` | `string` | The tool slug that ran. |
| `width` | `int` | Line-wrap width; 0 = single line. |

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
| `background` | `array` | Custom background/universe gene symbols. |
| `collections` | `array` | Which term collections to test. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `array` | Query gene symbols (human, e.g. |
| `maxTermSize` | `int` | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | `int` | Skip terms/pathways with fewer than this many background genes. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `compareToNamedSet` | `string` | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | `string` | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `overhangs` | `array` | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `riskThreshold` | `float` | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `variant` | `string` | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `from` | `string` | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | `array` | The ids to map, up to 1000 (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `taxId` | `string` | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | `string` | Target id type. |
| `tool` | `string` | The tool slug that ran. |

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
| `circular` | `bool` | Treat the template as circular (plasmid). |
| `forwardPrimer` | `string` | Primer 1, 5'→3'. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | Mismatches tolerated per primer. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `reversePrimer` | `string` | Primer 2, 5'→3' (order does not matter). |
| `template` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `addSecondaryMismatch` | `bool` | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | `string` | First allele (single base) — gets the FAM tail. |
| `alleleB` | `string` | Second allele (single base) — gets the HEX tail. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | `int` | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | `int` | Minimum amplicon length for the common reverse primer. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `snpPosition` | `int` | 1-based position of the SNP on the forward strand. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | `float` | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | `string` | The tool slug that ran. |

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
| `dntpMM` | `float` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `mixed` |  |
| `oligoNM` | `float` | Total strand concentration (nM). |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | `float` | Optional target Tm (°C). |
| `tmTolerance` | `float` | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | Maximum allowed mismatches per match. |
| `motif` | `string` | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `searchReverseStrand` | `bool` | Also search the reverse strand. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `dntpMM` | `float` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `mixed` |  |
| `oligoNM` | `float` | Total strand concentration (nM). |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sourceSpecies` | `string` | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | `array` | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | `string` | Ensembl species slug to find homologs in (e.g. |
| `tool` | `string` | The tool slug that ran. |
| `type` | `string` | Homology type to return. |

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
| `gap` | `float` | Linear gap penalty (per gap position). |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | `float` | Match score. |
| `mismatch` | `float` | Mismatch penalty. |
| `mode` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `seqA` | `string` | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | `string` | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `text` | `string` | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | `string` | The tool slug that ran. |

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
| `fileBase64` | `string` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | Optional original file name (echoed back). |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `circular` | `bool` | Treat the sequence as a circular plasmid (vs. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `circular` | `bool` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
| `topN` | `int` | How many top-ranked backbone candidates to report. |

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
| `circular` | `bool` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
| `topN` | `int` | How many top-ranked backbone candidates to report. |

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
| `editEnd` | `int` | 1-based inclusive end of the region being changed. |
| `editStart` | `int` | 1-based inclusive start of the region being changed. |
| `frameStart` | `int` | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | `string` | Replacement bases (forward strand). |
| `ok` | `mixed` |  |
| `pbsLength` | `int` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `rttHomology` | `int` | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | `string` | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | `string` | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` | `mixed` |  |
| `overlapLength` | `int` | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | `int` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `array` |  |
| `replaceEnd` | `int` | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | `int` | 1-based inclusive start of the region being replaced/deleted. |
| `result` | `array` | Tool-specific output object. |
| `target` | `string` | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | `string` | The tool slug that ran. |

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
| `dntpMM` | `float` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` | `float` |  |
| `gcMin` | `float` |  |
| `lenMax` | `int` |  |
| `lenMin` | `int` |  |
| `lenOpt` | `int` |  |
| `maxReturn` | `int` | Number of best pairs to return. |
| `mgMM` | `float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `mixed` |  |
| `oligoNM` | `float` | Total strand concentration (nM). |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `targetEnd` | `int` | 1-based inclusive end of the target region (optional). |
| `targetStart` | `int` | 1-based inclusive start of a region the product must span (optional). |
| `template` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` | `float` |  |
| `tmMaxDiff` | `float` | Max Tm difference within a pair (°C). |
| `tmMin` | `float` |  |
| `tmOpt` | `float` |  |
| `tool` | `string` | The tool slug that ran. |

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
| `forwardPrimer` | `string` | Forward primer, 5'→3'. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | `int` | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `reversePrimer` | `string` | Reverse primer, 5'→3'. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | `float` | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | `int` | Cap on the number of returned peptides. |
| `minMass` | `float` | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | `int` | Allowed missed internal cleavages (0–2). |
| `ok` | `mixed` |  |
| `protease` | `string` | Protease or chemical cleavage agent. |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` |  |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `appl` | `string` | Restrict to one member database (e.g. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | `bool` | Include GO-term cross-references. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `scale` | `string` | Amino-acid scale. |
| `sequence` | `string` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | The tool slug that ran. |
| `window` | `int` | Sliding-window size (clamped to an odd number ≥ 1). |

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
| `chargeStep` | `float` | pH step for the net-charge titration curve (0–14). |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | `float` | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` | `string` |  |
| `length` | `int` | Number of residues to generate. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `enzymes` | `array` | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` | `string` |  |
| `ok` | `mixed` |  |
| `organism` | `string` | Codon-usage host (ignored in degenerate mode). |
| `protein` | `string` | Protein sequence (one-letter codes; * for stop). |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `fileBase64` | `string` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | Optional original file name (echoed back). |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | `float` | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `read` | `string` | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | `string` | Expected reference sequence (FASTA or raw). |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `args` | `array` | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `qualityOffset` | `int` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `accession` | `string` | GenBank/RefSeq accession (e.g. |
| `db` | `string` | Database to query; auto-detects from the accession format. |
| `format` | `string` | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `from` | `string` | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | A FASTA or GenBank record to convert. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `to` | `string` | Output format. |
| `tool` | `string` | The tool slug that ran. |

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
| `endPrimerLength` | `int` | Length of the naive end primers taken from each end. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `int` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `int` | Minimum ORF length in amino acids. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Gene symbol/name, e.g. |
| `maxResults` | `int` | Up to 20. |
| `ok` | `mixed` |  |
| `organism` | `string` | Organism name, e.g. |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `term` | `string` | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | `int` | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `reads` | `string` | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | `string` | The claimed/expected reference sequence. |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `entries` | `array` | Initial named entries, e.g. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | `array` | Only return these entries; omit to return all of them. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sessionId` | `string` |  |
| `tool` | `string` | The tool slug that ran. |

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
| `args` | `array` | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | `array` | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sessionId` | `string` |  |
| `tool` | `string` | The tool slug that ran. |
| `writeBack` | `array` | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

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
| `entries` | `array` | Named entries to add/overwrite, e.g. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sessionId` | `string` |  |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | `int` | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `shRnaLoop` | `string` | Loop sequence used when assembling the shRNA cassette. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `armTmTarget` | `float` | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | `float` | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | `string` | Edit at the nucleotide or amino-acid level. |
| `frameStart` | `int` | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | `string` | Replacement base (editKind='nt'). |
| `ok` | `mixed` |  |
| `oligoNM` | `float` | Total strand concentration (nM). |
| `organism` | `string` | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | `int` | 1-based position to substitute (editKind='nt'). |
| `provenance` | `array` |  |
| `residue` | `int` | 1-based residue number to change (editKind='aa'). |
| `result` | `array` | Tool-specific output object. |
| `style` | `string` | Mutagenic primer style. |
| `targetAa` | `string` | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | `bool` | Stop at the first stop codon. |
| `tool` | `string` | The tool slug that ran. |

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
| `assembly` | `string` | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `variant` | `string` | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

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
| `coding` | `bool` | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | `int` | 1-based reading-frame start (used when coding is true). |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `query` | `string` | Query / variant sequence (raw or FASTA). |
| `reference` | `string` | Reference / wild-type sequence (raw or FASTA). |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `armTmTarget` | `float` | Target annealing Tm (°C) for primer arms. |
| `circular` | `bool` | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | `string` | The sequence you claim you ended up with. |
| `coding` | `bool` | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | `string` | Type IIS enzyme for Golden Gate. |
| `enzyme3` | `string` | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | 5′ enzyme (restriction method). |
| `fragmentPcrs` | `array` | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | `array` | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | `int` | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | Insert sequence (restriction method). |
| `insertPcr` | `array` | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | `string` | Assembly method used. |
| `names` | `array` | Optional labels for each fragment. |
| `ok` | `mixed` |  |
| `overlapLen` | `int` | Gibson homology-arm length (bp). |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `vector` | `string` | Vector sequence (restriction method). |
| `vectorPcr` | `array` | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

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
| `claimedConstruct` | `string` | The final sequence claimed to have been built. |
| `expectedFrameStart` | `int` | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | `string` | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | `string` | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | `string` | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | `int` | Mismatches tolerated per primer during PCR prediction. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `templateCircular` | `bool` | Treat insertTemplate as circular (e.g. |
| `tool` | `string` | The tool slug that ran. |

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
| `circular` | `bool` | Treat the sequence as circular (plasmid). |
| `enzymes` | `array` | Enzyme names to digest with. |
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | `string` | DNA ladder to plot alongside the sample lane. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `result` | `array` | Tool-specific output object. |
| `rows` | `array` | Differential expression rows, one per gene. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `mixed` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | `float` | Maximum number of results to return (default 5, max 10). |
| `ok` | `mixed` |  |
| `provenance` | `array` |  |
| `query` | `string` | The search query. |
| `result` | `array` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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

## Features

This SDK ships 4 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`ratelimit`](#ratelimit) | Client-side rate limiting via a token bucket |
| [`retry`](#retry) | Automatic retry of transient failures with exponential backoff |
| [`test`](#test) | In-memory mock transport for testing without a live server |
| [`timeout`](#timeout) | Per-request timeout with transport abort |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### ratelimit

Client-side rate limiting via a token bucket.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Automatic retry of transient failures with exponential backoff.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Per-request timeout with transport abort.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


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

- **RatelimitFeature**: Client-side rate limiting via a token bucket
- **RetryFeature**: Automatic retry of transient failures with exponential backoff
- **TestFeature**: In-memory mock transport for testing without a live server
- **TimeoutFeature**: Per-request timeout with transport abort

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
├── schema.php                     -- Generated option + entity specs
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
