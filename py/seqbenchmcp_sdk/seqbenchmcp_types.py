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
    frameStart: int
    gate: Any
    targetPosition: int


class BaseEditingDesignCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class BaseEditingDesignCreateData(BaseEditingDesignCreateDataRequired, total=False):
    editor: str
    frameStart: int
    gate: Any
    targetPosition: int


class BatchRequired(TypedDict):
    capped: bool
    columns: list
    count: int
    errors: int
    input: str
    limit: int
    provenance: dict
    rows: list
    tool: str


class Batch(BatchRequired, total=False):
    args: dict


class BatchLoadMatch(TypedDict, total=False):
    args: dict
    capped: bool
    columns: list
    count: int
    errors: int
    input: str
    limit: int
    provenance: dict
    rows: list
    tool: str


class BatchCreateDataRequired(TypedDict):
    capped: bool
    columns: list
    count: int
    errors: int
    input: str
    limit: int
    provenance: dict
    rows: list
    tool: str


class BatchCreateData(BatchCreateDataRequired, total=False):
    args: dict


class BatchWorkflow(TypedDict):
    capped: bool
    columns: list
    count: int
    errors: int
    input: str
    limit: int
    provenance: dict
    rows: list
    steps: list


class BatchWorkflowLoadMatch(TypedDict, total=False):
    capped: bool
    columns: list
    count: int
    errors: int
    input: str
    limit: int
    provenance: dict
    rows: list
    steps: list


class BatchWorkflowCreateData(TypedDict):
    capped: bool
    columns: list
    count: int
    errors: int
    input: str
    limit: int
    provenance: dict
    rows: list
    steps: list


class CharacterizeSequenceRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CharacterizeSequence(CharacterizeSequenceRequired, total=False):
    endPrimerLength: int
    gate: Any
    maxOrfs: int
    minOrfAa: int


class CharacterizeSequenceCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CharacterizeSequenceCreateData(CharacterizeSequenceCreateDataRequired, total=False):
    endPrimerLength: int
    gate: Any
    maxOrfs: int
    minOrfAa: int


class CloningSimulateRequired(TypedDict):
    method: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class CloningSimulate(CloningSimulateRequired, total=False):
    armTmTarget: float
    circular: bool
    enzyme: str
    enzyme3: str
    enzyme5: str
    fragments: list
    gate: Any
    insert: str
    names: list
    overlapLen: int
    vector: str


class CloningSimulateCreateDataRequired(TypedDict):
    method: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class CloningSimulateCreateData(CloningSimulateCreateDataRequired, total=False):
    armTmTarget: float
    circular: bool
    enzyme: str
    enzyme3: str
    enzyme5: str
    fragments: list
    gate: Any
    insert: str
    names: list
    overlapLen: int
    vector: str


class CodonAdaptationIndexRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CodonAdaptationIndex(CodonAdaptationIndexRequired, total=False):
    frameStart: int
    gate: Any
    organism: str
    rareThreshold: float


class CodonAdaptationIndexCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CodonAdaptationIndexCreateData(CodonAdaptationIndexCreateDataRequired, total=False):
    frameStart: int
    gate: Any
    organism: str
    rareThreshold: float


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
    avoidEnzymes: list
    crypticOrfMinAa: int
    frameStart: int
    gate: Any
    gcHigh: float
    gcLow: float
    gcWindow: int
    homopolymerMin: int
    maxPasses: int
    organism: str


class ConstructAutofixCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ConstructAutofixCreateData(ConstructAutofixCreateDataRequired, total=False):
    avoidEnzymes: list
    crypticOrfMinAa: int
    frameStart: int
    gate: Any
    gcHigh: float
    gcLow: float
    gcWindow: int
    homopolymerMin: int
    maxPasses: int
    organism: str


class ConstructQcRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ConstructQc(ConstructQcRequired, total=False):
    avoidEnzymes: list
    crypticOrfMinAa: int
    frameStart: int
    gate: Any
    gcHigh: float
    gcLow: float
    gcWindow: int
    homopolymerMin: int


class ConstructQcCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ConstructQcCreateData(ConstructQcCreateDataRequired, total=False):
    avoidEnzymes: list
    crypticOrfMinAa: int
    frameStart: int
    gate: Any
    gcHigh: float
    gcLow: float
    gcWindow: int
    homopolymerMin: int


class CrisprGrnaDesignRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CrisprGrnaDesign(CrisprGrnaDesignRequired, total=False):
    gate: Any
    minScore: float
    nuclease: str
    searchReverseStrand: bool


class CrisprGrnaDesignCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class CrisprGrnaDesignCreateData(CrisprGrnaDesignCreateDataRequired, total=False):
    gate: Any
    minScore: float
    nuclease: str
    searchReverseStrand: bool


class CrisprHdrDonorRequired(TypedDict):
    ok: Any
    provenance: dict
    replacement: str
    result: dict
    targetSequence: str
    tool: str


class CrisprHdrDonor(CrisprHdrDonorRequired, total=False):
    armLength: int
    blockPam: bool
    designGenotypingPrimers: bool
    editEnd: int
    editStart: int
    frameStart: int
    gate: Any
    guideEnd: int
    guideStart: int
    guideStrand: str
    nuclease: str


class CrisprHdrDonorCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    replacement: str
    result: dict
    targetSequence: str
    tool: str


class CrisprHdrDonorCreateData(CrisprHdrDonorCreateDataRequired, total=False):
    armLength: int
    blockPam: bool
    designGenotypingPrimers: bool
    editEnd: int
    editStart: int
    frameStart: int
    gate: Any
    guideEnd: int
    guideStart: int
    guideStrand: str
    nuclease: str


class CrisprOfftargetCheckRequired(TypedDict):
    ok: Any
    protospacer: str
    provenance: dict
    result: dict
    tool: str


class CrisprOfftargetCheck(CrisprOfftargetCheckRequired, total=False):
    gate: Any
    maxMismatches: int
    nuclease: str


class CrisprOfftargetCheckCreateDataRequired(TypedDict):
    ok: Any
    protospacer: str
    provenance: dict
    result: dict
    tool: str


class CrisprOfftargetCheckCreateData(CrisprOfftargetCheckCreateDataRequired, total=False):
    gate: Any
    maxMismatches: int
    nuclease: str


class CrossDimerRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequenceA: str
    sequenceB: str
    tool: str


class CrossDimer(CrossDimerRequired, total=False):
    gate: Any


class CrossDimerCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequenceA: str
    sequenceB: str
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
    massNg: float
    sequence: str
    type: str
    volumeUl: float


class DnaMolarityCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str


class DnaMolarityCreateData(DnaMolarityCreateDataRequired, total=False):
    gate: Any
    length: int
    massNg: float
    sequence: str
    type: str
    volumeUl: float


class DoubleDigestRequired(TypedDict):
    enzymeA: str
    enzymeB: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class DoubleDigest(DoubleDigestRequired, total=False):
    gate: Any


class DoubleDigestCreateDataRequired(TypedDict):
    enzymeA: str
    enzymeB: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class DoubleDigestCreateData(DoubleDigestCreateDataRequired, total=False):
    gate: Any


class ExportEchoPicklistRequired(TypedDict):
    ok: Any
    provenance: dict
    reactions: list
    result: dict
    tool: str


class ExportEchoPicklist(ExportEchoPicklistRequired, total=False):
    gate: Any


class ExportEchoPicklistCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    reactions: list
    result: dict
    tool: str


class ExportEchoPicklistCreateData(ExportEchoPicklistCreateDataRequired, total=False):
    gate: Any


class ExportOpentronsProtocolRequired(TypedDict):
    ok: Any
    provenance: dict
    reactions: list
    result: dict
    tool: str


class ExportOpentronsProtocol(ExportOpentronsProtocolRequired, total=False):
    gate: Any
    protocolName: str


class ExportOpentronsProtocolCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    reactions: list
    result: dict
    tool: str


class ExportOpentronsProtocolCreateData(ExportOpentronsProtocolCreateDataRequired, total=False):
    gate: Any
    protocolName: str


