# frozen_string_literal: true

# Typed models for the SeqbenchMcp SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# AlphafoldLookup entity data model.
#
# @!attribute [rw] accession
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
AlphafoldLookup = Struct.new(
  :accession,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for AlphafoldLookup#create.
#
# @!attribute [rw] accession
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
AlphafoldLookupCreateData = Struct.new(
  :accession,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# AsoDesign entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] length
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] wing
#   @return [Integer, nil]
AsoDesign = Struct.new(
  :gate,
  :length,
  :ok,
  :provenance,
  :result,
  :target,
  :tool,
  :wing,
  keyword_init: true
)

# Request payload for AsoDesign#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] length
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] wing
#   @return [Integer, nil]
AsoDesignCreateData = Struct.new(
  :gate,
  :length,
  :ok,
  :provenance,
  :result,
  :target,
  :tool,
  :wing,
  keyword_init: true
)

# BaseEditingDesign entity data model.
#
# @!attribute [rw] editor
#   @return [String, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] target_position
#   @return [Integer, nil]
#
# @!attribute [rw] tool
#   @return [String]
BaseEditingDesign = Struct.new(
  :editor,
  :frame_start,
  :gate,
  :ok,
  :provenance,
  :result,
  :target,
  :target_position,
  :tool,
  keyword_init: true
)

# Request payload for BaseEditingDesign#create.
#
# @!attribute [rw] editor
#   @return [String, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] target_position
#   @return [Integer, nil]
#
# @!attribute [rw] tool
#   @return [String]
BaseEditingDesignCreateData = Struct.new(
  :editor,
  :frame_start,
  :gate,
  :ok,
  :provenance,
  :result,
  :target,
  :target_position,
  :tool,
  keyword_init: true
)

# Batch entity data model.
#
# @!attribute [rw] arg
#   @return [Hash, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
Batch = Struct.new(
  :arg,
  :input,
  :ok,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for Batch#load.
#
# @!attribute [rw] arg
#   @return [Hash, nil]
#
# @!attribute [rw] input
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] tool
#   @return [String, nil]
BatchLoadMatch = Struct.new(
  :arg,
  :input,
  :ok,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for Batch#create.
#
# @!attribute [rw] arg
#   @return [Hash, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
BatchCreateData = Struct.new(
  :arg,
  :input,
  :ok,
  :result,
  :tool,
  keyword_init: true
)

# BatchWorkflow entity data model.
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] step
#   @return [Array]
BatchWorkflow = Struct.new(
  :input,
  :ok,
  :result,
  :step,
  keyword_init: true
)

# Request payload for BatchWorkflow#load.
#
# @!attribute [rw] input
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] step
#   @return [Array, nil]
BatchWorkflowLoadMatch = Struct.new(
  :input,
  :ok,
  :result,
  :step,
  keyword_init: true
)

# Request payload for BatchWorkflow#create.
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] step
#   @return [Array]
BatchWorkflowCreateData = Struct.new(
  :input,
  :ok,
  :result,
  :step,
  keyword_init: true
)

# CharacterizeSequence entity data model.
#
# @!attribute [rw] end_primer_length
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_orf
#   @return [Integer, nil]
#
# @!attribute [rw] min_orf_aa
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CharacterizeSequence = Struct.new(
  :end_primer_length,
  :gate,
  :max_orf,
  :min_orf_aa,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for CharacterizeSequence#create.
#
# @!attribute [rw] end_primer_length
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_orf
#   @return [Integer, nil]
#
# @!attribute [rw] min_orf_aa
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CharacterizeSequenceCreateData = Struct.new(
  :end_primer_length,
  :gate,
  :max_orf,
  :min_orf_aa,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# CloningSimulate entity data model.
#
# @!attribute [rw] arm_tm_target
#   @return [Float, nil]
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] enzyme
#   @return [String, nil]
#
# @!attribute [rw] enzyme3
#   @return [String, nil]
#
# @!attribute [rw] enzyme5
#   @return [String, nil]
#
# @!attribute [rw] fragment
#   @return [Array, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insert
#   @return [String, nil]
#
# @!attribute [rw] method
#   @return [String]
#
# @!attribute [rw] name
#   @return [Array, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlap_len
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] vector
#   @return [String, nil]
CloningSimulate = Struct.new(
  :arm_tm_target,
  :circular,
  :enzyme,
  :enzyme3,
  :enzyme5,
  :fragment,
  :gate,
  :insert,
  :method,
  :name,
  :ok,
  :overlap_len,
  :provenance,
  :result,
  :tool,
  :vector,
  keyword_init: true
)

# Request payload for CloningSimulate#create.
#
# @!attribute [rw] arm_tm_target
#   @return [Float, nil]
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] enzyme
#   @return [String, nil]
#
# @!attribute [rw] enzyme3
#   @return [String, nil]
#
# @!attribute [rw] enzyme5
#   @return [String, nil]
#
# @!attribute [rw] fragment
#   @return [Array, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insert
#   @return [String, nil]
#
# @!attribute [rw] method
#   @return [String]
#
# @!attribute [rw] name
#   @return [Array, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlap_len
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] vector
#   @return [String, nil]
CloningSimulateCreateData = Struct.new(
  :arm_tm_target,
  :circular,
  :enzyme,
  :enzyme3,
  :enzyme5,
  :fragment,
  :gate,
  :insert,
  :method,
  :name,
  :ok,
  :overlap_len,
  :provenance,
  :result,
  :tool,
  :vector,
  keyword_init: true
)

# CodonAdaptationIndex entity data model.
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] rare_threshold
#   @return [Float, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CodonAdaptationIndex = Struct.new(
  :frame_start,
  :gate,
  :ok,
  :organism,
  :provenance,
  :rare_threshold,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for CodonAdaptationIndex#create.
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] rare_threshold
#   @return [Float, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CodonAdaptationIndexCreateData = Struct.new(
  :frame_start,
  :gate,
  :ok,
  :organism,
  :provenance,
  :rare_threshold,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# CodonOptimize entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] protein
