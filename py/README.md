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
| `accession` | UniProt accession, e.g. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/alphafold_lookup`

#### AsoDesign

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | Total gapmer length (nt). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `target` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |
| `wing` | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

Operations: Create.

API path: `/aso_design`

#### BaseEditingDesign

| Field | Description |
| --- | --- |
| `editor` | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `target` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/base_editing_design`

#### Batch

| Field | Description |
| --- | --- |
| `args` | Shared tool arguments applied to every record. |
| `capped` | True if input exceeded the record limit. |
| `columns` |  |
| `count` |  |
| `errors` |  |
| `input` | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | Maximum records per call (500). |
| `provenance` |  |
| `rows` |  |
| `tool` | A batchable tool slug (see `GET /batch`). |

Operations: Create, Load.

API path: `/batch`

#### BatchWorkflow

| Field | Description |
| --- | --- |
| `capped` |  |
| `columns` | Flattened "<step>·<tool>·<key>" column headers. |
| `count` |  |
| `errors` |  |
| `input` | Multi-FASTA text or one sequence per line. |
| `limit` | Maximum records per call (200). |
| `provenance` |  |
| `rows` |  |
| `steps` |  |

Operations: Create, Load.

API path: `/workflow`

#### CharacterizeSequence

| Field | Description |
| --- | --- |
| `endPrimerLength` | Length of the naive end primers taken from each end. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/characterize_sequence`

#### CloningSimulate

| Field | Description |
| --- | --- |
| `armTmTarget` | Target annealing Tm (°C) for primer arms. |
| `circular` | Produce a circular product. |
| `enzyme` | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | 3′ enzyme (restriction method). |
| `enzyme5` | 5′ enzyme (restriction method). |
| `fragments` | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | Insert sequence (restriction method). |
| `method` | Assembly method. |
| `names` | Optional labels for each fragment. |
| `ok` |  |
| `overlapLen` | Gibson homology-arm length (bp). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |
| `vector` | Vector sequence (restriction method). |

Operations: Create.

API path: `/cloning_simulate`

#### CodonAdaptationIndex

| Field | Description |
| --- | --- |
| `frameStart` | 1-based position to start reading codons. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `organism` |  |
| `provenance` |  |
| `rareThreshold` | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | Tool-specific output object. |
| `sequence` | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/codon_adaptation_index`

#### CodonOptimize

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `organism` |  |
| `protein` | Protein sequence (one-letter codes). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/codon_optimize`

#### ConstructAutofix

| Field | Description |
| --- | --- |
| `avoidEnzymes` | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | 1-based nucleotide where the reading frame begins. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` |  |
| `gcLow` |  |
| `gcWindow` |  |
| `homopolymerMin` |  |
| `maxPasses` | Repeat full passes until clean or no further progress. |
| `ok` |  |
| `organism` | Codon-usage table to prefer among synonymous options. |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/construct_autofix`

#### ConstructQc

| Field | Description |
| --- | --- |
| `avoidEnzymes` | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | 1-based nucleotide where the reading frame begins. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | GC% above this flags a GC-rich window. |
| `gcLow` | GC% below this flags an AT-rich window. |
| `gcWindow` | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | Minimum run length to flag a homopolymer. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/construct_qc`

#### CrisprGrnaDesign

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | Nuclease id. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `searchReverseStrand` | Also scan the reverse strand for guides. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/crispr_grna_design`

#### CrisprHdrDonor

| Field | Description |
| --- | --- |
| `armLength` | Homology arm length (bp) on each side. |
| `blockPam` | When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut. |
| `designGenotypingPrimers` | Also design a primer pair (on the original targetSequence) whose product spans the edit site. |
| `editEnd` | 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. |
| `editStart` | 1-based start of the region being replaced. |
| `frameStart` | Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | Strand the guide's protospacer is on. |
| `nuclease` | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` |  |
| `provenance` |  |
| `replacement` | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | Tool-specific output object. |
| `targetSequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/crispr_hdr_donor`

#### CrisprOfftargetCheck

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` |  |
| `protospacer` | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/crispr_offtarget_check`

#### CrossDimer

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequenceA` | First oligo (5'→3'). |
| `sequenceB` | Second oligo (5'→3'). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/cross_dimer`

