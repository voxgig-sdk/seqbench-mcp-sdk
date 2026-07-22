# SeqbenchMcp Ruby SDK



The Ruby SDK for the SeqbenchMcp API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.AlphafoldLookup` — with named operations (`load`/`create`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/seqbench-mcp-sdk/releases](https://github.com/voxgig-sdk/seqbench-mcp-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "SeqbenchMcp_sdk"

client = SeqbenchMcpSDK.new({
  "apikey" => ENV["SEQBENCH_MCP_APIKEY"],
})
```

### 4. Create, update, and remove

```ruby
# create returns the bare created AlphafoldLookup record.
created = client.AlphafoldLookup.create({ "accession" => "example_accession", "ok" => "example_ok", "provenance" => {}, "result" => {}, "tool" => "example_tool" })

```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  batch = client.Batch.load()
rescue => err
  warn "load failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required:

```ruby
client = SeqbenchMcpSDK.test

# Entity ops return the bare mock record (raises on error).
batch = client.Batch.load()
puts batch
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = SeqbenchMcpSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
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
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### SeqbenchMcpSDK

```ruby
require_relative "SeqbenchMcp_sdk"
client = SeqbenchMcpSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = SeqbenchMcpSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### SeqbenchMcpSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
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
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `SeqbenchMcpError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

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

Create an instance: `alphafold_lookup = client.AlphafoldLookup`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `String` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
alphafold_lookup = client.AlphafoldLookup.create({
  "accession" => "example_accession", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### AsoDesign

Create an instance: `aso_design = client.AsoDesign`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `length` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `target` | `String` |  |
| `tool` | `String` |  |
| `wing` | `Integer` |  |

#### Example: Create

```ruby
aso_design = client.AsoDesign.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```


### BaseEditingDesign

Create an instance: `base_editing_design = client.BaseEditingDesign`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `editor` | `String` |  |
| `frame_start` | `Integer` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `target` | `String` |  |
| `target_position` | `Integer` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
base_editing_design = client.BaseEditingDesign.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```


### Batch

Create an instance: `batch = client.Batch`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `Hash` |  |
| `input` | `String` |  |
| `ok` | `Object` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Batch record (raises on error).
batch = client.Batch.load()
```

#### Example: Create

```ruby
batch = client.Batch.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### BatchWorkflow

Create an instance: `batch__workflow = client.BatchWorkflow`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `input` | `String` |  |
| `ok` | `Object` |  |
| `result` | `Hash` |  |
| `step` | `Array` |  |

#### Example: Load

```ruby
# load returns the bare BatchWorkflow record (raises on error).
batch__workflow = client.BatchWorkflow.load()
```

#### Example: Create

```ruby
batch__workflow = client.BatchWorkflow.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "result" => {}, # Hash
  "step" => [], # Array
})
```


### CharacterizeSequence

Create an instance: `characterize_sequence = client.CharacterizeSequence`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `end_primer_length` | `Integer` |  |
| `gate` | `Object` |  |
| `max_orf` | `Integer` |  |
| `min_orf_aa` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
characterize_sequence = client.CharacterizeSequence.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### CloningSimulate

Create an instance: `cloning_simulate = client.CloningSimulate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_tm_target` | `Float` |  |
| `circular` | `Boolean` |  |
| `enzyme` | `String` |  |
| `enzyme3` | `String` |  |
| `enzyme5` | `String` |  |
| `fragment` | `Array` |  |
| `gate` | `Object` |  |
| `insert` | `String` |  |
| `method` | `String` |  |
| `name` | `Array` |  |
| `ok` | `Object` |  |
| `overlap_len` | `Integer` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |
| `vector` | `String` |  |

#### Example: Create