class ExportPlateLayoutRequired(TypedDict):
    ok: Any
    provenance: dict
    reactions: list
    result: dict
    tool: str


class ExportPlateLayout(ExportPlateLayoutRequired, total=False):
    gate: Any


class ExportPlateLayoutCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    reactions: list
    result: dict
    tool: str


class ExportPlateLayoutCreateData(ExportPlateLayoutCreateDataRequired, total=False):
    gate: Any


class ExpressionHeatmapClusterRequired(TypedDict):
    genes: list
    ok: Any
    provenance: dict
    result: dict
    samples: list
    tool: str
    values: list


class ExpressionHeatmapCluster(ExpressionHeatmapClusterRequired, total=False):
    clusterCols: bool
    clusterRows: bool
    distanceMetric: str
    gate: Any
    linkage: str
    zScoreRows: bool


class ExpressionHeatmapClusterCreateDataRequired(TypedDict):
    genes: list
    ok: Any
    provenance: dict
    result: dict
    samples: list
    tool: str
    values: list


class ExpressionHeatmapClusterCreateData(ExpressionHeatmapClusterCreateDataRequired, total=False):
    clusterCols: bool
    clusterRows: bool
    distanceMetric: str
    gate: Any
    linkage: str
    zScoreRows: bool


class FastqQcReportRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FastqQcReport(FastqQcReportRequired, total=False):
    gate: Any
    qualityOffset: int


class FastqQcReportCreateDataRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FastqQcReportCreateData(FastqQcReportCreateDataRequired, total=False):
    gate: Any
    qualityOffset: int


class FastqTrimRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FastqTrim(FastqTrimRequired, total=False):
    gate: Any
    minLength: int
    qualityOffset: int
    qualityThreshold: int


class FastqTrimCreateDataRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FastqTrimCreateData(FastqTrimCreateDataRequired, total=False):
    gate: Any
    minLength: int
    qualityOffset: int
    qualityThreshold: int


class FindOrfRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class FindOrf(FindOrfRequired, total=False):
    gate: Any
    minAaLength: int
    requireStop: bool


class FindOrfCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class FindOrfCreateData(FindOrfCreateDataRequired, total=False):
    gate: Any
    minAaLength: int
    requireStop: bool


class FormatSequenceRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class FormatSequence(FormatSequenceRequired, total=False):
    caseMode: str
    convert: str
    gate: Any
    reverse: bool
    stripNonLetters: bool
    width: int


class FormatSequenceCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class FormatSequenceCreateData(FormatSequenceCreateDataRequired, total=False):
    caseMode: str
    convert: str
    gate: Any
    reverse: bool
    stripNonLetters: bool
    width: int


class FunctionalEnrichmentRequired(TypedDict):
    genes: list
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FunctionalEnrichment(FunctionalEnrichmentRequired, total=False):
    background: list
    collections: list
    gate: Any
    maxTermSize: int
    minTermSize: int


class FunctionalEnrichmentCreateDataRequired(TypedDict):
    genes: list
    ok: Any
    provenance: dict
    result: dict
    tool: str


class FunctionalEnrichmentCreateData(FunctionalEnrichmentCreateDataRequired, total=False):
    background: list
    collections: list
    gate: Any
    maxTermSize: int
    minTermSize: int


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
    overhangs: list
    provenance: dict
    result: dict
    tool: str


class GoldenGateFidelity(GoldenGateFidelityRequired, total=False):
    compareToNamedSet: str
    dataset: str
    gate: Any
    riskThreshold: float


class GoldenGateFidelityCreateDataRequired(TypedDict):
    ok: Any
    overhangs: list
    provenance: dict
    result: dict
    tool: str


class GoldenGateFidelityCreateData(GoldenGateFidelityCreateDataRequired, total=False):
    compareToNamedSet: str
    dataset: str
    gate: Any
    riskThreshold: float


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
    jobId: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class IdMapPoll(IdMapPollRequired, total=False):
    gate: Any


class IdMapPollCreateDataRequired(TypedDict):
    jobId: str
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
    taxId: str


class IdMapSubmitCreateDataRequired(TypedDict):
    ids: list
    ok: Any
    provenance: dict
    result: dict
    to: str
    tool: str


