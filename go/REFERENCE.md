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
| `accession` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `length` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `wing` | `int` | No |  |

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
| `editor` | `string` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `target` | `string` | Yes |  |
| `target_position` | `int` | No |  |
| `tool` | `string` | Yes |  |

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
| `arg` | `map[string]any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
    "input": "example_input",
    "ok": "example_ok",
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
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `step` | `[]any` | Yes |  |

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
    "input": "example_input",
    "ok": "example_ok",
    "result": map[string]any{},
    "step": []any{},
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
| `end_primer_length` | `int` | No |  |
| `gate` | `any` | No |  |
| `max_orf` | `int` | No |  |
| `min_orf_aa` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_tm_target` | `float64` | No |  |
| `circular` | `bool` | No |  |
| `enzyme` | `string` | No |  |
| `enzyme3` | `string` | No |  |
| `enzyme5` | `string` | No |  |
| `fragment` | `[]any` | No |  |
| `gate` | `any` | No |  |
| `insert` | `string` | No |  |
| `method` | `string` | Yes |  |
| `name` | `[]any` | No |  |
| `ok` | `any` | Yes |  |
| `overlap_len` | `int` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |
| `vector` | `string` | No |  |

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
| `frame_start` | `int` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `rare_threshold` | `float64` | No |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `avoid_enzyme` | `[]any` | No |  |
| `cryptic_orf_min_aa` | `int` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `any` | No |  |
| `gc_high` | `float64` | No |  |
| `gc_low` | `float64` | No |  |
| `gc_window` | `int` | No |  |
| `homopolymer_min` | `int` | No |  |
| `max_pass` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `avoid_enzyme` | `[]any` | No |  |
| `cryptic_orf_min_aa` | `int` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `any` | No |  |
| `gc_high` | `float64` | No |  |
| `gc_low` | `float64` | No |  |
| `gc_window` | `int` | No |  |
| `homopolymer_min` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `min_score` | `float64` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `search_reverse_strand` | `bool` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_length` | `int` | No |  |
| `block_pam` | `bool` | No |  |
| `design_genotyping_primer` | `bool` | No |  |
| `edit_end` | `int` | No |  |
| `edit_start` | `int` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `any` | No |  |
| `guide_end` | `int` | No |  |
| `guide_start` | `int` | No |  |
| `guide_strand` | `string` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `replacement` | `string` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `target_sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CrisprHdrDonor(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "replacement": "example_replacement",
    "result": map[string]any{},
    "target_sequence": "example_target_sequence",
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
| `gate` | `any` | No |  |
| `max_mismatch` | `int` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `protospacer` | `string` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence_a` | `string` | Yes |  |
| `sequence_b` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CrossDimer(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence_a": "example_sequence_a",
    "sequence_b": "example_sequence_b",
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
| `gate` | `any` | No |  |
| `length` | `int` | No |  |
| `mass_ng` | `float64` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | No |  |
| `tool` | `string` | Yes |  |
| `type` | `string` | No |  |
| `volume_ul` | `float64` | No |  |

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
| `enzyme_a` | `string` | Yes |  |
| `enzyme_b` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.DoubleDigest(nil).Create(map[string]any{
    "enzyme_a": "example_enzyme_a",
    "enzyme_b": "example_enzyme_b",
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `reaction` | `[]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ExportEchoPicklist(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reaction": []any{},
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `protocol_name` | `string` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `reaction` | `[]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ExportOpentronsProtocol(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reaction": []any{},
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `reaction` | `[]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ExportPlateLayout(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reaction": []any{},
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
| `cluster_col` | `bool` | No |  |
| `cluster_row` | `bool` | No |  |
| `distance_metric` | `string` | No |  |
| `gate` | `any` | No |  |
| `gene` | `[]any` | Yes |  |
| `linkage` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sample` | `[]any` | Yes |  |
| `tool` | `string` | Yes |  |
| `value` | `[]any` | Yes |  |
| `z_score_row` | `bool` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ExpressionHeatmapCluster(nil).Create(map[string]any{
    "gene": []any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sample": []any{},
    "tool": "example_tool",
    "value": []any{},
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
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `quality_offset` | `int` | No |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `min_length` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `quality_offset` | `int` | No |  |
| `quality_threshold` | `int` | No |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `min_aa_length` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `require_stop` | `bool` | No |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `case_mode` | `string` | No |  |
| `convert` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `reverse` | `bool` | No |  |
| `sequence` | `string` | Yes |  |
| `strip_non_letter` | `bool` | No |  |
| `tool` | `string` | Yes |  |
| `width` | `int` | No |  |

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
| `background` | `[]any` | No |  |
| `collection` | `[]any` | No |  |
| `gate` | `any` | No |  |
| `gene` | `[]any` | Yes |  |
| `max_term_size` | `int` | No |  |
| `min_term_size` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.FunctionalEnrichment(nil).Create(map[string]any{
    "gene": []any{},
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `compare_to_named_set` | `string` | No |  |
| `dataset` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `overhang` | `[]any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `risk_threshold` | `float64` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.GoldenGateFidelity(nil).Create(map[string]any{
    "ok": "example_ok",
    "overhang": []any{},
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |
| `variant` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `job_id` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.IdMapPoll(nil).Create(map[string]any{
    "job_id": "example_job_id",
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
| `from` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `ids` | `[]any` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tax_id` | `string` | No |  |
| `to` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `circular` | `bool` | No |  |
| `forward_primer` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `max_mismatch` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `reverse_primer` | `string` | Yes |  |
| `template` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.InSilicoPcr(nil).Create(map[string]any{
    "forward_primer": "example_forward_primer",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "reverse_primer": "example_reverse_primer",
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
| `add_secondary_mismatch` | `bool` | No |  |
| `allele_a` | `string` | Yes |  |
| `allele_b` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `max_amplicon` | `int` | No |  |
| `min_amplicon` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `snp_position` | `int` | Yes |  |
| `target` | `string` | Yes |  |
| `target_core_tm` | `float64` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.KaspPrimerDesign(nil).Create(map[string]any{
    "allele_a": "example_allele_a",
    "allele_b": "example_allele_b",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "snp_position": 1,
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
| `dntp_mm` | `float64` | No |  |
| `gate` | `any` | No |  |
| `mg_mm` | `float64` | No |  |
| `na_mm` | `float64` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `float64` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `target_tm` | `float64` | No |  |
| `tm_tolerance` | `float64` | No |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `max_mismatch` | `int` | No |  |
| `motif` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `search_reverse_strand` | `bool` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `dntp_mm` | `float64` | No |  |
| `gate` | `any` | No |  |
| `mg_mm` | `float64` | No |  |
| `na_mm` | `float64` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `float64` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `source_species` | `string` | No |  |
| `symbol` | `[]any` | Yes |  |
| `target_species` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `type` | `string` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.OrthologMap(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "symbol": []any{},
    "target_species": "example_target_species",
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
| `gap` | `float64` | No |  |
| `gate` | `any` | No |  |
| `match` | `float64` | No |  |
| `mismatch` | `float64` | No |  |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `seq_a` | `string` | Yes |  |
| `seq_b` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PairwiseAlignment(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "seq_a": "example_seq_a",
    "seq_b": "example_seq_b",
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `text` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `file_base64` | `string` | Yes |  |
| `file_name` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ParseSangerTrace(nil).Create(map[string]any{
    "file_base64": "example_file_base64",
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `circular` | `bool` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `circular` | `bool` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `top_n` | `int` | No |  |

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
| `circular` | `bool` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `top_n` | `int` | No |  |

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
| `edit_end` | `int` | Yes |  |
| `edit_start` | `int` | Yes |  |
| `frame_start` | `int` | No |  |
| `gate` | `any` | No |  |
| `inserted_seq` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `pbs_length` | `int` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `rtt_homology` | `int` | No |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PrimeEditingDesign(nil).Create(map[string]any{
    "edit_end": 1,
    "edit_start": 1,
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
| `gate` | `any` | No |  |
| `new_sequence` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `overlap_length` | `int` | No |  |
| `pbs_length` | `int` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `replace_end` | `int` | Yes |  |
| `replace_start` | `int` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PrimeEditingTwinDesign(nil).Create(map[string]any{
    "new_sequence": "example_new_sequence",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "replace_end": 1,
    "replace_start": 1,
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
| `amplicon_max` | `int` | No |  |
| `amplicon_min` | `int` | No |  |
| `dntp_mm` | `float64` | No |  |
| `gate` | `any` | No |  |
| `gc_max` | `float64` | No |  |
| `gc_min` | `float64` | No |  |
| `len_max` | `int` | No |  |
| `len_min` | `int` | No |  |
| `len_opt` | `int` | No |  |
| `max_return` | `int` | No |  |
| `mg_mm` | `float64` | No |  |
| `na_mm` | `float64` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `float64` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `target_end` | `int` | No |  |
| `target_start` | `int` | No |  |
| `template` | `string` | Yes |  |
| `tm_max` | `float64` | No |  |
| `tm_max_diff` | `float64` | No |  |
| `tm_min` | `float64` | No |  |
| `tm_opt` | `float64` | No |  |
| `tool` | `string` | Yes |  |

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
| `forward_primer` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `max_mismatch` | `int` | No |  |
| `max_product_length` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `reverse_primer` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PrimerSpecificity(nil).Create(map[string]any{
    "forward_primer": "example_forward_primer",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "reverse_primer": "example_reverse_primer",
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
| `gate` | `any` | No |  |
| `max_mass` | `float64` | No |  |
| `max_peptide` | `int` | No |  |
| `min_mass` | `float64` | No |  |
| `missed_cleavage` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `protease` | `string` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `job_id` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ProteinAnnotatePoll(nil).Create(map[string]any{
    "job_id": "example_job_id",
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
| `appl` | `string` | No |  |
| `gate` | `any` | No |  |
| `goterm` | `bool` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `scale` | `string` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `window` | `int` | No |  |

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
| `charge_step` | `float64` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `gc_content` | `float64` | No |  |
| `kind` | `string` | No |  |
| `length` | `int` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `enzyme` | `[]any` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
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
| `gate` | `any` | No |  |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `file_base64` | `string` | No |  |
| `file_name` | `string` | No |  |
| `gate` | `any` | No |  |
| `min_coverage` | `float64` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `read` | `string` | No |  |
| `reference` | `string` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arg` | `map[string]any` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SavePermalink(nil).Create(map[string]any{
    "arg": map[string]any{},
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
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `quality_offset` | `int` | No |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `accession` | `string` | Yes |  |
| `db` | `string` | No |  |
| `format` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `from` | `string` | No |  |
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `to` | `string` | No |  |
| `tool` | `string` | Yes |  |

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
| `end_primer_length` | `int` | No |  |
| `gate` | `any` | No |  |
| `max_orf` | `int` | No |  |
| `min_orf_aa` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `gene` | `string` | No |  |
| `max_result` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `term` | `string` | No |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `min_supporting_read` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `read` | `string` | Yes |  |
| `reference` | `string` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SequencingReadbackVerify(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "read": "example_read",
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
| `entry` | `map[string]any` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `name` | `[]any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SessionGet(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "session_id": "example_session_id",
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
| `arg` | `map[string]any` | No |  |
| `from_session` | `map[string]any` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `write_back` | `map[string]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SessionRun(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "session_id": "example_session_id",
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
| `entry` | `map[string]any` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.SessionSet(nil).Create(map[string]any{
    "entry": map[string]any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "session_id": "example_session_id",
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
| `gate` | `any` | No |  |
| `min_reynold` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sh_rna_loop` | `string` | No |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_tm_target` | `float64` | No |  |
| `dntp_mm` | `float64` | No |  |
| `edit_kind` | `string` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `any` | No |  |
| `mg_mm` | `float64` | No |  |
| `na_mm` | `float64` | No |  |
| `new_base` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `float64` | No |  |
| `organism` | `string` | No |  |
| `position` | `int` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `residue` | `int` | No |  |
| `result` | `map[string]any` | Yes |  |
| `style` | `string` | No |  |
| `target_aa` | `string` | No |  |
| `template` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `to_stop` | `bool` | No |  |
| `tool` | `string` | Yes |  |

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
| `assembly` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |
| `variant` | `string` | Yes |  |

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
| `coding` | `bool` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `query` | `string` | Yes |  |
| `reference` | `string` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_tm_target` | `float64` | No |  |
| `circular` | `bool` | No |  |
| `claimed_construct` | `string` | Yes |  |
| `coding` | `bool` | No |  |
| `enzyme` | `string` | No |  |
| `enzyme3` | `string` | No |  |
| `enzyme5` | `string` | No |  |
| `fragment` | `[]any` | No |  |
| `fragment_pcr` | `[]any` | No |  |
| `frame_start` | `int` | No |  |
| `gate` | `any` | No |  |
| `insert` | `string` | No |  |
| `insert_pcr` | `map[string]any` | No |  |
| `method` | `string` | Yes |  |
| `name` | `[]any` | No |  |
| `ok` | `any` | Yes |  |
| `overlap_len` | `int` | No |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |
| `vector` | `string` | No |  |
| `vector_pcr` | `map[string]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.VerifyAssembly(nil).Create(map[string]any{
    "claimed_construct": "example_claimed_construct",
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
| `claimed_construct` | `string` | Yes |  |
| `expected_frame_start` | `int` | No |  |
| `gate` | `any` | No |  |
| `insert_forward_primer` | `string` | Yes |  |
| `insert_reverse_primer` | `string` | Yes |  |
| `insert_template` | `string` | Yes |  |
| `max_primer_mismatch` | `int` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `template_circular` | `bool` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.VerifyConstruct(nil).Create(map[string]any{
    "claimed_construct": "example_claimed_construct",
    "insert_forward_primer": "example_insert_forward_primer",
    "insert_reverse_primer": "example_insert_reverse_primer",
    "insert_template": "example_insert_template",
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
| `circular` | `bool` | No |  |
| `enzyme` | `[]any` | No |  |
| `gate` | `any` | No |  |
| `ladder` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `row` | `[]any` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.VolcanoPlotData(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "row": []any{},
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
| `gate` | `any` | No |  |
| `max_result` | `float64` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `map[string]any` | Yes |  |
| `query` | `string` | Yes |  |
| `result` | `map[string]any` | Yes |  |
| `tool` | `string` | Yes |  |

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

