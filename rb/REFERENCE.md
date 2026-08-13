# SeqbenchMcp Ruby SDK Reference

Complete API reference for the SeqbenchMcp Ruby SDK.


## SeqbenchMcpSDK

### Constructor

```ruby
require_relative 'SeqbenchMcp_sdk'

client = SeqbenchMcpSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `SeqbenchMcpSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = SeqbenchMcpSDK.test
```


### Instance Methods

#### `AlphafoldLookup(data = nil)`

Create a new `AlphafoldLookup` entity instance. Pass `nil` for no initial data.

#### `AsoDesign(data = nil)`

Create a new `AsoDesign` entity instance. Pass `nil` for no initial data.

#### `BaseEditingDesign(data = nil)`

Create a new `BaseEditingDesign` entity instance. Pass `nil` for no initial data.

#### `Batch(data = nil)`

Create a new `Batch` entity instance. Pass `nil` for no initial data.

#### `BatchWorkflow(data = nil)`

Create a new `BatchWorkflow` entity instance. Pass `nil` for no initial data.

#### `CharacterizeSequence(data = nil)`

Create a new `CharacterizeSequence` entity instance. Pass `nil` for no initial data.

#### `CloningSimulate(data = nil)`

Create a new `CloningSimulate` entity instance. Pass `nil` for no initial data.

#### `CodonAdaptationIndex(data = nil)`

Create a new `CodonAdaptationIndex` entity instance. Pass `nil` for no initial data.

#### `CodonOptimize(data = nil)`

Create a new `CodonOptimize` entity instance. Pass `nil` for no initial data.

#### `ConstructAutofix(data = nil)`

Create a new `ConstructAutofix` entity instance. Pass `nil` for no initial data.

#### `ConstructQc(data = nil)`

Create a new `ConstructQc` entity instance. Pass `nil` for no initial data.

#### `CrisprGrnaDesign(data = nil)`

Create a new `CrisprGrnaDesign` entity instance. Pass `nil` for no initial data.

#### `CrisprHdrDonor(data = nil)`

Create a new `CrisprHdrDonor` entity instance. Pass `nil` for no initial data.

#### `CrisprOfftargetCheck(data = nil)`

Create a new `CrisprOfftargetCheck` entity instance. Pass `nil` for no initial data.

#### `CrossDimer(data = nil)`

Create a new `CrossDimer` entity instance. Pass `nil` for no initial data.

#### `DnaMolarity(data = nil)`

Create a new `DnaMolarity` entity instance. Pass `nil` for no initial data.

#### `DoubleDigest(data = nil)`

Create a new `DoubleDigest` entity instance. Pass `nil` for no initial data.

#### `ExportEchoPicklist(data = nil)`

Create a new `ExportEchoPicklist` entity instance. Pass `nil` for no initial data.

#### `ExportOpentronsProtocol(data = nil)`

Create a new `ExportOpentronsProtocol` entity instance. Pass `nil` for no initial data.

#### `ExportPlateLayout(data = nil)`

Create a new `ExportPlateLayout` entity instance. Pass `nil` for no initial data.

#### `ExpressionHeatmapCluster(data = nil)`

Create a new `ExpressionHeatmapCluster` entity instance. Pass `nil` for no initial data.

#### `FastqQcReport(data = nil)`

Create a new `FastqQcReport` entity instance. Pass `nil` for no initial data.

#### `FastqTrim(data = nil)`

Create a new `FastqTrim` entity instance. Pass `nil` for no initial data.

#### `FindOrf(data = nil)`

Create a new `FindOrf` entity instance. Pass `nil` for no initial data.

#### `FormatSequence(data = nil)`

Create a new `FormatSequence` entity instance. Pass `nil` for no initial data.

#### `FunctionalEnrichment(data = nil)`

Create a new `FunctionalEnrichment` entity instance. Pass `nil` for no initial data.

#### `GcContent(data = nil)`

Create a new `GcContent` entity instance. Pass `nil` for no initial data.

#### `GeneDossier(data = nil)`

Create a new `GeneDossier` entity instance. Pass `nil` for no initial data.

#### `GeneExpression(data = nil)`

Create a new `GeneExpression` entity instance. Pass `nil` for no initial data.

#### `GeneModel(data = nil)`

Create a new `GeneModel` entity instance. Pass `nil` for no initial data.

#### `GoldenGateFidelity(data = nil)`

Create a new `GoldenGateFidelity` entity instance. Pass `nil` for no initial data.

#### `HgvsConvert(data = nil)`

Create a new `HgvsConvert` entity instance. Pass `nil` for no initial data.

#### `IdMapPoll(data = nil)`

Create a new `IdMapPoll` entity instance. Pass `nil` for no initial data.

#### `IdMapSubmit(data = nil)`

Create a new `IdMapSubmit` entity instance. Pass `nil` for no initial data.

#### `InSilicoPcr(data = nil)`

Create a new `InSilicoPcr` entity instance. Pass `nil` for no initial data.

#### `KaspPrimerDesign(data = nil)`

Create a new `KaspPrimerDesign` entity instance. Pass `nil` for no initial data.

#### `ListTool(data = nil)`

Create a new `ListTool` entity instance. Pass `nil` for no initial data.

#### `MeltingTemperature(data = nil)`

Create a new `MeltingTemperature` entity instance. Pass `nil` for no initial data.

#### `MotifFinder(data = nil)`

Create a new `MotifFinder` entity instance. Pass `nil` for no initial data.

#### `MultipleSequenceAlignment(data = nil)`

Create a new `MultipleSequenceAlignment` entity instance. Pass `nil` for no initial data.

#### `OligoAnalysi(data = nil)`

Create a new `OligoAnalysi` entity instance. Pass `nil` for no initial data.

#### `OrthologMap(data = nil)`

Create a new `OrthologMap` entity instance. Pass `nil` for no initial data.

#### `PairwiseAlignment(data = nil)`

Create a new `PairwiseAlignment` entity instance. Pass `nil` for no initial data.

#### `ParseGenbank(data = nil)`

Create a new `ParseGenbank` entity instance. Pass `nil` for no initial data.

#### `ParseSangerTrace(data = nil)`

Create a new `ParseSangerTrace` entity instance. Pass `nil` for no initial data.

#### `PlasmidAnnotate(data = nil)`

Create a new `PlasmidAnnotate` entity instance. Pass `nil` for no initial data.

#### `PlasmidDeepAnnotate(data = nil)`

Create a new `PlasmidDeepAnnotate` entity instance. Pass `nil` for no initial data.

#### `PlasmidFullReport(data = nil)`

Create a new `PlasmidFullReport` entity instance. Pass `nil` for no initial data.

#### `PlasmidIdentify(data = nil)`

Create a new `PlasmidIdentify` entity instance. Pass `nil` for no initial data.

#### `PrimeEditingDesign(data = nil)`

Create a new `PrimeEditingDesign` entity instance. Pass `nil` for no initial data.

#### `PrimeEditingTwinDesign(data = nil)`

Create a new `PrimeEditingTwinDesign` entity instance. Pass `nil` for no initial data.

#### `PrimerDesign(data = nil)`

Create a new `PrimerDesign` entity instance. Pass `nil` for no initial data.

#### `PrimerSpecificity(data = nil)`

Create a new `PrimerSpecificity` entity instance. Pass `nil` for no initial data.

#### `ProteaseDigestion(data = nil)`

Create a new `ProteaseDigestion` entity instance. Pass `nil` for no initial data.

#### `ProteinAnnotatePoll(data = nil)`

Create a new `ProteinAnnotatePoll` entity instance. Pass `nil` for no initial data.

#### `ProteinAnnotateSubmit(data = nil)`

Create a new `ProteinAnnotateSubmit` entity instance. Pass `nil` for no initial data.

#### `ProteinHydrophobicity(data = nil)`

Create a new `ProteinHydrophobicity` entity instance. Pass `nil` for no initial data.

#### `ProteinProperty(data = nil)`

Create a new `ProteinProperty` entity instance. Pass `nil` for no initial data.

#### `RandomSequence(data = nil)`

Create a new `RandomSequence` entity instance. Pass `nil` for no initial data.

#### `RestrictionSite(data = nil)`

Create a new `RestrictionSite` entity instance. Pass `nil` for no initial data.

#### `ReverseComplement(data = nil)`

Create a new `ReverseComplement` entity instance. Pass `nil` for no initial data.

#### `ReverseTranslate(data = nil)`

Create a new `ReverseTranslate` entity instance. Pass `nil` for no initial data.

#### `RnaFold(data = nil)`

Create a new `RnaFold` entity instance. Pass `nil` for no initial data.

#### `SangerVsReference(data = nil)`

Create a new `SangerVsReference` entity instance. Pass `nil` for no initial data.

#### `SavePermalink(data = nil)`

Create a new `SavePermalink` entity instance. Pass `nil` for no initial data.

#### `SeqfileStat(data = nil)`

Create a new `SeqfileStat` entity instance. Pass `nil` for no initial data.

#### `SequenceFetch(data = nil)`

Create a new `SequenceFetch` entity instance. Pass `nil` for no initial data.

#### `SequenceFormatConvert(data = nil)`

Create a new `SequenceFormatConvert` entity instance. Pass `nil` for no initial data.

#### `SequenceReport(data = nil)`

Create a new `SequenceReport` entity instance. Pass `nil` for no initial data.

#### `SequenceSearch(data = nil)`

Create a new `SequenceSearch` entity instance. Pass `nil` for no initial data.

#### `SequencingReadbackVerify(data = nil)`

Create a new `SequencingReadbackVerify` entity instance. Pass `nil` for no initial data.

#### `SessionCreate(data = nil)`

Create a new `SessionCreate` entity instance. Pass `nil` for no initial data.

#### `SessionGet(data = nil)`

Create a new `SessionGet` entity instance. Pass `nil` for no initial data.

#### `SessionRun(data = nil)`

Create a new `SessionRun` entity instance. Pass `nil` for no initial data.

#### `SessionSet(data = nil)`

Create a new `SessionSet` entity instance. Pass `nil` for no initial data.

#### `SirnaDesign(data = nil)`

Create a new `SirnaDesign` entity instance. Pass `nil` for no initial data.

#### `SiteDirectedMutagenesi(data = nil)`

Create a new `SiteDirectedMutagenesi` entity instance. Pass `nil` for no initial data.

#### `Translate(data = nil)`

Create a new `Translate` entity instance. Pass `nil` for no initial data.

#### `VariantAnnotate(data = nil)`

Create a new `VariantAnnotate` entity instance. Pass `nil` for no initial data.

#### `VariantComparator(data = nil)`

Create a new `VariantComparator` entity instance. Pass `nil` for no initial data.

#### `VerifyAssembly(data = nil)`

Create a new `VerifyAssembly` entity instance. Pass `nil` for no initial data.

#### `VerifyConstruct(data = nil)`

Create a new `VerifyConstruct` entity instance. Pass `nil` for no initial data.

#### `VirtualGel(data = nil)`

Create a new `VirtualGel` entity instance. Pass `nil` for no initial data.

#### `VolcanoPlotData(data = nil)`

Create a new `VolcanoPlotData` entity instance. Pass `nil` for no initial data.

#### `WebSearch(data = nil)`

Create a new `WebSearch` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## AlphafoldLookupEntity

```ruby
alphafold_lookup = client.AlphafoldLookup
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `String` | Yes |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.AlphafoldLookup.create({
  "accession" => "example_accession", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AlphafoldLookupEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## AsoDesignEntity

```ruby
aso_design = client.AsoDesign
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `length` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `target` | `String` | Yes |  |
| `tool` | `String` | Yes |  |
| `wing` | `Integer` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.AsoDesign.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AsoDesignEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## BaseEditingDesignEntity

```ruby
base_editing_design = client.BaseEditingDesign
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `editor` | `String` | No |  |
| `frameStart` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `target` | `String` | Yes |  |
| `targetPosition` | `Integer` | No |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.BaseEditingDesign.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `BaseEditingDesignEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## BatchEntity

```ruby
batch = client.Batch
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `Hash` | No |  |
| `capped` | `Boolean` | Yes |  |
| `columns` | `Array` | Yes |  |
| `count` | `Integer` | Yes |  |
| `errors` | `Integer` | Yes |  |
| `input` | `String` | Yes |  |
| `limit` | `Integer` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `rows` | `Array` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Batch.create({
  "capped" => true, # Boolean
  "columns" => [], # Array
  "count" => 1, # Integer
  "errors" => 1, # Integer
  "input" => "example_input", # String
  "limit" => 1, # Integer
  "provenance" => {}, # Hash
  "rows" => [], # Array
  "tool" => "example_tool", # String
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Batch.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `BatchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## BatchWorkflowEntity

```ruby
batch__workflow = client.BatchWorkflow
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `capped` | `Boolean` | Yes |  |
| `columns` | `Array` | Yes |  |
| `count` | `Integer` | Yes |  |
| `errors` | `Integer` | Yes |  |
| `input` | `String` | Yes |  |
| `limit` | `Integer` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `rows` | `Array` | Yes |  |
| `steps` | `Array` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.BatchWorkflow.create({
  "capped" => true, # Boolean
  "columns" => [], # Array
  "count" => 1, # Integer
  "errors" => 1, # Integer
  "input" => "example_input", # String
  "limit" => 1, # Integer
  "provenance" => {}, # Hash
  "rows" => [], # Array
  "steps" => [], # Array
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.BatchWorkflow.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `BatchWorkflowEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CharacterizeSequenceEntity

```ruby
characterize_sequence = client.CharacterizeSequence
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `endPrimerLength` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `maxOrfs` | `Integer` | No |  |
| `minOrfAa` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.CharacterizeSequence.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CharacterizeSequenceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CloningSimulateEntity

```ruby
cloning_simulate = client.CloningSimulate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `Float` | No |  |
| `circular` | `Boolean` | No |  |
| `enzyme` | `String` | No |  |
| `enzyme3` | `String` | No |  |
| `enzyme5` | `String` | No |  |
| `fragments` | `Array` | No |  |
| `gate` | `Object` | No |  |
| `insert` | `String` | No |  |
| `method` | `String` | Yes |  |
| `names` | `Array` | No |  |
| `ok` | `Object` | Yes |  |
| `overlapLen` | `Integer` | No |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |
| `vector` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.CloningSimulate.create({
  "method" => "example_method", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CloningSimulateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CodonAdaptationIndexEntity

```ruby
codon_adaptation_index = client.CodonAdaptationIndex
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frameStart` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `organism` | `String` | No |  |
| `provenance` | `Hash` | Yes |  |
| `rareThreshold` | `Float` | No |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.CodonAdaptationIndex.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CodonAdaptationIndexEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CodonOptimizeEntity

```ruby
codon_optimize = client.CodonOptimize
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `organism` | `String` | No |  |
| `protein` | `String` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.CodonOptimize.create({
  "ok" => "example_ok", # Object
  "protein" => "example_protein", # String
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CodonOptimizeEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ConstructAutofixEntity

```ruby
construct_autofix = client.ConstructAutofix
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoidEnzymes` | `Array` | No |  |
| `crypticOrfMinAa` | `Integer` | No |  |
| `frameStart` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `gcHigh` | `Float` | No |  |
| `gcLow` | `Float` | No |  |
| `gcWindow` | `Integer` | No |  |
| `homopolymerMin` | `Integer` | No |  |
| `maxPasses` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `organism` | `String` | No |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ConstructAutofix.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ConstructAutofixEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ConstructQcEntity

```ruby
construct_qc = client.ConstructQc
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoidEnzymes` | `Array` | No |  |
| `crypticOrfMinAa` | `Integer` | No |  |
| `frameStart` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `gcHigh` | `Float` | No |  |
| `gcLow` | `Float` | No |  |
| `gcWindow` | `Integer` | No |  |
| `homopolymerMin` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ConstructQc.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ConstructQcEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CrisprGrnaDesignEntity

```ruby
crispr_grna_design = client.CrisprGrnaDesign
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `minScore` | `Float` | No |  |
| `nuclease` | `String` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `searchReverseStrand` | `Boolean` | No |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.CrisprGrnaDesign.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CrisprGrnaDesignEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CrisprHdrDonorEntity

```ruby
crispr_hdr_donor = client.CrisprHdrDonor
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armLength` | `Integer` | No |  |
| `blockPam` | `Boolean` | No |  |
| `designGenotypingPrimers` | `Boolean` | No |  |
| `editEnd` | `Integer` | No |  |
| `editStart` | `Integer` | No |  |
| `frameStart` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `guideEnd` | `Integer` | No |  |
| `guideStart` | `Integer` | No |  |
| `guideStrand` | `String` | No |  |
| `nuclease` | `String` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `replacement` | `String` | Yes |  |
| `result` | `Hash` | Yes |  |
| `targetSequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.CrisprHdrDonor.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "replacement" => "example_replacement", # String
  "result" => {}, # Hash
  "targetSequence" => "example_targetSequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CrisprHdrDonorEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CrisprOfftargetCheckEntity

```ruby
crispr_offtarget_check = client.CrisprOfftargetCheck
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `maxMismatches` | `Integer` | No |  |
| `nuclease` | `String` | No |  |
| `ok` | `Object` | Yes |  |
| `protospacer` | `String` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.CrisprOfftargetCheck.create({
  "ok" => "example_ok", # Object
  "protospacer" => "example_protospacer", # String
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CrisprOfftargetCheckEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CrossDimerEntity

```ruby
cross_dimer = client.CrossDimer
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequenceA` | `String` | Yes |  |
| `sequenceB` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.CrossDimer.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequenceA" => "example_sequenceA", # String
  "sequenceB" => "example_sequenceB", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CrossDimerEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## DnaMolarityEntity

```ruby
dna_molarity = client.DnaMolarity
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `length` | `Integer` | No |  |
| `massNg` | `Float` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | No |  |
| `tool` | `String` | Yes |  |
| `type` | `String` | No |  |
| `volumeUl` | `Float` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.DnaMolarity.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `DnaMolarityEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## DoubleDigestEntity

```ruby
double_digest = client.DoubleDigest
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzymeA` | `String` | Yes |  |
| `enzymeB` | `String` | Yes |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.DoubleDigest.create({
  "enzymeA" => "example_enzymeA", # String
  "enzymeB" => "example_enzymeB", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `DoubleDigestEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ExportEchoPicklistEntity

```ruby
export_echo_picklist = client.ExportEchoPicklist
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `reactions` | `Array` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ExportEchoPicklist.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reactions" => [], # Array
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ExportEchoPicklistEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ExportOpentronsProtocolEntity

```ruby
export_opentrons_protocol = client.ExportOpentronsProtocol
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `protocolName` | `String` | No |  |
| `provenance` | `Hash` | Yes |  |
| `reactions` | `Array` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ExportOpentronsProtocol.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reactions" => [], # Array
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ExportOpentronsProtocolEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ExportPlateLayoutEntity

```ruby
export_plate_layout = client.ExportPlateLayout
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `reactions` | `Array` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ExportPlateLayout.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reactions" => [], # Array
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ExportPlateLayoutEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ExpressionHeatmapClusterEntity

```ruby
expression_heatmap_cluster = client.ExpressionHeatmapCluster
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clusterCols` | `Boolean` | No |  |
| `clusterRows` | `Boolean` | No |  |
| `distanceMetric` | `String` | No |  |
| `gate` | `Object` | No |  |
| `genes` | `Array` | Yes |  |
| `linkage` | `String` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `samples` | `Array` | Yes |  |
| `tool` | `String` | Yes |  |
| `values` | `Array` | Yes |  |
| `zScoreRows` | `Boolean` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ExpressionHeatmapCluster.create({
  "genes" => [], # Array
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "samples" => [], # Array
  "tool" => "example_tool", # String
  "values" => [], # Array
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ExpressionHeatmapClusterEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## FastqQcReportEntity

```ruby
fastq_qc_report = client.FastqQcReport
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `input` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `qualityOffset` | `Integer` | No |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.FastqQcReport.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `FastqQcReportEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## FastqTrimEntity

```ruby
fastq_trim = client.FastqTrim
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `input` | `String` | Yes |  |
| `minLength` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `qualityOffset` | `Integer` | No |  |
| `qualityThreshold` | `Integer` | No |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.FastqTrim.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `FastqTrimEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## FindOrfEntity

```ruby
find_orf = client.FindOrf
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `minAaLength` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `requireStop` | `Boolean` | No |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.FindOrf.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `FindOrfEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## FormatSequenceEntity

```ruby
format_sequence = client.FormatSequence
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `caseMode` | `String` | No |  |
| `convert` | `String` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `reverse` | `Boolean` | No |  |
| `sequence` | `String` | Yes |  |
| `stripNonLetters` | `Boolean` | No |  |
| `tool` | `String` | Yes |  |
| `width` | `Integer` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.FormatSequence.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `FormatSequenceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## FunctionalEnrichmentEntity

```ruby
functional_enrichment = client.FunctionalEnrichment
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `background` | `Array` | No |  |
| `collections` | `Array` | No |  |
| `gate` | `Object` | No |  |
| `genes` | `Array` | Yes |  |
| `maxTermSize` | `Integer` | No |  |
| `minTermSize` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.FunctionalEnrichment.create({
  "genes" => [], # Array
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `FunctionalEnrichmentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## GcContentEntity

```ruby
gc_content = client.GcContent
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.GcContent.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `GcContentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## GeneDossierEntity

```ruby
gene_dossier = client.GeneDossier
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `gene` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.GeneDossier.create({
  "gene" => "example_gene", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `GeneDossierEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## GeneExpressionEntity

```ruby
gene_expression = client.GeneExpression
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `gene` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.GeneExpression.create({
  "gene" => "example_gene", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `GeneExpressionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## GeneModelEntity

```ruby
gene_model = client.GeneModel
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `gene` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.GeneModel.create({
  "gene" => "example_gene", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `GeneModelEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## GoldenGateFidelityEntity

```ruby
golden_gate_fidelity = client.GoldenGateFidelity
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `compareToNamedSet` | `String` | No |  |
| `dataset` | `String` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `overhangs` | `Array` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `riskThreshold` | `Float` | No |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.GoldenGateFidelity.create({
  "ok" => "example_ok", # Object
  "overhangs" => [], # Array
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `GoldenGateFidelityEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## HgvsConvertEntity

```ruby
hgvs_convert = client.HgvsConvert
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |
| `variant` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.HgvsConvert.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
  "variant" => "example_variant", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `HgvsConvertEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## IdMapPollEntity

```ruby
id_map_poll = client.IdMapPoll
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `jobId` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.IdMapPoll.create({
  "jobId" => "example_jobId", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `IdMapPollEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## IdMapSubmitEntity

```ruby
id_map_submit = client.IdMapSubmit
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `String` | Yes |  |
| `gate` | `Object` | No |  |
| `ids` | `Array` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `taxId` | `String` | No |  |
| `to` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.IdMapSubmit.create({
  "from" => "example_from", # String
  "ids" => [], # Array
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "to" => "example_to", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `IdMapSubmitEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## InSilicoPcrEntity

```ruby
in_silico_pcr = client.InSilicoPcr
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `Boolean` | No |  |
| `forwardPrimer` | `String` | Yes |  |
| `gate` | `Object` | No |  |
| `maxMismatches` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `reversePrimer` | `String` | Yes |  |
| `template` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.InSilicoPcr.create({
  "forwardPrimer" => "example_forwardPrimer", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "reversePrimer" => "example_reversePrimer", # String
  "template" => "example_template", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `InSilicoPcrEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## KaspPrimerDesignEntity

```ruby
kasp_primer_design = client.KaspPrimerDesign
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addSecondaryMismatch` | `Boolean` | No |  |
| `alleleA` | `String` | Yes |  |
| `alleleB` | `String` | Yes |  |
| `gate` | `Object` | No |  |
| `maxAmplicon` | `Integer` | No |  |
| `minAmplicon` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `snpPosition` | `Integer` | Yes |  |
| `target` | `String` | Yes |  |
| `targetCoreTm` | `Float` | No |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.KaspPrimerDesign.create({
  "alleleA" => "example_alleleA", # String
  "alleleB" => "example_alleleB", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "snpPosition" => 1, # Integer
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `KaspPrimerDesignEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ListToolEntity

```ruby
list_tool = client.ListTool
```

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.ListTool.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ListToolEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MeltingTemperatureEntity

```ruby
melting_temperature = client.MeltingTemperature
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntpMM` | `Float` | No |  |
| `gate` | `Object` | No |  |
| `mgMM` | `Float` | No |  |
| `naMM` | `Float` | No |  |
| `ok` | `Object` | Yes |  |
| `oligoNM` | `Float` | No |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `targetTm` | `Float` | No |  |
| `tmTolerance` | `Float` | No |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.MeltingTemperature.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `MeltingTemperatureEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MotifFinderEntity

```ruby
motif_finder = client.MotifFinder
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `maxMismatches` | `Integer` | No |  |
| `motif` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `searchReverseStrand` | `Boolean` | No |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.MotifFinder.create({
  "motif" => "example_motif", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `MotifFinderEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MultipleSequenceAlignmentEntity

```ruby
multiple_sequence_alignment = client.MultipleSequenceAlignment
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `input` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.MultipleSequenceAlignment.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `MultipleSequenceAlignmentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## OligoAnalysiEntity

```ruby
oligo_analysi = client.OligoAnalysi
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntpMM` | `Float` | No |  |
| `gate` | `Object` | No |  |
| `mgMM` | `Float` | No |  |
| `naMM` | `Float` | No |  |
| `ok` | `Object` | Yes |  |
| `oligoNM` | `Float` | No |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.OligoAnalysi.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `OligoAnalysiEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## OrthologMapEntity

```ruby
ortholog_map = client.OrthologMap
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sourceSpecies` | `String` | No |  |
| `symbols` | `Array` | Yes |  |
| `targetSpecies` | `String` | Yes |  |
| `tool` | `String` | Yes |  |
| `type` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.OrthologMap.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "symbols" => [], # Array
  "targetSpecies" => "example_targetSpecies", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `OrthologMapEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PairwiseAlignmentEntity

```ruby
pairwise_alignment = client.PairwiseAlignment
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gap` | `Float` | No |  |
| `gate` | `Object` | No |  |
| `match` | `Float` | No |  |
| `mismatch` | `Float` | No |  |
| `mode` | `String` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `seqA` | `String` | Yes |  |
| `seqB` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PairwiseAlignment.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "seqA" => "example_seqA", # String
  "seqB" => "example_seqB", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PairwiseAlignmentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ParseGenbankEntity

```ruby
parse_genbank = client.ParseGenbank
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `text` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ParseGenbank.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "text" => "example_text", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ParseGenbankEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ParseSangerTraceEntity

```ruby
parse_sanger_trace = client.ParseSangerTrace
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fileBase64` | `String` | Yes |  |
| `fileName` | `String` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ParseSangerTrace.create({
  "fileBase64" => "example_fileBase64", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ParseSangerTraceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PlasmidAnnotateEntity

```ruby
plasmid_annotate = client.PlasmidAnnotate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PlasmidAnnotate.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PlasmidAnnotateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PlasmidDeepAnnotateEntity

```ruby
plasmid_deep_annotate = client.PlasmidDeepAnnotate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `Boolean` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PlasmidDeepAnnotate.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PlasmidDeepAnnotateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PlasmidFullReportEntity

```ruby
plasmid_full_report = client.PlasmidFullReport
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `Boolean` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |
| `topN` | `Integer` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PlasmidFullReport.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PlasmidFullReportEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PlasmidIdentifyEntity

```ruby
plasmid_identify = client.PlasmidIdentify
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `Boolean` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |
| `topN` | `Integer` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PlasmidIdentify.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PlasmidIdentifyEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PrimeEditingDesignEntity

```ruby
prime_editing_design = client.PrimeEditingDesign
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `editEnd` | `Integer` | Yes |  |
| `editStart` | `Integer` | Yes |  |
| `frameStart` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `insertedSeq` | `String` | No |  |
| `ok` | `Object` | Yes |  |
| `pbsLength` | `Integer` | No |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `rttHomology` | `Integer` | No |  |
| `target` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PrimeEditingDesign.create({
  "editEnd" => 1, # Integer
  "editStart" => 1, # Integer
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PrimeEditingDesignEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PrimeEditingTwinDesignEntity

```ruby
prime_editing_twin_design = client.PrimeEditingTwinDesign
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `newSequence` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `overlapLength` | `Integer` | No |  |
| `pbsLength` | `Integer` | No |  |
| `provenance` | `Hash` | Yes |  |
| `replaceEnd` | `Integer` | Yes |  |
| `replaceStart` | `Integer` | Yes |  |
| `result` | `Hash` | Yes |  |
| `target` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PrimeEditingTwinDesign.create({
  "newSequence" => "example_newSequence", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "replaceEnd" => 1, # Integer
  "replaceStart" => 1, # Integer
  "result" => {}, # Hash
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PrimeEditingTwinDesignEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PrimerDesignEntity

```ruby
primer_design = client.PrimerDesign
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ampliconMax` | `Integer` | No |  |
| `ampliconMin` | `Integer` | No |  |
| `dntpMM` | `Float` | No |  |
| `gate` | `Object` | No |  |
| `gcMax` | `Float` | No |  |
| `gcMin` | `Float` | No |  |
| `lenMax` | `Integer` | No |  |
| `lenMin` | `Integer` | No |  |
| `lenOpt` | `Integer` | No |  |
| `maxReturn` | `Integer` | No |  |
| `mgMM` | `Float` | No |  |
| `naMM` | `Float` | No |  |
| `ok` | `Object` | Yes |  |
| `oligoNM` | `Float` | No |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `targetEnd` | `Integer` | No |  |
| `targetStart` | `Integer` | No |  |
| `template` | `String` | Yes |  |
| `tmMax` | `Float` | No |  |
| `tmMaxDiff` | `Float` | No |  |
| `tmMin` | `Float` | No |  |
| `tmOpt` | `Float` | No |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PrimerDesign.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "template" => "example_template", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PrimerDesignEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PrimerSpecificityEntity

```ruby
primer_specificity = client.PrimerSpecificity
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `forwardPrimer` | `String` | Yes |  |
| `gate` | `Object` | No |  |
| `maxMismatches` | `Integer` | No |  |
| `maxProductLength` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `reversePrimer` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PrimerSpecificity.create({
  "forwardPrimer" => "example_forwardPrimer", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "reversePrimer" => "example_reversePrimer", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PrimerSpecificityEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ProteaseDigestionEntity

```ruby
protease_digestion = client.ProteaseDigestion
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `maxMass` | `Float` | No |  |
| `maxPeptides` | `Integer` | No |  |
| `minMass` | `Float` | No |  |
| `missedCleavages` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `protease` | `String` | No |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ProteaseDigestion.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ProteaseDigestionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ProteinAnnotatePollEntity

```ruby
protein_annotate_poll = client.ProteinAnnotatePoll
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `jobId` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ProteinAnnotatePoll.create({
  "jobId" => "example_jobId", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ProteinAnnotatePollEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ProteinAnnotateSubmitEntity

```ruby
protein_annotate_submit = client.ProteinAnnotateSubmit
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `appl` | `String` | No |  |
| `gate` | `Object` | No |  |
| `goterms` | `Boolean` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ProteinAnnotateSubmit.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ProteinAnnotateSubmitEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ProteinHydrophobicityEntity

```ruby
protein_hydrophobicity = client.ProteinHydrophobicity
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `scale` | `String` | No |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |
| `window` | `Integer` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ProteinHydrophobicity.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ProteinHydrophobicityEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ProteinPropertyEntity

```ruby
protein_property = client.ProteinProperty
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `chargeStep` | `Float` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ProteinProperty.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ProteinPropertyEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RandomSequenceEntity

```ruby
random_sequence = client.RandomSequence
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `gcContent` | `Float` | No |  |
| `kind` | `String` | No |  |
| `length` | `Integer` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.RandomSequence.create({
  "length" => 1, # Integer
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RandomSequenceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RestrictionSiteEntity

```ruby
restriction_site = client.RestrictionSite
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzymes` | `Array` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.RestrictionSite.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RestrictionSiteEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ReverseComplementEntity

```ruby
reverse_complement = client.ReverseComplement
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |
| `type` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ReverseComplement.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ReverseComplementEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ReverseTranslateEntity

```ruby
reverse_translate = client.ReverseTranslate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `mode` | `String` | No |  |
| `ok` | `Object` | Yes |  |
| `organism` | `String` | No |  |
| `protein` | `String` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.ReverseTranslate.create({
  "ok" => "example_ok", # Object
  "protein" => "example_protein", # String
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ReverseTranslateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RnaFoldEntity

```ruby
rna_fold = client.RnaFold
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.RnaFold.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RnaFoldEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SangerVsReferenceEntity

```ruby
sanger_vs_reference = client.SangerVsReference
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fileBase64` | `String` | No |  |
| `fileName` | `String` | No |  |
| `gate` | `Object` | No |  |
| `minCoverage` | `Float` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `read` | `String` | No |  |
| `reference` | `String` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SangerVsReference.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reference" => "example_reference", # String
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SangerVsReferenceEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SavePermalinkEntity

```ruby
save_permalink = client.SavePermalink
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `Hash` | Yes |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SavePermalink.create({
  "args" => {}, # Hash
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SavePermalinkEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SeqfileStatEntity

```ruby
seqfile_stat = client.SeqfileStat
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `input` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `qualityOffset` | `Integer` | No |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SeqfileStat.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SeqfileStatEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SequenceFetchEntity

```ruby
sequence_fetch = client.SequenceFetch
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `String` | Yes |  |
| `db` | `String` | No |  |
| `format` | `String` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SequenceFetch.create({
  "accession" => "example_accession", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SequenceFetchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SequenceFormatConvertEntity

```ruby
sequence_format_convert = client.SequenceFormatConvert
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `String` | No |  |
| `gate` | `Object` | No |  |
| `input` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `to` | `String` | No |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SequenceFormatConvert.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SequenceFormatConvertEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SequenceReportEntity

```ruby
sequence_report = client.SequenceReport
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `endPrimerLength` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `maxOrfs` | `Integer` | No |  |
| `minOrfAa` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SequenceReport.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SequenceReportEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SequenceSearchEntity

```ruby
sequence_search = client.SequenceSearch
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `db` | `String` | No |  |
| `gate` | `Object` | No |  |
| `gene` | `String` | No |  |
| `maxResults` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `organism` | `String` | No |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `term` | `String` | No |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SequenceSearch.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SequenceSearchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SequencingReadbackVerifyEntity

```ruby
sequencing_readback_verify = client.SequencingReadbackVerify
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `minSupportingReads` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `reads` | `String` | Yes |  |
| `reference` | `String` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SequencingReadbackVerify.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reads" => "example_reads", # String
  "reference" => "example_reference", # String
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SequencingReadbackVerifyEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SessionCreateEntity

```ruby
session_create = client.SessionCreate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entries` | `Hash` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SessionCreate.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SessionCreateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SessionGetEntity

```ruby
session_get = client.SessionGet
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `names` | `Array` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sessionId` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SessionGet.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sessionId" => "example_sessionId", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SessionGetEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SessionRunEntity

```ruby
session_run = client.SessionRun
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `Hash` | No |  |
| `fromSession` | `Hash` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sessionId` | `String` | Yes |  |
| `tool` | `String` | Yes |  |
| `writeBack` | `Hash` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SessionRun.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sessionId" => "example_sessionId", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SessionRunEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SessionSetEntity

```ruby
session_set = client.SessionSet
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entries` | `Hash` | Yes |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sessionId` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SessionSet.create({
  "entries" => {}, # Hash
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sessionId" => "example_sessionId", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SessionSetEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SirnaDesignEntity

```ruby
sirna_design = client.SirnaDesign
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `minReynolds` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `shRnaLoop` | `String` | No |  |
| `target` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SirnaDesign.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SirnaDesignEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SiteDirectedMutagenesiEntity

```ruby
site_directed_mutagenesi = client.SiteDirectedMutagenesi
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `Float` | No |  |
| `dntpMM` | `Float` | No |  |
| `editKind` | `String` | No |  |
| `frameStart` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `mgMM` | `Float` | No |  |
| `naMM` | `Float` | No |  |
| `newBase` | `String` | No |  |
| `ok` | `Object` | Yes |  |
| `oligoNM` | `Float` | No |  |
| `organism` | `String` | No |  |
| `position` | `Integer` | No |  |
| `provenance` | `Hash` | Yes |  |
| `residue` | `Integer` | No |  |
| `result` | `Hash` | Yes |  |
| `style` | `String` | No |  |
| `targetAa` | `String` | No |  |
| `template` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SiteDirectedMutagenesi.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "template" => "example_template", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SiteDirectedMutagenesiEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TranslateEntity

```ruby
translate = client.Translate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frame` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `toStop` | `Boolean` | No |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Translate.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TranslateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## VariantAnnotateEntity

```ruby
variant_annotate = client.VariantAnnotate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `assembly` | `String` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |
| `variant` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.VariantAnnotate.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
  "variant" => "example_variant", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `VariantAnnotateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## VariantComparatorEntity

```ruby
variant_comparator = client.VariantComparator
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `coding` | `Boolean` | No |  |
| `frameStart` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `query` | `String` | Yes |  |
| `reference` | `String` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.VariantComparator.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "query" => "example_query", # String
  "reference" => "example_reference", # String
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `VariantComparatorEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## VerifyAssemblyEntity

```ruby
verify_assembly = client.VerifyAssembly
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `Float` | No |  |
| `circular` | `Boolean` | No |  |
| `claimedConstruct` | `String` | Yes |  |
| `coding` | `Boolean` | No |  |
| `enzyme` | `String` | No |  |
| `enzyme3` | `String` | No |  |
| `enzyme5` | `String` | No |  |
| `fragmentPcrs` | `Array` | No |  |
| `fragments` | `Array` | No |  |
| `frameStart` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `insert` | `String` | No |  |
| `insertPcr` | `Hash` | No |  |
| `method` | `String` | Yes |  |
| `names` | `Array` | No |  |
| `ok` | `Object` | Yes |  |
| `overlapLen` | `Integer` | No |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |
| `vector` | `String` | No |  |
| `vectorPcr` | `Hash` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.VerifyAssembly.create({
  "claimedConstruct" => "example_claimedConstruct", # String
  "method" => "example_method", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `VerifyAssemblyEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## VerifyConstructEntity

```ruby
verify_construct = client.VerifyConstruct
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `claimedConstruct` | `String` | Yes |  |
| `expectedFrameStart` | `Integer` | No |  |
| `gate` | `Object` | No |  |
| `insertForwardPrimer` | `String` | Yes |  |
| `insertReversePrimer` | `String` | Yes |  |
| `insertTemplate` | `String` | Yes |  |
| `maxPrimerMismatches` | `Integer` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `templateCircular` | `Boolean` | No |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.VerifyConstruct.create({
  "claimedConstruct" => "example_claimedConstruct", # String
  "insertForwardPrimer" => "example_insertForwardPrimer", # String
  "insertReversePrimer" => "example_insertReversePrimer", # String
  "insertTemplate" => "example_insertTemplate", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `VerifyConstructEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## VirtualGelEntity

```ruby
virtual_gel = client.VirtualGel
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `Boolean` | No |  |
| `enzymes` | `Array` | No |  |
| `gate` | `Object` | No |  |
| `ladder` | `String` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `sequence` | `String` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.VirtualGel.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `VirtualGelEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## VolcanoPlotDataEntity

```ruby
volcano_plot_data = client.VolcanoPlotData
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes |  |
| `rows` | `Array` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.VolcanoPlotData.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "rows" => [], # Array
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `VolcanoPlotDataEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## WebSearchEntity

```ruby
web_search = client.WebSearch
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Object` | No |  |
| `max_results` | `Float` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `query` | `String` | Yes |  |
| `result` | `Hash` | Yes |  |
| `tool` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.WebSearch.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "query" => "example_query", # String
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `WebSearchEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = SeqbenchMcpSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