class IdMapSubmitCreateData(IdMapSubmitCreateDataRequired, total=False):
    gate: Any
    taxId: str


class InSilicoPcrRequired(TypedDict):
    forwardPrimer: str
    ok: Any
    provenance: dict
    result: dict
    reversePrimer: str
    template: str
    tool: str


class InSilicoPcr(InSilicoPcrRequired, total=False):
    circular: bool
    gate: Any
    maxMismatches: int


class InSilicoPcrCreateDataRequired(TypedDict):
    forwardPrimer: str
    ok: Any
    provenance: dict
    result: dict
    reversePrimer: str
    template: str
    tool: str


class InSilicoPcrCreateData(InSilicoPcrCreateDataRequired, total=False):
    circular: bool
    gate: Any
    maxMismatches: int


class KaspPrimerDesignRequired(TypedDict):
    alleleA: str
    alleleB: str
    ok: Any
    provenance: dict
    result: dict
    snpPosition: int
    target: str
    tool: str


class KaspPrimerDesign(KaspPrimerDesignRequired, total=False):
    addSecondaryMismatch: bool
    gate: Any
    maxAmplicon: int
    minAmplicon: int
    targetCoreTm: float


class KaspPrimerDesignCreateDataRequired(TypedDict):
    alleleA: str
    alleleB: str
    ok: Any
    provenance: dict
    result: dict
    snpPosition: int
    target: str
    tool: str


class KaspPrimerDesignCreateData(KaspPrimerDesignCreateDataRequired, total=False):
    addSecondaryMismatch: bool
    gate: Any
    maxAmplicon: int
    minAmplicon: int
    targetCoreTm: float


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
    dntpMM: float
    gate: Any
    mgMM: float
    naMM: float
    oligoNM: float
    targetTm: float
    tmTolerance: float


class MeltingTemperatureCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class MeltingTemperatureCreateData(MeltingTemperatureCreateDataRequired, total=False):
    dntpMM: float
    gate: Any
    mgMM: float
    naMM: float
    oligoNM: float
    targetTm: float
    tmTolerance: float


class MotifFinderRequired(TypedDict):
    motif: str
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class MotifFinder(MotifFinderRequired, total=False):
    gate: Any
    maxMismatches: int
    searchReverseStrand: bool


class MotifFinderCreateDataRequired(TypedDict):
    motif: str
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class MotifFinderCreateData(MotifFinderCreateDataRequired, total=False):
    gate: Any
    maxMismatches: int
    searchReverseStrand: bool


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
    dntpMM: float
    gate: Any
    mgMM: float
    naMM: float
    oligoNM: float


class OligoAnalysiCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class OligoAnalysiCreateData(OligoAnalysiCreateDataRequired, total=False):
    dntpMM: float
    gate: Any
    mgMM: float
    naMM: float
    oligoNM: float


class OrthologMapRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    symbols: list
    targetSpecies: str
    tool: str


class OrthologMap(OrthologMapRequired, total=False):
    gate: Any
    sourceSpecies: str
    type: str


class OrthologMapCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    symbols: list
    targetSpecies: str
    tool: str


class OrthologMapCreateData(OrthologMapCreateDataRequired, total=False):
    gate: Any
    sourceSpecies: str
    type: str


class PairwiseAlignmentRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    seqA: str
    seqB: str
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
    seqA: str
    seqB: str
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
    fileBase64: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class ParseSangerTrace(ParseSangerTraceRequired, total=False):
    fileName: str
    gate: Any


class ParseSangerTraceCreateDataRequired(TypedDict):
    fileBase64: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class ParseSangerTraceCreateData(ParseSangerTraceCreateDataRequired, total=False):
    fileName: str
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
    topN: int


class PlasmidFullReportCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class PlasmidFullReportCreateData(PlasmidFullReportCreateDataRequired, total=False):
    circular: bool
    gate: Any
    topN: int


class PlasmidIdentifyRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class PlasmidIdentify(PlasmidIdentifyRequired, total=False):
    circular: bool
    gate: Any
    topN: int


class PlasmidIdentifyCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class PlasmidIdentifyCreateData(PlasmidIdentifyCreateDataRequired, total=False):
    circular: bool
    gate: Any
    topN: int


class PrimeEditingDesignRequired(TypedDict):
    editEnd: int
    editStart: int
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class PrimeEditingDesign(PrimeEditingDesignRequired, total=False):
    frameStart: int
    gate: Any
    insertedSeq: str
    pbsLength: int
    rttHomology: int


class PrimeEditingDesignCreateDataRequired(TypedDict):
    editEnd: int
    editStart: int
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class PrimeEditingDesignCreateData(PrimeEditingDesignCreateDataRequired, total=False):
    frameStart: int
    gate: Any
    insertedSeq: str
    pbsLength: int
    rttHomology: int


class PrimeEditingTwinDesignRequired(TypedDict):
    newSequence: str
    ok: Any
    provenance: dict
    replaceEnd: int
    replaceStart: int
    result: dict
    target: str
    tool: str


class PrimeEditingTwinDesign(PrimeEditingTwinDesignRequired, total=False):
    gate: Any
    overlapLength: int
    pbsLength: int


class PrimeEditingTwinDesignCreateDataRequired(TypedDict):
    newSequence: str
    ok: Any
    provenance: dict
    replaceEnd: int
    replaceStart: int
    result: dict
    target: str
    tool: str


class PrimeEditingTwinDesignCreateData(PrimeEditingTwinDesignCreateDataRequired, total=False):
    gate: Any
    overlapLength: int
    pbsLength: int


class PrimerDesignRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    template: str
    tool: str


class PrimerDesign(PrimerDesignRequired, total=False):
    ampliconMax: int
    ampliconMin: int
    dntpMM: float
    gate: Any
    gcMax: float
    gcMin: float
    lenMax: int
    lenMin: int
    lenOpt: int
    maxReturn: int
    mgMM: float
    naMM: float
    oligoNM: float
    targetEnd: int
    targetStart: int
    tmMax: float
    tmMaxDiff: float
    tmMin: float
    tmOpt: float


class PrimerDesignCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    template: str
    tool: str


class PrimerDesignCreateData(PrimerDesignCreateDataRequired, total=False):
    ampliconMax: int
    ampliconMin: int
    dntpMM: float
    gate: Any
    gcMax: float
    gcMin: float
    lenMax: int
    lenMin: int
    lenOpt: int
    maxReturn: int
    mgMM: float
    naMM: float
    oligoNM: float
    targetEnd: int
    targetStart: int
    tmMax: float
    tmMaxDiff: float
    tmMin: float
    tmOpt: float


class PrimerSpecificityRequired(TypedDict):
    forwardPrimer: str
    ok: Any
    provenance: dict
    result: dict
    reversePrimer: str
    tool: str


class PrimerSpecificity(PrimerSpecificityRequired, total=False):
    gate: Any
    maxMismatches: int
    maxProductLength: int


class PrimerSpecificityCreateDataRequired(TypedDict):
    forwardPrimer: str
    ok: Any
    provenance: dict
    result: dict
    reversePrimer: str
    tool: str


class PrimerSpecificityCreateData(PrimerSpecificityCreateDataRequired, total=False):
    gate: Any
    maxMismatches: int
    maxProductLength: int


class ProteaseDigestionRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteaseDigestion(ProteaseDigestionRequired, total=False):
    gate: Any
    maxMass: float
    maxPeptides: int
    minMass: float
    missedCleavages: int
    protease: str


class ProteaseDigestionCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteaseDigestionCreateData(ProteaseDigestionCreateDataRequired, total=False):
    gate: Any
    maxMass: float
    maxPeptides: int
    minMass: float
    missedCleavages: int
    protease: str


class ProteinAnnotatePollRequired(TypedDict):
    jobId: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class ProteinAnnotatePoll(ProteinAnnotatePollRequired, total=False):
    gate: Any


class ProteinAnnotatePollCreateDataRequired(TypedDict):
    jobId: str
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
    goterms: bool


class ProteinAnnotateSubmitCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteinAnnotateSubmitCreateData(ProteinAnnotateSubmitCreateDataRequired, total=False):
    appl: str
    gate: Any
    goterms: bool


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
    chargeStep: float
    gate: Any


class ProteinPropertyCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class ProteinPropertyCreateData(ProteinPropertyCreateDataRequired, total=False):
    chargeStep: float
    gate: Any


class RandomSequenceRequired(TypedDict):
    length: int
    ok: Any
    provenance: dict
    result: dict
    tool: str


class RandomSequence(RandomSequenceRequired, total=False):
    gate: Any
    gcContent: float
    kind: str


class RandomSequenceCreateDataRequired(TypedDict):
    length: int
    ok: Any
    provenance: dict
    result: dict
    tool: str


class RandomSequenceCreateData(RandomSequenceCreateDataRequired, total=False):
    gate: Any
    gcContent: float
    kind: str


class RestrictionSiteRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class RestrictionSite(RestrictionSiteRequired, total=False):
    enzymes: list
    gate: Any


class RestrictionSiteCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class RestrictionSiteCreateData(RestrictionSiteCreateDataRequired, total=False):
    enzymes: list
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
    fileBase64: str
    fileName: str
    gate: Any
    minCoverage: float
    read: str


class SangerVsReferenceCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    reference: str
    result: dict
    tool: str


class SangerVsReferenceCreateData(SangerVsReferenceCreateDataRequired, total=False):
    fileBase64: str
    fileName: str
    gate: Any
    minCoverage: float
    read: str


class SavePermalinkRequired(TypedDict):
    args: dict
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SavePermalink(SavePermalinkRequired, total=False):
    gate: Any


class SavePermalinkCreateDataRequired(TypedDict):
    args: dict
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
    qualityOffset: int


class SeqfileStatCreateDataRequired(TypedDict):
    input: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SeqfileStatCreateData(SeqfileStatCreateDataRequired, total=False):
    gate: Any
    qualityOffset: int


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
    endPrimerLength: int
    gate: Any
    maxOrfs: int
    minOrfAa: int


class SequenceReportCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class SequenceReportCreateData(SequenceReportCreateDataRequired, total=False):
    endPrimerLength: int
    gate: Any
    maxOrfs: int
    minOrfAa: int


class SequenceSearchRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SequenceSearch(SequenceSearchRequired, total=False):
    db: str
    gate: Any
    gene: str
    maxResults: int
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
    maxResults: int
    organism: str
    term: str


class SequencingReadbackVerifyRequired(TypedDict):
    ok: Any
    provenance: dict
    reads: str
    reference: str
    result: dict
    tool: str


class SequencingReadbackVerify(SequencingReadbackVerifyRequired, total=False):
    gate: Any
    minSupportingReads: int


class SequencingReadbackVerifyCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    reads: str
    reference: str
    result: dict
    tool: str


class SequencingReadbackVerifyCreateData(SequencingReadbackVerifyCreateDataRequired, total=False):
    gate: Any
    minSupportingReads: int


class SessionCreateRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SessionCreate(SessionCreateRequired, total=False):
    entries: dict
    gate: Any


class SessionCreateCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    tool: str


class SessionCreateCreateData(SessionCreateCreateDataRequired, total=False):
    entries: dict
    gate: Any


class SessionGetRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sessionId: str
    tool: str


class SessionGet(SessionGetRequired, total=False):
    gate: Any
    names: list


class SessionGetCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sessionId: str
    tool: str


class SessionGetCreateData(SessionGetCreateDataRequired, total=False):
    gate: Any
    names: list


class SessionRunRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sessionId: str
    tool: str


class SessionRun(SessionRunRequired, total=False):
    args: dict
    fromSession: dict
    gate: Any
    writeBack: dict


class SessionRunCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sessionId: str
    tool: str


class SessionRunCreateData(SessionRunCreateDataRequired, total=False):
    args: dict
    fromSession: dict
    gate: Any
    writeBack: dict


class SessionSetRequired(TypedDict):
    entries: dict
    ok: Any
    provenance: dict
    result: dict
    sessionId: str
    tool: str


class SessionSet(SessionSetRequired, total=False):
    gate: Any


