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

    local alphafold_lookup, err = client:AlphafoldLookup():load()
    if err then error(err) end
    -- alphafold_lookup is the loaded record

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
| `frame_start` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `target` |  |
| `target_position` |  |
| `tool` |  |

Operations: Create.

API path: `/base_editing_design`

#### Batch

| Field | Description |
| --- | --- |
| `arg` |  |
| `input` |  |
| `ok` |  |
| `result` |  |
| `tool` |  |

Operations: Create, Load.

API path: `/batch`

#### BatchWorkflow

| Field | Description |
| --- | --- |
| `input` |  |
| `ok` |  |
| `result` |  |
| `step` |  |

Operations: Create, Load.

API path: `/workflow`

#### CharacterizeSequence

| Field | Description |
| --- | --- |
| `end_primer_length` |  |
| `gate` |  |
| `max_orf` |  |
| `min_orf_aa` |  |
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
| `arm_tm_target` |  |
| `circular` |  |
| `enzyme` |  |
| `enzyme3` |  |
| `enzyme5` |  |
| `fragment` |  |
| `gate` |  |
| `insert` |  |
| `method` |  |
| `name` |  |
| `ok` |  |
| `overlap_len` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `vector` |  |

Operations: Create.

API path: `/cloning_simulate`

#### CodonAdaptationIndex

| Field | Description |
| --- | --- |
| `frame_start` |  |
| `gate` |  |
| `ok` |  |
| `organism` |  |
| `provenance` |  |
| `rare_threshold` |  |
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
| `avoid_enzyme` |  |
| `cryptic_orf_min_aa` |  |
| `frame_start` |  |
| `gate` |  |
| `gc_high` |  |
| `gc_low` |  |
| `gc_window` |  |
| `homopolymer_min` |  |
| `max_pass` |  |
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
| `avoid_enzyme` |  |
| `cryptic_orf_min_aa` |  |
| `frame_start` |  |
| `gate` |  |
| `gc_high` |  |
| `gc_low` |  |
| `gc_window` |  |
| `homopolymer_min` |  |
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
| `min_score` |  |
| `nuclease` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `search_reverse_strand` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/crispr_grna_design`

#### CrisprHdrDonor

| Field | Description |
| --- | --- |
| `arm_length` |  |
| `block_pam` |  |
| `design_genotyping_primer` |  |
| `edit_end` |  |
| `edit_start` |  |
| `frame_start` |  |
| `gate` |  |
| `guide_end` |  |
| `guide_start` |  |
| `guide_strand` |  |
| `nuclease` |  |
| `ok` |  |
| `provenance` |  |
| `replacement` |  |
| `result` |  |
| `target_sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/crispr_hdr_donor`

#### CrisprOfftargetCheck

| Field | Description |
| --- | --- |
| `gate` |  |
| `max_mismatch` |  |
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
| `sequence_a` |  |
| `sequence_b` |  |
| `tool` |  |

Operations: Create.

API path: `/cross_dimer`

#### DnaMolarity

| Field | Description |
| --- | --- |
| `gate` |  |
| `length` |  |
| `mass_ng` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |
| `type` |  |
| `volume_ul` |  |

Operations: Create.

API path: `/dna_molarity`

#### DoubleDigest

| Field | Description |
| --- | --- |
| `enzyme_a` |  |
| `enzyme_b` |  |
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
| `reaction` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/export_echo_picklist`

#### ExportOpentronsProtocol

| Field | Description |
| --- | --- |
| `gate` |  |
| `ok` |  |
| `protocol_name` |  |
| `provenance` |  |
| `reaction` |  |
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
| `reaction` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/export_plate_layout`

#### ExpressionHeatmapCluster

| Field | Description |
| --- | --- |
| `cluster_col` |  |
| `cluster_row` |  |
| `distance_metric` |  |
| `gate` |  |
| `gene` |  |
| `linkage` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sample` |  |
| `tool` |  |
| `value` |  |
| `z_score_row` |  |

Operations: Create.

API path: `/expression_heatmap_cluster`

#### FastqQcReport

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `ok` |  |
| `provenance` |  |
| `quality_offset` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/fastq_qc_report`

#### FastqTrim

| Field | Description |
| --- | --- |
| `gate` |  |
| `input` |  |
| `min_length` |  |
| `ok` |  |
| `provenance` |  |
| `quality_offset` |  |
| `quality_threshold` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/fastq_trim`

#### FindOrf

