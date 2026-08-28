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
| `accession` | `String` | Yes | UniProt accession, e.g. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `Integer` | No | Total gapmer length (nt). |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `target` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |
| `wing` | `Integer` | No | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

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
| `editor` | `String` | No | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | `Integer` | No | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `target` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | `Integer` | No | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `args` | `Hash` | No | Shared tool arguments applied to every record. |
| `capped` | `Boolean` | Yes | True if input exceeded the record limit. |
| `columns` | `Array` | Yes |  |
| `count` | `Integer` | Yes |  |
| `errors` | `Integer` | Yes |  |
| `input` | `String` | Yes | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | `Integer` | Yes | Maximum records per call (500). |
| `provenance` | `Hash` | Yes |  |
| `rows` | `Array` | Yes |  |
| `tool` | `String` | Yes | A batchable tool slug (see `GET /batch`). |

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
| `columns` | `Array` | Yes | Flattened "<step>·<tool>·<key>" column headers. |
| `count` | `Integer` | Yes |  |
| `errors` | `Integer` | Yes |  |
| `input` | `String` | Yes | Multi-FASTA text or one sequence per line. |
| `limit` | `Integer` | Yes | Maximum records per call (200). |
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
| `endPrimerLength` | `Integer` | No | Length of the naive end primers taken from each end. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `Integer` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `Integer` | No | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `armTmTarget` | `Float` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `Boolean` | No | Produce a circular product. |
| `enzyme` | `String` | No | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | `String` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `String` | No | 5′ enzyme (restriction method). |
| `fragments` | `Array` | No | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `String` | No | Insert sequence (restriction method). |
| `method` | `String` | Yes | Assembly method. |
| `names` | `Array` | No | Optional labels for each fragment. |
| `ok` | `Object` | Yes |  |
| `overlapLen` | `Integer` | No | Gibson homology-arm length (bp). |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |
| `vector` | `String` | No | Vector sequence (restriction method). |

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
| `frameStart` | `Integer` | No | 1-based position to start reading codons. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `organism` | `String` | No |  |
| `provenance` | `Hash` | Yes |  |
| `rareThreshold` | `Float` | No | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `organism` | `String` | No |  |
| `protein` | `String` | Yes | Protein sequence (one-letter codes). |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `avoidEnzymes` | `Array` | No | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | `Integer` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `Integer` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `Float` | No |  |
| `gcLow` | `Float` | No |  |
| `gcWindow` | `Integer` | No |  |
| `homopolymerMin` | `Integer` | No |  |
| `maxPasses` | `Integer` | No | Repeat full passes until clean or no further progress. |
| `ok` | `Object` | Yes |  |
| `organism` | `String` | No | Codon-usage table to prefer among synonymous options. |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `avoidEnzymes` | `Array` | No | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | `Integer` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `Integer` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `Float` | No | GC% above this flags a GC-rich window. |
| `gcLow` | `Float` | No | GC% below this flags an AT-rich window. |
| `gcWindow` | `Integer` | No | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | `Integer` | No | Minimum run length to flag a homopolymer. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | `Float` | No | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | `String` | No | Nuclease id. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `Boolean` | No | Also scan the reverse strand for guides. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `armLength` | `Integer` | No | Homology arm length (bp) on each side. |
| `blockPam` | `Boolean` | No | When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut. |
| `designGenotypingPrimers` | `Boolean` | No | Also design a primer pair (on the original targetSequence) whose product spans the edit site. |
| `editEnd` | `Integer` | No | 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. |
| `editStart` | `Integer` | No | 1-based start of the region being replaced. |
| `frameStart` | `Integer` | No | Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | `Integer` | No | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | `Integer` | No | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | `String` | No | Strand the guide's protospacer is on. |
| `nuclease` | `String` | No | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `replacement` | `String` | Yes | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `targetSequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `Integer` | No | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | `String` | No | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` | `Object` | Yes |  |
| `protospacer` | `String` | Yes | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequenceA` | `String` | Yes | First oligo (5'→3'). |
| `sequenceB` | `String` | Yes | Second oligo (5'→3'). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `Integer` | No | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | `Float` | No | Mass in nanograms. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | No | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | `String` | Yes | The tool slug that ran. |
| `type` | `String` | No | Molecule type. |
| `volumeUl` | `Float` | No | Volume in microlitres (0 = unknown; needed for concentration). |

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
| `enzymeA` | `String` | Yes | First enzyme name (e.g. |
| `enzymeB` | `String` | Yes | Second enzyme name (e.g. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `reactions` | `Array` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `protocolName` | `String` | No | Optional protocol name (used in the script's metadata). |
| `provenance` | `Hash` | Yes |  |
| `reactions` | `Array` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `reactions` | `Array` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `clusterCols` | `Boolean` | No | Cluster (reorder) samples. |
| `clusterRows` | `Boolean` | No | Cluster (reorder) genes. |
| `distanceMetric` | `String` | No | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `Array` | Yes | Row (gene) labels. |
| `linkage` | `String` | No | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `samples` | `Array` | Yes | Column (sample) labels. |
| `tool` | `String` | Yes | The tool slug that ran. |
| `values` | `Array` | Yes | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | `Boolean` | No | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `String` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `qualityOffset` | `Integer` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `String` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | `Integer` | No | Reads shorter than this after trimming are dropped. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `qualityOffset` | `Integer` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | `Integer` | No | 3' quality-trim threshold (Phred score). |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | `Integer` | No | Minimum protein length (aa) to report. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `requireStop` | `Boolean` | No | Only report ORFs terminated by a stop codon. |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `convert` | `String` | No | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `reverse` | `Boolean` | No | Reverse the sequence (no complement). |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | `Boolean` | No | Remove digits, spaces and gaps (keep letters only). |
| `tool` | `String` | Yes | The tool slug that ran. |
| `width` | `Integer` | No | Line-wrap width; 0 = single line. |

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
| `background` | `Array` | No | Custom background/universe gene symbols. |
| `collections` | `Array` | No | Which term collections to test. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `Array` | Yes | Query gene symbols (human, e.g. |
| `maxTermSize` | `Integer` | No | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | `Integer` | No | Skip terms/pathways with fewer than this many background genes. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `String` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `String` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `String` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `compareToNamedSet` | `String` | No | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | `String` | No | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `overhangs` | `Array` | Yes | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `riskThreshold` | `Float` | No | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |
| `variant` | `String` | Yes | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `from` | `String` | Yes | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | `Array` | Yes | The ids to map, up to 1000 (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `taxId` | `String` | No | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | `String` | Yes | Target id type. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `circular` | `Boolean` | No | Treat the template as circular (plasmid). |
| `forwardPrimer` | `String` | Yes | Primer 1, 5'→3'. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `Integer` | No | Mismatches tolerated per primer. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `reversePrimer` | `String` | Yes | Primer 2, 5'→3' (order does not matter). |
| `template` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `addSecondaryMismatch` | `Boolean` | No | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | `String` | Yes | First allele (single base) — gets the FAM tail. |
| `alleleB` | `String` | Yes | Second allele (single base) — gets the HEX tail. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | `Integer` | No | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | `Integer` | No | Minimum amplicon length for the common reverse primer. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `snpPosition` | `Integer` | Yes | 1-based position of the SNP on the forward strand. |
| `target` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | `Float` | No | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `dntpMM` | `Float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `Float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `Float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Object` | Yes |  |
| `oligoNM` | `Float` | No | Total strand concentration (nM). |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | `Float` | No | Optional target Tm (°C). |
| `tmTolerance` | `Float` | No | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `Integer` | No | Maximum allowed mismatches per match. |
| `motif` | `String` | Yes | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `Boolean` | No | Also search the reverse strand. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `String` | Yes | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `dntpMM` | `Float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `Float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `Float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Object` | Yes |  |
| `oligoNM` | `Float` | No | Total strand concentration (nM). |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sourceSpecies` | `String` | No | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | `Array` | Yes | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | `String` | Yes | Ensembl species slug to find homologs in (e.g. |
| `tool` | `String` | Yes | The tool slug that ran. |
| `type` | `String` | No | Homology type to return. |

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
| `gap` | `Float` | No | Linear gap penalty (per gap position). |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | `Float` | No | Match score. |
| `mismatch` | `Float` | No | Mismatch penalty. |
| `mode` | `String` | No |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `seqA` | `String` | Yes | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | `String` | Yes | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `text` | `String` | Yes | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `fileBase64` | `String` | Yes | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `String` | No | Optional original file name (echoed back). |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `circular` | `Boolean` | No | Treat the sequence as a circular plasmid (vs. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `circular` | `Boolean` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |
| `topN` | `Integer` | No | How many top-ranked backbone candidates to report. |

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
| `circular` | `Boolean` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |
| `topN` | `Integer` | No | How many top-ranked backbone candidates to report. |

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
| `editEnd` | `Integer` | Yes | 1-based inclusive end of the region being changed. |
| `editStart` | `Integer` | Yes | 1-based inclusive start of the region being changed. |
| `frameStart` | `Integer` | No | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | `String` | No | Replacement bases (forward strand). |
| `ok` | `Object` | Yes |  |
| `pbsLength` | `Integer` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `rttHomology` | `Integer` | No | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | `String` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | `String` | Yes | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` | `Object` | Yes |  |
| `overlapLength` | `Integer` | No | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | `Integer` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `Hash` | Yes |  |
| `replaceEnd` | `Integer` | Yes | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | `Integer` | Yes | 1-based inclusive start of the region being replaced/deleted. |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `target` | `String` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `dntpMM` | `Float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` | `Float` | No |  |
| `gcMin` | `Float` | No |  |
| `lenMax` | `Integer` | No |  |
| `lenMin` | `Integer` | No |  |
| `lenOpt` | `Integer` | No |  |
| `maxReturn` | `Integer` | No | Number of best pairs to return. |
| `mgMM` | `Float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `Float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Object` | Yes |  |
| `oligoNM` | `Float` | No | Total strand concentration (nM). |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `targetEnd` | `Integer` | No | 1-based inclusive end of the target region (optional). |
| `targetStart` | `Integer` | No | 1-based inclusive start of a region the product must span (optional). |
| `template` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` | `Float` | No |  |
| `tmMaxDiff` | `Float` | No | Max Tm difference within a pair (°C). |
| `tmMin` | `Float` | No |  |
| `tmOpt` | `Float` | No |  |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `forwardPrimer` | `String` | Yes | Forward primer, 5'→3'. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `Integer` | No | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | `Integer` | No | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `reversePrimer` | `String` | Yes | Reverse primer, 5'→3'. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | `Float` | No | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | `Integer` | No | Cap on the number of returned peptides. |
| `minMass` | `Float` | No | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | `Integer` | No | Allowed missed internal cleavages (0–2). |
| `ok` | `Object` | Yes |  |
| `protease` | `String` | No | Protease or chemical cleavage agent. |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `String` | Yes |  |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `appl` | `String` | No | Restrict to one member database (e.g. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | `Boolean` | No | Include GO-term cross-references. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `scale` | `String` | No | Amino-acid scale. |
| `sequence` | `String` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `String` | Yes | The tool slug that ran. |
| `window` | `Integer` | No | Sliding-window size (clamped to an odd number ≥ 1). |

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
| `chargeStep` | `Float` | No | pH step for the net-charge titration curve (0–14). |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | `Float` | No | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` | `String` | No |  |
| `length` | `Integer` | Yes | Number of residues to generate. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `enzymes` | `Array` | No | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |
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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` | `String` | No |  |
| `ok` | `Object` | Yes |  |
| `organism` | `String` | No | Codon-usage host (ignored in degenerate mode). |
| `protein` | `String` | Yes | Protein sequence (one-letter codes; * for stop). |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `fileBase64` | `String` | No | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `String` | No | Optional original file name (echoed back). |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | `Float` | No | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `read` | `String` | No | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | `String` | Yes | Expected reference sequence (FASTA or raw). |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `args` | `Hash` | Yes | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `String` | Yes | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `qualityOffset` | `Integer` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `accession` | `String` | Yes | GenBank/RefSeq accession (e.g. |
| `db` | `String` | No | Database to query; auto-detects from the accession format. |
| `format` | `String` | No | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `from` | `String` | No | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `String` | Yes | A FASTA or GenBank record to convert. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `to` | `String` | No | Output format. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `endPrimerLength` | `Integer` | No | Length of the naive end primers taken from each end. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `Integer` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `Integer` | No | Minimum ORF length in amino acids. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `String` | No | Gene symbol/name, e.g. |
| `maxResults` | `Integer` | No | Up to 20. |
| `ok` | `Object` | Yes |  |
| `organism` | `String` | No | Organism name, e.g. |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `term` | `String` | No | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | `Integer` | No | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `reads` | `String` | Yes | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | `String` | Yes | The claimed/expected reference sequence. |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `entries` | `Hash` | No | Initial named entries, e.g. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | `Array` | No | Only return these entries; omit to return all of them. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sessionId` | `String` | Yes |  |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `args` | `Hash` | No | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | `Hash` | No | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sessionId` | `String` | Yes |  |
| `tool` | `String` | Yes | The tool slug that ran. |
| `writeBack` | `Hash` | No | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

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
| `entries` | `Hash` | Yes | Named entries to add/overwrite, e.g. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sessionId` | `String` | Yes |  |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | `Integer` | No | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `shRnaLoop` | `String` | No | Loop sequence used when assembling the shRNA cassette. |
| `target` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `armTmTarget` | `Float` | No | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | `Float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | `String` | No | Edit at the nucleotide or amino-acid level. |
| `frameStart` | `Integer` | No | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `Float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `Float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | `String` | No | Replacement base (editKind='nt'). |
| `ok` | `Object` | Yes |  |
| `oligoNM` | `Float` | No | Total strand concentration (nM). |
| `organism` | `String` | No | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | `Integer` | No | 1-based position to substitute (editKind='nt'). |
| `provenance` | `Hash` | Yes |  |
| `residue` | `Integer` | No | 1-based residue number to change (editKind='aa'). |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `style` | `String` | No | Mutagenic primer style. |
| `targetAa` | `String` | No | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | `Boolean` | No | Stop at the first stop codon. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `assembly` | `String` | No | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |
| `variant` | `String` | Yes | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

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
| `coding` | `Boolean` | No | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | `Integer` | No | 1-based reading-frame start (used when coding is true). |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `query` | `String` | Yes | Query / variant sequence (raw or FASTA). |
| `reference` | `String` | Yes | Reference / wild-type sequence (raw or FASTA). |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `armTmTarget` | `Float` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `Boolean` | No | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | `String` | Yes | The sequence you claim you ended up with. |
| `coding` | `Boolean` | No | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | `String` | No | Type IIS enzyme for Golden Gate. |
| `enzyme3` | `String` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `String` | No | 5′ enzyme (restriction method). |
| `fragmentPcrs` | `Array` | No | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | `Array` | No | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | `Integer` | No | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `String` | No | Insert sequence (restriction method). |
| `insertPcr` | `Hash` | No | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | `String` | Yes | Assembly method used. |
| `names` | `Array` | No | Optional labels for each fragment. |
| `ok` | `Object` | Yes |  |
| `overlapLen` | `Integer` | No | Gibson homology-arm length (bp). |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |
| `vector` | `String` | No | Vector sequence (restriction method). |
| `vectorPcr` | `Hash` | No | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

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
| `claimedConstruct` | `String` | Yes | The final sequence claimed to have been built. |
| `expectedFrameStart` | `Integer` | No | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | `String` | Yes | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | `String` | Yes | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | `String` | Yes | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | `Integer` | No | Mismatches tolerated per primer during PCR prediction. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `templateCircular` | `Boolean` | No | Treat insertTemplate as circular (e.g. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `circular` | `Boolean` | No | Treat the sequence as circular (plasmid). |
| `enzymes` | `Array` | No | Enzyme names to digest with. |
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | `String` | No | DNA ladder to plot alongside the sample lane. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `sequence` | `String` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `rows` | `Array` | Yes | Differential expression rows, one per gene. |
| `tool` | `String` | Yes | The tool slug that ran. |

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
| `gate` | `Object` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | `Float` | No | Maximum number of results to return (default 5, max 10). |
| `ok` | `Object` | Yes |  |
| `provenance` | `Hash` | Yes |  |
| `query` | `String` | Yes | The search query. |
| `result` | `Hash` | Yes | Tool-specific output object. |
| `tool` | `String` | Yes | The tool slug that ran. |

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