class SessionSetCreateDataRequired(TypedDict):
    entries: dict
    ok: Any
    provenance: dict
    result: dict
    sessionId: str
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
    minReynolds: int
    shRnaLoop: str


class SirnaDesignCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    target: str
    tool: str


class SirnaDesignCreateData(SirnaDesignCreateDataRequired, total=False):
    gate: Any
    minReynolds: int
    shRnaLoop: str


class SiteDirectedMutagenesiRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    template: str
    tool: str


class SiteDirectedMutagenesi(SiteDirectedMutagenesiRequired, total=False):
    armTmTarget: float
    dntpMM: float
    editKind: str
    frameStart: int
    gate: Any
    mgMM: float
    naMM: float
    newBase: str
    oligoNM: float
    organism: str
    position: int
    residue: int
    style: str
    targetAa: str


class SiteDirectedMutagenesiCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    template: str
    tool: str


class SiteDirectedMutagenesiCreateData(SiteDirectedMutagenesiCreateDataRequired, total=False):
    armTmTarget: float
    dntpMM: float
    editKind: str
    frameStart: int
    gate: Any
    mgMM: float
    naMM: float
    newBase: str
    oligoNM: float
    organism: str
    position: int
    residue: int
    style: str
    targetAa: str


class TranslateRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class Translate(TranslateRequired, total=False):
    frame: int
    gate: Any
    toStop: bool


class TranslateCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class TranslateCreateData(TranslateCreateDataRequired, total=False):
    frame: int
    gate: Any
    toStop: bool


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
    frameStart: int
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
    frameStart: int
    gate: Any


class VerifyAssemblyRequired(TypedDict):
    claimedConstruct: str
    method: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class VerifyAssembly(VerifyAssemblyRequired, total=False):
    armTmTarget: float
    circular: bool
    coding: bool
    enzyme: str
    enzyme3: str
    enzyme5: str
    fragmentPcrs: list
    fragments: list
    frameStart: int
    gate: Any
    insert: str
    insertPcr: dict
    names: list
    overlapLen: int
    vector: str
    vectorPcr: dict


class VerifyAssemblyCreateDataRequired(TypedDict):
    claimedConstruct: str
    method: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class VerifyAssemblyCreateData(VerifyAssemblyCreateDataRequired, total=False):
    armTmTarget: float
    circular: bool
    coding: bool
    enzyme: str
    enzyme3: str
    enzyme5: str
    fragmentPcrs: list
    fragments: list
    frameStart: int
    gate: Any
    insert: str
    insertPcr: dict
    names: list
    overlapLen: int
    vector: str
    vectorPcr: dict


class VerifyConstructRequired(TypedDict):
    claimedConstruct: str
    insertForwardPrimer: str
    insertReversePrimer: str
    insertTemplate: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class VerifyConstruct(VerifyConstructRequired, total=False):
    expectedFrameStart: int
    gate: Any
    maxPrimerMismatches: int
    templateCircular: bool


class VerifyConstructCreateDataRequired(TypedDict):
    claimedConstruct: str
    insertForwardPrimer: str
    insertReversePrimer: str
    insertTemplate: str
    ok: Any
    provenance: dict
    result: dict
    tool: str


class VerifyConstructCreateData(VerifyConstructCreateDataRequired, total=False):
    expectedFrameStart: int
    gate: Any
    maxPrimerMismatches: int
    templateCircular: bool


class VirtualGelRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    sequence: str
    tool: str


class VirtualGel(VirtualGelRequired, total=False):
    circular: bool
    enzymes: list
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
    enzymes: list
    gate: Any
    ladder: str


class VolcanoPlotDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    rows: list
    tool: str


class VolcanoPlotData(VolcanoPlotDataRequired, total=False):
    gate: Any


class VolcanoPlotDataCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    result: dict
    rows: list
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
    max_results: float


class WebSearchCreateDataRequired(TypedDict):
    ok: Any
    provenance: dict
    query: str
    result: dict
    tool: str


class WebSearchCreateData(WebSearchCreateDataRequired, total=False):
    gate: Any
    max_results: float