```ruby
cloning_simulate = client.CloningSimulate.create({
  "method" => "example_method", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### CodonAdaptationIndex

Create an instance: `codon_adaptation_index = client.CodonAdaptationIndex`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `frame_start` | `Integer` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `organism` | `String` |  |
| `provenance` | `Hash` |  |
| `rare_threshold` | `Float` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
codon_adaptation_index = client.CodonAdaptationIndex.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### CodonOptimize

Create an instance: `codon_optimize = client.CodonOptimize`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `organism` | `String` |  |
| `protein` | `String` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
codon_optimize = client.CodonOptimize.create({
  "ok" => "example_ok", # Object
  "protein" => "example_protein", # String
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### ConstructAutofix

Create an instance: `construct_autofix = client.ConstructAutofix`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoid_enzyme` | `Array` |  |
| `cryptic_orf_min_aa` | `Integer` |  |
| `frame_start` | `Integer` |  |
| `gate` | `Object` |  |
| `gc_high` | `Float` |  |
| `gc_low` | `Float` |  |
| `gc_window` | `Integer` |  |
| `homopolymer_min` | `Integer` |  |
| `max_pass` | `Integer` |  |
| `ok` | `Object` |  |
| `organism` | `String` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
construct_autofix = client.ConstructAutofix.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### ConstructQc

Create an instance: `construct_qc = client.ConstructQc`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoid_enzyme` | `Array` |  |
| `cryptic_orf_min_aa` | `Integer` |  |
| `frame_start` | `Integer` |  |
| `gate` | `Object` |  |
| `gc_high` | `Float` |  |
| `gc_low` | `Float` |  |
| `gc_window` | `Integer` |  |
| `homopolymer_min` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
construct_qc = client.ConstructQc.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### CrisprGrnaDesign

Create an instance: `crispr_grna_design = client.CrisprGrnaDesign`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `min_score` | `Float` |  |
| `nuclease` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `search_reverse_strand` | `Boolean` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
crispr_grna_design = client.CrisprGrnaDesign.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### CrisprHdrDonor

Create an instance: `crispr_hdr_donor = client.CrisprHdrDonor`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_length` | `Integer` |  |
| `block_pam` | `Boolean` |  |
| `design_genotyping_primer` | `Boolean` |  |
| `edit_end` | `Integer` |  |
| `edit_start` | `Integer` |  |
| `frame_start` | `Integer` |  |
| `gate` | `Object` |  |
| `guide_end` | `Integer` |  |
| `guide_start` | `Integer` |  |
| `guide_strand` | `String` |  |
| `nuclease` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `replacement` | `String` |  |
| `result` | `Hash` |  |
| `target_sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
crispr_hdr_donor = client.CrisprHdrDonor.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "replacement" => "example_replacement", # String
  "result" => {}, # Hash
  "target_sequence" => "example_target_sequence", # String
  "tool" => "example_tool", # String
})
```


### CrisprOfftargetCheck

