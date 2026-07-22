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
    public ?int $frame_start = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $target;
    public ?int $target_position = null;
    public string $tool;
}

/** Request payload for BaseEditingDesign#create. */
class BaseEditingDesignCreateData
{
    public ?string $editor = null;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $target;
    public ?int $target_position = null;
    public string $tool;
}

/** Batch entity data model. */
class Batch
{
    public ?array $arg = null;
    public string $input;
    public mixed $ok;
    public array $result;
    public string $tool;
}

/** Request payload for Batch#load. */
class BatchLoadMatch
{
    public ?array $arg = null;
    public ?string $input = null;
    public mixed $ok = null;
    public ?array $result = null;
    public ?string $tool = null;
}

/** Request payload for Batch#create. */
class BatchCreateData
{
    public ?array $arg = null;
    public string $input;
    public mixed $ok;
    public array $result;
    public string $tool;
}

/** BatchWorkflow entity data model. */
class BatchWorkflow
{
    public string $input;
    public mixed $ok;
    public array $result;
    public array $step;
}

/** Request payload for BatchWorkflow#load. */
class BatchWorkflowLoadMatch
{
    public ?string $input = null;
    public mixed $ok = null;
    public ?array $result = null;
    public ?array $step = null;
}

/** Request payload for BatchWorkflow#create. */
class BatchWorkflowCreateData
{
    public string $input;
    public mixed $ok;
    public array $result;
    public array $step;
}

