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

Create an instance: `alphafold_lookup = client.AlphafoldLookup`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `String` | UniProt accession, e.g. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `Integer` | Total gapmer length (nt). |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `target` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |
| `wing` | `Integer` | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

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
| `editor` | `String` | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | `Integer` | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `target` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | `Integer` | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | `String` | The tool slug that ran. |

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
| `args` | `Hash` | Shared tool arguments applied to every record. |
| `capped` | `Boolean` | True if input exceeded the record limit. |
| `columns` | `Array` |  |
| `count` | `Integer` |  |
| `errors` | `Integer` |  |
| `input` | `String` | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | `Integer` | Maximum records per call (500). |
| `provenance` | `Hash` |  |
| `rows` | `Array` |  |
| `tool` | `String` | A batchable tool slug (see `GET /batch`). |

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
| `columns` | `Array` | Flattened "<step>·<tool>·<key>" column headers. |
| `count` | `Integer` |  |
| `errors` | `Integer` |  |
| `input` | `String` | Multi-FASTA text or one sequence per line. |
| `limit` | `Integer` | Maximum records per call (200). |
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
| `endPrimerLength` | `Integer` | Length of the naive end primers taken from each end. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `Integer` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `Integer` | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `armTmTarget` | `Float` | Target annealing Tm (°C) for primer arms. |
| `circular` | `Boolean` | Produce a circular product. |
| `enzyme` | `String` | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | `String` | 3′ enzyme (restriction method). |
| `enzyme5` | `String` | 5′ enzyme (restriction method). |
| `fragments` | `Array` | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `String` | Insert sequence (restriction method). |
| `method` | `String` | Assembly method. |
| `names` | `Array` | Optional labels for each fragment. |
| `ok` | `Object` |  |
| `overlapLen` | `Integer` | Gibson homology-arm length (bp). |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |
| `vector` | `String` | Vector sequence (restriction method). |

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
| `frameStart` | `Integer` | 1-based position to start reading codons. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `organism` | `String` |  |
| `provenance` | `Hash` |  |
| `rareThreshold` | `Float` | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `organism` | `String` |  |
| `protein` | `String` | Protein sequence (one-letter codes). |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `avoidEnzymes` | `Array` | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | `Integer` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `Integer` | 1-based nucleotide where the reading frame begins. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `Float` |  |
| `gcLow` | `Float` |  |
| `gcWindow` | `Integer` |  |
| `homopolymerMin` | `Integer` |  |
| `maxPasses` | `Integer` | Repeat full passes until clean or no further progress. |
| `ok` | `Object` |  |
| `organism` | `String` | Codon-usage table to prefer among synonymous options. |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `avoidEnzymes` | `Array` | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | `Integer` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `Integer` | 1-based nucleotide where the reading frame begins. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `Float` | GC% above this flags a GC-rich window. |
| `gcLow` | `Float` | GC% below this flags an AT-rich window. |
| `gcWindow` | `Integer` | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | `Integer` | Minimum run length to flag a homopolymer. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | `Float` | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | `String` | Nuclease id. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `searchReverseStrand` | `Boolean` | Also scan the reverse strand for guides. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `armLength` | `Integer` | Homology arm length (bp) on each side. |
| `blockPam` | `Boolean` | When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut. |
| `designGenotypingPrimers` | `Boolean` | Also design a primer pair (on the original targetSequence) whose product spans the edit site. |
| `editEnd` | `Integer` | 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. |
| `editStart` | `Integer` | 1-based start of the region being replaced. |
| `frameStart` | `Integer` | Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | `Integer` | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | `Integer` | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | `String` | Strand the guide's protospacer is on. |
| `nuclease` | `String` | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `replacement` | `String` | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | `Hash` | Tool-specific output object. |
| `targetSequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `Integer` | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | `String` | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` | `Object` |  |
| `protospacer` | `String` | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequenceA` | `String` | First oligo (5'→3'). |
| `sequenceB` | `String` | Second oligo (5'→3'). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `Integer` | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | `Float` | Mass in nanograms. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | `String` | The tool slug that ran. |
| `type` | `String` | Molecule type. |
| `volumeUl` | `Float` | Volume in microlitres (0 = unknown; needed for concentration). |

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
| `enzymeA` | `String` | First enzyme name (e.g. |
| `enzymeB` | `String` | Second enzyme name (e.g. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `reactions` | `Array` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `protocolName` | `String` | Optional protocol name (used in the script's metadata). |
| `provenance` | `Hash` |  |
| `reactions` | `Array` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `reactions` | `Array` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `clusterCols` | `Boolean` | Cluster (reorder) samples. |
| `clusterRows` | `Boolean` | Cluster (reorder) genes. |
| `distanceMetric` | `String` | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `Array` | Row (gene) labels. |
| `linkage` | `String` | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `samples` | `Array` | Column (sample) labels. |
| `tool` | `String` | The tool slug that ran. |
| `values` | `Array` | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | `Boolean` | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `String` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `qualityOffset` | `Integer` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `String` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | `Integer` | Reads shorter than this after trimming are dropped. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `qualityOffset` | `Integer` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | `Integer` | 3' quality-trim threshold (Phred score). |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | `Integer` | Minimum protein length (aa) to report. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `requireStop` | `Boolean` | Only report ORFs terminated by a stop codon. |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `convert` | `String` | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `reverse` | `Boolean` | Reverse the sequence (no complement). |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | `Boolean` | Remove digits, spaces and gaps (keep letters only). |
| `tool` | `String` | The tool slug that ran. |
| `width` | `Integer` | Line-wrap width; 0 = single line. |

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
| `background` | `Array` | Custom background/universe gene symbols. |
| `collections` | `Array` | Which term collections to test. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `Array` | Query gene symbols (human, e.g. |
| `maxTermSize` | `Integer` | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | `Integer` | Skip terms/pathways with fewer than this many background genes. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `String` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `String` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `String` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `compareToNamedSet` | `String` | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | `String` | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `overhangs` | `Array` | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `riskThreshold` | `Float` | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |
| `variant` | `String` | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `from` | `String` | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | `Array` | The ids to map, up to 1000 (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `taxId` | `String` | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | `String` | Target id type. |
| `tool` | `String` | The tool slug that ran. |

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
| `circular` | `Boolean` | Treat the template as circular (plasmid). |
| `forwardPrimer` | `String` | Primer 1, 5'→3'. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `Integer` | Mismatches tolerated per primer. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `reversePrimer` | `String` | Primer 2, 5'→3' (order does not matter). |
| `template` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `addSecondaryMismatch` | `Boolean` | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | `String` | First allele (single base) — gets the FAM tail. |
| `alleleB` | `String` | Second allele (single base) — gets the HEX tail. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | `Integer` | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | `Integer` | Minimum amplicon length for the common reverse primer. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `snpPosition` | `Integer` | 1-based position of the SNP on the forward strand. |
| `target` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | `Float` | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | `String` | The tool slug that ran. |

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
| `dntpMM` | `Float` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `Float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `Float` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Object` |  |
| `oligoNM` | `Float` | Total strand concentration (nM). |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | `Float` | Optional target Tm (°C). |
| `tmTolerance` | `Float` | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `Integer` | Maximum allowed mismatches per match. |
| `motif` | `String` | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `searchReverseStrand` | `Boolean` | Also search the reverse strand. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `String` | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `dntpMM` | `Float` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `Float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `Float` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Object` |  |
| `oligoNM` | `Float` | Total strand concentration (nM). |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sourceSpecies` | `String` | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | `Array` | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | `String` | Ensembl species slug to find homologs in (e.g. |
| `tool` | `String` | The tool slug that ran. |
| `type` | `String` | Homology type to return. |

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
| `gap` | `Float` | Linear gap penalty (per gap position). |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | `Float` | Match score. |
| `mismatch` | `Float` | Mismatch penalty. |
| `mode` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `seqA` | `String` | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | `String` | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `text` | `String` | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | `String` | The tool slug that ran. |

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
| `fileBase64` | `String` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `String` | Optional original file name (echoed back). |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `circular` | `Boolean` | Treat the sequence as a circular plasmid (vs. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `circular` | `Boolean` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |
| `topN` | `Integer` | How many top-ranked backbone candidates to report. |

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
| `circular` | `Boolean` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |
| `topN` | `Integer` | How many top-ranked backbone candidates to report. |

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
| `editEnd` | `Integer` | 1-based inclusive end of the region being changed. |
| `editStart` | `Integer` | 1-based inclusive start of the region being changed. |
| `frameStart` | `Integer` | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | `String` | Replacement bases (forward strand). |
| `ok` | `Object` |  |
| `pbsLength` | `Integer` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `rttHomology` | `Integer` | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | `String` | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | `String` | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` | `Object` |  |
| `overlapLength` | `Integer` | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | `Integer` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `Hash` |  |
| `replaceEnd` | `Integer` | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | `Integer` | 1-based inclusive start of the region being replaced/deleted. |
| `result` | `Hash` | Tool-specific output object. |
| `target` | `String` | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | `String` | The tool slug that ran. |

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
| `dntpMM` | `Float` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` | `Float` |  |
| `gcMin` | `Float` |  |
| `lenMax` | `Integer` |  |
| `lenMin` | `Integer` |  |
| `lenOpt` | `Integer` |  |
| `maxReturn` | `Integer` | Number of best pairs to return. |
| `mgMM` | `Float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `Float` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `Object` |  |
| `oligoNM` | `Float` | Total strand concentration (nM). |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `targetEnd` | `Integer` | 1-based inclusive end of the target region (optional). |
| `targetStart` | `Integer` | 1-based inclusive start of a region the product must span (optional). |
| `template` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` | `Float` |  |
| `tmMaxDiff` | `Float` | Max Tm difference within a pair (°C). |
| `tmMin` | `Float` |  |
| `tmOpt` | `Float` |  |
| `tool` | `String` | The tool slug that ran. |

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
| `forwardPrimer` | `String` | Forward primer, 5'→3'. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `Integer` | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | `Integer` | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `reversePrimer` | `String` | Reverse primer, 5'→3'. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | `Float` | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | `Integer` | Cap on the number of returned peptides. |
| `minMass` | `Float` | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | `Integer` | Allowed missed internal cleavages (0–2). |
| `ok` | `Object` |  |
| `protease` | `String` | Protease or chemical cleavage agent. |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `String` |  |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `appl` | `String` | Restrict to one member database (e.g. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | `Boolean` | Include GO-term cross-references. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `scale` | `String` | Amino-acid scale. |
| `sequence` | `String` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `String` | The tool slug that ran. |
| `window` | `Integer` | Sliding-window size (clamped to an odd number ≥ 1). |

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
| `chargeStep` | `Float` | pH step for the net-charge titration curve (0–14). |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | `Float` | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` | `String` |  |
| `length` | `Integer` | Number of residues to generate. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `enzymes` | `Array` | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |
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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` | `String` |  |
| `ok` | `Object` |  |
| `organism` | `String` | Codon-usage host (ignored in degenerate mode). |
| `protein` | `String` | Protein sequence (one-letter codes; * for stop). |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `fileBase64` | `String` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `String` | Optional original file name (echoed back). |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | `Float` | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `read` | `String` | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | `String` | Expected reference sequence (FASTA or raw). |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `args` | `Hash` | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `String` | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `qualityOffset` | `Integer` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `accession` | `String` | GenBank/RefSeq accession (e.g. |
| `db` | `String` | Database to query; auto-detects from the accession format. |
| `format` | `String` | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `from` | `String` | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `String` | A FASTA or GenBank record to convert. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `to` | `String` | Output format. |
| `tool` | `String` | The tool slug that ran. |

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
| `endPrimerLength` | `Integer` | Length of the naive end primers taken from each end. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `Integer` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `Integer` | Minimum ORF length in amino acids. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `String` | Gene symbol/name, e.g. |
| `maxResults` | `Integer` | Up to 20. |
| `ok` | `Object` |  |
| `organism` | `String` | Organism name, e.g. |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `term` | `String` | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | `Integer` | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `reads` | `String` | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | `String` | The claimed/expected reference sequence. |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `entries` | `Hash` | Initial named entries, e.g. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | `Array` | Only return these entries; omit to return all of them. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sessionId` | `String` |  |
| `tool` | `String` | The tool slug that ran. |

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
| `args` | `Hash` | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | `Hash` | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sessionId` | `String` |  |
| `tool` | `String` | The tool slug that ran. |
| `writeBack` | `Hash` | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

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
| `entries` | `Hash` | Named entries to add/overwrite, e.g. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sessionId` | `String` |  |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | `Integer` | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `shRnaLoop` | `String` | Loop sequence used when assembling the shRNA cassette. |
| `target` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `armTmTarget` | `Float` | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | `Float` | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | `String` | Edit at the nucleotide or amino-acid level. |
| `frameStart` | `Integer` | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `Float` | Divalent cation [Mg2+] (mM). |
| `naMM` | `Float` | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | `String` | Replacement base (editKind='nt'). |
| `ok` | `Object` |  |
| `oligoNM` | `Float` | Total strand concentration (nM). |
| `organism` | `String` | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | `Integer` | 1-based position to substitute (editKind='nt'). |
| `provenance` | `Hash` |  |
| `residue` | `Integer` | 1-based residue number to change (editKind='aa'). |
| `result` | `Hash` | Tool-specific output object. |
| `style` | `String` | Mutagenic primer style. |
| `targetAa` | `String` | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | `Boolean` | Stop at the first stop codon. |
| `tool` | `String` | The tool slug that ran. |

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
| `assembly` | `String` | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |
| `variant` | `String` | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

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
| `coding` | `Boolean` | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | `Integer` | 1-based reading-frame start (used when coding is true). |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `query` | `String` | Query / variant sequence (raw or FASTA). |
| `reference` | `String` | Reference / wild-type sequence (raw or FASTA). |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
| `armTmTarget` | `Float` | Target annealing Tm (°C) for primer arms. |
| `circular` | `Boolean` | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | `String` | The sequence you claim you ended up with. |
| `coding` | `Boolean` | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | `String` | Type IIS enzyme for Golden Gate. |
| `enzyme3` | `String` | 3′ enzyme (restriction method). |
| `enzyme5` | `String` | 5′ enzyme (restriction method). |
| `fragmentPcrs` | `Array` | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | `Array` | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | `Integer` | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `String` | Insert sequence (restriction method). |
| `insertPcr` | `Hash` | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | `String` | Assembly method used. |
| `names` | `Array` | Optional labels for each fragment. |
| `ok` | `Object` |  |
| `overlapLen` | `Integer` | Gibson homology-arm length (bp). |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |
| `vector` | `String` | Vector sequence (restriction method). |
| `vectorPcr` | `Hash` | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

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
| `claimedConstruct` | `String` | The final sequence claimed to have been built. |
| `expectedFrameStart` | `Integer` | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | `String` | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | `String` | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | `String` | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | `Integer` | Mismatches tolerated per primer during PCR prediction. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `templateCircular` | `Boolean` | Treat insertTemplate as circular (e.g. |
| `tool` | `String` | The tool slug that ran. |

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
| `circular` | `Boolean` | Treat the sequence as circular (plasmid). |
| `enzymes` | `Array` | Enzyme names to digest with. |
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | `String` | DNA ladder to plot alongside the sample lane. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `sequence` | `String` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `result` | `Hash` | Tool-specific output object. |
| `rows` | `Array` | Differential expression rows, one per gene. |
| `tool` | `String` | The tool slug that ran. |

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
| `gate` | `Object` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | `Float` | Maximum number of results to return (default 5, max 10). |
| `ok` | `Object` |  |
| `provenance` | `Hash` |  |
| `query` | `String` | The search query. |
| `result` | `Hash` | Tool-specific output object. |
| `tool` | `String` | The tool slug that ran. |

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
