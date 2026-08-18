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
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "length",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "wing",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "frameStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "targetPosition",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "capped",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "limit",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "limit",
						"req": true,
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
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxOrfs",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "minOrfAa",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "circular",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "enzyme",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enzyme3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enzyme5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fragments",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "insert",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "method",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "names",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "overlapLen",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "vector",
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
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "crypticOrfMinAa",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "frameStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$INTEGER`",
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
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "crypticOrfMinAa",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "frameStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "minScore",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "nuclease",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "searchReverseStrand",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "blockPam",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "designGenotypingPrimers",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "editEnd",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "editStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "frameStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "guideEnd",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "guideStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "guideStrand",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nuclease",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "targetSequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxMismatches",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "nuclease",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequenceA",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sequenceB",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "length",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "massNg",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "volumeUl",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enzymeB",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "protocolName",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "clusterRows",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "distanceMetric",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "genes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "linkage",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "samples",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "values",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "zScoreRows",
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "input",
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
						"name": "qualityOffset",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "input",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "minLength",
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
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "qualityThreshold",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "minAaLength",
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
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "reverse",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "stripNonLetters",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "width",
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
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "collections",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "genes",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "maxTermSize",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "minTermSize",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gene",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gene",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gene",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "dataset",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "riskThreshold",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "variant",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ids",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "taxId",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "to",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "forwardPrimer",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxMismatches",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "reversePrimer",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "template",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "alleleA",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "alleleB",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxAmplicon",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "minAmplicon",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "snpPosition",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "targetCoreTm",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "mgMM",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "naMM",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "oligoNM",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "targetTm",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tmTolerance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxMismatches",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "motif",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "searchReverseStrand",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "input",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "mgMM",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "naMM",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "oligoNM",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sourceSpecies",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "symbols",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "targetSpecies",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "match",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "mismatch",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "seqA",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "seqB",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "text",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fileName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "topN",
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
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "topN",
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
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "editStart",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "frameStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "insertedSeq",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "pbsLength",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "rttHomology",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "newSequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "overlapLength",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "pbsLength",
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
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "replaceStart",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "mgMM",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "naMM",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "oligoNM",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "targetEnd",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "targetStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "template",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tmMax",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "tmMaxDiff",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxMismatches",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "maxProductLength",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "reversePrimer",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxMass",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "maxPeptides",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "minMass",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "missedCleavages",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "protease",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "goterms",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "scale",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "window",
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gcContent",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "kind",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "length",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "protein",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fileName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "minCoverage",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "input",
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
						"name": "qualityOffset",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "db",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "input",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "to",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "maxOrfs",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "minOrfAa",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "gene",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "maxResults",
						"type": "`$INTEGER`",
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
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "term",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "minSupportingReads",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "names",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "fromSession",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "writeBack",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "minReynolds",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "shRnaLoop",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "target",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "dntpMM",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "editKind",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "frameStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "mgMM",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "naMM",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "newBase",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "oligoNM",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "organism",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "position",
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
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "style",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "targetAa",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "template",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "toStop",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "variant",
						"req": true,
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
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "frameStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "circular",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "claimedConstruct",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "coding",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "enzyme",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enzyme3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enzyme5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fragmentPcrs",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "fragments",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "frameStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "insert",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "insertPcr",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "method",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "names",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ok",
						"req": true,
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "overlapLen",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "vector",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "vectorPcr",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expectedFrameStart",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "insertForwardPrimer",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "insertReversePrimer",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "insertTemplate",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "maxPrimerMismatches",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "templateCircular",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "enzymes",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "gate",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "ladder",
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sequence",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "rows",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "max_results",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "tool",
						"req": true,
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