| Field | Description |
| --- | --- |
| `gate` |  |
| `min_aa_length` |  |
| `ok` |  |
| `provenance` |  |
| `require_stop` |  |
| `result` |  |
| `sequence` |  |
| `tool` |  |

Operations: Create.

API path: `/find_orfs`

#### FormatSequence

| Field | Description |
| --- | --- |
| `case_mode` |  |
| `convert` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reverse` |  |
| `sequence` |  |
| `strip_non_letter` |  |
| `tool` |  |
| `width` |  |

Operations: Create.

API path: `/format_sequence`

#### FunctionalEnrichment

| Field | Description |
| --- | --- |
| `background` |  |
| `collection` |  |
| `gate` |  |
| `gene` |  |
| `max_term_size` |  |
| `min_term_size` |  |
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
| `compare_to_named_set` |  |
| `dataset` |  |
| `gate` |  |
| `ok` |  |
| `overhang` |  |
| `provenance` |  |
| `result` |  |
| `risk_threshold` |  |
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
| `job_id` |  |
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
| `tax_id` |  |
| `to` |  |
| `tool` |  |

Operations: Create.

API path: `/id_map_submit`

#### InSilicoPcr

| Field | Description |
| --- | --- |
| `circular` |  |
| `forward_primer` |  |
| `gate` |  |
| `max_mismatch` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reverse_primer` |  |
| `template` |  |
| `tool` |  |

Operations: Create.

API path: `/in_silico_pcr`

#### KaspPrimerDesign

| Field | Description |
| --- | --- |
| `add_secondary_mismatch` |  |
| `allele_a` |  |
| `allele_b` |  |
| `gate` |  |
| `max_amplicon` |  |
| `min_amplicon` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `snp_position` |  |
| `target` |  |
| `target_core_tm` |  |
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
| `dntp_mm` |  |
| `gate` |  |
| `mg_mm` |  |
| `na_mm` |  |
| `ok` |  |
| `oligo_nm` |  |
| `provenance` |  |
| `result` |  |
| `sequence` |  |
| `target_tm` |  |
| `tm_tolerance` |  |
| `tool` |  |

Operations: Create.

API path: `/melting_temperature`

#### MotifFinder

| Field | Description |
| --- | --- |
| `gate` |  |
| `max_mismatch` |  |
| `motif` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `search_reverse_strand` |  |
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
| `dntp_mm` |  |
| `gate` |  |
| `mg_mm` |  |
| `na_mm` |  |
| `ok` |  |
| `oligo_nm` |  |
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
| `source_species` |  |
| `symbol` |  |
| `target_species` |  |
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
| `seq_a` |  |
| `seq_b` |  |
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
| `file_base64` |  |
| `file_name` |  |
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
| `top_n` |  |

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
| `top_n` |  |

Operations: Create.

API path: `/plasmid_identify`

#### PrimeEditingDesign

| Field | Description |
| --- | --- |
| `edit_end` |  |
| `edit_start` |  |
| `frame_start` |  |
| `gate` |  |
| `inserted_seq` |  |
| `ok` |  |
| `pbs_length` |  |
| `provenance` |  |
| `result` |  |
| `rtt_homology` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/prime_editing_design`

#### PrimeEditingTwinDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `new_sequence` |  |
| `ok` |  |
| `overlap_length` |  |
| `pbs_length` |  |
| `provenance` |  |
| `replace_end` |  |
| `replace_start` |  |
| `result` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/prime_editing_twin_design`

#### PrimerDesign

| Field | Description |
| --- | --- |
| `amplicon_max` |  |
| `amplicon_min` |  |
| `dntp_mm` |  |
| `gate` |  |
| `gc_max` |  |
| `gc_min` |  |
| `len_max` |  |
| `len_min` |  |
| `len_opt` |  |
| `max_return` |  |
| `mg_mm` |  |
| `na_mm` |  |
| `ok` |  |
| `oligo_nm` |  |
| `provenance` |  |
| `result` |  |
| `target_end` |  |
| `target_start` |  |
| `template` |  |
| `tm_max` |  |
| `tm_max_diff` |  |
| `tm_min` |  |
| `tm_opt` |  |
| `tool` |  |

Operations: Create.

API path: `/primer_design`

#### PrimerSpecificity

| Field | Description |
| --- | --- |
| `forward_primer` |  |
| `gate` |  |
| `max_mismatch` |  |
| `max_product_length` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `reverse_primer` |  |
| `tool` |  |

Operations: Create.

API path: `/primer_specificity`

#### ProteaseDigestion

