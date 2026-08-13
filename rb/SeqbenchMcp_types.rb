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
# @!attribute [rw] frameStart
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
# @!attribute [rw] targetPosition
#   @return [Integer, nil]
#
# @!attribute [rw] tool
#   @return [String]
BaseEditingDesign = Struct.new(
  :editor,
  :frameStart,
  :gate,
  :ok,
  :provenance,
  :result,
  :target,
  :targetPosition,
  :tool,
  keyword_init: true
)

# Request payload for BaseEditingDesign#create.
#
# @!attribute [rw] editor
#   @return [String, nil]
#
# @!attribute [rw] frameStart
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
# @!attribute [rw] targetPosition
#   @return [Integer, nil]
#
# @!attribute [rw] tool
#   @return [String]
BaseEditingDesignCreateData = Struct.new(
  :editor,
  :frameStart,
  :gate,
  :ok,
  :provenance,
  :result,
  :target,
  :targetPosition,
  :tool,
  keyword_init: true
)

# Batch entity data model.
#
# @!attribute [rw] args
#   @return [Hash, nil]
#
# @!attribute [rw] capped
#   @return [Boolean]
#
# @!attribute [rw] columns
#   @return [Array]
#
# @!attribute [rw] count
#   @return [Integer]
#
# @!attribute [rw] errors
#   @return [Integer]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] limit
#   @return [Integer]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] rows
#   @return [Array]
#
# @!attribute [rw] tool
#   @return [String]
Batch = Struct.new(
  :args,
  :capped,
  :columns,
  :count,
  :errors,
  :input,
  :limit,
  :provenance,
  :rows,
  :tool,
  keyword_init: true
)

# Request payload for Batch#load.
#
# @!attribute [rw] args
#   @return [Hash, nil]
#
# @!attribute [rw] capped
#   @return [Boolean, nil]
#
# @!attribute [rw] columns
#   @return [Array, nil]
#
# @!attribute [rw] count
#   @return [Integer, nil]
#
# @!attribute [rw] errors
#   @return [Integer, nil]
#
# @!attribute [rw] input
#   @return [String, nil]
#
# @!attribute [rw] limit
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash, nil]
#
# @!attribute [rw] rows
#   @return [Array, nil]
#
# @!attribute [rw] tool
#   @return [String, nil]
BatchLoadMatch = Struct.new(
  :args,
  :capped,
  :columns,
  :count,
  :errors,
  :input,
  :limit,
  :provenance,
  :rows,
  :tool,
  keyword_init: true
)

# Request payload for Batch#create.
#
# @!attribute [rw] args
#   @return [Hash, nil]
#
# @!attribute [rw] capped
#   @return [Boolean]
#
# @!attribute [rw] columns
#   @return [Array]
#
# @!attribute [rw] count
#   @return [Integer]
#
# @!attribute [rw] errors
#   @return [Integer]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] limit
#   @return [Integer]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] rows
#   @return [Array]
#
# @!attribute [rw] tool
#   @return [String]
BatchCreateData = Struct.new(
  :args,
  :capped,
  :columns,
  :count,
  :errors,
  :input,
  :limit,
  :provenance,
  :rows,
  :tool,
  keyword_init: true
)

# BatchWorkflow entity data model.
#
# @!attribute [rw] capped
#   @return [Boolean]
#
# @!attribute [rw] columns
#   @return [Array]
#
# @!attribute [rw] count
#   @return [Integer]
#
# @!attribute [rw] errors
#   @return [Integer]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] limit
#   @return [Integer]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] rows
#   @return [Array]
#
# @!attribute [rw] steps
#   @return [Array]
BatchWorkflow = Struct.new(
  :capped,
  :columns,
  :count,
  :errors,
  :input,
  :limit,
  :provenance,
  :rows,
  :steps,
  keyword_init: true
)

# Request payload for BatchWorkflow#load.
#
# @!attribute [rw] capped
#   @return [Boolean, nil]
#
# @!attribute [rw] columns
#   @return [Array, nil]
#
# @!attribute [rw] count
#   @return [Integer, nil]
#
# @!attribute [rw] errors
#   @return [Integer, nil]
#
# @!attribute [rw] input
#   @return [String, nil]
#
# @!attribute [rw] limit
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash, nil]
#
# @!attribute [rw] rows
#   @return [Array, nil]
#
# @!attribute [rw] steps
#   @return [Array, nil]
BatchWorkflowLoadMatch = Struct.new(
  :capped,
  :columns,
  :count,
  :errors,
  :input,
  :limit,
  :provenance,
  :rows,
  :steps,
  keyword_init: true
)

# Request payload for BatchWorkflow#create.
#
# @!attribute [rw] capped
#   @return [Boolean]
#
# @!attribute [rw] columns
#   @return [Array]
#
# @!attribute [rw] count
#   @return [Integer]
#
# @!attribute [rw] errors
#   @return [Integer]
#
# @!attribute [rw] input
#   @return [String]
#
# @!attribute [rw] limit
#   @return [Integer]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] rows
#   @return [Array]
#
# @!attribute [rw] steps
#   @return [Array]
BatchWorkflowCreateData = Struct.new(
  :capped,
  :columns,
  :count,
  :errors,
  :input,
  :limit,
  :provenance,
  :rows,
  :steps,
  keyword_init: true
)

# CharacterizeSequence entity data model.
#
# @!attribute [rw] endPrimerLength
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxOrfs
#   @return [Integer, nil]
#
# @!attribute [rw] minOrfAa
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
  :endPrimerLength,
  :gate,
  :maxOrfs,
  :minOrfAa,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for CharacterizeSequence#create.
#
# @!attribute [rw] endPrimerLength
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxOrfs
#   @return [Integer, nil]
#
# @!attribute [rw] minOrfAa
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
  :endPrimerLength,
  :gate,
  :maxOrfs,
  :minOrfAa,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# CloningSimulate entity data model.
