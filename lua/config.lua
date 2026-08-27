-- SeqbenchMcp SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "SeqbenchMcp",
      slug = "seqbench-mcp",
      version = "0.0.1",
      target = "lua",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
        ["transport"] = "base",
      },
    },
    options = {
      base = "https://seqbench.com/api/v1",
      auth = {
        prefix = "Bearer",
      },
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["alphafold_lookup"] = {},
        ["aso_design"] = {},
        ["base_editing_design"] = {},
        ["batch"] = {},
        ["batch__workflow"] = {},
        ["characterize_sequence"] = {},
        ["cloning_simulate"] = {},
        ["codon_adaptation_index"] = {},
        ["codon_optimize"] = {},
        ["construct_autofix"] = {},
        ["construct_qc"] = {},
        ["crispr_grna_design"] = {},
        ["crispr_hdr_donor"] = {},
        ["crispr_offtarget_check"] = {},
        ["cross_dimer"] = {},
        ["dna_molarity"] = {},
        ["double_digest"] = {},
        ["export_echo_picklist"] = {},
        ["export_opentrons_protocol"] = {},
        ["export_plate_layout"] = {},
        ["expression_heatmap_cluster"] = {},
        ["fastq_qc_report"] = {},
        ["fastq_trim"] = {},
        ["find_orf"] = {},
        ["format_sequence"] = {},
        ["functional_enrichment"] = {},
        ["gc_content"] = {},
        ["gene_dossier"] = {},
        ["gene_expression"] = {},
        ["gene_model"] = {},
        ["golden_gate_fidelity"] = {},
        ["hgvs_convert"] = {},
        ["id_map_poll"] = {},
        ["id_map_submit"] = {},
        ["in_silico_pcr"] = {},
        ["kasp_primer_design"] = {},
        ["list_tool"] = {},
        ["melting_temperature"] = {},
        ["motif_finder"] = {},
        ["multiple_sequence_alignment"] = {},
        ["oligo_analysi"] = {},
        ["ortholog_map"] = {},
        ["pairwise_alignment"] = {},
        ["parse_genbank"] = {},
        ["parse_sanger_trace"] = {},
        ["plasmid_annotate"] = {},
        ["plasmid_deep_annotate"] = {},
        ["plasmid_full_report"] = {},
        ["plasmid_identify"] = {},
        ["prime_editing_design"] = {},
        ["prime_editing_twin_design"] = {},
        ["primer_design"] = {},
        ["primer_specificity"] = {},
        ["protease_digestion"] = {},
        ["protein_annotate_poll"] = {},
        ["protein_annotate_submit"] = {},
        ["protein_hydrophobicity"] = {},
        ["protein_property"] = {},
        ["random_sequence"] = {},
        ["restriction_site"] = {},
        ["reverse_complement"] = {},
        ["reverse_translate"] = {},
        ["rna_fold"] = {},
        ["sanger_vs_reference"] = {},
        ["save_permalink"] = {},
        ["seqfile_stat"] = {},
        ["sequence_fetch"] = {},
        ["sequence_format_convert"] = {},
        ["sequence_report"] = {},
        ["sequence_search"] = {},
        ["sequencing_readback_verify"] = {},
        ["session_create"] = {},
        ["session_get"] = {},
        ["session_run"] = {},
        ["session_set"] = {},
        ["sirna_design"] = {},
        ["site_directed_mutagenesi"] = {},
        ["translate"] = {},
        ["variant_annotate"] = {},
        ["variant_comparator"] = {},
        ["verify_assembly"] = {},
        ["verify_construct"] = {},
        ["virtual_gel"] = {},
        ["volcano_plot_data"] = {},
        ["web_search"] = {},
      },
    },
    entity = {
      ["alphafold_lookup"] = {
        ["fields"] = {
          {
            ["name"] = "accession",
            ["req"] = true,
            ["short"] = "UniProt accession, e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "alphafold_lookup",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/alphafold_lookup",
                ["parts"] = {
                  "alphafold_lookup",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["accession"] = "`reqdata.accession`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["aso_design"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "length",
            ["short"] = "Total gapmer length (nt).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "target",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "wing",
            ["short"] = "Modified-wing length on each side (nt); the central gap = length − 2×wing.",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "aso_design",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/aso_design",
                ["parts"] = {
                  "aso_design",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["length"] = "`reqdata.length`",
                    ["target"] = "`reqdata.target`",
                    ["wing"] = "`reqdata.wing`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["base_editing_design"] = {
        ["fields"] = {
          {
            ["name"] = "editor",
            ["short"] = "Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "frameStart",
            ["short"] = "Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "target",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "targetPosition",
            ["short"] = "Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "base_editing_design",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/base_editing_design",
                ["parts"] = {
                  "base_editing_design",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["editor"] = "`reqdata.editor`",
                    ["frameStart"] = "`reqdata.frame_start`",
                    ["target"] = "`reqdata.target`",
                    ["targetPosition"] = "`reqdata.target_position`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["batch"] = {
        ["fields"] = {
          {
            ["name"] = "args",
            ["short"] = "Shared tool arguments applied to every record.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "capped",
            ["req"] = true,
            ["short"] = "True if input exceeded the record limit.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "columns",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "count",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "errors",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "input",
            ["req"] = true,
            ["short"] = "Multi-FASTA text or one sequence per line (max ~2,000,000 chars).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "limit",
            ["req"] = true,
            ["short"] = "Maximum records per call (500).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "rows",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "A batchable tool slug (see `GET /batch`).",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "batch",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/batch",
                ["parts"] = {
                  "batch",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.result`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/batch",
                ["parts"] = {
                  "batch",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["batch__workflow"] = {
        ["fields"] = {
          {
            ["name"] = "capped",
            ["req"] = true,
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "columns",
            ["req"] = true,
            ["short"] = "Flattened \"<step>·<tool>·<key>\" column headers.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "count",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "errors",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "input",
            ["req"] = true,
            ["short"] = "Multi-FASTA text or one sequence per line.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "limit",
            ["req"] = true,
            ["short"] = "Maximum records per call (200).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "rows",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "steps",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
        },
        ["name"] = "batch__workflow",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/workflow",
                ["parts"] = {
                  "workflow",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.result`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/workflow",
                ["parts"] = {
                  "workflow",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["characterize_sequence"] = {
        ["fields"] = {
          {
            ["name"] = "endPrimerLength",
            ["short"] = "Length of the naive end primers taken from each end.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "maxOrfs",
            ["short"] = "Maximum number of ORFs to return, longest first.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "minOrfAa",
            ["short"] = "Minimum ORF length in amino acids (nucleotide input only).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "characterize_sequence",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/characterize_sequence",
                ["parts"] = {
                  "characterize_sequence",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["endPrimerLength"] = "`reqdata.end_primer_length`",
                    ["maxOrfs"] = "`reqdata.max_orf`",
                    ["minOrfAa"] = "`reqdata.min_orf_aa`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["cloning_simulate"] = {
        ["fields"] = {
          {
            ["name"] = "armTmTarget",
            ["short"] = "Target annealing Tm (°C) for primer arms.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "circular",
            ["short"] = "Produce a circular product.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "enzyme",
            ["short"] = "Type IIS enzyme for Golden Gate (e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "enzyme3",
            ["short"] = "3′ enzyme (restriction method).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "enzyme5",
            ["short"] = "5′ enzyme (restriction method).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fragments",
            ["short"] = "Fragments (5′→3′), assembled head-to-tail.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "insert",
            ["short"] = "Insert sequence (restriction method).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "method",
            ["req"] = true,
            ["short"] = "Assembly method.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "names",
            ["short"] = "Optional labels for each fragment.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "overlapLen",
            ["short"] = "Gibson homology-arm length (bp).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "vector",
            ["short"] = "Vector sequence (restriction method).",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "cloning_simulate",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/cloning_simulate",
                ["parts"] = {
                  "cloning_simulate",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["armTmTarget"] = "`reqdata.arm_tm_target`",
                    ["circular"] = "`reqdata.circular`",
                    ["enzyme"] = "`reqdata.enzyme`",
                    ["enzyme3"] = "`reqdata.enzyme3`",
                    ["enzyme5"] = "`reqdata.enzyme5`",
                    ["fragments"] = "`reqdata.fragment`",
                    ["insert"] = "`reqdata.insert`",
                    ["method"] = "`reqdata.method`",
                    ["names"] = "`reqdata.name`",
                    ["overlapLen"] = "`reqdata.overlap_len`",
                    ["vector"] = "`reqdata.vector`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["codon_adaptation_index"] = {
        ["fields"] = {
          {
            ["name"] = "frameStart",
            ["short"] = "1-based position to start reading codons.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "organism",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "rareThreshold",
            ["short"] = "Relative adaptiveness (w) below this flags a codon as rare.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Coding sequence (DNA/RNA; should start in-frame at ATG).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "codon_adaptation_index",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/codon_adaptation_index",
                ["parts"] = {
                  "codon_adaptation_index",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["frameStart"] = "`reqdata.frame_start`",
                    ["organism"] = "`reqdata.organism`",
                    ["rareThreshold"] = "`reqdata.rare_threshold`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["codon_optimize"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "organism",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "protein",
            ["req"] = true,
            ["short"] = "Protein sequence (one-letter codes).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "codon_optimize",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/codon_optimize",
                ["parts"] = {
                  "codon_optimize",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["organism"] = "`reqdata.organism`",
                    ["protein"] = "`reqdata.protein`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["construct_autofix"] = {
        ["fields"] = {
          {
            ["name"] = "avoidEnzymes",
            ["short"] = "Enzyme names whose internal sites should be removed (e.g.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "crypticOrfMinAa",
            ["short"] = "Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "frameStart",
            ["short"] = "1-based nucleotide where the reading frame begins.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "gcHigh",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "gcLow",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "gcWindow",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "homopolymerMin",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "maxPasses",
            ["short"] = "Repeat full passes until clean or no further progress.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "organism",
            ["short"] = "Codon-usage table to prefer among synonymous options.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "construct_autofix",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/construct_autofix",
                ["parts"] = {
                  "construct_autofix",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["avoidEnzymes"] = "`reqdata.avoid_enzyme`",
                    ["crypticOrfMinAa"] = "`reqdata.cryptic_orf_min_aa`",
                    ["frameStart"] = "`reqdata.frame_start`",
                    ["gcHigh"] = "`reqdata.gc_high`",
                    ["gcLow"] = "`reqdata.gc_low`",
                    ["gcWindow"] = "`reqdata.gc_window`",
                    ["homopolymerMin"] = "`reqdata.homopolymer_min`",
                    ["maxPasses"] = "`reqdata.max_pass`",
                    ["organism"] = "`reqdata.organism`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["construct_qc"] = {
        ["fields"] = {
          {
            ["name"] = "avoidEnzymes",
            ["short"] = "Enzyme names whose internal sites should be flagged as errors.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "crypticOrfMinAa",
            ["short"] = "Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "frameStart",
            ["short"] = "1-based nucleotide where the reading frame begins.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "gcHigh",
            ["short"] = "GC% above this flags a GC-rich window.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "gcLow",
            ["short"] = "GC% below this flags an AT-rich window.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "gcWindow",
            ["short"] = "Sliding-window size (nt) for GC-extreme scanning.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "homopolymerMin",
            ["short"] = "Minimum run length to flag a homopolymer.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "construct_qc",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/construct_qc",
                ["parts"] = {
                  "construct_qc",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["avoidEnzymes"] = "`reqdata.avoid_enzyme`",
                    ["crypticOrfMinAa"] = "`reqdata.cryptic_orf_min_aa`",
                    ["frameStart"] = "`reqdata.frame_start`",
                    ["gcHigh"] = "`reqdata.gc_high`",
                    ["gcLow"] = "`reqdata.gc_low`",
                    ["gcWindow"] = "`reqdata.gc_window`",
                    ["homopolymerMin"] = "`reqdata.homopolymer_min`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["crispr_grna_design"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "minScore",
            ["short"] = "Only return guides with a heuristic score at least this high (0–100).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "nuclease",
            ["short"] = "Nuclease id.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "searchReverseStrand",
            ["short"] = "Also scan the reverse strand for guides.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "crispr_grna_design",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/crispr_grna_design",
                ["parts"] = {
                  "crispr_grna_design",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["minScore"] = "`reqdata.min_score`",
                    ["nuclease"] = "`reqdata.nuclease`",
                    ["searchReverseStrand"] = "`reqdata.search_reverse_strand`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["crispr_hdr_donor"] = {
        ["fields"] = {
          {
            ["name"] = "armLength",
            ["short"] = "Homology arm length (bp) on each side.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "blockPam",
            ["short"] = "When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "designGenotypingPrimers",
            ["short"] = "Also design a primer pair (on the original targetSequence) whose product spans the edit site.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "editEnd",
            ["short"] = "1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "editStart",
            ["short"] = "1-based start of the region being replaced.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "frameStart",
            ["short"] = "Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "guideEnd",
            ["short"] = "1-based forward-strand end of the guide's protospacer.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "guideStart",
            ["short"] = "1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "guideStrand",
            ["short"] = "Strand the guide's protospacer is on.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nuclease",
            ["short"] = "Needed only when deriving the cut site from guideStart/guideEnd/guideStrand.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "replacement",
            ["req"] = true,
            ["short"] = "Sequence to insert/substitute (\"\" for a pure deletion).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "targetSequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "crispr_hdr_donor",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/crispr_hdr_donor",
                ["parts"] = {
                  "crispr_hdr_donor",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["armLength"] = "`reqdata.arm_length`",
                    ["blockPam"] = "`reqdata.block_pam`",
                    ["designGenotypingPrimers"] = "`reqdata.design_genotyping_primer`",
                    ["editEnd"] = "`reqdata.edit_end`",
                    ["editStart"] = "`reqdata.edit_start`",
                    ["frameStart"] = "`reqdata.frame_start`",
                    ["guideEnd"] = "`reqdata.guide_end`",
                    ["guideStart"] = "`reqdata.guide_start`",
                    ["guideStrand"] = "`reqdata.guide_strand`",
                    ["nuclease"] = "`reqdata.nuclease`",
                    ["replacement"] = "`reqdata.replacement`",
                    ["targetSequence"] = "`reqdata.target_sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["crispr_offtarget_check"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "maxMismatches",
            ["short"] = "Mismatches tolerated between the protospacer and a candidate genomic site.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "nuclease",
            ["short"] = "Nuclease id — determines the PAM pattern/side required at each candidate site.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "protospacer",
            ["req"] = true,
            ["short"] = "The guide's protospacer sequence, 5'→3' (no PAM).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "crispr_offtarget_check",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/crispr_offtarget_check",
                ["parts"] = {
                  "crispr_offtarget_check",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["maxMismatches"] = "`reqdata.max_mismatch`",
                    ["nuclease"] = "`reqdata.nuclease`",
                    ["protospacer"] = "`reqdata.protospacer`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["cross_dimer"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequenceA",
            ["req"] = true,
            ["short"] = "First oligo (5'→3').",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sequenceB",
            ["req"] = true,
            ["short"] = "Second oligo (5'→3').",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "cross_dimer",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/cross_dimer",
                ["parts"] = {
                  "cross_dimer",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["sequenceA"] = "`reqdata.sequence_a`",
                    ["sequenceB"] = "`reqdata.sequence_b`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["dna_molarity"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "length",
            ["short"] = "Length in bp (dsDNA) or nt (ssDNA/ssRNA).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "massNg",
            ["short"] = "Mass in nanograms.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["short"] = "Optional sequence — overrides length and gives an exact molar mass from base composition.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "type",
            ["short"] = "Molecule type.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "volumeUl",
            ["short"] = "Volume in microlitres (0 = unknown; needed for concentration).",
            ["type"] = "`$NUMBER`",
          },
        },
        ["name"] = "dna_molarity",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/dna_molarity",
                ["parts"] = {
                  "dna_molarity",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["length"] = "`reqdata.length`",
                    ["massNg"] = "`reqdata.mass_ng`",
                    ["sequence"] = "`reqdata.sequence`",
                    ["type"] = "`reqdata.type`",
                    ["volumeUl"] = "`reqdata.volume_ul`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["double_digest"] = {
        ["fields"] = {
          {
            ["name"] = "enzymeA",
            ["req"] = true,
            ["short"] = "First enzyme name (e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "enzymeB",
            ["req"] = true,
            ["short"] = "Second enzyme name (e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "double_digest",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/double_digest",
                ["parts"] = {
                  "double_digest",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["enzymeA"] = "`reqdata.enzyme_a`",
                    ["enzymeB"] = "`reqdata.enzyme_b`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["export_echo_picklist"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "reactions",
            ["req"] = true,
            ["short"] = "One entry per PCR reaction, up to 96 (a single 96-well plate).",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "export_echo_picklist",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/export_echo_picklist",
                ["parts"] = {
                  "export_echo_picklist",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["reactions"] = "`reqdata.reaction`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["export_opentrons_protocol"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "protocolName",
            ["short"] = "Optional protocol name (used in the script's metadata).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "reactions",
            ["req"] = true,
            ["short"] = "One entry per PCR reaction, up to 96 (a single 96-well plate).",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "export_opentrons_protocol",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/export_opentrons_protocol",
                ["parts"] = {
                  "export_opentrons_protocol",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["protocolName"] = "`reqdata.protocol_name`",
                    ["reactions"] = "`reqdata.reaction`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["export_plate_layout"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "reactions",
            ["req"] = true,
            ["short"] = "One entry per PCR reaction, up to 96 (a single 96-well plate).",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "export_plate_layout",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/export_plate_layout",
                ["parts"] = {
                  "export_plate_layout",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["reactions"] = "`reqdata.reaction`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["expression_heatmap_cluster"] = {
        ["fields"] = {
          {
            ["name"] = "clusterCols",
            ["short"] = "Cluster (reorder) samples.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "clusterRows",
            ["short"] = "Cluster (reorder) genes.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "distanceMetric",
            ["short"] = "correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "genes",
            ["req"] = true,
            ["short"] = "Row (gene) labels.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "linkage",
            ["short"] = "average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "samples",
            ["req"] = true,
            ["short"] = "Column (sample) labels.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "values",
            ["req"] = true,
            ["short"] = "genes x samples numeric matrix — one row per gene, in the same order as `genes`.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "zScoreRows",
            ["short"] = "Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization).",
            ["type"] = "`$BOOLEAN`",
          },
        },
        ["name"] = "expression_heatmap_cluster",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/expression_heatmap_cluster",
                ["parts"] = {
                  "expression_heatmap_cluster",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["clusterCols"] = "`reqdata.cluster_col`",
                    ["clusterRows"] = "`reqdata.cluster_row`",
                    ["distanceMetric"] = "`reqdata.distance_metric`",
                    ["genes"] = "`reqdata.gene`",
                    ["linkage"] = "`reqdata.linkage`",
                    ["samples"] = "`reqdata.sample`",
                    ["values"] = "`reqdata.value`",
                    ["zScoreRows"] = "`reqdata.z_score_row`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["fastq_qc_report"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "input",
            ["req"] = true,
            ["short"] = "FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "qualityOffset",
            ["short"] = "FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "fastq_qc_report",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/fastq_qc_report",
                ["parts"] = {
                  "fastq_qc_report",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["input"] = "`reqdata.input`",
                    ["qualityOffset"] = "`reqdata.quality_offset`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["fastq_trim"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "input",
            ["req"] = true,
            ["short"] = "FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "minLength",
            ["short"] = "Reads shorter than this after trimming are dropped.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "qualityOffset",
            ["short"] = "FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "qualityThreshold",
            ["short"] = "3' quality-trim threshold (Phred score).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "fastq_trim",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/fastq_trim",
                ["parts"] = {
                  "fastq_trim",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["input"] = "`reqdata.input`",
                    ["minLength"] = "`reqdata.min_length`",
                    ["qualityOffset"] = "`reqdata.quality_offset`",
                    ["qualityThreshold"] = "`reqdata.quality_threshold`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["find_orf"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "minAaLength",
            ["short"] = "Minimum protein length (aa) to report.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "requireStop",
            ["short"] = "Only report ORFs terminated by a stop codon.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "find_orf",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/find_orfs",
                ["parts"] = {
                  "find_orfs",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["minAaLength"] = "`reqdata.min_aa_length`",
                    ["requireStop"] = "`reqdata.require_stop`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["format_sequence"] = {
        ["fields"] = {
          {
            ["name"] = "caseMode",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "convert",
            ["short"] = "DNA→RNA (T→U) or RNA→DNA (U→T).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "reverse",
            ["short"] = "Reverse the sequence (no complement).",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "stripNonLetters",
            ["short"] = "Remove digits, spaces and gaps (keep letters only).",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "width",
            ["short"] = "Line-wrap width; 0 = single line.",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "format_sequence",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/format_sequence",
                ["parts"] = {
                  "format_sequence",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["caseMode"] = "`reqdata.case_mode`",
                    ["convert"] = "`reqdata.convert`",
                    ["reverse"] = "`reqdata.reverse`",
                    ["sequence"] = "`reqdata.sequence`",
                    ["stripNonLetters"] = "`reqdata.strip_non_letter`",
                    ["width"] = "`reqdata.width`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["functional_enrichment"] = {
        ["fields"] = {
          {
            ["name"] = "background",
            ["short"] = "Custom background/universe gene symbols.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "collections",
            ["short"] = "Which term collections to test.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "genes",
            ["req"] = true,
            ["short"] = "Query gene symbols (human, e.g.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "maxTermSize",
            ["short"] = "Skip terms/pathways with more than this many background genes (matches clusterProfiler's default).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "minTermSize",
            ["short"] = "Skip terms/pathways with fewer than this many background genes.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "functional_enrichment",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/functional_enrichment",
                ["parts"] = {
                  "functional_enrichment",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["background"] = "`reqdata.background`",
                    ["collections"] = "`reqdata.collection`",
                    ["genes"] = "`reqdata.gene`",
                    ["maxTermSize"] = "`reqdata.max_term_size`",
                    ["minTermSize"] = "`reqdata.min_term_size`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["gc_content"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "gc_content",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/gc_content",
                ["parts"] = {
                  "gc_content",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["gene_dossier"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "gene",
            ["req"] = true,
            ["short"] = "A human gene symbol (\"TP53\") or Ensembl gene ID (\"ENSG00000141510\").",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "gene_dossier",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/gene_dossier",
                ["parts"] = {
                  "gene_dossier",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["gene"] = "`reqdata.gene`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["gene_expression"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "gene",
            ["req"] = true,
            ["short"] = "A human gene symbol (\"TP53\") or Ensembl gene ID (\"ENSG00000141510\").",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "gene_expression",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/gene_expression",
                ["parts"] = {
                  "gene_expression",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["gene"] = "`reqdata.gene`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["gene_model"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "gene",
            ["req"] = true,
            ["short"] = "A human gene symbol (\"TP53\") or Ensembl gene ID (\"ENSG00000141510\").",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "gene_model",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/gene_model",
                ["parts"] = {
                  "gene_model",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["gene"] = "`reqdata.gene`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["golden_gate_fidelity"] = {
        ["fields"] = {
          {
            ["name"] = "compareToNamedSet",
            ["short"] = "Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "dataset",
            ["short"] = "Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "overhangs",
            ["req"] = true,
            ["short"] = "The candidate 4-base overhangs for one assembly (e.g.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "riskThreshold",
            ["short"] = "Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "golden_gate_fidelity",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/golden_gate_fidelity",
                ["parts"] = {
                  "golden_gate_fidelity",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["compareToNamedSet"] = "`reqdata.compare_to_named_set`",
                    ["dataset"] = "`reqdata.dataset`",
                    ["overhangs"] = "`reqdata.overhang`",
                    ["riskThreshold"] = "`reqdata.risk_threshold`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["hgvs_convert"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "variant",
            ["req"] = true,
            ["short"] = "A full HGVS \"c.\" variant description: \"<accession or gene symbol>:c.<edit>\", e.g.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "hgvs_convert",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/hgvs_convert",
                ["parts"] = {
                  "hgvs_convert",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["variant"] = "`reqdata.variant`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["id_map_poll"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "jobId",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "id_map_poll",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/id_map_poll",
                ["parts"] = {
                  "id_map_poll",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["jobId"] = "`reqdata.job_id`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["id_map_submit"] = {
        ["fields"] = {
          {
            ["name"] = "from",
            ["req"] = true,
            ["short"] = "Source id type: \"Gene_Name\", \"Ensembl\", \"GeneID\", \"RefSeq_Protein\", or \"UniProtKB_AC-ID\".",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ids",
            ["req"] = true,
            ["short"] = "The ids to map, up to 1000 (e.g.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "taxId",
            ["short"] = "NCBI taxonomy id to disambiguate a gene symbol (only used when from=\"Gene_Name\").",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "to",
            ["req"] = true,
            ["short"] = "Target id type.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "id_map_submit",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/id_map_submit",
                ["parts"] = {
                  "id_map_submit",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["from"] = "`reqdata.from`",
                    ["ids"] = "`reqdata.ids`",
                    ["taxId"] = "`reqdata.tax_id`",
                    ["to"] = "`reqdata.to`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["in_silico_pcr"] = {
        ["fields"] = {
          {
            ["name"] = "circular",
            ["short"] = "Treat the template as circular (plasmid).",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "forwardPrimer",
            ["req"] = true,
            ["short"] = "Primer 1, 5'→3'.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "maxMismatches",
            ["short"] = "Mismatches tolerated per primer.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "reversePrimer",
            ["req"] = true,
            ["short"] = "Primer 2, 5'→3' (order does not matter).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "template",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "in_silico_pcr",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/in_silico_pcr",
                ["parts"] = {
                  "in_silico_pcr",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["circular"] = "`reqdata.circular`",
                    ["forwardPrimer"] = "`reqdata.forward_primer`",
                    ["maxMismatches"] = "`reqdata.max_mismatch`",
                    ["reversePrimer"] = "`reqdata.reverse_primer`",
                    ["template"] = "`reqdata.template`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["kasp_primer_design"] = {
        ["fields"] = {
          {
            ["name"] = "addSecondaryMismatch",
            ["short"] = "Engineer the internal ARMS destabilising mismatch near the 3' end.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "alleleA",
            ["req"] = true,
            ["short"] = "First allele (single base) — gets the FAM tail.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "alleleB",
            ["req"] = true,
            ["short"] = "Second allele (single base) — gets the HEX tail.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "maxAmplicon",
            ["short"] = "Maximum amplicon length for the common reverse primer.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "minAmplicon",
            ["short"] = "Minimum amplicon length for the common reverse primer.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "snpPosition",
            ["req"] = true,
            ["short"] = "1-based position of the SNP on the forward strand.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "target",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "targetCoreTm",
            ["short"] = "Target Tm (°C) for the allele-specific primer core (before the universal tail).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "kasp_primer_design",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/kasp_primer_design",
                ["parts"] = {
                  "kasp_primer_design",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["addSecondaryMismatch"] = "`reqdata.add_secondary_mismatch`",
                    ["alleleA"] = "`reqdata.allele_a`",
                    ["alleleB"] = "`reqdata.allele_b`",
                    ["maxAmplicon"] = "`reqdata.max_amplicon`",
                    ["minAmplicon"] = "`reqdata.min_amplicon`",
                    ["snpPosition"] = "`reqdata.snp_position`",
                    ["target"] = "`reqdata.target`",
                    ["targetCoreTm"] = "`reqdata.target_core_tm`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["list_tool"] = {
        ["fields"] = {},
        ["name"] = "list_tool",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/",
                ["parts"] = {},
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["melting_temperature"] = {
        ["fields"] = {
          {
            ["name"] = "dntpMM",
            ["short"] = "Total [dNTP] (mM), chelates Mg2+.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "mgMM",
            ["short"] = "Divalent cation [Mg2+] (mM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "naMM",
            ["short"] = "Monovalent cation [Na+]/[K+] (mM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "oligoNM",
            ["short"] = "Total strand concentration (nM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "targetTm",
            ["short"] = "Optional target Tm (°C).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "tmTolerance",
            ["short"] = "Allowed +/- window (°C) around targetTm for the gate.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "melting_temperature",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/melting_temperature",
                ["parts"] = {
                  "melting_temperature",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["dntpMM"] = "`reqdata.dntp_mm`",
                    ["mgMM"] = "`reqdata.mg_mm`",
                    ["naMM"] = "`reqdata.na_mm`",
                    ["oligoNM"] = "`reqdata.oligo_nm`",
                    ["sequence"] = "`reqdata.sequence`",
                    ["targetTm"] = "`reqdata.target_tm`",
                    ["tmTolerance"] = "`reqdata.tm_tolerance`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["motif_finder"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "maxMismatches",
            ["short"] = "Maximum allowed mismatches per match.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "motif",
            ["req"] = true,
            ["short"] = "Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "searchReverseStrand",
            ["short"] = "Also search the reverse strand.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "motif_finder",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/motif_finder",
                ["parts"] = {
                  "motif_finder",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["maxMismatches"] = "`reqdata.max_mismatch`",
                    ["motif"] = "`reqdata.motif`",
                    ["searchReverseStrand"] = "`reqdata.search_reverse_strand`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["multiple_sequence_alignment"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "input",
            ["req"] = true,
            ["short"] = "Two or more sequences in multi-FASTA format (>name / sequence).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "multiple_sequence_alignment",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/multiple_sequence_alignment",
                ["parts"] = {
                  "multiple_sequence_alignment",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["input"] = "`reqdata.input`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["oligo_analysi"] = {
        ["fields"] = {
          {
            ["name"] = "dntpMM",
            ["short"] = "Total [dNTP] (mM), chelates Mg2+.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "mgMM",
            ["short"] = "Divalent cation [Mg2+] (mM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "naMM",
            ["short"] = "Monovalent cation [Na+]/[K+] (mM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "oligoNM",
            ["short"] = "Total strand concentration (nM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "oligo_analysi",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/oligo_analysis",
                ["parts"] = {
                  "oligo_analysis",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["dntpMM"] = "`reqdata.dntp_mm`",
                    ["mgMM"] = "`reqdata.mg_mm`",
                    ["naMM"] = "`reqdata.na_mm`",
                    ["oligoNM"] = "`reqdata.oligo_nm`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["ortholog_map"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sourceSpecies",
            ["short"] = "Ensembl species slug the symbols belong to (e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "symbols",
            ["req"] = true,
            ["short"] = "Gene symbols to look up, up to 50 (e.g.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "targetSpecies",
            ["req"] = true,
            ["short"] = "Ensembl species slug to find homologs in (e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "type",
            ["short"] = "Homology type to return.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "ortholog_map",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/ortholog_map",
                ["parts"] = {
                  "ortholog_map",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["sourceSpecies"] = "`reqdata.source_species`",
                    ["symbols"] = "`reqdata.symbol`",
                    ["targetSpecies"] = "`reqdata.target_species`",
                    ["type"] = "`reqdata.type`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["pairwise_alignment"] = {
        ["fields"] = {
          {
            ["name"] = "gap",
            ["short"] = "Linear gap penalty (per gap position).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "match",
            ["short"] = "Match score.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "mismatch",
            ["short"] = "Mismatch penalty.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "mode",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "seqA",
            ["req"] = true,
            ["short"] = "First sequence (raw or FASTA; nucleotide or protein).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "seqB",
            ["req"] = true,
            ["short"] = "Second sequence (raw or FASTA; nucleotide or protein).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "pairwise_alignment",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/pairwise_alignment",
                ["parts"] = {
                  "pairwise_alignment",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["gap"] = "`reqdata.gap`",
                    ["match"] = "`reqdata.match`",
                    ["mismatch"] = "`reqdata.mismatch`",
                    ["mode"] = "`reqdata.mode`",
                    ["seqA"] = "`reqdata.seq_a`",
                    ["seqB"] = "`reqdata.seq_b`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["parse_genbank"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "text",
            ["req"] = true,
            ["short"] = "A GenBank flat file (LOCUS … FEATURES … ORIGIN … //).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "parse_genbank",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/parse_genbank",
                ["parts"] = {
                  "parse_genbank",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["text"] = "`reqdata.text`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["parse_sanger_trace"] = {
        ["fields"] = {
          {
            ["name"] = "fileBase64",
            ["req"] = true,
            ["short"] = "The binary ABIF (.ab1 / .abi) trace file, base64-encoded.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fileName",
            ["short"] = "Optional original file name (echoed back).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "parse_sanger_trace",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/parse_sanger_trace",
                ["parts"] = {
                  "parse_sanger_trace",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["fileBase64"] = "`reqdata.file_base64`",
                    ["fileName"] = "`reqdata.file_name`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["plasmid_annotate"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "plasmid_annotate",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/plasmid_annotate",
                ["parts"] = {
                  "plasmid_annotate",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["plasmid_deep_annotate"] = {
        ["fields"] = {
          {
            ["name"] = "circular",
            ["short"] = "Treat the sequence as a circular plasmid (vs.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "plasmid_deep_annotate",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/plasmid_deep_annotate",
                ["parts"] = {
                  "plasmid_deep_annotate",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["circular"] = "`reqdata.circular`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["plasmid_full_report"] = {
        ["fields"] = {
          {
            ["name"] = "circular",
            ["short"] = "Treat the query as a circular molecule (most plasmids are).",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "topN",
            ["short"] = "How many top-ranked backbone candidates to report.",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "plasmid_full_report",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/plasmid_full_report",
                ["parts"] = {
                  "plasmid_full_report",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["circular"] = "`reqdata.circular`",
                    ["sequence"] = "`reqdata.sequence`",
                    ["topN"] = "`reqdata.top_n`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["plasmid_identify"] = {
        ["fields"] = {
          {
            ["name"] = "circular",
            ["short"] = "Treat the query as a circular molecule (most plasmids are).",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "topN",
            ["short"] = "How many top-ranked backbone candidates to report.",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "plasmid_identify",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/plasmid_identify",
                ["parts"] = {
                  "plasmid_identify",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["circular"] = "`reqdata.circular`",
                    ["sequence"] = "`reqdata.sequence`",
                    ["topN"] = "`reqdata.top_n`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["prime_editing_design"] = {
        ["fields"] = {
          {
            ["name"] = "editEnd",
            ["req"] = true,
            ["short"] = "1-based inclusive end of the region being changed.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "editStart",
            ["req"] = true,
            ["short"] = "1-based inclusive start of the region being changed.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "frameStart",
            ["short"] = "Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "insertedSeq",
            ["short"] = "Replacement bases (forward strand).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "pbsLength",
            ["short"] = "Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "rttHomology",
            ["short"] = "Homology length (nt) 3' of the edit that the RTT should include (typically 10-16).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "target",
            ["req"] = true,
            ["short"] = "Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "prime_editing_design",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/prime_editing_design",
                ["parts"] = {
                  "prime_editing_design",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["editEnd"] = "`reqdata.edit_end`",
                    ["editStart"] = "`reqdata.edit_start`",
                    ["frameStart"] = "`reqdata.frame_start`",
                    ["insertedSeq"] = "`reqdata.inserted_seq`",
                    ["pbsLength"] = "`reqdata.pbs_length`",
                    ["rttHomology"] = "`reqdata.rtt_homology`",
                    ["target"] = "`reqdata.target`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["prime_editing_twin_design"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "newSequence",
            ["req"] = true,
            ["short"] = "New sequence (forward strand) to install in place of [replaceStart, replaceEnd].",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "overlapLength",
            ["short"] = "Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "pbsLength",
            ["short"] = "Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "replaceEnd",
            ["req"] = true,
            ["short"] = "1-based inclusive end of the region being replaced/deleted.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "replaceStart",
            ["req"] = true,
            ["short"] = "1-based inclusive start of the region being replaced/deleted.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "target",
            ["req"] = true,
            ["short"] = "Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "prime_editing_twin_design",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/prime_editing_twin_design",
                ["parts"] = {
                  "prime_editing_twin_design",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["newSequence"] = "`reqdata.new_sequence`",
                    ["overlapLength"] = "`reqdata.overlap_length`",
                    ["pbsLength"] = "`reqdata.pbs_length`",
                    ["replaceEnd"] = "`reqdata.replace_end`",
                    ["replaceStart"] = "`reqdata.replace_start`",
                    ["target"] = "`reqdata.target`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["primer_design"] = {
        ["fields"] = {
          {
            ["name"] = "ampliconMax",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ampliconMin",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "dntpMM",
            ["short"] = "Total [dNTP] (mM), chelates Mg2+.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "gcMax",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "gcMin",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "lenMax",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "lenMin",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "lenOpt",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "maxReturn",
            ["short"] = "Number of best pairs to return.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "mgMM",
            ["short"] = "Divalent cation [Mg2+] (mM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "naMM",
            ["short"] = "Monovalent cation [Na+]/[K+] (mM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "oligoNM",
            ["short"] = "Total strand concentration (nM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "targetEnd",
            ["short"] = "1-based inclusive end of the target region (optional).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "targetStart",
            ["short"] = "1-based inclusive start of a region the product must span (optional).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "template",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tmMax",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "tmMaxDiff",
            ["short"] = "Max Tm difference within a pair (°C).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "tmMin",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "tmOpt",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "primer_design",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/primer_design",
                ["parts"] = {
                  "primer_design",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["ampliconMax"] = "`reqdata.amplicon_max`",
                    ["ampliconMin"] = "`reqdata.amplicon_min`",
                    ["dntpMM"] = "`reqdata.dntp_mm`",
                    ["gcMax"] = "`reqdata.gc_max`",
                    ["gcMin"] = "`reqdata.gc_min`",
                    ["lenMax"] = "`reqdata.len_max`",
                    ["lenMin"] = "`reqdata.len_min`",
                    ["lenOpt"] = "`reqdata.len_opt`",
                    ["maxReturn"] = "`reqdata.max_return`",
                    ["mgMM"] = "`reqdata.mg_mm`",
                    ["naMM"] = "`reqdata.na_mm`",
                    ["oligoNM"] = "`reqdata.oligo_nm`",
                    ["targetEnd"] = "`reqdata.target_end`",
                    ["targetStart"] = "`reqdata.target_start`",
                    ["template"] = "`reqdata.template`",
                    ["tmMax"] = "`reqdata.tm_max`",
                    ["tmMaxDiff"] = "`reqdata.tm_max_diff`",
                    ["tmMin"] = "`reqdata.tm_min`",
                    ["tmOpt"] = "`reqdata.tm_opt`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["primer_specificity"] = {
        ["fields"] = {
          {
            ["name"] = "forwardPrimer",
            ["req"] = true,
            ["short"] = "Forward primer, 5'→3'.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "maxMismatches",
            ["short"] = "Mismatches tolerated per primer against a reference genome.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "maxProductLength",
            ["short"] = "Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "reversePrimer",
            ["req"] = true,
            ["short"] = "Reverse primer, 5'→3'.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "primer_specificity",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/primer_specificity",
                ["parts"] = {
                  "primer_specificity",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["forwardPrimer"] = "`reqdata.forward_primer`",
                    ["maxMismatches"] = "`reqdata.max_mismatch`",
                    ["maxProductLength"] = "`reqdata.max_product_length`",
                    ["reversePrimer"] = "`reqdata.reverse_primer`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["protease_digestion"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "maxMass",
            ["short"] = "Optional upper bound on neutral monoisotopic mass (Da).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "maxPeptides",
            ["short"] = "Cap on the number of returned peptides.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "minMass",
            ["short"] = "Optional lower bound on neutral monoisotopic mass (Da).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "missedCleavages",
            ["short"] = "Allowed missed internal cleavages (0–2).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "protease",
            ["short"] = "Protease or chemical cleavage agent.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Protein sequence (one-letter amino-acid codes; non-AA characters ignored).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "protease_digestion",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/protease_digestion",
                ["parts"] = {
                  "protease_digestion",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["maxMass"] = "`reqdata.max_mass`",
                    ["maxPeptides"] = "`reqdata.max_peptide`",
                    ["minMass"] = "`reqdata.min_mass`",
                    ["missedCleavages"] = "`reqdata.missed_cleavage`",
                    ["protease"] = "`reqdata.protease`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["protein_annotate_poll"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "jobId",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "protein_annotate_poll",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/protein_annotate_poll",
                ["parts"] = {
                  "protein_annotate_poll",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["jobId"] = "`reqdata.job_id`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["protein_annotate_submit"] = {
        ["fields"] = {
          {
            ["name"] = "appl",
            ["short"] = "Restrict to one member database (e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "goterms",
            ["short"] = "Include GO-term cross-references.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Protein sequence, one-letter code (FASTA header, if any, is stripped).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "protein_annotate_submit",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/protein_annotate_submit",
                ["parts"] = {
                  "protein_annotate_submit",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["appl"] = "`reqdata.appl`",
                    ["goterms"] = "`reqdata.goterm`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["protein_hydrophobicity"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "scale",
            ["short"] = "Amino-acid scale.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Protein sequence (one-letter amino-acid codes; non-AA characters ignored).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "window",
            ["short"] = "Sliding-window size (clamped to an odd number ≥ 1).",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "protein_hydrophobicity",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/protein_hydrophobicity",
                ["parts"] = {
                  "protein_hydrophobicity",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["scale"] = "`reqdata.scale`",
                    ["sequence"] = "`reqdata.sequence`",
                    ["window"] = "`reqdata.window`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["protein_property"] = {
        ["fields"] = {
          {
            ["name"] = "chargeStep",
            ["short"] = "pH step for the net-charge titration curve (0–14).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Protein sequence (one-letter amino-acid codes; non-AA characters ignored).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "protein_property",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/protein_properties",
                ["parts"] = {
                  "protein_properties",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["chargeStep"] = "`reqdata.charge_step`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["random_sequence"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "gcContent",
            ["short"] = "Target GC percentage 0..100 (dna/rna only); omit for uniform.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "kind",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "length",
            ["req"] = true,
            ["short"] = "Number of residues to generate.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "random_sequence",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/random_sequence",
                ["parts"] = {
                  "random_sequence",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["gcContent"] = "`reqdata.gc_content`",
                    ["kind"] = "`reqdata.kind`",
                    ["length"] = "`reqdata.length`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["restriction_site"] = {
        ["fields"] = {
          {
            ["name"] = "enzymes",
            ["short"] = "Enzyme names to scan; omit to scan all curated enzymes.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "restriction_site",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/restriction_sites",
                ["parts"] = {
                  "restriction_sites",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["enzymes"] = "`reqdata.enzyme`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["reverse_complement"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "type",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "reverse_complement",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/reverse_complement",
                ["parts"] = {
                  "reverse_complement",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["sequence"] = "`reqdata.sequence`",
                    ["type"] = "`reqdata.type`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["reverse_translate"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "mode",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "organism",
            ["short"] = "Codon-usage host (ignored in degenerate mode).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "protein",
            ["req"] = true,
            ["short"] = "Protein sequence (one-letter codes; * for stop).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "reverse_translate",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/reverse_translate",
                ["parts"] = {
                  "reverse_translate",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["mode"] = "`reqdata.mode`",
                    ["organism"] = "`reqdata.organism`",
                    ["protein"] = "`reqdata.protein`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["rna_fold"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "rna_fold",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/rna_fold",
                ["parts"] = {
                  "rna_fold",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["sanger_vs_reference"] = {
        ["fields"] = {
          {
            ["name"] = "fileBase64",
            ["short"] = "The binary ABIF (.ab1 / .abi) trace file, base64-encoded.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fileName",
            ["short"] = "Optional original file name (echoed back).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "minCoverage",
            ["short"] = "Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "read",
            ["short"] = "Sanger read as FASTA or raw text (alternative to uploading an ABIF trace).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "reference",
            ["req"] = true,
            ["short"] = "Expected reference sequence (FASTA or raw).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "sanger_vs_reference",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/sanger_vs_reference",
                ["parts"] = {
                  "sanger_vs_reference",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["fileBase64"] = "`reqdata.file_base64`",
                    ["fileName"] = "`reqdata.file_name`",
                    ["minCoverage"] = "`reqdata.min_coverage`",
                    ["read"] = "`reqdata.read`",
                    ["reference"] = "`reqdata.reference`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["save_permalink"] = {
        ["fields"] = {
          {
            ["name"] = "args",
            ["req"] = true,
            ["short"] = "Arguments for that tool, exactly as you would pass to it directly.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "save_permalink",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/save_permalink",
                ["parts"] = {
                  "save_permalink",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["args"] = "`reqdata.arg`",
                    ["tool"] = "`reqdata.tool`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["seqfile_stat"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "input",
            ["req"] = true,
            ["short"] = "FASTA or FASTQ text (raw sequence is treated as single-record FASTA).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "qualityOffset",
            ["short"] = "FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "seqfile_stat",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/seqfile_stats",
                ["parts"] = {
                  "seqfile_stats",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["input"] = "`reqdata.input`",
                    ["qualityOffset"] = "`reqdata.quality_offset`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["sequence_fetch"] = {
        ["fields"] = {
          {
            ["name"] = "accession",
            ["req"] = true,
            ["short"] = "GenBank/RefSeq accession (e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "db",
            ["short"] = "Database to query; auto-detects from the accession format.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "format",
            ["short"] = "Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "sequence_fetch",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/sequence_fetch",
                ["parts"] = {
                  "sequence_fetch",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["accession"] = "`reqdata.accession`",
                    ["db"] = "`reqdata.db`",
                    ["format"] = "`reqdata.format`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["sequence_format_convert"] = {
        ["fields"] = {
          {
            ["name"] = "from",
            ["short"] = "Input format; 'auto' sniffs it from the first meaningful line.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "input",
            ["req"] = true,
            ["short"] = "A FASTA or GenBank record to convert.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "to",
            ["short"] = "Output format.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "sequence_format_convert",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/sequence_format_convert",
                ["parts"] = {
                  "sequence_format_convert",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["from"] = "`reqdata.from`",
                    ["input"] = "`reqdata.input`",
                    ["to"] = "`reqdata.to`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["sequence_report"] = {
        ["fields"] = {
          {
            ["name"] = "endPrimerLength",
            ["short"] = "Length of the naive end primers taken from each end.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "maxOrfs",
            ["short"] = "Maximum number of ORFs to return, longest first.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "minOrfAa",
            ["short"] = "Minimum ORF length in amino acids.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "sequence_report",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/sequence_report",
                ["parts"] = {
                  "sequence_report",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["endPrimerLength"] = "`reqdata.end_primer_length`",
                    ["maxOrfs"] = "`reqdata.max_orf`",
                    ["minOrfAa"] = "`reqdata.min_orf_aa`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["sequence_search"] = {
        ["fields"] = {
          {
            ["name"] = "db",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "gene",
            ["short"] = "Gene symbol/name, e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "maxResults",
            ["short"] = "Up to 20.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "organism",
            ["short"] = "Organism name, e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "term",
            ["short"] = "Raw NCBI search term (advanced) — overrides gene/organism when given, e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "sequence_search",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/sequence_search",
                ["parts"] = {
                  "sequence_search",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["db"] = "`reqdata.db`",
                    ["gene"] = "`reqdata.gene`",
                    ["maxResults"] = "`reqdata.max_result`",
                    ["organism"] = "`reqdata.organism`",
                    ["term"] = "`reqdata.term`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["sequencing_readback_verify"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "minSupportingReads",
            ["short"] = "Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "reads",
            ["req"] = true,
            ["short"] = "Raw reads in FASTA or FASTQ format (auto-detected).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "reference",
            ["req"] = true,
            ["short"] = "The claimed/expected reference sequence.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "sequencing_readback_verify",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/sequencing_readback_verify",
                ["parts"] = {
                  "sequencing_readback_verify",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["minSupportingReads"] = "`reqdata.min_supporting_read`",
                    ["reads"] = "`reqdata.read`",
                    ["reference"] = "`reqdata.reference`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["session_create"] = {
        ["fields"] = {
          {
            ["name"] = "entries",
            ["short"] = "Initial named entries, e.g.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "session_create",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/session_create",
                ["parts"] = {
                  "session_create",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["entries"] = "`reqdata.entry`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["session_get"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "names",
            ["short"] = "Only return these entries; omit to return all of them.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sessionId",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "session_get",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/session_get",
                ["parts"] = {
                  "session_get",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["names"] = "`reqdata.name`",
                    ["sessionId"] = "`reqdata.session_id`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["session_run"] = {
        ["fields"] = {
          {
            ["name"] = "args",
            ["short"] = "Additional literal arguments, merged with the ones resolved from the session.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "fromSession",
            ["short"] = "Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sessionId",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "writeBack",
            ["short"] = "Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names.",
            ["type"] = "`$OBJECT`",
          },
        },
        ["name"] = "session_run",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/session_run",
                ["parts"] = {
                  "session_run",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["args"] = "`reqdata.arg`",
                    ["fromSession"] = "`reqdata.from_session`",
                    ["sessionId"] = "`reqdata.session_id`",
                    ["tool"] = "`reqdata.tool`",
                    ["writeBack"] = "`reqdata.write_back`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["session_set"] = {
        ["fields"] = {
          {
            ["name"] = "entries",
            ["req"] = true,
            ["short"] = "Named entries to add/overwrite, e.g.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sessionId",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "session_set",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/session_set",
                ["parts"] = {
                  "session_set",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["entries"] = "`reqdata.entry`",
                    ["sessionId"] = "`reqdata.session_id`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["sirna_design"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "minReynolds",
            ["short"] = "Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "shRnaLoop",
            ["short"] = "Loop sequence used when assembling the shRNA cassette.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "target",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "sirna_design",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/sirna_design",
                ["parts"] = {
                  "sirna_design",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["minReynolds"] = "`reqdata.min_reynold`",
                    ["shRnaLoop"] = "`reqdata.sh_rna_loop`",
                    ["target"] = "`reqdata.target`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["site_directed_mutagenesi"] = {
        ["fields"] = {
          {
            ["name"] = "armTmTarget",
            ["short"] = "Target Tm (°C) for each template-binding arm.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "dntpMM",
            ["short"] = "Total [dNTP] (mM), chelates Mg2+.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "editKind",
            ["short"] = "Edit at the nucleotide or amino-acid level.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "frameStart",
            ["short"] = "1-based position of the first base of codon 1 (editKind='aa').",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "mgMM",
            ["short"] = "Divalent cation [Mg2+] (mM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "naMM",
            ["short"] = "Monovalent cation [Na+]/[K+] (mM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "newBase",
            ["short"] = "Replacement base (editKind='nt').",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "oligoNM",
            ["short"] = "Total strand concentration (nM).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "organism",
            ["short"] = "Codon-usage table for choosing the new codon (editKind='aa').",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "position",
            ["short"] = "1-based position to substitute (editKind='nt').",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "residue",
            ["short"] = "1-based residue number to change (editKind='aa').",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "style",
            ["short"] = "Mutagenic primer style.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "targetAa",
            ["short"] = "Target amino acid, one-letter code incl '*' (editKind='aa').",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "template",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "site_directed_mutagenesi",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/site_directed_mutagenesis",
                ["parts"] = {
                  "site_directed_mutagenesis",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["armTmTarget"] = "`reqdata.arm_tm_target`",
                    ["dntpMM"] = "`reqdata.dntp_mm`",
                    ["editKind"] = "`reqdata.edit_kind`",
                    ["frameStart"] = "`reqdata.frame_start`",
                    ["mgMM"] = "`reqdata.mg_mm`",
                    ["naMM"] = "`reqdata.na_mm`",
                    ["newBase"] = "`reqdata.new_base`",
                    ["oligoNM"] = "`reqdata.oligo_nm`",
                    ["organism"] = "`reqdata.organism`",
                    ["position"] = "`reqdata.position`",
                    ["residue"] = "`reqdata.residue`",
                    ["style"] = "`reqdata.style`",
                    ["targetAa"] = "`reqdata.target_aa`",
                    ["template"] = "`reqdata.template`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["translate"] = {
        ["fields"] = {
          {
            ["name"] = "frame",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "toStop",
            ["short"] = "Stop at the first stop codon.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "translate",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/translate",
                ["parts"] = {
                  "translate",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["frame"] = "`reqdata.frame`",
                    ["sequence"] = "`reqdata.sequence`",
                    ["toStop"] = "`reqdata.to_stop`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["variant_annotate"] = {
        ["fields"] = {
          {
            ["name"] = "assembly",
            ["short"] = "Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "variant",
            ["req"] = true,
            ["short"] = "An rsID (\"rs1042522\"), chrom:pos:ref:alt (\"17:7676154:G:C\", single-base substitutions only), genomic HGVS (\"chr17:g.7676154G>C\" or \"17:g.7676154G>C\"), or transcript HGVS c.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "variant_annotate",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/variant_annotate",
                ["parts"] = {
                  "variant_annotate",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["assembly"] = "`reqdata.assembly`",
                    ["variant"] = "`reqdata.variant`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["variant_comparator"] = {
        ["fields"] = {
          {
            ["name"] = "coding",
            ["short"] = "Treat as a coding sequence and report amino-acid effects.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "frameStart",
            ["short"] = "1-based reading-frame start (used when coding is true).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "query",
            ["req"] = true,
            ["short"] = "Query / variant sequence (raw or FASTA).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "reference",
            ["req"] = true,
            ["short"] = "Reference / wild-type sequence (raw or FASTA).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "variant_comparator",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/variant_comparator",
                ["parts"] = {
                  "variant_comparator",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["coding"] = "`reqdata.coding`",
                    ["frameStart"] = "`reqdata.frame_start`",
                    ["query"] = "`reqdata.query`",
                    ["reference"] = "`reqdata.reference`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["verify_assembly"] = {
        ["fields"] = {
          {
            ["name"] = "armTmTarget",
            ["short"] = "Target annealing Tm (°C) for primer arms.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "circular",
            ["short"] = "Treat the product/claimed construct as circular (most plasmids are).",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "claimedConstruct",
            ["req"] = true,
            ["short"] = "The sequence you claim you ended up with.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "coding",
            ["short"] = "Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "enzyme",
            ["short"] = "Type IIS enzyme for Golden Gate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "enzyme3",
            ["short"] = "3′ enzyme (restriction method).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "enzyme5",
            ["short"] = "5′ enzyme (restriction method).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fragmentPcrs",
            ["short"] = "Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "fragments",
            ["short"] = "Fragments (5′→3′), assembled head-to-tail (gibson/goldengate).",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "frameStart",
            ["short"] = "1-based reading-frame start on claimedConstruct, used when coding is true.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "insert",
            ["short"] = "Insert sequence (restriction method).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "insertPcr",
            ["short"] = "Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "method",
            ["req"] = true,
            ["short"] = "Assembly method used.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "names",
            ["short"] = "Optional labels for each fragment.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "overlapLen",
            ["short"] = "Gibson homology-arm length (bp).",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "vector",
            ["short"] = "Vector sequence (restriction method).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "vectorPcr",
            ["short"] = "Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}.",
            ["type"] = "`$OBJECT`",
          },
        },
        ["name"] = "verify_assembly",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/verify_assembly",
                ["parts"] = {
                  "verify_assembly",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["armTmTarget"] = "`reqdata.arm_tm_target`",
                    ["circular"] = "`reqdata.circular`",
                    ["claimedConstruct"] = "`reqdata.claimed_construct`",
                    ["coding"] = "`reqdata.coding`",
                    ["enzyme"] = "`reqdata.enzyme`",
                    ["enzyme3"] = "`reqdata.enzyme3`",
                    ["enzyme5"] = "`reqdata.enzyme5`",
                    ["fragmentPcrs"] = "`reqdata.fragment_pcr`",
                    ["fragments"] = "`reqdata.fragment`",
                    ["frameStart"] = "`reqdata.frame_start`",
                    ["insert"] = "`reqdata.insert`",
                    ["insertPcr"] = "`reqdata.insert_pcr`",
                    ["method"] = "`reqdata.method`",
                    ["names"] = "`reqdata.name`",
                    ["overlapLen"] = "`reqdata.overlap_len`",
                    ["vector"] = "`reqdata.vector`",
                    ["vectorPcr"] = "`reqdata.vector_pcr`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["verify_construct"] = {
        ["fields"] = {
          {
            ["name"] = "claimedConstruct",
            ["req"] = true,
            ["short"] = "The final sequence claimed to have been built.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "expectedFrameStart",
            ["short"] = "1-based position in claimedConstruct where the intended reading frame begins.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "insertForwardPrimer",
            ["req"] = true,
            ["short"] = "Forward primer used to amplify the insert, 5'→3'.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "insertReversePrimer",
            ["req"] = true,
            ["short"] = "Reverse primer used to amplify the insert, 5'→3'.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "insertTemplate",
            ["req"] = true,
            ["short"] = "PCR template the insert was amplified from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "maxPrimerMismatches",
            ["short"] = "Mismatches tolerated per primer during PCR prediction.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "templateCircular",
            ["short"] = "Treat insertTemplate as circular (e.g.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "verify_construct",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/verify_construct",
                ["parts"] = {
                  "verify_construct",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["claimedConstruct"] = "`reqdata.claimed_construct`",
                    ["expectedFrameStart"] = "`reqdata.expected_frame_start`",
                    ["insertForwardPrimer"] = "`reqdata.insert_forward_primer`",
                    ["insertReversePrimer"] = "`reqdata.insert_reverse_primer`",
                    ["insertTemplate"] = "`reqdata.insert_template`",
                    ["maxPrimerMismatches"] = "`reqdata.max_primer_mismatch`",
                    ["templateCircular"] = "`reqdata.template_circular`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["virtual_gel"] = {
        ["fields"] = {
          {
            ["name"] = "circular",
            ["short"] = "Treat the sequence as circular (plasmid).",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "enzymes",
            ["short"] = "Enzyme names to digest with.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ladder",
            ["short"] = "DNA ladder to plot alongside the sample lane.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sequence",
            ["req"] = true,
            ["short"] = "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "virtual_gel",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/virtual_gel",
                ["parts"] = {
                  "virtual_gel",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["circular"] = "`reqdata.circular`",
                    ["enzymes"] = "`reqdata.enzyme`",
                    ["ladder"] = "`reqdata.ladder`",
                    ["sequence"] = "`reqdata.sequence`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["volcano_plot_data"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "rows",
            ["req"] = true,
            ["short"] = "Differential expression rows, one per gene.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "volcano_plot_data",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/volcano_plot_data",
                ["parts"] = {
                  "volcano_plot_data",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["rows"] = "`reqdata.row`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["web_search"] = {
        ["fields"] = {
          {
            ["name"] = "gate",
            ["short"] = "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "max_results",
            ["short"] = "Maximum number of results to return (default 5, max 10).",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "ok",
            ["req"] = true,
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "provenance",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["union"] = {
              ["branches"] = 2,
              ["count"] = 1,
              ["depth"] = 2,
            },
          },
          {
            ["name"] = "query",
            ["req"] = true,
            ["short"] = "The search query.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "result",
            ["req"] = true,
            ["short"] = "Tool-specific output object.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "tool",
            ["req"] = true,
            ["short"] = "The tool slug that ran.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "web_search",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/web_search",
                ["parts"] = {
                  "web_search",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["max_results"] = "`reqdata.max_result`",
                    ["query"] = "`reqdata.query`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
