# SeqbenchMcp PHP SDK Reference

Complete API reference for the SeqbenchMcp PHP SDK.


## SeqbenchMcpSDK

### Constructor

```php
require_once __DIR__ . '/seqbenchmcp_sdk.php';

$client = new SeqbenchMcpSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `SeqbenchMcpSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = SeqbenchMcpSDK::test();
```


### Instance Methods

#### `AlphafoldLookup($data = null)`

Create a new `AlphafoldLookupEntity` instance. Pass `null` for no initial data.

#### `AsoDesign($data = null)`

Create a new `AsoDesignEntity` instance. Pass `null` for no initial data.

#### `BaseEditingDesign($data = null)`

Create a new `BaseEditingDesignEntity` instance. Pass `null` for no initial data.

#### `Batch($data = null)`

Create a new `BatchEntity` instance. Pass `null` for no initial data.

#### `BatchWorkflow($data = null)`

Create a new `BatchWorkflowEntity` instance. Pass `null` for no initial data.

#### `CharacterizeSequence($data = null)`

Create a new `CharacterizeSequenceEntity` instance. Pass `null` for no initial data.

#### `CloningSimulate($data = null)`

Create a new `CloningSimulateEntity` instance. Pass `null` for no initial data.

#### `CodonAdaptationIndex($data = null)`

Create a new `CodonAdaptationIndexEntity` instance. Pass `null` for no initial data.

#### `CodonOptimize($data = null)`

Create a new `CodonOptimizeEntity` instance. Pass `null` for no initial data.

#### `ConstructAutofix($data = null)`

Create a new `ConstructAutofixEntity` instance. Pass `null` for no initial data.

#### `ConstructQc($data = null)`

Create a new `ConstructQcEntity` instance. Pass `null` for no initial data.

#### `CrisprGrnaDesign($data = null)`

Create a new `CrisprGrnaDesignEntity` instance. Pass `null` for no initial data.

#### `CrisprHdrDonor($data = null)`

Create a new `CrisprHdrDonorEntity` instance. Pass `null` for no initial data.

#### `CrisprOfftargetCheck($data = null)`

Create a new `CrisprOfftargetCheckEntity` instance. Pass `null` for no initial data.

#### `CrossDimer($data = null)`

Create a new `CrossDimerEntity` instance. Pass `null` for no initial data.

#### `DnaMolarity($data = null)`

Create a new `DnaMolarityEntity` instance. Pass `null` for no initial data.

#### `DoubleDigest($data = null)`

Create a new `DoubleDigestEntity` instance. Pass `null` for no initial data.

#### `ExportEchoPicklist($data = null)`

Create a new `ExportEchoPicklistEntity` instance. Pass `null` for no initial data.

#### `ExportOpentronsProtocol($data = null)`

Create a new `ExportOpentronsProtocolEntity` instance. Pass `null` for no initial data.

#### `ExportPlateLayout($data = null)`

Create a new `ExportPlateLayoutEntity` instance. Pass `null` for no initial data.

#### `ExpressionHeatmapCluster($data = null)`

Create a new `ExpressionHeatmapClusterEntity` instance. Pass `null` for no initial data.

#### `FastqQcReport($data = null)`

Create a new `FastqQcReportEntity` instance. Pass `null` for no initial data.

#### `FastqTrim($data = null)`

Create a new `FastqTrimEntity` instance. Pass `null` for no initial data.

#### `FindOrf($data = null)`

Create a new `FindOrfEntity` instance. Pass `null` for no initial data.

#### `FormatSequence($data = null)`

Create a new `FormatSequenceEntity` instance. Pass `null` for no initial data.

#### `FunctionalEnrichment($data = null)`

Create a new `FunctionalEnrichmentEntity` instance. Pass `null` for no initial data.

#### `GcContent($data = null)`

Create a new `GcContentEntity` instance. Pass `null` for no initial data.

#### `GeneDossier($data = null)`

Create a new `GeneDossierEntity` instance. Pass `null` for no initial data.

#### `GeneExpression($data = null)`

Create a new `GeneExpressionEntity` instance. Pass `null` for no initial data.

#### `GeneModel($data = null)`

Create a new `GeneModelEntity` instance. Pass `null` for no initial data.

#### `GoldenGateFidelity($data = null)`

Create a new `GoldenGateFidelityEntity` instance. Pass `null` for no initial data.

#### `HgvsConvert($data = null)`

Create a new `HgvsConvertEntity` instance. Pass `null` for no initial data.

#### `IdMapPoll($data = null)`

Create a new `IdMapPollEntity` instance. Pass `null` for no initial data.

#### `IdMapSubmit($data = null)`

Create a new `IdMapSubmitEntity` instance. Pass `null` for no initial data.

#### `InSilicoPcr($data = null)`

Create a new `InSilicoPcrEntity` instance. Pass `null` for no initial data.

#### `KaspPrimerDesign($data = null)`

Create a new `KaspPrimerDesignEntity` instance. Pass `null` for no initial data.

#### `ListTool($data = null)`

Create a new `ListToolEntity` instance. Pass `null` for no initial data.

#### `MeltingTemperature($data = null)`

Create a new `MeltingTemperatureEntity` instance. Pass `null` for no initial data.

#### `MotifFinder($data = null)`

Create a new `MotifFinderEntity` instance. Pass `null` for no initial data.

#### `MultipleSequenceAlignment($data = null)`

Create a new `MultipleSequenceAlignmentEntity` instance. Pass `null` for no initial data.

#### `OligoAnalysi($data = null)`

Create a new `OligoAnalysiEntity` instance. Pass `null` for no initial data.

#### `OrthologMap($data = null)`

Create a new `OrthologMapEntity` instance. Pass `null` for no initial data.

#### `PairwiseAlignment($data = null)`

Create a new `PairwiseAlignmentEntity` instance. Pass `null` for no initial data.

#### `ParseGenbank($data = null)`

Create a new `ParseGenbankEntity` instance. Pass `null` for no initial data.

#### `ParseSangerTrace($data = null)`

Create a new `ParseSangerTraceEntity` instance. Pass `null` for no initial data.

#### `PlasmidAnnotate($data = null)`

Create a new `PlasmidAnnotateEntity` instance. Pass `null` for no initial data.

#### `PlasmidDeepAnnotate($data = null)`

Create a new `PlasmidDeepAnnotateEntity` instance. Pass `null` for no initial data.

#### `PlasmidFullReport($data = null)`

Create a new `PlasmidFullReportEntity` instance. Pass `null` for no initial data.

#### `PlasmidIdentify($data = null)`

Create a new `PlasmidIdentifyEntity` instance. Pass `null` for no initial data.

#### `PrimeEditingDesign($data = null)`

Create a new `PrimeEditingDesignEntity` instance. Pass `null` for no initial data.

#### `PrimeEditingTwinDesign($data = null)`

Create a new `PrimeEditingTwinDesignEntity` instance. Pass `null` for no initial data.

#### `PrimerDesign($data = null)`

Create a new `PrimerDesignEntity` instance. Pass `null` for no initial data.

#### `PrimerSpecificity($data = null)`

Create a new `PrimerSpecificityEntity` instance. Pass `null` for no initial data.

#### `ProteaseDigestion($data = null)`

Create a new `ProteaseDigestionEntity` instance. Pass `null` for no initial data.

#### `ProteinAnnotatePoll($data = null)`

Create a new `ProteinAnnotatePollEntity` instance. Pass `null` for no initial data.

#### `ProteinAnnotateSubmit($data = null)`

Create a new `ProteinAnnotateSubmitEntity` instance. Pass `null` for no initial data.

#### `ProteinHydrophobicity($data = null)`

Create a new `ProteinHydrophobicityEntity` instance. Pass `null` for no initial data.

#### `ProteinProperty($data = null)`

Create a new `ProteinPropertyEntity` instance. Pass `null` for no initial data.

#### `RandomSequence($data = null)`

Create a new `RandomSequenceEntity` instance. Pass `null` for no initial data.

#### `RestrictionSite($data = null)`

Create a new `RestrictionSiteEntity` instance. Pass `null` for no initial data.

#### `ReverseComplement($data = null)`

Create a new `ReverseComplementEntity` instance. Pass `null` for no initial data.

#### `ReverseTranslate($data = null)`

Create a new `ReverseTranslateEntity` instance. Pass `null` for no initial data.

#### `RnaFold($data = null)`

Create a new `RnaFoldEntity` instance. Pass `null` for no initial data.

#### `SangerVsReference($data = null)`

Create a new `SangerVsReferenceEntity` instance. Pass `null` for no initial data.

#### `SavePermalink($data = null)`

Create a new `SavePermalinkEntity` instance. Pass `null` for no initial data.

#### `SeqfileStat($data = null)`

Create a new `SeqfileStatEntity` instance. Pass `null` for no initial data.

#### `SequenceFetch($data = null)`

Create a new `SequenceFetchEntity` instance. Pass `null` for no initial data.

#### `SequenceFormatConvert($data = null)`

Create a new `SequenceFormatConvertEntity` instance. Pass `null` for no initial data.

#### `SequenceReport($data = null)`

Create a new `SequenceReportEntity` instance. Pass `null` for no initial data.

#### `SequenceSearch($data = null)`

Create a new `SequenceSearchEntity` instance. Pass `null` for no initial data.

#### `SequencingReadbackVerify($data = null)`

Create a new `SequencingReadbackVerifyEntity` instance. Pass `null` for no initial data.

#### `SessionCreate($data = null)`

Create a new `SessionCreateEntity` instance. Pass `null` for no initial data.

#### `SessionGet($data = null)`

Create a new `SessionGetEntity` instance. Pass `null` for no initial data.

#### `SessionRun($data = null)`

Create a new `SessionRunEntity` instance. Pass `null` for no initial data.

#### `SessionSet($data = null)`

Create a new `SessionSetEntity` instance. Pass `null` for no initial data.

#### `SirnaDesign($data = null)`

Create a new `SirnaDesignEntity` instance. Pass `null` for no initial data.

#### `SiteDirectedMutagenesi($data = null)`

Create a new `SiteDirectedMutagenesiEntity` instance. Pass `null` for no initial data.

#### `Translate($data = null)`

Create a new `TranslateEntity` instance. Pass `null` for no initial data.

#### `VariantAnnotate($data = null)`

Create a new `VariantAnnotateEntity` instance. Pass `null` for no initial data.

#### `VariantComparator($data = null)`

Create a new `VariantComparatorEntity` instance. Pass `null` for no initial data.

#### `VerifyAssembly($data = null)`

Create a new `VerifyAssemblyEntity` instance. Pass `null` for no initial data.

#### `VerifyConstruct($data = null)`

Create a new `VerifyConstructEntity` instance. Pass `null` for no initial data.

#### `VirtualGel($data = null)`

Create a new `VirtualGelEntity` instance. Pass `null` for no initial data.

#### `VolcanoPlotData($data = null)`

Create a new `VolcanoPlotDataEntity` instance. Pass `null` for no initial data.

#### `WebSearch($data = null)`

Create a new `WebSearchEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): SeqbenchMcpUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## AlphafoldLookupEntity