#   @return [String]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
CodonOptimize = Struct.new(
  :gate,
  :ok,
  :organism,
  :protein,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for CodonOptimize#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] protein
#   @return [String]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
CodonOptimizeCreateData = Struct.new(
  :gate,
  :ok,
  :organism,
  :protein,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# ConstructAutofix entity data model.
#
# @!attribute [rw] avoid_enzyme
#   @return [Array, nil]
#
# @!attribute [rw] cryptic_orf_min_aa
#   @return [Integer, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gc_high
#   @return [Float, nil]
#
# @!attribute [rw] gc_low
#   @return [Float, nil]
#
# @!attribute [rw] gc_window
#   @return [Integer, nil]
#
# @!attribute [rw] homopolymer_min
#   @return [Integer, nil]
#
# @!attribute [rw] max_pass
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ConstructAutofix = Struct.new(
  :avoid_enzyme,
  :cryptic_orf_min_aa,
  :frame_start,
  :gate,
  :gc_high,
  :gc_low,
  :gc_window,
  :homopolymer_min,
  :max_pass,
  :ok,
  :organism,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for ConstructAutofix#create.
#
# @!attribute [rw] avoid_enzyme
#   @return [Array, nil]
#
# @!attribute [rw] cryptic_orf_min_aa
#   @return [Integer, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gc_high
#   @return [Float, nil]
#
# @!attribute [rw] gc_low
#   @return [Float, nil]
#
# @!attribute [rw] gc_window
#   @return [Integer, nil]
#
# @!attribute [rw] homopolymer_min
#   @return [Integer, nil]
#
# @!attribute [rw] max_pass
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ConstructAutofixCreateData = Struct.new(
  :avoid_enzyme,
  :cryptic_orf_min_aa,
  :frame_start,
  :gate,
  :gc_high,
  :gc_low,
  :gc_window,
  :homopolymer_min,
  :max_pass,
  :ok,
  :organism,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# ConstructQc entity data model.
#
# @!attribute [rw] avoid_enzyme
#   @return [Array, nil]
#
# @!attribute [rw] cryptic_orf_min_aa
#   @return [Integer, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gc_high
#   @return [Float, nil]
#
# @!attribute [rw] gc_low
#   @return [Float, nil]
#
# @!attribute [rw] gc_window
#   @return [Integer, nil]
#
# @!attribute [rw] homopolymer_min
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ConstructQc = Struct.new(
  :avoid_enzyme,
  :cryptic_orf_min_aa,
  :frame_start,
  :gate,
  :gc_high,
  :gc_low,
  :gc_window,
  :homopolymer_min,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for ConstructQc#create.
#
# @!attribute [rw] avoid_enzyme
#   @return [Array, nil]
#
# @!attribute [rw] cryptic_orf_min_aa
#   @return [Integer, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gc_high
#   @return [Float, nil]
#
# @!attribute [rw] gc_low
#   @return [Float, nil]
#
# @!attribute [rw] gc_window
#   @return [Integer, nil]
#
# @!attribute [rw] homopolymer_min
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ConstructQcCreateData = Struct.new(
  :avoid_enzyme,
  :cryptic_orf_min_aa,
  :frame_start,
  :gate,
  :gc_high,
  :gc_low,
  :gc_window,
  :homopolymer_min,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# CrisprGrnaDesign entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] min_score
#   @return [Float, nil]
#
# @!attribute [rw] nuclease
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] search_reverse_strand
#   @return [Boolean, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrisprGrnaDesign = Struct.new(
  :gate,
  :min_score,
  :nuclease,
  :ok,
  :provenance,
  :result,
  :search_reverse_strand,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for CrisprGrnaDesign#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] min_score
#   @return [Float, nil]
#
# @!attribute [rw] nuclease
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] search_reverse_strand
#   @return [Boolean, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrisprGrnaDesignCreateData = Struct.new(
  :gate,
  :min_score,
  :nuclease,
  :ok,
  :provenance,
  :result,
  :search_reverse_strand,
  :sequence,
  :tool,
  keyword_init: true
)

# CrisprHdrDonor entity data model.
#
# @!attribute [rw] arm_length
#   @return [Integer, nil]
#
# @!attribute [rw] block_pam
#   @return [Boolean, nil]
#
# @!attribute [rw] design_genotyping_primer
#   @return [Boolean, nil]
#
# @!attribute [rw] edit_end
#   @return [Integer, nil]
#
# @!attribute [rw] edit_start
#   @return [Integer, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] guide_end
#   @return [Integer, nil]
#
# @!attribute [rw] guide_start
#   @return [Integer, nil]
#
# @!attribute [rw] guide_strand
#   @return [String, nil]
#
# @!attribute [rw] nuclease
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] replacement
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] target_sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrisprHdrDonor = Struct.new(
  :arm_length,
  :block_pam,
  :design_genotyping_primer,
  :edit_end,
  :edit_start,
  :frame_start,
  :gate,
  :guide_end,
  :guide_start,
  :guide_strand,
  :nuclease,
  :ok,
  :provenance,
  :replacement,
  :result,
  :target_sequence,
  :tool,
  keyword_init: true
)

# Request payload for CrisprHdrDonor#create.
#
# @!attribute [rw] arm_length
#   @return [Integer, nil]
#
# @!attribute [rw] block_pam
#   @return [Boolean, nil]
#
# @!attribute [rw] design_genotyping_primer
#   @return [Boolean, nil]
#
# @!attribute [rw] edit_end
#   @return [Integer, nil]
#
# @!attribute [rw] edit_start
#   @return [Integer, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] guide_end
#   @return [Integer, nil]
#
# @!attribute [rw] guide_start
#   @return [Integer, nil]
#
# @!attribute [rw] guide_strand
#   @return [String, nil]
#
# @!attribute [rw] nuclease
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] replacement
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] target_sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrisprHdrDonorCreateData = Struct.new(
  :arm_length,
  :block_pam,
  :design_genotyping_primer,
  :edit_end,
  :edit_start,
  :frame_start,
  :gate,
  :guide_end,
  :guide_start,
  :guide_strand,
  :nuclease,
  :ok,
  :provenance,
  :replacement,
  :result,
  :target_sequence,
  :tool,
  keyword_init: true
)

# CrisprOfftargetCheck entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_mismatch
#   @return [Integer, nil]
#
# @!attribute [rw] nuclease
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] protospacer
#   @return [String]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
CrisprOfftargetCheck = Struct.new(
  :gate,
  :max_mismatch,
  :nuclease,
  :ok,
  :protospacer,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for CrisprOfftargetCheck#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_mismatch
#   @return [Integer, nil]
#
# @!attribute [rw] nuclease
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] protospacer
#   @return [String]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
CrisprOfftargetCheckCreateData = Struct.new(
  :gate,
  :max_mismatch,
  :nuclease,
  :ok,
  :protospacer,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# CrossDimer entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence_a
#   @return [String]
#
# @!attribute [rw] sequence_b
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrossDimer = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence_a,
  :sequence_b,
  :tool,
  keyword_init: true
)

# Request payload for CrossDimer#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence_a
#   @return [String]
#
# @!attribute [rw] sequence_b
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrossDimerCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence_a,
  :sequence_b,
  :tool,
  keyword_init: true
)

# DnaMolarity entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] length
#   @return [Integer, nil]
#
# @!attribute [rw] mass_ng
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String, nil]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] volume_ul
#   @return [Float, nil]
DnaMolarity = Struct.new(
  :gate,
  :length,
  :mass_ng,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :type,
  :volume_ul,
  keyword_init: true
)

# Request payload for DnaMolarity#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] length
#   @return [Integer, nil]
#
# @!attribute [rw] mass_ng
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String, nil]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] volume_ul
#   @return [Float, nil]
DnaMolarityCreateData = Struct.new(
  :gate,
  :length,
  :mass_ng,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :type,
  :volume_ul,
  keyword_init: true
)

# DoubleDigest entity data model.
#
# @!attribute [rw] enzyme_a
#   @return [String]
#
# @!attribute [rw] enzyme_b
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
DoubleDigest = Struct.new(
  :enzyme_a,
  :enzyme_b,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for DoubleDigest#create.
#
# @!attribute [rw] enzyme_a
#   @return [String]
#
# @!attribute [rw] enzyme_b
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
DoubleDigestCreateData = Struct.new(
  :enzyme_a,
  :enzyme_b,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# ExportEchoPicklist entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] reaction
#   @return [Array]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ExportEchoPicklist = Struct.new(
  :gate,
  :ok,
  :provenance,
  :reaction,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for ExportEchoPicklist#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] reaction
#   @return [Array]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ExportEchoPicklistCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :reaction,
  :result,
  :tool,
  keyword_init: true
)

# ExportOpentronsProtocol entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] protocol_name
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] reaction
#   @return [Array]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ExportOpentronsProtocol = Struct.new(
  :gate,
  :ok,
  :protocol_name,
  :provenance,
  :reaction,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for ExportOpentronsProtocol#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] protocol_name
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] reaction
#   @return [Array]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ExportOpentronsProtocolCreateData = Struct.new(
  :gate,
  :ok,
  :protocol_name,
  :provenance,
  :reaction,
  :result,
  :tool,
  keyword_init: true
)