/** CharacterizeSequence entity data model. */
class CharacterizeSequence
{
    public ?int $end_primer_length = null;
    public mixed $gate = null;
    public ?int $max_orf = null;
    public ?int $min_orf_aa = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for CharacterizeSequence#create. */
class CharacterizeSequenceCreateData
{
    public ?int $end_primer_length = null;
    public mixed $gate = null;
    public ?int $max_orf = null;
    public ?int $min_orf_aa = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** CloningSimulate entity data model. */
class CloningSimulate
{
    public ?float $arm_tm_target = null;
    public ?bool $circular = null;
    public ?string $enzyme = null;
    public ?string $enzyme3 = null;
    public ?string $enzyme5 = null;
    public ?array $fragment = null;
    public mixed $gate = null;
    public ?string $insert = null;
    public string $method;
    public ?array $name = null;
    public mixed $ok;
    public ?int $overlap_len = null;
    public array $provenance;
    public array $result;
    public string $tool;
    public ?string $vector = null;
}

/** Request payload for CloningSimulate#create. */
class CloningSimulateCreateData
{
    public ?float $arm_tm_target = null;
    public ?bool $circular = null;
    public ?string $enzyme = null;
    public ?string $enzyme3 = null;
    public ?string $enzyme5 = null;
    public ?array $fragment = null;
    public mixed $gate = null;
    public ?string $insert = null;
    public string $method;
    public ?array $name = null;
    public mixed $ok;
    public ?int $overlap_len = null;
    public array $provenance;
    public array $result;
    public string $tool;
    public ?string $vector = null;
}

/** CodonAdaptationIndex entity data model. */
class CodonAdaptationIndex
{
    public ?int $frame_start = null;
    public mixed $gate = null;
    public mixed $ok;
    public ?string $organism = null;
    public array $provenance;
    public ?float $rare_threshold = null;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for CodonAdaptationIndex#create. */
class CodonAdaptationIndexCreateData
{
    public ?int $frame_start = null;
    public mixed $gate = null;
    public mixed $ok;
    public ?string $organism = null;
    public array $provenance;
    public ?float $rare_threshold = null;
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
    public ?array $avoid_enzyme = null;
    public ?int $cryptic_orf_min_aa = null;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?float $gc_high = null;
    public ?float $gc_low = null;
    public ?int $gc_window = null;
    public ?int $homopolymer_min = null;
    public ?int $max_pass = null;
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
    public ?array $avoid_enzyme = null;
    public ?int $cryptic_orf_min_aa = null;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?float $gc_high = null;
    public ?float $gc_low = null;
    public ?int $gc_window = null;
    public ?int $homopolymer_min = null;
    public ?int $max_pass = null;
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
    public ?array $avoid_enzyme = null;
    public ?int $cryptic_orf_min_aa = null;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?float $gc_high = null;
    public ?float $gc_low = null;
    public ?int $gc_window = null;
    public ?int $homopolymer_min = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for ConstructQc#create. */
class ConstructQcCreateData
{
    public ?array $avoid_enzyme = null;
    public ?int $cryptic_orf_min_aa = null;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?float $gc_high = null;
    public ?float $gc_low = null;
    public ?int $gc_window = null;
    public ?int $homopolymer_min = null;
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
    public ?float $min_score = null;
    public ?string $nuclease = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $search_reverse_strand = null;
    public string $sequence;
    public string $tool;
}

/** Request payload for CrisprGrnaDesign#create. */
class CrisprGrnaDesignCreateData
{
    public mixed $gate = null;
    public ?float $min_score = null;
    public ?string $nuclease = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $search_reverse_strand = null;
    public string $sequence;
    public string $tool;
}

/** CrisprHdrDonor entity data model. */
class CrisprHdrDonor
{
    public ?int $arm_length = null;
    public ?bool $block_pam = null;
    public ?bool $design_genotyping_primer = null;
    public ?int $edit_end = null;
    public ?int $edit_start = null;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?int $guide_end = null;
    public ?int $guide_start = null;
    public ?string $guide_strand = null;
    public ?string $nuclease = null;
    public mixed $ok;
    public array $provenance;
    public string $replacement;
    public array $result;
    public string $target_sequence;
    public string $tool;
}

/** Request payload for CrisprHdrDonor#create. */
class CrisprHdrDonorCreateData
{
    public ?int $arm_length = null;
    public ?bool $block_pam = null;
    public ?bool $design_genotyping_primer = null;
    public ?int $edit_end = null;
    public ?int $edit_start = null;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?int $guide_end = null;
    public ?int $guide_start = null;
    public ?string $guide_strand = null;
    public ?string $nuclease = null;
    public mixed $ok;
    public array $provenance;
    public string $replacement;
    public array $result;
    public string $target_sequence;
    public string $tool;
}

/** CrisprOfftargetCheck entity data model. */
class CrisprOfftargetCheck
{
    public mixed $gate = null;
    public ?int $max_mismatch = null;
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
    public ?int $max_mismatch = null;
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
    public string $sequence_a;
    public string $sequence_b;
    public string $tool;
}

/** Request payload for CrossDimer#create. */
class CrossDimerCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence_a;
    public string $sequence_b;
    public string $tool;
}

/** DnaMolarity entity data model. */
class DnaMolarity
{
    public mixed $gate = null;
    public ?int $length = null;
    public ?float $mass_ng = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $sequence = null;
    public string $tool;
    public ?string $type = null;
    public ?float $volume_ul = null;
}

/** Request payload for DnaMolarity#create. */
class DnaMolarityCreateData
{
    public mixed $gate = null;
    public ?int $length = null;
    public ?float $mass_ng = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $sequence = null;
    public string $tool;
    public ?string $type = null;
    public ?float $volume_ul = null;
}

/** DoubleDigest entity data model. */
class DoubleDigest
{
    public string $enzyme_a;
    public string $enzyme_b;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for DoubleDigest#create. */
class DoubleDigestCreateData
{
    public string $enzyme_a;
    public string $enzyme_b;
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
    public array $reaction;
    public array $result;
    public string $tool;
}

/** Request payload for ExportEchoPicklist#create. */
class ExportEchoPicklistCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $reaction;
    public array $result;
    public string $tool;
}

/** ExportOpentronsProtocol entity data model. */
class ExportOpentronsProtocol
{
    public mixed $gate = null;
    public mixed $ok;
    public ?string $protocol_name = null;
    public array $provenance;
    public array $reaction;
    public array $result;
    public string $tool;
}

/** Request payload for ExportOpentronsProtocol#create. */
class ExportOpentronsProtocolCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public ?string $protocol_name = null;
    public array $provenance;
    public array $reaction;
    public array $result;
    public string $tool;
}

/** ExportPlateLayout entity data model. */
class ExportPlateLayout
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $reaction;
    public array $result;
    public string $tool;
}

/** Request payload for ExportPlateLayout#create. */
class ExportPlateLayoutCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $reaction;
    public array $result;
    public string $tool;
}