Create an instance: `crispr_offtarget_check = client.CrisprOfftargetCheck`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `max_mismatch` | `Integer` |  |
| `nuclease` | `String` |  |
| `ok` | `Object` |  |
| `protospacer` | `String` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
crispr_offtarget_check = client.CrisprOfftargetCheck.create({
  "ok" => "example_ok", # Object
  "protospacer" => "example_protospacer", # String
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### CrossDimer

Create an instance: `cross_dimer = client.CrossDimer`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence_a` | `String` |  |
| `sequence_b` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
cross_dimer = client.CrossDimer.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence_a" => "example_sequence_a", # String
  "sequence_b" => "example_sequence_b", # String
  "tool" => "example_tool", # String
})
```


### DnaMolarity

Create an instance: `dna_molarity = client.DnaMolarity`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `length` | `Integer` |  |
| `mass_ng` | `Float` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |
| `type` | `String` |  |
| `volume_ul` | `Float` |  |

#### Example: Create

```ruby
dna_molarity = client.DnaMolarity.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### DoubleDigest

Create an instance: `double_digest = client.DoubleDigest`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enzyme_a` | `String` |  |
| `enzyme_b` | `String` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
double_digest = client.DoubleDigest.create({
  "enzyme_a" => "example_enzyme_a", # String
  "enzyme_b" => "example_enzyme_b", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### ExportEchoPicklist

Create an instance: `export_echo_picklist = client.ExportEchoPicklist`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `reaction` | `Array` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
export_echo_picklist = client.ExportEchoPicklist.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reaction" => [], # Array
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### ExportOpentronsProtocol

Create an instance: `export_opentrons_protocol = client.ExportOpentronsProtocol`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `protocol_name` | `String` |  |
| `provenance` | `Hash` |  |
| `reaction` | `Array` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
export_opentrons_protocol = client.ExportOpentronsProtocol.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reaction" => [], # Array
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### ExportPlateLayout

Create an instance: `export_plate_layout = client.ExportPlateLayout`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `reaction` | `Array` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
export_plate_layout = client.ExportPlateLayout.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reaction" => [], # Array
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### ExpressionHeatmapCluster

Create an instance: `expression_heatmap_cluster = client.ExpressionHeatmapCluster`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cluster_col` | `Boolean` |  |
| `cluster_row` | `Boolean` |  |
| `distance_metric` | `String` |  |
| `gate` | `Object` |  |
| `gene` | `Array` |  |
| `linkage` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sample` | `Array` |  |
| `tool` | `String` |  |
| `value` | `Array` |  |
| `z_score_row` | `Boolean` |  |

#### Example: Create

```ruby
expression_heatmap_cluster = client.ExpressionHeatmapCluster.create({
  "gene" => [], # Array
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sample" => [], # Array
  "tool" => "example_tool", # String
  "value" => [], # Array
})
```


### FastqQcReport

Create an instance: `fastq_qc_report = client.FastqQcReport`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `input` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `quality_offset` | `Integer` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
fastq_qc_report = client.FastqQcReport.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### FastqTrim

Create an instance: `fastq_trim = client.FastqTrim`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `input` | `String` |  |
| `min_length` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `quality_offset` | `Integer` |  |
| `quality_threshold` | `Integer` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
fastq_trim = client.FastqTrim.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### FindOrf

Create an instance: `find_orf = client.FindOrf`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `min_aa_length` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `require_stop` | `Boolean` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
find_orf = client.FindOrf.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### FormatSequence

Create an instance: `format_sequence = client.FormatSequence`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `case_mode` | `String` |  |
| `convert` | `String` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `reverse` | `Boolean` |  |
| `sequence` | `String` |  |
| `strip_non_letter` | `Boolean` |  |
| `tool` | `String` |  |
| `width` | `Integer` |  |

#### Example: Create

```ruby
format_sequence = client.FormatSequence.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### FunctionalEnrichment

Create an instance: `functional_enrichment = client.FunctionalEnrichment`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `background` | `Array` |  |
| `collection` | `Array` |  |
| `gate` | `Object` |  |
| `gene` | `Array` |  |
| `max_term_size` | `Integer` |  |
| `min_term_size` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
functional_enrichment = client.FunctionalEnrichment.create({
  "gene" => [], # Array
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### GcContent

Create an instance: `gc_content = client.GcContent`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
gc_content = client.GcContent.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### GeneDossier

Create an instance: `gene_dossier = client.GeneDossier`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `gene` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
gene_dossier = client.GeneDossier.create({
  "gene" => "example_gene", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### GeneExpression

Create an instance: `gene_expression = client.GeneExpression`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `gene` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
gene_expression = client.GeneExpression.create({
  "gene" => "example_gene", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### GeneModel

Create an instance: `gene_model = client.GeneModel`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `gene` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
gene_model = client.GeneModel.create({
  "gene" => "example_gene", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### GoldenGateFidelity

Create an instance: `golden_gate_fidelity = client.GoldenGateFidelity`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `compare_to_named_set` | `String` |  |
| `dataset` | `String` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `overhang` | `Array` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `risk_threshold` | `Float` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
golden_gate_fidelity = client.GoldenGateFidelity.create({
  "ok" => "example_ok", # Object
  "overhang" => [], # Array
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### HgvsConvert

Create an instance: `hgvs_convert = client.HgvsConvert`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |
| `variant` | `String` |  |

#### Example: Create

```ruby
hgvs_convert = client.HgvsConvert.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
  "variant" => "example_variant", # String
})
```


### IdMapPoll

Create an instance: `id_map_poll = client.IdMapPoll`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `job_id` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
id_map_poll = client.IdMapPoll.create({
  "job_id" => "example_job_id", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### IdMapSubmit

Create an instance: `id_map_submit = client.IdMapSubmit`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from` | `String` |  |
| `gate` | `Object` |  |
| `ids` | `Array` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tax_id` | `String` |  |
| `to` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
id_map_submit = client.IdMapSubmit.create({
  "from" => "example_from", # String
  "ids" => [], # Array
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "to" => "example_to", # String
  "tool" => "example_tool", # String
})
```


### InSilicoPcr

Create an instance: `in_silico_pcr = client.InSilicoPcr`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `Boolean` |  |
| `forward_primer` | `String` |  |
| `gate` | `Object` |  |
| `max_mismatch` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `reverse_primer` | `String` |  |
| `template` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
in_silico_pcr = client.InSilicoPcr.create({
  "forward_primer" => "example_forward_primer", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "reverse_primer" => "example_reverse_primer", # String
  "template" => "example_template", # String
  "tool" => "example_tool", # String
})
```


### KaspPrimerDesign

Create an instance: `kasp_primer_design = client.KaspPrimerDesign`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `add_secondary_mismatch` | `Boolean` |  |
| `allele_a` | `String` |  |
| `allele_b` | `String` |  |
| `gate` | `Object` |  |
| `max_amplicon` | `Integer` |  |
| `min_amplicon` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `snp_position` | `Integer` |  |
| `target` | `String` |  |
| `target_core_tm` | `Float` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
kasp_primer_design = client.KaspPrimerDesign.create({
  "allele_a" => "example_allele_a", # String
  "allele_b" => "example_allele_b", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "snp_position" => 1, # Integer
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```


### ListTool

Create an instance: `list_tool = client.ListTool`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ruby
# load returns the bare ListTool record (raises on error).
list_tool = client.ListTool.load()
```


### MeltingTemperature

Create an instance: `melting_temperature = client.MeltingTemperature`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dntp_mm` | `Float` |  |
| `gate` | `Object` |  |
| `mg_mm` | `Float` |  |
| `na_mm` | `Float` |  |
| `ok` | `Object` |  |
| `oligo_nm` | `Float` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `target_tm` | `Float` |  |
| `tm_tolerance` | `Float` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
melting_temperature = client.MeltingTemperature.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### MotifFinder

Create an instance: `motif_finder = client.MotifFinder`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `max_mismatch` | `Integer` |  |
| `motif` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `search_reverse_strand` | `Boolean` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
motif_finder = client.MotifFinder.create({
  "motif" => "example_motif", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### MultipleSequenceAlignment

Create an instance: `multiple_sequence_alignment = client.MultipleSequenceAlignment`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `input` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
multiple_sequence_alignment = client.MultipleSequenceAlignment.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### OligoAnalysi

Create an instance: `oligo_analysi = client.OligoAnalysi`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dntp_mm` | `Float` |  |
| `gate` | `Object` |  |
| `mg_mm` | `Float` |  |
| `na_mm` | `Float` |  |
| `ok` | `Object` |  |
| `oligo_nm` | `Float` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
oligo_analysi = client.OligoAnalysi.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### OrthologMap

Create an instance: `ortholog_map = client.OrthologMap`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `source_species` | `String` |  |
| `symbol` | `Array` |  |
| `target_species` | `String` |  |
| `tool` | `String` |  |
| `type` | `String` |  |

#### Example: Create

```ruby
ortholog_map = client.OrthologMap.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "symbol" => [], # Array
  "target_species" => "example_target_species", # String
  "tool" => "example_tool", # String
})
```


### PairwiseAlignment

Create an instance: `pairwise_alignment = client.PairwiseAlignment`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gap` | `Float` |  |
| `gate` | `Object` |  |
| `match` | `Float` |  |
| `mismatch` | `Float` |  |
| `mode` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `seq_a` | `String` |  |
| `seq_b` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
pairwise_alignment = client.PairwiseAlignment.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "seq_a" => "example_seq_a", # String
  "seq_b" => "example_seq_b", # String
  "tool" => "example_tool", # String
})
```


### ParseGenbank

Create an instance: `parse_genbank = client.ParseGenbank`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `text` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
parse_genbank = client.ParseGenbank.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "text" => "example_text", # String
  "tool" => "example_tool", # String
})
```


