// SeqbenchMcp Ts SDK

import { AlphafoldLookupEntity } from './entity/AlphafoldLookupEntity'
import { AsoDesignEntity } from './entity/AsoDesignEntity'
import { BaseEditingDesignEntity } from './entity/BaseEditingDesignEntity'
import { BatchEntity } from './entity/BatchEntity'
import { BatchWorkflowEntity } from './entity/BatchWorkflowEntity'
import { CharacterizeSequenceEntity } from './entity/CharacterizeSequenceEntity'
import { CloningSimulateEntity } from './entity/CloningSimulateEntity'
import { CodonAdaptationIndexEntity } from './entity/CodonAdaptationIndexEntity'
import { CodonOptimizeEntity } from './entity/CodonOptimizeEntity'
import { ConstructAutofixEntity } from './entity/ConstructAutofixEntity'
import { ConstructQcEntity } from './entity/ConstructQcEntity'
import { CrisprGrnaDesignEntity } from './entity/CrisprGrnaDesignEntity'
import { CrisprHdrDonorEntity } from './entity/CrisprHdrDonorEntity'
import { CrisprOfftargetCheckEntity } from './entity/CrisprOfftargetCheckEntity'
import { CrossDimerEntity } from './entity/CrossDimerEntity'
import { DnaMolarityEntity } from './entity/DnaMolarityEntity'
import { DoubleDigestEntity } from './entity/DoubleDigestEntity'
import { ExportEchoPicklistEntity } from './entity/ExportEchoPicklistEntity'
import { ExportOpentronsProtocolEntity } from './entity/ExportOpentronsProtocolEntity'
import { ExportPlateLayoutEntity } from './entity/ExportPlateLayoutEntity'
import { ExpressionHeatmapClusterEntity } from './entity/ExpressionHeatmapClusterEntity'
import { FastqQcReportEntity } from './entity/FastqQcReportEntity'
import { FastqTrimEntity } from './entity/FastqTrimEntity'
import { FindOrfEntity } from './entity/FindOrfEntity'
import { FormatSequenceEntity } from './entity/FormatSequenceEntity'
import { FunctionalEnrichmentEntity } from './entity/FunctionalEnrichmentEntity'
import { GcContentEntity } from './entity/GcContentEntity'
import { GeneDossierEntity } from './entity/GeneDossierEntity'
import { GeneExpressionEntity } from './entity/GeneExpressionEntity'
import { GeneModelEntity } from './entity/GeneModelEntity'
import { GoldenGateFidelityEntity } from './entity/GoldenGateFidelityEntity'
import { HgvsConvertEntity } from './entity/HgvsConvertEntity'
import { IdMapPollEntity } from './entity/IdMapPollEntity'
import { IdMapSubmitEntity } from './entity/IdMapSubmitEntity'
import { InSilicoPcrEntity } from './entity/InSilicoPcrEntity'
import { KaspPrimerDesignEntity } from './entity/KaspPrimerDesignEntity'
import { ListToolEntity } from './entity/ListToolEntity'
import { MeltingTemperatureEntity } from './entity/MeltingTemperatureEntity'
import { MotifFinderEntity } from './entity/MotifFinderEntity'
import { MultipleSequenceAlignmentEntity } from './entity/MultipleSequenceAlignmentEntity'
import { OligoAnalysiEntity } from './entity/OligoAnalysiEntity'
import { OrthologMapEntity } from './entity/OrthologMapEntity'
import { PairwiseAlignmentEntity } from './entity/PairwiseAlignmentEntity'
import { ParseGenbankEntity } from './entity/ParseGenbankEntity'
import { ParseSangerTraceEntity } from './entity/ParseSangerTraceEntity'
import { PlasmidAnnotateEntity } from './entity/PlasmidAnnotateEntity'
import { PlasmidDeepAnnotateEntity } from './entity/PlasmidDeepAnnotateEntity'
import { PlasmidFullReportEntity } from './entity/PlasmidFullReportEntity'
import { PlasmidIdentifyEntity } from './entity/PlasmidIdentifyEntity'
import { PrimeEditingDesignEntity } from './entity/PrimeEditingDesignEntity'
import { PrimeEditingTwinDesignEntity } from './entity/PrimeEditingTwinDesignEntity'
import { PrimerDesignEntity } from './entity/PrimerDesignEntity'
import { PrimerSpecificityEntity } from './entity/PrimerSpecificityEntity'
import { ProteaseDigestionEntity } from './entity/ProteaseDigestionEntity'
import { ProteinAnnotatePollEntity } from './entity/ProteinAnnotatePollEntity'
import { ProteinAnnotateSubmitEntity } from './entity/ProteinAnnotateSubmitEntity'
import { ProteinHydrophobicityEntity } from './entity/ProteinHydrophobicityEntity'
import { ProteinPropertyEntity } from './entity/ProteinPropertyEntity'
import { RandomSequenceEntity } from './entity/RandomSequenceEntity'
import { RestrictionSiteEntity } from './entity/RestrictionSiteEntity'
import { ReverseComplementEntity } from './entity/ReverseComplementEntity'
import { ReverseTranslateEntity } from './entity/ReverseTranslateEntity'
import { RnaFoldEntity } from './entity/RnaFoldEntity'
import { SangerVsReferenceEntity } from './entity/SangerVsReferenceEntity'
import { SavePermalinkEntity } from './entity/SavePermalinkEntity'
import { SeqfileStatEntity } from './entity/SeqfileStatEntity'
import { SequenceFetchEntity } from './entity/SequenceFetchEntity'
import { SequenceFormatConvertEntity } from './entity/SequenceFormatConvertEntity'
import { SequenceReportEntity } from './entity/SequenceReportEntity'
import { SequenceSearchEntity } from './entity/SequenceSearchEntity'
import { SequencingReadbackVerifyEntity } from './entity/SequencingReadbackVerifyEntity'
import { SessionCreateEntity } from './entity/SessionCreateEntity'
import { SessionGetEntity } from './entity/SessionGetEntity'
import { SessionRunEntity } from './entity/SessionRunEntity'
import { SessionSetEntity } from './entity/SessionSetEntity'
import { SirnaDesignEntity } from './entity/SirnaDesignEntity'
import { SiteDirectedMutagenesiEntity } from './entity/SiteDirectedMutagenesiEntity'
import { TranslateEntity } from './entity/TranslateEntity'
import { VariantAnnotateEntity } from './entity/VariantAnnotateEntity'
import { VariantComparatorEntity } from './entity/VariantComparatorEntity'
import { VerifyAssemblyEntity } from './entity/VerifyAssemblyEntity'
import { VerifyConstructEntity } from './entity/VerifyConstructEntity'
import { VirtualGelEntity } from './entity/VirtualGelEntity'
import { VolcanoPlotDataEntity } from './entity/VolcanoPlotDataEntity'
import { WebSearchEntity } from './entity/WebSearchEntity'

