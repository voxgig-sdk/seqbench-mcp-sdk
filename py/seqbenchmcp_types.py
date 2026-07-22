# Typed models for the SeqbenchMcp SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class AlphafoldLookupRequired(TypedDict):
    accession: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class AlphafoldLookup(AlphafoldLookupRequired, total=False):
    gate: Any


class AlphafoldLookupCreateDataRequired(TypedDict):
    accession: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class AlphafoldLookupCreateData(AlphafoldLookupCreateDataRequired, total=False):
    gate: Any


class AsoDesignRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class AsoDesign(AsoDesignRequired, total=False):
    gate: Any
    length: int
    wing: int


class AsoDesignCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class AsoDesignCreateData(AsoDesignCreateDataRequired, total=False):
    gate: Any
    length: int
    wing: int


class BaseEditingDesignRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class BaseEditingDesign(BaseEditingDesignRequired, total=False):
    editor: str
    frame_start: int
    gate: Any
    target_position: int


class BaseEditingDesignCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class BaseEditingDesignCreateData(BaseEditingDesignCreateDataRequired, total=False):
    editor: str
    frame_start: int
    gate: Any
    target_position: int


class BatchRequired(TypedDict):
    input: str
    ok: Any
    result: dict
    tool: str


class Batch(BatchRequired, total=False):
    arg: dict


class BatchLoadMatch(TypedDict, total=False):
    arg: dict
    input: str
    ok: Any
    result: dict
    tool: str


class BatchCreateDataRequired(TypedDict):
    input: str
    ok: Any
    result: dict
    tool: str


class BatchCreateData(BatchCreateDataRequired, total=False):
    arg: dict


class BatchWorkflow(TypedDict):
    input: str
    ok: Any
    result: dict
    step: list


class BatchWorkflowLoadMatch(TypedDict, total=False):
    input: str
    ok: Any
    result: dict
    step: list


class BatchWorkflowCreateData(TypedDict):
    input: str
    ok: Any
    result: dict
    step: list


class CharacterizeSequenceRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CharacterizeSequence(CharacterizeSequenceRequired, total=False):
    end_primer_length: int
    gate: Any
    max_orf: int
    min_orf_aa: int


class CharacterizeSequenceCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CharacterizeSequenceCreateData(CharacterizeSequenceCreateDataRequired, total=False):
    end_primer_length: int
    gate: Any
    max_orf: int
    min_orf_aa: int


class CloningSimulateRequired(TypedDict):
    method: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class CloningSimulate(CloningSimulateRequired, total=False):
    arm_tm_target: float
    circular: bool
    enzyme: str
    enzyme3: str
    enzyme5: str
    fragment: list
    gate: Any
    insert: str
    name: list
    overlap_len: int
    vector: str


class CloningSimulateCreateDataRequired(TypedDict):
    method: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class CloningSimulateCreateData(CloningSimulateCreateDataRequired, total=False):
    arm_tm_target: float
    circular: bool
    enzyme: str
    enzyme3: str
    enzyme5: str
    fragment: list
    gate: Any
    insert: str
    name: list
    overlap_len: int
    vector: str


class CodonAdaptationIndexRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CodonAdaptationIndex(CodonAdaptationIndexRequired, total=False):
    frame_start: int
    gate: Any
    organism: str
    rare_threshold: float


class CodonAdaptationIndexCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CodonAdaptationIndexCreateData(CodonAdaptationIndexCreateDataRequired, total=False):
    frame_start: int
    gate: Any
    organism: str
    rare_threshold: float


class CodonOptimizeRequired(TypedDict):
    ok: Any
    protein: str
    provenance: dict
    result: dict
    tool: str


class CodonOptimize(CodonOptimizeRequired, total=False):
    gate: Any
    organism: str


class CodonOptimizeCreateDataRequired(TypedDict):
    ok: Any
    protein: str
    provenance: dict
    result: dict
    tool: str


class CodonOptimizeCreateData(CodonOptimizeCreateDataRequired, total=False):
    gate: Any
    organism: str


class ConstructAutofixRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ConstructAutofix(ConstructAutofixRequired, total=False):
    avoid_enzyme: list
    cryptic_orf_min_aa: int
    frame_start: int
    gate: Any
    gc_high: float
    gc_low: float
    gc_window: int
    homopolymer_min: int
    max_pass: int
    organism: str


class ConstructAutofixCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ConstructAutofixCreateData(ConstructAutofixCreateDataRequired, total=False):
    avoid_enzyme: list
    cryptic_orf_min_aa: int
    frame_start: int
    gate: Any
    gc_high: float
    gc_low: float
    gc_window: int
    homopolymer_min: int
    max_pass: int
    organism: str


class ConstructQcRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ConstructQc(ConstructQcRequired, total=False):
    avoid_enzyme: list
    cryptic_orf_min_aa: int
    frame_start: int
    gate: Any
    gc_high: float
    gc_low: float
    gc_window: int
    homopolymer_min: int


class ConstructQcCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ConstructQcCreateData(ConstructQcCreateDataRequired, total=False):
    avoid_enzyme: list
    cryptic_orf_min_aa: int
    frame_start: int
    gate: Any
    gc_high: float
    gc_low: float
    gc_window: int
    homopolymer_min: int


class CrisprGrnaDesignRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CrisprGrnaDesign(CrisprGrnaDesignRequired, total=False):
    gate: Any
    min_score: float
    nuclease: str
    search_reverse_strand: bool


class CrisprGrnaDesignCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CrisprGrnaDesignCreateData(CrisprGrnaDesignCreateDataRequired, total=False):
    gate: Any
    min_score: float
    nuclease: str
    search_reverse_strand: bool


class CrisprHdrDonorRequired(TypedDict):
    ok: Any
    provenance: dict
    replacement: str
    result: dict
    target_sequence: str
    tool: str


class CrisprHdrDonor(CrisprHdrDonorRequired, total=False):
    arm_length: int
    block_pam: bool
    design_genotyping_primer: bool
    edit_end: int
    edit_start: int
    frame_start: int
    gate: Any
    guide_end: int
    guide_start: int
    guide_strand: str
    nuclease: str


class CrisprHdrDonorCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    replacement: str
    result: dict
    target_sequence: str
    tool: str


class CrisprHdrDonorCreateData(CrisprHdrDonorCreateDataRequired, total=False):
    arm_length: int
    block_pam: bool
    design_genotyping_primer: bool
    edit_end: int
    edit_start: int
    frame_start: int
    gate: Any
    guide_end: int
    guide_start: int
    guide_strand: str
    nuclease: str


class CrisprOfftargetCheckRequired(TypedDict):
    ok: Any
    protospacer: str
    provenance: dict
    result: dict
    tool: str


class CrisprOfftargetCheck(CrisprOfftargetCheckRequired, total=False):
    gate: Any
    max_mismatch: int
    nuclease: str


class CrisprOfftargetCheckCreateDataRequired(TypedDict):
    ok: Any
    protospacer: str
    provenance: dict
    result: dict
    tool: str


class CrisprOfftargetCheckCreateData(CrisprOfftargetCheckCreateDataRequired, total=False):
    gate: Any
    max_mismatch: int
    nuclease: str


class CrossDimerRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence_a: str
    sequence_b: str
    tool: str


class CrossDimer(CrossDimerRequired, total=False):
    gate: Any


class CrossDimerCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence_a: str
    sequence_b: str
    tool: str


class CrossDimerCreateData(CrossDimerCreateDataRequired, total=False):
    gate: Any


class DnaMolarityRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str


class DnaMolarity(DnaMolarityRequired, total=False):
    gate: Any
    length: int
    mass_ng: float
    sequence: str
    type: str
    volume_ul: float


class DnaMolarityCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str


class DnaMolarityCreateData(DnaMolarityCreateDataRequired, total=False):
    gate: Any
    length: int
    mass_ng: float
    sequence: str
    type: str
    volume_ul: float


class DoubleDigestRequired(TypedDict):
    enzyme_a: str
    enzyme_b: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class DoubleDigest(DoubleDigestRequired, total=False):
    gate: Any


class DoubleDigestCreateDataRequired(TypedDict):
    enzyme_a: str
    enzyme_b: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class DoubleDigestCreateData(DoubleDigestCreateDataRequired, total=False):
    gate: Any


class ExportEchoPicklistRequired(TypedDict):
    ok: Any
    provenance: dict
    reaction: list
    result: dict
    tool: str


class ExportEchoPicklist(ExportEchoPicklistRequired, total=False):
    gate: Any


class ExportEchoPicklistCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    reaction: list
    result: dict
    tool: str


class ExportEchoPicklistCreateData(ExportEchoPicklistCreateDataRequired, total=False):
    gate: Any


class ExportOpentronsProtocolRequired(TypedDict):
    ok: Any
    provenance: dict
    reaction: list
    result: dict
    tool: str


class ExportOpentronsProtocol(ExportOpentronsProtocolRequired, total=False):
    gate: Any
    protocol_name: str


class ExportOpentronsProtocolCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    reaction: list
    result: dict
    tool: str


class ExportOpentronsProtocolCreateData(ExportOpentronsProtocolCreateDataRequired, total=False):
    gate: Any
    protocol_name: str


class ExportPlateLayoutRequired(TypedDict):
    ok: Any
    provenance: dict
    reaction: list
    result: dict
    tool: str


