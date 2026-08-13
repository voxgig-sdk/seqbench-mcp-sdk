// Typed models for the SeqbenchMcp SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface AlphafoldLookup {
  accession: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface AlphafoldLookupCreateData {
  accession: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface AsoDesign {
  gate?: any
  length?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  target: string
  tool: string
  wing?: number
}

export interface AsoDesignCreateData {
  gate?: any
  length?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  target: string
  tool: string
  wing?: number
}

export interface BaseEditingDesign {
  editor?: string
  frameStart?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  target: string
  targetPosition?: number
  tool: string
}

export interface BaseEditingDesignCreateData {
  editor?: string
  frameStart?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  target: string
  targetPosition?: number
  tool: string
}

export interface Batch {
  args?: Record<string, any>
  capped: boolean
  columns: any[]
  count: number
  errors: number
  input: string
  limit: number
  provenance: Record<string, any>
  rows: any[]
  tool: string
}

export interface BatchLoadMatch {
  args?: Record<string, any>
  capped?: boolean
  columns?: any[]
  count?: number
  errors?: number
  input?: string
  limit?: number
  provenance?: Record<string, any>
  rows?: any[]
  tool?: string
}

export interface BatchCreateData {
  args?: Record<string, any>
  capped: boolean
  columns: any[]
  count: number
  errors: number
  input: string
  limit: number
  provenance: Record<string, any>
  rows: any[]
  tool: string
}

export interface BatchWorkflow {
  capped: boolean
  columns: any[]
  count: number
  errors: number
  input: string
  limit: number
  provenance: Record<string, any>
  rows: any[]
  steps: any[]
}

export interface BatchWorkflowLoadMatch {
  capped?: boolean
  columns?: any[]
  count?: number
  errors?: number
  input?: string
  limit?: number
  provenance?: Record<string, any>
  rows?: any[]
  steps?: any[]
}

export interface BatchWorkflowCreateData {
  capped: boolean
  columns: any[]
  count: number
  errors: number
  input: string
  limit: number
  provenance: Record<string, any>
  rows: any[]
  steps: any[]
}

export interface CharacterizeSequence {
  endPrimerLength?: number
  gate?: any
  maxOrfs?: number
  minOrfAa?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface CharacterizeSequenceCreateData {
  endPrimerLength?: number
  gate?: any
  maxOrfs?: number
  minOrfAa?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface CloningSimulate {
  armTmTarget?: number
  circular?: boolean
  enzyme?: string
  enzyme3?: string
  enzyme5?: string
  fragments?: any[]
  gate?: any
  insert?: string
  method: string
  names?: any[]
  ok: any
  overlapLen?: number
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  vector?: string
}

export interface CloningSimulateCreateData {
  armTmTarget?: number
  circular?: boolean
  enzyme?: string
  enzyme3?: string
  enzyme5?: string
  fragments?: any[]
  gate?: any
  insert?: string
  method: string
  names?: any[]
  ok: any
  overlapLen?: number
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  vector?: string
}

export interface CodonAdaptationIndex {
  frameStart?: number
  gate?: any
  ok: any
  organism?: string
  provenance: Record<string, any>
  rareThreshold?: number
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface CodonAdaptationIndexCreateData {
  frameStart?: number
  gate?: any
  ok: any
  organism?: string
  provenance: Record<string, any>
  rareThreshold?: number
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface CodonOptimize {
  gate?: any
  ok: any
  organism?: string
  protein: string
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface CodonOptimizeCreateData {
  gate?: any
  ok: any
  organism?: string
  protein: string
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface ConstructAutofix {
  avoidEnzymes?: any[]
  crypticOrfMinAa?: number
  frameStart?: number
  gate?: any
  gcHigh?: number
  gcLow?: number
  gcWindow?: number
  homopolymerMin?: number
  maxPasses?: number
  ok: any
  organism?: string
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ConstructAutofixCreateData {
  avoidEnzymes?: any[]
  crypticOrfMinAa?: number
  frameStart?: number
  gate?: any
  gcHigh?: number
  gcLow?: number
  gcWindow?: number
  homopolymerMin?: number
  maxPasses?: number
  ok: any
  organism?: string
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ConstructQc {
  avoidEnzymes?: any[]
  crypticOrfMinAa?: number
  frameStart?: number
  gate?: any
  gcHigh?: number
  gcLow?: number
  gcWindow?: number
  homopolymerMin?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ConstructQcCreateData {
  avoidEnzymes?: any[]
  crypticOrfMinAa?: number
  frameStart?: number
  gate?: any
  gcHigh?: number
  gcLow?: number
  gcWindow?: number
  homopolymerMin?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface CrisprGrnaDesign {
  gate?: any
  minScore?: number
  nuclease?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  searchReverseStrand?: boolean
  sequence: string
  tool: string
}

export interface CrisprGrnaDesignCreateData {
  gate?: any
  minScore?: number
  nuclease?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  searchReverseStrand?: boolean
  sequence: string
  tool: string
}

export interface CrisprHdrDonor {
  armLength?: number
  blockPam?: boolean
  designGenotypingPrimers?: boolean
  editEnd?: number
  editStart?: number
  frameStart?: number
  gate?: any
  guideEnd?: number
  guideStart?: number
  guideStrand?: string
  nuclease?: string
  ok: any
  provenance: Record<string, any>
  replacement: string
  result: Record<string, any>
  targetSequence: string
  tool: string
}

export interface CrisprHdrDonorCreateData {
  armLength?: number
  blockPam?: boolean
  designGenotypingPrimers?: boolean
  editEnd?: number
  editStart?: number
  frameStart?: number
  gate?: any
  guideEnd?: number
  guideStart?: number
  guideStrand?: string
  nuclease?: string
  ok: any
  provenance: Record<string, any>
  replacement: string
  result: Record<string, any>
  targetSequence: string
  tool: string
}

export interface CrisprOfftargetCheck {
  gate?: any
  maxMismatches?: number
  nuclease?: string
  ok: any
  protospacer: string
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface CrisprOfftargetCheckCreateData {
  gate?: any
  maxMismatches?: number
  nuclease?: string
  ok: any
  protospacer: string
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface CrossDimer {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequenceA: string
  sequenceB: string
  tool: string
}

export interface CrossDimerCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequenceA: string
  sequenceB: string
  tool: string
}

export interface DnaMolarity {
  gate?: any
  length?: number
  massNg?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence?: string
  tool: string
  type?: string
  volumeUl?: number
}

export interface DnaMolarityCreateData {
  gate?: any
  length?: number
  massNg?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence?: string
  tool: string
  type?: string
  volumeUl?: number
}

export interface DoubleDigest {
  enzymeA: string
  enzymeB: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface DoubleDigestCreateData {
  enzymeA: string
  enzymeB: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface ExportEchoPicklist {
  gate?: any
  ok: any
  provenance: Record<string, any>
  reactions: any[]
  result: Record<string, any>
  tool: string
}

export interface ExportEchoPicklistCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  reactions: any[]
  result: Record<string, any>
  tool: string
}

export interface ExportOpentronsProtocol {
  gate?: any
  ok: any
  protocolName?: string
  provenance: Record<string, any>
  reactions: any[]
  result: Record<string, any>
  tool: string
}

export interface ExportOpentronsProtocolCreateData {
  gate?: any
  ok: any
  protocolName?: string
  provenance: Record<string, any>
  reactions: any[]
  result: Record<string, any>
  tool: string
}

export interface ExportPlateLayout {
  gate?: any
  ok: any
  provenance: Record<string, any>
  reactions: any[]
  result: Record<string, any>
  tool: string
}

export interface ExportPlateLayoutCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  reactions: any[]
  result: Record<string, any>
  tool: string
}

export interface ExpressionHeatmapCluster {
  clusterCols?: boolean
  clusterRows?: boolean
  distanceMetric?: string
  gate?: any
  genes: any[]
  linkage?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  samples: any[]
  tool: string
  values: any[]
  zScoreRows?: boolean
}

export interface ExpressionHeatmapClusterCreateData {
  clusterCols?: boolean
  clusterRows?: boolean
  distanceMetric?: string
  gate?: any
  genes: any[]
  linkage?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  samples: any[]
  tool: string
  values: any[]
  zScoreRows?: boolean
}

export interface FastqQcReport {
  gate?: any
  input: string
  ok: any
  provenance: Record<string, any>
  qualityOffset?: number
  result: Record<string, any>
  tool: string
}

export interface FastqQcReportCreateData {
  gate?: any
  input: string
  ok: any
  provenance: Record<string, any>
  qualityOffset?: number
  result: Record<string, any>
  tool: string
}

export interface FastqTrim {
  gate?: any
  input: string
  minLength?: number
  ok: any
  provenance: Record<string, any>
  qualityOffset?: number
  qualityThreshold?: number
  result: Record<string, any>
  tool: string
}

export interface FastqTrimCreateData {
  gate?: any
  input: string
  minLength?: number
  ok: any
  provenance: Record<string, any>
  qualityOffset?: number
  qualityThreshold?: number
  result: Record<string, any>
  tool: string
}

export interface FindOrf {
  gate?: any
  minAaLength?: number
  ok: any
  provenance: Record<string, any>
  requireStop?: boolean
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface FindOrfCreateData {
  gate?: any
  minAaLength?: number
  ok: any
  provenance: Record<string, any>
  requireStop?: boolean
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface FormatSequence {
  caseMode?: string
  convert?: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reverse?: boolean
  sequence: string
  stripNonLetters?: boolean
  tool: string
  width?: number
}

export interface FormatSequenceCreateData {
  caseMode?: string
  convert?: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reverse?: boolean
  sequence: string
  stripNonLetters?: boolean
  tool: string
  width?: number
}

export interface FunctionalEnrichment {
  background?: any[]
  collections?: any[]
  gate?: any
  genes: any[]
  maxTermSize?: number
  minTermSize?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface FunctionalEnrichmentCreateData {
  background?: any[]
  collections?: any[]
  gate?: any
  genes: any[]
  maxTermSize?: number
  minTermSize?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface GcContent {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface GcContentCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface GeneDossier {
  gate?: any
  gene: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface GeneDossierCreateData {
  gate?: any
  gene: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface GeneExpression {
  gate?: any
  gene: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface GeneExpressionCreateData {
  gate?: any
  gene: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface GeneModel {
  gate?: any
  gene: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface GeneModelCreateData {
  gate?: any
  gene: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface GoldenGateFidelity {
  compareToNamedSet?: string
  dataset?: string
  gate?: any
  ok: any
  overhangs: any[]
  provenance: Record<string, any>
  result: Record<string, any>
  riskThreshold?: number
  tool: string
}

export interface GoldenGateFidelityCreateData {
  compareToNamedSet?: string
  dataset?: string
  gate?: any
  ok: any
  overhangs: any[]
  provenance: Record<string, any>
  result: Record<string, any>
  riskThreshold?: number
  tool: string
}

export interface HgvsConvert {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  variant: string
}

export interface HgvsConvertCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  variant: string
}

export interface IdMapPoll {
  gate?: any
  jobId: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface IdMapPollCreateData {
  gate?: any
  jobId: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface IdMapSubmit {
  from: string
  gate?: any
  ids: any[]
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  taxId?: string
  to: string
  tool: string
}

export interface IdMapSubmitCreateData {
  from: string
  gate?: any
  ids: any[]
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  taxId?: string
  to: string
  tool: string
}

export interface InSilicoPcr {
  circular?: boolean
  forwardPrimer: string
  gate?: any
  maxMismatches?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reversePrimer: string
  template: string
  tool: string
}

export interface InSilicoPcrCreateData {
  circular?: boolean
  forwardPrimer: string
  gate?: any
  maxMismatches?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reversePrimer: string
  template: string
  tool: string
}

export interface KaspPrimerDesign {
  addSecondaryMismatch?: boolean
  alleleA: string
  alleleB: string
  gate?: any
  maxAmplicon?: number
  minAmplicon?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  snpPosition: number
  target: string
  targetCoreTm?: number
  tool: string
}

export interface KaspPrimerDesignCreateData {
  addSecondaryMismatch?: boolean
  alleleA: string
  alleleB: string
  gate?: any
  maxAmplicon?: number
  minAmplicon?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  snpPosition: number
  target: string
  targetCoreTm?: number
  tool: string
}

export interface ListTool {
}

export interface ListToolLoadMatch {
}

export interface MeltingTemperature {
  dntpMM?: number
  gate?: any
  mgMM?: number
  naMM?: number
  ok: any
  oligoNM?: number
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  targetTm?: number
  tmTolerance?: number
  tool: string
}

export interface MeltingTemperatureCreateData {
  dntpMM?: number
  gate?: any
  mgMM?: number
  naMM?: number
  ok: any
  oligoNM?: number
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  targetTm?: number
  tmTolerance?: number
  tool: string
}

export interface MotifFinder {
  gate?: any
  maxMismatches?: number
  motif: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  searchReverseStrand?: boolean
  sequence: string
  tool: string
}

export interface MotifFinderCreateData {
  gate?: any
  maxMismatches?: number
  motif: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  searchReverseStrand?: boolean
  sequence: string
  tool: string
}

export interface MultipleSequenceAlignment {
  gate?: any
  input: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface MultipleSequenceAlignmentCreateData {
  gate?: any
  input: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface OligoAnalysi {
  dntpMM?: number
  gate?: any
  mgMM?: number
  naMM?: number
  ok: any
  oligoNM?: number
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface OligoAnalysiCreateData {
  dntpMM?: number
  gate?: any
  mgMM?: number
  naMM?: number
  ok: any
  oligoNM?: number
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface OrthologMap {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sourceSpecies?: string
  symbols: any[]
  targetSpecies: string
  tool: string
  type?: string
}

export interface OrthologMapCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sourceSpecies?: string
  symbols: any[]
  targetSpecies: string
  tool: string
  type?: string
}

export interface PairwiseAlignment {
  gap?: number
  gate?: any
  match?: number
  mismatch?: number
  mode?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  seqA: string
  seqB: string
  tool: string
}

export interface PairwiseAlignmentCreateData {
  gap?: number
  gate?: any
  match?: number
  mismatch?: number
  mode?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  seqA: string
  seqB: string
  tool: string
}

export interface ParseGenbank {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  text: string
  tool: string
}

export interface ParseGenbankCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  text: string
  tool: string
}

export interface ParseSangerTrace {
  fileBase64: string
  fileName?: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface ParseSangerTraceCreateData {
  fileBase64: string
  fileName?: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface PlasmidAnnotate {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface PlasmidAnnotateCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface PlasmidDeepAnnotate {
  circular?: boolean
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface PlasmidDeepAnnotateCreateData {
  circular?: boolean
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface PlasmidFullReport {
  circular?: boolean
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
  topN?: number
}

export interface PlasmidFullReportCreateData {
  circular?: boolean
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
  topN?: number
}

export interface PlasmidIdentify {
  circular?: boolean
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
  topN?: number
}

export interface PlasmidIdentifyCreateData {
  circular?: boolean
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
  topN?: number
}

export interface PrimeEditingDesign {
  editEnd: number
  editStart: number
  frameStart?: number
  gate?: any
  insertedSeq?: string
  ok: any
  pbsLength?: number
  provenance: Record<string, any>
  result: Record<string, any>
  rttHomology?: number
  target: string
  tool: string
}

export interface PrimeEditingDesignCreateData {
  editEnd: number
  editStart: number
  frameStart?: number
  gate?: any
  insertedSeq?: string
  ok: any
  pbsLength?: number
  provenance: Record<string, any>
  result: Record<string, any>
  rttHomology?: number
  target: string
  tool: string
}

export interface PrimeEditingTwinDesign {
  gate?: any
  newSequence: string
  ok: any
  overlapLength?: number
  pbsLength?: number
  provenance: Record<string, any>
  replaceEnd: number
  replaceStart: number
  result: Record<string, any>
  target: string
  tool: string
}

export interface PrimeEditingTwinDesignCreateData {
  gate?: any
  newSequence: string
  ok: any
  overlapLength?: number
  pbsLength?: number
  provenance: Record<string, any>
  replaceEnd: number
  replaceStart: number
  result: Record<string, any>
  target: string
  tool: string
}

export interface PrimerDesign {
  ampliconMax?: number
  ampliconMin?: number
  dntpMM?: number
  gate?: any
  gcMax?: number
  gcMin?: number
  lenMax?: number
  lenMin?: number
  lenOpt?: number
  maxReturn?: number
  mgMM?: number
  naMM?: number
  ok: any
  oligoNM?: number
  provenance: Record<string, any>
  result: Record<string, any>
  targetEnd?: number
  targetStart?: number
  template: string
  tmMax?: number
  tmMaxDiff?: number
  tmMin?: number
  tmOpt?: number
  tool: string
}

export interface PrimerDesignCreateData {
  ampliconMax?: number
  ampliconMin?: number
  dntpMM?: number
  gate?: any
  gcMax?: number
  gcMin?: number
  lenMax?: number
  lenMin?: number
  lenOpt?: number
  maxReturn?: number
  mgMM?: number
  naMM?: number
  ok: any
  oligoNM?: number
  provenance: Record<string, any>
  result: Record<string, any>
  targetEnd?: number
  targetStart?: number
  template: string
  tmMax?: number
  tmMaxDiff?: number
  tmMin?: number
  tmOpt?: number
  tool: string
}

export interface PrimerSpecificity {
  forwardPrimer: string
  gate?: any
  maxMismatches?: number
  maxProductLength?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reversePrimer: string
  tool: string
}

export interface PrimerSpecificityCreateData {
  forwardPrimer: string
  gate?: any
  maxMismatches?: number
  maxProductLength?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reversePrimer: string
  tool: string
}

export interface ProteaseDigestion {
  gate?: any
  maxMass?: number
  maxPeptides?: number
  minMass?: number
  missedCleavages?: number
  ok: any
  protease?: string
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ProteaseDigestionCreateData {
  gate?: any
  maxMass?: number
  maxPeptides?: number
  minMass?: number
  missedCleavages?: number
  ok: any
  protease?: string
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ProteinAnnotatePoll {
  gate?: any
  jobId: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface ProteinAnnotatePollCreateData {
  gate?: any
  jobId: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface ProteinAnnotateSubmit {
  appl?: string
  gate?: any
  goterms?: boolean
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ProteinAnnotateSubmitCreateData {
  appl?: string
  gate?: any
  goterms?: boolean
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ProteinHydrophobicity {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  scale?: string
  sequence: string
  tool: string
  window?: number
}

export interface ProteinHydrophobicityCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  scale?: string
  sequence: string
  tool: string
  window?: number
}

export interface ProteinProperty {
  chargeStep?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ProteinPropertyCreateData {
  chargeStep?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface RandomSequence {
  gate?: any
  gcContent?: number
  kind?: string
  length: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface RandomSequenceCreateData {
  gate?: any
  gcContent?: number
  kind?: string
  length: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface RestrictionSite {
  enzymes?: any[]
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface RestrictionSiteCreateData {
  enzymes?: any[]
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ReverseComplement {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
  type?: string
}

export interface ReverseComplementCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
  type?: string
}

export interface ReverseTranslate {
  gate?: any
  mode?: string
  ok: any
  organism?: string
  protein: string
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface ReverseTranslateCreateData {
  gate?: any
  mode?: string
  ok: any
  organism?: string
  protein: string
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface RnaFold {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface RnaFoldCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface SangerVsReference {
  fileBase64?: string
  fileName?: string
  gate?: any
  minCoverage?: number
  ok: any
  provenance: Record<string, any>
  read?: string
  reference: string
  result: Record<string, any>
  tool: string
}

export interface SangerVsReferenceCreateData {
  fileBase64?: string
  fileName?: string
  gate?: any
  minCoverage?: number
  ok: any
  provenance: Record<string, any>
  read?: string
  reference: string
  result: Record<string, any>
  tool: string
}

export interface SavePermalink {
  args: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface SavePermalinkCreateData {
  args: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface SeqfileStat {
  gate?: any
  input: string
  ok: any
  provenance: Record<string, any>
  qualityOffset?: number
  result: Record<string, any>
  tool: string
}

export interface SeqfileStatCreateData {
  gate?: any
  input: string
  ok: any
  provenance: Record<string, any>
  qualityOffset?: number
  result: Record<string, any>
  tool: string
}

export interface SequenceFetch {
  accession: string
  db?: string
  format?: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface SequenceFetchCreateData {
  accession: string
  db?: string
  format?: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface SequenceFormatConvert {
  from?: string
  gate?: any
  input: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  to?: string
  tool: string
}

export interface SequenceFormatConvertCreateData {
  from?: string
  gate?: any
  input: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  to?: string
  tool: string
}

export interface SequenceReport {
  endPrimerLength?: number
  gate?: any
  maxOrfs?: number
  minOrfAa?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface SequenceReportCreateData {
  endPrimerLength?: number
  gate?: any
  maxOrfs?: number
  minOrfAa?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface SequenceSearch {
  db?: string
  gate?: any
  gene?: string
  maxResults?: number
  ok: any
  organism?: string
  provenance: Record<string, any>
  result: Record<string, any>
  term?: string
  tool: string
}

export interface SequenceSearchCreateData {
  db?: string
  gate?: any
  gene?: string
  maxResults?: number
  ok: any
  organism?: string
  provenance: Record<string, any>
  result: Record<string, any>
  term?: string
  tool: string
}

export interface SequencingReadbackVerify {
  gate?: any
  minSupportingReads?: number
  ok: any
  provenance: Record<string, any>
  reads: string
  reference: string
  result: Record<string, any>
  tool: string
}

export interface SequencingReadbackVerifyCreateData {
  gate?: any
  minSupportingReads?: number
  ok: any
  provenance: Record<string, any>
  reads: string
  reference: string
  result: Record<string, any>
  tool: string
}

export interface SessionCreate {
  entries?: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface SessionCreateCreateData {
  entries?: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface SessionGet {
  gate?: any
  names?: any[]
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sessionId: string
  tool: string
}

export interface SessionGetCreateData {
  gate?: any
  names?: any[]
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sessionId: string
  tool: string
}

export interface SessionRun {
  args?: Record<string, any>
  fromSession?: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sessionId: string
  tool: string
  writeBack?: Record<string, any>
}

export interface SessionRunCreateData {
  args?: Record<string, any>
  fromSession?: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sessionId: string
  tool: string
  writeBack?: Record<string, any>
}

export interface SessionSet {
  entries: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sessionId: string
  tool: string
}

export interface SessionSetCreateData {
  entries: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sessionId: string
  tool: string
}

export interface SirnaDesign {
  gate?: any
  minReynolds?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  shRnaLoop?: string
  target: string
  tool: string
}

export interface SirnaDesignCreateData {
  gate?: any
  minReynolds?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  shRnaLoop?: string
  target: string
  tool: string
}

export interface SiteDirectedMutagenesi {
  armTmTarget?: number
  dntpMM?: number
  editKind?: string
  frameStart?: number
  gate?: any
  mgMM?: number
  naMM?: number
  newBase?: string
  ok: any
  oligoNM?: number
  organism?: string
  position?: number
  provenance: Record<string, any>
  residue?: number
  result: Record<string, any>
  style?: string
  targetAa?: string
  template: string
  tool: string
}

export interface SiteDirectedMutagenesiCreateData {
  armTmTarget?: number
  dntpMM?: number
  editKind?: string
  frameStart?: number
  gate?: any
  mgMM?: number
  naMM?: number
  newBase?: string
  ok: any
  oligoNM?: number
  organism?: string
  position?: number
  provenance: Record<string, any>
  residue?: number
  result: Record<string, any>
  style?: string
  targetAa?: string
  template: string
  tool: string
}

export interface Translate {
  frame?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  toStop?: boolean
  tool: string
}

export interface TranslateCreateData {
  frame?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  toStop?: boolean
  tool: string
}

export interface VariantAnnotate {
  assembly?: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  variant: string
}

export interface VariantAnnotateCreateData {
  assembly?: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  variant: string
}

export interface VariantComparator {
  coding?: boolean
  frameStart?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  query: string
  reference: string
  result: Record<string, any>
  tool: string
}

export interface VariantComparatorCreateData {
  coding?: boolean
  frameStart?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  query: string
  reference: string
  result: Record<string, any>
  tool: string
}

export interface VerifyAssembly {
  armTmTarget?: number
  circular?: boolean
  claimedConstruct: string
  coding?: boolean
  enzyme?: string
  enzyme3?: string
  enzyme5?: string
  fragmentPcrs?: any[]
  fragments?: any[]
  frameStart?: number
  gate?: any
  insert?: string
  insertPcr?: Record<string, any>
  method: string
  names?: any[]
  ok: any
  overlapLen?: number
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  vector?: string
  vectorPcr?: Record<string, any>
}

export interface VerifyAssemblyCreateData {
  armTmTarget?: number
  circular?: boolean
  claimedConstruct: string
  coding?: boolean
  enzyme?: string
  enzyme3?: string
  enzyme5?: string
  fragmentPcrs?: any[]
  fragments?: any[]
  frameStart?: number
  gate?: any
  insert?: string
  insertPcr?: Record<string, any>
  method: string
  names?: any[]
  ok: any
  overlapLen?: number
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  vector?: string
  vectorPcr?: Record<string, any>
}

export interface VerifyConstruct {
  claimedConstruct: string
  expectedFrameStart?: number
  gate?: any
  insertForwardPrimer: string
  insertReversePrimer: string
  insertTemplate: string
  maxPrimerMismatches?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  templateCircular?: boolean
  tool: string
}

export interface VerifyConstructCreateData {
  claimedConstruct: string
  expectedFrameStart?: number
  gate?: any
  insertForwardPrimer: string
  insertReversePrimer: string
  insertTemplate: string
  maxPrimerMismatches?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  templateCircular?: boolean
  tool: string
}

export interface VirtualGel {
  circular?: boolean
  enzymes?: any[]
  gate?: any
  ladder?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface VirtualGelCreateData {
  circular?: boolean
  enzymes?: any[]
  gate?: any
  ladder?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface VolcanoPlotData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  rows: any[]
  tool: string
}

export interface VolcanoPlotDataCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  rows: any[]
  tool: string
}

export interface WebSearch {
  gate?: any
  max_results?: number
  ok: any
  provenance: Record<string, any>
  query: string
  result: Record<string, any>
  tool: string
}

export interface WebSearchCreateData {
  gate?: any
  max_results?: number
  ok: any
  provenance: Record<string, any>
  query: string
  result: Record<string, any>
  tool: string
}