# ExportPlateLayout entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] reaction
#   @return [Array]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ExportPlateLayout = Struct.new(
  :gate,
  :ok,
  :provenance,
  :reaction,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for ExportPlateLayout#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] reaction
#   @return [Array]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ExportPlateLayoutCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :reaction,
  :result,
  :tool,
  keyword_init: true
)

# ExpressionHeatmapCluster entity data model.
#
# @!attribute [rw] cluster_col
#   @return [Boolean, nil]
#
# @!attribute [rw] cluster_row
#   @return [Boolean, nil]
#
# @!attribute [rw] distance_metric
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [Array]
#
# @!attribute [rw] linkage
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sample
#   @return [Array]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] value
#   @return [Array]
#
# @!attribute [rw] z_score_row
#   @return [Boolean, nil]
ExpressionHeatmapCluster = Struct.new(
  :cluster_col,
  :cluster_row,
  :distance_metric,
  :gate,
  :gene,
  :linkage,
  :ok,
  :provenance,
  :result,
  :sample,
  :tool,
  :value,
  :z_score_row,
  keyword_init: true
)

# Request payload for ExpressionHeatmapCluster#create.
#
# @!attribute [rw] cluster_col
#   @return [Boolean, nil]
#
# @!attribute [rw] cluster_row
#   @return [Boolean, nil]
#
# @!attribute [rw] distance_metric
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [Array]
#
# @!attribute [rw] linkage
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sample
#   @return [Array]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] value
#   @return [Array]
#
# @!attribute [rw] z_score_row
#   @return [Boolean, nil]
ExpressionHeatmapClusterCreateData = Struct.new(
  :cluster_col,
  :cluster_row,
  :distance_metric,
  :gate,
  :gene,
  :linkage,
  :ok,
  :provenance,
  :result,
  :sample,
  :tool,
  :value,
  :z_score_row,
  keyword_init: true
)

# FastqQcReport entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] quality_offset
#   @return [Integer, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
FastqQcReport = Struct.new(
  :gate,
  :input,
  :ok,
  :provenance,
  :quality_offset,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for FastqQcReport#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] quality_offset
#   @return [Integer, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
FastqQcReportCreateData = Struct.new(
  :gate,
  :input,
  :ok,
  :provenance,
  :quality_offset,
  :result,
  :tool,
  keyword_init: true
)

# FastqTrim entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] min_length
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] quality_offset
#   @return [Integer, nil]
#
# @!attribute [rw] quality_threshold
#   @return [Integer, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
FastqTrim = Struct.new(
  :gate,
  :input,
  :min_length,
  :ok,
  :provenance,
  :quality_offset,
  :quality_threshold,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for FastqTrim#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] min_length
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] quality_offset
#   @return [Integer, nil]
#
# @!attribute [rw] quality_threshold
#   @return [Integer, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
FastqTrimCreateData = Struct.new(
  :gate,
  :input,
  :min_length,
  :ok,
  :provenance,
  :quality_offset,
  :quality_threshold,
  :result,
  :tool,
  keyword_init: true
)

# FindOrf entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] min_aa_length
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] require_stop
#   @return [Boolean, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
FindOrf = Struct.new(
  :gate,
  :min_aa_length,
  :ok,
  :provenance,
  :require_stop,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for FindOrf#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] min_aa_length
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] require_stop
#   @return [Boolean, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
FindOrfCreateData = Struct.new(
  :gate,
  :min_aa_length,
  :ok,
  :provenance,
  :require_stop,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# FormatSequence entity data model.
#
# @!attribute [rw] case_mode
#   @return [String, nil]
#
# @!attribute [rw] convert
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] reverse
#   @return [Boolean, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] strip_non_letter
#   @return [Boolean, nil]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] width
#   @return [Integer, nil]
FormatSequence = Struct.new(
  :case_mode,
  :convert,
  :gate,
  :ok,
  :provenance,
  :result,
  :reverse,
  :sequence,
  :strip_non_letter,
  :tool,
  :width,
  keyword_init: true
)

# Request payload for FormatSequence#create.
#
# @!attribute [rw] case_mode
#   @return [String, nil]
#
# @!attribute [rw] convert
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] reverse
#   @return [Boolean, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] strip_non_letter
#   @return [Boolean, nil]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] width
#   @return [Integer, nil]
FormatSequenceCreateData = Struct.new(
  :case_mode,
  :convert,
  :gate,
  :ok,
  :provenance,
  :result,
  :reverse,
  :sequence,
  :strip_non_letter,
  :tool,
  :width,
  keyword_init: true
)

