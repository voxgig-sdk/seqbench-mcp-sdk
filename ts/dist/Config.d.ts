import { BaseFeature } from './feature/base/BaseFeature';
declare const FEATURE_PLUGINS: Record<string, any[]>;
declare class Config {
    makeFeature(this: any, fn: string): BaseFeature;
    hasFeature(this: any, fn: string): boolean;
    main: {
        name: string;
        slug: string;
        version: string;
        target: string;
    };
    feature: {
        ratelimit: {
            options: {
                active: boolean;
                burst: number;
                rate: number;
            };
            optspec: {
                now: string;
                sleep: string;
            };
            strict: boolean;
            transport: string;
        };
        retry: {
            options: {
                active: boolean;
                factor: number;
                maxDelay: number;
                minDelay: number;
                retries: number;
                statuses: number[];
            };
            optspec: {
                jitter: string;
                sleep: string;
            };
            strict: boolean;
            transport: string;
        };
        test: {
            options: {
                active: boolean;
            };
            optspec: {
                entity: string;
                net: string;
            };
            strict: boolean;
            transport: string;
        };
        timeout: {
            options: {
                active: boolean;
                ms: number;
            };
            optspec: {
                clearTimer: string;
                setTimer: string;
            };
            strict: boolean;
            transport: string;
        };
    };
    options: {
        base: string;
        auth: {
            prefix: string;
        };
        headers: {
            "content-type": string;
        };
        entity: {
            alphafold_lookup: {};
            aso_design: {};
            base_editing_design: {};
            batch: {};
            batch__workflow: {};
            characterize_sequence: {};
            cloning_simulate: {};
            codon_adaptation_index: {};
            codon_optimize: {};
            construct_autofix: {};
            construct_qc: {};
            crispr_grna_design: {};
            crispr_hdr_donor: {};
            crispr_offtarget_check: {};
            cross_dimer: {};
            dna_molarity: {};
            double_digest: {};
            export_echo_picklist: {};
            export_opentrons_protocol: {};
            export_plate_layout: {};
            expression_heatmap_cluster: {};
            fastq_qc_report: {};
            fastq_trim: {};
            find_orf: {};
            format_sequence: {};
            functional_enrichment: {};
            gc_content: {};
            gene_dossier: {};
            gene_expression: {};
            gene_model: {};
            golden_gate_fidelity: {};
            hgvs_convert: {};
            id_map_poll: {};
            id_map_submit: {};
            in_silico_pcr: {};
            kasp_primer_design: {};
            list_tool: {};
            melting_temperature: {};
            motif_finder: {};
            multiple_sequence_alignment: {};
            oligo_analysi: {};
            ortholog_map: {};
            pairwise_alignment: {};
            parse_genbank: {};
            parse_sanger_trace: {};
            plasmid_annotate: {};
            plasmid_deep_annotate: {};
            plasmid_full_report: {};
            plasmid_identify: {};
            prime_editing_design: {};
            prime_editing_twin_design: {};
            primer_design: {};
            primer_specificity: {};
            protease_digestion: {};
            protein_annotate_poll: {};
            protein_annotate_submit: {};
            protein_hydrophobicity: {};
            protein_property: {};
            random_sequence: {};
            restriction_site: {};
            reverse_complement: {};
            reverse_translate: {};
            rna_fold: {};
            sanger_vs_reference: {};
            save_permalink: {};
            seqfile_stat: {};
            sequence_fetch: {};
            sequence_format_convert: {};
            sequence_report: {};
            sequence_search: {};
            sequencing_readback_verify: {};
            session_create: {};
            session_get: {};
            session_run: {};
            session_set: {};
            sirna_design: {};
            site_directed_mutagenesi: {};
            translate: {};
            variant_annotate: {};
            variant_comparator: {};
            verify_assembly: {};
            verify_construct: {};
            virtual_gel: {};
            volcano_plot_data: {};
            web_search: {};
        };
    };
    entity: {
        alphafold_lookup: {
            fields: ({
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                accession: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        aso_design: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                length: string;
                                target: string;
                                wing: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        base_editing_design: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                editor: string;
                                frameStart: string;
                                target: string;
                                targetPosition: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        batch: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: string;
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
                load: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: string;
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        batch__workflow: {
            fields: ({
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: string;
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
                load: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: string;
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        characterize_sequence: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                endPrimerLength: string;
                                maxOrfs: string;
                                minOrfAa: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        cloning_simulate: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                armTmTarget: string;
                                circular: string;
                                enzyme: string;
                                enzyme3: string;
                                enzyme5: string;
                                fragments: string;
                                insert: string;
                                method: string;
                                names: string;
                                overlapLen: string;
                                vector: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        codon_adaptation_index: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                type: string;
                short?: undefined;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                frameStart: string;
                                organism: string;
                                rareThreshold: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        codon_optimize: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                type: string;
                short?: undefined;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                organism: string;
                                protein: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        construct_autofix: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                type: string;
                short?: undefined;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                avoidEnzymes: string;
                                crypticOrfMinAa: string;
                                frameStart: string;
                                gcHigh: string;
                                gcLow: string;
                                gcWindow: string;
                                homopolymerMin: string;
                                maxPasses: string;
                                organism: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        construct_qc: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                avoidEnzymes: string;
                                crypticOrfMinAa: string;
                                frameStart: string;
                                gcHigh: string;
                                gcLow: string;
                                gcWindow: string;
                                homopolymerMin: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        crispr_grna_design: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                minScore: string;
                                nuclease: string;
                                searchReverseStrand: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        crispr_hdr_donor: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                armLength: string;
                                blockPam: string;
                                designGenotypingPrimers: string;
                                editEnd: string;
                                editStart: string;
                                frameStart: string;
                                guideEnd: string;
                                guideStart: string;
                                guideStrand: string;
                                nuclease: string;
                                replacement: string;
                                targetSequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        crispr_offtarget_check: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                maxMismatches: string;
                                nuclease: string;
                                protospacer: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        cross_dimer: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                sequenceA: string;
                                sequenceB: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        dna_molarity: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                length: string;
                                massNg: string;
                                sequence: string;
                                type: string;
                                volumeUl: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        double_digest: {
            fields: ({
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                enzymeA: string;
                                enzymeB: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        export_echo_picklist: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                reactions: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        export_opentrons_protocol: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                protocolName: string;
                                reactions: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        export_plate_layout: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                reactions: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        expression_heatmap_cluster: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                clusterCols: string;
                                clusterRows: string;
                                distanceMetric: string;
                                genes: string;
                                linkage: string;
                                samples: string;
                                values: string;
                                zScoreRows: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        fastq_qc_report: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                input: string;
                                qualityOffset: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        fastq_trim: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                input: string;
                                minLength: string;
                                qualityOffset: string;
                                qualityThreshold: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        find_orf: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                minAaLength: string;
                                requireStop: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        format_sequence: {
            fields: ({
                name: string;
                type: string;
                short?: undefined;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                caseMode: string;
                                convert: string;
                                reverse: string;
                                sequence: string;
                                stripNonLetters: string;
                                width: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        functional_enrichment: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                background: string;
                                collections: string;
                                genes: string;
                                maxTermSize: string;
                                minTermSize: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        gc_content: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        gene_dossier: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                gene: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        gene_expression: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                gene: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        gene_model: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                gene: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        golden_gate_fidelity: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                compareToNamedSet: string;
                                dataset: string;
                                overhangs: string;
                                riskThreshold: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        hgvs_convert: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                variant: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        id_map_poll: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                jobId: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        id_map_submit: {
            fields: ({
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                from: string;
                                ids: string;
                                taxId: string;
                                to: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        in_silico_pcr: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                circular: string;
                                forwardPrimer: string;
                                maxMismatches: string;
                                reversePrimer: string;
                                template: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        kasp_primer_design: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                addSecondaryMismatch: string;
                                alleleA: string;
                                alleleB: string;
                                maxAmplicon: string;
                                minAmplicon: string;
                                snpPosition: string;
                                target: string;
                                targetCoreTm: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        list_tool: {
            fields: never[];
            name: string;
            op: {
                load: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: never[];
                        select: {};
                        transform: {
                            req: string;
                            res: string;
                        };
                        parts: never[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        melting_temperature: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                dntpMM: string;
                                mgMM: string;
                                naMM: string;
                                oligoNM: string;
                                sequence: string;
                                targetTm: string;
                                tmTolerance: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        motif_finder: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                maxMismatches: string;
                                motif: string;
                                searchReverseStrand: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        multiple_sequence_alignment: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                input: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        oligo_analysi: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                dntpMM: string;
                                mgMM: string;
                                naMM: string;
                                oligoNM: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        ortholog_map: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                sourceSpecies: string;
                                symbols: string;
                                targetSpecies: string;
                                type: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        pairwise_alignment: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                type: string;
                short?: undefined;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                gap: string;
                                match: string;
                                mismatch: string;
                                mode: string;
                                seqA: string;
                                seqB: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        parse_genbank: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                text: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        parse_sanger_trace: {
            fields: ({
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                fileBase64: string;
                                fileName: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        plasmid_annotate: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        plasmid_deep_annotate: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                circular: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        plasmid_full_report: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                circular: string;
                                sequence: string;
                                topN: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        plasmid_identify: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                circular: string;
                                sequence: string;
                                topN: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        prime_editing_design: {
            fields: ({
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                editEnd: string;
                                editStart: string;
                                frameStart: string;
                                insertedSeq: string;
                                pbsLength: string;
                                rttHomology: string;
                                target: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        prime_editing_twin_design: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                newSequence: string;
                                overlapLength: string;
                                pbsLength: string;
                                replaceEnd: string;
                                replaceStart: string;
                                target: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        primer_design: {
            fields: ({
                name: string;
                type: string;
                short?: undefined;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                ampliconMax: string;
                                ampliconMin: string;
                                dntpMM: string;
                                gcMax: string;
                                gcMin: string;
                                lenMax: string;
                                lenMin: string;
                                lenOpt: string;
                                maxReturn: string;
                                mgMM: string;
                                naMM: string;
                                oligoNM: string;
                                targetEnd: string;
                                targetStart: string;
                                template: string;
                                tmMax: string;
                                tmMaxDiff: string;
                                tmMin: string;
                                tmOpt: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        primer_specificity: {
            fields: ({
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                forwardPrimer: string;
                                maxMismatches: string;
                                maxProductLength: string;
                                reversePrimer: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        protease_digestion: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                maxMass: string;
                                maxPeptides: string;
                                minMass: string;
                                missedCleavages: string;
                                protease: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        protein_annotate_poll: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                jobId: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        protein_annotate_submit: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                appl: string;
                                goterms: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        protein_hydrophobicity: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                scale: string;
                                sequence: string;
                                window: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        protein_property: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                chargeStep: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        random_sequence: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                type: string;
                short?: undefined;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                gcContent: string;
                                kind: string;
                                length: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        restriction_site: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                enzymes: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        reverse_complement: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                type: string;
                short?: undefined;
                req?: undefined;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                sequence: string;
                                type: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        reverse_translate: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                type: string;
                short?: undefined;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                mode: string;
                                organism: string;
                                protein: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        rna_fold: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        sanger_vs_reference: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                fileBase64: string;
                                fileName: string;
                                minCoverage: string;
                                read: string;
                                reference: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        save_permalink: {
            fields: ({
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                args: string;
                                tool: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        seqfile_stat: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                input: string;
                                qualityOffset: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        sequence_fetch: {
            fields: ({
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                accession: string;
                                db: string;
                                format: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        sequence_format_convert: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                from: string;
                                input: string;
                                to: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        sequence_report: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                endPrimerLength: string;
                                maxOrfs: string;
                                minOrfAa: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        sequence_search: {
            fields: ({
                name: string;
                type: string;
                short?: undefined;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                db: string;
                                gene: string;
                                maxResults: string;
                                organism: string;
                                term: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        sequencing_readback_verify: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                minSupportingReads: string;
                                reads: string;
                                reference: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        session_create: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                entries: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        session_get: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                names: string;
                                sessionId: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        session_run: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                args: string;
                                fromSession: string;
                                sessionId: string;
                                tool: string;
                                writeBack: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        session_set: {
            fields: ({
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                entries: string;
                                sessionId: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        sirna_design: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                minReynolds: string;
                                shRnaLoop: string;
                                target: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        site_directed_mutagenesi: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                armTmTarget: string;
                                dntpMM: string;
                                editKind: string;
                                frameStart: string;
                                mgMM: string;
                                naMM: string;
                                newBase: string;
                                oligoNM: string;
                                organism: string;
                                position: string;
                                residue: string;
                                style: string;
                                targetAa: string;
                                template: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        translate: {
            fields: ({
                name: string;
                type: string;
                short?: undefined;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                frame: string;
                                sequence: string;
                                toStop: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        variant_annotate: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                assembly: string;
                                variant: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        variant_comparator: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                coding: string;
                                frameStart: string;
                                query: string;
                                reference: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        verify_assembly: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                armTmTarget: string;
                                circular: string;
                                claimedConstruct: string;
                                coding: string;
                                enzyme: string;
                                enzyme3: string;
                                enzyme5: string;
                                fragmentPcrs: string;
                                fragments: string;
                                frameStart: string;
                                insert: string;
                                insertPcr: string;
                                method: string;
                                names: string;
                                overlapLen: string;
                                vector: string;
                                vectorPcr: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        verify_construct: {
            fields: ({
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            } | {
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                claimedConstruct: string;
                                expectedFrameStart: string;
                                insertForwardPrimer: string;
                                insertReversePrimer: string;
                                insertTemplate: string;
                                maxPrimerMismatches: string;
                                templateCircular: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        virtual_gel: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                circular: string;
                                enzymes: string;
                                ladder: string;
                                sequence: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        volcano_plot_data: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                rows: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
        web_search: {
            fields: ({
                name: string;
                short: string;
                type: string;
                req?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                short?: undefined;
                union?: undefined;
            } | {
                name: string;
                req: boolean;
                type: string;
                union: {
                    branches: number;
                    count: number;
                    depth: number;
                };
                short?: undefined;
            } | {
                name: string;
                req: boolean;
                short: string;
                type: string;
                union?: undefined;
            })[];
            name: string;
            op: {
                create: {
                    input: string;
                    name: string;
                    points: {
                        args: {};
                        kind: string;
                        method: string;
                        orig: string;
                        segments: {
                            lit: string;
                        }[];
                        select: {};
                        transform: {
                            req: {
                                max_results: string;
                                query: string;
                            };
                            res: string;
                        };
                        parts: string[];
                    }[];
                };
            };
            relations: {
                ancestors: never[];
            };
        };
    };
}
declare const config: Config;
export { config, FEATURE_PLUGINS, };