#
# @!attribute [rw] armTmTarget
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
# @!attribute [rw] fragments
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
# @!attribute [rw] names
#   @return [Array, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlapLen
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
  :armTmTarget,
  :circular,
  :enzyme,
  :enzyme3,
  :enzyme5,
  :fragments,
  :gate,
  :insert,
  :method,
  :names,
  :ok,
  :overlapLen,
  :provenance,
  :result,
  :tool,
  :vector,
  keyword_init: true
)

# Request payload for CloningSimulate#create.
#
# @!attribute [rw] armTmTarget
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
# @!attribute [rw] fragments
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
# @!attribute [rw] names
#   @return [Array, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlapLen
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
  :armTmTarget,
  :circular,
  :enzyme,
  :enzyme3,
  :enzyme5,
  :fragments,
  :gate,
  :insert,
  :method,
  :names,
  :ok,
  :overlapLen,
  :provenance,
  :result,
  :tool,
  :vector,
  keyword_init: true
)

# CodonAdaptationIndex entity data model.
#
# @!attribute [rw] frameStart
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
# @!attribute [rw] rareThreshold
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
  :frameStart,
  :gate,
  :ok,
  :organism,
  :provenance,
  :rareThreshold,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for CodonAdaptationIndex#create.
#
# @!attribute [rw] frameStart
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
# @!attribute [rw] rareThreshold
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
  :frameStart,
  :gate,
  :ok,
  :organism,
  :provenance,
  :rareThreshold,
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
# @!attribute [rw] avoidEnzymes
#   @return [Array, nil]
#
# @!attribute [rw] crypticOrfMinAa
#   @return [Integer, nil]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gcHigh
#   @return [Float, nil]
#
# @!attribute [rw] gcLow
#   @return [Float, nil]
#
# @!attribute [rw] gcWindow
#   @return [Integer, nil]
#
# @!attribute [rw] homopolymerMin
#   @return [Integer, nil]
#
# @!attribute [rw] maxPasses
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
  :avoidEnzymes,
  :crypticOrfMinAa,
  :frameStart,
  :gate,
  :gcHigh,
  :gcLow,
  :gcWindow,
  :homopolymerMin,
  :maxPasses,
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
# @!attribute [rw] avoidEnzymes
#   @return [Array, nil]
#
# @!attribute [rw] crypticOrfMinAa
#   @return [Integer, nil]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gcHigh
#   @return [Float, nil]
#
# @!attribute [rw] gcLow
#   @return [Float, nil]
#
# @!attribute [rw] gcWindow
#   @return [Integer, nil]
#
# @!attribute [rw] homopolymerMin
#   @return [Integer, nil]
#
# @!attribute [rw] maxPasses
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
  :avoidEnzymes,
  :crypticOrfMinAa,
  :frameStart,
  :gate,
  :gcHigh,
  :gcLow,
  :gcWindow,
  :homopolymerMin,
  :maxPasses,
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
# @!attribute [rw] avoidEnzymes
#   @return [Array, nil]
#
# @!attribute [rw] crypticOrfMinAa
#   @return [Integer, nil]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gcHigh
#   @return [Float, nil]
#
# @!attribute [rw] gcLow
#   @return [Float, nil]
#
# @!attribute [rw] gcWindow
#   @return [Integer, nil]
#
# @!attribute [rw] homopolymerMin
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
  :avoidEnzymes,
  :crypticOrfMinAa,
  :frameStart,
  :gate,
  :gcHigh,
  :gcLow,
  :gcWindow,
  :homopolymerMin,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for ConstructQc#create.
#
# @!attribute [rw] avoidEnzymes
#   @return [Array, nil]
#
# @!attribute [rw] crypticOrfMinAa
#   @return [Integer, nil]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gcHigh
#   @return [Float, nil]
#
# @!attribute [rw] gcLow
#   @return [Float, nil]
#
# @!attribute [rw] gcWindow
#   @return [Integer, nil]
#
# @!attribute [rw] homopolymerMin
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
  :avoidEnzymes,
  :crypticOrfMinAa,
  :frameStart,
  :gate,
  :gcHigh,
  :gcLow,
  :gcWindow,
  :homopolymerMin,
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
# @!attribute [rw] minScore
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
# @!attribute [rw] searchReverseStrand
#   @return [Boolean, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrisprGrnaDesign = Struct.new(
  :gate,
  :minScore,
  :nuclease,
  :ok,
  :provenance,
  :result,
  :searchReverseStrand,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for CrisprGrnaDesign#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] minScore
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
# @!attribute [rw] searchReverseStrand
#   @return [Boolean, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrisprGrnaDesignCreateData = Struct.new(
  :gate,
  :minScore,
  :nuclease,
  :ok,
  :provenance,
  :result,
  :searchReverseStrand,
  :sequence,
  :tool,
  keyword_init: true
)

# CrisprHdrDonor entity data model.
#
# @!attribute [rw] armLength
#   @return [Integer, nil]
#
# @!attribute [rw] blockPam
#   @return [Boolean, nil]
#
# @!attribute [rw] designGenotypingPrimers
#   @return [Boolean, nil]
#
# @!attribute [rw] editEnd
#   @return [Integer, nil]
#
# @!attribute [rw] editStart
#   @return [Integer, nil]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] guideEnd
#   @return [Integer, nil]
#
# @!attribute [rw] guideStart
#   @return [Integer, nil]
#
# @!attribute [rw] guideStrand
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
# @!attribute [rw] targetSequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrisprHdrDonor = Struct.new(
  :armLength,
  :blockPam,
  :designGenotypingPrimers,
  :editEnd,
  :editStart,
  :frameStart,
  :gate,
  :guideEnd,
  :guideStart,
  :guideStrand,
  :nuclease,
  :ok,
  :provenance,
  :replacement,
  :result,
  :targetSequence,
  :tool,
  keyword_init: true
)

