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
| `accession` | `str` | Yes | UniProt accession, e.g. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `int` | No | Total gapmer length (nt). |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `target` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |
| `wing` | `int` | No | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

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
| `editor` | `str` | No | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | `int` | No | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `target` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | `int` | No | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `args` | `dict` | No | Shared tool arguments applied to every record. |
| `capped` | `bool` | Yes | True if input exceeded the record limit. |
| `columns` | `list` | Yes |  |
| `count` | `int` | Yes |  |
| `errors` | `int` | Yes |  |
| `input` | `str` | Yes | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | `int` | Yes | Maximum records per call (500). |
| `provenance` | `dict` | Yes |  |
| `rows` | `list` | Yes |  |
| `tool` | `str` | Yes | A batchable tool slug (see `GET /batch`). |

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
| `columns` | `list` | Yes | Flattened "<step>·<tool>·<key>" column headers. |
| `count` | `int` | Yes |  |
| `errors` | `int` | Yes |  |
| `input` | `str` | Yes | Multi-FASTA text or one sequence per line. |
| `limit` | `int` | Yes | Maximum records per call (200). |
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
| `endPrimerLength` | `int` | No | Length of the naive end primers taken from each end. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `int` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `int` | No | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `armTmTarget` | `float` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `bool` | No | Produce a circular product. |
| `enzyme` | `str` | No | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | `str` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `str` | No | 5′ enzyme (restriction method). |
| `fragments` | `list` | No | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `str` | No | Insert sequence (restriction method). |
| `method` | `str` | Yes | Assembly method. |
| `names` | `list` | No | Optional labels for each fragment. |
| `ok` | `Any` | Yes |  |
| `overlapLen` | `int` | No | Gibson homology-arm length (bp). |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |
| `vector` | `str` | No | Vector sequence (restriction method). |

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
| `frameStart` | `int` | No | 1-based position to start reading codons. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `organism` | `str` | No |  |
| `provenance` | `dict` | Yes |  |
| `rareThreshold` | `float` | No | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `organism` | `str` | No |  |
| `protein` | `str` | Yes | Protein sequence (one-letter codes). |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `avoidEnzymes` | `list` | No | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | `int` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `int` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `float` | No |  |
| `gcLow` | `float` | No |  |
| `gcWindow` | `int` | No |  |
| `homopolymerMin` | `int` | No |  |
| `maxPasses` | `int` | No | Repeat full passes until clean or no further progress. |
| `ok` | `Any` | Yes |  |
| `organism` | `str` | No | Codon-usage table to prefer among synonymous options. |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `avoidEnzymes` | `list` | No | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | `int` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `int` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `float` | No | GC% above this flags a GC-rich window. |
| `gcLow` | `float` | No | GC% below this flags an AT-rich window. |
| `gcWindow` | `int` | No | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | `int` | No | Minimum run length to flag a homopolymer. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | `float` | No | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | `str` | No | Nuclease id. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `bool` | No | Also scan the reverse strand for guides. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `armLength` | `int` | No | Homology arm length (bp) on each side. |
| `blockPam` | `bool` | No | When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut. |
| `designGenotypingPrimers` | `bool` | No | Also design a primer pair (on the original targetSequence) whose product spans the edit site. |
| `editEnd` | `int` | No | 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. |
| `editStart` | `int` | No | 1-based start of the region being replaced. |
| `frameStart` | `int` | No | Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | `int` | No | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | `int` | No | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | `str` | No | Strand the guide's protospacer is on. |
| `nuclease` | `str` | No | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `replacement` | `str` | Yes | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | `dict` | Yes | Tool-specific output object. |
| `targetSequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | `str` | No | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` | `Any` | Yes |  |
| `protospacer` | `str` | Yes | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequenceA` | `str` | Yes | First oligo (5'→3'). |
| `sequenceB` | `str` | Yes | Second oligo (5'→3'). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `int` | No | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | `float` | No | Mass in nanograms. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | No | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | `str` | Yes | The tool slug that ran. |
| `type` | `str` | No | Molecule type. |
| `volumeUl` | `float` | No | Volume in microlitres (0 = unknown; needed for concentration). |

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
| `enzymeA` | `str` | Yes | First enzyme name (e.g. |
| `enzymeB` | `str` | Yes | Second enzyme name (e.g. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `reactions` | `list` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `protocolName` | `str` | No | Optional protocol name (used in the script's metadata). |
| `provenance` | `dict` | Yes |  |
| `reactions` | `list` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `reactions` | `list` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `clusterCols` | `bool` | No | Cluster (reorder) samples. |
| `clusterRows` | `bool` | No | Cluster (reorder) genes. |
| `distanceMetric` | `str` | No | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `list` | Yes | Row (gene) labels. |
| `linkage` | `str` | No | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `samples` | `list` | Yes | Column (sample) labels. |
| `tool` | `str` | Yes | The tool slug that ran. |
| `values` | `list` | Yes | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | `bool` | No | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `str` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `qualityOffset` | `int` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `str` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | `int` | No | Reads shorter than this after trimming are dropped. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `qualityOffset` | `int` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | `int` | No | 3' quality-trim threshold (Phred score). |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | `int` | No | Minimum protein length (aa) to report. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `requireStop` | `bool` | No | Only report ORFs terminated by a stop codon. |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `convert` | `str` | No | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `reverse` | `bool` | No | Reverse the sequence (no complement). |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | `bool` | No | Remove digits, spaces and gaps (keep letters only). |
| `tool` | `str` | Yes | The tool slug that ran. |
| `width` | `int` | No | Line-wrap width; 0 = single line. |

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
| `background` | `list` | No | Custom background/universe gene symbols. |
| `collections` | `list` | No | Which term collections to test. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `list` | Yes | Query gene symbols (human, e.g. |
| `maxTermSize` | `int` | No | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | `int` | No | Skip terms/pathways with fewer than this many background genes. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `str` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `str` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `str` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `compareToNamedSet` | `str` | No | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | `str` | No | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `overhangs` | `list` | Yes | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `riskThreshold` | `float` | No | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |
| `variant` | `str` | Yes | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `from` | `str` | Yes | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | `list` | Yes | The ids to map, up to 1000 (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `taxId` | `str` | No | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | `str` | Yes | Target id type. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `circular` | `bool` | No | Treat the template as circular (plasmid). |
| `forwardPrimer` | `str` | Yes | Primer 1, 5'→3'. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Mismatches tolerated per primer. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `reversePrimer` | `str` | Yes | Primer 2, 5'→3' (order does not matter). |
| `template` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `addSecondaryMismatch` | `bool` | No | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | `str` | Yes | First allele (single base) — gets the FAM tail. |
| `alleleB` | `str` | Yes | Second allele (single base) — gets the HEX tail. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | `int` | No | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | `int` | No | Minimum amplicon length for the common reverse primer. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `snpPosition` | `int` | Yes | 1-based position of the SNP on the forward strand. |
| `target` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | `float` | No | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `dntpMM` | `float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Any` | Yes |  |
| `oligoNM` | `float` | No | Total strand concentration (nM). |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | `float` | No | Optional target Tm (°C). |
| `tmTolerance` | `float` | No | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Maximum allowed mismatches per match. |
| `motif` | `str` | Yes | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `bool` | No | Also search the reverse strand. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `str` | Yes | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `dntpMM` | `float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Any` | Yes |  |
| `oligoNM` | `float` | No | Total strand concentration (nM). |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sourceSpecies` | `str` | No | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | `list` | Yes | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | `str` | Yes | Ensembl species slug to find homologs in (e.g. |
| `tool` | `str` | Yes | The tool slug that ran. |
| `type` | `str` | No | Homology type to return. |

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
| `gap` | `float` | No | Linear gap penalty (per gap position). |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | `float` | No | Match score. |
| `mismatch` | `float` | No | Mismatch penalty. |
| `mode` | `str` | No |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `seqA` | `str` | Yes | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | `str` | Yes | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `text` | `str` | Yes | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `fileBase64` | `str` | Yes | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `str` | No | Optional original file name (echoed back). |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `circular` | `bool` | No | Treat the sequence as a circular plasmid (vs. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `circular` | `bool` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |
| `topN` | `int` | No | How many top-ranked backbone candidates to report. |

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
| `circular` | `bool` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |
| `topN` | `int` | No | How many top-ranked backbone candidates to report. |

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
| `editEnd` | `int` | Yes | 1-based inclusive end of the region being changed. |
| `editStart` | `int` | Yes | 1-based inclusive start of the region being changed. |
| `frameStart` | `int` | No | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | `str` | No | Replacement bases (forward strand). |
| `ok` | `Any` | Yes |  |
| `pbsLength` | `int` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `rttHomology` | `int` | No | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | `str` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | `str` | Yes | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` | `Any` | Yes |  |
| `overlapLength` | `int` | No | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | `int` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `dict` | Yes |  |
| `replaceEnd` | `int` | Yes | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | `int` | Yes | 1-based inclusive start of the region being replaced/deleted. |
| `result` | `dict` | Yes | Tool-specific output object. |
| `target` | `str` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `dntpMM` | `float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` | `float` | No |  |
| `gcMin` | `float` | No |  |
| `lenMax` | `int` | No |  |
| `lenMin` | `int` | No |  |
| `lenOpt` | `int` | No |  |
| `maxReturn` | `int` | No | Number of best pairs to return. |
| `mgMM` | `float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Any` | Yes |  |
| `oligoNM` | `float` | No | Total strand concentration (nM). |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `targetEnd` | `int` | No | 1-based inclusive end of the target region (optional). |
| `targetStart` | `int` | No | 1-based inclusive start of a region the product must span (optional). |
| `template` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` | `float` | No |  |
| `tmMaxDiff` | `float` | No | Max Tm difference within a pair (°C). |
| `tmMin` | `float` | No |  |
| `tmOpt` | `float` | No |  |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `forwardPrimer` | `str` | Yes | Forward primer, 5'→3'. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | No | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | `int` | No | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `reversePrimer` | `str` | Yes | Reverse primer, 5'→3'. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | `float` | No | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | `int` | No | Cap on the number of returned peptides. |
| `minMass` | `float` | No | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | `int` | No | Allowed missed internal cleavages (0–2). |
| `ok` | `Any` | Yes |  |
| `protease` | `str` | No | Protease or chemical cleavage agent. |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `str` | Yes |  |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `appl` | `str` | No | Restrict to one member database (e.g. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | `bool` | No | Include GO-term cross-references. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `scale` | `str` | No | Amino-acid scale. |
| `sequence` | `str` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `str` | Yes | The tool slug that ran. |
| `window` | `int` | No | Sliding-window size (clamped to an odd number ≥ 1). |

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
| `chargeStep` | `float` | No | pH step for the net-charge titration curve (0–14). |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | `float` | No | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` | `str` | No |  |
| `length` | `int` | Yes | Number of residues to generate. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `enzymes` | `list` | No | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |
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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` | `str` | No |  |
| `ok` | `Any` | Yes |  |
| `organism` | `str` | No | Codon-usage host (ignored in degenerate mode). |
| `protein` | `str` | Yes | Protein sequence (one-letter codes; * for stop). |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `fileBase64` | `str` | No | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `str` | No | Optional original file name (echoed back). |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | `float` | No | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `read` | `str` | No | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | `str` | Yes | Expected reference sequence (FASTA or raw). |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `args` | `dict` | Yes | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `str` | Yes | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `qualityOffset` | `int` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `accession` | `str` | Yes | GenBank/RefSeq accession (e.g. |
| `db` | `str` | No | Database to query; auto-detects from the accession format. |
| `format` | `str` | No | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `from` | `str` | No | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `str` | Yes | A FASTA or GenBank record to convert. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `to` | `str` | No | Output format. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `endPrimerLength` | `int` | No | Length of the naive end primers taken from each end. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `int` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `int` | No | Minimum ORF length in amino acids. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `str` | No | Gene symbol/name, e.g. |
| `maxResults` | `int` | No | Up to 20. |
| `ok` | `Any` | Yes |  |
| `organism` | `str` | No | Organism name, e.g. |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `term` | `str` | No | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | `int` | No | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `reads` | `str` | Yes | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | `str` | Yes | The claimed/expected reference sequence. |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `entries` | `dict` | No | Initial named entries, e.g. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | `list` | No | Only return these entries; omit to return all of them. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sessionId` | `str` | Yes |  |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `args` | `dict` | No | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | `dict` | No | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sessionId` | `str` | Yes |  |
| `tool` | `str` | Yes | The tool slug that ran. |
| `writeBack` | `dict` | No | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

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
| `entries` | `dict` | Yes | Named entries to add/overwrite, e.g. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sessionId` | `str` | Yes |  |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | `int` | No | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `shRnaLoop` | `str` | No | Loop sequence used when assembling the shRNA cassette. |
| `target` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `armTmTarget` | `float` | No | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | `float` | No | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | `str` | No | Edit at the nucleotide or amino-acid level. |
| `frameStart` | `int` | No | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | No | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | `str` | No | Replacement base (editKind='nt'). |
| `ok` | `Any` | Yes |  |
| `oligoNM` | `float` | No | Total strand concentration (nM). |
| `organism` | `str` | No | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | `int` | No | 1-based position to substitute (editKind='nt'). |
| `provenance` | `dict` | Yes |  |
| `residue` | `int` | No | 1-based residue number to change (editKind='aa'). |
| `result` | `dict` | Yes | Tool-specific output object. |
| `style` | `str` | No | Mutagenic primer style. |
| `targetAa` | `str` | No | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | `bool` | No | Stop at the first stop codon. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `assembly` | `str` | No | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |
| `variant` | `str` | Yes | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

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
| `coding` | `bool` | No | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | `int` | No | 1-based reading-frame start (used when coding is true). |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `query` | `str` | Yes | Query / variant sequence (raw or FASTA). |
| `reference` | `str` | Yes | Reference / wild-type sequence (raw or FASTA). |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `armTmTarget` | `float` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `bool` | No | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | `str` | Yes | The sequence you claim you ended up with. |
| `coding` | `bool` | No | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | `str` | No | Type IIS enzyme for Golden Gate. |
| `enzyme3` | `str` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `str` | No | 5′ enzyme (restriction method). |
| `fragmentPcrs` | `list` | No | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | `list` | No | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | `int` | No | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `str` | No | Insert sequence (restriction method). |
| `insertPcr` | `dict` | No | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | `str` | Yes | Assembly method used. |
| `names` | `list` | No | Optional labels for each fragment. |
| `ok` | `Any` | Yes |  |
| `overlapLen` | `int` | No | Gibson homology-arm length (bp). |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |
| `vector` | `str` | No | Vector sequence (restriction method). |
| `vectorPcr` | `dict` | No | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

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
| `claimedConstruct` | `str` | Yes | The final sequence claimed to have been built. |
| `expectedFrameStart` | `int` | No | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | `str` | Yes | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | `str` | Yes | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | `str` | Yes | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | `int` | No | Mismatches tolerated per primer during PCR prediction. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `templateCircular` | `bool` | No | Treat insertTemplate as circular (e.g. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `circular` | `bool` | No | Treat the sequence as circular (plasmid). |
| `enzymes` | `list` | No | Enzyme names to digest with. |
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | `str` | No | DNA ladder to plot alongside the sample lane. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `sequence` | `str` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `result` | `dict` | Yes | Tool-specific output object. |
| `rows` | `list` | Yes | Differential expression rows, one per gene. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `gate` | `Any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | `float` | No | Maximum number of results to return (default 5, max 10). |
| `ok` | `Any` | Yes |  |
| `provenance` | `dict` | Yes |  |
| `query` | `str` | Yes | The search query. |
| `result` | `dict` | Yes | Tool-specific output object. |
| `tool` | `str` | Yes | The tool slug that ran. |

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
| `ratelimit` | 0.0.1 | Client-side rate limiting via a token bucket |
| `retry` | 0.0.1 | Automatic retry of transient failures with exponential backoff |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |
| `timeout` | 0.0.1 | Per-request timeout with transport abort |


Features are activated via the `feature` option:

```python
client = SeqbenchMcpSDK({
    "feature": {
        "ratelimit": {"active": True},
        "retry": {"active": True},
        "test": {"active": True},
        "timeout": {"active": True},
    },
})
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