/** ExpressionHeatmapCluster entity data model. */
class ExpressionHeatmapCluster
{
    public ?bool $cluster_col = null;
    public ?bool $cluster_row = null;
    public ?string $distance_metric = null;
    public mixed $gate = null;
    public array $gene;
    public ?string $linkage = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public array $sample;
    public string $tool;
    public array $value;
    public ?bool $z_score_row = null;
}

/** Request payload for ExpressionHeatmapCluster#create. */
class ExpressionHeatmapClusterCreateData
{
    public ?bool $cluster_col = null;
    public ?bool $cluster_row = null;
    public ?string $distance_metric = null;
    public mixed $gate = null;
    public array $gene;
    public ?string $linkage = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public array $sample;
    public string $tool;
    public array $value;
    public ?bool $z_score_row = null;
}

/** FastqQcReport entity data model. */
class FastqQcReport
{
    public mixed $gate = null;
    public string $input;
    public mixed $ok;
    public array $provenance;
    public ?int $quality_offset = null;
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
    public ?int $quality_offset = null;
    public array $result;
    public string $tool;
}

/** FastqTrim entity data model. */
class FastqTrim
{
    public mixed $gate = null;
    public string $input;
    public ?int $min_length = null;
    public mixed $ok;
    public array $provenance;
    public ?int $quality_offset = null;
    public ?int $quality_threshold = null;
    public array $result;
    public string $tool;
}

/** Request payload for FastqTrim#create. */
class FastqTrimCreateData
{
    public mixed $gate = null;
    public string $input;
    public ?int $min_length = null;
    public mixed $ok;
    public array $provenance;
    public ?int $quality_offset = null;
    public ?int $quality_threshold = null;
    public array $result;
    public string $tool;
}

