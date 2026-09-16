

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


describe('GoldenGateFidelityEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when SEQBENCH_MCP_TEST_LIVE=TRUE.
  afterEach(liveDelay('SEQBENCH_MCP_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = SeqbenchMcpSDK.test()
    const ent = testsdk.GoldenGateFidelity()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.SEQBENCH_MCP_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'golden_gate_fidelity.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"compareToNamedSet","req":false,"short":"Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison.","type":"`$STRING`","index$":0},{"active":true,"name":"dataset","req":false,"short":"Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme.","type":"`$STRING`","index$":1},{"active":true,"name":"gate","req":false,"short":"Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.","type":"`$ANY`","index$":2},{"active":true,"name":"ok","req":true,"type":"`$ANY`","index$":3},{"active":true,"name":"overhangs","req":true,"short":"The candidate 4-base overhangs for one assembly (e.g.","type":"`$ARRAY`","index$":4},{"active":true,"name":"provenance","req":true,"type":"`$OBJECT`","union":{"branches":2,"count":1,"depth":2},"index$":5},{"active":true,"name":"result","req":true,"short":"Tool-specific output object.","type":"`$OBJECT`","index$":6},{"active":true,"name":"riskThreshold","req":false,"short":"Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal.","type":"`$NUMBER`","index$":7},{"active":true,"name":"tool","req":true,"short":"The tool slug that ran.","type":"`$STRING`","index$":8}],"name":"golden_gate_fidelity","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{},"contract":{"id":"POST /golden_gate_fidelity","json":"{\"operationId\":\"golden_gate_fidelity\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"compareToNamedSet\":{\"description\":\"Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison.\",\"enum\":[\"pryor-2020-plant-11\",\"pryor-2020-20set\",\"cidar-moclo\"],\"type\":\"string\"},\"dataset\":{\"default\":\"generic-t4-37c-1h\",\"description\":\"Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme.\",\"enum\":[\"generic-t4-37c-1h\",\"bsai-hfv2\"],\"type\":\"string\"},\"overhangs\":{\"description\":\"The candidate 4-base overhangs for one assembly (e.g. [\\\"GGAG\\\",\\\"TACT\\\",\\\"AATG\\\"]). At least 2, no duplicates.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"riskThreshold\":{\"default\":0.05,\"description\":\"Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal.\",\"type\":\"number\"}},\"required\":[\"overhangs\"],\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard success envelope returned by every single-tool call.\",\"properties\":{\"gate\":{\"description\":\"Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. no targetTm).\",\"oneOf\":[{\"properties\":{\"checks\":{\"items\":{\"properties\":{\"id\":{\"type\":\"string\"},\"label\":{\"type\":\"string\"},\"message\":{\"type\":\"string\"},\"pass\":{\"type\":\"boolean\"},\"severity\":{\"enum\":[\"hard\",\"soft\"],\"type\":\"string\"},\"threshold\":{\"type\":\"string\"},\"value\":{\"type\":\"number\"}},\"required\":[\"id\",\"label\",\"severity\",\"pass\",\"message\"],\"type\":\"object\"},\"type\":\"array\"},\"notChecked\":{\"description\":\"Honest list of what this gate does NOT verify.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"pass\":{\"description\":\"True if and only if every hard check passed.\",\"type\":\"boolean\"}},\"required\":[\"pass\",\"checks\",\"notChecked\"],\"type\":\"object\"},{\"type\":\"null\"}]},\"ok\":{\"const\":true},\"provenance\":{\"properties\":{\"apiVersion\":{\"example\":\"1.1.0\",\"type\":\"string\"},\"generatedAt\":{\"description\":\"ISO 8601 timestamp of when the result was generated.\",\"format\":\"date-time\",\"type\":\"string\"},\"tool\":{\"description\":\"Tool slug, or an array of slugs (one per step) for a workflow.\",\"oneOf\":[{\"type\":\"string\"},{\"items\":{\"type\":\"string\"},\"type\":\"array\"}]}},\"required\":[\"apiVersion\",\"tool\",\"generatedAt\"],\"type\":\"object\"},\"result\":{\"additionalProperties\":true,\"description\":\"Tool-specific output object. Its shape depends on the tool — call `GET /{tool}` or see the docs for each tool's fields.\",\"type\":\"object\"},\"tool\":{\"description\":\"The tool slug that ran.\",\"type\":\"string\"}},\"required\":[\"ok\",\"tool\",\"result\",\"provenance\"],\"type\":\"object\"}}},\"description\":\"Tool ran successfully.\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Invalid argument or the tool could not run.\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Unknown tool.\"}},\"security\":[{}],\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/golden_gate_fidelity","segments":[{"lit":"golden_gate_fidelity"}],"select":{},"transform":{"req":{"compareToNamedSet":"`reqdata.compare_to_named_set`","dataset":"`reqdata.dataset`","overhangs":"`reqdata.overhang`","riskThreshold":"`reqdata.risk_threshold`"},"res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"golden_gate_fidelity","name__orig":"golden_gate_fidelity","Name":"GoldenGateFidelity","name_":"golden_gate_fidelity","name-":"golden-gate-fidelity","NAME":"GOLDEN_GATE_FIDELITY","index$":30}, {"active":true,"entity":"golden_gate_fidelity","key$":"BasicGoldenGateFidelityFlow","kind":"basic","name":"BasicGoldenGateFidelityFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"golden_gate_fidelity_ref01"},"match":{},"op":"create","spec":[],"valid":[],"index$":0}]}, 'GoldenGateFidelity')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const golden_gate_fidelity_ref01_ent = client.GoldenGateFidelity()
    let golden_gate_fidelity_ref01_data = setup.data.new.golden_gate_fidelity['golden_gate_fidelity_ref01']

    golden_gate_fidelity_ref01_data = (await golden_gate_fidelity_ref01_ent.create(golden_gate_fidelity_ref01_data)).data()
    assert(null != golden_gate_fidelity_ref01_data)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/golden_gate_fidelity/GoldenGateFidelityTestData.json')

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
    ['golden_gate_fidelity01','golden_gate_fidelity02','golden_gate_fidelity03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'SEQBENCH_MCP_TEST_GOLDEN_GATE_FIDELITY_ENTID': idmap,
    'SEQBENCH_MCP_TEST_LIVE': 'FALSE',
    'SEQBENCH_MCP_TEST_EXPLAIN': 'FALSE',
    'SEQBENCH_MCP_APIKEY': '',
  })

  idmap = env['SEQBENCH_MCP_TEST_GOLDEN_GATE_FIDELITY_ENTID']

  const live = 'TRUE' === env.SEQBENCH_MCP_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['SEQBENCH_MCP_TEST_GOLDEN_GATE_FIDELITY_ENTID']
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
  
