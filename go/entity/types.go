// Typed models for the SeqbenchMcp SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// AlphafoldLookup is the typed data model for the alphafold_lookup entity.
type AlphafoldLookup struct {
	Accession string `json:"accession"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// AlphafoldLookupCreateData is the typed request payload for AlphafoldLookup.CreateTyped.
type AlphafoldLookupCreateData struct {
	Accession string `json:"accession"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// AsoDesign is the typed data model for the aso_design entity.
type AsoDesign struct {
	Gate *any `json:"gate,omitempty"`
	Length *int `json:"length,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Target string `json:"target"`
	Tool string `json:"tool"`
	Wing *int `json:"wing,omitempty"`
}

// AsoDesignCreateData is the typed request payload for AsoDesign.CreateTyped.
type AsoDesignCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Length *int `json:"length,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Target string `json:"target"`
	Tool string `json:"tool"`
	Wing *int `json:"wing,omitempty"`
}

// BaseEditingDesign is the typed data model for the base_editing_design entity.
type BaseEditingDesign struct {
	Editor *string `json:"editor,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Target string `json:"target"`
	TargetPosition *int `json:"target_position,omitempty"`
	Tool string `json:"tool"`
}

// BaseEditingDesignCreateData is the typed request payload for BaseEditingDesign.CreateTyped.
type BaseEditingDesignCreateData struct {
	Editor *string `json:"editor,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Target string `json:"target"`
	TargetPosition *int `json:"target_position,omitempty"`
	Tool string `json:"tool"`
}

// Batch is the typed data model for the batch entity.
type Batch struct {
	Arg *map[string]any `json:"arg,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// BatchLoadMatch is the typed request payload for Batch.LoadTyped.
type BatchLoadMatch struct {
	Arg *map[string]any `json:"arg,omitempty"`
	Input *string `json:"input,omitempty"`
	Ok *any `json:"ok,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Tool *string `json:"tool,omitempty"`
}

// BatchCreateData is the typed request payload for Batch.CreateTyped.
type BatchCreateData struct {
	Arg *map[string]any `json:"arg,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// BatchWorkflow is the typed data model for the batch__workflow entity.
type BatchWorkflow struct {
	Input string `json:"input"`
	Ok any `json:"ok"`
	Result map[string]any `json:"result"`
	Step []any `json:"step"`
}

// BatchWorkflowLoadMatch is the typed request payload for BatchWorkflow.LoadTyped.
type BatchWorkflowLoadMatch struct {
	Input *string `json:"input,omitempty"`
	Ok *any `json:"ok,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	Step *[]any `json:"step,omitempty"`
}

// BatchWorkflowCreateData is the typed request payload for BatchWorkflow.CreateTyped.
type BatchWorkflowCreateData struct {
	Input string `json:"input"`
	Ok any `json:"ok"`
	Result map[string]any `json:"result"`
	Step []any `json:"step"`
}

// CharacterizeSequence is the typed data model for the characterize_sequence entity.
type CharacterizeSequence struct {
	EndPrimerLength *int `json:"end_primer_length,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MaxOrf *int `json:"max_orf,omitempty"`
	MinOrfAa *int `json:"min_orf_aa,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CharacterizeSequenceCreateData is the typed request payload for CharacterizeSequence.CreateTyped.
type CharacterizeSequenceCreateData struct {
	EndPrimerLength *int `json:"end_primer_length,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MaxOrf *int `json:"max_orf,omitempty"`
	MinOrfAa *int `json:"min_orf_aa,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CloningSimulate is the typed data model for the cloning_simulate entity.
type CloningSimulate struct {
	ArmTmTarget *float64 `json:"arm_tm_target,omitempty"`
	Circular *bool `json:"circular,omitempty"`
	Enzyme *string `json:"enzyme,omitempty"`
	Enzyme3 *string `json:"enzyme3,omitempty"`
	Enzyme5 *string `json:"enzyme5,omitempty"`
	Fragment *[]any `json:"fragment,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Insert *string `json:"insert,omitempty"`
	Method string `json:"method"`
	Name *[]any `json:"name,omitempty"`
	Ok any `json:"ok"`
	OverlapLen *int `json:"overlap_len,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Vector *string `json:"vector,omitempty"`
}

// CloningSimulateCreateData is the typed request payload for CloningSimulate.CreateTyped.
type CloningSimulateCreateData struct {
	ArmTmTarget *float64 `json:"arm_tm_target,omitempty"`
	Circular *bool `json:"circular,omitempty"`
	Enzyme *string `json:"enzyme,omitempty"`
	Enzyme3 *string `json:"enzyme3,omitempty"`
	Enzyme5 *string `json:"enzyme5,omitempty"`
	Fragment *[]any `json:"fragment,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Insert *string `json:"insert,omitempty"`
	Method string `json:"method"`
	Name *[]any `json:"name,omitempty"`
	Ok any `json:"ok"`
	OverlapLen *int `json:"overlap_len,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Vector *string `json:"vector,omitempty"`
}

// CodonAdaptationIndex is the typed data model for the codon_adaptation_index entity.
type CodonAdaptationIndex struct {
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Provenance map[string]any `json:"provenance"`
	RareThreshold *float64 `json:"rare_threshold,omitempty"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CodonAdaptationIndexCreateData is the typed request payload for CodonAdaptationIndex.CreateTyped.
type CodonAdaptationIndexCreateData struct {
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Provenance map[string]any `json:"provenance"`
	RareThreshold *float64 `json:"rare_threshold,omitempty"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CodonOptimize is the typed data model for the codon_optimize entity.
type CodonOptimize struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Protein string `json:"protein"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// CodonOptimizeCreateData is the typed request payload for CodonOptimize.CreateTyped.
type CodonOptimizeCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Protein string `json:"protein"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ConstructAutofix is the typed data model for the construct_autofix entity.
type ConstructAutofix struct {
	AvoidEnzyme *[]any `json:"avoid_enzyme,omitempty"`
	CrypticOrfMinAa *int `json:"cryptic_orf_min_aa,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcHigh *float64 `json:"gc_high,omitempty"`
	GcLow *float64 `json:"gc_low,omitempty"`
	GcWindow *int `json:"gc_window,omitempty"`
	HomopolymerMin *int `json:"homopolymer_min,omitempty"`
	MaxPass *int `json:"max_pass,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ConstructAutofixCreateData is the typed request payload for ConstructAutofix.CreateTyped.
type ConstructAutofixCreateData struct {
	AvoidEnzyme *[]any `json:"avoid_enzyme,omitempty"`
	CrypticOrfMinAa *int `json:"cryptic_orf_min_aa,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcHigh *float64 `json:"gc_high,omitempty"`
	GcLow *float64 `json:"gc_low,omitempty"`
	GcWindow *int `json:"gc_window,omitempty"`
	HomopolymerMin *int `json:"homopolymer_min,omitempty"`
	MaxPass *int `json:"max_pass,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ConstructQc is the typed data model for the construct_qc entity.
type ConstructQc struct {
	AvoidEnzyme *[]any `json:"avoid_enzyme,omitempty"`
	CrypticOrfMinAa *int `json:"cryptic_orf_min_aa,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcHigh *float64 `json:"gc_high,omitempty"`
	GcLow *float64 `json:"gc_low,omitempty"`
	GcWindow *int `json:"gc_window,omitempty"`
	HomopolymerMin *int `json:"homopolymer_min,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ConstructQcCreateData is the typed request payload for ConstructQc.CreateTyped.
type ConstructQcCreateData struct {
	AvoidEnzyme *[]any `json:"avoid_enzyme,omitempty"`
	CrypticOrfMinAa *int `json:"cryptic_orf_min_aa,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcHigh *float64 `json:"gc_high,omitempty"`
	GcLow *float64 `json:"gc_low,omitempty"`
	GcWindow *int `json:"gc_window,omitempty"`
	HomopolymerMin *int `json:"homopolymer_min,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CrisprGrnaDesign is the typed data model for the crispr_grna_design entity.
type CrisprGrnaDesign struct {
	Gate *any `json:"gate,omitempty"`
	MinScore *float64 `json:"min_score,omitempty"`
	Nuclease *string `json:"nuclease,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SearchReverseStrand *bool `json:"search_reverse_strand,omitempty"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CrisprGrnaDesignCreateData is the typed request payload for CrisprGrnaDesign.CreateTyped.
type CrisprGrnaDesignCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MinScore *float64 `json:"min_score,omitempty"`
	Nuclease *string `json:"nuclease,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SearchReverseStrand *bool `json:"search_reverse_strand,omitempty"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CrisprHdrDonor is the typed data model for the crispr_hdr_donor entity.
type CrisprHdrDonor struct {
	ArmLength *int `json:"arm_length,omitempty"`
	BlockPam *bool `json:"block_pam,omitempty"`
	DesignGenotypingPrimer *bool `json:"design_genotyping_primer,omitempty"`
	EditEnd *int `json:"edit_end,omitempty"`
	EditStart *int `json:"edit_start,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GuideEnd *int `json:"guide_end,omitempty"`
	GuideStart *int `json:"guide_start,omitempty"`
	GuideStrand *string `json:"guide_strand,omitempty"`
	Nuclease *string `json:"nuclease,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Replacement string `json:"replacement"`
	Result map[string]any `json:"result"`
	TargetSequence string `json:"target_sequence"`
	Tool string `json:"tool"`
}

// CrisprHdrDonorCreateData is the typed request payload for CrisprHdrDonor.CreateTyped.
type CrisprHdrDonorCreateData struct {
	ArmLength *int `json:"arm_length,omitempty"`
	BlockPam *bool `json:"block_pam,omitempty"`
	DesignGenotypingPrimer *bool `json:"design_genotyping_primer,omitempty"`
	EditEnd *int `json:"edit_end,omitempty"`
	EditStart *int `json:"edit_start,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GuideEnd *int `json:"guide_end,omitempty"`
	GuideStart *int `json:"guide_start,omitempty"`
	GuideStrand *string `json:"guide_strand,omitempty"`
	Nuclease *string `json:"nuclease,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Replacement string `json:"replacement"`
	Result map[string]any `json:"result"`
	TargetSequence string `json:"target_sequence"`
	Tool string `json:"tool"`
}

// CrisprOfftargetCheck is the typed data model for the crispr_offtarget_check entity.
type CrisprOfftargetCheck struct {
	Gate *any `json:"gate,omitempty"`
	MaxMismatch *int `json:"max_mismatch,omitempty"`
	Nuclease *string `json:"nuclease,omitempty"`
	Ok any `json:"ok"`
	Protospacer string `json:"protospacer"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// CrisprOfftargetCheckCreateData is the typed request payload for CrisprOfftargetCheck.CreateTyped.
type CrisprOfftargetCheckCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MaxMismatch *int `json:"max_mismatch,omitempty"`
	Nuclease *string `json:"nuclease,omitempty"`
	Ok any `json:"ok"`
	Protospacer string `json:"protospacer"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// CrossDimer is the typed data model for the cross_dimer entity.
type CrossDimer struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SequenceA string `json:"sequence_a"`
	SequenceB string `json:"sequence_b"`
	Tool string `json:"tool"`
}

// CrossDimerCreateData is the typed request payload for CrossDimer.CreateTyped.
type CrossDimerCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SequenceA string `json:"sequence_a"`
	SequenceB string `json:"sequence_b"`
	Tool string `json:"tool"`
}

// DnaMolarity is the typed data model for the dna_molarity entity.
type DnaMolarity struct {
	Gate *any `json:"gate,omitempty"`
	Length *int `json:"length,omitempty"`
	MassNg *float64 `json:"mass_ng,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence *string `json:"sequence,omitempty"`
	Tool string `json:"tool"`
	Type *string `json:"type,omitempty"`
	VolumeUl *float64 `json:"volume_ul,omitempty"`
}

// DnaMolarityCreateData is the typed request payload for DnaMolarity.CreateTyped.
type DnaMolarityCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Length *int `json:"length,omitempty"`
	MassNg *float64 `json:"mass_ng,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence *string `json:"sequence,omitempty"`
	Tool string `json:"tool"`
	Type *string `json:"type,omitempty"`
	VolumeUl *float64 `json:"volume_ul,omitempty"`
}

// DoubleDigest is the typed data model for the double_digest entity.
type DoubleDigest struct {
	EnzymeA string `json:"enzyme_a"`
	EnzymeB string `json:"enzyme_b"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// DoubleDigestCreateData is the typed request payload for DoubleDigest.CreateTyped.
type DoubleDigestCreateData struct {
	EnzymeA string `json:"enzyme_a"`
	EnzymeB string `json:"enzyme_b"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExportEchoPicklist is the typed data model for the export_echo_picklist entity.
type ExportEchoPicklist struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Reaction []any `json:"reaction"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExportEchoPicklistCreateData is the typed request payload for ExportEchoPicklist.CreateTyped.
type ExportEchoPicklistCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Reaction []any `json:"reaction"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExportOpentronsProtocol is the typed data model for the export_opentrons_protocol entity.
type ExportOpentronsProtocol struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	ProtocolName *string `json:"protocol_name,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Reaction []any `json:"reaction"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExportOpentronsProtocolCreateData is the typed request payload for ExportOpentronsProtocol.CreateTyped.
type ExportOpentronsProtocolCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	ProtocolName *string `json:"protocol_name,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Reaction []any `json:"reaction"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExportPlateLayout is the typed data model for the export_plate_layout entity.
type ExportPlateLayout struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Reaction []any `json:"reaction"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExportPlateLayoutCreateData is the typed request payload for ExportPlateLayout.CreateTyped.
type ExportPlateLayoutCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Reaction []any `json:"reaction"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExpressionHeatmapCluster is the typed data model for the expression_heatmap_cluster entity.
type ExpressionHeatmapCluster struct {
	ClusterCol *bool `json:"cluster_col,omitempty"`
	ClusterRow *bool `json:"cluster_row,omitempty"`
	DistanceMetric *string `json:"distance_metric,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Gene []any `json:"gene"`
	Linkage *string `json:"linkage,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sample []any `json:"sample"`
	Tool string `json:"tool"`
	Value []any `json:"value"`
	ZScoreRow *bool `json:"z_score_row,omitempty"`
}

// ExpressionHeatmapClusterCreateData is the typed request payload for ExpressionHeatmapCluster.CreateTyped.
type ExpressionHeatmapClusterCreateData struct {
	ClusterCol *bool `json:"cluster_col,omitempty"`
	ClusterRow *bool `json:"cluster_row,omitempty"`
	DistanceMetric *string `json:"distance_metric,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Gene []any `json:"gene"`
	Linkage *string `json:"linkage,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sample []any `json:"sample"`
	Tool string `json:"tool"`
	Value []any `json:"value"`
	ZScoreRow *bool `json:"z_score_row,omitempty"`
}

// FastqQcReport is the typed data model for the fastq_qc_report entity.
type FastqQcReport struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	QualityOffset *int `json:"quality_offset,omitempty"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// FastqQcReportCreateData is the typed request payload for FastqQcReport.CreateTyped.
type FastqQcReportCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	QualityOffset *int `json:"quality_offset,omitempty"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// FastqTrim is the typed data model for the fastq_trim entity.
type FastqTrim struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	MinLength *int `json:"min_length,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	QualityOffset *int `json:"quality_offset,omitempty"`
	QualityThreshold *int `json:"quality_threshold,omitempty"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// FastqTrimCreateData is the typed request payload for FastqTrim.CreateTyped.
type FastqTrimCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	MinLength *int `json:"min_length,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	QualityOffset *int `json:"quality_offset,omitempty"`
	QualityThreshold *int `json:"quality_threshold,omitempty"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// FindOrf is the typed data model for the find_orf entity.
type FindOrf struct {
	Gate *any `json:"gate,omitempty"`
	MinAaLength *int `json:"min_aa_length,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	RequireStop *bool `json:"require_stop,omitempty"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// FindOrfCreateData is the typed request payload for FindOrf.CreateTyped.
type FindOrfCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MinAaLength *int `json:"min_aa_length,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	RequireStop *bool `json:"require_stop,omitempty"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// FormatSequence is the typed data model for the format_sequence entity.
type FormatSequence struct {
	CaseMode *string `json:"case_mode,omitempty"`
	Convert *string `json:"convert,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Reverse *bool `json:"reverse,omitempty"`
	Sequence string `json:"sequence"`
	StripNonLetter *bool `json:"strip_non_letter,omitempty"`
	Tool string `json:"tool"`
	Width *int `json:"width,omitempty"`
}

// FormatSequenceCreateData is the typed request payload for FormatSequence.CreateTyped.
type FormatSequenceCreateData struct {
	CaseMode *string `json:"case_mode,omitempty"`
	Convert *string `json:"convert,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Reverse *bool `json:"reverse,omitempty"`
	Sequence string `json:"sequence"`
	StripNonLetter *bool `json:"strip_non_letter,omitempty"`
	Tool string `json:"tool"`
	Width *int `json:"width,omitempty"`
}

// FunctionalEnrichment is the typed data model for the functional_enrichment entity.
type FunctionalEnrichment struct {
	Background *[]any `json:"background,omitempty"`
	Collection *[]any `json:"collection,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Gene []any `json:"gene"`
	MaxTermSize *int `json:"max_term_size,omitempty"`
	MinTermSize *int `json:"min_term_size,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// FunctionalEnrichmentCreateData is the typed request payload for FunctionalEnrichment.CreateTyped.
type FunctionalEnrichmentCreateData struct {
	Background *[]any `json:"background,omitempty"`
	Collection *[]any `json:"collection,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Gene []any `json:"gene"`
	MaxTermSize *int `json:"max_term_size,omitempty"`
	MinTermSize *int `json:"min_term_size,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// GcContent is the typed data model for the gc_content entity.
type GcContent struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// GcContentCreateData is the typed request payload for GcContent.CreateTyped.
type GcContentCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// GeneDossier is the typed data model for the gene_dossier entity.
type GeneDossier struct {
	Gate *any `json:"gate,omitempty"`
	Gene string `json:"gene"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// GeneDossierCreateData is the typed request payload for GeneDossier.CreateTyped.
type GeneDossierCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Gene string `json:"gene"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// GeneExpression is the typed data model for the gene_expression entity.
type GeneExpression struct {
	Gate *any `json:"gate,omitempty"`
	Gene string `json:"gene"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// GeneExpressionCreateData is the typed request payload for GeneExpression.CreateTyped.
type GeneExpressionCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Gene string `json:"gene"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// GeneModel is the typed data model for the gene_model entity.
type GeneModel struct {
	Gate *any `json:"gate,omitempty"`
	Gene string `json:"gene"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// GeneModelCreateData is the typed request payload for GeneModel.CreateTyped.
type GeneModelCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Gene string `json:"gene"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// GoldenGateFidelity is the typed data model for the golden_gate_fidelity entity.
type GoldenGateFidelity struct {
	CompareToNamedSet *string `json:"compare_to_named_set,omitempty"`
	Dataset *string `json:"dataset,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Overhang []any `json:"overhang"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	RiskThreshold *float64 `json:"risk_threshold,omitempty"`
	Tool string `json:"tool"`
}

// GoldenGateFidelityCreateData is the typed request payload for GoldenGateFidelity.CreateTyped.
type GoldenGateFidelityCreateData struct {
	CompareToNamedSet *string `json:"compare_to_named_set,omitempty"`
	Dataset *string `json:"dataset,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Overhang []any `json:"overhang"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	RiskThreshold *float64 `json:"risk_threshold,omitempty"`
	Tool string `json:"tool"`
}

// HgvsConvert is the typed data model for the hgvs_convert entity.
type HgvsConvert struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Variant string `json:"variant"`
}

// HgvsConvertCreateData is the typed request payload for HgvsConvert.CreateTyped.
type HgvsConvertCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Variant string `json:"variant"`
}

// IdMapPoll is the typed data model for the id_map_poll entity.
type IdMapPoll struct {
	Gate *any `json:"gate,omitempty"`
	JobId string `json:"job_id"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// IdMapPollCreateData is the typed request payload for IdMapPoll.CreateTyped.
type IdMapPollCreateData struct {
	Gate *any `json:"gate,omitempty"`
	JobId string `json:"job_id"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// IdMapSubmit is the typed data model for the id_map_submit entity.
type IdMapSubmit struct {
	From string `json:"from"`
	Gate *any `json:"gate,omitempty"`
	Ids []any `json:"ids"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	TaxId *string `json:"tax_id,omitempty"`
	To string `json:"to"`
	Tool string `json:"tool"`
}

// IdMapSubmitCreateData is the typed request payload for IdMapSubmit.CreateTyped.
type IdMapSubmitCreateData struct {
	From string `json:"from"`
	Gate *any `json:"gate,omitempty"`
	Ids []any `json:"ids"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	TaxId *string `json:"tax_id,omitempty"`
	To string `json:"to"`
	Tool string `json:"tool"`
}

// InSilicoPcr is the typed data model for the in_silico_pcr entity.
type InSilicoPcr struct {
	Circular *bool `json:"circular,omitempty"`
	ForwardPrimer string `json:"forward_primer"`
	Gate *any `json:"gate,omitempty"`
	MaxMismatch *int `json:"max_mismatch,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ReversePrimer string `json:"reverse_primer"`
	Template string `json:"template"`
	Tool string `json:"tool"`
}

// InSilicoPcrCreateData is the typed request payload for InSilicoPcr.CreateTyped.
type InSilicoPcrCreateData struct {
	Circular *bool `json:"circular,omitempty"`
	ForwardPrimer string `json:"forward_primer"`
	Gate *any `json:"gate,omitempty"`
	MaxMismatch *int `json:"max_mismatch,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ReversePrimer string `json:"reverse_primer"`
	Template string `json:"template"`
	Tool string `json:"tool"`
}

// KaspPrimerDesign is the typed data model for the kasp_primer_design entity.
type KaspPrimerDesign struct {
	AddSecondaryMismatch *bool `json:"add_secondary_mismatch,omitempty"`
	AlleleA string `json:"allele_a"`
	AlleleB string `json:"allele_b"`
	Gate *any `json:"gate,omitempty"`
	MaxAmplicon *int `json:"max_amplicon,omitempty"`
	MinAmplicon *int `json:"min_amplicon,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SnpPosition int `json:"snp_position"`
	Target string `json:"target"`
	TargetCoreTm *float64 `json:"target_core_tm,omitempty"`
	Tool string `json:"tool"`
}

// KaspPrimerDesignCreateData is the typed request payload for KaspPrimerDesign.CreateTyped.
type KaspPrimerDesignCreateData struct {
	AddSecondaryMismatch *bool `json:"add_secondary_mismatch,omitempty"`
	AlleleA string `json:"allele_a"`
	AlleleB string `json:"allele_b"`
	Gate *any `json:"gate,omitempty"`
	MaxAmplicon *int `json:"max_amplicon,omitempty"`
	MinAmplicon *int `json:"min_amplicon,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SnpPosition int `json:"snp_position"`
	Target string `json:"target"`
	TargetCoreTm *float64 `json:"target_core_tm,omitempty"`
	Tool string `json:"tool"`
}

// ListTool is the typed data model for the list_tool entity.
type ListTool struct {
}

// ListToolLoadMatch is the typed request payload for ListTool.LoadTyped.
type ListToolLoadMatch struct {
}

// MeltingTemperature is the typed data model for the melting_temperature entity.
type MeltingTemperature struct {
	DntpMm *float64 `json:"dntp_mm,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMm *float64 `json:"mg_mm,omitempty"`
	NaMm *float64 `json:"na_mm,omitempty"`
	Ok any `json:"ok"`
	OligoNm *float64 `json:"oligo_nm,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	TargetTm *float64 `json:"target_tm,omitempty"`
	TmTolerance *float64 `json:"tm_tolerance,omitempty"`
	Tool string `json:"tool"`
}

// MeltingTemperatureCreateData is the typed request payload for MeltingTemperature.CreateTyped.
type MeltingTemperatureCreateData struct {
	DntpMm *float64 `json:"dntp_mm,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMm *float64 `json:"mg_mm,omitempty"`
	NaMm *float64 `json:"na_mm,omitempty"`
	Ok any `json:"ok"`
	OligoNm *float64 `json:"oligo_nm,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	TargetTm *float64 `json:"target_tm,omitempty"`
	TmTolerance *float64 `json:"tm_tolerance,omitempty"`
	Tool string `json:"tool"`
}

// MotifFinder is the typed data model for the motif_finder entity.
type MotifFinder struct {
	Gate *any `json:"gate,omitempty"`
	MaxMismatch *int `json:"max_mismatch,omitempty"`
	Motif string `json:"motif"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SearchReverseStrand *bool `json:"search_reverse_strand,omitempty"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// MotifFinderCreateData is the typed request payload for MotifFinder.CreateTyped.
type MotifFinderCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MaxMismatch *int `json:"max_mismatch,omitempty"`
	Motif string `json:"motif"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SearchReverseStrand *bool `json:"search_reverse_strand,omitempty"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// MultipleSequenceAlignment is the typed data model for the multiple_sequence_alignment entity.
type MultipleSequenceAlignment struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// MultipleSequenceAlignmentCreateData is the typed request payload for MultipleSequenceAlignment.CreateTyped.
type MultipleSequenceAlignmentCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// OligoAnalysi is the typed data model for the oligo_analysi entity.
type OligoAnalysi struct {
	DntpMm *float64 `json:"dntp_mm,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMm *float64 `json:"mg_mm,omitempty"`
	NaMm *float64 `json:"na_mm,omitempty"`
	Ok any `json:"ok"`
	OligoNm *float64 `json:"oligo_nm,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// OligoAnalysiCreateData is the typed request payload for OligoAnalysi.CreateTyped.
type OligoAnalysiCreateData struct {
	DntpMm *float64 `json:"dntp_mm,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMm *float64 `json:"mg_mm,omitempty"`
	NaMm *float64 `json:"na_mm,omitempty"`
	Ok any `json:"ok"`
	OligoNm *float64 `json:"oligo_nm,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// OrthologMap is the typed data model for the ortholog_map entity.
type OrthologMap struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SourceSpecies *string `json:"source_species,omitempty"`
	Symbol []any `json:"symbol"`
	TargetSpecies string `json:"target_species"`
	Tool string `json:"tool"`
	Type *string `json:"type,omitempty"`
}

// OrthologMapCreateData is the typed request payload for OrthologMap.CreateTyped.
type OrthologMapCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SourceSpecies *string `json:"source_species,omitempty"`
	Symbol []any `json:"symbol"`
	TargetSpecies string `json:"target_species"`
	Tool string `json:"tool"`
	Type *string `json:"type,omitempty"`
}

// PairwiseAlignment is the typed data model for the pairwise_alignment entity.
type PairwiseAlignment struct {
	Gap *float64 `json:"gap,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Match *float64 `json:"match,omitempty"`
	Mismatch *float64 `json:"mismatch,omitempty"`
	Mode *string `json:"mode,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SeqA string `json:"seq_a"`
	SeqB string `json:"seq_b"`
	Tool string `json:"tool"`
}

// PairwiseAlignmentCreateData is the typed request payload for PairwiseAlignment.CreateTyped.
type PairwiseAlignmentCreateData struct {
	Gap *float64 `json:"gap,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Match *float64 `json:"match,omitempty"`
	Mismatch *float64 `json:"mismatch,omitempty"`
	Mode *string `json:"mode,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SeqA string `json:"seq_a"`
	SeqB string `json:"seq_b"`
	Tool string `json:"tool"`
}

// ParseGenbank is the typed data model for the parse_genbank entity.
type ParseGenbank struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Text string `json:"text"`
	Tool string `json:"tool"`
}

// ParseGenbankCreateData is the typed request payload for ParseGenbank.CreateTyped.
type ParseGenbankCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Text string `json:"text"`
	Tool string `json:"tool"`
}

// ParseSangerTrace is the typed data model for the parse_sanger_trace entity.
type ParseSangerTrace struct {
	FileBase64 string `json:"file_base64"`
	FileName *string `json:"file_name,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ParseSangerTraceCreateData is the typed request payload for ParseSangerTrace.CreateTyped.
type ParseSangerTraceCreateData struct {
	FileBase64 string `json:"file_base64"`
	FileName *string `json:"file_name,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// PlasmidAnnotate is the typed data model for the plasmid_annotate entity.
type PlasmidAnnotate struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// PlasmidAnnotateCreateData is the typed request payload for PlasmidAnnotate.CreateTyped.
type PlasmidAnnotateCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// PlasmidDeepAnnotate is the typed data model for the plasmid_deep_annotate entity.
type PlasmidDeepAnnotate struct {
	Circular *bool `json:"circular,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// PlasmidDeepAnnotateCreateData is the typed request payload for PlasmidDeepAnnotate.CreateTyped.
type PlasmidDeepAnnotateCreateData struct {
	Circular *bool `json:"circular,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// PlasmidFullReport is the typed data model for the plasmid_full_report entity.
type PlasmidFullReport struct {
	Circular *bool `json:"circular,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
	TopN *int `json:"top_n,omitempty"`
}

// PlasmidFullReportCreateData is the typed request payload for PlasmidFullReport.CreateTyped.
type PlasmidFullReportCreateData struct {
	Circular *bool `json:"circular,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
	TopN *int `json:"top_n,omitempty"`
}

// PlasmidIdentify is the typed data model for the plasmid_identify entity.
type PlasmidIdentify struct {
	Circular *bool `json:"circular,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
	TopN *int `json:"top_n,omitempty"`
}

// PlasmidIdentifyCreateData is the typed request payload for PlasmidIdentify.CreateTyped.
type PlasmidIdentifyCreateData struct {
	Circular *bool `json:"circular,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
	TopN *int `json:"top_n,omitempty"`
}

// PrimeEditingDesign is the typed data model for the prime_editing_design entity.
type PrimeEditingDesign struct {
	EditEnd int `json:"edit_end"`
	EditStart int `json:"edit_start"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	InsertedSeq *string `json:"inserted_seq,omitempty"`
	Ok any `json:"ok"`
	PbsLength *int `json:"pbs_length,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	RttHomology *int `json:"rtt_homology,omitempty"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// PrimeEditingDesignCreateData is the typed request payload for PrimeEditingDesign.CreateTyped.
type PrimeEditingDesignCreateData struct {
	EditEnd int `json:"edit_end"`
	EditStart int `json:"edit_start"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	InsertedSeq *string `json:"inserted_seq,omitempty"`
	Ok any `json:"ok"`
	PbsLength *int `json:"pbs_length,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	RttHomology *int `json:"rtt_homology,omitempty"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// PrimeEditingTwinDesign is the typed data model for the prime_editing_twin_design entity.
type PrimeEditingTwinDesign struct {
	Gate *any `json:"gate,omitempty"`
	NewSequence string `json:"new_sequence"`
	Ok any `json:"ok"`
	OverlapLength *int `json:"overlap_length,omitempty"`
	PbsLength *int `json:"pbs_length,omitempty"`
	Provenance map[string]any `json:"provenance"`
	ReplaceEnd int `json:"replace_end"`
	ReplaceStart int `json:"replace_start"`
	Result map[string]any `json:"result"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// PrimeEditingTwinDesignCreateData is the typed request payload for PrimeEditingTwinDesign.CreateTyped.
type PrimeEditingTwinDesignCreateData struct {
	Gate *any `json:"gate,omitempty"`
	NewSequence string `json:"new_sequence"`
	Ok any `json:"ok"`
	OverlapLength *int `json:"overlap_length,omitempty"`
	PbsLength *int `json:"pbs_length,omitempty"`
	Provenance map[string]any `json:"provenance"`
	ReplaceEnd int `json:"replace_end"`
	ReplaceStart int `json:"replace_start"`
	Result map[string]any `json:"result"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// PrimerDesign is the typed data model for the primer_design entity.
type PrimerDesign struct {
	AmpliconMax *int `json:"amplicon_max,omitempty"`
	AmpliconMin *int `json:"amplicon_min,omitempty"`
	DntpMm *float64 `json:"dntp_mm,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcMax *float64 `json:"gc_max,omitempty"`
	GcMin *float64 `json:"gc_min,omitempty"`
	LenMax *int `json:"len_max,omitempty"`
	LenMin *int `json:"len_min,omitempty"`
	LenOpt *int `json:"len_opt,omitempty"`
	MaxReturn *int `json:"max_return,omitempty"`
	MgMm *float64 `json:"mg_mm,omitempty"`
	NaMm *float64 `json:"na_mm,omitempty"`
	Ok any `json:"ok"`
	OligoNm *float64 `json:"oligo_nm,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	TargetEnd *int `json:"target_end,omitempty"`
	TargetStart *int `json:"target_start,omitempty"`
	Template string `json:"template"`
	TmMax *float64 `json:"tm_max,omitempty"`
	TmMaxDiff *float64 `json:"tm_max_diff,omitempty"`
	TmMin *float64 `json:"tm_min,omitempty"`
	TmOpt *float64 `json:"tm_opt,omitempty"`
	Tool string `json:"tool"`
}

// PrimerDesignCreateData is the typed request payload for PrimerDesign.CreateTyped.
type PrimerDesignCreateData struct {
	AmpliconMax *int `json:"amplicon_max,omitempty"`
	AmpliconMin *int `json:"amplicon_min,omitempty"`
	DntpMm *float64 `json:"dntp_mm,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcMax *float64 `json:"gc_max,omitempty"`
	GcMin *float64 `json:"gc_min,omitempty"`
	LenMax *int `json:"len_max,omitempty"`
	LenMin *int `json:"len_min,omitempty"`
	LenOpt *int `json:"len_opt,omitempty"`
	MaxReturn *int `json:"max_return,omitempty"`
	MgMm *float64 `json:"mg_mm,omitempty"`
	NaMm *float64 `json:"na_mm,omitempty"`
	Ok any `json:"ok"`
	OligoNm *float64 `json:"oligo_nm,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	TargetEnd *int `json:"target_end,omitempty"`
	TargetStart *int `json:"target_start,omitempty"`
	Template string `json:"template"`
	TmMax *float64 `json:"tm_max,omitempty"`
	TmMaxDiff *float64 `json:"tm_max_diff,omitempty"`
	TmMin *float64 `json:"tm_min,omitempty"`
	TmOpt *float64 `json:"tm_opt,omitempty"`
	Tool string `json:"tool"`
}

// PrimerSpecificity is the typed data model for the primer_specificity entity.
type PrimerSpecificity struct {
	ForwardPrimer string `json:"forward_primer"`
	Gate *any `json:"gate,omitempty"`
	MaxMismatch *int `json:"max_mismatch,omitempty"`
	MaxProductLength *int `json:"max_product_length,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ReversePrimer string `json:"reverse_primer"`
	Tool string `json:"tool"`
}

// PrimerSpecificityCreateData is the typed request payload for PrimerSpecificity.CreateTyped.
type PrimerSpecificityCreateData struct {
	ForwardPrimer string `json:"forward_primer"`
	Gate *any `json:"gate,omitempty"`
	MaxMismatch *int `json:"max_mismatch,omitempty"`
	MaxProductLength *int `json:"max_product_length,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ReversePrimer string `json:"reverse_primer"`
	Tool string `json:"tool"`
}

// ProteaseDigestion is the typed data model for the protease_digestion entity.
type ProteaseDigestion struct {
	Gate *any `json:"gate,omitempty"`
	MaxMass *float64 `json:"max_mass,omitempty"`
	MaxPeptide *int `json:"max_peptide,omitempty"`
	MinMass *float64 `json:"min_mass,omitempty"`
	MissedCleavage *int `json:"missed_cleavage,omitempty"`
	Ok any `json:"ok"`
	Protease *string `json:"protease,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ProteaseDigestionCreateData is the typed request payload for ProteaseDigestion.CreateTyped.
type ProteaseDigestionCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MaxMass *float64 `json:"max_mass,omitempty"`
	MaxPeptide *int `json:"max_peptide,omitempty"`
	MinMass *float64 `json:"min_mass,omitempty"`
	MissedCleavage *int `json:"missed_cleavage,omitempty"`
	Ok any `json:"ok"`
	Protease *string `json:"protease,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ProteinAnnotatePoll is the typed data model for the protein_annotate_poll entity.
type ProteinAnnotatePoll struct {
	Gate *any `json:"gate,omitempty"`
	JobId string `json:"job_id"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ProteinAnnotatePollCreateData is the typed request payload for ProteinAnnotatePoll.CreateTyped.
type ProteinAnnotatePollCreateData struct {
	Gate *any `json:"gate,omitempty"`
	JobId string `json:"job_id"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ProteinAnnotateSubmit is the typed data model for the protein_annotate_submit entity.
type ProteinAnnotateSubmit struct {
	Appl *string `json:"appl,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Goterm *bool `json:"goterm,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ProteinAnnotateSubmitCreateData is the typed request payload for ProteinAnnotateSubmit.CreateTyped.
type ProteinAnnotateSubmitCreateData struct {
	Appl *string `json:"appl,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Goterm *bool `json:"goterm,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ProteinHydrophobicity is the typed data model for the protein_hydrophobicity entity.
type ProteinHydrophobicity struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Scale *string `json:"scale,omitempty"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
	Window *int `json:"window,omitempty"`
}

// ProteinHydrophobicityCreateData is the typed request payload for ProteinHydrophobicity.CreateTyped.
type ProteinHydrophobicityCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Scale *string `json:"scale,omitempty"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
	Window *int `json:"window,omitempty"`
}

// ProteinProperty is the typed data model for the protein_property entity.
type ProteinProperty struct {
	ChargeStep *float64 `json:"charge_step,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ProteinPropertyCreateData is the typed request payload for ProteinProperty.CreateTyped.
type ProteinPropertyCreateData struct {
	ChargeStep *float64 `json:"charge_step,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// RandomSequence is the typed data model for the random_sequence entity.
type RandomSequence struct {
	Gate *any `json:"gate,omitempty"`
	GcContent *float64 `json:"gc_content,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Length int `json:"length"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// RandomSequenceCreateData is the typed request payload for RandomSequence.CreateTyped.
type RandomSequenceCreateData struct {
	Gate *any `json:"gate,omitempty"`
	GcContent *float64 `json:"gc_content,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Length int `json:"length"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// RestrictionSite is the typed data model for the restriction_site entity.
type RestrictionSite struct {
	Enzyme *[]any `json:"enzyme,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// RestrictionSiteCreateData is the typed request payload for RestrictionSite.CreateTyped.
type RestrictionSiteCreateData struct {
	Enzyme *[]any `json:"enzyme,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ReverseComplement is the typed data model for the reverse_complement entity.
type ReverseComplement struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
	Type *string `json:"type,omitempty"`
}

// ReverseComplementCreateData is the typed request payload for ReverseComplement.CreateTyped.
type ReverseComplementCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
	Type *string `json:"type,omitempty"`
}

// ReverseTranslate is the typed data model for the reverse_translate entity.
type ReverseTranslate struct {
	Gate *any `json:"gate,omitempty"`
	Mode *string `json:"mode,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Protein string `json:"protein"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ReverseTranslateCreateData is the typed request payload for ReverseTranslate.CreateTyped.
type ReverseTranslateCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Mode *string `json:"mode,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Protein string `json:"protein"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// RnaFold is the typed data model for the rna_fold entity.
type RnaFold struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// RnaFoldCreateData is the typed request payload for RnaFold.CreateTyped.
type RnaFoldCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// SangerVsReference is the typed data model for the sanger_vs_reference entity.
type SangerVsReference struct {
	FileBase64 *string `json:"file_base64,omitempty"`
	FileName *string `json:"file_name,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MinCoverage *float64 `json:"min_coverage,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Read *string `json:"read,omitempty"`
	Reference string `json:"reference"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SangerVsReferenceCreateData is the typed request payload for SangerVsReference.CreateTyped.
type SangerVsReferenceCreateData struct {
	FileBase64 *string `json:"file_base64,omitempty"`
	FileName *string `json:"file_name,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MinCoverage *float64 `json:"min_coverage,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Read *string `json:"read,omitempty"`
	Reference string `json:"reference"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SavePermalink is the typed data model for the save_permalink entity.
type SavePermalink struct {
	Arg map[string]any `json:"arg"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SavePermalinkCreateData is the typed request payload for SavePermalink.CreateTyped.
type SavePermalinkCreateData struct {
	Arg map[string]any `json:"arg"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SeqfileStat is the typed data model for the seqfile_stat entity.
type SeqfileStat struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	QualityOffset *int `json:"quality_offset,omitempty"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SeqfileStatCreateData is the typed request payload for SeqfileStat.CreateTyped.
type SeqfileStatCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	QualityOffset *int `json:"quality_offset,omitempty"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SequenceFetch is the typed data model for the sequence_fetch entity.
type SequenceFetch struct {
	Accession string `json:"accession"`
	Db *string `json:"db,omitempty"`
	Format *string `json:"format,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SequenceFetchCreateData is the typed request payload for SequenceFetch.CreateTyped.
type SequenceFetchCreateData struct {
	Accession string `json:"accession"`
	Db *string `json:"db,omitempty"`
	Format *string `json:"format,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SequenceFormatConvert is the typed data model for the sequence_format_convert entity.
type SequenceFormatConvert struct {
	From *string `json:"from,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	To *string `json:"to,omitempty"`
	Tool string `json:"tool"`
}

// SequenceFormatConvertCreateData is the typed request payload for SequenceFormatConvert.CreateTyped.
type SequenceFormatConvertCreateData struct {
	From *string `json:"from,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	To *string `json:"to,omitempty"`
	Tool string `json:"tool"`
}

// SequenceReport is the typed data model for the sequence_report entity.
type SequenceReport struct {
	EndPrimerLength *int `json:"end_primer_length,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MaxOrf *int `json:"max_orf,omitempty"`
	MinOrfAa *int `json:"min_orf_aa,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// SequenceReportCreateData is the typed request payload for SequenceReport.CreateTyped.
type SequenceReportCreateData struct {
	EndPrimerLength *int `json:"end_primer_length,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MaxOrf *int `json:"max_orf,omitempty"`
	MinOrfAa *int `json:"min_orf_aa,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// SequenceSearch is the typed data model for the sequence_search entity.
type SequenceSearch struct {
	Db *string `json:"db,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Gene *string `json:"gene,omitempty"`
	MaxResult *int `json:"max_result,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Term *string `json:"term,omitempty"`
	Tool string `json:"tool"`
}

// SequenceSearchCreateData is the typed request payload for SequenceSearch.CreateTyped.
type SequenceSearchCreateData struct {
	Db *string `json:"db,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Gene *string `json:"gene,omitempty"`
	MaxResult *int `json:"max_result,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Term *string `json:"term,omitempty"`
	Tool string `json:"tool"`
}

// SequencingReadbackVerify is the typed data model for the sequencing_readback_verify entity.
type SequencingReadbackVerify struct {
	Gate *any `json:"gate,omitempty"`
	MinSupportingRead *int `json:"min_supporting_read,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Read string `json:"read"`
	Reference string `json:"reference"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SequencingReadbackVerifyCreateData is the typed request payload for SequencingReadbackVerify.CreateTyped.
type SequencingReadbackVerifyCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MinSupportingRead *int `json:"min_supporting_read,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Read string `json:"read"`
	Reference string `json:"reference"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SessionCreate is the typed data model for the session_create entity.
type SessionCreate struct {
	Entry *map[string]any `json:"entry,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SessionCreateCreateData is the typed request payload for SessionCreate.CreateTyped.
type SessionCreateCreateData struct {
	Entry *map[string]any `json:"entry,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SessionGet is the typed data model for the session_get entity.
type SessionGet struct {
	Gate *any `json:"gate,omitempty"`
	Name *[]any `json:"name,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"session_id"`
	Tool string `json:"tool"`
}

// SessionGetCreateData is the typed request payload for SessionGet.CreateTyped.
type SessionGetCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Name *[]any `json:"name,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"session_id"`
	Tool string `json:"tool"`
}

// SessionRun is the typed data model for the session_run entity.
type SessionRun struct {
	Arg *map[string]any `json:"arg,omitempty"`
	FromSession *map[string]any `json:"from_session,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"session_id"`
	Tool string `json:"tool"`
	WriteBack *map[string]any `json:"write_back,omitempty"`
}

// SessionRunCreateData is the typed request payload for SessionRun.CreateTyped.
type SessionRunCreateData struct {
	Arg *map[string]any `json:"arg,omitempty"`
	FromSession *map[string]any `json:"from_session,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"session_id"`
	Tool string `json:"tool"`
	WriteBack *map[string]any `json:"write_back,omitempty"`
}

// SessionSet is the typed data model for the session_set entity.
type SessionSet struct {
	Entry map[string]any `json:"entry"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"session_id"`
	Tool string `json:"tool"`
}

// SessionSetCreateData is the typed request payload for SessionSet.CreateTyped.
type SessionSetCreateData struct {
	Entry map[string]any `json:"entry"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"session_id"`
	Tool string `json:"tool"`
}

// SirnaDesign is the typed data model for the sirna_design entity.
type SirnaDesign struct {
	Gate *any `json:"gate,omitempty"`
	MinReynold *int `json:"min_reynold,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ShRnaLoop *string `json:"sh_rna_loop,omitempty"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// SirnaDesignCreateData is the typed request payload for SirnaDesign.CreateTyped.
type SirnaDesignCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MinReynold *int `json:"min_reynold,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ShRnaLoop *string `json:"sh_rna_loop,omitempty"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// SiteDirectedMutagenesi is the typed data model for the site_directed_mutagenesi entity.
type SiteDirectedMutagenesi struct {
	ArmTmTarget *float64 `json:"arm_tm_target,omitempty"`
	DntpMm *float64 `json:"dntp_mm,omitempty"`
	EditKind *string `json:"edit_kind,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMm *float64 `json:"mg_mm,omitempty"`
	NaMm *float64 `json:"na_mm,omitempty"`
	NewBase *string `json:"new_base,omitempty"`
	Ok any `json:"ok"`
	OligoNm *float64 `json:"oligo_nm,omitempty"`
	Organism *string `json:"organism,omitempty"`
	Position *int `json:"position,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Residue *int `json:"residue,omitempty"`
	Result map[string]any `json:"result"`
	Style *string `json:"style,omitempty"`
	TargetAa *string `json:"target_aa,omitempty"`
	Template string `json:"template"`
	Tool string `json:"tool"`
}

// SiteDirectedMutagenesiCreateData is the typed request payload for SiteDirectedMutagenesi.CreateTyped.
type SiteDirectedMutagenesiCreateData struct {
	ArmTmTarget *float64 `json:"arm_tm_target,omitempty"`
	DntpMm *float64 `json:"dntp_mm,omitempty"`
	EditKind *string `json:"edit_kind,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMm *float64 `json:"mg_mm,omitempty"`
	NaMm *float64 `json:"na_mm,omitempty"`
	NewBase *string `json:"new_base,omitempty"`
	Ok any `json:"ok"`
	OligoNm *float64 `json:"oligo_nm,omitempty"`
	Organism *string `json:"organism,omitempty"`
	Position *int `json:"position,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Residue *int `json:"residue,omitempty"`
	Result map[string]any `json:"result"`
	Style *string `json:"style,omitempty"`
	TargetAa *string `json:"target_aa,omitempty"`
	Template string `json:"template"`
	Tool string `json:"tool"`
}

// Translate is the typed data model for the translate entity.
type Translate struct {
	Frame *int `json:"frame,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	ToStop *bool `json:"to_stop,omitempty"`
	Tool string `json:"tool"`
}

// TranslateCreateData is the typed request payload for Translate.CreateTyped.
type TranslateCreateData struct {
	Frame *int `json:"frame,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	ToStop *bool `json:"to_stop,omitempty"`
	Tool string `json:"tool"`
}

// VariantAnnotate is the typed data model for the variant_annotate entity.
type VariantAnnotate struct {
	Assembly *string `json:"assembly,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Variant string `json:"variant"`
}

// VariantAnnotateCreateData is the typed request payload for VariantAnnotate.CreateTyped.
type VariantAnnotateCreateData struct {
	Assembly *string `json:"assembly,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Variant string `json:"variant"`
}

// VariantComparator is the typed data model for the variant_comparator entity.
type VariantComparator struct {
	Coding *bool `json:"coding,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Query string `json:"query"`
	Reference string `json:"reference"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// VariantComparatorCreateData is the typed request payload for VariantComparator.CreateTyped.
type VariantComparatorCreateData struct {
	Coding *bool `json:"coding,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Query string `json:"query"`
	Reference string `json:"reference"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// VerifyAssembly is the typed data model for the verify_assembly entity.
type VerifyAssembly struct {
	ArmTmTarget *float64 `json:"arm_tm_target,omitempty"`
	Circular *bool `json:"circular,omitempty"`
	ClaimedConstruct string `json:"claimed_construct"`
	Coding *bool `json:"coding,omitempty"`
	Enzyme *string `json:"enzyme,omitempty"`
	Enzyme3 *string `json:"enzyme3,omitempty"`
	Enzyme5 *string `json:"enzyme5,omitempty"`
	Fragment *[]any `json:"fragment,omitempty"`
	FragmentPcr *[]any `json:"fragment_pcr,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Insert *string `json:"insert,omitempty"`
	InsertPcr *map[string]any `json:"insert_pcr,omitempty"`
	Method string `json:"method"`
	Name *[]any `json:"name,omitempty"`
	Ok any `json:"ok"`
	OverlapLen *int `json:"overlap_len,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Vector *string `json:"vector,omitempty"`
	VectorPcr *map[string]any `json:"vector_pcr,omitempty"`
}

// VerifyAssemblyCreateData is the typed request payload for VerifyAssembly.CreateTyped.
type VerifyAssemblyCreateData struct {
	ArmTmTarget *float64 `json:"arm_tm_target,omitempty"`
	Circular *bool `json:"circular,omitempty"`
	ClaimedConstruct string `json:"claimed_construct"`
	Coding *bool `json:"coding,omitempty"`
	Enzyme *string `json:"enzyme,omitempty"`
	Enzyme3 *string `json:"enzyme3,omitempty"`
	Enzyme5 *string `json:"enzyme5,omitempty"`
	Fragment *[]any `json:"fragment,omitempty"`
	FragmentPcr *[]any `json:"fragment_pcr,omitempty"`
	FrameStart *int `json:"frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Insert *string `json:"insert,omitempty"`
	InsertPcr *map[string]any `json:"insert_pcr,omitempty"`
	Method string `json:"method"`
	Name *[]any `json:"name,omitempty"`
	Ok any `json:"ok"`
	OverlapLen *int `json:"overlap_len,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Vector *string `json:"vector,omitempty"`
	VectorPcr *map[string]any `json:"vector_pcr,omitempty"`
}

// VerifyConstruct is the typed data model for the verify_construct entity.
type VerifyConstruct struct {
	ClaimedConstruct string `json:"claimed_construct"`
	ExpectedFrameStart *int `json:"expected_frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	InsertForwardPrimer string `json:"insert_forward_primer"`
	InsertReversePrimer string `json:"insert_reverse_primer"`
	InsertTemplate string `json:"insert_template"`
	MaxPrimerMismatch *int `json:"max_primer_mismatch,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	TemplateCircular *bool `json:"template_circular,omitempty"`
	Tool string `json:"tool"`
}

// VerifyConstructCreateData is the typed request payload for VerifyConstruct.CreateTyped.
type VerifyConstructCreateData struct {
	ClaimedConstruct string `json:"claimed_construct"`
	ExpectedFrameStart *int `json:"expected_frame_start,omitempty"`
	Gate *any `json:"gate,omitempty"`
	InsertForwardPrimer string `json:"insert_forward_primer"`
	InsertReversePrimer string `json:"insert_reverse_primer"`
	InsertTemplate string `json:"insert_template"`
	MaxPrimerMismatch *int `json:"max_primer_mismatch,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	TemplateCircular *bool `json:"template_circular,omitempty"`
	Tool string `json:"tool"`
}

// VirtualGel is the typed data model for the virtual_gel entity.
type VirtualGel struct {
	Circular *bool `json:"circular,omitempty"`
	Enzyme *[]any `json:"enzyme,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ladder *string `json:"ladder,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// VirtualGelCreateData is the typed request payload for VirtualGel.CreateTyped.
type VirtualGelCreateData struct {
	Circular *bool `json:"circular,omitempty"`
	Enzyme *[]any `json:"enzyme,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ladder *string `json:"ladder,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// VolcanoPlotData is the typed data model for the volcano_plot_data entity.
type VolcanoPlotData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Row []any `json:"row"`
	Tool string `json:"tool"`
}

// VolcanoPlotDataCreateData is the typed request payload for VolcanoPlotData.CreateTyped.
type VolcanoPlotDataCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Row []any `json:"row"`
	Tool string `json:"tool"`
}

// WebSearch is the typed data model for the web_search entity.
type WebSearch struct {
	Gate *any `json:"gate,omitempty"`
	MaxResult *float64 `json:"max_result,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Query string `json:"query"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// WebSearchCreateData is the typed request payload for WebSearch.CreateTyped.
type WebSearchCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MaxResult *float64 `json:"max_result,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Query string `json:"query"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
