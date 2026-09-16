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
(0, node_test_1.describe)('ProteinHydrophobicityEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when SEQBENCH_MCP_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('SEQBENCH_MCP_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.SeqbenchMcpSDK.test();
        const ent = testsdk.ProteinHydrophobicity();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.SEQBENCH_MCP_TEST_LIVE;
        for (const op of ['create']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'protein_hydrophobicity.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "gate", "req": false, "short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.", "type": "`$ANY`", "index$": 0 }, { "active": true, "name": "ok", "req": true, "type": "`$ANY`", "index$": 1 }, { "active": true, "name": "provenance", "req": true, "type": "`$OBJECT`", "union": { "branches": 2, "count": 1, "depth": 2 }, "index$": 2 }, { "active": true, "name": "result", "req": true, "short": "Tool-specific output object.", "type": "`$OBJECT`", "index$": 3 }, { "active": true, "name": "scale", "req": false, "short": "Amino-acid scale.", "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "sequence", "req": true, "short": "Protein sequence (one-letter amino-acid codes; non-AA characters ignored).", "type": "`$STRING`", "index$": 5 }, { "active": true, "name": "tool", "req": true, "short": "The tool slug that ran.", "type": "`$STRING`", "index$": 6 }, { "active": true, "name": "window", "req": false, "short": "Sliding-window size (clamped to an odd number ≥ 1).", "type": "`$INTEGER`", "index$": 7 }], "name": "protein_hydrophobicity", "op": { "create": { "input": "data", "name": "create", "points": [{ "active": true, "args": {}, "contract": { "id": "POST /protein_hydrophobicity", "json": "{\"operationId\":\"protein_hydrophobicity\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"scale\":{\"default\":\"Kyte-Doolittle\",\"description\":\"Amino-acid scale. Kyte-Doolittle and Eisenberg are hydrophobicity; Hopp-Woods is hydrophilicity.\",\"enum\":[\"Kyte-Doolittle\",\"Hopp-Woods\",\"Eisenberg\"],\"type\":\"string\"},\"sequence\":{\"description\":\"Protein sequence (one-letter amino-acid codes; non-AA characters ignored).\",\"type\":\"string\"},\"window\":{\"default\":9,\"description\":\"Sliding-window size (clamped to an odd number ≥ 1).\",\"type\":\"integer\"}},\"required\":[\"sequence\"],\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard success envelope returned by every single-tool call.\",\"properties\":{\"gate\":{\"description\":\"Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. no targetTm).\",\"oneOf\":[{\"properties\":{\"checks\":{\"items\":{\"properties\":{\"id\":{\"type\":\"string\"},\"label\":{\"type\":\"string\"},\"message\":{\"type\":\"string\"},\"pass\":{\"type\":\"boolean\"},\"severity\":{\"enum\":[\"hard\",\"soft\"],\"type\":\"string\"},\"threshold\":{\"type\":\"string\"},\"value\":{\"type\":\"number\"}},\"required\":[\"id\",\"label\",\"severity\",\"pass\",\"message\"],\"type\":\"object\"},\"type\":\"array\"},\"notChecked\":{\"description\":\"Honest list of what this gate does NOT verify.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"pass\":{\"description\":\"True if and only if every hard check passed.\",\"type\":\"boolean\"}},\"required\":[\"pass\",\"checks\",\"notChecked\"],\"type\":\"object\"},{\"type\":\"null\"}]},\"ok\":{\"const\":true},\"provenance\":{\"properties\":{\"apiVersion\":{\"example\":\"1.1.0\",\"type\":\"string\"},\"generatedAt\":{\"description\":\"ISO 8601 timestamp of when the result was generated.\",\"format\":\"date-time\",\"type\":\"string\"},\"tool\":{\"description\":\"Tool slug, or an array of slugs (one per step) for a workflow.\",\"oneOf\":[{\"type\":\"string\"},{\"items\":{\"type\":\"string\"},\"type\":\"array\"}]}},\"required\":[\"apiVersion\",\"tool\",\"generatedAt\"],\"type\":\"object\"},\"result\":{\"additionalProperties\":true,\"description\":\"Tool-specific output object. Its shape depends on the tool — call `GET /{tool}` or see the docs for each tool's fields.\",\"type\":\"object\"},\"tool\":{\"description\":\"The tool slug that ran.\",\"type\":\"string\"}},\"required\":[\"ok\",\"tool\",\"result\",\"provenance\"],\"type\":\"object\"}}},\"description\":\"Tool ran successfully.\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Invalid argument or the tool could not run.\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Unknown tool.\"}},\"security\":[{}],\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "POST", "orig": "/protein_hydrophobicity", "segments": [{ "lit": "protein_hydrophobicity" }], "select": {}, "transform": { "req": { "scale": "`reqdata.scale`", "sequence": "`reqdata.sequence`", "window": "`reqdata.window`" }, "res": "`body`" }, "index$": 0 }], "key$": "create" } }, "relations": { "ancestors": [] }, "key$": "protein_hydrophobicity", "name__orig": "protein_hydrophobicity", "Name": "ProteinHydrophobicity", "name_": "protein_hydrophobicity", "name-": "protein-hydrophobicity", "NAME": "PROTEIN_HYDROPHOBICITY", "index$": 56 }, { "active": true, "entity": "protein_hydrophobicity", "key$": "BasicProteinHydrophobicityFlow", "kind": "basic", "name": "BasicProteinHydrophobicityFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": { "ref": "protein_hydrophobicity_ref01" }, "match": {}, "op": "create", "spec": [], "valid": [], "index$": 0 }] }, 'ProteinHydrophobicity');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const protein_hydrophobicity_ref01_ent = client.ProteinHydrophobicity();
        let protein_hydrophobicity_ref01_data = setup.data.new.protein_hydrophobicity['protein_hydrophobicity_ref01'];
        protein_hydrophobicity_ref01_data = (await protein_hydrophobicity_ref01_ent.create(protein_hydrophobicity_ref01_data)).data();
        (0, node_assert_1.default)(null != protein_hydrophobicity_ref01_data);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/protein_hydrophobicity/ProteinHydrophobicityTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.SeqbenchMcpSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['protein_hydrophobicity01', 'protein_hydrophobicity02', 'protein_hydrophobicity03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'SEQBENCH_MCP_TEST_PROTEIN_HYDROPHOBICITY_ENTID': idmap,
        'SEQBENCH_MCP_TEST_LIVE': 'FALSE',
        'SEQBENCH_MCP_TEST_EXPLAIN': 'FALSE',
        'SEQBENCH_MCP_APIKEY': '',
    });
    idmap = env['SEQBENCH_MCP_TEST_PROTEIN_HYDROPHOBICITY_ENTID'];
    const live = 'TRUE' === env.SEQBENCH_MCP_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['SEQBENCH_MCP_TEST_PROTEIN_HYDROPHOBICITY_ENTID'];
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
//# sourceMappingURL=ProteinHydrophobicityEntity.test.js.map