class ExportPlateLayout(ExportPlateLayoutRequired, total=False):
    gate: Any


class ExportPlateLayoutCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    reaction: list
    result: dict
    tool: str


class ExportPlateLayoutCreateData(ExportPlateLayoutCreateDataRequired, total=False):
    gate: Any


class ExpressionHeatmapClusterRequired(TypedDict):
    gene: list
    ok: Any
    provenance: dict
    result: dict
    sample: list
    tool: str
    value: list


class ExpressionHeatmapCluster(ExpressionHeatmapClusterRequired, total=False):
    cluster_col: bool
    cluster_row: bool
    distance_metric: str
    gate: Any
    linkage: str
    z_score_row: bool


class ExpressionHeatmapClusterCreateDataRequired(TypedDict):
    gene: list
    ok: Any
    provenance: dict
    result: dict
    sample: list
    tool: str
    value: list


class ExpressionHeatmapClusterCreateData(ExpressionHeatmapClusterCreateDataRequired, total=False):
    cluster_col: bool
    cluster_row: bool
    distance_metric: str
    gate: Any
    linkage: str
    z_score_row: bool


class FastqQcReportRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FastqQcReport(FastqQcReportRequired, total=False):
    gate: Any
    quality_offset: int


class FastqQcReportCreateDataRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FastqQcReportCreateData(FastqQcReportCreateDataRequired, total=False):
    gate: Any
    quality_offset: int


class FastqTrimRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FastqTrim(FastqTrimRequired, total=False):
    gate: Any
    min_length: int
    quality_offset: int
    quality_threshold: int


class FastqTrimCreateDataRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FastqTrimCreateData(FastqTrimCreateDataRequired, total=False):
    gate: Any
    min_length: int
    quality_offset: int
    quality_threshold: int


class FindOrfRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class FindOrf(FindOrfRequired, total=False):
    gate: Any
    min_aa_length: int
    require_stop: bool


class FindOrfCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class FindOrfCreateData(FindOrfCreateDataRequired, total=False):
    gate: Any
    min_aa_length: int
    require_stop: bool


class FormatSequenceRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class FormatSequence(FormatSequenceRequired, total=False):
    case_mode: str
    convert: str
    gate: Any
    reverse: bool
    strip_non_letter: bool
    width: int


class FormatSequenceCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class FormatSequenceCreateData(FormatSequenceCreateDataRequired, total=False):
    case_mode: str
    convert: str
    gate: Any
    reverse: bool
    strip_non_letter: bool
    width: int


class FunctionalEnrichmentRequired(TypedDict):
    gene: list
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FunctionalEnrichment(FunctionalEnrichmentRequired, total=False):
    background: list
    collection: list
    gate: Any
    max_term_size: int
    min_term_size: int


class FunctionalEnrichmentCreateDataRequired(TypedDict):
    gene: list
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FunctionalEnrichmentCreateData(FunctionalEnrichmentCreateDataRequired, total=False):
    background: list
    collection: list
    gate: Any
    max_term_size: int
    min_term_size: int


class GcContentRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class GcContent(GcContentRequired, total=False):
    gate: Any


class GcContentCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class GcContentCreateData(GcContentCreateDataRequired, total=False):
    gate: Any


class GeneDossierRequired(TypedDict):
    gene: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class GeneDossier(GeneDossierRequired, total=False):
    gate: Any


class GeneDossierCreateDataRequired(TypedDict):
    gene: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class GeneDossierCreateData(GeneDossierCreateDataRequired, total=False):
    gate: Any


class GeneExpressionRequired(TypedDict):
    gene: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class GeneExpression(GeneExpressionRequired, total=False):
    gate: Any


class GeneExpressionCreateDataRequired(TypedDict):
    gene: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class GeneExpressionCreateData(GeneExpressionCreateDataRequired, total=False):
    gate: Any


class GeneModelRequired(TypedDict):
    gene: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class GeneModel(GeneModelRequired, total=False):
    gate: Any


class GeneModelCreateDataRequired(TypedDict):
    gene: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class GeneModelCreateData(GeneModelCreateDataRequired, total=False):
    gate: Any


class GoldenGateFidelityRequired(TypedDict):
    ok: Any
    overhang: list
    provenance: dict
    result: dict
    tool: str


class GoldenGateFidelity(GoldenGateFidelityRequired, total=False):
    compare_to_named_set: str
    dataset: str
    gate: Any
    risk_threshold: float


class GoldenGateFidelityCreateDataRequired(TypedDict):
    ok: Any
    overhang: list
    provenance: dict
    result: dict
    tool: str


class GoldenGateFidelityCreateData(GoldenGateFidelityCreateDataRequired, total=False):
    compare_to_named_set: str
    dataset: str
    gate: Any
    risk_threshold: float


class HgvsConvertRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str
    variant: str


class HgvsConvert(HgvsConvertRequired, total=False):
    gate: Any


class HgvsConvertCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str
    variant: str


class HgvsConvertCreateData(HgvsConvertCreateDataRequired, total=False):
    gate: Any


class IdMapPollRequired(TypedDict):
    job_id: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class IdMapPoll(IdMapPollRequired, total=False):
    gate: Any


class IdMapPollCreateDataRequired(TypedDict):
    job_id: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class IdMapPollCreateData(IdMapPollCreateDataRequired, total=False):
    gate: Any


class IdMapSubmitRequired(TypedDict):
    ids: list
    ok: Any
    provenance: dict
    result: dict
    to: str
    tool: str


class IdMapSubmit(IdMapSubmitRequired, total=False):
    gate: Any
    tax_id: str


class IdMapSubmitCreateDataRequired(TypedDict):
    ids: list
    ok: Any
    provenance: dict
    result: dict
    to: str
    tool: str


class IdMapSubmitCreateData(IdMapSubmitCreateDataRequired, total=False):
    gate: Any
    tax_id: str


class InSilicoPcrRequired(TypedDict):
    forward_primer: str
    ok: Any
    provenance: dict
    result: dict
    reverse_primer: str
    template: str
    tool: str


class InSilicoPcr(InSilicoPcrRequired, total=False):
    circular: bool
    gate: Any
    max_mismatch: int


class InSilicoPcrCreateDataRequired(TypedDict):
    forward_primer: str
    ok: Any
    provenance: dict
    result: dict
    reverse_primer: str
    template: str
    tool: str


class InSilicoPcrCreateData(InSilicoPcrCreateDataRequired, total=False):
    circular: bool
    gate: Any
    max_mismatch: int


class KaspPrimerDesignRequired(TypedDict):
    allele_a: str
    allele_b: str
    ok: Any
    provenance: dict
    result: dict
    snp_position: int
    target: str
    tool: str


class KaspPrimerDesign(KaspPrimerDesignRequired, total=False):
    add_secondary_mismatch: bool
    gate: Any
    max_amplicon: int
    min_amplicon: int
    target_core_tm: float


class KaspPrimerDesignCreateDataRequired(TypedDict):
    allele_a: str
    allele_b: str
    ok: Any
    provenance: dict
    result: dict
    snp_position: int
    target: str
    tool: str


class KaspPrimerDesignCreateData(KaspPrimerDesignCreateDataRequired, total=False):
    add_secondary_mismatch: bool
    gate: Any
    max_amplicon: int
    min_amplicon: int
    target_core_tm: float


class ListTool(TypedDict):
    pass


class ListToolLoadMatch(TypedDict):
    pass


class MeltingTemperatureRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class MeltingTemperature(MeltingTemperatureRequired, total=False):
    dntp_mm: float
    gate: Any
    mg_mm: float
    na_mm: float
    oligo_nm: float
    target_tm: float
    tm_tolerance: float


class MeltingTemperatureCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class MeltingTemperatureCreateData(MeltingTemperatureCreateDataRequired, total=False):
    dntp_mm: float
    gate: Any
    mg_mm: float
    na_mm: float
    oligo_nm: float
    target_tm: float
    tm_tolerance: float


class MotifFinderRequired(TypedDict):
    motif: str
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class MotifFinder(MotifFinderRequired, total=False):
    gate: Any
    max_mismatch: int
    search_reverse_strand: bool


class MotifFinderCreateDataRequired(TypedDict):
    motif: str
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class MotifFinderCreateData(MotifFinderCreateDataRequired, total=False):
    gate: Any
    max_mismatch: int
    search_reverse_strand: bool


class MultipleSequenceAlignmentRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class MultipleSequenceAlignment(MultipleSequenceAlignmentRequired, total=False):
    gate: Any


class MultipleSequenceAlignmentCreateDataRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class MultipleSequenceAlignmentCreateData(MultipleSequenceAlignmentCreateDataRequired, total=False):
    gate: Any


class OligoAnalysiRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class OligoAnalysi(OligoAnalysiRequired, total=False):
    dntp_mm: float
    gate: Any
    mg_mm: float
    na_mm: float
    oligo_nm: float


class OligoAnalysiCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class OligoAnalysiCreateData(OligoAnalysiCreateDataRequired, total=False):
    dntp_mm: float
    gate: Any
    mg_mm: float
    na_mm: float
    oligo_nm: float


class OrthologMapRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    symbol: list
    target_species: str
    tool: str


class OrthologMap(OrthologMapRequired, total=False):
    gate: Any
    source_species: str
    type: str


class OrthologMapCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    symbol: list
    target_species: str
    tool: str