# Request payload for CrisprHdrDonor#create.
#
# @!attribute [rw] armLength
#   @return [Integer, nil]
#
# @!attribute [rw] blockPam
#   @return [Boolean, nil]
#
# @!attribute [rw] designGenotypingPrimers
#   @return [Boolean, nil]
#
# @!attribute [rw] editEnd
#   @return [Integer, nil]
#
# @!attribute [rw] editStart
#   @return [Integer, nil]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] guideEnd
#   @return [Integer, nil]
#
# @!attribute [rw] guideStart
#   @return [Integer, nil]
#
# @!attribute [rw] guideStrand
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
# @!attribute [rw] targetSequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrisprHdrDonorCreateData = Struct.new(
  :armLength,
  :blockPam,
  :designGenotypingPrimers,
  :editEnd,
  :editStart,
  :frameStart,
  :gate,
  :guideEnd,
  :guideStart,
  :guideStrand,
  :nuclease,
  :ok,
  :provenance,
  :replacement,
  :result,
  :targetSequence,
  :tool,
  keyword_init: true
)

# CrisprOfftargetCheck entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxMismatches
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
  :maxMismatches,
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
# @!attribute [rw] maxMismatches
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
  :maxMismatches,
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
# @!attribute [rw] sequenceA
#   @return [String]
#
# @!attribute [rw] sequenceB
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrossDimer = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequenceA,
  :sequenceB,
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
# @!attribute [rw] sequenceA
#   @return [String]
#
# @!attribute [rw] sequenceB
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
CrossDimerCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :sequenceA,
  :sequenceB,
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
# @!attribute [rw] massNg
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
# @!attribute [rw] volumeUl
#   @return [Float, nil]
DnaMolarity = Struct.new(
  :gate,
  :length,
  :massNg,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :type,
  :volumeUl,
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
# @!attribute [rw] massNg
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
# @!attribute [rw] volumeUl
#   @return [Float, nil]
DnaMolarityCreateData = Struct.new(
  :gate,
  :length,
  :massNg,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :type,
  :volumeUl,
  keyword_init: true
)

# DoubleDigest entity data model.
#
# @!attribute [rw] enzymeA
#   @return [String]
#
# @!attribute [rw] enzymeB
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
  :enzymeA,
  :enzymeB,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for DoubleDigest#create.
#
# @!attribute [rw] enzymeA
#   @return [String]
#
# @!attribute [rw] enzymeB
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
  :enzymeA,
  :enzymeB,
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
# @!attribute [rw] reactions
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
  :reactions,
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
# @!attribute [rw] reactions
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
  :reactions,
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
# @!attribute [rw] protocolName
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] reactions
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
  :protocolName,
  :provenance,
  :reactions,
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
# @!attribute [rw] protocolName
#   @return [String, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] reactions
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
  :protocolName,
  :provenance,
  :reactions,
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
# @!attribute [rw] reactions
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
  :reactions,
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
# @!attribute [rw] reactions
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
  :reactions,
  :result,
  :tool,
  keyword_init: true
)

# ExpressionHeatmapCluster entity data model.
#
# @!attribute [rw] clusterCols
#   @return [Boolean, nil]
#
# @!attribute [rw] clusterRows
#   @return [Boolean, nil]
#
# @!attribute [rw] distanceMetric
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] genes
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
# @!attribute [rw] samples
#   @return [Array]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] values
#   @return [Array]
#
# @!attribute [rw] zScoreRows
#   @return [Boolean, nil]
ExpressionHeatmapCluster = Struct.new(
  :clusterCols,
  :clusterRows,
  :distanceMetric,
  :gate,
  :genes,
  :linkage,
  :ok,
  :provenance,
  :result,
  :samples,
  :tool,
  :values,
  :zScoreRows,
  keyword_init: true
)

# Request payload for ExpressionHeatmapCluster#create.
#
# @!attribute [rw] clusterCols
#   @return [Boolean, nil]
#
# @!attribute [rw] clusterRows
#   @return [Boolean, nil]
#
# @!attribute [rw] distanceMetric
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] genes
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
# @!attribute [rw] samples
#   @return [Array]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] values
#   @return [Array]
#
# @!attribute [rw] zScoreRows
#   @return [Boolean, nil]
ExpressionHeatmapClusterCreateData = Struct.new(
  :clusterCols,
  :clusterRows,
  :distanceMetric,
  :gate,
  :genes,
  :linkage,
  :ok,
  :provenance,
  :result,
  :samples,
  :tool,
  :values,
  :zScoreRows,
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
# @!attribute [rw] qualityOffset
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
  :qualityOffset,
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
# @!attribute [rw] qualityOffset
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
  :qualityOffset,
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
# @!attribute [rw] minLength
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] qualityOffset
#   @return [Integer, nil]
#
# @!attribute [rw] qualityThreshold
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
  :minLength,
  :ok,
  :provenance,
  :qualityOffset,
  :qualityThreshold,
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
# @!attribute [rw] minLength
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] qualityOffset
#   @return [Integer, nil]
#
# @!attribute [rw] qualityThreshold
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
  :minLength,
  :ok,
  :provenance,
  :qualityOffset,
  :qualityThreshold,
  :result,
  :tool,
  keyword_init: true
)

# FindOrf entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] minAaLength
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] requireStop
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
  :minAaLength,
  :ok,
  :provenance,
  :requireStop,
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
# @!attribute [rw] minAaLength
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] requireStop
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
  :minAaLength,
  :ok,
  :provenance,
  :requireStop,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# FormatSequence entity data model.
#
# @!attribute [rw] caseMode
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
# @!attribute [rw] stripNonLetters
#   @return [Boolean, nil]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] width
#   @return [Integer, nil]
FormatSequence = Struct.new(
  :caseMode,
  :convert,
  :gate,
  :ok,
  :provenance,
  :result,
  :reverse,
  :sequence,
  :stripNonLetters,
  :tool,
  :width,
  keyword_init: true
)