| Field | Description |
| --- | --- |
| `gate` |  |
| `max_mass` |  |
| `max_peptide` |  |
| `min_mass` |  |
| `missed_cleavage` |  |
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
| `job_id` |  |
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
| `goterm` |  |
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
| `charge_step` |  |
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
| `gc_content` |  |
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
| `enzyme` |  |
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
| `file_base64` |  |
| `file_name` |  |
| `gate` |  |
| `min_coverage` |  |
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
| `arg` |  |
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
| `quality_offset` |  |
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
| `end_primer_length` |  |
| `gate` |  |
| `max_orf` |  |
| `min_orf_aa` |  |
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
| `max_result` |  |
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
| `min_supporting_read` |  |
| `ok` |  |
| `provenance` |  |
| `read` |  |
| `reference` |  |
| `result` |  |
| `tool` |  |

Operations: Create.

API path: `/sequencing_readback_verify`

#### SessionCreate

| Field | Description |
| --- | --- |
| `entry` |  |
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
| `name` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `session_id` |  |
| `tool` |  |

Operations: Create.

API path: `/session_get`

#### SessionRun

| Field | Description |
| --- | --- |
| `arg` |  |
| `from_session` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `session_id` |  |
| `tool` |  |
| `write_back` |  |

Operations: Create.

API path: `/session_run`

#### SessionSet

| Field | Description |
| --- | --- |
| `entry` |  |
| `gate` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `session_id` |  |
| `tool` |  |

Operations: Create.

API path: `/session_set`

#### SirnaDesign

| Field | Description |
| --- | --- |
| `gate` |  |
| `min_reynold` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `sh_rna_loop` |  |
| `target` |  |
| `tool` |  |

Operations: Create.

API path: `/sirna_design`

#### SiteDirectedMutagenesi

| Field | Description |
| --- | --- |
| `arm_tm_target` |  |
| `dntp_mm` |  |
| `edit_kind` |  |
| `frame_start` |  |
| `gate` |  |
| `mg_mm` |  |
| `na_mm` |  |
| `new_base` |  |
| `ok` |  |
| `oligo_nm` |  |
| `organism` |  |
| `position` |  |
| `provenance` |  |
| `residue` |  |
| `result` |  |
| `style` |  |
| `target_aa` |  |
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
| `to_stop` |  |
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
| `frame_start` |  |
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
| `arm_tm_target` |  |
| `circular` |  |
| `claimed_construct` |  |
| `coding` |  |
| `enzyme` |  |
| `enzyme3` |  |
| `enzyme5` |  |
| `fragment` |  |
| `fragment_pcr` |  |
| `frame_start` |  |
| `gate` |  |
| `insert` |  |
| `insert_pcr` |  |
| `method` |  |
| `name` |  |
| `ok` |  |
| `overlap_len` |  |
| `provenance` |  |
| `result` |  |
| `tool` |  |
| `vector` |  |
| `vector_pcr` |  |

Operations: Create.

API path: `/verify_assembly`

#### VerifyConstruct

| Field | Description |
| --- | --- |
| `claimed_construct` |  |
| `expected_frame_start` |  |
| `gate` |  |
| `insert_forward_primer` |  |
| `insert_reverse_primer` |  |
| `insert_template` |  |
| `max_primer_mismatch` |  |
| `ok` |  |
| `provenance` |  |
| `result` |  |
| `template_circular` |  |
| `tool` |  |

Operations: Create.

API path: `/verify_construct`

#### VirtualGel

| Field | Description |
| --- | --- |
| `circular` |  |
| `enzyme` |  |
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
| `row` |  |
| `tool` |  |

Operations: Create.

API path: `/volcano_plot_data`

#### WebSearch

| Field | Description |
| --- | --- |
| `gate` |  |
| `max_result` |  |
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
| `frame_start` | `number` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `target` | `string` |  |
| `target_position` | `number` |  |
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
| `arg` | `table` |  |
| `input` | `string` |  |
| `ok` | `any` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

#### Example: Load

```lua
local batch, err = client:Batch():load()
```

#### Example: Create