### ParseSangerTrace

Create an instance: `parse_sanger_trace = client.ParseSangerTrace`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file_base64` | `String` |  |
| `file_name` | `String` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
parse_sanger_trace = client.ParseSangerTrace.create({
  "file_base64" => "example_file_base64", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### PlasmidAnnotate

Create an instance: `plasmid_annotate = client.PlasmidAnnotate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
plasmid_annotate = client.PlasmidAnnotate.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### PlasmidDeepAnnotate

Create an instance: `plasmid_deep_annotate = client.PlasmidDeepAnnotate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `Boolean` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
plasmid_deep_annotate = client.PlasmidDeepAnnotate.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### PlasmidFullReport

Create an instance: `plasmid_full_report = client.PlasmidFullReport`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `Boolean` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |
| `top_n` | `Integer` |  |

#### Example: Create

```ruby
plasmid_full_report = client.PlasmidFullReport.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### PlasmidIdentify

Create an instance: `plasmid_identify = client.PlasmidIdentify`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `Boolean` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |
| `top_n` | `Integer` |  |

#### Example: Create

```ruby
plasmid_identify = client.PlasmidIdentify.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### PrimeEditingDesign

Create an instance: `prime_editing_design = client.PrimeEditingDesign`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `edit_end` | `Integer` |  |
| `edit_start` | `Integer` |  |
| `frame_start` | `Integer` |  |
| `gate` | `Object` |  |
| `inserted_seq` | `String` |  |
| `ok` | `Object` |  |
| `pbs_length` | `Integer` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `rtt_homology` | `Integer` |  |
| `target` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
prime_editing_design = client.PrimeEditingDesign.create({
  "edit_end" => 1, # Integer
  "edit_start" => 1, # Integer
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```