# FunctionalEnrichment entity data model.
#
# @!attribute [rw] background
#   @return [Array, nil]
#
# @!attribute [rw] collection
#   @return [Array, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [Array]
#
# @!attribute [rw] max_term_size
#   @return [Integer, nil]
#
# @!attribute [rw] min_term_size
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
FunctionalEnrichment = Struct.new(
  :background,
  :collection,
  :gate,
  :gene,
  :max_term_size,
  :min_term_size,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for FunctionalEnrichment#create.
#
# @!attribute [rw] background
#   @return [Array, nil]
#
# @!attribute [rw] collection
#   @return [Array, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [Array]
#
# @!attribute [rw] max_term_size
#   @return [Integer, nil]
#
# @!attribute [rw] min_term_size
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
FunctionalEnrichmentCreateData = Struct.new(
  :background,
  :collection,
  :gate,
  :gene,
  :max_term_size,
  :min_term_size,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# GcContent entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
GcContent = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for GcContent#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
GcContentCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# GeneDossier entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
GeneDossier = Struct.new(
  :gate,
  :gene,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for GeneDossier#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
GeneDossierCreateData = Struct.new(
  :gate,
  :gene,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# GeneExpression entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
GeneExpression = Struct.new(
  :gate,
  :gene,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for GeneExpression#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
GeneExpressionCreateData = Struct.new(
  :gate,
  :gene,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# GeneModel entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
GeneModel = Struct.new(
  :gate,
  :gene,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for GeneModel#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
GeneModelCreateData = Struct.new(
  :gate,
  :gene,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# GoldenGateFidelity entity data model.
#
# @!attribute [rw] compare_to_named_set
#   @return [String, nil]
#
# @!attribute [rw] dataset
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overhang
#   @return [Array]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] risk_threshold
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
GoldenGateFidelity = Struct.new(
  :compare_to_named_set,
  :dataset,
  :gate,
  :ok,
  :overhang,
  :provenance,
  :result,
  :risk_threshold,
  :tool,
  keyword_init: true
)

# Request payload for GoldenGateFidelity#create.
#
# @!attribute [rw] compare_to_named_set
#   @return [String, nil]
#
# @!attribute [rw] dataset
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overhang
#   @return [Array]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] risk_threshold
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
GoldenGateFidelityCreateData = Struct.new(
  :compare_to_named_set,
  :dataset,
  :gate,
  :ok,
  :overhang,
  :provenance,
  :result,
  :risk_threshold,
  :tool,
  keyword_init: true
)

# HgvsConvert entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] variant
#   @return [String]
HgvsConvert = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  :variant,
  keyword_init: true
)

# Request payload for HgvsConvert#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] variant
#   @return [String]
HgvsConvertCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  :variant,
  keyword_init: true
)

# IdMapPoll entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] job_id
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
IdMapPoll = Struct.new(
  :gate,
  :job_id,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for IdMapPoll#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] job_id
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
IdMapPollCreateData = Struct.new(
  :gate,
  :job_id,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# IdMapSubmit entity data model.
#
# @!attribute [rw] from
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ids
#   @return [Array]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tax_id
#   @return [String, nil]
#
# @!attribute [rw] to
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
IdMapSubmit = Struct.new(
  :from,
  :gate,
  :ids,
  :ok,
  :provenance,
  :result,
  :tax_id,
  :to,
  :tool,
  keyword_init: true
)

# Request payload for IdMapSubmit#create.
#
# @!attribute [rw] from
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ids
#   @return [Array]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tax_id
#   @return [String, nil]
#
# @!attribute [rw] to
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
IdMapSubmitCreateData = Struct.new(
  :from,
  :gate,
  :ids,
  :ok,
  :provenance,
  :result,
  :tax_id,
  :to,
  :tool,
  keyword_init: true
)

# InSilicoPcr entity data model.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] forward_primer
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_mismatch
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] reverse_primer
#   @return [String]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
InSilicoPcr = Struct.new(
  :circular,
  :forward_primer,
  :gate,
  :max_mismatch,
  :ok,
  :provenance,
  :result,
  :reverse_primer,
  :template,
  :tool,
  keyword_init: true
)

# Request payload for InSilicoPcr#create.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] forward_primer
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_mismatch
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] reverse_primer
#   @return [String]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
InSilicoPcrCreateData = Struct.new(
  :circular,
  :forward_primer,
  :gate,
  :max_mismatch,
  :ok,
  :provenance,
  :result,
  :reverse_primer,
  :template,
  :tool,
  keyword_init: true
)

# KaspPrimerDesign entity data model.
#
# @!attribute [rw] add_secondary_mismatch
#   @return [Boolean, nil]
#
# @!attribute [rw] allele_a
#   @return [String]
#
# @!attribute [rw] allele_b
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_amplicon
#   @return [Integer, nil]
#
# @!attribute [rw] min_amplicon
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] snp_position
#   @return [Integer]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] target_core_tm
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
KaspPrimerDesign = Struct.new(
  :add_secondary_mismatch,
  :allele_a,
  :allele_b,
  :gate,
  :max_amplicon,
  :min_amplicon,
  :ok,
  :provenance,
  :result,
  :snp_position,
  :target,
  :target_core_tm,
  :tool,
  keyword_init: true
)

# Request payload for KaspPrimerDesign#create.
#
# @!attribute [rw] add_secondary_mismatch
#   @return [Boolean, nil]
#
# @!attribute [rw] allele_a
#   @return [String]
#
# @!attribute [rw] allele_b
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_amplicon
#   @return [Integer, nil]
#
# @!attribute [rw] min_amplicon
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] snp_position
#   @return [Integer]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] target_core_tm
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
KaspPrimerDesignCreateData = Struct.new(
  :add_secondary_mismatch,
  :allele_a,
  :allele_b,
  :gate,
  :max_amplicon,
  :min_amplicon,
  :ok,
  :provenance,
  :result,
  :snp_position,
  :target,
  :target_core_tm,
  :tool,
  keyword_init: true
)

# ListTool entity data model.
class ListTool
end

# Request payload for ListTool#load.
class ListToolLoadMatch
end

# MeltingTemperature entity data model.
#
# @!attribute [rw] dntp_mm
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mg_mm
#   @return [Float, nil]
#
# @!attribute [rw] na_mm
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligo_nm
#   @return [Float, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] target_tm
#   @return [Float, nil]
#
# @!attribute [rw] tm_tolerance
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
MeltingTemperature = Struct.new(
  :dntp_mm,
  :gate,
  :mg_mm,
  :na_mm,
  :ok,
  :oligo_nm,
  :provenance,
  :result,
  :sequence,
  :target_tm,
  :tm_tolerance,
  :tool,
  keyword_init: true
)

# Request payload for MeltingTemperature#create.
#
# @!attribute [rw] dntp_mm
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mg_mm
#   @return [Float, nil]
#
# @!attribute [rw] na_mm
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligo_nm
#   @return [Float, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] target_tm
#   @return [Float, nil]
#
# @!attribute [rw] tm_tolerance
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
MeltingTemperatureCreateData = Struct.new(
  :dntp_mm,
  :gate,
  :mg_mm,
  :na_mm,
  :ok,
  :oligo_nm,
  :provenance,
  :result,
  :sequence,
  :target_tm,
  :tm_tolerance,
  :tool,
  keyword_init: true
)

# MotifFinder entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_mismatch
#   @return [Integer, nil]
#
# @!attribute [rw] motif
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] search_reverse_strand
#   @return [Boolean, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
MotifFinder = Struct.new(
  :gate,
  :max_mismatch,
  :motif,
  :ok,
  :provenance,
  :result,
  :search_reverse_strand,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for MotifFinder#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_mismatch
#   @return [Integer, nil]
#
# @!attribute [rw] motif
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] search_reverse_strand
#   @return [Boolean, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
MotifFinderCreateData = Struct.new(
  :gate,
  :max_mismatch,
  :motif,
  :ok,
  :provenance,
  :result,
  :search_reverse_strand,
  :sequence,
  :tool,
  keyword_init: true
)

# MultipleSequenceAlignment entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
MultipleSequenceAlignment = Struct.new(
  :gate,
  :input,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for MultipleSequenceAlignment#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
MultipleSequenceAlignmentCreateData = Struct.new(
  :gate,
  :input,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# OligoAnalysi entity data model.
#
# @!attribute [rw] dntp_mm
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mg_mm
#   @return [Float, nil]
#
# @!attribute [rw] na_mm
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligo_nm
#   @return [Float, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
OligoAnalysi = Struct.new(
  :dntp_mm,
  :gate,
  :mg_mm,
  :na_mm,
  :ok,
  :oligo_nm,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for OligoAnalysi#create.
#
# @!attribute [rw] dntp_mm
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mg_mm
#   @return [Float, nil]
#
# @!attribute [rw] na_mm
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligo_nm
#   @return [Float, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
OligoAnalysiCreateData = Struct.new(
  :dntp_mm,
  :gate,
  :mg_mm,
  :na_mm,
  :ok,
  :oligo_nm,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# OrthologMap entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] source_species
#   @return [String, nil]
#
# @!attribute [rw] symbol
#   @return [Array]
#
# @!attribute [rw] target_species
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] type
#   @return [String, nil]
OrthologMap = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :source_species,
  :symbol,
  :target_species,
  :tool,
  :type,
  keyword_init: true
)

# Request payload for OrthologMap#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] source_species
#   @return [String, nil]
#
# @!attribute [rw] symbol
#   @return [Array]
#
# @!attribute [rw] target_species
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] type
#   @return [String, nil]
OrthologMapCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :source_species,
  :symbol,
  :target_species,
  :tool,
  :type,
  keyword_init: true
)

# PairwiseAlignment entity data model.
#
# @!attribute [rw] gap
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] match
#   @return [Float, nil]
#
# @!attribute [rw] mismatch
#   @return [Float, nil]
#
# @!attribute [rw] mode
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] seq_a
#   @return [String]
#
# @!attribute [rw] seq_b
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PairwiseAlignment = Struct.new(
  :gap,
  :gate,
  :match,
  :mismatch,
  :mode,
  :ok,
  :provenance,
  :result,
  :seq_a,
  :seq_b,
  :tool,
  keyword_init: true
)

# Request payload for PairwiseAlignment#create.
#
# @!attribute [rw] gap
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] match
#   @return [Float, nil]
#
# @!attribute [rw] mismatch
#   @return [Float, nil]
#
# @!attribute [rw] mode
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] seq_a
#   @return [String]
#
# @!attribute [rw] seq_b
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PairwiseAlignmentCreateData = Struct.new(
  :gap,
  :gate,
  :match,
  :mismatch,
  :mode,
  :ok,
  :provenance,
  :result,
  :seq_a,
  :seq_b,
  :tool,
  keyword_init: true
)