# Request payload for FormatSequence#create.
#
# @!attribute [rw] caseMode
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
# @!attribute [rw] stripNonLetters
#   @return [Boolean, nil]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] width
#   @return [Integer, nil]
FormatSequenceCreateData = Struct.new(
  :caseMode,
  :convert,
  :gate,
  :ok,
  :provenance,
  :result,
  :reverse,
  :sequence,
  :stripNonLetters,
  :tool,
  :width,
  keyword_init: true
)

# FunctionalEnrichment entity data model.
#
# @!attribute [rw] background
#   @return [Array, nil]
#
# @!attribute [rw] collections
#   @return [Array, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] genes
#   @return [Array]
#
# @!attribute [rw] maxTermSize
#   @return [Integer, nil]
#
# @!attribute [rw] minTermSize
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
  :collections,
  :gate,
  :genes,
  :maxTermSize,
  :minTermSize,
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
# @!attribute [rw] collections
#   @return [Array, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] genes
#   @return [Array]
#
# @!attribute [rw] maxTermSize
#   @return [Integer, nil]
#
# @!attribute [rw] minTermSize
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
  :collections,
  :gate,
  :genes,
  :maxTermSize,
  :minTermSize,
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
# @!attribute [rw] compareToNamedSet
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
# @!attribute [rw] overhangs
#   @return [Array]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] riskThreshold
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
GoldenGateFidelity = Struct.new(
  :compareToNamedSet,
  :dataset,
  :gate,
  :ok,
  :overhangs,
  :provenance,
  :result,
  :riskThreshold,
  :tool,
  keyword_init: true
)

# Request payload for GoldenGateFidelity#create.
#
# @!attribute [rw] compareToNamedSet
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
# @!attribute [rw] overhangs
#   @return [Array]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] riskThreshold
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
GoldenGateFidelityCreateData = Struct.new(
  :compareToNamedSet,
  :dataset,
  :gate,
  :ok,
  :overhangs,
  :provenance,
  :result,
  :riskThreshold,
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
# @!attribute [rw] jobId
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
  :jobId,
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
# @!attribute [rw] jobId
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
  :jobId,
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
# @!attribute [rw] taxId
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
  :taxId,
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
# @!attribute [rw] taxId
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
  :taxId,
  :to,
  :tool,
  keyword_init: true
)

# InSilicoPcr entity data model.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] forwardPrimer
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxMismatches
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
# @!attribute [rw] reversePrimer
#   @return [String]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
InSilicoPcr = Struct.new(
  :circular,
  :forwardPrimer,
  :gate,
  :maxMismatches,
  :ok,
  :provenance,
  :result,
  :reversePrimer,
  :template,
  :tool,
  keyword_init: true
)

# Request payload for InSilicoPcr#create.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] forwardPrimer
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxMismatches
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
# @!attribute [rw] reversePrimer
#   @return [String]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
InSilicoPcrCreateData = Struct.new(
  :circular,
  :forwardPrimer,
  :gate,
  :maxMismatches,
  :ok,
  :provenance,
  :result,
  :reversePrimer,
  :template,
  :tool,
  keyword_init: true
)

# KaspPrimerDesign entity data model.
#
# @!attribute [rw] addSecondaryMismatch
#   @return [Boolean, nil]
#
# @!attribute [rw] alleleA
#   @return [String]
#
# @!attribute [rw] alleleB
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxAmplicon
#   @return [Integer, nil]
#
# @!attribute [rw] minAmplicon
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
# @!attribute [rw] snpPosition
#   @return [Integer]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] targetCoreTm
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
KaspPrimerDesign = Struct.new(
  :addSecondaryMismatch,
  :alleleA,
  :alleleB,
  :gate,
  :maxAmplicon,
  :minAmplicon,
  :ok,
  :provenance,
  :result,
  :snpPosition,
  :target,
  :targetCoreTm,
  :tool,
  keyword_init: true
)

