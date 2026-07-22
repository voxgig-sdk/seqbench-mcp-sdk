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
  frame_start?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  target: string
  target_position?: number
  tool: string
}

export interface BaseEditingDesignCreateData {
  editor?: string
  frame_start?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  target: string
  target_position?: number
  tool: string
}

export interface Batch {
  arg?: Record<string, any>
  input: string
  ok: any
  result: Record<string, any>
  tool: string
}

export interface BatchLoadMatch {
  arg?: Record<string, any>
  input?: string
  ok?: any
  result?: Record<string, any>
  tool?: string
}

export interface BatchCreateData {
  arg?: Record<string, any>
  input: string
  ok: any
  result: Record<string, any>
  tool: string
}

export interface BatchWorkflow {
  input: string
  ok: any
  result: Record<string, any>
  step: any[]
}

export interface BatchWorkflowLoadMatch {
  input?: string
  ok?: any
  result?: Record<string, any>
  step?: any[]
}

export interface BatchWorkflowCreateData {
  input: string
  ok: any
  result: Record<string, any>
  step: any[]
}

export interface CharacterizeSequence {
  end_primer_length?: number
  gate?: any
  max_orf?: number
  min_orf_aa?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface CharacterizeSequenceCreateData {
  end_primer_length?: number
  gate?: any
  max_orf?: number
  min_orf_aa?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface CloningSimulate {
  arm_tm_target?: number
  circular?: boolean
  enzyme?: string
  enzyme3?: string
  enzyme5?: string
  fragment?: any[]
  gate?: any
  insert?: string
  method: string
  name?: any[]
  ok: any
  overlap_len?: number
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  vector?: string
}

export interface CloningSimulateCreateData {
  arm_tm_target?: number
  circular?: boolean
  enzyme?: string
  enzyme3?: string
  enzyme5?: string
  fragment?: any[]
  gate?: any
  insert?: string
  method: string
  name?: any[]
  ok: any
  overlap_len?: number
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  vector?: string
}

export interface CodonAdaptationIndex {
  frame_start?: number
  gate?: any
  ok: any
  organism?: string
  provenance: Record<string, any>
  rare_threshold?: number
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface CodonAdaptationIndexCreateData {
  frame_start?: number
  gate?: any
  ok: any
  organism?: string
  provenance: Record<string, any>
  rare_threshold?: number
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
  avoid_enzyme?: any[]
  cryptic_orf_min_aa?: number
  frame_start?: number
  gate?: any
  gc_high?: number
  gc_low?: number
  gc_window?: number
  homopolymer_min?: number
  max_pass?: number
  ok: any
  organism?: string
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ConstructAutofixCreateData {
  avoid_enzyme?: any[]
  cryptic_orf_min_aa?: number
  frame_start?: number
  gate?: any
  gc_high?: number
  gc_low?: number
  gc_window?: number
  homopolymer_min?: number
  max_pass?: number
  ok: any
  organism?: string
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ConstructQc {
  avoid_enzyme?: any[]
  cryptic_orf_min_aa?: number
  frame_start?: number
  gate?: any
  gc_high?: number
  gc_low?: number
  gc_window?: number
  homopolymer_min?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ConstructQcCreateData {
  avoid_enzyme?: any[]
  cryptic_orf_min_aa?: number
  frame_start?: number
  gate?: any
  gc_high?: number
  gc_low?: number
  gc_window?: number
  homopolymer_min?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface CrisprGrnaDesign {
  gate?: any
  min_score?: number
  nuclease?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  search_reverse_strand?: boolean
  sequence: string
  tool: string
}

export interface CrisprGrnaDesignCreateData {
  gate?: any
  min_score?: number
  nuclease?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  search_reverse_strand?: boolean
  sequence: string
  tool: string
}

export interface CrisprHdrDonor {
  arm_length?: number
  block_pam?: boolean
  design_genotyping_primer?: boolean
  edit_end?: number
  edit_start?: number
  frame_start?: number
  gate?: any
  guide_end?: number
  guide_start?: number
  guide_strand?: string
  nuclease?: string
  ok: any
  provenance: Record<string, any>
  replacement: string
  result: Record<string, any>
  target_sequence: string
  tool: string
}

export interface CrisprHdrDonorCreateData {
  arm_length?: number
  block_pam?: boolean
  design_genotyping_primer?: boolean
  edit_end?: number
  edit_start?: number
  frame_start?: number
  gate?: any
  guide_end?: number
  guide_start?: number
  guide_strand?: string
  nuclease?: string
  ok: any
  provenance: Record<string, any>
  replacement: string
  result: Record<string, any>
  target_sequence: string
  tool: string
}

export interface CrisprOfftargetCheck {
  gate?: any
  max_mismatch?: number
  nuclease?: string
  ok: any
  protospacer: string
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface CrisprOfftargetCheckCreateData {
  gate?: any
  max_mismatch?: number
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
  sequence_a: string
  sequence_b: string
  tool: string
}

export interface CrossDimerCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence_a: string
  sequence_b: string
  tool: string
}

export interface DnaMolarity {
  gate?: any
  length?: number
  mass_ng?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence?: string
  tool: string
  type?: string
  volume_ul?: number
}

export interface DnaMolarityCreateData {
  gate?: any
  length?: number
  mass_ng?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence?: string
  tool: string
  type?: string
  volume_ul?: number
}

export interface DoubleDigest {
  enzyme_a: string
  enzyme_b: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface DoubleDigestCreateData {
  enzyme_a: string
  enzyme_b: string
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
  reaction: any[]
  result: Record<string, any>
  tool: string
}

export interface ExportEchoPicklistCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  reaction: any[]
  result: Record<string, any>
  tool: string
}

export interface ExportOpentronsProtocol {
  gate?: any
  ok: any
  protocol_name?: string
  provenance: Record<string, any>
  reaction: any[]
  result: Record<string, any>
  tool: string
}

export interface ExportOpentronsProtocolCreateData {
  gate?: any
  ok: any
  protocol_name?: string
  provenance: Record<string, any>
  reaction: any[]
  result: Record<string, any>
  tool: string
}

export interface ExportPlateLayout {
  gate?: any
  ok: any
  provenance: Record<string, any>
  reaction: any[]
  result: Record<string, any>
  tool: string
}

export interface ExportPlateLayoutCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  reaction: any[]
  result: Record<string, any>
  tool: string
}

export interface ExpressionHeatmapCluster {
  cluster_col?: boolean
  cluster_row?: boolean
  distance_metric?: string
  gate?: any
  gene: any[]
  linkage?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sample: any[]
  tool: string
  value: any[]
  z_score_row?: boolean
}

export interface ExpressionHeatmapClusterCreateData {
  cluster_col?: boolean
  cluster_row?: boolean
  distance_metric?: string
  gate?: any
  gene: any[]
  linkage?: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sample: any[]
  tool: string
  value: any[]
  z_score_row?: boolean
}

export interface FastqQcReport {
  gate?: any
  input: string
  ok: any
  provenance: Record<string, any>
  quality_offset?: number
  result: Record<string, any>
  tool: string
}

export interface FastqQcReportCreateData {
  gate?: any
  input: string
  ok: any
  provenance: Record<string, any>
  quality_offset?: number
  result: Record<string, any>
  tool: string
}

export interface FastqTrim {
  gate?: any
  input: string
  min_length?: number
  ok: any
  provenance: Record<string, any>
  quality_offset?: number
  quality_threshold?: number
  result: Record<string, any>
  tool: string
}

export interface FastqTrimCreateData {
  gate?: any
  input: string
  min_length?: number
  ok: any
  provenance: Record<string, any>
  quality_offset?: number
  quality_threshold?: number
  result: Record<string, any>
  tool: string
}

export interface FindOrf {
  gate?: any
  min_aa_length?: number
  ok: any
  provenance: Record<string, any>
  require_stop?: boolean
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface FindOrfCreateData {
  gate?: any
  min_aa_length?: number
  ok: any
  provenance: Record<string, any>
  require_stop?: boolean
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface FormatSequence {
  case_mode?: string
  convert?: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reverse?: boolean
  sequence: string
  strip_non_letter?: boolean
  tool: string
  width?: number
}

export interface FormatSequenceCreateData {
  case_mode?: string
  convert?: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reverse?: boolean
  sequence: string
  strip_non_letter?: boolean
  tool: string
  width?: number
}

export interface FunctionalEnrichment {
  background?: any[]
  collection?: any[]
  gate?: any
  gene: any[]
  max_term_size?: number
  min_term_size?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface FunctionalEnrichmentCreateData {
  background?: any[]
  collection?: any[]
  gate?: any
  gene: any[]
  max_term_size?: number
  min_term_size?: number
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
  compare_to_named_set?: string
  dataset?: string
  gate?: any
  ok: any
  overhang: any[]
  provenance: Record<string, any>
  result: Record<string, any>
  risk_threshold?: number
  tool: string
}

export interface GoldenGateFidelityCreateData {
  compare_to_named_set?: string
  dataset?: string
  gate?: any
  ok: any
  overhang: any[]
  provenance: Record<string, any>
  result: Record<string, any>
  risk_threshold?: number
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
  job_id: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface IdMapPollCreateData {
  gate?: any
  job_id: string
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
  tax_id?: string
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
  tax_id?: string
  to: string
  tool: string
}

export interface InSilicoPcr {
  circular?: boolean
  forward_primer: string
  gate?: any
  max_mismatch?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reverse_primer: string
  template: string
  tool: string
}

export interface InSilicoPcrCreateData {
  circular?: boolean
  forward_primer: string
  gate?: any
  max_mismatch?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reverse_primer: string
  template: string
  tool: string
}

export interface KaspPrimerDesign {
  add_secondary_mismatch?: boolean
  allele_a: string
  allele_b: string
  gate?: any
  max_amplicon?: number
  min_amplicon?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  snp_position: number
  target: string
  target_core_tm?: number
  tool: string
}

export interface KaspPrimerDesignCreateData {
  add_secondary_mismatch?: boolean
  allele_a: string
  allele_b: string
  gate?: any
  max_amplicon?: number
  min_amplicon?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  snp_position: number
  target: string
  target_core_tm?: number
  tool: string
}

export interface ListTool {
}

export interface ListToolLoadMatch {
}

export interface MeltingTemperature {
  dntp_mm?: number
  gate?: any
  mg_mm?: number
  na_mm?: number
  ok: any
  oligo_nm?: number
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  target_tm?: number
  tm_tolerance?: number
  tool: string
}

export interface MeltingTemperatureCreateData {
  dntp_mm?: number
  gate?: any
  mg_mm?: number
  na_mm?: number
  ok: any
  oligo_nm?: number
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  target_tm?: number
  tm_tolerance?: number
  tool: string
}

export interface MotifFinder {
  gate?: any
  max_mismatch?: number
  motif: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  search_reverse_strand?: boolean
  sequence: string
  tool: string
}

export interface MotifFinderCreateData {
  gate?: any
  max_mismatch?: number
  motif: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  search_reverse_strand?: boolean
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
  dntp_mm?: number
  gate?: any
  mg_mm?: number
  na_mm?: number
  ok: any
  oligo_nm?: number
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface OligoAnalysiCreateData {
  dntp_mm?: number
  gate?: any
  mg_mm?: number
  na_mm?: number
  ok: any
  oligo_nm?: number
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
  source_species?: string
  symbol: any[]
  target_species: string
  tool: string
  type?: string
}

export interface OrthologMapCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  source_species?: string
  symbol: any[]
  target_species: string
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
  seq_a: string
  seq_b: string
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
  seq_a: string
  seq_b: string
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
  file_base64: string
  file_name?: string
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface ParseSangerTraceCreateData {
  file_base64: string
  file_name?: string
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
  top_n?: number
}

export interface PlasmidFullReportCreateData {
  circular?: boolean
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
  top_n?: number
}

export interface PlasmidIdentify {
  circular?: boolean
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
  top_n?: number
}

export interface PlasmidIdentifyCreateData {
  circular?: boolean
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
  top_n?: number
}

export interface PrimeEditingDesign {
  edit_end: number
  edit_start: number
  frame_start?: number
  gate?: any
  inserted_seq?: string
  ok: any
  pbs_length?: number
  provenance: Record<string, any>
  result: Record<string, any>
  rtt_homology?: number
  target: string
  tool: string
}

export interface PrimeEditingDesignCreateData {
  edit_end: number
  edit_start: number
  frame_start?: number
  gate?: any
  inserted_seq?: string
  ok: any
  pbs_length?: number
  provenance: Record<string, any>
  result: Record<string, any>
  rtt_homology?: number
  target: string
  tool: string
}

export interface PrimeEditingTwinDesign {
  gate?: any
  new_sequence: string
  ok: any
  overlap_length?: number
  pbs_length?: number
  provenance: Record<string, any>
  replace_end: number
  replace_start: number
  result: Record<string, any>
  target: string
  tool: string
}

export interface PrimeEditingTwinDesignCreateData {
  gate?: any
  new_sequence: string
  ok: any
  overlap_length?: number
  pbs_length?: number
  provenance: Record<string, any>
  replace_end: number
  replace_start: number
  result: Record<string, any>
  target: string
  tool: string
}

export interface PrimerDesign {
  amplicon_max?: number
  amplicon_min?: number
  dntp_mm?: number
  gate?: any
  gc_max?: number
  gc_min?: number
  len_max?: number
  len_min?: number
  len_opt?: number
  max_return?: number
  mg_mm?: number
  na_mm?: number
  ok: any
  oligo_nm?: number
  provenance: Record<string, any>
  result: Record<string, any>
  target_end?: number
  target_start?: number
  template: string
  tm_max?: number
  tm_max_diff?: number
  tm_min?: number
  tm_opt?: number
  tool: string
}

export interface PrimerDesignCreateData {
  amplicon_max?: number
  amplicon_min?: number
  dntp_mm?: number
  gate?: any
  gc_max?: number
  gc_min?: number
  len_max?: number
  len_min?: number
  len_opt?: number
  max_return?: number
  mg_mm?: number
  na_mm?: number
  ok: any
  oligo_nm?: number
  provenance: Record<string, any>
  result: Record<string, any>
  target_end?: number
  target_start?: number
  template: string
  tm_max?: number
  tm_max_diff?: number
  tm_min?: number
  tm_opt?: number
  tool: string
}

export interface PrimerSpecificity {
  forward_primer: string
  gate?: any
  max_mismatch?: number
  max_product_length?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reverse_primer: string
  tool: string
}

export interface PrimerSpecificityCreateData {
  forward_primer: string
  gate?: any
  max_mismatch?: number
  max_product_length?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  reverse_primer: string
  tool: string
}

export interface ProteaseDigestion {
  gate?: any
  max_mass?: number
  max_peptide?: number
  min_mass?: number
  missed_cleavage?: number
  ok: any
  protease?: string
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ProteaseDigestionCreateData {
  gate?: any
  max_mass?: number
  max_peptide?: number
  min_mass?: number
  missed_cleavage?: number
  ok: any
  protease?: string
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ProteinAnnotatePoll {
  gate?: any
  job_id: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface ProteinAnnotatePollCreateData {
  gate?: any
  job_id: string
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface ProteinAnnotateSubmit {
  appl?: string
  gate?: any
  goterm?: boolean
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ProteinAnnotateSubmitCreateData {
  appl?: string
  gate?: any
  goterm?: boolean
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
  charge_step?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface ProteinPropertyCreateData {
  charge_step?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface RandomSequence {
  gate?: any
  gc_content?: number
  kind?: string
  length: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface RandomSequenceCreateData {
  gate?: any
  gc_content?: number
  kind?: string
  length: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface RestrictionSite {
  enzyme?: any[]
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface RestrictionSiteCreateData {
  enzyme?: any[]
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
  file_base64?: string
  file_name?: string
  gate?: any
  min_coverage?: number
  ok: any
  provenance: Record<string, any>
  read?: string
  reference: string
  result: Record<string, any>
  tool: string
}

export interface SangerVsReferenceCreateData {
  file_base64?: string
  file_name?: string
  gate?: any
  min_coverage?: number
  ok: any
  provenance: Record<string, any>
  read?: string
  reference: string
  result: Record<string, any>
  tool: string
}

export interface SavePermalink {
  arg: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface SavePermalinkCreateData {
  arg: Record<string, any>
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
  quality_offset?: number
  result: Record<string, any>
  tool: string
}

export interface SeqfileStatCreateData {
  gate?: any
  input: string
  ok: any
  provenance: Record<string, any>
  quality_offset?: number
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
  end_primer_length?: number
  gate?: any
  max_orf?: number
  min_orf_aa?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  tool: string
}

export interface SequenceReportCreateData {
  end_primer_length?: number
  gate?: any
  max_orf?: number
  min_orf_aa?: number
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
  max_result?: number
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
  max_result?: number
  ok: any
  organism?: string
  provenance: Record<string, any>
  result: Record<string, any>
  term?: string
  tool: string
}

export interface SequencingReadbackVerify {
  gate?: any
  min_supporting_read?: number
  ok: any
  provenance: Record<string, any>
  read: string
  reference: string
  result: Record<string, any>
  tool: string
}

export interface SequencingReadbackVerifyCreateData {
  gate?: any
  min_supporting_read?: number
  ok: any
  provenance: Record<string, any>
  read: string
  reference: string
  result: Record<string, any>
  tool: string
}

export interface SessionCreate {
  entry?: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface SessionCreateCreateData {
  entry?: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
}

export interface SessionGet {
  gate?: any
  name?: any[]
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  session_id: string
  tool: string
}

export interface SessionGetCreateData {
  gate?: any
  name?: any[]
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  session_id: string
  tool: string
}

export interface SessionRun {
  arg?: Record<string, any>
  from_session?: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  session_id: string
  tool: string
  write_back?: Record<string, any>
}

export interface SessionRunCreateData {
  arg?: Record<string, any>
  from_session?: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  session_id: string
  tool: string
  write_back?: Record<string, any>
}

export interface SessionSet {
  entry: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  session_id: string
  tool: string
}

export interface SessionSetCreateData {
  entry: Record<string, any>
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  session_id: string
  tool: string
}

export interface SirnaDesign {
  gate?: any
  min_reynold?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sh_rna_loop?: string
  target: string
  tool: string
}

export interface SirnaDesignCreateData {
  gate?: any
  min_reynold?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sh_rna_loop?: string
  target: string
  tool: string
}

export interface SiteDirectedMutagenesi {
  arm_tm_target?: number
  dntp_mm?: number
  edit_kind?: string
  frame_start?: number
  gate?: any
  mg_mm?: number
  na_mm?: number
  new_base?: string
  ok: any
  oligo_nm?: number
  organism?: string
  position?: number
  provenance: Record<string, any>
  residue?: number
  result: Record<string, any>
  style?: string
  target_aa?: string
  template: string
  tool: string
}

export interface SiteDirectedMutagenesiCreateData {
  arm_tm_target?: number
  dntp_mm?: number
  edit_kind?: string
  frame_start?: number
  gate?: any
  mg_mm?: number
  na_mm?: number
  new_base?: string
  ok: any
  oligo_nm?: number
  organism?: string
  position?: number
  provenance: Record<string, any>
  residue?: number
  result: Record<string, any>
  style?: string
  target_aa?: string
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
  to_stop?: boolean
  tool: string
}

export interface TranslateCreateData {
  frame?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  sequence: string
  to_stop?: boolean
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
  frame_start?: number
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
  frame_start?: number
  gate?: any
  ok: any
  provenance: Record<string, any>
  query: string
  reference: string
  result: Record<string, any>
  tool: string
}

export interface VerifyAssembly {
  arm_tm_target?: number
  circular?: boolean
  claimed_construct: string
  coding?: boolean
  enzyme?: string
  enzyme3?: string
  enzyme5?: string
  fragment?: any[]
  fragment_pcr?: any[]
  frame_start?: number
  gate?: any
  insert?: string
  insert_pcr?: Record<string, any>
  method: string
  name?: any[]
  ok: any
  overlap_len?: number
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  vector?: string
  vector_pcr?: Record<string, any>
}

export interface VerifyAssemblyCreateData {
  arm_tm_target?: number
  circular?: boolean
  claimed_construct: string
  coding?: boolean
  enzyme?: string
  enzyme3?: string
  enzyme5?: string
  fragment?: any[]
  fragment_pcr?: any[]
  frame_start?: number
  gate?: any
  insert?: string
  insert_pcr?: Record<string, any>
  method: string
  name?: any[]
  ok: any
  overlap_len?: number
  provenance: Record<string, any>
  result: Record<string, any>
  tool: string
  vector?: string
  vector_pcr?: Record<string, any>
}

export interface VerifyConstruct {
  claimed_construct: string
  expected_frame_start?: number
  gate?: any
  insert_forward_primer: string
  insert_reverse_primer: string
  insert_template: string
  max_primer_mismatch?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  template_circular?: boolean
  tool: string
}

export interface VerifyConstructCreateData {
  claimed_construct: string
  expected_frame_start?: number
  gate?: any
  insert_forward_primer: string
  insert_reverse_primer: string
  insert_template: string
  max_primer_mismatch?: number
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  template_circular?: boolean
  tool: string
}

export interface VirtualGel {
  circular?: boolean
  enzyme?: any[]
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
  enzyme?: any[]
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
  row: any[]
  tool: string
}

export interface VolcanoPlotDataCreateData {
  gate?: any
  ok: any
  provenance: Record<string, any>
  result: Record<string, any>
  row: any[]
  tool: string
}

export interface WebSearch {
  gate?: any
  max_result?: number
  ok: any
  provenance: Record<string, any>
  query: string
  result: Record<string, any>
  tool: string
}

export interface WebSearchCreateData {
  gate?: any
  max_result?: number
  ok: any
  provenance: Record<string, any>
  query: string
  result: Record<string, any>
  tool: string
}

