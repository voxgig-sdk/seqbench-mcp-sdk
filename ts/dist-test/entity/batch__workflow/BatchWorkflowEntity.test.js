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
(0, node_test_1.describe)('BatchWorkflowEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when SEQBENCH_MCP_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('SEQBENCH_MCP_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.SeqbenchMcpSDK.test();
        const ent = testsdk.BatchWorkflow();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.SEQBENCH_MCP_TEST_LIVE;
        for (const op of ['create', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'batch__workflow.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "capped", "req": true, "type": "`$BOOLEAN`", "index$": 0 }, { "active": true, "name": "columns", "req": true, "short": "Flattened \"<step>·<tool>·<key>\" column headers.", "type": "`$ARRAY`", "index$": 1 }, { "active": true, "name": "count", "req": true, "type": "`$INTEGER`", "index$": 2 }, { "active": true, "name": "errors", "req": true, "type": "`$INTEGER`", "index$": 3 }, { "active": true, "name": "input", "req": true, "short": "Multi-FASTA text or one sequence per line.", "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "limit", "req": true, "short": "Maximum records per call (200).", "type": "`$INTEGER`", "index$": 5 }, { "active": true, "name": "provenance", "req": true, "type": "`$OBJECT`", "union": { "branches": 2, "count": 1, "depth": 2 }, "index$": 6 }, { "active": true, "name": "rows", "req": true, "type": "`$ARRAY`", "index$": 7 }, { "active": true, "name": "steps", "req": true, "type": "`$ARRAY`", "index$": 8 }], "name": "batch__workflow", "op": { "create": { "input": "data", "name": "create", "points": [{ "active": true, "args": {}, "contract": { "id": "POST /workflow", "json": "{\"operationId\":\"runWorkflow\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"example\":{\"input\":\">seq1\\nATGGCCTGA\",\"steps\":[{\"tool\":\"reverse_complement\"},{\"args\":{\"frame\":1},\"tool\":\"translate\"}]},\"schema\":{\"properties\":{\"input\":{\"description\":\"Multi-FASTA text or one sequence per line.\",\"type\":\"string\"},\"steps\":{\"items\":{\"properties\":{\"args\":{\"additionalProperties\":true,\"type\":\"object\"},\"from\":{\"description\":\"Source of this step's input. Defaults to the previous step.\",\"oneOf\":[{\"description\":\"0-based index of an earlier step.\",\"type\":\"integer\"},{\"const\":\"initial\",\"description\":\"The original input record.\"}]},\"tool\":{\"description\":\"A pipeline-capable tool slug.\",\"type\":\"string\"}},\"required\":[\"tool\"],\"type\":\"object\"},\"maxItems\":8,\"minItems\":1,\"type\":\"array\"}},\"required\":[\"steps\",\"input\"],\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"ok\":{\"const\":true},\"result\":{\"properties\":{\"capped\":{\"type\":\"boolean\"},\"columns\":{\"description\":\"Flattened \\\"<step>·<tool>·<key>\\\" column headers.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"count\":{\"type\":\"integer\"},\"errors\":{\"type\":\"integer\"},\"limit\":{\"description\":\"Maximum records per call (200).\",\"type\":\"integer\"},\"provenance\":{\"properties\":{\"apiVersion\":{\"example\":\"1.1.0\",\"type\":\"string\"},\"generatedAt\":{\"description\":\"ISO 8601 timestamp of when the result was generated.\",\"format\":\"date-time\",\"type\":\"string\"},\"tool\":{\"description\":\"Tool slug, or an array of slugs (one per step) for a workflow.\",\"oneOf\":[{\"type\":\"string\"},{\"items\":{\"type\":\"string\"},\"type\":\"array\"}]}},\"required\":[\"apiVersion\",\"tool\",\"generatedAt\"],\"type\":\"object\"},\"rows\":{\"items\":{\"type\":\"object\"},\"type\":\"array\"},\"steps\":{\"items\":{\"properties\":{\"title\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"required\":[\"steps\",\"count\",\"capped\",\"limit\",\"columns\",\"rows\",\"errors\",\"provenance\"],\"type\":\"object\"}},\"required\":[\"ok\",\"result\"],\"type\":\"object\"}}},\"description\":\"Workflow completed (individual steps may carry per-record errors).\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Error shape returned by the batch and workflow endpoints.\",\"properties\":{\"error\":{\"type\":\"string\"},\"ok\":{\"const\":false}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Invalid request (bad JSON, missing steps/input, or a step is invalid).\"}},\"security\":[{}],\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "POST", "orig": "/workflow", "segments": [{ "lit": "workflow" }], "select": {}, "transform": { "req": "`reqdata`", "res": "`body.result`" }, "index$": 0 }], "key$": "create" }, "load": { "input": "data", "name": "load", "points": [{ "active": true, "args": {}, "contract": { "id": "GET /workflow", "json": "{\"operationId\":\"workflowInfo\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"type\":\"object\"}}},\"description\":\"Pipeline-capable tool list and limits.\"}},\"security\":[{}],\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/workflow", "segments": [{ "lit": "workflow" }], "select": {}, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "batch__workflow", "name__orig": "batch__workflow", "Name": "BatchWorkflow", "name_": "batch_workflow", "name-": "batch-workflow", "NAME": "BATCH__WORKFLOW", "index$": 4 }, { "active": true, "entity": "batch__workflow", "key$": "BasicBatchWorkflowFlow", "kind": "basic", "name": "BasicBatchWorkflowFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": { "ref": "batch__workflow_ref01" }, "match": {}, "op": "create", "spec": [], "valid": [], "index$": 0 }, { "active": true, "data": {}, "input": { "ref": "batch__workflow_ref01", "srcdatavar": "batch__workflow_ref01_data", "suffix": "_dt0" }, "match": {}, "op": "load", "spec": [], "valid": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-batch__workflow_ref01" } }], "index$": 1 }] }, 'BatchWorkflow');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const batch__workflow_ref01_ent = client.BatchWorkflow();
        let batch__workflow_ref01_data = setup.data.new.batch__workflow['batch__workflow_ref01'];
        batch__workflow_ref01_data = (await batch__workflow_ref01_ent.create(batch__workflow_ref01_data)).data();
        (0, node_assert_1.default)(null != batch__workflow_ref01_data);
        // LOAD
        const batch__workflow_ref01_match_dt0 = {};
        const batch__workflow_ref01_data_dt0 = (await batch__workflow_ref01_ent.load(batch__workflow_ref01_match_dt0)).data();
        (0, node_assert_1.default)(null != batch__workflow_ref01_data_dt0);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/batch__workflow/BatchWorkflowTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.SeqbenchMcpSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['batch__workflow01', 'batch__workflow02', 'batch__workflow03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'SEQBENCH_MCP_TEST_BATCH_WORKFLOW_ENTID': idmap,
        'SEQBENCH_MCP_TEST_LIVE': 'FALSE',
        'SEQBENCH_MCP_TEST_EXPLAIN': 'FALSE',
        'SEQBENCH_MCP_APIKEY': '',
    });
    idmap = env['SEQBENCH_MCP_TEST_BATCH_WORKFLOW_ENTID'];
    const live = 'TRUE' === env.SEQBENCH_MCP_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['SEQBENCH_MCP_TEST_BATCH_WORKFLOW_ENTID'];
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
//# sourceMappingURL=BatchWorkflowEntity.test.js.map