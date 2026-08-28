# SeqbenchMcp TypeScript SDK



The TypeScript SDK for the SeqbenchMcp API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.AlphafoldLookup()` — each with a small set of operations (`load`, `create`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Also generated from this model: `go`, `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb` — see
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
// Create — returns the created AlphafoldLookup ENTITY (.data() for the record)
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
// batch is the entity, populated with mock response data
// — call batch.data() for the record itself
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
| `accession` | UniProt accession, e.g. |
| `gate` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` |  |
| `provenance` |  |
| `result` | Tool-specific output object. |
| `tool` | The tool slug that ran. |

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create, load.

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

Operations: create, load.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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

Operations: create.

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
| `accession` | `string` | UniProt accession, e.g. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `number` | Total gapmer length (nt). |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
| `wing` | `number` | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

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
| `editor` | `string` | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | `number` | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | `number` | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | `string` | The tool slug that ran. |

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
| `args` | `Record<string, any>` | Shared tool arguments applied to every record. |
| `capped` | `boolean` | True if input exceeded the record limit. |
| `columns` | `any[]` |  |
| `count` | `number` |  |
| `errors` | `number` |  |
| `input` | `string` | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | `number` | Maximum records per call (500). |
| `provenance` | `Record<string, any>` |  |
| `rows` | `any[]` |  |
| `tool` | `string` | A batchable tool slug (see `GET /batch`). |

#### Example: Load

```ts
const batch = await client.Batch().load()
```

#### Example: Create

