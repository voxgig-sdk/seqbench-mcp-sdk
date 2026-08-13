# SeqbenchMcp SDK

require_relative 'utility/struct/voxgig_struct'
require_relative 'core/utility_type'
require_relative 'core/spec'
require_relative 'core/helpers'

# Load utility registration
require_relative 'utility/register'

# Load config and features
require_relative 'config'
require_relative 'feature/base_feature'
require_relative 'features'

# Load typed models (Struct value objects).
require_relative 'SeqbenchMcp_types'


class SeqbenchMcpSDK
  attr_accessor :mode, :features, :options

  def initialize(options = {})
    @mode = "live"
    @features = []
    @options = nil

    utility = SeqbenchMcpUtility.new
    @_utility = utility

    config = SeqbenchMcpConfig.make_config

    @_rootctx = utility.make_context.call({
      "client" => self,
      "utility" => utility,
      "config" => config,
      "options" => options || {},
      "shared" => {},
    }, nil)

    @options = utility.make_options.call(@_rootctx)

    if VoxgigStruct.getpath(@options, "feature.test.active") == true
      @mode = "test"
    end

    @_rootctx.options = @options

    # Add features in the resolved order (make_options puts an explicit array
    # order first, else defaults to test-first). Ordering matters: the `test`
    # feature installs the base mock transport and the transport features
    # (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
    # must be added before them to sit at the base of the chain.
    feature_opts = SeqbenchMcpHelpers.to_map(VoxgigStruct.getprop(@options, "feature"))
    if feature_opts
      featureorder = VoxgigStruct.getpath(@options, "__derived__.featureorder")
      if featureorder.is_a?(Array)
        featureorder.each do |fname|
          fopts = SeqbenchMcpHelpers.to_map(feature_opts[fname])
          if fopts && fopts["active"] == true
            utility.feature_add.call(@_rootctx, SeqbenchMcpFeatures.make_feature(fname))
          end
        end
      end
    end

    # Add extension features.
    extend_val = VoxgigStruct.getprop(@options, "extend")
    if extend_val.is_a?(Array)
      extend_val.each do |f|
        if f.respond_to?(:get_name)
          utility.feature_add.call(@_rootctx, f)
        end
      end
    end

    # Initialize features.
    @features.each do |f|
      utility.feature_init.call(@_rootctx, f)
    end

    utility.feature_hook.call(@_rootctx, "PostConstruct")
  end

  def options_map
    out = VoxgigStruct.clone(@options)
    out.is_a?(Hash) ? out : {}
  end

  def get_utility
    SeqbenchMcpUtility.copy(@_utility)
  end

  def get_root_ctx
    @_rootctx
  end

  def prepare(fetchargs = {})
    utility = @_utility
    fetchargs ||= {}

    ctrl = SeqbenchMcpHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "prepare",
      "ctrl" => ctrl,
    }, @_rootctx)

    opts = @options
    path = VoxgigStruct.getprop(fetchargs, "path") || ""
    path = "" unless path.is_a?(String)
    method_val = VoxgigStruct.getprop(fetchargs, "method") || "GET"
    method_val = "GET" unless method_val.is_a?(String)
    params = SeqbenchMcpHelpers.to_map(VoxgigStruct.getprop(fetchargs, "params")) || {}
    query = SeqbenchMcpHelpers.to_map(VoxgigStruct.getprop(fetchargs, "query")) || {}
    headers = utility.prepare_headers.call(ctx)

    base = VoxgigStruct.getprop(opts, "base") || ""
    base = "" unless base.is_a?(String)
    prefix = VoxgigStruct.getprop(opts, "prefix") || ""
    prefix = "" unless prefix.is_a?(String)
    suffix = VoxgigStruct.getprop(opts, "suffix") || ""
    suffix = "" unless suffix.is_a?(String)

    ctx.spec = SeqbenchMcpSpec.new({
      "base" => base, "prefix" => prefix, "suffix" => suffix,
      "path" => path, "method" => method_val,
      "params" => params, "query" => query, "headers" => headers,
      "body" => VoxgigStruct.getprop(fetchargs, "body"),
      "step" => "start",
    })

    # Merge user-provided headers.
    uh = VoxgigStruct.getprop(fetchargs, "headers")
    if uh.is_a?(Hash)
      uh.each { |k, v| ctx.spec.headers[k] = v }
    end

    _, err = utility.prepare_auth.call(ctx)
    raise err if err

    # make_fetch_def returns a (fetchdef, err) tuple; destructure it and
    # return just the fetchdef Hash (raising on error) so callers — including
    # direct(), which indexes fetchdef["url"] — receive a Hash, mirroring the
    # ts/py prepare().
    fetchdef, fd_err = utility.make_fetch_def.call(ctx)
    raise fd_err if fd_err

    fetchdef
  end

  # Raw endpoint access is operator-controllable, like every entity op.
  # Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
  # either one reaches the same endpoint.
  def direct(fetchargs = {})
    return op_denied("direct") unless op_allowed?("direct")

    raw_request(fetchargs)
  end

  # Is this raw-access op permitted by the SDK's allow.op option?
  def op_allowed?(op)
    allow_op = VoxgigStruct.getpath(@options, "allow.op")
    allow_op.is_a?(String) && allow_op.include?(op)
  end

  def op_denied(op)
    allow_op = VoxgigStruct.getpath(@options, "allow.op")
    {
      "ok" => false,
      "err" => SeqbenchMcpError.new(
        "#{op}_allow",
        "SeqbenchMcpSDK: #{op}: operation not allowed by" \
        " SDK option allow.op value: \"#{allow_op}\""),
    }
  end

  # Ungated request path shared by direct and graphql, each of which checks
  # its own allow.op token first. Separate, rather than a flag on fetchargs:
  # a caller-supplied marker would let anyone opt straight back out of the
  # gate by passing it.
  def raw_request(fetchargs = {})
    utility = @_utility

    # direct() is the raw-HTTP escape hatch: it always returns a result hash
    # ({ "ok" => ..., ... }) and never raises. prepare() raises on error, so
    # trap that and surface it in the hash.
    begin
      fetchdef = prepare(fetchargs)
    rescue SeqbenchMcpError => err
      return { "ok" => false, "err" => err }
    end

    fetchargs ||= {}
    ctrl = SeqbenchMcpHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "direct",
      "ctrl" => ctrl,
    }, @_rootctx)

    url = fetchdef["url"] || ""
    fetched, fetch_err = utility.fetcher.call(ctx, url, fetchdef)

    return { "ok" => false, "err" => fetch_err } if fetch_err

    if fetched.nil?
      return {
        "ok" => false,
        "err" => ctx.make_error("direct_no_response", "response: undefined"),
      }
    end

    if fetched.is_a?(Hash)
      status = SeqbenchMcpHelpers.to_int(VoxgigStruct.getprop(fetched, "status"))
      headers = VoxgigStruct.getprop(fetched, "headers") || {}

      # No-body responses (204, 304) and explicit zero content-length must
      # skip JSON parsing — calling json() on an empty body errors.
      content_length = headers.is_a?(Hash) ? headers["content-length"] : nil
      no_body = status == 204 || status == 304 || content_length.to_s == "0"

      json_data = nil
      unless no_body
        jf = VoxgigStruct.getprop(fetched, "json")
        if jf.is_a?(Proc)
          begin
            json_data = jf.call
          rescue StandardError
            # Non-JSON body — leave data nil, keep status/headers.
            json_data = nil
          end
        end
      end

      return {
        "ok" => status >= 200 && status < 300,
        "status" => status,
        "headers" => headers,
        "data" => json_data,
      }
    end

    return {
      "ok" => false,
      "err" => ctx.make_error("direct_invalid", "invalid response type"),
    }
  end

  # Raw GraphQL access: the pressure valve that makes the generated surface's
  # deliberate omissions (per-call selection sets, typed filter builders,
  # batching, subscriptions) livable — the whole schema stays reachable.
  #
  # Thin wrapper over the same prepare/fetch path direct uses, with the one
  # thing raw direct cannot do for GraphQL: a GraphQL failure rides HTTP 200
  # as a top-level `errors` array, so status alone would report a failed
  # query as ok.
  #
  # NOTE: like direct, this bypasses the feature pipeline — no retry,
  # ratelimit or paging features apply.
  def graphql(query, variables = nil, ctrl = nil)
    return op_denied("graphql") unless op_allowed?("graphql")

    res = raw_request({
      "method" => "POST",
      "headers" => { "content-type" => "application/json" },
      "body" => { "query" => query, "variables" => variables || {} },
      "ctrl" => ctrl || {},
    })

    # Errors are read BEFORE any status check: a GraphQL parse or validation
    # failure comes back as HTTP 400 carrying the standard { errors: [...] }
    # body, and the raw path represents a non-2xx as ok:false with no err —
    # so returning early on status would discard the server's own
    # diagnostics, which are the only useful part of that response.
    errors = VoxgigStruct.getpath(res, "data.errors")

    if errors.is_a?(Array) && !errors.empty?
      first = errors[0].is_a?(Hash) ? errors[0] : {}
      msg = first["message"]
      msg = "graphql error" if msg.nil? || msg.to_s.empty?
      res["ok"] = false
      res["err"] = SeqbenchMcpError.new(
        "graphql_error", "SeqbenchMcpSDK: graphql: #{msg}")
      res["graphql"] = errors
    end

    res
  end


  # Canonical facade: client.AlphafoldLookup.list / client.AlphafoldLookup.load({ "id" => ... })
  def AlphafoldLookup(data = nil)
    require_relative 'entity/alphafold_lookup_entity'
    AlphafoldLookupEntity.new(self, data)
  end


  # Canonical facade: client.AsoDesign.list / client.AsoDesign.load({ "id" => ... })
  def AsoDesign(data = nil)
    require_relative 'entity/aso_design_entity'
    AsoDesignEntity.new(self, data)
  end


  # Canonical facade: client.BaseEditingDesign.list / client.BaseEditingDesign.load({ "id" => ... })
  def BaseEditingDesign(data = nil)
    require_relative 'entity/base_editing_design_entity'
    BaseEditingDesignEntity.new(self, data)
  end


  # Canonical facade: client.Batch.list / client.Batch.load({ "id" => ... })
  def Batch(data = nil)
    require_relative 'entity/batch_entity'
    BatchEntity.new(self, data)
  end


  # Canonical facade: client.BatchWorkflow.list / client.BatchWorkflow.load({ "id" => ... })
  def BatchWorkflow(data = nil)
    require_relative 'entity/batch__workflow_entity'
    BatchWorkflowEntity.new(self, data)
  end


  # Canonical facade: client.CharacterizeSequence.list / client.CharacterizeSequence.load({ "id" => ... })
  def CharacterizeSequence(data = nil)
    require_relative 'entity/characterize_sequence_entity'
    CharacterizeSequenceEntity.new(self, data)
  end


  # Canonical facade: client.CloningSimulate.list / client.CloningSimulate.load({ "id" => ... })
  def CloningSimulate(data = nil)
    require_relative 'entity/cloning_simulate_entity'
    CloningSimulateEntity.new(self, data)
  end


  # Canonical facade: client.CodonAdaptationIndex.list / client.CodonAdaptationIndex.load({ "id" => ... })
  def CodonAdaptationIndex(data = nil)
    require_relative 'entity/codon_adaptation_index_entity'
    CodonAdaptationIndexEntity.new(self, data)
  end


  # Canonical facade: client.CodonOptimize.list / client.CodonOptimize.load({ "id" => ... })
  def CodonOptimize(data = nil)
    require_relative 'entity/codon_optimize_entity'
    CodonOptimizeEntity.new(self, data)
  end


  # Canonical facade: client.ConstructAutofix.list / client.ConstructAutofix.load({ "id" => ... })
  def ConstructAutofix(data = nil)
    require_relative 'entity/construct_autofix_entity'
    ConstructAutofixEntity.new(self, data)
  end


  # Canonical facade: client.ConstructQc.list / client.ConstructQc.load({ "id" => ... })
  def ConstructQc(data = nil)
    require_relative 'entity/construct_qc_entity'
    ConstructQcEntity.new(self, data)
  end


  # Canonical facade: client.CrisprGrnaDesign.list / client.CrisprGrnaDesign.load({ "id" => ... })
  def CrisprGrnaDesign(data = nil)
    require_relative 'entity/crispr_grna_design_entity'
    CrisprGrnaDesignEntity.new(self, data)
  end


  # Canonical facade: client.CrisprHdrDonor.list / client.CrisprHdrDonor.load({ "id" => ... })
  def CrisprHdrDonor(data = nil)
    require_relative 'entity/crispr_hdr_donor_entity'
    CrisprHdrDonorEntity.new(self, data)
  end


  # Canonical facade: client.CrisprOfftargetCheck.list / client.CrisprOfftargetCheck.load({ "id" => ... })
  def CrisprOfftargetCheck(data = nil)
    require_relative 'entity/crispr_offtarget_check_entity'
    CrisprOfftargetCheckEntity.new(self, data)
  end


  # Canonical facade: client.CrossDimer.list / client.CrossDimer.load({ "id" => ... })
  def CrossDimer(data = nil)
    require_relative 'entity/cross_dimer_entity'
    CrossDimerEntity.new(self, data)
  end


  # Canonical facade: client.DnaMolarity.list / client.DnaMolarity.load({ "id" => ... })
  def DnaMolarity(data = nil)
    require_relative 'entity/dna_molarity_entity'
    DnaMolarityEntity.new(self, data)
  end


  # Canonical facade: client.DoubleDigest.list / client.DoubleDigest.load({ "id" => ... })
  def DoubleDigest(data = nil)
    require_relative 'entity/double_digest_entity'
    DoubleDigestEntity.new(self, data)
  end


  # Canonical facade: client.ExportEchoPicklist.list / client.ExportEchoPicklist.load({ "id" => ... })
  def ExportEchoPicklist(data = nil)
    require_relative 'entity/export_echo_picklist_entity'
    ExportEchoPicklistEntity.new(self, data)
  end


  # Canonical facade: client.ExportOpentronsProtocol.list / client.ExportOpentronsProtocol.load({ "id" => ... })
  def ExportOpentronsProtocol(data = nil)
    require_relative 'entity/export_opentrons_protocol_entity'
    ExportOpentronsProtocolEntity.new(self, data)
  end


  # Canonical facade: client.ExportPlateLayout.list / client.ExportPlateLayout.load({ "id" => ... })
  def ExportPlateLayout(data = nil)
    require_relative 'entity/export_plate_layout_entity'
    ExportPlateLayoutEntity.new(self, data)
  end


  # Canonical facade: client.ExpressionHeatmapCluster.list / client.ExpressionHeatmapCluster.load({ "id" => ... })
  def ExpressionHeatmapCluster(data = nil)
    require_relative 'entity/expression_heatmap_cluster_entity'
    ExpressionHeatmapClusterEntity.new(self, data)
  end


  # Canonical facade: client.FastqQcReport.list / client.FastqQcReport.load({ "id" => ... })
  def FastqQcReport(data = nil)
    require_relative 'entity/fastq_qc_report_entity'
    FastqQcReportEntity.new(self, data)
  end


  # Canonical facade: client.FastqTrim.list / client.FastqTrim.load({ "id" => ... })
  def FastqTrim(data = nil)
    require_relative 'entity/fastq_trim_entity'
    FastqTrimEntity.new(self, data)
  end


  # Canonical facade: client.FindOrf.list / client.FindOrf.load({ "id" => ... })
  def FindOrf(data = nil)
    require_relative 'entity/find_orf_entity'
    FindOrfEntity.new(self, data)
  end


  # Canonical facade: client.FormatSequence.list / client.FormatSequence.load({ "id" => ... })
  def FormatSequence(data = nil)
    require_relative 'entity/format_sequence_entity'
    FormatSequenceEntity.new(self, data)
  end


  # Canonical facade: client.FunctionalEnrichment.list / client.FunctionalEnrichment.load({ "id" => ... })
  def FunctionalEnrichment(data = nil)
    require_relative 'entity/functional_enrichment_entity'
    FunctionalEnrichmentEntity.new(self, data)
  end


  # Canonical facade: client.GcContent.list / client.GcContent.load({ "id" => ... })
  def GcContent(data = nil)
    require_relative 'entity/gc_content_entity'
    GcContentEntity.new(self, data)
  end


  # Canonical facade: client.GeneDossier.list / client.GeneDossier.load({ "id" => ... })
  def GeneDossier(data = nil)
    require_relative 'entity/gene_dossier_entity'
    GeneDossierEntity.new(self, data)
  end


  # Canonical facade: client.GeneExpression.list / client.GeneExpression.load({ "id" => ... })
  def GeneExpression(data = nil)
    require_relative 'entity/gene_expression_entity'
    GeneExpressionEntity.new(self, data)
  end


  # Canonical facade: client.GeneModel.list / client.GeneModel.load({ "id" => ... })
  def GeneModel(data = nil)
    require_relative 'entity/gene_model_entity'
    GeneModelEntity.new(self, data)
  end


  # Canonical facade: client.GoldenGateFidelity.list / client.GoldenGateFidelity.load({ "id" => ... })
  def GoldenGateFidelity(data = nil)
    require_relative 'entity/golden_gate_fidelity_entity'
    GoldenGateFidelityEntity.new(self, data)
  end


  # Canonical facade: client.HgvsConvert.list / client.HgvsConvert.load({ "id" => ... })
  def HgvsConvert(data = nil)
    require_relative 'entity/hgvs_convert_entity'
    HgvsConvertEntity.new(self, data)
  end


  # Canonical facade: client.IdMapPoll.list / client.IdMapPoll.load({ "id" => ... })
  def IdMapPoll(data = nil)
    require_relative 'entity/id_map_poll_entity'
    IdMapPollEntity.new(self, data)
  end


  # Canonical facade: client.IdMapSubmit.list / client.IdMapSubmit.load({ "id" => ... })
  def IdMapSubmit(data = nil)
    require_relative 'entity/id_map_submit_entity'
    IdMapSubmitEntity.new(self, data)
  end


  # Canonical facade: client.InSilicoPcr.list / client.InSilicoPcr.load({ "id" => ... })
  def InSilicoPcr(data = nil)
    require_relative 'entity/in_silico_pcr_entity'
    InSilicoPcrEntity.new(self, data)
  end


  # Canonical facade: client.KaspPrimerDesign.list / client.KaspPrimerDesign.load({ "id" => ... })
  def KaspPrimerDesign(data = nil)
    require_relative 'entity/kasp_primer_design_entity'
    KaspPrimerDesignEntity.new(self, data)
  end


  # Canonical facade: client.ListTool.list / client.ListTool.load({ "id" => ... })
  def ListTool(data = nil)
    require_relative 'entity/list_tool_entity'
    ListToolEntity.new(self, data)
  end


  # Canonical facade: client.MeltingTemperature.list / client.MeltingTemperature.load({ "id" => ... })
  def MeltingTemperature(data = nil)
    require_relative 'entity/melting_temperature_entity'
    MeltingTemperatureEntity.new(self, data)
  end


  # Canonical facade: client.MotifFinder.list / client.MotifFinder.load({ "id" => ... })
  def MotifFinder(data = nil)
    require_relative 'entity/motif_finder_entity'
    MotifFinderEntity.new(self, data)
  end


  # Canonical facade: client.MultipleSequenceAlignment.list / client.MultipleSequenceAlignment.load({ "id" => ... })
  def MultipleSequenceAlignment(data = nil)
    require_relative 'entity/multiple_sequence_alignment_entity'
    MultipleSequenceAlignmentEntity.new(self, data)
  end


  # Canonical facade: client.OligoAnalysi.list / client.OligoAnalysi.load({ "id" => ... })
  def OligoAnalysi(data = nil)
    require_relative 'entity/oligo_analysi_entity'
    OligoAnalysiEntity.new(self, data)
  end


  # Canonical facade: client.OrthologMap.list / client.OrthologMap.load({ "id" => ... })
  def OrthologMap(data = nil)
    require_relative 'entity/ortholog_map_entity'
    OrthologMapEntity.new(self, data)
  end


  # Canonical facade: client.PairwiseAlignment.list / client.PairwiseAlignment.load({ "id" => ... })
  def PairwiseAlignment(data = nil)
    require_relative 'entity/pairwise_alignment_entity'
    PairwiseAlignmentEntity.new(self, data)
  end


  # Canonical facade: client.ParseGenbank.list / client.ParseGenbank.load({ "id" => ... })
  def ParseGenbank(data = nil)
    require_relative 'entity/parse_genbank_entity'
    ParseGenbankEntity.new(self, data)
  end


  # Canonical facade: client.ParseSangerTrace.list / client.ParseSangerTrace.load({ "id" => ... })
  def ParseSangerTrace(data = nil)
    require_relative 'entity/parse_sanger_trace_entity'
    ParseSangerTraceEntity.new(self, data)
  end


  # Canonical facade: client.PlasmidAnnotate.list / client.PlasmidAnnotate.load({ "id" => ... })
  def PlasmidAnnotate(data = nil)
    require_relative 'entity/plasmid_annotate_entity'
    PlasmidAnnotateEntity.new(self, data)
  end


  # Canonical facade: client.PlasmidDeepAnnotate.list / client.PlasmidDeepAnnotate.load({ "id" => ... })
  def PlasmidDeepAnnotate(data = nil)
    require_relative 'entity/plasmid_deep_annotate_entity'
    PlasmidDeepAnnotateEntity.new(self, data)
  end


  # Canonical facade: client.PlasmidFullReport.list / client.PlasmidFullReport.load({ "id" => ... })
  def PlasmidFullReport(data = nil)
    require_relative 'entity/plasmid_full_report_entity'
    PlasmidFullReportEntity.new(self, data)
  end


  # Canonical facade: client.PlasmidIdentify.list / client.PlasmidIdentify.load({ "id" => ... })
  def PlasmidIdentify(data = nil)
    require_relative 'entity/plasmid_identify_entity'
    PlasmidIdentifyEntity.new(self, data)
  end


  # Canonical facade: client.PrimeEditingDesign.list / client.PrimeEditingDesign.load({ "id" => ... })
  def PrimeEditingDesign(data = nil)
    require_relative 'entity/prime_editing_design_entity'
    PrimeEditingDesignEntity.new(self, data)
  end


  # Canonical facade: client.PrimeEditingTwinDesign.list / client.PrimeEditingTwinDesign.load({ "id" => ... })
  def PrimeEditingTwinDesign(data = nil)
    require_relative 'entity/prime_editing_twin_design_entity'
    PrimeEditingTwinDesignEntity.new(self, data)
  end


  # Canonical facade: client.PrimerDesign.list / client.PrimerDesign.load({ "id" => ... })
  def PrimerDesign(data = nil)
    require_relative 'entity/primer_design_entity'
    PrimerDesignEntity.new(self, data)
  end


  # Canonical facade: client.PrimerSpecificity.list / client.PrimerSpecificity.load({ "id" => ... })
  def PrimerSpecificity(data = nil)
    require_relative 'entity/primer_specificity_entity'
    PrimerSpecificityEntity.new(self, data)
  end


  # Canonical facade: client.ProteaseDigestion.list / client.ProteaseDigestion.load({ "id" => ... })
  def ProteaseDigestion(data = nil)
    require_relative 'entity/protease_digestion_entity'
    ProteaseDigestionEntity.new(self, data)
  end


  # Canonical facade: client.ProteinAnnotatePoll.list / client.ProteinAnnotatePoll.load({ "id" => ... })
  def ProteinAnnotatePoll(data = nil)
    require_relative 'entity/protein_annotate_poll_entity'
    ProteinAnnotatePollEntity.new(self, data)
  end


  # Canonical facade: client.ProteinAnnotateSubmit.list / client.ProteinAnnotateSubmit.load({ "id" => ... })
  def ProteinAnnotateSubmit(data = nil)
    require_relative 'entity/protein_annotate_submit_entity'
    ProteinAnnotateSubmitEntity.new(self, data)
  end


  # Canonical facade: client.ProteinHydrophobicity.list / client.ProteinHydrophobicity.load({ "id" => ... })
  def ProteinHydrophobicity(data = nil)
    require_relative 'entity/protein_hydrophobicity_entity'
    ProteinHydrophobicityEntity.new(self, data)
  end


  # Canonical facade: client.ProteinProperty.list / client.ProteinProperty.load({ "id" => ... })
  def ProteinProperty(data = nil)
    require_relative 'entity/protein_property_entity'
    ProteinPropertyEntity.new(self, data)
  end


  # Canonical facade: client.RandomSequence.list / client.RandomSequence.load({ "id" => ... })
  def RandomSequence(data = nil)
    require_relative 'entity/random_sequence_entity'
    RandomSequenceEntity.new(self, data)
  end


  # Canonical facade: client.RestrictionSite.list / client.RestrictionSite.load({ "id" => ... })
  def RestrictionSite(data = nil)
    require_relative 'entity/restriction_site_entity'
    RestrictionSiteEntity.new(self, data)
  end


  # Canonical facade: client.ReverseComplement.list / client.ReverseComplement.load({ "id" => ... })
  def ReverseComplement(data = nil)
    require_relative 'entity/reverse_complement_entity'
    ReverseComplementEntity.new(self, data)
  end


  # Canonical facade: client.ReverseTranslate.list / client.ReverseTranslate.load({ "id" => ... })
  def ReverseTranslate(data = nil)
    require_relative 'entity/reverse_translate_entity'
    ReverseTranslateEntity.new(self, data)
  end


  # Canonical facade: client.RnaFold.list / client.RnaFold.load({ "id" => ... })
  def RnaFold(data = nil)
    require_relative 'entity/rna_fold_entity'
    RnaFoldEntity.new(self, data)
  end


  # Canonical facade: client.SangerVsReference.list / client.SangerVsReference.load({ "id" => ... })
  def SangerVsReference(data = nil)
    require_relative 'entity/sanger_vs_reference_entity'
    SangerVsReferenceEntity.new(self, data)
  end


  # Canonical facade: client.SavePermalink.list / client.SavePermalink.load({ "id" => ... })
  def SavePermalink(data = nil)
    require_relative 'entity/save_permalink_entity'
    SavePermalinkEntity.new(self, data)
  end


  # Canonical facade: client.SeqfileStat.list / client.SeqfileStat.load({ "id" => ... })
  def SeqfileStat(data = nil)
    require_relative 'entity/seqfile_stat_entity'
    SeqfileStatEntity.new(self, data)
  end


  # Canonical facade: client.SequenceFetch.list / client.SequenceFetch.load({ "id" => ... })
  def SequenceFetch(data = nil)
    require_relative 'entity/sequence_fetch_entity'
    SequenceFetchEntity.new(self, data)
  end


  # Canonical facade: client.SequenceFormatConvert.list / client.SequenceFormatConvert.load({ "id" => ... })
  def SequenceFormatConvert(data = nil)
    require_relative 'entity/sequence_format_convert_entity'
    SequenceFormatConvertEntity.new(self, data)
  end


  # Canonical facade: client.SequenceReport.list / client.SequenceReport.load({ "id" => ... })
  def SequenceReport(data = nil)
    require_relative 'entity/sequence_report_entity'
    SequenceReportEntity.new(self, data)
  end


  # Canonical facade: client.SequenceSearch.list / client.SequenceSearch.load({ "id" => ... })
  def SequenceSearch(data = nil)
    require_relative 'entity/sequence_search_entity'
    SequenceSearchEntity.new(self, data)
  end


  # Canonical facade: client.SequencingReadbackVerify.list / client.SequencingReadbackVerify.load({ "id" => ... })
  def SequencingReadbackVerify(data = nil)
    require_relative 'entity/sequencing_readback_verify_entity'
    SequencingReadbackVerifyEntity.new(self, data)
  end


  # Canonical facade: client.SessionCreate.list / client.SessionCreate.load({ "id" => ... })
  def SessionCreate(data = nil)
    require_relative 'entity/session_create_entity'
    SessionCreateEntity.new(self, data)
  end


  # Canonical facade: client.SessionGet.list / client.SessionGet.load({ "id" => ... })
  def SessionGet(data = nil)
    require_relative 'entity/session_get_entity'
    SessionGetEntity.new(self, data)
  end


  # Canonical facade: client.SessionRun.list / client.SessionRun.load({ "id" => ... })
  def SessionRun(data = nil)
    require_relative 'entity/session_run_entity'
    SessionRunEntity.new(self, data)
  end


  # Canonical facade: client.SessionSet.list / client.SessionSet.load({ "id" => ... })
  def SessionSet(data = nil)
    require_relative 'entity/session_set_entity'
    SessionSetEntity.new(self, data)
  end


  # Canonical facade: client.SirnaDesign.list / client.SirnaDesign.load({ "id" => ... })
  def SirnaDesign(data = nil)
    require_relative 'entity/sirna_design_entity'
    SirnaDesignEntity.new(self, data)
  end


  # Canonical facade: client.SiteDirectedMutagenesi.list / client.SiteDirectedMutagenesi.load({ "id" => ... })
  def SiteDirectedMutagenesi(data = nil)
    require_relative 'entity/site_directed_mutagenesi_entity'
    SiteDirectedMutagenesiEntity.new(self, data)
  end


  # Canonical facade: client.Translate.list / client.Translate.load({ "id" => ... })
  def Translate(data = nil)
    require_relative 'entity/translate_entity'
    TranslateEntity.new(self, data)
  end


  # Canonical facade: client.VariantAnnotate.list / client.VariantAnnotate.load({ "id" => ... })
  def VariantAnnotate(data = nil)
    require_relative 'entity/variant_annotate_entity'
    VariantAnnotateEntity.new(self, data)
  end


  # Canonical facade: client.VariantComparator.list / client.VariantComparator.load({ "id" => ... })
  def VariantComparator(data = nil)
    require_relative 'entity/variant_comparator_entity'
    VariantComparatorEntity.new(self, data)
  end


  # Canonical facade: client.VerifyAssembly.list / client.VerifyAssembly.load({ "id" => ... })
  def VerifyAssembly(data = nil)
    require_relative 'entity/verify_assembly_entity'
    VerifyAssemblyEntity.new(self, data)
  end


  # Canonical facade: client.VerifyConstruct.list / client.VerifyConstruct.load({ "id" => ... })
  def VerifyConstruct(data = nil)
    require_relative 'entity/verify_construct_entity'
    VerifyConstructEntity.new(self, data)
  end


  # Canonical facade: client.VirtualGel.list / client.VirtualGel.load({ "id" => ... })
  def VirtualGel(data = nil)
    require_relative 'entity/virtual_gel_entity'
    VirtualGelEntity.new(self, data)
  end


  # Canonical facade: client.VolcanoPlotData.list / client.VolcanoPlotData.load({ "id" => ... })
  def VolcanoPlotData(data = nil)
    require_relative 'entity/volcano_plot_data_entity'
    VolcanoPlotDataEntity.new(self, data)
  end


  # Canonical facade: client.WebSearch.list / client.WebSearch.load({ "id" => ... })
  def WebSearch(data = nil)
    require_relative 'entity/web_search_entity'
    WebSearchEntity.new(self, data)
  end



  def self.test(testopts = nil, sdkopts = nil)
    sdkopts = sdkopts || {}
    sdkopts = VoxgigStruct.clone(sdkopts)
    sdkopts = {} unless sdkopts.is_a?(Hash)

    testopts = testopts || {}
    testopts = VoxgigStruct.clone(testopts)
    testopts = {} unless testopts.is_a?(Hash)
    testopts["active"] = true

    VoxgigStruct.setpath(sdkopts, "feature.test", testopts)

    sdk = SeqbenchMcpSDK.new(sdkopts)
    sdk.mode = "test"
    sdk
  end
end