# ParseGenbank entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] text
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ParseGenbank = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :text,
  :tool,
  keyword_init: true
)

# Request payload for ParseGenbank#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] text
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ParseGenbankCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :text,
  :tool,
  keyword_init: true
)

# ParseSangerTrace entity data model.
#
# @!attribute [rw] file_base64
#   @return [String]
#
# @!attribute [rw] file_name
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ParseSangerTrace = Struct.new(
  :file_base64,
  :file_name,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for ParseSangerTrace#create.
#
# @!attribute [rw] file_base64
#   @return [String]
#
# @!attribute [rw] file_name
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ParseSangerTraceCreateData = Struct.new(
  :file_base64,
  :file_name,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# PlasmidAnnotate entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PlasmidAnnotate = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for PlasmidAnnotate#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PlasmidAnnotateCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# PlasmidDeepAnnotate entity data model.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PlasmidDeepAnnotate = Struct.new(
  :circular,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for PlasmidDeepAnnotate#create.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PlasmidDeepAnnotateCreateData = Struct.new(
  :circular,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# PlasmidFullReport entity data model.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] top_n
#   @return [Integer, nil]
PlasmidFullReport = Struct.new(
  :circular,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :top_n,
  keyword_init: true
)

# Request payload for PlasmidFullReport#create.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] top_n
#   @return [Integer, nil]
PlasmidFullReportCreateData = Struct.new(
  :circular,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :top_n,
  keyword_init: true
)

# PlasmidIdentify entity data model.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] top_n
#   @return [Integer, nil]
PlasmidIdentify = Struct.new(
  :circular,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :top_n,
  keyword_init: true
)

# Request payload for PlasmidIdentify#create.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] top_n
#   @return [Integer, nil]
PlasmidIdentifyCreateData = Struct.new(
  :circular,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :top_n,
  keyword_init: true
)

# PrimeEditingDesign entity data model.
#
# @!attribute [rw] edit_end
#   @return [Integer]
#
# @!attribute [rw] edit_start
#   @return [Integer]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] inserted_seq
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] pbs_length
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] rtt_homology
#   @return [Integer, nil]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PrimeEditingDesign = Struct.new(
  :edit_end,
  :edit_start,
  :frame_start,
  :gate,
  :inserted_seq,
  :ok,
  :pbs_length,
  :provenance,
  :result,
  :rtt_homology,
  :target,
  :tool,
  keyword_init: true
)

# Request payload for PrimeEditingDesign#create.
#
# @!attribute [rw] edit_end
#   @return [Integer]
#
# @!attribute [rw] edit_start
#   @return [Integer]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] inserted_seq
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] pbs_length
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] rtt_homology
#   @return [Integer, nil]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PrimeEditingDesignCreateData = Struct.new(
  :edit_end,
  :edit_start,
  :frame_start,
  :gate,
  :inserted_seq,
  :ok,
  :pbs_length,
  :provenance,
  :result,
  :rtt_homology,
  :target,
  :tool,
  keyword_init: true
)

# PrimeEditingTwinDesign entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] new_sequence
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlap_length
#   @return [Integer, nil]
#
# @!attribute [rw] pbs_length
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] replace_end
#   @return [Integer]
#
# @!attribute [rw] replace_start
#   @return [Integer]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PrimeEditingTwinDesign = Struct.new(
  :gate,
  :new_sequence,
  :ok,
  :overlap_length,
  :pbs_length,
  :provenance,
  :replace_end,
  :replace_start,
  :result,
  :target,
  :tool,
  keyword_init: true
)

# Request payload for PrimeEditingTwinDesign#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] new_sequence
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlap_length
#   @return [Integer, nil]
#
# @!attribute [rw] pbs_length
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] replace_end
#   @return [Integer]
#
# @!attribute [rw] replace_start
#   @return [Integer]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PrimeEditingTwinDesignCreateData = Struct.new(
  :gate,
  :new_sequence,
  :ok,
  :overlap_length,
  :pbs_length,
  :provenance,
  :replace_end,
  :replace_start,
  :result,
  :target,
  :tool,
  keyword_init: true
)

# PrimerDesign entity data model.
#
# @!attribute [rw] amplicon_max
#   @return [Integer, nil]
#
# @!attribute [rw] amplicon_min
#   @return [Integer, nil]
#
# @!attribute [rw] dntp_mm
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gc_max
#   @return [Float, nil]
#
# @!attribute [rw] gc_min
#   @return [Float, nil]
#
# @!attribute [rw] len_max
#   @return [Integer, nil]
#
# @!attribute [rw] len_min
#   @return [Integer, nil]
#
# @!attribute [rw] len_opt
#   @return [Integer, nil]
#
# @!attribute [rw] max_return
#   @return [Integer, nil]
#
# @!attribute [rw] mg_mm
#   @return [Float, nil]
#
# @!attribute [rw] na_mm
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligo_nm
#   @return [Float, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] target_end
#   @return [Integer, nil]
#
# @!attribute [rw] target_start
#   @return [Integer, nil]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tm_max
#   @return [Float, nil]
#
# @!attribute [rw] tm_max_diff
#   @return [Float, nil]
#
# @!attribute [rw] tm_min
#   @return [Float, nil]
#
# @!attribute [rw] tm_opt
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
PrimerDesign = Struct.new(
  :amplicon_max,
  :amplicon_min,
  :dntp_mm,
  :gate,
  :gc_max,
  :gc_min,
  :len_max,
  :len_min,
  :len_opt,
  :max_return,
  :mg_mm,
  :na_mm,
  :ok,
  :oligo_nm,
  :provenance,
  :result,
  :target_end,
  :target_start,
  :template,
  :tm_max,
  :tm_max_diff,
  :tm_min,
  :tm_opt,
  :tool,
  keyword_init: true
)

# Request payload for PrimerDesign#create.
#
# @!attribute [rw] amplicon_max
#   @return [Integer, nil]
#
# @!attribute [rw] amplicon_min
#   @return [Integer, nil]
#
# @!attribute [rw] dntp_mm
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gc_max
#   @return [Float, nil]
#
# @!attribute [rw] gc_min
#   @return [Float, nil]
#
# @!attribute [rw] len_max
#   @return [Integer, nil]
#
# @!attribute [rw] len_min
#   @return [Integer, nil]
#
# @!attribute [rw] len_opt
#   @return [Integer, nil]
#
# @!attribute [rw] max_return
#   @return [Integer, nil]
#
# @!attribute [rw] mg_mm
#   @return [Float, nil]
#
# @!attribute [rw] na_mm
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligo_nm
#   @return [Float, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] target_end
#   @return [Integer, nil]
#
# @!attribute [rw] target_start
#   @return [Integer, nil]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tm_max
#   @return [Float, nil]
#
# @!attribute [rw] tm_max_diff
#   @return [Float, nil]
#
# @!attribute [rw] tm_min
#   @return [Float, nil]
#
# @!attribute [rw] tm_opt
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
PrimerDesignCreateData = Struct.new(
  :amplicon_max,
  :amplicon_min,
  :dntp_mm,
  :gate,
  :gc_max,
  :gc_min,
  :len_max,
  :len_min,
  :len_opt,
  :max_return,
  :mg_mm,
  :na_mm,
  :ok,
  :oligo_nm,
  :provenance,
  :result,
  :target_end,
  :target_start,
  :template,
  :tm_max,
  :tm_max_diff,
  :tm_min,
  :tm_opt,
  :tool,
  keyword_init: true
)

