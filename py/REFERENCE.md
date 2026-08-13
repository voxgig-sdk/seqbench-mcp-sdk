# SeqbenchMcp Python SDK Reference

Complete API reference for the SeqbenchMcp Python SDK.


## SeqbenchMcpSDK

### Constructor

```python
from seqbenchmcp_sdk import SeqbenchMcpSDK

client = SeqbenchMcpSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `SeqbenchMcpSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = SeqbenchMcpSDK.test()
```


### Instance Methods

#### `AlphafoldLookup(data=None)`

Create a new `AlphafoldLookupEntity` instance. Pass `None` for no initial data.

#### `AsoDesign(data=None)`

Create a new `AsoDesignEntity` instance. Pass `None` for no initial data.

#### `BaseEditingDesign(data=None)`

Create a new `BaseEditingDesignEntity` instance. Pass `None` for no initial data.

#### `Batch(data=None)`

Create a new `BatchEntity` instance. Pass `None` for no initial data.

#### `BatchWorkflow(data=None)`

Create a new `BatchWorkflowEntity` instance. Pass `None` for no initial data.

#### `CharacterizeSequence(data=None)`

Create a new `CharacterizeSequenceEntity` instance. Pass `None` for no initial data.

#### `CloningSimulate(data=None)`

Create a new `CloningSimulateEntity` instance. Pass `None` for no initial data.

#### `CodonAdaptationIndex(data=None)`

Create a new `CodonAdaptationIndexEntity` instance. Pass `None` for no initial data.

#### `CodonOptimize(data=None)`

Create a new `CodonOptimizeEntity` instance. Pass `None` for no initial data.

#### `ConstructAutofix(data=None)`

Create a new `ConstructAutofixEntity` instance. Pass `None` for no initial data.

#### `ConstructQc(data=None)`

Create a new `ConstructQcEntity` instance. Pass `None` for no initial data.

#### `CrisprGrnaDesign(data=None)`

Create a new `CrisprGrnaDesignEntity` instance. Pass `None` for no initial data.

#### `CrisprHdrDonor(data=None)`

Create a new `CrisprHdrDonorEntity` instance. Pass `None` for no initial data.

#### `CrisprOfftargetCheck(data=None)`

Create a new `CrisprOfftargetCheckEntity` instance. Pass `None` for no initial data.

#### `CrossDimer(data=None)`

Create a new `CrossDimerEntity` instance. Pass `None` for no initial data.

#### `DnaMolarity(data=None)`

Create a new `DnaMolarityEntity` instance. Pass `None` for no initial data.

#### `DoubleDigest(data=None)`

Create a new `DoubleDigestEntity` instance. Pass `None` for no initial data.

#### `ExportEchoPicklist(data=None)`

Create a new `ExportEchoPicklistEntity` instance. Pass `None` for no initial data.

#### `ExportOpentronsProtocol(data=None)`

Create a new `ExportOpentronsProtocolEntity` instance. Pass `None` for no initial data.

#### `ExportPlateLayout(data=None)`

Create a new `ExportPlateLayoutEntity` instance. Pass `None` for no initial data.

#### `ExpressionHeatmapCluster(data=None)`

Create a new `ExpressionHeatmapClusterEntity` instance. Pass `None` for no initial data.

#### `FastqQcReport(data=None)`

Create a new `FastqQcReportEntity` instance. Pass `None` for no initial data.

#### `FastqTrim(data=None)`

Create a new `FastqTrimEntity` instance. Pass `None` for no initial data.

#### `FindOrf(data=None)`

Create a new `FindOrfEntity` instance. Pass `None` for no initial data.

#### `FormatSequence(data=None)`

Create a new `FormatSequenceEntity` instance. Pass `None` for no initial data.

#### `FunctionalEnrichment(data=None)`

Create a new `FunctionalEnrichmentEntity` instance. Pass `None` for no initial data.

#### `GcContent(data=None)`

Create a new `GcContentEntity` instance. Pass `None` for no initial data.

#### `GeneDossier(data=None)`

Create a new `GeneDossierEntity` instance. Pass `None` for no initial data.

#### `GeneExpression(data=None)`

Create a new `GeneExpressionEntity` instance. Pass `None` for no initial data.

#### `GeneModel(data=None)`

Create a new `GeneModelEntity` instance. Pass `None` for no initial data.

#### `GoldenGateFidelity(data=None)`

Create a new `GoldenGateFidelityEntity` instance. Pass `None` for no initial data.

#### `HgvsConvert(data=None)`

Create a new `HgvsConvertEntity` instance. Pass `None` for no initial data.

#### `IdMapPoll(data=None)`

Create a new `IdMapPollEntity` instance. Pass `None` for no initial data.

#### `IdMapSubmit(data=None)`

Create a new `IdMapSubmitEntity` instance. Pass `None` for no initial data.

#### `InSilicoPcr(data=None)`

Create a new `InSilicoPcrEntity` instance. Pass `None` for no initial data.

#### `KaspPrimerDesign(data=None)`

Create a new `KaspPrimerDesignEntity` instance. Pass `None` for no initial data.

#### `ListTool(data=None)`

Create a new `ListToolEntity` instance. Pass `None` for no initial data.

#### `MeltingTemperature(data=None)`

Create a new `MeltingTemperatureEntity` instance. Pass `None` for no initial data.

#### `MotifFinder(data=None)`

Create a new `MotifFinderEntity` instance. Pass `None` for no initial data.

#### `MultipleSequenceAlignment(data=None)`

Create a new `MultipleSequenceAlignmentEntity` instance. Pass `None` for no initial data.

#### `OligoAnalysi(data=None)`

Create a new `OligoAnalysiEntity` instance. Pass `None` for no initial data.

#### `OrthologMap(data=None)`

Create a new `OrthologMapEntity` instance. Pass `None` for no initial data.

#### `PairwiseAlignment(data=None)`

Create a new `PairwiseAlignmentEntity` instance. Pass `None` for no initial data.

#### `ParseGenbank(data=None)`

Create a new `ParseGenbankEntity` instance. Pass `None` for no initial data.

#### `ParseSangerTrace(data=None)`

Create a new `ParseSangerTraceEntity` instance. Pass `None` for no initial data.

#### `PlasmidAnnotate(data=None)`

Create a new `PlasmidAnnotateEntity` instance. Pass `None` for no initial data.

#### `PlasmidDeepAnnotate(data=None)`

Create a new `PlasmidDeepAnnotateEntity` instance. Pass `None` for no initial data.

#### `PlasmidFullReport(data=None)`

Create a new `PlasmidFullReportEntity` instance. Pass `None` for no initial data.

#### `PlasmidIdentify(data=None)`

Create a new `PlasmidIdentifyEntity` instance. Pass `None` for no initial data.

#### `PrimeEditingDesign(data=None)`

Create a new `PrimeEditingDesignEntity` instance. Pass `None` for no initial data.

#### `PrimeEditingTwinDesign(data=None)`

Create a new `PrimeEditingTwinDesignEntity` instance. Pass `None` for no initial data.

#### `PrimerDesign(data=None)`

Create a new `PrimerDesignEntity` instance. Pass `None` for no initial data.

#### `PrimerSpecificity(data=None)`

Create a new `PrimerSpecificityEntity` instance. Pass `None` for no initial data.

#### `ProteaseDigestion(data=None)`

Create a new `ProteaseDigestionEntity` instance. Pass `None` for no initial data.

#### `ProteinAnnotatePoll(data=None)`

Create a new `ProteinAnnotatePollEntity` instance. Pass `None` for no initial data.

#### `ProteinAnnotateSubmit(data=None)`

Create a new `ProteinAnnotateSubmitEntity` instance. Pass `None` for no initial data.

#### `ProteinHydrophobicity(data=None)`

Create a new `ProteinHydrophobicityEntity` instance. Pass `None` for no initial data.

#### `ProteinProperty(data=None)`

Create a new `ProteinPropertyEntity` instance. Pass `None` for no initial data.

#### `RandomSequence(data=None)`

Create a new `RandomSequenceEntity` instance. Pass `None` for no initial data.

#### `RestrictionSite(data=None)`

Create a new `RestrictionSiteEntity` instance. Pass `None` for no initial data.

#### `ReverseComplement(data=None)`

Create a new `ReverseComplementEntity` instance. Pass `None` for no initial data.

#### `ReverseTranslate(data=None)`

Create a new `ReverseTranslateEntity` instance. Pass `None` for no initial data.

#### `RnaFold(data=None)`

Create a new `RnaFoldEntity` instance. Pass `None` for no initial data.

#### `SangerVsReference(data=None)`

Create a new `SangerVsReferenceEntity` instance. Pass `None` for no initial data.

#### `SavePermalink(data=None)`

Create a new `SavePermalinkEntity` instance. Pass `None` for no initial data.

#### `SeqfileStat(data=None)`

Create a new `SeqfileStatEntity` instance. Pass `None` for no initial data.

#### `SequenceFetch(data=None)`

Create a new `SequenceFetchEntity` instance. Pass `None` for no initial data.

#### `SequenceFormatConvert(data=None)`

Create a new `SequenceFormatConvertEntity` instance. Pass `None` for no initial data.

#### `SequenceReport(data=None)`

Create a new `SequenceReportEntity` instance. Pass `None` for no initial data.

#### `SequenceSearch(data=None)`

Create a new `SequenceSearchEntity` instance. Pass `None` for no initial data.

#### `SequencingReadbackVerify(data=None)`

Create a new `SequencingReadbackVerifyEntity` instance. Pass `None` for no initial data.

#### `SessionCreate(data=None)`

Create a new `SessionCreateEntity` instance. Pass `None` for no initial data.

#### `SessionGet(data=None)`

Create a new `SessionGetEntity` instance. Pass `None` for no initial data.

#### `SessionRun(data=None)`

Create a new `SessionRunEntity` instance. Pass `None` for no initial data.

#### `SessionSet(data=None)`

Create a new `SessionSetEntity` instance. Pass `None` for no initial data.

#### `SirnaDesign(data=None)`

Create a new `SirnaDesignEntity` instance. Pass `None` for no initial data.

#### `SiteDirectedMutagenesi(data=None)`

Create a new `SiteDirectedMutagenesiEntity` instance. Pass `None` for no initial data.

#### `Translate(data=None)`

Create a new `TranslateEntity` instance. Pass `None` for no initial data.

#### `VariantAnnotate(data=None)`

Create a new `VariantAnnotateEntity` instance. Pass `None` for no initial data.

#### `VariantComparator(data=None)`

Create a new `VariantComparatorEntity` instance. Pass `None` for no initial data.

#### `VerifyAssembly(data=None)`

Create a new `VerifyAssemblyEntity` instance. Pass `None` for no initial data.

#### `VerifyConstruct(data=None)`

Create a new `VerifyConstructEntity` instance. Pass `None` for no initial data.

#### `VirtualGel(data=None)`

Create a new `VirtualGelEntity` instance. Pass `None` for no initial data.

#### `VolcanoPlotData(data=None)`

Create a new `VolcanoPlotDataEntity` instance. Pass `None` for no initial data.

#### `WebSearch(data=None)`

Create a new `WebSearchEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## AlphafoldLookupEntity

