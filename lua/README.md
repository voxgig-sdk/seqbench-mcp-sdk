# SeqbenchMcp Lua SDK



The Lua SDK for the SeqbenchMcp API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:AlphafoldLookup()` — each with the same small set of operations (`load`, `create`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/seqbench-mcp-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("seqbench-mcp_sdk")

local client = sdk.new({
  apikey = os.getenv("SEQBENCH_MCP_APIKEY"),
})
```

### 4. Create, update, and remove

```lua
-- Create
local created, err = client:AlphafoldLookup():create({ accession = "example_accession", ok = "example_ok", provenance = {}, result = {}, tool = "example_tool" })
if err then error(err) end

```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local batch, err = client:Batch():load()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Batch():load()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
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
cd lua && busted test/
```


## Reference

### SeqbenchMcpSDK

```lua
local sdk = require("seqbench-mcp_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### SeqbenchMcpSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
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
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `create` | `(reqdata, ctrl) -> any, err` | Create a new entity. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` / `create` | the entity record (a `table`) |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local batch, err = client:Batch():load()
    if err then error(err) end
    -- batch is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

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

Create an instance: `local alphafold_lookup = client:AlphafoldLookup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `string` | UniProt accession, e.g. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local alphafold_lookup, err = client:AlphafoldLookup():create({
  accession = "example_accession", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### AsoDesign

Create an instance: `local aso_design = client:AsoDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `number` | Total gapmer length (nt). |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
| `wing` | `number` | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

#### Example: Create

```lua
local aso_design, err = client:AsoDesign():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  target = "example_target", -- string
  tool = "example_tool", -- string
})
```


### BaseEditingDesign

Create an instance: `local base_editing_design = client:BaseEditingDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `editor` | `string` | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | `number` | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | `number` | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local base_editing_design, err = client:BaseEditingDesign():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  target = "example_target", -- string
  tool = "example_tool", -- string
})
```


### Batch

Create an instance: `local batch = client:Batch(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `args` | `table` | Shared tool arguments applied to every record. |
| `capped` | `boolean` | True if input exceeded the record limit. |
| `columns` | `table` |  |
| `count` | `number` |  |
| `errors` | `number` |  |
| `input` | `string` | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | `number` | Maximum records per call (500). |
| `provenance` | `table` |  |
| `rows` | `table` |  |
| `tool` | `string` | A batchable tool slug (see `GET /batch`). |

#### Example: Load

```lua
local batch, err = client:Batch():load()
```

#### Example: Create

```lua
local batch, err = client:Batch():create({
  capped = true, -- boolean
  columns = {}, -- table
  count = 1, -- number
  errors = 1, -- number
  input = "example_input", -- string
  limit = 1, -- number
  provenance = {}, -- table
  rows = {}, -- table
  tool = "example_tool", -- string
})
```


### BatchWorkflow

Create an instance: `local batch__workflow = client:BatchWorkflow(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `capped` | `boolean` |  |
| `columns` | `table` | Flattened "<step>·<tool>·<key>" column headers. |
| `count` | `number` |  |
| `errors` | `number` |  |
| `input` | `string` | Multi-FASTA text or one sequence per line. |
| `limit` | `number` | Maximum records per call (200). |
| `provenance` | `table` |  |
| `rows` | `table` |  |
| `steps` | `table` |  |

#### Example: Load

```lua
local batch__workflow, err = client:BatchWorkflow():load()
```

#### Example: Create

```lua
local batch__workflow, err = client:BatchWorkflow():create({
  capped = true, -- boolean
  columns = {}, -- table
  count = 1, -- number
  errors = 1, -- number
  input = "example_input", -- string
  limit = 1, -- number
  provenance = {}, -- table
  rows = {}, -- table
  steps = {}, -- table
})
```


### CharacterizeSequence

Create an instance: `local characterize_sequence = client:CharacterizeSequence(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `endPrimerLength` | `number` | Length of the naive end primers taken from each end. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `number` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `number` | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local characterize_sequence, err = client:CharacterizeSequence():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### CloningSimulate

Create an instance: `local cloning_simulate = client:CloningSimulate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `armTmTarget` | `number` | Target annealing Tm (°C) for primer arms. |
| `circular` | `boolean` | Produce a circular product. |
| `enzyme` | `string` | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | `string` | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | 5′ enzyme (restriction method). |
| `fragments` | `table` | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | Insert sequence (restriction method). |
| `method` | `string` | Assembly method. |
| `names` | `table` | Optional labels for each fragment. |
| `ok` | `any` |  |
| `overlapLen` | `number` | Gibson homology-arm length (bp). |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `vector` | `string` | Vector sequence (restriction method). |

#### Example: Create

```lua
local cloning_simulate, err = client:CloningSimulate():create({
  method = "example_method", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### CodonAdaptationIndex

Create an instance: `local codon_adaptation_index = client:CodonAdaptationIndex(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `frameStart` | `number` | 1-based position to start reading codons. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `provenance` | `table` |  |
| `rareThreshold` | `number` | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local codon_adaptation_index, err = client:CodonAdaptationIndex():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### CodonOptimize

Create an instance: `local codon_optimize = client:CodonOptimize(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `protein` | `string` | Protein sequence (one-letter codes). |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local codon_optimize, err = client:CodonOptimize():create({
  ok = "example_ok", -- any
  protein = "example_protein", -- string
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### ConstructAutofix

Create an instance: `local construct_autofix = client:ConstructAutofix(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoidEnzymes` | `table` | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | `number` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `number` | 1-based nucleotide where the reading frame begins. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `number` |  |
| `gcLow` | `number` |  |
| `gcWindow` | `number` |  |
| `homopolymerMin` | `number` |  |
| `maxPasses` | `number` | Repeat full passes until clean or no further progress. |
| `ok` | `any` |  |
| `organism` | `string` | Codon-usage table to prefer among synonymous options. |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local construct_autofix, err = client:ConstructAutofix():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### ConstructQc

Create an instance: `local construct_qc = client:ConstructQc(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoidEnzymes` | `table` | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | `number` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `number` | 1-based nucleotide where the reading frame begins. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `number` | GC% above this flags a GC-rich window. |
| `gcLow` | `number` | GC% below this flags an AT-rich window. |
| `gcWindow` | `number` | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | `number` | Minimum run length to flag a homopolymer. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local construct_qc, err = client:ConstructQc():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### CrisprGrnaDesign

Create an instance: `local crispr_grna_design = client:CrisprGrnaDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | `number` | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | `string` | Nuclease id. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `searchReverseStrand` | `boolean` | Also scan the reverse strand for guides. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local crispr_grna_design, err = client:CrisprGrnaDesign():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### CrisprHdrDonor

Create an instance: `local crispr_hdr_donor = client:CrisprHdrDonor(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `armLength` | `number` | Homology arm length (bp) on each side. |
| `blockPam` | `boolean` | When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut. |
| `designGenotypingPrimers` | `boolean` | Also design a primer pair (on the original targetSequence) whose product spans the edit site. |
| `editEnd` | `number` | 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. |
| `editStart` | `number` | 1-based start of the region being replaced. |
| `frameStart` | `number` | Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | `number` | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | `number` | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | `string` | Strand the guide's protospacer is on. |
| `nuclease` | `string` | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `replacement` | `string` | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | `table` | Tool-specific output object. |
| `targetSequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local crispr_hdr_donor, err = client:CrisprHdrDonor():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  replacement = "example_replacement", -- string
  result = {}, -- table
  targetSequence = "example_targetSequence", -- string
  tool = "example_tool", -- string
})
```


### CrisprOfftargetCheck

Create an instance: `local crispr_offtarget_check = client:CrisprOfftargetCheck(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | `string` | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` | `any` |  |
| `protospacer` | `string` | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local crispr_offtarget_check, err = client:CrisprOfftargetCheck():create({
  ok = "example_ok", -- any
  protospacer = "example_protospacer", -- string
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### CrossDimer

Create an instance: `local cross_dimer = client:CrossDimer(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequenceA` | `string` | First oligo (5'→3'). |
| `sequenceB` | `string` | Second oligo (5'→3'). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local cross_dimer, err = client:CrossDimer():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequenceA = "example_sequenceA", -- string
  sequenceB = "example_sequenceB", -- string
  tool = "example_tool", -- string
})
```


### DnaMolarity

Create an instance: `local dna_molarity = client:DnaMolarity(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `number` | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | `number` | Mass in nanograms. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | `string` | The tool slug that ran. |
| `type` | `string` | Molecule type. |
| `volumeUl` | `number` | Volume in microlitres (0 = unknown; needed for concentration). |

#### Example: Create

```lua
local dna_molarity, err = client:DnaMolarity():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### DoubleDigest

Create an instance: `local double_digest = client:DoubleDigest(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enzymeA` | `string` | First enzyme name (e.g. |
| `enzymeB` | `string` | Second enzyme name (e.g. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local double_digest, err = client:DoubleDigest():create({
  enzymeA = "example_enzymeA", -- string
  enzymeB = "example_enzymeB", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### ExportEchoPicklist

Create an instance: `local export_echo_picklist = client:ExportEchoPicklist(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `reactions` | `table` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local export_echo_picklist, err = client:ExportEchoPicklist():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  reactions = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### ExportOpentronsProtocol

Create an instance: `local export_opentrons_protocol = client:ExportOpentronsProtocol(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `protocolName` | `string` | Optional protocol name (used in the script's metadata). |
| `provenance` | `table` |  |
| `reactions` | `table` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local export_opentrons_protocol, err = client:ExportOpentronsProtocol():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  reactions = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### ExportPlateLayout

Create an instance: `local export_plate_layout = client:ExportPlateLayout(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `reactions` | `table` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local export_plate_layout, err = client:ExportPlateLayout():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  reactions = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### ExpressionHeatmapCluster

Create an instance: `local expression_heatmap_cluster = client:ExpressionHeatmapCluster(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `clusterCols` | `boolean` | Cluster (reorder) samples. |
| `clusterRows` | `boolean` | Cluster (reorder) genes. |
| `distanceMetric` | `string` | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `table` | Row (gene) labels. |
| `linkage` | `string` | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `samples` | `table` | Column (sample) labels. |
| `tool` | `string` | The tool slug that ran. |
| `values` | `table` | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | `boolean` | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

#### Example: Create

```lua
local expression_heatmap_cluster, err = client:ExpressionHeatmapCluster():create({
  genes = {}, -- table
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  samples = {}, -- table
  tool = "example_tool", -- string
  values = {}, -- table
})
```


### FastqQcReport

Create an instance: `local fastq_qc_report = client:FastqQcReport(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `qualityOffset` | `number` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local fastq_qc_report, err = client:FastqQcReport():create({
  input = "example_input", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### FastqTrim

Create an instance: `local fastq_trim = client:FastqTrim(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | `number` | Reads shorter than this after trimming are dropped. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `qualityOffset` | `number` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | `number` | 3' quality-trim threshold (Phred score). |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local fastq_trim, err = client:FastqTrim():create({
  input = "example_input", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### FindOrf

Create an instance: `local find_orf = client:FindOrf(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | `number` | Minimum protein length (aa) to report. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `requireStop` | `boolean` | Only report ORFs terminated by a stop codon. |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local find_orf, err = client:FindOrf():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### FormatSequence

Create an instance: `local format_sequence = client:FormatSequence(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `caseMode` | `string` |  |
| `convert` | `string` | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `reverse` | `boolean` | Reverse the sequence (no complement). |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | `boolean` | Remove digits, spaces and gaps (keep letters only). |
| `tool` | `string` | The tool slug that ran. |
| `width` | `number` | Line-wrap width; 0 = single line. |

#### Example: Create

```lua
local format_sequence, err = client:FormatSequence():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### FunctionalEnrichment

Create an instance: `local functional_enrichment = client:FunctionalEnrichment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `background` | `table` | Custom background/universe gene symbols. |
| `collections` | `table` | Which term collections to test. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `table` | Query gene symbols (human, e.g. |
| `maxTermSize` | `number` | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | `number` | Skip terms/pathways with fewer than this many background genes. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local functional_enrichment, err = client:FunctionalEnrichment():create({
  genes = {}, -- table
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### GcContent

Create an instance: `local gc_content = client:GcContent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local gc_content, err = client:GcContent():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### GeneDossier

Create an instance: `local gene_dossier = client:GeneDossier(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local gene_dossier, err = client:GeneDossier():create({
  gene = "example_gene", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### GeneExpression

Create an instance: `local gene_expression = client:GeneExpression(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local gene_expression, err = client:GeneExpression():create({
  gene = "example_gene", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### GeneModel

Create an instance: `local gene_model = client:GeneModel(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local gene_model, err = client:GeneModel():create({
  gene = "example_gene", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### GoldenGateFidelity

Create an instance: `local golden_gate_fidelity = client:GoldenGateFidelity(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `compareToNamedSet` | `string` | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | `string` | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `overhangs` | `table` | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `riskThreshold` | `number` | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local golden_gate_fidelity, err = client:GoldenGateFidelity():create({
  ok = "example_ok", -- any
  overhangs = {}, -- table
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### HgvsConvert

Create an instance: `local hgvs_convert = client:HgvsConvert(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `variant` | `string` | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

#### Example: Create

```lua
local hgvs_convert, err = client:HgvsConvert():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
  variant = "example_variant", -- string
})
```


### IdMapPoll

Create an instance: `local id_map_poll = client:IdMapPoll(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local id_map_poll, err = client:IdMapPoll():create({
  jobId = "example_jobId", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### IdMapSubmit

Create an instance: `local id_map_submit = client:IdMapSubmit(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from` | `string` | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | `table` | The ids to map, up to 1000 (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `taxId` | `string` | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | `string` | Target id type. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local id_map_submit, err = client:IdMapSubmit():create({
  from = "example_from", -- string
  ids = {}, -- table
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  to = "example_to", -- string
  tool = "example_tool", -- string
})
```


### InSilicoPcr

Create an instance: `local in_silico_pcr = client:InSilicoPcr(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `boolean` | Treat the template as circular (plasmid). |
| `forwardPrimer` | `string` | Primer 1, 5'→3'. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | Mismatches tolerated per primer. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `reversePrimer` | `string` | Primer 2, 5'→3' (order does not matter). |
| `template` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local in_silico_pcr, err = client:InSilicoPcr():create({
  forwardPrimer = "example_forwardPrimer", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  reversePrimer = "example_reversePrimer", -- string
  template = "example_template", -- string
  tool = "example_tool", -- string
})
```


### KaspPrimerDesign

Create an instance: `local kasp_primer_design = client:KaspPrimerDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addSecondaryMismatch` | `boolean` | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | `string` | First allele (single base) — gets the FAM tail. |
| `alleleB` | `string` | Second allele (single base) — gets the HEX tail. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | `number` | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | `number` | Minimum amplicon length for the common reverse primer. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `snpPosition` | `number` | 1-based position of the SNP on the forward strand. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | `number` | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local kasp_primer_design, err = client:KaspPrimerDesign():create({
  alleleA = "example_alleleA", -- string
  alleleB = "example_alleleB", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  snpPosition = 1, -- number
  target = "example_target", -- string
  tool = "example_tool", -- string
})
```


### ListTool

Create an instance: `local list_tool = client:ListTool(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```lua
local list_tool, err = client:ListTool():load()
```


### MeltingTemperature

Create an instance: `local melting_temperature = client:MeltingTemperature(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dntpMM` | `number` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `number` | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` |  |
| `oligoNM` | `number` | Total strand concentration (nM). |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | `number` | Optional target Tm (°C). |
| `tmTolerance` | `number` | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local melting_temperature, err = client:MeltingTemperature():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### MotifFinder

Create an instance: `local motif_finder = client:MotifFinder(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | Maximum allowed mismatches per match. |
| `motif` | `string` | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `searchReverseStrand` | `boolean` | Also search the reverse strand. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local motif_finder, err = client:MotifFinder():create({
  motif = "example_motif", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### MultipleSequenceAlignment

Create an instance: `local multiple_sequence_alignment = client:MultipleSequenceAlignment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local multiple_sequence_alignment, err = client:MultipleSequenceAlignment():create({
  input = "example_input", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### OligoAnalysi

Create an instance: `local oligo_analysi = client:OligoAnalysi(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dntpMM` | `number` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `number` | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` |  |
| `oligoNM` | `number` | Total strand concentration (nM). |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local oligo_analysi, err = client:OligoAnalysi():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### OrthologMap

Create an instance: `local ortholog_map = client:OrthologMap(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sourceSpecies` | `string` | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | `table` | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | `string` | Ensembl species slug to find homologs in (e.g. |
| `tool` | `string` | The tool slug that ran. |
| `type` | `string` | Homology type to return. |

#### Example: Create

```lua
local ortholog_map, err = client:OrthologMap():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  symbols = {}, -- table
  targetSpecies = "example_targetSpecies", -- string
  tool = "example_tool", -- string
})
```


### PairwiseAlignment

Create an instance: `local pairwise_alignment = client:PairwiseAlignment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gap` | `number` | Linear gap penalty (per gap position). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | `number` | Match score. |
| `mismatch` | `number` | Mismatch penalty. |
| `mode` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `seqA` | `string` | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | `string` | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local pairwise_alignment, err = client:PairwiseAlignment():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  seqA = "example_seqA", -- string
  seqB = "example_seqB", -- string
  tool = "example_tool", -- string
})
```


### ParseGenbank

Create an instance: `local parse_genbank = client:ParseGenbank(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `text` | `string` | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local parse_genbank, err = client:ParseGenbank():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  text = "example_text", -- string
  tool = "example_tool", -- string
})
```


### ParseSangerTrace

Create an instance: `local parse_sanger_trace = client:ParseSangerTrace(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `fileBase64` | `string` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | Optional original file name (echoed back). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local parse_sanger_trace, err = client:ParseSangerTrace():create({
  fileBase64 = "example_fileBase64", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### PlasmidAnnotate

Create an instance: `local plasmid_annotate = client:PlasmidAnnotate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local plasmid_annotate, err = client:PlasmidAnnotate():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### PlasmidDeepAnnotate

Create an instance: `local plasmid_deep_annotate = client:PlasmidDeepAnnotate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `boolean` | Treat the sequence as a circular plasmid (vs. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local plasmid_deep_annotate, err = client:PlasmidDeepAnnotate():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### PlasmidFullReport

Create an instance: `local plasmid_full_report = client:PlasmidFullReport(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `boolean` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
| `topN` | `number` | How many top-ranked backbone candidates to report. |

#### Example: Create

```lua
local plasmid_full_report, err = client:PlasmidFullReport():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### PlasmidIdentify

Create an instance: `local plasmid_identify = client:PlasmidIdentify(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `boolean` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
| `topN` | `number` | How many top-ranked backbone candidates to report. |

#### Example: Create

```lua
local plasmid_identify, err = client:PlasmidIdentify():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### PrimeEditingDesign

Create an instance: `local prime_editing_design = client:PrimeEditingDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `editEnd` | `number` | 1-based inclusive end of the region being changed. |
| `editStart` | `number` | 1-based inclusive start of the region being changed. |
| `frameStart` | `number` | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | `string` | Replacement bases (forward strand). |
| `ok` | `any` |  |
| `pbsLength` | `number` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `rttHomology` | `number` | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | `string` | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local prime_editing_design, err = client:PrimeEditingDesign():create({
  editEnd = 1, -- number
  editStart = 1, -- number
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  target = "example_target", -- string
  tool = "example_tool", -- string
})
```


### PrimeEditingTwinDesign

Create an instance: `local prime_editing_twin_design = client:PrimeEditingTwinDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | `string` | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` | `any` |  |
| `overlapLength` | `number` | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | `number` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `table` |  |
| `replaceEnd` | `number` | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | `number` | 1-based inclusive start of the region being replaced/deleted. |
| `result` | `table` | Tool-specific output object. |
| `target` | `string` | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local prime_editing_twin_design, err = client:PrimeEditingTwinDesign():create({
  newSequence = "example_newSequence", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  replaceEnd = 1, -- number
  replaceStart = 1, -- number
  result = {}, -- table
  target = "example_target", -- string
  tool = "example_tool", -- string
})
```


### PrimerDesign

Create an instance: `local primer_design = client:PrimerDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `ampliconMax` | `number` |  |
| `ampliconMin` | `number` |  |
| `dntpMM` | `number` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` | `number` |  |
| `gcMin` | `number` |  |
| `lenMax` | `number` |  |
| `lenMin` | `number` |  |
| `lenOpt` | `number` |  |
| `maxReturn` | `number` | Number of best pairs to return. |
| `mgMM` | `number` | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` |  |
| `oligoNM` | `number` | Total strand concentration (nM). |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `targetEnd` | `number` | 1-based inclusive end of the target region (optional). |
| `targetStart` | `number` | 1-based inclusive start of a region the product must span (optional). |
| `template` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` | `number` |  |
| `tmMaxDiff` | `number` | Max Tm difference within a pair (°C). |
| `tmMin` | `number` |  |
| `tmOpt` | `number` |  |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local primer_design, err = client:PrimerDesign():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  template = "example_template", -- string
  tool = "example_tool", -- string
})
```


### PrimerSpecificity

Create an instance: `local primer_specificity = client:PrimerSpecificity(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `forwardPrimer` | `string` | Forward primer, 5'→3'. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | `number` | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `reversePrimer` | `string` | Reverse primer, 5'→3'. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local primer_specificity, err = client:PrimerSpecificity():create({
  forwardPrimer = "example_forwardPrimer", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  reversePrimer = "example_reversePrimer", -- string
  tool = "example_tool", -- string
})
```


### ProteaseDigestion

Create an instance: `local protease_digestion = client:ProteaseDigestion(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | `number` | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | `number` | Cap on the number of returned peptides. |
| `minMass` | `number` | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | `number` | Allowed missed internal cleavages (0–2). |
| `ok` | `any` |  |
| `protease` | `string` | Protease or chemical cleavage agent. |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local protease_digestion, err = client:ProteaseDigestion():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### ProteinAnnotatePoll

Create an instance: `local protein_annotate_poll = client:ProteinAnnotatePoll(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local protein_annotate_poll, err = client:ProteinAnnotatePoll():create({
  jobId = "example_jobId", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### ProteinAnnotateSubmit

Create an instance: `local protein_annotate_submit = client:ProteinAnnotateSubmit(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `appl` | `string` | Restrict to one member database (e.g. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | `boolean` | Include GO-term cross-references. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local protein_annotate_submit, err = client:ProteinAnnotateSubmit():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### ProteinHydrophobicity

Create an instance: `local protein_hydrophobicity = client:ProteinHydrophobicity(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `scale` | `string` | Amino-acid scale. |
| `sequence` | `string` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | The tool slug that ran. |
| `window` | `number` | Sliding-window size (clamped to an odd number ≥ 1). |

#### Example: Create

```lua
local protein_hydrophobicity, err = client:ProteinHydrophobicity():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### ProteinProperty

Create an instance: `local protein_property = client:ProteinProperty(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `chargeStep` | `number` | pH step for the net-charge titration curve (0–14). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local protein_property, err = client:ProteinProperty():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### RandomSequence

Create an instance: `local random_sequence = client:RandomSequence(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | `number` | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` | `string` |  |
| `length` | `number` | Number of residues to generate. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local random_sequence, err = client:RandomSequence():create({
  length = 1, -- number
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### RestrictionSite

Create an instance: `local restriction_site = client:RestrictionSite(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enzymes` | `table` | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local restriction_site, err = client:RestrictionSite():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### ReverseComplement

Create an instance: `local reverse_complement = client:ReverseComplement(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
| `type` | `string` |  |

#### Example: Create

```lua
local reverse_complement, err = client:ReverseComplement():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### ReverseTranslate

Create an instance: `local reverse_translate = client:ReverseTranslate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` | `string` |  |
| `ok` | `any` |  |
| `organism` | `string` | Codon-usage host (ignored in degenerate mode). |
| `protein` | `string` | Protein sequence (one-letter codes; * for stop). |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local reverse_translate, err = client:ReverseTranslate():create({
  ok = "example_ok", -- any
  protein = "example_protein", -- string
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### RnaFold

Create an instance: `local rna_fold = client:RnaFold(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local rna_fold, err = client:RnaFold():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### SangerVsReference

Create an instance: `local sanger_vs_reference = client:SangerVsReference(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `fileBase64` | `string` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | Optional original file name (echoed back). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | `number` | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `read` | `string` | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | `string` | Expected reference sequence (FASTA or raw). |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local sanger_vs_reference, err = client:SangerVsReference():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  reference = "example_reference", -- string
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### SavePermalink

Create an instance: `local save_permalink = client:SavePermalink(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `args` | `table` | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local save_permalink, err = client:SavePermalink():create({
  args = {}, -- table
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### SeqfileStat

Create an instance: `local seqfile_stat = client:SeqfileStat(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `qualityOffset` | `number` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local seqfile_stat, err = client:SeqfileStat():create({
  input = "example_input", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### SequenceFetch

Create an instance: `local sequence_fetch = client:SequenceFetch(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `string` | GenBank/RefSeq accession (e.g. |
| `db` | `string` | Database to query; auto-detects from the accession format. |
| `format` | `string` | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local sequence_fetch, err = client:SequenceFetch():create({
  accession = "example_accession", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### SequenceFormatConvert

Create an instance: `local sequence_format_convert = client:SequenceFormatConvert(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from` | `string` | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | A FASTA or GenBank record to convert. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `to` | `string` | Output format. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local sequence_format_convert, err = client:SequenceFormatConvert():create({
  input = "example_input", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### SequenceReport

Create an instance: `local sequence_report = client:SequenceReport(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `endPrimerLength` | `number` | Length of the naive end primers taken from each end. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `number` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `number` | Minimum ORF length in amino acids. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local sequence_report, err = client:SequenceReport():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### SequenceSearch

Create an instance: `local sequence_search = client:SequenceSearch(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `db` | `string` |  |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Gene symbol/name, e.g. |
| `maxResults` | `number` | Up to 20. |
| `ok` | `any` |  |
| `organism` | `string` | Organism name, e.g. |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `term` | `string` | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local sequence_search, err = client:SequenceSearch():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### SequencingReadbackVerify

Create an instance: `local sequencing_readback_verify = client:SequencingReadbackVerify(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | `number` | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `reads` | `string` | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | `string` | The claimed/expected reference sequence. |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local sequencing_readback_verify, err = client:SequencingReadbackVerify():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  reads = "example_reads", -- string
  reference = "example_reference", -- string
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### SessionCreate

Create an instance: `local session_create = client:SessionCreate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entries` | `table` | Initial named entries, e.g. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local session_create, err = client:SessionCreate():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### SessionGet

Create an instance: `local session_get = client:SessionGet(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | `table` | Only return these entries; omit to return all of them. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sessionId` | `string` |  |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local session_get, err = client:SessionGet():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sessionId = "example_sessionId", -- string
  tool = "example_tool", -- string
})
```


### SessionRun

Create an instance: `local session_run = client:SessionRun(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `args` | `table` | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | `table` | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sessionId` | `string` |  |
| `tool` | `string` | The tool slug that ran. |
| `writeBack` | `table` | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

#### Example: Create

```lua
local session_run, err = client:SessionRun():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sessionId = "example_sessionId", -- string
  tool = "example_tool", -- string
})
```


### SessionSet

Create an instance: `local session_set = client:SessionSet(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entries` | `table` | Named entries to add/overwrite, e.g. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sessionId` | `string` |  |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local session_set, err = client:SessionSet():create({
  entries = {}, -- table
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sessionId = "example_sessionId", -- string
  tool = "example_tool", -- string
})
```


### SirnaDesign

Create an instance: `local sirna_design = client:SirnaDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | `number` | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `shRnaLoop` | `string` | Loop sequence used when assembling the shRNA cassette. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local sirna_design, err = client:SirnaDesign():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  target = "example_target", -- string
  tool = "example_tool", -- string
})
```


### SiteDirectedMutagenesi

Create an instance: `local site_directed_mutagenesi = client:SiteDirectedMutagenesi(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `armTmTarget` | `number` | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | `number` | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | `string` | Edit at the nucleotide or amino-acid level. |
| `frameStart` | `number` | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `number` | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | `string` | Replacement base (editKind='nt'). |
| `ok` | `any` |  |
| `oligoNM` | `number` | Total strand concentration (nM). |
| `organism` | `string` | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | `number` | 1-based position to substitute (editKind='nt'). |
| `provenance` | `table` |  |
| `residue` | `number` | 1-based residue number to change (editKind='aa'). |
| `result` | `table` | Tool-specific output object. |
| `style` | `string` | Mutagenic primer style. |
| `targetAa` | `string` | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local site_directed_mutagenesi, err = client:SiteDirectedMutagenesi():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  template = "example_template", -- string
  tool = "example_tool", -- string
})
```


### Translate

Create an instance: `local translate = client:Translate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `frame` | `number` |  |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | `boolean` | Stop at the first stop codon. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local translate, err = client:Translate():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### VariantAnnotate

Create an instance: `local variant_annotate = client:VariantAnnotate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `assembly` | `string` | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `variant` | `string` | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

#### Example: Create

```lua
local variant_annotate, err = client:VariantAnnotate():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
  variant = "example_variant", -- string
})
```


### VariantComparator

Create an instance: `local variant_comparator = client:VariantComparator(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `coding` | `boolean` | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | `number` | 1-based reading-frame start (used when coding is true). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `query` | `string` | Query / variant sequence (raw or FASTA). |
| `reference` | `string` | Reference / wild-type sequence (raw or FASTA). |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local variant_comparator, err = client:VariantComparator():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  query = "example_query", -- string
  reference = "example_reference", -- string
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### VerifyAssembly

Create an instance: `local verify_assembly = client:VerifyAssembly(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `armTmTarget` | `number` | Target annealing Tm (°C) for primer arms. |
| `circular` | `boolean` | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | `string` | The sequence you claim you ended up with. |
| `coding` | `boolean` | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | `string` | Type IIS enzyme for Golden Gate. |
| `enzyme3` | `string` | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | 5′ enzyme (restriction method). |
| `fragmentPcrs` | `table` | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | `table` | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | `number` | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | Insert sequence (restriction method). |
| `insertPcr` | `table` | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | `string` | Assembly method used. |
| `names` | `table` | Optional labels for each fragment. |
| `ok` | `any` |  |
| `overlapLen` | `number` | Gibson homology-arm length (bp). |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `vector` | `string` | Vector sequence (restriction method). |
| `vectorPcr` | `table` | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

#### Example: Create

```lua
local verify_assembly, err = client:VerifyAssembly():create({
  claimedConstruct = "example_claimedConstruct", -- string
  method = "example_method", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### VerifyConstruct

Create an instance: `local verify_construct = client:VerifyConstruct(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `claimedConstruct` | `string` | The final sequence claimed to have been built. |
| `expectedFrameStart` | `number` | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | `string` | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | `string` | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | `string` | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | `number` | Mismatches tolerated per primer during PCR prediction. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `templateCircular` | `boolean` | Treat insertTemplate as circular (e.g. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local verify_construct, err = client:VerifyConstruct():create({
  claimedConstruct = "example_claimedConstruct", -- string
  insertForwardPrimer = "example_insertForwardPrimer", -- string
  insertReversePrimer = "example_insertReversePrimer", -- string
  insertTemplate = "example_insertTemplate", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  tool = "example_tool", -- string
})
```


### VirtualGel

Create an instance: `local virtual_gel = client:VirtualGel(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `boolean` | Treat the sequence as circular (plasmid). |
| `enzymes` | `table` | Enzyme names to digest with. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | `string` | DNA ladder to plot alongside the sample lane. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local virtual_gel, err = client:VirtualGel():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence = "example_sequence", -- string
  tool = "example_tool", -- string
})
```


### VolcanoPlotData

Create an instance: `local volcano_plot_data = client:VolcanoPlotData(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` | Tool-specific output object. |
| `rows` | `table` | Differential expression rows, one per gene. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local volcano_plot_data, err = client:VolcanoPlotData():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  rows = {}, -- table
  tool = "example_tool", -- string
})
```


### WebSearch

Create an instance: `local web_search = client:WebSearch(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | `number` | Maximum number of results to return (default 5, max 10). |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `query` | `string` | The search query. |
| `result` | `table` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```lua
local web_search, err = client:WebSearch():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  query = "example_query", -- string
  result = {}, -- table
  tool = "example_tool", -- string
})
```

## Features

This SDK ships 1 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`test`](#test) | In-memory mock transport for testing without a live server |

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.


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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── seqbench-mcp_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`seqbench-mcp_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```lua
local batch = client:Batch()
batch:load()

-- batch:data_get() now returns the batch data from the last load
-- batch:match_get() returns the last match criteria
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
