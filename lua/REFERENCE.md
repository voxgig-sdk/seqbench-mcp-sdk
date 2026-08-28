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
| `accession` | `string` | Yes | UniProt accession, e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `number` | No | Total gapmer length (nt). |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `wing` | `number` | No | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

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
| `editor` | `string` | No | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | `number` | No | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | `number` | No | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `args` | `table` | No | Shared tool arguments applied to every record. |
| `capped` | `boolean` | Yes | True if input exceeded the record limit. |
| `columns` | `table` | Yes |  |
| `count` | `number` | Yes |  |
| `errors` | `number` | Yes |  |
| `input` | `string` | Yes | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | `number` | Yes | Maximum records per call (500). |
| `provenance` | `table` | Yes |  |
| `rows` | `table` | Yes |  |
| `tool` | `string` | Yes | A batchable tool slug (see `GET /batch`). |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Batch():create({
  capped = --[[ boolean ]],
  columns = --[[ table ]],
  count = --[[ number ]],
  errors = --[[ number ]],
  input = --[[ string ]],
  limit = --[[ number ]],
  provenance = --[[ table ]],
  rows = --[[ table ]],
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
| `capped` | `boolean` | Yes |  |
| `columns` | `table` | Yes | Flattened "<step>·<tool>·<key>" column headers. |
| `count` | `number` | Yes |  |
| `errors` | `number` | Yes |  |
| `input` | `string` | Yes | Multi-FASTA text or one sequence per line. |
| `limit` | `number` | Yes | Maximum records per call (200). |
| `provenance` | `table` | Yes |  |
| `rows` | `table` | Yes |  |
| `steps` | `table` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:BatchWorkflow():create({
  capped = --[[ boolean ]],
  columns = --[[ table ]],
  count = --[[ number ]],
  errors = --[[ number ]],
  input = --[[ string ]],
  limit = --[[ number ]],
  provenance = --[[ table ]],
  rows = --[[ table ]],
  steps = --[[ table ]],
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
| `endPrimerLength` | `number` | No | Length of the naive end primers taken from each end. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `number` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `number` | No | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `armTmTarget` | `number` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `boolean` | No | Produce a circular product. |
| `enzyme` | `string` | No | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | `string` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | No | 5′ enzyme (restriction method). |
| `fragments` | `table` | No | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | No | Insert sequence (restriction method). |
| `method` | `string` | Yes | Assembly method. |
| `names` | `table` | No | Optional labels for each fragment. |
| `ok` | `any` | Yes |  |
| `overlapLen` | `number` | No | Gibson homology-arm length (bp). |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `vector` | `string` | No | Vector sequence (restriction method). |

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
| `frameStart` | `number` | No | 1-based position to start reading codons. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `table` | Yes |  |
| `rareThreshold` | `number` | No | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes | Protein sequence (one-letter codes). |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `avoidEnzymes` | `table` | No | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | `number` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `number` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `number` | No |  |
| `gcLow` | `number` | No |  |
| `gcWindow` | `number` | No |  |
| `homopolymerMin` | `number` | No |  |
| `maxPasses` | `number` | No | Repeat full passes until clean or no further progress. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No | Codon-usage table to prefer among synonymous options. |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `avoidEnzymes` | `table` | No | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | `number` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `number` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `number` | No | GC% above this flags a GC-rich window. |
| `gcLow` | `number` | No | GC% below this flags an AT-rich window. |
| `gcWindow` | `number` | No | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | `number` | No | Minimum run length to flag a homopolymer. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | `number` | No | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | `string` | No | Nuclease id. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `boolean` | No | Also scan the reverse strand for guides. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `armLength` | `number` | No | Homology arm length (bp) on each side. |
| `blockPam` | `boolean` | No | When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut. |
| `designGenotypingPrimers` | `boolean` | No | Also design a primer pair (on the original targetSequence) whose product spans the edit site. |
| `editEnd` | `number` | No | 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. |
| `editStart` | `number` | No | 1-based start of the region being replaced. |
| `frameStart` | `number` | No | Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | `number` | No | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | `number` | No | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | `string` | No | Strand the guide's protospacer is on. |
| `nuclease` | `string` | No | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `replacement` | `string` | Yes | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | `table` | Yes | Tool-specific output object. |
| `targetSequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CrisprHdrDonor():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  replacement = --[[ string ]],
  result = --[[ table ]],
  targetSequence = --[[ string ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | No | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | `string` | No | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` | `any` | Yes |  |
| `protospacer` | `string` | Yes | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequenceA` | `string` | Yes | First oligo (5'→3'). |
| `sequenceB` | `string` | Yes | Second oligo (5'→3'). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CrossDimer():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sequenceA = --[[ string ]],
  sequenceB = --[[ string ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `number` | No | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | `number` | No | Mass in nanograms. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | No | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `type` | `string` | No | Molecule type. |
| `volumeUl` | `number` | No | Volume in microlitres (0 = unknown; needed for concentration). |

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
| `enzymeA` | `string` | Yes | First enzyme name (e.g. |
| `enzymeB` | `string` | Yes | Second enzyme name (e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:DoubleDigest():create({
  enzymeA = --[[ string ]],
  enzymeB = --[[ string ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `reactions` | `table` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ExportEchoPicklist():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  reactions = --[[ table ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `protocolName` | `string` | No | Optional protocol name (used in the script's metadata). |
| `provenance` | `table` | Yes |  |
| `reactions` | `table` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ExportOpentronsProtocol():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  reactions = --[[ table ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `reactions` | `table` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ExportPlateLayout():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  reactions = --[[ table ]],
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
| `clusterCols` | `boolean` | No | Cluster (reorder) samples. |
| `clusterRows` | `boolean` | No | Cluster (reorder) genes. |
| `distanceMetric` | `string` | No | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `table` | Yes | Row (gene) labels. |
| `linkage` | `string` | No | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `samples` | `table` | Yes | Column (sample) labels. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `values` | `table` | Yes | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | `boolean` | No | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ExpressionHeatmapCluster():create({
  genes = --[[ table ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  samples = --[[ table ]],
  tool = --[[ string ]],
  values = --[[ table ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `qualityOffset` | `number` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | `number` | No | Reads shorter than this after trimming are dropped. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `qualityOffset` | `number` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | `number` | No | 3' quality-trim threshold (Phred score). |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | `number` | No | Minimum protein length (aa) to report. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `requireStop` | `boolean` | No | Only report ORFs terminated by a stop codon. |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `caseMode` | `string` | No |  |
| `convert` | `string` | No | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `reverse` | `boolean` | No | Reverse the sequence (no complement). |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | `boolean` | No | Remove digits, spaces and gaps (keep letters only). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `width` | `number` | No | Line-wrap width; 0 = single line. |

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
| `background` | `table` | No | Custom background/universe gene symbols. |
| `collections` | `table` | No | Which term collections to test. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `table` | Yes | Query gene symbols (human, e.g. |
| `maxTermSize` | `number` | No | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | `number` | No | Skip terms/pathways with fewer than this many background genes. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:FunctionalEnrichment():create({
  genes = --[[ table ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `compareToNamedSet` | `string` | No | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | `string` | No | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `overhangs` | `table` | Yes | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `riskThreshold` | `number` | No | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:GoldenGateFidelity():create({
  ok = --[[ any ]],
  overhangs = --[[ table ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `variant` | `string` | Yes | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:IdMapPoll():create({
  jobId = --[[ string ]],
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
| `from` | `string` | Yes | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | `table` | Yes | The ids to map, up to 1000 (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `taxId` | `string` | No | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | `string` | Yes | Target id type. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `circular` | `boolean` | No | Treat the template as circular (plasmid). |
| `forwardPrimer` | `string` | Yes | Primer 1, 5'→3'. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | No | Mismatches tolerated per primer. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `reversePrimer` | `string` | Yes | Primer 2, 5'→3' (order does not matter). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:InSilicoPcr():create({
  forwardPrimer = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  reversePrimer = --[[ string ]],
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
| `addSecondaryMismatch` | `boolean` | No | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | `string` | Yes | First allele (single base) — gets the FAM tail. |
| `alleleB` | `string` | Yes | Second allele (single base) — gets the HEX tail. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | `number` | No | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | `number` | No | Minimum amplicon length for the common reverse primer. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `snpPosition` | `number` | Yes | 1-based position of the SNP on the forward strand. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | `number` | No | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:KaspPrimerDesign():create({
  alleleA = --[[ string ]],
  alleleB = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  snpPosition = --[[ number ]],
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
| `dntpMM` | `number` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `number` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `number` | No | Total strand concentration (nM). |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | `number` | No | Optional target Tm (°C). |
| `tmTolerance` | `number` | No | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | No | Maximum allowed mismatches per match. |
| `motif` | `string` | Yes | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `boolean` | No | Also search the reverse strand. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `dntpMM` | `number` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `number` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `number` | No | Total strand concentration (nM). |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sourceSpecies` | `string` | No | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | `table` | Yes | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | `string` | Yes | Ensembl species slug to find homologs in (e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `type` | `string` | No | Homology type to return. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:OrthologMap():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  symbols = --[[ table ]],
  targetSpecies = --[[ string ]],
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
| `gap` | `number` | No | Linear gap penalty (per gap position). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | `number` | No | Match score. |
| `mismatch` | `number` | No | Mismatch penalty. |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `seqA` | `string` | Yes | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | `string` | Yes | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PairwiseAlignment():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  seqA = --[[ string ]],
  seqB = --[[ string ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `text` | `string` | Yes | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `fileBase64` | `string` | Yes | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | No | Optional original file name (echoed back). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ParseSangerTrace():create({
  fileBase64 = --[[ string ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `circular` | `boolean` | No | Treat the sequence as a circular plasmid (vs. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `circular` | `boolean` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `topN` | `number` | No | How many top-ranked backbone candidates to report. |

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
| `circular` | `boolean` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `topN` | `number` | No | How many top-ranked backbone candidates to report. |

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
| `editEnd` | `number` | Yes | 1-based inclusive end of the region being changed. |
| `editStart` | `number` | Yes | 1-based inclusive start of the region being changed. |
| `frameStart` | `number` | No | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | `string` | No | Replacement bases (forward strand). |
| `ok` | `any` | Yes |  |
| `pbsLength` | `number` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `rttHomology` | `number` | No | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | `string` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PrimeEditingDesign():create({
  editEnd = --[[ number ]],
  editStart = --[[ number ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | `string` | Yes | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` | `any` | Yes |  |
| `overlapLength` | `number` | No | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | `number` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `table` | Yes |  |
| `replaceEnd` | `number` | Yes | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | `number` | Yes | 1-based inclusive start of the region being replaced/deleted. |
| `result` | `table` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PrimeEditingTwinDesign():create({
  newSequence = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  replaceEnd = --[[ number ]],
  replaceStart = --[[ number ]],
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
| `ampliconMax` | `number` | No |  |
| `ampliconMin` | `number` | No |  |
| `dntpMM` | `number` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` | `number` | No |  |
| `gcMin` | `number` | No |  |
| `lenMax` | `number` | No |  |
| `lenMin` | `number` | No |  |
| `lenOpt` | `number` | No |  |
| `maxReturn` | `number` | No | Number of best pairs to return. |
| `mgMM` | `number` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `number` | No | Total strand concentration (nM). |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `targetEnd` | `number` | No | 1-based inclusive end of the target region (optional). |
| `targetStart` | `number` | No | 1-based inclusive start of a region the product must span (optional). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` | `number` | No |  |
| `tmMaxDiff` | `number` | No | Max Tm difference within a pair (°C). |
| `tmMin` | `number` | No |  |
| `tmOpt` | `number` | No |  |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `forwardPrimer` | `string` | Yes | Forward primer, 5'→3'. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | No | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | `number` | No | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `reversePrimer` | `string` | Yes | Reverse primer, 5'→3'. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PrimerSpecificity():create({
  forwardPrimer = --[[ string ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  reversePrimer = --[[ string ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | `number` | No | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | `number` | No | Cap on the number of returned peptides. |
| `minMass` | `number` | No | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | `number` | No | Allowed missed internal cleavages (0–2). |
| `ok` | `any` | Yes |  |
| `protease` | `string` | No | Protease or chemical cleavage agent. |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ProteinAnnotatePoll():create({
  jobId = --[[ string ]],
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
| `appl` | `string` | No | Restrict to one member database (e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | `boolean` | No | Include GO-term cross-references. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `scale` | `string` | No | Amino-acid scale. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `window` | `number` | No | Sliding-window size (clamped to an odd number ≥ 1). |

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
| `chargeStep` | `number` | No | pH step for the net-charge titration curve (0–14). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | `number` | No | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` | `string` | No |  |
| `length` | `number` | Yes | Number of residues to generate. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `enzymes` | `table` | No | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No | Codon-usage host (ignored in degenerate mode). |
| `protein` | `string` | Yes | Protein sequence (one-letter codes; * for stop). |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `fileBase64` | `string` | No | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | No | Optional original file name (echoed back). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | `number` | No | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `read` | `string` | No | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | `string` | Yes | Expected reference sequence (FASTA or raw). |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `args` | `table` | Yes | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SavePermalink():create({
  args = --[[ table ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `qualityOffset` | `number` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `accession` | `string` | Yes | GenBank/RefSeq accession (e.g. |
| `db` | `string` | No | Database to query; auto-detects from the accession format. |
| `format` | `string` | No | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `from` | `string` | No | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | A FASTA or GenBank record to convert. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `to` | `string` | No | Output format. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `endPrimerLength` | `number` | No | Length of the naive end primers taken from each end. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `number` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `number` | No | Minimum ORF length in amino acids. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | No | Gene symbol/name, e.g. |
| `maxResults` | `number` | No | Up to 20. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No | Organism name, e.g. |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `term` | `string` | No | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | `number` | No | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `reads` | `string` | Yes | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | `string` | Yes | The claimed/expected reference sequence. |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SequencingReadbackVerify():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  reads = --[[ string ]],
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
| `entries` | `table` | No | Initial named entries, e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | `table` | No | Only return these entries; omit to return all of them. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SessionGet():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sessionId = --[[ string ]],
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
| `args` | `table` | No | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | `table` | No | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |
| `writeBack` | `table` | No | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SessionRun():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sessionId = --[[ string ]],
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
| `entries` | `table` | Yes | Named entries to add/overwrite, e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SessionSet():create({
  entries = --[[ table ]],
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  sessionId = --[[ string ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | `number` | No | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `shRnaLoop` | `string` | No | Loop sequence used when assembling the shRNA cassette. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `armTmTarget` | `number` | No | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | `number` | No | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | `string` | No | Edit at the nucleotide or amino-acid level. |
| `frameStart` | `number` | No | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `number` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | No | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | `string` | No | Replacement base (editKind='nt'). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `number` | No | Total strand concentration (nM). |
| `organism` | `string` | No | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | `number` | No | 1-based position to substitute (editKind='nt'). |
| `provenance` | `table` | Yes |  |
| `residue` | `number` | No | 1-based residue number to change (editKind='aa'). |
| `result` | `table` | Yes | Tool-specific output object. |
| `style` | `string` | No | Mutagenic primer style. |
| `targetAa` | `string` | No | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | `boolean` | No | Stop at the first stop codon. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `assembly` | `string` | No | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `variant` | `string` | Yes | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

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
| `coding` | `boolean` | No | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | `number` | No | 1-based reading-frame start (used when coding is true). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `query` | `string` | Yes | Query / variant sequence (raw or FASTA). |
| `reference` | `string` | Yes | Reference / wild-type sequence (raw or FASTA). |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `armTmTarget` | `number` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `boolean` | No | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | `string` | Yes | The sequence you claim you ended up with. |
| `coding` | `boolean` | No | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | `string` | No | Type IIS enzyme for Golden Gate. |
| `enzyme3` | `string` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | No | 5′ enzyme (restriction method). |
| `fragmentPcrs` | `table` | No | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | `table` | No | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | `number` | No | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | No | Insert sequence (restriction method). |
| `insertPcr` | `table` | No | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | `string` | Yes | Assembly method used. |
| `names` | `table` | No | Optional labels for each fragment. |
| `ok` | `any` | Yes |  |
| `overlapLen` | `number` | No | Gibson homology-arm length (bp). |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `vector` | `string` | No | Vector sequence (restriction method). |
| `vectorPcr` | `table` | No | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:VerifyAssembly():create({
  claimedConstruct = --[[ string ]],
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
| `claimedConstruct` | `string` | Yes | The final sequence claimed to have been built. |
| `expectedFrameStart` | `number` | No | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | `string` | Yes | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | `string` | Yes | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | `string` | Yes | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | `number` | No | Mismatches tolerated per primer during PCR prediction. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `templateCircular` | `boolean` | No | Treat insertTemplate as circular (e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:VerifyConstruct():create({
  claimedConstruct = --[[ string ]],
  insertForwardPrimer = --[[ string ]],
  insertReversePrimer = --[[ string ]],
  insertTemplate = --[[ string ]],
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
| `circular` | `boolean` | No | Treat the sequence as circular (plasmid). |
| `enzymes` | `table` | No | Enzyme names to digest with. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | `string` | No | DNA ladder to plot alongside the sample lane. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `result` | `table` | Yes | Tool-specific output object. |
| `rows` | `table` | Yes | Differential expression rows, one per gene. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:VolcanoPlotData():create({
  ok = --[[ any ]],
  provenance = --[[ table ]],
  result = --[[ table ]],
  rows = --[[ table ]],
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | `number` | No | Maximum number of results to return (default 5, max 10). |
| `ok` | `any` | Yes |  |
| `provenance` | `table` | Yes |  |
| `query` | `string` | Yes | The search query. |
| `result` | `table` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

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


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

