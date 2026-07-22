# SeqbenchMcp SDK

from utility.voxgig_struct import voxgig_struct as vs
from core.utility_type import SeqbenchMcpUtility
from core.spec import SeqbenchMcpSpec
from core import helpers

# Load utility registration (populates Utility._registrar)
from utility import register

# Load features
from feature.base_feature import SeqbenchMcpBaseFeature
from features import _make_feature


class SeqbenchMcpSDK:

    def __init__(self, options=None):
        self.mode = "live"
        self.features = []
        self.options = None

        utility = SeqbenchMcpUtility()
        self._utility = utility

        from config import make_config
        config = make_config()

        self._rootctx = utility.make_context({
            "client": self,
            "utility": utility,
            "config": config,
            "options": options if options is not None else {},
            "shared": {},
        }, None)

        self.options = utility.make_options(self._rootctx)

        if vs.getpath(self.options, "feature.test.active") is True:
            self.mode = "test"

        self._rootctx.options = self.options

        # Add features in the resolved order (make_options puts an explicit
        # list order first, else defaults to test-first). Ordering matters: the
        # `test` feature installs the base mock transport and the transport
        # features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        # current, so `test` must be added before them to sit at the base.
        feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
        if feature_opts is not None:
            featureorder = vs.getpath(self.options, "__derived__.featureorder")
            if isinstance(featureorder, list):
                for fname in featureorder:
                    fopts = helpers.to_map(feature_opts.get(fname))
                    if fopts is not None and fopts.get("active") is True:
                        utility.feature_add(self._rootctx, _make_feature(fname))

        # Add extension features.
        extend = vs.getprop(self.options, "extend")
        if isinstance(extend, list):
            for f in extend:
                if isinstance(f, dict) or (hasattr(f, "get_name") and callable(f.get_name)):
                    utility.feature_add(self._rootctx, f)

        # Initialize features.
        for f in self.features:
            utility.feature_init(self._rootctx, f)

        utility.feature_hook(self._rootctx, "PostConstruct")

        # #BuildFeatures

    def options_map(self):
        out = vs.clone(self.options)
        if isinstance(out, dict):
            return out
        return {}

    def get_utility(self):
        return SeqbenchMcpUtility.copy(self._utility)

    def get_root_ctx(self):
        return self._rootctx

    def prepare(self, fetchargs=None):
        utility = self._utility

        if fetchargs is None:
            fetchargs = {}

        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "prepare",
            "ctrl": ctrl,
        }, self._rootctx)

        options = self.options

        path = vs.getprop(fetchargs, "path") or ""
        if not isinstance(path, str):
            path = ""

        method = vs.getprop(fetchargs, "method") or "GET"
        if not isinstance(method, str):
            method = "GET"

        params = helpers.to_map(vs.getprop(fetchargs, "params"))
        if params is None:
            params = {}
        query = helpers.to_map(vs.getprop(fetchargs, "query"))
        if query is None:
            query = {}

        headers = utility.prepare_headers(ctx)

        base = vs.getprop(options, "base") or ""
        if not isinstance(base, str):
            base = ""
        prefix = vs.getprop(options, "prefix") or ""
        if not isinstance(prefix, str):
            prefix = ""
        suffix = vs.getprop(options, "suffix") or ""
        if not isinstance(suffix, str):
            suffix = ""

        ctx.spec = SeqbenchMcpSpec({
            "base": base,
            "prefix": prefix,
            "suffix": suffix,
            "path": path,
            "method": method,
            "params": params,
            "query": query,
            "headers": headers,
            "body": vs.getprop(fetchargs, "body"),
            "step": "start",
        })

        # Merge user-provided headers.
        uh = vs.getprop(fetchargs, "headers")
        if isinstance(uh, dict):
            for k, v in uh.items():
                ctx.spec.headers[k] = v

        _, err = utility.prepare_auth(ctx)
        if err is not None:
            raise err

        fetchdef, err = utility.make_fetch_def(ctx)
        if err is not None:
            raise err

        return fetchdef

    def direct(self, fetchargs=None):
        utility = self._utility

        try:
            fetchdef = self.prepare(fetchargs)
        except Exception as err:
            # direct() is the raw-HTTP escape hatch: it never raises, it
            # returns a result object callers branch on via result["ok"].
            return {"ok": False, "err": err}

        if fetchargs is None:
            fetchargs = {}
        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "direct",
            "ctrl": ctrl,
        }, self._rootctx)

        url = fetchdef.get("url", "")
        fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

        if fetch_err is not None:
            return {"ok": False, "err": fetch_err}

        if fetched is None:
            return {
                "ok": False,
                "err": ctx.make_error("direct_no_response", "response: undefined"),
            }

        if isinstance(fetched, dict):
            status = helpers.to_int(vs.getprop(fetched, "status"))
            headers = vs.getprop(fetched, "headers") or {}

            # No-body responses (204, 304) and explicit zero content-length
            # must skip JSON parsing — calling json() on an empty body raises.
            content_length = None
            if isinstance(headers, dict):
                content_length = headers.get("content-length")
            no_body = status in (204, 304) or str(content_length) == "0"

            json_data = None
            if not no_body:
                jf = vs.getprop(fetched, "json")
                if callable(jf):
                    try:
                        json_data = jf()
                    except Exception:
                        # Non-JSON body (e.g. text/plain, text/html). Surface
                        # status + headers but leave data as None.
                        json_data = None

            return {
                "ok": status >= 200 and status < 300,
                "status": status,
                "headers": headers,
                "data": json_data,
            }

        return {
            "ok": False,
            "err": ctx.make_error("direct_invalid", "invalid response type"),
        }


    def AlphafoldLookup(self, data=None) -> "AlphafoldLookupEntity":
        """Entity factory: client.AlphafoldLookup().list() / client.AlphafoldLookup().load({"id": ...})."""
        from entity.alphafold_lookup_entity import AlphafoldLookupEntity
        return AlphafoldLookupEntity(self, data)


    def AsoDesign(self, data=None) -> "AsoDesignEntity":
        """Entity factory: client.AsoDesign().list() / client.AsoDesign().load({"id": ...})."""
        from entity.aso_design_entity import AsoDesignEntity
        return AsoDesignEntity(self, data)


    def BaseEditingDesign(self, data=None) -> "BaseEditingDesignEntity":
        """Entity factory: client.BaseEditingDesign().list() / client.BaseEditingDesign().load({"id": ...})."""
        from entity.base_editing_design_entity import BaseEditingDesignEntity
        return BaseEditingDesignEntity(self, data)


    def Batch(self, data=None) -> "BatchEntity":
        """Entity factory: client.Batch().list() / client.Batch().load({"id": ...})."""
        from entity.batch_entity import BatchEntity
        return BatchEntity(self, data)


    def BatchWorkflow(self, data=None) -> "BatchWorkflowEntity":
        """Entity factory: client.BatchWorkflow().list() / client.BatchWorkflow().load({"id": ...})."""
        from entity.batch__workflow_entity import BatchWorkflowEntity
        return BatchWorkflowEntity(self, data)


    def CharacterizeSequence(self, data=None) -> "CharacterizeSequenceEntity":
        """Entity factory: client.CharacterizeSequence().list() / client.CharacterizeSequence().load({"id": ...})."""
        from entity.characterize_sequence_entity import CharacterizeSequenceEntity
        return CharacterizeSequenceEntity(self, data)


    def CloningSimulate(self, data=None) -> "CloningSimulateEntity":
        """Entity factory: client.CloningSimulate().list() / client.CloningSimulate().load({"id": ...})."""
        from entity.cloning_simulate_entity import CloningSimulateEntity
        return CloningSimulateEntity(self, data)


    def CodonAdaptationIndex(self, data=None) -> "CodonAdaptationIndexEntity":
        """Entity factory: client.CodonAdaptationIndex().list() / client.CodonAdaptationIndex().load({"id": ...})."""
        from entity.codon_adaptation_index_entity import CodonAdaptationIndexEntity
        return CodonAdaptationIndexEntity(self, data)


    def CodonOptimize(self, data=None) -> "CodonOptimizeEntity":
        """Entity factory: client.CodonOptimize().list() / client.CodonOptimize().load({"id": ...})."""
        from entity.codon_optimize_entity import CodonOptimizeEntity
        return CodonOptimizeEntity(self, data)


    def ConstructAutofix(self, data=None) -> "ConstructAutofixEntity":
        """Entity factory: client.ConstructAutofix().list() / client.ConstructAutofix().load({"id": ...})."""
        from entity.construct_autofix_entity import ConstructAutofixEntity
        return ConstructAutofixEntity(self, data)


    def ConstructQc(self, data=None) -> "ConstructQcEntity":
        """Entity factory: client.ConstructQc().list() / client.ConstructQc().load({"id": ...})."""
        from entity.construct_qc_entity import ConstructQcEntity
        return ConstructQcEntity(self, data)


    def CrisprGrnaDesign(self, data=None) -> "CrisprGrnaDesignEntity":
        """Entity factory: client.CrisprGrnaDesign().list() / client.CrisprGrnaDesign().load({"id": ...})."""
        from entity.crispr_grna_design_entity import CrisprGrnaDesignEntity
        return CrisprGrnaDesignEntity(self, data)


    def CrisprHdrDonor(self, data=None) -> "CrisprHdrDonorEntity":
        """Entity factory: client.CrisprHdrDonor().list() / client.CrisprHdrDonor().load({"id": ...})."""
        from entity.crispr_hdr_donor_entity import CrisprHdrDonorEntity
        return CrisprHdrDonorEntity(self, data)


    def CrisprOfftargetCheck(self, data=None) -> "CrisprOfftargetCheckEntity":
        """Entity factory: client.CrisprOfftargetCheck().list() / client.CrisprOfftargetCheck().load({"id": ...})."""
        from entity.crispr_offtarget_check_entity import CrisprOfftargetCheckEntity
        return CrisprOfftargetCheckEntity(self, data)


    def CrossDimer(self, data=None) -> "CrossDimerEntity":
        """Entity factory: client.CrossDimer().list() / client.CrossDimer().load({"id": ...})."""
        from entity.cross_dimer_entity import CrossDimerEntity
        return CrossDimerEntity(self, data)


    def DnaMolarity(self, data=None) -> "DnaMolarityEntity":
        """Entity factory: client.DnaMolarity().list() / client.DnaMolarity().load({"id": ...})."""
        from entity.dna_molarity_entity import DnaMolarityEntity
        return DnaMolarityEntity(self, data)


    def DoubleDigest(self, data=None) -> "DoubleDigestEntity":
        """Entity factory: client.DoubleDigest().list() / client.DoubleDigest().load({"id": ...})."""
        from entity.double_digest_entity import DoubleDigestEntity
        return DoubleDigestEntity(self, data)


    def ExportEchoPicklist(self, data=None) -> "ExportEchoPicklistEntity":
        """Entity factory: client.ExportEchoPicklist().list() / client.ExportEchoPicklist().load({"id": ...})."""
        from entity.export_echo_picklist_entity import ExportEchoPicklistEntity
        return ExportEchoPicklistEntity(self, data)


    def ExportOpentronsProtocol(self, data=None) -> "ExportOpentronsProtocolEntity":
        """Entity factory: client.ExportOpentronsProtocol().list() / client.ExportOpentronsProtocol().load({"id": ...})."""
        from entity.export_opentrons_protocol_entity import ExportOpentronsProtocolEntity
        return ExportOpentronsProtocolEntity(self, data)


    def ExportPlateLayout(self, data=None) -> "ExportPlateLayoutEntity":
        """Entity factory: client.ExportPlateLayout().list() / client.ExportPlateLayout().load({"id": ...})."""
        from entity.export_plate_layout_entity import ExportPlateLayoutEntity
        return ExportPlateLayoutEntity(self, data)


    def ExpressionHeatmapCluster(self, data=None) -> "ExpressionHeatmapClusterEntity":
        """Entity factory: client.ExpressionHeatmapCluster().list() / client.ExpressionHeatmapCluster().load({"id": ...})."""
        from entity.expression_heatmap_cluster_entity import ExpressionHeatmapClusterEntity
        return ExpressionHeatmapClusterEntity(self, data)


    def FastqQcReport(self, data=None) -> "FastqQcReportEntity":
        """Entity factory: client.FastqQcReport().list() / client.FastqQcReport().load({"id": ...})."""
        from entity.fastq_qc_report_entity import FastqQcReportEntity
        return FastqQcReportEntity(self, data)


    def FastqTrim(self, data=None) -> "FastqTrimEntity":
        """Entity factory: client.FastqTrim().list() / client.FastqTrim().load({"id": ...})."""
        from entity.fastq_trim_entity import FastqTrimEntity
        return FastqTrimEntity(self, data)


    def FindOrf(self, data=None) -> "FindOrfEntity":
        """Entity factory: client.FindOrf().list() / client.FindOrf().load({"id": ...})."""
        from entity.find_orf_entity import FindOrfEntity
        return FindOrfEntity(self, data)


    def FormatSequence(self, data=None) -> "FormatSequenceEntity":
        """Entity factory: client.FormatSequence().list() / client.FormatSequence().load({"id": ...})."""
        from entity.format_sequence_entity import FormatSequenceEntity
        return FormatSequenceEntity(self, data)


    def FunctionalEnrichment(self, data=None) -> "FunctionalEnrichmentEntity":
        """Entity factory: client.FunctionalEnrichment().list() / client.FunctionalEnrichment().load({"id": ...})."""
        from entity.functional_enrichment_entity import FunctionalEnrichmentEntity
        return FunctionalEnrichmentEntity(self, data)


    def GcContent(self, data=None) -> "GcContentEntity":
        """Entity factory: client.GcContent().list() / client.GcContent().load({"id": ...})."""
        from entity.gc_content_entity import GcContentEntity
        return GcContentEntity(self, data)


    def GeneDossier(self, data=None) -> "GeneDossierEntity":
        """Entity factory: client.GeneDossier().list() / client.GeneDossier().load({"id": ...})."""
        from entity.gene_dossier_entity import GeneDossierEntity
        return GeneDossierEntity(self, data)


    def GeneExpression(self, data=None) -> "GeneExpressionEntity":
        """Entity factory: client.GeneExpression().list() / client.GeneExpression().load({"id": ...})."""
        from entity.gene_expression_entity import GeneExpressionEntity
        return GeneExpressionEntity(self, data)


    def GeneModel(self, data=None) -> "GeneModelEntity":
        """Entity factory: client.GeneModel().list() / client.GeneModel().load({"id": ...})."""
        from entity.gene_model_entity import GeneModelEntity
        return GeneModelEntity(self, data)


    def GoldenGateFidelity(self, data=None) -> "GoldenGateFidelityEntity":
        """Entity factory: client.GoldenGateFidelity().list() / client.GoldenGateFidelity().load({"id": ...})."""
        from entity.golden_gate_fidelity_entity import GoldenGateFidelityEntity
        return GoldenGateFidelityEntity(self, data)


    def HgvsConvert(self, data=None) -> "HgvsConvertEntity":
        """Entity factory: client.HgvsConvert().list() / client.HgvsConvert().load({"id": ...})."""
        from entity.hgvs_convert_entity import HgvsConvertEntity
        return HgvsConvertEntity(self, data)


    def IdMapPoll(self, data=None) -> "IdMapPollEntity":
        """Entity factory: client.IdMapPoll().list() / client.IdMapPoll().load({"id": ...})."""
        from entity.id_map_poll_entity import IdMapPollEntity
        return IdMapPollEntity(self, data)


    def IdMapSubmit(self, data=None) -> "IdMapSubmitEntity":
        """Entity factory: client.IdMapSubmit().list() / client.IdMapSubmit().load({"id": ...})."""
        from entity.id_map_submit_entity import IdMapSubmitEntity
        return IdMapSubmitEntity(self, data)


    def InSilicoPcr(self, data=None) -> "InSilicoPcrEntity":
        """Entity factory: client.InSilicoPcr().list() / client.InSilicoPcr().load({"id": ...})."""
        from entity.in_silico_pcr_entity import InSilicoPcrEntity
        return InSilicoPcrEntity(self, data)


    def KaspPrimerDesign(self, data=None) -> "KaspPrimerDesignEntity":
        """Entity factory: client.KaspPrimerDesign().list() / client.KaspPrimerDesign().load({"id": ...})."""
        from entity.kasp_primer_design_entity import KaspPrimerDesignEntity
        return KaspPrimerDesignEntity(self, data)


    def ListTool(self, data=None) -> "ListToolEntity":
        """Entity factory: client.ListTool().list() / client.ListTool().load({"id": ...})."""
        from entity.list_tool_entity import ListToolEntity
        return ListToolEntity(self, data)


    def MeltingTemperature(self, data=None) -> "MeltingTemperatureEntity":
        """Entity factory: client.MeltingTemperature().list() / client.MeltingTemperature().load({"id": ...})."""
        from entity.melting_temperature_entity import MeltingTemperatureEntity
        return MeltingTemperatureEntity(self, data)


    def MotifFinder(self, data=None) -> "MotifFinderEntity":
        """Entity factory: client.MotifFinder().list() / client.MotifFinder().load({"id": ...})."""
        from entity.motif_finder_entity import MotifFinderEntity
        return MotifFinderEntity(self, data)


    def MultipleSequenceAlignment(self, data=None) -> "MultipleSequenceAlignmentEntity":
        """Entity factory: client.MultipleSequenceAlignment().list() / client.MultipleSequenceAlignment().load({"id": ...})."""
        from entity.multiple_sequence_alignment_entity import MultipleSequenceAlignmentEntity
        return MultipleSequenceAlignmentEntity(self, data)


    def OligoAnalysi(self, data=None) -> "OligoAnalysiEntity":
        """Entity factory: client.OligoAnalysi().list() / client.OligoAnalysi().load({"id": ...})."""
        from entity.oligo_analysi_entity import OligoAnalysiEntity
        return OligoAnalysiEntity(self, data)


    def OrthologMap(self, data=None) -> "OrthologMapEntity":
        """Entity factory: client.OrthologMap().list() / client.OrthologMap().load({"id": ...})."""
        from entity.ortholog_map_entity import OrthologMapEntity
        return OrthologMapEntity(self, data)


    def PairwiseAlignment(self, data=None) -> "PairwiseAlignmentEntity":
        """Entity factory: client.PairwiseAlignment().list() / client.PairwiseAlignment().load({"id": ...})."""
        from entity.pairwise_alignment_entity import PairwiseAlignmentEntity
        return PairwiseAlignmentEntity(self, data)


    def ParseGenbank(self, data=None) -> "ParseGenbankEntity":
        """Entity factory: client.ParseGenbank().list() / client.ParseGenbank().load({"id": ...})."""
        from entity.parse_genbank_entity import ParseGenbankEntity
        return ParseGenbankEntity(self, data)


    def ParseSangerTrace(self, data=None) -> "ParseSangerTraceEntity":
        """Entity factory: client.ParseSangerTrace().list() / client.ParseSangerTrace().load({"id": ...})."""
        from entity.parse_sanger_trace_entity import ParseSangerTraceEntity
        return ParseSangerTraceEntity(self, data)


    def PlasmidAnnotate(self, data=None) -> "PlasmidAnnotateEntity":
        """Entity factory: client.PlasmidAnnotate().list() / client.PlasmidAnnotate().load({"id": ...})."""
        from entity.plasmid_annotate_entity import PlasmidAnnotateEntity
        return PlasmidAnnotateEntity(self, data)


    def PlasmidDeepAnnotate(self, data=None) -> "PlasmidDeepAnnotateEntity":
        """Entity factory: client.PlasmidDeepAnnotate().list() / client.PlasmidDeepAnnotate().load({"id": ...})."""
        from entity.plasmid_deep_annotate_entity import PlasmidDeepAnnotateEntity
        return PlasmidDeepAnnotateEntity(self, data)


    def PlasmidFullReport(self, data=None) -> "PlasmidFullReportEntity":
        """Entity factory: client.PlasmidFullReport().list() / client.PlasmidFullReport().load({"id": ...})."""
        from entity.plasmid_full_report_entity import PlasmidFullReportEntity
        return PlasmidFullReportEntity(self, data)


    def PlasmidIdentify(self, data=None) -> "PlasmidIdentifyEntity":
        """Entity factory: client.PlasmidIdentify().list() / client.PlasmidIdentify().load({"id": ...})."""
        from entity.plasmid_identify_entity import PlasmidIdentifyEntity
        return PlasmidIdentifyEntity(self, data)


    def PrimeEditingDesign(self, data=None) -> "PrimeEditingDesignEntity":
        """Entity factory: client.PrimeEditingDesign().list() / client.PrimeEditingDesign().load({"id": ...})."""
        from entity.prime_editing_design_entity import PrimeEditingDesignEntity
        return PrimeEditingDesignEntity(self, data)


    def PrimeEditingTwinDesign(self, data=None) -> "PrimeEditingTwinDesignEntity":
        """Entity factory: client.PrimeEditingTwinDesign().list() / client.PrimeEditingTwinDesign().load({"id": ...})."""
        from entity.prime_editing_twin_design_entity import PrimeEditingTwinDesignEntity
        return PrimeEditingTwinDesignEntity(self, data)


    def PrimerDesign(self, data=None) -> "PrimerDesignEntity":
        """Entity factory: client.PrimerDesign().list() / client.PrimerDesign().load({"id": ...})."""
        from entity.primer_design_entity import PrimerDesignEntity
        return PrimerDesignEntity(self, data)


    def PrimerSpecificity(self, data=None) -> "PrimerSpecificityEntity":
        """Entity factory: client.PrimerSpecificity().list() / client.PrimerSpecificity().load({"id": ...})."""
        from entity.primer_specificity_entity import PrimerSpecificityEntity
        return PrimerSpecificityEntity(self, data)


    def ProteaseDigestion(self, data=None) -> "ProteaseDigestionEntity":
        """Entity factory: client.ProteaseDigestion().list() / client.ProteaseDigestion().load({"id": ...})."""
        from entity.protease_digestion_entity import ProteaseDigestionEntity
        return ProteaseDigestionEntity(self, data)


    def ProteinAnnotatePoll(self, data=None) -> "ProteinAnnotatePollEntity":
        """Entity factory: client.ProteinAnnotatePoll().list() / client.ProteinAnnotatePoll().load({"id": ...})."""
        from entity.protein_annotate_poll_entity import ProteinAnnotatePollEntity
        return ProteinAnnotatePollEntity(self, data)


    def ProteinAnnotateSubmit(self, data=None) -> "ProteinAnnotateSubmitEntity":
        """Entity factory: client.ProteinAnnotateSubmit().list() / client.ProteinAnnotateSubmit().load({"id": ...})."""
        from entity.protein_annotate_submit_entity import ProteinAnnotateSubmitEntity
        return ProteinAnnotateSubmitEntity(self, data)


    def ProteinHydrophobicity(self, data=None) -> "ProteinHydrophobicityEntity":
        """Entity factory: client.ProteinHydrophobicity().list() / client.ProteinHydrophobicity().load({"id": ...})."""
        from entity.protein_hydrophobicity_entity import ProteinHydrophobicityEntity
        return ProteinHydrophobicityEntity(self, data)


    def ProteinProperty(self, data=None) -> "ProteinPropertyEntity":
        """Entity factory: client.ProteinProperty().list() / client.ProteinProperty().load({"id": ...})."""
        from entity.protein_property_entity import ProteinPropertyEntity
        return ProteinPropertyEntity(self, data)


    def RandomSequence(self, data=None) -> "RandomSequenceEntity":
        """Entity factory: client.RandomSequence().list() / client.RandomSequence().load({"id": ...})."""
        from entity.random_sequence_entity import RandomSequenceEntity
        return RandomSequenceEntity(self, data)


    def RestrictionSite(self, data=None) -> "RestrictionSiteEntity":
        """Entity factory: client.RestrictionSite().list() / client.RestrictionSite().load({"id": ...})."""
        from entity.restriction_site_entity import RestrictionSiteEntity
        return RestrictionSiteEntity(self, data)


    def ReverseComplement(self, data=None) -> "ReverseComplementEntity":
        """Entity factory: client.ReverseComplement().list() / client.ReverseComplement().load({"id": ...})."""
        from entity.reverse_complement_entity import ReverseComplementEntity
        return ReverseComplementEntity(self, data)


    def ReverseTranslate(self, data=None) -> "ReverseTranslateEntity":
        """Entity factory: client.ReverseTranslate().list() / client.ReverseTranslate().load({"id": ...})."""
        from entity.reverse_translate_entity import ReverseTranslateEntity
        return ReverseTranslateEntity(self, data)


    def RnaFold(self, data=None) -> "RnaFoldEntity":
        """Entity factory: client.RnaFold().list() / client.RnaFold().load({"id": ...})."""
        from entity.rna_fold_entity import RnaFoldEntity
        return RnaFoldEntity(self, data)


    def SangerVsReference(self, data=None) -> "SangerVsReferenceEntity":
        """Entity factory: client.SangerVsReference().list() / client.SangerVsReference().load({"id": ...})."""
        from entity.sanger_vs_reference_entity import SangerVsReferenceEntity
        return SangerVsReferenceEntity(self, data)


    def SavePermalink(self, data=None) -> "SavePermalinkEntity":
        """Entity factory: client.SavePermalink().list() / client.SavePermalink().load({"id": ...})."""
        from entity.save_permalink_entity import SavePermalinkEntity
        return SavePermalinkEntity(self, data)


    def SeqfileStat(self, data=None) -> "SeqfileStatEntity":
        """Entity factory: client.SeqfileStat().list() / client.SeqfileStat().load({"id": ...})."""
        from entity.seqfile_stat_entity import SeqfileStatEntity
        return SeqfileStatEntity(self, data)


    def SequenceFetch(self, data=None) -> "SequenceFetchEntity":
        """Entity factory: client.SequenceFetch().list() / client.SequenceFetch().load({"id": ...})."""
        from entity.sequence_fetch_entity import SequenceFetchEntity
        return SequenceFetchEntity(self, data)


    def SequenceFormatConvert(self, data=None) -> "SequenceFormatConvertEntity":
        """Entity factory: client.SequenceFormatConvert().list() / client.SequenceFormatConvert().load({"id": ...})."""
        from entity.sequence_format_convert_entity import SequenceFormatConvertEntity
        return SequenceFormatConvertEntity(self, data)


    def SequenceReport(self, data=None) -> "SequenceReportEntity":
        """Entity factory: client.SequenceReport().list() / client.SequenceReport().load({"id": ...})."""
        from entity.sequence_report_entity import SequenceReportEntity
        return SequenceReportEntity(self, data)


    def SequenceSearch(self, data=None) -> "SequenceSearchEntity":
        """Entity factory: client.SequenceSearch().list() / client.SequenceSearch().load({"id": ...})."""
        from entity.sequence_search_entity import SequenceSearchEntity
        return SequenceSearchEntity(self, data)


    def SequencingReadbackVerify(self, data=None) -> "SequencingReadbackVerifyEntity":
        """Entity factory: client.SequencingReadbackVerify().list() / client.SequencingReadbackVerify().load({"id": ...})."""
        from entity.sequencing_readback_verify_entity import SequencingReadbackVerifyEntity
        return SequencingReadbackVerifyEntity(self, data)


    def SessionCreate(self, data=None) -> "SessionCreateEntity":
        """Entity factory: client.SessionCreate().list() / client.SessionCreate().load({"id": ...})."""
        from entity.session_create_entity import SessionCreateEntity
        return SessionCreateEntity(self, data)


    def SessionGet(self, data=None) -> "SessionGetEntity":
        """Entity factory: client.SessionGet().list() / client.SessionGet().load({"id": ...})."""
        from entity.session_get_entity import SessionGetEntity
        return SessionGetEntity(self, data)


    def SessionRun(self, data=None) -> "SessionRunEntity":
        """Entity factory: client.SessionRun().list() / client.SessionRun().load({"id": ...})."""
        from entity.session_run_entity import SessionRunEntity
        return SessionRunEntity(self, data)


    def SessionSet(self, data=None) -> "SessionSetEntity":
        """Entity factory: client.SessionSet().list() / client.SessionSet().load({"id": ...})."""
        from entity.session_set_entity import SessionSetEntity
        return SessionSetEntity(self, data)


    def SirnaDesign(self, data=None) -> "SirnaDesignEntity":
        """Entity factory: client.SirnaDesign().list() / client.SirnaDesign().load({"id": ...})."""
        from entity.sirna_design_entity import SirnaDesignEntity
        return SirnaDesignEntity(self, data)


    def SiteDirectedMutagenesi(self, data=None) -> "SiteDirectedMutagenesiEntity":
        """Entity factory: client.SiteDirectedMutagenesi().list() / client.SiteDirectedMutagenesi().load({"id": ...})."""
        from entity.site_directed_mutagenesi_entity import SiteDirectedMutagenesiEntity
        return SiteDirectedMutagenesiEntity(self, data)


    def Translate(self, data=None) -> "TranslateEntity":
        """Entity factory: client.Translate().list() / client.Translate().load({"id": ...})."""
        from entity.translate_entity import TranslateEntity
        return TranslateEntity(self, data)


    def VariantAnnotate(self, data=None) -> "VariantAnnotateEntity":
        """Entity factory: client.VariantAnnotate().list() / client.VariantAnnotate().load({"id": ...})."""
        from entity.variant_annotate_entity import VariantAnnotateEntity
        return VariantAnnotateEntity(self, data)


    def VariantComparator(self, data=None) -> "VariantComparatorEntity":
        """Entity factory: client.VariantComparator().list() / client.VariantComparator().load({"id": ...})."""
        from entity.variant_comparator_entity import VariantComparatorEntity
        return VariantComparatorEntity(self, data)


    def VerifyAssembly(self, data=None) -> "VerifyAssemblyEntity":
        """Entity factory: client.VerifyAssembly().list() / client.VerifyAssembly().load({"id": ...})."""
        from entity.verify_assembly_entity import VerifyAssemblyEntity
        return VerifyAssemblyEntity(self, data)


    def VerifyConstruct(self, data=None) -> "VerifyConstructEntity":
        """Entity factory: client.VerifyConstruct().list() / client.VerifyConstruct().load({"id": ...})."""
        from entity.verify_construct_entity import VerifyConstructEntity
        return VerifyConstructEntity(self, data)


    def VirtualGel(self, data=None) -> "VirtualGelEntity":
        """Entity factory: client.VirtualGel().list() / client.VirtualGel().load({"id": ...})."""
        from entity.virtual_gel_entity import VirtualGelEntity
        return VirtualGelEntity(self, data)


    def VolcanoPlotData(self, data=None) -> "VolcanoPlotDataEntity":
        """Entity factory: client.VolcanoPlotData().list() / client.VolcanoPlotData().load({"id": ...})."""
        from entity.volcano_plot_data_entity import VolcanoPlotDataEntity
        return VolcanoPlotDataEntity(self, data)


    def WebSearch(self, data=None) -> "WebSearchEntity":
        """Entity factory: client.WebSearch().list() / client.WebSearch().load({"id": ...})."""
        from entity.web_search_entity import WebSearchEntity
        return WebSearchEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None) -> "SeqbenchMcpSDK":
        if sdkopts is None:
            sdkopts = {}
        sdkopts = vs.clone(sdkopts)
        if not isinstance(sdkopts, dict):
            sdkopts = {}

        if testopts is None:
            testopts = {}
        testopts = vs.clone(testopts)
        if not isinstance(testopts, dict):
            testopts = {}
        testopts["active"] = True

        vs.setpath(sdkopts, "feature.test", testopts)

        sdk = cls(sdkopts)
        sdk.mode = "test"

        return sdk


