// Typed models for the SeqbenchMcp SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/seqbench-mcp-sdk/go/core"
)

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
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Target string `json:"target"`
	TargetPosition *int `json:"targetPosition,omitempty"`
	Tool string `json:"tool"`
}

// BaseEditingDesignCreateData is the typed request payload for BaseEditingDesign.CreateTyped.
type BaseEditingDesignCreateData struct {
	Editor *string `json:"editor,omitempty"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Target string `json:"target"`
	TargetPosition *int `json:"targetPosition,omitempty"`
	Tool string `json:"tool"`
}

// Batch is the typed data model for the batch entity.
type Batch struct {
	Args *map[string]any `json:"args,omitempty"`
	Capped bool `json:"capped"`
	Columns []any `json:"columns"`
	Count int `json:"count"`
	Errors int `json:"errors"`
	Input string `json:"input"`
	Limit int `json:"limit"`
	Provenance map[string]any `json:"provenance"`
	Rows []any `json:"rows"`
	Tool string `json:"tool"`
}

// BatchLoadMatch is the typed request payload for Batch.LoadTyped.
type BatchLoadMatch struct {
	Args *map[string]any `json:"args,omitempty"`
	Capped *bool `json:"capped,omitempty"`
	Columns *[]any `json:"columns,omitempty"`
	Count *int `json:"count,omitempty"`
	Errors *int `json:"errors,omitempty"`
	Input *string `json:"input,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Provenance *map[string]any `json:"provenance,omitempty"`
	Rows *[]any `json:"rows,omitempty"`
	Tool *string `json:"tool,omitempty"`
}

// BatchCreateData is the typed request payload for Batch.CreateTyped.
type BatchCreateData struct {
	Args *map[string]any `json:"args,omitempty"`
	Capped bool `json:"capped"`
	Columns []any `json:"columns"`
	Count int `json:"count"`
	Errors int `json:"errors"`
	Input string `json:"input"`
	Limit int `json:"limit"`
	Provenance map[string]any `json:"provenance"`
	Rows []any `json:"rows"`
	Tool string `json:"tool"`
}

// BatchWorkflow is the typed data model for the batch__workflow entity.
type BatchWorkflow struct {
	Capped bool `json:"capped"`
	Columns []any `json:"columns"`
	Count int `json:"count"`
	Errors int `json:"errors"`
	Input string `json:"input"`
	Limit int `json:"limit"`
	Provenance map[string]any `json:"provenance"`
	Rows []any `json:"rows"`
	Steps []any `json:"steps"`
}

// BatchWorkflowLoadMatch is the typed request payload for BatchWorkflow.LoadTyped.
type BatchWorkflowLoadMatch struct {
	Capped *bool `json:"capped,omitempty"`
	Columns *[]any `json:"columns,omitempty"`
	Count *int `json:"count,omitempty"`
	Errors *int `json:"errors,omitempty"`
	Input *string `json:"input,omitempty"`
	Limit *int `json:"limit,omitempty"`
	Provenance *map[string]any `json:"provenance,omitempty"`
	Rows *[]any `json:"rows,omitempty"`
	Steps *[]any `json:"steps,omitempty"`
}

// BatchWorkflowCreateData is the typed request payload for BatchWorkflow.CreateTyped.
type BatchWorkflowCreateData struct {
	Capped bool `json:"capped"`
	Columns []any `json:"columns"`
	Count int `json:"count"`
	Errors int `json:"errors"`
	Input string `json:"input"`
	Limit int `json:"limit"`
	Provenance map[string]any `json:"provenance"`
	Rows []any `json:"rows"`
	Steps []any `json:"steps"`
}