# PrimerSpecificity entity data model.
#
# @!attribute [rw] forward_primer
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_mismatch
#   @return [Integer, nil]
#
# @!attribute [rw] max_product_length
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] reverse_primer
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PrimerSpecificity = Struct.new(
  :forward_primer,
  :gate,
  :max_mismatch,
  :max_product_length,
  :ok,
  :provenance,
  :result,
  :reverse_primer,
  :tool,
  keyword_init: true
)

# Request payload for PrimerSpecificity#create.
#
# @!attribute [rw] forward_primer
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_mismatch
#   @return [Integer, nil]
#
# @!attribute [rw] max_product_length
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] reverse_primer
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PrimerSpecificityCreateData = Struct.new(
  :forward_primer,
  :gate,
  :max_mismatch,
  :max_product_length,
  :ok,
  :provenance,
  :result,
  :reverse_primer,
  :tool,
  keyword_init: true
)

# ProteaseDigestion entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_mass
#   @return [Float, nil]
#
# @!attribute [rw] max_peptide
#   @return [Integer, nil]
#
# @!attribute [rw] min_mass
#   @return [Float, nil]
#
# @!attribute [rw] missed_cleavage
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] protease
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ProteaseDigestion = Struct.new(
  :gate,
  :max_mass,
  :max_peptide,
  :min_mass,
  :missed_cleavage,
  :ok,
  :protease,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for ProteaseDigestion#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_mass
#   @return [Float, nil]
#
# @!attribute [rw] max_peptide
#   @return [Integer, nil]
#
# @!attribute [rw] min_mass
#   @return [Float, nil]
#
# @!attribute [rw] missed_cleavage
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] protease
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ProteaseDigestionCreateData = Struct.new(
  :gate,
  :max_mass,
  :max_peptide,
  :min_mass,
  :missed_cleavage,
  :ok,
  :protease,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# ProteinAnnotatePoll entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] job_id
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ProteinAnnotatePoll = Struct.new(
  :gate,
  :job_id,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for ProteinAnnotatePoll#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] job_id
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ProteinAnnotatePollCreateData = Struct.new(
  :gate,
  :job_id,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# ProteinAnnotateSubmit entity data model.
#
# @!attribute [rw] appl
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] goterm
#   @return [Boolean, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ProteinAnnotateSubmit = Struct.new(
  :appl,
  :gate,
  :goterm,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for ProteinAnnotateSubmit#create.
#
# @!attribute [rw] appl
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] goterm
#   @return [Boolean, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ProteinAnnotateSubmitCreateData = Struct.new(
  :appl,
  :gate,
  :goterm,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# ProteinHydrophobicity entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] scale
#   @return [String, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] window
#   @return [Integer, nil]
ProteinHydrophobicity = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :scale,
  :sequence,
  :tool,
  :window,
  keyword_init: true
)

# Request payload for ProteinHydrophobicity#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] scale
#   @return [String, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] window
#   @return [Integer, nil]
ProteinHydrophobicityCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :scale,
  :sequence,
  :tool,
  :window,
  keyword_init: true
)

# ProteinProperty entity data model.
#
# @!attribute [rw] charge_step
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ProteinProperty = Struct.new(
  :charge_step,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for ProteinProperty#create.
#
# @!attribute [rw] charge_step
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
ProteinPropertyCreateData = Struct.new(
  :charge_step,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# RandomSequence entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gc_content
#   @return [Float, nil]
#
# @!attribute [rw] kind
#   @return [String, nil]
#
# @!attribute [rw] length
#   @return [Integer]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
RandomSequence = Struct.new(
  :gate,
  :gc_content,
  :kind,
  :length,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for RandomSequence#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gc_content
#   @return [Float, nil]
#
# @!attribute [rw] kind
#   @return [String, nil]
#
# @!attribute [rw] length
#   @return [Integer]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
RandomSequenceCreateData = Struct.new(
  :gate,
  :gc_content,
  :kind,
  :length,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# RestrictionSite entity data model.
#
# @!attribute [rw] enzyme
#   @return [Array, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
RestrictionSite = Struct.new(
  :enzyme,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for RestrictionSite#create.
#
# @!attribute [rw] enzyme
#   @return [Array, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
RestrictionSiteCreateData = Struct.new(
  :enzyme,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# ReverseComplement entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] type
#   @return [String, nil]
ReverseComplement = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :type,
  keyword_init: true
)

# Request payload for ReverseComplement#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] type
#   @return [String, nil]
ReverseComplementCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :type,
  keyword_init: true
)

# ReverseTranslate entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mode
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] protein
#   @return [String]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ReverseTranslate = Struct.new(
  :gate,
  :mode,
  :ok,
  :organism,
  :protein,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for ReverseTranslate#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mode
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] protein
#   @return [String]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
ReverseTranslateCreateData = Struct.new(
  :gate,
  :mode,
  :ok,
  :organism,
  :protein,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# RnaFold entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
RnaFold = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for RnaFold#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
RnaFoldCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# SangerVsReference entity data model.
#
# @!attribute [rw] file_base64
#   @return [String, nil]
#
# @!attribute [rw] file_name
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] min_coverage
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] read
#   @return [String, nil]
#
# @!attribute [rw] reference
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SangerVsReference = Struct.new(
  :file_base64,
  :file_name,
  :gate,
  :min_coverage,
  :ok,
  :provenance,
  :read,
  :reference,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for SangerVsReference#create.
#
# @!attribute [rw] file_base64
#   @return [String, nil]
#
# @!attribute [rw] file_name
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] min_coverage
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] read
#   @return [String, nil]
#
# @!attribute [rw] reference
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SangerVsReferenceCreateData = Struct.new(
  :file_base64,
  :file_name,
  :gate,
  :min_coverage,
  :ok,
  :provenance,
  :read,
  :reference,
  :result,
  :tool,
  keyword_init: true
)

# SavePermalink entity data model.
#
# @!attribute [rw] arg
#   @return [Hash]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SavePermalink = Struct.new(
  :arg,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for SavePermalink#create.
#
# @!attribute [rw] arg
#   @return [Hash]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SavePermalinkCreateData = Struct.new(
  :arg,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# SeqfileStat entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] quality_offset
#   @return [Integer, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SeqfileStat = Struct.new(
  :gate,
  :input,
  :ok,
  :provenance,
  :quality_offset,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for SeqfileStat#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] quality_offset
#   @return [Integer, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SeqfileStatCreateData = Struct.new(
  :gate,
  :input,
  :ok,
  :provenance,
  :quality_offset,
  :result,
  :tool,
  keyword_init: true
)

# SequenceFetch entity data model.
#
# @!attribute [rw] accession
#   @return [String]
#
# @!attribute [rw] db
#   @return [String, nil]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SequenceFetch = Struct.new(
  :accession,
  :db,
  :format,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for SequenceFetch#create.
#
# @!attribute [rw] accession
#   @return [String]
#
# @!attribute [rw] db
#   @return [String, nil]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SequenceFetchCreateData = Struct.new(
  :accession,
  :db,
  :format,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# SequenceFormatConvert entity data model.
