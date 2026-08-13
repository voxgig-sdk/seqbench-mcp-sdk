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


  // Raw endpoint access is operator-controllable, like every entity op.
  // Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
  // either one reaches the same endpoint.
  async direct(fetchargs?: any) {
    if (!this._options.allow.op.includes('direct')) {
      return {
        ok: false,
        err: new Error('SeqbenchMcpSDK: direct: operation not allowed by' +
          ' SDK option allow.op value: "' + this._options.allow.op + '"'),
      }
    }

    return this._rawRequest(fetchargs)
  }


  // Ungated request path shared by direct() and graphql(), each of which
  // checks its own allow.op token first. Private, rather than a flag on
  // fetchargs: a caller-supplied marker would let anyone opt straight back
  // out of the gate by passing it.
  async _rawRequest(fetchargs?: any) {
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



  // Raw GraphQL access: the pressure valve that makes the generated
  // surface's deliberate omissions (per-call selection sets, typed filter
  // builders, batching, subscriptions) livable — the whole schema stays
  // reachable.
  //
  // Thin wrapper over the same prepare/fetch path `direct` uses, with the
  // one thing raw `direct` cannot do for GraphQL: a GraphQL failure rides
  // HTTP 200 as a top-level `errors` array, so status alone would report a
  // failed query as ok.
  //
  // NOTE: like `direct`, this bypasses the feature pipeline — no retry,
  // ratelimit or paging features apply.
  async graphql(query: string, variables?: any, ctrl?: any) {
    const options = this._options

    if (!options.allow.op.includes('graphql')) {
      return {
        ok: false,
        err: new Error('SeqbenchMcpSDK: graphql: operation not allowed by' +
          ' SDK option allow.op value: "' + options.allow.op + '"'),
      }
    }

    const res: any = await this._rawRequest({
      method: 'POST',
      headers: { 'content-type': 'application/json' },
      body: { query, variables: variables || {} },
      ctrl,
    })

    if (res instanceof Error) {
      return res
    }

    // Errors are read BEFORE any status check: a GraphQL parse or validation
    // failure comes back as HTTP 400 carrying the standard { errors: [...] }
    // body, and the raw path represents a non-2xx as { ok: false } with no
    // err — so returning early on status would discard the server's own
    // diagnostics, which are the only useful part of that response.
    const errors = null == res.data ? undefined : res.data.errors

    if (null != errors && Array.isArray(errors) && 0 < errors.length) {
      const first = errors[0] || {}
      const err: any = new Error('SeqbenchMcpSDK: graphql: ' +
        (first.message || 'graphql error'))
      err.graphql = errors
      return { ok: false, status: res.status, headers: res.headers, err, data: res.data }
    }

    return res
  }



  // Entity access: `client.AlphafoldLookup().list()` / `client.AlphafoldLookup().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  AlphafoldLookup(entopts?: Record<string, any>) {
    const self = this
    return new AlphafoldLookupEntity(self, entopts)
  }


  // Entity access: `client.AsoDesign().list()` / `client.AsoDesign().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  AsoDesign(entopts?: Record<string, any>) {
    const self = this
    return new AsoDesignEntity(self, entopts)
  }


  // Entity access: `client.BaseEditingDesign().list()` / `client.BaseEditingDesign().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  BaseEditingDesign(entopts?: Record<string, any>) {
    const self = this
    return new BaseEditingDesignEntity(self, entopts)
  }


  // Entity access: `client.Batch().list()` / `client.Batch().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Batch(entopts?: Record<string, any>) {
    const self = this
    return new BatchEntity(self, entopts)
  }


  // Entity access: `client.BatchWorkflow().list()` / `client.BatchWorkflow().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  BatchWorkflow(entopts?: Record<string, any>) {
    const self = this
    return new BatchWorkflowEntity(self, entopts)
  }


  // Entity access: `client.CharacterizeSequence().list()` / `client.CharacterizeSequence().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CharacterizeSequence(entopts?: Record<string, any>) {
    const self = this
    return new CharacterizeSequenceEntity(self, entopts)
  }


  // Entity access: `client.CloningSimulate().list()` / `client.CloningSimulate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CloningSimulate(entopts?: Record<string, any>) {
    const self = this
    return new CloningSimulateEntity(self, entopts)
  }


  // Entity access: `client.CodonAdaptationIndex().list()` / `client.CodonAdaptationIndex().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CodonAdaptationIndex(entopts?: Record<string, any>) {
    const self = this
    return new CodonAdaptationIndexEntity(self, entopts)
  }


  // Entity access: `client.CodonOptimize().list()` / `client.CodonOptimize().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CodonOptimize(entopts?: Record<string, any>) {
    const self = this
    return new CodonOptimizeEntity(self, entopts)
  }


  // Entity access: `client.ConstructAutofix().list()` / `client.ConstructAutofix().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ConstructAutofix(entopts?: Record<string, any>) {
    const self = this
    return new ConstructAutofixEntity(self, entopts)
  }


  // Entity access: `client.ConstructQc().list()` / `client.ConstructQc().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ConstructQc(entopts?: Record<string, any>) {
    const self = this
    return new ConstructQcEntity(self, entopts)
  }


  // Entity access: `client.CrisprGrnaDesign().list()` / `client.CrisprGrnaDesign().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CrisprGrnaDesign(entopts?: Record<string, any>) {
    const self = this
    return new CrisprGrnaDesignEntity(self, entopts)
  }


  // Entity access: `client.CrisprHdrDonor().list()` / `client.CrisprHdrDonor().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CrisprHdrDonor(entopts?: Record<string, any>) {
    const self = this
    return new CrisprHdrDonorEntity(self, entopts)
  }


  // Entity access: `client.CrisprOfftargetCheck().list()` / `client.CrisprOfftargetCheck().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CrisprOfftargetCheck(entopts?: Record<string, any>) {
    const self = this
    return new CrisprOfftargetCheckEntity(self, entopts)
  }


  // Entity access: `client.CrossDimer().list()` / `client.CrossDimer().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CrossDimer(entopts?: Record<string, any>) {
    const self = this
    return new CrossDimerEntity(self, entopts)
  }


  // Entity access: `client.DnaMolarity().list()` / `client.DnaMolarity().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  DnaMolarity(entopts?: Record<string, any>) {
    const self = this
    return new DnaMolarityEntity(self, entopts)
  }


  // Entity access: `client.DoubleDigest().list()` / `client.DoubleDigest().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  DoubleDigest(entopts?: Record<string, any>) {
    const self = this
    return new DoubleDigestEntity(self, entopts)
  }


  // Entity access: `client.ExportEchoPicklist().list()` / `client.ExportEchoPicklist().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ExportEchoPicklist(entopts?: Record<string, any>) {
    const self = this
    return new ExportEchoPicklistEntity(self, entopts)
  }


  // Entity access: `client.ExportOpentronsProtocol().list()` / `client.ExportOpentronsProtocol().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ExportOpentronsProtocol(entopts?: Record<string, any>) {
    const self = this
    return new ExportOpentronsProtocolEntity(self, entopts)
  }


  // Entity access: `client.ExportPlateLayout().list()` / `client.ExportPlateLayout().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ExportPlateLayout(entopts?: Record<string, any>) {
    const self = this
    return new ExportPlateLayoutEntity(self, entopts)
  }


  // Entity access: `client.ExpressionHeatmapCluster().list()` / `client.ExpressionHeatmapCluster().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ExpressionHeatmapCluster(entopts?: Record<string, any>) {
    const self = this
    return new ExpressionHeatmapClusterEntity(self, entopts)
  }


  // Entity access: `client.FastqQcReport().list()` / `client.FastqQcReport().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  FastqQcReport(entopts?: Record<string, any>) {
    const self = this
    return new FastqQcReportEntity(self, entopts)
  }


  // Entity access: `client.FastqTrim().list()` / `client.FastqTrim().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  FastqTrim(entopts?: Record<string, any>) {
    const self = this
    return new FastqTrimEntity(self, entopts)
  }


  // Entity access: `client.FindOrf().list()` / `client.FindOrf().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  FindOrf(entopts?: Record<string, any>) {
    const self = this
    return new FindOrfEntity(self, entopts)
  }


  // Entity access: `client.FormatSequence().list()` / `client.FormatSequence().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  FormatSequence(entopts?: Record<string, any>) {
    const self = this
    return new FormatSequenceEntity(self, entopts)
  }


  // Entity access: `client.FunctionalEnrichment().list()` / `client.FunctionalEnrichment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  FunctionalEnrichment(entopts?: Record<string, any>) {
    const self = this
    return new FunctionalEnrichmentEntity(self, entopts)
  }


  // Entity access: `client.GcContent().list()` / `client.GcContent().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  GcContent(entopts?: Record<string, any>) {
    const self = this
    return new GcContentEntity(self, entopts)
  }


  // Entity access: `client.GeneDossier().list()` / `client.GeneDossier().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  GeneDossier(entopts?: Record<string, any>) {
    const self = this
    return new GeneDossierEntity(self, entopts)
  }


  // Entity access: `client.GeneExpression().list()` / `client.GeneExpression().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  GeneExpression(entopts?: Record<string, any>) {
    const self = this
    return new GeneExpressionEntity(self, entopts)
  }


  // Entity access: `client.GeneModel().list()` / `client.GeneModel().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  GeneModel(entopts?: Record<string, any>) {
    const self = this
    return new GeneModelEntity(self, entopts)
  }


  // Entity access: `client.GoldenGateFidelity().list()` / `client.GoldenGateFidelity().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  GoldenGateFidelity(entopts?: Record<string, any>) {
    const self = this
    return new GoldenGateFidelityEntity(self, entopts)
  }


  // Entity access: `client.HgvsConvert().list()` / `client.HgvsConvert().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  HgvsConvert(entopts?: Record<string, any>) {
    const self = this
    return new HgvsConvertEntity(self, entopts)
  }


  // Entity access: `client.IdMapPoll().list()` / `client.IdMapPoll().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  IdMapPoll(entopts?: Record<string, any>) {
    const self = this
    return new IdMapPollEntity(self, entopts)
  }


  // Entity access: `client.IdMapSubmit().list()` / `client.IdMapSubmit().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  IdMapSubmit(entopts?: Record<string, any>) {
    const self = this
    return new IdMapSubmitEntity(self, entopts)
  }


  // Entity access: `client.InSilicoPcr().list()` / `client.InSilicoPcr().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  InSilicoPcr(entopts?: Record<string, any>) {
    const self = this
    return new InSilicoPcrEntity(self, entopts)
  }


  // Entity access: `client.KaspPrimerDesign().list()` / `client.KaspPrimerDesign().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  KaspPrimerDesign(entopts?: Record<string, any>) {
    const self = this
    return new KaspPrimerDesignEntity(self, entopts)
  }


  // Entity access: `client.ListTool().list()` / `client.ListTool().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ListTool(entopts?: Record<string, any>) {
    const self = this
    return new ListToolEntity(self, entopts)
  }


  // Entity access: `client.MeltingTemperature().list()` / `client.MeltingTemperature().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  MeltingTemperature(entopts?: Record<string, any>) {
    const self = this
    return new MeltingTemperatureEntity(self, entopts)
  }


  // Entity access: `client.MotifFinder().list()` / `client.MotifFinder().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  MotifFinder(entopts?: Record<string, any>) {
    const self = this
    return new MotifFinderEntity(self, entopts)
  }


  // Entity access: `client.MultipleSequenceAlignment().list()` / `client.MultipleSequenceAlignment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  MultipleSequenceAlignment(entopts?: Record<string, any>) {
    const self = this
    return new MultipleSequenceAlignmentEntity(self, entopts)
  }


  // Entity access: `client.OligoAnalysi().list()` / `client.OligoAnalysi().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  OligoAnalysi(entopts?: Record<string, any>) {
    const self = this
    return new OligoAnalysiEntity(self, entopts)
  }


  // Entity access: `client.OrthologMap().list()` / `client.OrthologMap().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  OrthologMap(entopts?: Record<string, any>) {
    const self = this
    return new OrthologMapEntity(self, entopts)
  }


  // Entity access: `client.PairwiseAlignment().list()` / `client.PairwiseAlignment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PairwiseAlignment(entopts?: Record<string, any>) {
    const self = this
    return new PairwiseAlignmentEntity(self, entopts)
  }


  // Entity access: `client.ParseGenbank().list()` / `client.ParseGenbank().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ParseGenbank(entopts?: Record<string, any>) {
    const self = this
    return new ParseGenbankEntity(self, entopts)
  }


  // Entity access: `client.ParseSangerTrace().list()` / `client.ParseSangerTrace().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ParseSangerTrace(entopts?: Record<string, any>) {
    const self = this
    return new ParseSangerTraceEntity(self, entopts)
  }


  // Entity access: `client.PlasmidAnnotate().list()` / `client.PlasmidAnnotate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PlasmidAnnotate(entopts?: Record<string, any>) {
    const self = this
    return new PlasmidAnnotateEntity(self, entopts)
  }


  // Entity access: `client.PlasmidDeepAnnotate().list()` / `client.PlasmidDeepAnnotate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PlasmidDeepAnnotate(entopts?: Record<string, any>) {
    const self = this
    return new PlasmidDeepAnnotateEntity(self, entopts)
  }


  // Entity access: `client.PlasmidFullReport().list()` / `client.PlasmidFullReport().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PlasmidFullReport(entopts?: Record<string, any>) {
    const self = this
    return new PlasmidFullReportEntity(self, entopts)
  }


  // Entity access: `client.PlasmidIdentify().list()` / `client.PlasmidIdentify().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PlasmidIdentify(entopts?: Record<string, any>) {
    const self = this
    return new PlasmidIdentifyEntity(self, entopts)
  }


  // Entity access: `client.PrimeEditingDesign().list()` / `client.PrimeEditingDesign().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PrimeEditingDesign(entopts?: Record<string, any>) {
    const self = this
    return new PrimeEditingDesignEntity(self, entopts)
  }


  // Entity access: `client.PrimeEditingTwinDesign().list()` / `client.PrimeEditingTwinDesign().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PrimeEditingTwinDesign(entopts?: Record<string, any>) {
    const self = this
    return new PrimeEditingTwinDesignEntity(self, entopts)
  }


  // Entity access: `client.PrimerDesign().list()` / `client.PrimerDesign().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PrimerDesign(entopts?: Record<string, any>) {
    const self = this
    return new PrimerDesignEntity(self, entopts)
  }


  // Entity access: `client.PrimerSpecificity().list()` / `client.PrimerSpecificity().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PrimerSpecificity(entopts?: Record<string, any>) {
    const self = this
    return new PrimerSpecificityEntity(self, entopts)
  }


  // Entity access: `client.ProteaseDigestion().list()` / `client.ProteaseDigestion().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ProteaseDigestion(entopts?: Record<string, any>) {
    const self = this
    return new ProteaseDigestionEntity(self, entopts)
  }


  // Entity access: `client.ProteinAnnotatePoll().list()` / `client.ProteinAnnotatePoll().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ProteinAnnotatePoll(entopts?: Record<string, any>) {
    const self = this
    return new ProteinAnnotatePollEntity(self, entopts)
  }


  // Entity access: `client.ProteinAnnotateSubmit().list()` / `client.ProteinAnnotateSubmit().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ProteinAnnotateSubmit(entopts?: Record<string, any>) {
    const self = this
    return new ProteinAnnotateSubmitEntity(self, entopts)
  }


  // Entity access: `client.ProteinHydrophobicity().list()` / `client.ProteinHydrophobicity().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ProteinHydrophobicity(entopts?: Record<string, any>) {
    const self = this
    return new ProteinHydrophobicityEntity(self, entopts)
  }


  // Entity access: `client.ProteinProperty().list()` / `client.ProteinProperty().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ProteinProperty(entopts?: Record<string, any>) {
    const self = this
    return new ProteinPropertyEntity(self, entopts)
  }


  // Entity access: `client.RandomSequence().list()` / `client.RandomSequence().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  RandomSequence(entopts?: Record<string, any>) {
    const self = this
    return new RandomSequenceEntity(self, entopts)
  }


  // Entity access: `client.RestrictionSite().list()` / `client.RestrictionSite().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  RestrictionSite(entopts?: Record<string, any>) {
    const self = this
    return new RestrictionSiteEntity(self, entopts)
  }


  // Entity access: `client.ReverseComplement().list()` / `client.ReverseComplement().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ReverseComplement(entopts?: Record<string, any>) {
    const self = this
    return new ReverseComplementEntity(self, entopts)
  }


  // Entity access: `client.ReverseTranslate().list()` / `client.ReverseTranslate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ReverseTranslate(entopts?: Record<string, any>) {
    const self = this
    return new ReverseTranslateEntity(self, entopts)
  }


  // Entity access: `client.RnaFold().list()` / `client.RnaFold().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  RnaFold(entopts?: Record<string, any>) {
    const self = this
    return new RnaFoldEntity(self, entopts)
  }


  // Entity access: `client.SangerVsReference().list()` / `client.SangerVsReference().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SangerVsReference(entopts?: Record<string, any>) {
    const self = this
    return new SangerVsReferenceEntity(self, entopts)
  }


  // Entity access: `client.SavePermalink().list()` / `client.SavePermalink().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SavePermalink(entopts?: Record<string, any>) {
    const self = this
    return new SavePermalinkEntity(self, entopts)
  }


  // Entity access: `client.SeqfileStat().list()` / `client.SeqfileStat().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SeqfileStat(entopts?: Record<string, any>) {
    const self = this
    return new SeqfileStatEntity(self, entopts)
  }


  // Entity access: `client.SequenceFetch().list()` / `client.SequenceFetch().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SequenceFetch(entopts?: Record<string, any>) {
    const self = this
    return new SequenceFetchEntity(self, entopts)
  }


  // Entity access: `client.SequenceFormatConvert().list()` / `client.SequenceFormatConvert().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SequenceFormatConvert(entopts?: Record<string, any>) {
    const self = this
    return new SequenceFormatConvertEntity(self, entopts)
  }


  // Entity access: `client.SequenceReport().list()` / `client.SequenceReport().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SequenceReport(entopts?: Record<string, any>) {
    const self = this
    return new SequenceReportEntity(self, entopts)
  }


  // Entity access: `client.SequenceSearch().list()` / `client.SequenceSearch().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SequenceSearch(entopts?: Record<string, any>) {
    const self = this
    return new SequenceSearchEntity(self, entopts)
  }


  // Entity access: `client.SequencingReadbackVerify().list()` / `client.SequencingReadbackVerify().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SequencingReadbackVerify(entopts?: Record<string, any>) {
    const self = this
    return new SequencingReadbackVerifyEntity(self, entopts)
  }


  // Entity access: `client.SessionCreate().list()` / `client.SessionCreate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SessionCreate(entopts?: Record<string, any>) {
    const self = this
    return new SessionCreateEntity(self, entopts)
  }


  // Entity access: `client.SessionGet().list()` / `client.SessionGet().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SessionGet(entopts?: Record<string, any>) {
    const self = this
    return new SessionGetEntity(self, entopts)
  }


  // Entity access: `client.SessionRun().list()` / `client.SessionRun().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SessionRun(entopts?: Record<string, any>) {
    const self = this
    return new SessionRunEntity(self, entopts)
  }


  // Entity access: `client.SessionSet().list()` / `client.SessionSet().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SessionSet(entopts?: Record<string, any>) {
    const self = this
    return new SessionSetEntity(self, entopts)
  }


  // Entity access: `client.SirnaDesign().list()` / `client.SirnaDesign().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SirnaDesign(entopts?: Record<string, any>) {
    const self = this
    return new SirnaDesignEntity(self, entopts)
  }


  // Entity access: `client.SiteDirectedMutagenesi().list()` / `client.SiteDirectedMutagenesi().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  SiteDirectedMutagenesi(entopts?: Record<string, any>) {
    const self = this
    return new SiteDirectedMutagenesiEntity(self, entopts)
  }


  // Entity access: `client.Translate().list()` / `client.Translate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Translate(entopts?: Record<string, any>) {
    const self = this
    return new TranslateEntity(self, entopts)
  }


  // Entity access: `client.VariantAnnotate().list()` / `client.VariantAnnotate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  VariantAnnotate(entopts?: Record<string, any>) {
    const self = this
    return new VariantAnnotateEntity(self, entopts)
  }


  // Entity access: `client.VariantComparator().list()` / `client.VariantComparator().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  VariantComparator(entopts?: Record<string, any>) {
    const self = this
    return new VariantComparatorEntity(self, entopts)
  }


  // Entity access: `client.VerifyAssembly().list()` / `client.VerifyAssembly().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  VerifyAssembly(entopts?: Record<string, any>) {
    const self = this
    return new VerifyAssemblyEntity(self, entopts)
  }


  // Entity access: `client.VerifyConstruct().list()` / `client.VerifyConstruct().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  VerifyConstruct(entopts?: Record<string, any>) {
    const self = this
    return new VerifyConstructEntity(self, entopts)
  }


  // Entity access: `client.VirtualGel().list()` / `client.VirtualGel().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  VirtualGel(entopts?: Record<string, any>) {
    const self = this
    return new VirtualGelEntity(self, entopts)
  }


  // Entity access: `client.VolcanoPlotData().list()` / `client.VolcanoPlotData().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  VolcanoPlotData(entopts?: Record<string, any>) {
    const self = this
    return new VolcanoPlotDataEntity(self, entopts)
  }


  // Entity access: `client.WebSearch().list()` / `client.WebSearch().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  WebSearch(entopts?: Record<string, any>) {
    const self = this
    return new WebSearchEntity(self, entopts)
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