```php
$alphafold_lookup = $client->AlphafoldLookup();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `string` | Yes | UniProt accession, e.g. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->AlphafoldLookup()->create([
  "accession" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AlphafoldLookupEntity`

Create a new `AlphafoldLookupEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## AsoDesignEntity

```php
$aso_design = $client->AsoDesign();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `int` | No | Total gapmer length (nt). |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `wing` | `int` | No | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->AsoDesign()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "target" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AsoDesignEntity`

Create a new `AsoDesignEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## BaseEditingDesignEntity

```php
$base_editing_design = $client->BaseEditingDesign();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `editor` | `string` | No | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | `int` | No | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | `int` | No | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->BaseEditingDesign()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "target" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): BaseEditingDesignEntity`

Create a new `BaseEditingDesignEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## BatchEntity

```php
$batch = $client->Batch();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `array` | No | Shared tool arguments applied to every record. |
| `capped` | `bool` | Yes | True if input exceeded the record limit. |
| `columns` | `array` | Yes |  |
| `count` | `int` | Yes |  |
| `errors` | `int` | Yes |  |
| `input` | `string` | Yes | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | `int` | Yes | Maximum records per call (500). |
| `provenance` | `array` | Yes |  |
| `rows` | `array` | Yes |  |
| `tool` | `string` | Yes | A batchable tool slug (see `GET /batch`). |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Batch()->create([
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

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Batch()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): BatchEntity`

Create a new `BatchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## BatchWorkflowEntity

```php
$batch__workflow = $client->BatchWorkflow();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `capped` | `bool` | Yes |  |
| `columns` | `array` | Yes | Flattened "<step>·<tool>·<key>" column headers. |
| `count` | `int` | Yes |  |
| `errors` | `int` | Yes |  |
| `input` | `string` | Yes | Multi-FASTA text or one sequence per line. |
| `limit` | `int` | Yes | Maximum records per call (200). |
| `provenance` | `array` | Yes |  |
| `rows` | `array` | Yes |  |
| `steps` | `array` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->BatchWorkflow()->create([
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

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->BatchWorkflow()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): BatchWorkflowEntity`

Create a new `BatchWorkflowEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CharacterizeSequenceEntity

```php
$characterize_sequence = $client->CharacterizeSequence();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `endPrimerLength` | `int` | No | Length of the naive end primers taken from each end. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `int` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `int` | No | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CharacterizeSequence()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CharacterizeSequenceEntity`

Create a new `CharacterizeSequenceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CloningSimulateEntity

```php
$cloning_simulate = $client->CloningSimulate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `float` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `bool` | No | Produce a circular product. |
| `enzyme` | `string` | No | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | `string` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | No | 5′ enzyme (restriction method). |
| `fragments` | `array` | No | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | No | Insert sequence (restriction method). |
| `method` | `string` | Yes | Assembly method. |
| `names` | `array` | No | Optional labels for each fragment. |
| `ok` | `mixed` | Yes |  |
| `overlapLen` | `int` | No | Gibson homology-arm length (bp). |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `vector` | `string` | No | Vector sequence (restriction method). |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CloningSimulate()->create([
  "method" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CloningSimulateEntity`

Create a new `CloningSimulateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CodonAdaptationIndexEntity

```php
$codon_adaptation_index = $client->CodonAdaptationIndex();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frameStart` | `int` | No | 1-based position to start reading codons. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `array` | Yes |  |
| `rareThreshold` | `float` | No | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CodonAdaptationIndex()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CodonAdaptationIndexEntity`

Create a new `CodonAdaptationIndexEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CodonOptimizeEntity

```php
$codon_optimize = $client->CodonOptimize();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes | Protein sequence (one-letter codes). |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CodonOptimize()->create([
  "ok" => null, // mixed
  "protein" => null, // string
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CodonOptimizeEntity`

Create a new `CodonOptimizeEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ConstructAutofixEntity

```php
$construct_autofix = $client->ConstructAutofix();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoidEnzymes` | `array` | No | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | `int` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `int` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `float` | No |  |
| `gcLow` | `float` | No |  |
| `gcWindow` | `int` | No |  |
| `homopolymerMin` | `int` | No |  |
| `maxPasses` | `int` | No | Repeat full passes until clean or no further progress. |
| `ok` | `mixed` | Yes |  |
| `organism` | `string` | No | Codon-usage table to prefer among synonymous options. |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ConstructAutofix()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ConstructAutofixEntity`

Create a new `ConstructAutofixEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ConstructQcEntity

```php
$construct_qc = $client->ConstructQc();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoidEnzymes` | `array` | No | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | `int` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `int` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `float` | No | GC% above this flags a GC-rich window. |
| `gcLow` | `float` | No | GC% below this flags an AT-rich window. |
| `gcWindow` | `int` | No | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | `int` | No | Minimum run length to flag a homopolymer. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ConstructQc()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ConstructQcEntity`

Create a new `ConstructQcEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CrisprGrnaDesignEntity

```php
$crispr_grna_design = $client->CrisprGrnaDesign();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | `float` | No | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | `string` | No | Nuclease id. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `bool` | No | Also scan the reverse strand for guides. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CrisprGrnaDesign()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CrisprGrnaDesignEntity`

Create a new `CrisprGrnaDesignEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CrisprHdrDonorEntity

```php
$crispr_hdr_donor = $client->CrisprHdrDonor();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armLength` | `int` | No | Homology arm length (bp) on each side. |
| `blockPam` | `bool` | No | When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut. |
| `designGenotypingPrimers` | `bool` | No | Also design a primer pair (on the original targetSequence) whose product spans the edit site. |
| `editEnd` | `int` | No | 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. |
| `editStart` | `int` | No | 1-based start of the region being replaced. |
| `frameStart` | `int` | No | Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | `int` | No | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | `int` | No | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | `string` | No | Strand the guide's protospacer is on. |
| `nuclease` | `string` | No | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `replacement` | `string` | Yes | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | `array` | Yes | Tool-specific output object. |
| `targetSequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CrisprHdrDonor()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "replacement" => null, // string
  "result" => null, // array
  "targetSequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CrisprHdrDonorEntity`

Create a new `CrisprHdrDonorEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CrisprOfftargetCheckEntity

```php
$crispr_offtarget_check = $client->CrisprOfftargetCheck();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | `string` | No | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` | `mixed` | Yes |  |
| `protospacer` | `string` | Yes | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CrisprOfftargetCheck()->create([
  "ok" => null, // mixed
  "protospacer" => null, // string
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CrisprOfftargetCheckEntity`

Create a new `CrisprOfftargetCheckEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CrossDimerEntity

```php
$cross_dimer = $client->CrossDimer();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequenceA` | `string` | Yes | First oligo (5'→3'). |
| `sequenceB` | `string` | Yes | Second oligo (5'→3'). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CrossDimer()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequenceA" => null, // string
  "sequenceB" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CrossDimerEntity`

Create a new `CrossDimerEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## DnaMolarityEntity

```php
$dna_molarity = $client->DnaMolarity();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `int` | No | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | `float` | No | Mass in nanograms. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | No | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `type` | `string` | No | Molecule type. |
| `volumeUl` | `float` | No | Volume in microlitres (0 = unknown; needed for concentration). |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->DnaMolarity()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DnaMolarityEntity`

Create a new `DnaMolarityEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## DoubleDigestEntity

```php
$double_digest = $client->DoubleDigest();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzymeA` | `string` | Yes | First enzyme name (e.g. |
| `enzymeB` | `string` | Yes | Second enzyme name (e.g. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->DoubleDigest()->create([
  "enzymeA" => null, // string
  "enzymeB" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DoubleDigestEntity`

Create a new `DoubleDigestEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ExportEchoPicklistEntity

```php
$export_echo_picklist = $client->ExportEchoPicklist();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `reactions` | `array` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ExportEchoPicklist()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "reactions" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ExportEchoPicklistEntity`

Create a new `ExportEchoPicklistEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ExportOpentronsProtocolEntity

```php
$export_opentrons_protocol = $client->ExportOpentronsProtocol();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `protocolName` | `string` | No | Optional protocol name (used in the script's metadata). |
| `provenance` | `array` | Yes |  |
| `reactions` | `array` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ExportOpentronsProtocol()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "reactions" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ExportOpentronsProtocolEntity`

Create a new `ExportOpentronsProtocolEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ExportPlateLayoutEntity

```php
$export_plate_layout = $client->ExportPlateLayout();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `reactions` | `array` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ExportPlateLayout()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "reactions" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ExportPlateLayoutEntity`

Create a new `ExportPlateLayoutEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ExpressionHeatmapClusterEntity

```php
$expression_heatmap_cluster = $client->ExpressionHeatmapCluster();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clusterCols` | `bool` | No | Cluster (reorder) samples. |
| `clusterRows` | `bool` | No | Cluster (reorder) genes. |
| `distanceMetric` | `string` | No | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `array` | Yes | Row (gene) labels. |
| `linkage` | `string` | No | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `samples` | `array` | Yes | Column (sample) labels. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `values` | `array` | Yes | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | `bool` | No | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ExpressionHeatmapCluster()->create([
  "genes" => null, // array
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "samples" => null, // array
  "tool" => null, // string
  "values" => null, // array
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ExpressionHeatmapClusterEntity`

Create a new `ExpressionHeatmapClusterEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## FastqQcReportEntity

```php
$fastq_qc_report = $client->FastqQcReport();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `qualityOffset` | `int` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->FastqQcReport()->create([
  "input" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): FastqQcReportEntity`

Create a new `FastqQcReportEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## FastqTrimEntity

```php
$fastq_trim = $client->FastqTrim();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | `int` | No | Reads shorter than this after trimming are dropped. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `qualityOffset` | `int` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | `int` | No | 3' quality-trim threshold (Phred score). |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->FastqTrim()->create([
  "input" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): FastqTrimEntity`

Create a new `FastqTrimEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## FindOrfEntity

```php
$find_orf = $client->FindOrf();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | `int` | No | Minimum protein length (aa) to report. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `requireStop` | `bool` | No | Only report ORFs terminated by a stop codon. |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->FindOrf()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): FindOrfEntity`

Create a new `FindOrfEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## FormatSequenceEntity

```php
$format_sequence = $client->FormatSequence();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `caseMode` | `string` | No |  |
| `convert` | `string` | No | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `reverse` | `bool` | No | Reverse the sequence (no complement). |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | `bool` | No | Remove digits, spaces and gaps (keep letters only). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `width` | `int` | No | Line-wrap width; 0 = single line. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->FormatSequence()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): FormatSequenceEntity`

Create a new `FormatSequenceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## FunctionalEnrichmentEntity

```php
$functional_enrichment = $client->FunctionalEnrichment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `background` | `array` | No | Custom background/universe gene symbols. |
| `collections` | `array` | No | Which term collections to test. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `array` | Yes | Query gene symbols (human, e.g. |
| `maxTermSize` | `int` | No | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | `int` | No | Skip terms/pathways with fewer than this many background genes. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->FunctionalEnrichment()->create([
  "genes" => null, // array
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): FunctionalEnrichmentEntity`

Create a new `FunctionalEnrichmentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GcContentEntity

```php
$gc_content = $client->GcContent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->GcContent()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GcContentEntity`

Create a new `GcContentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GeneDossierEntity

```php
$gene_dossier = $client->GeneDossier();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->GeneDossier()->create([
  "gene" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GeneDossierEntity`

Create a new `GeneDossierEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GeneExpressionEntity

```php
$gene_expression = $client->GeneExpression();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->GeneExpression()->create([
  "gene" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GeneExpressionEntity`

Create a new `GeneExpressionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GeneModelEntity

```php
$gene_model = $client->GeneModel();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->GeneModel()->create([
  "gene" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GeneModelEntity`

Create a new `GeneModelEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GoldenGateFidelityEntity

```php
$golden_gate_fidelity = $client->GoldenGateFidelity();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `compareToNamedSet` | `string` | No | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | `string` | No | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `overhangs` | `array` | Yes | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `riskThreshold` | `float` | No | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->GoldenGateFidelity()->create([
  "ok" => null, // mixed
  "overhangs" => null, // array
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GoldenGateFidelityEntity`

Create a new `GoldenGateFidelityEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## HgvsConvertEntity

```php
$hgvs_convert = $client->HgvsConvert();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `variant` | `string` | Yes | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->HgvsConvert()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
  "variant" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): HgvsConvertEntity`

Create a new `HgvsConvertEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## IdMapPollEntity

```php
$id_map_poll = $client->IdMapPoll();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->IdMapPoll()->create([
  "jobId" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): IdMapPollEntity`

Create a new `IdMapPollEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## IdMapSubmitEntity

```php
$id_map_submit = $client->IdMapSubmit();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | Yes | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | `array` | Yes | The ids to map, up to 1000 (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `taxId` | `string` | No | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | `string` | Yes | Target id type. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->IdMapSubmit()->create([
  "from" => null, // string
  "ids" => null, // array
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "to" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): IdMapSubmitEntity`

Create a new `IdMapSubmitEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## InSilicoPcrEntity

```php
$in_silico_pcr = $client->InSilicoPcr();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No | Treat the template as circular (plasmid). |
| `forwardPrimer` | `string` | Yes | Primer 1, 5'→3'. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Mismatches tolerated per primer. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `reversePrimer` | `string` | Yes | Primer 2, 5'→3' (order does not matter). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->InSilicoPcr()->create([
  "forwardPrimer" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "reversePrimer" => null, // string
  "template" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): InSilicoPcrEntity`

Create a new `InSilicoPcrEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## KaspPrimerDesignEntity

```php
$kasp_primer_design = $client->KaspPrimerDesign();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addSecondaryMismatch` | `bool` | No | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | `string` | Yes | First allele (single base) — gets the FAM tail. |
| `alleleB` | `string` | Yes | Second allele (single base) — gets the HEX tail. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | `int` | No | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | `int` | No | Minimum amplicon length for the common reverse primer. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `snpPosition` | `int` | Yes | 1-based position of the SNP on the forward strand. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | `float` | No | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->KaspPrimerDesign()->create([
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

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): KaspPrimerDesignEntity`

Create a new `KaspPrimerDesignEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ListToolEntity

```php
$list_tool = $client->ListTool();
```

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->ListTool()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ListToolEntity`

Create a new `ListToolEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MeltingTemperatureEntity

```php
$melting_temperature = $client->MeltingTemperature();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntpMM` | `float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `mixed` | Yes |  |
| `oligoNM` | `float` | No | Total strand concentration (nM). |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | `float` | No | Optional target Tm (°C). |
| `tmTolerance` | `float` | No | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->MeltingTemperature()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MeltingTemperatureEntity`

Create a new `MeltingTemperatureEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MotifFinderEntity

```php
$motif_finder = $client->MotifFinder();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Maximum allowed mismatches per match. |
| `motif` | `string` | Yes | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `bool` | No | Also search the reverse strand. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->MotifFinder()->create([
  "motif" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MotifFinderEntity`

Create a new `MotifFinderEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MultipleSequenceAlignmentEntity

```php
$multiple_sequence_alignment = $client->MultipleSequenceAlignment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->MultipleSequenceAlignment()->create([
  "input" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MultipleSequenceAlignmentEntity`

Create a new `MultipleSequenceAlignmentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## OligoAnalysiEntity

```php
$oligo_analysi = $client->OligoAnalysi();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntpMM` | `float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `mixed` | Yes |  |
| `oligoNM` | `float` | No | Total strand concentration (nM). |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->OligoAnalysi()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): OligoAnalysiEntity`

Create a new `OligoAnalysiEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## OrthologMapEntity

```php
$ortholog_map = $client->OrthologMap();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sourceSpecies` | `string` | No | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | `array` | Yes | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | `string` | Yes | Ensembl species slug to find homologs in (e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `type` | `string` | No | Homology type to return. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->OrthologMap()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "symbols" => null, // array
  "targetSpecies" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): OrthologMapEntity`

Create a new `OrthologMapEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PairwiseAlignmentEntity

```php
$pairwise_alignment = $client->PairwiseAlignment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gap` | `float` | No | Linear gap penalty (per gap position). |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | `float` | No | Match score. |
| `mismatch` | `float` | No | Mismatch penalty. |
| `mode` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `seqA` | `string` | Yes | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | `string` | Yes | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PairwiseAlignment()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "seqA" => null, // string
  "seqB" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PairwiseAlignmentEntity`

Create a new `PairwiseAlignmentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ParseGenbankEntity

```php
$parse_genbank = $client->ParseGenbank();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `text` | `string` | Yes | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ParseGenbank()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "text" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ParseGenbankEntity`

Create a new `ParseGenbankEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ParseSangerTraceEntity

```php
$parse_sanger_trace = $client->ParseSangerTrace();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fileBase64` | `string` | Yes | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | No | Optional original file name (echoed back). |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ParseSangerTrace()->create([
  "fileBase64" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ParseSangerTraceEntity`

Create a new `ParseSangerTraceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PlasmidAnnotateEntity

```php
$plasmid_annotate = $client->PlasmidAnnotate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PlasmidAnnotate()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PlasmidAnnotateEntity`

Create a new `PlasmidAnnotateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PlasmidDeepAnnotateEntity

```php
$plasmid_deep_annotate = $client->PlasmidDeepAnnotate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No | Treat the sequence as a circular plasmid (vs. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PlasmidDeepAnnotate()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PlasmidDeepAnnotateEntity`

Create a new `PlasmidDeepAnnotateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PlasmidFullReportEntity

```php
$plasmid_full_report = $client->PlasmidFullReport();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `topN` | `int` | No | How many top-ranked backbone candidates to report. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PlasmidFullReport()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PlasmidFullReportEntity`

Create a new `PlasmidFullReportEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PlasmidIdentifyEntity

```php
$plasmid_identify = $client->PlasmidIdentify();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `topN` | `int` | No | How many top-ranked backbone candidates to report. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PlasmidIdentify()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PlasmidIdentifyEntity`

Create a new `PlasmidIdentifyEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PrimeEditingDesignEntity

```php
$prime_editing_design = $client->PrimeEditingDesign();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `editEnd` | `int` | Yes | 1-based inclusive end of the region being changed. |
| `editStart` | `int` | Yes | 1-based inclusive start of the region being changed. |
| `frameStart` | `int` | No | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | `string` | No | Replacement bases (forward strand). |
| `ok` | `mixed` | Yes |  |
| `pbsLength` | `int` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `rttHomology` | `int` | No | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | `string` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PrimeEditingDesign()->create([
  "editEnd" => null, // int
  "editStart" => null, // int
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "target" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PrimeEditingDesignEntity`

Create a new `PrimeEditingDesignEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PrimeEditingTwinDesignEntity

```php
$prime_editing_twin_design = $client->PrimeEditingTwinDesign();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | `string` | Yes | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` | `mixed` | Yes |  |
| `overlapLength` | `int` | No | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | `int` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `array` | Yes |  |
| `replaceEnd` | `int` | Yes | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | `int` | Yes | 1-based inclusive start of the region being replaced/deleted. |
| `result` | `array` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PrimeEditingTwinDesign()->create([
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

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PrimeEditingTwinDesignEntity`

Create a new `PrimeEditingTwinDesignEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PrimerDesignEntity

```php
$primer_design = $client->PrimerDesign();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ampliconMax` | `int` | No |  |
| `ampliconMin` | `int` | No |  |
| `dntpMM` | `float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` | `float` | No |  |
| `gcMin` | `float` | No |  |
| `lenMax` | `int` | No |  |
| `lenMin` | `int` | No |  |
| `lenOpt` | `int` | No |  |
| `maxReturn` | `int` | No | Number of best pairs to return. |
| `mgMM` | `float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `mixed` | Yes |  |
| `oligoNM` | `float` | No | Total strand concentration (nM). |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `targetEnd` | `int` | No | 1-based inclusive end of the target region (optional). |
| `targetStart` | `int` | No | 1-based inclusive start of a region the product must span (optional). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` | `float` | No |  |
| `tmMaxDiff` | `float` | No | Max Tm difference within a pair (°C). |
| `tmMin` | `float` | No |  |
| `tmOpt` | `float` | No |  |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PrimerDesign()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "template" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PrimerDesignEntity`

Create a new `PrimerDesignEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PrimerSpecificityEntity

```php
$primer_specificity = $client->PrimerSpecificity();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `forwardPrimer` | `string` | Yes | Forward primer, 5'→3'. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | `int` | No | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `reversePrimer` | `string` | Yes | Reverse primer, 5'→3'. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PrimerSpecificity()->create([
  "forwardPrimer" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "reversePrimer" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PrimerSpecificityEntity`

Create a new `PrimerSpecificityEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ProteaseDigestionEntity

```php
$protease_digestion = $client->ProteaseDigestion();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | `float` | No | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | `int` | No | Cap on the number of returned peptides. |
| `minMass` | `float` | No | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | `int` | No | Allowed missed internal cleavages (0–2). |
| `ok` | `mixed` | Yes |  |
| `protease` | `string` | No | Protease or chemical cleavage agent. |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ProteaseDigestion()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ProteaseDigestionEntity`

Create a new `ProteaseDigestionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ProteinAnnotatePollEntity

```php
$protein_annotate_poll = $client->ProteinAnnotatePoll();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ProteinAnnotatePoll()->create([
  "jobId" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ProteinAnnotatePollEntity`

Create a new `ProteinAnnotatePollEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ProteinAnnotateSubmitEntity

```php
$protein_annotate_submit = $client->ProteinAnnotateSubmit();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `appl` | `string` | No | Restrict to one member database (e.g. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | `bool` | No | Include GO-term cross-references. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ProteinAnnotateSubmit()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ProteinAnnotateSubmitEntity`

Create a new `ProteinAnnotateSubmitEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ProteinHydrophobicityEntity

```php
$protein_hydrophobicity = $client->ProteinHydrophobicity();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `scale` | `string` | No | Amino-acid scale. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `window` | `int` | No | Sliding-window size (clamped to an odd number ≥ 1). |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ProteinHydrophobicity()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ProteinHydrophobicityEntity`

Create a new `ProteinHydrophobicityEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ProteinPropertyEntity

```php
$protein_property = $client->ProteinProperty();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `chargeStep` | `float` | No | pH step for the net-charge titration curve (0–14). |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ProteinProperty()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ProteinPropertyEntity`

Create a new `ProteinPropertyEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RandomSequenceEntity

```php
$random_sequence = $client->RandomSequence();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | `float` | No | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` | `string` | No |  |
| `length` | `int` | Yes | Number of residues to generate. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->RandomSequence()->create([
  "length" => null, // int
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RandomSequenceEntity`

Create a new `RandomSequenceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RestrictionSiteEntity

```php
$restriction_site = $client->RestrictionSite();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzymes` | `array` | No | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->RestrictionSite()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RestrictionSiteEntity`

Create a new `RestrictionSiteEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ReverseComplementEntity

```php
$reverse_complement = $client->ReverseComplement();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `type` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ReverseComplement()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ReverseComplementEntity`

Create a new `ReverseComplementEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ReverseTranslateEntity

```php
$reverse_translate = $client->ReverseTranslate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `organism` | `string` | No | Codon-usage host (ignored in degenerate mode). |
| `protein` | `string` | Yes | Protein sequence (one-letter codes; * for stop). |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ReverseTranslate()->create([
  "ok" => null, // mixed
  "protein" => null, // string
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ReverseTranslateEntity`

Create a new `ReverseTranslateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RnaFoldEntity

```php
$rna_fold = $client->RnaFold();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->RnaFold()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RnaFoldEntity`

Create a new `RnaFoldEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SangerVsReferenceEntity

```php
$sanger_vs_reference = $client->SangerVsReference();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fileBase64` | `string` | No | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | No | Optional original file name (echoed back). |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | `float` | No | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `read` | `string` | No | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | `string` | Yes | Expected reference sequence (FASTA or raw). |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SangerVsReference()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "reference" => null, // string
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SangerVsReferenceEntity`

Create a new `SangerVsReferenceEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SavePermalinkEntity

```php
$save_permalink = $client->SavePermalink();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `array` | Yes | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SavePermalink()->create([
  "args" => null, // array
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SavePermalinkEntity`

Create a new `SavePermalinkEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SeqfileStatEntity

```php
$seqfile_stat = $client->SeqfileStat();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `qualityOffset` | `int` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SeqfileStat()->create([
  "input" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SeqfileStatEntity`

Create a new `SeqfileStatEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SequenceFetchEntity

```php
$sequence_fetch = $client->SequenceFetch();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `string` | Yes | GenBank/RefSeq accession (e.g. |
| `db` | `string` | No | Database to query; auto-detects from the accession format. |
| `format` | `string` | No | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SequenceFetch()->create([
  "accession" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SequenceFetchEntity`

Create a new `SequenceFetchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SequenceFormatConvertEntity

```php
$sequence_format_convert = $client->SequenceFormatConvert();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | No | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | A FASTA or GenBank record to convert. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `to` | `string` | No | Output format. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SequenceFormatConvert()->create([
  "input" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SequenceFormatConvertEntity`

Create a new `SequenceFormatConvertEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SequenceReportEntity

```php
$sequence_report = $client->SequenceReport();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `endPrimerLength` | `int` | No | Length of the naive end primers taken from each end. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `int` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `int` | No | Minimum ORF length in amino acids. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SequenceReport()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SequenceReportEntity`

Create a new `SequenceReportEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SequenceSearchEntity

```php
$sequence_search = $client->SequenceSearch();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `db` | `string` | No |  |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | No | Gene symbol/name, e.g. |
| `maxResults` | `int` | No | Up to 20. |
| `ok` | `mixed` | Yes |  |
| `organism` | `string` | No | Organism name, e.g. |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `term` | `string` | No | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SequenceSearch()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SequenceSearchEntity`

Create a new `SequenceSearchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SequencingReadbackVerifyEntity

```php
$sequencing_readback_verify = $client->SequencingReadbackVerify();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | `int` | No | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `reads` | `string` | Yes | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | `string` | Yes | The claimed/expected reference sequence. |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SequencingReadbackVerify()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "reads" => null, // string
  "reference" => null, // string
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SequencingReadbackVerifyEntity`

Create a new `SequencingReadbackVerifyEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SessionCreateEntity

```php
$session_create = $client->SessionCreate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entries` | `array` | No | Initial named entries, e.g. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SessionCreate()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SessionCreateEntity`

Create a new `SessionCreateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SessionGetEntity

```php
$session_get = $client->SessionGet();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | `array` | No | Only return these entries; omit to return all of them. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SessionGet()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sessionId" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SessionGetEntity`

Create a new `SessionGetEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SessionRunEntity

```php
$session_run = $client->SessionRun();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `array` | No | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | `array` | No | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |
| `writeBack` | `array` | No | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SessionRun()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sessionId" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SessionRunEntity`

Create a new `SessionRunEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SessionSetEntity

```php
$session_set = $client->SessionSet();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entries` | `array` | Yes | Named entries to add/overwrite, e.g. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SessionSet()->create([
  "entries" => null, // array
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sessionId" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SessionSetEntity`

Create a new `SessionSetEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SirnaDesignEntity

```php
$sirna_design = $client->SirnaDesign();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | `int` | No | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `shRnaLoop` | `string` | No | Loop sequence used when assembling the shRNA cassette. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SirnaDesign()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "target" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SirnaDesignEntity`

Create a new `SirnaDesignEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SiteDirectedMutagenesiEntity

```php
$site_directed_mutagenesi = $client->SiteDirectedMutagenesi();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `float` | No | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | `float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | `string` | No | Edit at the nucleotide or amino-acid level. |
| `frameStart` | `int` | No | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | `string` | No | Replacement base (editKind='nt'). |
| `ok` | `mixed` | Yes |  |
| `oligoNM` | `float` | No | Total strand concentration (nM). |
| `organism` | `string` | No | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | `int` | No | 1-based position to substitute (editKind='nt'). |
| `provenance` | `array` | Yes |  |
| `residue` | `int` | No | 1-based residue number to change (editKind='aa'). |
| `result` | `array` | Yes | Tool-specific output object. |
| `style` | `string` | No | Mutagenic primer style. |
| `targetAa` | `string` | No | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SiteDirectedMutagenesi()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "template" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SiteDirectedMutagenesiEntity`

Create a new `SiteDirectedMutagenesiEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TranslateEntity

```php
$translate = $client->Translate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frame` | `int` | No |  |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | `bool` | No | Stop at the first stop codon. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Translate()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TranslateEntity`

Create a new `TranslateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## VariantAnnotateEntity

```php
$variant_annotate = $client->VariantAnnotate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `assembly` | `string` | No | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `variant` | `string` | Yes | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->VariantAnnotate()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
  "variant" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): VariantAnnotateEntity`

Create a new `VariantAnnotateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## VariantComparatorEntity

```php
$variant_comparator = $client->VariantComparator();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `coding` | `bool` | No | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | `int` | No | 1-based reading-frame start (used when coding is true). |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `query` | `string` | Yes | Query / variant sequence (raw or FASTA). |
| `reference` | `string` | Yes | Reference / wild-type sequence (raw or FASTA). |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->VariantComparator()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "query" => null, // string
  "reference" => null, // string
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): VariantComparatorEntity`

Create a new `VariantComparatorEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## VerifyAssemblyEntity

```php
$verify_assembly = $client->VerifyAssembly();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `float` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `bool` | No | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | `string` | Yes | The sequence you claim you ended up with. |
| `coding` | `bool` | No | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | `string` | No | Type IIS enzyme for Golden Gate. |
| `enzyme3` | `string` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | No | 5′ enzyme (restriction method). |
| `fragmentPcrs` | `array` | No | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | `array` | No | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | `int` | No | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | No | Insert sequence (restriction method). |
| `insertPcr` | `array` | No | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | `string` | Yes | Assembly method used. |
| `names` | `array` | No | Optional labels for each fragment. |
| `ok` | `mixed` | Yes |  |
| `overlapLen` | `int` | No | Gibson homology-arm length (bp). |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `vector` | `string` | No | Vector sequence (restriction method). |
| `vectorPcr` | `array` | No | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->VerifyAssembly()->create([
  "claimedConstruct" => null, // string
  "method" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): VerifyAssemblyEntity`

Create a new `VerifyAssemblyEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## VerifyConstructEntity

```php
$verify_construct = $client->VerifyConstruct();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `claimedConstruct` | `string` | Yes | The final sequence claimed to have been built. |
| `expectedFrameStart` | `int` | No | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | `string` | Yes | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | `string` | Yes | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | `string` | Yes | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | `int` | No | Mismatches tolerated per primer during PCR prediction. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `templateCircular` | `bool` | No | Treat insertTemplate as circular (e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->VerifyConstruct()->create([
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

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): VerifyConstructEntity`

Create a new `VerifyConstructEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## VirtualGelEntity

```php
$virtual_gel = $client->VirtualGel();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No | Treat the sequence as circular (plasmid). |
| `enzymes` | `array` | No | Enzyme names to digest with. |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | `string` | No | DNA ladder to plot alongside the sample lane. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->VirtualGel()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence" => null, // string
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): VirtualGelEntity`

Create a new `VirtualGelEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## VolcanoPlotDataEntity

```php
$volcano_plot_data = $client->VolcanoPlotData();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes | Tool-specific output object. |
| `rows` | `array` | Yes | Differential expression rows, one per gene. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->VolcanoPlotData()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "rows" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): VolcanoPlotDataEntity`

Create a new `VolcanoPlotDataEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## WebSearchEntity

```php
$web_search = $client->WebSearch();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `mixed` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | `float` | No | Maximum number of results to return (default 5, max 10). |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `query` | `string` | Yes | The search query. |
| `result` | `array` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->WebSearch()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "query" => null, // string
  "result" => null, // array
  "tool" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): WebSearchEntity`

Create a new `WebSearchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `ratelimit` | 0.0.1 | Client-side rate limiting via a token bucket |
| `retry` | 0.0.1 | Automatic retry of transient failures with exponential backoff |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |
| `timeout` | 0.0.1 | Per-request timeout with transport abort |


Features are activated via the `feature` option:

```php
$client = new SeqbenchMcpSDK([
  "feature" => [
    "ratelimit" => ["active" => true],
    "retry" => ["active" => true],
    "test" => ["active" => true],
    "timeout" => ["active" => true],
  ],
]);
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `ratelimit`

Client-side rate limiting via a token bucket.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Automatic retry of transient failures with exponential backoff.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Per-request timeout with transport abort.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