// CharacterizeSequence is the typed data model for the characterize_sequence entity.
type CharacterizeSequence struct {
	EndPrimerLength *int `json:"endPrimerLength,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MaxOrfs *int `json:"maxOrfs,omitempty"`
	MinOrfAa *int `json:"minOrfAa,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CharacterizeSequenceCreateData is the typed request payload for CharacterizeSequence.CreateTyped.
type CharacterizeSequenceCreateData struct {
	EndPrimerLength *int `json:"endPrimerLength,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MaxOrfs *int `json:"maxOrfs,omitempty"`
	MinOrfAa *int `json:"minOrfAa,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CloningSimulate is the typed data model for the cloning_simulate entity.
type CloningSimulate struct {
	ArmTmTarget *float64 `json:"armTmTarget,omitempty"`
	Circular *bool `json:"circular,omitempty"`
	Enzyme *string `json:"enzyme,omitempty"`
	Enzyme3 *string `json:"enzyme3,omitempty"`
	Enzyme5 *string `json:"enzyme5,omitempty"`
	Fragments *[]any `json:"fragments,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Insert *string `json:"insert,omitempty"`
	Method string `json:"method"`
	Names *[]any `json:"names,omitempty"`
	Ok any `json:"ok"`
	OverlapLen *int `json:"overlapLen,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Vector *string `json:"vector,omitempty"`
}

// CloningSimulateCreateData is the typed request payload for CloningSimulate.CreateTyped.
type CloningSimulateCreateData struct {
	ArmTmTarget *float64 `json:"armTmTarget,omitempty"`
	Circular *bool `json:"circular,omitempty"`
	Enzyme *string `json:"enzyme,omitempty"`
	Enzyme3 *string `json:"enzyme3,omitempty"`
	Enzyme5 *string `json:"enzyme5,omitempty"`
	Fragments *[]any `json:"fragments,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Insert *string `json:"insert,omitempty"`
	Method string `json:"method"`
	Names *[]any `json:"names,omitempty"`
	Ok any `json:"ok"`
	OverlapLen *int `json:"overlapLen,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Vector *string `json:"vector,omitempty"`
}

// CodonAdaptationIndex is the typed data model for the codon_adaptation_index entity.
type CodonAdaptationIndex struct {
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Provenance map[string]any `json:"provenance"`
	RareThreshold *float64 `json:"rareThreshold,omitempty"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CodonAdaptationIndexCreateData is the typed request payload for CodonAdaptationIndex.CreateTyped.
type CodonAdaptationIndexCreateData struct {
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Provenance map[string]any `json:"provenance"`
	RareThreshold *float64 `json:"rareThreshold,omitempty"`
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
	AvoidEnzymes *[]any `json:"avoidEnzymes,omitempty"`
	CrypticOrfMinAa *int `json:"crypticOrfMinAa,omitempty"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcHigh *float64 `json:"gcHigh,omitempty"`
	GcLow *float64 `json:"gcLow,omitempty"`
	GcWindow *int `json:"gcWindow,omitempty"`
	HomopolymerMin *int `json:"homopolymerMin,omitempty"`
	MaxPasses *int `json:"maxPasses,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ConstructAutofixCreateData is the typed request payload for ConstructAutofix.CreateTyped.
type ConstructAutofixCreateData struct {
	AvoidEnzymes *[]any `json:"avoidEnzymes,omitempty"`
	CrypticOrfMinAa *int `json:"crypticOrfMinAa,omitempty"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcHigh *float64 `json:"gcHigh,omitempty"`
	GcLow *float64 `json:"gcLow,omitempty"`
	GcWindow *int `json:"gcWindow,omitempty"`
	HomopolymerMin *int `json:"homopolymerMin,omitempty"`
	MaxPasses *int `json:"maxPasses,omitempty"`
	Ok any `json:"ok"`
	Organism *string `json:"organism,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ConstructQc is the typed data model for the construct_qc entity.
type ConstructQc struct {
	AvoidEnzymes *[]any `json:"avoidEnzymes,omitempty"`
	CrypticOrfMinAa *int `json:"crypticOrfMinAa,omitempty"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcHigh *float64 `json:"gcHigh,omitempty"`
	GcLow *float64 `json:"gcLow,omitempty"`
	GcWindow *int `json:"gcWindow,omitempty"`
	HomopolymerMin *int `json:"homopolymerMin,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ConstructQcCreateData is the typed request payload for ConstructQc.CreateTyped.
type ConstructQcCreateData struct {
	AvoidEnzymes *[]any `json:"avoidEnzymes,omitempty"`
	CrypticOrfMinAa *int `json:"crypticOrfMinAa,omitempty"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcHigh *float64 `json:"gcHigh,omitempty"`
	GcLow *float64 `json:"gcLow,omitempty"`
	GcWindow *int `json:"gcWindow,omitempty"`
	HomopolymerMin *int `json:"homopolymerMin,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CrisprGrnaDesign is the typed data model for the crispr_grna_design entity.
type CrisprGrnaDesign struct {
	Gate *any `json:"gate,omitempty"`
	MinScore *float64 `json:"minScore,omitempty"`
	Nuclease *string `json:"nuclease,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SearchReverseStrand *bool `json:"searchReverseStrand,omitempty"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CrisprGrnaDesignCreateData is the typed request payload for CrisprGrnaDesign.CreateTyped.
type CrisprGrnaDesignCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MinScore *float64 `json:"minScore,omitempty"`
	Nuclease *string `json:"nuclease,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SearchReverseStrand *bool `json:"searchReverseStrand,omitempty"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// CrisprHdrDonor is the typed data model for the crispr_hdr_donor entity.
type CrisprHdrDonor struct {
	ArmLength *int `json:"armLength,omitempty"`
	BlockPam *bool `json:"blockPam,omitempty"`
	DesignGenotypingPrimers *bool `json:"designGenotypingPrimers,omitempty"`
	EditEnd *int `json:"editEnd,omitempty"`
	EditStart *int `json:"editStart,omitempty"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GuideEnd *int `json:"guideEnd,omitempty"`
	GuideStart *int `json:"guideStart,omitempty"`
	GuideStrand *string `json:"guideStrand,omitempty"`
	Nuclease *string `json:"nuclease,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Replacement string `json:"replacement"`
	Result map[string]any `json:"result"`
	TargetSequence string `json:"targetSequence"`
	Tool string `json:"tool"`
}

// CrisprHdrDonorCreateData is the typed request payload for CrisprHdrDonor.CreateTyped.
type CrisprHdrDonorCreateData struct {
	ArmLength *int `json:"armLength,omitempty"`
	BlockPam *bool `json:"blockPam,omitempty"`
	DesignGenotypingPrimers *bool `json:"designGenotypingPrimers,omitempty"`
	EditEnd *int `json:"editEnd,omitempty"`
	EditStart *int `json:"editStart,omitempty"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GuideEnd *int `json:"guideEnd,omitempty"`
	GuideStart *int `json:"guideStart,omitempty"`
	GuideStrand *string `json:"guideStrand,omitempty"`
	Nuclease *string `json:"nuclease,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Replacement string `json:"replacement"`
	Result map[string]any `json:"result"`
	TargetSequence string `json:"targetSequence"`
	Tool string `json:"tool"`
}

// CrisprOfftargetCheck is the typed data model for the crispr_offtarget_check entity.
type CrisprOfftargetCheck struct {
	Gate *any `json:"gate,omitempty"`
	MaxMismatches *int `json:"maxMismatches,omitempty"`
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
	MaxMismatches *int `json:"maxMismatches,omitempty"`
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
	SequenceA string `json:"sequenceA"`
	SequenceB string `json:"sequenceB"`
	Tool string `json:"tool"`
}

// CrossDimerCreateData is the typed request payload for CrossDimer.CreateTyped.
type CrossDimerCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SequenceA string `json:"sequenceA"`
	SequenceB string `json:"sequenceB"`
	Tool string `json:"tool"`
}

// DnaMolarity is the typed data model for the dna_molarity entity.
type DnaMolarity struct {
	Gate *any `json:"gate,omitempty"`
	Length *int `json:"length,omitempty"`
	MassNg *float64 `json:"massNg,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence *string `json:"sequence,omitempty"`
	Tool string `json:"tool"`
	Type *string `json:"type,omitempty"`
	VolumeUl *float64 `json:"volumeUl,omitempty"`
}

// DnaMolarityCreateData is the typed request payload for DnaMolarity.CreateTyped.
type DnaMolarityCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Length *int `json:"length,omitempty"`
	MassNg *float64 `json:"massNg,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence *string `json:"sequence,omitempty"`
	Tool string `json:"tool"`
	Type *string `json:"type,omitempty"`
	VolumeUl *float64 `json:"volumeUl,omitempty"`
}

// DoubleDigest is the typed data model for the double_digest entity.
type DoubleDigest struct {
	EnzymeA string `json:"enzymeA"`
	EnzymeB string `json:"enzymeB"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// DoubleDigestCreateData is the typed request payload for DoubleDigest.CreateTyped.
type DoubleDigestCreateData struct {
	EnzymeA string `json:"enzymeA"`
	EnzymeB string `json:"enzymeB"`
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
	Reactions []any `json:"reactions"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExportEchoPicklistCreateData is the typed request payload for ExportEchoPicklist.CreateTyped.
type ExportEchoPicklistCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Reactions []any `json:"reactions"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExportOpentronsProtocol is the typed data model for the export_opentrons_protocol entity.
type ExportOpentronsProtocol struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	ProtocolName *string `json:"protocolName,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Reactions []any `json:"reactions"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExportOpentronsProtocolCreateData is the typed request payload for ExportOpentronsProtocol.CreateTyped.
type ExportOpentronsProtocolCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	ProtocolName *string `json:"protocolName,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Reactions []any `json:"reactions"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExportPlateLayout is the typed data model for the export_plate_layout entity.
type ExportPlateLayout struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Reactions []any `json:"reactions"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExportPlateLayoutCreateData is the typed request payload for ExportPlateLayout.CreateTyped.
type ExportPlateLayoutCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Reactions []any `json:"reactions"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ExpressionHeatmapCluster is the typed data model for the expression_heatmap_cluster entity.
type ExpressionHeatmapCluster struct {
	ClusterCols *bool `json:"clusterCols,omitempty"`
	ClusterRows *bool `json:"clusterRows,omitempty"`
	DistanceMetric *string `json:"distanceMetric,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Genes []any `json:"genes"`
	Linkage *string `json:"linkage,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Samples []any `json:"samples"`
	Tool string `json:"tool"`
	Values []any `json:"values"`
	ZScoreRows *bool `json:"zScoreRows,omitempty"`
}

// ExpressionHeatmapClusterCreateData is the typed request payload for ExpressionHeatmapCluster.CreateTyped.
type ExpressionHeatmapClusterCreateData struct {
	ClusterCols *bool `json:"clusterCols,omitempty"`
	ClusterRows *bool `json:"clusterRows,omitempty"`
	DistanceMetric *string `json:"distanceMetric,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Genes []any `json:"genes"`
	Linkage *string `json:"linkage,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Samples []any `json:"samples"`
	Tool string `json:"tool"`
	Values []any `json:"values"`
	ZScoreRows *bool `json:"zScoreRows,omitempty"`
}

// FastqQcReport is the typed data model for the fastq_qc_report entity.
type FastqQcReport struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	QualityOffset *int `json:"qualityOffset,omitempty"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// FastqQcReportCreateData is the typed request payload for FastqQcReport.CreateTyped.
type FastqQcReportCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	QualityOffset *int `json:"qualityOffset,omitempty"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// FastqTrim is the typed data model for the fastq_trim entity.
type FastqTrim struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	MinLength *int `json:"minLength,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	QualityOffset *int `json:"qualityOffset,omitempty"`
	QualityThreshold *int `json:"qualityThreshold,omitempty"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// FastqTrimCreateData is the typed request payload for FastqTrim.CreateTyped.
type FastqTrimCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	MinLength *int `json:"minLength,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	QualityOffset *int `json:"qualityOffset,omitempty"`
	QualityThreshold *int `json:"qualityThreshold,omitempty"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// FindOrf is the typed data model for the find_orf entity.
type FindOrf struct {
	Gate *any `json:"gate,omitempty"`
	MinAaLength *int `json:"minAaLength,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	RequireStop *bool `json:"requireStop,omitempty"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// FindOrfCreateData is the typed request payload for FindOrf.CreateTyped.
type FindOrfCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MinAaLength *int `json:"minAaLength,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	RequireStop *bool `json:"requireStop,omitempty"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// FormatSequence is the typed data model for the format_sequence entity.
type FormatSequence struct {
	CaseMode *string `json:"caseMode,omitempty"`
	Convert *string `json:"convert,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Reverse *bool `json:"reverse,omitempty"`
	Sequence string `json:"sequence"`
	StripNonLetters *bool `json:"stripNonLetters,omitempty"`
	Tool string `json:"tool"`
	Width *int `json:"width,omitempty"`
}

// FormatSequenceCreateData is the typed request payload for FormatSequence.CreateTyped.
type FormatSequenceCreateData struct {
	CaseMode *string `json:"caseMode,omitempty"`
	Convert *string `json:"convert,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Reverse *bool `json:"reverse,omitempty"`
	Sequence string `json:"sequence"`
	StripNonLetters *bool `json:"stripNonLetters,omitempty"`
	Tool string `json:"tool"`
	Width *int `json:"width,omitempty"`
}

// FunctionalEnrichment is the typed data model for the functional_enrichment entity.
type FunctionalEnrichment struct {
	Background *[]any `json:"background,omitempty"`
	Collections *[]any `json:"collections,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Genes []any `json:"genes"`
	MaxTermSize *int `json:"maxTermSize,omitempty"`
	MinTermSize *int `json:"minTermSize,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// FunctionalEnrichmentCreateData is the typed request payload for FunctionalEnrichment.CreateTyped.
type FunctionalEnrichmentCreateData struct {
	Background *[]any `json:"background,omitempty"`
	Collections *[]any `json:"collections,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Genes []any `json:"genes"`
	MaxTermSize *int `json:"maxTermSize,omitempty"`
	MinTermSize *int `json:"minTermSize,omitempty"`
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
	CompareToNamedSet *string `json:"compareToNamedSet,omitempty"`
	Dataset *string `json:"dataset,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Overhangs []any `json:"overhangs"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	RiskThreshold *float64 `json:"riskThreshold,omitempty"`
	Tool string `json:"tool"`
}

// GoldenGateFidelityCreateData is the typed request payload for GoldenGateFidelity.CreateTyped.
type GoldenGateFidelityCreateData struct {
	CompareToNamedSet *string `json:"compareToNamedSet,omitempty"`
	Dataset *string `json:"dataset,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Overhangs []any `json:"overhangs"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	RiskThreshold *float64 `json:"riskThreshold,omitempty"`
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
	JobId string `json:"jobId"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// IdMapPollCreateData is the typed request payload for IdMapPoll.CreateTyped.
type IdMapPollCreateData struct {
	Gate *any `json:"gate,omitempty"`
	JobId string `json:"jobId"`
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
	TaxId *string `json:"taxId,omitempty"`
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
	TaxId *string `json:"taxId,omitempty"`
	To string `json:"to"`
	Tool string `json:"tool"`
}

// InSilicoPcr is the typed data model for the in_silico_pcr entity.
type InSilicoPcr struct {
	Circular *bool `json:"circular,omitempty"`
	ForwardPrimer string `json:"forwardPrimer"`
	Gate *any `json:"gate,omitempty"`
	MaxMismatches *int `json:"maxMismatches,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ReversePrimer string `json:"reversePrimer"`
	Template string `json:"template"`
	Tool string `json:"tool"`
}

// InSilicoPcrCreateData is the typed request payload for InSilicoPcr.CreateTyped.
type InSilicoPcrCreateData struct {
	Circular *bool `json:"circular,omitempty"`
	ForwardPrimer string `json:"forwardPrimer"`
	Gate *any `json:"gate,omitempty"`
	MaxMismatches *int `json:"maxMismatches,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ReversePrimer string `json:"reversePrimer"`
	Template string `json:"template"`
	Tool string `json:"tool"`
}

// KaspPrimerDesign is the typed data model for the kasp_primer_design entity.
type KaspPrimerDesign struct {
	AddSecondaryMismatch *bool `json:"addSecondaryMismatch,omitempty"`
	AlleleA string `json:"alleleA"`
	AlleleB string `json:"alleleB"`
	Gate *any `json:"gate,omitempty"`
	MaxAmplicon *int `json:"maxAmplicon,omitempty"`
	MinAmplicon *int `json:"minAmplicon,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SnpPosition int `json:"snpPosition"`
	Target string `json:"target"`
	TargetCoreTm *float64 `json:"targetCoreTm,omitempty"`
	Tool string `json:"tool"`
}

// KaspPrimerDesignCreateData is the typed request payload for KaspPrimerDesign.CreateTyped.
type KaspPrimerDesignCreateData struct {
	AddSecondaryMismatch *bool `json:"addSecondaryMismatch,omitempty"`
	AlleleA string `json:"alleleA"`
	AlleleB string `json:"alleleB"`
	Gate *any `json:"gate,omitempty"`
	MaxAmplicon *int `json:"maxAmplicon,omitempty"`
	MinAmplicon *int `json:"minAmplicon,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SnpPosition int `json:"snpPosition"`
	Target string `json:"target"`
	TargetCoreTm *float64 `json:"targetCoreTm,omitempty"`
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
	DntpMM *float64 `json:"dntpMM,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMM *float64 `json:"mgMM,omitempty"`
	NaMM *float64 `json:"naMM,omitempty"`
	Ok any `json:"ok"`
	OligoNM *float64 `json:"oligoNM,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	TargetTm *float64 `json:"targetTm,omitempty"`
	TmTolerance *float64 `json:"tmTolerance,omitempty"`
	Tool string `json:"tool"`
}

// MeltingTemperatureCreateData is the typed request payload for MeltingTemperature.CreateTyped.
type MeltingTemperatureCreateData struct {
	DntpMM *float64 `json:"dntpMM,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMM *float64 `json:"mgMM,omitempty"`
	NaMM *float64 `json:"naMM,omitempty"`
	Ok any `json:"ok"`
	OligoNM *float64 `json:"oligoNM,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	TargetTm *float64 `json:"targetTm,omitempty"`
	TmTolerance *float64 `json:"tmTolerance,omitempty"`
	Tool string `json:"tool"`
}

// MotifFinder is the typed data model for the motif_finder entity.
type MotifFinder struct {
	Gate *any `json:"gate,omitempty"`
	MaxMismatches *int `json:"maxMismatches,omitempty"`
	Motif string `json:"motif"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SearchReverseStrand *bool `json:"searchReverseStrand,omitempty"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// MotifFinderCreateData is the typed request payload for MotifFinder.CreateTyped.
type MotifFinderCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MaxMismatches *int `json:"maxMismatches,omitempty"`
	Motif string `json:"motif"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SearchReverseStrand *bool `json:"searchReverseStrand,omitempty"`
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
	DntpMM *float64 `json:"dntpMM,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMM *float64 `json:"mgMM,omitempty"`
	NaMM *float64 `json:"naMM,omitempty"`
	Ok any `json:"ok"`
	OligoNM *float64 `json:"oligoNM,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// OligoAnalysiCreateData is the typed request payload for OligoAnalysi.CreateTyped.
type OligoAnalysiCreateData struct {
	DntpMM *float64 `json:"dntpMM,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMM *float64 `json:"mgMM,omitempty"`
	NaMM *float64 `json:"naMM,omitempty"`
	Ok any `json:"ok"`
	OligoNM *float64 `json:"oligoNM,omitempty"`
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
	SourceSpecies *string `json:"sourceSpecies,omitempty"`
	Symbols []any `json:"symbols"`
	TargetSpecies string `json:"targetSpecies"`
	Tool string `json:"tool"`
	Type *string `json:"type,omitempty"`
}

// OrthologMapCreateData is the typed request payload for OrthologMap.CreateTyped.
type OrthologMapCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SourceSpecies *string `json:"sourceSpecies,omitempty"`
	Symbols []any `json:"symbols"`
	TargetSpecies string `json:"targetSpecies"`
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
	SeqA string `json:"seqA"`
	SeqB string `json:"seqB"`
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
	SeqA string `json:"seqA"`
	SeqB string `json:"seqB"`
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
	FileBase64 string `json:"fileBase64"`
	FileName *string `json:"fileName,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ParseSangerTraceCreateData is the typed request payload for ParseSangerTrace.CreateTyped.
type ParseSangerTraceCreateData struct {
	FileBase64 string `json:"fileBase64"`
	FileName *string `json:"fileName,omitempty"`
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
	TopN *int `json:"topN,omitempty"`
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
	TopN *int `json:"topN,omitempty"`
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
	TopN *int `json:"topN,omitempty"`
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
	TopN *int `json:"topN,omitempty"`
}

// PrimeEditingDesign is the typed data model for the prime_editing_design entity.
type PrimeEditingDesign struct {
	EditEnd int `json:"editEnd"`
	EditStart int `json:"editStart"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	InsertedSeq *string `json:"insertedSeq,omitempty"`
	Ok any `json:"ok"`
	PbsLength *int `json:"pbsLength,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	RttHomology *int `json:"rttHomology,omitempty"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// PrimeEditingDesignCreateData is the typed request payload for PrimeEditingDesign.CreateTyped.
type PrimeEditingDesignCreateData struct {
	EditEnd int `json:"editEnd"`
	EditStart int `json:"editStart"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	InsertedSeq *string `json:"insertedSeq,omitempty"`
	Ok any `json:"ok"`
	PbsLength *int `json:"pbsLength,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	RttHomology *int `json:"rttHomology,omitempty"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// PrimeEditingTwinDesign is the typed data model for the prime_editing_twin_design entity.
type PrimeEditingTwinDesign struct {
	Gate *any `json:"gate,omitempty"`
	NewSequence string `json:"newSequence"`
	Ok any `json:"ok"`
	OverlapLength *int `json:"overlapLength,omitempty"`
	PbsLength *int `json:"pbsLength,omitempty"`
	Provenance map[string]any `json:"provenance"`
	ReplaceEnd int `json:"replaceEnd"`
	ReplaceStart int `json:"replaceStart"`
	Result map[string]any `json:"result"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// PrimeEditingTwinDesignCreateData is the typed request payload for PrimeEditingTwinDesign.CreateTyped.
type PrimeEditingTwinDesignCreateData struct {
	Gate *any `json:"gate,omitempty"`
	NewSequence string `json:"newSequence"`
	Ok any `json:"ok"`
	OverlapLength *int `json:"overlapLength,omitempty"`
	PbsLength *int `json:"pbsLength,omitempty"`
	Provenance map[string]any `json:"provenance"`
	ReplaceEnd int `json:"replaceEnd"`
	ReplaceStart int `json:"replaceStart"`
	Result map[string]any `json:"result"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// PrimerDesign is the typed data model for the primer_design entity.
type PrimerDesign struct {
	AmpliconMax *int `json:"ampliconMax,omitempty"`
	AmpliconMin *int `json:"ampliconMin,omitempty"`
	DntpMM *float64 `json:"dntpMM,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcMax *float64 `json:"gcMax,omitempty"`
	GcMin *float64 `json:"gcMin,omitempty"`
	LenMax *int `json:"lenMax,omitempty"`
	LenMin *int `json:"lenMin,omitempty"`
	LenOpt *int `json:"lenOpt,omitempty"`
	MaxReturn *int `json:"maxReturn,omitempty"`
	MgMM *float64 `json:"mgMM,omitempty"`
	NaMM *float64 `json:"naMM,omitempty"`
	Ok any `json:"ok"`
	OligoNM *float64 `json:"oligoNM,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	TargetEnd *int `json:"targetEnd,omitempty"`
	TargetStart *int `json:"targetStart,omitempty"`
	Template string `json:"template"`
	TmMax *float64 `json:"tmMax,omitempty"`
	TmMaxDiff *float64 `json:"tmMaxDiff,omitempty"`
	TmMin *float64 `json:"tmMin,omitempty"`
	TmOpt *float64 `json:"tmOpt,omitempty"`
	Tool string `json:"tool"`
}

// PrimerDesignCreateData is the typed request payload for PrimerDesign.CreateTyped.
type PrimerDesignCreateData struct {
	AmpliconMax *int `json:"ampliconMax,omitempty"`
	AmpliconMin *int `json:"ampliconMin,omitempty"`
	DntpMM *float64 `json:"dntpMM,omitempty"`
	Gate *any `json:"gate,omitempty"`
	GcMax *float64 `json:"gcMax,omitempty"`
	GcMin *float64 `json:"gcMin,omitempty"`
	LenMax *int `json:"lenMax,omitempty"`
	LenMin *int `json:"lenMin,omitempty"`
	LenOpt *int `json:"lenOpt,omitempty"`
	MaxReturn *int `json:"maxReturn,omitempty"`
	MgMM *float64 `json:"mgMM,omitempty"`
	NaMM *float64 `json:"naMM,omitempty"`
	Ok any `json:"ok"`
	OligoNM *float64 `json:"oligoNM,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	TargetEnd *int `json:"targetEnd,omitempty"`
	TargetStart *int `json:"targetStart,omitempty"`
	Template string `json:"template"`
	TmMax *float64 `json:"tmMax,omitempty"`
	TmMaxDiff *float64 `json:"tmMaxDiff,omitempty"`
	TmMin *float64 `json:"tmMin,omitempty"`
	TmOpt *float64 `json:"tmOpt,omitempty"`
	Tool string `json:"tool"`
}

// PrimerSpecificity is the typed data model for the primer_specificity entity.
type PrimerSpecificity struct {
	ForwardPrimer string `json:"forwardPrimer"`
	Gate *any `json:"gate,omitempty"`
	MaxMismatches *int `json:"maxMismatches,omitempty"`
	MaxProductLength *int `json:"maxProductLength,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ReversePrimer string `json:"reversePrimer"`
	Tool string `json:"tool"`
}

// PrimerSpecificityCreateData is the typed request payload for PrimerSpecificity.CreateTyped.
type PrimerSpecificityCreateData struct {
	ForwardPrimer string `json:"forwardPrimer"`
	Gate *any `json:"gate,omitempty"`
	MaxMismatches *int `json:"maxMismatches,omitempty"`
	MaxProductLength *int `json:"maxProductLength,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ReversePrimer string `json:"reversePrimer"`
	Tool string `json:"tool"`
}

// ProteaseDigestion is the typed data model for the protease_digestion entity.
type ProteaseDigestion struct {
	Gate *any `json:"gate,omitempty"`
	MaxMass *float64 `json:"maxMass,omitempty"`
	MaxPeptides *int `json:"maxPeptides,omitempty"`
	MinMass *float64 `json:"minMass,omitempty"`
	MissedCleavages *int `json:"missedCleavages,omitempty"`
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
	MaxMass *float64 `json:"maxMass,omitempty"`
	MaxPeptides *int `json:"maxPeptides,omitempty"`
	MinMass *float64 `json:"minMass,omitempty"`
	MissedCleavages *int `json:"missedCleavages,omitempty"`
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
	JobId string `json:"jobId"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ProteinAnnotatePollCreateData is the typed request payload for ProteinAnnotatePoll.CreateTyped.
type ProteinAnnotatePollCreateData struct {
	Gate *any `json:"gate,omitempty"`
	JobId string `json:"jobId"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// ProteinAnnotateSubmit is the typed data model for the protein_annotate_submit entity.
type ProteinAnnotateSubmit struct {
	Appl *string `json:"appl,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Goterms *bool `json:"goterms,omitempty"`
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
	Goterms *bool `json:"goterms,omitempty"`
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
	ChargeStep *float64 `json:"chargeStep,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// ProteinPropertyCreateData is the typed request payload for ProteinProperty.CreateTyped.
type ProteinPropertyCreateData struct {
	ChargeStep *float64 `json:"chargeStep,omitempty"`
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
	GcContent *float64 `json:"gcContent,omitempty"`
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
	GcContent *float64 `json:"gcContent,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Length int `json:"length"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// RestrictionSite is the typed data model for the restriction_site entity.
type RestrictionSite struct {
	Enzymes *[]any `json:"enzymes,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// RestrictionSiteCreateData is the typed request payload for RestrictionSite.CreateTyped.
type RestrictionSiteCreateData struct {
	Enzymes *[]any `json:"enzymes,omitempty"`
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
	FileBase64 *string `json:"fileBase64,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MinCoverage *float64 `json:"minCoverage,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Read *string `json:"read,omitempty"`
	Reference string `json:"reference"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SangerVsReferenceCreateData is the typed request payload for SangerVsReference.CreateTyped.
type SangerVsReferenceCreateData struct {
	FileBase64 *string `json:"fileBase64,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MinCoverage *float64 `json:"minCoverage,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Read *string `json:"read,omitempty"`
	Reference string `json:"reference"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SavePermalink is the typed data model for the save_permalink entity.
type SavePermalink struct {
	Args map[string]any `json:"args"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SavePermalinkCreateData is the typed request payload for SavePermalink.CreateTyped.
type SavePermalinkCreateData struct {
	Args map[string]any `json:"args"`
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
	QualityOffset *int `json:"qualityOffset,omitempty"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SeqfileStatCreateData is the typed request payload for SeqfileStat.CreateTyped.
type SeqfileStatCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Input string `json:"input"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	QualityOffset *int `json:"qualityOffset,omitempty"`
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
	EndPrimerLength *int `json:"endPrimerLength,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MaxOrfs *int `json:"maxOrfs,omitempty"`
	MinOrfAa *int `json:"minOrfAa,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Sequence string `json:"sequence"`
	Tool string `json:"tool"`
}

// SequenceReportCreateData is the typed request payload for SequenceReport.CreateTyped.
type SequenceReportCreateData struct {
	EndPrimerLength *int `json:"endPrimerLength,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MaxOrfs *int `json:"maxOrfs,omitempty"`
	MinOrfAa *int `json:"minOrfAa,omitempty"`
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
	MaxResults *int `json:"maxResults,omitempty"`
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
	MaxResults *int `json:"maxResults,omitempty"`
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
	MinSupportingReads *int `json:"minSupportingReads,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Reads string `json:"reads"`
	Reference string `json:"reference"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SequencingReadbackVerifyCreateData is the typed request payload for SequencingReadbackVerify.CreateTyped.
type SequencingReadbackVerifyCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MinSupportingReads *int `json:"minSupportingReads,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Reads string `json:"reads"`
	Reference string `json:"reference"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SessionCreate is the typed data model for the session_create entity.
type SessionCreate struct {
	Entries *map[string]any `json:"entries,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SessionCreateCreateData is the typed request payload for SessionCreate.CreateTyped.
type SessionCreateCreateData struct {
	Entries *map[string]any `json:"entries,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// SessionGet is the typed data model for the session_get entity.
type SessionGet struct {
	Gate *any `json:"gate,omitempty"`
	Names *[]any `json:"names,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"sessionId"`
	Tool string `json:"tool"`
}

// SessionGetCreateData is the typed request payload for SessionGet.CreateTyped.
type SessionGetCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Names *[]any `json:"names,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"sessionId"`
	Tool string `json:"tool"`
}

// SessionRun is the typed data model for the session_run entity.
type SessionRun struct {
	Args *map[string]any `json:"args,omitempty"`
	FromSession *map[string]any `json:"fromSession,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"sessionId"`
	Tool string `json:"tool"`
	WriteBack *map[string]any `json:"writeBack,omitempty"`
}

// SessionRunCreateData is the typed request payload for SessionRun.CreateTyped.
type SessionRunCreateData struct {
	Args *map[string]any `json:"args,omitempty"`
	FromSession *map[string]any `json:"fromSession,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"sessionId"`
	Tool string `json:"tool"`
	WriteBack *map[string]any `json:"writeBack,omitempty"`
}

// SessionSet is the typed data model for the session_set entity.
type SessionSet struct {
	Entries map[string]any `json:"entries"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"sessionId"`
	Tool string `json:"tool"`
}

// SessionSetCreateData is the typed request payload for SessionSet.CreateTyped.
type SessionSetCreateData struct {
	Entries map[string]any `json:"entries"`
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	SessionId string `json:"sessionId"`
	Tool string `json:"tool"`
}

// SirnaDesign is the typed data model for the sirna_design entity.
type SirnaDesign struct {
	Gate *any `json:"gate,omitempty"`
	MinReynolds *int `json:"minReynolds,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ShRnaLoop *string `json:"shRnaLoop,omitempty"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// SirnaDesignCreateData is the typed request payload for SirnaDesign.CreateTyped.
type SirnaDesignCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MinReynolds *int `json:"minReynolds,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	ShRnaLoop *string `json:"shRnaLoop,omitempty"`
	Target string `json:"target"`
	Tool string `json:"tool"`
}

// SiteDirectedMutagenesi is the typed data model for the site_directed_mutagenesi entity.
type SiteDirectedMutagenesi struct {
	ArmTmTarget *float64 `json:"armTmTarget,omitempty"`
	DntpMM *float64 `json:"dntpMM,omitempty"`
	EditKind *string `json:"editKind,omitempty"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMM *float64 `json:"mgMM,omitempty"`
	NaMM *float64 `json:"naMM,omitempty"`
	NewBase *string `json:"newBase,omitempty"`
	Ok any `json:"ok"`
	OligoNM *float64 `json:"oligoNM,omitempty"`
	Organism *string `json:"organism,omitempty"`
	Position *int `json:"position,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Residue *int `json:"residue,omitempty"`
	Result map[string]any `json:"result"`
	Style *string `json:"style,omitempty"`
	TargetAa *string `json:"targetAa,omitempty"`
	Template string `json:"template"`
	Tool string `json:"tool"`
}

// SiteDirectedMutagenesiCreateData is the typed request payload for SiteDirectedMutagenesi.CreateTyped.
type SiteDirectedMutagenesiCreateData struct {
	ArmTmTarget *float64 `json:"armTmTarget,omitempty"`
	DntpMM *float64 `json:"dntpMM,omitempty"`
	EditKind *string `json:"editKind,omitempty"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	MgMM *float64 `json:"mgMM,omitempty"`
	NaMM *float64 `json:"naMM,omitempty"`
	NewBase *string `json:"newBase,omitempty"`
	Ok any `json:"ok"`
	OligoNM *float64 `json:"oligoNM,omitempty"`
	Organism *string `json:"organism,omitempty"`
	Position *int `json:"position,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Residue *int `json:"residue,omitempty"`
	Result map[string]any `json:"result"`
	Style *string `json:"style,omitempty"`
	TargetAa *string `json:"targetAa,omitempty"`
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
	ToStop *bool `json:"toStop,omitempty"`
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
	ToStop *bool `json:"toStop,omitempty"`
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
	FrameStart *int `json:"frameStart,omitempty"`
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
	FrameStart *int `json:"frameStart,omitempty"`
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
	ArmTmTarget *float64 `json:"armTmTarget,omitempty"`
	Circular *bool `json:"circular,omitempty"`
	ClaimedConstruct string `json:"claimedConstruct"`
	Coding *bool `json:"coding,omitempty"`
	Enzyme *string `json:"enzyme,omitempty"`
	Enzyme3 *string `json:"enzyme3,omitempty"`
	Enzyme5 *string `json:"enzyme5,omitempty"`
	FragmentPcrs *[]any `json:"fragmentPcrs,omitempty"`
	Fragments *[]any `json:"fragments,omitempty"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Insert *string `json:"insert,omitempty"`
	InsertPcr *map[string]any `json:"insertPcr,omitempty"`
	Method string `json:"method"`
	Names *[]any `json:"names,omitempty"`
	Ok any `json:"ok"`
	OverlapLen *int `json:"overlapLen,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Vector *string `json:"vector,omitempty"`
	VectorPcr *map[string]any `json:"vectorPcr,omitempty"`
}

// VerifyAssemblyCreateData is the typed request payload for VerifyAssembly.CreateTyped.
type VerifyAssemblyCreateData struct {
	ArmTmTarget *float64 `json:"armTmTarget,omitempty"`
	Circular *bool `json:"circular,omitempty"`
	ClaimedConstruct string `json:"claimedConstruct"`
	Coding *bool `json:"coding,omitempty"`
	Enzyme *string `json:"enzyme,omitempty"`
	Enzyme3 *string `json:"enzyme3,omitempty"`
	Enzyme5 *string `json:"enzyme5,omitempty"`
	FragmentPcrs *[]any `json:"fragmentPcrs,omitempty"`
	Fragments *[]any `json:"fragments,omitempty"`
	FrameStart *int `json:"frameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	Insert *string `json:"insert,omitempty"`
	InsertPcr *map[string]any `json:"insertPcr,omitempty"`
	Method string `json:"method"`
	Names *[]any `json:"names,omitempty"`
	Ok any `json:"ok"`
	OverlapLen *int `json:"overlapLen,omitempty"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
	Vector *string `json:"vector,omitempty"`
	VectorPcr *map[string]any `json:"vectorPcr,omitempty"`
}

// VerifyConstruct is the typed data model for the verify_construct entity.
type VerifyConstruct struct {
	ClaimedConstruct string `json:"claimedConstruct"`
	ExpectedFrameStart *int `json:"expectedFrameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	InsertForwardPrimer string `json:"insertForwardPrimer"`
	InsertReversePrimer string `json:"insertReversePrimer"`
	InsertTemplate string `json:"insertTemplate"`
	MaxPrimerMismatches *int `json:"maxPrimerMismatches,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	TemplateCircular *bool `json:"templateCircular,omitempty"`
	Tool string `json:"tool"`
}

// VerifyConstructCreateData is the typed request payload for VerifyConstruct.CreateTyped.
type VerifyConstructCreateData struct {
	ClaimedConstruct string `json:"claimedConstruct"`
	ExpectedFrameStart *int `json:"expectedFrameStart,omitempty"`
	Gate *any `json:"gate,omitempty"`
	InsertForwardPrimer string `json:"insertForwardPrimer"`
	InsertReversePrimer string `json:"insertReversePrimer"`
	InsertTemplate string `json:"insertTemplate"`
	MaxPrimerMismatches *int `json:"maxPrimerMismatches,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	TemplateCircular *bool `json:"templateCircular,omitempty"`
	Tool string `json:"tool"`
}

// VirtualGel is the typed data model for the virtual_gel entity.
type VirtualGel struct {
	Circular *bool `json:"circular,omitempty"`
	Enzymes *[]any `json:"enzymes,omitempty"`
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
	Enzymes *[]any `json:"enzymes,omitempty"`
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
	Rows []any `json:"rows"`
	Tool string `json:"tool"`
}

// VolcanoPlotDataCreateData is the typed request payload for VolcanoPlotData.CreateTyped.
type VolcanoPlotDataCreateData struct {
	Gate *any `json:"gate,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Result map[string]any `json:"result"`
	Rows []any `json:"rows"`
	Tool string `json:"tool"`
}

// WebSearch is the typed data model for the web_search entity.
type WebSearch struct {
	Gate *any `json:"gate,omitempty"`
	MaxResults *float64 `json:"max_results,omitempty"`
	Ok any `json:"ok"`
	Provenance map[string]any `json:"provenance"`
	Query string `json:"query"`
	Result map[string]any `json:"result"`
	Tool string `json:"tool"`
}

// WebSearchCreateData is the typed request payload for WebSearch.CreateTyped.
type WebSearchCreateData struct {
	Gate *any `json:"gate,omitempty"`
	MaxResults *float64 `json:"max_results,omitempty"`
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

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
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

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