from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from entity.alphafold_lookup_entity import AlphafoldLookupEntity
    from entity.aso_design_entity import AsoDesignEntity
    from entity.base_editing_design_entity import BaseEditingDesignEntity
    from entity.batch_entity import BatchEntity
    from entity.batch__workflow_entity import BatchWorkflowEntity
    from entity.characterize_sequence_entity import CharacterizeSequenceEntity
    from entity.cloning_simulate_entity import CloningSimulateEntity
    from entity.codon_adaptation_index_entity import CodonAdaptationIndexEntity
    from entity.codon_optimize_entity import CodonOptimizeEntity
    from entity.construct_autofix_entity import ConstructAutofixEntity
    from entity.construct_qc_entity import ConstructQcEntity
    from entity.crispr_grna_design_entity import CrisprGrnaDesignEntity
    from entity.crispr_hdr_donor_entity import CrisprHdrDonorEntity
    from entity.crispr_offtarget_check_entity import CrisprOfftargetCheckEntity
    from entity.cross_dimer_entity import CrossDimerEntity
    from entity.dna_molarity_entity import DnaMolarityEntity
    from entity.double_digest_entity import DoubleDigestEntity
    from entity.export_echo_picklist_entity import ExportEchoPicklistEntity
    from entity.export_opentrons_protocol_entity import ExportOpentronsProtocolEntity
    from entity.export_plate_layout_entity import ExportPlateLayoutEntity
    from entity.expression_heatmap_cluster_entity import ExpressionHeatmapClusterEntity
    from entity.fastq_qc_report_entity import FastqQcReportEntity
    from entity.fastq_trim_entity import FastqTrimEntity
    from entity.find_orf_entity import FindOrfEntity
    from entity.format_sequence_entity import FormatSequenceEntity
    from entity.functional_enrichment_entity import FunctionalEnrichmentEntity
    from entity.gc_content_entity import GcContentEntity
    from entity.gene_dossier_entity import GeneDossierEntity
    from entity.gene_expression_entity import GeneExpressionEntity
    from entity.gene_model_entity import GeneModelEntity
    from entity.golden_gate_fidelity_entity import GoldenGateFidelityEntity
    from entity.hgvs_convert_entity import HgvsConvertEntity
    from entity.id_map_poll_entity import IdMapPollEntity
    from entity.id_map_submit_entity import IdMapSubmitEntity
    from entity.in_silico_pcr_entity import InSilicoPcrEntity
    from entity.kasp_primer_design_entity import KaspPrimerDesignEntity
    from entity.list_tool_entity import ListToolEntity
    from entity.melting_temperature_entity import MeltingTemperatureEntity
    from entity.motif_finder_entity import MotifFinderEntity
    from entity.multiple_sequence_alignment_entity import MultipleSequenceAlignmentEntity
    from entity.oligo_analysi_entity import OligoAnalysiEntity
    from entity.ortholog_map_entity import OrthologMapEntity
    from entity.pairwise_alignment_entity import PairwiseAlignmentEntity
    from entity.parse_genbank_entity import ParseGenbankEntity
    from entity.parse_sanger_trace_entity import ParseSangerTraceEntity
    from entity.plasmid_annotate_entity import PlasmidAnnotateEntity
    from entity.plasmid_deep_annotate_entity import PlasmidDeepAnnotateEntity
    from entity.plasmid_full_report_entity import PlasmidFullReportEntity
    from entity.plasmid_identify_entity import PlasmidIdentifyEntity
    from entity.prime_editing_design_entity import PrimeEditingDesignEntity
    from entity.prime_editing_twin_design_entity import PrimeEditingTwinDesignEntity
    from entity.primer_design_entity import PrimerDesignEntity
    from entity.primer_specificity_entity import PrimerSpecificityEntity
    from entity.protease_digestion_entity import ProteaseDigestionEntity
    from entity.protein_annotate_poll_entity import ProteinAnnotatePollEntity
    from entity.protein_annotate_submit_entity import ProteinAnnotateSubmitEntity
    from entity.protein_hydrophobicity_entity import ProteinHydrophobicityEntity
    from entity.protein_property_entity import ProteinPropertyEntity
    from entity.random_sequence_entity import RandomSequenceEntity
    from entity.restriction_site_entity import RestrictionSiteEntity
    from entity.reverse_complement_entity import ReverseComplementEntity
    from entity.reverse_translate_entity import ReverseTranslateEntity
    from entity.rna_fold_entity import RnaFoldEntity
    from entity.sanger_vs_reference_entity import SangerVsReferenceEntity
    from entity.save_permalink_entity import SavePermalinkEntity
    from entity.seqfile_stat_entity import SeqfileStatEntity
    from entity.sequence_fetch_entity import SequenceFetchEntity
    from entity.sequence_format_convert_entity import SequenceFormatConvertEntity
    from entity.sequence_report_entity import SequenceReportEntity
    from entity.sequence_search_entity import SequenceSearchEntity
    from entity.sequencing_readback_verify_entity import SequencingReadbackVerifyEntity
    from entity.session_create_entity import SessionCreateEntity
    from entity.session_get_entity import SessionGetEntity
    from entity.session_run_entity import SessionRunEntity
    from entity.session_set_entity import SessionSetEntity
    from entity.sirna_design_entity import SirnaDesignEntity
    from entity.site_directed_mutagenesi_entity import SiteDirectedMutagenesiEntity
    from entity.translate_entity import TranslateEntity
    from entity.variant_annotate_entity import VariantAnnotateEntity
    from entity.variant_comparator_entity import VariantComparatorEntity
    from entity.verify_assembly_entity import VerifyAssemblyEntity
    from entity.verify_construct_entity import VerifyConstructEntity
    from entity.virtual_gel_entity import VirtualGelEntity
    from entity.volcano_plot_data_entity import VolcanoPlotDataEntity
    from entity.web_search_entity import WebSearchEntity
