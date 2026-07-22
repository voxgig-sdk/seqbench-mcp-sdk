-- SeqbenchMcp SDK

local vs = require("utility.struct.struct")
local Utility = require("core.utility_type")
local Spec = require("core.spec")
local helpers = require("core.helpers")

-- Load utility registration (populates Utility._registrar)
require("utility.register")

-- Load features
local BaseFeature = require("feature.base_feature")
local features_factory = require("features")


local SeqbenchMcpSDK = {}
SeqbenchMcpSDK.__index = SeqbenchMcpSDK


local function _make_feature(name)
  local factory = features_factory[name]
  if factory ~= nil then
    return factory()
  end
  return features_factory.base()
end

SeqbenchMcpSDK._make_feature = _make_feature


function SeqbenchMcpSDK.new(options)
  local self = setmetatable({}, SeqbenchMcpSDK)
  self.mode = "live"
  self.features = {}
  self.options = nil

  local utility = Utility.new()
  self._utility = utility

  local config = require("config")()

  self._rootctx = utility.make_context({
    client = self,
    utility = utility,
    config = config,
    options = options or {},
    shared = {},
  }, nil)

  self.options = utility.make_options(self._rootctx)

  if vs.getpath(self.options, "feature.test.active") == true then
    self.mode = "test"
  end

  self._rootctx.options = self.options

  -- Add features in the resolved order (make_options puts an explicit list
  -- order first, else defaults to test-first). Ordering matters: the `test`
  -- feature installs the base mock transport and the transport features
  -- (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
  -- must be added before them to sit at the base of the chain.
  local feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
  if feature_opts ~= nil then
    local featureorder = vs.getpath(self.options, "__derived__.featureorder")
    if type(featureorder) == "table" then
      for _, fname in ipairs(featureorder) do
        local fopts = helpers.to_map(feature_opts[fname])
        if fopts ~= nil and fopts["active"] == true then
          utility.feature_add(self._rootctx, _make_feature(fname))
        end
      end
    end
  end

  -- Add extension features.
  local extend = vs.getprop(self.options, "extend")
  if type(extend) == "table" then
    for _, f in ipairs(extend) do
      if type(f) == "table" and type(f.get_name) == "function" then
        utility.feature_add(self._rootctx, f)
      end
    end
  end

  -- Initialize features.
  for _, f in ipairs(self.features) do
    utility.feature_init(self._rootctx, f)
  end

  utility.feature_hook(self._rootctx, "PostConstruct")

  -- #BuildFeatures

  return self
end


function SeqbenchMcpSDK:options_map()
  local out = vs.clone(self.options)
  if type(out) == "table" then
    return out
  end
  return {}
end


function SeqbenchMcpSDK:get_utility()
  return Utility.copy(self._utility)
end


function SeqbenchMcpSDK:get_root_ctx()
  return self._rootctx
end


function SeqbenchMcpSDK:prepare(fetchargs)
  local utility = self._utility

  fetchargs = fetchargs or {}

  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "prepare",
    ctrl = ctrl,
  }, self._rootctx)

  local options = self.options

  local path = vs.getprop(fetchargs, "path") or ""
  if type(path) ~= "string" then path = "" end

  local method = vs.getprop(fetchargs, "method") or "GET"
  if type(method) ~= "string" then method = "GET" end

  local params = helpers.to_map(vs.getprop(fetchargs, "params")) or {}
  local query = helpers.to_map(vs.getprop(fetchargs, "query")) or {}

  local headers = utility.prepare_headers(ctx)

  local base = vs.getprop(options, "base") or ""
  if type(base) ~= "string" then base = "" end
  local prefix = vs.getprop(options, "prefix") or ""
  if type(prefix) ~= "string" then prefix = "" end
  local suffix = vs.getprop(options, "suffix") or ""
  if type(suffix) ~= "string" then suffix = "" end

  ctx.spec = Spec.new({
    base = base,
    prefix = prefix,
    suffix = suffix,
    path = path,
    method = method,
    params = params,
    query = query,
    headers = headers,
    body = vs.getprop(fetchargs, "body"),
    step = "start",
  })

  -- Merge user-provided headers.
  local uh = vs.getprop(fetchargs, "headers")
  if type(uh) == "table" then
    for k, v in pairs(uh) do
      ctx.spec.headers[k] = v
    end
  end

  local _, err = utility.prepare_auth(ctx)
  if err ~= nil then
    return nil, err
  end

  return utility.make_fetch_def(ctx)
