<?php
declare(strict_types=1);

// Typed models for the SeqbenchMcp SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** AlphafoldLookup entity data model. */
class AlphafoldLookup
{
    public string $accession;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for AlphafoldLookup#create. */
class AlphafoldLookupCreateData
{
    public string $accession;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** AsoDesign entity data model. */
class AsoDesign
{
    public mixed $gate = null;
    public ?int $length = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $target;
    public string $tool;
    public ?int $wing = null;
}

/** Request payload for AsoDesign#create. */
class AsoDesignCreateData
{
    public mixed $gate = null;
    public ?int $length = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $target;
    public string $tool;
    public ?int $wing = null;
}

/** BaseEditingDesign entity data model. */
class BaseEditingDesign
{
    public ?string $editor = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $target;
    public ?int $targetPosition = null;
    public string $tool;
}

/** Request payload for BaseEditingDesign#create. */
class BaseEditingDesignCreateData
{
    public ?string $editor = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $target;
    public ?int $targetPosition = null;
    public string $tool;
}

/** Batch entity data model. */
class Batch
{
    public ?array $args = null;
    public bool $capped;
    public array $columns;
    public int $count;
    public int $errors;
    public string $input;
    public int $limit;
    public array $provenance;
    public array $rows;
    public string $tool;
}

/** Request payload for Batch#load. */
class BatchLoadMatch
{
    public ?array $args = null;
    public ?bool $capped = null;
    public ?array $columns = null;
    public ?int $count = null;
    public ?int $errors = null;
    public ?string $input = null;
    public ?int $limit = null;
    public ?array $provenance = null;
    public ?array $rows = null;
    public ?string $tool = null;
}

/** Request payload for Batch#create. */
class BatchCreateData
{
    public ?array $args = null;
    public bool $capped;
    public array $columns;
    public int $count;
    public int $errors;
    public string $input;
    public int $limit;
    public array $provenance;
    public array $rows;
    public string $tool;
}

/** BatchWorkflow entity data model. */
class BatchWorkflow
{
    public bool $capped;
    public array $columns;
    public int $count;
    public int $errors;
    public string $input;
    public int $limit;
    public array $provenance;
    public array $rows;
    public array $steps;
}

/** Request payload for BatchWorkflow#load. */
class BatchWorkflowLoadMatch
{
    public ?bool $capped = null;
    public ?array $columns = null;
    public ?int $count = null;
    public ?int $errors = null;
    public ?string $input = null;
    public ?int $limit = null;
    public ?array $provenance = null;
    public ?array $rows = null;
    public ?array $steps = null;
}

/** Request payload for BatchWorkflow#create. */
class BatchWorkflowCreateData
{
    public bool $capped;
    public array $columns;
    public int $count;
    public int $errors;
    public string $input;
    public int $limit;
    public array $provenance;
    public array $rows;
    public array $steps;
}

/** CharacterizeSequence entity data model. */
class CharacterizeSequence
{
    public ?int $endPrimerLength = null;
    public mixed $gate = null;
    public ?int $maxOrfs = null;
    public ?int $minOrfAa = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for CharacterizeSequence#create. */
class CharacterizeSequenceCreateData
{
    public ?int $endPrimerLength = null;
    public mixed $gate = null;
    public ?int $maxOrfs = null;
    public ?int $minOrfAa = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** CloningSimulate entity data model. */
class CloningSimulate
{
    public ?float $armTmTarget = null;
    public ?bool $circular = null;
    public ?string $enzyme = null;
    public ?string $enzyme3 = null;
    public ?string $enzyme5 = null;
    public ?array $fragments = null;
    public mixed $gate = null;
    public ?string $insert = null;
    public string $method;
    public ?array $names = null;
    public mixed $ok;
    public ?int $overlapLen = null;
    public array $provenance;
    public array $result;
    public string $tool;
    public ?string $vector = null;
}

/** Request payload for CloningSimulate#create. */
class CloningSimulateCreateData
{
    public ?float $armTmTarget = null;
    public ?bool $circular = null;
    public ?string $enzyme = null;
    public ?string $enzyme3 = null;
    public ?string $enzyme5 = null;
    public ?array $fragments = null;
    public mixed $gate = null;
    public ?string $insert = null;
    public string $method;
    public ?array $names = null;
    public mixed $ok;
    public ?int $overlapLen = null;
    public array $provenance;
    public array $result;
    public string $tool;
    public ?string $vector = null;
}

/** CodonAdaptationIndex entity data model. */
class CodonAdaptationIndex
{
    public ?int $frameStart = null;
    public mixed $gate = null;
    public mixed $ok;
    public ?string $organism = null;
    public array $provenance;
    public ?float $rareThreshold = null;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for CodonAdaptationIndex#create. */
class CodonAdaptationIndexCreateData
{
    public ?int $frameStart = null;
    public mixed $gate = null;
    public mixed $ok;
    public ?string $organism = null;
    public array $provenance;
    public ?float $rareThreshold = null;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** CodonOptimize entity data model. */
class CodonOptimize
{
    public mixed $gate = null;
    public mixed $ok;
    public ?string $organism = null;
    public string $protein;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for CodonOptimize#create. */
class CodonOptimizeCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public ?string $organism = null;
    public string $protein;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** ConstructAutofix entity data model. */
class ConstructAutofix
{
    public ?array $avoidEnzymes = null;
    public ?int $crypticOrfMinAa = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?float $gcHigh = null;
    public ?float $gcLow = null;
    public ?int $gcWindow = null;
    public ?int $homopolymerMin = null;
    public ?int $maxPasses = null;
    public mixed $ok;
    public ?string $organism = null;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for ConstructAutofix#create. */
class ConstructAutofixCreateData
{
    public ?array $avoidEnzymes = null;
    public ?int $crypticOrfMinAa = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?float $gcHigh = null;
    public ?float $gcLow = null;
    public ?int $gcWindow = null;
    public ?int $homopolymerMin = null;
    public ?int $maxPasses = null;
    public mixed $ok;
    public ?string $organism = null;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** ConstructQc entity data model. */
class ConstructQc
{
    public ?array $avoidEnzymes = null;
    public ?int $crypticOrfMinAa = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?float $gcHigh = null;
    public ?float $gcLow = null;
    public ?int $gcWindow = null;
    public ?int $homopolymerMin = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for ConstructQc#create. */
class ConstructQcCreateData
{
    public ?array $avoidEnzymes = null;
    public ?int $crypticOrfMinAa = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?float $gcHigh = null;
    public ?float $gcLow = null;
    public ?int $gcWindow = null;
    public ?int $homopolymerMin = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** CrisprGrnaDesign entity data model. */
class CrisprGrnaDesign
{
    public mixed $gate = null;
    public ?float $minScore = null;
    public ?string $nuclease = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $searchReverseStrand = null;
    public string $sequence;
    public string $tool;
}

/** Request payload for CrisprGrnaDesign#create. */
class CrisprGrnaDesignCreateData
{
    public mixed $gate = null;
    public ?float $minScore = null;
    public ?string $nuclease = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $searchReverseStrand = null;
    public string $sequence;
    public string $tool;
}

/** CrisprHdrDonor entity data model. */
class CrisprHdrDonor
{
    public ?int $armLength = null;
    public ?bool $blockPam = null;
    public ?bool $designGenotypingPrimers = null;
    public ?int $editEnd = null;
    public ?int $editStart = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?int $guideEnd = null;
    public ?int $guideStart = null;
    public ?string $guideStrand = null;
    public ?string $nuclease = null;
    public mixed $ok;
    public array $provenance;
    public string $replacement;
    public array $result;
    public string $targetSequence;
    public string $tool;
}

/** Request payload for CrisprHdrDonor#create. */
class CrisprHdrDonorCreateData
{
    public ?int $armLength = null;
    public ?bool $blockPam = null;
    public ?bool $designGenotypingPrimers = null;
    public ?int $editEnd = null;
    public ?int $editStart = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?int $guideEnd = null;
    public ?int $guideStart = null;
    public ?string $guideStrand = null;
    public ?string $nuclease = null;
    public mixed $ok;
    public array $provenance;
    public string $replacement;
    public array $result;
    public string $targetSequence;
    public string $tool;
}

/** CrisprOfftargetCheck entity data model. */
class CrisprOfftargetCheck
{
    public mixed $gate = null;
    public ?int $maxMismatches = null;
    public ?string $nuclease = null;
    public mixed $ok;
    public string $protospacer;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for CrisprOfftargetCheck#create. */
class CrisprOfftargetCheckCreateData
{
    public mixed $gate = null;
    public ?int $maxMismatches = null;
    public ?string $nuclease = null;
    public mixed $ok;
    public string $protospacer;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** CrossDimer entity data model. */
class CrossDimer
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequenceA;
    public string $sequenceB;
    public string $tool;
}

/** Request payload for CrossDimer#create. */
class CrossDimerCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequenceA;
    public string $sequenceB;
    public string $tool;
}

/** DnaMolarity entity data model. */
class DnaMolarity
{
    public mixed $gate = null;
    public ?int $length = null;
    public ?float $massNg = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $sequence = null;
    public string $tool;
    public ?string $type = null;
    public ?float $volumeUl = null;
}

/** Request payload for DnaMolarity#create. */
class DnaMolarityCreateData
{
    public mixed $gate = null;
    public ?int $length = null;
    public ?float $massNg = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $sequence = null;
    public string $tool;
    public ?string $type = null;
    public ?float $volumeUl = null;
}

/** DoubleDigest entity data model. */
class DoubleDigest
{
    public string $enzymeA;
    public string $enzymeB;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for DoubleDigest#create. */
class DoubleDigestCreateData
{
    public string $enzymeA;
    public string $enzymeB;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** ExportEchoPicklist entity data model. */
class ExportEchoPicklist
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $reactions;
    public array $result;
    public string $tool;
}

/** Request payload for ExportEchoPicklist#create. */
class ExportEchoPicklistCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $reactions;
    public array $result;
    public string $tool;
}

/** ExportOpentronsProtocol entity data model. */
class ExportOpentronsProtocol
{
    public mixed $gate = null;
    public mixed $ok;
    public ?string $protocolName = null;
    public array $provenance;
    public array $reactions;
    public array $result;
    public string $tool;
}

/** Request payload for ExportOpentronsProtocol#create. */
class ExportOpentronsProtocolCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public ?string $protocolName = null;
    public array $provenance;
    public array $reactions;
    public array $result;
    public string $tool;
}

/** ExportPlateLayout entity data model. */
class ExportPlateLayout
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $reactions;
    public array $result;
    public string $tool;
}

/** Request payload for ExportPlateLayout#create. */
class ExportPlateLayoutCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $reactions;
    public array $result;
    public string $tool;
}

/** ExpressionHeatmapCluster entity data model. */
class ExpressionHeatmapCluster
{
    public ?bool $clusterCols = null;
    public ?bool $clusterRows = null;
    public ?string $distanceMetric = null;
    public mixed $gate = null;
    public array $genes;
    public ?string $linkage = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public array $samples;
    public string $tool;
    public array $values;
    public ?bool $zScoreRows = null;
}

/** Request payload for ExpressionHeatmapCluster#create. */
class ExpressionHeatmapClusterCreateData
{
    public ?bool $clusterCols = null;
    public ?bool $clusterRows = null;
    public ?string $distanceMetric = null;
    public mixed $gate = null;
    public array $genes;
    public ?string $linkage = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public array $samples;
    public string $tool;
    public array $values;
    public ?bool $zScoreRows = null;
}

/** FastqQcReport entity data model. */
class FastqQcReport
{
    public mixed $gate = null;
    public string $input;
    public mixed $ok;
    public array $provenance;
    public ?int $qualityOffset = null;
    public array $result;
    public string $tool;
}

/** Request payload for FastqQcReport#create. */
class FastqQcReportCreateData
{
    public mixed $gate = null;
    public string $input;
    public mixed $ok;
    public array $provenance;
    public ?int $qualityOffset = null;
    public array $result;
    public string $tool;
}

/** FastqTrim entity data model. */
class FastqTrim
{
    public mixed $gate = null;
    public string $input;
    public ?int $minLength = null;
    public mixed $ok;
    public array $provenance;
    public ?int $qualityOffset = null;
    public ?int $qualityThreshold = null;
    public array $result;
    public string $tool;
}

/** Request payload for FastqTrim#create. */
class FastqTrimCreateData
{
    public mixed $gate = null;
    public string $input;
    public ?int $minLength = null;
    public mixed $ok;
    public array $provenance;
    public ?int $qualityOffset = null;
    public ?int $qualityThreshold = null;
    public array $result;
    public string $tool;
}

/** FindOrf entity data model. */
class FindOrf
{
    public mixed $gate = null;
    public ?int $minAaLength = null;
    public mixed $ok;
    public array $provenance;
    public ?bool $requireStop = null;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for FindOrf#create. */
class FindOrfCreateData
{
    public mixed $gate = null;
    public ?int $minAaLength = null;
    public mixed $ok;
    public array $provenance;
    public ?bool $requireStop = null;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** FormatSequence entity data model. */
class FormatSequence
{
    public ?string $caseMode = null;
    public ?string $convert = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $reverse = null;
    public string $sequence;
    public ?bool $stripNonLetters = null;
    public string $tool;
    public ?int $width = null;
}

/** Request payload for FormatSequence#create. */
class FormatSequenceCreateData
{
    public ?string $caseMode = null;
    public ?string $convert = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $reverse = null;
    public string $sequence;
    public ?bool $stripNonLetters = null;
    public string $tool;
    public ?int $width = null;
}

/** FunctionalEnrichment entity data model. */
class FunctionalEnrichment
{
    public ?array $background = null;
    public ?array $collections = null;
    public mixed $gate = null;
    public array $genes;
    public ?int $maxTermSize = null;
    public ?int $minTermSize = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for FunctionalEnrichment#create. */
class FunctionalEnrichmentCreateData
{
    public ?array $background = null;
    public ?array $collections = null;
    public mixed $gate = null;
    public array $genes;
    public ?int $maxTermSize = null;
    public ?int $minTermSize = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** GcContent entity data model. */
class GcContent
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for GcContent#create. */
class GcContentCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** GeneDossier entity data model. */
class GeneDossier
{
    public mixed $gate = null;
    public string $gene;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for GeneDossier#create. */
class GeneDossierCreateData
{
    public mixed $gate = null;
    public string $gene;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** GeneExpression entity data model. */
class GeneExpression
{
    public mixed $gate = null;
    public string $gene;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for GeneExpression#create. */
class GeneExpressionCreateData
{
    public mixed $gate = null;
    public string $gene;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** GeneModel entity data model. */
class GeneModel
{
    public mixed $gate = null;
    public string $gene;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for GeneModel#create. */
class GeneModelCreateData
{
    public mixed $gate = null;
    public string $gene;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** GoldenGateFidelity entity data model. */
class GoldenGateFidelity
{
    public ?string $compareToNamedSet = null;
    public ?string $dataset = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $overhangs;
    public array $provenance;
    public array $result;
    public ?float $riskThreshold = null;
    public string $tool;
}

/** Request payload for GoldenGateFidelity#create. */
class GoldenGateFidelityCreateData
{
    public ?string $compareToNamedSet = null;
    public ?string $dataset = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $overhangs;
    public array $provenance;
    public array $result;
    public ?float $riskThreshold = null;
    public string $tool;
}

/** HgvsConvert entity data model. */
class HgvsConvert
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
    public string $variant;
}

/** Request payload for HgvsConvert#create. */
class HgvsConvertCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
    public string $variant;
}

/** IdMapPoll entity data model. */
class IdMapPoll
{
    public mixed $gate = null;
    public string $jobId;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for IdMapPoll#create. */
class IdMapPollCreateData
{
    public mixed $gate = null;
    public string $jobId;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** IdMapSubmit entity data model. */
class IdMapSubmit
{
    public string $from;
    public mixed $gate = null;
    public array $ids;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $taxId = null;
    public string $to;
    public string $tool;
}

/** Request payload for IdMapSubmit#create. */
class IdMapSubmitCreateData
{
    public string $from;
    public mixed $gate = null;
    public array $ids;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $taxId = null;
    public string $to;
    public string $tool;
}

/** InSilicoPcr entity data model. */
class InSilicoPcr
{
    public ?bool $circular = null;
    public string $forwardPrimer;
    public mixed $gate = null;
    public ?int $maxMismatches = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $reversePrimer;
    public string $template;
    public string $tool;
}

/** Request payload for InSilicoPcr#create. */
class InSilicoPcrCreateData
{
    public ?bool $circular = null;
    public string $forwardPrimer;
    public mixed $gate = null;
    public ?int $maxMismatches = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $reversePrimer;
    public string $template;
    public string $tool;
}

/** KaspPrimerDesign entity data model. */
class KaspPrimerDesign
{
    public ?bool $addSecondaryMismatch = null;
    public string $alleleA;
    public string $alleleB;
    public mixed $gate = null;
    public ?int $maxAmplicon = null;
    public ?int $minAmplicon = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public int $snpPosition;
    public string $target;
    public ?float $targetCoreTm = null;
    public string $tool;
}

/** Request payload for KaspPrimerDesign#create. */
class KaspPrimerDesignCreateData
{
    public ?bool $addSecondaryMismatch = null;
    public string $alleleA;
    public string $alleleB;
    public mixed $gate = null;
    public ?int $maxAmplicon = null;
    public ?int $minAmplicon = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public int $snpPosition;
    public string $target;
    public ?float $targetCoreTm = null;
    public string $tool;
}

/** ListTool entity data model. */
class ListTool
{
}

/** Request payload for ListTool#load. */
class ListToolLoadMatch
{
}

/** MeltingTemperature entity data model. */
class MeltingTemperature
{
    public ?float $dntpMM = null;
    public mixed $gate = null;
    public ?float $mgMM = null;
    public ?float $naMM = null;
    public mixed $ok;
    public ?float $oligoNM = null;
    public array $provenance;
    public array $result;
    public string $sequence;
    public ?float $targetTm = null;
    public ?float $tmTolerance = null;
    public string $tool;
}

/** Request payload for MeltingTemperature#create. */
class MeltingTemperatureCreateData
{
    public ?float $dntpMM = null;
    public mixed $gate = null;
    public ?float $mgMM = null;
    public ?float $naMM = null;
    public mixed $ok;
    public ?float $oligoNM = null;
    public array $provenance;
    public array $result;
    public string $sequence;
    public ?float $targetTm = null;
    public ?float $tmTolerance = null;
    public string $tool;
}

/** MotifFinder entity data model. */
class MotifFinder
{
    public mixed $gate = null;
    public ?int $maxMismatches = null;
    public string $motif;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $searchReverseStrand = null;
    public string $sequence;
    public string $tool;
}

/** Request payload for MotifFinder#create. */
class MotifFinderCreateData
{
    public mixed $gate = null;
    public ?int $maxMismatches = null;
    public string $motif;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $searchReverseStrand = null;
    public string $sequence;
    public string $tool;
}

/** MultipleSequenceAlignment entity data model. */
class MultipleSequenceAlignment
{
    public mixed $gate = null;
    public string $input;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for MultipleSequenceAlignment#create. */
class MultipleSequenceAlignmentCreateData
{
    public mixed $gate = null;
    public string $input;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** OligoAnalysi entity data model. */
class OligoAnalysi
{
    public ?float $dntpMM = null;
    public mixed $gate = null;
    public ?float $mgMM = null;
    public ?float $naMM = null;
    public mixed $ok;
    public ?float $oligoNM = null;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for OligoAnalysi#create. */
class OligoAnalysiCreateData
{
    public ?float $dntpMM = null;
    public mixed $gate = null;
    public ?float $mgMM = null;
    public ?float $naMM = null;
    public mixed $ok;
    public ?float $oligoNM = null;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** OrthologMap entity data model. */
class OrthologMap
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $sourceSpecies = null;
    public array $symbols;
    public string $targetSpecies;
    public string $tool;
    public ?string $type = null;
}

/** Request payload for OrthologMap#create. */
class OrthologMapCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $sourceSpecies = null;
    public array $symbols;
    public string $targetSpecies;
    public string $tool;
    public ?string $type = null;
}

/** PairwiseAlignment entity data model. */
class PairwiseAlignment
{
    public ?float $gap = null;
    public mixed $gate = null;
    public ?float $match = null;
    public ?float $mismatch = null;
    public ?string $mode = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $seqA;
    public string $seqB;
    public string $tool;
}

/** Request payload for PairwiseAlignment#create. */
class PairwiseAlignmentCreateData
{
    public ?float $gap = null;
    public mixed $gate = null;
    public ?float $match = null;
    public ?float $mismatch = null;
    public ?string $mode = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $seqA;
    public string $seqB;
    public string $tool;
}

/** ParseGenbank entity data model. */
class ParseGenbank
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $text;
    public string $tool;
}

/** Request payload for ParseGenbank#create. */
class ParseGenbankCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $text;
    public string $tool;
}

/** ParseSangerTrace entity data model. */
class ParseSangerTrace
{
    public string $fileBase64;
    public ?string $fileName = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for ParseSangerTrace#create. */
class ParseSangerTraceCreateData
{
    public string $fileBase64;
    public ?string $fileName = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** PlasmidAnnotate entity data model. */
class PlasmidAnnotate
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for PlasmidAnnotate#create. */
class PlasmidAnnotateCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** PlasmidDeepAnnotate entity data model. */
class PlasmidDeepAnnotate
{
    public ?bool $circular = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for PlasmidDeepAnnotate#create. */
class PlasmidDeepAnnotateCreateData
{
    public ?bool $circular = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** PlasmidFullReport entity data model. */
class PlasmidFullReport
{
    public ?bool $circular = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
    public ?int $topN = null;
}

/** Request payload for PlasmidFullReport#create. */
class PlasmidFullReportCreateData
{
    public ?bool $circular = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
    public ?int $topN = null;
}

/** PlasmidIdentify entity data model. */
class PlasmidIdentify
{
    public ?bool $circular = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
    public ?int $topN = null;
}

/** Request payload for PlasmidIdentify#create. */
class PlasmidIdentifyCreateData
{
    public ?bool $circular = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
    public ?int $topN = null;
}

/** PrimeEditingDesign entity data model. */
class PrimeEditingDesign
{
    public int $editEnd;
    public int $editStart;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?string $insertedSeq = null;
    public mixed $ok;
    public ?int $pbsLength = null;
    public array $provenance;
    public array $result;
    public ?int $rttHomology = null;
    public string $target;
    public string $tool;
}

/** Request payload for PrimeEditingDesign#create. */
class PrimeEditingDesignCreateData
{
    public int $editEnd;
    public int $editStart;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?string $insertedSeq = null;
    public mixed $ok;
    public ?int $pbsLength = null;
    public array $provenance;
    public array $result;
    public ?int $rttHomology = null;
    public string $target;
    public string $tool;
}

/** PrimeEditingTwinDesign entity data model. */
class PrimeEditingTwinDesign
{
    public mixed $gate = null;
    public string $newSequence;
    public mixed $ok;
    public ?int $overlapLength = null;
    public ?int $pbsLength = null;
    public array $provenance;
    public int $replaceEnd;
    public int $replaceStart;
    public array $result;
    public string $target;
    public string $tool;
}

/** Request payload for PrimeEditingTwinDesign#create. */
class PrimeEditingTwinDesignCreateData
{
    public mixed $gate = null;
    public string $newSequence;
    public mixed $ok;
    public ?int $overlapLength = null;
    public ?int $pbsLength = null;
    public array $provenance;
    public int $replaceEnd;
    public int $replaceStart;
    public array $result;
    public string $target;
    public string $tool;
}

/** PrimerDesign entity data model. */
class PrimerDesign
{
    public ?int $ampliconMax = null;
    public ?int $ampliconMin = null;
    public ?float $dntpMM = null;
    public mixed $gate = null;
    public ?float $gcMax = null;
    public ?float $gcMin = null;
    public ?int $lenMax = null;
    public ?int $lenMin = null;
    public ?int $lenOpt = null;
    public ?int $maxReturn = null;
    public ?float $mgMM = null;
    public ?float $naMM = null;
    public mixed $ok;
    public ?float $oligoNM = null;
    public array $provenance;
    public array $result;
    public ?int $targetEnd = null;
    public ?int $targetStart = null;
    public string $template;
    public ?float $tmMax = null;
    public ?float $tmMaxDiff = null;
    public ?float $tmMin = null;
    public ?float $tmOpt = null;
    public string $tool;
}

/** Request payload for PrimerDesign#create. */
class PrimerDesignCreateData
{
    public ?int $ampliconMax = null;
    public ?int $ampliconMin = null;
    public ?float $dntpMM = null;
    public mixed $gate = null;
    public ?float $gcMax = null;
    public ?float $gcMin = null;
    public ?int $lenMax = null;
    public ?int $lenMin = null;
    public ?int $lenOpt = null;
    public ?int $maxReturn = null;
    public ?float $mgMM = null;
    public ?float $naMM = null;
    public mixed $ok;
    public ?float $oligoNM = null;
    public array $provenance;
    public array $result;
    public ?int $targetEnd = null;
    public ?int $targetStart = null;
    public string $template;
    public ?float $tmMax = null;
    public ?float $tmMaxDiff = null;
    public ?float $tmMin = null;
    public ?float $tmOpt = null;
    public string $tool;
}

/** PrimerSpecificity entity data model. */
class PrimerSpecificity
{
    public string $forwardPrimer;
    public mixed $gate = null;
    public ?int $maxMismatches = null;
    public ?int $maxProductLength = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $reversePrimer;
    public string $tool;
}

/** Request payload for PrimerSpecificity#create. */
class PrimerSpecificityCreateData
{
    public string $forwardPrimer;
    public mixed $gate = null;
    public ?int $maxMismatches = null;
    public ?int $maxProductLength = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $reversePrimer;
    public string $tool;
}

/** ProteaseDigestion entity data model. */
class ProteaseDigestion
{
    public mixed $gate = null;
    public ?float $maxMass = null;
    public ?int $maxPeptides = null;
    public ?float $minMass = null;
    public ?int $missedCleavages = null;
    public mixed $ok;
    public ?string $protease = null;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for ProteaseDigestion#create. */
class ProteaseDigestionCreateData
{
    public mixed $gate = null;
    public ?float $maxMass = null;
    public ?int $maxPeptides = null;
    public ?float $minMass = null;
    public ?int $missedCleavages = null;
    public mixed $ok;
    public ?string $protease = null;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** ProteinAnnotatePoll entity data model. */
class ProteinAnnotatePoll
{
    public mixed $gate = null;
    public string $jobId;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for ProteinAnnotatePoll#create. */
class ProteinAnnotatePollCreateData
{
    public mixed $gate = null;
    public string $jobId;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** ProteinAnnotateSubmit entity data model. */
class ProteinAnnotateSubmit
{
    public ?string $appl = null;
    public mixed $gate = null;
    public ?bool $goterms = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for ProteinAnnotateSubmit#create. */
class ProteinAnnotateSubmitCreateData
{
    public ?string $appl = null;
    public mixed $gate = null;
    public ?bool $goterms = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** ProteinHydrophobicity entity data model. */
class ProteinHydrophobicity
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $scale = null;
    public string $sequence;
    public string $tool;
    public ?int $window = null;
}

/** Request payload for ProteinHydrophobicity#create. */
class ProteinHydrophobicityCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $scale = null;
    public string $sequence;
    public string $tool;
    public ?int $window = null;
}

/** ProteinProperty entity data model. */
class ProteinProperty
{
    public ?float $chargeStep = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for ProteinProperty#create. */
class ProteinPropertyCreateData
{
    public ?float $chargeStep = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** RandomSequence entity data model. */
class RandomSequence
{
    public mixed $gate = null;
    public ?float $gcContent = null;
    public ?string $kind = null;
    public int $length;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for RandomSequence#create. */
class RandomSequenceCreateData
{
    public mixed $gate = null;
    public ?float $gcContent = null;
    public ?string $kind = null;
    public int $length;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** RestrictionSite entity data model. */
class RestrictionSite
{
    public ?array $enzymes = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for RestrictionSite#create. */
class RestrictionSiteCreateData
{
    public ?array $enzymes = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** ReverseComplement entity data model. */
class ReverseComplement
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
    public ?string $type = null;
}

/** Request payload for ReverseComplement#create. */
class ReverseComplementCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
    public ?string $type = null;
}

/** ReverseTranslate entity data model. */
class ReverseTranslate
{
    public mixed $gate = null;
    public ?string $mode = null;
    public mixed $ok;
    public ?string $organism = null;
    public string $protein;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for ReverseTranslate#create. */
class ReverseTranslateCreateData
{
    public mixed $gate = null;
    public ?string $mode = null;
    public mixed $ok;
    public ?string $organism = null;
    public string $protein;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** RnaFold entity data model. */
class RnaFold
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for RnaFold#create. */
class RnaFoldCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** SangerVsReference entity data model. */
class SangerVsReference
{
    public ?string $fileBase64 = null;
    public ?string $fileName = null;
    public mixed $gate = null;
    public ?float $minCoverage = null;
    public mixed $ok;
    public array $provenance;
    public ?string $read = null;
    public string $reference;
    public array $result;
    public string $tool;
}

/** Request payload for SangerVsReference#create. */
class SangerVsReferenceCreateData
{
    public ?string $fileBase64 = null;
    public ?string $fileName = null;
    public mixed $gate = null;
    public ?float $minCoverage = null;
    public mixed $ok;
    public array $provenance;
    public ?string $read = null;
    public string $reference;
    public array $result;
    public string $tool;
}

/** SavePermalink entity data model. */
class SavePermalink
{
    public array $args;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for SavePermalink#create. */
class SavePermalinkCreateData
{
    public array $args;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** SeqfileStat entity data model. */
class SeqfileStat
{
    public mixed $gate = null;
    public string $input;
    public mixed $ok;
    public array $provenance;
    public ?int $qualityOffset = null;
    public array $result;
    public string $tool;
}

/** Request payload for SeqfileStat#create. */
class SeqfileStatCreateData
{
    public mixed $gate = null;
    public string $input;
    public mixed $ok;
    public array $provenance;
    public ?int $qualityOffset = null;
    public array $result;
    public string $tool;
}

/** SequenceFetch entity data model. */
class SequenceFetch
{
    public string $accession;
    public ?string $db = null;
    public ?string $format = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for SequenceFetch#create. */
class SequenceFetchCreateData
{
    public string $accession;
    public ?string $db = null;
    public ?string $format = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** SequenceFormatConvert entity data model. */
class SequenceFormatConvert
{
    public ?string $from = null;
    public mixed $gate = null;
    public string $input;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $to = null;
    public string $tool;
}

/** Request payload for SequenceFormatConvert#create. */
class SequenceFormatConvertCreateData
{
    public ?string $from = null;
    public mixed $gate = null;
    public string $input;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $to = null;
    public string $tool;
}

/** SequenceReport entity data model. */
class SequenceReport
{
    public ?int $endPrimerLength = null;
    public mixed $gate = null;
    public ?int $maxOrfs = null;
    public ?int $minOrfAa = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for SequenceReport#create. */
class SequenceReportCreateData
{
    public ?int $endPrimerLength = null;
    public mixed $gate = null;
    public ?int $maxOrfs = null;
    public ?int $minOrfAa = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** SequenceSearch entity data model. */
class SequenceSearch
{
    public ?string $db = null;
    public mixed $gate = null;
    public ?string $gene = null;
    public ?int $maxResults = null;
    public mixed $ok;
    public ?string $organism = null;
    public array $provenance;
    public array $result;
    public ?string $term = null;
    public string $tool;
}

/** Request payload for SequenceSearch#create. */
class SequenceSearchCreateData
{
    public ?string $db = null;
    public mixed $gate = null;
    public ?string $gene = null;
    public ?int $maxResults = null;
    public mixed $ok;
    public ?string $organism = null;
    public array $provenance;
    public array $result;
    public ?string $term = null;
    public string $tool;
}

/** SequencingReadbackVerify entity data model. */
class SequencingReadbackVerify
{
    public mixed $gate = null;
    public ?int $minSupportingReads = null;
    public mixed $ok;
    public array $provenance;
    public string $reads;
    public string $reference;
    public array $result;
    public string $tool;
}

/** Request payload for SequencingReadbackVerify#create. */
class SequencingReadbackVerifyCreateData
{
    public mixed $gate = null;
    public ?int $minSupportingReads = null;
    public mixed $ok;
    public array $provenance;
    public string $reads;
    public string $reference;
    public array $result;
    public string $tool;
}

/** SessionCreate entity data model. */
class SessionCreate
{
    public ?array $entries = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for SessionCreate#create. */
class SessionCreateCreateData
{
    public ?array $entries = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** SessionGet entity data model. */
class SessionGet
{
    public mixed $gate = null;
    public ?array $names = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sessionId;
    public string $tool;
}

/** Request payload for SessionGet#create. */
class SessionGetCreateData
{
    public mixed $gate = null;
    public ?array $names = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sessionId;
    public string $tool;
}

/** SessionRun entity data model. */
class SessionRun
{
    public ?array $args = null;
    public ?array $fromSession = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sessionId;
    public string $tool;
    public ?array $writeBack = null;
}

/** Request payload for SessionRun#create. */
class SessionRunCreateData
{
    public ?array $args = null;
    public ?array $fromSession = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sessionId;
    public string $tool;
    public ?array $writeBack = null;
}

/** SessionSet entity data model. */
class SessionSet
{
    public array $entries;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sessionId;
    public string $tool;
}

/** Request payload for SessionSet#create. */
class SessionSetCreateData
{
    public array $entries;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sessionId;
    public string $tool;
}

/** SirnaDesign entity data model. */
class SirnaDesign
{
    public mixed $gate = null;
    public ?int $minReynolds = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $shRnaLoop = null;
    public string $target;
    public string $tool;
}

/** Request payload for SirnaDesign#create. */
class SirnaDesignCreateData
{
    public mixed $gate = null;
    public ?int $minReynolds = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $shRnaLoop = null;
    public string $target;
    public string $tool;
}

/** SiteDirectedMutagenesi entity data model. */
class SiteDirectedMutagenesi
{
    public ?float $armTmTarget = null;
    public ?float $dntpMM = null;
    public ?string $editKind = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?float $mgMM = null;
    public ?float $naMM = null;
    public ?string $newBase = null;
    public mixed $ok;
    public ?float $oligoNM = null;
    public ?string $organism = null;
    public ?int $position = null;
    public array $provenance;
    public ?int $residue = null;
    public array $result;
    public ?string $style = null;
    public ?string $targetAa = null;
    public string $template;
    public string $tool;
}

/** Request payload for SiteDirectedMutagenesi#create. */
class SiteDirectedMutagenesiCreateData
{
    public ?float $armTmTarget = null;
    public ?float $dntpMM = null;
    public ?string $editKind = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?float $mgMM = null;
    public ?float $naMM = null;
    public ?string $newBase = null;
    public mixed $ok;
    public ?float $oligoNM = null;
    public ?string $organism = null;
    public ?int $position = null;
    public array $provenance;
    public ?int $residue = null;
    public array $result;
    public ?string $style = null;
    public ?string $targetAa = null;
    public string $template;
    public string $tool;
}

/** Translate entity data model. */
class Translate
{
    public ?int $frame = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public ?bool $toStop = null;
    public string $tool;
}

/** Request payload for Translate#create. */
class TranslateCreateData
{
    public ?int $frame = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public ?bool $toStop = null;
    public string $tool;
}

/** VariantAnnotate entity data model. */
class VariantAnnotate
{
    public ?string $assembly = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
    public string $variant;
}

/** Request payload for VariantAnnotate#create. */
class VariantAnnotateCreateData
{
    public ?string $assembly = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
    public string $variant;
}

/** VariantComparator entity data model. */
class VariantComparator
{
    public ?bool $coding = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public string $query;
    public string $reference;
    public array $result;
    public string $tool;
}

/** Request payload for VariantComparator#create. */
class VariantComparatorCreateData
{
    public ?bool $coding = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public string $query;
    public string $reference;
    public array $result;
    public string $tool;
}

/** VerifyAssembly entity data model. */
class VerifyAssembly
{
    public ?float $armTmTarget = null;
    public ?bool $circular = null;
    public string $claimedConstruct;
    public ?bool $coding = null;
    public ?string $enzyme = null;
    public ?string $enzyme3 = null;
    public ?string $enzyme5 = null;
    public ?array $fragmentPcrs = null;
    public ?array $fragments = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?string $insert = null;
    public ?array $insertPcr = null;
    public string $method;
    public ?array $names = null;
    public mixed $ok;
    public ?int $overlapLen = null;
    public array $provenance;
    public array $result;
    public string $tool;
    public ?string $vector = null;
    public ?array $vectorPcr = null;
}

/** Request payload for VerifyAssembly#create. */
class VerifyAssemblyCreateData
{
    public ?float $armTmTarget = null;
    public ?bool $circular = null;
    public string $claimedConstruct;
    public ?bool $coding = null;
    public ?string $enzyme = null;
    public ?string $enzyme3 = null;
    public ?string $enzyme5 = null;
    public ?array $fragmentPcrs = null;
    public ?array $fragments = null;
    public ?int $frameStart = null;
    public mixed $gate = null;
    public ?string $insert = null;
    public ?array $insertPcr = null;
    public string $method;
    public ?array $names = null;
    public mixed $ok;
    public ?int $overlapLen = null;
    public array $provenance;
    public array $result;
    public string $tool;
    public ?string $vector = null;
    public ?array $vectorPcr = null;
}

/** VerifyConstruct entity data model. */
class VerifyConstruct
{
    public string $claimedConstruct;
    public ?int $expectedFrameStart = null;
    public mixed $gate = null;
    public string $insertForwardPrimer;
    public string $insertReversePrimer;
    public string $insertTemplate;
    public ?int $maxPrimerMismatches = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $templateCircular = null;
    public string $tool;
}

/** Request payload for VerifyConstruct#create. */
class VerifyConstructCreateData
{
    public string $claimedConstruct;
    public ?int $expectedFrameStart = null;
    public mixed $gate = null;
    public string $insertForwardPrimer;
    public string $insertReversePrimer;
    public string $insertTemplate;
    public ?int $maxPrimerMismatches = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $templateCircular = null;
    public string $tool;
}

/** VirtualGel entity data model. */
class VirtualGel
{
    public ?bool $circular = null;
    public ?array $enzymes = null;
    public mixed $gate = null;
    public ?string $ladder = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for VirtualGel#create. */
class VirtualGelCreateData
{
    public ?bool $circular = null;
    public ?array $enzymes = null;
    public mixed $gate = null;
    public ?string $ladder = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** VolcanoPlotData entity data model. */
class VolcanoPlotData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public array $rows;
    public string $tool;
}

/** Request payload for VolcanoPlotData#create. */
class VolcanoPlotDataCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public array $rows;
    public string $tool;
}

/** WebSearch entity data model. */
class WebSearch
{
    public mixed $gate = null;
    public ?float $max_results = null;
    public mixed $ok;
    public array $provenance;
    public string $query;
    public array $result;
    public string $tool;
}

/** Request payload for WebSearch#create. */
class WebSearchCreateData
{
    public mixed $gate = null;
    public ?float $max_results = null;
    public mixed $ok;
    public array $provenance;
    public string $query;
    public array $result;
    public string $tool;
}