# Request payload for KaspPrimerDesign#create.
#
# @!attribute [rw] addSecondaryMismatch
#   @return [Boolean, nil]
#
# @!attribute [rw] alleleA
#   @return [String]
#
# @!attribute [rw] alleleB
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxAmplicon
#   @return [Integer, nil]
#
# @!attribute [rw] minAmplicon
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
# @!attribute [rw] snpPosition
#   @return [Integer]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] targetCoreTm
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
KaspPrimerDesignCreateData = Struct.new(
  :addSecondaryMismatch,
  :alleleA,
  :alleleB,
  :gate,
  :maxAmplicon,
  :minAmplicon,
  :ok,
  :provenance,
  :result,
  :snpPosition,
  :target,
  :targetCoreTm,
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
# @!attribute [rw] dntpMM
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mgMM
#   @return [Float, nil]
#
# @!attribute [rw] naMM
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligoNM
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
# @!attribute [rw] targetTm
#   @return [Float, nil]
#
# @!attribute [rw] tmTolerance
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
MeltingTemperature = Struct.new(
  :dntpMM,
  :gate,
  :mgMM,
  :naMM,
  :ok,
  :oligoNM,
  :provenance,
  :result,
  :sequence,
  :targetTm,
  :tmTolerance,
  :tool,
  keyword_init: true
)

# Request payload for MeltingTemperature#create.
#
# @!attribute [rw] dntpMM
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mgMM
#   @return [Float, nil]
#
# @!attribute [rw] naMM
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligoNM
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
# @!attribute [rw] targetTm
#   @return [Float, nil]
#
# @!attribute [rw] tmTolerance
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
MeltingTemperatureCreateData = Struct.new(
  :dntpMM,
  :gate,
  :mgMM,
  :naMM,
  :ok,
  :oligoNM,
  :provenance,
  :result,
  :sequence,
  :targetTm,
  :tmTolerance,
  :tool,
  keyword_init: true
)

# MotifFinder entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxMismatches
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
# @!attribute [rw] searchReverseStrand
#   @return [Boolean, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
MotifFinder = Struct.new(
  :gate,
  :maxMismatches,
  :motif,
  :ok,
  :provenance,
  :result,
  :searchReverseStrand,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for MotifFinder#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxMismatches
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
# @!attribute [rw] searchReverseStrand
#   @return [Boolean, nil]
#
# @!attribute [rw] sequence
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
MotifFinderCreateData = Struct.new(
  :gate,
  :maxMismatches,
  :motif,
  :ok,
  :provenance,
  :result,
  :searchReverseStrand,
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
# @!attribute [rw] dntpMM
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mgMM
#   @return [Float, nil]
#
# @!attribute [rw] naMM
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligoNM
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
  :dntpMM,
  :gate,
  :mgMM,
  :naMM,
  :ok,
  :oligoNM,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for OligoAnalysi#create.
#
# @!attribute [rw] dntpMM
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mgMM
#   @return [Float, nil]
#
# @!attribute [rw] naMM
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligoNM
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
  :dntpMM,
  :gate,
  :mgMM,
  :naMM,
  :ok,
  :oligoNM,
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
# @!attribute [rw] sourceSpecies
#   @return [String, nil]
#
# @!attribute [rw] symbols
#   @return [Array]
#
# @!attribute [rw] targetSpecies
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
  :sourceSpecies,
  :symbols,
  :targetSpecies,
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
# @!attribute [rw] sourceSpecies
#   @return [String, nil]
#
# @!attribute [rw] symbols
#   @return [Array]
#
# @!attribute [rw] targetSpecies
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
  :sourceSpecies,
  :symbols,
  :targetSpecies,
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
# @!attribute [rw] seqA
#   @return [String]
#
# @!attribute [rw] seqB
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
  :seqA,
  :seqB,
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
# @!attribute [rw] seqA
#   @return [String]
#
# @!attribute [rw] seqB
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
  :seqA,
  :seqB,
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
# @!attribute [rw] fileBase64
#   @return [String]
#
# @!attribute [rw] fileName
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
  :fileBase64,
  :fileName,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for ParseSangerTrace#create.
#
# @!attribute [rw] fileBase64
#   @return [String]
#
# @!attribute [rw] fileName
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
  :fileBase64,
  :fileName,
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
# @!attribute [rw] topN
#   @return [Integer, nil]
PlasmidFullReport = Struct.new(
  :circular,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :topN,
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
# @!attribute [rw] topN
#   @return [Integer, nil]
PlasmidFullReportCreateData = Struct.new(
  :circular,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :topN,
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
# @!attribute [rw] topN
#   @return [Integer, nil]
PlasmidIdentify = Struct.new(
  :circular,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :topN,
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
# @!attribute [rw] topN
#   @return [Integer, nil]
PlasmidIdentifyCreateData = Struct.new(
  :circular,
  :gate,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  :topN,
  keyword_init: true
)

# PrimeEditingDesign entity data model.
#
# @!attribute [rw] editEnd
#   @return [Integer]
#
# @!attribute [rw] editStart
#   @return [Integer]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insertedSeq
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] pbsLength
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] rttHomology
#   @return [Integer, nil]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PrimeEditingDesign = Struct.new(
  :editEnd,
  :editStart,
  :frameStart,
  :gate,
  :insertedSeq,
  :ok,
  :pbsLength,
  :provenance,
  :result,
  :rttHomology,
  :target,
  :tool,
  keyword_init: true
)

# Request payload for PrimeEditingDesign#create.
#
# @!attribute [rw] editEnd
#   @return [Integer]
#
# @!attribute [rw] editStart
#   @return [Integer]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insertedSeq
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] pbsLength
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] rttHomology
#   @return [Integer, nil]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PrimeEditingDesignCreateData = Struct.new(
  :editEnd,
  :editStart,
  :frameStart,
  :gate,
  :insertedSeq,
  :ok,
  :pbsLength,
  :provenance,
  :result,
  :rttHomology,
  :target,
  :tool,
  keyword_init: true
)

# PrimeEditingTwinDesign entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] newSequence
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlapLength
#   @return [Integer, nil]
#
# @!attribute [rw] pbsLength
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] replaceEnd
#   @return [Integer]
#
# @!attribute [rw] replaceStart
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
  :newSequence,
  :ok,
  :overlapLength,
  :pbsLength,
  :provenance,
  :replaceEnd,
  :replaceStart,
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
# @!attribute [rw] newSequence
#   @return [String]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlapLength
#   @return [Integer, nil]
#
# @!attribute [rw] pbsLength
#   @return [Integer, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] replaceEnd
#   @return [Integer]
#
# @!attribute [rw] replaceStart
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
  :newSequence,
  :ok,
  :overlapLength,
  :pbsLength,
  :provenance,
  :replaceEnd,
  :replaceStart,
  :result,
  :target,
  :tool,
  keyword_init: true
)