#### DnaMolarity

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | Mass in nanograms. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | The tool slug that ran. |
| `type` | Molecule type. |
| `volumeUl` | Volume in microlitres (0 = unknown; needed for concentration). |

Operations: Create.

API path: `/dna_molarity`

#### DoubleDigest

| Field | Description |
| --- | --- |
| `enzymeA` | First enzyme name (e.g. |
| `enzymeB` | Second enzyme name (e.g. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/double_digest`

#### ExportEchoPicklist

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `reactions` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/export_echo_picklist`

#### ExportOpentronsProtocol

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `protocolName` | Optional protocol name (used in the script's metadata). |
| `provenance` |  |
| `reactions` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/export_opentrons_protocol`

#### ExportPlateLayout

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `reactions` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/export_plate_layout`

#### ExpressionHeatmapCluster

| Field | Description |
| --- | --- |
| `clusterCols` | Cluster (reorder) samples. |
| `clusterRows` | Cluster (reorder) genes. |
| `distanceMetric` | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | Row (gene) labels. |
| `linkage` | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `samples` | Column (sample) labels. |
| `tool` | The tool slug that ran. |
| `values` | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

Operations: Create.

API path: `/expression_heatmap_cluster`

#### FastqQcReport

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` |  |
| `provenance` |  |
| `qualityOffset` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/fastq_qc_report`

#### FastqTrim

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | Reads shorter than this after trimming are dropped. |
| `ok` |  |
| `provenance` |  |
| `qualityOffset` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | 3' quality-trim threshold (Phred score). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/fastq_trim`

#### FindOrf

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | Minimum protein length (aa) to report. |
| `ok` |  |
| `provenance` |  |
| `requireStop` | Only report ORFs terminated by a stop codon. |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/find_orfs`

#### FormatSequence

| Field | Description |
| --- | --- |
| `caseMode` |  |
| `convert` | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `reverse` | Reverse the sequence (no complement). |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | Remove digits, spaces and gaps (keep letters only). |
| `tool` | The tool slug that ran. |
| `width` | Line-wrap width; 0 = single line. |

Operations: Create.

API path: `/format_sequence`

#### FunctionalEnrichment

| Field | Description |
| --- | --- |
| `background` | Custom background/universe gene symbols. |
| `collections` | Which term collections to test. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | Query gene symbols (human, e.g. |
| `maxTermSize` | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | Skip terms/pathways with fewer than this many background genes. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/functional_enrichment`

#### GcContent

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/gc_content`

#### GeneDossier

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/gene_dossier`

#### GeneExpression

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/gene_expression`

#### GeneModel

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/gene_model`

#### GoldenGateFidelity

| Field | Description |
| --- | --- |
| `compareToNamedSet` | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `overhangs` | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `riskThreshold` | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/golden_gate_fidelity`

#### HgvsConvert

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |
| `variant` | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

Operations: Create.

API path: `/hgvs_convert`

#### IdMapPoll

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` |  |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/id_map_poll`

#### IdMapSubmit

| Field | Description |
| --- | --- |
| `from` | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | The ids to map, up to 1000 (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `taxId` | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | Target id type. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/id_map_submit`

#### InSilicoPcr

| Field | Description |
| --- | --- |
| `circular` | Treat the template as circular (plasmid). |
| `forwardPrimer` | Primer 1, 5'→3'. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | Mismatches tolerated per primer. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `reversePrimer` | Primer 2, 5'→3' (order does not matter). |
| `template` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/in_silico_pcr`

#### KaspPrimerDesign

| Field | Description |
| --- | --- |
| `addSecondaryMismatch` | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | First allele (single base) — gets the FAM tail. |
| `alleleB` | Second allele (single base) — gets the HEX tail. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | Minimum amplicon length for the common reverse primer. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `snpPosition` | 1-based position of the SNP on the forward strand. |
| `target` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | The tool slug that ran. |

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
| `dntpMM` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | Divalent cation [Mg2+] (mM). |
| `naMM` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` |  |
| `oligoNM` | Total strand concentration (nM). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | Optional target Tm (°C). |
| `tmTolerance` | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/melting_temperature`

#### MotifFinder

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | Maximum allowed mismatches per match. |
| `motif` | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `searchReverseStrand` | Also search the reverse strand. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/motif_finder`

#### MultipleSequenceAlignment

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/multiple_sequence_alignment`

#### OligoAnalysi

| Field | Description |
| --- | --- |
| `dntpMM` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | Divalent cation [Mg2+] (mM). |
| `naMM` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` |  |
| `oligoNM` | Total strand concentration (nM). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/oligo_analysis`

#### OrthologMap

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sourceSpecies` | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | Ensembl species slug to find homologs in (e.g. |
| `tool` | The tool slug that ran. |
| `type` | Homology type to return. |

Operations: Create.

API path: `/ortholog_map`

#### PairwiseAlignment

| Field | Description |
| --- | --- |
| `gap` | Linear gap penalty (per gap position). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | Match score. |
| `mismatch` | Mismatch penalty. |
| `mode` |  |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `seqA` | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/pairwise_alignment`

#### ParseGenbank

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `text` | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/parse_genbank`

#### ParseSangerTrace

| Field | Description |
| --- | --- |
| `fileBase64` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | Optional original file name (echoed back). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/parse_sanger_trace`

#### PlasmidAnnotate

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/plasmid_annotate`

#### PlasmidDeepAnnotate

| Field | Description |
| --- | --- |
| `circular` | Treat the sequence as a circular plasmid (vs. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/plasmid_deep_annotate`

#### PlasmidFullReport

| Field | Description |
| --- | --- |
| `circular` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |
| `topN` | How many top-ranked backbone candidates to report. |

Operations: Create.

API path: `/plasmid_full_report`

#### PlasmidIdentify

| Field | Description |
| --- | --- |
| `circular` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |
| `topN` | How many top-ranked backbone candidates to report. |

Operations: Create.

API path: `/plasmid_identify`

#### PrimeEditingDesign

| Field | Description |
| --- | --- |
| `editEnd` | 1-based inclusive end of the region being changed. |
| `editStart` | 1-based inclusive start of the region being changed. |
| `frameStart` | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | Replacement bases (forward strand). |
| `ok` |  |
| `pbsLength` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `rttHomology` | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/prime_editing_design`

#### PrimeEditingTwinDesign

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` |  |
| `overlapLength` | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` |  |
| `replaceEnd` | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | 1-based inclusive start of the region being replaced/deleted. |
| `result` | Tool-specific output object. |
| `target` | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/prime_editing_twin_design`

#### PrimerDesign

| Field | Description |
| --- | --- |
| `ampliconMax` |  |
| `ampliconMin` |  |
| `dntpMM` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` |  |
| `gcMin` |  |
| `lenMax` |  |
| `lenMin` |  |
| `lenOpt` |  |
| `maxReturn` | Number of best pairs to return. |
| `mgMM` | Divalent cation [Mg2+] (mM). |
| `naMM` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` |  |
| `oligoNM` | Total strand concentration (nM). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `targetEnd` | 1-based inclusive end of the target region (optional). |
| `targetStart` | 1-based inclusive start of a region the product must span (optional). |
| `template` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` |  |
| `tmMaxDiff` | Max Tm difference within a pair (°C). |
| `tmMin` |  |
| `tmOpt` |  |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/primer_design`

#### PrimerSpecificity

| Field | Description |
| --- | --- |
| `forwardPrimer` | Forward primer, 5'→3'. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `reversePrimer` | Reverse primer, 5'→3'. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/primer_specificity`

#### ProteaseDigestion

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | Cap on the number of returned peptides. |
| `minMass` | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | Allowed missed internal cleavages (0–2). |
| `ok` |  |
| `protease` | Protease or chemical cleavage agent. |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/protease_digestion`

#### ProteinAnnotatePoll

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` |  |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/protein_annotate_poll`

#### ProteinAnnotateSubmit

| Field | Description |
| --- | --- |
| `appl` | Restrict to one member database (e.g. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | Include GO-term cross-references. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/protein_annotate_submit`

#### ProteinHydrophobicity

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `scale` | Amino-acid scale. |
| `sequence` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | The tool slug that ran. |
| `window` | Sliding-window size (clamped to an odd number ≥ 1). |

Operations: Create.

API path: `/protein_hydrophobicity`

#### ProteinProperty

| Field | Description |
| --- | --- |
| `chargeStep` | pH step for the net-charge titration curve (0–14). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/protein_properties`

#### RandomSequence

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` |  |
| `length` | Number of residues to generate. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/random_sequence`

#### RestrictionSite

| Field | Description |
| --- | --- |
| `enzymes` | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/restriction_sites`

#### ReverseComplement

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |
| `type` |  |

Operations: Create.

API path: `/reverse_complement`

#### ReverseTranslate

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` |  |
| `ok` |  |
| `organism` | Codon-usage host (ignored in degenerate mode). |
| `protein` | Protein sequence (one-letter codes; * for stop). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/reverse_translate`

#### RnaFold

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/rna_fold`

#### SangerVsReference

| Field | Description |
| --- | --- |
| `fileBase64` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | Optional original file name (echoed back). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` |  |
| `provenance` |  |
| `read` | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | Expected reference sequence (FASTA or raw). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sanger_vs_reference`

#### SavePermalink

| Field | Description |
| --- | --- |
| `args` | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/save_permalink`

#### SeqfileStat

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` |  |
| `provenance` |  |
| `qualityOffset` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/seqfile_stats`

#### SequenceFetch

| Field | Description |
| --- | --- |
| `accession` | GenBank/RefSeq accession (e.g. |
| `db` | Database to query; auto-detects from the accession format. |
| `format` | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sequence_fetch`

#### SequenceFormatConvert

| Field | Description |
| --- | --- |
| `from` | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | A FASTA or GenBank record to convert. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `to` | Output format. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sequence_format_convert`

#### SequenceReport

| Field | Description |
| --- | --- |
| `endPrimerLength` | Length of the naive end primers taken from each end. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | Minimum ORF length in amino acids. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sequence_report`

#### SequenceSearch

| Field | Description |
| --- | --- |
| `db` |  |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | Gene symbol/name, e.g. |
| `maxResults` | Up to 20. |
| `ok` |  |
| `organism` | Organism name, e.g. |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `term` | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sequence_search`

#### SequencingReadbackVerify

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` |  |
| `provenance` |  |
| `reads` | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | The claimed/expected reference sequence. |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sequencing_readback_verify`

#### SessionCreate

| Field | Description |
| --- | --- |
| `entries` | Initial named entries, e.g. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/session_create`

#### SessionGet

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | Only return these entries; omit to return all of them. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sessionId` |  |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/session_get`

#### SessionRun

| Field | Description |
| --- | --- |
| `args` | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sessionId` |  |
| `tool` | The tool slug that ran. |
| `writeBack` | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

Operations: Create.

API path: `/session_run`

#### SessionSet

| Field | Description |
| --- | --- |
| `entries` | Named entries to add/overwrite, e.g. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sessionId` |  |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/session_set`

#### SirnaDesign

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `shRnaLoop` | Loop sequence used when assembling the shRNA cassette. |
| `target` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/sirna_design`

#### SiteDirectedMutagenesi

| Field | Description |
| --- | --- |
| `armTmTarget` | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | Edit at the nucleotide or amino-acid level. |
| `frameStart` | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | Divalent cation [Mg2+] (mM). |
| `naMM` | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | Replacement base (editKind='nt'). |
| `ok` |  |
| `oligoNM` | Total strand concentration (nM). |
| `organism` | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | 1-based position to substitute (editKind='nt'). |
| `provenance` |  |
| `residue` | 1-based residue number to change (editKind='aa'). |
| `result` | Tool-specific output object. |
| `style` | Mutagenic primer style. |
| `targetAa` | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/site_directed_mutagenesis`

#### Translate

| Field | Description |
| --- | --- |
| `frame` |  |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | Stop at the first stop codon. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/translate`

#### VariantAnnotate

| Field | Description |
| --- | --- |
| `assembly` | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |
| `variant` | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

Operations: Create.

API path: `/variant_annotate`

#### VariantComparator

| Field | Description |
| --- | --- |
| `coding` | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | 1-based reading-frame start (used when coding is true). |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `query` | Query / variant sequence (raw or FASTA). |
| `reference` | Reference / wild-type sequence (raw or FASTA). |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/variant_comparator`

#### VerifyAssembly

| Field | Description |
| --- | --- |
| `armTmTarget` | Target annealing Tm (°C) for primer arms. |
| `circular` | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | The sequence you claim you ended up with. |
| `coding` | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | Type IIS enzyme for Golden Gate. |
| `enzyme3` | 3′ enzyme (restriction method). |
| `enzyme5` | 5′ enzyme (restriction method). |
| `fragmentPcrs` | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | Insert sequence (restriction method). |
| `insertPcr` | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | Assembly method used. |
| `names` | Optional labels for each fragment. |
| `ok` |  |
| `overlapLen` | Gibson homology-arm length (bp). |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |
| `vector` | Vector sequence (restriction method). |
| `vectorPcr` | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

Operations: Create.

API path: `/verify_assembly`

#### VerifyConstruct

| Field | Description |
| --- | --- |
| `claimedConstruct` | The final sequence claimed to have been built. |
| `expectedFrameStart` | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | Mismatches tolerated per primer during PCR prediction. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `templateCircular` | Treat insertTemplate as circular (e.g. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/verify_construct`

#### VirtualGel

| Field | Description |
| --- | --- |
| `circular` | Treat the sequence as circular (plasmid). |
| `enzymes` | Enzyme names to digest with. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | DNA ladder to plot alongside the sample lane. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `sequence` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/virtual_gel`

#### VolcanoPlotData

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `rows` | Differential expression rows, one per gene. |
| `tool` | The tool slug that ran. |

Operations: Create.

API path: `/volcano_plot_data`

#### WebSearch

| Field | Description |
| --- | --- |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | Maximum number of results to return (default 5, max 10). |
| `ok` |  |
| `provenance` |  |
| `query` | The search query. |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

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
| `accession` | `str` | UniProt accession, e.g. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `int` | Total gapmer length (nt). |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `target` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |
| `wing` | `int` | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

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
| `editor` | `str` | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | `int` | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `target` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | `int` | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | `str` | The tool slug that ran. |

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
| `args` | `dict` | Shared tool arguments applied to every record. |
| `capped` | `bool` | True if input exceeded the record limit. |
| `columns` | `list` |  |
| `count` | `int` |  |
| `errors` | `int` |  |
| `input` | `str` | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | `int` | Maximum records per call (500). |
| `provenance` | `dict` |  |
| `rows` | `list` |  |
| `tool` | `str` | A batchable tool slug (see `GET /batch`). |

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
| `columns` | `list` | Flattened "<step>·<tool>·<key>" column headers. |
| `count` | `int` |  |
| `errors` | `int` |  |
| `input` | `str` | Multi-FASTA text or one sequence per line. |
| `limit` | `int` | Maximum records per call (200). |
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
| `endPrimerLength` | `int` | Length of the naive end primers taken from each end. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `int` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `int` | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `armTmTarget` | `float` | Target annealing Tm (°C) for primer arms. |
| `circular` | `bool` | Produce a circular product. |
| `enzyme` | `str` | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | `str` | 3′ enzyme (restriction method). |
| `enzyme5` | `str` | 5′ enzyme (restriction method). |
| `fragments` | `list` | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `str` | Insert sequence (restriction method). |
| `method` | `str` | Assembly method. |
| `names` | `list` | Optional labels for each fragment. |
| `ok` | `Any` |  |
| `overlapLen` | `int` | Gibson homology-arm length (bp). |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |
| `vector` | `str` | Vector sequence (restriction method). |

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
| `frameStart` | `int` | 1-based position to start reading codons. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `organism` | `str` |  |
| `provenance` | `dict` |  |
| `rareThreshold` | `float` | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `organism` | `str` |  |
| `protein` | `str` | Protein sequence (one-letter codes). |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `avoidEnzymes` | `list` | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | `int` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `int` | 1-based nucleotide where the reading frame begins. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `float` |  |
| `gcLow` | `float` |  |
| `gcWindow` | `int` |  |
| `homopolymerMin` | `int` |  |
| `maxPasses` | `int` | Repeat full passes until clean or no further progress. |
| `ok` | `Any` |  |
| `organism` | `str` | Codon-usage table to prefer among synonymous options. |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `avoidEnzymes` | `list` | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | `int` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `int` | 1-based nucleotide where the reading frame begins. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `float` | GC% above this flags a GC-rich window. |
| `gcLow` | `float` | GC% below this flags an AT-rich window. |
| `gcWindow` | `int` | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | `int` | Minimum run length to flag a homopolymer. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | `float` | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | `str` | Nuclease id. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `searchReverseStrand` | `bool` | Also scan the reverse strand for guides. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `armLength` | `int` | Homology arm length (bp) on each side. |
| `blockPam` | `bool` | When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut. |
| `designGenotypingPrimers` | `bool` | Also design a primer pair (on the original targetSequence) whose product spans the edit site. |
| `editEnd` | `int` | 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. |
| `editStart` | `int` | 1-based start of the region being replaced. |
| `frameStart` | `int` | Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | `int` | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | `int` | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | `str` | Strand the guide's protospacer is on. |
| `nuclease` | `str` | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `replacement` | `str` | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | `dict` | Tool-specific output object. |
| `targetSequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | `str` | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` | `Any` |  |
| `protospacer` | `str` | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequenceA` | `str` | First oligo (5'→3'). |
| `sequenceB` | `str` | Second oligo (5'→3'). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `int` | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | `float` | Mass in nanograms. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | `str` | The tool slug that ran. |
| `type` | `str` | Molecule type. |
| `volumeUl` | `float` | Volume in microlitres (0 = unknown; needed for concentration). |

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
| `enzymeA` | `str` | First enzyme name (e.g. |
| `enzymeB` | `str` | Second enzyme name (e.g. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `reactions` | `list` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `protocolName` | `str` | Optional protocol name (used in the script's metadata). |
| `provenance` | `dict` |  |
| `reactions` | `list` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `reactions` | `list` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `clusterCols` | `bool` | Cluster (reorder) samples. |
| `clusterRows` | `bool` | Cluster (reorder) genes. |
| `distanceMetric` | `str` | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `list` | Row (gene) labels. |
| `linkage` | `str` | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `samples` | `list` | Column (sample) labels. |
| `tool` | `str` | The tool slug that ran. |
| `values` | `list` | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | `bool` | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `str` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `qualityOffset` | `int` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `str` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | `int` | Reads shorter than this after trimming are dropped. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `qualityOffset` | `int` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | `int` | 3' quality-trim threshold (Phred score). |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | `int` | Minimum protein length (aa) to report. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `requireStop` | `bool` | Only report ORFs terminated by a stop codon. |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `convert` | `str` | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `reverse` | `bool` | Reverse the sequence (no complement). |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | `bool` | Remove digits, spaces and gaps (keep letters only). |
| `tool` | `str` | The tool slug that ran. |
| `width` | `int` | Line-wrap width; 0 = single line. |

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
| `background` | `list` | Custom background/universe gene symbols. |
| `collections` | `list` | Which term collections to test. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `list` | Query gene symbols (human, e.g. |
| `maxTermSize` | `int` | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | `int` | Skip terms/pathways with fewer than this many background genes. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `str` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `str` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `str` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `compareToNamedSet` | `str` | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | `str` | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `overhangs` | `list` | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `riskThreshold` | `float` | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |
| `variant` | `str` | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `from` | `str` | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | `list` | The ids to map, up to 1000 (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `taxId` | `str` | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | `str` | Target id type. |
| `tool` | `str` | The tool slug that ran. |

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
| `circular` | `bool` | Treat the template as circular (plasmid). |
| `forwardPrimer` | `str` | Primer 1, 5'→3'. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | Mismatches tolerated per primer. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `reversePrimer` | `str` | Primer 2, 5'→3' (order does not matter). |
| `template` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `addSecondaryMismatch` | `bool` | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | `str` | First allele (single base) — gets the FAM tail. |
| `alleleB` | `str` | Second allele (single base) — gets the HEX tail. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | `int` | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | `int` | Minimum amplicon length for the common reverse primer. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `snpPosition` | `int` | 1-based position of the SNP on the forward strand. |
| `target` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | `float` | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | `str` | The tool slug that ran. |

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
| `dntpMM` | `float` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Any` |  |
| `oligoNM` | `float` | Total strand concentration (nM). |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | `float` | Optional target Tm (°C). |
| `tmTolerance` | `float` | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | Maximum allowed mismatches per match. |
| `motif` | `str` | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `searchReverseStrand` | `bool` | Also search the reverse strand. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `str` | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `dntpMM` | `float` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Any` |  |
| `oligoNM` | `float` | Total strand concentration (nM). |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sourceSpecies` | `str` | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | `list` | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | `str` | Ensembl species slug to find homologs in (e.g. |
| `tool` | `str` | The tool slug that ran. |
| `type` | `str` | Homology type to return. |

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
| `gap` | `float` | Linear gap penalty (per gap position). |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | `float` | Match score. |
| `mismatch` | `float` | Mismatch penalty. |
| `mode` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `seqA` | `str` | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | `str` | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `text` | `str` | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | `str` | The tool slug that ran. |

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
| `fileBase64` | `str` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `str` | Optional original file name (echoed back). |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `circular` | `bool` | Treat the sequence as a circular plasmid (vs. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `circular` | `bool` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |
| `topN` | `int` | How many top-ranked backbone candidates to report. |

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
| `circular` | `bool` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |
| `topN` | `int` | How many top-ranked backbone candidates to report. |

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
| `editEnd` | `int` | 1-based inclusive end of the region being changed. |
| `editStart` | `int` | 1-based inclusive start of the region being changed. |
| `frameStart` | `int` | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | `str` | Replacement bases (forward strand). |
| `ok` | `Any` |  |
| `pbsLength` | `int` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `rttHomology` | `int` | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | `str` | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | `str` | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` | `Any` |  |
| `overlapLength` | `int` | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | `int` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `dict` |  |
| `replaceEnd` | `int` | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | `int` | 1-based inclusive start of the region being replaced/deleted. |
| `result` | `dict` | Tool-specific output object. |
| `target` | `str` | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | `str` | The tool slug that ran. |

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
| `dntpMM` | `float` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` | `float` |  |
| `gcMin` | `float` |  |
| `lenMax` | `int` |  |
| `lenMin` | `int` |  |
| `lenOpt` | `int` |  |
| `maxReturn` | `int` | Number of best pairs to return. |
| `mgMM` | `float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Any` |  |
| `oligoNM` | `float` | Total strand concentration (nM). |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `targetEnd` | `int` | 1-based inclusive end of the target region (optional). |
| `targetStart` | `int` | 1-based inclusive start of a region the product must span (optional). |
| `template` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` | `float` |  |
| `tmMaxDiff` | `float` | Max Tm difference within a pair (°C). |
| `tmMin` | `float` |  |
| `tmOpt` | `float` |  |
| `tool` | `str` | The tool slug that ran. |

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
| `forwardPrimer` | `str` | Forward primer, 5'→3'. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `int` | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | `int` | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `reversePrimer` | `str` | Reverse primer, 5'→3'. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | `float` | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | `int` | Cap on the number of returned peptides. |
| `minMass` | `float` | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | `int` | Allowed missed internal cleavages (0–2). |
| `ok` | `Any` |  |
| `protease` | `str` | Protease or chemical cleavage agent. |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `str` |  |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `appl` | `str` | Restrict to one member database (e.g. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | `bool` | Include GO-term cross-references. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `scale` | `str` | Amino-acid scale. |
| `sequence` | `str` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `str` | The tool slug that ran. |
| `window` | `int` | Sliding-window size (clamped to an odd number ≥ 1). |

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
| `chargeStep` | `float` | pH step for the net-charge titration curve (0–14). |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | `float` | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` | `str` |  |
| `length` | `int` | Number of residues to generate. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `enzymes` | `list` | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |
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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` | `str` |  |
| `ok` | `Any` |  |
| `organism` | `str` | Codon-usage host (ignored in degenerate mode). |
| `protein` | `str` | Protein sequence (one-letter codes; * for stop). |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `fileBase64` | `str` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `str` | Optional original file name (echoed back). |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | `float` | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `read` | `str` | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | `str` | Expected reference sequence (FASTA or raw). |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `args` | `dict` | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `str` | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `qualityOffset` | `int` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `accession` | `str` | GenBank/RefSeq accession (e.g. |
| `db` | `str` | Database to query; auto-detects from the accession format. |
| `format` | `str` | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `from` | `str` | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `str` | A FASTA or GenBank record to convert. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `to` | `str` | Output format. |
| `tool` | `str` | The tool slug that ran. |

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
| `endPrimerLength` | `int` | Length of the naive end primers taken from each end. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `int` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `int` | Minimum ORF length in amino acids. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `str` | Gene symbol/name, e.g. |
| `maxResults` | `int` | Up to 20. |
| `ok` | `Any` |  |
| `organism` | `str` | Organism name, e.g. |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `term` | `str` | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | `int` | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `reads` | `str` | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | `str` | The claimed/expected reference sequence. |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `entries` | `dict` | Initial named entries, e.g. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | `list` | Only return these entries; omit to return all of them. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sessionId` | `str` |  |
| `tool` | `str` | The tool slug that ran. |

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
| `args` | `dict` | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | `dict` | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sessionId` | `str` |  |
| `tool` | `str` | The tool slug that ran. |
| `writeBack` | `dict` | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

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
| `entries` | `dict` | Named entries to add/overwrite, e.g. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sessionId` | `str` |  |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | `int` | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `shRnaLoop` | `str` | Loop sequence used when assembling the shRNA cassette. |
| `target` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `armTmTarget` | `float` | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | `float` | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | `str` | Edit at the nucleotide or amino-acid level. |
| `frameStart` | `int` | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `float` | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | `str` | Replacement base (editKind='nt'). |
| `ok` | `Any` |  |
| `oligoNM` | `float` | Total strand concentration (nM). |
| `organism` | `str` | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | `int` | 1-based position to substitute (editKind='nt'). |
| `provenance` | `dict` |  |
| `residue` | `int` | 1-based residue number to change (editKind='aa'). |
| `result` | `dict` | Tool-specific output object. |
| `style` | `str` | Mutagenic primer style. |
| `targetAa` | `str` | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | `bool` | Stop at the first stop codon. |
| `tool` | `str` | The tool slug that ran. |

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
| `assembly` | `str` | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |
| `variant` | `str` | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

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
| `coding` | `bool` | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | `int` | 1-based reading-frame start (used when coding is true). |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `query` | `str` | Query / variant sequence (raw or FASTA). |
| `reference` | `str` | Reference / wild-type sequence (raw or FASTA). |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
| `armTmTarget` | `float` | Target annealing Tm (°C) for primer arms. |
| `circular` | `bool` | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | `str` | The sequence you claim you ended up with. |
| `coding` | `bool` | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | `str` | Type IIS enzyme for Golden Gate. |
| `enzyme3` | `str` | 3′ enzyme (restriction method). |
| `enzyme5` | `str` | 5′ enzyme (restriction method). |
| `fragmentPcrs` | `list` | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | `list` | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | `int` | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `str` | Insert sequence (restriction method). |
| `insertPcr` | `dict` | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | `str` | Assembly method used. |
| `names` | `list` | Optional labels for each fragment. |
| `ok` | `Any` |  |
| `overlapLen` | `int` | Gibson homology-arm length (bp). |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |
| `vector` | `str` | Vector sequence (restriction method). |
| `vectorPcr` | `dict` | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

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
| `claimedConstruct` | `str` | The final sequence claimed to have been built. |
| `expectedFrameStart` | `int` | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | `str` | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | `str` | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | `str` | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | `int` | Mismatches tolerated per primer during PCR prediction. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `templateCircular` | `bool` | Treat insertTemplate as circular (e.g. |
| `tool` | `str` | The tool slug that ran. |

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
| `circular` | `bool` | Treat the sequence as circular (plasmid). |
| `enzymes` | `list` | Enzyme names to digest with. |
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | `str` | DNA ladder to plot alongside the sample lane. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `sequence` | `str` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `result` | `dict` | Tool-specific output object. |
| `rows` | `list` | Differential expression rows, one per gene. |
| `tool` | `str` | The tool slug that ran. |

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
| `gate` | `Any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | `float` | Maximum number of results to return (default 5, max 10). |
| `ok` | `Any` |  |
| `provenance` | `dict` |  |
| `query` | `str` | The search query. |
| `result` | `dict` | Tool-specific output object. |
| `tool` | `str` | The tool slug that ran. |

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