```ts
const batch = await client.Batch().create({
  capped: true,
  columns: [],
  count: 1,
  errors: 1,
  input: 'example_input',
  limit: 1,
  provenance: {},
  rows: [],
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
| `capped` | `boolean` |  |
| `columns` | `any[]` | Flattened "<step>·<tool>·<key>" column headers. |
| `count` | `number` |  |
| `errors` | `number` |  |
| `input` | `string` | Multi-FASTA text or one sequence per line. |
| `limit` | `number` | Maximum records per call (200). |
| `provenance` | `Record<string, any>` |  |
| `rows` | `any[]` |  |
| `steps` | `any[]` |  |

#### Example: Load

```ts
const batch__workflow = await client.BatchWorkflow().load()
```

#### Example: Create

```ts
const batch__workflow = await client.BatchWorkflow().create({
  capped: true,
  columns: [],
  count: 1,
  errors: 1,
  input: 'example_input',
  limit: 1,
  provenance: {},
  rows: [],
  steps: [],
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
| `endPrimerLength` | `number` | Length of the naive end primers taken from each end. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `number` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `number` | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `armTmTarget` | `number` | Target annealing Tm (°C) for primer arms. |
| `circular` | `boolean` | Produce a circular product. |
| `enzyme` | `string` | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | `string` | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | 5′ enzyme (restriction method). |
| `fragments` | `any[]` | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | Insert sequence (restriction method). |
| `method` | `string` | Assembly method. |
| `names` | `any[]` | Optional labels for each fragment. |
| `ok` | `any` |  |
| `overlapLen` | `number` | Gibson homology-arm length (bp). |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `vector` | `string` | Vector sequence (restriction method). |

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
| `frameStart` | `number` | 1-based position to start reading codons. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `provenance` | `Record<string, any>` |  |
| `rareThreshold` | `number` | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `protein` | `string` | Protein sequence (one-letter codes). |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `avoidEnzymes` | `any[]` | Enzyme names whose internal sites should be removed (e.g. |
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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `avoidEnzymes` | `any[]` | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | `number` | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `number` | 1-based nucleotide where the reading frame begins. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `number` | GC% above this flags a GC-rich window. |
| `gcLow` | `number` | GC% below this flags an AT-rich window. |
| `gcWindow` | `number` | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | `number` | Minimum run length to flag a homopolymer. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | `number` | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | `string` | Nuclease id. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `searchReverseStrand` | `boolean` | Also scan the reverse strand for guides. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `provenance` | `Record<string, any>` |  |
| `replacement` | `string` | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `targetSequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const crispr_hdr_donor = await client.CrisprHdrDonor().create({
  ok: 'example_ok',
  provenance: {},
  replacement: 'example_replacement',
  result: {},
  targetSequence: 'example_targetSequence',
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | `string` | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` | `any` |  |
| `protospacer` | `string` | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequenceA` | `string` | First oligo (5'→3'). |
| `sequenceB` | `string` | Second oligo (5'→3'). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const cross_dimer = await client.CrossDimer().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequenceA: 'example_sequenceA',
  sequenceB: 'example_sequenceB',
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `number` | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | `number` | Mass in nanograms. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | `string` | The tool slug that ran. |
| `type` | `string` | Molecule type. |
| `volumeUl` | `number` | Volume in microlitres (0 = unknown; needed for concentration). |

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
| `enzymeA` | `string` | First enzyme name (e.g. |
| `enzymeB` | `string` | Second enzyme name (e.g. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const double_digest = await client.DoubleDigest().create({
  enzymeA: 'example_enzymeA',
  enzymeB: 'example_enzymeB',
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `reactions` | `any[]` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const export_echo_picklist = await client.ExportEchoPicklist().create({
  ok: 'example_ok',
  provenance: {},
  reactions: [],
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `protocolName` | `string` | Optional protocol name (used in the script's metadata). |
| `provenance` | `Record<string, any>` |  |
| `reactions` | `any[]` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const export_opentrons_protocol = await client.ExportOpentronsProtocol().create({
  ok: 'example_ok',
  provenance: {},
  reactions: [],
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `reactions` | `any[]` | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const export_plate_layout = await client.ExportPlateLayout().create({
  ok: 'example_ok',
  provenance: {},
  reactions: [],
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
| `clusterCols` | `boolean` | Cluster (reorder) samples. |
| `clusterRows` | `boolean` | Cluster (reorder) genes. |
| `distanceMetric` | `string` | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `any[]` | Row (gene) labels. |
| `linkage` | `string` | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `samples` | `any[]` | Column (sample) labels. |
| `tool` | `string` | The tool slug that ran. |
| `values` | `any[]` | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | `boolean` | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

#### Example: Create

```ts
const expression_heatmap_cluster = await client.ExpressionHeatmapCluster().create({
  genes: [],
  ok: 'example_ok',
  provenance: {},
  result: {},
  samples: [],
  tool: 'example_tool',
  values: [],
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `qualityOffset` | `number` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | `number` | Reads shorter than this after trimming are dropped. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `qualityOffset` | `number` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | `number` | 3' quality-trim threshold (Phred score). |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | `number` | Minimum protein length (aa) to report. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `requireStop` | `boolean` | Only report ORFs terminated by a stop codon. |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `caseMode` | `string` |  |
| `convert` | `string` | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `reverse` | `boolean` | Reverse the sequence (no complement). |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | `boolean` | Remove digits, spaces and gaps (keep letters only). |
| `tool` | `string` | The tool slug that ran. |
| `width` | `number` | Line-wrap width; 0 = single line. |

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
| `background` | `any[]` | Custom background/universe gene symbols. |
| `collections` | `any[]` | Which term collections to test. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `any[]` | Query gene symbols (human, e.g. |
| `maxTermSize` | `number` | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | `number` | Skip terms/pathways with fewer than this many background genes. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const functional_enrichment = await client.FunctionalEnrichment().create({
  genes: [],
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `compareToNamedSet` | `string` | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | `string` | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `overhangs` | `any[]` | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `riskThreshold` | `number` | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const golden_gate_fidelity = await client.GoldenGateFidelity().create({
  ok: 'example_ok',
  overhangs: [],
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `variant` | `string` | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const id_map_poll = await client.IdMapPoll().create({
  jobId: 'example_jobId',
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
| `from` | `string` | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | `any[]` | The ids to map, up to 1000 (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `taxId` | `string` | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | `string` | Target id type. |
| `tool` | `string` | The tool slug that ran. |

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
| `circular` | `boolean` | Treat the template as circular (plasmid). |
| `forwardPrimer` | `string` | Primer 1, 5'→3'. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | Mismatches tolerated per primer. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `reversePrimer` | `string` | Primer 2, 5'→3' (order does not matter). |
| `template` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const in_silico_pcr = await client.InSilicoPcr().create({
  forwardPrimer: 'example_forwardPrimer',
  ok: 'example_ok',
  provenance: {},
  result: {},
  reversePrimer: 'example_reversePrimer',
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
| `addSecondaryMismatch` | `boolean` | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | `string` | First allele (single base) — gets the FAM tail. |
| `alleleB` | `string` | Second allele (single base) — gets the HEX tail. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | `number` | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | `number` | Minimum amplicon length for the common reverse primer. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `snpPosition` | `number` | 1-based position of the SNP on the forward strand. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | `number` | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const kasp_primer_design = await client.KaspPrimerDesign().create({
  alleleA: 'example_alleleA',
  alleleB: 'example_alleleB',
  ok: 'example_ok',
  provenance: {},
  result: {},
  snpPosition: 1,
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
| `dntpMM` | `number` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `number` | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` |  |
| `oligoNM` | `number` | Total strand concentration (nM). |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | `number` | Optional target Tm (°C). |
| `tmTolerance` | `number` | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | Maximum allowed mismatches per match. |
| `motif` | `string` | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `searchReverseStrand` | `boolean` | Also search the reverse strand. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `dntpMM` | `number` | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `number` | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` |  |
| `oligoNM` | `number` | Total strand concentration (nM). |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sourceSpecies` | `string` | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | `any[]` | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | `string` | Ensembl species slug to find homologs in (e.g. |
| `tool` | `string` | The tool slug that ran. |
| `type` | `string` | Homology type to return. |

#### Example: Create

```ts
const ortholog_map = await client.OrthologMap().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  symbols: [],
  targetSpecies: 'example_targetSpecies',
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
| `gap` | `number` | Linear gap penalty (per gap position). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | `number` | Match score. |
| `mismatch` | `number` | Mismatch penalty. |
| `mode` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `seqA` | `string` | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | `string` | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const pairwise_alignment = await client.PairwiseAlignment().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  seqA: 'example_seqA',
  seqB: 'example_seqB',
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `text` | `string` | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | `string` | The tool slug that ran. |

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
| `fileBase64` | `string` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | Optional original file name (echoed back). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const parse_sanger_trace = await client.ParseSangerTrace().create({
  fileBase64: 'example_fileBase64',
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `circular` | `boolean` | Treat the sequence as a circular plasmid (vs. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `circular` | `boolean` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
| `topN` | `number` | How many top-ranked backbone candidates to report. |

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
| `circular` | `boolean` | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
| `topN` | `number` | How many top-ranked backbone candidates to report. |

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
| `editEnd` | `number` | 1-based inclusive end of the region being changed. |
| `editStart` | `number` | 1-based inclusive start of the region being changed. |
| `frameStart` | `number` | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | `string` | Replacement bases (forward strand). |
| `ok` | `any` |  |
| `pbsLength` | `number` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `rttHomology` | `number` | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | `string` | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const prime_editing_design = await client.PrimeEditingDesign().create({
  editEnd: 1,
  editStart: 1,
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | `string` | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` | `any` |  |
| `overlapLength` | `number` | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | `number` | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `Record<string, any>` |  |
| `replaceEnd` | `number` | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | `number` | 1-based inclusive start of the region being replaced/deleted. |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `target` | `string` | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const prime_editing_twin_design = await client.PrimeEditingTwinDesign().create({
  newSequence: 'example_newSequence',
  ok: 'example_ok',
  provenance: {},
  replaceEnd: 1,
  replaceStart: 1,
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
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `targetEnd` | `number` | 1-based inclusive end of the target region (optional). |
| `targetStart` | `number` | 1-based inclusive start of a region the product must span (optional). |
| `template` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` | `number` |  |
| `tmMaxDiff` | `number` | Max Tm difference within a pair (°C). |
| `tmMin` | `number` |  |
| `tmOpt` | `number` |  |
| `tool` | `string` | The tool slug that ran. |

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
| `forwardPrimer` | `string` | Forward primer, 5'→3'. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | `number` | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `reversePrimer` | `string` | Reverse primer, 5'→3'. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const primer_specificity = await client.PrimerSpecificity().create({
  forwardPrimer: 'example_forwardPrimer',
  ok: 'example_ok',
  provenance: {},
  result: {},
  reversePrimer: 'example_reversePrimer',
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | `number` | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | `number` | Cap on the number of returned peptides. |
| `minMass` | `number` | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | `number` | Allowed missed internal cleavages (0–2). |
| `ok` | `any` |  |
| `protease` | `string` | Protease or chemical cleavage agent. |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const protein_annotate_poll = await client.ProteinAnnotatePoll().create({
  jobId: 'example_jobId',
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
| `appl` | `string` | Restrict to one member database (e.g. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | `boolean` | Include GO-term cross-references. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `scale` | `string` | Amino-acid scale. |
| `sequence` | `string` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | The tool slug that ran. |
| `window` | `number` | Sliding-window size (clamped to an odd number ≥ 1). |

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
| `chargeStep` | `number` | pH step for the net-charge titration curve (0–14). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | `number` | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` | `string` |  |
| `length` | `number` | Number of residues to generate. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `enzymes` | `any[]` | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` | `string` |  |
| `ok` | `any` |  |
| `organism` | `string` | Codon-usage host (ignored in degenerate mode). |
| `protein` | `string` | Protein sequence (one-letter codes; * for stop). |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `fileBase64` | `string` | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | Optional original file name (echoed back). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | `number` | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `read` | `string` | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | `string` | Expected reference sequence (FASTA or raw). |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `args` | `Record<string, any>` | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const save_permalink = await client.SavePermalink().create({
  args: {},
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `qualityOffset` | `number` | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `accession` | `string` | GenBank/RefSeq accession (e.g. |
| `db` | `string` | Database to query; auto-detects from the accession format. |
| `format` | `string` | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `from` | `string` | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | A FASTA or GenBank record to convert. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `to` | `string` | Output format. |
| `tool` | `string` | The tool slug that ran. |

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
| `endPrimerLength` | `number` | Length of the naive end primers taken from each end. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `number` | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `number` | Minimum ORF length in amino acids. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Gene symbol/name, e.g. |
| `maxResults` | `number` | Up to 20. |
| `ok` | `any` |  |
| `organism` | `string` | Organism name, e.g. |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `term` | `string` | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | `number` | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `reads` | `string` | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | `string` | The claimed/expected reference sequence. |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const sequencing_readback_verify = await client.SequencingReadbackVerify().create({
  ok: 'example_ok',
  provenance: {},
  reads: 'example_reads',
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
| `entries` | `Record<string, any>` | Initial named entries, e.g. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | `any[]` | Only return these entries; omit to return all of them. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sessionId` | `string` |  |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const session_get = await client.SessionGet().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sessionId: 'example_sessionId',
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
| `args` | `Record<string, any>` | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | `Record<string, any>` | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sessionId` | `string` |  |
| `tool` | `string` | The tool slug that ran. |
| `writeBack` | `Record<string, any>` | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

#### Example: Create

```ts
const session_run = await client.SessionRun().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sessionId: 'example_sessionId',
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
| `entries` | `Record<string, any>` | Named entries to add/overwrite, e.g. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sessionId` | `string` |  |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const session_set = await client.SessionSet().create({
  entries: {},
  ok: 'example_ok',
  provenance: {},
  result: {},
  sessionId: 'example_sessionId',
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | `number` | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `shRnaLoop` | `string` | Loop sequence used when assembling the shRNA cassette. |
| `target` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `provenance` | `Record<string, any>` |  |
| `residue` | `number` | 1-based residue number to change (editKind='aa'). |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `style` | `string` | Mutagenic primer style. |
| `targetAa` | `string` | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | `boolean` | Stop at the first stop codon. |
| `tool` | `string` | The tool slug that ran. |

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
| `assembly` | `string` | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `variant` | `string` | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

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
| `coding` | `boolean` | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | `number` | 1-based reading-frame start (used when coding is true). |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `query` | `string` | Query / variant sequence (raw or FASTA). |
| `reference` | `string` | Reference / wild-type sequence (raw or FASTA). |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
| `armTmTarget` | `number` | Target annealing Tm (°C) for primer arms. |
| `circular` | `boolean` | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | `string` | The sequence you claim you ended up with. |
| `coding` | `boolean` | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | `string` | Type IIS enzyme for Golden Gate. |
| `enzyme3` | `string` | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | 5′ enzyme (restriction method). |
| `fragmentPcrs` | `any[]` | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | `any[]` | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | `number` | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | Insert sequence (restriction method). |
| `insertPcr` | `Record<string, any>` | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | `string` | Assembly method used. |
| `names` | `any[]` | Optional labels for each fragment. |
| `ok` | `any` |  |
| `overlapLen` | `number` | Gibson homology-arm length (bp). |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |
| `vector` | `string` | Vector sequence (restriction method). |
| `vectorPcr` | `Record<string, any>` | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

#### Example: Create

```ts
const verify_assembly = await client.VerifyAssembly().create({
  claimedConstruct: 'example_claimedConstruct',
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
| `claimedConstruct` | `string` | The final sequence claimed to have been built. |
| `expectedFrameStart` | `number` | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | `string` | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | `string` | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | `string` | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | `number` | Mismatches tolerated per primer during PCR prediction. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `templateCircular` | `boolean` | Treat insertTemplate as circular (e.g. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const verify_construct = await client.VerifyConstruct().create({
  claimedConstruct: 'example_claimedConstruct',
  insertForwardPrimer: 'example_insertForwardPrimer',
  insertReversePrimer: 'example_insertReversePrimer',
  insertTemplate: 'example_insertTemplate',
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
| `circular` | `boolean` | Treat the sequence as circular (plasmid). |
| `enzymes` | `any[]` | Enzyme names to digest with. |
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | `string` | DNA ladder to plot alongside the sample lane. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `sequence` | `string` | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | The tool slug that ran. |

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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `rows` | `any[]` | Differential expression rows, one per gene. |
| `tool` | `string` | The tool slug that ran. |

#### Example: Create

```ts
const volcano_plot_data = await client.VolcanoPlotData().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  rows: [],
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
| `gate` | `any` | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | `number` | Maximum number of results to return (default 5, max 10). |
| `ok` | `any` |  |
| `provenance` | `Record<string, any>` |  |
| `query` | `string` | The search query. |
| `result` | `Record<string, any>` | Tool-specific output object. |
| `tool` | `string` | The tool slug that ran. |

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
