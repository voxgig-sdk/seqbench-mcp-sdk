# SeqbenchMcp TypeScript SDK Reference

Complete API reference for the SeqbenchMcp TypeScript SDK.


## SeqbenchMcpSDK

### Constructor

```ts
new SeqbenchMcpSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `SeqbenchMcpSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = SeqbenchMcpSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `SeqbenchMcpSDK` instance in test mode.


### Instance Methods

#### `AlphafoldLookup(data?: object)`

Create a new `AlphafoldLookup` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AlphafoldLookupEntity` instance.

#### `AsoDesign(data?: object)`

Create a new `AsoDesign` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AsoDesignEntity` instance.

#### `BaseEditingDesign(data?: object)`

Create a new `BaseEditingDesign` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BaseEditingDesignEntity` instance.

#### `Batch(data?: object)`

Create a new `Batch` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BatchEntity` instance.

#### `BatchWorkflow(data?: object)`

Create a new `BatchWorkflow` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BatchWorkflowEntity` instance.

#### `CharacterizeSequence(data?: object)`

Create a new `CharacterizeSequence` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CharacterizeSequenceEntity` instance.

#### `CloningSimulate(data?: object)`

Create a new `CloningSimulate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CloningSimulateEntity` instance.

#### `CodonAdaptationIndex(data?: object)`

Create a new `CodonAdaptationIndex` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CodonAdaptationIndexEntity` instance.

#### `CodonOptimize(data?: object)`

Create a new `CodonOptimize` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CodonOptimizeEntity` instance.

#### `ConstructAutofix(data?: object)`

Create a new `ConstructAutofix` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ConstructAutofixEntity` instance.

#### `ConstructQc(data?: object)`

Create a new `ConstructQc` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ConstructQcEntity` instance.

#### `CrisprGrnaDesign(data?: object)`

Create a new `CrisprGrnaDesign` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CrisprGrnaDesignEntity` instance.

#### `CrisprHdrDonor(data?: object)`

Create a new `CrisprHdrDonor` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CrisprHdrDonorEntity` instance.

#### `CrisprOfftargetCheck(data?: object)`

Create a new `CrisprOfftargetCheck` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CrisprOfftargetCheckEntity` instance.

#### `CrossDimer(data?: object)`

Create a new `CrossDimer` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CrossDimerEntity` instance.

#### `DnaMolarity(data?: object)`

Create a new `DnaMolarity` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DnaMolarityEntity` instance.

#### `DoubleDigest(data?: object)`

Create a new `DoubleDigest` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DoubleDigestEntity` instance.

#### `ExportEchoPicklist(data?: object)`

Create a new `ExportEchoPicklist` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ExportEchoPicklistEntity` instance.

#### `ExportOpentronsProtocol(data?: object)`

Create a new `ExportOpentronsProtocol` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ExportOpentronsProtocolEntity` instance.

#### `ExportPlateLayout(data?: object)`

Create a new `ExportPlateLayout` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ExportPlateLayoutEntity` instance.

#### `ExpressionHeatmapCluster(data?: object)`

Create a new `ExpressionHeatmapCluster` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ExpressionHeatmapClusterEntity` instance.

#### `FastqQcReport(data?: object)`

Create a new `FastqQcReport` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `FastqQcReportEntity` instance.

#### `FastqTrim(data?: object)`

Create a new `FastqTrim` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `FastqTrimEntity` instance.

#### `FindOrf(data?: object)`

Create a new `FindOrf` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `FindOrfEntity` instance.

#### `FormatSequence(data?: object)`

Create a new `FormatSequence` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `FormatSequenceEntity` instance.

#### `FunctionalEnrichment(data?: object)`

Create a new `FunctionalEnrichment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `FunctionalEnrichmentEntity` instance.

#### `GcContent(data?: object)`

Create a new `GcContent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GcContentEntity` instance.

#### `GeneDossier(data?: object)`

Create a new `GeneDossier` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GeneDossierEntity` instance.

#### `GeneExpression(data?: object)`

Create a new `GeneExpression` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GeneExpressionEntity` instance.

#### `GeneModel(data?: object)`

Create a new `GeneModel` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GeneModelEntity` instance.

#### `GoldenGateFidelity(data?: object)`

Create a new `GoldenGateFidelity` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GoldenGateFidelityEntity` instance.

#### `HgvsConvert(data?: object)`

Create a new `HgvsConvert` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `HgvsConvertEntity` instance.

#### `IdMapPoll(data?: object)`

Create a new `IdMapPoll` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `IdMapPollEntity` instance.

#### `IdMapSubmit(data?: object)`

Create a new `IdMapSubmit` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `IdMapSubmitEntity` instance.

#### `InSilicoPcr(data?: object)`

Create a new `InSilicoPcr` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `InSilicoPcrEntity` instance.

#### `KaspPrimerDesign(data?: object)`

Create a new `KaspPrimerDesign` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `KaspPrimerDesignEntity` instance.

#### `ListTool(data?: object)`

Create a new `ListTool` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ListToolEntity` instance.

#### `MeltingTemperature(data?: object)`

Create a new `MeltingTemperature` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MeltingTemperatureEntity` instance.

#### `MotifFinder(data?: object)`

Create a new `MotifFinder` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MotifFinderEntity` instance.

#### `MultipleSequenceAlignment(data?: object)`

Create a new `MultipleSequenceAlignment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MultipleSequenceAlignmentEntity` instance.

#### `OligoAnalysi(data?: object)`

Create a new `OligoAnalysi` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `OligoAnalysiEntity` instance.

#### `OrthologMap(data?: object)`

Create a new `OrthologMap` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `OrthologMapEntity` instance.

#### `PairwiseAlignment(data?: object)`

Create a new `PairwiseAlignment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PairwiseAlignmentEntity` instance.

#### `ParseGenbank(data?: object)`

Create a new `ParseGenbank` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ParseGenbankEntity` instance.

#### `ParseSangerTrace(data?: object)`

Create a new `ParseSangerTrace` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ParseSangerTraceEntity` instance.

#### `PlasmidAnnotate(data?: object)`

Create a new `PlasmidAnnotate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PlasmidAnnotateEntity` instance.

#### `PlasmidDeepAnnotate(data?: object)`

Create a new `PlasmidDeepAnnotate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PlasmidDeepAnnotateEntity` instance.

