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
| `accession` | `string` | Yes |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `length` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `wing` | `int` | No |  |

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
| `editor` | `string` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `target` | `string` | Yes |  |
| `target_position` | `int` | No |  |
| `tool` | `string` | Yes |  |

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
| `arg` | `array` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Batch()->create([
  "input" => null, // string
  "ok" => null, // mixed
  "result" => null, // array
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
| `input` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `result` | `array` | Yes |  |
| `step` | `array` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->BatchWorkflow()->create([
  "input" => null, // string
  "ok" => null, // mixed
  "result" => null, // array
  "step" => null, // array
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
| `end_primer_length` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `max_orf` | `int` | No |  |
| `min_orf_aa` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_tm_target` | `float` | No |  |
| `circular` | `bool` | No |  |
| `enzyme` | `string` | No |  |
| `enzyme3` | `string` | No |  |
| `enzyme5` | `string` | No |  |
| `fragment` | `array` | No |  |
| `gate` | `mixed` | No |  |
| `insert` | `string` | No |  |
| `method` | `string` | Yes |  |
| `name` | `array` | No |  |
| `ok` | `mixed` | Yes |  |
| `overlap_len` | `int` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |
| `vector` | `string` | No |  |

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
| `frame_start` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `array` | Yes |  |
| `rare_threshold` | `float` | No |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `avoid_enzyme` | `array` | No |  |
| `cryptic_orf_min_aa` | `int` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `gc_high` | `float` | No |  |
| `gc_low` | `float` | No |  |
| `gc_window` | `int` | No |  |
| `homopolymer_min` | `int` | No |  |
| `max_pass` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `avoid_enzyme` | `array` | No |  |
| `cryptic_orf_min_aa` | `int` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `gc_high` | `float` | No |  |
| `gc_low` | `float` | No |  |
| `gc_window` | `int` | No |  |
| `homopolymer_min` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `min_score` | `float` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `search_reverse_strand` | `bool` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_length` | `int` | No |  |
| `block_pam` | `bool` | No |  |
| `design_genotyping_primer` | `bool` | No |  |
| `edit_end` | `int` | No |  |
| `edit_start` | `int` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `guide_end` | `int` | No |  |
| `guide_start` | `int` | No |  |
| `guide_strand` | `string` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `replacement` | `string` | Yes |  |
| `result` | `array` | Yes |  |
| `target_sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CrisprHdrDonor()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "replacement" => null, // string
  "result" => null, // array
  "target_sequence" => null, // string
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
| `gate` | `mixed` | No |  |
| `max_mismatch` | `int` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `protospacer` | `string` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence_a` | `string` | Yes |  |
| `sequence_b` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CrossDimer()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sequence_a" => null, // string
  "sequence_b" => null, // string
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
| `gate` | `mixed` | No |  |
| `length` | `int` | No |  |
| `mass_ng` | `float` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | No |  |
| `tool` | `string` | Yes |  |
| `type` | `string` | No |  |
| `volume_ul` | `float` | No |  |

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
| `enzyme_a` | `string` | Yes |  |
| `enzyme_b` | `string` | Yes |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->DoubleDigest()->create([
  "enzyme_a" => null, // string
  "enzyme_b" => null, // string
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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `reaction` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ExportEchoPicklist()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "reaction" => null, // array
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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `protocol_name` | `string` | No |  |
| `provenance` | `array` | Yes |  |
| `reaction` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ExportOpentronsProtocol()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "reaction" => null, // array
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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `reaction` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ExportPlateLayout()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "reaction" => null, // array
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
| `cluster_col` | `bool` | No |  |
| `cluster_row` | `bool` | No |  |
| `distance_metric` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `gene` | `array` | Yes |  |
| `linkage` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sample` | `array` | Yes |  |
| `tool` | `string` | Yes |  |
| `value` | `array` | Yes |  |
| `z_score_row` | `bool` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ExpressionHeatmapCluster()->create([
  "gene" => null, // array
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "sample" => null, // array
  "tool" => null, // string
  "value" => null, // array
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
| `gate` | `mixed` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `quality_offset` | `int` | No |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `input` | `string` | Yes |  |
| `min_length` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `quality_offset` | `int` | No |  |
| `quality_threshold` | `int` | No |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `min_aa_length` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `require_stop` | `bool` | No |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `case_mode` | `string` | No |  |
| `convert` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `reverse` | `bool` | No |  |
| `sequence` | `string` | Yes |  |
| `strip_non_letter` | `bool` | No |  |
| `tool` | `string` | Yes |  |
| `width` | `int` | No |  |

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
| `background` | `array` | No |  |
| `collection` | `array` | No |  |
| `gate` | `mixed` | No |  |
| `gene` | `array` | Yes |  |
| `max_term_size` | `int` | No |  |
| `min_term_size` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->FunctionalEnrichment()->create([
  "gene" => null, // array
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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `compare_to_named_set` | `string` | No |  |
| `dataset` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `overhang` | `array` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `risk_threshold` | `float` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->GoldenGateFidelity()->create([
  "ok" => null, // mixed
  "overhang" => null, // array
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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |
| `variant` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `job_id` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->IdMapPoll()->create([
  "job_id" => null, // string
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
| `from` | `string` | Yes |  |
| `gate` | `mixed` | No |  |
| `ids` | `array` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tax_id` | `string` | No |  |
| `to` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `circular` | `bool` | No |  |
| `forward_primer` | `string` | Yes |  |
| `gate` | `mixed` | No |  |
| `max_mismatch` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `reverse_primer` | `string` | Yes |  |
| `template` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->InSilicoPcr()->create([
  "forward_primer" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "reverse_primer" => null, // string
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
| `add_secondary_mismatch` | `bool` | No |  |
| `allele_a` | `string` | Yes |  |
| `allele_b` | `string` | Yes |  |
| `gate` | `mixed` | No |  |
| `max_amplicon` | `int` | No |  |
| `min_amplicon` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `snp_position` | `int` | Yes |  |
| `target` | `string` | Yes |  |
| `target_core_tm` | `float` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->KaspPrimerDesign()->create([
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
| `dntp_mm` | `float` | No |  |
| `gate` | `mixed` | No |  |
| `mg_mm` | `float` | No |  |
| `na_mm` | `float` | No |  |
| `ok` | `mixed` | Yes |  |
| `oligo_nm` | `float` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `target_tm` | `float` | No |  |
| `tm_tolerance` | `float` | No |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `max_mismatch` | `int` | No |  |
| `motif` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `search_reverse_strand` | `bool` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `dntp_mm` | `float` | No |  |
| `gate` | `mixed` | No |  |
| `mg_mm` | `float` | No |  |
| `na_mm` | `float` | No |  |
| `ok` | `mixed` | Yes |  |
| `oligo_nm` | `float` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `source_species` | `string` | No |  |
| `symbol` | `array` | Yes |  |
| `target_species` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `type` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->OrthologMap()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "symbol" => null, // array
  "target_species" => null, // string
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
| `gap` | `float` | No |  |
| `gate` | `mixed` | No |  |
| `match` | `float` | No |  |
| `mismatch` | `float` | No |  |
| `mode` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `seq_a` | `string` | Yes |  |
| `seq_b` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PairwiseAlignment()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "seq_a" => null, // string
  "seq_b" => null, // string
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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `text` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `file_base64` | `string` | Yes |  |
| `file_name` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ParseSangerTrace()->create([
  "file_base64" => null, // string
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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `circular` | `bool` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `circular` | `bool` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `top_n` | `int` | No |  |

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
| `circular` | `bool` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `top_n` | `int` | No |  |

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
| `edit_end` | `int` | Yes |  |
| `edit_start` | `int` | Yes |  |
| `frame_start` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `inserted_seq` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `pbs_length` | `int` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `rtt_homology` | `int` | No |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PrimeEditingDesign()->create([
  "edit_end" => null, // int
  "edit_start" => null, // int
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
| `gate` | `mixed` | No |  |
| `new_sequence` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `overlap_length` | `int` | No |  |
| `pbs_length` | `int` | No |  |
| `provenance` | `array` | Yes |  |
| `replace_end` | `int` | Yes |  |
| `replace_start` | `int` | Yes |  |
| `result` | `array` | Yes |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PrimeEditingTwinDesign()->create([
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
| `amplicon_max` | `int` | No |  |
| `amplicon_min` | `int` | No |  |
| `dntp_mm` | `float` | No |  |
| `gate` | `mixed` | No |  |
| `gc_max` | `float` | No |  |
| `gc_min` | `float` | No |  |
| `len_max` | `int` | No |  |
| `len_min` | `int` | No |  |
| `len_opt` | `int` | No |  |
| `max_return` | `int` | No |  |
| `mg_mm` | `float` | No |  |
| `na_mm` | `float` | No |  |
| `ok` | `mixed` | Yes |  |
| `oligo_nm` | `float` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `target_end` | `int` | No |  |
| `target_start` | `int` | No |  |
| `template` | `string` | Yes |  |
| `tm_max` | `float` | No |  |
| `tm_max_diff` | `float` | No |  |
| `tm_min` | `float` | No |  |
| `tm_opt` | `float` | No |  |
| `tool` | `string` | Yes |  |

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
| `forward_primer` | `string` | Yes |  |
| `gate` | `mixed` | No |  |
| `max_mismatch` | `int` | No |  |
| `max_product_length` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `reverse_primer` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PrimerSpecificity()->create([
  "forward_primer" => null, // string
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "reverse_primer" => null, // string
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
| `gate` | `mixed` | No |  |
| `max_mass` | `float` | No |  |
| `max_peptide` | `int` | No |  |
| `min_mass` | `float` | No |  |
| `missed_cleavage` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `protease` | `string` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `job_id` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ProteinAnnotatePoll()->create([
  "job_id" => null, // string
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
| `appl` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `goterm` | `bool` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `scale` | `string` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `window` | `int` | No |  |

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
| `charge_step` | `float` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `gc_content` | `float` | No |  |
| `kind` | `string` | No |  |
| `length` | `int` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `enzyme` | `array` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
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
| `gate` | `mixed` | No |  |
| `mode` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `file_base64` | `string` | No |  |
| `file_name` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `min_coverage` | `float` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `read` | `string` | No |  |
| `reference` | `string` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arg` | `array` | Yes |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SavePermalink()->create([
  "arg" => null, // array
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
| `gate` | `mixed` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `quality_offset` | `int` | No |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `accession` | `string` | Yes |  |
| `db` | `string` | No |  |
| `format` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `from` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `to` | `string` | No |  |
| `tool` | `string` | Yes |  |

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
| `end_primer_length` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `max_orf` | `int` | No |  |
| `min_orf_aa` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `gene` | `string` | No |  |
| `max_result` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `term` | `string` | No |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `min_supporting_read` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `read` | `string` | Yes |  |
| `reference` | `string` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SequencingReadbackVerify()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "read" => null, // string
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
| `entry` | `array` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `name` | `array` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SessionGet()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "session_id" => null, // string
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
| `arg` | `array` | No |  |
| `from_session` | `array` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `write_back` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SessionRun()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "session_id" => null, // string
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
| `entry` | `array` | Yes |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SessionSet()->create([
  "entry" => null, // array
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "session_id" => null, // string
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
| `gate` | `mixed` | No |  |
| `min_reynold` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sh_rna_loop` | `string` | No |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_tm_target` | `float` | No |  |
| `dntp_mm` | `float` | No |  |
| `edit_kind` | `string` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `mg_mm` | `float` | No |  |
| `na_mm` | `float` | No |  |
| `new_base` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `oligo_nm` | `float` | No |  |
| `organism` | `string` | No |  |
| `position` | `int` | No |  |
| `provenance` | `array` | Yes |  |
| `residue` | `int` | No |  |
| `result` | `array` | Yes |  |
| `style` | `string` | No |  |
| `target_aa` | `string` | No |  |
| `template` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `to_stop` | `bool` | No |  |
| `tool` | `string` | Yes |  |

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
| `assembly` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |
| `variant` | `string` | Yes |  |

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
| `coding` | `bool` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `query` | `string` | Yes |  |
| `reference` | `string` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_tm_target` | `float` | No |  |
| `circular` | `bool` | No |  |
| `claimed_construct` | `string` | Yes |  |
| `coding` | `bool` | No |  |
| `enzyme` | `string` | No |  |
| `enzyme3` | `string` | No |  |
| `enzyme5` | `string` | No |  |
| `fragment` | `array` | No |  |
| `fragment_pcr` | `array` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `insert` | `string` | No |  |
| `insert_pcr` | `array` | No |  |
| `method` | `string` | Yes |  |
| `name` | `array` | No |  |
| `ok` | `mixed` | Yes |  |
| `overlap_len` | `int` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |
| `vector` | `string` | No |  |
| `vector_pcr` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->VerifyAssembly()->create([
  "claimed_construct" => null, // string
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
| `claimed_construct` | `string` | Yes |  |
| `expected_frame_start` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `insert_forward_primer` | `string` | Yes |  |
| `insert_reverse_primer` | `string` | Yes |  |
| `insert_template` | `string` | Yes |  |
| `max_primer_mismatch` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `template_circular` | `bool` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->VerifyConstruct()->create([
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
| `circular` | `bool` | No |  |
| `enzyme` | `array` | No |  |
| `gate` | `mixed` | No |  |
| `ladder` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `row` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->VolcanoPlotData()->create([
  "ok" => null, // mixed
  "provenance" => null, // array
  "result" => null, // array
  "row" => null, // array
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
| `gate` | `mixed` | No |  |
| `max_result` | `float` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `query` | `string` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new SeqbenchMcpSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

