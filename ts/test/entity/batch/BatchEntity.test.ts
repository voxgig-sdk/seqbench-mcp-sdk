

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { SeqbenchMcpSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('BatchEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when SEQBENCH_MCP_TEST_LIVE=TRUE.
  afterEach(liveDelay('SEQBENCH_MCP_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = SeqbenchMcpSDK.test()
    const ent = testsdk.Batch()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.SEQBENCH_MCP_TEST_LIVE
    for (const op of ['create', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'batch.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"args","req":false,"short":"Shared tool arguments applied to every record.","type":"`$OBJECT`","index$":0},{"active":true,"name":"capped","req":true,"short":"True if input exceeded the record limit.","type":"`$BOOLEAN`","index$":1},{"active":true,"name":"columns","req":true,"type":"`$ARRAY`","index$":2},{"active":true,"name":"count","req":true,"type":"`$INTEGER`","index$":3},{"active":true,"name":"errors","req":true,"type":"`$INTEGER`","index$":4},{"active":true,"name":"input","req":true,"short":"Multi-FASTA text or one sequence per line (max ~2,000,000 chars).","type":"`$STRING`","index$":5},{"active":true,"name":"limit","req":true,"short":"Maximum records per call (500).","type":"`$INTEGER`","index$":6},{"active":true,"name":"provenance","req":true,"type":"`$OBJECT`","union":{"branches":2,"count":1,"depth":2},"index$":7},{"active":true,"name":"rows","req":true,"type":"`$ARRAY`","index$":8},{"active":true,"name":"tool","req":true,"short":"A batchable tool slug (see `GET /batch`).","type":"`$STRING`","index$":9}],"name":"batch","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{},"contract":{"id":"POST /batch","json":"{\"operationId\":\"runBatch\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"example\":{\"input\":\">seq1\\nATGGCGCGCTAA\\n>seq2\\nTTTTAAAACCCC\",\"tool\":\"gc_content\"},\"schema\":{\"properties\":{\"args\":{\"additionalProperties\":true,\"description\":\"Shared tool arguments applied to every record.\",\"type\":\"object\"},\"input\":{\"description\":\"Multi-FASTA text or one sequence per line (max ~2,000,000 chars).\",\"type\":\"string\"},\"tool\":{\"description\":\"A batchable tool slug (see `GET /batch`).\",\"type\":\"string\"}},\"required\":[\"tool\",\"input\"],\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"ok\":{\"const\":true},\"result\":{\"properties\":{\"capped\":{\"description\":\"True if input exceeded the record limit.\",\"type\":\"boolean\"},\"columns\":{\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"count\":{\"type\":\"integer\"},\"errors\":{\"type\":\"integer\"},\"limit\":{\"description\":\"Maximum records per call (500).\",\"type\":\"integer\"},\"provenance\":{\"properties\":{\"apiVersion\":{\"example\":\"1.1.0\",\"type\":\"string\"},\"generatedAt\":{\"description\":\"ISO 8601 timestamp of when the result was generated.\",\"format\":\"date-time\",\"type\":\"string\"},\"tool\":{\"description\":\"Tool slug, or an array of slugs (one per step) for a workflow.\",\"oneOf\":[{\"type\":\"string\"},{\"items\":{\"type\":\"string\"},\"type\":\"array\"}]}},\"required\":[\"apiVersion\",\"tool\",\"generatedAt\"],\"type\":\"object\"},\"rows\":{\"items\":{\"properties\":{\"cells\":{\"additionalProperties\":{\"type\":[\"string\",\"number\"]},\"type\":\"object\"},\"error\":{\"type\":\"string\"},\"errorCode\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"retryable\":{\"type\":\"boolean\"}},\"required\":[\"name\",\"cells\"],\"type\":\"object\"},\"type\":\"array\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"tool\",\"count\",\"capped\",\"limit\",\"columns\",\"rows\",\"errors\",\"provenance\"],\"type\":\"object\"}},\"required\":[\"ok\",\"result\"],\"type\":\"object\"}}},\"description\":\"Batch completed (individual rows may carry per-record errors).\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Error shape returned by the batch and workflow endpoints.\",\"properties\":{\"error\":{\"type\":\"string\"},\"ok\":{\"const\":false}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Invalid request (bad JSON, missing tool/input, or tool not batchable).\"}},\"security\":[{}],\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/batch","segments":[{"lit":"batch"}],"select":{},"transform":{"req":"`reqdata`","res":"`body.result`"},"index$":0}],"key$":"create"},"load":{"input":"data","name":"load","points":[{"active":true,"args":{},"contract":{"id":"GET /batch","json":"{\"operationId\":\"batchInfo\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"type\":\"object\"}}},\"description\":\"Batchable tool list and limits.\"}},\"security\":[{}],\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/batch","segments":[{"lit":"batch"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"batch","name__orig":"batch","Name":"Batch","name_":"batch","name-":"batch","NAME":"BATCH","index$":3}, {"active":true,"entity":"batch","key$":"BasicBatchFlow","kind":"basic","name":"BasicBatchFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"batch_ref01"},"match":{},"op":"create","spec":[],"valid":[],"index$":0},{"active":true,"data":{},"input":{"ref":"batch_ref01","srcdatavar":"batch_ref01_data","suffix":"_dt0"},"match":{},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-batch_ref01"}}],"index$":1}]}, 'Batch')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const batch_ref01_ent = client.Batch()
    let batch_ref01_data = setup.data.new.batch['batch_ref01']

    batch_ref01_data = (await batch_ref01_ent.create(batch_ref01_data)).data()
    assert(null != batch_ref01_data)


    // LOAD
    const batch_ref01_match_dt0: any = {}
    const batch_ref01_data_dt0 = (await batch_ref01_ent.load(batch_ref01_match_dt0)).data()
    assert(null != batch_ref01_data_dt0)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/batch/BatchTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = SeqbenchMcpSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['batch01','batch02','batch03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'SEQBENCH_MCP_TEST_BATCH_ENTID': idmap,
    'SEQBENCH_MCP_TEST_LIVE': 'FALSE',
    'SEQBENCH_MCP_TEST_EXPLAIN': 'FALSE',
    'SEQBENCH_MCP_APIKEY': '',
  })

  idmap = env['SEQBENCH_MCP_TEST_BATCH_ENTID']

  const live = 'TRUE' === env.SEQBENCH_MCP_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['SEQBENCH_MCP_TEST_BATCH_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new SeqbenchMcpSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
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
    ]))
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
  }

  return setup
}
  