# PrimerDesign entity data model.
#
# @!attribute [rw] ampliconMax
#   @return [Integer, nil]
#
# @!attribute [rw] ampliconMin
#   @return [Integer, nil]
#
# @!attribute [rw] dntpMM
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gcMax
#   @return [Float, nil]
#
# @!attribute [rw] gcMin
#   @return [Float, nil]
#
# @!attribute [rw] lenMax
#   @return [Integer, nil]
#
# @!attribute [rw] lenMin
#   @return [Integer, nil]
#
# @!attribute [rw] lenOpt
#   @return [Integer, nil]
#
# @!attribute [rw] maxReturn
#   @return [Integer, nil]
#
# @!attribute [rw] mgMM
#   @return [Float, nil]
#
# @!attribute [rw] naMM
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligoNM
#   @return [Float, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] targetEnd
#   @return [Integer, nil]
#
# @!attribute [rw] targetStart
#   @return [Integer, nil]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tmMax
#   @return [Float, nil]
#
# @!attribute [rw] tmMaxDiff
#   @return [Float, nil]
#
# @!attribute [rw] tmMin
#   @return [Float, nil]
#
# @!attribute [rw] tmOpt
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
PrimerDesign = Struct.new(
  :ampliconMax,
  :ampliconMin,
  :dntpMM,
  :gate,
  :gcMax,
  :gcMin,
  :lenMax,
  :lenMin,
  :lenOpt,
  :maxReturn,
  :mgMM,
  :naMM,
  :ok,
  :oligoNM,
  :provenance,
  :result,
  :targetEnd,
  :targetStart,
  :template,
  :tmMax,
  :tmMaxDiff,
  :tmMin,
  :tmOpt,
  :tool,
  keyword_init: true
)

# Request payload for PrimerDesign#create.
#
# @!attribute [rw] ampliconMax
#   @return [Integer, nil]
#
# @!attribute [rw] ampliconMin
#   @return [Integer, nil]
#
# @!attribute [rw] dntpMM
#   @return [Float, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] gcMax
#   @return [Float, nil]
#
# @!attribute [rw] gcMin
#   @return [Float, nil]
#
# @!attribute [rw] lenMax
#   @return [Integer, nil]
#
# @!attribute [rw] lenMin
#   @return [Integer, nil]
#
# @!attribute [rw] lenOpt
#   @return [Integer, nil]
#
# @!attribute [rw] maxReturn
#   @return [Integer, nil]
#
# @!attribute [rw] mgMM
#   @return [Float, nil]
#
# @!attribute [rw] naMM
#   @return [Float, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligoNM
#   @return [Float, nil]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] result
#   @return [Hash]
#
# @!attribute [rw] targetEnd
#   @return [Integer, nil]
#
# @!attribute [rw] targetStart
#   @return [Integer, nil]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tmMax
#   @return [Float, nil]
#
# @!attribute [rw] tmMaxDiff
#   @return [Float, nil]
#
# @!attribute [rw] tmMin
#   @return [Float, nil]
#
# @!attribute [rw] tmOpt
#   @return [Float, nil]
#
# @!attribute [rw] tool
#   @return [String]
PrimerDesignCreateData = Struct.new(
  :ampliconMax,
  :ampliconMin,
  :dntpMM,
  :gate,
  :gcMax,
  :gcMin,
  :lenMax,
  :lenMin,
  :lenOpt,
  :maxReturn,
  :mgMM,
  :naMM,
  :ok,
  :oligoNM,
  :provenance,
  :result,
  :targetEnd,
  :targetStart,
  :template,
  :tmMax,
  :tmMaxDiff,
  :tmMin,
  :tmOpt,
  :tool,
  keyword_init: true
)

# PrimerSpecificity entity data model.
#
# @!attribute [rw] forwardPrimer
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxMismatches
#   @return [Integer, nil]
#
# @!attribute [rw] maxProductLength
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
# @!attribute [rw] reversePrimer
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PrimerSpecificity = Struct.new(
  :forwardPrimer,
  :gate,
  :maxMismatches,
  :maxProductLength,
  :ok,
  :provenance,
  :result,
  :reversePrimer,
  :tool,
  keyword_init: true
)

# Request payload for PrimerSpecificity#create.
#
# @!attribute [rw] forwardPrimer
#   @return [String]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxMismatches
#   @return [Integer, nil]
#
# @!attribute [rw] maxProductLength
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
# @!attribute [rw] reversePrimer
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
PrimerSpecificityCreateData = Struct.new(
  :forwardPrimer,
  :gate,
  :maxMismatches,
  :maxProductLength,
  :ok,
  :provenance,
  :result,
  :reversePrimer,
  :tool,
  keyword_init: true
)

# ProteaseDigestion entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxMass
#   @return [Float, nil]
#
# @!attribute [rw] maxPeptides
#   @return [Integer, nil]
#
# @!attribute [rw] minMass
#   @return [Float, nil]
#
# @!attribute [rw] missedCleavages
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
  :maxMass,
  :maxPeptides,
  :minMass,
  :missedCleavages,
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
# @!attribute [rw] maxMass
#   @return [Float, nil]
#
# @!attribute [rw] maxPeptides
#   @return [Integer, nil]
#
# @!attribute [rw] minMass
#   @return [Float, nil]
#
# @!attribute [rw] missedCleavages
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
  :maxMass,
  :maxPeptides,
  :minMass,
  :missedCleavages,
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
# @!attribute [rw] jobId
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
  :jobId,
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
# @!attribute [rw] jobId
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
  :jobId,
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
# @!attribute [rw] goterms
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
  :goterms,
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
# @!attribute [rw] goterms
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
  :goterms,
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
# @!attribute [rw] chargeStep
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
  :chargeStep,
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
# @!attribute [rw] chargeStep
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
  :chargeStep,
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
# @!attribute [rw] gcContent
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
  :gcContent,
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
# @!attribute [rw] gcContent
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
  :gcContent,
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
# @!attribute [rw] enzymes
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
  :enzymes,
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
# @!attribute [rw] enzymes
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
  :enzymes,
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
# @!attribute [rw] fileBase64
#   @return [String, nil]
#
# @!attribute [rw] fileName
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] minCoverage
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
  :fileBase64,
  :fileName,
  :gate,
  :minCoverage,
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
# @!attribute [rw] fileBase64
#   @return [String, nil]
#
# @!attribute [rw] fileName
#   @return [String, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] minCoverage
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
  :fileBase64,
  :fileName,
  :gate,
  :minCoverage,
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
# @!attribute [rw] args
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
  :args,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for SavePermalink#create.