end


function SeqbenchMcpSDK:direct(fetchargs)
  local utility = self._utility

  local fetchdef, err = self:prepare(fetchargs)
  if err ~= nil then
    return { ok = false, err = err }, nil
  end

  fetchargs = fetchargs or {}
  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "direct",
    ctrl = ctrl,
  }, self._rootctx)

  local url = fetchdef["url"] or ""
  local fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

  if fetch_err ~= nil then
    return { ok = false, err = fetch_err }, nil
  end

  if fetched == nil then
    return {
      ok = false,
      err = ctx:make_error("direct_no_response", "response: undefined"),
    }, nil
  end

  if type(fetched) == "table" then
    local status = helpers.to_int(vs.getprop(fetched, "status"))
    local headers = vs.getprop(fetched, "headers") or {}

    -- No-body responses (204, 304) and explicit zero content-length
    -- must skip JSON parsing — calling json() on an empty body errors.
    local content_length = nil
    if type(headers) == "table" then
      content_length = headers["content-length"]
    end
    local no_body = status == 204 or status == 304 or tostring(content_length) == "0"

    local json_data = nil
    if not no_body then
      local jf = vs.getprop(fetched, "json")
      if type(jf) == "function" then
        local ok, result = pcall(jf)
        if ok then
          json_data = result
        end
        -- Non-JSON body: json_data stays nil, status/headers preserved.
      end
    end

    return {
      ok = status >= 200 and status < 300,
      status = status,
      headers = headers,
      data = json_data,
    }, nil
  end

  return {
    ok = false,
    err = ctx:make_error("direct_invalid", "invalid response type"),
  }, nil
end



