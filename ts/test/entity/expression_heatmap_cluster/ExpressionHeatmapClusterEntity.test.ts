

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


describe('ExpressionHeatmapClusterEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when SEQBENCH_MCP_TEST_LIVE=TRUE.
  afterEach(liveDelay('SEQBENCH_MCP_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = SeqbenchMcpSDK.test()
    const ent = testsdk.ExpressionHeatmapCluster()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.SEQBENCH_MCP_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'expression_heatmap_cluster.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"clusterCols","req":false,"short":"Cluster (reorder) samples.","type":"`$BOOLEAN`","index$":0},{"active":true,"name":"clusterRows","req":false,"short":"Cluster (reorder) genes.","type":"`$BOOLEAN`","index$":1},{"active":true,"name":"distanceMetric","req":false,"short":"correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance.","type":"`$STRING`","index$":2},{"active":true,"name":"gate","req":false,"short":"Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.","type":"`$ANY`","index$":3},{"active":true,"name":"genes","req":true,"short":"Row (gene) labels.","type":"`$ARRAY`","index$":4},{"active":true,"name":"linkage","req":false,"short":"average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor.","type":"`$STRING`","index$":5},{"active":true,"name":"ok","req":true,"type":"`$ANY`","index$":6},{"active":true,"name":"provenance","req":true,"type":"`$OBJECT`","union":{"branches":2,"count":1,"depth":2},"index$":7},{"active":true,"name":"result","req":true,"short":"Tool-specific output object.","type":"`$OBJECT`","index$":8},{"active":true,"name":"samples","req":true,"short":"Column (sample) labels.","type":"`$ARRAY`","index$":9},{"active":true,"name":"tool","req":true,"short":"The tool slug that ran.","type":"`$STRING`","index$":10},{"active":true,"name":"values","req":true,"short":"genes x samples numeric matrix — one row per gene, in the same order as `genes`.","type":"`$ARRAY`","index$":11},{"active":true,"name":"zScoreRows","req":false,"short":"Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization).","type":"`$BOOLEAN`","index$":12}],"name":"expression_heatmap_cluster","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{},"contract":{"id":"POST /expression_heatmap_cluster","json":"{\"operationId\":\"expression_heatmap_cluster\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"clusterCols\":{\"default\":true,\"description\":\"Cluster (reorder) samples.\",\"type\":\"boolean\"},\"clusterRows\":{\"default\":true,\"description\":\"Cluster (reorder) genes.\",\"type\":\"boolean\"},\"distanceMetric\":{\"default\":\"correlation\",\"description\":\"correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance.\",\"enum\":[\"euclidean\",\"correlation\"],\"type\":\"string\"},\"genes\":{\"description\":\"Row (gene) labels.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"linkage\":{\"default\":\"average\",\"description\":\"average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor.\",\"enum\":[\"average\",\"complete\",\"single\"],\"type\":\"string\"},\"samples\":{\"description\":\"Column (sample) labels.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"values\":{\"description\":\"genes x samples numeric matrix — one row per gene, in the same order as `genes`.\",\"items\":{\"items\":{\"type\":\"number\"},\"type\":\"array\"},\"type\":\"array\"},\"zScoreRows\":{\"default\":true,\"description\":\"Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization).\",\"type\":\"boolean\"}},\"required\":[\"genes\",\"samples\",\"values\"],\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard success envelope returned by every single-tool call.\",\"properties\":{\"gate\":{\"description\":\"Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g. no targetTm).\",\"oneOf\":[{\"properties\":{\"checks\":{\"items\":{\"properties\":{\"id\":{\"type\":\"string\"},\"label\":{\"type\":\"string\"},\"message\":{\"type\":\"string\"},\"pass\":{\"type\":\"boolean\"},\"severity\":{\"enum\":[\"hard\",\"soft\"],\"type\":\"string\"},\"threshold\":{\"type\":\"string\"},\"value\":{\"type\":\"number\"}},\"required\":[\"id\",\"label\",\"severity\",\"pass\",\"message\"],\"type\":\"object\"},\"type\":\"array\"},\"notChecked\":{\"description\":\"Honest list of what this gate does NOT verify.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"pass\":{\"description\":\"True if and only if every hard check passed.\",\"type\":\"boolean\"}},\"required\":[\"pass\",\"checks\",\"notChecked\"],\"type\":\"object\"},{\"type\":\"null\"}]},\"ok\":{\"const\":true},\"provenance\":{\"properties\":{\"apiVersion\":{\"example\":\"1.1.0\",\"type\":\"string\"},\"generatedAt\":{\"description\":\"ISO 8601 timestamp of when the result was generated.\",\"format\":\"date-time\",\"type\":\"string\"},\"tool\":{\"description\":\"Tool slug, or an array of slugs (one per step) for a workflow.\",\"oneOf\":[{\"type\":\"string\"},{\"items\":{\"type\":\"string\"},\"type\":\"array\"}]}},\"required\":[\"apiVersion\",\"tool\",\"generatedAt\"],\"type\":\"object\"},\"result\":{\"additionalProperties\":true,\"description\":\"Tool-specific output object. Its shape depends on the tool — call `GET /{tool}` or see the docs for each tool's fields.\",\"type\":\"object\"},\"tool\":{\"description\":\"The tool slug that ran.\",\"type\":\"string\"}},\"required\":[\"ok\",\"tool\",\"result\",\"provenance\"],\"type\":\"object\"}}},\"description\":\"Tool ran successfully.\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Invalid argument or the tool could not run.\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Standard error envelope for single-tool calls.\",\"properties\":{\"code\":{\"enum\":[\"invalid_argument\",\"unsupported\",\"upstream_unavailable\",\"upstream_timeout\",\"rate_limited\",\"internal_error\"],\"type\":\"string\"},\"error\":{\"description\":\"Human-readable error message.\",\"type\":\"string\"},\"ok\":{\"const\":false},\"retryable\":{\"type\":\"boolean\"},\"suggestedAction\":{\"type\":\"string\"},\"tool\":{\"type\":\"string\"}},\"required\":[\"ok\",\"error\"],\"type\":\"object\"}}},\"description\":\"Unknown tool.\"}},\"security\":[{}],\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/expression_heatmap_cluster","segments":[{"lit":"expression_heatmap_cluster"}],"select":{},"transform":{"req":{"clusterCols":"`reqdata.cluster_col`","clusterRows":"`reqdata.cluster_row`","distanceMetric":"`reqdata.distance_metric`","genes":"`reqdata.gene`","linkage":"`reqdata.linkage`","samples":"`reqdata.sample`","values":"`reqdata.value`","zScoreRows":"`reqdata.z_score_row`"},"res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"expression_heatmap_cluster","name__orig":"expression_heatmap_cluster","Name":"ExpressionHeatmapCluster","name_":"expression_heatmap_cluster","name-":"expression-heatmap-cluster","NAME":"EXPRESSION_HEATMAP_CLUSTER","index$":20}, {"active":true,"entity":"expression_heatmap_cluster","key$":"BasicExpressionHeatmapClusterFlow","kind":"basic","name":"BasicExpressionHeatmapClusterFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"expression_heatmap_cluster_ref01"},"match":{},"op":"create","spec":[],"valid":[],"index$":0}]}, 'ExpressionHeatmapCluster')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const expression_heatmap_cluster_ref01_ent = client.ExpressionHeatmapCluster()
    let expression_heatmap_cluster_ref01_data = setup.data.new.expression_heatmap_cluster['expression_heatmap_cluster_ref01']

    expression_heatmap_cluster_ref01_data = (await expression_heatmap_cluster_ref01_ent.create(expression_heatmap_cluster_ref01_data)).data()
    assert(null != expression_heatmap_cluster_ref01_data)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/expression_heatmap_cluster/ExpressionHeatmapClusterTestData.json')

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
    ['expression_heatmap_cluster01','expression_heatmap_cluster02','expression_heatmap_cluster03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'SEQBENCH_MCP_TEST_EXPRESSION_HEATMAP_CLUSTER_ENTID': idmap,
    'SEQBENCH_MCP_TEST_LIVE': 'FALSE',
    'SEQBENCH_MCP_TEST_EXPLAIN': 'FALSE',
    'SEQBENCH_MCP_APIKEY': '',
  })

  idmap = env['SEQBENCH_MCP_TEST_EXPRESSION_HEATMAP_CLUSTER_ENTID']

  const live = 'TRUE' === env.SEQBENCH_MCP_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['SEQBENCH_MCP_TEST_EXPRESSION_HEATMAP_CLUSTER_ENTID']
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
  
