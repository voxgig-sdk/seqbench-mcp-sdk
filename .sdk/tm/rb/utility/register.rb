# SeqbenchMcp SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

SeqbenchMcpUtility.registrar = ->(u) {
  u.clean = SeqbenchMcpUtilities::Clean
  u.done = SeqbenchMcpUtilities::Done
  u.make_error = SeqbenchMcpUtilities::MakeError
  u.feature_add = SeqbenchMcpUtilities::FeatureAdd
  u.feature_hook = SeqbenchMcpUtilities::FeatureHook
  u.feature_init = SeqbenchMcpUtilities::FeatureInit
  u.fetcher = SeqbenchMcpUtilities::Fetcher
  u.make_fetch_def = SeqbenchMcpUtilities::MakeFetchDef
  u.make_context = SeqbenchMcpUtilities::MakeContext
  u.make_options = SeqbenchMcpUtilities::MakeOptions
  u.make_request = SeqbenchMcpUtilities::MakeRequest
  u.make_response = SeqbenchMcpUtilities::MakeResponse
  u.make_result = SeqbenchMcpUtilities::MakeResult
  u.make_point = SeqbenchMcpUtilities::MakePoint
  u.make_spec = SeqbenchMcpUtilities::MakeSpec
  u.make_url = SeqbenchMcpUtilities::MakeUrl
  u.param = SeqbenchMcpUtilities::Param
  u.prepare_auth = SeqbenchMcpUtilities::PrepareAuth
  u.prepare_body = SeqbenchMcpUtilities::PrepareBody
  u.prepare_headers = SeqbenchMcpUtilities::PrepareHeaders
  u.prepare_method = SeqbenchMcpUtilities::PrepareMethod
  u.prepare_params = SeqbenchMcpUtilities::PrepareParams
  u.prepare_path = SeqbenchMcpUtilities::PreparePath
  u.prepare_query = SeqbenchMcpUtilities::PrepareQuery
  u.graphql_body = SeqbenchMcpUtilities::GraphqlBody
  u.graphql_errors = SeqbenchMcpUtilities::GraphqlErrors
  u.result_basic = SeqbenchMcpUtilities::ResultBasic
  u.result_body = SeqbenchMcpUtilities::ResultBody
  u.result_headers = SeqbenchMcpUtilities::ResultHeaders
  u.transform_request = SeqbenchMcpUtilities::TransformRequest
  u.transform_response = SeqbenchMcpUtilities::TransformResponse
}