-- Idiomatic facade: client:AlphafoldLookup():list() / client:AlphafoldLookup():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:AlphafoldLookup(data)
  local EntityMod = require("entity.alphafold_lookup_entity")
  if data == nil then
    if self._alphafold_lookup == nil then
      self._alphafold_lookup = EntityMod.new(self, nil)
    end
    return self._alphafold_lookup
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:AsoDesign():list() / client:AsoDesign():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:AsoDesign(data)
  local EntityMod = require("entity.aso_design_entity")
  if data == nil then
    if self._aso_design == nil then
      self._aso_design = EntityMod.new(self, nil)
    end
    return self._aso_design
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:BaseEditingDesign():list() / client:BaseEditingDesign():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:BaseEditingDesign(data)
  local EntityMod = require("entity.base_editing_design_entity")
  if data == nil then
    if self._base_editing_design == nil then
      self._base_editing_design = EntityMod.new(self, nil)
    end
    return self._base_editing_design
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Batch():list() / client:Batch():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:Batch(data)
  local EntityMod = require("entity.batch_entity")
  if data == nil then
    if self._batch == nil then
      self._batch = EntityMod.new(self, nil)
    end
    return self._batch
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:BatchWorkflow():list() / client:BatchWorkflow():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:BatchWorkflow(data)
  local EntityMod = require("entity.batch__workflow_entity")
  if data == nil then
    if self._batch__workflow == nil then
      self._batch__workflow = EntityMod.new(self, nil)
    end
    return self._batch__workflow
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CharacterizeSequence():list() / client:CharacterizeSequence():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:CharacterizeSequence(data)
  local EntityMod = require("entity.characterize_sequence_entity")
  if data == nil then
    if self._characterize_sequence == nil then
      self._characterize_sequence = EntityMod.new(self, nil)
    end
    return self._characterize_sequence
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CloningSimulate():list() / client:CloningSimulate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:CloningSimulate(data)
  local EntityMod = require("entity.cloning_simulate_entity")
  if data == nil then
    if self._cloning_simulate == nil then
      self._cloning_simulate = EntityMod.new(self, nil)
    end
    return self._cloning_simulate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CodonAdaptationIndex():list() / client:CodonAdaptationIndex():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:CodonAdaptationIndex(data)
  local EntityMod = require("entity.codon_adaptation_index_entity")
  if data == nil then
    if self._codon_adaptation_index == nil then
      self._codon_adaptation_index = EntityMod.new(self, nil)
    end
    return self._codon_adaptation_index
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CodonOptimize():list() / client:CodonOptimize():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:CodonOptimize(data)
  local EntityMod = require("entity.codon_optimize_entity")
  if data == nil then
    if self._codon_optimize == nil then
      self._codon_optimize = EntityMod.new(self, nil)
    end
    return self._codon_optimize
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ConstructAutofix():list() / client:ConstructAutofix():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ConstructAutofix(data)
  local EntityMod = require("entity.construct_autofix_entity")
  if data == nil then
    if self._construct_autofix == nil then
      self._construct_autofix = EntityMod.new(self, nil)
    end
    return self._construct_autofix
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ConstructQc():list() / client:ConstructQc():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ConstructQc(data)
  local EntityMod = require("entity.construct_qc_entity")
  if data == nil then
    if self._construct_qc == nil then
      self._construct_qc = EntityMod.new(self, nil)
    end
    return self._construct_qc
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CrisprGrnaDesign():list() / client:CrisprGrnaDesign():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:CrisprGrnaDesign(data)
  local EntityMod = require("entity.crispr_grna_design_entity")
  if data == nil then
    if self._crispr_grna_design == nil then
      self._crispr_grna_design = EntityMod.new(self, nil)
    end
    return self._crispr_grna_design
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CrisprHdrDonor():list() / client:CrisprHdrDonor():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:CrisprHdrDonor(data)
  local EntityMod = require("entity.crispr_hdr_donor_entity")
  if data == nil then
    if self._crispr_hdr_donor == nil then
      self._crispr_hdr_donor = EntityMod.new(self, nil)
    end
    return self._crispr_hdr_donor
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CrisprOfftargetCheck():list() / client:CrisprOfftargetCheck():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:CrisprOfftargetCheck(data)
  local EntityMod = require("entity.crispr_offtarget_check_entity")
  if data == nil then
    if self._crispr_offtarget_check == nil then
      self._crispr_offtarget_check = EntityMod.new(self, nil)
    end
    return self._crispr_offtarget_check
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CrossDimer():list() / client:CrossDimer():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:CrossDimer(data)
  local EntityMod = require("entity.cross_dimer_entity")
  if data == nil then
    if self._cross_dimer == nil then
      self._cross_dimer = EntityMod.new(self, nil)
    end
    return self._cross_dimer
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:DnaMolarity():list() / client:DnaMolarity():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:DnaMolarity(data)
  local EntityMod = require("entity.dna_molarity_entity")
  if data == nil then
    if self._dna_molarity == nil then
      self._dna_molarity = EntityMod.new(self, nil)
    end
    return self._dna_molarity
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:DoubleDigest():list() / client:DoubleDigest():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:DoubleDigest(data)
  local EntityMod = require("entity.double_digest_entity")
  if data == nil then
    if self._double_digest == nil then
      self._double_digest = EntityMod.new(self, nil)
    end
    return self._double_digest
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ExportEchoPicklist():list() / client:ExportEchoPicklist():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ExportEchoPicklist(data)
  local EntityMod = require("entity.export_echo_picklist_entity")
  if data == nil then
    if self._export_echo_picklist == nil then
      self._export_echo_picklist = EntityMod.new(self, nil)
    end
    return self._export_echo_picklist
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ExportOpentronsProtocol():list() / client:ExportOpentronsProtocol():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ExportOpentronsProtocol(data)
  local EntityMod = require("entity.export_opentrons_protocol_entity")
  if data == nil then
    if self._export_opentrons_protocol == nil then
      self._export_opentrons_protocol = EntityMod.new(self, nil)
    end
    return self._export_opentrons_protocol
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ExportPlateLayout():list() / client:ExportPlateLayout():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ExportPlateLayout(data)
  local EntityMod = require("entity.export_plate_layout_entity")
  if data == nil then
    if self._export_plate_layout == nil then
      self._export_plate_layout = EntityMod.new(self, nil)
    end
    return self._export_plate_layout
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ExpressionHeatmapCluster():list() / client:ExpressionHeatmapCluster():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ExpressionHeatmapCluster(data)
  local EntityMod = require("entity.expression_heatmap_cluster_entity")
  if data == nil then
    if self._expression_heatmap_cluster == nil then
      self._expression_heatmap_cluster = EntityMod.new(self, nil)
    end
    return self._expression_heatmap_cluster
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:FastqQcReport():list() / client:FastqQcReport():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:FastqQcReport(data)
  local EntityMod = require("entity.fastq_qc_report_entity")
  if data == nil then
    if self._fastq_qc_report == nil then
      self._fastq_qc_report = EntityMod.new(self, nil)
    end
    return self._fastq_qc_report
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:FastqTrim():list() / client:FastqTrim():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:FastqTrim(data)
  local EntityMod = require("entity.fastq_trim_entity")
  if data == nil then
    if self._fastq_trim == nil then
      self._fastq_trim = EntityMod.new(self, nil)
    end
    return self._fastq_trim
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:FindOrf():list() / client:FindOrf():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:FindOrf(data)
  local EntityMod = require("entity.find_orf_entity")
  if data == nil then
    if self._find_orf == nil then
      self._find_orf = EntityMod.new(self, nil)
    end
    return self._find_orf
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:FormatSequence():list() / client:FormatSequence():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:FormatSequence(data)
  local EntityMod = require("entity.format_sequence_entity")
  if data == nil then
    if self._format_sequence == nil then
      self._format_sequence = EntityMod.new(self, nil)
    end
    return self._format_sequence
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:FunctionalEnrichment():list() / client:FunctionalEnrichment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:FunctionalEnrichment(data)
  local EntityMod = require("entity.functional_enrichment_entity")
  if data == nil then
    if self._functional_enrichment == nil then
      self._functional_enrichment = EntityMod.new(self, nil)
    end
    return self._functional_enrichment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:GcContent():list() / client:GcContent():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:GcContent(data)
  local EntityMod = require("entity.gc_content_entity")
  if data == nil then
    if self._gc_content == nil then
      self._gc_content = EntityMod.new(self, nil)
    end
    return self._gc_content
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:GeneDossier():list() / client:GeneDossier():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:GeneDossier(data)
  local EntityMod = require("entity.gene_dossier_entity")
  if data == nil then
    if self._gene_dossier == nil then
      self._gene_dossier = EntityMod.new(self, nil)
    end
    return self._gene_dossier
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:GeneExpression():list() / client:GeneExpression():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:GeneExpression(data)
  local EntityMod = require("entity.gene_expression_entity")
  if data == nil then
    if self._gene_expression == nil then
      self._gene_expression = EntityMod.new(self, nil)
    end
    return self._gene_expression
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:GeneModel():list() / client:GeneModel():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:GeneModel(data)
  local EntityMod = require("entity.gene_model_entity")
  if data == nil then
    if self._gene_model == nil then
      self._gene_model = EntityMod.new(self, nil)
    end
    return self._gene_model
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:GoldenGateFidelity():list() / client:GoldenGateFidelity():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:GoldenGateFidelity(data)
  local EntityMod = require("entity.golden_gate_fidelity_entity")
  if data == nil then
    if self._golden_gate_fidelity == nil then
      self._golden_gate_fidelity = EntityMod.new(self, nil)
    end
    return self._golden_gate_fidelity
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:HgvsConvert():list() / client:HgvsConvert():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:HgvsConvert(data)
  local EntityMod = require("entity.hgvs_convert_entity")
  if data == nil then
    if self._hgvs_convert == nil then
      self._hgvs_convert = EntityMod.new(self, nil)
    end
    return self._hgvs_convert
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:IdMapPoll():list() / client:IdMapPoll():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:IdMapPoll(data)
  local EntityMod = require("entity.id_map_poll_entity")
  if data == nil then
    if self._id_map_poll == nil then
      self._id_map_poll = EntityMod.new(self, nil)
    end
    return self._id_map_poll
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:IdMapSubmit():list() / client:IdMapSubmit():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:IdMapSubmit(data)
  local EntityMod = require("entity.id_map_submit_entity")
  if data == nil then
    if self._id_map_submit == nil then
      self._id_map_submit = EntityMod.new(self, nil)
    end
    return self._id_map_submit
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:InSilicoPcr():list() / client:InSilicoPcr():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:InSilicoPcr(data)
  local EntityMod = require("entity.in_silico_pcr_entity")
  if data == nil then
    if self._in_silico_pcr == nil then
      self._in_silico_pcr = EntityMod.new(self, nil)
    end
    return self._in_silico_pcr
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:KaspPrimerDesign():list() / client:KaspPrimerDesign():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:KaspPrimerDesign(data)
  local EntityMod = require("entity.kasp_primer_design_entity")
  if data == nil then
    if self._kasp_primer_design == nil then
      self._kasp_primer_design = EntityMod.new(self, nil)
    end
    return self._kasp_primer_design
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ListTool():list() / client:ListTool():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ListTool(data)
  local EntityMod = require("entity.list_tool_entity")
  if data == nil then
    if self._list_tool == nil then
      self._list_tool = EntityMod.new(self, nil)
    end
    return self._list_tool
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:MeltingTemperature():list() / client:MeltingTemperature():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:MeltingTemperature(data)
  local EntityMod = require("entity.melting_temperature_entity")
  if data == nil then
    if self._melting_temperature == nil then
      self._melting_temperature = EntityMod.new(self, nil)
    end
    return self._melting_temperature
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:MotifFinder():list() / client:MotifFinder():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:MotifFinder(data)
  local EntityMod = require("entity.motif_finder_entity")
  if data == nil then
    if self._motif_finder == nil then
      self._motif_finder = EntityMod.new(self, nil)
    end
    return self._motif_finder
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:MultipleSequenceAlignment():list() / client:MultipleSequenceAlignment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:MultipleSequenceAlignment(data)
  local EntityMod = require("entity.multiple_sequence_alignment_entity")
  if data == nil then
    if self._multiple_sequence_alignment == nil then
      self._multiple_sequence_alignment = EntityMod.new(self, nil)
    end
    return self._multiple_sequence_alignment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:OligoAnalysi():list() / client:OligoAnalysi():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:OligoAnalysi(data)
  local EntityMod = require("entity.oligo_analysi_entity")
  if data == nil then
    if self._oligo_analysi == nil then
      self._oligo_analysi = EntityMod.new(self, nil)
    end
    return self._oligo_analysi
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:OrthologMap():list() / client:OrthologMap():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:OrthologMap(data)
  local EntityMod = require("entity.ortholog_map_entity")
  if data == nil then
    if self._ortholog_map == nil then
      self._ortholog_map = EntityMod.new(self, nil)
    end
    return self._ortholog_map
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PairwiseAlignment():list() / client:PairwiseAlignment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:PairwiseAlignment(data)
  local EntityMod = require("entity.pairwise_alignment_entity")
  if data == nil then
    if self._pairwise_alignment == nil then
      self._pairwise_alignment = EntityMod.new(self, nil)
    end
    return self._pairwise_alignment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ParseGenbank():list() / client:ParseGenbank():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ParseGenbank(data)
  local EntityMod = require("entity.parse_genbank_entity")
  if data == nil then
    if self._parse_genbank == nil then
      self._parse_genbank = EntityMod.new(self, nil)
    end
    return self._parse_genbank
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ParseSangerTrace():list() / client:ParseSangerTrace():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ParseSangerTrace(data)
  local EntityMod = require("entity.parse_sanger_trace_entity")
  if data == nil then
    if self._parse_sanger_trace == nil then
      self._parse_sanger_trace = EntityMod.new(self, nil)
    end
    return self._parse_sanger_trace
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PlasmidAnnotate():list() / client:PlasmidAnnotate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:PlasmidAnnotate(data)
  local EntityMod = require("entity.plasmid_annotate_entity")
  if data == nil then
    if self._plasmid_annotate == nil then
      self._plasmid_annotate = EntityMod.new(self, nil)
    end
    return self._plasmid_annotate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PlasmidDeepAnnotate():list() / client:PlasmidDeepAnnotate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:PlasmidDeepAnnotate(data)
  local EntityMod = require("entity.plasmid_deep_annotate_entity")
  if data == nil then
    if self._plasmid_deep_annotate == nil then
      self._plasmid_deep_annotate = EntityMod.new(self, nil)
    end
    return self._plasmid_deep_annotate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PlasmidFullReport():list() / client:PlasmidFullReport():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:PlasmidFullReport(data)
  local EntityMod = require("entity.plasmid_full_report_entity")
  if data == nil then
    if self._plasmid_full_report == nil then
      self._plasmid_full_report = EntityMod.new(self, nil)
    end
    return self._plasmid_full_report
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PlasmidIdentify():list() / client:PlasmidIdentify():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:PlasmidIdentify(data)
  local EntityMod = require("entity.plasmid_identify_entity")
  if data == nil then
    if self._plasmid_identify == nil then
      self._plasmid_identify = EntityMod.new(self, nil)
    end
    return self._plasmid_identify
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PrimeEditingDesign():list() / client:PrimeEditingDesign():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:PrimeEditingDesign(data)
  local EntityMod = require("entity.prime_editing_design_entity")
  if data == nil then
    if self._prime_editing_design == nil then
      self._prime_editing_design = EntityMod.new(self, nil)
    end
    return self._prime_editing_design
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PrimeEditingTwinDesign():list() / client:PrimeEditingTwinDesign():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:PrimeEditingTwinDesign(data)
  local EntityMod = require("entity.prime_editing_twin_design_entity")
  if data == nil then
    if self._prime_editing_twin_design == nil then
      self._prime_editing_twin_design = EntityMod.new(self, nil)
    end
    return self._prime_editing_twin_design
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PrimerDesign():list() / client:PrimerDesign():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:PrimerDesign(data)
  local EntityMod = require("entity.primer_design_entity")
  if data == nil then
    if self._primer_design == nil then
      self._primer_design = EntityMod.new(self, nil)
    end
    return self._primer_design
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PrimerSpecificity():list() / client:PrimerSpecificity():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:PrimerSpecificity(data)
  local EntityMod = require("entity.primer_specificity_entity")
  if data == nil then
    if self._primer_specificity == nil then
      self._primer_specificity = EntityMod.new(self, nil)
    end
    return self._primer_specificity
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ProteaseDigestion():list() / client:ProteaseDigestion():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ProteaseDigestion(data)
  local EntityMod = require("entity.protease_digestion_entity")
  if data == nil then
    if self._protease_digestion == nil then
      self._protease_digestion = EntityMod.new(self, nil)
    end
    return self._protease_digestion
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ProteinAnnotatePoll():list() / client:ProteinAnnotatePoll():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ProteinAnnotatePoll(data)
  local EntityMod = require("entity.protein_annotate_poll_entity")
  if data == nil then
    if self._protein_annotate_poll == nil then
      self._protein_annotate_poll = EntityMod.new(self, nil)
    end
    return self._protein_annotate_poll
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ProteinAnnotateSubmit():list() / client:ProteinAnnotateSubmit():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ProteinAnnotateSubmit(data)
  local EntityMod = require("entity.protein_annotate_submit_entity")
  if data == nil then
    if self._protein_annotate_submit == nil then
      self._protein_annotate_submit = EntityMod.new(self, nil)
    end
    return self._protein_annotate_submit
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ProteinHydrophobicity():list() / client:ProteinHydrophobicity():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ProteinHydrophobicity(data)
  local EntityMod = require("entity.protein_hydrophobicity_entity")
  if data == nil then
    if self._protein_hydrophobicity == nil then
      self._protein_hydrophobicity = EntityMod.new(self, nil)
    end
    return self._protein_hydrophobicity
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ProteinProperty():list() / client:ProteinProperty():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ProteinProperty(data)
  local EntityMod = require("entity.protein_property_entity")
  if data == nil then
    if self._protein_property == nil then
      self._protein_property = EntityMod.new(self, nil)
    end
    return self._protein_property
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:RandomSequence():list() / client:RandomSequence():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:RandomSequence(data)
  local EntityMod = require("entity.random_sequence_entity")
  if data == nil then
    if self._random_sequence == nil then
      self._random_sequence = EntityMod.new(self, nil)
    end
    return self._random_sequence
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:RestrictionSite():list() / client:RestrictionSite():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:RestrictionSite(data)
  local EntityMod = require("entity.restriction_site_entity")
  if data == nil then
    if self._restriction_site == nil then
      self._restriction_site = EntityMod.new(self, nil)
    end
    return self._restriction_site
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ReverseComplement():list() / client:ReverseComplement():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ReverseComplement(data)
  local EntityMod = require("entity.reverse_complement_entity")
  if data == nil then
    if self._reverse_complement == nil then
      self._reverse_complement = EntityMod.new(self, nil)
    end
    return self._reverse_complement
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ReverseTranslate():list() / client:ReverseTranslate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:ReverseTranslate(data)
  local EntityMod = require("entity.reverse_translate_entity")
  if data == nil then
    if self._reverse_translate == nil then
      self._reverse_translate = EntityMod.new(self, nil)
    end
    return self._reverse_translate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:RnaFold():list() / client:RnaFold():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:RnaFold(data)
  local EntityMod = require("entity.rna_fold_entity")
  if data == nil then
    if self._rna_fold == nil then
      self._rna_fold = EntityMod.new(self, nil)
    end
    return self._rna_fold
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SangerVsReference():list() / client:SangerVsReference():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SangerVsReference(data)
  local EntityMod = require("entity.sanger_vs_reference_entity")
  if data == nil then
    if self._sanger_vs_reference == nil then
      self._sanger_vs_reference = EntityMod.new(self, nil)
    end
    return self._sanger_vs_reference
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SavePermalink():list() / client:SavePermalink():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SavePermalink(data)
  local EntityMod = require("entity.save_permalink_entity")
  if data == nil then
    if self._save_permalink == nil then
      self._save_permalink = EntityMod.new(self, nil)
    end
    return self._save_permalink
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SeqfileStat():list() / client:SeqfileStat():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SeqfileStat(data)
  local EntityMod = require("entity.seqfile_stat_entity")
  if data == nil then
    if self._seqfile_stat == nil then
      self._seqfile_stat = EntityMod.new(self, nil)
    end
    return self._seqfile_stat
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SequenceFetch():list() / client:SequenceFetch():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SequenceFetch(data)
  local EntityMod = require("entity.sequence_fetch_entity")
  if data == nil then
    if self._sequence_fetch == nil then
      self._sequence_fetch = EntityMod.new(self, nil)
    end
    return self._sequence_fetch
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SequenceFormatConvert():list() / client:SequenceFormatConvert():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SequenceFormatConvert(data)
  local EntityMod = require("entity.sequence_format_convert_entity")
  if data == nil then
    if self._sequence_format_convert == nil then
      self._sequence_format_convert = EntityMod.new(self, nil)
    end
    return self._sequence_format_convert
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SequenceReport():list() / client:SequenceReport():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SequenceReport(data)
  local EntityMod = require("entity.sequence_report_entity")
  if data == nil then
    if self._sequence_report == nil then
      self._sequence_report = EntityMod.new(self, nil)
    end
    return self._sequence_report
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SequenceSearch():list() / client:SequenceSearch():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SequenceSearch(data)
  local EntityMod = require("entity.sequence_search_entity")
  if data == nil then
    if self._sequence_search == nil then
      self._sequence_search = EntityMod.new(self, nil)
    end
    return self._sequence_search
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SequencingReadbackVerify():list() / client:SequencingReadbackVerify():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SequencingReadbackVerify(data)
  local EntityMod = require("entity.sequencing_readback_verify_entity")
  if data == nil then
    if self._sequencing_readback_verify == nil then
      self._sequencing_readback_verify = EntityMod.new(self, nil)
    end
    return self._sequencing_readback_verify
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SessionCreate():list() / client:SessionCreate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SessionCreate(data)
  local EntityMod = require("entity.session_create_entity")
  if data == nil then
    if self._session_create == nil then
      self._session_create = EntityMod.new(self, nil)
    end
    return self._session_create
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SessionGet():list() / client:SessionGet():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SessionGet(data)
  local EntityMod = require("entity.session_get_entity")
  if data == nil then
    if self._session_get == nil then
      self._session_get = EntityMod.new(self, nil)
    end
    return self._session_get
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SessionRun():list() / client:SessionRun():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SessionRun(data)
  local EntityMod = require("entity.session_run_entity")
  if data == nil then
    if self._session_run == nil then
      self._session_run = EntityMod.new(self, nil)
    end
    return self._session_run
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SessionSet():list() / client:SessionSet():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SessionSet(data)
  local EntityMod = require("entity.session_set_entity")
  if data == nil then
    if self._session_set == nil then
      self._session_set = EntityMod.new(self, nil)
    end
    return self._session_set
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SirnaDesign():list() / client:SirnaDesign():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SirnaDesign(data)
  local EntityMod = require("entity.sirna_design_entity")
  if data == nil then
    if self._sirna_design == nil then
      self._sirna_design = EntityMod.new(self, nil)
    end
    return self._sirna_design
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:SiteDirectedMutagenesi():list() / client:SiteDirectedMutagenesi():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:SiteDirectedMutagenesi(data)
  local EntityMod = require("entity.site_directed_mutagenesi_entity")
  if data == nil then
    if self._site_directed_mutagenesi == nil then
      self._site_directed_mutagenesi = EntityMod.new(self, nil)
    end
    return self._site_directed_mutagenesi
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Translate():list() / client:Translate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:Translate(data)
  local EntityMod = require("entity.translate_entity")
  if data == nil then
    if self._translate == nil then
      self._translate = EntityMod.new(self, nil)
    end
    return self._translate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:VariantAnnotate():list() / client:VariantAnnotate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:VariantAnnotate(data)
  local EntityMod = require("entity.variant_annotate_entity")
  if data == nil then
    if self._variant_annotate == nil then
      self._variant_annotate = EntityMod.new(self, nil)
    end
    return self._variant_annotate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:VariantComparator():list() / client:VariantComparator():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:VariantComparator(data)
  local EntityMod = require("entity.variant_comparator_entity")
  if data == nil then
    if self._variant_comparator == nil then
      self._variant_comparator = EntityMod.new(self, nil)
    end
    return self._variant_comparator
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:VerifyAssembly():list() / client:VerifyAssembly():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:VerifyAssembly(data)
  local EntityMod = require("entity.verify_assembly_entity")
  if data == nil then
    if self._verify_assembly == nil then
      self._verify_assembly = EntityMod.new(self, nil)
    end
    return self._verify_assembly
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:VerifyConstruct():list() / client:VerifyConstruct():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:VerifyConstruct(data)
  local EntityMod = require("entity.verify_construct_entity")
  if data == nil then
    if self._verify_construct == nil then
      self._verify_construct = EntityMod.new(self, nil)
    end
    return self._verify_construct
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:VirtualGel():list() / client:VirtualGel():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:VirtualGel(data)
  local EntityMod = require("entity.virtual_gel_entity")
  if data == nil then
    if self._virtual_gel == nil then
      self._virtual_gel = EntityMod.new(self, nil)
    end
    return self._virtual_gel
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:VolcanoPlotData():list() / client:VolcanoPlotData():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:VolcanoPlotData(data)
  local EntityMod = require("entity.volcano_plot_data_entity")
  if data == nil then
    if self._volcano_plot_data == nil then
      self._volcano_plot_data = EntityMod.new(self, nil)
    end
    return self._volcano_plot_data
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:WebSearch():list() / client:WebSearch():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function SeqbenchMcpSDK:WebSearch(data)
  local EntityMod = require("entity.web_search_entity")
  if data == nil then
    if self._web_search == nil then
      self._web_search = EntityMod.new(self, nil)
    end
    return self._web_search
  end
  return EntityMod.new(self, data)
end




function SeqbenchMcpSDK.test(testopts, sdkopts)
  sdkopts = sdkopts or {}
  sdkopts = vs.clone(sdkopts)
  if type(sdkopts) ~= "table" then
    sdkopts = {}
  end

  testopts = testopts or {}
  testopts = vs.clone(testopts)
  if type(testopts) ~= "table" then
    testopts = {}
  end
  testopts["active"] = true

  vs.setpath(sdkopts, "feature.test", testopts)

  local sdk = SeqbenchMcpSDK.new(sdkopts)
  sdk.mode = "test"

  return sdk
end


return SeqbenchMcpSDK
