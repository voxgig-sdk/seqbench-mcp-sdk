# seqbench-mcp-cli

AQL-driven command-line client **and** interactive REPL for the SeqbenchMcp
SDK. Each command line is parsed as a single [AQL](https://github.com/aql-lang/aql)
expression and evaluated against the live API; run it with no arguments to drop
into a REPL. Built on `github.com/aql-lang/aql/eng/go` and the sibling Go SDK
at `../go`.

## Examples

```sh
# 1. Build a native binary (-> dist/<os>-<arch>/seqbench-mcp-cli)
make build

# 2. See usage (words, entities, env vars)
./seqbench-mcp-cli --help

# 3. Provide credentials once, via the environment
export SEQBENCH_MCP_APIKEY=sk_live_xxx

# 4. Each command line is ONE AQL expression, run against the API:

# 5. Override the API base URL for a single call
SEQBENCH_MCP_BASE=https://api.example.com ./seqbench-mcp-cli --help

# 6. No arguments -> interactive REPL
./seqbench-mcp-cli
seqbench-mcp> /help
seqbench-mcp> /quit
```

> The rest of this guide follows the [Diátaxis](https://diataxis.fr) framework:
> a hands-on **Tutorial**, task-focused **How-to guides**, a factual
> **Reference**, and background **Explanation**.

## Tutorial: your first query in under a minute

1. **Build the binary.** From this `go-cli/` directory:

   ```sh
   make build          # -> dist/<os>-<arch>/seqbench-mcp-cli
   ```

2. **Set your API key** (read from the environment):

   ```sh
   export SEQBENCH_MCP_APIKEY=sk_live_xxx
   ```

3. **Run a query.** Evaluate an AQL expression against the API (or run with no
   arguments to open the REPL):

   ```sh
   ./dist/*/seqbench-mcp-cli --help
   ```

4. **Go interactive.** Run the binary with no arguments to open the REPL, then
   type `/help` for the word and entity lists and `/quit` to leave.

That is the whole loop: *build → set key → evaluate AQL expressions*.

## How-to guides

### Authenticate and choose an environment

Configuration is read from the environment — nothing is written to disk:

```sh
export SEQBENCH_MCP_APIKEY=sk_live_xxx            # API key
export SEQBENCH_MCP_BASE=https://api.example.com  # optional: override the API base URL
./seqbench-mcp-cli --help
```

Both are injectable by a secrets vault, so the key never has to be typed inline.

### Explore interactively with the REPL

Run with no arguments to open a REPL (prompt `seqbench-mcp>`). Each line is
evaluated as its own AQL expression:

```text
$ ./seqbench-mcp-cli
seqbench-mcp> /help
seqbench-mcp> /quit
```

### Cross-compile release binaries

```sh
make build       # native binary for this machine
make build-all   # linux/darwin/windows x amd64/arm64, under dist/<os>-<arch>/
```

### Discover the available entities

`/help` in the REPL prints the full entity list, or see [Entities](#entities)
below — this SDK exposes 85 entities.

## Reference

### Words

The CLI registers these AQL words, each bound to the SDK:

| Word     | Signatures                                    | Returns                        |
|----------|-----------------------------------------------|--------------------------------|
| `load`   | `load <entity>` · `load <query> <entity>`     | A single record                |

- `<entity>` is a bareword, auto-quoted as an AQL atom (e.g. `alphafold_lookup`).
- `<query>` is either a **Map** (`{id:1}`) or a **Scalar** (`1`, treated as
  `{id:1}`). A scalar is always wrapped as `{id:<value>}`.

### Environment variables

| Variable | Purpose |
|----------|---------|
| `SEQBENCH_MCP_APIKEY` | API key sent with every request. |
| `SEQBENCH_MCP_BASE` | Optional override of the API base URL. |

Unset variables fall back to the SDK's built-in defaults.

### CLI flags

- `--help` / `-h` — print usage (words, entities, env vars) and exit.

### REPL commands

Meta-commands use the `/` prefix (everything else on a line is evaluated as AQL):

- `/quit` / `/q` / `/exit` — exit the REPL
- `/help` / `/h` / `/?`     — show the word list, entity list and meta commands

### Exit codes

| Code | Meaning |
|------|---------|
| `0` | Success (also the normal REPL exit). |
| `1` | Parse error, word-registration error, or an API/evaluation error. |

### Build targets

| Target | Result |
|--------|--------|
| `make build` | Native binary at `dist/<os>-<arch>/seqbench-mcp-cli`. |
| `make build-all` | linux/darwin/windows x amd64/arm64, each under its own `dist/<os>-<arch>/`. |
| `make clean` | Remove `dist/` and any stray binaries. |

### Entities

The 85 entities this SDK exposes (any is valid as `<entity>`):

alphafold_lookup aso_design base_editing_design batch batch__workflow characterize_sequence cloning_simulate codon_adaptation_index codon_optimize construct_autofix construct_qc crispr_grna_design crispr_hdr_donor crispr_offtarget_check cross_dimer dna_molarity double_digest export_echo_picklist export_opentrons_protocol export_plate_layout expression_heatmap_cluster fastq_qc_report fastq_trim find_orf format_sequence functional_enrichment gc_content gene_dossier gene_expression gene_model golden_gate_fidelity hgvs_convert id_map_poll id_map_submit in_silico_pcr kasp_primer_design list_tool melting_temperature motif_finder multiple_sequence_alignment oligo_analysi ortholog_map pairwise_alignment parse_genbank parse_sanger_trace plasmid_annotate plasmid_deep_annotate plasmid_full_report plasmid_identify prime_editing_design prime_editing_twin_design primer_design primer_specificity protease_digestion protein_annotate_poll protein_annotate_submit protein_hydrophobicity protein_property random_sequence restriction_site reverse_complement reverse_translate rna_fold sanger_vs_reference save_permalink seqfile_stat sequence_fetch sequence_format_convert sequence_report sequence_search sequencing_readback_verify session_create session_get session_run session_set sirna_design site_directed_mutagenesi translate variant_annotate variant_comparator verify_assembly verify_construct virtual_gel volcano_plot_data web_search

## Explanation

### Why AQL?

The whole command line is one [AQL](https://github.com/aql-lang/aql) expression,
not a fixed `verb --flag` grammar. That means the same binary works one-shot
(`./seqbench-mcp-cli <expr>`) and interactively (the REPL), and expressions compose the
same way in both. `list` / `load` / `update` are ordinary AQL *words* bound to
the SDK — adding SDK operations is adding words, not re-parsing flags.

### How it is wired

`main.go` builds the SDK client (configured from the environment), creates an
AQL registry, and `words.go` registers `list` / `load` / `update` as native
words that dispatch on the entity atom and call the sibling Go SDK at `../go`.
Results are unwrapped from their `Entity` wrappers to plain data before being
printed.

### Output format

Each result value is printed as its AQL string form (a JSON-like rendering of
the record or list of records). One-shot mode prints to stdout; errors go to
stderr with a non-zero exit code.

## Generated by

sdkgen `go-cli` target. See the target source under `.sdk/src/cmp/go-cli/` in
this repo, or upstream at
`github.com/voxgig/sdkgen/project/.sdk/src/cmp/go-cli/`.