### PrimeEditingTwinDesign

Create an instance: `prime_editing_twin_design = client.PrimeEditingTwinDesign`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `new_sequence` | `String` |  |
| `ok` | `Object` |  |
| `overlap_length` | `Integer` |  |
| `pbs_length` | `Integer` |  |
| `provenance` | `Hash` |  |
| `replace_end` | `Integer` |  |
| `replace_start` | `Integer` |  |
| `result` | `Hash` |  |
| `target` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
prime_editing_twin_design = client.PrimeEditingTwinDesign.create({
  "new_sequence" => "example_new_sequence", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "replace_end" => 1, # Integer
  "replace_start" => 1, # Integer
  "result" => {}, # Hash
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```


### PrimerDesign

Create an instance: `primer_design = client.PrimerDesign`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amplicon_max` | `Integer` |  |
| `amplicon_min` | `Integer` |  |
| `dntp_mm` | `Float` |  |
| `gate` | `Object` |  |
| `gc_max` | `Float` |  |
| `gc_min` | `Float` |  |
| `len_max` | `Integer` |  |
| `len_min` | `Integer` |  |
| `len_opt` | `Integer` |  |
| `max_return` | `Integer` |  |
| `mg_mm` | `Float` |  |
| `na_mm` | `Float` |  |
| `ok` | `Object` |  |
| `oligo_nm` | `Float` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `target_end` | `Integer` |  |
| `target_start` | `Integer` |  |
| `template` | `String` |  |
| `tm_max` | `Float` |  |
| `tm_max_diff` | `Float` |  |
| `tm_min` | `Float` |  |
| `tm_opt` | `Float` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
primer_design = client.PrimerDesign.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "template" => "example_template", # String
  "tool" => "example_tool", # String
})
```


### PrimerSpecificity

Create an instance: `primer_specificity = client.PrimerSpecificity`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `forward_primer` | `String` |  |
| `gate` | `Object` |  |
| `max_mismatch` | `Integer` |  |
| `max_product_length` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `reverse_primer` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
primer_specificity = client.PrimerSpecificity.create({
  "forward_primer" => "example_forward_primer", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "reverse_primer" => "example_reverse_primer", # String
  "tool" => "example_tool", # String
})
```


