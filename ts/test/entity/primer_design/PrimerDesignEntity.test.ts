

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


describe('PrimerDesignEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when SEQBENCH_MCP_TEST_LIVE=TRUE.
  afterEach(liveDelay('SEQBENCH_MCP_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = SeqbenchMcpSDK.test()
    const ent = testsdk.PrimerDesign()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.SEQBENCH_MCP_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'primer_design.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"ampliconMax","req":false,"type":"`$INTEGER`","index$":0},{"active":true,"name":"ampliconMin","req":false,"type":"`$INTEGER`","index$":1},{"active":true,"name":"dntpMM","req":false,"short":"Total [dNTP] (mM), chelates Mg2+.","type":"`$NUMBER`","index$":2},{"active":true,"name":"gate","req":false,"short":"Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.","type":"`$ANY`","index$":3},{"active":true,"name":"gcMax","req":false,"type":"`$NUMBER`","index$":4},{"active":true,"name":"gcMin","req":false,"type":"`$NUMBER`","index$":5},{"active":true,"name":"lenMax","req":false,"type":"`$INTEGER`","index$":6},{"active":true,"name":"lenMin","req":false,"type":"`$INTEGER`","index$":7},{"active":true,"name":"lenOpt","req":false,"type":"`$INTEGER`","index$":8},{"active":true,"name":"maxReturn","req":false,"short":"Number of best pairs to return.","type":"`$INTEGER`","index$":9},{"active":true,"name":"mgMM","req":false,"short":"Divalent cation [Mg2+] (mM).","type":"`$NUMBER`","index$":10},{"active":true,"name":"naMM","req":false,"short":"Monovalent cation [Na+]/[K+] (mM).","type":"`$NUMBER`","index$":11},{"active":true,"name":"ok","req":true,"type":"`$ANY`","index$":12},{"active":true,"name":"oligoNM","req":false,"short":"Total strand concentration (nM).","type":"`$NUMBER`","index$":13},{"active":true,"name":"provenance","req":true,"type":"`$OBJECT`","union":{"branches":2,"count":1,"depth":2},"index$":14},{"active":true,"name":"result","req":true,"short":"Tool-specific output object.","type":"`$OBJECT`","index$":15},{"active":true,"name":"targetEnd","req":false,"short":"1-based inclusive end of the target region (optional).","type":"`$INTEGER`","index$":16},{"active":true,"name":"targetStart","req":false,"short":"1-based inclusive start of a region the product must span (optional).","type":"`$INTEGER`","index$":17},{"active":true,"name":"template","req":true,"short":"Nucleotide sequence (raw or FASTA; IUPAC accepted).","type":"`$STRING`","index$":18},{"active":true,"name":"tmMax","req":false,"type":"`$NUMBER`","index$":19},{"active":true,"name":"tmMaxDiff","req":false,"short":"Max Tm difference within a pair (°C).","type":"`$NUMBER`","index$":20},{"active":true,"name":"tmMin","req":false,"type":"`$NUMBER`","index$":21},{"active":true,"name":"tmOpt","req":false,"type":"`$NUMBER`","index$":22},{"active":true,"name":"tool","req":true,"short":"The tool slug that ran.","type":"`$STRING`","index$":23}],"name":"primer_design","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{},"contract":{"id":"POST /primer_design","json":"{\"operationId\":\"primer_design\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"ampliconMax\":{\"default\":1000,\"type\":\"integer\"},\"ampliconMin\":{\"default\":100,\"type\":\"integer\"},\"dntpMM\":{\"default\":0.2,\"description\":\"Total [dNTP] (mM), chelates Mg2+.\",\"type\":\"number\"},\"gcMax\":{\"default\":60,\"type\":\"number\"},\"gcMin\":{\"default\":40,\"type\":\"number\"},\"lenMax\":{\"default\":25,\"type\":\"integer\"},\"lenMin\":{\"default\":18,\"type\":\"integer\"},\"lenOpt\":{\"default\":20,\"type\":\"integer\"},\"maxReturn\":{\"default\":5,\"description\":\"Number of best pairs to return.\",\"type\":\"integer\"},\"mgMM\":{\"default\":1.5,\"description\":\"Divalent cation [Mg2+] (mM).\",\"type\":\"number\"},\"naMM\":{\"default\":50,\"description\":\"Monovalent cation [Na+]/[K+] (mM).\",\"type\":\"number\"},\"oligoNM\":{\"default\":250,\"description\":\"Total strand concentration (nM).\",\"type\":\"number\"},\"targetEnd\":{\"description\":\"1-based inclusive end of the target region (optional).\",\"type\":\"integer\"},\"targetStart\":{\"description\":\"1-based inclusive start of a region the product must span (optional).\",\"type\":\"integer\"},\"template\":{\"description\":\"Nucleotide sequence (raw or FASTA; IUPAC accepted).\",\"type\":\"string\"},\"tmMax\":{\"default\":63,\"type\":\"number\"},\"tmMaxDiff\":{\"default\":3,\"description\":\"Max Tm difference within a pair (°C).\",\"type\":\"number\"},\"tmMin\":{\"default\":57,\"type\":\"number\"},\"tmOpt\":{\"default\":60,\"type\":\"number\"}},\"required\":[\"template\"],\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard success envelope returned by every single-tool call.\",\"properties\":{\"gate\":{\"description\":\"Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. no targetTm).\",\"oneOf\":[{\"properties\":{\"checks\":{\"items\":{\"properties\":{\"id\":{\"type\":\"string\"},\"label\":{\"type\":\"string\"},\"message\":{\"type\":\"string\"},\"pass\":{\"type\":\"boolean\"},\"severity\":{\"enum\":[\"hard\",\"soft\"],\"type\":\"string\"},\"threshold\":{\"type\":\"string\"},\"value\":{\"type\":\"number\"}},\"required\":[\"id\",\"label\",\"severity\",\"pass\",\"message\"],\"type\":\"object\"},\"type\":\"array\"},\"notChecked\":{\"description\":\"Honest list of what this gate does NOT verify.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"pass\":{\"description\":\"True if and only if every hard check passed.\",\"type\":\"boolean\"}},\"required\":[\"pass\",\"checks\",\"notChecked\"],\"type\":\"object\"},{\"type\":\"null\"}]},\"ok\":{\"const\":true},\"provenance\":{\"properties\":{\"apiVersion\":{\"example\":\"1.1.0\",\"type\":\"string\"},\"generatedAt\":{\"description\":\"ISO 8601 timestamp of when the result was generated.\",\"format\":\"date-time\",\"type\":\"string\"},\"tool\":{\"description\":\"Tool slug, or an array of slugs (one per step) for a workflow.\",\"oneOf\":[{\"type\":\"string\"},{\"items\":{\"type\":\"string\"},\"type\":\"array\"}]}},\"required\":[\"apiVersion\",\"tool\",\"generatedAt\"],\"type\":\"object\"},\"result\":{\"additionalProperties\":true,\"description\":\"Tool-specific output object. Its shape depends on the tool — call `GET /{tool}` or see the docs for each tool's fields.\",\"type\":\"object\"},\"tool\":{\"description\":\"The tool slug that ran.\",\"type\":\"string\"}},\"required\":[\"ok\",\"tool\",\"result\",\"provenance\"],\"type\":\"object\"}}},\"description\":\"Tool ran successfully.\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Invalid argument or the tool could not run.\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Unknown tool.\"}},\"security\":[{}],\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/primer_design","segments":[{"lit":"primer_design"}],"select":{},"transform":{"req":{"ampliconMax":"`reqdata.amplicon_max`","ampliconMin":"`reqdata.amplicon_min`","dntpMM":"`reqdata.dntp_mm`","gcMax":"`reqdata.gc_max`","gcMin":"`reqdata.gc_min`","lenMax":"`reqdata.len_max`","lenMin":"`reqdata.len_min`","lenOpt":"`reqdata.len_opt`","maxReturn":"`reqdata.max_return`","mgMM":"`reqdata.mg_mm`","naMM":"`reqdata.na_mm`","oligoNM":"`reqdata.oligo_nm`","targetEnd":"`reqdata.target_end`","targetStart":"`reqdata.target_start`","template":"`reqdata.template`","tmMax":"`reqdata.tm_max`","tmMaxDiff":"`reqdata.tm_max_diff`","tmMin":"`reqdata.tm_min`","tmOpt":"`reqdata.tm_opt`"},"res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"primer_design","name__orig":"primer_design","Name":"PrimerDesign","name_":"primer_design","name-":"primer-design","NAME":"PRIMER_DESIGN","index$":51}, {"active":true,"entity":"primer_design","key$":"BasicPrimerDesignFlow","kind":"basic","name":"BasicPrimerDesignFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"primer_design_ref01"},"match":{},"op":"create","spec":[],"valid":[],"index$":0}]}, 'PrimerDesign')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const primer_design_ref01_ent = client.PrimerDesign()
    let primer_design_ref01_data = setup.data.new.primer_design['primer_design_ref01']

    primer_design_ref01_data = (await primer_design_ref01_ent.create(primer_design_ref01_data)).data()
    assert(null != primer_design_ref01_data)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/primer_design/PrimerDesignTestData.json')

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
    ['primer_design01','primer_design02','primer_design03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'SEQBENCH_MCP_TEST_PRIMER_DESIGN_ENTID': idmap,
    'SEQBENCH_MCP_TEST_LIVE': 'FALSE',
    'SEQBENCH_MCP_TEST_EXPLAIN': 'FALSE',
    'SEQBENCH_MCP_APIKEY': '',
  })

  idmap = env['SEQBENCH_MCP_TEST_PRIMER_DESIGN_ENTID']

  const live = 'TRUE' === env.SEQBENCH_MCP_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['SEQBENCH_MCP_TEST_PRIMER_DESIGN_ENTID']
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
  