export type * from './SeqbenchMcpTypes'


import { inspect } from 'node:util'

import type { Context, Feature } from './types'

import { config } from './Config'
import { SeqbenchMcpEntityBase } from './SeqbenchMcpEntityBase'
import { Utility } from './utility/Utility'


import { BaseFeature } from './feature/base/BaseFeature'


const stdutil = new Utility()


class SeqbenchMcpSDK {
  _mode: string = 'live'
  _options: any
  _utility = new Utility()
  _features: Feature[]
  _rootctx: Context

  constructor(options?: any) {

    this._rootctx = this._utility.makeContext({
      client: this,
      utility: this._utility,
      config,
      options,
      shared: new WeakMap()
    })

    this._options = this._utility.makeOptions(this._rootctx)

    const struct = this._utility.struct
    const getpath = struct.getpath

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    // Add features in the resolved order (makeOptions puts an explicit
    // array order first, else defaults to test-first). Ordering matters:
    // the `test` feature installs the base mock transport and the transport
    // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
    // so `test` must be added before them to sit at the base of the chain.
    const featureorder = getpath(this._options, '__derived__.featureorder') || []
    for (const fname of featureorder) {
      const fopts = this._options.feature[fname] || {}
      if (fopts.active) {
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    }

    if (null != this._options.extend) {
      for (let f of this._options.extend) {
        featureAdd(this._rootctx, f)
      }
    }

    for (let f of this._features) {
      featureInit(this._rootctx, f)
    }

    const featureHook = this._utility.featureHook
    featureHook(this._rootctx, 'PostConstruct')
  }


  options() {
    return this._utility.struct.clone(this._options)
  }


  utility() {
    return this._utility.struct.clone(this._utility)
  }


  async prepare(fetchargs?: any) {
    const utility = this._utility
    const struct = utility.struct
    const clone = struct.clone

    const {
      makeContext,
      makeFetchDef,
      prepareHeaders,
      prepareAuth,
    } = utility

    fetchargs = fetchargs || {}

    let ctx: Context = makeContext({
      opname: 'prepare',
      ctrl: fetchargs.ctrl || {},
    }, this._rootctx)

    const options = this._options

    // Build spec directly from SDK options + user-provided fetch args.
    const spec: any = {
      base: options.base,
      prefix: options.prefix,
      suffix: options.suffix,
      path: fetchargs.path || '',
      method: fetchargs.method || 'GET',
      params: fetchargs.params || {},
      query: fetchargs.query || {},
      headers: prepareHeaders(ctx),
      body: fetchargs.body,
      step: 'start',
    }

    ctx.spec = spec

    // Merge user-provided headers over SDK defaults.
    if (fetchargs.headers) {
      const uheaders = fetchargs.headers
      for (let key in uheaders) {
        spec.headers[key] = uheaders[key]
      }
    }

    // Apply SDK auth (apikey, auth prefix, etc.)
    const authResult = prepareAuth(ctx)
    if (authResult instanceof Error) {
      return authResult
    }

    return makeFetchDef(ctx)
  }


  async direct(fetchargs?: any) {
    const utility = this._utility
    const fetcher = utility.fetcher
    const makeContext = utility.makeContext

    const fetchdef = await this.prepare(fetchargs)
    if (fetchdef instanceof Error) {
      return fetchdef
    }

    let ctx: Context = makeContext({
      opname: 'direct',
      ctrl: (fetchargs || {}).ctrl || {},
    }, this._rootctx)

    try {
      const fetched = await fetcher(ctx, fetchdef.url, fetchdef)

      if (null == fetched) {
        return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') }
      }
      else if (fetched instanceof Error) {
        return { ok: false, err: fetched }
      }

      const status = fetched.status

      // No body responses (204 No Content, 304 Not Modified) and explicit
      // zero content-length must skip JSON parsing — fetched.json() would
      // throw `Unexpected end of JSON input` on an empty body.
      const headers = fetched.headers
      const contentLength = headers && 'function' === typeof headers.get
        ? headers.get('content-length')
        : (headers || {})['content-length']
      const noBody = 204 === status || 304 === status || '0' === String(contentLength)

      let json: any = undefined
      if (!noBody) {
        try {
          json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json
        }
        catch (parseErr) {
          // Body wasn't valid JSON — surface the raw response rather than
          // throwing. data stays undefined; callers can inspect status/headers.
          json = undefined
        }
      }

      return {
        ok: status >= 200 && status < 300,
        status,
        headers: fetched.headers,
        data: json,
      }
    }
    catch (err: any) {
      return { ok: false, err }
    }
  }



  // Entity access: `client.AlphafoldLookup().list()` / `client.AlphafoldLookup().load({ id })`.
  AlphafoldLookup(data?: any) {
    const self = this
    return new AlphafoldLookupEntity(self,data)
  }


  // Entity access: `client.AsoDesign().list()` / `client.AsoDesign().load({ id })`.
  AsoDesign(data?: any) {
    const self = this
    return new AsoDesignEntity(self,data)
  }


  // Entity access: `client.BaseEditingDesign().list()` / `client.BaseEditingDesign().load({ id })`.
  BaseEditingDesign(data?: any) {
    const self = this
    return new BaseEditingDesignEntity(self,data)
  }


  // Entity access: `client.Batch().list()` / `client.Batch().load({ id })`.
  Batch(data?: any) {
    const self = this
    return new BatchEntity(self,data)
  }


  // Entity access: `client.BatchWorkflow().list()` / `client.BatchWorkflow().load({ id })`.
  BatchWorkflow(data?: any) {
    const self = this
    return new BatchWorkflowEntity(self,data)
  }


  // Entity access: `client.CharacterizeSequence().list()` / `client.CharacterizeSequence().load({ id })`.
  CharacterizeSequence(data?: any) {
    const self = this
    return new CharacterizeSequenceEntity(self,data)
  }


  // Entity access: `client.CloningSimulate().list()` / `client.CloningSimulate().load({ id })`.
  CloningSimulate(data?: any) {
    const self = this
    return new CloningSimulateEntity(self,data)
  }


  // Entity access: `client.CodonAdaptationIndex().list()` / `client.CodonAdaptationIndex().load({ id })`.
  CodonAdaptationIndex(data?: any) {
    const self = this
    return new CodonAdaptationIndexEntity(self,data)
  }


  // Entity access: `client.CodonOptimize().list()` / `client.CodonOptimize().load({ id })`.
  CodonOptimize(data?: any) {
    const self = this
    return new CodonOptimizeEntity(self,data)
  }


  // Entity access: `client.ConstructAutofix().list()` / `client.ConstructAutofix().load({ id })`.
  ConstructAutofix(data?: any) {
    const self = this
    return new ConstructAutofixEntity(self,data)
  }


  // Entity access: `client.ConstructQc().list()` / `client.ConstructQc().load({ id })`.
  ConstructQc(data?: any) {
    const self = this
    return new ConstructQcEntity(self,data)
  }


  // Entity access: `client.CrisprGrnaDesign().list()` / `client.CrisprGrnaDesign().load({ id })`.
  CrisprGrnaDesign(data?: any) {
    const self = this
    return new CrisprGrnaDesignEntity(self,data)
  }


  // Entity access: `client.CrisprHdrDonor().list()` / `client.CrisprHdrDonor().load({ id })`.
  CrisprHdrDonor(data?: any) {
    const self = this
    return new CrisprHdrDonorEntity(self,data)
  }


  // Entity access: `client.CrisprOfftargetCheck().list()` / `client.CrisprOfftargetCheck().load({ id })`.
  CrisprOfftargetCheck(data?: any) {
    const self = this
    return new CrisprOfftargetCheckEntity(self,data)
  }


  // Entity access: `client.CrossDimer().list()` / `client.CrossDimer().load({ id })`.
  CrossDimer(data?: any) {
    const self = this
    return new CrossDimerEntity(self,data)
  }


  // Entity access: `client.DnaMolarity().list()` / `client.DnaMolarity().load({ id })`.
  DnaMolarity(data?: any) {
    const self = this
    return new DnaMolarityEntity(self,data)
  }


  // Entity access: `client.DoubleDigest().list()` / `client.DoubleDigest().load({ id })`.
  DoubleDigest(data?: any) {
    const self = this
    return new DoubleDigestEntity(self,data)
  }


  // Entity access: `client.ExportEchoPicklist().list()` / `client.ExportEchoPicklist().load({ id })`.
  ExportEchoPicklist(data?: any) {
    const self = this
    return new ExportEchoPicklistEntity(self,data)
  }


  // Entity access: `client.ExportOpentronsProtocol().list()` / `client.ExportOpentronsProtocol().load({ id })`.
  ExportOpentronsProtocol(data?: any) {
    const self = this
    return new ExportOpentronsProtocolEntity(self,data)
  }


  // Entity access: `client.ExportPlateLayout().list()` / `client.ExportPlateLayout().load({ id })`.
  ExportPlateLayout(data?: any) {
    const self = this
    return new ExportPlateLayoutEntity(self,data)
  }


  // Entity access: `client.ExpressionHeatmapCluster().list()` / `client.ExpressionHeatmapCluster().load({ id })`.
  ExpressionHeatmapCluster(data?: any) {
    const self = this
    return new ExpressionHeatmapClusterEntity(self,data)
  }


  // Entity access: `client.FastqQcReport().list()` / `client.FastqQcReport().load({ id })`.
  FastqQcReport(data?: any) {
    const self = this
    return new FastqQcReportEntity(self,data)
  }


  // Entity access: `client.FastqTrim().list()` / `client.FastqTrim().load({ id })`.
  FastqTrim(data?: any) {
    const self = this
    return new FastqTrimEntity(self,data)
  }


  // Entity access: `client.FindOrf().list()` / `client.FindOrf().load({ id })`.
  FindOrf(data?: any) {
    const self = this
    return new FindOrfEntity(self,data)
  }


  // Entity access: `client.FormatSequence().list()` / `client.FormatSequence().load({ id })`.
  FormatSequence(data?: any) {
    const self = this
    return new FormatSequenceEntity(self,data)
  }


  // Entity access: `client.FunctionalEnrichment().list()` / `client.FunctionalEnrichment().load({ id })`.
  FunctionalEnrichment(data?: any) {
    const self = this
    return new FunctionalEnrichmentEntity(self,data)
  }


  // Entity access: `client.GcContent().list()` / `client.GcContent().load({ id })`.
  GcContent(data?: any) {
    const self = this
    return new GcContentEntity(self,data)
  }


  // Entity access: `client.GeneDossier().list()` / `client.GeneDossier().load({ id })`.
  GeneDossier(data?: any) {
    const self = this
    return new GeneDossierEntity(self,data)
  }


  // Entity access: `client.GeneExpression().list()` / `client.GeneExpression().load({ id })`.
  GeneExpression(data?: any) {
    const self = this
    return new GeneExpressionEntity(self,data)
  }


  // Entity access: `client.GeneModel().list()` / `client.GeneModel().load({ id })`.
  GeneModel(data?: any) {
    const self = this
    return new GeneModelEntity(self,data)
  }


  // Entity access: `client.GoldenGateFidelity().list()` / `client.GoldenGateFidelity().load({ id })`.
  GoldenGateFidelity(data?: any) {
    const self = this
    return new GoldenGateFidelityEntity(self,data)
  }


  // Entity access: `client.HgvsConvert().list()` / `client.HgvsConvert().load({ id })`.
  HgvsConvert(data?: any) {
    const self = this
    return new HgvsConvertEntity(self,data)
  }


  // Entity access: `client.IdMapPoll().list()` / `client.IdMapPoll().load({ id })`.
  IdMapPoll(data?: any) {
    const self = this
    return new IdMapPollEntity(self,data)
  }


  // Entity access: `client.IdMapSubmit().list()` / `client.IdMapSubmit().load({ id })`.
  IdMapSubmit(data?: any) {
    const self = this
    return new IdMapSubmitEntity(self,data)
  }


  // Entity access: `client.InSilicoPcr().list()` / `client.InSilicoPcr().load({ id })`.
  InSilicoPcr(data?: any) {
    const self = this
    return new InSilicoPcrEntity(self,data)
  }


  // Entity access: `client.KaspPrimerDesign().list()` / `client.KaspPrimerDesign().load({ id })`.
  KaspPrimerDesign(data?: any) {
    const self = this
    return new KaspPrimerDesignEntity(self,data)
  }


  // Entity access: `client.ListTool().list()` / `client.ListTool().load({ id })`.
  ListTool(data?: any) {
    const self = this
    return new ListToolEntity(self,data)
  }


  // Entity access: `client.MeltingTemperature().list()` / `client.MeltingTemperature().load({ id })`.
  MeltingTemperature(data?: any) {
    const self = this
    return new MeltingTemperatureEntity(self,data)
  }


  // Entity access: `client.MotifFinder().list()` / `client.MotifFinder().load({ id })`.
  MotifFinder(data?: any) {
    const self = this
    return new MotifFinderEntity(self,data)
  }


  // Entity access: `client.MultipleSequenceAlignment().list()` / `client.MultipleSequenceAlignment().load({ id })`.
  MultipleSequenceAlignment(data?: any) {
    const self = this
    return new MultipleSequenceAlignmentEntity(self,data)
  }


  // Entity access: `client.OligoAnalysi().list()` / `client.OligoAnalysi().load({ id })`.
  OligoAnalysi(data?: any) {
    const self = this
    return new OligoAnalysiEntity(self,data)
  }


  // Entity access: `client.OrthologMap().list()` / `client.OrthologMap().load({ id })`.
  OrthologMap(data?: any) {
    const self = this
    return new OrthologMapEntity(self,data)
  }


  // Entity access: `client.PairwiseAlignment().list()` / `client.PairwiseAlignment().load({ id })`.
  PairwiseAlignment(data?: any) {
    const self = this
    return new PairwiseAlignmentEntity(self,data)
  }


  // Entity access: `client.ParseGenbank().list()` / `client.ParseGenbank().load({ id })`.
  ParseGenbank(data?: any) {
    const self = this
    return new ParseGenbankEntity(self,data)
  }


  // Entity access: `client.ParseSangerTrace().list()` / `client.ParseSangerTrace().load({ id })`.
  ParseSangerTrace(data?: any) {
    const self = this
    return new ParseSangerTraceEntity(self,data)
  }


  // Entity access: `client.PlasmidAnnotate().list()` / `client.PlasmidAnnotate().load({ id })`.
  PlasmidAnnotate(data?: any) {
    const self = this
    return new PlasmidAnnotateEntity(self,data)
  }


  // Entity access: `client.PlasmidDeepAnnotate().list()` / `client.PlasmidDeepAnnotate().load({ id })`.
  PlasmidDeepAnnotate(data?: any) {
    const self = this
    return new PlasmidDeepAnnotateEntity(self,data)
  }


  // Entity access: `client.PlasmidFullReport().list()` / `client.PlasmidFullReport().load({ id })`.
  PlasmidFullReport(data?: any) {
    const self = this
    return new PlasmidFullReportEntity(self,data)
  }


  // Entity access: `client.PlasmidIdentify().list()` / `client.PlasmidIdentify().load({ id })`.
  PlasmidIdentify(data?: any) {
    const self = this
    return new PlasmidIdentifyEntity(self,data)
  }


  // Entity access: `client.PrimeEditingDesign().list()` / `client.PrimeEditingDesign().load({ id })`.
  PrimeEditingDesign(data?: any) {
    const self = this
    return new PrimeEditingDesignEntity(self,data)
  }


  // Entity access: `client.PrimeEditingTwinDesign().list()` / `client.PrimeEditingTwinDesign().load({ id })`.
  PrimeEditingTwinDesign(data?: any) {
    const self = this
    return new PrimeEditingTwinDesignEntity(self,data)
  }


  // Entity access: `client.PrimerDesign().list()` / `client.PrimerDesign().load({ id })`.
  PrimerDesign(data?: any) {
    const self = this
    return new PrimerDesignEntity(self,data)
  }


  // Entity access: `client.PrimerSpecificity().list()` / `client.PrimerSpecificity().load({ id })`.
  PrimerSpecificity(data?: any) {
    const self = this
    return new PrimerSpecificityEntity(self,data)
  }


  // Entity access: `client.ProteaseDigestion().list()` / `client.ProteaseDigestion().load({ id })`.
  ProteaseDigestion(data?: any) {
    const self = this
    return new ProteaseDigestionEntity(self,data)
  }


  // Entity access: `client.ProteinAnnotatePoll().list()` / `client.ProteinAnnotatePoll().load({ id })`.
  ProteinAnnotatePoll(data?: any) {
    const self = this
    return new ProteinAnnotatePollEntity(self,data)
  }


  // Entity access: `client.ProteinAnnotateSubmit().list()` / `client.ProteinAnnotateSubmit().load({ id })`.
  ProteinAnnotateSubmit(data?: any) {
    const self = this
    return new ProteinAnnotateSubmitEntity(self,data)
  }


  // Entity access: `client.ProteinHydrophobicity().list()` / `client.ProteinHydrophobicity().load({ id })`.
  ProteinHydrophobicity(data?: any) {
    const self = this
    return new ProteinHydrophobicityEntity(self,data)
  }


  // Entity access: `client.ProteinProperty().list()` / `client.ProteinProperty().load({ id })`.
  ProteinProperty(data?: any) {
    const self = this
    return new ProteinPropertyEntity(self,data)
  }


  // Entity access: `client.RandomSequence().list()` / `client.RandomSequence().load({ id })`.
  RandomSequence(data?: any) {
    const self = this
    return new RandomSequenceEntity(self,data)
  }


  // Entity access: `client.RestrictionSite().list()` / `client.RestrictionSite().load({ id })`.
  RestrictionSite(data?: any) {
    const self = this
    return new RestrictionSiteEntity(self,data)
  }


  // Entity access: `client.ReverseComplement().list()` / `client.ReverseComplement().load({ id })`.
  ReverseComplement(data?: any) {
    const self = this
    return new ReverseComplementEntity(self,data)
  }


  // Entity access: `client.ReverseTranslate().list()` / `client.ReverseTranslate().load({ id })`.
  ReverseTranslate(data?: any) {
    const self = this
    return new ReverseTranslateEntity(self,data)
  }


  // Entity access: `client.RnaFold().list()` / `client.RnaFold().load({ id })`.
  RnaFold(data?: any) {
    const self = this
    return new RnaFoldEntity(self,data)
  }


  // Entity access: `client.SangerVsReference().list()` / `client.SangerVsReference().load({ id })`.
  SangerVsReference(data?: any) {
    const self = this
    return new SangerVsReferenceEntity(self,data)
  }


  // Entity access: `client.SavePermalink().list()` / `client.SavePermalink().load({ id })`.
  SavePermalink(data?: any) {
    const self = this
    return new SavePermalinkEntity(self,data)
  }


  // Entity access: `client.SeqfileStat().list()` / `client.SeqfileStat().load({ id })`.
  SeqfileStat(data?: any) {
    const self = this
    return new SeqfileStatEntity(self,data)
  }


  // Entity access: `client.SequenceFetch().list()` / `client.SequenceFetch().load({ id })`.
  SequenceFetch(data?: any) {
    const self = this
    return new SequenceFetchEntity(self,data)
  }


  // Entity access: `client.SequenceFormatConvert().list()` / `client.SequenceFormatConvert().load({ id })`.
  SequenceFormatConvert(data?: any) {
    const self = this
    return new SequenceFormatConvertEntity(self,data)
  }


  // Entity access: `client.SequenceReport().list()` / `client.SequenceReport().load({ id })`.
  SequenceReport(data?: any) {
    const self = this
    return new SequenceReportEntity(self,data)
  }


  // Entity access: `client.SequenceSearch().list()` / `client.SequenceSearch().load({ id })`.
  SequenceSearch(data?: any) {
    const self = this
    return new SequenceSearchEntity(self,data)
  }


  // Entity access: `client.SequencingReadbackVerify().list()` / `client.SequencingReadbackVerify().load({ id })`.
  SequencingReadbackVerify(data?: any) {
    const self = this
    return new SequencingReadbackVerifyEntity(self,data)
  }


  // Entity access: `client.SessionCreate().list()` / `client.SessionCreate().load({ id })`.
  SessionCreate(data?: any) {
    const self = this
    return new SessionCreateEntity(self,data)
  }


  // Entity access: `client.SessionGet().list()` / `client.SessionGet().load({ id })`.
  SessionGet(data?: any) {
    const self = this
    return new SessionGetEntity(self,data)
  }


  // Entity access: `client.SessionRun().list()` / `client.SessionRun().load({ id })`.
  SessionRun(data?: any) {
    const self = this
    return new SessionRunEntity(self,data)
  }


  // Entity access: `client.SessionSet().list()` / `client.SessionSet().load({ id })`.
  SessionSet(data?: any) {
    const self = this
    return new SessionSetEntity(self,data)
  }


  // Entity access: `client.SirnaDesign().list()` / `client.SirnaDesign().load({ id })`.
  SirnaDesign(data?: any) {
    const self = this
    return new SirnaDesignEntity(self,data)
  }


  // Entity access: `client.SiteDirectedMutagenesi().list()` / `client.SiteDirectedMutagenesi().load({ id })`.
  SiteDirectedMutagenesi(data?: any) {
    const self = this
    return new SiteDirectedMutagenesiEntity(self,data)
  }


  // Entity access: `client.Translate().list()` / `client.Translate().load({ id })`.
  Translate(data?: any) {
    const self = this
    return new TranslateEntity(self,data)
  }


  // Entity access: `client.VariantAnnotate().list()` / `client.VariantAnnotate().load({ id })`.
  VariantAnnotate(data?: any) {
    const self = this
    return new VariantAnnotateEntity(self,data)
  }


  // Entity access: `client.VariantComparator().list()` / `client.VariantComparator().load({ id })`.
  VariantComparator(data?: any) {
    const self = this
    return new VariantComparatorEntity(self,data)
  }


  // Entity access: `client.VerifyAssembly().list()` / `client.VerifyAssembly().load({ id })`.
  VerifyAssembly(data?: any) {
    const self = this
    return new VerifyAssemblyEntity(self,data)
  }


  // Entity access: `client.VerifyConstruct().list()` / `client.VerifyConstruct().load({ id })`.
  VerifyConstruct(data?: any) {
    const self = this
    return new VerifyConstructEntity(self,data)
  }


  // Entity access: `client.VirtualGel().list()` / `client.VirtualGel().load({ id })`.
  VirtualGel(data?: any) {
    const self = this
    return new VirtualGelEntity(self,data)
  }


  // Entity access: `client.VolcanoPlotData().list()` / `client.VolcanoPlotData().load({ id })`.
  VolcanoPlotData(data?: any) {
    const self = this
    return new VolcanoPlotDataEntity(self,data)
  }


  // Entity access: `client.WebSearch().list()` / `client.WebSearch().load({ id })`.
  WebSearch(data?: any) {
    const self = this
    return new WebSearchEntity(self,data)
  }




  static test(testoptsarg?: any, sdkoptsarg?: any) {
    const struct = stdutil.struct
    const setpath = struct.setpath
    const getdef = struct.getdef
    const clone = struct.clone
    const setprop = struct.setprop

    const sdkopts = getdef(clone(sdkoptsarg), {})
    const testopts = getdef(clone(testoptsarg), {})
    setprop(testopts, 'active', true)
    setpath(sdkopts, 'feature.test', testopts)

    const testsdk = new SeqbenchMcpSDK(sdkopts)
    testsdk._mode = 'test'

    return testsdk
  }


  tester(testopts?: any, sdkopts?: any) {
    return SeqbenchMcpSDK.test(testopts, sdkopts)
  }


  toJSON() {
    return { name: 'SeqbenchMcp' }
  }

  toString() {
    return 'SeqbenchMcp ' + this._utility.struct.jsonify(this.toJSON())
  }

  [inspect.custom]() {
    return this.toString()
  }

}




const SDK = SeqbenchMcpSDK


export {
  stdutil,
  config,

  BaseFeature,
  SeqbenchMcpEntityBase,

  SeqbenchMcpSDK,
  SDK,
}