### ProteaseDigestion

Create an instance: `protease_digestion = client.ProteaseDigestion`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `max_mass` | `Float` |  |
| `max_peptide` | `Integer` |  |
| `min_mass` | `Float` |  |
| `missed_cleavage` | `Integer` |  |
| `ok` | `Object` |  |
| `protease` | `String` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
protease_digestion = client.ProteaseDigestion.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### ProteinAnnotatePoll

Create an instance: `protein_annotate_poll = client.ProteinAnnotatePoll`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `job_id` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
protein_annotate_poll = client.ProteinAnnotatePoll.create({
  "job_id" => "example_job_id", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### ProteinAnnotateSubmit

Create an instance: `protein_annotate_submit = client.ProteinAnnotateSubmit`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `appl` | `String` |  |
| `gate` | `Object` |  |
| `goterm` | `Boolean` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
protein_annotate_submit = client.ProteinAnnotateSubmit.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### ProteinHydrophobicity

Create an instance: `protein_hydrophobicity = client.ProteinHydrophobicity`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `scale` | `String` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |
| `window` | `Integer` |  |

#### Example: Create

```ruby
protein_hydrophobicity = client.ProteinHydrophobicity.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### ProteinProperty

Create an instance: `protein_property = client.ProteinProperty`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `charge_step` | `Float` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
protein_property = client.ProteinProperty.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### RandomSequence

Create an instance: `random_sequence = client.RandomSequence`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `gc_content` | `Float` |  |
| `kind` | `String` |  |
| `length` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
random_sequence = client.RandomSequence.create({
  "length" => 1, # Integer
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### RestrictionSite

Create an instance: `restriction_site = client.RestrictionSite`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enzyme` | `Array` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
restriction_site = client.RestrictionSite.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### ReverseComplement

Create an instance: `reverse_complement = client.ReverseComplement`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |
| `type` | `String` |  |

#### Example: Create

```ruby
reverse_complement = client.ReverseComplement.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### ReverseTranslate

Create an instance: `reverse_translate = client.ReverseTranslate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `mode` | `String` |  |
| `ok` | `Object` |  |
| `organism` | `String` |  |
| `protein` | `String` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
reverse_translate = client.ReverseTranslate.create({
  "ok" => "example_ok", # Object
  "protein" => "example_protein", # String
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### RnaFold

Create an instance: `rna_fold = client.RnaFold`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
rna_fold = client.RnaFold.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### SangerVsReference

Create an instance: `sanger_vs_reference = client.SangerVsReference`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file_base64` | `String` |  |
| `file_name` | `String` |  |
| `gate` | `Object` |  |
| `min_coverage` | `Float` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `read` | `String` |  |
| `reference` | `String` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
sanger_vs_reference = client.SangerVsReference.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reference" => "example_reference", # String
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### SavePermalink

Create an instance: `save_permalink = client.SavePermalink`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `Hash` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
save_permalink = client.SavePermalink.create({
  "arg" => {}, # Hash
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### SeqfileStat

Create an instance: `seqfile_stat = client.SeqfileStat`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `input` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `quality_offset` | `Integer` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
seqfile_stat = client.SeqfileStat.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### SequenceFetch

Create an instance: `sequence_fetch = client.SequenceFetch`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `String` |  |
| `db` | `String` |  |
| `format` | `String` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
sequence_fetch = client.SequenceFetch.create({
  "accession" => "example_accession", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### SequenceFormatConvert

Create an instance: `sequence_format_convert = client.SequenceFormatConvert`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from` | `String` |  |
| `gate` | `Object` |  |
| `input` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `to` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
sequence_format_convert = client.SequenceFormatConvert.create({
  "input" => "example_input", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### SequenceReport

Create an instance: `sequence_report = client.SequenceReport`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `end_primer_length` | `Integer` |  |
| `gate` | `Object` |  |
| `max_orf` | `Integer` |  |
| `min_orf_aa` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
sequence_report = client.SequenceReport.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### SequenceSearch

Create an instance: `sequence_search = client.SequenceSearch`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `db` | `String` |  |
| `gate` | `Object` |  |
| `gene` | `String` |  |
| `max_result` | `Integer` |  |
| `ok` | `Object` |  |
| `organism` | `String` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `term` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
sequence_search = client.SequenceSearch.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### SequencingReadbackVerify

Create an instance: `sequencing_readback_verify = client.SequencingReadbackVerify`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `min_supporting_read` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `read` | `String` |  |
| `reference` | `String` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
sequencing_readback_verify = client.SequencingReadbackVerify.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "read" => "example_read", # String
  "reference" => "example_reference", # String
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### SessionCreate

Create an instance: `session_create = client.SessionCreate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entry` | `Hash` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
session_create = client.SessionCreate.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### SessionGet

Create an instance: `session_get = client.SessionGet`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `name` | `Array` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `session_id` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
session_get = client.SessionGet.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "session_id" => "example_session_id", # String
  "tool" => "example_tool", # String
})
```


### SessionRun

Create an instance: `session_run = client.SessionRun`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `Hash` |  |
| `from_session` | `Hash` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `session_id` | `String` |  |
| `tool` | `String` |  |
| `write_back` | `Hash` |  |

#### Example: Create

```ruby
session_run = client.SessionRun.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "session_id" => "example_session_id", # String
  "tool" => "example_tool", # String
})
```


### SessionSet

Create an instance: `session_set = client.SessionSet`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entry` | `Hash` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `session_id` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
session_set = client.SessionSet.create({
  "entry" => {}, # Hash
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "session_id" => "example_session_id", # String
  "tool" => "example_tool", # String
})
```


### SirnaDesign

Create an instance: `sirna_design = client.SirnaDesign`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `min_reynold` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sh_rna_loop` | `String` |  |
| `target` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
sirna_design = client.SirnaDesign.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "target" => "example_target", # String
  "tool" => "example_tool", # String
})
```


### SiteDirectedMutagenesi

Create an instance: `site_directed_mutagenesi = client.SiteDirectedMutagenesi`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_tm_target` | `Float` |  |
| `dntp_mm` | `Float` |  |
| `edit_kind` | `String` |  |
| `frame_start` | `Integer` |  |
| `gate` | `Object` |  |
| `mg_mm` | `Float` |  |
| `na_mm` | `Float` |  |
| `new_base` | `String` |  |
| `ok` | `Object` |  |
| `oligo_nm` | `Float` |  |
| `organism` | `String` |  |
| `position` | `Integer` |  |
| `provenance` | `Hash` |  |
| `residue` | `Integer` |  |
| `result` | `Hash` |  |
| `style` | `String` |  |
| `target_aa` | `String` |  |
| `template` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
site_directed_mutagenesi = client.SiteDirectedMutagenesi.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "template" => "example_template", # String
  "tool" => "example_tool", # String
})
```


### Translate

Create an instance: `translate = client.Translate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `frame` | `Integer` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `to_stop` | `Boolean` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
translate = client.Translate.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### VariantAnnotate

