package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewAlphafoldLookupEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewAsoDesignEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewBaseEditingDesignEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewBatchEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewBatchWorkflowEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewCharacterizeSequenceEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewCloningSimulateEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewCodonAdaptationIndexEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewCodonOptimizeEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewConstructAutofixEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewConstructQcEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewCrisprGrnaDesignEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewCrisprHdrDonorEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewCrisprOfftargetCheckEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewCrossDimerEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewDnaMolarityEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewDoubleDigestEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewExportEchoPicklistEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewExportOpentronsProtocolEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewExportPlateLayoutEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewExpressionHeatmapClusterEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewFastqQcReportEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewFastqTrimEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewFindOrfEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewFormatSequenceEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewFunctionalEnrichmentEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewGcContentEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewGeneDossierEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewGeneExpressionEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewGeneModelEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewGoldenGateFidelityEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewHgvsConvertEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewIdMapPollEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewIdMapSubmitEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewInSilicoPcrEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewKaspPrimerDesignEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewListToolEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewMeltingTemperatureEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewMotifFinderEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewMultipleSequenceAlignmentEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewOligoAnalysiEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewOrthologMapEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewPairwiseAlignmentEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewParseGenbankEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewParseSangerTraceEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewPlasmidAnnotateEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewPlasmidDeepAnnotateEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewPlasmidFullReportEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewPlasmidIdentifyEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewPrimeEditingDesignEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewPrimeEditingTwinDesignEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewPrimerDesignEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewPrimerSpecificityEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewProteaseDigestionEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewProteinAnnotatePollEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewProteinAnnotateSubmitEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewProteinHydrophobicityEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewProteinPropertyEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewRandomSequenceEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewRestrictionSiteEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewReverseComplementEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewReverseTranslateEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewRnaFoldEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSangerVsReferenceEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSavePermalinkEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSeqfileStatEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSequenceFetchEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSequenceFormatConvertEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSequenceReportEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSequenceSearchEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSequencingReadbackVerifyEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSessionCreateEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSessionGetEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSessionRunEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSessionSetEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSirnaDesignEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewSiteDirectedMutagenesiEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewTranslateEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewVariantAnnotateEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewVariantComparatorEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewVerifyAssemblyEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewVerifyConstructEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewVirtualGelEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewVolcanoPlotDataEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

var NewWebSearchEntityFunc func(client *SeqbenchMcpSDK, entopts map[string]any) SeqbenchMcpEntity

