<?php
declare(strict_types=1);

// SeqbenchMcp SDK

require_once __DIR__ . '/utility/struct/Struct.php';
require_once __DIR__ . '/core/UtilityType.php';
require_once __DIR__ . '/core/Spec.php';
require_once __DIR__ . '/core/Helpers.php';

// Load utility registration
require_once __DIR__ . '/utility/Register.php';

// Load config and features
require_once __DIR__ . '/config.php';
require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/features.php';

use Voxgig\Struct\Struct;

// Features record diagnostic state on the client as dynamic properties
// (_retry, _cache, _metrics, ...); allow them explicitly (PHP 8.2+
// deprecates implicit dynamic properties).
#[\AllowDynamicProperties]
class SeqbenchMcpSDK
{
    public string $mode;
    public array $features;
    public ?array $options;

    private $_utility;
    private $_rootctx;

    public function __construct(array $options = [])
    {
        $this->mode = "live";
        $this->features = [];
        $this->options = null;

        $utility = new SeqbenchMcpUtility();
        $this->_utility = $utility;

        $config = SeqbenchMcpConfig::make_config();

        $this->_rootctx = ($utility->make_context)([
            "client" => $this,
            "utility" => $utility,
            "config" => $config,
            "options" => $options ?? [],
            "shared" => [],
        ], null);

        $this->options = ($utility->make_options)($this->_rootctx);

        if (Struct::getpath($this->options, "feature.test.active") === true) {
            $this->mode = "test";
        }

        $this->_rootctx->options = $this->options;

        // Add features in the resolved order (make_options puts an explicit
        // list order first, else defaults to test-first). Ordering matters: the
        // `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        // current, so `test` must be added before them to sit at the base.
        $feature_opts = SeqbenchMcpHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $featureorder = Struct::getpath($this->options, "__derived__.featureorder");
            if (is_array($featureorder)) {
                foreach ($featureorder as $fname) {
                    $fopts = SeqbenchMcpHelpers::to_map($feature_opts[$fname] ?? null);
                    if ($fopts && isset($fopts["active"]) && $fopts["active"] === true) {
                        ($utility->feature_add)($this->_rootctx, SeqbenchMcpFeatures::make_feature($fname));
                    }
                }
            }
        }

        // Add extension features.
        $extend_val = Struct::getprop($this->options, "extend");
        if (is_array($extend_val)) {
            foreach ($extend_val as $f) {
                if (is_object($f) && method_exists($f, 'get_name')) {
                    ($utility->feature_add)($this->_rootctx, $f);
                }
            }
        }

        // Initialize features.
        foreach ($this->features as $f) {
            ($utility->feature_init)($this->_rootctx, $f);
        }