#
# @!attribute [rw] args
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
  :args,
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
# @!attribute [rw] qualityOffset
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
  :qualityOffset,
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
# @!attribute [rw] qualityOffset
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
  :qualityOffset,
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
# @!attribute [rw] endPrimerLength
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxOrfs
#   @return [Integer, nil]
#
# @!attribute [rw] minOrfAa
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
  :endPrimerLength,
  :gate,
  :maxOrfs,
  :minOrfAa,
  :ok,
  :provenance,
  :result,
  :sequence,
  :tool,
  keyword_init: true
)

# Request payload for SequenceReport#create.
#
# @!attribute [rw] endPrimerLength
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] maxOrfs
#   @return [Integer, nil]
#
# @!attribute [rw] minOrfAa
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
  :endPrimerLength,
  :gate,
  :maxOrfs,
  :minOrfAa,
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
# @!attribute [rw] maxResults
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
  :maxResults,
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
# @!attribute [rw] maxResults
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
  :maxResults,
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
# @!attribute [rw] minSupportingReads
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] reads
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
  :minSupportingReads,
  :ok,
  :provenance,
  :reads,
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
# @!attribute [rw] minSupportingReads
#   @return [Integer, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] provenance
#   @return [Hash]
#
# @!attribute [rw] reads
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
  :minSupportingReads,
  :ok,
  :provenance,
  :reads,
  :reference,
  :result,
  :tool,
  keyword_init: true
)

# SessionCreate entity data model.
#
# @!attribute [rw] entries
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
  :entries,
  :gate,
  :ok,
  :provenance,
  :result,
  :tool,
  keyword_init: true
)

# Request payload for SessionCreate#create.
#
# @!attribute [rw] entries
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
  :entries,
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
# @!attribute [rw] names
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
# @!attribute [rw] sessionId
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SessionGet = Struct.new(
  :gate,
  :names,
  :ok,
  :provenance,
  :result,
  :sessionId,
  :tool,
  keyword_init: true
)

# Request payload for SessionGet#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] names
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
# @!attribute [rw] sessionId
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SessionGetCreateData = Struct.new(
  :gate,
  :names,
  :ok,
  :provenance,
  :result,
  :sessionId,
  :tool,
  keyword_init: true
)

# SessionRun entity data model.
#
# @!attribute [rw] args
#   @return [Hash, nil]
#
# @!attribute [rw] fromSession
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
# @!attribute [rw] sessionId
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] writeBack
#   @return [Hash, nil]
SessionRun = Struct.new(
  :args,
  :fromSession,
  :gate,
  :ok,
  :provenance,
  :result,
  :sessionId,
  :tool,
  :writeBack,
  keyword_init: true
)

# Request payload for SessionRun#create.
#
# @!attribute [rw] args
#   @return [Hash, nil]
#
# @!attribute [rw] fromSession
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
# @!attribute [rw] sessionId
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
#
# @!attribute [rw] writeBack
#   @return [Hash, nil]
SessionRunCreateData = Struct.new(
  :args,
  :fromSession,
  :gate,
  :ok,
  :provenance,
  :result,
  :sessionId,
  :tool,
  :writeBack,
  keyword_init: true
)

# SessionSet entity data model.
#
# @!attribute [rw] entries
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
# @!attribute [rw] sessionId
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SessionSet = Struct.new(
  :entries,
  :gate,
  :ok,
  :provenance,
  :result,
  :sessionId,
  :tool,
  keyword_init: true
)

# Request payload for SessionSet#create.
#
# @!attribute [rw] entries
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
# @!attribute [rw] sessionId
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SessionSetCreateData = Struct.new(
  :entries,
  :gate,
  :ok,
  :provenance,
  :result,
  :sessionId,
  :tool,
  keyword_init: true
)

# SirnaDesign entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] minReynolds
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
# @!attribute [rw] shRnaLoop
#   @return [String, nil]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SirnaDesign = Struct.new(
  :gate,
  :minReynolds,
  :ok,
  :provenance,
  :result,
  :shRnaLoop,
  :target,
  :tool,
  keyword_init: true
)

# Request payload for SirnaDesign#create.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] minReynolds
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
# @!attribute [rw] shRnaLoop
#   @return [String, nil]
#
# @!attribute [rw] target
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SirnaDesignCreateData = Struct.new(
  :gate,
  :minReynolds,
  :ok,
  :provenance,
  :result,
  :shRnaLoop,
  :target,
  :tool,
  keyword_init: true
)

# SiteDirectedMutagenesi entity data model.
#
# @!attribute [rw] armTmTarget
#   @return [Float, nil]
#
# @!attribute [rw] dntpMM
#   @return [Float, nil]
#
# @!attribute [rw] editKind
#   @return [String, nil]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mgMM
#   @return [Float, nil]
#
# @!attribute [rw] naMM
#   @return [Float, nil]
#
# @!attribute [rw] newBase
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligoNM
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
# @!attribute [rw] targetAa
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SiteDirectedMutagenesi = Struct.new(
  :armTmTarget,
  :dntpMM,
  :editKind,
  :frameStart,
  :gate,
  :mgMM,
  :naMM,
  :newBase,
  :ok,
  :oligoNM,
  :organism,
  :position,
  :provenance,
  :residue,
  :result,
  :style,
  :targetAa,
  :template,
  :tool,
  keyword_init: true
)

