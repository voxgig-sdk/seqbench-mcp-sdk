# SeqbenchMcp Golang SDK Reference

Complete API reference for the SeqbenchMcp Golang SDK.


## SeqbenchMcpSDK

### Constructor

```go
func NewSeqbenchMcpSDK(options map[string]any) *SeqbenchMcpSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *SeqbenchMcpSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *SeqbenchMcpSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `AlphafoldLookup(data map[string]any) SeqbenchMcpEntity`

Create a new `AlphafoldLookup` entity instance. Pass `nil` for no initial data.

#### `AsoDesign(data map[string]any) SeqbenchMcpEntity`

Create a new `AsoDesign` entity instance. Pass `nil` for no initial data.

#### `BaseEditingDesign(data map[string]any) SeqbenchMcpEntity`

Create a new `BaseEditingDesign` entity instance. Pass `nil` for no initial data.

#### `Batch(data map[string]any) SeqbenchMcpEntity`

Create a new `Batch` entity instance. Pass `nil` for no initial data.

#### `BatchWorkflow(data map[string]any) SeqbenchMcpEntity`

Create a new `BatchWorkflow` entity instance. Pass `nil` for no initial data.

#### `CharacterizeSequence(data map[string]any) SeqbenchMcpEntity`

Create a new `CharacterizeSequence` entity instance. Pass `nil` for no initial data.

#### `CloningSimulate(data map[string]any) SeqbenchMcpEntity`

Create a new `CloningSimulate` entity instance. Pass `nil` for no initial data.

#### `CodonAdaptationIndex(data map[string]any) SeqbenchMcpEntity`

Create a new `CodonAdaptationIndex` entity instance. Pass `nil` for no initial data.

#### `CodonOptimize(data map[string]any) SeqbenchMcpEntity`

Create a new `CodonOptimize` entity instance. Pass `nil` for no initial data.

#### `ConstructAutofix(data map[string]any) SeqbenchMcpEntity`

Create a new `ConstructAutofix` entity instance. Pass `nil` for no initial data.

#### `ConstructQc(data map[string]any) SeqbenchMcpEntity`

Create a new `ConstructQc` entity instance. Pass `nil` for no initial data.

#### `CrisprGrnaDesign(data map[string]any) SeqbenchMcpEntity`

Create a new `CrisprGrnaDesign` entity instance. Pass `nil` for no initial data.

#### `CrisprHdrDonor(data map[string]any) SeqbenchMcpEntity`

Create a new `CrisprHdrDonor` entity instance. Pass `nil` for no initial data.

#### `CrisprOfftargetCheck(data map[string]any) SeqbenchMcpEntity`

Create a new `CrisprOfftargetCheck` entity instance. Pass `nil` for no initial data.

#### `CrossDimer(data map[string]any) SeqbenchMcpEntity`

Create a new `CrossDimer` entity instance. Pass `nil` for no initial data.

#### `DnaMolarity(data map[string]any) SeqbenchMcpEntity`

Create a new `DnaMolarity` entity instance. Pass `nil` for no initial data.

#### `DoubleDigest(data map[string]any) SeqbenchMcpEntity`

Create a new `DoubleDigest` entity instance. Pass `nil` for no initial data.

#### `ExportEchoPicklist(data map[string]any) SeqbenchMcpEntity`

Create a new `ExportEchoPicklist` entity instance. Pass `nil` for no initial data.

#### `ExportOpentronsProtocol(data map[string]any) SeqbenchMcpEntity`

Create a new `ExportOpentronsProtocol` entity instance. Pass `nil` for no initial data.

#### `ExportPlateLayout(data map[string]any) SeqbenchMcpEntity`

Create a new `ExportPlateLayout` entity instance. Pass `nil` for no initial data.

#### `ExpressionHeatmapCluster(data map[string]any) SeqbenchMcpEntity`

Create a new `ExpressionHeatmapCluster` entity instance. Pass `nil` for no initial data.

#### `FastqQcReport(data map[string]any) SeqbenchMcpEntity`

Create a new `FastqQcReport` entity instance. Pass `nil` for no initial data.

#### `FastqTrim(data map[string]any) SeqbenchMcpEntity`

Create a new `FastqTrim` entity instance. Pass `nil` for no initial data.

#### `FindOrf(data map[string]any) SeqbenchMcpEntity`

Create a new `FindOrf` entity instance. Pass `nil` for no initial data.

#### `FormatSequence(data map[string]any) SeqbenchMcpEntity`

Create a new `FormatSequence` entity instance. Pass `nil` for no initial data.

#### `FunctionalEnrichment(data map[string]any) SeqbenchMcpEntity`

Create a new `FunctionalEnrichment` entity instance. Pass `nil` for no initial data.

#### `GcContent(data map[string]any) SeqbenchMcpEntity`

Create a new `GcContent` entity instance. Pass `nil` for no initial data.

#### `GeneDossier(data map[string]any) SeqbenchMcpEntity`

Create a new `GeneDossier` entity instance. Pass `nil` for no initial data.

#### `GeneExpression(data map[string]any) SeqbenchMcpEntity`

Create a new `GeneExpression` entity instance. Pass `nil` for no initial data.

#### `GeneModel(data map[string]any) SeqbenchMcpEntity`

Create a new `GeneModel` entity instance. Pass `nil` for no initial data.

#### `GoldenGateFidelity(data map[string]any) SeqbenchMcpEntity`

Create a new `GoldenGateFidelity` entity instance. Pass `nil` for no initial data.

#### `HgvsConvert(data map[string]any) SeqbenchMcpEntity`

Create a new `HgvsConvert` entity instance. Pass `nil` for no initial data.

#### `IdMapPoll(data map[string]any) SeqbenchMcpEntity`

Create a new `IdMapPoll` entity instance. Pass `nil` for no initial data.

#### `IdMapSubmit(data map[string]any) SeqbenchMcpEntity`

Create a new `IdMapSubmit` entity instance. Pass `nil` for no initial data.

#### `InSilicoPcr(data map[string]any) SeqbenchMcpEntity`

Create a new `InSilicoPcr` entity instance. Pass `nil` for no initial data.

#### `KaspPrimerDesign(data map[string]any) SeqbenchMcpEntity`

Create a new `KaspPrimerDesign` entity instance. Pass `nil` for no initial data.

#### `ListTool(data map[string]any) SeqbenchMcpEntity`

Create a new `ListTool` entity instance. Pass `nil` for no initial data.

#### `MeltingTemperature(data map[string]any) SeqbenchMcpEntity`

Create a new `MeltingTemperature` entity instance. Pass `nil` for no initial data.

#### `MotifFinder(data map[string]any) SeqbenchMcpEntity`

Create a new `MotifFinder` entity instance. Pass `nil` for no initial data.

#### `MultipleSequenceAlignment(data map[string]any) SeqbenchMcpEntity`

Create a new `MultipleSequenceAlignment` entity instance. Pass `nil` for no initial data.

#### `OligoAnalysi(data map[string]any) SeqbenchMcpEntity`

Create a new `OligoAnalysi` entity instance. Pass `nil` for no initial data.

#### `OrthologMap(data map[string]any) SeqbenchMcpEntity`

Create a new `OrthologMap` entity instance. Pass `nil` for no initial data.

#### `PairwiseAlignment(data map[string]any) SeqbenchMcpEntity`

Create a new `PairwiseAlignment` entity instance. Pass `nil` for no initial data.

#### `ParseGenbank(data map[string]any) SeqbenchMcpEntity`

Create a new `ParseGenbank` entity instance. Pass `nil` for no initial data.

#### `ParseSangerTrace(data map[string]any) SeqbenchMcpEntity`

Create a new `ParseSangerTrace` entity instance. Pass `nil` for no initial data.

#### `PlasmidAnnotate(data map[string]any) SeqbenchMcpEntity`

Create a new `PlasmidAnnotate` entity instance. Pass `nil` for no initial data.

#### `PlasmidDeepAnnotate(data map[string]any) SeqbenchMcpEntity`

Create a new `PlasmidDeepAnnotate` entity instance. Pass `nil` for no initial data.

#### `PlasmidFullReport(data map[string]any) SeqbenchMcpEntity`

Create a new `PlasmidFullReport` entity instance. Pass `nil` for no initial data.

#### `PlasmidIdentify(data map[string]any) SeqbenchMcpEntity`

Create a new `PlasmidIdentify` entity instance. Pass `nil` for no initial data.

#### `PrimeEditingDesign(data map[string]any) SeqbenchMcpEntity`

Create a new `PrimeEditingDesign` entity instance. Pass `nil` for no initial data.

#### `PrimeEditingTwinDesign(data map[string]any) SeqbenchMcpEntity`

Create a new `PrimeEditingTwinDesign` entity instance. Pass `nil` for no initial data.

#### `PrimerDesign(data map[string]any) SeqbenchMcpEntity`

Create a new `PrimerDesign` entity instance. Pass `nil` for no initial data.

#### `PrimerSpecificity(data map[string]any) SeqbenchMcpEntity`

Create a new `PrimerSpecificity` entity instance. Pass `nil` for no initial data.

#### `ProteaseDigestion(data map[string]any) SeqbenchMcpEntity`

Create a new `ProteaseDigestion` entity instance. Pass `nil` for no initial data.

#### `ProteinAnnotatePoll(data map[string]any) SeqbenchMcpEntity`

Create a new `ProteinAnnotatePoll` entity instance. Pass `nil` for no initial data.

#### `ProteinAnnotateSubmit(data map[string]any) SeqbenchMcpEntity`

Create a new `ProteinAnnotateSubmit` entity instance. Pass `nil` for no initial data.

#### `ProteinHydrophobicity(data map[string]any) SeqbenchMcpEntity`

Create a new `ProteinHydrophobicity` entity instance. Pass `nil` for no initial data.

#### `ProteinProperty(data map[string]any) SeqbenchMcpEntity`

Create a new `ProteinProperty` entity instance. Pass `nil` for no initial data.

#### `RandomSequence(data map[string]any) SeqbenchMcpEntity`

Create a new `RandomSequence` entity instance. Pass `nil` for no initial data.

#### `RestrictionSite(data map[string]any) SeqbenchMcpEntity`

Create a new `RestrictionSite` entity instance. Pass `nil` for no initial data.

#### `ReverseComplement(data map[string]any) SeqbenchMcpEntity`

Create a new `ReverseComplement` entity instance. Pass `nil` for no initial data.

#### `ReverseTranslate(data map[string]any) SeqbenchMcpEntity`

Create a new `ReverseTranslate` entity instance. Pass `nil` for no initial data.

#### `RnaFold(data map[string]any) SeqbenchMcpEntity`

Create a new `RnaFold` entity instance. Pass `nil` for no initial data.

#### `SangerVsReference(data map[string]any) SeqbenchMcpEntity`

Create a new `SangerVsReference` entity instance. Pass `nil` for no initial data.

#### `SavePermalink(data map[string]any) SeqbenchMcpEntity`

Create a new `SavePermalink` entity instance. Pass `nil` for no initial data.

#### `SeqfileStat(data map[string]any) SeqbenchMcpEntity`

Create a new `SeqfileStat` entity instance. Pass `nil` for no initial data.

#### `SequenceFetch(data map[string]any) SeqbenchMcpEntity`

Create a new `SequenceFetch` entity instance. Pass `nil` for no initial data.

#### `SequenceFormatConvert(data map[string]any) SeqbenchMcpEntity`

Create a new `SequenceFormatConvert` entity instance. Pass `nil` for no initial data.

#### `SequenceReport(data map[string]any) SeqbenchMcpEntity`

Create a new `SequenceReport` entity instance. Pass `nil` for no initial data.

#### `SequenceSearch(data map[string]any) SeqbenchMcpEntity`

Create a new `SequenceSearch` entity instance. Pass `nil` for no initial data.

#### `SequencingReadbackVerify(data map[string]any) SeqbenchMcpEntity`

Create a new `SequencingReadbackVerify` entity instance. Pass `nil` for no initial data.

#### `SessionCreate(data map[string]any) SeqbenchMcpEntity`

Create a new `SessionCreate` entity instance. Pass `nil` for no initial data.

#### `SessionGet(data map[string]any) SeqbenchMcpEntity`

Create a new `SessionGet` entity instance. Pass `nil` for no initial data.

#### `SessionRun(data map[string]any) SeqbenchMcpEntity`

Create a new `SessionRun` entity instance. Pass `nil` for no initial data.

#### `SessionSet(data map[string]any) SeqbenchMcpEntity`

Create a new `SessionSet` entity instance. Pass `nil` for no initial data.

#### `SirnaDesign(data map[string]any) SeqbenchMcpEntity`

Create a new `SirnaDesign` entity instance. Pass `nil` for no initial data.

#### `SiteDirectedMutagenesi(data map[string]any) SeqbenchMcpEntity`

Create a new `SiteDirectedMutagenesi` entity instance. Pass `nil` for no initial data.

#### `Translate(data map[string]any) SeqbenchMcpEntity`

Create a new `Translate` entity instance. Pass `nil` for no initial data.

#### `VariantAnnotate(data map[string]any) SeqbenchMcpEntity`

Create a new `VariantAnnotate` entity instance. Pass `nil` for no initial data.

#### `VariantComparator(data map[string]any) SeqbenchMcpEntity`

Create a new `VariantComparator` entity instance. Pass `nil` for no initial data.

#### `VerifyAssembly(data map[string]any) SeqbenchMcpEntity`

Create a new `VerifyAssembly` entity instance. Pass `nil` for no initial data.

#### `VerifyConstruct(data map[string]any) SeqbenchMcpEntity`

Create a new `VerifyConstruct` entity instance. Pass `nil` for no initial data.

#### `VirtualGel(data map[string]any) SeqbenchMcpEntity`

Create a new `VirtualGel` entity instance. Pass `nil` for no initial data.

#### `VolcanoPlotData(data map[string]any) SeqbenchMcpEntity`

Create a new `VolcanoPlotData` entity instance. Pass `nil` for no initial data.

#### `WebSearch(data map[string]any) SeqbenchMcpEntity`

Create a new `WebSearch` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## AlphafoldLookupEntity

```go
alphafoldLookup := client.AlphafoldLookup(nil)
fmt.Println(alphafoldLookup.GetName()) // "alphafold_lookup"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `string` | Yes | UniProt accession, e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.AlphafoldLookup(nil).Create(map[string]any{
    "accession": "example_accession",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AlphafoldLookupEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## AsoDesignEntity