        ($utility->feature_hook)($this->_rootctx, "PostConstruct");
    }

    public function options_map(): array
    {
        $out = Struct::clone($this->options);
        return is_array($out) ? $out : [];
    }

    public function get_utility()
    {
        return SeqbenchMcpUtility::copy($this->_utility);
    }

    public function get_root_ctx()
    {
        return $this->_rootctx;
    }

    public function prepare(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;
        $fetchargs = $fetchargs ?? [];

        $ctrl = SeqbenchMcpHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "prepare",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $opts = $this->options;
        $path = Struct::getprop($fetchargs, "path") ?? "";
        $path = is_string($path) ? $path : "";
        $method_val = Struct::getprop($fetchargs, "method") ?? "GET";
        $method_val = is_string($method_val) ? $method_val : "GET";
        $params = SeqbenchMcpHelpers::to_map(Struct::getprop($fetchargs, "params")) ?? [];
        $query = SeqbenchMcpHelpers::to_map(Struct::getprop($fetchargs, "query")) ?? [];
        $headers = ($utility->prepare_headers)($ctx);

        $base = Struct::getprop($opts, "base") ?? "";
        $base = is_string($base) ? $base : "";
        $prefix = Struct::getprop($opts, "prefix") ?? "";
        $prefix = is_string($prefix) ? $prefix : "";
        $suffix = Struct::getprop($opts, "suffix") ?? "";
        $suffix = is_string($suffix) ? $suffix : "";

        $ctx->spec = new SeqbenchMcpSpec([
            "base" => $base, "prefix" => $prefix, "suffix" => $suffix,
            "path" => $path, "method" => $method_val,
            "params" => $params, "query" => $query, "headers" => $headers,
            "body" => Struct::getprop($fetchargs, "body"),
            "step" => "start",
        ]);

        // Merge user-provided headers.
        $uh = Struct::getprop($fetchargs, "headers");
        if (is_array($uh)) {
            foreach ($uh as $k => $v) {
                $ctx->spec->headers[$k] = $v;
            }
        }

        [$_, $err] = ($utility->prepare_auth)($ctx);
        if ($err) {
            return ($utility->make_error)($ctx, $err);
        }

        [$fetchdef, $fd_err] = ($utility->make_fetch_def)($ctx);
        if ($fd_err) {
            return ($utility->make_error)($ctx, $fd_err);
        }
        return $fetchdef;
    }

    public function direct(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;

        // direct() is the raw-HTTP escape hatch: it never throws, it returns
        // an {ok, err, ...} dict. prepare() now raises on error, so catch it
        // and surface the failure through the dict instead.
        try {
            $fetchdef = $this->prepare($fetchargs);
        } catch (\Throwable $err) {
            return ["ok" => false, "err" => $err];
        }

        $fetchargs = $fetchargs ?? [];
        $ctrl = SeqbenchMcpHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "direct",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $url = $fetchdef["url"] ?? "";
        [$fetched, $fetch_err] = ($utility->fetcher)($ctx, $url, $fetchdef);

        if ($fetch_err) {
            return ["ok" => false, "err" => $fetch_err];
        }

        if ($fetched === null) {
            return [
                "ok" => false,
                "err" => $ctx->make_error("direct_no_response", "response: undefined"),
            ];
        }

        if (is_array($fetched)) {
            $status = SeqbenchMcpHelpers::to_int(Struct::getprop($fetched, "status"));
            $headers = Struct::getprop($fetched, "headers") ?? [];

            // No-body responses (204, 304) and explicit zero content-length
            // must skip JSON parsing — calling json() on an empty body errors.
            $content_length = is_array($headers) ? ($headers["content-length"] ?? null) : null;
            $no_body = $status === 204 || $status === 304 || (string)$content_length === "0";

            $json_data = null;
            if (!$no_body) {
                $jf = Struct::getprop($fetched, "json");
                if (is_callable($jf)) {
                    try {
                        $json_data = $jf();
                    } catch (\Throwable $e) {
                        // Non-JSON body — leave data null but keep status/ok.
                        $json_data = null;
                    }
                }
            }

            return [
                "ok" => $status >= 200 && $status < 300,
                "status" => $status,
                "headers" => Struct::getprop($fetched, "headers"),
                "data" => $json_data,
            ];
        }

        return [
            "ok" => false,
            "err" => $ctx->make_error("direct_invalid", "invalid response type"),
        ];
    }


    private $_alphafold_lookup = null;

    // Canonical facade: $client->AlphafoldLookup()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->alphafold_lookup()
    // resolves here too.
    public function AlphafoldLookup($data = null)
    {
        require_once __DIR__ . '/entity/alphafold_lookup_entity.php';
        if ($data === null) {
            if ($this->_alphafold_lookup === null) {
                $this->_alphafold_lookup = new AlphafoldLookupEntity($this, null);
            }
            return $this->_alphafold_lookup;
        }
        return new AlphafoldLookupEntity($this, $data);
    }


    private $_aso_design = null;

    // Canonical facade: $client->AsoDesign()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->aso_design()
    // resolves here too.
    public function AsoDesign($data = null)
    {
        require_once __DIR__ . '/entity/aso_design_entity.php';
        if ($data === null) {
            if ($this->_aso_design === null) {
                $this->_aso_design = new AsoDesignEntity($this, null);
            }
            return $this->_aso_design;
        }
        return new AsoDesignEntity($this, $data);
    }


    private $_base_editing_design = null;

    // Canonical facade: $client->BaseEditingDesign()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->base_editing_design()
    // resolves here too.
    public function BaseEditingDesign($data = null)
    {
        require_once __DIR__ . '/entity/base_editing_design_entity.php';
        if ($data === null) {
            if ($this->_base_editing_design === null) {
                $this->_base_editing_design = new BaseEditingDesignEntity($this, null);
            }
            return $this->_base_editing_design;
        }
        return new BaseEditingDesignEntity($this, $data);
    }


    private $_batch = null;

    // Canonical facade: $client->Batch()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->batch()
    // resolves here too.
    public function Batch($data = null)
    {
        require_once __DIR__ . '/entity/batch_entity.php';
        if ($data === null) {
            if ($this->_batch === null) {
                $this->_batch = new BatchEntity($this, null);
            }
            return $this->_batch;
        }
        return new BatchEntity($this, $data);
    }


    private $_batch__workflow = null;

    // Canonical facade: $client->BatchWorkflow()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->batch__workflow()
    // resolves here too.
    public function BatchWorkflow($data = null)
    {
        require_once __DIR__ . '/entity/batch__workflow_entity.php';
        if ($data === null) {
            if ($this->_batch__workflow === null) {
                $this->_batch__workflow = new BatchWorkflowEntity($this, null);
            }
            return $this->_batch__workflow;
        }
        return new BatchWorkflowEntity($this, $data);
    }


    private $_characterize_sequence = null;

    // Canonical facade: $client->CharacterizeSequence()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->characterize_sequence()
    // resolves here too.
    public function CharacterizeSequence($data = null)
    {
        require_once __DIR__ . '/entity/characterize_sequence_entity.php';
        if ($data === null) {
            if ($this->_characterize_sequence === null) {
                $this->_characterize_sequence = new CharacterizeSequenceEntity($this, null);
            }
            return $this->_characterize_sequence;
        }
        return new CharacterizeSequenceEntity($this, $data);
    }


    private $_cloning_simulate = null;

    // Canonical facade: $client->CloningSimulate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->cloning_simulate()
    // resolves here too.
    public function CloningSimulate($data = null)
    {
        require_once __DIR__ . '/entity/cloning_simulate_entity.php';
        if ($data === null) {
            if ($this->_cloning_simulate === null) {
                $this->_cloning_simulate = new CloningSimulateEntity($this, null);
            }
            return $this->_cloning_simulate;
        }
        return new CloningSimulateEntity($this, $data);
    }


    private $_codon_adaptation_index = null;

    // Canonical facade: $client->CodonAdaptationIndex()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->codon_adaptation_index()
    // resolves here too.
    public function CodonAdaptationIndex($data = null)
    {
        require_once __DIR__ . '/entity/codon_adaptation_index_entity.php';
        if ($data === null) {
            if ($this->_codon_adaptation_index === null) {
                $this->_codon_adaptation_index = new CodonAdaptationIndexEntity($this, null);
            }
            return $this->_codon_adaptation_index;
        }
        return new CodonAdaptationIndexEntity($this, $data);
    }


    private $_codon_optimize = null;

    // Canonical facade: $client->CodonOptimize()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->codon_optimize()
    // resolves here too.
    public function CodonOptimize($data = null)
    {
        require_once __DIR__ . '/entity/codon_optimize_entity.php';
        if ($data === null) {
            if ($this->_codon_optimize === null) {
                $this->_codon_optimize = new CodonOptimizeEntity($this, null);
            }
            return $this->_codon_optimize;
        }
        return new CodonOptimizeEntity($this, $data);
    }


    private $_construct_autofix = null;

    // Canonical facade: $client->ConstructAutofix()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->construct_autofix()
    // resolves here too.
    public function ConstructAutofix($data = null)
    {
        require_once __DIR__ . '/entity/construct_autofix_entity.php';
        if ($data === null) {
            if ($this->_construct_autofix === null) {
                $this->_construct_autofix = new ConstructAutofixEntity($this, null);
            }
            return $this->_construct_autofix;
        }
        return new ConstructAutofixEntity($this, $data);
    }


    private $_construct_qc = null;

    // Canonical facade: $client->ConstructQc()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->construct_qc()
    // resolves here too.
    public function ConstructQc($data = null)
    {
        require_once __DIR__ . '/entity/construct_qc_entity.php';
        if ($data === null) {
            if ($this->_construct_qc === null) {
                $this->_construct_qc = new ConstructQcEntity($this, null);
            }
            return $this->_construct_qc;
        }
        return new ConstructQcEntity($this, $data);
    }


    private $_crispr_grna_design = null;

    // Canonical facade: $client->CrisprGrnaDesign()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->crispr_grna_design()
    // resolves here too.
    public function CrisprGrnaDesign($data = null)
    {
        require_once __DIR__ . '/entity/crispr_grna_design_entity.php';
        if ($data === null) {
            if ($this->_crispr_grna_design === null) {
                $this->_crispr_grna_design = new CrisprGrnaDesignEntity($this, null);
            }
            return $this->_crispr_grna_design;
        }
        return new CrisprGrnaDesignEntity($this, $data);
    }


    private $_crispr_hdr_donor = null;

    // Canonical facade: $client->CrisprHdrDonor()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->crispr_hdr_donor()
    // resolves here too.
    public function CrisprHdrDonor($data = null)
    {
        require_once __DIR__ . '/entity/crispr_hdr_donor_entity.php';
        if ($data === null) {
            if ($this->_crispr_hdr_donor === null) {
                $this->_crispr_hdr_donor = new CrisprHdrDonorEntity($this, null);
            }
            return $this->_crispr_hdr_donor;
        }
        return new CrisprHdrDonorEntity($this, $data);
    }


    private $_crispr_offtarget_check = null;

    // Canonical facade: $client->CrisprOfftargetCheck()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->crispr_offtarget_check()
    // resolves here too.
    public function CrisprOfftargetCheck($data = null)
    {
        require_once __DIR__ . '/entity/crispr_offtarget_check_entity.php';
        if ($data === null) {
            if ($this->_crispr_offtarget_check === null) {
                $this->_crispr_offtarget_check = new CrisprOfftargetCheckEntity($this, null);
            }
            return $this->_crispr_offtarget_check;
        }
        return new CrisprOfftargetCheckEntity($this, $data);
    }


    private $_cross_dimer = null;

    // Canonical facade: $client->CrossDimer()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->cross_dimer()
    // resolves here too.
    public function CrossDimer($data = null)
    {
        require_once __DIR__ . '/entity/cross_dimer_entity.php';
        if ($data === null) {
            if ($this->_cross_dimer === null) {
                $this->_cross_dimer = new CrossDimerEntity($this, null);
            }
            return $this->_cross_dimer;
        }
        return new CrossDimerEntity($this, $data);
    }


    private $_dna_molarity = null;

    // Canonical facade: $client->DnaMolarity()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->dna_molarity()
    // resolves here too.
    public function DnaMolarity($data = null)
    {
        require_once __DIR__ . '/entity/dna_molarity_entity.php';
        if ($data === null) {
            if ($this->_dna_molarity === null) {
                $this->_dna_molarity = new DnaMolarityEntity($this, null);
            }
            return $this->_dna_molarity;
        }
        return new DnaMolarityEntity($this, $data);
    }


    private $_double_digest = null;

    // Canonical facade: $client->DoubleDigest()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->double_digest()
    // resolves here too.
    public function DoubleDigest($data = null)
    {
        require_once __DIR__ . '/entity/double_digest_entity.php';
        if ($data === null) {
            if ($this->_double_digest === null) {
                $this->_double_digest = new DoubleDigestEntity($this, null);
            }
            return $this->_double_digest;
        }
        return new DoubleDigestEntity($this, $data);
    }


    private $_export_echo_picklist = null;

    // Canonical facade: $client->ExportEchoPicklist()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->export_echo_picklist()
    // resolves here too.
    public function ExportEchoPicklist($data = null)
    {
        require_once __DIR__ . '/entity/export_echo_picklist_entity.php';
        if ($data === null) {
            if ($this->_export_echo_picklist === null) {
                $this->_export_echo_picklist = new ExportEchoPicklistEntity($this, null);
            }
            return $this->_export_echo_picklist;
        }
        return new ExportEchoPicklistEntity($this, $data);
    }


    private $_export_opentrons_protocol = null;

    // Canonical facade: $client->ExportOpentronsProtocol()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->export_opentrons_protocol()
    // resolves here too.
    public function ExportOpentronsProtocol($data = null)
    {
        require_once __DIR__ . '/entity/export_opentrons_protocol_entity.php';
        if ($data === null) {
            if ($this->_export_opentrons_protocol === null) {
                $this->_export_opentrons_protocol = new ExportOpentronsProtocolEntity($this, null);
            }
            return $this->_export_opentrons_protocol;
        }
        return new ExportOpentronsProtocolEntity($this, $data);
    }


    private $_export_plate_layout = null;

    // Canonical facade: $client->ExportPlateLayout()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->export_plate_layout()
    // resolves here too.
    public function ExportPlateLayout($data = null)
    {
        require_once __DIR__ . '/entity/export_plate_layout_entity.php';
        if ($data === null) {
            if ($this->_export_plate_layout === null) {
                $this->_export_plate_layout = new ExportPlateLayoutEntity($this, null);
            }
            return $this->_export_plate_layout;
        }
        return new ExportPlateLayoutEntity($this, $data);
    }


    private $_expression_heatmap_cluster = null;

    // Canonical facade: $client->ExpressionHeatmapCluster()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->expression_heatmap_cluster()
    // resolves here too.
    public function ExpressionHeatmapCluster($data = null)
    {
        require_once __DIR__ . '/entity/expression_heatmap_cluster_entity.php';
        if ($data === null) {
            if ($this->_expression_heatmap_cluster === null) {
                $this->_expression_heatmap_cluster = new ExpressionHeatmapClusterEntity($this, null);
            }
            return $this->_expression_heatmap_cluster;
        }
        return new ExpressionHeatmapClusterEntity($this, $data);
    }


    private $_fastq_qc_report = null;

    // Canonical facade: $client->FastqQcReport()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->fastq_qc_report()
    // resolves here too.
    public function FastqQcReport($data = null)
    {
        require_once __DIR__ . '/entity/fastq_qc_report_entity.php';
        if ($data === null) {
            if ($this->_fastq_qc_report === null) {
                $this->_fastq_qc_report = new FastqQcReportEntity($this, null);
            }
            return $this->_fastq_qc_report;
        }
        return new FastqQcReportEntity($this, $data);
    }


    private $_fastq_trim = null;

    // Canonical facade: $client->FastqTrim()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->fastq_trim()
    // resolves here too.
    public function FastqTrim($data = null)
    {
        require_once __DIR__ . '/entity/fastq_trim_entity.php';
        if ($data === null) {
            if ($this->_fastq_trim === null) {
                $this->_fastq_trim = new FastqTrimEntity($this, null);
            }
            return $this->_fastq_trim;
        }
        return new FastqTrimEntity($this, $data);
    }


    private $_find_orf = null;

    // Canonical facade: $client->FindOrf()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->find_orf()
    // resolves here too.
    public function FindOrf($data = null)
    {
        require_once __DIR__ . '/entity/find_orf_entity.php';
        if ($data === null) {
            if ($this->_find_orf === null) {
                $this->_find_orf = new FindOrfEntity($this, null);
            }
            return $this->_find_orf;
        }
        return new FindOrfEntity($this, $data);
    }


    private $_format_sequence = null;

    // Canonical facade: $client->FormatSequence()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->format_sequence()
    // resolves here too.
    public function FormatSequence($data = null)
    {
        require_once __DIR__ . '/entity/format_sequence_entity.php';
        if ($data === null) {
            if ($this->_format_sequence === null) {
                $this->_format_sequence = new FormatSequenceEntity($this, null);
            }
            return $this->_format_sequence;
        }
        return new FormatSequenceEntity($this, $data);
    }


    private $_functional_enrichment = null;

    // Canonical facade: $client->FunctionalEnrichment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->functional_enrichment()
    // resolves here too.
    public function FunctionalEnrichment($data = null)
    {
        require_once __DIR__ . '/entity/functional_enrichment_entity.php';
        if ($data === null) {
            if ($this->_functional_enrichment === null) {
                $this->_functional_enrichment = new FunctionalEnrichmentEntity($this, null);
            }
            return $this->_functional_enrichment;
        }
        return new FunctionalEnrichmentEntity($this, $data);
    }


    private $_gc_content = null;

    // Canonical facade: $client->GcContent()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->gc_content()
    // resolves here too.
    public function GcContent($data = null)
    {
        require_once __DIR__ . '/entity/gc_content_entity.php';
        if ($data === null) {
            if ($this->_gc_content === null) {
                $this->_gc_content = new GcContentEntity($this, null);
            }
            return $this->_gc_content;
        }
        return new GcContentEntity($this, $data);
    }


    private $_gene_dossier = null;

    // Canonical facade: $client->GeneDossier()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->gene_dossier()
    // resolves here too.
    public function GeneDossier($data = null)
    {
        require_once __DIR__ . '/entity/gene_dossier_entity.php';
        if ($data === null) {
            if ($this->_gene_dossier === null) {
                $this->_gene_dossier = new GeneDossierEntity($this, null);
            }
            return $this->_gene_dossier;
        }
        return new GeneDossierEntity($this, $data);
    }


    private $_gene_expression = null;

    // Canonical facade: $client->GeneExpression()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->gene_expression()
    // resolves here too.
    public function GeneExpression($data = null)
    {
        require_once __DIR__ . '/entity/gene_expression_entity.php';
        if ($data === null) {
            if ($this->_gene_expression === null) {
                $this->_gene_expression = new GeneExpressionEntity($this, null);
            }
            return $this->_gene_expression;
        }
        return new GeneExpressionEntity($this, $data);
    }


    private $_gene_model = null;

    // Canonical facade: $client->GeneModel()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->gene_model()
    // resolves here too.
    public function GeneModel($data = null)
    {
        require_once __DIR__ . '/entity/gene_model_entity.php';
        if ($data === null) {
            if ($this->_gene_model === null) {
                $this->_gene_model = new GeneModelEntity($this, null);
            }
            return $this->_gene_model;
        }
        return new GeneModelEntity($this, $data);
    }


    private $_golden_gate_fidelity = null;

    // Canonical facade: $client->GoldenGateFidelity()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->golden_gate_fidelity()
    // resolves here too.
    public function GoldenGateFidelity($data = null)
    {
        require_once __DIR__ . '/entity/golden_gate_fidelity_entity.php';
        if ($data === null) {
            if ($this->_golden_gate_fidelity === null) {
                $this->_golden_gate_fidelity = new GoldenGateFidelityEntity($this, null);
            }
            return $this->_golden_gate_fidelity;
        }
        return new GoldenGateFidelityEntity($this, $data);
    }


    private $_hgvs_convert = null;

    // Canonical facade: $client->HgvsConvert()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->hgvs_convert()
    // resolves here too.
    public function HgvsConvert($data = null)
    {
        require_once __DIR__ . '/entity/hgvs_convert_entity.php';
        if ($data === null) {
            if ($this->_hgvs_convert === null) {
                $this->_hgvs_convert = new HgvsConvertEntity($this, null);
            }
            return $this->_hgvs_convert;
        }
        return new HgvsConvertEntity($this, $data);
    }


    private $_id_map_poll = null;

    // Canonical facade: $client->IdMapPoll()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->id_map_poll()
    // resolves here too.
    public function IdMapPoll($data = null)
    {
        require_once __DIR__ . '/entity/id_map_poll_entity.php';
        if ($data === null) {
            if ($this->_id_map_poll === null) {
                $this->_id_map_poll = new IdMapPollEntity($this, null);
            }
            return $this->_id_map_poll;
        }
        return new IdMapPollEntity($this, $data);
    }


    private $_id_map_submit = null;

    // Canonical facade: $client->IdMapSubmit()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->id_map_submit()
    // resolves here too.
    public function IdMapSubmit($data = null)
    {
        require_once __DIR__ . '/entity/id_map_submit_entity.php';
        if ($data === null) {
            if ($this->_id_map_submit === null) {
                $this->_id_map_submit = new IdMapSubmitEntity($this, null);
            }
            return $this->_id_map_submit;
        }
        return new IdMapSubmitEntity($this, $data);
    }


    private $_in_silico_pcr = null;

    // Canonical facade: $client->InSilicoPcr()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->in_silico_pcr()
    // resolves here too.
    public function InSilicoPcr($data = null)
    {
        require_once __DIR__ . '/entity/in_silico_pcr_entity.php';
        if ($data === null) {
            if ($this->_in_silico_pcr === null) {
                $this->_in_silico_pcr = new InSilicoPcrEntity($this, null);
            }
            return $this->_in_silico_pcr;
        }
        return new InSilicoPcrEntity($this, $data);
    }


    private $_kasp_primer_design = null;

    // Canonical facade: $client->KaspPrimerDesign()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->kasp_primer_design()
    // resolves here too.
    public function KaspPrimerDesign($data = null)
    {
        require_once __DIR__ . '/entity/kasp_primer_design_entity.php';
        if ($data === null) {
            if ($this->_kasp_primer_design === null) {
                $this->_kasp_primer_design = new KaspPrimerDesignEntity($this, null);
            }
            return $this->_kasp_primer_design;
        }
        return new KaspPrimerDesignEntity($this, $data);
    }


    private $_list_tool = null;

    // Canonical facade: $client->ListTool()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->list_tool()
    // resolves here too.
    public function ListTool($data = null)
    {
        require_once __DIR__ . '/entity/list_tool_entity.php';
        if ($data === null) {
            if ($this->_list_tool === null) {
                $this->_list_tool = new ListToolEntity($this, null);
            }
            return $this->_list_tool;
        }
        return new ListToolEntity($this, $data);
    }


    private $_melting_temperature = null;

    // Canonical facade: $client->MeltingTemperature()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->melting_temperature()
    // resolves here too.
    public function MeltingTemperature($data = null)
    {
        require_once __DIR__ . '/entity/melting_temperature_entity.php';
        if ($data === null) {
            if ($this->_melting_temperature === null) {
                $this->_melting_temperature = new MeltingTemperatureEntity($this, null);
            }
            return $this->_melting_temperature;
        }
        return new MeltingTemperatureEntity($this, $data);
    }


    private $_motif_finder = null;

    // Canonical facade: $client->MotifFinder()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->motif_finder()
    // resolves here too.
    public function MotifFinder($data = null)
    {
        require_once __DIR__ . '/entity/motif_finder_entity.php';
        if ($data === null) {
            if ($this->_motif_finder === null) {
                $this->_motif_finder = new MotifFinderEntity($this, null);
            }
            return $this->_motif_finder;
        }
        return new MotifFinderEntity($this, $data);
    }


    private $_multiple_sequence_alignment = null;

    // Canonical facade: $client->MultipleSequenceAlignment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->multiple_sequence_alignment()
    // resolves here too.
    public function MultipleSequenceAlignment($data = null)
    {
        require_once __DIR__ . '/entity/multiple_sequence_alignment_entity.php';
        if ($data === null) {
            if ($this->_multiple_sequence_alignment === null) {
                $this->_multiple_sequence_alignment = new MultipleSequenceAlignmentEntity($this, null);
            }
            return $this->_multiple_sequence_alignment;
        }
        return new MultipleSequenceAlignmentEntity($this, $data);
    }


    private $_oligo_analysi = null;

    // Canonical facade: $client->OligoAnalysi()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->oligo_analysi()
    // resolves here too.
    public function OligoAnalysi($data = null)
    {
        require_once __DIR__ . '/entity/oligo_analysi_entity.php';
        if ($data === null) {
            if ($this->_oligo_analysi === null) {
                $this->_oligo_analysi = new OligoAnalysiEntity($this, null);
            }
            return $this->_oligo_analysi;
        }
        return new OligoAnalysiEntity($this, $data);
    }


    private $_ortholog_map = null;

    // Canonical facade: $client->OrthologMap()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->ortholog_map()
    // resolves here too.
    public function OrthologMap($data = null)
    {
        require_once __DIR__ . '/entity/ortholog_map_entity.php';
        if ($data === null) {
            if ($this->_ortholog_map === null) {
                $this->_ortholog_map = new OrthologMapEntity($this, null);
            }
            return $this->_ortholog_map;
        }
        return new OrthologMapEntity($this, $data);
    }


    private $_pairwise_alignment = null;

    // Canonical facade: $client->PairwiseAlignment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->pairwise_alignment()
    // resolves here too.
    public function PairwiseAlignment($data = null)
    {
        require_once __DIR__ . '/entity/pairwise_alignment_entity.php';
        if ($data === null) {
            if ($this->_pairwise_alignment === null) {
                $this->_pairwise_alignment = new PairwiseAlignmentEntity($this, null);
            }
            return $this->_pairwise_alignment;
        }
        return new PairwiseAlignmentEntity($this, $data);
    }


    private $_parse_genbank = null;

    // Canonical facade: $client->ParseGenbank()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->parse_genbank()
    // resolves here too.
    public function ParseGenbank($data = null)
    {
        require_once __DIR__ . '/entity/parse_genbank_entity.php';
        if ($data === null) {
            if ($this->_parse_genbank === null) {
                $this->_parse_genbank = new ParseGenbankEntity($this, null);
            }
            return $this->_parse_genbank;
        }
        return new ParseGenbankEntity($this, $data);
    }


    private $_parse_sanger_trace = null;

    // Canonical facade: $client->ParseSangerTrace()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->parse_sanger_trace()
    // resolves here too.
    public function ParseSangerTrace($data = null)
    {
        require_once __DIR__ . '/entity/parse_sanger_trace_entity.php';
        if ($data === null) {
            if ($this->_parse_sanger_trace === null) {
                $this->_parse_sanger_trace = new ParseSangerTraceEntity($this, null);
            }
            return $this->_parse_sanger_trace;
        }
        return new ParseSangerTraceEntity($this, $data);
    }


    private $_plasmid_annotate = null;

    // Canonical facade: $client->PlasmidAnnotate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->plasmid_annotate()
    // resolves here too.
    public function PlasmidAnnotate($data = null)
    {
        require_once __DIR__ . '/entity/plasmid_annotate_entity.php';
        if ($data === null) {
            if ($this->_plasmid_annotate === null) {
                $this->_plasmid_annotate = new PlasmidAnnotateEntity($this, null);
            }
            return $this->_plasmid_annotate;
        }
        return new PlasmidAnnotateEntity($this, $data);
    }


    private $_plasmid_deep_annotate = null;

    // Canonical facade: $client->PlasmidDeepAnnotate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->plasmid_deep_annotate()
    // resolves here too.
    public function PlasmidDeepAnnotate($data = null)
    {
        require_once __DIR__ . '/entity/plasmid_deep_annotate_entity.php';
        if ($data === null) {
            if ($this->_plasmid_deep_annotate === null) {
                $this->_plasmid_deep_annotate = new PlasmidDeepAnnotateEntity($this, null);
            }
            return $this->_plasmid_deep_annotate;
        }
        return new PlasmidDeepAnnotateEntity($this, $data);
    }


    private $_plasmid_full_report = null;

    // Canonical facade: $client->PlasmidFullReport()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->plasmid_full_report()
    // resolves here too.
    public function PlasmidFullReport($data = null)
    {
        require_once __DIR__ . '/entity/plasmid_full_report_entity.php';
        if ($data === null) {
            if ($this->_plasmid_full_report === null) {
                $this->_plasmid_full_report = new PlasmidFullReportEntity($this, null);
            }
            return $this->_plasmid_full_report;
        }
        return new PlasmidFullReportEntity($this, $data);
    }


    private $_plasmid_identify = null;

    // Canonical facade: $client->PlasmidIdentify()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->plasmid_identify()
    // resolves here too.
    public function PlasmidIdentify($data = null)
    {
        require_once __DIR__ . '/entity/plasmid_identify_entity.php';
        if ($data === null) {
            if ($this->_plasmid_identify === null) {
                $this->_plasmid_identify = new PlasmidIdentifyEntity($this, null);
            }
            return $this->_plasmid_identify;
        }
        return new PlasmidIdentifyEntity($this, $data);
    }


    private $_prime_editing_design = null;

    // Canonical facade: $client->PrimeEditingDesign()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->prime_editing_design()
    // resolves here too.
    public function PrimeEditingDesign($data = null)
    {
        require_once __DIR__ . '/entity/prime_editing_design_entity.php';
        if ($data === null) {
            if ($this->_prime_editing_design === null) {
                $this->_prime_editing_design = new PrimeEditingDesignEntity($this, null);
            }
            return $this->_prime_editing_design;
        }
        return new PrimeEditingDesignEntity($this, $data);
    }


    private $_prime_editing_twin_design = null;

    // Canonical facade: $client->PrimeEditingTwinDesign()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->prime_editing_twin_design()
    // resolves here too.
    public function PrimeEditingTwinDesign($data = null)
    {
        require_once __DIR__ . '/entity/prime_editing_twin_design_entity.php';
        if ($data === null) {
            if ($this->_prime_editing_twin_design === null) {
                $this->_prime_editing_twin_design = new PrimeEditingTwinDesignEntity($this, null);
            }
            return $this->_prime_editing_twin_design;
        }
        return new PrimeEditingTwinDesignEntity($this, $data);
    }


    private $_primer_design = null;

    // Canonical facade: $client->PrimerDesign()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->primer_design()
    // resolves here too.
    public function PrimerDesign($data = null)
    {
        require_once __DIR__ . '/entity/primer_design_entity.php';
        if ($data === null) {
            if ($this->_primer_design === null) {
                $this->_primer_design = new PrimerDesignEntity($this, null);
            }
            return $this->_primer_design;
        }
        return new PrimerDesignEntity($this, $data);
    }


    private $_primer_specificity = null;

    // Canonical facade: $client->PrimerSpecificity()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->primer_specificity()
    // resolves here too.
    public function PrimerSpecificity($data = null)
    {
        require_once __DIR__ . '/entity/primer_specificity_entity.php';
        if ($data === null) {
            if ($this->_primer_specificity === null) {
                $this->_primer_specificity = new PrimerSpecificityEntity($this, null);
            }
            return $this->_primer_specificity;
        }
        return new PrimerSpecificityEntity($this, $data);
    }


    private $_protease_digestion = null;

    // Canonical facade: $client->ProteaseDigestion()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->protease_digestion()
    // resolves here too.
    public function ProteaseDigestion($data = null)
    {
        require_once __DIR__ . '/entity/protease_digestion_entity.php';
        if ($data === null) {
            if ($this->_protease_digestion === null) {
                $this->_protease_digestion = new ProteaseDigestionEntity($this, null);
            }
            return $this->_protease_digestion;
        }
        return new ProteaseDigestionEntity($this, $data);
    }


    private $_protein_annotate_poll = null;

    // Canonical facade: $client->ProteinAnnotatePoll()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->protein_annotate_poll()
    // resolves here too.
    public function ProteinAnnotatePoll($data = null)
    {
        require_once __DIR__ . '/entity/protein_annotate_poll_entity.php';
        if ($data === null) {
            if ($this->_protein_annotate_poll === null) {
                $this->_protein_annotate_poll = new ProteinAnnotatePollEntity($this, null);
            }
            return $this->_protein_annotate_poll;
        }
        return new ProteinAnnotatePollEntity($this, $data);
    }


    private $_protein_annotate_submit = null;

    // Canonical facade: $client->ProteinAnnotateSubmit()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->protein_annotate_submit()
    // resolves here too.
    public function ProteinAnnotateSubmit($data = null)
    {
        require_once __DIR__ . '/entity/protein_annotate_submit_entity.php';
        if ($data === null) {
            if ($this->_protein_annotate_submit === null) {
                $this->_protein_annotate_submit = new ProteinAnnotateSubmitEntity($this, null);
            }
            return $this->_protein_annotate_submit;
        }
        return new ProteinAnnotateSubmitEntity($this, $data);
    }


    private $_protein_hydrophobicity = null;

    // Canonical facade: $client->ProteinHydrophobicity()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->protein_hydrophobicity()
    // resolves here too.
    public function ProteinHydrophobicity($data = null)
    {
        require_once __DIR__ . '/entity/protein_hydrophobicity_entity.php';
        if ($data === null) {
            if ($this->_protein_hydrophobicity === null) {
                $this->_protein_hydrophobicity = new ProteinHydrophobicityEntity($this, null);
            }
            return $this->_protein_hydrophobicity;
        }
        return new ProteinHydrophobicityEntity($this, $data);
    }


    private $_protein_property = null;

    // Canonical facade: $client->ProteinProperty()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->protein_property()
    // resolves here too.
    public function ProteinProperty($data = null)
    {
        require_once __DIR__ . '/entity/protein_property_entity.php';
        if ($data === null) {
            if ($this->_protein_property === null) {
                $this->_protein_property = new ProteinPropertyEntity($this, null);
            }
            return $this->_protein_property;
        }
        return new ProteinPropertyEntity($this, $data);
    }


    private $_random_sequence = null;

    // Canonical facade: $client->RandomSequence()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->random_sequence()
    // resolves here too.
    public function RandomSequence($data = null)
    {
        require_once __DIR__ . '/entity/random_sequence_entity.php';
        if ($data === null) {
            if ($this->_random_sequence === null) {
                $this->_random_sequence = new RandomSequenceEntity($this, null);
            }
            return $this->_random_sequence;
        }
        return new RandomSequenceEntity($this, $data);
    }


    private $_restriction_site = null;

    // Canonical facade: $client->RestrictionSite()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->restriction_site()
    // resolves here too.
    public function RestrictionSite($data = null)
    {
        require_once __DIR__ . '/entity/restriction_site_entity.php';
        if ($data === null) {
            if ($this->_restriction_site === null) {
                $this->_restriction_site = new RestrictionSiteEntity($this, null);
            }
            return $this->_restriction_site;
        }
        return new RestrictionSiteEntity($this, $data);
    }


    private $_reverse_complement = null;

    // Canonical facade: $client->ReverseComplement()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->reverse_complement()
    // resolves here too.
    public function ReverseComplement($data = null)
    {
        require_once __DIR__ . '/entity/reverse_complement_entity.php';
        if ($data === null) {
            if ($this->_reverse_complement === null) {
                $this->_reverse_complement = new ReverseComplementEntity($this, null);
            }
            return $this->_reverse_complement;
        }
        return new ReverseComplementEntity($this, $data);
    }


    private $_reverse_translate = null;

    // Canonical facade: $client->ReverseTranslate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->reverse_translate()
    // resolves here too.
    public function ReverseTranslate($data = null)
    {
        require_once __DIR__ . '/entity/reverse_translate_entity.php';
        if ($data === null) {
            if ($this->_reverse_translate === null) {
                $this->_reverse_translate = new ReverseTranslateEntity($this, null);
            }
            return $this->_reverse_translate;
        }
        return new ReverseTranslateEntity($this, $data);
    }


    private $_rna_fold = null;

    // Canonical facade: $client->RnaFold()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->rna_fold()
    // resolves here too.
    public function RnaFold($data = null)
    {
        require_once __DIR__ . '/entity/rna_fold_entity.php';
        if ($data === null) {
            if ($this->_rna_fold === null) {
                $this->_rna_fold = new RnaFoldEntity($this, null);
            }
            return $this->_rna_fold;
        }
        return new RnaFoldEntity($this, $data);
    }


    private $_sanger_vs_reference = null;

    // Canonical facade: $client->SangerVsReference()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->sanger_vs_reference()
    // resolves here too.
    public function SangerVsReference($data = null)
    {
        require_once __DIR__ . '/entity/sanger_vs_reference_entity.php';
        if ($data === null) {
            if ($this->_sanger_vs_reference === null) {
                $this->_sanger_vs_reference = new SangerVsReferenceEntity($this, null);
            }
            return $this->_sanger_vs_reference;
        }
        return new SangerVsReferenceEntity($this, $data);
    }


    private $_save_permalink = null;

    // Canonical facade: $client->SavePermalink()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->save_permalink()
    // resolves here too.
    public function SavePermalink($data = null)
    {
        require_once __DIR__ . '/entity/save_permalink_entity.php';
        if ($data === null) {
            if ($this->_save_permalink === null) {
                $this->_save_permalink = new SavePermalinkEntity($this, null);
            }
            return $this->_save_permalink;
        }
        return new SavePermalinkEntity($this, $data);
    }


    private $_seqfile_stat = null;

    // Canonical facade: $client->SeqfileStat()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->seqfile_stat()
    // resolves here too.
    public function SeqfileStat($data = null)
    {
        require_once __DIR__ . '/entity/seqfile_stat_entity.php';
        if ($data === null) {
            if ($this->_seqfile_stat === null) {
                $this->_seqfile_stat = new SeqfileStatEntity($this, null);
            }
            return $this->_seqfile_stat;
        }
        return new SeqfileStatEntity($this, $data);
    }


    private $_sequence_fetch = null;

    // Canonical facade: $client->SequenceFetch()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->sequence_fetch()
    // resolves here too.
    public function SequenceFetch($data = null)
    {
        require_once __DIR__ . '/entity/sequence_fetch_entity.php';
        if ($data === null) {
            if ($this->_sequence_fetch === null) {
                $this->_sequence_fetch = new SequenceFetchEntity($this, null);
            }
            return $this->_sequence_fetch;
        }
        return new SequenceFetchEntity($this, $data);
    }


    private $_sequence_format_convert = null;

    // Canonical facade: $client->SequenceFormatConvert()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->sequence_format_convert()
    // resolves here too.
    public function SequenceFormatConvert($data = null)
    {
        require_once __DIR__ . '/entity/sequence_format_convert_entity.php';
        if ($data === null) {
            if ($this->_sequence_format_convert === null) {
                $this->_sequence_format_convert = new SequenceFormatConvertEntity($this, null);
            }
            return $this->_sequence_format_convert;
        }
        return new SequenceFormatConvertEntity($this, $data);
    }


    private $_sequence_report = null;

    // Canonical facade: $client->SequenceReport()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->sequence_report()
    // resolves here too.
    public function SequenceReport($data = null)
    {
        require_once __DIR__ . '/entity/sequence_report_entity.php';
        if ($data === null) {
            if ($this->_sequence_report === null) {
                $this->_sequence_report = new SequenceReportEntity($this, null);
            }
            return $this->_sequence_report;
        }
        return new SequenceReportEntity($this, $data);
    }


    private $_sequence_search = null;

    // Canonical facade: $client->SequenceSearch()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->sequence_search()
    // resolves here too.
    public function SequenceSearch($data = null)
    {
        require_once __DIR__ . '/entity/sequence_search_entity.php';
        if ($data === null) {
            if ($this->_sequence_search === null) {
                $this->_sequence_search = new SequenceSearchEntity($this, null);
            }
            return $this->_sequence_search;
        }
        return new SequenceSearchEntity($this, $data);
    }


    private $_sequencing_readback_verify = null;

    // Canonical facade: $client->SequencingReadbackVerify()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->sequencing_readback_verify()
    // resolves here too.
    public function SequencingReadbackVerify($data = null)
    {
        require_once __DIR__ . '/entity/sequencing_readback_verify_entity.php';
        if ($data === null) {
            if ($this->_sequencing_readback_verify === null) {
                $this->_sequencing_readback_verify = new SequencingReadbackVerifyEntity($this, null);
            }
            return $this->_sequencing_readback_verify;
        }
        return new SequencingReadbackVerifyEntity($this, $data);
    }


    private $_session_create = null;

    // Canonical facade: $client->SessionCreate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->session_create()
    // resolves here too.
    public function SessionCreate($data = null)
    {
        require_once __DIR__ . '/entity/session_create_entity.php';
        if ($data === null) {
            if ($this->_session_create === null) {
                $this->_session_create = new SessionCreateEntity($this, null);
            }
            return $this->_session_create;
        }
        return new SessionCreateEntity($this, $data);
    }


    private $_session_get = null;

    // Canonical facade: $client->SessionGet()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->session_get()
    // resolves here too.
    public function SessionGet($data = null)
    {
        require_once __DIR__ . '/entity/session_get_entity.php';
        if ($data === null) {
            if ($this->_session_get === null) {
                $this->_session_get = new SessionGetEntity($this, null);
            }
            return $this->_session_get;
        }
        return new SessionGetEntity($this, $data);
    }


    private $_session_run = null;

    // Canonical facade: $client->SessionRun()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->session_run()
    // resolves here too.
    public function SessionRun($data = null)
    {
        require_once __DIR__ . '/entity/session_run_entity.php';
        if ($data === null) {
            if ($this->_session_run === null) {
                $this->_session_run = new SessionRunEntity($this, null);
            }
            return $this->_session_run;
        }
        return new SessionRunEntity($this, $data);
    }


    private $_session_set = null;

    // Canonical facade: $client->SessionSet()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->session_set()
    // resolves here too.
    public function SessionSet($data = null)
    {
        require_once __DIR__ . '/entity/session_set_entity.php';
        if ($data === null) {
            if ($this->_session_set === null) {
                $this->_session_set = new SessionSetEntity($this, null);
            }
            return $this->_session_set;
        }
        return new SessionSetEntity($this, $data);
    }


    private $_sirna_design = null;

    // Canonical facade: $client->SirnaDesign()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->sirna_design()
    // resolves here too.
    public function SirnaDesign($data = null)
    {
        require_once __DIR__ . '/entity/sirna_design_entity.php';
        if ($data === null) {
            if ($this->_sirna_design === null) {
                $this->_sirna_design = new SirnaDesignEntity($this, null);
            }
            return $this->_sirna_design;
        }
        return new SirnaDesignEntity($this, $data);
    }


    private $_site_directed_mutagenesi = null;

    // Canonical facade: $client->SiteDirectedMutagenesi()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->site_directed_mutagenesi()
    // resolves here too.
    public function SiteDirectedMutagenesi($data = null)
    {
        require_once __DIR__ . '/entity/site_directed_mutagenesi_entity.php';
        if ($data === null) {
            if ($this->_site_directed_mutagenesi === null) {
                $this->_site_directed_mutagenesi = new SiteDirectedMutagenesiEntity($this, null);
            }
            return $this->_site_directed_mutagenesi;
        }
        return new SiteDirectedMutagenesiEntity($this, $data);
    }


    private $_translate = null;

    // Canonical facade: $client->Translate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->translate()
    // resolves here too.
    public function Translate($data = null)
    {
        require_once __DIR__ . '/entity/translate_entity.php';
        if ($data === null) {
            if ($this->_translate === null) {
                $this->_translate = new TranslateEntity($this, null);
            }
            return $this->_translate;
        }
        return new TranslateEntity($this, $data);
    }


    private $_variant_annotate = null;

    // Canonical facade: $client->VariantAnnotate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->variant_annotate()
    // resolves here too.
    public function VariantAnnotate($data = null)
    {
        require_once __DIR__ . '/entity/variant_annotate_entity.php';
        if ($data === null) {
            if ($this->_variant_annotate === null) {
                $this->_variant_annotate = new VariantAnnotateEntity($this, null);
            }
            return $this->_variant_annotate;
        }
        return new VariantAnnotateEntity($this, $data);
    }


    private $_variant_comparator = null;

    // Canonical facade: $client->VariantComparator()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->variant_comparator()
    // resolves here too.
    public function VariantComparator($data = null)
    {
        require_once __DIR__ . '/entity/variant_comparator_entity.php';
        if ($data === null) {
            if ($this->_variant_comparator === null) {
                $this->_variant_comparator = new VariantComparatorEntity($this, null);
            }
            return $this->_variant_comparator;
        }
        return new VariantComparatorEntity($this, $data);
    }


    private $_verify_assembly = null;

    // Canonical facade: $client->VerifyAssembly()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->verify_assembly()
    // resolves here too.
    public function VerifyAssembly($data = null)
    {
        require_once __DIR__ . '/entity/verify_assembly_entity.php';
        if ($data === null) {
            if ($this->_verify_assembly === null) {
                $this->_verify_assembly = new VerifyAssemblyEntity($this, null);
            }
            return $this->_verify_assembly;
        }
        return new VerifyAssemblyEntity($this, $data);
    }


    private $_verify_construct = null;

    // Canonical facade: $client->VerifyConstruct()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->verify_construct()
    // resolves here too.
    public function VerifyConstruct($data = null)
    {
        require_once __DIR__ . '/entity/verify_construct_entity.php';
        if ($data === null) {
            if ($this->_verify_construct === null) {
                $this->_verify_construct = new VerifyConstructEntity($this, null);
            }
            return $this->_verify_construct;
        }
        return new VerifyConstructEntity($this, $data);
    }


    private $_virtual_gel = null;

    // Canonical facade: $client->VirtualGel()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->virtual_gel()
    // resolves here too.
    public function VirtualGel($data = null)
    {
        require_once __DIR__ . '/entity/virtual_gel_entity.php';
        if ($data === null) {
            if ($this->_virtual_gel === null) {
                $this->_virtual_gel = new VirtualGelEntity($this, null);
            }
            return $this->_virtual_gel;
        }
        return new VirtualGelEntity($this, $data);
    }


    private $_volcano_plot_data = null;

    // Canonical facade: $client->VolcanoPlotData()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->volcano_plot_data()
    // resolves here too.
    public function VolcanoPlotData($data = null)
    {
        require_once __DIR__ . '/entity/volcano_plot_data_entity.php';
        if ($data === null) {
            if ($this->_volcano_plot_data === null) {
                $this->_volcano_plot_data = new VolcanoPlotDataEntity($this, null);
            }
            return $this->_volcano_plot_data;
        }
        return new VolcanoPlotDataEntity($this, $data);
    }


    private $_web_search = null;

    // Canonical facade: $client->WebSearch()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->web_search()
    // resolves here too.
    public function WebSearch($data = null)
    {
        require_once __DIR__ . '/entity/web_search_entity.php';
        if ($data === null) {
            if ($this->_web_search === null) {
                $this->_web_search = new WebSearchEntity($this, null);
            }
            return $this->_web_search;
        }
        return new WebSearchEntity($this, $data);
    }



    public static function test(?array $testopts = null, ?array $sdkopts = null): self
    {
        $sdkopts = $sdkopts ?? [];
        $sdkopts = Struct::clone($sdkopts);
        $sdkopts = is_array($sdkopts) ? $sdkopts : [];

        $testopts = $testopts ?? [];
        $testopts = Struct::clone($testopts);
        $testopts = is_array($testopts) ? $testopts : [];
        $testopts["active"] = true;

        if (!isset($sdkopts["feature"])) {
            $sdkopts["feature"] = [];
        }
        $sdkopts["feature"]["test"] = $testopts;

        $sdk = new SeqbenchMcpSDK($sdkopts);
        $sdk->mode = "test";
        return $sdk;
    }
}