# Request payload for SiteDirectedMutagenesi#create.
#
# @!attribute [rw] armTmTarget
#   @return [Float, nil]
#
# @!attribute [rw] dntpMM
#   @return [Float, nil]
#
# @!attribute [rw] editKind
#   @return [String, nil]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] mgMM
#   @return [Float, nil]
#
# @!attribute [rw] naMM
#   @return [Float, nil]
#
# @!attribute [rw] newBase
#   @return [String, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] oligoNM
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
# @!attribute [rw] targetAa
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [String]
#
# @!attribute [rw] tool
#   @return [String]
SiteDirectedMutagenesiCreateData = Struct.new(
  :armTmTarget,
  :dntpMM,
  :editKind,
  :frameStart,
  :gate,
  :mgMM,
  :naMM,
  :newBase,
  :ok,
  :oligoNM,
  :organism,
  :position,
  :provenance,
  :residue,
  :result,
  :style,
  :targetAa,
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
# @!attribute [rw] toStop
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
  :toStop,
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
# @!attribute [rw] toStop
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
  :toStop,
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
# @!attribute [rw] frameStart
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
  :frameStart,
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
# @!attribute [rw] frameStart
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
  :frameStart,
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
# @!attribute [rw] armTmTarget
#   @return [Float, nil]
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] claimedConstruct
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
# @!attribute [rw] fragmentPcrs
#   @return [Array, nil]
#
# @!attribute [rw] fragments
#   @return [Array, nil]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insert
#   @return [String, nil]
#
# @!attribute [rw] insertPcr
#   @return [Hash, nil]
#
# @!attribute [rw] method
#   @return [String]
#
# @!attribute [rw] names
#   @return [Array, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlapLen
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
# @!attribute [rw] vectorPcr
#   @return [Hash, nil]
VerifyAssembly = Struct.new(
  :armTmTarget,
  :circular,
  :claimedConstruct,
  :coding,
  :enzyme,
  :enzyme3,
  :enzyme5,
  :fragmentPcrs,
  :fragments,
  :frameStart,
  :gate,
  :insert,
  :insertPcr,
  :method,
  :names,
  :ok,
  :overlapLen,
  :provenance,
  :result,
  :tool,
  :vector,
  :vectorPcr,
  keyword_init: true
)

# Request payload for VerifyAssembly#create.
#
# @!attribute [rw] armTmTarget
#   @return [Float, nil]
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] claimedConstruct
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
# @!attribute [rw] fragmentPcrs
#   @return [Array, nil]
#
# @!attribute [rw] fragments
#   @return [Array, nil]
#
# @!attribute [rw] frameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insert
#   @return [String, nil]
#
# @!attribute [rw] insertPcr
#   @return [Hash, nil]
#
# @!attribute [rw] method
#   @return [String]
#
# @!attribute [rw] names
#   @return [Array, nil]
#
# @!attribute [rw] ok
#   @return [Object]
#
# @!attribute [rw] overlapLen
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
# @!attribute [rw] vectorPcr
#   @return [Hash, nil]
VerifyAssemblyCreateData = Struct.new(
  :armTmTarget,
  :circular,
  :claimedConstruct,
  :coding,
  :enzyme,
  :enzyme3,
  :enzyme5,
  :fragmentPcrs,
  :fragments,
  :frameStart,
  :gate,
  :insert,
  :insertPcr,
  :method,
  :names,
  :ok,
  :overlapLen,
  :provenance,
  :result,
  :tool,
  :vector,
  :vectorPcr,
  keyword_init: true
)

# VerifyConstruct entity data model.
#
# @!attribute [rw] claimedConstruct
#   @return [String]
#
# @!attribute [rw] expectedFrameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insertForwardPrimer
#   @return [String]
#
# @!attribute [rw] insertReversePrimer
#   @return [String]
#
# @!attribute [rw] insertTemplate
#   @return [String]
#
# @!attribute [rw] maxPrimerMismatches
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
# @!attribute [rw] templateCircular
#   @return [Boolean, nil]
#
# @!attribute [rw] tool
#   @return [String]
VerifyConstruct = Struct.new(
  :claimedConstruct,
  :expectedFrameStart,
  :gate,
  :insertForwardPrimer,
  :insertReversePrimer,
  :insertTemplate,
  :maxPrimerMismatches,
  :ok,
  :provenance,
  :result,
  :templateCircular,
  :tool,
  keyword_init: true
)

# Request payload for VerifyConstruct#create.
#
# @!attribute [rw] claimedConstruct
#   @return [String]
#
# @!attribute [rw] expectedFrameStart
#   @return [Integer, nil]
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] insertForwardPrimer
#   @return [String]
#
# @!attribute [rw] insertReversePrimer
#   @return [String]
#
# @!attribute [rw] insertTemplate
#   @return [String]
#
# @!attribute [rw] maxPrimerMismatches
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
# @!attribute [rw] templateCircular
#   @return [Boolean, nil]
#
# @!attribute [rw] tool
#   @return [String]
VerifyConstructCreateData = Struct.new(
  :claimedConstruct,
  :expectedFrameStart,
  :gate,
  :insertForwardPrimer,
  :insertReversePrimer,
  :insertTemplate,
  :maxPrimerMismatches,
  :ok,
  :provenance,
  :result,
  :templateCircular,
  :tool,
  keyword_init: true
)

# VirtualGel entity data model.
#
# @!attribute [rw] circular
#   @return [Boolean, nil]
#
# @!attribute [rw] enzymes
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
  :enzymes,
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
# @!attribute [rw] enzymes
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
  :enzymes,
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
# @!attribute [rw] rows
#   @return [Array]
#
# @!attribute [rw] tool
#   @return [String]
VolcanoPlotData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :rows,
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
# @!attribute [rw] rows
#   @return [Array]
#
# @!attribute [rw] tool
#   @return [String]
VolcanoPlotDataCreateData = Struct.new(
  :gate,
  :ok,
  :provenance,
  :result,
  :rows,
  :tool,
  keyword_init: true
)

# WebSearch entity data model.
#
# @!attribute [rw] gate
#   @return [Object, nil]
#
# @!attribute [rw] max_results
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
  :max_results,
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
# @!attribute [rw] max_results
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
  :max_results,
  :ok,
  :provenance,
  :query,
  :result,
  :tool,
  keyword_init: true
)