class OrthologMapCreateData(OrthologMapCreateDataRequired, total=False):
    gate: Any
    source_species: str
    type: str


class PairwiseAlignmentRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    seq_a: str
    seq_b: str
    tool: str


class PairwiseAlignment(PairwiseAlignmentRequired, total=False):
    gap: float
    gate: Any
    match: float
    mismatch: float
    mode: str


class PairwiseAlignmentCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    seq_a: str
    seq_b: str
    tool: str


class PairwiseAlignmentCreateData(PairwiseAlignmentCreateDataRequired, total=False):
    gap: float
    gate: Any
    match: float
    mismatch: float
    mode: str


class ParseGenbankRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    text: str
    tool: str


class ParseGenbank(ParseGenbankRequired, total=False):
    gate: Any


class ParseGenbankCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    text: str
    tool: str


class ParseGenbankCreateData(ParseGenbankCreateDataRequired, total=False):
    gate: Any


class ParseSangerTraceRequired(TypedDict):
    file_base64: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class ParseSangerTrace(ParseSangerTraceRequired, total=False):
    file_name: str
    gate: Any


class ParseSangerTraceCreateDataRequired(TypedDict):
    file_base64: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class ParseSangerTraceCreateData(ParseSangerTraceCreateDataRequired, total=False):
    file_name: str
    gate: Any


class PlasmidAnnotateRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class PlasmidAnnotate(PlasmidAnnotateRequired, total=False):
    gate: Any


class PlasmidAnnotateCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class PlasmidAnnotateCreateData(PlasmidAnnotateCreateDataRequired, total=False):
    gate: Any


class PlasmidDeepAnnotateRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class PlasmidDeepAnnotate(PlasmidDeepAnnotateRequired, total=False):
    circular: bool
    gate: Any


class PlasmidDeepAnnotateCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class PlasmidDeepAnnotateCreateData(PlasmidDeepAnnotateCreateDataRequired, total=False):
    circular: bool
    gate: Any


class PlasmidFullReportRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class PlasmidFullReport(PlasmidFullReportRequired, total=False):
    circular: bool
    gate: Any
    top_n: int


class PlasmidFullReportCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class PlasmidFullReportCreateData(PlasmidFullReportCreateDataRequired, total=False):
    circular: bool
    gate: Any
    top_n: int


class PlasmidIdentifyRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class PlasmidIdentify(PlasmidIdentifyRequired, total=False):
    circular: bool
    gate: Any
    top_n: int


class PlasmidIdentifyCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class PlasmidIdentifyCreateData(PlasmidIdentifyCreateDataRequired, total=False):
    circular: bool
    gate: Any
    top_n: int


class PrimeEditingDesignRequired(TypedDict):
    edit_end: int
    edit_start: int
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class PrimeEditingDesign(PrimeEditingDesignRequired, total=False):
    frame_start: int
    gate: Any
    inserted_seq: str
    pbs_length: int
    rtt_homology: int


class PrimeEditingDesignCreateDataRequired(TypedDict):
    edit_end: int
    edit_start: int
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class PrimeEditingDesignCreateData(PrimeEditingDesignCreateDataRequired, total=False):
    frame_start: int
    gate: Any
    inserted_seq: str
    pbs_length: int
    rtt_homology: int


class PrimeEditingTwinDesignRequired(TypedDict):
    new_sequence: str
    ok: Any
    provenance: dict
    replace_end: int
    replace_start: int
    result: dict
    target: str
    tool: str


class PrimeEditingTwinDesign(PrimeEditingTwinDesignRequired, total=False):
    gate: Any
    overlap_length: int
    pbs_length: int


class PrimeEditingTwinDesignCreateDataRequired(TypedDict):
    new_sequence: str
    ok: Any
    provenance: dict
    replace_end: int
    replace_start: int
    result: dict
    target: str
    tool: str


class PrimeEditingTwinDesignCreateData(PrimeEditingTwinDesignCreateDataRequired, total=False):
    gate: Any
    overlap_length: int
    pbs_length: int


class PrimerDesignRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    template: str
    tool: str


class PrimerDesign(PrimerDesignRequired, total=False):
    amplicon_max: int
    amplicon_min: int
    dntp_mm: float
    gate: Any
    gc_max: float
    gc_min: float
    len_max: int
    len_min: int
    len_opt: int
    max_return: int
    mg_mm: float
    na_mm: float
    oligo_nm: float
    target_end: int
    target_start: int
    tm_max: float
    tm_max_diff: float
    tm_min: float
    tm_opt: float


class PrimerDesignCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    template: str
    tool: str