```go
asoDesign := client.AsoDesign(nil)
fmt.Println(asoDesign.GetName()) // "aso_design"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `int` | No | Total gapmer length (nt). |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `wing` | `int` | No | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.AsoDesign(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AsoDesignEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## BaseEditingDesignEntity

```go
baseEditingDesign := client.BaseEditingDesign(nil)
fmt.Println(baseEditingDesign.GetName()) // "base_editing_design"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `editor` | `string` | No | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | `int` | No | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | `int` | No | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.BaseEditingDesign(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BaseEditingDesignEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## BatchEntity

```go
batch := client.Batch(nil)
fmt.Println(batch.GetName()) // "batch"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `map[string]any` | No | Shared tool arguments applied to every record. |
| `capped` | `bool` | Yes | True if input exceeded the record limit. |
| `columns` | `[]any` | Yes |  |
| `count` | `int` | Yes |  |
| `errors` | `int` | Yes |  |
| `input` | `string` | Yes | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | `int` | Yes | Maximum records per call (500). |
| `provenance` | `map[string]any` | Yes |  |
| `rows` | `[]any` | Yes |  |
| `tool` | `string` | Yes | A batchable tool slug (see `GET /batch`). |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Batch(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Batch(nil).Create(map[string]any{
    "capped": true,
    "columns": []any{},
    "count": 1,
    "errors": 1,
    "input": "example_input",
    "limit": 1,
    "provenance": map[string]any{},
    "rows": []any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BatchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## BatchWorkflowEntity

```go
batchWorkflow := client.BatchWorkflow(nil)
fmt.Println(batchWorkflow.GetName()) // "batch__workflow"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `capped` | `bool` | Yes |  |
| `columns` | `[]any` | Yes | Flattened "<step>·<tool>·<key>" column headers. |
| `count` | `int` | Yes |  |
| `errors` | `int` | Yes |  |
| `input` | `string` | Yes | Multi-FASTA text or one sequence per line. |
| `limit` | `int` | Yes | Maximum records per call (200). |
| `provenance` | `map[string]any` | Yes |  |
| `rows` | `[]any` | Yes |  |
| `steps` | `[]any` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.BatchWorkflow(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.BatchWorkflow(nil).Create(map[string]any{
    "capped": true,
    "columns": []any{},
    "count": 1,
    "errors": 1,
    "input": "example_input",
    "limit": 1,
    "provenance": map[string]any{},
    "rows": []any{},
    "steps": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BatchWorkflowEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CharacterizeSequenceEntity

```go
characterizeSequence := client.CharacterizeSequence(nil)
fmt.Println(characterizeSequence.GetName()) // "characterize_sequence"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `endPrimerLength` | `int` | No | Length of the naive end primers taken from each end. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `int` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `int` | No | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CharacterizeSequence(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CharacterizeSequenceEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CloningSimulateEntity

```go
cloningSimulate := client.CloningSimulate(nil)
fmt.Println(cloningSimulate.GetName()) // "cloning_simulate"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `float64` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `bool` | No | Produce a circular product. |
| `enzyme` | `string` | No | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | `string` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | No | 5′ enzyme (restriction method). |
| `fragments` | `[]any` | No | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | No | Insert sequence (restriction method). |
| `method` | `string` | Yes | Assembly method. |
| `names` | `[]any` | No | Optional labels for each fragment. |
| `ok` | `any` | Yes |  |
| `overlapLen` | `int` | No | Gibson homology-arm length (bp). |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `vector` | `string` | No | Vector sequence (restriction method). |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CloningSimulate(nil).Create(map[string]any{
    "method": "example_method",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CloningSimulateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CodonAdaptationIndexEntity

```go
codonAdaptationIndex := client.CodonAdaptationIndex(nil)
fmt.Println(codonAdaptationIndex.GetName()) // "codon_adaptation_index"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frameStart` | `int` | No | 1-based position to start reading codons. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `rareThreshold` | `float64` | No | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CodonAdaptationIndex(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CodonAdaptationIndexEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CodonOptimizeEntity

```go
codonOptimize := client.CodonOptimize(nil)
fmt.Println(codonOptimize.GetName()) // "codon_optimize"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes | Protein sequence (one-letter codes). |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CodonOptimize(nil).Create(map[string]any{
    "ok": "example_ok",
    "protein": "example_protein",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CodonOptimizeEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ConstructAutofixEntity

```go
constructAutofix := client.ConstructAutofix(nil)
fmt.Println(constructAutofix.GetName()) // "construct_autofix"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoidEnzymes` | `[]any` | No | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | `int` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `int` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `float64` | No |  |
| `gcLow` | `float64` | No |  |
| `gcWindow` | `int` | No |  |
| `homopolymerMin` | `int` | No |  |
| `maxPasses` | `int` | No | Repeat full passes until clean or no further progress. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No | Codon-usage table to prefer among synonymous options. |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ConstructAutofix(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ConstructAutofixEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ConstructQcEntity

```go
constructQc := client.ConstructQc(nil)
fmt.Println(constructQc.GetName()) // "construct_qc"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoidEnzymes` | `[]any` | No | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | `int` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `int` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `float64` | No | GC% above this flags a GC-rich window. |
| `gcLow` | `float64` | No | GC% below this flags an AT-rich window. |
| `gcWindow` | `int` | No | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | `int` | No | Minimum run length to flag a homopolymer. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ConstructQc(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ConstructQcEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CrisprGrnaDesignEntity

```go
crisprGrnaDesign := client.CrisprGrnaDesign(nil)
fmt.Println(crisprGrnaDesign.GetName()) // "crispr_grna_design"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | `float64` | No | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | `string` | No | Nuclease id. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `bool` | No | Also scan the reverse strand for guides. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CrisprGrnaDesign(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CrisprGrnaDesignEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CrisprHdrDonorEntity

```go
crisprHdrDonor := client.CrisprHdrDonor(nil)
fmt.Println(crisprHdrDonor.GetName()) // "crispr_hdr_donor"
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
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | `int` | No | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | `int` | No | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | `string` | No | Strand the guide's protospacer is on. |
| `nuclease` | `string` | No | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `replacement` | `string` | Yes | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `targetSequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CrisprHdrDonor(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "replacement": "example_replacement",
    "result": map[string]any{},
    "targetSequence": "example_targetSequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CrisprHdrDonorEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CrisprOfftargetCheckEntity

```go
crisprOfftargetCheck := client.CrisprOfftargetCheck(nil)
fmt.Println(crisprOfftargetCheck.GetName()) // "crispr_offtarget_check"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | `string` | No | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` | `any` | Yes |  |
| `protospacer` | `string` | Yes | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CrisprOfftargetCheck(nil).Create(map[string]any{
    "ok": "example_ok",
    "protospacer": "example_protospacer",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CrisprOfftargetCheckEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CrossDimerEntity

```go
crossDimer := client.CrossDimer(nil)
fmt.Println(crossDimer.GetName()) // "cross_dimer"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequenceA` | `string` | Yes | First oligo (5'→3'). |
| `sequenceB` | `string` | Yes | Second oligo (5'→3'). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CrossDimer(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequenceA": "example_sequenceA",
    "sequenceB": "example_sequenceB",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CrossDimerEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## DnaMolarityEntity

```go
dnaMolarity := client.DnaMolarity(nil)
fmt.Println(dnaMolarity.GetName()) // "dna_molarity"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `int` | No | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | `float64` | No | Mass in nanograms. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | No | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `type` | `string` | No | Molecule type. |
| `volumeUl` | `float64` | No | Volume in microlitres (0 = unknown; needed for concentration). |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.DnaMolarity(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DnaMolarityEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## DoubleDigestEntity

```go
doubleDigest := client.DoubleDigest(nil)
fmt.Println(doubleDigest.GetName()) // "double_digest"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzymeA` | `string` | Yes | First enzyme name (e.g. |
| `enzymeB` | `string` | Yes | Second enzyme name (e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.DoubleDigest(nil).Create(map[string]any{
    "enzymeA": "example_enzymeA",
    "enzymeB": "example_enzymeB",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DoubleDigestEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ExportEchoPicklistEntity

```go
exportEchoPicklist := client.ExportEchoPicklist(nil)
fmt.Println(exportEchoPicklist.GetName()) // "export_echo_picklist"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `reactions` | `[]any` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ExportEchoPicklist(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reactions": []any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ExportEchoPicklistEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ExportOpentronsProtocolEntity

```go
exportOpentronsProtocol := client.ExportOpentronsProtocol(nil)
fmt.Println(exportOpentronsProtocol.GetName()) // "export_opentrons_protocol"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `protocolName` | `string` | No | Optional protocol name (used in the script's metadata). |
| `provenance` | `map[string]any` | Yes |  |
| `reactions` | `[]any` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ExportOpentronsProtocol(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reactions": []any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ExportOpentronsProtocolEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ExportPlateLayoutEntity

```go
exportPlateLayout := client.ExportPlateLayout(nil)
fmt.Println(exportPlateLayout.GetName()) // "export_plate_layout"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `reactions` | `[]any` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ExportPlateLayout(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reactions": []any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ExportPlateLayoutEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ExpressionHeatmapClusterEntity

```go
expressionHeatmapCluster := client.ExpressionHeatmapCluster(nil)
fmt.Println(expressionHeatmapCluster.GetName()) // "expression_heatmap_cluster"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clusterCols` | `bool` | No | Cluster (reorder) samples. |
| `clusterRows` | `bool` | No | Cluster (reorder) genes. |
| `distanceMetric` | `string` | No | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `[]any` | Yes | Row (gene) labels. |
| `linkage` | `string` | No | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `samples` | `[]any` | Yes | Column (sample) labels. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `values` | `[]any` | Yes | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | `bool` | No | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ExpressionHeatmapCluster(nil).Create(map[string]any{
    "genes": []any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "samples": []any{},
    "tool": "example_tool",
    "values": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ExpressionHeatmapClusterEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## FastqQcReportEntity

```go
fastqQcReport := client.FastqQcReport(nil)
fmt.Println(fastqQcReport.GetName()) // "fastq_qc_report"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `qualityOffset` | `int` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.FastqQcReport(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `FastqQcReportEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## FastqTrimEntity

```go
fastqTrim := client.FastqTrim(nil)
fmt.Println(fastqTrim.GetName()) // "fastq_trim"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | `int` | No | Reads shorter than this after trimming are dropped. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `qualityOffset` | `int` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | `int` | No | 3' quality-trim threshold (Phred score). |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.FastqTrim(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `FastqTrimEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## FindOrfEntity

```go
findOrf := client.FindOrf(nil)
fmt.Println(findOrf.GetName()) // "find_orf"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | `int` | No | Minimum protein length (aa) to report. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `requireStop` | `bool` | No | Only report ORFs terminated by a stop codon. |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.FindOrf(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `FindOrfEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## FormatSequenceEntity

```go
formatSequence := client.FormatSequence(nil)
fmt.Println(formatSequence.GetName()) // "format_sequence"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `caseMode` | `string` | No |  |
| `convert` | `string` | No | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `reverse` | `bool` | No | Reverse the sequence (no complement). |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | `bool` | No | Remove digits, spaces and gaps (keep letters only). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `width` | `int` | No | Line-wrap width; 0 = single line. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.FormatSequence(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `FormatSequenceEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## FunctionalEnrichmentEntity

```go
functionalEnrichment := client.FunctionalEnrichment(nil)
fmt.Println(functionalEnrichment.GetName()) // "functional_enrichment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `background` | `[]any` | No | Custom background/universe gene symbols. |
| `collections` | `[]any` | No | Which term collections to test. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `[]any` | Yes | Query gene symbols (human, e.g. |
| `maxTermSize` | `int` | No | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | `int` | No | Skip terms/pathways with fewer than this many background genes. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.FunctionalEnrichment(nil).Create(map[string]any{
    "genes": []any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `FunctionalEnrichmentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GcContentEntity

```go
gcContent := client.GcContent(nil)
fmt.Println(gcContent.GetName()) // "gc_content"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.GcContent(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GcContentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GeneDossierEntity

```go
geneDossier := client.GeneDossier(nil)
fmt.Println(geneDossier.GetName()) // "gene_dossier"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.GeneDossier(nil).Create(map[string]any{
    "gene": "example_gene",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GeneDossierEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GeneExpressionEntity

```go
geneExpression := client.GeneExpression(nil)
fmt.Println(geneExpression.GetName()) // "gene_expression"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.GeneExpression(nil).Create(map[string]any{
    "gene": "example_gene",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GeneExpressionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GeneModelEntity

```go
geneModel := client.GeneModel(nil)
fmt.Println(geneModel.GetName()) // "gene_model"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.GeneModel(nil).Create(map[string]any{
    "gene": "example_gene",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GeneModelEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GoldenGateFidelityEntity

```go
goldenGateFidelity := client.GoldenGateFidelity(nil)
fmt.Println(goldenGateFidelity.GetName()) // "golden_gate_fidelity"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `compareToNamedSet` | `string` | No | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | `string` | No | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `overhangs` | `[]any` | Yes | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `riskThreshold` | `float64` | No | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.GoldenGateFidelity(nil).Create(map[string]any{
    "ok": "example_ok",
    "overhangs": []any{},
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GoldenGateFidelityEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## HgvsConvertEntity

```go
hgvsConvert := client.HgvsConvert(nil)
fmt.Println(hgvsConvert.GetName()) // "hgvs_convert"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `variant` | `string` | Yes | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.HgvsConvert(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
    "variant": "example_variant",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `HgvsConvertEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## IdMapPollEntity

```go
idMapPoll := client.IdMapPoll(nil)
fmt.Println(idMapPoll.GetName()) // "id_map_poll"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.IdMapPoll(nil).Create(map[string]any{
    "jobId": "example_jobId",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `IdMapPollEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## IdMapSubmitEntity

```go
idMapSubmit := client.IdMapSubmit(nil)
fmt.Println(idMapSubmit.GetName()) // "id_map_submit"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | Yes | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | `[]any` | Yes | The ids to map, up to 1000 (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `taxId` | `string` | No | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | `string` | Yes | Target id type. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.IdMapSubmit(nil).Create(map[string]any{
    "from": "example_from",
    "ids": []any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "to": "example_to",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `IdMapSubmitEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## InSilicoPcrEntity

```go
inSilicoPcr := client.InSilicoPcr(nil)
fmt.Println(inSilicoPcr.GetName()) // "in_silico_pcr"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No | Treat the template as circular (plasmid). |
| `forwardPrimer` | `string` | Yes | Primer 1, 5'→3'. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Mismatches tolerated per primer. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `reversePrimer` | `string` | Yes | Primer 2, 5'→3' (order does not matter). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.InSilicoPcr(nil).Create(map[string]any{
    "forwardPrimer": "example_forwardPrimer",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "reversePrimer": "example_reversePrimer",
    "template": "example_template",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `InSilicoPcrEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## KaspPrimerDesignEntity

```go
kaspPrimerDesign := client.KaspPrimerDesign(nil)
fmt.Println(kaspPrimerDesign.GetName()) // "kasp_primer_design"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addSecondaryMismatch` | `bool` | No | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | `string` | Yes | First allele (single base) — gets the FAM tail. |
| `alleleB` | `string` | Yes | Second allele (single base) — gets the HEX tail. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | `int` | No | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | `int` | No | Minimum amplicon length for the common reverse primer. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `snpPosition` | `int` | Yes | 1-based position of the SNP on the forward strand. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | `float64` | No | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.KaspPrimerDesign(nil).Create(map[string]any{
    "alleleA": "example_alleleA",
    "alleleB": "example_alleleB",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "snpPosition": 1,
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `KaspPrimerDesignEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ListToolEntity

```go
listTool := client.ListTool(nil)
fmt.Println(listTool.GetName()) // "list_tool"
```

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.ListTool(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ListToolEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MeltingTemperatureEntity

```go
meltingTemperature := client.MeltingTemperature(nil)
fmt.Println(meltingTemperature.GetName()) // "melting_temperature"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntpMM` | `float64` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float64` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float64` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `float64` | No | Total strand concentration (nM). |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | `float64` | No | Optional target Tm (°C). |
| `tmTolerance` | `float64` | No | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.MeltingTemperature(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MeltingTemperatureEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MotifFinderEntity

```go
motifFinder := client.MotifFinder(nil)
fmt.Println(motifFinder.GetName()) // "motif_finder"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Maximum allowed mismatches per match. |
| `motif` | `string` | Yes | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `bool` | No | Also search the reverse strand. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.MotifFinder(nil).Create(map[string]any{
    "motif": "example_motif",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MotifFinderEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MultipleSequenceAlignmentEntity

```go
multipleSequenceAlignment := client.MultipleSequenceAlignment(nil)
fmt.Println(multipleSequenceAlignment.GetName()) // "multiple_sequence_alignment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.MultipleSequenceAlignment(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MultipleSequenceAlignmentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## OligoAnalysiEntity

```go
oligoAnalysi := client.OligoAnalysi(nil)
fmt.Println(oligoAnalysi.GetName()) // "oligo_analysi"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntpMM` | `float64` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float64` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float64` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `float64` | No | Total strand concentration (nM). |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.OligoAnalysi(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `OligoAnalysiEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## OrthologMapEntity

```go
orthologMap := client.OrthologMap(nil)
fmt.Println(orthologMap.GetName()) // "ortholog_map"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sourceSpecies` | `string` | No | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | `[]any` | Yes | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | `string` | Yes | Ensembl species slug to find homologs in (e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `type` | `string` | No | Homology type to return. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.OrthologMap(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "symbols": []any{},
    "targetSpecies": "example_targetSpecies",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `OrthologMapEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PairwiseAlignmentEntity

```go
pairwiseAlignment := client.PairwiseAlignment(nil)
fmt.Println(pairwiseAlignment.GetName()) // "pairwise_alignment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gap` | `float64` | No | Linear gap penalty (per gap position). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | `float64` | No | Match score. |
| `mismatch` | `float64` | No | Mismatch penalty. |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `seqA` | `string` | Yes | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | `string` | Yes | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PairwiseAlignment(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "seqA": "example_seqA",
    "seqB": "example_seqB",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PairwiseAlignmentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ParseGenbankEntity

```go
parseGenbank := client.ParseGenbank(nil)
fmt.Println(parseGenbank.GetName()) // "parse_genbank"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `text` | `string` | Yes | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ParseGenbank(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "text": "example_text",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ParseGenbankEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ParseSangerTraceEntity

```go
parseSangerTrace := client.ParseSangerTrace(nil)
fmt.Println(parseSangerTrace.GetName()) // "parse_sanger_trace"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fileBase64` | `string` | Yes | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | No | Optional original file name (echoed back). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ParseSangerTrace(nil).Create(map[string]any{
    "fileBase64": "example_fileBase64",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ParseSangerTraceEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PlasmidAnnotateEntity

```go
plasmidAnnotate := client.PlasmidAnnotate(nil)
fmt.Println(plasmidAnnotate.GetName()) // "plasmid_annotate"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PlasmidAnnotate(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PlasmidAnnotateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PlasmidDeepAnnotateEntity

```go
plasmidDeepAnnotate := client.PlasmidDeepAnnotate(nil)
fmt.Println(plasmidDeepAnnotate.GetName()) // "plasmid_deep_annotate"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No | Treat the sequence as a circular plasmid (vs. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PlasmidDeepAnnotate(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PlasmidDeepAnnotateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PlasmidFullReportEntity

```go
plasmidFullReport := client.PlasmidFullReport(nil)
fmt.Println(plasmidFullReport.GetName()) // "plasmid_full_report"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `topN` | `int` | No | How many top-ranked backbone candidates to report. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PlasmidFullReport(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PlasmidFullReportEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PlasmidIdentifyEntity

```go
plasmidIdentify := client.PlasmidIdentify(nil)
fmt.Println(plasmidIdentify.GetName()) // "plasmid_identify"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `topN` | `int` | No | How many top-ranked backbone candidates to report. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PlasmidIdentify(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PlasmidIdentifyEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PrimeEditingDesignEntity

```go
primeEditingDesign := client.PrimeEditingDesign(nil)
fmt.Println(primeEditingDesign.GetName()) // "prime_editing_design"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `editEnd` | `int` | Yes | 1-based inclusive end of the region being changed. |
| `editStart` | `int` | Yes | 1-based inclusive start of the region being changed. |
| `frameStart` | `int` | No | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | `string` | No | Replacement bases (forward strand). |
| `ok` | `any` | Yes |  |
| `pbsLength` | `int` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `rttHomology` | `int` | No | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | `string` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PrimeEditingDesign(nil).Create(map[string]any{
    "editEnd": 1,
    "editStart": 1,
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PrimeEditingDesignEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PrimeEditingTwinDesignEntity

```go
primeEditingTwinDesign := client.PrimeEditingTwinDesign(nil)
fmt.Println(primeEditingTwinDesign.GetName()) // "prime_editing_twin_design"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | `string` | Yes | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` | `any` | Yes |  |
| `overlapLength` | `int` | No | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | `int` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `map[string]any` | Yes |  |
| `replaceEnd` | `int` | Yes | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | `int` | Yes | 1-based inclusive start of the region being replaced/deleted. |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PrimeEditingTwinDesign(nil).Create(map[string]any{
    "newSequence": "example_newSequence",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "replaceEnd": 1,
    "replaceStart": 1,
    "result": map[string]any{},
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PrimeEditingTwinDesignEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PrimerDesignEntity

```go
primerDesign := client.PrimerDesign(nil)
fmt.Println(primerDesign.GetName()) // "primer_design"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ampliconMax` | `int` | No |  |
| `ampliconMin` | `int` | No |  |
| `dntpMM` | `float64` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` | `float64` | No |  |
| `gcMin` | `float64` | No |  |
| `lenMax` | `int` | No |  |
| `lenMin` | `int` | No |  |
| `lenOpt` | `int` | No |  |
| `maxReturn` | `int` | No | Number of best pairs to return. |
| `mgMM` | `float64` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float64` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `float64` | No | Total strand concentration (nM). |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `targetEnd` | `int` | No | 1-based inclusive end of the target region (optional). |
| `targetStart` | `int` | No | 1-based inclusive start of a region the product must span (optional). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` | `float64` | No |  |
| `tmMaxDiff` | `float64` | No | Max Tm difference within a pair (°C). |
| `tmMin` | `float64` | No |  |
| `tmOpt` | `float64` | No |  |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PrimerDesign(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "template": "example_template",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PrimerDesignEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PrimerSpecificityEntity

```go
primerSpecificity := client.PrimerSpecificity(nil)
fmt.Println(primerSpecificity.GetName()) // "primer_specificity"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `forwardPrimer` | `string` | Yes | Forward primer, 5'→3'. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | `int` | No | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `reversePrimer` | `string` | Yes | Reverse primer, 5'→3'. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PrimerSpecificity(nil).Create(map[string]any{
    "forwardPrimer": "example_forwardPrimer",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "reversePrimer": "example_reversePrimer",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PrimerSpecificityEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ProteaseDigestionEntity

```go
proteaseDigestion := client.ProteaseDigestion(nil)
fmt.Println(proteaseDigestion.GetName()) // "protease_digestion"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | `float64` | No | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | `int` | No | Cap on the number of returned peptides. |
| `minMass` | `float64` | No | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | `int` | No | Allowed missed internal cleavages (0–2). |
| `ok` | `any` | Yes |  |
| `protease` | `string` | No | Protease or chemical cleavage agent. |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ProteaseDigestion(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ProteaseDigestionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ProteinAnnotatePollEntity

```go
proteinAnnotatePoll := client.ProteinAnnotatePoll(nil)
fmt.Println(proteinAnnotatePoll.GetName()) // "protein_annotate_poll"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ProteinAnnotatePoll(nil).Create(map[string]any{
    "jobId": "example_jobId",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ProteinAnnotatePollEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ProteinAnnotateSubmitEntity

```go
proteinAnnotateSubmit := client.ProteinAnnotateSubmit(nil)
fmt.Println(proteinAnnotateSubmit.GetName()) // "protein_annotate_submit"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `appl` | `string` | No | Restrict to one member database (e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | `bool` | No | Include GO-term cross-references. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ProteinAnnotateSubmit(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ProteinAnnotateSubmitEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ProteinHydrophobicityEntity

```go
proteinHydrophobicity := client.ProteinHydrophobicity(nil)
fmt.Println(proteinHydrophobicity.GetName()) // "protein_hydrophobicity"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `scale` | `string` | No | Amino-acid scale. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `window` | `int` | No | Sliding-window size (clamped to an odd number ≥ 1). |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ProteinHydrophobicity(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ProteinHydrophobicityEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ProteinPropertyEntity

```go
proteinProperty := client.ProteinProperty(nil)
fmt.Println(proteinProperty.GetName()) // "protein_property"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `chargeStep` | `float64` | No | pH step for the net-charge titration curve (0–14). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ProteinProperty(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ProteinPropertyEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RandomSequenceEntity

```go
randomSequence := client.RandomSequence(nil)
fmt.Println(randomSequence.GetName()) // "random_sequence"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | `float64` | No | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` | `string` | No |  |
| `length` | `int` | Yes | Number of residues to generate. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.RandomSequence(nil).Create(map[string]any{
    "length": 1,
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RandomSequenceEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RestrictionSiteEntity

```go
restrictionSite := client.RestrictionSite(nil)
fmt.Println(restrictionSite.GetName()) // "restriction_site"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzymes` | `[]any` | No | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.RestrictionSite(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RestrictionSiteEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ReverseComplementEntity

```go
reverseComplement := client.ReverseComplement(nil)
fmt.Println(reverseComplement.GetName()) // "reverse_complement"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `type` | `string` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ReverseComplement(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ReverseComplementEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ReverseTranslateEntity

```go
reverseTranslate := client.ReverseTranslate(nil)
fmt.Println(reverseTranslate.GetName()) // "reverse_translate"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No | Codon-usage host (ignored in degenerate mode). |
| `protein` | `string` | Yes | Protein sequence (one-letter codes; * for stop). |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ReverseTranslate(nil).Create(map[string]any{
    "ok": "example_ok",
    "protein": "example_protein",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ReverseTranslateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RnaFoldEntity

```go
rnaFold := client.RnaFold(nil)
fmt.Println(rnaFold.GetName()) // "rna_fold"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.RnaFold(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RnaFoldEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SangerVsReferenceEntity

```go
sangerVsReference := client.SangerVsReference(nil)
fmt.Println(sangerVsReference.GetName()) // "sanger_vs_reference"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fileBase64` | `string` | No | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | No | Optional original file name (echoed back). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | `float64` | No | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `read` | `string` | No | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | `string` | Yes | Expected reference sequence (FASTA or raw). |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SangerVsReference(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reference": "example_reference",
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SangerVsReferenceEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SavePermalinkEntity

```go
savePermalink := client.SavePermalink(nil)
fmt.Println(savePermalink.GetName()) // "save_permalink"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `map[string]any` | Yes | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SavePermalink(nil).Create(map[string]any{
    "args": map[string]any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SavePermalinkEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SeqfileStatEntity

```go
seqfileStat := client.SeqfileStat(nil)
fmt.Println(seqfileStat.GetName()) // "seqfile_stat"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `qualityOffset` | `int` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SeqfileStat(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SeqfileStatEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SequenceFetchEntity

```go
sequenceFetch := client.SequenceFetch(nil)
fmt.Println(sequenceFetch.GetName()) // "sequence_fetch"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `string` | Yes | GenBank/RefSeq accession (e.g. |
| `db` | `string` | No | Database to query; auto-detects from the accession format. |
| `format` | `string` | No | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SequenceFetch(nil).Create(map[string]any{
    "accession": "example_accession",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SequenceFetchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SequenceFormatConvertEntity

```go
sequenceFormatConvert := client.SequenceFormatConvert(nil)
fmt.Println(sequenceFormatConvert.GetName()) // "sequence_format_convert"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | No | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | A FASTA or GenBank record to convert. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `to` | `string` | No | Output format. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SequenceFormatConvert(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SequenceFormatConvertEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SequenceReportEntity

```go
sequenceReport := client.SequenceReport(nil)
fmt.Println(sequenceReport.GetName()) // "sequence_report"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `endPrimerLength` | `int` | No | Length of the naive end primers taken from each end. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `int` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `int` | No | Minimum ORF length in amino acids. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SequenceReport(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SequenceReportEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SequenceSearchEntity

```go
sequenceSearch := client.SequenceSearch(nil)
fmt.Println(sequenceSearch.GetName()) // "sequence_search"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `db` | `string` | No |  |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | No | Gene symbol/name, e.g. |
| `maxResults` | `int` | No | Up to 20. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No | Organism name, e.g. |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `term` | `string` | No | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SequenceSearch(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SequenceSearchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SequencingReadbackVerifyEntity

```go
sequencingReadbackVerify := client.SequencingReadbackVerify(nil)
fmt.Println(sequencingReadbackVerify.GetName()) // "sequencing_readback_verify"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | `int` | No | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `reads` | `string` | Yes | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | `string` | Yes | The claimed/expected reference sequence. |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SequencingReadbackVerify(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reads": "example_reads",
    "reference": "example_reference",
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SequencingReadbackVerifyEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SessionCreateEntity

```go
sessionCreate := client.SessionCreate(nil)
fmt.Println(sessionCreate.GetName()) // "session_create"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entries` | `map[string]any` | No | Initial named entries, e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SessionCreate(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SessionCreateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SessionGetEntity

```go
sessionGet := client.SessionGet(nil)
fmt.Println(sessionGet.GetName()) // "session_get"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | `[]any` | No | Only return these entries; omit to return all of them. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SessionGet(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sessionId": "example_sessionId",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SessionGetEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SessionRunEntity

```go
sessionRun := client.SessionRun(nil)
fmt.Println(sessionRun.GetName()) // "session_run"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `map[string]any` | No | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | `map[string]any` | No | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |
| `writeBack` | `map[string]any` | No | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SessionRun(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sessionId": "example_sessionId",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SessionRunEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SessionSetEntity

```go
sessionSet := client.SessionSet(nil)
fmt.Println(sessionSet.GetName()) // "session_set"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entries` | `map[string]any` | Yes | Named entries to add/overwrite, e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SessionSet(nil).Create(map[string]any{
    "entries": map[string]any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sessionId": "example_sessionId",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SessionSetEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SirnaDesignEntity

```go
sirnaDesign := client.SirnaDesign(nil)
fmt.Println(sirnaDesign.GetName()) // "sirna_design"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | `int` | No | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `shRnaLoop` | `string` | No | Loop sequence used when assembling the shRNA cassette. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SirnaDesign(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SirnaDesignEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SiteDirectedMutagenesiEntity

```go
siteDirectedMutagenesi := client.SiteDirectedMutagenesi(nil)
fmt.Println(siteDirectedMutagenesi.GetName()) // "site_directed_mutagenesi"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `float64` | No | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | `float64` | No | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | `string` | No | Edit at the nucleotide or amino-acid level. |
| `frameStart` | `int` | No | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float64` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float64` | No | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | `string` | No | Replacement base (editKind='nt'). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `float64` | No | Total strand concentration (nM). |
| `organism` | `string` | No | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | `int` | No | 1-based position to substitute (editKind='nt'). |
| `provenance` | `map[string]any` | Yes |  |
| `residue` | `int` | No | 1-based residue number to change (editKind='aa'). |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `style` | `string` | No | Mutagenic primer style. |
| `targetAa` | `string` | No | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SiteDirectedMutagenesi(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "template": "example_template",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SiteDirectedMutagenesiEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TranslateEntity

```go
translate := client.Translate(nil)
fmt.Println(translate.GetName()) // "translate"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frame` | `int` | No |  |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | `bool` | No | Stop at the first stop codon. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Translate(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TranslateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## VariantAnnotateEntity

```go
variantAnnotate := client.VariantAnnotate(nil)
fmt.Println(variantAnnotate.GetName()) // "variant_annotate"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `assembly` | `string` | No | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `variant` | `string` | Yes | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.VariantAnnotate(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
    "variant": "example_variant",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `VariantAnnotateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## VariantComparatorEntity

```go
variantComparator := client.VariantComparator(nil)
fmt.Println(variantComparator.GetName()) // "variant_comparator"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `coding` | `bool` | No | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | `int` | No | 1-based reading-frame start (used when coding is true). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `query` | `string` | Yes | Query / variant sequence (raw or FASTA). |
| `reference` | `string` | Yes | Reference / wild-type sequence (raw or FASTA). |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.VariantComparator(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "query": "example_query",
    "reference": "example_reference",
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `VariantComparatorEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## VerifyAssemblyEntity

```go
verifyAssembly := client.VerifyAssembly(nil)
fmt.Println(verifyAssembly.GetName()) // "verify_assembly"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `float64` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `bool` | No | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | `string` | Yes | The sequence you claim you ended up with. |
| `coding` | `bool` | No | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | `string` | No | Type IIS enzyme for Golden Gate. |
| `enzyme3` | `string` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | No | 5′ enzyme (restriction method). |
| `fragmentPcrs` | `[]any` | No | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | `[]any` | No | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | `int` | No | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | No | Insert sequence (restriction method). |
| `insertPcr` | `map[string]any` | No | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | `string` | Yes | Assembly method used. |
| `names` | `[]any` | No | Optional labels for each fragment. |
| `ok` | `any` | Yes |  |
| `overlapLen` | `int` | No | Gibson homology-arm length (bp). |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `vector` | `string` | No | Vector sequence (restriction method). |
| `vectorPcr` | `map[string]any` | No | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.VerifyAssembly(nil).Create(map[string]any{
    "claimedConstruct": "example_claimedConstruct",
    "method": "example_method",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `VerifyAssemblyEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## VerifyConstructEntity

```go
verifyConstruct := client.VerifyConstruct(nil)
fmt.Println(verifyConstruct.GetName()) // "verify_construct"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `claimedConstruct` | `string` | Yes | The final sequence claimed to have been built. |
| `expectedFrameStart` | `int` | No | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | `string` | Yes | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | `string` | Yes | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | `string` | Yes | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | `int` | No | Mismatches tolerated per primer during PCR prediction. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `templateCircular` | `bool` | No | Treat insertTemplate as circular (e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.VerifyConstruct(nil).Create(map[string]any{
    "claimedConstruct": "example_claimedConstruct",
    "insertForwardPrimer": "example_insertForwardPrimer",
    "insertReversePrimer": "example_insertReversePrimer",
    "insertTemplate": "example_insertTemplate",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `VerifyConstructEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## VirtualGelEntity

```go
virtualGel := client.VirtualGel(nil)
fmt.Println(virtualGel.GetName()) // "virtual_gel"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No | Treat the sequence as circular (plasmid). |
| `enzymes` | `[]any` | No | Enzyme names to digest with. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | `string` | No | DNA ladder to plot alongside the sample lane. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.VirtualGel(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `VirtualGelEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## VolcanoPlotDataEntity

```go
volcanoPlotData := client.VolcanoPlotData(nil)
fmt.Println(volcanoPlotData.GetName()) // "volcano_plot_data"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `rows` | `[]any` | Yes | Differential expression rows, one per gene. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.VolcanoPlotData(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "rows": []any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `VolcanoPlotDataEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## WebSearchEntity

```go
webSearch := client.WebSearch(nil)
fmt.Println(webSearch.GetName()) // "web_search"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | `float64` | No | Maximum number of results to return (default 5, max 10). |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `query` | `string` | Yes | The search query. |
| `result` | `map[string]any` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.WebSearch(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "query": "example_query",
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `WebSearchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewSeqbenchMcpSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

