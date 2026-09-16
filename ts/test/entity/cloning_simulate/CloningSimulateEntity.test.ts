

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


describe('CloningSimulateEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when SEQBENCH_MCP_TEST_LIVE=TRUE.
  afterEach(liveDelay('SEQBENCH_MCP_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = SeqbenchMcpSDK.test()
    const ent = testsdk.CloningSimulate()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.SEQBENCH_MCP_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'cloning_simulate.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"armTmTarget","req":false,"short":"Target annealing Tm (°C) for primer arms.","type":"`$NUMBER`","index$":0},{"active":true,"name":"circular","req":false,"short":"Produce a circular product.","type":"`$BOOLEAN`","index$":1},{"active":true,"name":"enzyme","req":false,"short":"Type IIS enzyme for Golden Gate (e.g.","type":"`$STRING`","index$":2},{"active":true,"name":"enzyme3","req":false,"short":"3′ enzyme (restriction method).","type":"`$STRING`","index$":3},{"active":true,"name":"enzyme5","req":false,"short":"5′ enzyme (restriction method).","type":"`$STRING`","index$":4},{"active":true,"name":"fragments","req":false,"short":"Fragments (5′→3′), assembled head-to-tail.","type":"`$ARRAY`","index$":5},{"active":true,"name":"gate","req":false,"short":"Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.","type":"`$ANY`","index$":6},{"active":true,"name":"insert","req":false,"short":"Insert sequence (restriction method).","type":"`$STRING`","index$":7},{"active":true,"name":"method","req":true,"short":"Assembly method.","type":"`$STRING`","index$":8},{"active":true,"name":"names","req":false,"short":"Optional labels for each fragment.","type":"`$ARRAY`","index$":9},{"active":true,"name":"ok","req":true,"type":"`$ANY`","index$":10},{"active":true,"name":"overlapLen","req":false,"short":"Gibson homology-arm length (bp).","type":"`$INTEGER`","index$":11},{"active":true,"name":"provenance","req":true,"type":"`$OBJECT`","union":{"branches":2,"count":1,"depth":2},"index$":12},{"active":true,"name":"result","req":true,"short":"Tool-specific output object.","type":"`$OBJECT`","index$":13},{"active":true,"name":"tool","req":true,"short":"The tool slug that ran.","type":"`$STRING`","index$":14},{"active":true,"name":"vector","req":false,"short":"Vector sequence (restriction method).","type":"`$STRING`","index$":15}],"name":"cloning_simulate","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{},"contract":{"id":"POST /cloning_simulate","json":"{\"operationId\":\"cloning_simulate\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"example\":{\"circular\":true,\"fragments\":[\"ATGGCCAAGCTGACCGAACTGAAAGCCGCCGTGGAAACCCTG\",\"GAAACCCTGGATAAAGCCTTCAAAGATGCCCTGAAAGCCTAA\"],\"method\":\"gibson\",\"overlapLen\":20},\"schema\":{\"additionalProperties\":false,\"properties\":{\"armTmTarget\":{\"default\":60,\"description\":\"Target annealing Tm (°C) for primer arms.\",\"type\":\"number\"},\"circular\":{\"default\":true,\"description\":\"Produce a circular product.\",\"type\":\"boolean\"},\"enzyme\":{\"default\":\"BsaI\",\"description\":\"Type IIS enzyme for Golden Gate (e.g. BsaI, BbsI, Esp3I (BsmBI)).\",\"type\":\"string\"},\"enzyme3\":{\"default\":\"BamHI\",\"description\":\"3′ enzyme (restriction method).\",\"type\":\"string\"},\"enzyme5\":{\"default\":\"EcoRI\",\"description\":\"5′ enzyme (restriction method).\",\"type\":\"string\"},\"fragments\":{\"description\":\"Fragments (5′→3′), assembled head-to-tail. Used by gibson/goldengate.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"insert\":{\"description\":\"Insert sequence (restriction method).\",\"type\":\"string\"},\"method\":{\"default\":\"gibson\",\"description\":\"Assembly method.\",\"enum\":[\"gibson\",\"goldengate\",\"restriction\"],\"type\":\"string\"},\"names\":{\"description\":\"Optional labels for each fragment.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"overlapLen\":{\"default\":20,\"description\":\"Gibson homology-arm length (bp).\",\"type\":\"integer\"},\"vector\":{\"description\":\"Vector sequence (restriction method).\",\"type\":\"string\"}},\"required\":[\"method\"],\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard success envelope returned by every single-tool call.\",\"properties\":{\"gate\":{\"description\":\"Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. no targetTm).\",\"oneOf\":[{\"properties\":{\"checks\":{\"items\":{\"properties\":{\"id\":{\"type\":\"string\"},\"label\":{\"type\":\"string\"},\"message\":{\"type\":\"string\"},\"pass\":{\"type\":\"boolean\"},\"severity\":{\"enum\":[\"hard\",\"soft\"],\"type\":\"string\"},\"threshold\":{\"type\":\"string\"},\"value\":{\"type\":\"number\"}},\"required\":[\"id\",\"label\",\"severity\",\"pass\",\"message\"],\"type\":\"object\"},\"type\":\"array\"},\"notChecked\":{\"description\":\"Honest list of what this gate does NOT verify.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"pass\":{\"description\":\"True if and only if every hard check passed.\",\"type\":\"boolean\"}},\"required\":[\"pass\",\"checks\",\"notChecked\"],\"type\":\"object\"},{\"type\":\"null\"}]},\"ok\":{\"const\":true},\"provenance\":{\"properties\":{\"apiVersion\":{\"example\":\"1.1.0\",\"type\":\"string\"},\"generatedAt\":{\"description\":\"ISO 8601 timestamp of when the result was generated.\",\"format\":\"date-time\",\"type\":\"string\"},\"tool\":{\"description\":\"Tool slug, or an array of slugs (one per step) for a workflow.\",\"oneOf\":[{\"type\":\"string\"},{\"items\":{\"type\":\"string\"},\"type\":\"array\"}]}},\"required\":[\"apiVersion\",\"tool\",\"generatedAt\"],\"type\":\"object\"},\"result\":{\"additionalProperties\":true,\"description\":\"Tool-specific output object. Its shape depends on the tool — call `GET /{tool}` or see the docs for each tool's fields.\",\"type\":\"object\"},\"tool\":{\"description\":\"The tool slug that ran.\",\"type\":\"string\"}},\"required\":[\"ok\",\"tool\",\"result\",\"provenance\"],\"type\":\"object\"}}},\"description\":\"Tool ran successfully.\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Invalid argument or the tool could not run.\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Unknown tool.\"}},\"security\":[{}],\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/cloning_simulate","segments":[{"lit":"cloning_simulate"}],"select":{},"transform":{"req":{"armTmTarget":"`reqdata.arm_tm_target`","circular":"`reqdata.circular`","enzyme":"`reqdata.enzyme`","enzyme3":"`reqdata.enzyme3`","enzyme5":"`reqdata.enzyme5`","fragments":"`reqdata.fragment`","insert":"`reqdata.insert`","method":"`reqdata.method`","names":"`reqdata.name`","overlapLen":"`reqdata.overlap_len`","vector":"`reqdata.vector`"},"res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"cloning_simulate","name__orig":"cloning_simulate","Name":"CloningSimulate","name_":"cloning_simulate","name-":"cloning-simulate","NAME":"CLONING_SIMULATE","index$":6}, {"active":true,"entity":"cloning_simulate","key$":"BasicCloningSimulateFlow","kind":"basic","name":"BasicCloningSimulateFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"cloning_simulate_ref01"},"match":{},"op":"create","spec":[],"valid":[],"index$":0}]}, 'CloningSimulate')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const cloning_simulate_ref01_ent = client.CloningSimulate()
    let cloning_simulate_ref01_data = setup.data.new.cloning_simulate['cloning_simulate_ref01']

    cloning_simulate_ref01_data = (await cloning_simulate_ref01_ent.create(cloning_simulate_ref01_data)).data()
    assert(null != cloning_simulate_ref01_data)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/cloning_simulate/CloningSimulateTestData.json')

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
    ['cloning_simulate01','cloning_simulate02','cloning_simulate03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'SEQBENCH_MCP_TEST_CLONING_SIMULATE_ENTID': idmap,
    'SEQBENCH_MCP_TEST_LIVE': 'FALSE',
    'SEQBENCH_MCP_TEST_EXPLAIN': 'FALSE',
    'SEQBENCH_MCP_APIKEY': '',
  })

  idmap = env['SEQBENCH_MCP_TEST_CLONING_SIMULATE_ENTID']

  const live = 'TRUE' === env.SEQBENCH_MCP_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['SEQBENCH_MCP_TEST_CLONING_SIMULATE_ENTID']
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
  