class PrimerDesignCreateData(PrimerDesignCreateDataRequired, total=False):
    amplicon_max: int
    amplicon_min: int
    dntp_mm: float
    gate: Any
    gc_max: float
    gc_min: float
    len_max: int
    len_min: int
    len_opt: int
    max_return: int
    mg_mm: float
    na_mm: float
    oligo_nm: float
    target_end: int
    target_start: int
    tm_max: float
    tm_max_diff: float
    tm_min: float
    tm_opt: float


class PrimerSpecificityRequired(TypedDict):
    forward_primer: str
    ok: Any
    provenance: dict
    result: dict
    reverse_primer: str
    tool: str


class PrimerSpecificity(PrimerSpecificityRequired, total=False):
    gate: Any
    max_mismatch: int
    max_product_length: int


class PrimerSpecificityCreateDataRequired(TypedDict):
    forward_primer: str
    ok: Any
    provenance: dict
    result: dict
    reverse_primer: str
    tool: str


class PrimerSpecificityCreateData(PrimerSpecificityCreateDataRequired, total=False):
    gate: Any
    max_mismatch: int
    max_product_length: int


class ProteaseDigestionRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteaseDigestion(ProteaseDigestionRequired, total=False):
    gate: Any
    max_mass: float
    max_peptide: int
    min_mass: float
    missed_cleavage: int
    protease: str


class ProteaseDigestionCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteaseDigestionCreateData(ProteaseDigestionCreateDataRequired, total=False):
    gate: Any
    max_mass: float
    max_peptide: int
    min_mass: float
    missed_cleavage: int
    protease: str


class ProteinAnnotatePollRequired(TypedDict):
    job_id: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class ProteinAnnotatePoll(ProteinAnnotatePollRequired, total=False):
    gate: Any


class ProteinAnnotatePollCreateDataRequired(TypedDict):
    job_id: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class ProteinAnnotatePollCreateData(ProteinAnnotatePollCreateDataRequired, total=False):
    gate: Any


class ProteinAnnotateSubmitRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteinAnnotateSubmit(ProteinAnnotateSubmitRequired, total=False):
    appl: str
    gate: Any
    goterm: bool


class ProteinAnnotateSubmitCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteinAnnotateSubmitCreateData(ProteinAnnotateSubmitCreateDataRequired, total=False):
    appl: str
    gate: Any
    goterm: bool


class ProteinHydrophobicityRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteinHydrophobicity(ProteinHydrophobicityRequired, total=False):
    gate: Any
    scale: str
    window: int


class ProteinHydrophobicityCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteinHydrophobicityCreateData(ProteinHydrophobicityCreateDataRequired, total=False):
    gate: Any
    scale: str
    window: int


class ProteinPropertyRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteinProperty(ProteinPropertyRequired, total=False):
    charge_step: float
    gate: Any


class ProteinPropertyCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteinPropertyCreateData(ProteinPropertyCreateDataRequired, total=False):
    charge_step: float
    gate: Any


class RandomSequenceRequired(TypedDict):
    length: int
    ok: Any
    provenance: dict
    result: dict
    tool: str


class RandomSequence(RandomSequenceRequired, total=False):
    gate: Any
    gc_content: float
    kind: str


class RandomSequenceCreateDataRequired(TypedDict):
    length: int
    ok: Any
    provenance: dict
    result: dict
    tool: str


class RandomSequenceCreateData(RandomSequenceCreateDataRequired, total=False):
    gate: Any
    gc_content: float
    kind: str


class RestrictionSiteRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class RestrictionSite(RestrictionSiteRequired, total=False):
    enzyme: list
    gate: Any


class RestrictionSiteCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class RestrictionSiteCreateData(RestrictionSiteCreateDataRequired, total=False):
    enzyme: list
    gate: Any


class ReverseComplementRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ReverseComplement(ReverseComplementRequired, total=False):
    gate: Any
    type: str


class ReverseComplementCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ReverseComplementCreateData(ReverseComplementCreateDataRequired, total=False):
    gate: Any
    type: str


class ReverseTranslateRequired(TypedDict):
    ok: Any
    protein: str
    provenance: dict
    result: dict
    tool: str


class ReverseTranslate(ReverseTranslateRequired, total=False):
    gate: Any
    mode: str
    organism: str


class ReverseTranslateCreateDataRequired(TypedDict):
    ok: Any
    protein: str
    provenance: dict
    result: dict
    tool: str


class ReverseTranslateCreateData(ReverseTranslateCreateDataRequired, total=False):
    gate: Any
    mode: str
    organism: str


class RnaFoldRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class RnaFold(RnaFoldRequired, total=False):
    gate: Any


class RnaFoldCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class RnaFoldCreateData(RnaFoldCreateDataRequired, total=False):
    gate: Any


class SangerVsReferenceRequired(TypedDict):
    ok: Any
    provenance: dict
    reference: str
    result: dict
    tool: str