#
# @!attribute [rw] from
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] to
#   @return [String, nil]
#
# @!attribute [rw] tool
#   @return [String]
SequenceFormatConvert = Struct.new(
  :from,
  :gate,
  :input,
  :ok,
  :provenance,
  :result,
  :to,
  :tool,
  keyword_init: true
)

# Request payload for SequenceFormatConvert#create.
#
# @!attribute [rw] from
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] to
#   @return [String, nil]
#
# @!attribute [rw] tool
#   @return [String]
SequenceFormatConvertCreateData = Struct.new(
  :from,
  :gate,
  :input,
  :ok,
  :provenance,
  :result,
  :to,
  :tool,
  keyword_init: true
)

# SequenceReport entity data model.
#
# @!attribute [rw] end_primer_length
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_orf
#   @return [Integer, nil]
#
# @!attribute [rw] min_orf_aa
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SequenceReport = Struct.new(
  :end_primer_length,
  :gate,
  :max_orf,
  :min_orf_aa,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for SequenceReport#create.
#
# @!attribute [rw] end_primer_length
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_orf
#   @return [Integer, nil]
#
# @!attribute [rw] min_orf_aa
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SequenceReportCreateData = Struct.new(
  :end_primer_length,
  :gate,
  :max_orf,
  :min_orf_aa,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# SequenceSearch entity data model.
#
# @!attribute [rw] db
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [String, nil]
#
# @!attribute [rw] max_result
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] term
#   @return [String, nil]
#
# @!attribute [rw] tool
#   @return [String]
SequenceSearch = Struct.new(
  :db,
  :gate,
  :gene,
  :max_result,
  :ok,
  :organism,
  :provenance,
  :result,
  :term,
  :tool,
  keyword_init: true
)

# Request payload for SequenceSearch#create.
#
# @!attribute [rw] db
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gene
#   @return [String, nil]
#
# @!attribute [rw] max_result
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] term
#   @return [String, nil]
#
# @!attribute [rw] tool
#   @return [String]
SequenceSearchCreateData = Struct.new(
  :db,
  :gate,
  :gene,
  :max_result,
  :ok,
  :organism,
  :provenance,
  :result,
  :term,
  :tool,
  keyword_init: true
)

# SequencingReadbackVerify entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] min_supporting_read
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] read
#   @return [String]
#
# @!attribute [rw] reference
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SequencingReadbackVerify = Struct.new(
  :gate,
  :min_supporting_read,
  :ok,
  :provenance,
  :read,
  :reference,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for SequencingReadbackVerify#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] min_supporting_read
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] read
#   @return [String]
#
# @!attribute [rw] reference
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SequencingReadbackVerifyCreateData = Struct.new(
  :gate,
  :min_supporting_read,
  :ok,
  :provenance,
  :read,
  :reference,
  :result,
  :tool,
  keyword_init: true
)

# SessionCreate entity data model.
#
# @!attribute [rw] entry
#   @return [Hash, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SessionCreate = Struct.new(
  :entry,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for SessionCreate#create.
#
# @!attribute [rw] entry
#   @return [Hash, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
SessionCreateCreateData = Struct.new(
  :entry,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# SessionGet entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] name
#   @return [Array, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] session_id
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SessionGet = Struct.new(
  :gate,
  :name,
  :ok,
  :provenance,
  :result,
  :session_id,
  :tool,
  keyword_init: true
)

# Request payload for SessionGet#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] name
#   @return [Array, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] session_id
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SessionGetCreateData = Struct.new(
  :gate,
  :name,
  :ok,
  :provenance,
  :result,
  :session_id,
  :tool,
  keyword_init: true
)

# SessionRun entity data model.
#
# @!attribute [rw] arg
#   @return [Hash, nil]
#
# @!attribute [rw] from_session
#   @return [Hash, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] session_id
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] write_back
#   @return [Hash, nil]
SessionRun = Struct.new(
  :arg,
  :from_session,
  :gate,
  :ok,
  :provenance,
  :result,
  :session_id,
  :tool,
  :write_back,
  keyword_init: true
)

# Request payload for SessionRun#create.
#
# @!attribute [rw] arg
#   @return [Hash, nil]
#
# @!attribute [rw] from_session
#   @return [Hash, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] session_id
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] write_back
#   @return [Hash, nil]
SessionRunCreateData = Struct.new(
  :arg,
  :from_session,
  :gate,
  :ok,
  :provenance,
  :result,
  :session_id,
  :tool,
  :write_back,
  keyword_init: true
)

# SessionSet entity data model.
#
# @!attribute [rw] entry
#   @return [Hash]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] session_id
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SessionSet = Struct.new(
  :entry,
  :gate,
  :ok,
  :provenance,
  :result,
  :session_id,
  :tool,
  keyword_init: true
)

# Request payload for SessionSet#create.
#
# @!attribute [rw] entry
#   @return [Hash]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] session_id
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SessionSetCreateData = Struct.new(
  :entry,
  :gate,
  :ok,
  :provenance,
  :result,
  :session_id,
  :tool,
  keyword_init: true
)

# SirnaDesign entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] min_reynold
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sh_rna_loop
#   @return [String, nil]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SirnaDesign = Struct.new(
  :gate,
  :min_reynold,
  :ok,
  :provenance,
  :result,
  :sh_rna_loop,
  :target,
  :tool,
  keyword_init: true
)

# Request payload for SirnaDesign#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] min_reynold
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sh_rna_loop
#   @return [String, nil]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SirnaDesignCreateData = Struct.new(
  :gate,
  :min_reynold,
  :ok,
  :provenance,
  :result,
  :sh_rna_loop,
  :target,
  :tool,
  keyword_init: true
)

# SiteDirectedMutagenesi entity data model.
#
# @!attribute [rw] arm_tm_target
#   @return [Float, nil]
#
# @!attribute [rw] dntp_mm
#   @return [Float, nil]
#
# @!attribute [rw] edit_kind
#   @return [String, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mg_mm
#   @return [Float, nil]
#
# @!attribute [rw] na_mm
#   @return [Float, nil]
#
# @!attribute [rw] new_base
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligo_nm
#   @return [Float, nil]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] position
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] residue
#   @return [Integer, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] style
#   @return [String, nil]
#
# @!attribute [rw] target_aa
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SiteDirectedMutagenesi = Struct.new(
  :arm_tm_target,
  :dntp_mm,
  :edit_kind,
  :frame_start,
  :gate,
  :mg_mm,
  :na_mm,
  :new_base,
  :ok,
  :oligo_nm,
  :organism,
  :position,
  :provenance,
  :residue,
  :result,
  :style,
  :target_aa,
  :template,
  :tool,
  keyword_init: true
)

# Request payload for SiteDirectedMutagenesi#create.
#
# @!attribute [rw] arm_tm_target
#   @return [Float, nil]
#
# @!attribute [rw] dntp_mm
#   @return [Float, nil]
#
# @!attribute [rw] edit_kind
#   @return [String, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mg_mm
#   @return [Float, nil]
#
# @!attribute [rw] na_mm
#   @return [Float, nil]
#
# @!attribute [rw] new_base
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligo_nm
#   @return [Float, nil]
#
# @!attribute [rw] organism
#   @return [String, nil]
#
# @!attribute [rw] position
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] residue
#   @return [Integer, nil]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] style
#   @return [String, nil]
#
# @!attribute [rw] target_aa
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SiteDirectedMutagenesiCreateData = Struct.new(
  :arm_tm_target,
  :dntp_mm,
  :edit_kind,
  :frame_start,
  :gate,
  :mg_mm,
  :na_mm,
  :new_base,
  :ok,
  :oligo_nm,
  :organism,
  :position,
  :provenance,
  :residue,
  :result,
  :style,
  :target_aa,
  :template,
  :tool,
  keyword_init: true
)

