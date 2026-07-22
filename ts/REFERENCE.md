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
| `accession` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `length` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `wing` | `number` | No |  |

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
| `editor` | `string` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `target` | `string` | Yes |  |
| `target_position` | `number` | No |  |
| `tool` | `string` | Yes |  |

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
| `arg` | `Record<string, any>` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Batch().create({
  input: 'example_input',
  ok: 'example_ok',
  result: {},
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
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `step` | `any[]` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.BatchWorkflow().create({
  input: 'example_input',
  ok: 'example_ok',
  result: {},
  step: [],
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
| `end_primer_length` | `number` | No |  |
| `gate` | `any` | No |  |
| `max_orf` | `number` | No |  |
| `min_orf_aa` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_tm_target` | `number` | No |  |
| `circular` | `boolean` | No |  |
| `enzyme` | `string` | No |  |
| `enzyme3` | `string` | No |  |
| `enzyme5` | `string` | No |  |
| `fragment` | `any[]` | No |  |
| `gate` | `any` | No |  |
| `insert` | `string` | No |  |
| `method` | `string` | Yes |  |
| `name` | `any[]` | No |  |
| `ok` | `any` | Yes |  |
| `overlap_len` | `number` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |
| `vector` | `string` | No |  |

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
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `rare_threshold` | `number` | No |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `avoid_enzyme` | `any[]` | No |  |
| `cryptic_orf_min_aa` | `number` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `gc_high` | `number` | No |  |
| `gc_low` | `number` | No |  |
| `gc_window` | `number` | No |  |
| `homopolymer_min` | `number` | No |  |
| `max_pass` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `avoid_enzyme` | `any[]` | No |  |
| `cryptic_orf_min_aa` | `number` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `gc_high` | `number` | No |  |
| `gc_low` | `number` | No |  |
| `gc_window` | `number` | No |  |
| `homopolymer_min` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `min_score` | `number` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `search_reverse_strand` | `boolean` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_length` | `number` | No |  |
| `block_pam` | `boolean` | No |  |
| `design_genotyping_primer` | `boolean` | No |  |
| `edit_end` | `number` | No |  |
| `edit_start` | `number` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `guide_end` | `number` | No |  |
| `guide_start` | `number` | No |  |
| `guide_strand` | `string` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `replacement` | `string` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `target_sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CrisprHdrDonor().create({
  ok: 'example_ok',
  provenance: {},
  replacement: 'example_replacement',
  result: {},
  target_sequence: 'example_target_sequence',
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
| `gate` | `any` | No |  |
| `max_mismatch` | `number` | No |  |
| `nuclease` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `protospacer` | `string` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence_a` | `string` | Yes |  |
| `sequence_b` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CrossDimer().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  sequence_a: 'example_sequence_a',
  sequence_b: 'example_sequence_b',
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
| `gate` | `any` | No |  |
| `length` | `number` | No |  |
| `mass_ng` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | No |  |
| `tool` | `string` | Yes |  |
| `type` | `string` | No |  |
| `volume_ul` | `number` | No |  |

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
| `enzyme_a` | `string` | Yes |  |
| `enzyme_b` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.DoubleDigest().create({
  enzyme_a: 'example_enzyme_a',
  enzyme_b: 'example_enzyme_b',
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `reaction` | `any[]` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ExportEchoPicklist().create({
  ok: 'example_ok',
  provenance: {},
  reaction: [],
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `protocol_name` | `string` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `reaction` | `any[]` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ExportOpentronsProtocol().create({
  ok: 'example_ok',
  provenance: {},
  reaction: [],
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `reaction` | `any[]` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ExportPlateLayout().create({
  ok: 'example_ok',
  provenance: {},
  reaction: [],
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
| `cluster_col` | `boolean` | No |  |
| `cluster_row` | `boolean` | No |  |
| `distance_metric` | `string` | No |  |
| `gate` | `any` | No |  |
| `gene` | `any[]` | Yes |  |
| `linkage` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sample` | `any[]` | Yes |  |
| `tool` | `string` | Yes |  |
| `value` | `any[]` | Yes |  |
| `z_score_row` | `boolean` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ExpressionHeatmapCluster().create({
  gene: [],
  ok: 'example_ok',
  provenance: {},
  result: {},
  sample: [],
  tool: 'example_tool',
  value: [],
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
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `quality_offset` | `number` | No |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `min_length` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `quality_offset` | `number` | No |  |
| `quality_threshold` | `number` | No |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `min_aa_length` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `require_stop` | `boolean` | No |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `case_mode` | `string` | No |  |
| `convert` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `reverse` | `boolean` | No |  |
| `sequence` | `string` | Yes |  |
| `strip_non_letter` | `boolean` | No |  |
| `tool` | `string` | Yes |  |
| `width` | `number` | No |  |

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
| `background` | `any[]` | No |  |
| `collection` | `any[]` | No |  |
| `gate` | `any` | No |  |
| `gene` | `any[]` | Yes |  |
| `max_term_size` | `number` | No |  |
| `min_term_size` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.FunctionalEnrichment().create({
  gene: [],
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `gene` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `compare_to_named_set` | `string` | No |  |
| `dataset` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `overhang` | `any[]` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `risk_threshold` | `number` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.GoldenGateFidelity().create({
  ok: 'example_ok',
  overhang: [],
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |
| `variant` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `job_id` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.IdMapPoll().create({
  job_id: 'example_job_id',
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
| `from` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `ids` | `any[]` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tax_id` | `string` | No |  |
| `to` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `circular` | `boolean` | No |  |
| `forward_primer` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `max_mismatch` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `reverse_primer` | `string` | Yes |  |
| `template` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.InSilicoPcr().create({
  forward_primer: 'example_forward_primer',
  ok: 'example_ok',
  provenance: {},
  result: {},
  reverse_primer: 'example_reverse_primer',
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
| `add_secondary_mismatch` | `boolean` | No |  |
| `allele_a` | `string` | Yes |  |
| `allele_b` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `max_amplicon` | `number` | No |  |
| `min_amplicon` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `snp_position` | `number` | Yes |  |
| `target` | `string` | Yes |  |
| `target_core_tm` | `number` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.KaspPrimerDesign().create({
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
| `dntp_mm` | `number` | No |  |
| `gate` | `any` | No |  |
| `mg_mm` | `number` | No |  |
| `na_mm` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `number` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `target_tm` | `number` | No |  |
| `tm_tolerance` | `number` | No |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `max_mismatch` | `number` | No |  |
| `motif` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `search_reverse_strand` | `boolean` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `dntp_mm` | `number` | No |  |
| `gate` | `any` | No |  |
| `mg_mm` | `number` | No |  |
| `na_mm` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `number` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `source_species` | `string` | No |  |
| `symbol` | `any[]` | Yes |  |
| `target_species` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `type` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.OrthologMap().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  symbol: [],
  target_species: 'example_target_species',
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
| `gap` | `number` | No |  |
| `gate` | `any` | No |  |
| `match` | `number` | No |  |
| `mismatch` | `number` | No |  |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `seq_a` | `string` | Yes |  |
| `seq_b` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PairwiseAlignment().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  seq_a: 'example_seq_a',
  seq_b: 'example_seq_b',
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `text` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `file_base64` | `string` | Yes |  |
| `file_name` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ParseSangerTrace().create({
  file_base64: 'example_file_base64',
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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `circular` | `boolean` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `circular` | `boolean` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `top_n` | `number` | No |  |

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
| `circular` | `boolean` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `top_n` | `number` | No |  |

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
| `edit_end` | `number` | Yes |  |
| `edit_start` | `number` | Yes |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `inserted_seq` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `pbs_length` | `number` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `rtt_homology` | `number` | No |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PrimeEditingDesign().create({
  edit_end: 1,
  edit_start: 1,
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
| `gate` | `any` | No |  |
| `new_sequence` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `overlap_length` | `number` | No |  |
| `pbs_length` | `number` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `replace_end` | `number` | Yes |  |
| `replace_start` | `number` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PrimeEditingTwinDesign().create({
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
| `amplicon_max` | `number` | No |  |
| `amplicon_min` | `number` | No |  |
| `dntp_mm` | `number` | No |  |
| `gate` | `any` | No |  |
| `gc_max` | `number` | No |  |
| `gc_min` | `number` | No |  |
| `len_max` | `number` | No |  |
| `len_min` | `number` | No |  |
| `len_opt` | `number` | No |  |
| `max_return` | `number` | No |  |
| `mg_mm` | `number` | No |  |
| `na_mm` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `number` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `target_end` | `number` | No |  |
| `target_start` | `number` | No |  |
| `template` | `string` | Yes |  |
| `tm_max` | `number` | No |  |
| `tm_max_diff` | `number` | No |  |
| `tm_min` | `number` | No |  |
| `tm_opt` | `number` | No |  |
| `tool` | `string` | Yes |  |

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
| `forward_primer` | `string` | Yes |  |
| `gate` | `any` | No |  |
| `max_mismatch` | `number` | No |  |
| `max_product_length` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `reverse_primer` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PrimerSpecificity().create({
  forward_primer: 'example_forward_primer',
  ok: 'example_ok',
  provenance: {},
  result: {},
  reverse_primer: 'example_reverse_primer',
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
| `gate` | `any` | No |  |
| `max_mass` | `number` | No |  |
| `max_peptide` | `number` | No |  |
| `min_mass` | `number` | No |  |
| `missed_cleavage` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `protease` | `string` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `job_id` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ProteinAnnotatePoll().create({
  job_id: 'example_job_id',
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
| `appl` | `string` | No |  |
| `gate` | `any` | No |  |
| `goterm` | `boolean` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `scale` | `string` | No |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `window` | `number` | No |  |

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
| `charge_step` | `number` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `gc_content` | `number` | No |  |
| `kind` | `string` | No |  |
| `length` | `number` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `enzyme` | `any[]` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
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
| `gate` | `any` | No |  |
| `mode` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `protein` | `string` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `file_base64` | `string` | No |  |
| `file_name` | `string` | No |  |
| `gate` | `any` | No |  |
| `min_coverage` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `read` | `string` | No |  |
| `reference` | `string` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arg` | `Record<string, any>` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SavePermalink().create({
  arg: {},
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
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `quality_offset` | `number` | No |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `accession` | `string` | Yes |  |
| `db` | `string` | No |  |
| `format` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `from` | `string` | No |  |
| `gate` | `any` | No |  |
| `input` | `string` | Yes |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `to` | `string` | No |  |
| `tool` | `string` | Yes |  |

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
| `end_primer_length` | `number` | No |  |
| `gate` | `any` | No |  |
| `max_orf` | `number` | No |  |
| `min_orf_aa` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `gene` | `string` | No |  |
| `max_result` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `organism` | `string` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `term` | `string` | No |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `min_supporting_read` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `read` | `string` | Yes |  |
| `reference` | `string` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SequencingReadbackVerify().create({
  ok: 'example_ok',
  provenance: {},
  read: 'example_read',
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
| `entry` | `Record<string, any>` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `name` | `any[]` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SessionGet().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  session_id: 'example_session_id',
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
| `arg` | `Record<string, any>` | No |  |
| `from_session` | `Record<string, any>` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |
| `write_back` | `Record<string, any>` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SessionRun().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  session_id: 'example_session_id',
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
| `entry` | `Record<string, any>` | Yes |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `session_id` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SessionSet().create({
  entry: {},
  ok: 'example_ok',
  provenance: {},
  result: {},
  session_id: 'example_session_id',
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
| `gate` | `any` | No |  |
| `min_reynold` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sh_rna_loop` | `string` | No |  |
| `target` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_tm_target` | `number` | No |  |
| `dntp_mm` | `number` | No |  |
| `edit_kind` | `string` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `mg_mm` | `number` | No |  |
| `na_mm` | `number` | No |  |
| `new_base` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `oligo_nm` | `number` | No |  |
| `organism` | `string` | No |  |
| `position` | `number` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `residue` | `number` | No |  |
| `result` | `Record<string, any>` | Yes |  |
| `style` | `string` | No |  |
| `target_aa` | `string` | No |  |
| `template` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `to_stop` | `boolean` | No |  |
| `tool` | `string` | Yes |  |

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
| `assembly` | `string` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |
| `variant` | `string` | Yes |  |

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
| `coding` | `boolean` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `query` | `string` | Yes |  |
| `reference` | `string` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `arm_tm_target` | `number` | No |  |
| `circular` | `boolean` | No |  |
| `claimed_construct` | `string` | Yes |  |
| `coding` | `boolean` | No |  |
| `enzyme` | `string` | No |  |
| `enzyme3` | `string` | No |  |
| `enzyme5` | `string` | No |  |
| `fragment` | `any[]` | No |  |
| `fragment_pcr` | `any[]` | No |  |
| `frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `insert` | `string` | No |  |
| `insert_pcr` | `Record<string, any>` | No |  |
| `method` | `string` | Yes |  |
| `name` | `any[]` | No |  |
| `ok` | `any` | Yes |  |
| `overlap_len` | `number` | No |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |
| `vector` | `string` | No |  |
| `vector_pcr` | `Record<string, any>` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.VerifyAssembly().create({
  claimed_construct: 'example_claimed_construct',
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
| `claimed_construct` | `string` | Yes |  |
| `expected_frame_start` | `number` | No |  |
| `gate` | `any` | No |  |
| `insert_forward_primer` | `string` | Yes |  |
| `insert_reverse_primer` | `string` | Yes |  |
| `insert_template` | `string` | Yes |  |
| `max_primer_mismatch` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `template_circular` | `boolean` | No |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.VerifyConstruct().create({
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
| `circular` | `boolean` | No |  |
| `enzyme` | `any[]` | No |  |
| `gate` | `any` | No |  |
| `ladder` | `string` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `sequence` | `string` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `gate` | `any` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `row` | `any[]` | Yes |  |
| `tool` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.VolcanoPlotData().create({
  ok: 'example_ok',
  provenance: {},
  result: {},
  row: [],
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
| `gate` | `any` | No |  |
| `max_result` | `number` | No |  |
| `ok` | `any` | Yes |  |
| `provenance` | `Record<string, any>` | Yes |  |
| `query` | `string` | Yes |  |
| `result` | `Record<string, any>` | Yes |  |
| `tool` | `string` | Yes |  |

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
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new SeqbenchMcpSDK({
  feature: {
    test: { active: true },
  }
})
```