class SangerVsReference(SangerVsReferenceRequired, total=False):
    file_base64: str
    file_name: str
    gate: Any
    min_coverage: float
    read: str


class SangerVsReferenceCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    reference: str
    result: dict
    tool: str


class SangerVsReferenceCreateData(SangerVsReferenceCreateDataRequired, total=False):
    file_base64: str
    file_name: str
    gate: Any
    min_coverage: float
    read: str


class SavePermalinkRequired(TypedDict):
    arg: dict
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SavePermalink(SavePermalinkRequired, total=False):
    gate: Any


class SavePermalinkCreateDataRequired(TypedDict):
    arg: dict
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SavePermalinkCreateData(SavePermalinkCreateDataRequired, total=False):
    gate: Any


class SeqfileStatRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SeqfileStat(SeqfileStatRequired, total=False):
    gate: Any
    quality_offset: int


class SeqfileStatCreateDataRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SeqfileStatCreateData(SeqfileStatCreateDataRequired, total=False):
    gate: Any
    quality_offset: int


class SequenceFetchRequired(TypedDict):
    accession: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SequenceFetch(SequenceFetchRequired, total=False):
    db: str
    format: str
    gate: Any


class SequenceFetchCreateDataRequired(TypedDict):
    accession: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SequenceFetchCreateData(SequenceFetchCreateDataRequired, total=False):
    db: str
    format: str
    gate: Any


class SequenceFormatConvertRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SequenceFormatConvert(SequenceFormatConvertRequired, total=False):
    gate: Any
    to: str


class SequenceFormatConvertCreateDataRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SequenceFormatConvertCreateData(SequenceFormatConvertCreateDataRequired, total=False):
    gate: Any
    to: str


class SequenceReportRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class SequenceReport(SequenceReportRequired, total=False):
    end_primer_length: int
    gate: Any
    max_orf: int
    min_orf_aa: int


class SequenceReportCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class SequenceReportCreateData(SequenceReportCreateDataRequired, total=False):
    end_primer_length: int
    gate: Any
    max_orf: int
    min_orf_aa: int


class SequenceSearchRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SequenceSearch(SequenceSearchRequired, total=False):
    db: str
    gate: Any
    gene: str
    max_result: int
    organism: str
    term: str


class SequenceSearchCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SequenceSearchCreateData(SequenceSearchCreateDataRequired, total=False):
    db: str
    gate: Any
    gene: str
    max_result: int
    organism: str
    term: str


class SequencingReadbackVerifyRequired(TypedDict):
    ok: Any
    provenance: dict
    read: str
    reference: str
    result: dict
    tool: str


class SequencingReadbackVerify(SequencingReadbackVerifyRequired, total=False):
    gate: Any
    min_supporting_read: int


class SequencingReadbackVerifyCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    read: str
    reference: str
    result: dict
    tool: str


class SequencingReadbackVerifyCreateData(SequencingReadbackVerifyCreateDataRequired, total=False):
    gate: Any
    min_supporting_read: int


class SessionCreateRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SessionCreate(SessionCreateRequired, total=False):
    entry: dict
    gate: Any


class SessionCreateCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SessionCreateCreateData(SessionCreateCreateDataRequired, total=False):
    entry: dict
    gate: Any


class SessionGetRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    session_id: str
    tool: str


class SessionGet(SessionGetRequired, total=False):
    gate: Any
    name: list


class SessionGetCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    session_id: str
    tool: str


class SessionGetCreateData(SessionGetCreateDataRequired, total=False):
    gate: Any
    name: list


class SessionRunRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    session_id: str
    tool: str


class SessionRun(SessionRunRequired, total=False):
    arg: dict
    from_session: dict
    gate: Any
    write_back: dict


class SessionRunCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    session_id: str
    tool: str


class SessionRunCreateData(SessionRunCreateDataRequired, total=False):
    arg: dict
    from_session: dict
    gate: Any
    write_back: dict


class SessionSetRequired(TypedDict):
    entry: dict
    ok: Any
    provenance: dict
    result: dict
    session_id: str
    tool: str


class SessionSet(SessionSetRequired, total=False):
    gate: Any


class SessionSetCreateDataRequired(TypedDict):
    entry: dict
    ok: Any
    provenance: dict
    result: dict
    session_id: str
    tool: str


class SessionSetCreateData(SessionSetCreateDataRequired, total=False):
    gate: Any


class SirnaDesignRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class SirnaDesign(SirnaDesignRequired, total=False):
    gate: Any
    min_reynold: int
    sh_rna_loop: str


class SirnaDesignCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class SirnaDesignCreateData(SirnaDesignCreateDataRequired, total=False):
    gate: Any
    min_reynold: int
    sh_rna_loop: str


class SiteDirectedMutagenesiRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    template: str
    tool: str


class SiteDirectedMutagenesi(SiteDirectedMutagenesiRequired, total=False):
    arm_tm_target: float
    dntp_mm: float
    edit_kind: str
    frame_start: int
    gate: Any
    mg_mm: float
    na_mm: float
    new_base: str
    oligo_nm: float
    organism: str
    position: int
    residue: int
    style: str
    target_aa: str


class SiteDirectedMutagenesiCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    template: str
    tool: str


class SiteDirectedMutagenesiCreateData(SiteDirectedMutagenesiCreateDataRequired, total=False):
    arm_tm_target: float
    dntp_mm: float
    edit_kind: str
    frame_start: int
    gate: Any
    mg_mm: float
    na_mm: float
    new_base: str
    oligo_nm: float
    organism: str
    position: int
    residue: int
    style: str
    target_aa: str


class TranslateRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class Translate(TranslateRequired, total=False):
    frame: int
    gate: Any
    to_stop: bool


class TranslateCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class TranslateCreateData(TranslateCreateDataRequired, total=False):
    frame: int
    gate: Any
    to_stop: bool


class VariantAnnotateRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str
    variant: str


class VariantAnnotate(VariantAnnotateRequired, total=False):
    assembly: str
    gate: Any


class VariantAnnotateCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str
    variant: str


class VariantAnnotateCreateData(VariantAnnotateCreateDataRequired, total=False):
    assembly: str
    gate: Any


class VariantComparatorRequired(TypedDict):
    ok: Any
    provenance: dict
    query: str
    reference: str
    result: dict
    tool: str


class VariantComparator(VariantComparatorRequired, total=False):
    coding: bool
    frame_start: int
    gate: Any


class VariantComparatorCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    query: str
    reference: str
    result: dict
    tool: str


class VariantComparatorCreateData(VariantComparatorCreateDataRequired, total=False):
    coding: bool
    frame_start: int
    gate: Any


class VerifyAssemblyRequired(TypedDict):
    claimed_construct: str
    method: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class VerifyAssembly(VerifyAssemblyRequired, total=False):
    arm_tm_target: float
    circular: bool
    coding: bool
    enzyme: str
    enzyme3: str
    enzyme5: str
    fragment: list
    fragment_pcr: list
    frame_start: int
    gate: Any
    insert: str
    insert_pcr: dict
    name: list
    overlap_len: int
    vector: str
    vector_pcr: dict


class VerifyAssemblyCreateDataRequired(TypedDict):
    claimed_construct: str
    method: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class VerifyAssemblyCreateData(VerifyAssemblyCreateDataRequired, total=False):
    arm_tm_target: float
    circular: bool
    coding: bool
    enzyme: str
    enzyme3: str
    enzyme5: str
    fragment: list
    fragment_pcr: list
    frame_start: int
    gate: Any
    insert: str
    insert_pcr: dict
    name: list
    overlap_len: int
    vector: str
    vector_pcr: dict


class VerifyConstructRequired(TypedDict):
    claimed_construct: str
    insert_forward_primer: str
    insert_reverse_primer: str
    insert_template: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class VerifyConstruct(VerifyConstructRequired, total=False):
    expected_frame_start: int
    gate: Any
    max_primer_mismatch: int
    template_circular: bool


class VerifyConstructCreateDataRequired(TypedDict):
    claimed_construct: str
    insert_forward_primer: str
    insert_reverse_primer: str
    insert_template: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class VerifyConstructCreateData(VerifyConstructCreateDataRequired, total=False):
    expected_frame_start: int
    gate: Any
    max_primer_mismatch: int
    template_circular: bool


class VirtualGelRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class VirtualGel(VirtualGelRequired, total=False):
    circular: bool
    enzyme: list
    gate: Any
    ladder: str


class VirtualGelCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class VirtualGelCreateData(VirtualGelCreateDataRequired, total=False):
    circular: bool
    enzyme: list
    gate: Any
    ladder: str


class VolcanoPlotDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    row: list
    tool: str


class VolcanoPlotData(VolcanoPlotDataRequired, total=False):
    gate: Any


class VolcanoPlotDataCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    row: list
    tool: str


class VolcanoPlotDataCreateData(VolcanoPlotDataCreateDataRequired, total=False):
    gate: Any


class WebSearchRequired(TypedDict):
    ok: Any
    provenance: dict
    query: str
    result: dict
    tool: str


class WebSearch(WebSearchRequired, total=False):
    gate: Any
    max_result: float


class WebSearchCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    query: str
    result: dict
    tool: str


class WebSearchCreateData(WebSearchCreateDataRequired, total=False):
    gate: Any
    max_result: float
