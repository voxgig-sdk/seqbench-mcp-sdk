# SeqbenchMcp TypeScript SDK



The TypeScript SDK for the SeqbenchMcp API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.AlphafoldLookup()` — each with a small set of operations (`load`, `create`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/seqbench-mcp-sdk/releases](https://github.com/voxgig-sdk/seqbench-mcp-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { SeqbenchMcpSDK } from '@voxgig-sdk/seqbench-mcp'

const client = new SeqbenchMcpSDK({
  apikey: process.env.SEQBENCH_MCP_APIKEY,
})
```

### 4. Create, update, and remove

```ts
// Create — returns the created AlphafoldLookup
const created = await client.AlphafoldLookup().create({
  accession: 'example_accession',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})

```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const batch = await client.Batch().load()
  console.log(batch)
} catch (err) {
  console.error('load failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = SeqbenchMcpSDK.test()

const batch = await client.Batch().load()
// batch is a bare entity populated with mock response data
console.log(batch)
```

You can also use the instance method:

```ts
const client = new SeqbenchMcpSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Batch()

// First call runs the operation and stores its result
await entity.load()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new SeqbenchMcpSDK({
  apikey: '...',
  extend: [logger],
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
cd ts && npm test
```


## Reference

### SeqbenchMcpSDK

#### Constructor

```ts
new SeqbenchMcpSDK(options?: {
  apikey?: string
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `AlphafoldLookup(data?)` | `AlphafoldLookupEntity` | Create an AlphafoldLookup entity instance. |
| `AsoDesign(data?)` | `AsoDesignEntity` | Create an AsoDesign entity instance. |
| `BaseEditingDesign(data?)` | `BaseEditingDesignEntity` | Create a BaseEditingDesign entity instance. |
| `Batch(data?)` | `BatchEntity` | Create a Batch entity instance. |
| `BatchWorkflow(data?)` | `BatchWorkflowEntity` | Create a BatchWorkflow entity instance. |
| `CharacterizeSequence(data?)` | `CharacterizeSequenceEntity` | Create a CharacterizeSequence entity instance. |
| `CloningSimulate(data?)` | `CloningSimulateEntity` | Create a CloningSimulate entity instance. |
| `CodonAdaptationIndex(data?)` | `CodonAdaptationIndexEntity` | Create a CodonAdaptationIndex entity instance. |
| `CodonOptimize(data?)` | `CodonOptimizeEntity` | Create a CodonOptimize entity instance. |
| `ConstructAutofix(data?)` | `ConstructAutofixEntity` | Create a ConstructAutofix entity instance. |
| `ConstructQc(data?)` | `ConstructQcEntity` | Create a ConstructQc entity instance. |
| `CrisprGrnaDesign(data?)` | `CrisprGrnaDesignEntity` | Create a CrisprGrnaDesign entity instance. |
| `CrisprHdrDonor(data?)` | `CrisprHdrDonorEntity` | Create a CrisprHdrDonor entity instance. |
| `CrisprOfftargetCheck(data?)` | `CrisprOfftargetCheckEntity` | Create a CrisprOfftargetCheck entity instance. |
| `CrossDimer(data?)` | `CrossDimerEntity` | Create a CrossDimer entity instance. |
| `DnaMolarity(data?)` | `DnaMolarityEntity` | Create a DnaMolarity entity instance. |
| `DoubleDigest(data?)` | `DoubleDigestEntity` | Create a DoubleDigest entity instance. |
| `ExportEchoPicklist(data?)` | `ExportEchoPicklistEntity` | Create an ExportEchoPicklist entity instance. |
| `ExportOpentronsProtocol(data?)` | `ExportOpentronsProtocolEntity` | Create an ExportOpentronsProtocol entity instance. |
| `ExportPlateLayout(data?)` | `ExportPlateLayoutEntity` | Create an ExportPlateLayout entity instance. |
| `ExpressionHeatmapCluster(data?)` | `ExpressionHeatmapClusterEntity` | Create an ExpressionHeatmapCluster entity instance. |
| `FastqQcReport(data?)` | `FastqQcReportEntity` | Create a FastqQcReport entity instance. |
| `FastqTrim(data?)` | `FastqTrimEntity` | Create a FastqTrim entity instance. |
| `FindOrf(data?)` | `FindOrfEntity` | Create a FindOrf entity instance. |
| `FormatSequence(data?)` | `FormatSequenceEntity` | Create a FormatSequence entity instance. |
| `FunctionalEnrichment(data?)` | `FunctionalEnrichmentEntity` | Create a FunctionalEnrichment entity instance. |
| `GcContent(data?)` | `GcContentEntity` | Create a GcContent entity instance. |
| `GeneDossier(data?)` | `GeneDossierEntity` | Create a GeneDossier entity instance. |
| `GeneExpression(data?)` | `GeneExpressionEntity` | Create a GeneExpression entity instance. |
| `GeneModel(data?)` | `GeneModelEntity` | Create a GeneModel entity instance. |
| `GoldenGateFidelity(data?)` | `GoldenGateFidelityEntity` | Create a GoldenGateFidelity entity instance. |
| `HgvsConvert(data?)` | `HgvsConvertEntity` | Create a HgvsConvert entity instance. |
| `IdMapPoll(data?)` | `IdMapPollEntity` | Create an IdMapPoll entity instance. |
| `IdMapSubmit(data?)` | `IdMapSubmitEntity` | Create an IdMapSubmit entity instance. |
| `InSilicoPcr(data?)` | `InSilicoPcrEntity` | Create an InSilicoPcr entity instance. |
| `KaspPrimerDesign(data?)` | `KaspPrimerDesignEntity` | Create a KaspPrimerDesign entity instance. |
| `ListTool(data?)` | `ListToolEntity` | Create a ListTool entity instance. |
| `MeltingTemperature(data?)` | `MeltingTemperatureEntity` | Create a MeltingTemperature entity instance. |
| `MotifFinder(data?)` | `MotifFinderEntity` | Create a MotifFinder entity instance. |
| `MultipleSequenceAlignment(data?)` | `MultipleSequenceAlignmentEntity` | Create a MultipleSequenceAlignment entity instance. |
| `OligoAnalysi(data?)` | `OligoAnalysiEntity` | Create an OligoAnalysi entity instance. |
| `OrthologMap(data?)` | `OrthologMapEntity` | Create an OrthologMap entity instance. |
| `PairwiseAlignment(data?)` | `PairwiseAlignmentEntity` | Create a PairwiseAlignment entity instance. |
| `ParseGenbank(data?)` | `ParseGenbankEntity` | Create a ParseGenbank entity instance. |
| `ParseSangerTrace(data?)` | `ParseSangerTraceEntity` | Create a ParseSangerTrace entity instance. |
| `PlasmidAnnotate(data?)` | `PlasmidAnnotateEntity` | Create a PlasmidAnnotate entity instance. |
| `PlasmidDeepAnnotate(data?)` | `PlasmidDeepAnnotateEntity` | Create a PlasmidDeepAnnotate entity instance. |
| `PlasmidFullReport(data?)` | `PlasmidFullReportEntity` | Create a PlasmidFullReport entity instance. |
| `PlasmidIdentify(data?)` | `PlasmidIdentifyEntity` | Create a PlasmidIdentify entity instance. |
| `PrimeEditingDesign(data?)` | `PrimeEditingDesignEntity` | Create a PrimeEditingDesign entity instance. |
| `PrimeEditingTwinDesign(data?)` | `PrimeEditingTwinDesignEntity` | Create a PrimeEditingTwinDesign entity instance. |
| `PrimerDesign(data?)` | `PrimerDesignEntity` | Create a PrimerDesign entity instance. |
| `PrimerSpecificity(data?)` | `PrimerSpecificityEntity` | Create a PrimerSpecificity entity instance. |
| `ProteaseDigestion(data?)` | `ProteaseDigestionEntity` | Create a ProteaseDigestion entity instance. |
| `ProteinAnnotatePoll(data?)` | `ProteinAnnotatePollEntity` | Create a ProteinAnnotatePoll entity instance. |
| `ProteinAnnotateSubmit(data?)` | `ProteinAnnotateSubmitEntity` | Create a ProteinAnnotateSubmit entity instance. |
| `ProteinHydrophobicity(data?)` | `ProteinHydrophobicityEntity` | Create a ProteinHydrophobicity entity instance. |
| `ProteinProperty(data?)` | `ProteinPropertyEntity` | Create a ProteinProperty entity instance. |
| `RandomSequence(data?)` | `RandomSequenceEntity` | Create a RandomSequence entity instance. |
| `RestrictionSite(data?)` | `RestrictionSiteEntity` | Create a RestrictionSite entity instance. |
| `ReverseComplement(data?)` | `ReverseComplementEntity` | Create a ReverseComplement entity instance. |
| `ReverseTranslate(data?)` | `ReverseTranslateEntity` | Create a ReverseTranslate entity instance. |
| `RnaFold(data?)` | `RnaFoldEntity` | Create a RnaFold entity instance. |
| `SangerVsReference(data?)` | `SangerVsReferenceEntity` | Create a SangerVsReference entity instance. |
| `SavePermalink(data?)` | `SavePermalinkEntity` | Create a SavePermalink entity instance. |
| `SeqfileStat(data?)` | `SeqfileStatEntity` | Create a SeqfileStat entity instance. |
| `SequenceFetch(data?)` | `SequenceFetchEntity` | Create a SequenceFetch entity instance. |
| `SequenceFormatConvert(data?)` | `SequenceFormatConvertEntity` | Create a SequenceFormatConvert entity instance. |
| `SequenceReport(data?)` | `SequenceReportEntity` | Create a SequenceReport entity instance. |
| `SequenceSearch(data?)` | `SequenceSearchEntity` | Create a SequenceSearch entity instance. |
| `SequencingReadbackVerify(data?)` | `SequencingReadbackVerifyEntity` | Create a SequencingReadbackVerify entity instance. |
| `SessionCreate(data?)` | `SessionCreateEntity` | Create a SessionCreate entity instance. |
| `SessionGet(data?)` | `SessionGetEntity` | Create a SessionGet entity instance. |
| `SessionRun(data?)` | `SessionRunEntity` | Create a SessionRun entity instance. |
| `SessionSet(data?)` | `SessionSetEntity` | Create a SessionSet entity instance. |
| `SirnaDesign(data?)` | `SirnaDesignEntity` | Create a SirnaDesign entity instance. |
| `SiteDirectedMutagenesi(data?)` | `SiteDirectedMutagenesiEntity` | Create a SiteDirectedMutagenesi entity instance. |
| `Translate(data?)` | `TranslateEntity` | Create a Translate entity instance. |
| `VariantAnnotate(data?)` | `VariantAnnotateEntity` | Create a VariantAnnotate entity instance. |
| `VariantComparator(data?)` | `VariantComparatorEntity` | Create a VariantComparator entity instance. |
| `VerifyAssembly(data?)` | `VerifyAssemblyEntity` | Create a VerifyAssembly entity instance. |
| `VerifyConstruct(data?)` | `VerifyConstructEntity` | Create a VerifyConstruct entity instance. |
| `VirtualGel(data?)` | `VirtualGelEntity` | Create a VirtualGel entity instance. |
| `VolcanoPlotData(data?)` | `VolcanoPlotDataEntity` | Create a VolcanoPlotData entity instance. |
| `WebSearch(data?)` | `WebSearchEntity` | Create a WebSearch entity instance. |
| `tester(testopts?, sdkopts?)` | `SeqbenchMcpSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `SeqbenchMcpSDK.test(testopts?, sdkopts?)` | `SeqbenchMcpSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): SeqbenchMcpSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load` and `create` resolve to a single entity object.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

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

Operations: create.

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

Operations: create.

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

Operations: create.

API path: `/base_editing_design`

#### Batch

| Field | Description |
| --- | --- |
| `arg` |  |
| `input` |  |
| `ok` |  |
| `result` |  |
| `tool` |  |

Operations: create, load.

API path: `/batch`

#### BatchWorkflow

| Field | Description |
| --- | --- |
| `input` |  |
| `ok` |  |
| `result` |  |
| `step` |  |

Operations: create, load.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

API path: `/kasp_primer_design`

#### ListTool

| Field | Description |
| --- | --- |

Operations: load.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

API path: `/web_search`



## Entities


### AlphafoldLookup

Create an instance: `const alphafold_lookup = client.AlphafoldLookup()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const alphafold_lookup = await client.AlphafoldLookup().create({
  accession: 'example_accession',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### AsoDesign

Create an instance: `const aso_design = client.AsoDesign()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `target` | `string` |  |
| `tool` | `string` |  |
| `wing` | `number` |  |

#### Example: Create

```ts
const aso_design = await client.AsoDesign().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  target: 'example_target',
  tool: 'example_tool',
})
```


### BaseEditingDesign

Create an instance: `const base_editing_design = client.BaseEditingDesign()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `target` | `string` |  |
| `target_position` | `number` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const base_editing_design = await client.BaseEditingDesign().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  target: 'example_target',
  tool: 'example_tool',
})
```


### Batch

Create an instance: `const batch = client.Batch()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `Record<string, any>` |  |
| `input` | `string` |  |
| `ok` | `any` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Load

```ts
const batch = await client.Batch().load()
```

#### Example: Create

```ts
const batch = await client.Batch().create({
  input: 'example_input',
  ok: 'example_ok',
  result: {},
  tool: 'example_tool',
})
```


### BatchWorkflow

Create an instance: `const batch__workflow = client.BatchWorkflow()`

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
| `result` | `Record<string, any>` |  |
| `step` | `any[]` |  |

#### Example: Load

```ts
const batch__workflow = await client.BatchWorkflow().load()
```

#### Example: Create

```ts
const batch__workflow = await client.BatchWorkflow().create({
  input: 'example_input',
  ok: 'example_ok',
  result: {},
  step: [],
})
```


### CharacterizeSequence

Create an instance: `const characterize_sequence = client.CharacterizeSequence()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const characterize_sequence = await client.CharacterizeSequence().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### CloningSimulate

Create an instance: `const cloning_simulate = client.CloningSimulate()`

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
| `fragment` | `any[]` |  |
| `gate` | `any` |  |
| `insert` | `string` |  |
| `method` | `string` |  |
| `name` | `any[]` |  |
| `ok` | `any` |  |
| `overlap_len` | `number` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |
| `vector` | `string` |  |

#### Example: Create

```ts
const cloning_simulate = await client.CloningSimulate().create({
  method: 'example_method',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### CodonAdaptationIndex

Create an instance: `const codon_adaptation_index = client.CodonAdaptationIndex()`

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
| `provenance` | `Record<string, any>` |  |
| `rare_threshold` | `number` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const codon_adaptation_index = await client.CodonAdaptationIndex().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### CodonOptimize

Create an instance: `const codon_optimize = client.CodonOptimize()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const codon_optimize = await client.CodonOptimize().create({
  ok: 'example_ok',
  protein: 'example_protein',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### ConstructAutofix

Create an instance: `const construct_autofix = client.ConstructAutofix()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoid_enzyme` | `any[]` |  |
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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const construct_autofix = await client.ConstructAutofix().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### ConstructQc

Create an instance: `const construct_qc = client.ConstructQc()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoid_enzyme` | `any[]` |  |
| `cryptic_orf_min_aa` | `number` |  |
| `frame_start` | `number` |  |
| `gate` | `any` |  |
| `gc_high` | `number` |  |
| `gc_low` | `number` |  |
| `gc_window` | `number` |  |
| `homopolymer_min` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const construct_qc = await client.ConstructQc().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### CrisprGrnaDesign

Create an instance: `const crispr_grna_design = client.CrisprGrnaDesign()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `search_reverse_strand` | `boolean` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const crispr_grna_design = await client.CrisprGrnaDesign().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### CrisprHdrDonor

Create an instance: `const crispr_hdr_donor = client.CrisprHdrDonor()`

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
| `provenance` | `Record<string, any>` |  |
| `replacement` | `string` |  |
| `result` | `Record<string, any>` |  |
| `target_sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const crispr_hdr_donor = await client.CrisprHdrDonor().create({
  ok: 'example_ok',
  provenance: {},
  replacement: 'example_replacement',
  result: {},
  target_sequence: 'example_target_sequence',
  tool: 'example_tool',
})
```


### CrisprOfftargetCheck

Create an instance: `const crispr_offtarget_check = client.CrisprOfftargetCheck()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const crispr_offtarget_check = await client.CrisprOfftargetCheck().create({
  ok: 'example_ok',
  protospacer: 'example_protospacer',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### CrossDimer

Create an instance: `const cross_dimer = client.CrossDimer()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence_a` | `string` |  |
| `sequence_b` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const cross_dimer = await client.CrossDimer().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence_a: 'example_sequence_a',
  sequence_b: 'example_sequence_b',
  tool: 'example_tool',
})
```


### DnaMolarity

Create an instance: `const dna_molarity = client.DnaMolarity()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |
| `volume_ul` | `number` |  |

#### Example: Create

```ts
const dna_molarity = await client.DnaMolarity().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### DoubleDigest

Create an instance: `const double_digest = client.DoubleDigest()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const double_digest = await client.DoubleDigest().create({
  enzyme_a: 'example_enzyme_a',
  enzyme_b: 'example_enzyme_b',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### ExportEchoPicklist

Create an instance: `const export_echo_picklist = client.ExportEchoPicklist()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `reaction` | `any[]` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const export_echo_picklist = await client.ExportEchoPicklist().create({
  ok: 'example_ok',
  provenance: {},
  reaction: [],
  result: {},
  tool: 'example_tool',
})
```


### ExportOpentronsProtocol

Create an instance: `const export_opentrons_protocol = client.ExportOpentronsProtocol()`

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
| `provenance` | `Record<string, any>` |  |
| `reaction` | `any[]` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const export_opentrons_protocol = await client.ExportOpentronsProtocol().create({
  ok: 'example_ok',
  provenance: {},
  reaction: [],
  result: {},
  tool: 'example_tool',
})
```


### ExportPlateLayout

Create an instance: `const export_plate_layout = client.ExportPlateLayout()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `reaction` | `any[]` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const export_plate_layout = await client.ExportPlateLayout().create({
  ok: 'example_ok',
  provenance: {},
  reaction: [],
  result: {},
  tool: 'example_tool',
})
```


### ExpressionHeatmapCluster

Create an instance: `const expression_heatmap_cluster = client.ExpressionHeatmapCluster()`

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
| `gene` | `any[]` |  |
| `linkage` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sample` | `any[]` |  |
| `tool` | `string` |  |
| `value` | `any[]` |  |
| `z_score_row` | `boolean` |  |

#### Example: Create

```ts
const expression_heatmap_cluster = await client.ExpressionHeatmapCluster().create({
  gene: [],
  ok: 'example_ok',
  provenance: {},
  result: {},
  sample: [],
  tool: 'example_tool',
  value: [],
})
```


### FastqQcReport

Create an instance: `const fastq_qc_report = client.FastqQcReport()`

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
| `provenance` | `Record<string, any>` |  |
| `quality_offset` | `number` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const fastq_qc_report = await client.FastqQcReport().create({
  input: 'example_input',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### FastqTrim

Create an instance: `const fastq_trim = client.FastqTrim()`

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
| `provenance` | `Record<string, any>` |  |
| `quality_offset` | `number` |  |
| `quality_threshold` | `number` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const fastq_trim = await client.FastqTrim().create({
  input: 'example_input',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### FindOrf

Create an instance: `const find_orf = client.FindOrf()`

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
| `provenance` | `Record<string, any>` |  |
| `require_stop` | `boolean` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const find_orf = await client.FindOrf().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### FormatSequence

Create an instance: `const format_sequence = client.FormatSequence()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `reverse` | `boolean` |  |
| `sequence` | `string` |  |
| `strip_non_letter` | `boolean` |  |
| `tool` | `string` |  |
| `width` | `number` |  |

#### Example: Create

```ts
const format_sequence = await client.FormatSequence().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### FunctionalEnrichment

Create an instance: `const functional_enrichment = client.FunctionalEnrichment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `background` | `any[]` |  |
| `collection` | `any[]` |  |
| `gate` | `any` |  |
| `gene` | `any[]` |  |
| `max_term_size` | `number` |  |
| `min_term_size` | `number` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const functional_enrichment = await client.FunctionalEnrichment().create({
  gene: [],
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### GcContent

Create an instance: `const gc_content = client.GcContent()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const gc_content = await client.GcContent().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### GeneDossier

Create an instance: `const gene_dossier = client.GeneDossier()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const gene_dossier = await client.GeneDossier().create({
  gene: 'example_gene',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### GeneExpression

Create an instance: `const gene_expression = client.GeneExpression()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const gene_expression = await client.GeneExpression().create({
  gene: 'example_gene',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### GeneModel

Create an instance: `const gene_model = client.GeneModel()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const gene_model = await client.GeneModel().create({
  gene: 'example_gene',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### GoldenGateFidelity

Create an instance: `const golden_gate_fidelity = client.GoldenGateFidelity()`

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
| `overhang` | `any[]` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `risk_threshold` | `number` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const golden_gate_fidelity = await client.GoldenGateFidelity().create({
  ok: 'example_ok',
  overhang: [],
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### HgvsConvert

Create an instance: `const hgvs_convert = client.HgvsConvert()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |
| `variant` | `string` |  |

#### Example: Create

```ts
const hgvs_convert = await client.HgvsConvert().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
  variant: 'example_variant',
})
```


### IdMapPoll

Create an instance: `const id_map_poll = client.IdMapPoll()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const id_map_poll = await client.IdMapPoll().create({
  job_id: 'example_job_id',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### IdMapSubmit

Create an instance: `const id_map_submit = client.IdMapSubmit()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from` | `string` |  |
| `gate` | `any` |  |
| `ids` | `any[]` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tax_id` | `string` |  |
| `to` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const id_map_submit = await client.IdMapSubmit().create({
  from: 'example_from',
  ids: [],
  ok: 'example_ok',
  provenance: {},
  result: {},
  to: 'example_to',
  tool: 'example_tool',
})
```


### InSilicoPcr

Create an instance: `const in_silico_pcr = client.InSilicoPcr()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `reverse_primer` | `string` |  |
| `template` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const in_silico_pcr = await client.InSilicoPcr().create({
  forward_primer: 'example_forward_primer',
  ok: 'example_ok',
  provenance: {},
  result: {},
  reverse_primer: 'example_reverse_primer',
  template: 'example_template',
  tool: 'example_tool',
})
```


### KaspPrimerDesign

Create an instance: `const kasp_primer_design = client.KaspPrimerDesign()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `snp_position` | `number` |  |
| `target` | `string` |  |
| `target_core_tm` | `number` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const kasp_primer_design = await client.KaspPrimerDesign().create({
  allele_a: 'example_allele_a',
  allele_b: 'example_allele_b',
  ok: 'example_ok',
  provenance: {},
  result: {},
  snp_position: 1,
  target: 'example_target',
  tool: 'example_tool',
})
```


### ListTool

Create an instance: `const list_tool = client.ListTool()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ts
const list_tool = await client.ListTool().load()
```


### MeltingTemperature

Create an instance: `const melting_temperature = client.MeltingTemperature()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `target_tm` | `number` |  |
| `tm_tolerance` | `number` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const melting_temperature = await client.MeltingTemperature().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### MotifFinder

Create an instance: `const motif_finder = client.MotifFinder()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `search_reverse_strand` | `boolean` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const motif_finder = await client.MotifFinder().create({
  motif: 'example_motif',
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### MultipleSequenceAlignment

Create an instance: `const multiple_sequence_alignment = client.MultipleSequenceAlignment()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const multiple_sequence_alignment = await client.MultipleSequenceAlignment().create({
  input: 'example_input',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### OligoAnalysi

Create an instance: `const oligo_analysi = client.OligoAnalysi()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const oligo_analysi = await client.OligoAnalysi().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### OrthologMap

Create an instance: `const ortholog_map = client.OrthologMap()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `source_species` | `string` |  |
| `symbol` | `any[]` |  |
| `target_species` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |

#### Example: Create

```ts
const ortholog_map = await client.OrthologMap().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  symbol: [],
  target_species: 'example_target_species',
  tool: 'example_tool',
})
```


### PairwiseAlignment

Create an instance: `const pairwise_alignment = client.PairwiseAlignment()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `seq_a` | `string` |  |
| `seq_b` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const pairwise_alignment = await client.PairwiseAlignment().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  seq_a: 'example_seq_a',
  seq_b: 'example_seq_b',
  tool: 'example_tool',
})
```


### ParseGenbank

Create an instance: `const parse_genbank = client.ParseGenbank()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `text` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const parse_genbank = await client.ParseGenbank().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  text: 'example_text',
  tool: 'example_tool',
})
```


### ParseSangerTrace

Create an instance: `const parse_sanger_trace = client.ParseSangerTrace()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const parse_sanger_trace = await client.ParseSangerTrace().create({
  file_base64: 'example_file_base64',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### PlasmidAnnotate

Create an instance: `const plasmid_annotate = client.PlasmidAnnotate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const plasmid_annotate = await client.PlasmidAnnotate().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### PlasmidDeepAnnotate

Create an instance: `const plasmid_deep_annotate = client.PlasmidDeepAnnotate()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const plasmid_deep_annotate = await client.PlasmidDeepAnnotate().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### PlasmidFullReport

Create an instance: `const plasmid_full_report = client.PlasmidFullReport()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `top_n` | `number` |  |

#### Example: Create

```ts
const plasmid_full_report = await client.PlasmidFullReport().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### PlasmidIdentify

Create an instance: `const plasmid_identify = client.PlasmidIdentify()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `top_n` | `number` |  |

#### Example: Create

```ts
const plasmid_identify = await client.PlasmidIdentify().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### PrimeEditingDesign

Create an instance: `const prime_editing_design = client.PrimeEditingDesign()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `rtt_homology` | `number` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const prime_editing_design = await client.PrimeEditingDesign().create({
  edit_end: 1,
  edit_start: 1,
  ok: 'example_ok',
  provenance: {},
  result: {},
  target: 'example_target',
  tool: 'example_tool',
})
```


### PrimeEditingTwinDesign

Create an instance: `const prime_editing_twin_design = client.PrimeEditingTwinDesign()`

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
| `provenance` | `Record<string, any>` |  |
| `replace_end` | `number` |  |
| `replace_start` | `number` |  |
| `result` | `Record<string, any>` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const prime_editing_twin_design = await client.PrimeEditingTwinDesign().create({
  new_sequence: 'example_new_sequence',
  ok: 'example_ok',
  provenance: {},
  replace_end: 1,
  replace_start: 1,
  result: {},
  target: 'example_target',
  tool: 'example_tool',
})
```


### PrimerDesign

Create an instance: `const primer_design = client.PrimerDesign()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `target_end` | `number` |  |
| `target_start` | `number` |  |
| `template` | `string` |  |
| `tm_max` | `number` |  |
| `tm_max_diff` | `number` |  |
| `tm_min` | `number` |  |
| `tm_opt` | `number` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const primer_design = await client.PrimerDesign().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  template: 'example_template',
  tool: 'example_tool',
})
```


### PrimerSpecificity

Create an instance: `const primer_specificity = client.PrimerSpecificity()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `reverse_primer` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const primer_specificity = await client.PrimerSpecificity().create({
  forward_primer: 'example_forward_primer',
  ok: 'example_ok',
  provenance: {},
  result: {},
  reverse_primer: 'example_reverse_primer',
  tool: 'example_tool',
})
```


### ProteaseDigestion

Create an instance: `const protease_digestion = client.ProteaseDigestion()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const protease_digestion = await client.ProteaseDigestion().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### ProteinAnnotatePoll

Create an instance: `const protein_annotate_poll = client.ProteinAnnotatePoll()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const protein_annotate_poll = await client.ProteinAnnotatePoll().create({
  job_id: 'example_job_id',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### ProteinAnnotateSubmit

Create an instance: `const protein_annotate_submit = client.ProteinAnnotateSubmit()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const protein_annotate_submit = await client.ProteinAnnotateSubmit().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### ProteinHydrophobicity

Create an instance: `const protein_hydrophobicity = client.ProteinHydrophobicity()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `scale` | `string` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `window` | `number` |  |

#### Example: Create

```ts
const protein_hydrophobicity = await client.ProteinHydrophobicity().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### ProteinProperty

Create an instance: `const protein_property = client.ProteinProperty()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const protein_property = await client.ProteinProperty().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### RandomSequence

Create an instance: `const random_sequence = client.RandomSequence()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const random_sequence = await client.RandomSequence().create({
  length: 1,
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### RestrictionSite

Create an instance: `const restriction_site = client.RestrictionSite()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enzyme` | `any[]` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const restriction_site = await client.RestrictionSite().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### ReverseComplement

Create an instance: `const reverse_complement = client.ReverseComplement()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |

#### Example: Create

```ts
const reverse_complement = await client.ReverseComplement().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### ReverseTranslate

Create an instance: `const reverse_translate = client.ReverseTranslate()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const reverse_translate = await client.ReverseTranslate().create({
  ok: 'example_ok',
  protein: 'example_protein',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### RnaFold

Create an instance: `const rna_fold = client.RnaFold()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const rna_fold = await client.RnaFold().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### SangerVsReference

Create an instance: `const sanger_vs_reference = client.SangerVsReference()`

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
| `provenance` | `Record<string, any>` |  |
| `read` | `string` |  |
| `reference` | `string` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const sanger_vs_reference = await client.SangerVsReference().create({
  ok: 'example_ok',
  provenance: {},
  reference: 'example_reference',
  result: {},
  tool: 'example_tool',
})
```


### SavePermalink

Create an instance: `const save_permalink = client.SavePermalink()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `Record<string, any>` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const save_permalink = await client.SavePermalink().create({
  arg: {},
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### SeqfileStat

Create an instance: `const seqfile_stat = client.SeqfileStat()`

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
| `provenance` | `Record<string, any>` |  |
| `quality_offset` | `number` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const seqfile_stat = await client.SeqfileStat().create({
  input: 'example_input',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### SequenceFetch

Create an instance: `const sequence_fetch = client.SequenceFetch()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const sequence_fetch = await client.SequenceFetch().create({
  accession: 'example_accession',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### SequenceFormatConvert

Create an instance: `const sequence_format_convert = client.SequenceFormatConvert()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `to` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const sequence_format_convert = await client.SequenceFormatConvert().create({
  input: 'example_input',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### SequenceReport

Create an instance: `const sequence_report = client.SequenceReport()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const sequence_report = await client.SequenceReport().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### SequenceSearch

Create an instance: `const sequence_search = client.SequenceSearch()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `term` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const sequence_search = await client.SequenceSearch().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### SequencingReadbackVerify

Create an instance: `const sequencing_readback_verify = client.SequencingReadbackVerify()`

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
| `provenance` | `Record<string, any>` |  |
| `read` | `string` |  |
| `reference` | `string` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const sequencing_readback_verify = await client.SequencingReadbackVerify().create({
  ok: 'example_ok',
  provenance: {},
  read: 'example_read',
  reference: 'example_reference',
  result: {},
  tool: 'example_tool',
})
```


### SessionCreate

Create an instance: `const session_create = client.SessionCreate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entry` | `Record<string, any>` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const session_create = await client.SessionCreate().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### SessionGet

Create an instance: `const session_get = client.SessionGet()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `name` | `any[]` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const session_get = await client.SessionGet().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  session_id: 'example_session_id',
  tool: 'example_tool',
})
```


### SessionRun

Create an instance: `const session_run = client.SessionRun()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `Record<string, any>` |  |
| `from_session` | `Record<string, any>` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |
| `write_back` | `Record<string, any>` |  |

#### Example: Create

```ts
const session_run = await client.SessionRun().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  session_id: 'example_session_id',
  tool: 'example_tool',
})
```


### SessionSet

Create an instance: `const session_set = client.SessionSet()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entry` | `Record<string, any>` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const session_set = await client.SessionSet().create({
  entry: {},
  ok: 'example_ok',
  provenance: {},
  result: {},
  session_id: 'example_session_id',
  tool: 'example_tool',
})
```


### SirnaDesign

Create an instance: `const sirna_design = client.SirnaDesign()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sh_rna_loop` | `string` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const sirna_design = await client.SirnaDesign().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  target: 'example_target',
  tool: 'example_tool',
})
```


### SiteDirectedMutagenesi

Create an instance: `const site_directed_mutagenesi = client.SiteDirectedMutagenesi()`

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
| `provenance` | `Record<string, any>` |  |
| `residue` | `number` |  |
| `result` | `Record<string, any>` |  |
| `style` | `string` |  |
| `target_aa` | `string` |  |
| `template` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const site_directed_mutagenesi = await client.SiteDirectedMutagenesi().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  template: 'example_template',
  tool: 'example_tool',
})
```


### Translate

Create an instance: `const translate = client.Translate()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `to_stop` | `boolean` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const translate = await client.Translate().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### VariantAnnotate

Create an instance: `const variant_annotate = client.VariantAnnotate()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |
| `variant` | `string` |  |

#### Example: Create

```ts
const variant_annotate = await client.VariantAnnotate().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
  variant: 'example_variant',
})
```


### VariantComparator

Create an instance: `const variant_comparator = client.VariantComparator()`

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
| `provenance` | `Record<string, any>` |  |
| `query` | `string` |  |
| `reference` | `string` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const variant_comparator = await client.VariantComparator().create({
  ok: 'example_ok',
  provenance: {},
  query: 'example_query',
  reference: 'example_reference',
  result: {},
  tool: 'example_tool',
})
```


### VerifyAssembly

Create an instance: `const verify_assembly = client.VerifyAssembly()`

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
| `fragment` | `any[]` |  |
| `fragment_pcr` | `any[]` |  |
| `frame_start` | `number` |  |
| `gate` | `any` |  |
| `insert` | `string` |  |
| `insert_pcr` | `Record<string, any>` |  |
| `method` | `string` |  |
| `name` | `any[]` |  |
| `ok` | `any` |  |
| `overlap_len` | `number` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |
| `vector` | `string` |  |
| `vector_pcr` | `Record<string, any>` |  |

#### Example: Create

```ts
const verify_assembly = await client.VerifyAssembly().create({
  claimed_construct: 'example_claimed_construct',
  method: 'example_method',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### VerifyConstruct

Create an instance: `const verify_construct = client.VerifyConstruct()`

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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `template_circular` | `boolean` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const verify_construct = await client.VerifyConstruct().create({
  claimed_construct: 'example_claimed_construct',
  insert_forward_primer: 'example_insert_forward_primer',
  insert_reverse_primer: 'example_insert_reverse_primer',
  insert_template: 'example_insert_template',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```


### VirtualGel

Create an instance: `const virtual_gel = client.VirtualGel()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `boolean` |  |
| `enzyme` | `any[]` |  |
| `gate` | `any` |  |
| `ladder` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const virtual_gel = await client.VirtualGel().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```


### VolcanoPlotData

Create an instance: `const volcano_plot_data = client.VolcanoPlotData()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` |  |
| `row` | `any[]` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const volcano_plot_data = await client.VolcanoPlotData().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  row: [],
  tool: 'example_tool',
})
```


### WebSearch

Create an instance: `const web_search = client.WebSearch()`

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
| `provenance` | `Record<string, any>` |  |
| `query` | `string` |  |
| `result` | `Record<string, any>` |  |
| `tool` | `string` |  |

#### Example: Create

```ts
const web_search = await client.WebSearch().create({
  ok: 'example_ok',
  provenance: {},
  query: 'example_query',
  result: {},
  tool: 'example_tool',
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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
seqbench-mcp/
├── src/
│   ├── SeqbenchMcpSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { SeqbenchMcpSDK } from '@voxgig-sdk/seqbench-mcp'
```

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const batch = client.Batch()
await batch.load()

// batch.data() now returns the batch data from the last `load`
// batch.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
