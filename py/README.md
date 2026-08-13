# SeqbenchMcp Python SDK



The Python SDK for the SeqbenchMcp API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.AlphafoldLookup()` — each
carrying a small, uniform set of operations (`load`, `create`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/seqbench-mcp-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from seqbenchmcp_sdk import SeqbenchMcpSDK

client = SeqbenchMcpSDK({
    "apikey": os.environ.get("SEQBENCH_MCP_APIKEY"),
})
```

### 4. Create, update, and remove

```python
# Create — returns the ENTITY (call data_get() for the record)
created = client.AlphafoldLookup().create({"accession": "example_accession", "ok": "example_ok", "provenance": {}, "result": {}, "tool": "example_tool"})

```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    batch = client.Batch().load()
    print(batch)
except Exception as err:
    print(f"load failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = SeqbenchMcpSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
batch = client.Batch().load()
# batch contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = SeqbenchMcpSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
SEQBENCH_MCP_TEST_LIVE=TRUE
SEQBENCH_MCP_APIKEY=<your-key>
```

Then run:

```bash
cd py && pytest test/
```


## Reference

### SeqbenchMcpSDK

```python
from seqbenchmcp_sdk import SeqbenchMcpSDK

client = SeqbenchMcpSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = SeqbenchMcpSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### SeqbenchMcpSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
| `AlphafoldLookup` | `(data) -> AlphafoldLookupEntity` | Create an AlphafoldLookup entity instance. |
| `AsoDesign` | `(data) -> AsoDesignEntity` | Create an AsoDesign entity instance. |
| `BaseEditingDesign` | `(data) -> BaseEditingDesignEntity` | Create a BaseEditingDesign entity instance. |
| `Batch` | `(data) -> BatchEntity` | Create a Batch entity instance. |
| `BatchWorkflow` | `(data) -> BatchWorkflowEntity` | Create a BatchWorkflow entity instance. |
| `CharacterizeSequence` | `(data) -> CharacterizeSequenceEntity` | Create a CharacterizeSequence entity instance. |
| `CloningSimulate` | `(data) -> CloningSimulateEntity` | Create a CloningSimulate entity instance. |
| `CodonAdaptationIndex` | `(data) -> CodonAdaptationIndexEntity` | Create a CodonAdaptationIndex entity instance. |
| `CodonOptimize` | `(data) -> CodonOptimizeEntity` | Create a CodonOptimize entity instance. |
| `ConstructAutofix` | `(data) -> ConstructAutofixEntity` | Create a ConstructAutofix entity instance. |
| `ConstructQc` | `(data) -> ConstructQcEntity` | Create a ConstructQc entity instance. |
| `CrisprGrnaDesign` | `(data) -> CrisprGrnaDesignEntity` | Create a CrisprGrnaDesign entity instance. |
| `CrisprHdrDonor` | `(data) -> CrisprHdrDonorEntity` | Create a CrisprHdrDonor entity instance. |
| `CrisprOfftargetCheck` | `(data) -> CrisprOfftargetCheckEntity` | Create a CrisprOfftargetCheck entity instance. |
| `CrossDimer` | `(data) -> CrossDimerEntity` | Create a CrossDimer entity instance. |
| `DnaMolarity` | `(data) -> DnaMolarityEntity` | Create a DnaMolarity entity instance. |
| `DoubleDigest` | `(data) -> DoubleDigestEntity` | Create a DoubleDigest entity instance. |
| `ExportEchoPicklist` | `(data) -> ExportEchoPicklistEntity` | Create an ExportEchoPicklist entity instance. |
| `ExportOpentronsProtocol` | `(data) -> ExportOpentronsProtocolEntity` | Create an ExportOpentronsProtocol entity instance. |
| `ExportPlateLayout` | `(data) -> ExportPlateLayoutEntity` | Create an ExportPlateLayout entity instance. |
| `ExpressionHeatmapCluster` | `(data) -> ExpressionHeatmapClusterEntity` | Create an ExpressionHeatmapCluster entity instance. |
| `FastqQcReport` | `(data) -> FastqQcReportEntity` | Create a FastqQcReport entity instance. |
| `FastqTrim` | `(data) -> FastqTrimEntity` | Create a FastqTrim entity instance. |
| `FindOrf` | `(data) -> FindOrfEntity` | Create a FindOrf entity instance. |
| `FormatSequence` | `(data) -> FormatSequenceEntity` | Create a FormatSequence entity instance. |
| `FunctionalEnrichment` | `(data) -> FunctionalEnrichmentEntity` | Create a FunctionalEnrichment entity instance. |
| `GcContent` | `(data) -> GcContentEntity` | Create a GcContent entity instance. |
| `GeneDossier` | `(data) -> GeneDossierEntity` | Create a GeneDossier entity instance. |
| `GeneExpression` | `(data) -> GeneExpressionEntity` | Create a GeneExpression entity instance. |
| `GeneModel` | `(data) -> GeneModelEntity` | Create a GeneModel entity instance. |
| `GoldenGateFidelity` | `(data) -> GoldenGateFidelityEntity` | Create a GoldenGateFidelity entity instance. |
| `HgvsConvert` | `(data) -> HgvsConvertEntity` | Create a HgvsConvert entity instance. |
| `IdMapPoll` | `(data) -> IdMapPollEntity` | Create an IdMapPoll entity instance. |
| `IdMapSubmit` | `(data) -> IdMapSubmitEntity` | Create an IdMapSubmit entity instance. |
| `InSilicoPcr` | `(data) -> InSilicoPcrEntity` | Create an InSilicoPcr entity instance. |
| `KaspPrimerDesign` | `(data) -> KaspPrimerDesignEntity` | Create a KaspPrimerDesign entity instance. |
| `ListTool` | `(data) -> ListToolEntity` | Create a ListTool entity instance. |
| `MeltingTemperature` | `(data) -> MeltingTemperatureEntity` | Create a MeltingTemperature entity instance. |
| `MotifFinder` | `(data) -> MotifFinderEntity` | Create a MotifFinder entity instance. |
| `MultipleSequenceAlignment` | `(data) -> MultipleSequenceAlignmentEntity` | Create a MultipleSequenceAlignment entity instance. |
| `OligoAnalysi` | `(data) -> OligoAnalysiEntity` | Create an OligoAnalysi entity instance. |
| `OrthologMap` | `(data) -> OrthologMapEntity` | Create an OrthologMap entity instance. |
| `PairwiseAlignment` | `(data) -> PairwiseAlignmentEntity` | Create a PairwiseAlignment entity instance. |
| `ParseGenbank` | `(data) -> ParseGenbankEntity` | Create a ParseGenbank entity instance. |
| `ParseSangerTrace` | `(data) -> ParseSangerTraceEntity` | Create a ParseSangerTrace entity instance. |
| `PlasmidAnnotate` | `(data) -> PlasmidAnnotateEntity` | Create a PlasmidAnnotate entity instance. |
| `PlasmidDeepAnnotate` | `(data) -> PlasmidDeepAnnotateEntity` | Create a PlasmidDeepAnnotate entity instance. |
| `PlasmidFullReport` | `(data) -> PlasmidFullReportEntity` | Create a PlasmidFullReport entity instance. |
| `PlasmidIdentify` | `(data) -> PlasmidIdentifyEntity` | Create a PlasmidIdentify entity instance. |
| `PrimeEditingDesign` | `(data) -> PrimeEditingDesignEntity` | Create a PrimeEditingDesign entity instance. |
| `PrimeEditingTwinDesign` | `(data) -> PrimeEditingTwinDesignEntity` | Create a PrimeEditingTwinDesign entity instance. |
| `PrimerDesign` | `(data) -> PrimerDesignEntity` | Create a PrimerDesign entity instance. |
| `PrimerSpecificity` | `(data) -> PrimerSpecificityEntity` | Create a PrimerSpecificity entity instance. |
| `ProteaseDigestion` | `(data) -> ProteaseDigestionEntity` | Create a ProteaseDigestion entity instance. |
| `ProteinAnnotatePoll` | `(data) -> ProteinAnnotatePollEntity` | Create a ProteinAnnotatePoll entity instance. |
| `ProteinAnnotateSubmit` | `(data) -> ProteinAnnotateSubmitEntity` | Create a ProteinAnnotateSubmit entity instance. |
| `ProteinHydrophobicity` | `(data) -> ProteinHydrophobicityEntity` | Create a ProteinHydrophobicity entity instance. |
| `ProteinProperty` | `(data) -> ProteinPropertyEntity` | Create a ProteinProperty entity instance. |
| `RandomSequence` | `(data) -> RandomSequenceEntity` | Create a RandomSequence entity instance. |
| `RestrictionSite` | `(data) -> RestrictionSiteEntity` | Create a RestrictionSite entity instance. |
| `ReverseComplement` | `(data) -> ReverseComplementEntity` | Create a ReverseComplement entity instance. |
| `ReverseTranslate` | `(data) -> ReverseTranslateEntity` | Create a ReverseTranslate entity instance. |
| `RnaFold` | `(data) -> RnaFoldEntity` | Create a RnaFold entity instance. |
| `SangerVsReference` | `(data) -> SangerVsReferenceEntity` | Create a SangerVsReference entity instance. |
| `SavePermalink` | `(data) -> SavePermalinkEntity` | Create a SavePermalink entity instance. |
| `SeqfileStat` | `(data) -> SeqfileStatEntity` | Create a SeqfileStat entity instance. |
| `SequenceFetch` | `(data) -> SequenceFetchEntity` | Create a SequenceFetch entity instance. |
| `SequenceFormatConvert` | `(data) -> SequenceFormatConvertEntity` | Create a SequenceFormatConvert entity instance. |
| `SequenceReport` | `(data) -> SequenceReportEntity` | Create a SequenceReport entity instance. |
| `SequenceSearch` | `(data) -> SequenceSearchEntity` | Create a SequenceSearch entity instance. |
| `SequencingReadbackVerify` | `(data) -> SequencingReadbackVerifyEntity` | Create a SequencingReadbackVerify entity instance. |
| `SessionCreate` | `(data) -> SessionCreateEntity` | Create a SessionCreate entity instance. |
| `SessionGet` | `(data) -> SessionGetEntity` | Create a SessionGet entity instance. |
| `SessionRun` | `(data) -> SessionRunEntity` | Create a SessionRun entity instance. |
| `SessionSet` | `(data) -> SessionSetEntity` | Create a SessionSet entity instance. |
| `SirnaDesign` | `(data) -> SirnaDesignEntity` | Create a SirnaDesign entity instance. |
| `SiteDirectedMutagenesi` | `(data) -> SiteDirectedMutagenesiEntity` | Create a SiteDirectedMutagenesi entity instance. |
| `Translate` | `(data) -> TranslateEntity` | Create a Translate entity instance. |
| `VariantAnnotate` | `(data) -> VariantAnnotateEntity` | Create a VariantAnnotate entity instance. |
| `VariantComparator` | `(data) -> VariantComparatorEntity` | Create a VariantComparator entity instance. |
| `VerifyAssembly` | `(data) -> VerifyAssemblyEntity` | Create a VerifyAssembly entity instance. |
| `VerifyConstruct` | `(data) -> VerifyConstructEntity` | Create a VerifyConstruct entity instance. |
| `VirtualGel` | `(data) -> VirtualGelEntity` | Create a VirtualGel entity instance. |
| `VolcanoPlotData` | `(data) -> VolcanoPlotDataEntity` | Create a VolcanoPlotData entity instance. |
| `WebSearch` | `(data) -> WebSearchEntity` | Create a WebSearch entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

### Entities

#### AlphafoldLookup

| Field | Description |
| --- | --- |
| `accession` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/alphafold_lookup`

#### AsoDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `length` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `target` |  |
| `tool` |  |
| `wing` |  |

Operations: Create.

API path: `/aso_design`

#### BaseEditingDesign

| Field | Description |
| --- | --- |
| `editor` |  |
| `frameStart` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `target` |  |
| `targetPosition` |  |
| `tool` |  |

Operations: Create.

API path: `/base_editing_design`

#### Batch

| Field | Description |
| --- | --- |
| `args` |  |
| `capped` |  |
| `columns` |  |
| `count` |  |
| `errors` |  |
| `input` |  |
| `limit` |  |
| `provenance` |  |
| `rows` |  |
| `tool` |  |

Operations: Create, Load.

API path: `/batch`

#### BatchWorkflow

| Field | Description |
| --- | --- |
| `capped` |  |
| `columns` |  |
| `count` |  |
| `errors` |  |
| `input` |  |
| `limit` |  |
| `provenance` |  |
| `rows` |  |
| `steps` |  |

Operations: Create, Load.

API path: `/workflow`

#### CharacterizeSequence

| Field | Description |
| --- | --- |
| `endPrimerLength` |  |
| `gate` |  |
| `maxOrfs` |  |
| `minOrfAa` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/characterize_sequence`

#### CloningSimulate

| Field | Description |
| --- | --- |
| `armTmTarget` |  |
| `circular` |  |
| `enzyme` |  |
| `enzyme3` |  |
| `enzyme5` |  |
| `fragments` |  |
| `gate` |  |
| `insert` |  |
| `method` |  |
| `names` |  |
| `ok` |  |
| `overlapLen` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `vector` |  |

Operations: Create.

API path: `/cloning_simulate`

#### CodonAdaptationIndex

| Field | Description |
| --- | --- |
| `frameStart` |  |
| `gate` |  |
| `ok` |  |
| `organism` |  |
| `provenance` |  |
| `rareThreshold` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/codon_adaptation_index`

#### CodonOptimize

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `organism` |  |
| `protein` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/codon_optimize`

#### ConstructAutofix

| Field | Description |
| --- | --- |
| `avoidEnzymes` |  |
| `crypticOrfMinAa` |  |
| `frameStart` |  |
| `gate` |  |
| `gcHigh` |  |
| `gcLow` |  |
| `gcWindow` |  |
| `homopolymerMin` |  |
| `maxPasses` |  |
| `ok` |  |
| `organism` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/construct_autofix`

#### ConstructQc

| Field | Description |
| --- | --- |
| `avoidEnzymes` |  |
| `crypticOrfMinAa` |  |
| `frameStart` |  |
| `gate` |  |
| `gcHigh` |  |
| `gcLow` |  |
| `gcWindow` |  |
| `homopolymerMin` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/construct_qc`

#### CrisprGrnaDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `minScore` |  |
| `nuclease` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `searchReverseStrand` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/crispr_grna_design`

#### CrisprHdrDonor

| Field | Description |
| --- | --- |
| `armLength` |  |
| `blockPam` |  |
| `designGenotypingPrimers` |  |
| `editEnd` |  |
| `editStart` |  |
| `frameStart` |  |
| `gate` |  |
| `guideEnd` |  |
| `guideStart` |  |
| `guideStrand` |  |
| `nuclease` |  |
| `ok` |  |
| `provenance` |  |
| `replacement` |  |
| `result` |  |
| `targetSequence` |  |
| `tool` |  |

Operations: Create.

API path: `/crispr_hdr_donor`

#### CrisprOfftargetCheck

| Field | Description |
| --- | --- |
| `gate` |  |
| `maxMismatches` |  |
| `nuclease` |  |
| `ok` |  |
| `protospacer` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/crispr_offtarget_check`

#### CrossDimer

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequenceA` |  |
| `sequenceB` |  |
| `tool` |  |

Operations: Create.

API path: `/cross_dimer`

#### DnaMolarity

| Field | Description |
| --- | --- |
| `gate` |  |
| `length` |  |
| `massNg` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |
| `type` |  |
| `volumeUl` |  |

Operations: Create.

API path: `/dna_molarity`

#### DoubleDigest

| Field | Description |
| --- | --- |
| `enzymeA` |  |
| `enzymeB` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/double_digest`

#### ExportEchoPicklist

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `reactions` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/export_echo_picklist`

#### ExportOpentronsProtocol

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `protocolName` |  |
| `provenance` |  |
| `reactions` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/export_opentrons_protocol`

#### ExportPlateLayout

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `reactions` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/export_plate_layout`

#### ExpressionHeatmapCluster

| Field | Description |
| --- | --- |
| `clusterCols` |  |
| `clusterRows` |  |
| `distanceMetric` |  |
| `gate` |  |
| `genes` |  |
| `linkage` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `samples` |  |
| `tool` |  |
| `values` |  |
| `zScoreRows` |  |

Operations: Create.

API path: `/expression_heatmap_cluster`

#### FastqQcReport

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `ok` |  |
| `provenance` |  |
| `qualityOffset` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/fastq_qc_report`

#### FastqTrim

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `minLength` |  |
| `ok` |  |
| `provenance` |  |
| `qualityOffset` |  |
| `qualityThreshold` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/fastq_trim`

#### FindOrf

| Field | Description |
| --- | --- |
| `gate` |  |
| `minAaLength` |  |
| `ok` |  |
| `provenance` |  |
| `requireStop` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/find_orfs`

#### FormatSequence

| Field | Description |
| --- | --- |
| `caseMode` |  |
| `convert` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reverse` |  |
| `sequence` |  |
| `stripNonLetters` |  |
| `tool` |  |
| `width` |  |

Operations: Create.

API path: `/format_sequence`

#### FunctionalEnrichment

| Field | Description |
| --- | --- |
| `background` |  |
| `collections` |  |
| `gate` |  |
| `genes` |  |
| `maxTermSize` |  |
| `minTermSize` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/functional_enrichment`

#### GcContent

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/gc_content`

#### GeneDossier

| Field | Description |
| --- | --- |
| `gate` |  |
| `gene` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/gene_dossier`

#### GeneExpression

| Field | Description |
| --- | --- |
| `gate` |  |
| `gene` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/gene_expression`

#### GeneModel

| Field | Description |
| --- | --- |
| `gate` |  |
| `gene` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/gene_model`

#### GoldenGateFidelity

| Field | Description |
| --- | --- |
| `compareToNamedSet` |  |
| `dataset` |  |
| `gate` |  |
| `ok` |  |
| `overhangs` |  |
| `provenance` |  |
| `result` |  |
| `riskThreshold` |  |
| `tool` |  |

Operations: Create.

API path: `/golden_gate_fidelity`

#### HgvsConvert

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `variant` |  |

Operations: Create.

API path: `/hgvs_convert`

#### IdMapPoll

| Field | Description |
| --- | --- |
| `gate` |  |
| `jobId` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/id_map_poll`

#### IdMapSubmit

| Field | Description |
| --- | --- |
| `from` |  |
| `gate` |  |
| `ids` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `taxId` |  |
| `to` |  |
| `tool` |  |

Operations: Create.

API path: `/id_map_submit`

#### InSilicoPcr

| Field | Description |
| --- | --- |
| `circular` |  |
| `forwardPrimer` |  |
| `gate` |  |
| `maxMismatches` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reversePrimer` |  |
| `template` |  |
| `tool` |  |

Operations: Create.

API path: `/in_silico_pcr`

#### KaspPrimerDesign

| Field | Description |
| --- | --- |
| `addSecondaryMismatch` |  |
| `alleleA` |  |
| `alleleB` |  |
| `gate` |  |
| `maxAmplicon` |  |
| `minAmplicon` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `snpPosition` |  |
| `target` |  |
| `targetCoreTm` |  |
| `tool` |  |

Operations: Create.

API path: `/kasp_primer_design`

#### ListTool

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/`

#### MeltingTemperature

| Field | Description |
| --- | --- |
| `dntpMM` |  |
| `gate` |  |
| `mgMM` |  |
| `naMM` |  |
| `ok` |  |
| `oligoNM` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `targetTm` |  |
| `tmTolerance` |  |
| `tool` |  |

Operations: Create.

API path: `/melting_temperature`

#### MotifFinder

| Field | Description |
| --- | --- |
| `gate` |  |
| `maxMismatches` |  |
| `motif` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `searchReverseStrand` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/motif_finder`

#### MultipleSequenceAlignment

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/multiple_sequence_alignment`

#### OligoAnalysi

| Field | Description |
| --- | --- |
| `dntpMM` |  |
| `gate` |  |
| `mgMM` |  |
| `naMM` |  |
| `ok` |  |
| `oligoNM` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/oligo_analysis`

#### OrthologMap

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sourceSpecies` |  |
| `symbols` |  |
| `targetSpecies` |  |
| `tool` |  |
| `type` |  |

Operations: Create.

API path: `/ortholog_map`

#### PairwiseAlignment

| Field | Description |
| --- | --- |
| `gap` |  |
| `gate` |  |
| `match` |  |
| `mismatch` |  |
| `mode` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `seqA` |  |
| `seqB` |  |
| `tool` |  |

Operations: Create.

API path: `/pairwise_alignment`

#### ParseGenbank

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `text` |  |
| `tool` |  |

Operations: Create.

API path: `/parse_genbank`

#### ParseSangerTrace

| Field | Description |
| --- | --- |
| `fileBase64` |  |
| `fileName` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/parse_sanger_trace`

#### PlasmidAnnotate

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/plasmid_annotate`

#### PlasmidDeepAnnotate

| Field | Description |
| --- | --- |
| `circular` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/plasmid_deep_annotate`

#### PlasmidFullReport

| Field | Description |
| --- | --- |
| `circular` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |
| `topN` |  |

Operations: Create.

API path: `/plasmid_full_report`

#### PlasmidIdentify

| Field | Description |
| --- | --- |
| `circular` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |
| `topN` |  |

Operations: Create.

API path: `/plasmid_identify`

#### PrimeEditingDesign

| Field | Description |
| --- | --- |
| `editEnd` |  |
| `editStart` |  |
| `frameStart` |  |
| `gate` |  |
| `insertedSeq` |  |
| `ok` |  |
| `pbsLength` |  |
| `provenance` |  |
| `result` |  |
| `rttHomology` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/prime_editing_design`

#### PrimeEditingTwinDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `newSequence` |  |
| `ok` |  |
| `overlapLength` |  |
| `pbsLength` |  |
| `provenance` |  |
| `replaceEnd` |  |
| `replaceStart` |  |
| `result` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/prime_editing_twin_design`

#### PrimerDesign

| Field | Description |
| --- | --- |
| `ampliconMax` |  |
| `ampliconMin` |  |
| `dntpMM` |  |
| `gate` |  |
| `gcMax` |  |
| `gcMin` |  |
| `lenMax` |  |
| `lenMin` |  |
| `lenOpt` |  |
| `maxReturn` |  |
| `mgMM` |  |
| `naMM` |  |
| `ok` |  |
| `oligoNM` |  |
| `provenance` |  |
| `result` |  |
| `targetEnd` |  |
| `targetStart` |  |
| `template` |  |
| `tmMax` |  |
| `tmMaxDiff` |  |
| `tmMin` |  |
| `tmOpt` |  |
| `tool` |  |

Operations: Create.

API path: `/primer_design`

#### PrimerSpecificity

| Field | Description |
| --- | --- |
| `forwardPrimer` |  |
| `gate` |  |
| `maxMismatches` |  |
| `maxProductLength` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reversePrimer` |  |
| `tool` |  |

Operations: Create.

API path: `/primer_specificity`

#### ProteaseDigestion

| Field | Description |
| --- | --- |
| `gate` |  |
| `maxMass` |  |
| `maxPeptides` |  |
| `minMass` |  |
| `missedCleavages` |  |
| `ok` |  |
| `protease` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/protease_digestion`

#### ProteinAnnotatePoll

| Field | Description |
| --- | --- |
| `gate` |  |
| `jobId` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/protein_annotate_poll`

#### ProteinAnnotateSubmit

| Field | Description |
| --- | --- |
| `appl` |  |
| `gate` |  |
| `goterms` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/protein_annotate_submit`

#### ProteinHydrophobicity

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `scale` |  |
| `sequence` |  |
| `tool` |  |
| `window` |  |

Operations: Create.

API path: `/protein_hydrophobicity`

#### ProteinProperty

| Field | Description |
| --- | --- |
| `chargeStep` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/protein_properties`

#### RandomSequence

| Field | Description |
| --- | --- |
| `gate` |  |
| `gcContent` |  |
| `kind` |  |
| `length` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/random_sequence`

#### RestrictionSite

| Field | Description |
| --- | --- |
| `enzymes` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/restriction_sites`

#### ReverseComplement

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |
| `type` |  |

Operations: Create.

API path: `/reverse_complement`

#### ReverseTranslate

| Field | Description |
| --- | --- |
| `gate` |  |
| `mode` |  |
| `ok` |  |
| `organism` |  |
| `protein` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/reverse_translate`

#### RnaFold

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/rna_fold`

#### SangerVsReference

| Field | Description |
| --- | --- |
| `fileBase64` |  |
| `fileName` |  |
| `gate` |  |
| `minCoverage` |  |
| `ok` |  |
| `provenance` |  |
| `read` |  |
| `reference` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/sanger_vs_reference`

#### SavePermalink

| Field | Description |
| --- | --- |
| `args` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/save_permalink`

#### SeqfileStat

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `ok` |  |
| `provenance` |  |
| `qualityOffset` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/seqfile_stats`

#### SequenceFetch

| Field | Description |
| --- | --- |
| `accession` |  |
| `db` |  |
| `format` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/sequence_fetch`

#### SequenceFormatConvert

| Field | Description |
| --- | --- |
| `from` |  |
| `gate` |  |
| `input` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `to` |  |
| `tool` |  |

Operations: Create.

API path: `/sequence_format_convert`

#### SequenceReport

| Field | Description |
| --- | --- |
| `endPrimerLength` |  |
| `gate` |  |
| `maxOrfs` |  |
| `minOrfAa` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/sequence_report`

#### SequenceSearch

| Field | Description |
| --- | --- |
| `db` |  |
| `gate` |  |
| `gene` |  |
| `maxResults` |  |
| `ok` |  |
| `organism` |  |
| `provenance` |  |
| `result` |  |
| `term` |  |
| `tool` |  |

Operations: Create.

API path: `/sequence_search`

#### SequencingReadbackVerify

| Field | Description |
| --- | --- |
| `gate` |  |
| `minSupportingReads` |  |
| `ok` |  |
| `provenance` |  |
| `reads` |  |
| `reference` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/sequencing_readback_verify`

#### SessionCreate

| Field | Description |
| --- | --- |
| `entries` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/session_create`

#### SessionGet

| Field | Description |
| --- | --- |
| `gate` |  |
| `names` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sessionId` |  |
| `tool` |  |

Operations: Create.

API path: `/session_get`

#### SessionRun

| Field | Description |
| --- | --- |
| `args` |  |
| `fromSession` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sessionId` |  |
| `tool` |  |
| `writeBack` |  |

Operations: Create.

API path: `/session_run`

#### SessionSet

| Field | Description |
| --- | --- |
| `entries` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sessionId` |  |
| `tool` |  |

Operations: Create.

API path: `/session_set`

#### SirnaDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `minReynolds` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `shRnaLoop` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/sirna_design`

#### SiteDirectedMutagenesi

| Field | Description |
| --- | --- |
| `armTmTarget` |  |
| `dntpMM` |  |
| `editKind` |  |
| `frameStart` |  |
| `gate` |  |
| `mgMM` |  |
| `naMM` |  |
| `newBase` |  |
| `ok` |  |
| `oligoNM` |  |
| `organism` |  |
| `position` |  |
| `provenance` |  |
| `residue` |  |
| `result` |  |
| `style` |  |
| `targetAa` |  |
| `template` |  |
| `tool` |  |

Operations: Create.

API path: `/site_directed_mutagenesis`

#### Translate

| Field | Description |
| --- | --- |
| `frame` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `toStop` |  |
| `tool` |  |

Operations: Create.

API path: `/translate`

#### VariantAnnotate

| Field | Description |
| --- | --- |
| `assembly` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `variant` |  |

Operations: Create.

API path: `/variant_annotate`

#### VariantComparator

| Field | Description |
| --- | --- |
| `coding` |  |
| `frameStart` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `query` |  |
| `reference` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/variant_comparator`

#### VerifyAssembly

| Field | Description |
| --- | --- |
| `armTmTarget` |  |
| `circular` |  |
| `claimedConstruct` |  |
| `coding` |  |
| `enzyme` |  |
| `enzyme3` |  |
| `enzyme5` |  |
| `fragmentPcrs` |  |
| `fragments` |  |
| `frameStart` |  |
| `gate` |  |
| `insert` |  |
| `insertPcr` |  |
| `method` |  |
| `names` |  |
| `ok` |  |
| `overlapLen` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `vector` |  |
| `vectorPcr` |  |

Operations: Create.

API path: `/verify_assembly`

#### VerifyConstruct

| Field | Description |
| --- | --- |
| `claimedConstruct` |  |
| `expectedFrameStart` |  |
| `gate` |  |
| `insertForwardPrimer` |  |
| `insertReversePrimer` |  |
| `insertTemplate` |  |
| `maxPrimerMismatches` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `templateCircular` |  |
| `tool` |  |

Operations: Create.

API path: `/verify_construct`

#### VirtualGel

| Field | Description |
| --- | --- |
| `circular` |  |
| `enzymes` |  |
| `gate` |  |
| `ladder` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/virtual_gel`

#### VolcanoPlotData

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `rows` |  |
| `tool` |  |

Operations: Create.

API path: `/volcano_plot_data`

#### WebSearch

| Field | Description |
| --- | --- |
| `gate` |  |
| `max_results` |  |
| `ok` |  |
| `provenance` |  |
| `query` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/web_search`



## Entities


### AlphafoldLookup

Create an instance: `alphafold_lookup = client.AlphafoldLookup()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `str` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
alphafold_lookup = client.AlphafoldLookup().create({
    "accession": "example_accession",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### AsoDesign

Create an instance: `aso_design = client.AsoDesign()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `length` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `target` | `str` |  |
| `tool` | `str` |  |
| `wing` | `int` |  |

#### Example: Create

```python
aso_design = client.AsoDesign().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "target": "example_target",  # str
    "tool": "example_tool",  # str
})
```


### BaseEditingDesign

Create an instance: `base_editing_design = client.BaseEditingDesign()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `editor` | `str` |  |
| `frameStart` | `int` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `target` | `str` |  |
| `targetPosition` | `int` |  |
| `tool` | `str` |  |

#### Example: Create

```python
base_editing_design = client.BaseEditingDesign().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "target": "example_target",  # str
    "tool": "example_tool",  # str
})
```


### Batch

Create an instance: `batch = client.Batch()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `args` | `dict` |  |
| `capped` | `bool` |  |
| `columns` | `list` |  |
| `count` | `int` |  |
| `errors` | `int` |  |
| `input` | `str` |  |
| `limit` | `int` |  |
| `provenance` | `dict` |  |
| `rows` | `list` |  |
| `tool` | `str` |  |

#### Example: Load

```python
batch = client.Batch().load()
```

#### Example: Create

```python
batch = client.Batch().create({
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


### BatchWorkflow

Create an instance: `batch__workflow = client.BatchWorkflow()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `capped` | `bool` |  |
| `columns` | `list` |  |
| `count` | `int` |  |
| `errors` | `int` |  |
| `input` | `str` |  |
| `limit` | `int` |  |
| `provenance` | `dict` |  |
| `rows` | `list` |  |
| `steps` | `list` |  |

#### Example: Load

```python
batch__workflow = client.BatchWorkflow().load()
```

#### Example: Create

```python
batch__workflow = client.BatchWorkflow().create({
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


### CharacterizeSequence

Create an instance: `characterize_sequence = client.CharacterizeSequence()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `endPrimerLength` | `int` |  |
| `gate` | `Any` |  |
| `maxOrfs` | `int` |  |
| `minOrfAa` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
characterize_sequence = client.CharacterizeSequence().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### CloningSimulate

Create an instance: `cloning_simulate = client.CloningSimulate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `armTmTarget` | `float` |  |
| `circular` | `bool` |  |
| `enzyme` | `str` |  |
| `enzyme3` | `str` |  |
| `enzyme5` | `str` |  |
| `fragments` | `list` |  |
| `gate` | `Any` |  |
| `insert` | `str` |  |
| `method` | `str` |  |
| `names` | `list` |  |
| `ok` | `Any` |  |
| `overlapLen` | `int` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |
| `vector` | `str` |  |

#### Example: Create

```python
cloning_simulate = client.CloningSimulate().create({
    "method": "example_method",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### CodonAdaptationIndex

Create an instance: `codon_adaptation_index = client.CodonAdaptationIndex()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `frameStart` | `int` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `organism` | `str` |  |
| `provenance` | `dict` |  |
| `rareThreshold` | `float` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
codon_adaptation_index = client.CodonAdaptationIndex().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### CodonOptimize

Create an instance: `codon_optimize = client.CodonOptimize()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `organism` | `str` |  |
| `protein` | `str` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
codon_optimize = client.CodonOptimize().create({
    "ok": "example_ok",  # Any
    "protein": "example_protein",  # str
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### ConstructAutofix

Create an instance: `construct_autofix = client.ConstructAutofix()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoidEnzymes` | `list` |  |
| `crypticOrfMinAa` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `Any` |  |
| `gcHigh` | `float` |  |
| `gcLow` | `float` |  |
| `gcWindow` | `int` |  |
| `homopolymerMin` | `int` |  |
| `maxPasses` | `int` |  |
| `ok` | `Any` |  |
| `organism` | `str` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
construct_autofix = client.ConstructAutofix().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### ConstructQc

Create an instance: `construct_qc = client.ConstructQc()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoidEnzymes` | `list` |  |
| `crypticOrfMinAa` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `Any` |  |
| `gcHigh` | `float` |  |
| `gcLow` | `float` |  |
| `gcWindow` | `int` |  |
| `homopolymerMin` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
construct_qc = client.ConstructQc().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### CrisprGrnaDesign

Create an instance: `crispr_grna_design = client.CrisprGrnaDesign()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `minScore` | `float` |  |
| `nuclease` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `searchReverseStrand` | `bool` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
crispr_grna_design = client.CrisprGrnaDesign().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### CrisprHdrDonor

Create an instance: `crispr_hdr_donor = client.CrisprHdrDonor()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `armLength` | `int` |  |
| `blockPam` | `bool` |  |
| `designGenotypingPrimers` | `bool` |  |
| `editEnd` | `int` |  |
| `editStart` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `Any` |  |
| `guideEnd` | `int` |  |
| `guideStart` | `int` |  |
| `guideStrand` | `str` |  |
| `nuclease` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `replacement` | `str` |  |
| `result` | `dict` |  |
| `targetSequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
crispr_hdr_donor = client.CrisprHdrDonor().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "replacement": "example_replacement",  # str
    "result": {},  # dict
    "targetSequence": "example_targetSequence",  # str
    "tool": "example_tool",  # str
})
```


### CrisprOfftargetCheck

Create an instance: `crispr_offtarget_check = client.CrisprOfftargetCheck()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `maxMismatches` | `int` |  |
| `nuclease` | `str` |  |
| `ok` | `Any` |  |
| `protospacer` | `str` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
crispr_offtarget_check = client.CrisprOfftargetCheck().create({
    "ok": "example_ok",  # Any
    "protospacer": "example_protospacer",  # str
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### CrossDimer

Create an instance: `cross_dimer = client.CrossDimer()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequenceA` | `str` |  |
| `sequenceB` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
cross_dimer = client.CrossDimer().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequenceA": "example_sequenceA",  # str
    "sequenceB": "example_sequenceB",  # str
    "tool": "example_tool",  # str
})
```


### DnaMolarity

Create an instance: `dna_molarity = client.DnaMolarity()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `length` | `int` |  |
| `massNg` | `float` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |
| `type` | `str` |  |
| `volumeUl` | `float` |  |

#### Example: Create

```python
dna_molarity = client.DnaMolarity().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### DoubleDigest

Create an instance: `double_digest = client.DoubleDigest()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enzymeA` | `str` |  |
| `enzymeB` | `str` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
double_digest = client.DoubleDigest().create({
    "enzymeA": "example_enzymeA",  # str
    "enzymeB": "example_enzymeB",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### ExportEchoPicklist

Create an instance: `export_echo_picklist = client.ExportEchoPicklist()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `reactions` | `list` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
export_echo_picklist = client.ExportEchoPicklist().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "reactions": [],  # list
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### ExportOpentronsProtocol

Create an instance: `export_opentrons_protocol = client.ExportOpentronsProtocol()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `protocolName` | `str` |  |
| `provenance` | `dict` |  |
| `reactions` | `list` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
export_opentrons_protocol = client.ExportOpentronsProtocol().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "reactions": [],  # list
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### ExportPlateLayout

Create an instance: `export_plate_layout = client.ExportPlateLayout()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `reactions` | `list` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
export_plate_layout = client.ExportPlateLayout().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "reactions": [],  # list
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### ExpressionHeatmapCluster

Create an instance: `expression_heatmap_cluster = client.ExpressionHeatmapCluster()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `clusterCols` | `bool` |  |
| `clusterRows` | `bool` |  |
| `distanceMetric` | `str` |  |
| `gate` | `Any` |  |
| `genes` | `list` |  |
| `linkage` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `samples` | `list` |  |
| `tool` | `str` |  |
| `values` | `list` |  |
| `zScoreRows` | `bool` |  |

#### Example: Create

```python
expression_heatmap_cluster = client.ExpressionHeatmapCluster().create({
    "genes": [],  # list
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "samples": [],  # list
    "tool": "example_tool",  # str
    "values": [],  # list
})
```


### FastqQcReport

Create an instance: `fastq_qc_report = client.FastqQcReport()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `input` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `qualityOffset` | `int` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
fastq_qc_report = client.FastqQcReport().create({
    "input": "example_input",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### FastqTrim

Create an instance: `fastq_trim = client.FastqTrim()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `input` | `str` |  |
| `minLength` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `qualityOffset` | `int` |  |
| `qualityThreshold` | `int` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
fastq_trim = client.FastqTrim().create({
    "input": "example_input",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### FindOrf

Create an instance: `find_orf = client.FindOrf()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `minAaLength` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `requireStop` | `bool` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
find_orf = client.FindOrf().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### FormatSequence

Create an instance: `format_sequence = client.FormatSequence()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `caseMode` | `str` |  |
| `convert` | `str` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `reverse` | `bool` |  |
| `sequence` | `str` |  |
| `stripNonLetters` | `bool` |  |
| `tool` | `str` |  |
| `width` | `int` |  |

#### Example: Create

```python
format_sequence = client.FormatSequence().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### FunctionalEnrichment

Create an instance: `functional_enrichment = client.FunctionalEnrichment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `background` | `list` |  |
| `collections` | `list` |  |
| `gate` | `Any` |  |
| `genes` | `list` |  |
| `maxTermSize` | `int` |  |
| `minTermSize` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
functional_enrichment = client.FunctionalEnrichment().create({
    "genes": [],  # list
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### GcContent

Create an instance: `gc_content = client.GcContent()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
gc_content = client.GcContent().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### GeneDossier

Create an instance: `gene_dossier = client.GeneDossier()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `gene` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
gene_dossier = client.GeneDossier().create({
    "gene": "example_gene",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### GeneExpression

Create an instance: `gene_expression = client.GeneExpression()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `gene` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
gene_expression = client.GeneExpression().create({
    "gene": "example_gene",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### GeneModel

Create an instance: `gene_model = client.GeneModel()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `gene` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
gene_model = client.GeneModel().create({
    "gene": "example_gene",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### GoldenGateFidelity

Create an instance: `golden_gate_fidelity = client.GoldenGateFidelity()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `compareToNamedSet` | `str` |  |
| `dataset` | `str` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `overhangs` | `list` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `riskThreshold` | `float` |  |
| `tool` | `str` |  |

#### Example: Create

```python
golden_gate_fidelity = client.GoldenGateFidelity().create({
    "ok": "example_ok",  # Any
    "overhangs": [],  # list
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### HgvsConvert

Create an instance: `hgvs_convert = client.HgvsConvert()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |
| `variant` | `str` |  |

#### Example: Create

```python
hgvs_convert = client.HgvsConvert().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
    "variant": "example_variant",  # str
})
```


### IdMapPoll

Create an instance: `id_map_poll = client.IdMapPoll()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `jobId` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
id_map_poll = client.IdMapPoll().create({
    "jobId": "example_jobId",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### IdMapSubmit

Create an instance: `id_map_submit = client.IdMapSubmit()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from` | `str` |  |
| `gate` | `Any` |  |
| `ids` | `list` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `taxId` | `str` |  |
| `to` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
id_map_submit = client.IdMapSubmit().create({
    "from": "example_from",  # str
    "ids": [],  # list
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "to": "example_to",  # str
    "tool": "example_tool",  # str
})
```


### InSilicoPcr

Create an instance: `in_silico_pcr = client.InSilicoPcr()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `forwardPrimer` | `str` |  |
| `gate` | `Any` |  |
| `maxMismatches` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `reversePrimer` | `str` |  |
| `template` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
in_silico_pcr = client.InSilicoPcr().create({
    "forwardPrimer": "example_forwardPrimer",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "reversePrimer": "example_reversePrimer",  # str
    "template": "example_template",  # str
    "tool": "example_tool",  # str
})
```


### KaspPrimerDesign

Create an instance: `kasp_primer_design = client.KaspPrimerDesign()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addSecondaryMismatch` | `bool` |  |
| `alleleA` | `str` |  |
| `alleleB` | `str` |  |
| `gate` | `Any` |  |
| `maxAmplicon` | `int` |  |
| `minAmplicon` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `snpPosition` | `int` |  |
| `target` | `str` |  |
| `targetCoreTm` | `float` |  |
| `tool` | `str` |  |

#### Example: Create

```python
kasp_primer_design = client.KaspPrimerDesign().create({
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


### ListTool

Create an instance: `list_tool = client.ListTool()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```python
list_tool = client.ListTool().load()
```


### MeltingTemperature

Create an instance: `melting_temperature = client.MeltingTemperature()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dntpMM` | `float` |  |
| `gate` | `Any` |  |
| `mgMM` | `float` |  |
| `naMM` | `float` |  |
| `ok` | `Any` |  |
| `oligoNM` | `float` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `targetTm` | `float` |  |
| `tmTolerance` | `float` |  |
| `tool` | `str` |  |

#### Example: Create

```python
melting_temperature = client.MeltingTemperature().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### MotifFinder

Create an instance: `motif_finder = client.MotifFinder()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `maxMismatches` | `int` |  |
| `motif` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `searchReverseStrand` | `bool` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
motif_finder = client.MotifFinder().create({
    "motif": "example_motif",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### MultipleSequenceAlignment

Create an instance: `multiple_sequence_alignment = client.MultipleSequenceAlignment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `input` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
multiple_sequence_alignment = client.MultipleSequenceAlignment().create({
    "input": "example_input",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### OligoAnalysi

Create an instance: `oligo_analysi = client.OligoAnalysi()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dntpMM` | `float` |  |
| `gate` | `Any` |  |
| `mgMM` | `float` |  |
| `naMM` | `float` |  |
| `ok` | `Any` |  |
| `oligoNM` | `float` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
oligo_analysi = client.OligoAnalysi().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### OrthologMap

Create an instance: `ortholog_map = client.OrthologMap()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sourceSpecies` | `str` |  |
| `symbols` | `list` |  |
| `targetSpecies` | `str` |  |
| `tool` | `str` |  |
| `type` | `str` |  |

#### Example: Create

```python
ortholog_map = client.OrthologMap().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "symbols": [],  # list
    "targetSpecies": "example_targetSpecies",  # str
    "tool": "example_tool",  # str
})
```


### PairwiseAlignment

Create an instance: `pairwise_alignment = client.PairwiseAlignment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gap` | `float` |  |
| `gate` | `Any` |  |
| `match` | `float` |  |
| `mismatch` | `float` |  |
| `mode` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `seqA` | `str` |  |
| `seqB` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
pairwise_alignment = client.PairwiseAlignment().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "seqA": "example_seqA",  # str
    "seqB": "example_seqB",  # str
    "tool": "example_tool",  # str
})
```


### ParseGenbank

Create an instance: `parse_genbank = client.ParseGenbank()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `text` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
parse_genbank = client.ParseGenbank().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "text": "example_text",  # str
    "tool": "example_tool",  # str
})
```


### ParseSangerTrace

Create an instance: `parse_sanger_trace = client.ParseSangerTrace()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `fileBase64` | `str` |  |
| `fileName` | `str` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
parse_sanger_trace = client.ParseSangerTrace().create({
    "fileBase64": "example_fileBase64",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### PlasmidAnnotate

Create an instance: `plasmid_annotate = client.PlasmidAnnotate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
plasmid_annotate = client.PlasmidAnnotate().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### PlasmidDeepAnnotate

Create an instance: `plasmid_deep_annotate = client.PlasmidDeepAnnotate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
plasmid_deep_annotate = client.PlasmidDeepAnnotate().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### PlasmidFullReport

Create an instance: `plasmid_full_report = client.PlasmidFullReport()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |
| `topN` | `int` |  |

#### Example: Create

```python
plasmid_full_report = client.PlasmidFullReport().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### PlasmidIdentify

Create an instance: `plasmid_identify = client.PlasmidIdentify()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |
| `topN` | `int` |  |

#### Example: Create

```python
plasmid_identify = client.PlasmidIdentify().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### PrimeEditingDesign

Create an instance: `prime_editing_design = client.PrimeEditingDesign()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `editEnd` | `int` |  |
| `editStart` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `Any` |  |
| `insertedSeq` | `str` |  |
| `ok` | `Any` |  |
| `pbsLength` | `int` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `rttHomology` | `int` |  |
| `target` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
prime_editing_design = client.PrimeEditingDesign().create({
    "editEnd": 1,  # int
    "editStart": 1,  # int
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "target": "example_target",  # str
    "tool": "example_tool",  # str
})
```


### PrimeEditingTwinDesign

Create an instance: `prime_editing_twin_design = client.PrimeEditingTwinDesign()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `newSequence` | `str` |  |
| `ok` | `Any` |  |
| `overlapLength` | `int` |  |
| `pbsLength` | `int` |  |
| `provenance` | `dict` |  |
| `replaceEnd` | `int` |  |
| `replaceStart` | `int` |  |
| `result` | `dict` |  |
| `target` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
prime_editing_twin_design = client.PrimeEditingTwinDesign().create({
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


### PrimerDesign

Create an instance: `primer_design = client.PrimerDesign()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ampliconMax` | `int` |  |
| `ampliconMin` | `int` |  |
| `dntpMM` | `float` |  |
| `gate` | `Any` |  |
| `gcMax` | `float` |  |
| `gcMin` | `float` |  |
| `lenMax` | `int` |  |
| `lenMin` | `int` |  |
| `lenOpt` | `int` |  |
| `maxReturn` | `int` |  |
| `mgMM` | `float` |  |
| `naMM` | `float` |  |
| `ok` | `Any` |  |
| `oligoNM` | `float` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `targetEnd` | `int` |  |
| `targetStart` | `int` |  |
| `template` | `str` |  |
| `tmMax` | `float` |  |
| `tmMaxDiff` | `float` |  |
| `tmMin` | `float` |  |
| `tmOpt` | `float` |  |
| `tool` | `str` |  |

#### Example: Create

```python
primer_design = client.PrimerDesign().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "template": "example_template",  # str
    "tool": "example_tool",  # str
})
```


### PrimerSpecificity

Create an instance: `primer_specificity = client.PrimerSpecificity()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `forwardPrimer` | `str` |  |
| `gate` | `Any` |  |
| `maxMismatches` | `int` |  |
| `maxProductLength` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `reversePrimer` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
primer_specificity = client.PrimerSpecificity().create({
    "forwardPrimer": "example_forwardPrimer",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "reversePrimer": "example_reversePrimer",  # str
    "tool": "example_tool",  # str
})
```


### ProteaseDigestion

Create an instance: `protease_digestion = client.ProteaseDigestion()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `maxMass` | `float` |  |
| `maxPeptides` | `int` |  |
| `minMass` | `float` |  |
| `missedCleavages` | `int` |  |
| `ok` | `Any` |  |
| `protease` | `str` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
protease_digestion = client.ProteaseDigestion().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### ProteinAnnotatePoll

Create an instance: `protein_annotate_poll = client.ProteinAnnotatePoll()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `jobId` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
protein_annotate_poll = client.ProteinAnnotatePoll().create({
    "jobId": "example_jobId",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### ProteinAnnotateSubmit

Create an instance: `protein_annotate_submit = client.ProteinAnnotateSubmit()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `appl` | `str` |  |
| `gate` | `Any` |  |
| `goterms` | `bool` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
protein_annotate_submit = client.ProteinAnnotateSubmit().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### ProteinHydrophobicity

Create an instance: `protein_hydrophobicity = client.ProteinHydrophobicity()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `scale` | `str` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |
| `window` | `int` |  |

#### Example: Create

```python
protein_hydrophobicity = client.ProteinHydrophobicity().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### ProteinProperty

Create an instance: `protein_property = client.ProteinProperty()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `chargeStep` | `float` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
protein_property = client.ProteinProperty().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### RandomSequence

Create an instance: `random_sequence = client.RandomSequence()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `gcContent` | `float` |  |
| `kind` | `str` |  |
| `length` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
random_sequence = client.RandomSequence().create({
    "length": 1,  # int
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### RestrictionSite

Create an instance: `restriction_site = client.RestrictionSite()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enzymes` | `list` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
restriction_site = client.RestrictionSite().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### ReverseComplement

Create an instance: `reverse_complement = client.ReverseComplement()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |
| `type` | `str` |  |

#### Example: Create

```python
reverse_complement = client.ReverseComplement().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### ReverseTranslate

Create an instance: `reverse_translate = client.ReverseTranslate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `mode` | `str` |  |
| `ok` | `Any` |  |
| `organism` | `str` |  |
| `protein` | `str` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
reverse_translate = client.ReverseTranslate().create({
    "ok": "example_ok",  # Any
    "protein": "example_protein",  # str
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### RnaFold

Create an instance: `rna_fold = client.RnaFold()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
rna_fold = client.RnaFold().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### SangerVsReference

Create an instance: `sanger_vs_reference = client.SangerVsReference()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `fileBase64` | `str` |  |
| `fileName` | `str` |  |
| `gate` | `Any` |  |
| `minCoverage` | `float` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `read` | `str` |  |
| `reference` | `str` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
sanger_vs_reference = client.SangerVsReference().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "reference": "example_reference",  # str
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### SavePermalink

Create an instance: `save_permalink = client.SavePermalink()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `args` | `dict` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
save_permalink = client.SavePermalink().create({
    "args": {},  # dict
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### SeqfileStat

Create an instance: `seqfile_stat = client.SeqfileStat()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `input` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `qualityOffset` | `int` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
seqfile_stat = client.SeqfileStat().create({
    "input": "example_input",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### SequenceFetch

Create an instance: `sequence_fetch = client.SequenceFetch()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `str` |  |
| `db` | `str` |  |
| `format` | `str` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
sequence_fetch = client.SequenceFetch().create({
    "accession": "example_accession",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### SequenceFormatConvert

Create an instance: `sequence_format_convert = client.SequenceFormatConvert()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from` | `str` |  |
| `gate` | `Any` |  |
| `input` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `to` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
sequence_format_convert = client.SequenceFormatConvert().create({
    "input": "example_input",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### SequenceReport

Create an instance: `sequence_report = client.SequenceReport()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `endPrimerLength` | `int` |  |
| `gate` | `Any` |  |
| `maxOrfs` | `int` |  |
| `minOrfAa` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
sequence_report = client.SequenceReport().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### SequenceSearch

Create an instance: `sequence_search = client.SequenceSearch()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `db` | `str` |  |
| `gate` | `Any` |  |
| `gene` | `str` |  |
| `maxResults` | `int` |  |
| `ok` | `Any` |  |
| `organism` | `str` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `term` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
sequence_search = client.SequenceSearch().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### SequencingReadbackVerify

Create an instance: `sequencing_readback_verify = client.SequencingReadbackVerify()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `minSupportingReads` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `reads` | `str` |  |
| `reference` | `str` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
sequencing_readback_verify = client.SequencingReadbackVerify().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "reads": "example_reads",  # str
    "reference": "example_reference",  # str
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### SessionCreate

Create an instance: `session_create = client.SessionCreate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entries` | `dict` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
session_create = client.SessionCreate().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### SessionGet

Create an instance: `session_get = client.SessionGet()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `names` | `list` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sessionId` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
session_get = client.SessionGet().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sessionId": "example_sessionId",  # str
    "tool": "example_tool",  # str
})
```


### SessionRun

Create an instance: `session_run = client.SessionRun()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `args` | `dict` |  |
| `fromSession` | `dict` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sessionId` | `str` |  |
| `tool` | `str` |  |
| `writeBack` | `dict` |  |

#### Example: Create

```python
session_run = client.SessionRun().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sessionId": "example_sessionId",  # str
    "tool": "example_tool",  # str
})
```


### SessionSet

Create an instance: `session_set = client.SessionSet()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entries` | `dict` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sessionId` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
session_set = client.SessionSet().create({
    "entries": {},  # dict
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sessionId": "example_sessionId",  # str
    "tool": "example_tool",  # str
})
```


### SirnaDesign

Create an instance: `sirna_design = client.SirnaDesign()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `minReynolds` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `shRnaLoop` | `str` |  |
| `target` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
sirna_design = client.SirnaDesign().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "target": "example_target",  # str
    "tool": "example_tool",  # str
})
```


### SiteDirectedMutagenesi

Create an instance: `site_directed_mutagenesi = client.SiteDirectedMutagenesi()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `armTmTarget` | `float` |  |
| `dntpMM` | `float` |  |
| `editKind` | `str` |  |
| `frameStart` | `int` |  |
| `gate` | `Any` |  |
| `mgMM` | `float` |  |
| `naMM` | `float` |  |
| `newBase` | `str` |  |
| `ok` | `Any` |  |
| `oligoNM` | `float` |  |
| `organism` | `str` |  |
| `position` | `int` |  |
| `provenance` | `dict` |  |
| `residue` | `int` |  |
| `result` | `dict` |  |
| `style` | `str` |  |
| `targetAa` | `str` |  |
| `template` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
site_directed_mutagenesi = client.SiteDirectedMutagenesi().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "template": "example_template",  # str
    "tool": "example_tool",  # str
})
```


### Translate

Create an instance: `translate = client.Translate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `frame` | `int` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `toStop` | `bool` |  |
| `tool` | `str` |  |

#### Example: Create

```python
translate = client.Translate().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### VariantAnnotate

Create an instance: `variant_annotate = client.VariantAnnotate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `assembly` | `str` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |
| `variant` | `str` |  |

#### Example: Create

```python
variant_annotate = client.VariantAnnotate().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
    "variant": "example_variant",  # str
})
```


### VariantComparator

Create an instance: `variant_comparator = client.VariantComparator()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `coding` | `bool` |  |
| `frameStart` | `int` |  |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `query` | `str` |  |
| `reference` | `str` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
variant_comparator = client.VariantComparator().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "query": "example_query",  # str
    "reference": "example_reference",  # str
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### VerifyAssembly

Create an instance: `verify_assembly = client.VerifyAssembly()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `armTmTarget` | `float` |  |
| `circular` | `bool` |  |
| `claimedConstruct` | `str` |  |
| `coding` | `bool` |  |
| `enzyme` | `str` |  |
| `enzyme3` | `str` |  |
| `enzyme5` | `str` |  |
| `fragmentPcrs` | `list` |  |
| `fragments` | `list` |  |
| `frameStart` | `int` |  |
| `gate` | `Any` |  |
| `insert` | `str` |  |
| `insertPcr` | `dict` |  |
| `method` | `str` |  |
| `names` | `list` |  |
| `ok` | `Any` |  |
| `overlapLen` | `int` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |
| `vector` | `str` |  |
| `vectorPcr` | `dict` |  |

#### Example: Create

```python
verify_assembly = client.VerifyAssembly().create({
    "claimedConstruct": "example_claimedConstruct",  # str
    "method": "example_method",  # str
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


### VerifyConstruct

Create an instance: `verify_construct = client.VerifyConstruct()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `claimedConstruct` | `str` |  |
| `expectedFrameStart` | `int` |  |
| `gate` | `Any` |  |
| `insertForwardPrimer` | `str` |  |
| `insertReversePrimer` | `str` |  |
| `insertTemplate` | `str` |  |
| `maxPrimerMismatches` | `int` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `templateCircular` | `bool` |  |
| `tool` | `str` |  |

#### Example: Create

```python
verify_construct = client.VerifyConstruct().create({
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


### VirtualGel

Create an instance: `virtual_gel = client.VirtualGel()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `enzymes` | `list` |  |
| `gate` | `Any` |  |
| `ladder` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `sequence` | `str` |  |
| `tool` | `str` |  |

#### Example: Create

```python
virtual_gel = client.VirtualGel().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "sequence": "example_sequence",  # str
    "tool": "example_tool",  # str
})
```


### VolcanoPlotData

Create an instance: `volcano_plot_data = client.VolcanoPlotData()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` |  |
| `rows` | `list` |  |
| `tool` | `str` |  |

#### Example: Create

```python
volcano_plot_data = client.VolcanoPlotData().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "result": {},  # dict
    "rows": [],  # list
    "tool": "example_tool",  # str
})
```


### WebSearch

Create an instance: `web_search = client.WebSearch()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Any` |  |
| `max_results` | `float` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `query` | `str` |  |
| `result` | `dict` |  |
| `tool` | `str` |  |

#### Example: Create

```python
web_search = client.WebSearch().create({
    "ok": "example_ok",  # Any
    "provenance": {},  # dict
    "query": "example_query",  # str
    "result": {},  # dict
    "tool": "example_tool",  # str
})
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── seqbenchmcp_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`seqbenchmcp_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```python
batch = client.Batch()
batch.load()

# batch.data_get() now returns the batch data from the last load
# batch.match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