# Translate entity data model.
#
# @!attribute [rw] frame
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] to_stop
#   @return [Boolean, nil]
#
# @!attribute [rw] tool
#   @return [String]
Translate = Struct.new(
  :frame,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :to_stop,
  :tool,
  keyword_init: true
)

# Request payload for Translate#create.
#
# @!attribute [rw] frame
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] to_stop
#   @return [Boolean, nil]
#
# @!attribute [rw] tool
#   @return [String]
TranslateCreateData = Struct.new(
  :frame,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :to_stop,
  :tool,
  keyword_init: true
)

# VariantAnnotate entity data model.
#
# @!attribute [rw] assembly
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] variant
#   @return [String]
VariantAnnotate = Struct.new(
  :assembly,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  :variant,
  keyword_init: true
)

# Request payload for VariantAnnotate#create.
#
# @!attribute [rw] assembly
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] variant
#   @return [String]
VariantAnnotateCreateData = Struct.new(
  :assembly,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  :variant,
  keyword_init: true
)

# VariantComparator entity data model.
#
# @!attribute [rw] coding
#   @return [Boolean, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] query
#   @return [String]
#
# @!attribute [rw] reference
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
VariantComparator = Struct.new(
  :coding,
  :frame_start,
  :gate,
  :ok,
  :provenance,
  :query,
  :reference,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for VariantComparator#create.
#
# @!attribute [rw] coding
#   @return [Boolean, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] query
#   @return [String]
#
# @!attribute [rw] reference
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
VariantComparatorCreateData = Struct.new(
  :coding,
  :frame_start,
  :gate,
  :ok,
  :provenance,
  :query,
  :reference,
  :result,
  :tool,
  keyword_init: true
)

# VerifyAssembly entity data model.
#
# @!attribute [rw] arm_tm_target
#   @return [Float, nil]
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] claimed_construct
#   @return [String]
#
# @!attribute [rw] coding
#   @return [Boolean, nil]
#
# @!attribute [rw] enzyme
#   @return [String, nil]
#
# @!attribute [rw] enzyme3
#   @return [String, nil]
#
# @!attribute [rw] enzyme5
#   @return [String, nil]
#
# @!attribute [rw] fragment
#   @return [Array, nil]
#
# @!attribute [rw] fragment_pcr
#   @return [Array, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insert
#   @return [String, nil]
#
# @!attribute [rw] insert_pcr
#   @return [Hash, nil]
#
# @!attribute [rw] method
#   @return [String]
#
# @!attribute [rw] name
#   @return [Array, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlap_len
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] vector
#   @return [String, nil]
#
# @!attribute [rw] vector_pcr
#   @return [Hash, nil]
VerifyAssembly = Struct.new(
  :arm_tm_target,
  :circular,
  :claimed_construct,
  :coding,
  :enzyme,
  :enzyme3,
  :enzyme5,
  :fragment,
  :fragment_pcr,
  :frame_start,
  :gate,
  :insert,
  :insert_pcr,
  :method,
  :name,
  :ok,
  :overlap_len,
  :provenance,
  :result,
  :tool,
  :vector,
  :vector_pcr,
  keyword_init: true
)

# Request payload for VerifyAssembly#create.
#
# @!attribute [rw] arm_tm_target
#   @return [Float, nil]
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] claimed_construct
#   @return [String]
#
# @!attribute [rw] coding
#   @return [Boolean, nil]
#
# @!attribute [rw] enzyme
#   @return [String, nil]
#
# @!attribute [rw] enzyme3
#   @return [String, nil]
#
# @!attribute [rw] enzyme5
#   @return [String, nil]
#
# @!attribute [rw] fragment
#   @return [Array, nil]
#
# @!attribute [rw] fragment_pcr
#   @return [Array, nil]
#
# @!attribute [rw] frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insert
#   @return [String, nil]
#
# @!attribute [rw] insert_pcr
#   @return [Hash, nil]
#
# @!attribute [rw] method
#   @return [String]
#
# @!attribute [rw] name
#   @return [Array, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlap_len
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] vector
#   @return [String, nil]
#
# @!attribute [rw] vector_pcr
#   @return [Hash, nil]
VerifyAssemblyCreateData = Struct.new(
  :arm_tm_target,
  :circular,
  :claimed_construct,
  :coding,
  :enzyme,
  :enzyme3,
  :enzyme5,
  :fragment,
  :fragment_pcr,
  :frame_start,
  :gate,
  :insert,
  :insert_pcr,
  :method,
  :name,
  :ok,
  :overlap_len,
  :provenance,
  :result,
  :tool,
  :vector,
  :vector_pcr,
  keyword_init: true
)

# VerifyConstruct entity data model.
#
# @!attribute [rw] claimed_construct
#   @return [String]
#
# @!attribute [rw] expected_frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insert_forward_primer
#   @return [String]
#
# @!attribute [rw] insert_reverse_primer
#   @return [String]
#
# @!attribute [rw] insert_template
#   @return [String]
#
# @!attribute [rw] max_primer_mismatch
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] template_circular
#   @return [Boolean, nil]
#
# @!attribute [rw] tool
#   @return [String]
VerifyConstruct = Struct.new(
  :claimed_construct,
  :expected_frame_start,
  :gate,
  :insert_forward_primer,
  :insert_reverse_primer,
  :insert_template,
  :max_primer_mismatch,
  :ok,
  :provenance,
  :result,
  :template_circular,
  :tool,
  keyword_init: true
)

# Request payload for VerifyConstruct#create.
#
# @!attribute [rw] claimed_construct
#   @return [String]
#
# @!attribute [rw] expected_frame_start
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insert_forward_primer
#   @return [String]
#
# @!attribute [rw] insert_reverse_primer
#   @return [String]
#
# @!attribute [rw] insert_template
#   @return [String]
#
# @!attribute [rw] max_primer_mismatch
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] template_circular
#   @return [Boolean, nil]
#
# @!attribute [rw] tool
#   @return [String]
VerifyConstructCreateData = Struct.new(
  :claimed_construct,
  :expected_frame_start,
  :gate,
  :insert_forward_primer,
  :insert_reverse_primer,
  :insert_template,
  :max_primer_mismatch,
  :ok,
  :provenance,
  :result,
  :template_circular,
  :tool,
  keyword_init: true
)

# VirtualGel entity data model.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] enzyme
#   @return [Array, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ladder
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
VirtualGel = Struct.new(
  :circular,
  :enzyme,
  :gate,
  :ladder,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for VirtualGel#create.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] enzyme
#   @return [Array, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ladder
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
VirtualGelCreateData = Struct.new(
  :circular,
  :enzyme,
  :gate,
  :ladder,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# VolcanoPlotData entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] row
#   @return [Array]
#
# @!attribute [rw] tool
#   @return [String]
VolcanoPlotData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :row,
  :tool,
  keyword_init: true
)

# Request payload for VolcanoPlotData#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] row
#   @return [Array]
#
# @!attribute [rw] tool
#   @return [String]
VolcanoPlotDataCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :row,
  :tool,
  keyword_init: true
)

# WebSearch entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_result
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] query
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
WebSearch = Struct.new(
  :gate,
  :max_result,
  :ok,
  :provenance,
  :query,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for WebSearch#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_result
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] query
#   @return [String]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] tool
#   @return [String]
WebSearchCreateData = Struct.new(
  :gate,
  :max_result,
  :ok,
  :provenance,
  :query,
  :result,
  :tool,
  keyword_init: true
)

