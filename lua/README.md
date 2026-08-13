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

Create an instance: `local alphafold_lookup = client:AlphafoldLookup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `length` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `target` | `string` |  |
| `tool` | `string` |  |
| `wing` | `number` |  |

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
| `editor` | `string` |  |
| `frameStart` | `number` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `target` | `string` |  |
| `targetPosition` | `number` |  |
| `tool` | `string` |  |

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
| `args` | `table` |  |
| `capped` | `boolean` |  |
| `columns` | `table` |  |
| `count` | `number` |  |
| `errors` | `number` |  |
| `input` | `string` |  |
| `limit` | `number` |  |
| `provenance` | `table` |  |
| `rows` | `table` |  |
| `tool` | `string` |  |

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
| `columns` | `table` |  |
| `count` | `number` |  |
| `errors` | `number` |  |
| `input` | `string` |  |
| `limit` | `number` |  |
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
| `endPrimerLength` | `number` |  |
| `gate` | `any` |  |
| `maxOrfs` | `number` |  |
| `minOrfAa` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `armTmTarget` | `number` |  |
| `circular` | `boolean` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragments` | `table` |  |
| `gate` | `any` |  |
| `insert` | `string` |  |
| `method` | `string` |  |
| `names` | `table` |  |
| `ok` | `any` |  |
| `overlapLen` | `number` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |
| `vector` | `string` |  |

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
| `frameStart` | `number` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `provenance` | `table` |  |
| `rareThreshold` | `number` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `protein` | `string` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `avoidEnzymes` | `table` |  |
| `crypticOrfMinAa` | `number` |  |
| `frameStart` | `number` |  |
| `gate` | `any` |  |
| `gcHigh` | `number` |  |
| `gcLow` | `number` |  |
| `gcWindow` | `number` |  |
| `homopolymerMin` | `number` |  |
| `maxPasses` | `number` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `avoidEnzymes` | `table` |  |
| `crypticOrfMinAa` | `number` |  |
| `frameStart` | `number` |  |
| `gate` | `any` |  |
| `gcHigh` | `number` |  |
| `gcLow` | `number` |  |
| `gcWindow` | `number` |  |
| `homopolymerMin` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `minScore` | `number` |  |
| `nuclease` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `searchReverseStrand` | `boolean` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `armLength` | `number` |  |
| `blockPam` | `boolean` |  |
| `designGenotypingPrimers` | `boolean` |  |
| `editEnd` | `number` |  |
| `editStart` | `number` |  |
| `frameStart` | `number` |  |
| `gate` | `any` |  |
| `guideEnd` | `number` |  |
| `guideStart` | `number` |  |
| `guideStrand` | `string` |  |
| `nuclease` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `replacement` | `string` |  |
| `result` | `table` |  |
| `targetSequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `maxMismatches` | `number` |  |
| `nuclease` | `string` |  |
| `ok` | `any` |  |
| `protospacer` | `string` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequenceA` | `string` |  |
| `sequenceB` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `length` | `number` |  |
| `massNg` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |
| `volumeUl` | `number` |  |

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
| `enzymeA` | `string` |  |
| `enzymeB` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `reactions` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `protocolName` | `string` |  |
| `provenance` | `table` |  |
| `reactions` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `reactions` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `clusterCols` | `boolean` |  |
| `clusterRows` | `boolean` |  |
| `distanceMetric` | `string` |  |
| `gate` | `any` |  |
| `genes` | `table` |  |
| `linkage` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `samples` | `table` |  |
| `tool` | `string` |  |
| `values` | `table` |  |
| `zScoreRows` | `boolean` |  |

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
| `gate` | `any` |  |
| `input` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `qualityOffset` | `number` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `input` | `string` |  |
| `minLength` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `qualityOffset` | `number` |  |
| `qualityThreshold` | `number` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `minAaLength` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `requireStop` | `boolean` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `convert` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `reverse` | `boolean` |  |
| `sequence` | `string` |  |
| `stripNonLetters` | `boolean` |  |
| `tool` | `string` |  |
| `width` | `number` |  |

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
| `background` | `table` |  |
| `collections` | `table` |  |
| `gate` | `any` |  |
| `genes` | `table` |  |
| `maxTermSize` | `number` |  |
| `minTermSize` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `gene` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `gene` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `gene` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `compareToNamedSet` | `string` |  |
| `dataset` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `overhangs` | `table` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `riskThreshold` | `number` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |
| `variant` | `string` |  |

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
| `gate` | `any` |  |
| `jobId` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `from` | `string` |  |
| `gate` | `any` |  |
| `ids` | `table` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `taxId` | `string` |  |
| `to` | `string` |  |
| `tool` | `string` |  |

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
| `circular` | `boolean` |  |
| `forwardPrimer` | `string` |  |
| `gate` | `any` |  |
| `maxMismatches` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `reversePrimer` | `string` |  |
| `template` | `string` |  |
| `tool` | `string` |  |

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
| `addSecondaryMismatch` | `boolean` |  |
| `alleleA` | `string` |  |
| `alleleB` | `string` |  |
| `gate` | `any` |  |
| `maxAmplicon` | `number` |  |
| `minAmplicon` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `snpPosition` | `number` |  |
| `target` | `string` |  |
| `targetCoreTm` | `number` |  |
| `tool` | `string` |  |

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
| `dntpMM` | `number` |  |
| `gate` | `any` |  |
| `mgMM` | `number` |  |
| `naMM` | `number` |  |
| `ok` | `any` |  |
| `oligoNM` | `number` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `targetTm` | `number` |  |
| `tmTolerance` | `number` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `maxMismatches` | `number` |  |
| `motif` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `searchReverseStrand` | `boolean` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `input` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `dntpMM` | `number` |  |
| `gate` | `any` |  |
| `mgMM` | `number` |  |
| `naMM` | `number` |  |
| `ok` | `any` |  |
| `oligoNM` | `number` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sourceSpecies` | `string` |  |
| `symbols` | `table` |  |
| `targetSpecies` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |

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
| `gap` | `number` |  |
| `gate` | `any` |  |
| `match` | `number` |  |
| `mismatch` | `number` |  |
| `mode` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `seqA` | `string` |  |
| `seqB` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `text` | `string` |  |
| `tool` | `string` |  |

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
| `fileBase64` | `string` |  |
| `fileName` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `circular` | `boolean` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `circular` | `boolean` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `topN` | `number` |  |

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
| `circular` | `boolean` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `topN` | `number` |  |

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
| `editEnd` | `number` |  |
| `editStart` | `number` |  |
| `frameStart` | `number` |  |
| `gate` | `any` |  |
| `insertedSeq` | `string` |  |
| `ok` | `any` |  |
| `pbsLength` | `number` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `rttHomology` | `number` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `newSequence` | `string` |  |
| `ok` | `any` |  |
| `overlapLength` | `number` |  |
| `pbsLength` | `number` |  |
| `provenance` | `table` |  |
| `replaceEnd` | `number` |  |
| `replaceStart` | `number` |  |
| `result` | `table` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

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
| `dntpMM` | `number` |  |
| `gate` | `any` |  |
| `gcMax` | `number` |  |
| `gcMin` | `number` |  |
| `lenMax` | `number` |  |
| `lenMin` | `number` |  |
| `lenOpt` | `number` |  |
| `maxReturn` | `number` |  |
| `mgMM` | `number` |  |
| `naMM` | `number` |  |
| `ok` | `any` |  |
| `oligoNM` | `number` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `targetEnd` | `number` |  |
| `targetStart` | `number` |  |
| `template` | `string` |  |
| `tmMax` | `number` |  |
| `tmMaxDiff` | `number` |  |
| `tmMin` | `number` |  |
| `tmOpt` | `number` |  |
| `tool` | `string` |  |

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
| `forwardPrimer` | `string` |  |
| `gate` | `any` |  |
| `maxMismatches` | `number` |  |
| `maxProductLength` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `reversePrimer` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `maxMass` | `number` |  |
| `maxPeptides` | `number` |  |
| `minMass` | `number` |  |
| `missedCleavages` | `number` |  |
| `ok` | `any` |  |
| `protease` | `string` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `jobId` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `appl` | `string` |  |
| `gate` | `any` |  |
| `goterms` | `boolean` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `scale` | `string` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `window` | `number` |  |

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
| `chargeStep` | `number` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `gcContent` | `number` |  |
| `kind` | `string` |  |
| `length` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `enzymes` | `table` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
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
| `gate` | `any` |  |
| `mode` | `string` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `protein` | `string` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `fileBase64` | `string` |  |
| `fileName` | `string` |  |
| `gate` | `any` |  |
| `minCoverage` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `read` | `string` |  |
| `reference` | `string` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `args` | `table` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `input` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `qualityOffset` | `number` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `accession` | `string` |  |
| `db` | `string` |  |
| `format` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `from` | `string` |  |
| `gate` | `any` |  |
| `input` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `to` | `string` |  |
| `tool` | `string` |  |

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
| `endPrimerLength` | `number` |  |
| `gate` | `any` |  |
| `maxOrfs` | `number` |  |
| `minOrfAa` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `gene` | `string` |  |
| `maxResults` | `number` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `term` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `minSupportingReads` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `reads` | `string` |  |
| `reference` | `string` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `entries` | `table` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `names` | `table` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sessionId` | `string` |  |
| `tool` | `string` |  |

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
| `args` | `table` |  |
| `fromSession` | `table` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sessionId` | `string` |  |
| `tool` | `string` |  |
| `writeBack` | `table` |  |

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
| `entries` | `table` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sessionId` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `minReynolds` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `shRnaLoop` | `string` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

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
| `armTmTarget` | `number` |  |
| `dntpMM` | `number` |  |
| `editKind` | `string` |  |
| `frameStart` | `number` |  |
| `gate` | `any` |  |
| `mgMM` | `number` |  |
| `naMM` | `number` |  |
| `newBase` | `string` |  |
| `ok` | `any` |  |
| `oligoNM` | `number` |  |
| `organism` | `string` |  |
| `position` | `number` |  |
| `provenance` | `table` |  |
| `residue` | `number` |  |
| `result` | `table` |  |
| `style` | `string` |  |
| `targetAa` | `string` |  |
| `template` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `toStop` | `boolean` |  |
| `tool` | `string` |  |

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
| `assembly` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |
| `variant` | `string` |  |

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
| `coding` | `boolean` |  |
| `frameStart` | `number` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `query` | `string` |  |
| `reference` | `string` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
| `armTmTarget` | `number` |  |
| `circular` | `boolean` |  |
| `claimedConstruct` | `string` |  |
| `coding` | `boolean` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragmentPcrs` | `table` |  |
| `fragments` | `table` |  |
| `frameStart` | `number` |  |
| `gate` | `any` |  |
| `insert` | `string` |  |
| `insertPcr` | `table` |  |
| `method` | `string` |  |
| `names` | `table` |  |
| `ok` | `any` |  |
| `overlapLen` | `number` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |
| `vector` | `string` |  |
| `vectorPcr` | `table` |  |

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
| `claimedConstruct` | `string` |  |
| `expectedFrameStart` | `number` |  |
| `gate` | `any` |  |
| `insertForwardPrimer` | `string` |  |
| `insertReversePrimer` | `string` |  |
| `insertTemplate` | `string` |  |
| `maxPrimerMismatches` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `templateCircular` | `boolean` |  |
| `tool` | `string` |  |

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
| `circular` | `boolean` |  |
| `enzymes` | `table` |  |
| `gate` | `any` |  |
| `ladder` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `rows` | `table` |  |
| `tool` | `string` |  |

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
| `gate` | `any` |  |
| `max_results` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `query` | `string` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

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
