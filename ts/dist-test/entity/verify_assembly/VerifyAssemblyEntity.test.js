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
(0, node_test_1.describe)('VerifyAssemblyEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when SEQBENCH_MCP_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('SEQBENCH_MCP_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.SeqbenchMcpSDK.test();
        const ent = testsdk.VerifyAssembly();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.SEQBENCH_MCP_TEST_LIVE;
        for (const op of ['create']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'verify_assembly.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "armTmTarget", "req": false, "short": "Target annealing Tm (°C) for primer arms.", "type": "`$NUMBER`", "index$": 0 }, { "active": true, "name": "circular", "req": false, "short": "Treat the product/claimed construct as circular (most plasmids are).", "type": "`$BOOLEAN`", "index$": 1 }, { "active": true, "name": "claimedConstruct", "req": true, "short": "The sequence you claim you ended up with.", "type": "`$STRING`", "index$": 2 }, { "active": true, "name": "coding", "req": false, "short": "Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence.", "type": "`$BOOLEAN`", "index$": 3 }, { "active": true, "name": "enzyme", "req": false, "short": "Type IIS enzyme for Golden Gate.", "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "enzyme3", "req": false, "short": "3′ enzyme (restriction method).", "type": "`$STRING`", "index$": 5 }, { "active": true, "name": "enzyme5", "req": false, "short": "5′ enzyme (restriction method).", "type": "`$STRING`", "index$": 6 }, { "active": true, "name": "fragmentPcrs", "req": false, "short": "Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead.", "type": "`$ARRAY`", "index$": 7 }, { "active": true, "name": "fragments", "req": false, "short": "Fragments (5′→3′), assembled head-to-tail (gibson/goldengate).", "type": "`$ARRAY`", "index$": 8 }, { "active": true, "name": "frameStart", "req": false, "short": "1-based reading-frame start on claimedConstruct, used when coding is true.", "type": "`$INTEGER`", "index$": 9 }, { "active": true, "name": "gate", "req": false, "short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.", "type": "`$ANY`", "index$": 10 }, { "active": true, "name": "insert", "req": false, "short": "Insert sequence (restriction method).", "type": "`$STRING`", "index$": 11 }, { "active": true, "name": "insertPcr", "req": false, "short": "Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}.", "type": "`$OBJECT`", "index$": 12 }, { "active": true, "name": "method", "req": true, "short": "Assembly method used.", "type": "`$STRING`", "index$": 13 }, { "active": true, "name": "names", "req": false, "short": "Optional labels for each fragment.", "type": "`$ARRAY`", "index$": 14 }, { "active": true, "name": "ok", "req": true, "type": "`$ANY`", "index$": 15 }, { "active": true, "name": "overlapLen", "req": false, "short": "Gibson homology-arm length (bp).", "type": "`$INTEGER`", "index$": 16 }, { "active": true, "name": "provenance", "req": true, "type": "`$OBJECT`", "union": { "branches": 2, "count": 1, "depth": 2 }, "index$": 17 }, { "active": true, "name": "result", "req": true, "short": "Tool-specific output object.", "type": "`$OBJECT`", "index$": 18 }, { "active": true, "name": "tool", "req": true, "short": "The tool slug that ran.", "type": "`$STRING`", "index$": 19 }, { "active": true, "name": "vector", "req": false, "short": "Vector sequence (restriction method).", "type": "`$STRING`", "index$": 20 }, { "active": true, "name": "vectorPcr", "req": false, "short": "Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}.", "type": "`$OBJECT`", "index$": 21 }], "name": "verify_assembly", "op": { "create": { "input": "data", "name": "create", "points": [{ "active": true, "args": {}, "contract": { "id": "POST /verify_assembly", "json": "{\"operationId\":\"verify_assembly\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"armTmTarget\":{\"default\":60,\"description\":\"Target annealing Tm (°C) for primer arms.\",\"type\":\"number\"},\"circular\":{\"default\":true,\"description\":\"Treat the product/claimed construct as circular (most plasmids are).\",\"type\":\"boolean\"},\"claimedConstruct\":{\"description\":\"The sequence you claim you ended up with.\",\"type\":\"string\"},\"coding\":{\"default\":false,\"description\":\"Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence.\",\"type\":\"boolean\"},\"enzyme\":{\"default\":\"BsaI\",\"description\":\"Type IIS enzyme for Golden Gate.\",\"type\":\"string\"},\"enzyme3\":{\"default\":\"BamHI\",\"description\":\"3′ enzyme (restriction method).\",\"type\":\"string\"},\"enzyme5\":{\"default\":\"EcoRI\",\"description\":\"5′ enzyme (restriction method).\",\"type\":\"string\"},\"fragmentPcrs\":{\"description\":\"Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead.\",\"items\":{\"type\":\"object\"},\"type\":\"array\"},\"fragments\":{\"description\":\"Fragments (5′→3′), assembled head-to-tail (gibson/goldengate). Use \\\"\\\" as a placeholder for any fragment supplied instead via the matching fragmentPcrs[i].\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"frameStart\":{\"default\":1,\"description\":\"1-based reading-frame start on claimedConstruct, used when coding is true.\",\"type\":\"integer\"},\"insert\":{\"description\":\"Insert sequence (restriction method). Omit if insertPcr is given.\",\"type\":\"string\"},\"insertPcr\":{\"description\":\"Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}.\",\"type\":\"object\"},\"method\":{\"description\":\"Assembly method used.\",\"enum\":[\"restriction\",\"gibson\",\"goldengate\"],\"type\":\"string\"},\"names\":{\"description\":\"Optional labels for each fragment.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"overlapLen\":{\"default\":20,\"description\":\"Gibson homology-arm length (bp).\",\"type\":\"integer\"},\"vector\":{\"description\":\"Vector sequence (restriction method). Omit if vectorPcr is given.\",\"type\":\"string\"},\"vectorPcr\":{\"description\":\"Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}.\",\"type\":\"object\"}},\"required\":[\"method\",\"claimedConstruct\"],\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard success envelope returned by every single-tool call.\",\"properties\":{\"gate\":{\"description\":\"Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. no targetTm).\",\"oneOf\":[{\"properties\":{\"checks\":{\"items\":{\"properties\":{\"id\":{\"type\":\"string\"},\"label\":{\"type\":\"string\"},\"message\":{\"type\":\"string\"},\"pass\":{\"type\":\"boolean\"},\"severity\":{\"enum\":[\"hard\",\"soft\"],\"type\":\"string\"},\"threshold\":{\"type\":\"string\"},\"value\":{\"type\":\"number\"}},\"required\":[\"id\",\"label\",\"severity\",\"pass\",\"message\"],\"type\":\"object\"},\"type\":\"array\"},\"notChecked\":{\"description\":\"Honest list of what this gate does NOT verify.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"pass\":{\"description\":\"True if and only if every hard check passed.\",\"type\":\"boolean\"}},\"required\":[\"pass\",\"checks\",\"notChecked\"],\"type\":\"object\"},{\"type\":\"null\"}]},\"ok\":{\"const\":true},\"provenance\":{\"properties\":{\"apiVersion\":{\"example\":\"1.1.0\",\"type\":\"string\"},\"generatedAt\":{\"description\":\"ISO 8601 timestamp of when the result was generated.\",\"format\":\"date-time\",\"type\":\"string\"},\"tool\":{\"description\":\"Tool slug, or an array of slugs (one per step) for a workflow.\",\"oneOf\":[{\"type\":\"string\"},{\"items\":{\"type\":\"string\"},\"type\":\"array\"}]}},\"required\":[\"apiVersion\",\"tool\",\"generatedAt\"],\"type\":\"object\"},\"result\":{\"additionalProperties\":true,\"description\":\"Tool-specific output object. Its shape depends on the tool — call `GET /{tool}` or see the docs for each tool's fields.\",\"type\":\"object\"},\"tool\":{\"description\":\"The tool slug that ran.\",\"type\":\"string\"}},\"required\":[\"ok\",\"tool\",\"result\",\"provenance\"],\"type\":\"object\"}}},\"description\":\"Tool ran successfully.\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Invalid argument or the tool could not run.\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Unknown tool.\"}},\"security\":[{}],\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "POST", "orig": "/verify_assembly", "segments": [{ "lit": "verify_assembly" }], "select": {}, "transform": { "req": { "armTmTarget": "`reqdata.arm_tm_target`", "circular": "`reqdata.circular`", "claimedConstruct": "`reqdata.claimed_construct`", "coding": "`reqdata.coding`", "enzyme": "`reqdata.enzyme`", "enzyme3": "`reqdata.enzyme3`", "enzyme5": "`reqdata.enzyme5`", "fragmentPcrs": "`reqdata.fragment_pcr`", "fragments": "`reqdata.fragment`", "frameStart": "`reqdata.frame_start`", "insert": "`reqdata.insert`", "insertPcr": "`reqdata.insert_pcr`", "method": "`reqdata.method`", "names": "`reqdata.name`", "overlapLen": "`reqdata.overlap_len`", "vector": "`reqdata.vector`", "vectorPcr": "`reqdata.vector_pcr`" }, "res": "`body`" }, "index$": 0 }], "key$": "create" } }, "relations": { "ancestors": [] }, "key$": "verify_assembly", "name__orig": "verify_assembly", "Name": "VerifyAssembly", "name_": "verify_assembly", "name-": "verify-assembly", "NAME": "VERIFY_ASSEMBLY", "index$": 80 }, { "active": true, "entity": "verify_assembly", "key$": "BasicVerifyAssemblyFlow", "kind": "basic", "name": "BasicVerifyAssemblyFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": { "ref": "verify_assembly_ref01" }, "match": {}, "op": "create", "spec": [], "valid": [], "index$": 0 }] }, 'VerifyAssembly');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const verify_assembly_ref01_ent = client.VerifyAssembly();
        let verify_assembly_ref01_data = setup.data.new.verify_assembly['verify_assembly_ref01'];
        verify_assembly_ref01_data = (await verify_assembly_ref01_ent.create(verify_assembly_ref01_data)).data();
        (0, node_assert_1.default)(null != verify_assembly_ref01_data);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/verify_assembly/VerifyAssemblyTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.SeqbenchMcpSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['verify_assembly01', 'verify_assembly02', 'verify_assembly03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'SEQBENCH_MCP_TEST_VERIFY_ASSEMBLY_ENTID': idmap,
        'SEQBENCH_MCP_TEST_LIVE': 'FALSE',
        'SEQBENCH_MCP_TEST_EXPLAIN': 'FALSE',
        'SEQBENCH_MCP_APIKEY': '',
    });
    idmap = env['SEQBENCH_MCP_TEST_VERIFY_ASSEMBLY_ENTID'];
    const live = 'TRUE' === env.SEQBENCH_MCP_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['SEQBENCH_MCP_TEST_VERIFY_ASSEMBLY_ENTID'];
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
//# sourceMappingURL=VerifyAssemblyEntity.test.js.map