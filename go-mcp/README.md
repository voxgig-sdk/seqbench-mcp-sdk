# seqbench-mcp-mcp

[MCP](https://modelcontextprotocol.io) server exposing the SeqbenchMcp SDK as
two agent tools — `seqbench-mcp_list` and `seqbench-mcp_load` — built on the
[official Go MCP SDK](https://github.com/modelcontextprotocol/go-sdk) and the
sibling Go SDK at `../go`. Runs over **stdio** (default, for spawnable installs)
or **streamable HTTP** (one shared server for several agents).

## Examples

```sh
# 1. Build a native binary (-> dist/<os>-<arch>/seqbench-mcp-mcp)
make build

# 2. Provide credentials via the environment
export SEQBENCH_MCP_APIKEY=sk_live_xxx

# 3a. Install into Claude Code over stdio (most common)
claude mcp add --scope user seqbench-mcp \
  -- /absolute/path/to/seqbench-mcp-mcp -transport stdio

# 3b. …or run a shared HTTP server instead
./seqbench-mcp-mcp -transport http -addr :8080
```

Tool-call arguments (what an agent sends):

```jsonc
// seqbench-mcp_list: first page of records
{ "entity": "alphafold_lookup" }
{ "entity": "alphafold_lookup", "query": { } }

// seqbench-mcp_load: one record by id
{ "entity": "batch", "query": { "id": 1 } }
```

> The rest of this guide follows the [Diátaxis](https://diataxis.fr) framework:
> a hands-on **Tutorial**, task-focused **How-to guides**, a factual
> **Reference**, and background **Explanation**.

## Tutorial: install and call a tool

1. **Build** the server from this `go-mcp/` directory:

   ```sh
   make build          # -> dist/<os>-<arch>/seqbench-mcp-mcp
   ```

2. **Set your API key:**

   ```sh
   export SEQBENCH_MCP_APIKEY=sk_live_xxx
   ```

3. **Install it into Claude Code** (stdio transport):

   ```sh
   claude mcp add --scope user seqbench-mcp \
     -- "$PWD"/dist/*/seqbench-mcp-mcp -transport stdio
   ```

4. **Restart Claude Code.** The `seqbench-mcp_list` and `seqbench-mcp_load` tools now appear
   in new sessions. Ask the agent to *"list alphafold_lookup using seqbench-mcp"*
   and it calls `seqbench-mcp_list` with `{"entity":"alphafold_lookup"}`.

## How-to guides

### Authenticate and choose an environment

Configuration is read from the environment — nothing is written to disk:

```sh
export SEQBENCH_MCP_APIKEY=sk_live_xxx            # API key
export SEQBENCH_MCP_BASE=https://api.example.com  # optional: override the API base URL
```

Set these in the shell that launches the server (or in the `claude mcp add`
environment) so every tool call is authenticated.

### Run as a shared HTTP server

```sh
./seqbench-mcp-mcp -transport http -addr :8080
```

Streamable HTTP lets several agents share one running process; stdio (the
default) spawns a fresh process per client.

### Call the `seqbench-mcp_list` tool

Args: `entity` (required), `query` (optional filter map). Returns the first
page of records as JSON:

```jsonc
{ "entity": "alphafold_lookup" }
```

### Call the `seqbench-mcp_load` tool

Args: `entity` (required), `query` = `{"id":N}` (required). Returns the single
record as JSON:

```jsonc
{ "entity": "batch", "query": { "id": 1 } }
```

### Cross-compile release binaries

```sh
make build       # native binary for this machine
make build-all   # linux/darwin/windows x amd64/arm64, under dist/<os>-<arch>/
```

## Reference

### Tools

| Tool | Args | Returns |
|------|------|---------|
| `seqbench-mcp_list` | `entity` (required), `query` (optional map) | First page of records as JSON |
| `seqbench-mcp_load` | `entity` (required), `query` = `{id:N}` | Single record as JSON |

On error, a tool returns an MCP error result (`isError: true`) whose text is the
failure message (e.g. unknown entity, or an API error).

### `Args` schema

Both tools take the same argument object:

| Field | Type | Notes |
|-------|------|-------|
| `entity` | string | One of the 85 supported entities (see below). |
| `query` | object | Optional match map. `{"id":N}` for load; omit or `{}` for list. |

JSON schemas are emitted by the SDK from the `Args` struct's `json` /
`jsonschema` tags — no schema is hand-written.

### Transports & flags

| Flag | Default | Purpose |
|------|---------|---------|
| `-transport` | `stdio` | `stdio` (spawnable) or `http` (streamable HTTP). |
| `-addr` | `:8080` | Listen address for the `http` transport. |

### Environment variables

| Variable | Purpose |
|----------|---------|
| `SEQBENCH_MCP_APIKEY` | API key sent with every request. |
| `SEQBENCH_MCP_BASE` | Optional override of the API base URL. |

### Entities

The 85 entities valid as the `entity` argument:

alphafold_lookup | aso_design | base_editing_design | batch | batch__workflow | characterize_sequence | cloning_simulate | codon_adaptation_index | codon_optimize | construct_autofix | construct_qc | crispr_grna_design | crispr_hdr_donor | crispr_offtarget_check | cross_dimer | dna_molarity | double_digest | export_echo_picklist | export_opentrons_protocol | export_plate_layout | expression_heatmap_cluster | fastq_qc_report | fastq_trim | find_orf | format_sequence | functional_enrichment | gc_content | gene_dossier | gene_expression | gene_model | golden_gate_fidelity | hgvs_convert | id_map_poll | id_map_submit | in_silico_pcr | kasp_primer_design | list_tool | melting_temperature | motif_finder | multiple_sequence_alignment | oligo_analysi | ortholog_map | pairwise_alignment | parse_genbank | parse_sanger_trace | plasmid_annotate | plasmid_deep_annotate | plasmid_full_report | plasmid_identify | prime_editing_design | prime_editing_twin_design | primer_design | primer_specificity | protease_digestion | protein_annotate_poll | protein_annotate_submit | protein_hydrophobicity | protein_property | random_sequence | restriction_site | reverse_complement | reverse_translate | rna_fold | sanger_vs_reference | save_permalink | seqfile_stat | sequence_fetch | sequence_format_convert | sequence_report | sequence_search | sequencing_readback_verify | session_create | session_get | session_run | session_set | sirna_design | site_directed_mutagenesi | translate | variant_annotate | variant_comparator | verify_assembly | verify_construct | virtual_gel | volcano_plot_data | web_search

### Smoke test via HTTP (raw JSON-RPC)

```sh
./seqbench-mcp-mcp -transport http -addr :18080 &

# initialize, grab the session id
curl -sN -X POST http://localhost:18080 \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json, text/event-stream' \
  -D headers \
  -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-06-18","capabilities":{},"clientInfo":{"name":"smoke","version":"0"}}}'

SESSION=$(awk '/Mcp-Session-Id/ {print $2}' headers | tr -d '\r')

curl -sN -X POST http://localhost:18080 \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json, text/event-stream' \
  -H "Mcp-Session-Id: $SESSION" \
  -d '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"seqbench-mcp_load","arguments":{"entity":"batch","query":{"id":1}}}}'
```

## Explanation

### How tools map to the SDK

`main.go` builds the SDK client (configured from the environment) and registers
two tools. Each dispatches on the `entity` argument to the matching entity in
the sibling Go SDK at `../go`, calls `List` or `Load`, unwraps the `Entity`
wrappers to plain data, and returns it as pretty-printed JSON.

### Why two transports

**stdio** is the standard for agent hosts that spawn a server per client
(Claude Code's `claude mcp add`). **streamable HTTP** keeps one process running
that many agents can share — handy for a long-lived deployment.

### Schema generation

The input schema is derived from the `Args` Go struct's `json` / `jsonschema`
tags at registration time, so the advertised tool schema can never drift from
the code that consumes it.

## Generated by

sdkgen `go-mcp` target. See the target source under `.sdk/src/cmp/go-mcp/` in
this repo, or upstream at
`github.com/voxgig/sdkgen/project/.sdk/src/cmp/go-mcp/`.