#### `PlasmidFullReport(data?: object)`

Create a new `PlasmidFullReport` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PlasmidFullReportEntity` instance.

#### `PlasmidIdentify(data?: object)`

Create a new `PlasmidIdentify` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PlasmidIdentifyEntity` instance.

#### `PrimeEditingDesign(data?: object)`

Create a new `PrimeEditingDesign` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PrimeEditingDesignEntity` instance.

#### `PrimeEditingTwinDesign(data?: object)`

Create a new `PrimeEditingTwinDesign` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PrimeEditingTwinDesignEntity` instance.

#### `PrimerDesign(data?: object)`

Create a new `PrimerDesign` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PrimerDesignEntity` instance.

#### `PrimerSpecificity(data?: object)`

Create a new `PrimerSpecificity` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PrimerSpecificityEntity` instance.

#### `ProteaseDigestion(data?: object)`

Create a new `ProteaseDigestion` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ProteaseDigestionEntity` instance.

#### `ProteinAnnotatePoll(data?: object)`

Create a new `ProteinAnnotatePoll` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ProteinAnnotatePollEntity` instance.

#### `ProteinAnnotateSubmit(data?: object)`

Create a new `ProteinAnnotateSubmit` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ProteinAnnotateSubmitEntity` instance.

#### `ProteinHydrophobicity(data?: object)`

Create a new `ProteinHydrophobicity` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ProteinHydrophobicityEntity` instance.

#### `ProteinProperty(data?: object)`

Create a new `ProteinProperty` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ProteinPropertyEntity` instance.

#### `RandomSequence(data?: object)`

Create a new `RandomSequence` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RandomSequenceEntity` instance.

#### `RestrictionSite(data?: object)`

Create a new `RestrictionSite` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RestrictionSiteEntity` instance.

#### `ReverseComplement(data?: object)`

Create a new `ReverseComplement` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ReverseComplementEntity` instance.

#### `ReverseTranslate(data?: object)`

Create a new `ReverseTranslate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ReverseTranslateEntity` instance.

#### `RnaFold(data?: object)`

Create a new `RnaFold` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RnaFoldEntity` instance.

#### `SangerVsReference(data?: object)`

Create a new `SangerVsReference` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SangerVsReferenceEntity` instance.

#### `SavePermalink(data?: object)`

Create a new `SavePermalink` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SavePermalinkEntity` instance.

#### `SeqfileStat(data?: object)`

Create a new `SeqfileStat` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SeqfileStatEntity` instance.

#### `SequenceFetch(data?: object)`

Create a new `SequenceFetch` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SequenceFetchEntity` instance.

#### `SequenceFormatConvert(data?: object)`

Create a new `SequenceFormatConvert` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SequenceFormatConvertEntity` instance.

#### `SequenceReport(data?: object)`

Create a new `SequenceReport` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SequenceReportEntity` instance.

#### `SequenceSearch(data?: object)`

Create a new `SequenceSearch` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SequenceSearchEntity` instance.

#### `SequencingReadbackVerify(data?: object)`

Create a new `SequencingReadbackVerify` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SequencingReadbackVerifyEntity` instance.

#### `SessionCreate(data?: object)`

Create a new `SessionCreate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SessionCreateEntity` instance.

#### `SessionGet(data?: object)`

Create a new `SessionGet` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SessionGetEntity` instance.

#### `SessionRun(data?: object)`

Create a new `SessionRun` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SessionRunEntity` instance.

#### `SessionSet(data?: object)`

Create a new `SessionSet` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SessionSetEntity` instance.

#### `SirnaDesign(data?: object)`

Create a new `SirnaDesign` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SirnaDesignEntity` instance.

#### `SiteDirectedMutagenesi(data?: object)`

Create a new `SiteDirectedMutagenesi` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SiteDirectedMutagenesiEntity` instance.

#### `Translate(data?: object)`

Create a new `Translate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TranslateEntity` instance.

#### `VariantAnnotate(data?: object)`

Create a new `VariantAnnotate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `VariantAnnotateEntity` instance.

#### `VariantComparator(data?: object)`

Create a new `VariantComparator` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `VariantComparatorEntity` instance.

#### `VerifyAssembly(data?: object)`

Create a new `VerifyAssembly` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `VerifyAssemblyEntity` instance.

#### `VerifyConstruct(data?: object)`

Create a new `VerifyConstruct` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `VerifyConstructEntity` instance.

#### `VirtualGel(data?: object)`

Create a new `VirtualGel` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `VirtualGelEntity` instance.

#### `VolcanoPlotData(data?: object)`

Create a new `VolcanoPlotData` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `VolcanoPlotDataEntity` instance.

#### `WebSearch(data?: object)`

Create a new `WebSearch` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `WebSearchEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `SeqbenchMcpSDK.test()`.

**Returns:** `SeqbenchMcpSDK` instance in test mode.


---

## AlphafoldLookupEntity

```ts
const alphafold_lookup = client.AlphafoldLookup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `string` | Yes | UniProt accession, e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.AlphafoldLookup().create({
  accession: 'example_accession',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AlphafoldLookupEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## AsoDesignEntity

```ts
const aso_design = client.AsoDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `number` | No | Total gapmer length (nt). |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `wing` | `number` | No | Modified-wing length on each side (nt); the central gap = length − 2×wing. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.AsoDesign().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  target: 'example_target',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AsoDesignEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## BaseEditingDesignEntity

```ts
const base_editing_design = client.BaseEditingDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `editor` | `string` | No | Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G). |
| `frameStart` | `number` | No | Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetPosition` | `number` | No | Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.BaseEditingDesign().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  target: 'example_target',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BaseEditingDesignEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## BatchEntity