Create an instance: `variant_annotate = client.VariantAnnotate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `assembly` | `String` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |
| `variant` | `String` |  |

#### Example: Create

```ruby
variant_annotate = client.VariantAnnotate.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
  "variant" => "example_variant", # String
})
```


### VariantComparator

Create an instance: `variant_comparator = client.VariantComparator`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `coding` | `Boolean` |  |
| `frame_start` | `Integer` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `query` | `String` |  |
| `reference` | `String` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
variant_comparator = client.VariantComparator.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "query" => "example_query", # String
  "reference" => "example_reference", # String
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### VerifyAssembly

Create an instance: `verify_assembly = client.VerifyAssembly`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_tm_target` | `Float` |  |
| `circular` | `Boolean` |  |
| `claimed_construct` | `String` |  |
| `coding` | `Boolean` |  |
| `enzyme` | `String` |  |
| `enzyme3` | `String` |  |
| `enzyme5` | `String` |  |
| `fragment` | `Array` |  |
| `fragment_pcr` | `Array` |  |
| `frame_start` | `Integer` |  |
| `gate` | `Object` |  |
| `insert` | `String` |  |
| `insert_pcr` | `Hash` |  |
| `method` | `String` |  |
| `name` | `Array` |  |
| `ok` | `Object` |  |
| `overlap_len` | `Integer` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |
| `vector` | `String` |  |
| `vector_pcr` | `Hash` |  |

