"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('CrisprHdrDonorEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when SEQBENCH_MCP_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('SEQBENCH_MCP_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.SeqbenchMcpSDK.test();
        const ent = testsdk.CrisprHdrDonor();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.SEQBENCH_MCP_TEST_LIVE;
        for (const op of ['create']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'crispr_hdr_donor.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "armLength", "req": false, "short": "Homology arm length (bp) on each side.", "type": "`$INTEGER`", "index$": 0 }, { "active": true, "name": "blockPam", "req": false, "short": "When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut.", "type": "`$BOOLEAN`", "index$": 1 }, { "active": true, "name": "designGenotypingPrimers", "req": false, "short": "Also design a primer pair (on the original targetSequence) whose product spans the edit site.", "type": "`$BOOLEAN`", "index$": 2 }, { "active": true, "name": "editEnd", "req": false, "short": "1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed.", "type": "`$INTEGER`", "index$": 3 }, { "active": true, "name": "editStart", "req": false, "short": "1-based start of the region being replaced.", "type": "`$INTEGER`", "index$": 4 }, { "active": true, "name": "frameStart", "req": false, "short": "Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible.", "type": "`$INTEGER`", "index$": 5 }, { "active": true, "name": "gate", "req": false, "short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.", "type": "`$ANY`", "index$": 6 }, { "active": true, "name": "guideEnd", "req": false, "short": "1-based forward-strand end of the guide's protospacer.", "type": "`$INTEGER`", "index$": 7 }, { "active": true, "name": "guideStart", "req": false, "short": "1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site).", "type": "`$INTEGER`", "index$": 8 }, { "active": true, "name": "guideStrand", "req": false, "short": "Strand the guide's protospacer is on.", "type": "`$STRING`", "index$": 9 }, { "active": true, "name": "nuclease", "req": false, "short": "Needed only when deriving the cut site from guideStart/guideEnd/guideStrand.", "type": "`$STRING`", "index$": 10 }, { "active": true, "name": "ok", "req": true, "type": "`$ANY`", "index$": 11 }, { "active": true, "name": "provenance", "req": true, "type": "`$OBJECT`", "union": { "branches": 2, "count": 1, "depth": 2 }, "index$": 12 }, { "active": true, "name": "replacement", "req": true, "short": "Sequence to insert/substitute (\"\" for a pure deletion).", "type": "`$STRING`", "index$": 13 }, { "active": true, "name": "result", "req": true, "short": "Tool-specific output object.", "type": "`$OBJECT`", "index$": 14 }, { "active": true, "name": "targetSequence", "req": true, "short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).", "type": "`$STRING`", "index$": 15 }, { "active": true, "name": "tool", "req": true, "short": "The tool slug that ran.", "type": "`$STRING`", "index$": 16 }], "name": "crispr_hdr_donor", "op": { "create": { "input": "data", "name": "create", "points": [{ "active": true, "args": {}, "contract": { "id": "POST /crispr_hdr_donor", "json": "{\"operationId\":\"crispr_hdr_donor\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"armLength\":{\"default\":500,\"description\":\"Homology arm length (bp) on each side. Use ~30–60 for an ssODN donor, ~500–1000 for a dsDNA donor plasmid.\",\"type\":\"integer\"},\"blockPam\":{\"default\":true,\"description\":\"When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut.\",\"type\":\"boolean\"},\"designGenotypingPrimers\":{\"default\":true,\"description\":\"Also design a primer pair (on the original targetSequence) whose product spans the edit site.\",\"type\":\"boolean\"},\"editEnd\":{\"description\":\"1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed. Omit to derive from the guide's cut site.\",\"type\":\"integer\"},\"editStart\":{\"description\":\"1-based start of the region being replaced. Omit to derive from guideStart/guideEnd/guideStrand instead.\",\"type\":\"integer\"},\"frameStart\":{\"description\":\"Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible.\",\"type\":\"integer\"},\"guideEnd\":{\"description\":\"1-based forward-strand end of the guide's protospacer.\",\"type\":\"integer\"},\"guideStart\":{\"description\":\"1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site).\",\"type\":\"integer\"},\"guideStrand\":{\"description\":\"Strand the guide's protospacer is on.\",\"enum\":[\"+\",\"-\"],\"type\":\"string\"},\"nuclease\":{\"default\":\"spcas9\",\"description\":\"Needed only when deriving the cut site from guideStart/guideEnd/guideStrand.\",\"enum\":[\"spcas9\",\"spcas9ng\",\"sacas9\",\"cas12a\"],\"type\":\"string\"},\"replacement\":{\"default\":\"\",\"description\":\"Sequence to insert/substitute (\\\"\\\" for a pure deletion).\",\"type\":\"string\"},\"targetSequence\":{\"description\":\"Nucleotide sequence (raw or FASTA; IUPAC accepted).\",\"type\":\"string\"}},\"required\":[\"targetSequence\",\"replacement\"],\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard success envelope returned by every single-tool call.\",\"properties\":{\"gate\":{\"description\":\"Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. no targetTm).\",\"oneOf\":[{\"properties\":{\"checks\":{\"items\":{\"properties\":{\"id\":{\"type\":\"string\"},\"label\":{\"type\":\"string\"},\"message\":{\"type\":\"string\"},\"pass\":{\"type\":\"boolean\"},\"severity\":{\"enum\":[\"hard\",\"soft\"],\"type\":\"string\"},\"threshold\":{\"type\":\"string\"},\"value\":{\"type\":\"number\"}},\"required\":[\"id\",\"label\",\"severity\",\"pass\",\"message\"],\"type\":\"object\"},\"type\":\"array\"},\"notChecked\":{\"description\":\"Honest list of what this gate does NOT verify.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"pass\":{\"description\":\"True if and only if every hard check passed.\",\"type\":\"boolean\"}},\"required\":[\"pass\",\"checks\",\"notChecked\"],\"type\":\"object\"},{\"type\":\"null\"}]},\"ok\":{\"const\":true},\"provenance\":{\"properties\":{\"apiVersion\":{\"example\":\"1.1.0\",\"type\":\"string\"},\"generatedAt\":{\"description\":\"ISO 8601 timestamp of when the result was generated.\",\"format\":\"date-time\",\"type\":\"string\"},\"tool\":{\"description\":\"Tool slug, or an array of slugs (one per step) for a workflow.\",\"oneOf\":[{\"type\":\"string\"},{\"items\":{\"type\":\"string\"},\"type\":\"array\"}]}},\"required\":[\"apiVersion\",\"tool\",\"generatedAt\"],\"type\":\"object\"},\"result\":{\"additionalProperties\":true,\"description\":\"Tool-specific output object. Its shape depends on the tool — call `GET /{tool}` or see the docs for each tool's fields.\",\"type\":\"object\"},\"tool\":{\"description\":\"The tool slug that ran.\",\"type\":\"string\"}},\"required\":[\"ok\",\"tool\",\"result\",\"provenance\"],\"type\":\"object\"}}},\"description\":\"Tool ran successfully.\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Invalid argument or the tool could not run.\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Unknown tool.\"}},\"security\":[{}],\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "POST", "orig": "/crispr_hdr_donor", "segments": [{ "lit": "crispr_hdr_donor" }], "select": {}, "transform": { "req": { "armLength": "`reqdata.arm_length`", "blockPam": "`reqdata.block_pam`", "designGenotypingPrimers": "`reqdata.design_genotyping_primer`", "editEnd": "`reqdata.edit_end`", "editStart": "`reqdata.edit_start`", "frameStart": "`reqdata.frame_start`", "guideEnd": "`reqdata.guide_end`", "guideStart": "`reqdata.guide_start`", "guideStrand": "`reqdata.guide_strand`", "nuclease": "`reqdata.nuclease`", "replacement": "`reqdata.replacement`", "targetSequence": "`reqdata.target_sequence`" }, "res": "`body`" }, "index$": 0 }], "key$": "create" } }, "relations": { "ancestors": [] }, "key$": "crispr_hdr_donor", "name__orig": "crispr_hdr_donor", "Name": "CrisprHdrDonor", "name_": "crispr_hdr_donor", "name-": "crispr-hdr-donor", "NAME": "CRISPR_HDR_DONOR", "index$": 12 }, { "active": true, "entity": "crispr_hdr_donor", "key$": "BasicCrisprHdrDonorFlow", "kind": "basic", "name": "BasicCrisprHdrDonorFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": { "ref": "crispr_hdr_donor_ref01" }, "match": {}, "op": "create", "spec": [], "valid": [], "index$": 0 }] }, 'CrisprHdrDonor');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const crispr_hdr_donor_ref01_ent = client.CrisprHdrDonor();
        let crispr_hdr_donor_ref01_data = setup.data.new.crispr_hdr_donor['crispr_hdr_donor_ref01'];
        crispr_hdr_donor_ref01_data = (await crispr_hdr_donor_ref01_ent.create(crispr_hdr_donor_ref01_data)).data();
        (0, node_assert_1.default)(null != crispr_hdr_donor_ref01_data);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/crispr_hdr_donor/CrisprHdrDonorTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.SeqbenchMcpSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['crispr_hdr_donor01', 'crispr_hdr_donor02', 'crispr_hdr_donor03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'SEQBENCH_MCP_TEST_CRISPR_HDR_DONOR_ENTID': idmap,
        'SEQBENCH_MCP_TEST_LIVE': 'FALSE',
        'SEQBENCH_MCP_TEST_EXPLAIN': 'FALSE',
        'SEQBENCH_MCP_APIKEY': '',
    });
    idmap = env['SEQBENCH_MCP_TEST_CRISPR_HDR_DONOR_ENTID'];
    const live = 'TRUE' === env.SEQBENCH_MCP_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['SEQBENCH_MCP_TEST_CRISPR_HDR_DONOR_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.SeqbenchMcpSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {
                apikey: env.SEQBENCH_MCP_APIKEY,
            },
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.SEQBENCH_MCP_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=CrisprHdrDonorEntity.test.js.map