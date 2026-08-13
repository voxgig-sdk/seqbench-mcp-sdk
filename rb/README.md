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
# create returns the ENTITY — call data_get for the created AlphafoldLookup record.
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

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
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
| `frameStart` | `Integer` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `target` | `String` |  |
| `targetPosition` | `Integer` |  |
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
| `args` | `Hash` |  |
| `capped` | `Boolean` |  |
| `columns` | `Array` |  |
| `count` | `Integer` |  |
| `errors` | `Integer` |  |
| `input` | `String` |  |
| `limit` | `Integer` |  |
| `provenance` | `Hash` |  |
| `rows` | `Array` |  |
| `tool` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Batch record (raises on error).
batch = client.Batch.load()
```

#### Example: Create

```ruby
batch = client.Batch.create({
  "capped" => true, # Boolean
  "columns" => [], # Array
  "count" => 1, # Integer
  "errors" => 1, # Integer
  "input" => "example_input", # String
  "limit" => 1, # Integer
  "provenance" => {}, # Hash
  "rows" => [], # Array
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
| `capped` | `Boolean` |  |
| `columns` | `Array` |  |
| `count` | `Integer` |  |
| `errors` | `Integer` |  |
| `input` | `String` |  |
| `limit` | `Integer` |  |
| `provenance` | `Hash` |  |
| `rows` | `Array` |  |
| `steps` | `Array` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the BatchWorkflow record (raises on error).
batch__workflow = client.BatchWorkflow.load()
```

#### Example: Create

```ruby
batch__workflow = client.BatchWorkflow.create({
  "capped" => true, # Boolean
  "columns" => [], # Array
  "count" => 1, # Integer
  "errors" => 1, # Integer
  "input" => "example_input", # String
  "limit" => 1, # Integer
  "provenance" => {}, # Hash
  "rows" => [], # Array
  "steps" => [], # Array
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
| `endPrimerLength` | `Integer` |  |
| `gate` | `Object` |  |
| `maxOrfs` | `Integer` |  |
| `minOrfAa` | `Integer` |  |
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
| `armTmTarget` | `Float` |  |
| `circular` | `Boolean` |  |
| `enzyme` | `String` |  |
| `enzyme3` | `String` |  |
| `enzyme5` | `String` |  |
| `fragments` | `Array` |  |
| `gate` | `Object` |  |
| `insert` | `String` |  |
| `method` | `String` |  |
| `names` | `Array` |  |
| `ok` | `Object` |  |
| `overlapLen` | `Integer` |  |
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
| `frameStart` | `Integer` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `organism` | `String` |  |
| `provenance` | `Hash` |  |
| `rareThreshold` | `Float` |  |
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
| `avoidEnzymes` | `Array` |  |
| `crypticOrfMinAa` | `Integer` |  |
| `frameStart` | `Integer` |  |
| `gate` | `Object` |  |
| `gcHigh` | `Float` |  |
| `gcLow` | `Float` |  |
| `gcWindow` | `Integer` |  |
| `homopolymerMin` | `Integer` |  |
| `maxPasses` | `Integer` |  |
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
| `avoidEnzymes` | `Array` |  |
| `crypticOrfMinAa` | `Integer` |  |
| `frameStart` | `Integer` |  |
| `gate` | `Object` |  |
| `gcHigh` | `Float` |  |
| `gcLow` | `Float` |  |
| `gcWindow` | `Integer` |  |
| `homopolymerMin` | `Integer` |  |
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
| `minScore` | `Float` |  |
| `nuclease` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `searchReverseStrand` | `Boolean` |  |
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
| `armLength` | `Integer` |  |
| `blockPam` | `Boolean` |  |
| `designGenotypingPrimers` | `Boolean` |  |
| `editEnd` | `Integer` |  |
| `editStart` | `Integer` |  |
| `frameStart` | `Integer` |  |
| `gate` | `Object` |  |
| `guideEnd` | `Integer` |  |
| `guideStart` | `Integer` |  |
| `guideStrand` | `String` |  |
| `nuclease` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `replacement` | `String` |  |
| `result` | `Hash` |  |
| `targetSequence` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
crispr_hdr_donor = client.CrisprHdrDonor.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "replacement" => "example_replacement", # String
  "result" => {}, # Hash
  "targetSequence" => "example_targetSequence", # String
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
| `maxMismatches` | `Integer` |  |
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
| `sequenceA` | `String` |  |
| `sequenceB` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
cross_dimer = client.CrossDimer.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sequenceA" => "example_sequenceA", # String
  "sequenceB" => "example_sequenceB", # String
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
| `massNg` | `Float` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `tool` | `String` |  |
| `type` | `String` |  |
| `volumeUl` | `Float` |  |

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
| `enzymeA` | `String` |  |
| `enzymeB` | `String` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
double_digest = client.DoubleDigest.create({
  "enzymeA" => "example_enzymeA", # String
  "enzymeB" => "example_enzymeB", # String
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
| `reactions` | `Array` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
export_echo_picklist = client.ExportEchoPicklist.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reactions" => [], # Array
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
| `protocolName` | `String` |  |
| `provenance` | `Hash` |  |
| `reactions` | `Array` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
export_opentrons_protocol = client.ExportOpentronsProtocol.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reactions" => [], # Array
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
| `reactions` | `Array` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
export_plate_layout = client.ExportPlateLayout.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reactions" => [], # Array
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
| `clusterCols` | `Boolean` |  |
| `clusterRows` | `Boolean` |  |
| `distanceMetric` | `String` |  |
| `gate` | `Object` |  |
| `genes` | `Array` |  |
| `linkage` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `samples` | `Array` |  |
| `tool` | `String` |  |
| `values` | `Array` |  |
| `zScoreRows` | `Boolean` |  |

#### Example: Create

```ruby
expression_heatmap_cluster = client.ExpressionHeatmapCluster.create({
  "genes" => [], # Array
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "samples" => [], # Array
  "tool" => "example_tool", # String
  "values" => [], # Array
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
| `qualityOffset` | `Integer` |  |
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
| `minLength` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `qualityOffset` | `Integer` |  |
| `qualityThreshold` | `Integer` |  |
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
| `minAaLength` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `requireStop` | `Boolean` |  |
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
| `caseMode` | `String` |  |
| `convert` | `String` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `reverse` | `Boolean` |  |
| `sequence` | `String` |  |
| `stripNonLetters` | `Boolean` |  |
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
| `collections` | `Array` |  |
| `gate` | `Object` |  |
| `genes` | `Array` |  |
| `maxTermSize` | `Integer` |  |
| `minTermSize` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
functional_enrichment = client.FunctionalEnrichment.create({
  "genes" => [], # Array
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
| `compareToNamedSet` | `String` |  |
| `dataset` | `String` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `overhangs` | `Array` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `riskThreshold` | `Float` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
golden_gate_fidelity = client.GoldenGateFidelity.create({
  "ok" => "example_ok", # Object
  "overhangs" => [], # Array
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
| `jobId` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
id_map_poll = client.IdMapPoll.create({
  "jobId" => "example_jobId", # String
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
| `taxId` | `String` |  |
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
| `forwardPrimer` | `String` |  |
| `gate` | `Object` |  |
| `maxMismatches` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `reversePrimer` | `String` |  |
| `template` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
in_silico_pcr = client.InSilicoPcr.create({
  "forwardPrimer" => "example_forwardPrimer", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "reversePrimer" => "example_reversePrimer", # String
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
| `addSecondaryMismatch` | `Boolean` |  |
| `alleleA` | `String` |  |
| `alleleB` | `String` |  |
| `gate` | `Object` |  |
| `maxAmplicon` | `Integer` |  |
| `minAmplicon` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `snpPosition` | `Integer` |  |
| `target` | `String` |  |
| `targetCoreTm` | `Float` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
kasp_primer_design = client.KaspPrimerDesign.create({
  "alleleA" => "example_alleleA", # String
  "alleleB" => "example_alleleB", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "snpPosition" => 1, # Integer
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
# load returns the ENTITY — call data_get for the ListTool record (raises on error).
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
| `dntpMM` | `Float` |  |
| `gate` | `Object` |  |
| `mgMM` | `Float` |  |
| `naMM` | `Float` |  |
| `ok` | `Object` |  |
| `oligoNM` | `Float` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sequence` | `String` |  |
| `targetTm` | `Float` |  |
| `tmTolerance` | `Float` |  |
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
| `maxMismatches` | `Integer` |  |
| `motif` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `searchReverseStrand` | `Boolean` |  |
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
| `dntpMM` | `Float` |  |
| `gate` | `Object` |  |
| `mgMM` | `Float` |  |
| `naMM` | `Float` |  |
| `ok` | `Object` |  |
| `oligoNM` | `Float` |  |
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
| `sourceSpecies` | `String` |  |
| `symbols` | `Array` |  |
| `targetSpecies` | `String` |  |
| `tool` | `String` |  |
| `type` | `String` |  |

#### Example: Create

```ruby
ortholog_map = client.OrthologMap.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "symbols" => [], # Array
  "targetSpecies" => "example_targetSpecies", # String
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
| `seqA` | `String` |  |
| `seqB` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
pairwise_alignment = client.PairwiseAlignment.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "seqA" => "example_seqA", # String
  "seqB" => "example_seqB", # String
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
| `fileBase64` | `String` |  |
| `fileName` | `String` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
parse_sanger_trace = client.ParseSangerTrace.create({
  "fileBase64" => "example_fileBase64", # String
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
| `topN` | `Integer` |  |

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
| `topN` | `Integer` |  |

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
| `editEnd` | `Integer` |  |
| `editStart` | `Integer` |  |
| `frameStart` | `Integer` |  |
| `gate` | `Object` |  |
| `insertedSeq` | `String` |  |
| `ok` | `Object` |  |
| `pbsLength` | `Integer` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `rttHomology` | `Integer` |  |
| `target` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
prime_editing_design = client.PrimeEditingDesign.create({
  "editEnd" => 1, # Integer
  "editStart" => 1, # Integer
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
| `newSequence` | `String` |  |
| `ok` | `Object` |  |
| `overlapLength` | `Integer` |  |
| `pbsLength` | `Integer` |  |
| `provenance` | `Hash` |  |
| `replaceEnd` | `Integer` |  |
| `replaceStart` | `Integer` |  |
| `result` | `Hash` |  |
| `target` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
prime_editing_twin_design = client.PrimeEditingTwinDesign.create({
  "newSequence" => "example_newSequence", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "replaceEnd" => 1, # Integer
  "replaceStart" => 1, # Integer
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
| `ampliconMax` | `Integer` |  |
| `ampliconMin` | `Integer` |  |
| `dntpMM` | `Float` |  |
| `gate` | `Object` |  |
| `gcMax` | `Float` |  |
| `gcMin` | `Float` |  |
| `lenMax` | `Integer` |  |
| `lenMin` | `Integer` |  |
| `lenOpt` | `Integer` |  |
| `maxReturn` | `Integer` |  |
| `mgMM` | `Float` |  |
| `naMM` | `Float` |  |
| `ok` | `Object` |  |
| `oligoNM` | `Float` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `targetEnd` | `Integer` |  |
| `targetStart` | `Integer` |  |
| `template` | `String` |  |
| `tmMax` | `Float` |  |
| `tmMaxDiff` | `Float` |  |
| `tmMin` | `Float` |  |
| `tmOpt` | `Float` |  |
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
| `forwardPrimer` | `String` |  |
| `gate` | `Object` |  |
| `maxMismatches` | `Integer` |  |
| `maxProductLength` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `reversePrimer` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
primer_specificity = client.PrimerSpecificity.create({
  "forwardPrimer" => "example_forwardPrimer", # String
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "reversePrimer" => "example_reversePrimer", # String
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
| `maxMass` | `Float` |  |
| `maxPeptides` | `Integer` |  |
| `minMass` | `Float` |  |
| `missedCleavages` | `Integer` |  |
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
| `jobId` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
protein_annotate_poll = client.ProteinAnnotatePoll.create({
  "jobId" => "example_jobId", # String
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
| `goterms` | `Boolean` |  |
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
| `chargeStep` | `Float` |  |
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
| `gcContent` | `Float` |  |
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
| `enzymes` | `Array` |  |
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
| `fileBase64` | `String` |  |
| `fileName` | `String` |  |
| `gate` | `Object` |  |
| `minCoverage` | `Float` |  |
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
| `args` | `Hash` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
save_permalink = client.SavePermalink.create({
  "args" => {}, # Hash
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
| `qualityOffset` | `Integer` |  |
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
| `endPrimerLength` | `Integer` |  |
| `gate` | `Object` |  |
| `maxOrfs` | `Integer` |  |
| `minOrfAa` | `Integer` |  |
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
| `maxResults` | `Integer` |  |
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
| `minSupportingReads` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `reads` | `String` |  |
| `reference` | `String` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
sequencing_readback_verify = client.SequencingReadbackVerify.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "reads" => "example_reads", # String
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
| `entries` | `Hash` |  |
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
| `names` | `Array` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sessionId` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
session_get = client.SessionGet.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sessionId" => "example_sessionId", # String
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
| `args` | `Hash` |  |
| `fromSession` | `Hash` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sessionId` | `String` |  |
| `tool` | `String` |  |
| `writeBack` | `Hash` |  |

#### Example: Create

```ruby
session_run = client.SessionRun.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sessionId" => "example_sessionId", # String
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
| `entries` | `Hash` |  |
| `gate` | `Object` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `sessionId` | `String` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
session_set = client.SessionSet.create({
  "entries" => {}, # Hash
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "sessionId" => "example_sessionId", # String
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
| `minReynolds` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `shRnaLoop` | `String` |  |
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
| `armTmTarget` | `Float` |  |
| `dntpMM` | `Float` |  |
| `editKind` | `String` |  |
| `frameStart` | `Integer` |  |
| `gate` | `Object` |  |
| `mgMM` | `Float` |  |
| `naMM` | `Float` |  |
| `newBase` | `String` |  |
| `ok` | `Object` |  |
| `oligoNM` | `Float` |  |
| `organism` | `String` |  |
| `position` | `Integer` |  |
| `provenance` | `Hash` |  |
| `residue` | `Integer` |  |
| `result` | `Hash` |  |
| `style` | `String` |  |
| `targetAa` | `String` |  |
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
| `toStop` | `Boolean` |  |
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
| `frameStart` | `Integer` |  |
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
| `armTmTarget` | `Float` |  |
| `circular` | `Boolean` |  |
| `claimedConstruct` | `String` |  |
| `coding` | `Boolean` |  |
| `enzyme` | `String` |  |
| `enzyme3` | `String` |  |
| `enzyme5` | `String` |  |
| `fragmentPcrs` | `Array` |  |
| `fragments` | `Array` |  |
| `frameStart` | `Integer` |  |
| `gate` | `Object` |  |
| `insert` | `String` |  |
| `insertPcr` | `Hash` |  |
| `method` | `String` |  |
| `names` | `Array` |  |
| `ok` | `Object` |  |
| `overlapLen` | `Integer` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `tool` | `String` |  |
| `vector` | `String` |  |
| `vectorPcr` | `Hash` |  |

#### Example: Create

```ruby
verify_assembly = client.VerifyAssembly.create({
  "claimedConstruct" => "example_claimedConstruct", # String
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
| `claimedConstruct` | `String` |  |
| `expectedFrameStart` | `Integer` |  |
| `gate` | `Object` |  |
| `insertForwardPrimer` | `String` |  |
| `insertReversePrimer` | `String` |  |
| `insertTemplate` | `String` |  |
| `maxPrimerMismatches` | `Integer` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` |  |
| `templateCircular` | `Boolean` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
verify_construct = client.VerifyConstruct.create({
  "claimedConstruct" => "example_claimedConstruct", # String
  "insertForwardPrimer" => "example_insertForwardPrimer", # String
  "insertReversePrimer" => "example_insertReversePrimer", # String
  "insertTemplate" => "example_insertTemplate", # String
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
| `enzymes` | `Array` |  |
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
| `rows` | `Array` |  |
| `tool` | `String` |  |

#### Example: Create

```ruby
volcano_plot_data = client.VolcanoPlotData.create({
  "ok" => "example_ok", # Object
  "provenance" => {}, # Hash
  "result" => {}, # Hash
  "rows" => [], # Array
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
| `max_results` | `Float` |  |
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