#### Example: Create

```ruby
verify_assembly = client.VerifyAssembly.create({
  "claimed_construct" => "example_claimed_construct", # String
  "method" => "example_method", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### VerifyConstruct

Create an instance: `verify_construct = client.VerifyConstruct`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `claimed_construct` | `String` |  |
| `expected_frame_start` | `Integer` |  |
| `gate` | `Object` |  |
| `insert_forward_primer` | `String` |  |
| `insert_reverse_primer` | `String` |  |
| `insert_template` | `String` |  |
| `max_primer_mismatch` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `template_circular` | `Boolean` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
verify_construct = client.VerifyConstruct.create({
  "claimed_construct" => "example_claimed_construct", # String
  "insert_forward_primer" => "example_insert_forward_primer", # String
  "insert_reverse_primer" => "example_insert_reverse_primer", # String
  "insert_template" => "example_insert_template", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "tool" => "example_tool", # String
})
```


### VirtualGel

Create an instance: `virtual_gel = client.VirtualGel`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `Boolean` |  |
| `enzyme` | `Array` |  |
| `gate` | `Object` |  |
| `ladder` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
virtual_gel = client.VirtualGel.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequence" => "example_sequence", # String
  "tool" => "example_tool", # String
})
```


### VolcanoPlotData

Create an instance: `volcano_plot_data = client.VolcanoPlotData`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `row` | `Array` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
volcano_plot_data = client.VolcanoPlotData.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "row" => [], # Array
  "tool" => "example_tool", # String
})
```


### WebSearch

Create an instance: `web_search = client.WebSearch`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `Object` |  |
| `max_result` | `Float` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `query` | `String` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
web_search = client.WebSearch.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "query" => "example_query", # String
  "result" => {}, # Hash
  "tool" => "example_tool", # String
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

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── SeqbenchMcp_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`SeqbenchMcp_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```ruby
batch = client.Batch
batch.load()

# batch.data_get now returns the batch data from the last load
# batch.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