/** FindOrf entity data model. */
class FindOrf
{
    public mixed $gate = null;
    public ?int $min_aa_length = null;
    public mixed $ok;
    public array $provenance;
    public ?bool $require_stop = null;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for FindOrf#create. */
class FindOrfCreateData
{
    public mixed $gate = null;
    public ?int $min_aa_length = null;
    public mixed $ok;
    public array $provenance;
    public ?bool $require_stop = null;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** FormatSequence entity data model. */
class FormatSequence
{
    public ?string $case_mode = null;
    public ?string $convert = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $reverse = null;
    public string $sequence;
    public ?bool $strip_non_letter = null;
    public string $tool;
    public ?int $width = null;
}

/** Request payload for FormatSequence#create. */
class FormatSequenceCreateData
{
    public ?string $case_mode = null;
    public ?string $convert = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $reverse = null;
    public string $sequence;
    public ?bool $strip_non_letter = null;
    public string $tool;
    public ?int $width = null;
}

/** FunctionalEnrichment entity data model. */
class FunctionalEnrichment
{
    public ?array $background = null;
    public ?array $collection = null;
    public mixed $gate = null;
    public array $gene;
    public ?int $max_term_size = null;
    public ?int $min_term_size = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for FunctionalEnrichment#create. */
class FunctionalEnrichmentCreateData
{
    public ?array $background = null;
    public ?array $collection = null;
    public mixed $gate = null;
    public array $gene;
    public ?int $max_term_size = null;
    public ?int $min_term_size = null;
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
    public ?string $compare_to_named_set = null;
    public ?string $dataset = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $overhang;
    public array $provenance;
    public array $result;
    public ?float $risk_threshold = null;
    public string $tool;
}

/** Request payload for GoldenGateFidelity#create. */
class GoldenGateFidelityCreateData
{
    public ?string $compare_to_named_set = null;
    public ?string $dataset = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $overhang;
    public array $provenance;
    public array $result;
    public ?float $risk_threshold = null;
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
    public string $job_id;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for IdMapPoll#create. */
class IdMapPollCreateData
{
    public mixed $gate = null;
    public string $job_id;
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
    public ?string $tax_id = null;
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
    public ?string $tax_id = null;
    public string $to;
    public string $tool;
}

/** InSilicoPcr entity data model. */
class InSilicoPcr
{
    public ?bool $circular = null;
    public string $forward_primer;
    public mixed $gate = null;
    public ?int $max_mismatch = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $reverse_primer;
    public string $template;
    public string $tool;
}

/** Request payload for InSilicoPcr#create. */
class InSilicoPcrCreateData
{
    public ?bool $circular = null;
    public string $forward_primer;
    public mixed $gate = null;
    public ?int $max_mismatch = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $reverse_primer;
    public string $template;
    public string $tool;
}

/** KaspPrimerDesign entity data model. */
class KaspPrimerDesign
{
    public ?bool $add_secondary_mismatch = null;
    public string $allele_a;
    public string $allele_b;
    public mixed $gate = null;
    public ?int $max_amplicon = null;
    public ?int $min_amplicon = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public int $snp_position;
    public string $target;
    public ?float $target_core_tm = null;
    public string $tool;
}

/** Request payload for KaspPrimerDesign#create. */
class KaspPrimerDesignCreateData
{
    public ?bool $add_secondary_mismatch = null;
    public string $allele_a;
    public string $allele_b;
    public mixed $gate = null;
    public ?int $max_amplicon = null;
    public ?int $min_amplicon = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public int $snp_position;
    public string $target;
    public ?float $target_core_tm = null;
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
    public ?float $dntp_mm = null;
    public mixed $gate = null;
    public ?float $mg_mm = null;
    public ?float $na_mm = null;
    public mixed $ok;
    public ?float $oligo_nm = null;
    public array $provenance;
    public array $result;
    public string $sequence;
    public ?float $target_tm = null;
    public ?float $tm_tolerance = null;
    public string $tool;
}

/** Request payload for MeltingTemperature#create. */
class MeltingTemperatureCreateData
{
    public ?float $dntp_mm = null;
    public mixed $gate = null;
    public ?float $mg_mm = null;
    public ?float $na_mm = null;
    public mixed $ok;
    public ?float $oligo_nm = null;
    public array $provenance;
    public array $result;
    public string $sequence;
    public ?float $target_tm = null;
    public ?float $tm_tolerance = null;
    public string $tool;
}

/** MotifFinder entity data model. */
class MotifFinder
{
    public mixed $gate = null;
    public ?int $max_mismatch = null;
    public string $motif;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $search_reverse_strand = null;
    public string $sequence;
    public string $tool;
}

/** Request payload for MotifFinder#create. */
class MotifFinderCreateData
{
    public mixed $gate = null;
    public ?int $max_mismatch = null;
    public string $motif;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $search_reverse_strand = null;
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
    public ?float $dntp_mm = null;
    public mixed $gate = null;
    public ?float $mg_mm = null;
    public ?float $na_mm = null;
    public mixed $ok;
    public ?float $oligo_nm = null;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for OligoAnalysi#create. */
class OligoAnalysiCreateData
{
    public ?float $dntp_mm = null;
    public mixed $gate = null;
    public ?float $mg_mm = null;
    public ?float $na_mm = null;
    public mixed $ok;
    public ?float $oligo_nm = null;
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
    public ?string $source_species = null;
    public array $symbol;
    public string $target_species;
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
    public ?string $source_species = null;
    public array $symbol;
    public string $target_species;
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
    public string $seq_a;
    public string $seq_b;
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
    public string $seq_a;
    public string $seq_b;
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
    public string $file_base64;
    public ?string $file_name = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for ParseSangerTrace#create. */
class ParseSangerTraceCreateData
{
    public string $file_base64;
    public ?string $file_name = null;
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
    public ?int $top_n = null;
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
    public ?int $top_n = null;
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
    public ?int $top_n = null;
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
    public ?int $top_n = null;
}

/** PrimeEditingDesign entity data model. */
class PrimeEditingDesign
{
    public int $edit_end;
    public int $edit_start;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?string $inserted_seq = null;
    public mixed $ok;
    public ?int $pbs_length = null;
    public array $provenance;
    public array $result;
    public ?int $rtt_homology = null;
    public string $target;
    public string $tool;
}

/** Request payload for PrimeEditingDesign#create. */
class PrimeEditingDesignCreateData
{
    public int $edit_end;
    public int $edit_start;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?string $inserted_seq = null;
    public mixed $ok;
    public ?int $pbs_length = null;
    public array $provenance;
    public array $result;
    public ?int $rtt_homology = null;
    public string $target;
    public string $tool;
}

/** PrimeEditingTwinDesign entity data model. */
class PrimeEditingTwinDesign
{
    public mixed $gate = null;
    public string $new_sequence;
    public mixed $ok;
    public ?int $overlap_length = null;
    public ?int $pbs_length = null;
    public array $provenance;
    public int $replace_end;
    public int $replace_start;
    public array $result;
    public string $target;
    public string $tool;
}

/** Request payload for PrimeEditingTwinDesign#create. */
class PrimeEditingTwinDesignCreateData
{
    public mixed $gate = null;
    public string $new_sequence;
    public mixed $ok;
    public ?int $overlap_length = null;
    public ?int $pbs_length = null;
    public array $provenance;
    public int $replace_end;
    public int $replace_start;
    public array $result;
    public string $target;
    public string $tool;
}

/** PrimerDesign entity data model. */
class PrimerDesign
{
    public ?int $amplicon_max = null;
    public ?int $amplicon_min = null;
    public ?float $dntp_mm = null;
    public mixed $gate = null;
    public ?float $gc_max = null;
    public ?float $gc_min = null;
    public ?int $len_max = null;
    public ?int $len_min = null;
    public ?int $len_opt = null;
    public ?int $max_return = null;
    public ?float $mg_mm = null;
    public ?float $na_mm = null;
    public mixed $ok;
    public ?float $oligo_nm = null;
    public array $provenance;
    public array $result;
    public ?int $target_end = null;
    public ?int $target_start = null;
    public string $template;
    public ?float $tm_max = null;
    public ?float $tm_max_diff = null;
    public ?float $tm_min = null;
    public ?float $tm_opt = null;
    public string $tool;
}

/** Request payload for PrimerDesign#create. */
class PrimerDesignCreateData
{
    public ?int $amplicon_max = null;
    public ?int $amplicon_min = null;
    public ?float $dntp_mm = null;
    public mixed $gate = null;
    public ?float $gc_max = null;
    public ?float $gc_min = null;
    public ?int $len_max = null;
    public ?int $len_min = null;
    public ?int $len_opt = null;
    public ?int $max_return = null;
    public ?float $mg_mm = null;
    public ?float $na_mm = null;
    public mixed $ok;
    public ?float $oligo_nm = null;
    public array $provenance;
    public array $result;
    public ?int $target_end = null;
    public ?int $target_start = null;
    public string $template;
    public ?float $tm_max = null;
    public ?float $tm_max_diff = null;
    public ?float $tm_min = null;
    public ?float $tm_opt = null;
    public string $tool;
}

/** PrimerSpecificity entity data model. */
class PrimerSpecificity
{
    public string $forward_primer;
    public mixed $gate = null;
    public ?int $max_mismatch = null;
    public ?int $max_product_length = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $reverse_primer;
    public string $tool;
}

/** Request payload for PrimerSpecificity#create. */
class PrimerSpecificityCreateData
{
    public string $forward_primer;
    public mixed $gate = null;
    public ?int $max_mismatch = null;
    public ?int $max_product_length = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $reverse_primer;
    public string $tool;
}

/** ProteaseDigestion entity data model. */
class ProteaseDigestion
{
    public mixed $gate = null;
    public ?float $max_mass = null;
    public ?int $max_peptide = null;
    public ?float $min_mass = null;
    public ?int $missed_cleavage = null;
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
    public ?float $max_mass = null;
    public ?int $max_peptide = null;
    public ?float $min_mass = null;
    public ?int $missed_cleavage = null;
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
    public string $job_id;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for ProteinAnnotatePoll#create. */
class ProteinAnnotatePollCreateData
{
    public mixed $gate = null;
    public string $job_id;
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
    public ?bool $goterm = null;
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
    public ?bool $goterm = null;
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
    public ?float $charge_step = null;
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
    public ?float $charge_step = null;
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
    public ?float $gc_content = null;
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
    public ?float $gc_content = null;
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
    public ?array $enzyme = null;
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
    public ?array $enzyme = null;
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
    public ?string $file_base64 = null;
    public ?string $file_name = null;
    public mixed $gate = null;
    public ?float $min_coverage = null;
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
    public ?string $file_base64 = null;
    public ?string $file_name = null;
    public mixed $gate = null;
    public ?float $min_coverage = null;
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
    public array $arg;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for SavePermalink#create. */
class SavePermalinkCreateData
{
    public array $arg;
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
    public ?int $quality_offset = null;
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
    public ?int $quality_offset = null;
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
    public ?int $end_primer_length = null;
    public mixed $gate = null;
    public ?int $max_orf = null;
    public ?int $min_orf_aa = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $sequence;
    public string $tool;
}

/** Request payload for SequenceReport#create. */
class SequenceReportCreateData
{
    public ?int $end_primer_length = null;
    public mixed $gate = null;
    public ?int $max_orf = null;
    public ?int $min_orf_aa = null;
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
    public ?int $max_result = null;
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
    public ?int $max_result = null;
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
    public ?int $min_supporting_read = null;
    public mixed $ok;
    public array $provenance;
    public string $read;
    public string $reference;
    public array $result;
    public string $tool;
}

/** Request payload for SequencingReadbackVerify#create. */
class SequencingReadbackVerifyCreateData
{
    public mixed $gate = null;
    public ?int $min_supporting_read = null;
    public mixed $ok;
    public array $provenance;
    public string $read;
    public string $reference;
    public array $result;
    public string $tool;
}

/** SessionCreate entity data model. */
class SessionCreate
{
    public ?array $entry = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $tool;
}

/** Request payload for SessionCreate#create. */
class SessionCreateCreateData
{
    public ?array $entry = null;
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
    public ?array $name = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $session_id;
    public string $tool;
}

/** Request payload for SessionGet#create. */
class SessionGetCreateData
{
    public mixed $gate = null;
    public ?array $name = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $session_id;
    public string $tool;
}

/** SessionRun entity data model. */
class SessionRun
{
    public ?array $arg = null;
    public ?array $from_session = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $session_id;
    public string $tool;
    public ?array $write_back = null;
}

/** Request payload for SessionRun#create. */
class SessionRunCreateData
{
    public ?array $arg = null;
    public ?array $from_session = null;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $session_id;
    public string $tool;
    public ?array $write_back = null;
}

/** SessionSet entity data model. */
class SessionSet
{
    public array $entry;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $session_id;
    public string $tool;
}

/** Request payload for SessionSet#create. */
class SessionSetCreateData
{
    public array $entry;
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public string $session_id;
    public string $tool;
}

/** SirnaDesign entity data model. */
class SirnaDesign
{
    public mixed $gate = null;
    public ?int $min_reynold = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $sh_rna_loop = null;
    public string $target;
    public string $tool;
}

/** Request payload for SirnaDesign#create. */
class SirnaDesignCreateData
{
    public mixed $gate = null;
    public ?int $min_reynold = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?string $sh_rna_loop = null;
    public string $target;
    public string $tool;
}

/** SiteDirectedMutagenesi entity data model. */
class SiteDirectedMutagenesi
{
    public ?float $arm_tm_target = null;
    public ?float $dntp_mm = null;
    public ?string $edit_kind = null;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?float $mg_mm = null;
    public ?float $na_mm = null;
    public ?string $new_base = null;
    public mixed $ok;
    public ?float $oligo_nm = null;
    public ?string $organism = null;
    public ?int $position = null;
    public array $provenance;
    public ?int $residue = null;
    public array $result;
    public ?string $style = null;
    public ?string $target_aa = null;
    public string $template;
    public string $tool;
}

/** Request payload for SiteDirectedMutagenesi#create. */
class SiteDirectedMutagenesiCreateData
{
    public ?float $arm_tm_target = null;
    public ?float $dntp_mm = null;
    public ?string $edit_kind = null;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?float $mg_mm = null;
    public ?float $na_mm = null;
    public ?string $new_base = null;
    public mixed $ok;
    public ?float $oligo_nm = null;
    public ?string $organism = null;
    public ?int $position = null;
    public array $provenance;
    public ?int $residue = null;
    public array $result;
    public ?string $style = null;
    public ?string $target_aa = null;
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
    public ?bool $to_stop = null;
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
    public ?bool $to_stop = null;
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
    public ?int $frame_start = null;
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
    public ?int $frame_start = null;
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
    public ?float $arm_tm_target = null;
    public ?bool $circular = null;
    public string $claimed_construct;
    public ?bool $coding = null;
    public ?string $enzyme = null;
    public ?string $enzyme3 = null;
    public ?string $enzyme5 = null;
    public ?array $fragment = null;
    public ?array $fragment_pcr = null;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?string $insert = null;
    public ?array $insert_pcr = null;
    public string $method;
    public ?array $name = null;
    public mixed $ok;
    public ?int $overlap_len = null;
    public array $provenance;
    public array $result;
    public string $tool;
    public ?string $vector = null;
    public ?array $vector_pcr = null;
}

/** Request payload for VerifyAssembly#create. */
class VerifyAssemblyCreateData
{
    public ?float $arm_tm_target = null;
    public ?bool $circular = null;
    public string $claimed_construct;
    public ?bool $coding = null;
    public ?string $enzyme = null;
    public ?string $enzyme3 = null;
    public ?string $enzyme5 = null;
    public ?array $fragment = null;
    public ?array $fragment_pcr = null;
    public ?int $frame_start = null;
    public mixed $gate = null;
    public ?string $insert = null;
    public ?array $insert_pcr = null;
    public string $method;
    public ?array $name = null;
    public mixed $ok;
    public ?int $overlap_len = null;
    public array $provenance;
    public array $result;
    public string $tool;
    public ?string $vector = null;
    public ?array $vector_pcr = null;
}

/** VerifyConstruct entity data model. */
class VerifyConstruct
{
    public string $claimed_construct;
    public ?int $expected_frame_start = null;
    public mixed $gate = null;
    public string $insert_forward_primer;
    public string $insert_reverse_primer;
    public string $insert_template;
    public ?int $max_primer_mismatch = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $template_circular = null;
    public string $tool;
}

/** Request payload for VerifyConstruct#create. */
class VerifyConstructCreateData
{
    public string $claimed_construct;
    public ?int $expected_frame_start = null;
    public mixed $gate = null;
    public string $insert_forward_primer;
    public string $insert_reverse_primer;
    public string $insert_template;
    public ?int $max_primer_mismatch = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public ?bool $template_circular = null;
    public string $tool;
}

/** VirtualGel entity data model. */
class VirtualGel
{
    public ?bool $circular = null;
    public ?array $enzyme = null;
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
    public ?array $enzyme = null;
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
    public array $row;
    public string $tool;
}

/** Request payload for VolcanoPlotData#create. */
class VolcanoPlotDataCreateData
{
    public mixed $gate = null;
    public mixed $ok;
    public array $provenance;
    public array $result;
    public array $row;
    public string $tool;
}

/** WebSearch entity data model. */
class WebSearch
{
    public mixed $gate = null;
    public ?float $max_result = null;
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
    public ?float $max_result = null;
    public mixed $ok;
    public array $provenance;
    public string $query;
    public array $result;
    public string $tool;
}