```ts
const batch = client.Batch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `Record<string, any>` | No | Shared tool arguments applied to every record. |
| `capped` | `boolean` | Yes | True if input exceeded the record limit. |
| `columns` | `any[]` | Yes |  |
| `count` | `number` | Yes |  |
| `errors` | `number` | Yes |  |
| `input` | `string` | Yes | Multi-FASTA text or one sequence per line (max ~2,000,000 chars). |
| `limit` | `number` | Yes | Maximum records per call (500). |
| `provenance` | `Record<string, any>` | Yes |  |
| `rows` | `any[]` | Yes |  |
| `tool` | `string` | Yes | A batchable tool slug (see `GET /batch`). |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Batch().create({
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

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Batch().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BatchEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## BatchWorkflowEntity

```ts
const batch__workflow = client.BatchWorkflow()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `capped` | `boolean` | Yes |  |
| `columns` | `any[]` | Yes | Flattened "<step>·<tool>·<key>" column headers. |
| `count` | `number` | Yes |  |
| `errors` | `number` | Yes |  |
| `input` | `string` | Yes | Multi-FASTA text or one sequence per line. |
| `limit` | `number` | Yes | Maximum records per call (200). |
| `provenance` | `Record<string, any>` | Yes |  |
| `rows` | `any[]` | Yes |  |
| `steps` | `any[]` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.BatchWorkflow().create({
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

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.BatchWorkflow().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BatchWorkflowEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CharacterizeSequenceEntity

```ts
const characterize_sequence = client.CharacterizeSequence()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `endPrimerLength` | `number` | No | Length of the naive end primers taken from each end. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `number` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `number` | No | Minimum ORF length in amino acids (nucleotide input only). |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CharacterizeSequence().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CharacterizeSequenceEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CloningSimulateEntity

```ts
const cloning_simulate = client.CloningSimulate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `number` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `boolean` | No | Produce a circular product. |
| `enzyme` | `string` | No | Type IIS enzyme for Golden Gate (e.g. |
| `enzyme3` | `string` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | No | 5′ enzyme (restriction method). |
| `fragments` | `any[]` | No | Fragments (5′→3′), assembled head-to-tail. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | No | Insert sequence (restriction method). |
| `method` | `string` | Yes | Assembly method. |
| `names` | `any[]` | No | Optional labels for each fragment. |
| `ok` | `any` | Yes |  |
| `overlapLen` | `number` | No | Gibson homology-arm length (bp). |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `vector` | `string` | No | Vector sequence (restriction method). |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CloningSimulate().create({
  method: 'example_method',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CloningSimulateEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CodonAdaptationIndexEntity

```ts
const codon_adaptation_index = client.CodonAdaptationIndex()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frameStart` | `number` | No | 1-based position to start reading codons. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `rareThreshold` | `number` | No | Relative adaptiveness (w) below this flags a codon as rare. |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Coding sequence (DNA/RNA; should start in-frame at ATG). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CodonAdaptationIndex().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CodonAdaptationIndexEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CodonOptimizeEntity

```ts
const codon_optimize = client.CodonOptimize()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes | Protein sequence (one-letter codes). |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CodonOptimize().create({
  ok: 'example_ok',
  protein: 'example_protein',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CodonOptimizeEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ConstructAutofixEntity

```ts
const construct_autofix = client.ConstructAutofix()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoidEnzymes` | `any[]` | No | Enzyme names whose internal sites should be removed (e.g. |
| `crypticOrfMinAa` | `number` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `number` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `number` | No |  |
| `gcLow` | `number` | No |  |
| `gcWindow` | `number` | No |  |
| `homopolymerMin` | `number` | No |  |
| `maxPasses` | `number` | No | Repeat full passes until clean or no further progress. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No | Codon-usage table to prefer among synonymous options. |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ConstructAutofix().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ConstructAutofixEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ConstructQcEntity

```ts
const construct_qc = client.ConstructQc()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avoidEnzymes` | `any[]` | No | Enzyme names whose internal sites should be flagged as errors. |
| `crypticOrfMinAa` | `number` | No | Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged. |
| `frameStart` | `number` | No | 1-based nucleotide where the reading frame begins. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcHigh` | `number` | No | GC% above this flags a GC-rich window. |
| `gcLow` | `number` | No | GC% below this flags an AT-rich window. |
| `gcWindow` | `number` | No | Sliding-window size (nt) for GC-extreme scanning. |
| `homopolymerMin` | `number` | No | Minimum run length to flag a homopolymer. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ConstructQc().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ConstructQcEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CrisprGrnaDesignEntity

```ts
const crispr_grna_design = client.CrisprGrnaDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minScore` | `number` | No | Only return guides with a heuristic score at least this high (0–100). |
| `nuclease` | `string` | No | Nuclease id. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `boolean` | No | Also scan the reverse strand for guides. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CrisprGrnaDesign().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CrisprGrnaDesignEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CrisprHdrDonorEntity

```ts
const crispr_hdr_donor = client.CrisprHdrDonor()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armLength` | `number` | No | Homology arm length (bp) on each side. |
| `blockPam` | `boolean` | No | When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut. |
| `designGenotypingPrimers` | `boolean` | No | Also design a primer pair (on the original targetSequence) whose product spans the edit site. |
| `editEnd` | `number` | No | 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. |
| `editStart` | `number` | No | 1-based start of the region being replaced. |
| `frameStart` | `number` | No | Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `guideEnd` | `number` | No | 1-based forward-strand end of the guide's protospacer. |
| `guideStart` | `number` | No | 1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site). |
| `guideStrand` | `string` | No | Strand the guide's protospacer is on. |
| `nuclease` | `string` | No | Needed only when deriving the cut site from guideStart/guideEnd/guideStrand. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `replacement` | `string` | Yes | Sequence to insert/substitute ("" for a pure deletion). |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `targetSequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CrisprHdrDonor().create({
  ok: 'example_ok',
  provenance: {},
  replacement: 'example_replacement',
  result: {},
  targetSequence: 'example_targetSequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CrisprHdrDonorEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CrisprOfftargetCheckEntity

```ts
const crispr_offtarget_check = client.CrisprOfftargetCheck()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | No | Mismatches tolerated between the protospacer and a candidate genomic site. |
| `nuclease` | `string` | No | Nuclease id — determines the PAM pattern/side required at each candidate site. |
| `ok` | `any` | Yes |  |
| `protospacer` | `string` | Yes | The guide's protospacer sequence, 5'→3' (no PAM). |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CrisprOfftargetCheck().create({
  ok: 'example_ok',
  protospacer: 'example_protospacer',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CrisprOfftargetCheckEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CrossDimerEntity

```ts
const cross_dimer = client.CrossDimer()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequenceA` | `string` | Yes | First oligo (5'→3'). |
| `sequenceB` | `string` | Yes | Second oligo (5'→3'). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CrossDimer().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequenceA: 'example_sequenceA',
  sequenceB: 'example_sequenceB',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CrossDimerEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DnaMolarityEntity

```ts
const dna_molarity = client.DnaMolarity()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `length` | `number` | No | Length in bp (dsDNA) or nt (ssDNA/ssRNA). |
| `massNg` | `number` | No | Mass in nanograms. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | No | Optional sequence — overrides length and gives an exact molar mass from base composition. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `type` | `string` | No | Molecule type. |
| `volumeUl` | `number` | No | Volume in microlitres (0 = unknown; needed for concentration). |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.DnaMolarity().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DnaMolarityEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DoubleDigestEntity

```ts
const double_digest = client.DoubleDigest()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzymeA` | `string` | Yes | First enzyme name (e.g. |
| `enzymeB` | `string` | Yes | Second enzyme name (e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.DoubleDigest().create({
  enzymeA: 'example_enzymeA',
  enzymeB: 'example_enzymeB',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DoubleDigestEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ExportEchoPicklistEntity

```ts
const export_echo_picklist = client.ExportEchoPicklist()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `reactions` | `any[]` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ExportEchoPicklist().create({
  ok: 'example_ok',
  provenance: {},
  reactions: [],
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ExportEchoPicklistEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ExportOpentronsProtocolEntity

```ts
const export_opentrons_protocol = client.ExportOpentronsProtocol()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `protocolName` | `string` | No | Optional protocol name (used in the script's metadata). |
| `provenance` | `Record<string, any>` | Yes |  |
| `reactions` | `any[]` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ExportOpentronsProtocol().create({
  ok: 'example_ok',
  provenance: {},
  reactions: [],
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ExportOpentronsProtocolEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ExportPlateLayoutEntity

```ts
const export_plate_layout = client.ExportPlateLayout()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `reactions` | `any[]` | Yes | One entry per PCR reaction, up to 96 (a single 96-well plate). |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ExportPlateLayout().create({
  ok: 'example_ok',
  provenance: {},
  reactions: [],
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ExportPlateLayoutEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ExpressionHeatmapClusterEntity

```ts
const expression_heatmap_cluster = client.ExpressionHeatmapCluster()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clusterCols` | `boolean` | No | Cluster (reorder) samples. |
| `clusterRows` | `boolean` | No | Cluster (reorder) genes. |
| `distanceMetric` | `string` | No | correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `any[]` | Yes | Row (gene) labels. |
| `linkage` | `string` | No | average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `samples` | `any[]` | Yes | Column (sample) labels. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `values` | `any[]` | Yes | genes x samples numeric matrix — one row per gene, in the same order as `genes`. |
| `zScoreRows` | `boolean` | No | Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization). |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ExpressionHeatmapCluster().create({
  genes: [],
  ok: 'example_ok',
  provenance: {},
  result: {},
  samples: [],
  tool: 'example_tool',
  values: [],
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ExpressionHeatmapClusterEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## FastqQcReportEntity

```ts
const fastq_qc_report = client.FastqQcReport()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `qualityOffset` | `number` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.FastqQcReport().create({
  input: 'example_input',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `FastqQcReportEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## FastqTrimEntity

```ts
const fastq_trim = client.FastqTrim()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each). |
| `minLength` | `number` | No | Reads shorter than this after trimming are dropped. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `qualityOffset` | `number` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7). |
| `qualityThreshold` | `number` | No | 3' quality-trim threshold (Phred score). |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.FastqTrim().create({
  input: 'example_input',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `FastqTrimEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## FindOrfEntity

```ts
const find_orf = client.FindOrf()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minAaLength` | `number` | No | Minimum protein length (aa) to report. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `requireStop` | `boolean` | No | Only report ORFs terminated by a stop codon. |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.FindOrf().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `FindOrfEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## FormatSequenceEntity

```ts
const format_sequence = client.FormatSequence()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `caseMode` | `string` | No |  |
| `convert` | `string` | No | DNA→RNA (T→U) or RNA→DNA (U→T). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `reverse` | `boolean` | No | Reverse the sequence (no complement). |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `stripNonLetters` | `boolean` | No | Remove digits, spaces and gaps (keep letters only). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `width` | `number` | No | Line-wrap width; 0 = single line. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.FormatSequence().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `FormatSequenceEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## FunctionalEnrichmentEntity

```ts
const functional_enrichment = client.FunctionalEnrichment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `background` | `any[]` | No | Custom background/universe gene symbols. |
| `collections` | `any[]` | No | Which term collections to test. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `genes` | `any[]` | Yes | Query gene symbols (human, e.g. |
| `maxTermSize` | `number` | No | Skip terms/pathways with more than this many background genes (matches clusterProfiler's default). |
| `minTermSize` | `number` | No | Skip terms/pathways with fewer than this many background genes. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.FunctionalEnrichment().create({
  genes: [],
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `FunctionalEnrichmentEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GcContentEntity

```ts
const gc_content = client.GcContent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.GcContent().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GcContentEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GeneDossierEntity

```ts
const gene_dossier = client.GeneDossier()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.GeneDossier().create({
  gene: 'example_gene',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GeneDossierEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GeneExpressionEntity

```ts
const gene_expression = client.GeneExpression()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.GeneExpression().create({
  gene: 'example_gene',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GeneExpressionEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GeneModelEntity

```ts
const gene_model = client.GeneModel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | Yes | A human gene symbol ("TP53") or Ensembl gene ID ("ENSG00000141510"). |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.GeneModel().create({
  gene: 'example_gene',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GeneModelEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GoldenGateFidelityEntity

```ts
const golden_gate_fidelity = client.GoldenGateFidelity()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `compareToNamedSet` | `string` | No | Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison. |
| `dataset` | `string` | No | Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `overhangs` | `any[]` | Yes | The candidate 4-base overhangs for one assembly (e.g. |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `riskThreshold` | `number` | No | Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.GoldenGateFidelity().create({
  ok: 'example_ok',
  overhangs: [],
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GoldenGateFidelityEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## HgvsConvertEntity

```ts
const hgvs_convert = client.HgvsConvert()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `variant` | `string` | Yes | A full HGVS "c." variant description: "<accession or gene symbol>:c.<edit>", e.g. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.HgvsConvert().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
  variant: 'example_variant',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `HgvsConvertEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## IdMapPollEntity

```ts
const id_map_poll = client.IdMapPoll()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.IdMapPoll().create({
  jobId: 'example_jobId',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `IdMapPollEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## IdMapSubmitEntity

```ts
const id_map_submit = client.IdMapSubmit()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | Yes | Source id type: "Gene_Name", "Ensembl", "GeneID", "RefSeq_Protein", or "UniProtKB_AC-ID". |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ids` | `any[]` | Yes | The ids to map, up to 1000 (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `taxId` | `string` | No | NCBI taxonomy id to disambiguate a gene symbol (only used when from="Gene_Name"). |
| `to` | `string` | Yes | Target id type. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.IdMapSubmit().create({
  from: 'example_from',
  ids: [],
  ok: 'example_ok',
  provenance: {},
  result: {},
  to: 'example_to',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `IdMapSubmitEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## InSilicoPcrEntity

```ts
const in_silico_pcr = client.InSilicoPcr()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `boolean` | No | Treat the template as circular (plasmid). |
| `forwardPrimer` | `string` | Yes | Primer 1, 5'→3'. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | No | Mismatches tolerated per primer. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `reversePrimer` | `string` | Yes | Primer 2, 5'→3' (order does not matter). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.InSilicoPcr().create({
  forwardPrimer: 'example_forwardPrimer',
  ok: 'example_ok',
  provenance: {},
  result: {},
  reversePrimer: 'example_reversePrimer',
  template: 'example_template',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `InSilicoPcrEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## KaspPrimerDesignEntity

```ts
const kasp_primer_design = client.KaspPrimerDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addSecondaryMismatch` | `boolean` | No | Engineer the internal ARMS destabilising mismatch near the 3' end. |
| `alleleA` | `string` | Yes | First allele (single base) — gets the FAM tail. |
| `alleleB` | `string` | Yes | Second allele (single base) — gets the HEX tail. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxAmplicon` | `number` | No | Maximum amplicon length for the common reverse primer. |
| `minAmplicon` | `number` | No | Minimum amplicon length for the common reverse primer. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `snpPosition` | `number` | Yes | 1-based position of the SNP on the forward strand. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetCoreTm` | `number` | No | Target Tm (°C) for the allele-specific primer core (before the universal tail). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.KaspPrimerDesign().create({
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

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `KaspPrimerDesignEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ListToolEntity

```ts
const list_tool = client.ListTool()
```

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.ListTool().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ListToolEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MeltingTemperatureEntity

```ts
const melting_temperature = client.MeltingTemperature()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntpMM` | `number` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `number` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `number` | No | Total strand concentration (nM). |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `targetTm` | `number` | No | Optional target Tm (°C). |
| `tmTolerance` | `number` | No | Allowed +/- window (°C) around targetTm for the gate. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.MeltingTemperature().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MeltingTemperatureEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MotifFinderEntity

```ts
const motif_finder = client.MotifFinder()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | No | Maximum allowed mismatches per match. |
| `motif` | `string` | Yes | Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `searchReverseStrand` | `boolean` | No | Also search the reverse strand. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.MotifFinder().create({
  motif: 'example_motif',
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MotifFinderEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MultipleSequenceAlignmentEntity

```ts
const multiple_sequence_alignment = client.MultipleSequenceAlignment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | Two or more sequences in multi-FASTA format (>name / sequence). |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.MultipleSequenceAlignment().create({
  input: 'example_input',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MultipleSequenceAlignmentEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## OligoAnalysiEntity

```ts
const oligo_analysi = client.OligoAnalysi()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dntpMM` | `number` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `number` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `number` | No | Total strand concentration (nM). |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.OligoAnalysi().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `OligoAnalysiEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## OrthologMapEntity

```ts
const ortholog_map = client.OrthologMap()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sourceSpecies` | `string` | No | Ensembl species slug the symbols belong to (e.g. |
| `symbols` | `any[]` | Yes | Gene symbols to look up, up to 50 (e.g. |
| `targetSpecies` | `string` | Yes | Ensembl species slug to find homologs in (e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `type` | `string` | No | Homology type to return. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.OrthologMap().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  symbols: [],
  targetSpecies: 'example_targetSpecies',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `OrthologMapEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PairwiseAlignmentEntity

```ts
const pairwise_alignment = client.PairwiseAlignment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gap` | `number` | No | Linear gap penalty (per gap position). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `match` | `number` | No | Match score. |
| `mismatch` | `number` | No | Mismatch penalty. |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `seqA` | `string` | Yes | First sequence (raw or FASTA; nucleotide or protein). |
| `seqB` | `string` | Yes | Second sequence (raw or FASTA; nucleotide or protein). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PairwiseAlignment().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  seqA: 'example_seqA',
  seqB: 'example_seqB',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PairwiseAlignmentEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ParseGenbankEntity

```ts
const parse_genbank = client.ParseGenbank()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `text` | `string` | Yes | A GenBank flat file (LOCUS … FEATURES … ORIGIN … //). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ParseGenbank().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  text: 'example_text',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ParseGenbankEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ParseSangerTraceEntity

```ts
const parse_sanger_trace = client.ParseSangerTrace()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fileBase64` | `string` | Yes | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | No | Optional original file name (echoed back). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ParseSangerTrace().create({
  fileBase64: 'example_fileBase64',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ParseSangerTraceEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PlasmidAnnotateEntity

```ts
const plasmid_annotate = client.PlasmidAnnotate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PlasmidAnnotate().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PlasmidAnnotateEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PlasmidDeepAnnotateEntity

```ts
const plasmid_deep_annotate = client.PlasmidDeepAnnotate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `boolean` | No | Treat the sequence as a circular plasmid (vs. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PlasmidDeepAnnotate().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PlasmidDeepAnnotateEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PlasmidFullReportEntity

```ts
const plasmid_full_report = client.PlasmidFullReport()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `boolean` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `topN` | `number` | No | How many top-ranked backbone candidates to report. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PlasmidFullReport().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PlasmidFullReportEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PlasmidIdentifyEntity

```ts
const plasmid_identify = client.PlasmidIdentify()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `boolean` | No | Treat the query as a circular molecule (most plasmids are). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `topN` | `number` | No | How many top-ranked backbone candidates to report. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PlasmidIdentify().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PlasmidIdentifyEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PrimeEditingDesignEntity

```ts
const prime_editing_design = client.PrimeEditingDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `editEnd` | `number` | Yes | 1-based inclusive end of the region being changed. |
| `editStart` | `number` | Yes | 1-based inclusive start of the region being changed. |
| `frameStart` | `number` | No | Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertedSeq` | `string` | No | Replacement bases (forward strand). |
| `ok` | `any` | Yes |  |
| `pbsLength` | `number` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `rttHomology` | `number` | No | Homology length (nt) 3' of the edit that the RTT should include (typically 10-16). |
| `target` | `string` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PrimeEditingDesign().create({
  editEnd: 1,
  editStart: 1,
  ok: 'example_ok',
  provenance: {},
  result: {},
  target: 'example_target',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PrimeEditingDesignEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PrimeEditingTwinDesignEntity

```ts
const prime_editing_twin_design = client.PrimeEditingTwinDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `newSequence` | `string` | Yes | New sequence (forward strand) to install in place of [replaceStart, replaceEnd]. |
| `ok` | `any` | Yes |  |
| `overlapLength` | `number` | No | Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal. |
| `pbsLength` | `number` | No | Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned. |
| `provenance` | `Record<string, any>` | Yes |  |
| `replaceEnd` | `number` | Yes | 1-based inclusive end of the region being replaced/deleted. |
| `replaceStart` | `number` | Yes | 1-based inclusive start of the region being replaced/deleted. |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `target` | `string` | Yes | Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PrimeEditingTwinDesign().create({
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

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PrimeEditingTwinDesignEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PrimerDesignEntity

```ts
const primer_design = client.PrimerDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ampliconMax` | `number` | No |  |
| `ampliconMin` | `number` | No |  |
| `dntpMM` | `number` | No | Total [dNTP] (mM), chelates Mg2+. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcMax` | `number` | No |  |
| `gcMin` | `number` | No |  |
| `lenMax` | `number` | No |  |
| `lenMin` | `number` | No |  |
| `lenOpt` | `number` | No |  |
| `maxReturn` | `number` | No | Number of best pairs to return. |
| `mgMM` | `number` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | No | Monovalent cation [Na+]/[K+] (mM). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `number` | No | Total strand concentration (nM). |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `targetEnd` | `number` | No | 1-based inclusive end of the target region (optional). |
| `targetStart` | `number` | No | 1-based inclusive start of a region the product must span (optional). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tmMax` | `number` | No |  |
| `tmMaxDiff` | `number` | No | Max Tm difference within a pair (°C). |
| `tmMin` | `number` | No |  |
| `tmOpt` | `number` | No |  |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PrimerDesign().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  template: 'example_template',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PrimerDesignEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PrimerSpecificityEntity

```ts
const primer_specificity = client.PrimerSpecificity()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `forwardPrimer` | `string` | Yes | Forward primer, 5'→3'. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMismatches` | `number` | No | Mismatches tolerated per primer against a reference genome. |
| `maxProductLength` | `number` | No | Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `reversePrimer` | `string` | Yes | Reverse primer, 5'→3'. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PrimerSpecificity().create({
  forwardPrimer: 'example_forwardPrimer',
  ok: 'example_ok',
  provenance: {},
  result: {},
  reversePrimer: 'example_reversePrimer',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PrimerSpecificityEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ProteaseDigestionEntity

```ts
const protease_digestion = client.ProteaseDigestion()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxMass` | `number` | No | Optional upper bound on neutral monoisotopic mass (Da). |
| `maxPeptides` | `number` | No | Cap on the number of returned peptides. |
| `minMass` | `number` | No | Optional lower bound on neutral monoisotopic mass (Da). |
| `missedCleavages` | `number` | No | Allowed missed internal cleavages (0–2). |
| `ok` | `any` | Yes |  |
| `protease` | `string` | No | Protease or chemical cleavage agent. |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ProteaseDigestion().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ProteaseDigestionEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ProteinAnnotatePollEntity

```ts
const protein_annotate_poll = client.ProteinAnnotatePoll()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `jobId` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ProteinAnnotatePoll().create({
  jobId: 'example_jobId',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ProteinAnnotatePollEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ProteinAnnotateSubmitEntity

```ts
const protein_annotate_submit = client.ProteinAnnotateSubmit()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `appl` | `string` | No | Restrict to one member database (e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `goterms` | `boolean` | No | Include GO-term cross-references. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence, one-letter code (FASTA header, if any, is stripped). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ProteinAnnotateSubmit().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ProteinAnnotateSubmitEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ProteinHydrophobicityEntity

```ts
const protein_hydrophobicity = client.ProteinHydrophobicity()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `scale` | `string` | No | Amino-acid scale. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `window` | `number` | No | Sliding-window size (clamped to an odd number ≥ 1). |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ProteinHydrophobicity().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ProteinHydrophobicityEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ProteinPropertyEntity

```ts
const protein_property = client.ProteinProperty()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `chargeStep` | `number` | No | pH step for the net-charge titration curve (0–14). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Protein sequence (one-letter amino-acid codes; non-AA characters ignored). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ProteinProperty().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ProteinPropertyEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RandomSequenceEntity

```ts
const random_sequence = client.RandomSequence()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gcContent` | `number` | No | Target GC percentage 0..100 (dna/rna only); omit for uniform. |
| `kind` | `string` | No |  |
| `length` | `number` | Yes | Number of residues to generate. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.RandomSequence().create({
  length: 1,
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RandomSequenceEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RestrictionSiteEntity

```ts
const restriction_site = client.RestrictionSite()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `enzymes` | `any[]` | No | Enzyme names to scan; omit to scan all curated enzymes. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.RestrictionSite().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RestrictionSiteEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ReverseComplementEntity

```ts
const reverse_complement = client.ReverseComplement()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |
| `type` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ReverseComplement().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ReverseComplementEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ReverseTranslateEntity

```ts
const reverse_translate = client.ReverseTranslate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No | Codon-usage host (ignored in degenerate mode). |
| `protein` | `string` | Yes | Protein sequence (one-letter codes; * for stop). |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ReverseTranslate().create({
  ok: 'example_ok',
  protein: 'example_protein',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ReverseTranslateEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RnaFoldEntity

```ts
const rna_fold = client.RnaFold()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.RnaFold().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RnaFoldEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SangerVsReferenceEntity

```ts
const sanger_vs_reference = client.SangerVsReference()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `fileBase64` | `string` | No | The binary ABIF (.ab1 / .abi) trace file, base64-encoded. |
| `fileName` | `string` | No | Optional original file name (echoed back). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minCoverage` | `number` | No | Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `read` | `string` | No | Sanger read as FASTA or raw text (alternative to uploading an ABIF trace). |
| `reference` | `string` | Yes | Expected reference sequence (FASTA or raw). |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SangerVsReference().create({
  ok: 'example_ok',
  provenance: {},
  reference: 'example_reference',
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SangerVsReferenceEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SavePermalinkEntity

```ts
const save_permalink = client.SavePermalink()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `Record<string, any>` | Yes | Arguments for that tool, exactly as you would pass to it directly. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SavePermalink().create({
  args: {},
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SavePermalinkEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SeqfileStatEntity

```ts
const seqfile_stat = client.SeqfileStat()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | FASTA or FASTQ text (raw sequence is treated as single-record FASTA). |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `qualityOffset` | `number` | No | FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7). |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SeqfileStat().create({
  input: 'example_input',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SeqfileStatEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SequenceFetchEntity

```ts
const sequence_fetch = client.SequenceFetch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accession` | `string` | Yes | GenBank/RefSeq accession (e.g. |
| `db` | `string` | No | Database to query; auto-detects from the accession format. |
| `format` | `string` | No | Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SequenceFetch().create({
  accession: 'example_accession',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SequenceFetchEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SequenceFormatConvertEntity

```ts
const sequence_format_convert = client.SequenceFormatConvert()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | No | Input format; 'auto' sniffs it from the first meaningful line. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `input` | `string` | Yes | A FASTA or GenBank record to convert. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `to` | `string` | No | Output format. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SequenceFormatConvert().create({
  input: 'example_input',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SequenceFormatConvertEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SequenceReportEntity

```ts
const sequence_report = client.SequenceReport()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `endPrimerLength` | `number` | No | Length of the naive end primers taken from each end. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `maxOrfs` | `number` | No | Maximum number of ORFs to return, longest first. |
| `minOrfAa` | `number` | No | Minimum ORF length in amino acids. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SequenceReport().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SequenceReportEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SequenceSearchEntity

```ts
const sequence_search = client.SequenceSearch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `db` | `string` | No |  |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `gene` | `string` | No | Gene symbol/name, e.g. |
| `maxResults` | `number` | No | Up to 20. |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No | Organism name, e.g. |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `term` | `string` | No | Raw NCBI search term (advanced) — overrides gene/organism when given, e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SequenceSearch().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SequenceSearchEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SequencingReadbackVerifyEntity

```ts
const sequencing_readback_verify = client.SequencingReadbackVerify()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minSupportingReads` | `number` | No | Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `reads` | `string` | Yes | Raw reads in FASTA or FASTQ format (auto-detected). |
| `reference` | `string` | Yes | The claimed/expected reference sequence. |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SequencingReadbackVerify().create({
  ok: 'example_ok',
  provenance: {},
  reads: 'example_reads',
  reference: 'example_reference',
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SequencingReadbackVerifyEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SessionCreateEntity

```ts
const session_create = client.SessionCreate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entries` | `Record<string, any>` | No | Initial named entries, e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SessionCreate().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SessionCreateEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SessionGetEntity

```ts
const session_get = client.SessionGet()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `names` | `any[]` | No | Only return these entries; omit to return all of them. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SessionGet().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sessionId: 'example_sessionId',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SessionGetEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SessionRunEntity

```ts
const session_run = client.SessionRun()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `args` | `Record<string, any>` | No | Additional literal arguments, merged with the ones resolved from the session. |
| `fromSession` | `Record<string, any>` | No | Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |
| `writeBack` | `Record<string, any>` | No | Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SessionRun().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sessionId: 'example_sessionId',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SessionRunEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SessionSetEntity

```ts
const session_set = client.SessionSet()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `entries` | `Record<string, any>` | Yes | Named entries to add/overwrite, e.g. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sessionId` | `string` | Yes |  |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SessionSet().create({
  entries: {},
  ok: 'example_ok',
  provenance: {},
  result: {},
  sessionId: 'example_sessionId',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SessionSetEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SirnaDesignEntity

```ts
const sirna_design = client.SirnaDesign()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `minReynolds` | `number` | No | Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `shRnaLoop` | `string` | No | Loop sequence used when assembling the shRNA cassette. |
| `target` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SirnaDesign().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  target: 'example_target',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SirnaDesignEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SiteDirectedMutagenesiEntity

```ts
const site_directed_mutagenesi = client.SiteDirectedMutagenesi()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `number` | No | Target Tm (°C) for each template-binding arm. |
| `dntpMM` | `number` | No | Total [dNTP] (mM), chelates Mg2+. |
| `editKind` | `string` | No | Edit at the nucleotide or amino-acid level. |
| `frameStart` | `number` | No | 1-based position of the first base of codon 1 (editKind='aa'). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `mgMM` | `number` | No | Divalent cation [Mg2+] (mM). |
| `naMM` | `number` | No | Monovalent cation [Na+]/[K+] (mM). |
| `newBase` | `string` | No | Replacement base (editKind='nt'). |
| `ok` | `any` | Yes |  |
| `oligoNM` | `number` | No | Total strand concentration (nM). |
| `organism` | `string` | No | Codon-usage table for choosing the new codon (editKind='aa'). |
| `position` | `number` | No | 1-based position to substitute (editKind='nt'). |
| `provenance` | `Record<string, any>` | Yes |  |
| `residue` | `number` | No | 1-based residue number to change (editKind='aa'). |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `style` | `string` | No | Mutagenic primer style. |
| `targetAa` | `string` | No | Target amino acid, one-letter code incl '*' (editKind='aa'). |
| `template` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SiteDirectedMutagenesi().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  template: 'example_template',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SiteDirectedMutagenesiEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TranslateEntity

```ts
const translate = client.Translate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `frame` | `number` | No |  |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `toStop` | `boolean` | No | Stop at the first stop codon. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Translate().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TranslateEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## VariantAnnotateEntity

```ts
const variant_annotate = client.VariantAnnotate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `assembly` | `string` | No | Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `variant` | `string` | Yes | An rsID ("rs1042522"), chrom:pos:ref:alt ("17:7676154:G:C", single-base substitutions only), genomic HGVS ("chr17:g.7676154G>C" or "17:g.7676154G>C"), or transcript HGVS c. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.VariantAnnotate().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
  variant: 'example_variant',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `VariantAnnotateEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## VariantComparatorEntity

```ts
const variant_comparator = client.VariantComparator()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `coding` | `boolean` | No | Treat as a coding sequence and report amino-acid effects. |
| `frameStart` | `number` | No | 1-based reading-frame start (used when coding is true). |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `query` | `string` | Yes | Query / variant sequence (raw or FASTA). |
| `reference` | `string` | Yes | Reference / wild-type sequence (raw or FASTA). |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.VariantComparator().create({
  ok: 'example_ok',
  provenance: {},
  query: 'example_query',
  reference: 'example_reference',
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `VariantComparatorEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## VerifyAssemblyEntity

```ts
const verify_assembly = client.VerifyAssembly()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `armTmTarget` | `number` | No | Target annealing Tm (°C) for primer arms. |
| `circular` | `boolean` | No | Treat the product/claimed construct as circular (most plasmids are). |
| `claimedConstruct` | `string` | Yes | The sequence you claim you ended up with. |
| `coding` | `boolean` | No | Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence. |
| `enzyme` | `string` | No | Type IIS enzyme for Golden Gate. |
| `enzyme3` | `string` | No | 3′ enzyme (restriction method). |
| `enzyme5` | `string` | No | 5′ enzyme (restriction method). |
| `fragmentPcrs` | `any[]` | No | Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead. |
| `fragments` | `any[]` | No | Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). |
| `frameStart` | `number` | No | 1-based reading-frame start on claimedConstruct, used when coding is true. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insert` | `string` | No | Insert sequence (restriction method). |
| `insertPcr` | `Record<string, any>` | No | Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |
| `method` | `string` | Yes | Assembly method used. |
| `names` | `any[]` | No | Optional labels for each fragment. |
| `ok` | `any` | Yes |  |
| `overlapLen` | `number` | No | Gibson homology-arm length (bp). |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |
| `vector` | `string` | No | Vector sequence (restriction method). |
| `vectorPcr` | `Record<string, any>` | No | Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.VerifyAssembly().create({
  claimedConstruct: 'example_claimedConstruct',
  method: 'example_method',
  ok: 'example_ok',
  provenance: {},
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `VerifyAssemblyEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## VerifyConstructEntity

```ts
const verify_construct = client.VerifyConstruct()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `claimedConstruct` | `string` | Yes | The final sequence claimed to have been built. |
| `expectedFrameStart` | `number` | No | 1-based position in claimedConstruct where the intended reading frame begins. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `insertForwardPrimer` | `string` | Yes | Forward primer used to amplify the insert, 5'→3'. |
| `insertReversePrimer` | `string` | Yes | Reverse primer used to amplify the insert, 5'→3'. |
| `insertTemplate` | `string` | Yes | PCR template the insert was amplified from. |
| `maxPrimerMismatches` | `number` | No | Mismatches tolerated per primer during PCR prediction. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `templateCircular` | `boolean` | No | Treat insertTemplate as circular (e.g. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.VerifyConstruct().create({
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

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `VerifyConstructEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## VirtualGelEntity

```ts
const virtual_gel = client.VirtualGel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `circular` | `boolean` | No | Treat the sequence as circular (plasmid). |
| `enzymes` | `any[]` | No | Enzyme names to digest with. |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ladder` | `string` | No | DNA ladder to plot alongside the sample lane. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `sequence` | `string` | Yes | Nucleotide sequence (raw or FASTA; IUPAC accepted). |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.VirtualGel().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence: 'example_sequence',
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `VirtualGelEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## VolcanoPlotDataEntity

```ts
const volcano_plot_data = client.VolcanoPlotData()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `rows` | `any[]` | Yes | Differential expression rows, one per gene. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.VolcanoPlotData().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  rows: [],
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `VolcanoPlotDataEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## WebSearchEntity

```ts
const web_search = client.WebSearch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `gate` | `any` | No | Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. |
| `max_results` | `number` | No | Maximum number of results to return (default 5, max 10). |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `query` | `string` | Yes | The search query. |
| `result` | `Record<string, any>` | Yes | Tool-specific output object. |
| `tool` | `string` | Yes | The tool slug that ran. |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.WebSearch().create({
  ok: 'example_ok',
  provenance: {},
  query: 'example_query',
  result: {},
  tool: 'example_tool',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `WebSearchEntity` instance with the same client and
options.

#### `client()`

Return the parent `SeqbenchMcpSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `ratelimit` | 0.0.1 | Client-side rate limiting via a token bucket |
| `retry` | 0.0.1 | Automatic retry of transient failures with exponential backoff |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |
| `timeout` | 0.0.1 | Per-request timeout with transport abort |


Features are activated via the `feature` option:

```ts
const client = new SeqbenchMcpSDK({
  feature: {
    ratelimit: { active: true },
    retry: { active: true },
    test: { active: true },
    timeout: { active: true },
  }
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### Ordering

`ratelimit`, `retry`, `timeout` wrap the transport. Each
wraps whatever is already installed, so **activation order is nesting order**:
a feature activated later sits OUTSIDE one activated earlier, and sees the call
first.

That decides behaviour, not just sequence: a feature that short-circuits the
call, such as a cache serving a hit, stops every feature nested inside it from
ever seeing that call.

`test` attach to pipeline hooks
rather than the transport, so their order does not affect what they observe.

#### `ratelimit`

Client-side rate limiting via a token bucket.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

| Option | Type |
|---|---|
| `now` | function |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.ratelimit.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `retry`

Automatic retry of transient failures with exponential backoff.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

| Option | Type |
|---|---|
| `jitter` | boolean |
| `sleep` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.retry.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

#### `timeout`

Per-request timeout with transport abort.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

| Option | Type |
|---|---|
| `clearTimer` | function |
| `setTimer` | function |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.timeout.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Wraps the transport: its place in the activation order decides what it
  sees. See [Ordering](#ordering) above.
- Inactive by default: leaving it out costs nothing at runtime.

