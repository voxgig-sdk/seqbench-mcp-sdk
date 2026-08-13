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
| `frameStart` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `target` | `string` | Yes |  |
| `targetPosition` | `int` | No |  |
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
| `args` | `array` | No |  |
| `capped` | `bool` | Yes |  |
| `columns` | `array` | Yes |  |
| `count` | `int` | Yes |  |
| `errors` | `int` | Yes |  |
| `input` | `string` | Yes |  |
| `limit` | `int` | Yes |  |
| `provenance` | `array` | Yes |  |
| `rows` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `columns` | `array` | Yes |  |
| `count` | `int` | Yes |  |
| `errors` | `int` | Yes |  |
| `input` | `string` | Yes |  |
| `limit` | `int` | Yes |  |
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
| `endPrimerLength` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `maxOrfs` | `int` | No |  |
| `minOrfAa` | `int` | No |  |
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
| `armTmTarget` | `float` | No |  |
| `circular` | `bool` | No |  |
| `enzyme` | `string` | No |  |
| `enzyme3` | `string` | No |  |
| `enzyme5` | `string` | No |  |
| `fragments` | `array` | No |  |
| `gate` | `mixed` | No |  |
| `insert` | `string` | No |  |
| `method` | `string` | Yes |  |
| `names` | `array` | No |  |
| `ok` | `mixed` | Yes |  |
| `overlapLen` | `int` | No |  |
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
| `frameStart` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `array` | Yes |  |
| `rareThreshold` | `float` | No |  |
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
| `avoidEnzymes` | `array` | No |  |
| `crypticOrfMinAa` | `int` | No |  |
| `frameStart` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `gcHigh` | `float` | No |  |
| `gcLow` | `float` | No |  |
| `gcWindow` | `int` | No |  |
| `homopolymerMin` | `int` | No |  |
| `maxPasses` | `int` | No |  |
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
| `avoidEnzymes` | `array` | No |  |
| `crypticOrfMinAa` | `int` | No |  |
| `frameStart` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `gcHigh` | `float` | No |  |
| `gcLow` | `float` | No |  |
| `gcWindow` | `int` | No |  |
| `homopolymerMin` | `int` | No |  |
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
| `minScore` | `float` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `searchReverseStrand` | `bool` | No |  |
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
| `armLength` | `int` | No |  |
| `blockPam` | `bool` | No |  |
| `designGenotypingPrimers` | `bool` | No |  |
| `editEnd` | `int` | No |  |
| `editStart` | `int` | No |  |
| `frameStart` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `guideEnd` | `int` | No |  |
| `guideStart` | `int` | No |  |
| `guideStrand` | `string` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `replacement` | `string` | Yes |  |
| `result` | `array` | Yes |  |
| `targetSequence` | `string` | Yes |  |
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
| `gate` | `mixed` | No |  |
| `maxMismatches` | `int` | No |  |
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
| `sequenceA` | `string` | Yes |  |
| `sequenceB` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `length` | `int` | No |  |
| `massNg` | `float` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | No |  |
| `tool` | `string` | Yes |  |
| `type` | `string` | No |  |
| `volumeUl` | `float` | No |  |

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
| `enzymeA` | `string` | Yes |  |
| `enzymeB` | `string` | Yes |  |
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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `reactions` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `protocolName` | `string` | No |  |
| `provenance` | `array` | Yes |  |
| `reactions` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `reactions` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `clusterCols` | `bool` | No |  |
| `clusterRows` | `bool` | No |  |
| `distanceMetric` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `genes` | `array` | Yes |  |
| `linkage` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `samples` | `array` | Yes |  |
| `tool` | `string` | Yes |  |
| `values` | `array` | Yes |  |
| `zScoreRows` | `bool` | No |  |

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
| `gate` | `mixed` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `qualityOffset` | `int` | No |  |
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
| `minLength` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `qualityOffset` | `int` | No |  |
| `qualityThreshold` | `int` | No |  |
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
| `minAaLength` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `requireStop` | `bool` | No |  |
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
| `caseMode` | `string` | No |  |
| `convert` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `reverse` | `bool` | No |  |
| `sequence` | `string` | Yes |  |
| `stripNonLetters` | `bool` | No |  |
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
| `collections` | `array` | No |  |
| `gate` | `mixed` | No |  |
| `genes` | `array` | Yes |  |
| `maxTermSize` | `int` | No |  |
| `minTermSize` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `compareToNamedSet` | `string` | No |  |
| `dataset` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `overhangs` | `array` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `riskThreshold` | `float` | No |  |
| `tool` | `string` | Yes |  |

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
| `jobId` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `from` | `string` | Yes |  |
| `gate` | `mixed` | No |  |
| `ids` | `array` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `taxId` | `string` | No |  |
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
| `forwardPrimer` | `string` | Yes |  |
| `gate` | `mixed` | No |  |
| `maxMismatches` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `reversePrimer` | `string` | Yes |  |
| `template` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `addSecondaryMismatch` | `bool` | No |  |
| `alleleA` | `string` | Yes |  |
| `alleleB` | `string` | Yes |  |
| `gate` | `mixed` | No |  |
| `maxAmplicon` | `int` | No |  |
| `minAmplicon` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `snpPosition` | `int` | Yes |  |
| `target` | `string` | Yes |  |
| `targetCoreTm` | `float` | No |  |
| `tool` | `string` | Yes |  |

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
| `dntpMM` | `float` | No |  |
| `gate` | `mixed` | No |  |
| `mgMM` | `float` | No |  |
| `naMM` | `float` | No |  |
| `ok` | `mixed` | Yes |  |
| `oligoNM` | `float` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sequence` | `string` | Yes |  |
| `targetTm` | `float` | No |  |
| `tmTolerance` | `float` | No |  |
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
| `maxMismatches` | `int` | No |  |
| `motif` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `searchReverseStrand` | `bool` | No |  |
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
| `dntpMM` | `float` | No |  |
| `gate` | `mixed` | No |  |
| `mgMM` | `float` | No |  |
| `naMM` | `float` | No |  |
| `ok` | `mixed` | Yes |  |
| `oligoNM` | `float` | No |  |
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
| `sourceSpecies` | `string` | No |  |
| `symbols` | `array` | Yes |  |
| `targetSpecies` | `string` | Yes |  |
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
| `gap` | `float` | No |  |
| `gate` | `mixed` | No |  |
| `match` | `float` | No |  |
| `mismatch` | `float` | No |  |
| `mode` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `seqA` | `string` | Yes |  |
| `seqB` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `fileBase64` | `string` | Yes |  |
| `fileName` | `string` | No |  |
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
| `topN` | `int` | No |  |

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
| `topN` | `int` | No |  |

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
| `editEnd` | `int` | Yes |  |
| `editStart` | `int` | Yes |  |
| `frameStart` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `insertedSeq` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `pbsLength` | `int` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `rttHomology` | `int` | No |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `newSequence` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `overlapLength` | `int` | No |  |
| `pbsLength` | `int` | No |  |
| `provenance` | `array` | Yes |  |
| `replaceEnd` | `int` | Yes |  |
| `replaceStart` | `int` | Yes |  |
| `result` | `array` | Yes |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `dntpMM` | `float` | No |  |
| `gate` | `mixed` | No |  |
| `gcMax` | `float` | No |  |
| `gcMin` | `float` | No |  |
| `lenMax` | `int` | No |  |
| `lenMin` | `int` | No |  |
| `lenOpt` | `int` | No |  |
| `maxReturn` | `int` | No |  |
| `mgMM` | `float` | No |  |
| `naMM` | `float` | No |  |
| `ok` | `mixed` | Yes |  |
| `oligoNM` | `float` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `targetEnd` | `int` | No |  |
| `targetStart` | `int` | No |  |
| `template` | `string` | Yes |  |
| `tmMax` | `float` | No |  |
| `tmMaxDiff` | `float` | No |  |
| `tmMin` | `float` | No |  |
| `tmOpt` | `float` | No |  |
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
| `forwardPrimer` | `string` | Yes |  |
| `gate` | `mixed` | No |  |
| `maxMismatches` | `int` | No |  |
| `maxProductLength` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `reversePrimer` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `maxMass` | `float` | No |  |
| `maxPeptides` | `int` | No |  |
| `minMass` | `float` | No |  |
| `missedCleavages` | `int` | No |  |
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
| `jobId` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `appl` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `goterms` | `bool` | No |  |
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
| `chargeStep` | `float` | No |  |
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
| `gcContent` | `float` | No |  |
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
| `enzymes` | `array` | No |  |
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
| `fileBase64` | `string` | No |  |
| `fileName` | `string` | No |  |
| `gate` | `mixed` | No |  |
| `minCoverage` | `float` | No |  |
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
| `args` | `array` | Yes |  |
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
| `gate` | `mixed` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `qualityOffset` | `int` | No |  |
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
| `endPrimerLength` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `maxOrfs` | `int` | No |  |
| `minOrfAa` | `int` | No |  |
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
| `maxResults` | `int` | No |  |
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
| `minSupportingReads` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `reads` | `string` | Yes |  |
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
| `entries` | `array` | No |  |
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
| `names` | `array` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `args` | `array` | No |  |
| `fromSession` | `array` | No |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `writeBack` | `array` | No |  |

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
| `entries` | `array` | Yes |  |
| `gate` | `mixed` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `minReynolds` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `shRnaLoop` | `string` | No |  |
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
| `armTmTarget` | `float` | No |  |
| `dntpMM` | `float` | No |  |
| `editKind` | `string` | No |  |
| `frameStart` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `mgMM` | `float` | No |  |
| `naMM` | `float` | No |  |
| `newBase` | `string` | No |  |
| `ok` | `mixed` | Yes |  |
| `oligoNM` | `float` | No |  |
| `organism` | `string` | No |  |
| `position` | `int` | No |  |
| `provenance` | `array` | Yes |  |
| `residue` | `int` | No |  |
| `result` | `array` | Yes |  |
| `style` | `string` | No |  |
| `targetAa` | `string` | No |  |
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
| `toStop` | `bool` | No |  |
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
| `frameStart` | `int` | No |  |
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
| `armTmTarget` | `float` | No |  |
| `circular` | `bool` | No |  |
| `claimedConstruct` | `string` | Yes |  |
| `coding` | `bool` | No |  |
| `enzyme` | `string` | No |  |
| `enzyme3` | `string` | No |  |
| `enzyme5` | `string` | No |  |
| `fragmentPcrs` | `array` | No |  |
| `fragments` | `array` | No |  |
| `frameStart` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `insert` | `string` | No |  |
| `insertPcr` | `array` | No |  |
| `method` | `string` | Yes |  |
| `names` | `array` | No |  |
| `ok` | `mixed` | Yes |  |
| `overlapLen` | `int` | No |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `tool` | `string` | Yes |  |
| `vector` | `string` | No |  |
| `vectorPcr` | `array` | No |  |

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
| `claimedConstruct` | `string` | Yes |  |
| `expectedFrameStart` | `int` | No |  |
| `gate` | `mixed` | No |  |
| `insertForwardPrimer` | `string` | Yes |  |
| `insertReversePrimer` | `string` | Yes |  |
| `insertTemplate` | `string` | Yes |  |
| `maxPrimerMismatches` | `int` | No |  |
| `ok` | `mixed` | Yes |  |
| `provenance` | `array` | Yes |  |
| `result` | `array` | Yes |  |
| `templateCircular` | `bool` | No |  |
| `tool` | `string` | Yes |  |

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
| `circular` | `bool` | No |  |
| `enzymes` | `array` | No |  |
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
| `rows` | `array` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `mixed` | No |  |
| `max_results` | `float` | No |  |
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