```lua
local batch, err = client:Batch():create({
  input = "example_input", -- string
  ok = "example_ok", -- any
  result = {}, -- table
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
| `input` | `string` |  |
| `ok` | `any` |  |
| `result` | `table` |  |
| `step` | `table` |  |

#### Example: Load

```lua
local batch__workflow, err = client:BatchWorkflow():load()
```

#### Example: Create

```lua
local batch__workflow, err = client:BatchWorkflow():create({
  input = "example_input", -- string
  ok = "example_ok", -- any
  result = {}, -- table
  step = {}, -- table
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
| `end_primer_length` | `number` |  |
| `gate` | `any` |  |
| `max_orf` | `number` |  |
| `min_orf_aa` | `number` |  |
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
| `arm_tm_target` | `number` |  |
| `circular` | `boolean` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragment` | `table` |  |
| `gate` | `any` |  |
| `insert` | `string` |  |
| `method` | `string` |  |
| `name` | `table` |  |
| `ok` | `any` |  |
| `overlap_len` | `number` |  |
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
| `frame_start` | `number` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `provenance` | `table` |  |
| `rare_threshold` | `number` |  |
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
| `avoid_enzyme` | `table` |  |
| `cryptic_orf_min_aa` | `number` |  |
| `frame_start` | `number` |  |
| `gate` | `any` |  |
| `gc_high` | `number` |  |
| `gc_low` | `number` |  |
| `gc_window` | `number` |  |
| `homopolymer_min` | `number` |  |
| `max_pass` | `number` |  |
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
| `avoid_enzyme` | `table` |  |
| `cryptic_orf_min_aa` | `number` |  |
| `frame_start` | `number` |  |
| `gate` | `any` |  |
| `gc_high` | `number` |  |
| `gc_low` | `number` |  |
| `gc_window` | `number` |  |
| `homopolymer_min` | `number` |  |
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
| `min_score` | `number` |  |
| `nuclease` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `search_reverse_strand` | `boolean` |  |
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
| `arm_length` | `number` |  |
| `block_pam` | `boolean` |  |
| `design_genotyping_primer` | `boolean` |  |
| `edit_end` | `number` |  |
| `edit_start` | `number` |  |
| `frame_start` | `number` |  |
| `gate` | `any` |  |
| `guide_end` | `number` |  |
| `guide_start` | `number` |  |
| `guide_strand` | `string` |  |
| `nuclease` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `replacement` | `string` |  |
| `result` | `table` |  |
| `target_sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local crispr_hdr_donor, err = client:CrisprHdrDonor():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  replacement = "example_replacement", -- string
  result = {}, -- table
  target_sequence = "example_target_sequence", -- string
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
| `max_mismatch` | `number` |  |
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
| `sequence_a` | `string` |  |
| `sequence_b` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local cross_dimer, err = client:CrossDimer():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sequence_a = "example_sequence_a", -- string
  sequence_b = "example_sequence_b", -- string
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
| `mass_ng` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |
| `volume_ul` | `number` |  |

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
| `enzyme_a` | `string` |  |
| `enzyme_b` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local double_digest, err = client:DoubleDigest():create({
  enzyme_a = "example_enzyme_a", -- string
  enzyme_b = "example_enzyme_b", -- string
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
| `reaction` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local export_echo_picklist, err = client:ExportEchoPicklist():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  reaction = {}, -- table
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
| `protocol_name` | `string` |  |
| `provenance` | `table` |  |
| `reaction` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local export_opentrons_protocol, err = client:ExportOpentronsProtocol():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  reaction = {}, -- table
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
| `reaction` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local export_plate_layout, err = client:ExportPlateLayout():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  reaction = {}, -- table
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
| `cluster_col` | `boolean` |  |
| `cluster_row` | `boolean` |  |
| `distance_metric` | `string` |  |
| `gate` | `any` |  |
| `gene` | `table` |  |
| `linkage` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sample` | `table` |  |
| `tool` | `string` |  |
| `value` | `table` |  |
| `z_score_row` | `boolean` |  |

#### Example: Create

```lua
local expression_heatmap_cluster, err = client:ExpressionHeatmapCluster():create({
  gene = {}, -- table
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  sample = {}, -- table
  tool = "example_tool", -- string
  value = {}, -- table
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
| `quality_offset` | `number` |  |
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
| `min_length` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `quality_offset` | `number` |  |
| `quality_threshold` | `number` |  |
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
| `min_aa_length` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `require_stop` | `boolean` |  |
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
| `case_mode` | `string` |  |
| `convert` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `reverse` | `boolean` |  |
| `sequence` | `string` |  |
| `strip_non_letter` | `boolean` |  |
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
| `collection` | `table` |  |
| `gate` | `any` |  |
| `gene` | `table` |  |
| `max_term_size` | `number` |  |
| `min_term_size` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local functional_enrichment, err = client:FunctionalEnrichment():create({
  gene = {}, -- table
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
| `compare_to_named_set` | `string` |  |
| `dataset` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `overhang` | `table` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `risk_threshold` | `number` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local golden_gate_fidelity, err = client:GoldenGateFidelity():create({
  ok = "example_ok", -- any
  overhang = {}, -- table
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
| `job_id` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local id_map_poll, err = client:IdMapPoll():create({
  job_id = "example_job_id", -- string
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
| `tax_id` | `string` |  |
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
| `forward_primer` | `string` |  |
| `gate` | `any` |  |
| `max_mismatch` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `reverse_primer` | `string` |  |
| `template` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local in_silico_pcr, err = client:InSilicoPcr():create({
  forward_primer = "example_forward_primer", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  reverse_primer = "example_reverse_primer", -- string
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
| `add_secondary_mismatch` | `boolean` |  |
| `allele_a` | `string` |  |
| `allele_b` | `string` |  |
| `gate` | `any` |  |
| `max_amplicon` | `number` |  |
| `min_amplicon` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `snp_position` | `number` |  |
| `target` | `string` |  |
| `target_core_tm` | `number` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local kasp_primer_design, err = client:KaspPrimerDesign():create({
  allele_a = "example_allele_a", -- string
  allele_b = "example_allele_b", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  snp_position = 1, -- number
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
| `dntp_mm` | `number` |  |
| `gate` | `any` |  |
| `mg_mm` | `number` |  |
| `na_mm` | `number` |  |
| `ok` | `any` |  |
| `oligo_nm` | `number` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sequence` | `string` |  |
| `target_tm` | `number` |  |
| `tm_tolerance` | `number` |  |
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
| `max_mismatch` | `number` |  |
| `motif` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `search_reverse_strand` | `boolean` |  |
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
| `dntp_mm` | `number` |  |
| `gate` | `any` |  |
| `mg_mm` | `number` |  |
| `na_mm` | `number` |  |
| `ok` | `any` |  |
| `oligo_nm` | `number` |  |
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
| `source_species` | `string` |  |
| `symbol` | `table` |  |
| `target_species` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |

#### Example: Create

```lua
local ortholog_map, err = client:OrthologMap():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  symbol = {}, -- table
  target_species = "example_target_species", -- string
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
| `seq_a` | `string` |  |
| `seq_b` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local pairwise_alignment, err = client:PairwiseAlignment():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  seq_a = "example_seq_a", -- string
  seq_b = "example_seq_b", -- string
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
| `file_base64` | `string` |  |
| `file_name` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local parse_sanger_trace, err = client:ParseSangerTrace():create({
  file_base64 = "example_file_base64", -- string
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
| `top_n` | `number` |  |

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
| `top_n` | `number` |  |

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
| `edit_end` | `number` |  |
| `edit_start` | `number` |  |
| `frame_start` | `number` |  |
| `gate` | `any` |  |
| `inserted_seq` | `string` |  |
| `ok` | `any` |  |
| `pbs_length` | `number` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `rtt_homology` | `number` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local prime_editing_design, err = client:PrimeEditingDesign():create({
  edit_end = 1, -- number
  edit_start = 1, -- number
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
| `new_sequence` | `string` |  |
| `ok` | `any` |  |
| `overlap_length` | `number` |  |
| `pbs_length` | `number` |  |
| `provenance` | `table` |  |
| `replace_end` | `number` |  |
| `replace_start` | `number` |  |
| `result` | `table` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local prime_editing_twin_design, err = client:PrimeEditingTwinDesign():create({
  new_sequence = "example_new_sequence", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  replace_end = 1, -- number
  replace_start = 1, -- number
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
| `amplicon_max` | `number` |  |
| `amplicon_min` | `number` |  |
| `dntp_mm` | `number` |  |
| `gate` | `any` |  |
| `gc_max` | `number` |  |
| `gc_min` | `number` |  |
| `len_max` | `number` |  |
| `len_min` | `number` |  |
| `len_opt` | `number` |  |
| `max_return` | `number` |  |
| `mg_mm` | `number` |  |
| `na_mm` | `number` |  |
| `ok` | `any` |  |
| `oligo_nm` | `number` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `target_end` | `number` |  |
| `target_start` | `number` |  |
| `template` | `string` |  |
| `tm_max` | `number` |  |
| `tm_max_diff` | `number` |  |
| `tm_min` | `number` |  |
| `tm_opt` | `number` |  |
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
| `forward_primer` | `string` |  |
| `gate` | `any` |  |
| `max_mismatch` | `number` |  |
| `max_product_length` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `reverse_primer` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local primer_specificity, err = client:PrimerSpecificity():create({
  forward_primer = "example_forward_primer", -- string
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  reverse_primer = "example_reverse_primer", -- string
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
| `max_mass` | `number` |  |
| `max_peptide` | `number` |  |
| `min_mass` | `number` |  |
| `missed_cleavage` | `number` |  |
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
| `job_id` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local protein_annotate_poll, err = client:ProteinAnnotatePoll():create({
  job_id = "example_job_id", -- string
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
| `goterm` | `boolean` |  |
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
| `charge_step` | `number` |  |
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
| `gc_content` | `number` |  |
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
| `enzyme` | `table` |  |
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
| `file_base64` | `string` |  |
| `file_name` | `string` |  |
| `gate` | `any` |  |
| `min_coverage` | `number` |  |
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
| `arg` | `table` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local save_permalink, err = client:SavePermalink():create({
  arg = {}, -- table
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
| `quality_offset` | `number` |  |
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
| `end_primer_length` | `number` |  |
| `gate` | `any` |  |
| `max_orf` | `number` |  |
| `min_orf_aa` | `number` |  |
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
| `max_result` | `number` |  |
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
| `min_supporting_read` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `read` | `string` |  |
| `reference` | `string` |  |
| `result` | `table` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local sequencing_readback_verify, err = client:SequencingReadbackVerify():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  read = "example_read", -- string
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
| `entry` | `table` |  |
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
| `name` | `table` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local session_get, err = client:SessionGet():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  session_id = "example_session_id", -- string
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
| `arg` | `table` |  |
| `from_session` | `table` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |
| `write_back` | `table` |  |

#### Example: Create

```lua
local session_run, err = client:SessionRun():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  session_id = "example_session_id", -- string
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
| `entry` | `table` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local session_set, err = client:SessionSet():create({
  entry = {}, -- table
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  session_id = "example_session_id", -- string
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
| `min_reynold` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `sh_rna_loop` | `string` |  |
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
| `arm_tm_target` | `number` |  |
| `dntp_mm` | `number` |  |
| `edit_kind` | `string` |  |
| `frame_start` | `number` |  |
| `gate` | `any` |  |
| `mg_mm` | `number` |  |
| `na_mm` | `number` |  |
| `new_base` | `string` |  |
| `ok` | `any` |  |
| `oligo_nm` | `number` |  |
| `organism` | `string` |  |
| `position` | `number` |  |
| `provenance` | `table` |  |
| `residue` | `number` |  |
| `result` | `table` |  |
| `style` | `string` |  |
| `target_aa` | `string` |  |
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
| `to_stop` | `boolean` |  |
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
| `frame_start` | `number` |  |
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
| `arm_tm_target` | `number` |  |
| `circular` | `boolean` |  |
| `claimed_construct` | `string` |  |
| `coding` | `boolean` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragment` | `table` |  |
| `fragment_pcr` | `table` |  |
| `frame_start` | `number` |  |
| `gate` | `any` |  |
| `insert` | `string` |  |
| `insert_pcr` | `table` |  |
| `method` | `string` |  |
| `name` | `table` |  |
| `ok` | `any` |  |
| `overlap_len` | `number` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `tool` | `string` |  |
| `vector` | `string` |  |
| `vector_pcr` | `table` |  |

#### Example: Create

```lua
local verify_assembly, err = client:VerifyAssembly():create({
  claimed_construct = "example_claimed_construct", -- string
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
| `claimed_construct` | `string` |  |
| `expected_frame_start` | `number` |  |
| `gate` | `any` |  |
| `insert_forward_primer` | `string` |  |
| `insert_reverse_primer` | `string` |  |
| `insert_template` | `string` |  |
| `max_primer_mismatch` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `table` |  |
| `result` | `table` |  |
| `template_circular` | `boolean` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local verify_construct, err = client:VerifyConstruct():create({
  claimed_construct = "example_claimed_construct", -- string
  insert_forward_primer = "example_insert_forward_primer", -- string
  insert_reverse_primer = "example_insert_reverse_primer", -- string
  insert_template = "example_insert_template", -- string
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
| `enzyme` | `table` |  |
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
| `row` | `table` |  |
| `tool` | `string` |  |

#### Example: Create

```lua
local volcano_plot_data, err = client:VolcanoPlotData():create({
  ok = "example_ok", -- any
  provenance = {}, -- table
  result = {}, -- table
  row = {}, -- table
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
| `max_result` | `number` |  |
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