```python
alphafold_lookup = client.AlphafoldLookup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `str` | Yes |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.AlphafoldLookup().create({
    "accession": "example_accession",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AlphafoldLookupEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## AsoDesignEntity

```python
aso_design = client.AsoDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `length` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `target` | `str` | Yes |  |
| `tool` | `str` | Yes |  |
| `wing` | `int` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.AsoDesign().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "target": "example_target",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AsoDesignEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BaseEditingDesignEntity

```python
base_editing_design = client.BaseEditingDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `editor` | `str` | No |  |
| `frameStart` | `int` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `target` | `str` | Yes |  |
| `targetPosition` | `int` | No |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.BaseEditingDesign().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "target": "example_target",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BaseEditingDesignEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BatchEntity

```python
batch = client.Batch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `dict` | No |  |
| `capped` | `bool` | Yes |  |
| `columns` | `list` | Yes |  |
| `count` | `int` | Yes |  |
| `errors` | `int` | Yes |  |
| `input` | `str` | Yes |  |
| `limit` | `int` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `rows` | `list` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Batch().create({
    "capped": True,  # bool
    "columns": [],  # list
    "count": 1,  # int
    "errors": 1,  # int
    "input": "example_input",  # str
    "limit": 1,  # int
    "provenance": {},  # dict
    "rows": [],  # list
    "tool": "example_tool",  # str
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Batch().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BatchEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BatchWorkflowEntity

```python
batch__workflow = client.BatchWorkflow()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `capped` | `bool` | Yes |  |
| `columns` | `list` | Yes |  |
| `count` | `int` | Yes |  |
| `errors` | `int` | Yes |  |
| `input` | `str` | Yes |  |
| `limit` | `int` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `rows` | `list` | Yes |  |
| `steps` | `list` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.BatchWorkflow().create({
    "capped": True,  # bool
    "columns": [],  # list
    "count": 1,  # int
    "errors": 1,  # int
    "input": "example_input",  # str
    "limit": 1,  # int
    "provenance": {},  # dict
    "rows": [],  # list
    "steps": [],  # list
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.BatchWorkflow().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BatchWorkflowEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CharacterizeSequenceEntity

```python
characterize_sequence = client.CharacterizeSequence()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `endPrimerLength` | `int` | No |  |
| `gate` | `Any` | No |  |
| `maxOrfs` | `int` | No |  |
| `minOrfAa` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CharacterizeSequence().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CharacterizeSequenceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CloningSimulateEntity

```python
cloning_simulate = client.CloningSimulate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `float` | No |  |
| `circular` | `bool` | No |  |
| `enzyme` | `str` | No |  |
| `enzyme3` | `str` | No |  |
| `enzyme5` | `str` | No |  |
| `fragments` | `list` | No |  |
| `gate` | `Any` | No |  |
| `insert` | `str` | No |  |
| `method` | `str` | Yes |  |
| `names` | `list` | No |  |
| `ok` | `Any` | Yes |  |
| `overlapLen` | `int` | No |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |
| `vector` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CloningSimulate().create({
    "method": "example_method",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CloningSimulateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CodonAdaptationIndexEntity

```python
codon_adaptation_index = client.CodonAdaptationIndex()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frameStart` | `int` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `organism` | `str` | No |  |
| `provenance` | `dict` | Yes |  |
| `rareThreshold` | `float` | No |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CodonAdaptationIndex().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CodonAdaptationIndexEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CodonOptimizeEntity

```python
codon_optimize = client.CodonOptimize()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `organism` | `str` | No |  |
| `protein` | `str` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CodonOptimize().create({
    "ok": "example_ok",  # Any
    "protein": "example_protein",  # str
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CodonOptimizeEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ConstructAutofixEntity

```python
construct_autofix = client.ConstructAutofix()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoidEnzymes` | `list` | No |  |
| `crypticOrfMinAa` | `int` | No |  |
| `frameStart` | `int` | No |  |
| `gate` | `Any` | No |  |
| `gcHigh` | `float` | No |  |
| `gcLow` | `float` | No |  |
| `gcWindow` | `int` | No |  |
| `homopolymerMin` | `int` | No |  |
| `maxPasses` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `organism` | `str` | No |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ConstructAutofix().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ConstructAutofixEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ConstructQcEntity

```python
construct_qc = client.ConstructQc()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoidEnzymes` | `list` | No |  |
| `crypticOrfMinAa` | `int` | No |  |
| `frameStart` | `int` | No |  |
| `gate` | `Any` | No |  |
| `gcHigh` | `float` | No |  |
| `gcLow` | `float` | No |  |
| `gcWindow` | `int` | No |  |
| `homopolymerMin` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ConstructQc().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ConstructQcEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CrisprGrnaDesignEntity

```python
crispr_grna_design = client.CrisprGrnaDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `minScore` | `float` | No |  |
| `nuclease` | `str` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `searchReverseStrand` | `bool` | No |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CrisprGrnaDesign().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CrisprGrnaDesignEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CrisprHdrDonorEntity

```python
crispr_hdr_donor = client.CrisprHdrDonor()
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
| `gate` | `Any` | No |  |
| `guideEnd` | `int` | No |  |
| `guideStart` | `int` | No |  |
| `guideStrand` | `str` | No |  |
| `nuclease` | `str` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `replacement` | `str` | Yes |  |
| `result` | `dict` | Yes |  |
| `targetSequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CrisprHdrDonor().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "replacement": "example_replacement",  # str
    "result": {},  # dict
    "targetSequence": "example_targetSequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CrisprHdrDonorEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CrisprOfftargetCheckEntity

```python
crispr_offtarget_check = client.CrisprOfftargetCheck()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `maxMismatches` | `int` | No |  |
| `nuclease` | `str` | No |  |
| `ok` | `Any` | Yes |  |
| `protospacer` | `str` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CrisprOfftargetCheck().create({
    "ok": "example_ok",  # Any
    "protospacer": "example_protospacer",  # str
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CrisprOfftargetCheckEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CrossDimerEntity

```python
cross_dimer = client.CrossDimer()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequenceA` | `str` | Yes |  |
| `sequenceB` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CrossDimer().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequenceA": "example_sequenceA",  # str
    "sequenceB": "example_sequenceB",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CrossDimerEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DnaMolarityEntity

```python
dna_molarity = client.DnaMolarity()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `length` | `int` | No |  |
| `massNg` | `float` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | No |  |
| `tool` | `str` | Yes |  |
| `type` | `str` | No |  |
| `volumeUl` | `float` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.DnaMolarity().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DnaMolarityEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DoubleDigestEntity

```python
double_digest = client.DoubleDigest()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzymeA` | `str` | Yes |  |
| `enzymeB` | `str` | Yes |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.DoubleDigest().create({
    "enzymeA": "example_enzymeA",  # str
    "enzymeB": "example_enzymeB",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DoubleDigestEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ExportEchoPicklistEntity

```python
export_echo_picklist = client.ExportEchoPicklist()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `reactions` | `list` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ExportEchoPicklist().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "reactions": [],  # list
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ExportEchoPicklistEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ExportOpentronsProtocolEntity

```python
export_opentrons_protocol = client.ExportOpentronsProtocol()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `protocolName` | `str` | No |  |
| `provenance` | `dict` | Yes |  |
| `reactions` | `list` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ExportOpentronsProtocol().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "reactions": [],  # list
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ExportOpentronsProtocolEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ExportPlateLayoutEntity

```python
export_plate_layout = client.ExportPlateLayout()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `reactions` | `list` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ExportPlateLayout().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "reactions": [],  # list
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ExportPlateLayoutEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ExpressionHeatmapClusterEntity

```python
expression_heatmap_cluster = client.ExpressionHeatmapCluster()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clusterCols` | `bool` | No |  |
| `clusterRows` | `bool` | No |  |
| `distanceMetric` | `str` | No |  |
| `gate` | `Any` | No |  |
| `genes` | `list` | Yes |  |
| `linkage` | `str` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `samples` | `list` | Yes |  |
| `tool` | `str` | Yes |  |
| `values` | `list` | Yes |  |
| `zScoreRows` | `bool` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ExpressionHeatmapCluster().create({
    "genes": [],  # list
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "samples": [],  # list
    "tool": "example_tool",  # str
    "values": [],  # list
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ExpressionHeatmapClusterEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## FastqQcReportEntity

```python
fastq_qc_report = client.FastqQcReport()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `input` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `qualityOffset` | `int` | No |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.FastqQcReport().create({
    "input": "example_input",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FastqQcReportEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## FastqTrimEntity

```python
fastq_trim = client.FastqTrim()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `input` | `str` | Yes |  |
| `minLength` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `qualityOffset` | `int` | No |  |
| `qualityThreshold` | `int` | No |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.FastqTrim().create({
    "input": "example_input",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FastqTrimEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## FindOrfEntity

```python
find_orf = client.FindOrf()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `minAaLength` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `requireStop` | `bool` | No |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.FindOrf().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FindOrfEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## FormatSequenceEntity

```python
format_sequence = client.FormatSequence()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `caseMode` | `str` | No |  |
| `convert` | `str` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `reverse` | `bool` | No |  |
| `sequence` | `str` | Yes |  |
| `stripNonLetters` | `bool` | No |  |
| `tool` | `str` | Yes |  |
| `width` | `int` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.FormatSequence().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FormatSequenceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## FunctionalEnrichmentEntity

```python
functional_enrichment = client.FunctionalEnrichment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `background` | `list` | No |  |
| `collections` | `list` | No |  |
| `gate` | `Any` | No |  |
| `genes` | `list` | Yes |  |
| `maxTermSize` | `int` | No |  |
| `minTermSize` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.FunctionalEnrichment().create({
    "genes": [],  # list
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FunctionalEnrichmentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GcContentEntity

```python
gc_content = client.GcContent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.GcContent().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GcContentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GeneDossierEntity

```python
gene_dossier = client.GeneDossier()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `gene` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.GeneDossier().create({
    "gene": "example_gene",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GeneDossierEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GeneExpressionEntity

```python
gene_expression = client.GeneExpression()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `gene` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.GeneExpression().create({
    "gene": "example_gene",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GeneExpressionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GeneModelEntity

```python
gene_model = client.GeneModel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `gene` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.GeneModel().create({
    "gene": "example_gene",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GeneModelEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GoldenGateFidelityEntity

```python
golden_gate_fidelity = client.GoldenGateFidelity()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `compareToNamedSet` | `str` | No |  |
| `dataset` | `str` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `overhangs` | `list` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `riskThreshold` | `float` | No |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.GoldenGateFidelity().create({
    "ok": "example_ok",  # Any
    "overhangs": [],  # list
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GoldenGateFidelityEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## HgvsConvertEntity

```python
hgvs_convert = client.HgvsConvert()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |
| `variant` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.HgvsConvert().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
    "variant": "example_variant",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HgvsConvertEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## IdMapPollEntity

```python
id_map_poll = client.IdMapPoll()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `jobId` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.IdMapPoll().create({
    "jobId": "example_jobId",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IdMapPollEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## IdMapSubmitEntity

```python
id_map_submit = client.IdMapSubmit()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `str` | Yes |  |
| `gate` | `Any` | No |  |
| `ids` | `list` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `taxId` | `str` | No |  |
| `to` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.IdMapSubmit().create({
    "from": "example_from",  # str
    "ids": [],  # list
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "to": "example_to",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IdMapSubmitEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InSilicoPcrEntity

```python
in_silico_pcr = client.InSilicoPcr()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No |  |
| `forwardPrimer` | `str` | Yes |  |
| `gate` | `Any` | No |  |
| `maxMismatches` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `reversePrimer` | `str` | Yes |  |
| `template` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.InSilicoPcr().create({
    "forwardPrimer": "example_forwardPrimer",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "reversePrimer": "example_reversePrimer",  # str
    "template": "example_template",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InSilicoPcrEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## KaspPrimerDesignEntity

```python
kasp_primer_design = client.KaspPrimerDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addSecondaryMismatch` | `bool` | No |  |
| `alleleA` | `str` | Yes |  |
| `alleleB` | `str` | Yes |  |
| `gate` | `Any` | No |  |
| `maxAmplicon` | `int` | No |  |
| `minAmplicon` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `snpPosition` | `int` | Yes |  |
| `target` | `str` | Yes |  |
| `targetCoreTm` | `float` | No |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.KaspPrimerDesign().create({
    "alleleA": "example_alleleA",  # str
    "alleleB": "example_alleleB",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "snpPosition": 1,  # int
    "target": "example_target",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `KaspPrimerDesignEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ListToolEntity

```python
list_tool = client.ListTool()
```

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.ListTool().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ListToolEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MeltingTemperatureEntity

```python
melting_temperature = client.MeltingTemperature()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntpMM` | `float` | No |  |
| `gate` | `Any` | No |  |
| `mgMM` | `float` | No |  |
| `naMM` | `float` | No |  |
| `ok` | `Any` | Yes |  |
| `oligoNM` | `float` | No |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `targetTm` | `float` | No |  |
| `tmTolerance` | `float` | No |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.MeltingTemperature().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MeltingTemperatureEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MotifFinderEntity

```python
motif_finder = client.MotifFinder()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `maxMismatches` | `int` | No |  |
| `motif` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `searchReverseStrand` | `bool` | No |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.MotifFinder().create({
    "motif": "example_motif",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MotifFinderEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MultipleSequenceAlignmentEntity

```python
multiple_sequence_alignment = client.MultipleSequenceAlignment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `input` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.MultipleSequenceAlignment().create({
    "input": "example_input",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MultipleSequenceAlignmentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## OligoAnalysiEntity

```python
oligo_analysi = client.OligoAnalysi()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntpMM` | `float` | No |  |
| `gate` | `Any` | No |  |
| `mgMM` | `float` | No |  |
| `naMM` | `float` | No |  |
| `ok` | `Any` | Yes |  |
| `oligoNM` | `float` | No |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.OligoAnalysi().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OligoAnalysiEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## OrthologMapEntity

```python
ortholog_map = client.OrthologMap()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sourceSpecies` | `str` | No |  |
| `symbols` | `list` | Yes |  |
| `targetSpecies` | `str` | Yes |  |
| `tool` | `str` | Yes |  |
| `type` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.OrthologMap().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "symbols": [],  # list
    "targetSpecies": "example_targetSpecies",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrthologMapEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PairwiseAlignmentEntity

```python
pairwise_alignment = client.PairwiseAlignment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gap` | `float` | No |  |
| `gate` | `Any` | No |  |
| `match` | `float` | No |  |
| `mismatch` | `float` | No |  |
| `mode` | `str` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `seqA` | `str` | Yes |  |
| `seqB` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PairwiseAlignment().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "seqA": "example_seqA",  # str
    "seqB": "example_seqB",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PairwiseAlignmentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ParseGenbankEntity

```python
parse_genbank = client.ParseGenbank()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `text` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ParseGenbank().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "text": "example_text",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ParseGenbankEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ParseSangerTraceEntity

```python
parse_sanger_trace = client.ParseSangerTrace()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fileBase64` | `str` | Yes |  |
| `fileName` | `str` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ParseSangerTrace().create({
    "fileBase64": "example_fileBase64",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ParseSangerTraceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PlasmidAnnotateEntity

```python
plasmid_annotate = client.PlasmidAnnotate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PlasmidAnnotate().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlasmidAnnotateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PlasmidDeepAnnotateEntity

```python
plasmid_deep_annotate = client.PlasmidDeepAnnotate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PlasmidDeepAnnotate().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlasmidDeepAnnotateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PlasmidFullReportEntity

```python
plasmid_full_report = client.PlasmidFullReport()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |
| `topN` | `int` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PlasmidFullReport().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlasmidFullReportEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PlasmidIdentifyEntity

```python
plasmid_identify = client.PlasmidIdentify()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |
| `topN` | `int` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PlasmidIdentify().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PlasmidIdentifyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PrimeEditingDesignEntity

```python
prime_editing_design = client.PrimeEditingDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `editEnd` | `int` | Yes |  |
| `editStart` | `int` | Yes |  |
| `frameStart` | `int` | No |  |
| `gate` | `Any` | No |  |
| `insertedSeq` | `str` | No |  |
| `ok` | `Any` | Yes |  |
| `pbsLength` | `int` | No |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `rttHomology` | `int` | No |  |
| `target` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PrimeEditingDesign().create({
    "editEnd": 1,  # int
    "editStart": 1,  # int
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "target": "example_target",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PrimeEditingDesignEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PrimeEditingTwinDesignEntity

```python
prime_editing_twin_design = client.PrimeEditingTwinDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `newSequence` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `overlapLength` | `int` | No |  |
| `pbsLength` | `int` | No |  |
| `provenance` | `dict` | Yes |  |
| `replaceEnd` | `int` | Yes |  |
| `replaceStart` | `int` | Yes |  |
| `result` | `dict` | Yes |  |
| `target` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PrimeEditingTwinDesign().create({
    "newSequence": "example_newSequence",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "replaceEnd": 1,  # int
    "replaceStart": 1,  # int
    "result": {},  # dict
    "target": "example_target",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PrimeEditingTwinDesignEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PrimerDesignEntity

```python
primer_design = client.PrimerDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ampliconMax` | `int` | No |  |
| `ampliconMin` | `int` | No |  |
| `dntpMM` | `float` | No |  |
| `gate` | `Any` | No |  |
| `gcMax` | `float` | No |  |
| `gcMin` | `float` | No |  |
| `lenMax` | `int` | No |  |
| `lenMin` | `int` | No |  |
| `lenOpt` | `int` | No |  |
| `maxReturn` | `int` | No |  |
| `mgMM` | `float` | No |  |
| `naMM` | `float` | No |  |
| `ok` | `Any` | Yes |  |
| `oligoNM` | `float` | No |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `targetEnd` | `int` | No |  |
| `targetStart` | `int` | No |  |
| `template` | `str` | Yes |  |
| `tmMax` | `float` | No |  |
| `tmMaxDiff` | `float` | No |  |
| `tmMin` | `float` | No |  |
| `tmOpt` | `float` | No |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PrimerDesign().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "template": "example_template",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PrimerDesignEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PrimerSpecificityEntity

```python
primer_specificity = client.PrimerSpecificity()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `forwardPrimer` | `str` | Yes |  |
| `gate` | `Any` | No |  |
| `maxMismatches` | `int` | No |  |
| `maxProductLength` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `reversePrimer` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PrimerSpecificity().create({
    "forwardPrimer": "example_forwardPrimer",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "reversePrimer": "example_reversePrimer",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PrimerSpecificityEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ProteaseDigestionEntity

```python
protease_digestion = client.ProteaseDigestion()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `maxMass` | `float` | No |  |
| `maxPeptides` | `int` | No |  |
| `minMass` | `float` | No |  |
| `missedCleavages` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `protease` | `str` | No |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ProteaseDigestion().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProteaseDigestionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ProteinAnnotatePollEntity

```python
protein_annotate_poll = client.ProteinAnnotatePoll()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `jobId` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ProteinAnnotatePoll().create({
    "jobId": "example_jobId",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProteinAnnotatePollEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ProteinAnnotateSubmitEntity

```python
protein_annotate_submit = client.ProteinAnnotateSubmit()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `appl` | `str` | No |  |
| `gate` | `Any` | No |  |
| `goterms` | `bool` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ProteinAnnotateSubmit().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProteinAnnotateSubmitEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ProteinHydrophobicityEntity

```python
protein_hydrophobicity = client.ProteinHydrophobicity()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `scale` | `str` | No |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |
| `window` | `int` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ProteinHydrophobicity().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProteinHydrophobicityEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ProteinPropertyEntity

```python
protein_property = client.ProteinProperty()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `chargeStep` | `float` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ProteinProperty().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ProteinPropertyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RandomSequenceEntity

```python
random_sequence = client.RandomSequence()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `gcContent` | `float` | No |  |
| `kind` | `str` | No |  |
| `length` | `int` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.RandomSequence().create({
    "length": 1,  # int
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RandomSequenceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RestrictionSiteEntity

```python
restriction_site = client.RestrictionSite()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzymes` | `list` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.RestrictionSite().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RestrictionSiteEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ReverseComplementEntity

```python
reverse_complement = client.ReverseComplement()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |
| `type` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ReverseComplement().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReverseComplementEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ReverseTranslateEntity

```python
reverse_translate = client.ReverseTranslate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `mode` | `str` | No |  |
| `ok` | `Any` | Yes |  |
| `organism` | `str` | No |  |
| `protein` | `str` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ReverseTranslate().create({
    "ok": "example_ok",  # Any
    "protein": "example_protein",  # str
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReverseTranslateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RnaFoldEntity

```python
rna_fold = client.RnaFold()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.RnaFold().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RnaFoldEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SangerVsReferenceEntity

```python
sanger_vs_reference = client.SangerVsReference()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fileBase64` | `str` | No |  |
| `fileName` | `str` | No |  |
| `gate` | `Any` | No |  |
| `minCoverage` | `float` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `read` | `str` | No |  |
| `reference` | `str` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SangerVsReference().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "reference": "example_reference",  # str
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SangerVsReferenceEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SavePermalinkEntity

```python
save_permalink = client.SavePermalink()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `dict` | Yes |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SavePermalink().create({
    "args": {},  # dict
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SavePermalinkEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SeqfileStatEntity

```python
seqfile_stat = client.SeqfileStat()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `input` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `qualityOffset` | `int` | No |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SeqfileStat().create({
    "input": "example_input",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SeqfileStatEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SequenceFetchEntity

```python
sequence_fetch = client.SequenceFetch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `str` | Yes |  |
| `db` | `str` | No |  |
| `format` | `str` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SequenceFetch().create({
    "accession": "example_accession",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SequenceFetchEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SequenceFormatConvertEntity

```python
sequence_format_convert = client.SequenceFormatConvert()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `str` | No |  |
| `gate` | `Any` | No |  |
| `input` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `to` | `str` | No |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SequenceFormatConvert().create({
    "input": "example_input",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SequenceFormatConvertEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SequenceReportEntity

```python
sequence_report = client.SequenceReport()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `endPrimerLength` | `int` | No |  |
| `gate` | `Any` | No |  |
| `maxOrfs` | `int` | No |  |
| `minOrfAa` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SequenceReport().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SequenceReportEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SequenceSearchEntity

```python
sequence_search = client.SequenceSearch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `db` | `str` | No |  |
| `gate` | `Any` | No |  |
| `gene` | `str` | No |  |
| `maxResults` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `organism` | `str` | No |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `term` | `str` | No |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SequenceSearch().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SequenceSearchEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SequencingReadbackVerifyEntity

```python
sequencing_readback_verify = client.SequencingReadbackVerify()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `minSupportingReads` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `reads` | `str` | Yes |  |
| `reference` | `str` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SequencingReadbackVerify().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "reads": "example_reads",  # str
    "reference": "example_reference",  # str
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SequencingReadbackVerifyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SessionCreateEntity

```python
session_create = client.SessionCreate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entries` | `dict` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SessionCreate().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SessionCreateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SessionGetEntity

```python
session_get = client.SessionGet()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `names` | `list` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sessionId` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SessionGet().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sessionId": "example_sessionId",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SessionGetEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SessionRunEntity

```python
session_run = client.SessionRun()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `dict` | No |  |
| `fromSession` | `dict` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sessionId` | `str` | Yes |  |
| `tool` | `str` | Yes |  |
| `writeBack` | `dict` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SessionRun().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sessionId": "example_sessionId",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SessionRunEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SessionSetEntity

```python
session_set = client.SessionSet()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entries` | `dict` | Yes |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sessionId` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SessionSet().create({
    "entries": {},  # dict
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sessionId": "example_sessionId",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SessionSetEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SirnaDesignEntity

```python
sirna_design = client.SirnaDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `minReynolds` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `shRnaLoop` | `str` | No |  |
| `target` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SirnaDesign().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "target": "example_target",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SirnaDesignEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SiteDirectedMutagenesiEntity

```python
site_directed_mutagenesi = client.SiteDirectedMutagenesi()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `float` | No |  |
| `dntpMM` | `float` | No |  |
| `editKind` | `str` | No |  |
| `frameStart` | `int` | No |  |
| `gate` | `Any` | No |  |
| `mgMM` | `float` | No |  |
| `naMM` | `float` | No |  |
| `newBase` | `str` | No |  |
| `ok` | `Any` | Yes |  |
| `oligoNM` | `float` | No |  |
| `organism` | `str` | No |  |
| `position` | `int` | No |  |
| `provenance` | `dict` | Yes |  |
| `residue` | `int` | No |  |
| `result` | `dict` | Yes |  |
| `style` | `str` | No |  |
| `targetAa` | `str` | No |  |
| `template` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SiteDirectedMutagenesi().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "template": "example_template",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SiteDirectedMutagenesiEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TranslateEntity

```python
translate = client.Translate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frame` | `int` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `toStop` | `bool` | No |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Translate().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TranslateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## VariantAnnotateEntity

```python
variant_annotate = client.VariantAnnotate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `assembly` | `str` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |
| `variant` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.VariantAnnotate().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
    "variant": "example_variant",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VariantAnnotateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## VariantComparatorEntity

```python
variant_comparator = client.VariantComparator()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `coding` | `bool` | No |  |
| `frameStart` | `int` | No |  |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `query` | `str` | Yes |  |
| `reference` | `str` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.VariantComparator().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "query": "example_query",  # str
    "reference": "example_reference",  # str
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VariantComparatorEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## VerifyAssemblyEntity

```python
verify_assembly = client.VerifyAssembly()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `float` | No |  |
| `circular` | `bool` | No |  |
| `claimedConstruct` | `str` | Yes |  |
| `coding` | `bool` | No |  |
| `enzyme` | `str` | No |  |
| `enzyme3` | `str` | No |  |
| `enzyme5` | `str` | No |  |
| `fragmentPcrs` | `list` | No |  |
| `fragments` | `list` | No |  |
| `frameStart` | `int` | No |  |
| `gate` | `Any` | No |  |
| `insert` | `str` | No |  |
| `insertPcr` | `dict` | No |  |
| `method` | `str` | Yes |  |
| `names` | `list` | No |  |
| `ok` | `Any` | Yes |  |
| `overlapLen` | `int` | No |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |
| `vector` | `str` | No |  |
| `vectorPcr` | `dict` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.VerifyAssembly().create({
    "claimedConstruct": "example_claimedConstruct",  # str
    "method": "example_method",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VerifyAssemblyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## VerifyConstructEntity

```python
verify_construct = client.VerifyConstruct()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `claimedConstruct` | `str` | Yes |  |
| `expectedFrameStart` | `int` | No |  |
| `gate` | `Any` | No |  |
| `insertForwardPrimer` | `str` | Yes |  |
| `insertReversePrimer` | `str` | Yes |  |
| `insertTemplate` | `str` | Yes |  |
| `maxPrimerMismatches` | `int` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `templateCircular` | `bool` | No |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.VerifyConstruct().create({
    "claimedConstruct": "example_claimedConstruct",  # str
    "insertForwardPrimer": "example_insertForwardPrimer",  # str
    "insertReversePrimer": "example_insertReversePrimer",  # str
    "insertTemplate": "example_insertTemplate",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VerifyConstructEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## VirtualGelEntity

```python
virtual_gel = client.VirtualGel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `bool` | No |  |
| `enzymes` | `list` | No |  |
| `gate` | `Any` | No |  |
| `ladder` | `str` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `sequence` | `str` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.VirtualGel().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VirtualGelEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## VolcanoPlotDataEntity

```python
volcano_plot_data = client.VolcanoPlotData()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes |  |
| `rows` | `list` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.VolcanoPlotData().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "rows": [],  # list
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VolcanoPlotDataEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## WebSearchEntity

```python
web_search = client.WebSearch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `Any` | No |  |
| `max_results` | `float` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `query` | `str` | Yes |  |
| `result` | `dict` | Yes |  |
| `tool` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.WebSearch().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "query": "example_query",  # str
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WebSearchEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = SeqbenchMcpSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

