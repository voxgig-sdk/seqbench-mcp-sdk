# SeqbenchMcp Lua SDK Reference

Complete API reference for the SeqbenchMcp Lua SDK.


## SeqbenchMcpSDK

### Constructor

```lua
local sdk = require("seqbench-mcp_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `AlphafoldLookup(data)`

Create a new `AlphafoldLookup` entity instance. Pass `nil` for no initial data.

#### `AsoDesign(data)`

Create a new `AsoDesign` entity instance. Pass `nil` for no initial data.

#### `BaseEditingDesign(data)`

Create a new `BaseEditingDesign` entity instance. Pass `nil` for no initial data.

#### `Batch(data)`

Create a new `Batch` entity instance. Pass `nil` for no initial data.

#### `BatchWorkflow(data)`

Create a new `BatchWorkflow` entity instance. Pass `nil` for no initial data.

#### `CharacterizeSequence(data)`

Create a new `CharacterizeSequence` entity instance. Pass `nil` for no initial data.

#### `CloningSimulate(data)`

Create a new `CloningSimulate` entity instance. Pass `nil` for no initial data.

#### `CodonAdaptationIndex(data)`

Create a new `CodonAdaptationIndex` entity instance. Pass `nil` for no initial data.

#### `CodonOptimize(data)`

Create a new `CodonOptimize` entity instance. Pass `nil` for no initial data.

#### `ConstructAutofix(data)`

Create a new `ConstructAutofix` entity instance. Pass `nil` for no initial data.

#### `ConstructQc(data)`

Create a new `ConstructQc` entity instance. Pass `nil` for no initial data.

#### `CrisprGrnaDesign(data)`

Create a new `CrisprGrnaDesign` entity instance. Pass `nil` for no initial data.

#### `CrisprHdrDonor(data)`

Create a new `CrisprHdrDonor` entity instance. Pass `nil` for no initial data.

#### `CrisprOfftargetCheck(data)`

Create a new `CrisprOfftargetCheck` entity instance. Pass `nil` for no initial data.

#### `CrossDimer(data)`

Create a new `CrossDimer` entity instance. Pass `nil` for no initial data.

#### `DnaMolarity(data)`

Create a new `DnaMolarity` entity instance. Pass `nil` for no initial data.

#### `DoubleDigest(data)`

Create a new `DoubleDigest` entity instance. Pass `nil` for no initial data.

#### `ExportEchoPicklist(data)`

Create a new `ExportEchoPicklist` entity instance. Pass `nil` for no initial data.

#### `ExportOpentronsProtocol(data)`

Create a new `ExportOpentronsProtocol` entity instance. Pass `nil` for no initial data.

#### `ExportPlateLayout(data)`

Create a new `ExportPlateLayout` entity instance. Pass `nil` for no initial data.

#### `ExpressionHeatmapCluster(data)`

Create a new `ExpressionHeatmapCluster` entity instance. Pass `nil` for no initial data.

#### `FastqQcReport(data)`

Create a new `FastqQcReport` entity instance. Pass `nil` for no initial data.

#### `FastqTrim(data)`

Create a new `FastqTrim` entity instance. Pass `nil` for no initial data.

#### `FindOrf(data)`

Create a new `FindOrf` entity instance. Pass `nil` for no initial data.

#### `FormatSequence(data)`

Create a new `FormatSequence` entity instance. Pass `nil` for no initial data.

#### `FunctionalEnrichment(data)`

Create a new `FunctionalEnrichment` entity instance. Pass `nil` for no initial data.

#### `GcContent(data)`

Create a new `GcContent` entity instance. Pass `nil` for no initial data.

#### `GeneDossier(data)`

Create a new `GeneDossier` entity instance. Pass `nil` for no initial data.

#### `GeneExpression(data)`

Create a new `GeneExpression` entity instance. Pass `nil` for no initial data.

#### `GeneModel(data)`

Create a new `GeneModel` entity instance. Pass `nil` for no initial data.

#### `GoldenGateFidelity(data)`

Create a new `GoldenGateFidelity` entity instance. Pass `nil` for no initial data.

#### `HgvsConvert(data)`

Create a new `HgvsConvert` entity instance. Pass `nil` for no initial data.

#### `IdMapPoll(data)`

Create a new `IdMapPoll` entity instance. Pass `nil` for no initial data.

#### `IdMapSubmit(data)`

Create a new `IdMapSubmit` entity instance. Pass `nil` for no initial data.

#### `InSilicoPcr(data)`

Create a new `InSilicoPcr` entity instance. Pass `nil` for no initial data.

#### `KaspPrimerDesign(data)`

Create a new `KaspPrimerDesign` entity instance. Pass `nil` for no initial data.

#### `ListTool(data)`

Create a new `ListTool` entity instance. Pass `nil` for no initial data.

#### `MeltingTemperature(data)`

Create a new `MeltingTemperature` entity instance. Pass `nil` for no initial data.

#### `MotifFinder(data)`

Create a new `MotifFinder` entity instance. Pass `nil` for no initial data.

#### `MultipleSequenceAlignment(data)`

Create a new `MultipleSequenceAlignment` entity instance. Pass `nil` for no initial data.

#### `OligoAnalysi(data)`

Create a new `OligoAnalysi` entity instance. Pass `nil` for no initial data.

#### `OrthologMap(data)`

Create a new `OrthologMap` entity instance. Pass `nil` for no initial data.

#### `PairwiseAlignment(data)`

Create a new `PairwiseAlignment` entity instance. Pass `nil` for no initial data.

#### `ParseGenbank(data)`

Create a new `ParseGenbank` entity instance. Pass `nil` for no initial data.

#### `ParseSangerTrace(data)`

Create a new `ParseSangerTrace` entity instance. Pass `nil` for no initial data.

#### `PlasmidAnnotate(data)`

Create a new `PlasmidAnnotate` entity instance. Pass `nil` for no initial data.

#### `PlasmidDeepAnnotate(data)`

Create a new `PlasmidDeepAnnotate` entity instance. Pass `nil` for no initial data.

#### `PlasmidFullReport(data)`

Create a new `PlasmidFullReport` entity instance. Pass `nil` for no initial data.

#### `PlasmidIdentify(data)`

Create a new `PlasmidIdentify` entity instance. Pass `nil` for no initial data.

#### `PrimeEditingDesign(data)`

Create a new `PrimeEditingDesign` entity instance. Pass `nil` for no initial data.

#### `PrimeEditingTwinDesign(data)`

Create a new `PrimeEditingTwinDesign` entity instance. Pass `nil` for no initial data.

#### `PrimerDesign(data)`

Create a new `PrimerDesign` entity instance. Pass `nil` for no initial data.

#### `PrimerSpecificity(data)`

Create a new `PrimerSpecificity` entity instance. Pass `nil` for no initial data.

#### `ProteaseDigestion(data)`

Create a new `ProteaseDigestion` entity instance. Pass `nil` for no initial data.

#### `ProteinAnnotatePoll(data)`

Create a new `ProteinAnnotatePoll` entity instance. Pass `nil` for no initial data.

#### `ProteinAnnotateSubmit(data)`

Create a new `ProteinAnnotateSubmit` entity instance. Pass `nil` for no initial data.

#### `ProteinHydrophobicity(data)`

Create a new `ProteinHydrophobicity` entity instance. Pass `nil` for no initial data.

#### `ProteinProperty(data)`

Create a new `ProteinProperty` entity instance. Pass `nil` for no initial data.

#### `RandomSequence(data)`

Create a new `RandomSequence` entity instance. Pass `nil` for no initial data.

#### `RestrictionSite(data)`

Create a new `RestrictionSite` entity instance. Pass `nil` for no initial data.

#### `ReverseComplement(data)`

Create a new `ReverseComplement` entity instance. Pass `nil` for no initial data.

#### `ReverseTranslate(data)`

Create a new `ReverseTranslate` entity instance. Pass `nil` for no initial data.

#### `RnaFold(data)`

Create a new `RnaFold` entity instance. Pass `nil` for no initial data.

#### `SangerVsReference(data)`

Create a new `SangerVsReference` entity instance. Pass `nil` for no initial data.

#### `SavePermalink(data)`

Create a new `SavePermalink` entity instance. Pass `nil` for no initial data.

#### `SeqfileStat(data)`

Create a new `SeqfileStat` entity instance. Pass `nil` for no initial data.

#### `SequenceFetch(data)`

Create a new `SequenceFetch` entity instance. Pass `nil` for no initial data.

#### `SequenceFormatConvert(data)`

Create a new `SequenceFormatConvert` entity instance. Pass `nil` for no initial data.

#### `SequenceReport(data)`

Create a new `SequenceReport` entity instance. Pass `nil` for no initial data.

#### `SequenceSearch(data)`

Create a new `SequenceSearch` entity instance. Pass `nil` for no initial data.

#### `SequencingReadbackVerify(data)`

Create a new `SequencingReadbackVerify` entity instance. Pass `nil` for no initial data.

#### `SessionCreate(data)`

Create a new `SessionCreate` entity instance. Pass `nil` for no initial data.

#### `SessionGet(data)`

Create a new `SessionGet` entity instance. Pass `nil` for no initial data.

#### `SessionRun(data)`

Create a new `SessionRun` entity instance. Pass `nil` for no initial data.

#### `SessionSet(data)`

Create a new `SessionSet` entity instance. Pass `nil` for no initial data.

#### `SirnaDesign(data)`

Create a new `SirnaDesign` entity instance. Pass `nil` for no initial data.

#### `SiteDirectedMutagenesi(data)`

Create a new `SiteDirectedMutagenesi` entity instance. Pass `nil` for no initial data.

#### `Translate(data)`

Create a new `Translate` entity instance. Pass `nil` for no initial data.

#### `VariantAnnotate(data)`

Create a new `VariantAnnotate` entity instance. Pass `nil` for no initial data.

#### `VariantComparator(data)`

Create a new `VariantComparator` entity instance. Pass `nil` for no initial data.

#### `VerifyAssembly(data)`

Create a new `VerifyAssembly` entity instance. Pass `nil` for no initial data.

#### `VerifyConstruct(data)`

Create a new `VerifyConstruct` entity instance. Pass `nil` for no initial data.

#### `VirtualGel(data)`

Create a new `VirtualGel` entity instance. Pass `nil` for no initial data.

#### `VolcanoPlotData(data)`

Create a new `VolcanoPlotData` entity instance. Pass `nil` for no initial data.

#### `WebSearch(data)`

Create a new `WebSearch` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## AlphafoldLookupEntity

```lua
local alphafold_lookup = client:AlphafoldLookup(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:AlphafoldLookup():create({
  accession = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AlphafoldLookupEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## AsoDesignEntity

```lua
local aso_design = client:AsoDesign(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `length` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `wing` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:AsoDesign():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  target = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AsoDesignEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## BaseEditingDesignEntity

```lua
local base_editing_design = client:BaseEditingDesign(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `editor` | `string` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `target` | `string` | Yes |  |
| `target_position` | `number` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:BaseEditingDesign():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  target = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BaseEditingDesignEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## BatchEntity

```lua
local batch = client:Batch(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `arg` | `table` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Batch():create({
  input = --[[ string ]],
  ok = --[[ any ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Batch():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BatchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## BatchWorkflowEntity

```lua
local batch__workflow = client:BatchWorkflow(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `result` | `table` | Yes |  |
| `step` | `table` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:BatchWorkflow():create({
  input = --[[ string ]],
  ok = --[[ any ]],
  result = --[[ table ]],
  step = --[[ table ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:BatchWorkflow():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BatchWorkflowEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CharacterizeSequenceEntity

```lua
local characterize_sequence = client:CharacterizeSequence(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `end_primer_length` | `number` | No |  |
| `gate` | `any` | No |  |
| `max_orf` | `number` | No |  |
| `min_orf_aa` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CharacterizeSequence():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CharacterizeSequenceEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CloningSimulateEntity

```lua
local cloning_simulate = client:CloningSimulate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `arm_tm_target` | `number` | No |  |
| `circular` | `boolean` | No |  |
| `enzyme` | `string` | No |  |
| `enzyme3` | `string` | No |  |
| `enzyme5` | `string` | No |  |
| `fragment` | `table` | No |  |
| `gate` | `any` | No |  |
| `insert` | `string` | No |  |
| `method` | `string` | Yes |  |
| `name` | `table` | No |  |
| `ok` | `any` | Yes |  |
| `overlap_len` | `number` | No |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |
| `vector` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CloningSimulate():create({
  method = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CloningSimulateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CodonAdaptationIndexEntity

```lua
local codon_adaptation_index = client:CodonAdaptationIndex(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `table` | Yes |  |
| `rare_threshold` | `number` | No |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CodonAdaptationIndex():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CodonAdaptationIndexEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CodonOptimizeEntity

```lua
local codon_optimize = client:CodonOptimize(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CodonOptimize():create({
  ok = --[[ any ]],
  protein = --[[ string ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CodonOptimizeEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ConstructAutofixEntity

```lua
local construct_autofix = client:ConstructAutofix(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoid_enzyme` | `table` | No |  |
| `cryptic_orf_min_aa` | `number` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `gc_high` | `number` | No |  |
| `gc_low` | `number` | No |  |
| `gc_window` | `number` | No |  |
| `homopolymer_min` | `number` | No |  |
| `max_pass` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ConstructAutofix():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ConstructAutofixEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ConstructQcEntity

```lua
local construct_qc = client:ConstructQc(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoid_enzyme` | `table` | No |  |
| `cryptic_orf_min_aa` | `number` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `gc_high` | `number` | No |  |
| `gc_low` | `number` | No |  |
| `gc_window` | `number` | No |  |
| `homopolymer_min` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ConstructQc():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ConstructQcEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CrisprGrnaDesignEntity

```lua
local crispr_grna_design = client:CrisprGrnaDesign(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `min_score` | `number` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `search_reverse_strand` | `boolean` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CrisprGrnaDesign():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CrisprGrnaDesignEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CrisprHdrDonorEntity

```lua
local crispr_hdr_donor = client:CrisprHdrDonor(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `arm_length` | `number` | No |  |
| `block_pam` | `boolean` | No |  |
| `design_genotyping_primer` | `boolean` | No |  |
| `edit_end` | `number` | No |  |
| `edit_start` | `number` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `guide_end` | `number` | No |  |
| `guide_start` | `number` | No |  |
| `guide_strand` | `string` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `replacement` | `string` | Yes |  |
| `result` | `table` | Yes |  |
| `target_sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CrisprHdrDonor():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  replacement = --[[ string ]],
  result = --[[ table ]],
  target_sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CrisprHdrDonorEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CrisprOfftargetCheckEntity

```lua
local crispr_offtarget_check = client:CrisprOfftargetCheck(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `max_mismatch` | `number` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `protospacer` | `string` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CrisprOfftargetCheck():create({
  ok = --[[ any ]],
  protospacer = --[[ string ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CrisprOfftargetCheckEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CrossDimerEntity

```lua
local cross_dimer = client:CrossDimer(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence_a` | `string` | Yes |  |
| `sequence_b` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CrossDimer():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence_a = --[[ string ]],
  sequence_b = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CrossDimerEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DnaMolarityEntity

```lua
local dna_molarity = client:DnaMolarity(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `length` | `number` | No |  |
| `mass_ng` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | No |  |
| `tool` | `string` | Yes |  |
| `type` | `string` | No |  |
| `volume_ul` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:DnaMolarity():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DnaMolarityEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DoubleDigestEntity

```lua
local double_digest = client:DoubleDigest(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzyme_a` | `string` | Yes |  |
| `enzyme_b` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:DoubleDigest():create({
  enzyme_a = --[[ string ]],
  enzyme_b = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DoubleDigestEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ExportEchoPicklistEntity

```lua
local export_echo_picklist = client:ExportEchoPicklist(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `reaction` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ExportEchoPicklist():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  reaction = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ExportEchoPicklistEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ExportOpentronsProtocolEntity

```lua
local export_opentrons_protocol = client:ExportOpentronsProtocol(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `protocol_name` | `string` | No |  |
| `provenance` | `table` | Yes |  |
| `reaction` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ExportOpentronsProtocol():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  reaction = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ExportOpentronsProtocolEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ExportPlateLayoutEntity

```lua
local export_plate_layout = client:ExportPlateLayout(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `reaction` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ExportPlateLayout():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  reaction = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ExportPlateLayoutEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ExpressionHeatmapClusterEntity

```lua
local expression_heatmap_cluster = client:ExpressionHeatmapCluster(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cluster_col` | `boolean` | No |  |
| `cluster_row` | `boolean` | No |  |
| `distance_metric` | `string` | No |  |
| `gate` | `any` | No |  |
| `gene` | `table` | Yes |  |
| `linkage` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sample` | `table` | Yes |  |
| `tool` | `string` | Yes |  |
| `value` | `table` | Yes |  |
| `z_score_row` | `boolean` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ExpressionHeatmapCluster():create({
  gene = --[[ table ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sample = --[[ table ]],
  tool = --[[ string ]],
  value = --[[ table ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ExpressionHeatmapClusterEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## FastqQcReportEntity

```lua
local fastq_qc_report = client:FastqQcReport(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `quality_offset` | `number` | No |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:FastqQcReport():create({
  input = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FastqQcReportEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## FastqTrimEntity

```lua
local fastq_trim = client:FastqTrim(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `min_length` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `quality_offset` | `number` | No |  |
| `quality_threshold` | `number` | No |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:FastqTrim():create({
  input = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FastqTrimEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## FindOrfEntity

```lua
local find_orf = client:FindOrf(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `min_aa_length` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `require_stop` | `boolean` | No |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:FindOrf():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FindOrfEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## FormatSequenceEntity

```lua
local format_sequence = client:FormatSequence(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `case_mode` | `string` | No |  |
| `convert` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `reverse` | `boolean` | No |  |
| `sequence` | `string` | Yes |  |
| `strip_non_letter` | `boolean` | No |  |
| `tool` | `string` | Yes |  |
| `width` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:FormatSequence():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FormatSequenceEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## FunctionalEnrichmentEntity

```lua
local functional_enrichment = client:FunctionalEnrichment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `background` | `table` | No |  |
| `collection` | `table` | No |  |
| `gate` | `any` | No |  |
| `gene` | `table` | Yes |  |
| `max_term_size` | `number` | No |  |
| `min_term_size` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:FunctionalEnrichment():create({
  gene = --[[ table ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FunctionalEnrichmentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## GcContentEntity

```lua
local gc_content = client:GcContent(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:GcContent():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GcContentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## GeneDossierEntity

```lua
local gene_dossier = client:GeneDossier(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:GeneDossier():create({
  gene = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GeneDossierEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## GeneExpressionEntity

```lua
local gene_expression = client:GeneExpression(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:GeneExpression():create({
  gene = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GeneExpressionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## GeneModelEntity

```lua
local gene_model = client:GeneModel(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:GeneModel():create({
  gene = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GeneModelEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## GoldenGateFidelityEntity

```lua
local golden_gate_fidelity = client:GoldenGateFidelity(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `compare_to_named_set` | `string` | No |  |
| `dataset` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `overhang` | `table` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `risk_threshold` | `number` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:GoldenGateFidelity():create({
  ok = --[[ any ]],
  overhang = --[[ table ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GoldenGateFidelityEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## HgvsConvertEntity

```lua
local hgvs_convert = client:HgvsConvert(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |
| `variant` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:HgvsConvert():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
  variant = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HgvsConvertEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## IdMapPollEntity

```lua
local id_map_poll = client:IdMapPoll(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `job_id` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:IdMapPoll():create({
  job_id = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IdMapPollEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## IdMapSubmitEntity

```lua
local id_map_submit = client:IdMapSubmit(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `ids` | `table` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tax_id` | `string` | No |  |
| `to` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:IdMapSubmit():create({
  from = --[[ string ]],
  ids = --[[ table ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  to = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IdMapSubmitEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## InSilicoPcrEntity

```lua
local in_silico_pcr = client:InSilicoPcr(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `boolean` | No |  |
| `forward_primer` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `max_mismatch` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `reverse_primer` | `string` | Yes |  |
| `template` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:InSilicoPcr():create({
  forward_primer = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  reverse_primer = --[[ string ]],
  template = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InSilicoPcrEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## KaspPrimerDesignEntity

```lua
local kasp_primer_design = client:KaspPrimerDesign(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `add_secondary_mismatch` | `boolean` | No |  |
| `allele_a` | `string` | Yes |  |
| `allele_b` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `max_amplicon` | `number` | No |  |
| `min_amplicon` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `snp_position` | `number` | Yes |  |
| `target` | `string` | Yes |  |
| `target_core_tm` | `number` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:KaspPrimerDesign():create({
  allele_a = --[[ string ]],
  allele_b = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  snp_position = --[[ number ]],
  target = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `KaspPrimerDesignEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ListToolEntity

```lua
local list_tool = client:ListTool(nil)
```

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:ListTool():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListToolEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MeltingTemperatureEntity

```lua
local melting_temperature = client:MeltingTemperature(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntp_mm` | `number` | No |  |
| `gate` | `any` | No |  |
| `mg_mm` | `number` | No |  |
| `na_mm` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `number` | No |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `target_tm` | `number` | No |  |
| `tm_tolerance` | `number` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:MeltingTemperature():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MeltingTemperatureEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MotifFinderEntity

```lua
local motif_finder = client:MotifFinder(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `max_mismatch` | `number` | No |  |
| `motif` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `search_reverse_strand` | `boolean` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:MotifFinder():create({
  motif = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MotifFinderEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MultipleSequenceAlignmentEntity

```lua
local multiple_sequence_alignment = client:MultipleSequenceAlignment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:MultipleSequenceAlignment():create({
  input = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MultipleSequenceAlignmentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## OligoAnalysiEntity

```lua
local oligo_analysi = client:OligoAnalysi(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntp_mm` | `number` | No |  |
| `gate` | `any` | No |  |
| `mg_mm` | `number` | No |  |
| `na_mm` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `number` | No |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:OligoAnalysi():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OligoAnalysiEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## OrthologMapEntity

```lua
local ortholog_map = client:OrthologMap(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `source_species` | `string` | No |  |
| `symbol` | `table` | Yes |  |
| `target_species` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `type` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:OrthologMap():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  symbol = --[[ table ]],
  target_species = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrthologMapEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PairwiseAlignmentEntity

```lua
local pairwise_alignment = client:PairwiseAlignment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gap` | `number` | No |  |
| `gate` | `any` | No |  |
| `match` | `number` | No |  |
| `mismatch` | `number` | No |  |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `seq_a` | `string` | Yes |  |
| `seq_b` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PairwiseAlignment():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  seq_a = --[[ string ]],
  seq_b = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PairwiseAlignmentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ParseGenbankEntity

```lua
local parse_genbank = client:ParseGenbank(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `text` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ParseGenbank():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  text = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ParseGenbankEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ParseSangerTraceEntity

```lua
local parse_sanger_trace = client:ParseSangerTrace(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `file_base64` | `string` | Yes |  |
| `file_name` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ParseSangerTrace():create({
  file_base64 = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ParseSangerTraceEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PlasmidAnnotateEntity

```lua
local plasmid_annotate = client:PlasmidAnnotate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PlasmidAnnotate():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlasmidAnnotateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PlasmidDeepAnnotateEntity

```lua
local plasmid_deep_annotate = client:PlasmidDeepAnnotate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `boolean` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PlasmidDeepAnnotate():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlasmidDeepAnnotateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PlasmidFullReportEntity

```lua
local plasmid_full_report = client:PlasmidFullReport(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `boolean` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `top_n` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PlasmidFullReport():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlasmidFullReportEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PlasmidIdentifyEntity

```lua
local plasmid_identify = client:PlasmidIdentify(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `boolean` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `top_n` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PlasmidIdentify():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlasmidIdentifyEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PrimeEditingDesignEntity

```lua
local prime_editing_design = client:PrimeEditingDesign(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `edit_end` | `number` | Yes |  |
| `edit_start` | `number` | Yes |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `inserted_seq` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `pbs_length` | `number` | No |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `rtt_homology` | `number` | No |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PrimeEditingDesign():create({
  edit_end = --[[ number ]],
  edit_start = --[[ number ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  target = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PrimeEditingDesignEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PrimeEditingTwinDesignEntity

```lua
local prime_editing_twin_design = client:PrimeEditingTwinDesign(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `new_sequence` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `overlap_length` | `number` | No |  |
| `pbs_length` | `number` | No |  |
| `provenance` | `table` | Yes |  |
| `replace_end` | `number` | Yes |  |
| `replace_start` | `number` | Yes |  |
| `result` | `table` | Yes |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PrimeEditingTwinDesign():create({
  new_sequence = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  replace_end = --[[ number ]],
  replace_start = --[[ number ]],
  result = --[[ table ]],
  target = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PrimeEditingTwinDesignEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PrimerDesignEntity

```lua
local primer_design = client:PrimerDesign(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amplicon_max` | `number` | No |  |
| `amplicon_min` | `number` | No |  |
| `dntp_mm` | `number` | No |  |
| `gate` | `any` | No |  |
| `gc_max` | `number` | No |  |
| `gc_min` | `number` | No |  |
| `len_max` | `number` | No |  |
| `len_min` | `number` | No |  |
| `len_opt` | `number` | No |  |
| `max_return` | `number` | No |  |
| `mg_mm` | `number` | No |  |
| `na_mm` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `number` | No |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `target_end` | `number` | No |  |
| `target_start` | `number` | No |  |
| `template` | `string` | Yes |  |
| `tm_max` | `number` | No |  |
| `tm_max_diff` | `number` | No |  |
| `tm_min` | `number` | No |  |
| `tm_opt` | `number` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PrimerDesign():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  template = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PrimerDesignEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PrimerSpecificityEntity

```lua
local primer_specificity = client:PrimerSpecificity(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `forward_primer` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `max_mismatch` | `number` | No |  |
| `max_product_length` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `reverse_primer` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PrimerSpecificity():create({
  forward_primer = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  reverse_primer = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PrimerSpecificityEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ProteaseDigestionEntity

```lua
local protease_digestion = client:ProteaseDigestion(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `max_mass` | `number` | No |  |
| `max_peptide` | `number` | No |  |
| `min_mass` | `number` | No |  |
| `missed_cleavage` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `protease` | `string` | No |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ProteaseDigestion():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProteaseDigestionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ProteinAnnotatePollEntity

```lua
local protein_annotate_poll = client:ProteinAnnotatePoll(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `job_id` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ProteinAnnotatePoll():create({
  job_id = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProteinAnnotatePollEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ProteinAnnotateSubmitEntity

```lua
local protein_annotate_submit = client:ProteinAnnotateSubmit(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `appl` | `string` | No |  |
| `gate` | `any` | No |  |
| `goterm` | `boolean` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ProteinAnnotateSubmit():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProteinAnnotateSubmitEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ProteinHydrophobicityEntity

```lua
local protein_hydrophobicity = client:ProteinHydrophobicity(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `scale` | `string` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `window` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ProteinHydrophobicity():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProteinHydrophobicityEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ProteinPropertyEntity

```lua
local protein_property = client:ProteinProperty(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `charge_step` | `number` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ProteinProperty():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProteinPropertyEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RandomSequenceEntity

```lua
local random_sequence = client:RandomSequence(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `gc_content` | `number` | No |  |
| `kind` | `string` | No |  |
| `length` | `number` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:RandomSequence():create({
  length = --[[ number ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RandomSequenceEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RestrictionSiteEntity

```lua
local restriction_site = client:RestrictionSite(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzyme` | `table` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:RestrictionSite():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RestrictionSiteEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ReverseComplementEntity

```lua
local reverse_complement = client:ReverseComplement(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `type` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ReverseComplement():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReverseComplementEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ReverseTranslateEntity

```lua
local reverse_translate = client:ReverseTranslate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ReverseTranslate():create({
  ok = --[[ any ]],
  protein = --[[ string ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReverseTranslateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RnaFoldEntity

```lua
local rna_fold = client:RnaFold(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:RnaFold():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RnaFoldEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SangerVsReferenceEntity

```lua
local sanger_vs_reference = client:SangerVsReference(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `file_base64` | `string` | No |  |
| `file_name` | `string` | No |  |
| `gate` | `any` | No |  |
| `min_coverage` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `read` | `string` | No |  |
| `reference` | `string` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SangerVsReference():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  reference = --[[ string ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SangerVsReferenceEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SavePermalinkEntity

```lua
local save_permalink = client:SavePermalink(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `arg` | `table` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SavePermalink():create({
  arg = --[[ table ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SavePermalinkEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SeqfileStatEntity

```lua
local seqfile_stat = client:SeqfileStat(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `quality_offset` | `number` | No |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SeqfileStat():create({
  input = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SeqfileStatEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SequenceFetchEntity

```lua
local sequence_fetch = client:SequenceFetch(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `string` | Yes |  |
| `db` | `string` | No |  |
| `format` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SequenceFetch():create({
  accession = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SequenceFetchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SequenceFormatConvertEntity

```lua
local sequence_format_convert = client:SequenceFormatConvert(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | No |  |
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `to` | `string` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SequenceFormatConvert():create({
  input = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SequenceFormatConvertEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SequenceReportEntity

```lua
local sequence_report = client:SequenceReport(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `end_primer_length` | `number` | No |  |
| `gate` | `any` | No |  |
| `max_orf` | `number` | No |  |
| `min_orf_aa` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SequenceReport():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SequenceReportEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SequenceSearchEntity

```lua
local sequence_search = client:SequenceSearch(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `db` | `string` | No |  |
| `gate` | `any` | No |  |
| `gene` | `string` | No |  |
| `max_result` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `term` | `string` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SequenceSearch():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SequenceSearchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SequencingReadbackVerifyEntity

```lua
local sequencing_readback_verify = client:SequencingReadbackVerify(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `min_supporting_read` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `read` | `string` | Yes |  |
| `reference` | `string` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SequencingReadbackVerify():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  read = --[[ string ]],
  reference = --[[ string ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SequencingReadbackVerifyEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SessionCreateEntity

```lua
local session_create = client:SessionCreate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entry` | `table` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SessionCreate():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SessionCreateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SessionGetEntity

```lua
local session_get = client:SessionGet(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `name` | `table` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SessionGet():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  session_id = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SessionGetEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SessionRunEntity

```lua
local session_run = client:SessionRun(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `arg` | `table` | No |  |
| `from_session` | `table` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `write_back` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SessionRun():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  session_id = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SessionRunEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SessionSetEntity

```lua
local session_set = client:SessionSet(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entry` | `table` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SessionSet():create({
  entry = --[[ table ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  session_id = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SessionSetEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SirnaDesignEntity

```lua
local sirna_design = client:SirnaDesign(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `min_reynold` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sh_rna_loop` | `string` | No |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SirnaDesign():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  target = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SirnaDesignEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SiteDirectedMutagenesiEntity

```lua
local site_directed_mutagenesi = client:SiteDirectedMutagenesi(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `arm_tm_target` | `number` | No |  |
| `dntp_mm` | `number` | No |  |
| `edit_kind` | `string` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `mg_mm` | `number` | No |  |
| `na_mm` | `number` | No |  |
| `new_base` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `number` | No |  |
| `organism` | `string` | No |  |
| `position` | `number` | No |  |
| `provenance` | `table` | Yes |  |
| `residue` | `number` | No |  |
| `result` | `table` | Yes |  |
| `style` | `string` | No |  |
| `target_aa` | `string` | No |  |
| `template` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SiteDirectedMutagenesi():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  template = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SiteDirectedMutagenesiEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TranslateEntity

```lua
local translate = client:Translate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frame` | `number` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `to_stop` | `boolean` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Translate():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TranslateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## VariantAnnotateEntity

```lua
local variant_annotate = client:VariantAnnotate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `assembly` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |
| `variant` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:VariantAnnotate():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
  variant = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VariantAnnotateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## VariantComparatorEntity

```lua
local variant_comparator = client:VariantComparator(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `coding` | `boolean` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `query` | `string` | Yes |  |
| `reference` | `string` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:VariantComparator():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  query = --[[ string ]],
  reference = --[[ string ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VariantComparatorEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## VerifyAssemblyEntity

```lua
local verify_assembly = client:VerifyAssembly(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `arm_tm_target` | `number` | No |  |
| `circular` | `boolean` | No |  |
| `claimed_construct` | `string` | Yes |  |
| `coding` | `boolean` | No |  |
| `enzyme` | `string` | No |  |
| `enzyme3` | `string` | No |  |
| `enzyme5` | `string` | No |  |
| `fragment` | `table` | No |  |
| `fragment_pcr` | `table` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `insert` | `string` | No |  |
| `insert_pcr` | `table` | No |  |
| `method` | `string` | Yes |  |
| `name` | `table` | No |  |
| `ok` | `any` | Yes |  |
| `overlap_len` | `number` | No |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |
| `vector` | `string` | No |  |
| `vector_pcr` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:VerifyAssembly():create({
  claimed_construct = --[[ string ]],
  method = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VerifyAssemblyEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## VerifyConstructEntity

```lua
local verify_construct = client:VerifyConstruct(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `claimed_construct` | `string` | Yes |  |
| `expected_frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `insert_forward_primer` | `string` | Yes |  |
| `insert_reverse_primer` | `string` | Yes |  |
| `insert_template` | `string` | Yes |  |
| `max_primer_mismatch` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `template_circular` | `boolean` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:VerifyConstruct():create({
  claimed_construct = --[[ string ]],
  insert_forward_primer = --[[ string ]],
  insert_reverse_primer = --[[ string ]],
  insert_template = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VerifyConstructEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## VirtualGelEntity

```lua
local virtual_gel = client:VirtualGel(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `boolean` | No |  |
| `enzyme` | `table` | No |  |
| `gate` | `any` | No |  |
| `ladder` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:VirtualGel():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequence = --[[ string ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VirtualGelEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## VolcanoPlotDataEntity

```lua
local volcano_plot_data = client:VolcanoPlotData(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes |  |
| `row` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:VolcanoPlotData():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  row = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VolcanoPlotDataEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## WebSearchEntity

```lua
local web_search = client:WebSearch(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No |  |
| `max_result` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `query` | `string` | Yes |  |
| `result` | `table` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:WebSearch():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  query = --[[ string ]],
  result = --[[ table ]],
  tool = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WebSearchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

