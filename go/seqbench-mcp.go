package voxgigseqbenchmcpsdk

import (
	"github.com/voxgig-sdk/seqbench-mcp-sdk/go/core"
	"github.com/voxgig-sdk/seqbench-mcp-sdk/go/entity"
	"github.com/voxgig-sdk/seqbench-mcp-sdk/go/feature"
	_ "github.com/voxgig-sdk/seqbench-mcp-sdk/go/utility"
)

// Type aliases preserve external API.
type SeqbenchMcpSDK = core.SeqbenchMcpSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type SeqbenchMcpEntity = core.SeqbenchMcpEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type SeqbenchMcpError = core.SeqbenchMcpError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewAlphafoldLookupEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewAlphafoldLookupEntity(client, entopts)
	}
	core.NewAsoDesignEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewAsoDesignEntity(client, entopts)
	}
	core.NewBaseEditingDesignEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewBaseEditingDesignEntity(client, entopts)
	}
	core.NewBatchEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewBatchEntity(client, entopts)
	}
	core.NewBatchWorkflowEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewBatchWorkflowEntity(client, entopts)
	}
	core.NewCharacterizeSequenceEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewCharacterizeSequenceEntity(client, entopts)
	}
	core.NewCloningSimulateEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewCloningSimulateEntity(client, entopts)
	}
	core.NewCodonAdaptationIndexEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewCodonAdaptationIndexEntity(client, entopts)
	}
	core.NewCodonOptimizeEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewCodonOptimizeEntity(client, entopts)
	}
	core.NewConstructAutofixEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewConstructAutofixEntity(client, entopts)
	}
	core.NewConstructQcEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewConstructQcEntity(client, entopts)
	}
	core.NewCrisprGrnaDesignEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewCrisprGrnaDesignEntity(client, entopts)
	}
	core.NewCrisprHdrDonorEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewCrisprHdrDonorEntity(client, entopts)
	}
	core.NewCrisprOfftargetCheckEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewCrisprOfftargetCheckEntity(client, entopts)
	}
	core.NewCrossDimerEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewCrossDimerEntity(client, entopts)
	}
	core.NewDnaMolarityEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewDnaMolarityEntity(client, entopts)
	}
	core.NewDoubleDigestEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewDoubleDigestEntity(client, entopts)
	}
	core.NewExportEchoPicklistEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewExportEchoPicklistEntity(client, entopts)
	}
	core.NewExportOpentronsProtocolEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewExportOpentronsProtocolEntity(client, entopts)
	}
	core.NewExportPlateLayoutEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewExportPlateLayoutEntity(client, entopts)
	}
	core.NewExpressionHeatmapClusterEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewExpressionHeatmapClusterEntity(client, entopts)
	}
	core.NewFastqQcReportEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewFastqQcReportEntity(client, entopts)
	}
	core.NewFastqTrimEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewFastqTrimEntity(client, entopts)
	}
	core.NewFindOrfEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewFindOrfEntity(client, entopts)
	}
	core.NewFormatSequenceEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewFormatSequenceEntity(client, entopts)
	}
	core.NewFunctionalEnrichmentEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewFunctionalEnrichmentEntity(client, entopts)
	}
	core.NewGcContentEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewGcContentEntity(client, entopts)
	}
	core.NewGeneDossierEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewGeneDossierEntity(client, entopts)
	}
	core.NewGeneExpressionEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewGeneExpressionEntity(client, entopts)
	}
	core.NewGeneModelEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewGeneModelEntity(client, entopts)
	}
	core.NewGoldenGateFidelityEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewGoldenGateFidelityEntity(client, entopts)
	}
	core.NewHgvsConvertEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewHgvsConvertEntity(client, entopts)
	}
	core.NewIdMapPollEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewIdMapPollEntity(client, entopts)
	}
	core.NewIdMapSubmitEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewIdMapSubmitEntity(client, entopts)
	}
	core.NewInSilicoPcrEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewInSilicoPcrEntity(client, entopts)
	}
	core.NewKaspPrimerDesignEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewKaspPrimerDesignEntity(client, entopts)
	}
	core.NewListToolEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewListToolEntity(client, entopts)
	}
	core.NewMeltingTemperatureEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewMeltingTemperatureEntity(client, entopts)
	}
	core.NewMotifFinderEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewMotifFinderEntity(client, entopts)
	}
	core.NewMultipleSequenceAlignmentEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewMultipleSequenceAlignmentEntity(client, entopts)
	}
	core.NewOligoAnalysiEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewOligoAnalysiEntity(client, entopts)
	}
	core.NewOrthologMapEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewOrthologMapEntity(client, entopts)
	}
	core.NewPairwiseAlignmentEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewPairwiseAlignmentEntity(client, entopts)
	}
	core.NewParseGenbankEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewParseGenbankEntity(client, entopts)
	}
	core.NewParseSangerTraceEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewParseSangerTraceEntity(client, entopts)
	}
	core.NewPlasmidAnnotateEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewPlasmidAnnotateEntity(client, entopts)
	}
	core.NewPlasmidDeepAnnotateEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewPlasmidDeepAnnotateEntity(client, entopts)
	}
	core.NewPlasmidFullReportEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewPlasmidFullReportEntity(client, entopts)
	}
	core.NewPlasmidIdentifyEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewPlasmidIdentifyEntity(client, entopts)
	}
	core.NewPrimeEditingDesignEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewPrimeEditingDesignEntity(client, entopts)
	}
	core.NewPrimeEditingTwinDesignEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewPrimeEditingTwinDesignEntity(client, entopts)
	}
	core.NewPrimerDesignEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewPrimerDesignEntity(client, entopts)
	}
	core.NewPrimerSpecificityEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewPrimerSpecificityEntity(client, entopts)
	}
	core.NewProteaseDigestionEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewProteaseDigestionEntity(client, entopts)
	}
	core.NewProteinAnnotatePollEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewProteinAnnotatePollEntity(client, entopts)
	}
	core.NewProteinAnnotateSubmitEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewProteinAnnotateSubmitEntity(client, entopts)
	}
	core.NewProteinHydrophobicityEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewProteinHydrophobicityEntity(client, entopts)
	}
	core.NewProteinPropertyEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewProteinPropertyEntity(client, entopts)
	}
	core.NewRandomSequenceEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewRandomSequenceEntity(client, entopts)
	}
	core.NewRestrictionSiteEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewRestrictionSiteEntity(client, entopts)
	}
	core.NewReverseComplementEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewReverseComplementEntity(client, entopts)
	}
	core.NewReverseTranslateEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewReverseTranslateEntity(client, entopts)
	}
	core.NewRnaFoldEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewRnaFoldEntity(client, entopts)
	}
	core.NewSangerVsReferenceEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSangerVsReferenceEntity(client, entopts)
	}
	core.NewSavePermalinkEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSavePermalinkEntity(client, entopts)
	}
	core.NewSeqfileStatEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSeqfileStatEntity(client, entopts)
	}
	core.NewSequenceFetchEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSequenceFetchEntity(client, entopts)
	}
	core.NewSequenceFormatConvertEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSequenceFormatConvertEntity(client, entopts)
	}
	core.NewSequenceReportEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSequenceReportEntity(client, entopts)
	}
	core.NewSequenceSearchEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSequenceSearchEntity(client, entopts)
	}
	core.NewSequencingReadbackVerifyEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSequencingReadbackVerifyEntity(client, entopts)
	}
	core.NewSessionCreateEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSessionCreateEntity(client, entopts)
	}
	core.NewSessionGetEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSessionGetEntity(client, entopts)
	}
	core.NewSessionRunEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSessionRunEntity(client, entopts)
	}
	core.NewSessionSetEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSessionSetEntity(client, entopts)
	}
	core.NewSirnaDesignEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSirnaDesignEntity(client, entopts)
	}
	core.NewSiteDirectedMutagenesiEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewSiteDirectedMutagenesiEntity(client, entopts)
	}
	core.NewTranslateEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewTranslateEntity(client, entopts)
	}
	core.NewVariantAnnotateEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewVariantAnnotateEntity(client, entopts)
	}
	core.NewVariantComparatorEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewVariantComparatorEntity(client, entopts)
	}
	core.NewVerifyAssemblyEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewVerifyAssemblyEntity(client, entopts)
	}
	core.NewVerifyConstructEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewVerifyConstructEntity(client, entopts)
	}
	core.NewVirtualGelEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewVirtualGelEntity(client, entopts)
	}
	core.NewVolcanoPlotDataEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewVolcanoPlotDataEntity(client, entopts)
	}
	core.NewWebSearchEntityFunc = func(client *core.SeqbenchMcpSDK, entopts map[string]any) core.SeqbenchMcpEntity {
		return entity.NewWebSearchEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewSeqbenchMcpSDK = core.NewSeqbenchMcpSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewSeqbenchMcpSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *SeqbenchMcpSDK  { return NewSeqbenchMcpSDK(nil) }
func Test() *SeqbenchMcpSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
