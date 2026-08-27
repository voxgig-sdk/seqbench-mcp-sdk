package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "SeqbenchMcp",
			"slug": "seqbench-mcp",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://seqbench.com/api/v1",
			"auth": map[string]any{
				"prefix": "Bearer",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"alphafold_lookup": map[string]any{},
				"aso_design": map[string]any{},
				"base_editing_design": map[string]any{},
				"batch": map[string]any{},
				"batch__workflow": map[string]any{},
				"characterize_sequence": map[string]any{},
				"cloning_simulate": map[string]any{},
				"codon_adaptation_index": map[string]any{},
				"codon_optimize": map[string]any{},
				"construct_autofix": map[string]any{},
				"construct_qc": map[string]any{},
				"crispr_grna_design": map[string]any{},
				"crispr_hdr_donor": map[string]any{},
				"crispr_offtarget_check": map[string]any{},
				"cross_dimer": map[string]any{},
				"dna_molarity": map[string]any{},
				"double_digest": map[string]any{},
				"export_echo_picklist": map[string]any{},
				"export_opentrons_protocol": map[string]any{},
				"export_plate_layout": map[string]any{},
				"expression_heatmap_cluster": map[string]any{},
				"fastq_qc_report": map[string]any{},
				"fastq_trim": map[string]any{},
				"find_orf": map[string]any{},
				"format_sequence": map[string]any{},
				"functional_enrichment": map[string]any{},
				"gc_content": map[string]any{},
				"gene_dossier": map[string]any{},
				"gene_expression": map[string]any{},
				"gene_model": map[string]any{},
				"golden_gate_fidelity": map[string]any{},
				"hgvs_convert": map[string]any{},
				"id_map_poll": map[string]any{},
				"id_map_submit": map[string]any{},
				"in_silico_pcr": map[string]any{},
				"kasp_primer_design": map[string]any{},
				"list_tool": map[string]any{},
				"melting_temperature": map[string]any{},
				"motif_finder": map[string]any{},
				"multiple_sequence_alignment": map[string]any{},
				"oligo_analysi": map[string]any{},
				"ortholog_map": map[string]any{},
				"pairwise_alignment": map[string]any{},
				"parse_genbank": map[string]any{},
				"parse_sanger_trace": map[string]any{},
				"plasmid_annotate": map[string]any{},
				"plasmid_deep_annotate": map[string]any{},
				"plasmid_full_report": map[string]any{},
				"plasmid_identify": map[string]any{},
				"prime_editing_design": map[string]any{},
				"prime_editing_twin_design": map[string]any{},
				"primer_design": map[string]any{},
				"primer_specificity": map[string]any{},
				"protease_digestion": map[string]any{},
				"protein_annotate_poll": map[string]any{},
				"protein_annotate_submit": map[string]any{},
				"protein_hydrophobicity": map[string]any{},
				"protein_property": map[string]any{},
				"random_sequence": map[string]any{},
				"restriction_site": map[string]any{},
				"reverse_complement": map[string]any{},
				"reverse_translate": map[string]any{},
				"rna_fold": map[string]any{},
				"sanger_vs_reference": map[string]any{},
				"save_permalink": map[string]any{},
				"seqfile_stat": map[string]any{},
				"sequence_fetch": map[string]any{},
				"sequence_format_convert": map[string]any{},
				"sequence_report": map[string]any{},
				"sequence_search": map[string]any{},
				"sequencing_readback_verify": map[string]any{},
				"session_create": map[string]any{},
				"session_get": map[string]any{},
				"session_run": map[string]any{},
				"session_set": map[string]any{},
				"sirna_design": map[string]any{},
				"site_directed_mutagenesi": map[string]any{},
				"translate": map[string]any{},
				"variant_annotate": map[string]any{},
				"variant_comparator": map[string]any{},
				"verify_assembly": map[string]any{},
				"verify_construct": map[string]any{},
				"virtual_gel": map[string]any{},
				"volcano_plot_data": map[string]any{},
				"web_search": map[string]any{},
			},
		},
		"entity": map[string]any{
			"alphafold_lookup": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accession",
						"req": true,
						"short": "UniProt accession, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "alphafold_lookup",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/alphafold_lookup",
								"parts": []any{
									"alphafold_lookup",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"accession": "`reqdata.accession`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"aso_design": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "length",
						"short": "Total gapmer length (nt).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "wing",
						"short": "Modified-wing length on each side (nt); the central gap = length − 2×wing.",
						"type": "`$INTEGER`",
					},
				},
				"name": "aso_design",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/aso_design",
								"parts": []any{
									"aso_design",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"length": "`reqdata.length`",
										"target": "`reqdata.target`",
										"wing": "`reqdata.wing`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"base_editing_design": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "editor",
						"short": "Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "frameStart",
						"short": "Optional 1-based CDS reading-frame start, to classify each edit's amino-acid consequence.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "targetPosition",
						"short": "Optional 1-based forward-strand position of the base you intend to edit; only guides whose window covers it are returned.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "base_editing_design",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/base_editing_design",
								"parts": []any{
									"base_editing_design",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"editor": "`reqdata.editor`",
										"frameStart": "`reqdata.frame_start`",
										"target": "`reqdata.target`",
										"targetPosition": "`reqdata.target_position`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"batch": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "args",
						"short": "Shared tool arguments applied to every record.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "capped",
						"req": true,
						"short": "True if input exceeded the record limit.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "columns",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "count",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "errors",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "input",
						"req": true,
						"short": "Multi-FASTA text or one sequence per line (max ~2,000,000 chars).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "limit",
						"req": true,
						"short": "Maximum records per call (500).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "rows",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "A batchable tool slug (see `GET /batch`).",
						"type": "`$STRING`",
					},
				},
				"name": "batch",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/batch",
								"parts": []any{
									"batch",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/batch",
								"parts": []any{
									"batch",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"batch__workflow": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "capped",
						"req": true,
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "columns",
						"req": true,
						"short": "Flattened \"<step>·<tool>·<key>\" column headers.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "count",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "errors",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "input",
						"req": true,
						"short": "Multi-FASTA text or one sequence per line.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "limit",
						"req": true,
						"short": "Maximum records per call (200).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "rows",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "steps",
						"req": true,
						"type": "`$ARRAY`",
					},
				},
				"name": "batch__workflow",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/workflow",
								"parts": []any{
									"workflow",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.result`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/workflow",
								"parts": []any{
									"workflow",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"characterize_sequence": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "endPrimerLength",
						"short": "Length of the naive end primers taken from each end.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxOrfs",
						"short": "Maximum number of ORFs to return, longest first.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "minOrfAa",
						"short": "Minimum ORF length in amino acids (nucleotide input only).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "characterize_sequence",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/characterize_sequence",
								"parts": []any{
									"characterize_sequence",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"endPrimerLength": "`reqdata.end_primer_length`",
										"maxOrfs": "`reqdata.max_orf`",
										"minOrfAa": "`reqdata.min_orf_aa`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"cloning_simulate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "armTmTarget",
						"short": "Target annealing Tm (°C) for primer arms.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "circular",
						"short": "Produce a circular product.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "enzyme",
						"short": "Type IIS enzyme for Golden Gate (e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enzyme3",
						"short": "3′ enzyme (restriction method).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enzyme5",
						"short": "5′ enzyme (restriction method).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fragments",
						"short": "Fragments (5′→3′), assembled head-to-tail.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "insert",
						"short": "Insert sequence (restriction method).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "method",
						"req": true,
						"short": "Assembly method.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "names",
						"short": "Optional labels for each fragment.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "overlapLen",
						"short": "Gibson homology-arm length (bp).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "vector",
						"short": "Vector sequence (restriction method).",
						"type": "`$STRING`",
					},
				},
				"name": "cloning_simulate",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/cloning_simulate",
								"parts": []any{
									"cloning_simulate",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"armTmTarget": "`reqdata.arm_tm_target`",
										"circular": "`reqdata.circular`",
										"enzyme": "`reqdata.enzyme`",
										"enzyme3": "`reqdata.enzyme3`",
										"enzyme5": "`reqdata.enzyme5`",
										"fragments": "`reqdata.fragment`",
										"insert": "`reqdata.insert`",
										"method": "`reqdata.method`",
										"names": "`reqdata.name`",
										"overlapLen": "`reqdata.overlap_len`",
										"vector": "`reqdata.vector`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"codon_adaptation_index": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "frameStart",
						"short": "1-based position to start reading codons.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "organism",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "rareThreshold",
						"short": "Relative adaptiveness (w) below this flags a codon as rare.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Coding sequence (DNA/RNA; should start in-frame at ATG).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "codon_adaptation_index",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/codon_adaptation_index",
								"parts": []any{
									"codon_adaptation_index",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"frameStart": "`reqdata.frame_start`",
										"organism": "`reqdata.organism`",
										"rareThreshold": "`reqdata.rare_threshold`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"codon_optimize": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "organism",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "protein",
						"req": true,
						"short": "Protein sequence (one-letter codes).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "codon_optimize",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/codon_optimize",
								"parts": []any{
									"codon_optimize",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"organism": "`reqdata.organism`",
										"protein": "`reqdata.protein`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"construct_autofix": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "avoidEnzymes",
						"short": "Enzyme names whose internal sites should be removed (e.g.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "crypticOrfMinAa",
						"short": "Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "frameStart",
						"short": "1-based nucleotide where the reading frame begins.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gcHigh",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gcLow",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gcWindow",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "homopolymerMin",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "maxPasses",
						"short": "Repeat full passes until clean or no further progress.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "organism",
						"short": "Codon-usage table to prefer among synonymous options.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "construct_autofix",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/construct_autofix",
								"parts": []any{
									"construct_autofix",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"avoidEnzymes": "`reqdata.avoid_enzyme`",
										"crypticOrfMinAa": "`reqdata.cryptic_orf_min_aa`",
										"frameStart": "`reqdata.frame_start`",
										"gcHigh": "`reqdata.gc_high`",
										"gcLow": "`reqdata.gc_low`",
										"gcWindow": "`reqdata.gc_window`",
										"homopolymerMin": "`reqdata.homopolymer_min`",
										"maxPasses": "`reqdata.max_pass`",
										"organism": "`reqdata.organism`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"construct_qc": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "avoidEnzymes",
						"short": "Enzyme names whose internal sites should be flagged as errors.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "crypticOrfMinAa",
						"short": "Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "frameStart",
						"short": "1-based nucleotide where the reading frame begins.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gcHigh",
						"short": "GC% above this flags a GC-rich window.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gcLow",
						"short": "GC% below this flags an AT-rich window.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gcWindow",
						"short": "Sliding-window size (nt) for GC-extreme scanning.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "homopolymerMin",
						"short": "Minimum run length to flag a homopolymer.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "construct_qc",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/construct_qc",
								"parts": []any{
									"construct_qc",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"avoidEnzymes": "`reqdata.avoid_enzyme`",
										"crypticOrfMinAa": "`reqdata.cryptic_orf_min_aa`",
										"frameStart": "`reqdata.frame_start`",
										"gcHigh": "`reqdata.gc_high`",
										"gcLow": "`reqdata.gc_low`",
										"gcWindow": "`reqdata.gc_window`",
										"homopolymerMin": "`reqdata.homopolymer_min`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"crispr_grna_design": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "minScore",
						"short": "Only return guides with a heuristic score at least this high (0–100).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "nuclease",
						"short": "Nuclease id.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "searchReverseStrand",
						"short": "Also scan the reverse strand for guides.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "crispr_grna_design",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/crispr_grna_design",
								"parts": []any{
									"crispr_grna_design",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"minScore": "`reqdata.min_score`",
										"nuclease": "`reqdata.nuclease`",
										"searchReverseStrand": "`reqdata.search_reverse_strand`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"crispr_hdr_donor": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "armLength",
						"short": "Homology arm length (bp) on each side.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "blockPam",
						"short": "When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can't be re-cut.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "designGenotypingPrimers",
						"short": "Also design a primer pair (on the original targetSequence) whose product spans the edit site.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "editEnd",
						"short": "1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "editStart",
						"short": "1-based start of the region being replaced.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "frameStart",
						"short": "Optional 1-based CDS reading-frame start; makes the PAM-blocking mutation synonymous where possible.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "guideEnd",
						"short": "1-based forward-strand end of the guide's protospacer.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "guideStart",
						"short": "1-based forward-strand start of the guide's protospacer (alternative to editStart/editEnd, for an insertion exactly at the cut site).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "guideStrand",
						"short": "Strand the guide's protospacer is on.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nuclease",
						"short": "Needed only when deriving the cut site from guideStart/guideEnd/guideStrand.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "replacement",
						"req": true,
						"short": "Sequence to insert/substitute (\"\" for a pure deletion).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "targetSequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "crispr_hdr_donor",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/crispr_hdr_donor",
								"parts": []any{
									"crispr_hdr_donor",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"armLength": "`reqdata.arm_length`",
										"blockPam": "`reqdata.block_pam`",
										"designGenotypingPrimers": "`reqdata.design_genotyping_primer`",
										"editEnd": "`reqdata.edit_end`",
										"editStart": "`reqdata.edit_start`",
										"frameStart": "`reqdata.frame_start`",
										"guideEnd": "`reqdata.guide_end`",
										"guideStart": "`reqdata.guide_start`",
										"guideStrand": "`reqdata.guide_strand`",
										"nuclease": "`reqdata.nuclease`",
										"replacement": "`reqdata.replacement`",
										"targetSequence": "`reqdata.target_sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"crispr_offtarget_check": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxMismatches",
						"short": "Mismatches tolerated between the protospacer and a candidate genomic site.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "nuclease",
						"short": "Nuclease id — determines the PAM pattern/side required at each candidate site.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "protospacer",
						"req": true,
						"short": "The guide's protospacer sequence, 5'→3' (no PAM).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "crispr_offtarget_check",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/crispr_offtarget_check",
								"parts": []any{
									"crispr_offtarget_check",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"maxMismatches": "`reqdata.max_mismatch`",
										"nuclease": "`reqdata.nuclease`",
										"protospacer": "`reqdata.protospacer`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"cross_dimer": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequenceA",
						"req": true,
						"short": "First oligo (5'→3').",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sequenceB",
						"req": true,
						"short": "Second oligo (5'→3').",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "cross_dimer",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/cross_dimer",
								"parts": []any{
									"cross_dimer",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"sequenceA": "`reqdata.sequence_a`",
										"sequenceB": "`reqdata.sequence_b`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"dna_molarity": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "length",
						"short": "Length in bp (dsDNA) or nt (ssDNA/ssRNA).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "massNg",
						"short": "Mass in nanograms.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"short": "Optional sequence — overrides length and gives an exact molar mass from base composition.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"short": "Molecule type.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "volumeUl",
						"short": "Volume in microlitres (0 = unknown; needed for concentration).",
						"type": "`$NUMBER`",
					},
				},
				"name": "dna_molarity",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/dna_molarity",
								"parts": []any{
									"dna_molarity",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"length": "`reqdata.length`",
										"massNg": "`reqdata.mass_ng`",
										"sequence": "`reqdata.sequence`",
										"type": "`reqdata.type`",
										"volumeUl": "`reqdata.volume_ul`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"double_digest": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "enzymeA",
						"req": true,
						"short": "First enzyme name (e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enzymeB",
						"req": true,
						"short": "Second enzyme name (e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "double_digest",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/double_digest",
								"parts": []any{
									"double_digest",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"enzymeA": "`reqdata.enzyme_a`",
										"enzymeB": "`reqdata.enzyme_b`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"export_echo_picklist": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "reactions",
						"req": true,
						"short": "One entry per PCR reaction, up to 96 (a single 96-well plate).",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "export_echo_picklist",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/export_echo_picklist",
								"parts": []any{
									"export_echo_picklist",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"reactions": "`reqdata.reaction`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"export_opentrons_protocol": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "protocolName",
						"short": "Optional protocol name (used in the script's metadata).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "reactions",
						"req": true,
						"short": "One entry per PCR reaction, up to 96 (a single 96-well plate).",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "export_opentrons_protocol",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/export_opentrons_protocol",
								"parts": []any{
									"export_opentrons_protocol",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"protocolName": "`reqdata.protocol_name`",
										"reactions": "`reqdata.reaction`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"export_plate_layout": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "reactions",
						"req": true,
						"short": "One entry per PCR reaction, up to 96 (a single 96-well plate).",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "export_plate_layout",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/export_plate_layout",
								"parts": []any{
									"export_plate_layout",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"reactions": "`reqdata.reaction`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"expression_heatmap_cluster": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "clusterCols",
						"short": "Cluster (reorder) samples.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "clusterRows",
						"short": "Cluster (reorder) genes.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "distanceMetric",
						"short": "correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "genes",
						"req": true,
						"short": "Row (gene) labels.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "linkage",
						"short": "average = UPGMA (standard default), complete = farthest-neighbor, single = nearest-neighbor.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "samples",
						"req": true,
						"short": "Column (sample) labels.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "values",
						"req": true,
						"short": "genes x samples numeric matrix — one row per gene, in the same order as `genes`.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "zScoreRows",
						"short": "Row-wise z-score each gene's values before returning (the conventional 'relative expression' heatmap normalization).",
						"type": "`$BOOLEAN`",
					},
				},
				"name": "expression_heatmap_cluster",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/expression_heatmap_cluster",
								"parts": []any{
									"expression_heatmap_cluster",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"clusterCols": "`reqdata.cluster_col`",
										"clusterRows": "`reqdata.cluster_row`",
										"distanceMetric": "`reqdata.distance_metric`",
										"genes": "`reqdata.gene`",
										"linkage": "`reqdata.linkage`",
										"samples": "`reqdata.sample`",
										"values": "`reqdata.value`",
										"zScoreRows": "`reqdata.z_score_row`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"fastq_qc_report": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "input",
						"req": true,
						"short": "FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "qualityOffset",
						"short": "FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "fastq_qc_report",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/fastq_qc_report",
								"parts": []any{
									"fastq_qc_report",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"input": "`reqdata.input`",
										"qualityOffset": "`reqdata.quality_offset`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"fastq_trim": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "input",
						"req": true,
						"short": "FASTQ text: records of an '@id' header, sequence, '+' separator and quality line (four lines each).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "minLength",
						"short": "Reads shorter than this after trimming are dropped.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "qualityOffset",
						"short": "FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "qualityThreshold",
						"short": "3' quality-trim threshold (Phred score).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "fastq_trim",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/fastq_trim",
								"parts": []any{
									"fastq_trim",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"input": "`reqdata.input`",
										"minLength": "`reqdata.min_length`",
										"qualityOffset": "`reqdata.quality_offset`",
										"qualityThreshold": "`reqdata.quality_threshold`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"find_orf": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "minAaLength",
						"short": "Minimum protein length (aa) to report.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "requireStop",
						"short": "Only report ORFs terminated by a stop codon.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "find_orf",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/find_orfs",
								"parts": []any{
									"find_orfs",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"minAaLength": "`reqdata.min_aa_length`",
										"requireStop": "`reqdata.require_stop`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"format_sequence": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "caseMode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "convert",
						"short": "DNA→RNA (T→U) or RNA→DNA (U→T).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "reverse",
						"short": "Reverse the sequence (no complement).",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "stripNonLetters",
						"short": "Remove digits, spaces and gaps (keep letters only).",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "width",
						"short": "Line-wrap width; 0 = single line.",
						"type": "`$INTEGER`",
					},
				},
				"name": "format_sequence",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/format_sequence",
								"parts": []any{
									"format_sequence",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"caseMode": "`reqdata.case_mode`",
										"convert": "`reqdata.convert`",
										"reverse": "`reqdata.reverse`",
										"sequence": "`reqdata.sequence`",
										"stripNonLetters": "`reqdata.strip_non_letter`",
										"width": "`reqdata.width`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"functional_enrichment": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "background",
						"short": "Custom background/universe gene symbols.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "collections",
						"short": "Which term collections to test.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "genes",
						"req": true,
						"short": "Query gene symbols (human, e.g.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "maxTermSize",
						"short": "Skip terms/pathways with more than this many background genes (matches clusterProfiler's default).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "minTermSize",
						"short": "Skip terms/pathways with fewer than this many background genes.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "functional_enrichment",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/functional_enrichment",
								"parts": []any{
									"functional_enrichment",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"background": "`reqdata.background`",
										"collections": "`reqdata.collection`",
										"genes": "`reqdata.gene`",
										"maxTermSize": "`reqdata.max_term_size`",
										"minTermSize": "`reqdata.min_term_size`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"gc_content": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "gc_content",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/gc_content",
								"parts": []any{
									"gc_content",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"gene_dossier": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gene",
						"req": true,
						"short": "A human gene symbol (\"TP53\") or Ensembl gene ID (\"ENSG00000141510\").",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "gene_dossier",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/gene_dossier",
								"parts": []any{
									"gene_dossier",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"gene": "`reqdata.gene`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"gene_expression": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gene",
						"req": true,
						"short": "A human gene symbol (\"TP53\") or Ensembl gene ID (\"ENSG00000141510\").",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "gene_expression",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/gene_expression",
								"parts": []any{
									"gene_expression",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"gene": "`reqdata.gene`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"gene_model": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gene",
						"req": true,
						"short": "A human gene symbol (\"TP53\") or Ensembl gene ID (\"ENSG00000141510\").",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "gene_model",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/gene_model",
								"parts": []any{
									"gene_model",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"gene": "`reqdata.gene`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"golden_gate_fidelity": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "compareToNamedSet",
						"short": "Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "dataset",
						"short": "Which real ligation dataset to score against — generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "overhangs",
						"req": true,
						"short": "The candidate 4-base overhangs for one assembly (e.g.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "riskThreshold",
						"short": "Flag a pair as risky when the cross-reaction is at least this fraction of that pair's own total signal.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "golden_gate_fidelity",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/golden_gate_fidelity",
								"parts": []any{
									"golden_gate_fidelity",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"compareToNamedSet": "`reqdata.compare_to_named_set`",
										"dataset": "`reqdata.dataset`",
										"overhangs": "`reqdata.overhang`",
										"riskThreshold": "`reqdata.risk_threshold`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"hgvs_convert": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "variant",
						"req": true,
						"short": "A full HGVS \"c.\" variant description: \"<accession or gene symbol>:c.<edit>\", e.g.",
						"type": "`$STRING`",
					},
				},
				"name": "hgvs_convert",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/hgvs_convert",
								"parts": []any{
									"hgvs_convert",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"variant": "`reqdata.variant`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"id_map_poll": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "jobId",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "id_map_poll",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/id_map_poll",
								"parts": []any{
									"id_map_poll",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"jobId": "`reqdata.job_id`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"id_map_submit": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "from",
						"req": true,
						"short": "Source id type: \"Gene_Name\", \"Ensembl\", \"GeneID\", \"RefSeq_Protein\", or \"UniProtKB_AC-ID\".",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ids",
						"req": true,
						"short": "The ids to map, up to 1000 (e.g.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "taxId",
						"short": "NCBI taxonomy id to disambiguate a gene symbol (only used when from=\"Gene_Name\").",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "to",
						"req": true,
						"short": "Target id type.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "id_map_submit",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/id_map_submit",
								"parts": []any{
									"id_map_submit",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"from": "`reqdata.from`",
										"ids": "`reqdata.ids`",
										"taxId": "`reqdata.tax_id`",
										"to": "`reqdata.to`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"in_silico_pcr": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "circular",
						"short": "Treat the template as circular (plasmid).",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "forwardPrimer",
						"req": true,
						"short": "Primer 1, 5'→3'.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxMismatches",
						"short": "Mismatches tolerated per primer.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "reversePrimer",
						"req": true,
						"short": "Primer 2, 5'→3' (order does not matter).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "template",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "in_silico_pcr",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/in_silico_pcr",
								"parts": []any{
									"in_silico_pcr",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"circular": "`reqdata.circular`",
										"forwardPrimer": "`reqdata.forward_primer`",
										"maxMismatches": "`reqdata.max_mismatch`",
										"reversePrimer": "`reqdata.reverse_primer`",
										"template": "`reqdata.template`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"kasp_primer_design": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "addSecondaryMismatch",
						"short": "Engineer the internal ARMS destabilising mismatch near the 3' end.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "alleleA",
						"req": true,
						"short": "First allele (single base) — gets the FAM tail.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "alleleB",
						"req": true,
						"short": "Second allele (single base) — gets the HEX tail.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxAmplicon",
						"short": "Maximum amplicon length for the common reverse primer.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "minAmplicon",
						"short": "Minimum amplicon length for the common reverse primer.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "snpPosition",
						"req": true,
						"short": "1-based position of the SNP on the forward strand.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "targetCoreTm",
						"short": "Target Tm (°C) for the allele-specific primer core (before the universal tail).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "kasp_primer_design",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/kasp_primer_design",
								"parts": []any{
									"kasp_primer_design",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"addSecondaryMismatch": "`reqdata.add_secondary_mismatch`",
										"alleleA": "`reqdata.allele_a`",
										"alleleB": "`reqdata.allele_b`",
										"maxAmplicon": "`reqdata.max_amplicon`",
										"minAmplicon": "`reqdata.min_amplicon`",
										"snpPosition": "`reqdata.snp_position`",
										"target": "`reqdata.target`",
										"targetCoreTm": "`reqdata.target_core_tm`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"list_tool": map[string]any{
				"fields": []any{},
				"name": "list_tool",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/",
								"parts": []any{},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"melting_temperature": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "dntpMM",
						"short": "Total [dNTP] (mM), chelates Mg2+.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "mgMM",
						"short": "Divalent cation [Mg2+] (mM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "naMM",
						"short": "Monovalent cation [Na+]/[K+] (mM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "oligoNM",
						"short": "Total strand concentration (nM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "targetTm",
						"short": "Optional target Tm (°C).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tmTolerance",
						"short": "Allowed +/- window (°C) around targetTm for the gate.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "melting_temperature",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/melting_temperature",
								"parts": []any{
									"melting_temperature",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"dntpMM": "`reqdata.dntp_mm`",
										"mgMM": "`reqdata.mg_mm`",
										"naMM": "`reqdata.na_mm`",
										"oligoNM": "`reqdata.oligo_nm`",
										"sequence": "`reqdata.sequence`",
										"targetTm": "`reqdata.target_tm`",
										"tmTolerance": "`reqdata.tm_tolerance`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"motif_finder": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxMismatches",
						"short": "Maximum allowed mismatches per match.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "motif",
						"req": true,
						"short": "Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "searchReverseStrand",
						"short": "Also search the reverse strand.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "motif_finder",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/motif_finder",
								"parts": []any{
									"motif_finder",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"maxMismatches": "`reqdata.max_mismatch`",
										"motif": "`reqdata.motif`",
										"searchReverseStrand": "`reqdata.search_reverse_strand`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"multiple_sequence_alignment": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "input",
						"req": true,
						"short": "Two or more sequences in multi-FASTA format (>name / sequence).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "multiple_sequence_alignment",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/multiple_sequence_alignment",
								"parts": []any{
									"multiple_sequence_alignment",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"input": "`reqdata.input`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"oligo_analysi": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "dntpMM",
						"short": "Total [dNTP] (mM), chelates Mg2+.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "mgMM",
						"short": "Divalent cation [Mg2+] (mM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "naMM",
						"short": "Monovalent cation [Na+]/[K+] (mM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "oligoNM",
						"short": "Total strand concentration (nM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "oligo_analysi",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/oligo_analysis",
								"parts": []any{
									"oligo_analysis",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"dntpMM": "`reqdata.dntp_mm`",
										"mgMM": "`reqdata.mg_mm`",
										"naMM": "`reqdata.na_mm`",
										"oligoNM": "`reqdata.oligo_nm`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"ortholog_map": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sourceSpecies",
						"short": "Ensembl species slug the symbols belong to (e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "symbols",
						"req": true,
						"short": "Gene symbols to look up, up to 50 (e.g.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "targetSpecies",
						"req": true,
						"short": "Ensembl species slug to find homologs in (e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"short": "Homology type to return.",
						"type": "`$STRING`",
					},
				},
				"name": "ortholog_map",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/ortholog_map",
								"parts": []any{
									"ortholog_map",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"sourceSpecies": "`reqdata.source_species`",
										"symbols": "`reqdata.symbol`",
										"targetSpecies": "`reqdata.target_species`",
										"type": "`reqdata.type`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"pairwise_alignment": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gap",
						"short": "Linear gap penalty (per gap position).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "match",
						"short": "Match score.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "mismatch",
						"short": "Mismatch penalty.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "mode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "seqA",
						"req": true,
						"short": "First sequence (raw or FASTA; nucleotide or protein).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "seqB",
						"req": true,
						"short": "Second sequence (raw or FASTA; nucleotide or protein).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "pairwise_alignment",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/pairwise_alignment",
								"parts": []any{
									"pairwise_alignment",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"gap": "`reqdata.gap`",
										"match": "`reqdata.match`",
										"mismatch": "`reqdata.mismatch`",
										"mode": "`reqdata.mode`",
										"seqA": "`reqdata.seq_a`",
										"seqB": "`reqdata.seq_b`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"parse_genbank": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "text",
						"req": true,
						"short": "A GenBank flat file (LOCUS … FEATURES … ORIGIN … //).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "parse_genbank",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/parse_genbank",
								"parts": []any{
									"parse_genbank",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"text": "`reqdata.text`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"parse_sanger_trace": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "fileBase64",
						"req": true,
						"short": "The binary ABIF (.ab1 / .abi) trace file, base64-encoded.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fileName",
						"short": "Optional original file name (echoed back).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "parse_sanger_trace",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/parse_sanger_trace",
								"parts": []any{
									"parse_sanger_trace",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"fileBase64": "`reqdata.file_base64`",
										"fileName": "`reqdata.file_name`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"plasmid_annotate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "plasmid_annotate",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/plasmid_annotate",
								"parts": []any{
									"plasmid_annotate",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"plasmid_deep_annotate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "circular",
						"short": "Treat the sequence as a circular plasmid (vs.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "plasmid_deep_annotate",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/plasmid_deep_annotate",
								"parts": []any{
									"plasmid_deep_annotate",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"circular": "`reqdata.circular`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"plasmid_full_report": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "circular",
						"short": "Treat the query as a circular molecule (most plasmids are).",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "topN",
						"short": "How many top-ranked backbone candidates to report.",
						"type": "`$INTEGER`",
					},
				},
				"name": "plasmid_full_report",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/plasmid_full_report",
								"parts": []any{
									"plasmid_full_report",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"circular": "`reqdata.circular`",
										"sequence": "`reqdata.sequence`",
										"topN": "`reqdata.top_n`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"plasmid_identify": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "circular",
						"short": "Treat the query as a circular molecule (most plasmids are).",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "topN",
						"short": "How many top-ranked backbone candidates to report.",
						"type": "`$INTEGER`",
					},
				},
				"name": "plasmid_identify",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/plasmid_identify",
								"parts": []any{
									"plasmid_identify",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"circular": "`reqdata.circular`",
										"sequence": "`reqdata.sequence`",
										"topN": "`reqdata.top_n`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"prime_editing_design": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "editEnd",
						"req": true,
						"short": "1-based inclusive end of the region being changed.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "editStart",
						"req": true,
						"short": "1-based inclusive start of the region being changed.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "frameStart",
						"short": "Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "insertedSeq",
						"short": "Replacement bases (forward strand).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "pbsLength",
						"short": "Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "rttHomology",
						"short": "Homology length (nt) 3' of the edit that the RTT should include (typically 10-16).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"short": "Forward-strand target DNA (raw or FASTA), with flanking sequence around the intended edit.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "prime_editing_design",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/prime_editing_design",
								"parts": []any{
									"prime_editing_design",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"editEnd": "`reqdata.edit_end`",
										"editStart": "`reqdata.edit_start`",
										"frameStart": "`reqdata.frame_start`",
										"insertedSeq": "`reqdata.inserted_seq`",
										"pbsLength": "`reqdata.pbs_length`",
										"rttHomology": "`reqdata.rtt_homology`",
										"target": "`reqdata.target`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"prime_editing_twin_design": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "newSequence",
						"req": true,
						"short": "New sequence (forward strand) to install in place of [replaceStart, replaceEnd].",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "overlapLength",
						"short": "Length (bp) of the shared overlap built into both pegRNAs' 3' flaps where they meet and anneal.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "pbsLength",
						"short": "Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "replaceEnd",
						"req": true,
						"short": "1-based inclusive end of the region being replaced/deleted.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "replaceStart",
						"req": true,
						"short": "1-based inclusive start of the region being replaced/deleted.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"short": "Forward-strand target DNA (raw or FASTA), with flanking sequence on both sides of the replacement window.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "prime_editing_twin_design",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/prime_editing_twin_design",
								"parts": []any{
									"prime_editing_twin_design",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"newSequence": "`reqdata.new_sequence`",
										"overlapLength": "`reqdata.overlap_length`",
										"pbsLength": "`reqdata.pbs_length`",
										"replaceEnd": "`reqdata.replace_end`",
										"replaceStart": "`reqdata.replace_start`",
										"target": "`reqdata.target`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"primer_design": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ampliconMax",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ampliconMin",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "dntpMM",
						"short": "Total [dNTP] (mM), chelates Mg2+.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gcMax",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gcMin",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "lenMax",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "lenMin",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "lenOpt",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "maxReturn",
						"short": "Number of best pairs to return.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "mgMM",
						"short": "Divalent cation [Mg2+] (mM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "naMM",
						"short": "Monovalent cation [Na+]/[K+] (mM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "oligoNM",
						"short": "Total strand concentration (nM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "targetEnd",
						"short": "1-based inclusive end of the target region (optional).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "targetStart",
						"short": "1-based inclusive start of a region the product must span (optional).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "template",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tmMax",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tmMaxDiff",
						"short": "Max Tm difference within a pair (°C).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tmMin",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tmOpt",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "primer_design",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/primer_design",
								"parts": []any{
									"primer_design",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"ampliconMax": "`reqdata.amplicon_max`",
										"ampliconMin": "`reqdata.amplicon_min`",
										"dntpMM": "`reqdata.dntp_mm`",
										"gcMax": "`reqdata.gc_max`",
										"gcMin": "`reqdata.gc_min`",
										"lenMax": "`reqdata.len_max`",
										"lenMin": "`reqdata.len_min`",
										"lenOpt": "`reqdata.len_opt`",
										"maxReturn": "`reqdata.max_return`",
										"mgMM": "`reqdata.mg_mm`",
										"naMM": "`reqdata.na_mm`",
										"oligoNM": "`reqdata.oligo_nm`",
										"targetEnd": "`reqdata.target_end`",
										"targetStart": "`reqdata.target_start`",
										"template": "`reqdata.template`",
										"tmMax": "`reqdata.tm_max`",
										"tmMaxDiff": "`reqdata.tm_max_diff`",
										"tmMin": "`reqdata.tm_min`",
										"tmOpt": "`reqdata.tm_opt`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"primer_specificity": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "forwardPrimer",
						"req": true,
						"short": "Forward primer, 5'→3'.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxMismatches",
						"short": "Mismatches tolerated per primer against a reference genome.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "maxProductLength",
						"short": "Ignore candidate off-target products longer than this (bp) — a search-window cap, not a biological claim.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "reversePrimer",
						"req": true,
						"short": "Reverse primer, 5'→3'.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "primer_specificity",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/primer_specificity",
								"parts": []any{
									"primer_specificity",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"forwardPrimer": "`reqdata.forward_primer`",
										"maxMismatches": "`reqdata.max_mismatch`",
										"maxProductLength": "`reqdata.max_product_length`",
										"reversePrimer": "`reqdata.reverse_primer`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"protease_digestion": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxMass",
						"short": "Optional upper bound on neutral monoisotopic mass (Da).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "maxPeptides",
						"short": "Cap on the number of returned peptides.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "minMass",
						"short": "Optional lower bound on neutral monoisotopic mass (Da).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "missedCleavages",
						"short": "Allowed missed internal cleavages (0–2).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "protease",
						"short": "Protease or chemical cleavage agent.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Protein sequence (one-letter amino-acid codes; non-AA characters ignored).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "protease_digestion",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/protease_digestion",
								"parts": []any{
									"protease_digestion",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"maxMass": "`reqdata.max_mass`",
										"maxPeptides": "`reqdata.max_peptide`",
										"minMass": "`reqdata.min_mass`",
										"missedCleavages": "`reqdata.missed_cleavage`",
										"protease": "`reqdata.protease`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"protein_annotate_poll": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "jobId",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "protein_annotate_poll",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/protein_annotate_poll",
								"parts": []any{
									"protein_annotate_poll",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"jobId": "`reqdata.job_id`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"protein_annotate_submit": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "appl",
						"short": "Restrict to one member database (e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "goterms",
						"short": "Include GO-term cross-references.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Protein sequence, one-letter code (FASTA header, if any, is stripped).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "protein_annotate_submit",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/protein_annotate_submit",
								"parts": []any{
									"protein_annotate_submit",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"appl": "`reqdata.appl`",
										"goterms": "`reqdata.goterm`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"protein_hydrophobicity": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "scale",
						"short": "Amino-acid scale.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Protein sequence (one-letter amino-acid codes; non-AA characters ignored).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "window",
						"short": "Sliding-window size (clamped to an odd number ≥ 1).",
						"type": "`$INTEGER`",
					},
				},
				"name": "protein_hydrophobicity",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/protein_hydrophobicity",
								"parts": []any{
									"protein_hydrophobicity",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"scale": "`reqdata.scale`",
										"sequence": "`reqdata.sequence`",
										"window": "`reqdata.window`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"protein_property": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "chargeStep",
						"short": "pH step for the net-charge titration curve (0–14).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Protein sequence (one-letter amino-acid codes; non-AA characters ignored).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "protein_property",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/protein_properties",
								"parts": []any{
									"protein_properties",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"chargeStep": "`reqdata.charge_step`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"random_sequence": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gcContent",
						"short": "Target GC percentage 0..100 (dna/rna only); omit for uniform.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "kind",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "length",
						"req": true,
						"short": "Number of residues to generate.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "random_sequence",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/random_sequence",
								"parts": []any{
									"random_sequence",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"gcContent": "`reqdata.gc_content`",
										"kind": "`reqdata.kind`",
										"length": "`reqdata.length`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"restriction_site": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "enzymes",
						"short": "Enzyme names to scan; omit to scan all curated enzymes.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "restriction_site",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/restriction_sites",
								"parts": []any{
									"restriction_sites",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"enzymes": "`reqdata.enzyme`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"reverse_complement": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
				},
				"name": "reverse_complement",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/reverse_complement",
								"parts": []any{
									"reverse_complement",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"sequence": "`reqdata.sequence`",
										"type": "`reqdata.type`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"reverse_translate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "mode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "organism",
						"short": "Codon-usage host (ignored in degenerate mode).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "protein",
						"req": true,
						"short": "Protein sequence (one-letter codes; * for stop).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "reverse_translate",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/reverse_translate",
								"parts": []any{
									"reverse_translate",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"mode": "`reqdata.mode`",
										"organism": "`reqdata.organism`",
										"protein": "`reqdata.protein`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"rna_fold": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "rna_fold",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/rna_fold",
								"parts": []any{
									"rna_fold",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"sanger_vs_reference": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "fileBase64",
						"short": "The binary ABIF (.ab1 / .abi) trace file, base64-encoded.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fileName",
						"short": "Optional original file name (echoed back).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "minCoverage",
						"short": "Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is 'ambiguous_low_coverage' regardless of identity.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "read",
						"short": "Sanger read as FASTA or raw text (alternative to uploading an ABIF trace).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"req": true,
						"short": "Expected reference sequence (FASTA or raw).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "sanger_vs_reference",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/sanger_vs_reference",
								"parts": []any{
									"sanger_vs_reference",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"fileBase64": "`reqdata.file_base64`",
										"fileName": "`reqdata.file_name`",
										"minCoverage": "`reqdata.min_coverage`",
										"read": "`reqdata.read`",
										"reference": "`reqdata.reference`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"save_permalink": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "args",
						"req": true,
						"short": "Arguments for that tool, exactly as you would pass to it directly.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "save_permalink",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/save_permalink",
								"parts": []any{
									"save_permalink",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"args": "`reqdata.arg`",
										"tool": "`reqdata.tool`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"seqfile_stat": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "input",
						"req": true,
						"short": "FASTA or FASTQ text (raw sequence is treated as single-record FASTA).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "qualityOffset",
						"short": "FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "seqfile_stat",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/seqfile_stats",
								"parts": []any{
									"seqfile_stats",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"input": "`reqdata.input`",
										"qualityOffset": "`reqdata.quality_offset`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"sequence_fetch": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accession",
						"req": true,
						"short": "GenBank/RefSeq accession (e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "db",
						"short": "Database to query; auto-detects from the accession format.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "format",
						"short": "Output format (GenBank is only available for NCBI accessions — UniProt and Ensembl are FASTA-only).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "sequence_fetch",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/sequence_fetch",
								"parts": []any{
									"sequence_fetch",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"accession": "`reqdata.accession`",
										"db": "`reqdata.db`",
										"format": "`reqdata.format`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"sequence_format_convert": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "from",
						"short": "Input format; 'auto' sniffs it from the first meaningful line.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "input",
						"req": true,
						"short": "A FASTA or GenBank record to convert.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "to",
						"short": "Output format.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "sequence_format_convert",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/sequence_format_convert",
								"parts": []any{
									"sequence_format_convert",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"from": "`reqdata.from`",
										"input": "`reqdata.input`",
										"to": "`reqdata.to`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"sequence_report": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "endPrimerLength",
						"short": "Length of the naive end primers taken from each end.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxOrfs",
						"short": "Maximum number of ORFs to return, longest first.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "minOrfAa",
						"short": "Minimum ORF length in amino acids.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "sequence_report",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/sequence_report",
								"parts": []any{
									"sequence_report",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"endPrimerLength": "`reqdata.end_primer_length`",
										"maxOrfs": "`reqdata.max_orf`",
										"minOrfAa": "`reqdata.min_orf_aa`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"sequence_search": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "db",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gene",
						"short": "Gene symbol/name, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "maxResults",
						"short": "Up to 20.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "organism",
						"short": "Organism name, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "term",
						"short": "Raw NCBI search term (advanced) — overrides gene/organism when given, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "sequence_search",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/sequence_search",
								"parts": []any{
									"sequence_search",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"db": "`reqdata.db`",
										"gene": "`reqdata.gene`",
										"maxResults": "`reqdata.max_result`",
										"organism": "`reqdata.organism`",
										"term": "`reqdata.term`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"sequencing_readback_verify": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "minSupportingReads",
						"short": "Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "reads",
						"req": true,
						"short": "Raw reads in FASTA or FASTQ format (auto-detected).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"req": true,
						"short": "The claimed/expected reference sequence.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "sequencing_readback_verify",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/sequencing_readback_verify",
								"parts": []any{
									"sequencing_readback_verify",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"minSupportingReads": "`reqdata.min_supporting_read`",
										"reads": "`reqdata.read`",
										"reference": "`reqdata.reference`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"session_create": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "entries",
						"short": "Initial named entries, e.g.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "session_create",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/session_create",
								"parts": []any{
									"session_create",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"entries": "`reqdata.entry`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"session_get": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "names",
						"short": "Only return these entries; omit to return all of them.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sessionId",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "session_get",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/session_get",
								"parts": []any{
									"session_get",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"names": "`reqdata.name`",
										"sessionId": "`reqdata.session_id`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"session_run": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "args",
						"short": "Additional literal arguments, merged with the ones resolved from the session.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "fromSession",
						"short": "Map of { toolArgName: sessionEntryName } — resolves each named tool argument from the session before running.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sessionId",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "writeBack",
						"short": "Map of { resultFieldName: sessionEntryName } — stores selected fields of the result back into the session under these names.",
						"type": "`$OBJECT`",
					},
				},
				"name": "session_run",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/session_run",
								"parts": []any{
									"session_run",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"args": "`reqdata.arg`",
										"fromSession": "`reqdata.from_session`",
										"sessionId": "`reqdata.session_id`",
										"tool": "`reqdata.tool`",
										"writeBack": "`reqdata.write_back`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"session_set": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "entries",
						"req": true,
						"short": "Named entries to add/overwrite, e.g.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sessionId",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "session_set",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/session_set",
								"parts": []any{
									"session_set",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"entries": "`reqdata.entry`",
										"sessionId": "`reqdata.session_id`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"sirna_design": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "minReynolds",
						"short": "Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "shRnaLoop",
						"short": "Loop sequence used when assembling the shRNA cassette.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "sirna_design",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/sirna_design",
								"parts": []any{
									"sirna_design",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"minReynolds": "`reqdata.min_reynold`",
										"shRnaLoop": "`reqdata.sh_rna_loop`",
										"target": "`reqdata.target`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"site_directed_mutagenesi": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "armTmTarget",
						"short": "Target Tm (°C) for each template-binding arm.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "dntpMM",
						"short": "Total [dNTP] (mM), chelates Mg2+.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "editKind",
						"short": "Edit at the nucleotide or amino-acid level.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "frameStart",
						"short": "1-based position of the first base of codon 1 (editKind='aa').",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "mgMM",
						"short": "Divalent cation [Mg2+] (mM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "naMM",
						"short": "Monovalent cation [Na+]/[K+] (mM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "newBase",
						"short": "Replacement base (editKind='nt').",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "oligoNM",
						"short": "Total strand concentration (nM).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "organism",
						"short": "Codon-usage table for choosing the new codon (editKind='aa').",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "position",
						"short": "1-based position to substitute (editKind='nt').",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "residue",
						"short": "1-based residue number to change (editKind='aa').",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "style",
						"short": "Mutagenic primer style.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "targetAa",
						"short": "Target amino acid, one-letter code incl '*' (editKind='aa').",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "template",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "site_directed_mutagenesi",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/site_directed_mutagenesis",
								"parts": []any{
									"site_directed_mutagenesis",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"armTmTarget": "`reqdata.arm_tm_target`",
										"dntpMM": "`reqdata.dntp_mm`",
										"editKind": "`reqdata.edit_kind`",
										"frameStart": "`reqdata.frame_start`",
										"mgMM": "`reqdata.mg_mm`",
										"naMM": "`reqdata.na_mm`",
										"newBase": "`reqdata.new_base`",
										"oligoNM": "`reqdata.oligo_nm`",
										"organism": "`reqdata.organism`",
										"position": "`reqdata.position`",
										"residue": "`reqdata.residue`",
										"style": "`reqdata.style`",
										"targetAa": "`reqdata.target_aa`",
										"template": "`reqdata.template`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"translate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "frame",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "toStop",
						"short": "Stop at the first stop codon.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "translate",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/translate",
								"parts": []any{
									"translate",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"frame": "`reqdata.frame`",
										"sequence": "`reqdata.sequence`",
										"toStop": "`reqdata.to_stop`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"variant_annotate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "assembly",
						"short": "Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info's native default is hg19).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "variant",
						"req": true,
						"short": "An rsID (\"rs1042522\"), chrom:pos:ref:alt (\"17:7676154:G:C\", single-base substitutions only), genomic HGVS (\"chr17:g.7676154G>C\" or \"17:g.7676154G>C\"), or transcript HGVS c.",
						"type": "`$STRING`",
					},
				},
				"name": "variant_annotate",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/variant_annotate",
								"parts": []any{
									"variant_annotate",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"assembly": "`reqdata.assembly`",
										"variant": "`reqdata.variant`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"variant_comparator": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "coding",
						"short": "Treat as a coding sequence and report amino-acid effects.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "frameStart",
						"short": "1-based reading-frame start (used when coding is true).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "query",
						"req": true,
						"short": "Query / variant sequence (raw or FASTA).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"req": true,
						"short": "Reference / wild-type sequence (raw or FASTA).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "variant_comparator",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/variant_comparator",
								"parts": []any{
									"variant_comparator",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"coding": "`reqdata.coding`",
										"frameStart": "`reqdata.frame_start`",
										"query": "`reqdata.query`",
										"reference": "`reqdata.reference`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"verify_assembly": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "armTmTarget",
						"short": "Target annealing Tm (°C) for primer arms.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "circular",
						"short": "Treat the product/claimed construct as circular (most plasmids are).",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "claimedConstruct",
						"req": true,
						"short": "The sequence you claim you ended up with.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "coding",
						"short": "Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "enzyme",
						"short": "Type IIS enzyme for Golden Gate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enzyme3",
						"short": "3′ enzyme (restriction method).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enzyme5",
						"short": "5′ enzyme (restriction method).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fragmentPcrs",
						"short": "Parallel to fragments, same length: null (or omit) to use fragments[i] directly, or a PCR spec {template, forwardPrimer, reversePrimer, maxMismatches?, circular?} to derive that fragment instead.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "fragments",
						"short": "Fragments (5′→3′), assembled head-to-tail (gibson/goldengate).",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "frameStart",
						"short": "1-based reading-frame start on claimedConstruct, used when coding is true.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "insert",
						"short": "Insert sequence (restriction method).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "insertPcr",
						"short": "Derive the insert by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "method",
						"req": true,
						"short": "Assembly method used.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "names",
						"short": "Optional labels for each fragment.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "overlapLen",
						"short": "Gibson homology-arm length (bp).",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "vector",
						"short": "Vector sequence (restriction method).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "vectorPcr",
						"short": "Derive the vector by PCR instead: {template, forwardPrimer, reversePrimer, maxMismatches?, circular?}.",
						"type": "`$OBJECT`",
					},
				},
				"name": "verify_assembly",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/verify_assembly",
								"parts": []any{
									"verify_assembly",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"armTmTarget": "`reqdata.arm_tm_target`",
										"circular": "`reqdata.circular`",
										"claimedConstruct": "`reqdata.claimed_construct`",
										"coding": "`reqdata.coding`",
										"enzyme": "`reqdata.enzyme`",
										"enzyme3": "`reqdata.enzyme3`",
										"enzyme5": "`reqdata.enzyme5`",
										"fragmentPcrs": "`reqdata.fragment_pcr`",
										"fragments": "`reqdata.fragment`",
										"frameStart": "`reqdata.frame_start`",
										"insert": "`reqdata.insert`",
										"insertPcr": "`reqdata.insert_pcr`",
										"method": "`reqdata.method`",
										"names": "`reqdata.name`",
										"overlapLen": "`reqdata.overlap_len`",
										"vector": "`reqdata.vector`",
										"vectorPcr": "`reqdata.vector_pcr`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"verify_construct": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "claimedConstruct",
						"req": true,
						"short": "The final sequence claimed to have been built.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expectedFrameStart",
						"short": "1-based position in claimedConstruct where the intended reading frame begins.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "insertForwardPrimer",
						"req": true,
						"short": "Forward primer used to amplify the insert, 5'→3'.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "insertReversePrimer",
						"req": true,
						"short": "Reverse primer used to amplify the insert, 5'→3'.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "insertTemplate",
						"req": true,
						"short": "PCR template the insert was amplified from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "maxPrimerMismatches",
						"short": "Mismatches tolerated per primer during PCR prediction.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "templateCircular",
						"short": "Treat insertTemplate as circular (e.g.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "verify_construct",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/verify_construct",
								"parts": []any{
									"verify_construct",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"claimedConstruct": "`reqdata.claimed_construct`",
										"expectedFrameStart": "`reqdata.expected_frame_start`",
										"insertForwardPrimer": "`reqdata.insert_forward_primer`",
										"insertReversePrimer": "`reqdata.insert_reverse_primer`",
										"insertTemplate": "`reqdata.insert_template`",
										"maxPrimerMismatches": "`reqdata.max_primer_mismatch`",
										"templateCircular": "`reqdata.template_circular`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"virtual_gel": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "circular",
						"short": "Treat the sequence as circular (plasmid).",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "enzymes",
						"short": "Enzyme names to digest with.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ladder",
						"short": "DNA ladder to plot alongside the sample lane.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"short": "Nucleotide sequence (raw or FASTA; IUPAC accepted).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "virtual_gel",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/virtual_gel",
								"parts": []any{
									"virtual_gel",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"circular": "`reqdata.circular`",
										"enzymes": "`reqdata.enzyme`",
										"ladder": "`reqdata.ladder`",
										"sequence": "`reqdata.sequence`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"volcano_plot_data": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "rows",
						"req": true,
						"short": "Differential expression rows, one per gene.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "volcano_plot_data",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/volcano_plot_data",
								"parts": []any{
									"volcano_plot_data",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"rows": "`reqdata.row`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"web_search": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "gate",
						"short": "Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (e.g.",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "max_results",
						"short": "Maximum number of results to return (default 5, max 10).",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "provenance",
						"req": true,
						"type": "`$OBJECT`",
						"union": map[string]any{
							"branches": 2,
							"count": 1,
							"depth": 2,
						},
					},
					map[string]any{
						"name": "query",
						"req": true,
						"short": "The search query.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"short": "Tool-specific output object.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"short": "The tool slug that ran.",
						"type": "`$STRING`",
					},
				},
				"name": "web_search",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/web_search",
								"parts": []any{
									"web_search",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"max_results": "`reqdata.max_result`",
										"query": "`reqdata.query`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
