"use strict";
// SeqbenchMcp Ts SDK
Object.defineProperty(exports, "__esModule", { value: true });
exports.SDK = exports.SeqbenchMcpSDK = exports.SeqbenchMcpEntityBase = exports.BaseFeature = exports.config = exports.stdutil = void 0;
const AlphafoldLookupEntity_1 = require("./entity/AlphafoldLookupEntity");
const AsoDesignEntity_1 = require("./entity/AsoDesignEntity");
const BaseEditingDesignEntity_1 = require("./entity/BaseEditingDesignEntity");
const BatchEntity_1 = require("./entity/BatchEntity");
const BatchWorkflowEntity_1 = require("./entity/BatchWorkflowEntity");
const CharacterizeSequenceEntity_1 = require("./entity/CharacterizeSequenceEntity");
const CloningSimulateEntity_1 = require("./entity/CloningSimulateEntity");
const CodonAdaptationIndexEntity_1 = require("./entity/CodonAdaptationIndexEntity");
const CodonOptimizeEntity_1 = require("./entity/CodonOptimizeEntity");
const ConstructAutofixEntity_1 = require("./entity/ConstructAutofixEntity");
const ConstructQcEntity_1 = require("./entity/ConstructQcEntity");
const CrisprGrnaDesignEntity_1 = require("./entity/CrisprGrnaDesignEntity");
const CrisprHdrDonorEntity_1 = require("./entity/CrisprHdrDonorEntity");
const CrisprOfftargetCheckEntity_1 = require("./entity/CrisprOfftargetCheckEntity");
const CrossDimerEntity_1 = require("./entity/CrossDimerEntity");
const DnaMolarityEntity_1 = require("./entity/DnaMolarityEntity");
const DoubleDigestEntity_1 = require("./entity/DoubleDigestEntity");
const ExportEchoPicklistEntity_1 = require("./entity/ExportEchoPicklistEntity");
const ExportOpentronsProtocolEntity_1 = require("./entity/ExportOpentronsProtocolEntity");
const ExportPlateLayoutEntity_1 = require("./entity/ExportPlateLayoutEntity");
const ExpressionHeatmapClusterEntity_1 = require("./entity/ExpressionHeatmapClusterEntity");
const FastqQcReportEntity_1 = require("./entity/FastqQcReportEntity");
const FastqTrimEntity_1 = require("./entity/FastqTrimEntity");
const FindOrfEntity_1 = require("./entity/FindOrfEntity");
const FormatSequenceEntity_1 = require("./entity/FormatSequenceEntity");
const FunctionalEnrichmentEntity_1 = require("./entity/FunctionalEnrichmentEntity");
const GcContentEntity_1 = require("./entity/GcContentEntity");
const GeneDossierEntity_1 = require("./entity/GeneDossierEntity");
const GeneExpressionEntity_1 = require("./entity/GeneExpressionEntity");
const GeneModelEntity_1 = require("./entity/GeneModelEntity");
const GoldenGateFidelityEntity_1 = require("./entity/GoldenGateFidelityEntity");
const HgvsConvertEntity_1 = require("./entity/HgvsConvertEntity");
const IdMapPollEntity_1 = require("./entity/IdMapPollEntity");
const IdMapSubmitEntity_1 = require("./entity/IdMapSubmitEntity");
const InSilicoPcrEntity_1 = require("./entity/InSilicoPcrEntity");
const KaspPrimerDesignEntity_1 = require("./entity/KaspPrimerDesignEntity");
const ListToolEntity_1 = require("./entity/ListToolEntity");
const MeltingTemperatureEntity_1 = require("./entity/MeltingTemperatureEntity");
const MotifFinderEntity_1 = require("./entity/MotifFinderEntity");
const MultipleSequenceAlignmentEntity_1 = require("./entity/MultipleSequenceAlignmentEntity");
const OligoAnalysiEntity_1 = require("./entity/OligoAnalysiEntity");
const OrthologMapEntity_1 = require("./entity/OrthologMapEntity");
const PairwiseAlignmentEntity_1 = require("./entity/PairwiseAlignmentEntity");
const ParseGenbankEntity_1 = require("./entity/ParseGenbankEntity");
const ParseSangerTraceEntity_1 = require("./entity/ParseSangerTraceEntity");
const PlasmidAnnotateEntity_1 = require("./entity/PlasmidAnnotateEntity");
const PlasmidDeepAnnotateEntity_1 = require("./entity/PlasmidDeepAnnotateEntity");
const PlasmidFullReportEntity_1 = require("./entity/PlasmidFullReportEntity");
const PlasmidIdentifyEntity_1 = require("./entity/PlasmidIdentifyEntity");
const PrimeEditingDesignEntity_1 = require("./entity/PrimeEditingDesignEntity");
const PrimeEditingTwinDesignEntity_1 = require("./entity/PrimeEditingTwinDesignEntity");
const PrimerDesignEntity_1 = require("./entity/PrimerDesignEntity");
const PrimerSpecificityEntity_1 = require("./entity/PrimerSpecificityEntity");
const ProteaseDigestionEntity_1 = require("./entity/ProteaseDigestionEntity");
const ProteinAnnotatePollEntity_1 = require("./entity/ProteinAnnotatePollEntity");
const ProteinAnnotateSubmitEntity_1 = require("./entity/ProteinAnnotateSubmitEntity");
const ProteinHydrophobicityEntity_1 = require("./entity/ProteinHydrophobicityEntity");
const ProteinPropertyEntity_1 = require("./entity/ProteinPropertyEntity");
const RandomSequenceEntity_1 = require("./entity/RandomSequenceEntity");
const RestrictionSiteEntity_1 = require("./entity/RestrictionSiteEntity");
const ReverseComplementEntity_1 = require("./entity/ReverseComplementEntity");
const ReverseTranslateEntity_1 = require("./entity/ReverseTranslateEntity");
const RnaFoldEntity_1 = require("./entity/RnaFoldEntity");
const SangerVsReferenceEntity_1 = require("./entity/SangerVsReferenceEntity");
const SavePermalinkEntity_1 = require("./entity/SavePermalinkEntity");
const SeqfileStatEntity_1 = require("./entity/SeqfileStatEntity");
const SequenceFetchEntity_1 = require("./entity/SequenceFetchEntity");
const SequenceFormatConvertEntity_1 = require("./entity/SequenceFormatConvertEntity");
const SequenceReportEntity_1 = require("./entity/SequenceReportEntity");
const SequenceSearchEntity_1 = require("./entity/SequenceSearchEntity");
const SequencingReadbackVerifyEntity_1 = require("./entity/SequencingReadbackVerifyEntity");
const SessionCreateEntity_1 = require("./entity/SessionCreateEntity");
const SessionGetEntity_1 = require("./entity/SessionGetEntity");
const SessionRunEntity_1 = require("./entity/SessionRunEntity");
const SessionSetEntity_1 = require("./entity/SessionSetEntity");
const SirnaDesignEntity_1 = require("./entity/SirnaDesignEntity");
const SiteDirectedMutagenesiEntity_1 = require("./entity/SiteDirectedMutagenesiEntity");
const TranslateEntity_1 = require("./entity/TranslateEntity");
const VariantAnnotateEntity_1 = require("./entity/VariantAnnotateEntity");
const VariantComparatorEntity_1 = require("./entity/VariantComparatorEntity");
const VerifyAssemblyEntity_1 = require("./entity/VerifyAssemblyEntity");
const VerifyConstructEntity_1 = require("./entity/VerifyConstructEntity");
const VirtualGelEntity_1 = require("./entity/VirtualGelEntity");
const VolcanoPlotDataEntity_1 = require("./entity/VolcanoPlotDataEntity");
const WebSearchEntity_1 = require("./entity/WebSearchEntity");
const node_util_1 = require("node:util");
const Config_1 = require("./Config");
Object.defineProperty(exports, "config", { enumerable: true, get: function () { return Config_1.config; } });
const SeqbenchMcpEntityBase_1 = require("./SeqbenchMcpEntityBase");
Object.defineProperty(exports, "SeqbenchMcpEntityBase", { enumerable: true, get: function () { return SeqbenchMcpEntityBase_1.SeqbenchMcpEntityBase; } });
const Utility_1 = require("./utility/Utility");
const BaseFeature_1 = require("./feature/base/BaseFeature");
Object.defineProperty(exports, "BaseFeature", { enumerable: true, get: function () { return BaseFeature_1.BaseFeature; } });
const stdutil = new Utility_1.Utility();
exports.stdutil = stdutil;
class SeqbenchMcpSDK {
    _mode = 'live';
    _options;
    _utility = new Utility_1.Utility();
    _features;
    _rootctx;
    constructor(options) {
        this._rootctx = this._utility.makeContext({
            client: this,
            utility: this._utility,
            config: Config_1.config,
            options,
            shared: new WeakMap()
        });
        this._options = this._utility.makeOptions(this._rootctx);
        const struct = this._utility.struct;
        const getpath = struct.getpath;
        if (true === getpath(this._options.feature, 'test.active')) {
            this._mode = 'test';
        }
        this._rootctx.options = this._options;
        this._features = [];
        const featureAdd = this._utility.featureAdd;
        const featureInit = this._utility.featureInit;
        // Add features in the resolved order (makeOptions puts an explicit
        // array order first, else defaults to test-first). Ordering matters:
        // the `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
        // so `test` must be added before them to sit at the base of the chain.
        const extend = this._options.extend || [];
        const featureorder = getpath(this._options, '__derived__.featureorder') || [];
        for (const fname of featureorder) {
            const fopts = this._options.feature[fname] || {};
            if (fopts.active) {
                // An active name with no generated class is legal when an
                // extend-supplied instance carries that name (station's adopt
                // path): the instance is added below, positioned by its own
                // __after__ entry, so skip it here rather than fail construction.
                if (!this._rootctx.config.hasFeature(fname) &&
                    extend.some((f) => fname === f.name)) {
                    continue;
                }
                featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname));
            }
        }
        for (let f of extend) {
            featureAdd(this._rootctx, f);
        }
        for (let f of this._features) {
            featureInit(this._rootctx, f);
        }
        const featureHook = this._utility.featureHook;
        featureHook(this._rootctx, 'PostConstruct');
    }
    options() {
        return this._utility.struct.clone(this._options);
    }
    utility() {
        return this._utility.struct.clone(this._utility);
    }
    async prepare(fetchargs) {
        const utility = this._utility;
        const struct = utility.struct;
        const clone = struct.clone;
        const { makeContext, makeFetchDef, prepareHeaders, prepareAuth, } = utility;
        fetchargs = fetchargs || {};
        let ctx = makeContext({
            opname: 'prepare',
            ctrl: fetchargs.ctrl || {},
        }, this._rootctx);
        const options = this._options;
        // Build spec directly from SDK options + user-provided fetch args.
        const spec = {
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
        };
        ctx.spec = spec;
        // Merge user-provided headers over SDK defaults.
        if (fetchargs.headers) {
            const uheaders = fetchargs.headers;
            for (let key in uheaders) {
                spec.headers[key] = uheaders[key];
            }
        }
        // Apply SDK auth (apikey, auth prefix, etc.)
        const authResult = prepareAuth(ctx);
        if (authResult instanceof Error) {
            return authResult;
        }
        return makeFetchDef(ctx);
    }
    // Raw endpoint access is operator-controllable, like every entity op.
    // Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
    // either one reaches the same endpoint.
    async direct(fetchargs) {
        if (!this._options.allow.op.includes('direct')) {
            return {
                ok: false,
                err: new Error('SeqbenchMcpSDK: direct: operation not allowed by' +
                    ' SDK option allow.op value: "' + this._options.allow.op + '"'),
            };
        }
        return this._rawRequest(fetchargs);
    }
    // Ungated request path shared by direct() and graphql(), each of which
    // checks its own allow.op token first. Private, rather than a flag on
    // fetchargs: a caller-supplied marker would let anyone opt straight back
    // out of the gate by passing it.
    async _rawRequest(fetchargs) {
        const utility = this._utility;
        const fetcher = utility.fetcher;
        const makeContext = utility.makeContext;
        const fetchdef = await this.prepare(fetchargs);
        if (fetchdef instanceof Error) {
            return fetchdef;
        }
        let ctx = makeContext({
            opname: 'direct',
            ctrl: (fetchargs || {}).ctrl || {},
        }, this._rootctx);
        try {
            const fetched = await fetcher(ctx, fetchdef.url, fetchdef);
            if (null == fetched) {
                return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') };
            }
            else if (fetched instanceof Error) {
                return { ok: false, err: fetched };
            }
            const status = fetched.status;
            // No body responses (204 No Content, 304 Not Modified) and explicit
            // zero content-length must skip JSON parsing — fetched.json() would
            // throw `Unexpected end of JSON input` on an empty body.
            const headers = fetched.headers;
            const contentLength = headers && 'function' === typeof headers.get
                ? headers.get('content-length')
                : (headers || {})['content-length'];
            const noBody = 204 === status || 304 === status || '0' === String(contentLength);
            let json = undefined;
            if (!noBody) {
                try {
                    json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json;
                }
                catch (parseErr) {
                    // Body wasn't valid JSON — surface the raw response rather than
                    // throwing. data stays undefined; callers can inspect status/headers.
                    json = undefined;
                }
            }
            return {
                ok: status >= 200 && status < 300,
                status,
                headers: fetched.headers,
                data: json,
            };
        }
        catch (err) {
            return { ok: false, err };
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
    async graphql(query, variables, ctrl) {
        const options = this._options;
        if (!options.allow.op.includes('graphql')) {
            return {
                ok: false,
                err: new Error('SeqbenchMcpSDK: graphql: operation not allowed by' +
                    ' SDK option allow.op value: "' + options.allow.op + '"'),
            };
        }
        const res = await this._rawRequest({
            method: 'POST',
            headers: { 'content-type': 'application/json' },
            body: { query, variables: variables || {} },
            ctrl,
        });
        if (res instanceof Error) {
            return res;
        }
        // Errors are read BEFORE any status check: a GraphQL parse or validation
        // failure comes back as HTTP 400 carrying the standard { errors: [...] }
        // body, and the raw path represents a non-2xx as { ok: false } with no
        // err — so returning early on status would discard the server's own
        // diagnostics, which are the only useful part of that response.
        const errors = null == res.data ? undefined : res.data.errors;
        if (null != errors && Array.isArray(errors) && 0 < errors.length) {
            const first = errors[0] || {};
            const err = new Error('SeqbenchMcpSDK: graphql: ' +
                (first.message || 'graphql error'));
            err.graphql = errors;
            return { ok: false, status: res.status, headers: res.headers, err, data: res.data };
        }
        return res;
    }
    // Entity access: `client.AlphafoldLookup().list()` / `client.AlphafoldLookup().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    AlphafoldLookup(entopts) {
        const self = this;
        return new AlphafoldLookupEntity_1.AlphafoldLookupEntity(self, entopts);
    }
    // Entity access: `client.AsoDesign().list()` / `client.AsoDesign().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    AsoDesign(entopts) {
        const self = this;
        return new AsoDesignEntity_1.AsoDesignEntity(self, entopts);
    }
    // Entity access: `client.BaseEditingDesign().list()` / `client.BaseEditingDesign().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    BaseEditingDesign(entopts) {
        const self = this;
        return new BaseEditingDesignEntity_1.BaseEditingDesignEntity(self, entopts);
    }
    // Entity access: `client.Batch().list()` / `client.Batch().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Batch(entopts) {
        const self = this;
        return new BatchEntity_1.BatchEntity(self, entopts);
    }
    // Entity access: `client.BatchWorkflow().list()` / `client.BatchWorkflow().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    BatchWorkflow(entopts) {
        const self = this;
        return new BatchWorkflowEntity_1.BatchWorkflowEntity(self, entopts);
    }
    // Entity access: `client.CharacterizeSequence().list()` / `client.CharacterizeSequence().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CharacterizeSequence(entopts) {
        const self = this;
        return new CharacterizeSequenceEntity_1.CharacterizeSequenceEntity(self, entopts);
    }
    // Entity access: `client.CloningSimulate().list()` / `client.CloningSimulate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CloningSimulate(entopts) {
        const self = this;
        return new CloningSimulateEntity_1.CloningSimulateEntity(self, entopts);
    }
    // Entity access: `client.CodonAdaptationIndex().list()` / `client.CodonAdaptationIndex().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CodonAdaptationIndex(entopts) {
        const self = this;
        return new CodonAdaptationIndexEntity_1.CodonAdaptationIndexEntity(self, entopts);
    }
    // Entity access: `client.CodonOptimize().list()` / `client.CodonOptimize().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CodonOptimize(entopts) {
        const self = this;
        return new CodonOptimizeEntity_1.CodonOptimizeEntity(self, entopts);
    }
    // Entity access: `client.ConstructAutofix().list()` / `client.ConstructAutofix().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ConstructAutofix(entopts) {
        const self = this;
        return new ConstructAutofixEntity_1.ConstructAutofixEntity(self, entopts);
    }
    // Entity access: `client.ConstructQc().list()` / `client.ConstructQc().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ConstructQc(entopts) {
        const self = this;
        return new ConstructQcEntity_1.ConstructQcEntity(self, entopts);
    }
    // Entity access: `client.CrisprGrnaDesign().list()` / `client.CrisprGrnaDesign().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CrisprGrnaDesign(entopts) {
        const self = this;
        return new CrisprGrnaDesignEntity_1.CrisprGrnaDesignEntity(self, entopts);
    }
    // Entity access: `client.CrisprHdrDonor().list()` / `client.CrisprHdrDonor().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CrisprHdrDonor(entopts) {
        const self = this;
        return new CrisprHdrDonorEntity_1.CrisprHdrDonorEntity(self, entopts);
    }
    // Entity access: `client.CrisprOfftargetCheck().list()` / `client.CrisprOfftargetCheck().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CrisprOfftargetCheck(entopts) {
        const self = this;
        return new CrisprOfftargetCheckEntity_1.CrisprOfftargetCheckEntity(self, entopts);
    }
    // Entity access: `client.CrossDimer().list()` / `client.CrossDimer().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CrossDimer(entopts) {
        const self = this;
        return new CrossDimerEntity_1.CrossDimerEntity(self, entopts);
    }
    // Entity access: `client.DnaMolarity().list()` / `client.DnaMolarity().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    DnaMolarity(entopts) {
        const self = this;
        return new DnaMolarityEntity_1.DnaMolarityEntity(self, entopts);
    }
    // Entity access: `client.DoubleDigest().list()` / `client.DoubleDigest().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    DoubleDigest(entopts) {
        const self = this;
        return new DoubleDigestEntity_1.DoubleDigestEntity(self, entopts);
    }
    // Entity access: `client.ExportEchoPicklist().list()` / `client.ExportEchoPicklist().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ExportEchoPicklist(entopts) {
        const self = this;
        return new ExportEchoPicklistEntity_1.ExportEchoPicklistEntity(self, entopts);
    }
    // Entity access: `client.ExportOpentronsProtocol().list()` / `client.ExportOpentronsProtocol().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ExportOpentronsProtocol(entopts) {
        const self = this;
        return new ExportOpentronsProtocolEntity_1.ExportOpentronsProtocolEntity(self, entopts);
    }
    // Entity access: `client.ExportPlateLayout().list()` / `client.ExportPlateLayout().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ExportPlateLayout(entopts) {
        const self = this;
        return new ExportPlateLayoutEntity_1.ExportPlateLayoutEntity(self, entopts);
    }
    // Entity access: `client.ExpressionHeatmapCluster().list()` / `client.ExpressionHeatmapCluster().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ExpressionHeatmapCluster(entopts) {
        const self = this;
        return new ExpressionHeatmapClusterEntity_1.ExpressionHeatmapClusterEntity(self, entopts);
    }
    // Entity access: `client.FastqQcReport().list()` / `client.FastqQcReport().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    FastqQcReport(entopts) {
        const self = this;
        return new FastqQcReportEntity_1.FastqQcReportEntity(self, entopts);
    }
    // Entity access: `client.FastqTrim().list()` / `client.FastqTrim().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    FastqTrim(entopts) {
        const self = this;
        return new FastqTrimEntity_1.FastqTrimEntity(self, entopts);
    }
    // Entity access: `client.FindOrf().list()` / `client.FindOrf().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    FindOrf(entopts) {
        const self = this;
        return new FindOrfEntity_1.FindOrfEntity(self, entopts);
    }
    // Entity access: `client.FormatSequence().list()` / `client.FormatSequence().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    FormatSequence(entopts) {
        const self = this;
        return new FormatSequenceEntity_1.FormatSequenceEntity(self, entopts);
    }
    // Entity access: `client.FunctionalEnrichment().list()` / `client.FunctionalEnrichment().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    FunctionalEnrichment(entopts) {
        const self = this;
        return new FunctionalEnrichmentEntity_1.FunctionalEnrichmentEntity(self, entopts);
    }
    // Entity access: `client.GcContent().list()` / `client.GcContent().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    GcContent(entopts) {
        const self = this;
        return new GcContentEntity_1.GcContentEntity(self, entopts);
    }
    // Entity access: `client.GeneDossier().list()` / `client.GeneDossier().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    GeneDossier(entopts) {
        const self = this;
        return new GeneDossierEntity_1.GeneDossierEntity(self, entopts);
    }
    // Entity access: `client.GeneExpression().list()` / `client.GeneExpression().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    GeneExpression(entopts) {
        const self = this;
        return new GeneExpressionEntity_1.GeneExpressionEntity(self, entopts);
    }
    // Entity access: `client.GeneModel().list()` / `client.GeneModel().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    GeneModel(entopts) {
        const self = this;
        return new GeneModelEntity_1.GeneModelEntity(self, entopts);
    }
    // Entity access: `client.GoldenGateFidelity().list()` / `client.GoldenGateFidelity().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    GoldenGateFidelity(entopts) {
        const self = this;
        return new GoldenGateFidelityEntity_1.GoldenGateFidelityEntity(self, entopts);
    }
    // Entity access: `client.HgvsConvert().list()` / `client.HgvsConvert().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    HgvsConvert(entopts) {
        const self = this;
        return new HgvsConvertEntity_1.HgvsConvertEntity(self, entopts);
    }
    // Entity access: `client.IdMapPoll().list()` / `client.IdMapPoll().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    IdMapPoll(entopts) {
        const self = this;
        return new IdMapPollEntity_1.IdMapPollEntity(self, entopts);
    }
    // Entity access: `client.IdMapSubmit().list()` / `client.IdMapSubmit().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    IdMapSubmit(entopts) {
        const self = this;
        return new IdMapSubmitEntity_1.IdMapSubmitEntity(self, entopts);
    }
    // Entity access: `client.InSilicoPcr().list()` / `client.InSilicoPcr().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    InSilicoPcr(entopts) {
        const self = this;
        return new InSilicoPcrEntity_1.InSilicoPcrEntity(self, entopts);
    }
    // Entity access: `client.KaspPrimerDesign().list()` / `client.KaspPrimerDesign().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    KaspPrimerDesign(entopts) {
        const self = this;
        return new KaspPrimerDesignEntity_1.KaspPrimerDesignEntity(self, entopts);
    }
    // Entity access: `client.ListTool().list()` / `client.ListTool().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ListTool(entopts) {
        const self = this;
        return new ListToolEntity_1.ListToolEntity(self, entopts);
    }
    // Entity access: `client.MeltingTemperature().list()` / `client.MeltingTemperature().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    MeltingTemperature(entopts) {
        const self = this;
        return new MeltingTemperatureEntity_1.MeltingTemperatureEntity(self, entopts);
    }
    // Entity access: `client.MotifFinder().list()` / `client.MotifFinder().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    MotifFinder(entopts) {
        const self = this;
        return new MotifFinderEntity_1.MotifFinderEntity(self, entopts);
    }
    // Entity access: `client.MultipleSequenceAlignment().list()` / `client.MultipleSequenceAlignment().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    MultipleSequenceAlignment(entopts) {
        const self = this;
        return new MultipleSequenceAlignmentEntity_1.MultipleSequenceAlignmentEntity(self, entopts);
    }
    // Entity access: `client.OligoAnalysi().list()` / `client.OligoAnalysi().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    OligoAnalysi(entopts) {
        const self = this;
        return new OligoAnalysiEntity_1.OligoAnalysiEntity(self, entopts);
    }
    // Entity access: `client.OrthologMap().list()` / `client.OrthologMap().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    OrthologMap(entopts) {
        const self = this;
        return new OrthologMapEntity_1.OrthologMapEntity(self, entopts);
    }
    // Entity access: `client.PairwiseAlignment().list()` / `client.PairwiseAlignment().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PairwiseAlignment(entopts) {
        const self = this;
        return new PairwiseAlignmentEntity_1.PairwiseAlignmentEntity(self, entopts);
    }
    // Entity access: `client.ParseGenbank().list()` / `client.ParseGenbank().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ParseGenbank(entopts) {
        const self = this;
        return new ParseGenbankEntity_1.ParseGenbankEntity(self, entopts);
    }
    // Entity access: `client.ParseSangerTrace().list()` / `client.ParseSangerTrace().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ParseSangerTrace(entopts) {
        const self = this;
        return new ParseSangerTraceEntity_1.ParseSangerTraceEntity(self, entopts);
    }
    // Entity access: `client.PlasmidAnnotate().list()` / `client.PlasmidAnnotate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PlasmidAnnotate(entopts) {
        const self = this;
        return new PlasmidAnnotateEntity_1.PlasmidAnnotateEntity(self, entopts);
    }
    // Entity access: `client.PlasmidDeepAnnotate().list()` / `client.PlasmidDeepAnnotate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PlasmidDeepAnnotate(entopts) {
        const self = this;
        return new PlasmidDeepAnnotateEntity_1.PlasmidDeepAnnotateEntity(self, entopts);
    }
    // Entity access: `client.PlasmidFullReport().list()` / `client.PlasmidFullReport().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PlasmidFullReport(entopts) {
        const self = this;
        return new PlasmidFullReportEntity_1.PlasmidFullReportEntity(self, entopts);
    }
    // Entity access: `client.PlasmidIdentify().list()` / `client.PlasmidIdentify().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PlasmidIdentify(entopts) {
        const self = this;
        return new PlasmidIdentifyEntity_1.PlasmidIdentifyEntity(self, entopts);
    }
    // Entity access: `client.PrimeEditingDesign().list()` / `client.PrimeEditingDesign().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PrimeEditingDesign(entopts) {
        const self = this;
        return new PrimeEditingDesignEntity_1.PrimeEditingDesignEntity(self, entopts);
    }
    // Entity access: `client.PrimeEditingTwinDesign().list()` / `client.PrimeEditingTwinDesign().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PrimeEditingTwinDesign(entopts) {
        const self = this;
        return new PrimeEditingTwinDesignEntity_1.PrimeEditingTwinDesignEntity(self, entopts);
    }
    // Entity access: `client.PrimerDesign().list()` / `client.PrimerDesign().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PrimerDesign(entopts) {
        const self = this;
        return new PrimerDesignEntity_1.PrimerDesignEntity(self, entopts);
    }
    // Entity access: `client.PrimerSpecificity().list()` / `client.PrimerSpecificity().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PrimerSpecificity(entopts) {
        const self = this;
        return new PrimerSpecificityEntity_1.PrimerSpecificityEntity(self, entopts);
    }
    // Entity access: `client.ProteaseDigestion().list()` / `client.ProteaseDigestion().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ProteaseDigestion(entopts) {
        const self = this;
        return new ProteaseDigestionEntity_1.ProteaseDigestionEntity(self, entopts);
    }
    // Entity access: `client.ProteinAnnotatePoll().list()` / `client.ProteinAnnotatePoll().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ProteinAnnotatePoll(entopts) {
        const self = this;
        return new ProteinAnnotatePollEntity_1.ProteinAnnotatePollEntity(self, entopts);
    }
    // Entity access: `client.ProteinAnnotateSubmit().list()` / `client.ProteinAnnotateSubmit().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ProteinAnnotateSubmit(entopts) {
        const self = this;
        return new ProteinAnnotateSubmitEntity_1.ProteinAnnotateSubmitEntity(self, entopts);
    }
    // Entity access: `client.ProteinHydrophobicity().list()` / `client.ProteinHydrophobicity().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ProteinHydrophobicity(entopts) {
        const self = this;
        return new ProteinHydrophobicityEntity_1.ProteinHydrophobicityEntity(self, entopts);
    }
    // Entity access: `client.ProteinProperty().list()` / `client.ProteinProperty().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ProteinProperty(entopts) {
        const self = this;
        return new ProteinPropertyEntity_1.ProteinPropertyEntity(self, entopts);
    }
    // Entity access: `client.RandomSequence().list()` / `client.RandomSequence().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    RandomSequence(entopts) {
        const self = this;
        return new RandomSequenceEntity_1.RandomSequenceEntity(self, entopts);
    }
    // Entity access: `client.RestrictionSite().list()` / `client.RestrictionSite().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    RestrictionSite(entopts) {
        const self = this;
        return new RestrictionSiteEntity_1.RestrictionSiteEntity(self, entopts);
    }
    // Entity access: `client.ReverseComplement().list()` / `client.ReverseComplement().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ReverseComplement(entopts) {
        const self = this;
        return new ReverseComplementEntity_1.ReverseComplementEntity(self, entopts);
    }
    // Entity access: `client.ReverseTranslate().list()` / `client.ReverseTranslate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ReverseTranslate(entopts) {
        const self = this;
        return new ReverseTranslateEntity_1.ReverseTranslateEntity(self, entopts);
    }
    // Entity access: `client.RnaFold().list()` / `client.RnaFold().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    RnaFold(entopts) {
        const self = this;
        return new RnaFoldEntity_1.RnaFoldEntity(self, entopts);
    }
    // Entity access: `client.SangerVsReference().list()` / `client.SangerVsReference().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SangerVsReference(entopts) {
        const self = this;
        return new SangerVsReferenceEntity_1.SangerVsReferenceEntity(self, entopts);
    }
    // Entity access: `client.SavePermalink().list()` / `client.SavePermalink().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SavePermalink(entopts) {
        const self = this;
        return new SavePermalinkEntity_1.SavePermalinkEntity(self, entopts);
    }
    // Entity access: `client.SeqfileStat().list()` / `client.SeqfileStat().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SeqfileStat(entopts) {
        const self = this;
        return new SeqfileStatEntity_1.SeqfileStatEntity(self, entopts);
    }
    // Entity access: `client.SequenceFetch().list()` / `client.SequenceFetch().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SequenceFetch(entopts) {
        const self = this;
        return new SequenceFetchEntity_1.SequenceFetchEntity(self, entopts);
    }
    // Entity access: `client.SequenceFormatConvert().list()` / `client.SequenceFormatConvert().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SequenceFormatConvert(entopts) {
        const self = this;
        return new SequenceFormatConvertEntity_1.SequenceFormatConvertEntity(self, entopts);
    }
    // Entity access: `client.SequenceReport().list()` / `client.SequenceReport().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SequenceReport(entopts) {
        const self = this;
        return new SequenceReportEntity_1.SequenceReportEntity(self, entopts);
    }
    // Entity access: `client.SequenceSearch().list()` / `client.SequenceSearch().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SequenceSearch(entopts) {
        const self = this;
        return new SequenceSearchEntity_1.SequenceSearchEntity(self, entopts);
    }
    // Entity access: `client.SequencingReadbackVerify().list()` / `client.SequencingReadbackVerify().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SequencingReadbackVerify(entopts) {
        const self = this;
        return new SequencingReadbackVerifyEntity_1.SequencingReadbackVerifyEntity(self, entopts);
    }
    // Entity access: `client.SessionCreate().list()` / `client.SessionCreate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SessionCreate(entopts) {
        const self = this;
        return new SessionCreateEntity_1.SessionCreateEntity(self, entopts);
    }
    // Entity access: `client.SessionGet().list()` / `client.SessionGet().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SessionGet(entopts) {
        const self = this;
        return new SessionGetEntity_1.SessionGetEntity(self, entopts);
    }
    // Entity access: `client.SessionRun().list()` / `client.SessionRun().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SessionRun(entopts) {
        const self = this;
        return new SessionRunEntity_1.SessionRunEntity(self, entopts);
    }
    // Entity access: `client.SessionSet().list()` / `client.SessionSet().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SessionSet(entopts) {
        const self = this;
        return new SessionSetEntity_1.SessionSetEntity(self, entopts);
    }
    // Entity access: `client.SirnaDesign().list()` / `client.SirnaDesign().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SirnaDesign(entopts) {
        const self = this;
        return new SirnaDesignEntity_1.SirnaDesignEntity(self, entopts);
    }
    // Entity access: `client.SiteDirectedMutagenesi().list()` / `client.SiteDirectedMutagenesi().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    SiteDirectedMutagenesi(entopts) {
        const self = this;
        return new SiteDirectedMutagenesiEntity_1.SiteDirectedMutagenesiEntity(self, entopts);
    }
    // Entity access: `client.Translate().list()` / `client.Translate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Translate(entopts) {
        const self = this;
        return new TranslateEntity_1.TranslateEntity(self, entopts);
    }
    // Entity access: `client.VariantAnnotate().list()` / `client.VariantAnnotate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    VariantAnnotate(entopts) {
        const self = this;
        return new VariantAnnotateEntity_1.VariantAnnotateEntity(self, entopts);
    }
    // Entity access: `client.VariantComparator().list()` / `client.VariantComparator().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    VariantComparator(entopts) {
        const self = this;
        return new VariantComparatorEntity_1.VariantComparatorEntity(self, entopts);
    }
    // Entity access: `client.VerifyAssembly().list()` / `client.VerifyAssembly().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    VerifyAssembly(entopts) {
        const self = this;
        return new VerifyAssemblyEntity_1.VerifyAssemblyEntity(self, entopts);
    }
    // Entity access: `client.VerifyConstruct().list()` / `client.VerifyConstruct().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    VerifyConstruct(entopts) {
        const self = this;
        return new VerifyConstructEntity_1.VerifyConstructEntity(self, entopts);
    }
    // Entity access: `client.VirtualGel().list()` / `client.VirtualGel().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    VirtualGel(entopts) {
        const self = this;
        return new VirtualGelEntity_1.VirtualGelEntity(self, entopts);
    }
    // Entity access: `client.VolcanoPlotData().list()` / `client.VolcanoPlotData().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    VolcanoPlotData(entopts) {
        const self = this;
        return new VolcanoPlotDataEntity_1.VolcanoPlotDataEntity(self, entopts);
    }
    // Entity access: `client.WebSearch().list()` / `client.WebSearch().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    WebSearch(entopts) {
        const self = this;
        return new WebSearchEntity_1.WebSearchEntity(self, entopts);
    }
    static test(testoptsarg, sdkoptsarg) {
        const struct = stdutil.struct;
        const setpath = struct.setpath;
        const getdef = struct.getdef;
        const clone = struct.clone;
        const setprop = struct.setprop;
        const sdkopts = getdef(clone(sdkoptsarg), {});
        const testopts = getdef(clone(testoptsarg), {});
        setprop(testopts, 'active', true);
        setpath(sdkopts, 'feature.test', testopts);
        const testsdk = new SeqbenchMcpSDK(sdkopts);
        testsdk._mode = 'test';
        return testsdk;
    }
    tester(testopts, sdkopts) {
        return SeqbenchMcpSDK.test(testopts, sdkopts);
    }
    toJSON() {
        return { name: 'SeqbenchMcp' };
    }
    toString() {
        return 'SeqbenchMcp ' + this._utility.struct.jsonify(this.toJSON());
    }
    [node_util_1.inspect.custom]() {
        return this.toString();
    }
}
exports.SeqbenchMcpSDK = SeqbenchMcpSDK;
const SDK = SeqbenchMcpSDK;
exports.SDK = SDK;
//# sourceMappingURL=SeqbenchMcpSDK.js.map