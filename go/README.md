# SeqbenchMcp Golang SDK



The Golang SDK for the SeqbenchMcp API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.AlphafoldLookup(nil)` — each with the same small set of operations (`Load`, `Create`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/seqbench-mcp-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/seqbench-mcp-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/seqbench-mcp-sdk/go=../seqbench-mcp-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/seqbench-mcp-sdk/go"
)

func main() {
    client := sdk.NewSeqbenchMcpSDK(map[string]any{
        "apikey": os.Getenv("SEQBENCH_MCP_APIKEY"),
    })

    // Create a alphafoldLookup.
    created, err := client.AlphafoldLookup(nil).Create(map[string]any{"accession": "example_accession", "ok": "example_ok", "provenance": map[string]any{}, "result": map[string]any{}, "tool": "example_tool"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created)
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
batch, err := client.Batch(nil).Load(nil, nil)
if err != nil {
    // handle err
    return
}
_ = batch
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

batch, err := client.Batch(nil).Load(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(batch) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewSeqbenchMcpSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
SEQBENCH_MCP_TEST_LIVE=TRUE
SEQBENCH_MCP_APIKEY=<your-key>
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewSeqbenchMcpSDK

```go
func NewSeqbenchMcpSDK(options map[string]any) *SeqbenchMcpSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *SeqbenchMcpSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### SeqbenchMcpSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `AlphafoldLookup` | `(data map[string]any) SeqbenchMcpEntity` | Create an AlphafoldLookup entity instance. |
| `AsoDesign` | `(data map[string]any) SeqbenchMcpEntity` | Create an AsoDesign entity instance. |
| `BaseEditingDesign` | `(data map[string]any) SeqbenchMcpEntity` | Create a BaseEditingDesign entity instance. |
| `Batch` | `(data map[string]any) SeqbenchMcpEntity` | Create a Batch entity instance. |
| `BatchWorkflow` | `(data map[string]any) SeqbenchMcpEntity` | Create a BatchWorkflow entity instance. |
| `CharacterizeSequence` | `(data map[string]any) SeqbenchMcpEntity` | Create a CharacterizeSequence entity instance. |
| `CloningSimulate` | `(data map[string]any) SeqbenchMcpEntity` | Create a CloningSimulate entity instance. |
| `CodonAdaptationIndex` | `(data map[string]any) SeqbenchMcpEntity` | Create a CodonAdaptationIndex entity instance. |
| `CodonOptimize` | `(data map[string]any) SeqbenchMcpEntity` | Create a CodonOptimize entity instance. |
| `ConstructAutofix` | `(data map[string]any) SeqbenchMcpEntity` | Create a ConstructAutofix entity instance. |
| `ConstructQc` | `(data map[string]any) SeqbenchMcpEntity` | Create a ConstructQc entity instance. |
| `CrisprGrnaDesign` | `(data map[string]any) SeqbenchMcpEntity` | Create a CrisprGrnaDesign entity instance. |
| `CrisprHdrDonor` | `(data map[string]any) SeqbenchMcpEntity` | Create a CrisprHdrDonor entity instance. |
| `CrisprOfftargetCheck` | `(data map[string]any) SeqbenchMcpEntity` | Create a CrisprOfftargetCheck entity instance. |
| `CrossDimer` | `(data map[string]any) SeqbenchMcpEntity` | Create a CrossDimer entity instance. |
| `DnaMolarity` | `(data map[string]any) SeqbenchMcpEntity` | Create a DnaMolarity entity instance. |
| `DoubleDigest` | `(data map[string]any) SeqbenchMcpEntity` | Create a DoubleDigest entity instance. |
| `ExportEchoPicklist` | `(data map[string]any) SeqbenchMcpEntity` | Create an ExportEchoPicklist entity instance. |
| `ExportOpentronsProtocol` | `(data map[string]any) SeqbenchMcpEntity` | Create an ExportOpentronsProtocol entity instance. |
| `ExportPlateLayout` | `(data map[string]any) SeqbenchMcpEntity` | Create an ExportPlateLayout entity instance. |
| `ExpressionHeatmapCluster` | `(data map[string]any) SeqbenchMcpEntity` | Create an ExpressionHeatmapCluster entity instance. |
| `FastqQcReport` | `(data map[string]any) SeqbenchMcpEntity` | Create a FastqQcReport entity instance. |
| `FastqTrim` | `(data map[string]any) SeqbenchMcpEntity` | Create a FastqTrim entity instance. |
| `FindOrf` | `(data map[string]any) SeqbenchMcpEntity` | Create a FindOrf entity instance. |
| `FormatSequence` | `(data map[string]any) SeqbenchMcpEntity` | Create a FormatSequence entity instance. |
| `FunctionalEnrichment` | `(data map[string]any) SeqbenchMcpEntity` | Create a FunctionalEnrichment entity instance. |
| `GcContent` | `(data map[string]any) SeqbenchMcpEntity` | Create a GcContent entity instance. |
| `GeneDossier` | `(data map[string]any) SeqbenchMcpEntity` | Create a GeneDossier entity instance. |
| `GeneExpression` | `(data map[string]any) SeqbenchMcpEntity` | Create a GeneExpression entity instance. |
| `GeneModel` | `(data map[string]any) SeqbenchMcpEntity` | Create a GeneModel entity instance. |
| `GoldenGateFidelity` | `(data map[string]any) SeqbenchMcpEntity` | Create a GoldenGateFidelity entity instance. |
| `HgvsConvert` | `(data map[string]any) SeqbenchMcpEntity` | Create a HgvsConvert entity instance. |
| `IdMapPoll` | `(data map[string]any) SeqbenchMcpEntity` | Create an IdMapPoll entity instance. |
| `IdMapSubmit` | `(data map[string]any) SeqbenchMcpEntity` | Create an IdMapSubmit entity instance. |
| `InSilicoPcr` | `(data map[string]any) SeqbenchMcpEntity` | Create an InSilicoPcr entity instance. |
| `KaspPrimerDesign` | `(data map[string]any) SeqbenchMcpEntity` | Create a KaspPrimerDesign entity instance. |
| `ListTool` | `(data map[string]any) SeqbenchMcpEntity` | Create a ListTool entity instance. |
| `MeltingTemperature` | `(data map[string]any) SeqbenchMcpEntity` | Create a MeltingTemperature entity instance. |
| `MotifFinder` | `(data map[string]any) SeqbenchMcpEntity` | Create a MotifFinder entity instance. |
| `MultipleSequenceAlignment` | `(data map[string]any) SeqbenchMcpEntity` | Create a MultipleSequenceAlignment entity instance. |
| `OligoAnalysi` | `(data map[string]any) SeqbenchMcpEntity` | Create an OligoAnalysi entity instance. |
| `OrthologMap` | `(data map[string]any) SeqbenchMcpEntity` | Create an OrthologMap entity instance. |
| `PairwiseAlignment` | `(data map[string]any) SeqbenchMcpEntity` | Create a PairwiseAlignment entity instance. |
| `ParseGenbank` | `(data map[string]any) SeqbenchMcpEntity` | Create a ParseGenbank entity instance. |
| `ParseSangerTrace` | `(data map[string]any) SeqbenchMcpEntity` | Create a ParseSangerTrace entity instance. |
| `PlasmidAnnotate` | `(data map[string]any) SeqbenchMcpEntity` | Create a PlasmidAnnotate entity instance. |
| `PlasmidDeepAnnotate` | `(data map[string]any) SeqbenchMcpEntity` | Create a PlasmidDeepAnnotate entity instance. |
| `PlasmidFullReport` | `(data map[string]any) SeqbenchMcpEntity` | Create a PlasmidFullReport entity instance. |
| `PlasmidIdentify` | `(data map[string]any) SeqbenchMcpEntity` | Create a PlasmidIdentify entity instance. |
| `PrimeEditingDesign` | `(data map[string]any) SeqbenchMcpEntity` | Create a PrimeEditingDesign entity instance. |
| `PrimeEditingTwinDesign` | `(data map[string]any) SeqbenchMcpEntity` | Create a PrimeEditingTwinDesign entity instance. |
| `PrimerDesign` | `(data map[string]any) SeqbenchMcpEntity` | Create a PrimerDesign entity instance. |
| `PrimerSpecificity` | `(data map[string]any) SeqbenchMcpEntity` | Create a PrimerSpecificity entity instance. |
| `ProteaseDigestion` | `(data map[string]any) SeqbenchMcpEntity` | Create a ProteaseDigestion entity instance. |
| `ProteinAnnotatePoll` | `(data map[string]any) SeqbenchMcpEntity` | Create a ProteinAnnotatePoll entity instance. |
| `ProteinAnnotateSubmit` | `(data map[string]any) SeqbenchMcpEntity` | Create a ProteinAnnotateSubmit entity instance. |
| `ProteinHydrophobicity` | `(data map[string]any) SeqbenchMcpEntity` | Create a ProteinHydrophobicity entity instance. |
| `ProteinProperty` | `(data map[string]any) SeqbenchMcpEntity` | Create a ProteinProperty entity instance. |
| `RandomSequence` | `(data map[string]any) SeqbenchMcpEntity` | Create a RandomSequence entity instance. |
| `RestrictionSite` | `(data map[string]any) SeqbenchMcpEntity` | Create a RestrictionSite entity instance. |
| `ReverseComplement` | `(data map[string]any) SeqbenchMcpEntity` | Create a ReverseComplement entity instance. |
| `ReverseTranslate` | `(data map[string]any) SeqbenchMcpEntity` | Create a ReverseTranslate entity instance. |
| `RnaFold` | `(data map[string]any) SeqbenchMcpEntity` | Create a RnaFold entity instance. |
| `SangerVsReference` | `(data map[string]any) SeqbenchMcpEntity` | Create a SangerVsReference entity instance. |
| `SavePermalink` | `(data map[string]any) SeqbenchMcpEntity` | Create a SavePermalink entity instance. |
| `SeqfileStat` | `(data map[string]any) SeqbenchMcpEntity` | Create a SeqfileStat entity instance. |
| `SequenceFetch` | `(data map[string]any) SeqbenchMcpEntity` | Create a SequenceFetch entity instance. |
| `SequenceFormatConvert` | `(data map[string]any) SeqbenchMcpEntity` | Create a SequenceFormatConvert entity instance. |
| `SequenceReport` | `(data map[string]any) SeqbenchMcpEntity` | Create a SequenceReport entity instance. |
| `SequenceSearch` | `(data map[string]any) SeqbenchMcpEntity` | Create a SequenceSearch entity instance. |
| `SequencingReadbackVerify` | `(data map[string]any) SeqbenchMcpEntity` | Create a SequencingReadbackVerify entity instance. |
| `SessionCreate` | `(data map[string]any) SeqbenchMcpEntity` | Create a SessionCreate entity instance. |
| `SessionGet` | `(data map[string]any) SeqbenchMcpEntity` | Create a SessionGet entity instance. |
| `SessionRun` | `(data map[string]any) SeqbenchMcpEntity` | Create a SessionRun entity instance. |
| `SessionSet` | `(data map[string]any) SeqbenchMcpEntity` | Create a SessionSet entity instance. |
| `SirnaDesign` | `(data map[string]any) SeqbenchMcpEntity` | Create a SirnaDesign entity instance. |
| `SiteDirectedMutagenesi` | `(data map[string]any) SeqbenchMcpEntity` | Create a SiteDirectedMutagenesi entity instance. |
| `Translate` | `(data map[string]any) SeqbenchMcpEntity` | Create a Translate entity instance. |
| `VariantAnnotate` | `(data map[string]any) SeqbenchMcpEntity` | Create a VariantAnnotate entity instance. |
| `VariantComparator` | `(data map[string]any) SeqbenchMcpEntity` | Create a VariantComparator entity instance. |
| `VerifyAssembly` | `(data map[string]any) SeqbenchMcpEntity` | Create a VerifyAssembly entity instance. |
| `VerifyConstruct` | `(data map[string]any) SeqbenchMcpEntity` | Create a VerifyConstruct entity instance. |
| `VirtualGel` | `(data map[string]any) SeqbenchMcpEntity` | Create a VirtualGel entity instance. |
| `VolcanoPlotData` | `(data map[string]any) SeqbenchMcpEntity` | Create a VolcanoPlotData entity instance. |
| `WebSearch` | `(data map[string]any) SeqbenchMcpEntity` | Create a WebSearch entity instance. |

### Entity interface (SeqbenchMcpEntity)

All entities implement the `SeqbenchMcpEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` | the entity record (`map[string]any`) |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    alphafoldLookup, err := client.AlphafoldLookup(nil).Create(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // alphafoldLookup is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### AlphafoldLookup

| Field | Description |
| --- | --- |
| `"accession"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/alphafold_lookup`

#### AsoDesign

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"length"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"target"` |  |
| `"tool"` |  |
| `"wing"` |  |

Operations: Create.

API path: `/aso_design`

#### BaseEditingDesign

| Field | Description |
| --- | --- |
| `"editor"` |  |
| `"frame_start"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"target"` |  |
| `"target_position"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/base_editing_design`

#### Batch

| Field | Description |
| --- | --- |
| `"arg"` |  |
| `"input"` |  |
| `"ok"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create, Load.

API path: `/batch`

#### BatchWorkflow

| Field | Description |
| --- | --- |
| `"input"` |  |
| `"ok"` |  |
| `"result"` |  |
| `"step"` |  |

Operations: Create, Load.

API path: `/workflow`

#### CharacterizeSequence

| Field | Description |
| --- | --- |
| `"end_primer_length"` |  |
| `"gate"` |  |
| `"max_orf"` |  |
| `"min_orf_aa"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/characterize_sequence`

#### CloningSimulate

| Field | Description |
| --- | --- |
| `"arm_tm_target"` |  |
| `"circular"` |  |
| `"enzyme"` |  |
| `"enzyme3"` |  |
| `"enzyme5"` |  |
| `"fragment"` |  |
| `"gate"` |  |
| `"insert"` |  |
| `"method"` |  |
| `"name"` |  |
| `"ok"` |  |
| `"overlap_len"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |
| `"vector"` |  |

Operations: Create.

API path: `/cloning_simulate`

#### CodonAdaptationIndex

| Field | Description |
| --- | --- |
| `"frame_start"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"organism"` |  |
| `"provenance"` |  |
| `"rare_threshold"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/codon_adaptation_index`

#### CodonOptimize

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"organism"` |  |
| `"protein"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/codon_optimize`

#### ConstructAutofix

| Field | Description |
| --- | --- |
| `"avoid_enzyme"` |  |
| `"cryptic_orf_min_aa"` |  |
| `"frame_start"` |  |
| `"gate"` |  |
| `"gc_high"` |  |
| `"gc_low"` |  |
| `"gc_window"` |  |
| `"homopolymer_min"` |  |
| `"max_pass"` |  |
| `"ok"` |  |
| `"organism"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/construct_autofix`

#### ConstructQc

| Field | Description |
| --- | --- |
| `"avoid_enzyme"` |  |
| `"cryptic_orf_min_aa"` |  |
| `"frame_start"` |  |
| `"gate"` |  |
| `"gc_high"` |  |
| `"gc_low"` |  |
| `"gc_window"` |  |
| `"homopolymer_min"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/construct_qc`

#### CrisprGrnaDesign

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"min_score"` |  |
| `"nuclease"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"search_reverse_strand"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/crispr_grna_design`

#### CrisprHdrDonor

| Field | Description |
| --- | --- |
| `"arm_length"` |  |
| `"block_pam"` |  |
| `"design_genotyping_primer"` |  |
| `"edit_end"` |  |
| `"edit_start"` |  |
| `"frame_start"` |  |
| `"gate"` |  |
| `"guide_end"` |  |
| `"guide_start"` |  |
| `"guide_strand"` |  |
| `"nuclease"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"replacement"` |  |
| `"result"` |  |
| `"target_sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/crispr_hdr_donor`

#### CrisprOfftargetCheck

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"max_mismatch"` |  |
| `"nuclease"` |  |
| `"ok"` |  |
| `"protospacer"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/crispr_offtarget_check`

#### CrossDimer

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence_a"` |  |
| `"sequence_b"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/cross_dimer`

#### DnaMolarity

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"length"` |  |
| `"mass_ng"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |
| `"type"` |  |
| `"volume_ul"` |  |

Operations: Create.

API path: `/dna_molarity`

#### DoubleDigest

| Field | Description |
| --- | --- |
| `"enzyme_a"` |  |
| `"enzyme_b"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/double_digest`

#### ExportEchoPicklist

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"reaction"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/export_echo_picklist`

#### ExportOpentronsProtocol

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"protocol_name"` |  |
| `"provenance"` |  |
| `"reaction"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/export_opentrons_protocol`

#### ExportPlateLayout

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"reaction"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/export_plate_layout`

#### ExpressionHeatmapCluster

| Field | Description |
| --- | --- |
| `"cluster_col"` |  |
| `"cluster_row"` |  |
| `"distance_metric"` |  |
| `"gate"` |  |
| `"gene"` |  |
| `"linkage"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sample"` |  |
| `"tool"` |  |
| `"value"` |  |
| `"z_score_row"` |  |

Operations: Create.

API path: `/expression_heatmap_cluster`

#### FastqQcReport

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"input"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"quality_offset"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/fastq_qc_report`

#### FastqTrim

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"input"` |  |
| `"min_length"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"quality_offset"` |  |
| `"quality_threshold"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/fastq_trim`

#### FindOrf

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"min_aa_length"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"require_stop"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/find_orfs`

#### FormatSequence

| Field | Description |
| --- | --- |
| `"case_mode"` |  |
| `"convert"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"reverse"` |  |
| `"sequence"` |  |
| `"strip_non_letter"` |  |
| `"tool"` |  |
| `"width"` |  |

Operations: Create.

API path: `/format_sequence`

#### FunctionalEnrichment

| Field | Description |
| --- | --- |
| `"background"` |  |
| `"collection"` |  |
| `"gate"` |  |
| `"gene"` |  |
| `"max_term_size"` |  |
| `"min_term_size"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/functional_enrichment`

#### GcContent

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/gc_content`

#### GeneDossier

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"gene"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/gene_dossier`

#### GeneExpression

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"gene"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/gene_expression`

#### GeneModel

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"gene"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/gene_model`

#### GoldenGateFidelity

| Field | Description |
| --- | --- |
| `"compare_to_named_set"` |  |
| `"dataset"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"overhang"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"risk_threshold"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/golden_gate_fidelity`

#### HgvsConvert

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |
| `"variant"` |  |

Operations: Create.

API path: `/hgvs_convert`

#### IdMapPoll

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"job_id"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/id_map_poll`

#### IdMapSubmit

| Field | Description |
| --- | --- |
| `"from"` |  |
| `"gate"` |  |
| `"ids"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tax_id"` |  |
| `"to"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/id_map_submit`

#### InSilicoPcr

| Field | Description |
| --- | --- |
| `"circular"` |  |
| `"forward_primer"` |  |
| `"gate"` |  |
| `"max_mismatch"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"reverse_primer"` |  |
| `"template"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/in_silico_pcr`

#### KaspPrimerDesign

| Field | Description |
| --- | --- |
| `"add_secondary_mismatch"` |  |
| `"allele_a"` |  |
| `"allele_b"` |  |
| `"gate"` |  |
| `"max_amplicon"` |  |
| `"min_amplicon"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"snp_position"` |  |
| `"target"` |  |
| `"target_core_tm"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/kasp_primer_design`

#### ListTool

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/`

#### MeltingTemperature

| Field | Description |
| --- | --- |
| `"dntp_mm"` |  |
| `"gate"` |  |
| `"mg_mm"` |  |
| `"na_mm"` |  |
| `"ok"` |  |
| `"oligo_nm"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"target_tm"` |  |
| `"tm_tolerance"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/melting_temperature`

#### MotifFinder

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"max_mismatch"` |  |
| `"motif"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"search_reverse_strand"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/motif_finder`

#### MultipleSequenceAlignment

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"input"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/multiple_sequence_alignment`

#### OligoAnalysi

| Field | Description |
| --- | --- |
| `"dntp_mm"` |  |
| `"gate"` |  |
| `"mg_mm"` |  |
| `"na_mm"` |  |
| `"ok"` |  |
| `"oligo_nm"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/oligo_analysis`

#### OrthologMap

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"source_species"` |  |
| `"symbol"` |  |
| `"target_species"` |  |
| `"tool"` |  |
| `"type"` |  |

Operations: Create.

API path: `/ortholog_map`

#### PairwiseAlignment

| Field | Description |
| --- | --- |
| `"gap"` |  |
| `"gate"` |  |
| `"match"` |  |
| `"mismatch"` |  |
| `"mode"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"seq_a"` |  |
| `"seq_b"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/pairwise_alignment`

#### ParseGenbank

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"text"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/parse_genbank`

#### ParseSangerTrace

| Field | Description |
| --- | --- |
| `"file_base64"` |  |
| `"file_name"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/parse_sanger_trace`

#### PlasmidAnnotate

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/plasmid_annotate`

#### PlasmidDeepAnnotate

| Field | Description |
| --- | --- |
| `"circular"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/plasmid_deep_annotate`

#### PlasmidFullReport

| Field | Description |
| --- | --- |
| `"circular"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |
| `"top_n"` |  |

Operations: Create.

API path: `/plasmid_full_report`

#### PlasmidIdentify

| Field | Description |
| --- | --- |
| `"circular"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |
| `"top_n"` |  |

Operations: Create.

API path: `/plasmid_identify`

#### PrimeEditingDesign

| Field | Description |
| --- | --- |
| `"edit_end"` |  |
| `"edit_start"` |  |
| `"frame_start"` |  |
| `"gate"` |  |
| `"inserted_seq"` |  |
| `"ok"` |  |
| `"pbs_length"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"rtt_homology"` |  |
| `"target"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/prime_editing_design`

#### PrimeEditingTwinDesign

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"new_sequence"` |  |
| `"ok"` |  |
| `"overlap_length"` |  |
| `"pbs_length"` |  |
| `"provenance"` |  |
| `"replace_end"` |  |
| `"replace_start"` |  |
| `"result"` |  |
| `"target"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/prime_editing_twin_design`

#### PrimerDesign

| Field | Description |
| --- | --- |
| `"amplicon_max"` |  |
| `"amplicon_min"` |  |
| `"dntp_mm"` |  |
| `"gate"` |  |
| `"gc_max"` |  |
| `"gc_min"` |  |
| `"len_max"` |  |
| `"len_min"` |  |
| `"len_opt"` |  |
| `"max_return"` |  |
| `"mg_mm"` |  |
| `"na_mm"` |  |
| `"ok"` |  |
| `"oligo_nm"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"target_end"` |  |
| `"target_start"` |  |
| `"template"` |  |
| `"tm_max"` |  |
| `"tm_max_diff"` |  |
| `"tm_min"` |  |
| `"tm_opt"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/primer_design`

#### PrimerSpecificity

| Field | Description |
| --- | --- |
| `"forward_primer"` |  |
| `"gate"` |  |
| `"max_mismatch"` |  |
| `"max_product_length"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"reverse_primer"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/primer_specificity`

#### ProteaseDigestion

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"max_mass"` |  |
| `"max_peptide"` |  |
| `"min_mass"` |  |
| `"missed_cleavage"` |  |
| `"ok"` |  |
| `"protease"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/protease_digestion`

#### ProteinAnnotatePoll

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"job_id"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/protein_annotate_poll`

#### ProteinAnnotateSubmit

| Field | Description |
| --- | --- |
| `"appl"` |  |
| `"gate"` |  |
| `"goterm"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/protein_annotate_submit`

#### ProteinHydrophobicity

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"scale"` |  |
| `"sequence"` |  |
| `"tool"` |  |
| `"window"` |  |

Operations: Create.

API path: `/protein_hydrophobicity`

#### ProteinProperty

| Field | Description |
| --- | --- |
| `"charge_step"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/protein_properties`

#### RandomSequence

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"gc_content"` |  |
| `"kind"` |  |
| `"length"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/random_sequence`

#### RestrictionSite

| Field | Description |
| --- | --- |
| `"enzyme"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/restriction_sites`

#### ReverseComplement

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |
| `"type"` |  |

Operations: Create.

API path: `/reverse_complement`

#### ReverseTranslate

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"mode"` |  |
| `"ok"` |  |
| `"organism"` |  |
| `"protein"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/reverse_translate`

#### RnaFold

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/rna_fold`

#### SangerVsReference

| Field | Description |
| --- | --- |
| `"file_base64"` |  |
| `"file_name"` |  |
| `"gate"` |  |
| `"min_coverage"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"read"` |  |
| `"reference"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/sanger_vs_reference`

#### SavePermalink

| Field | Description |
| --- | --- |
| `"arg"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/save_permalink`

#### SeqfileStat

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"input"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"quality_offset"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/seqfile_stats`

#### SequenceFetch

| Field | Description |
| --- | --- |
| `"accession"` |  |
| `"db"` |  |
| `"format"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/sequence_fetch`

#### SequenceFormatConvert

| Field | Description |
| --- | --- |
| `"from"` |  |
| `"gate"` |  |
| `"input"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"to"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/sequence_format_convert`

#### SequenceReport

| Field | Description |
| --- | --- |
| `"end_primer_length"` |  |
| `"gate"` |  |
| `"max_orf"` |  |
| `"min_orf_aa"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/sequence_report`

#### SequenceSearch

| Field | Description |
| --- | --- |
| `"db"` |  |
| `"gate"` |  |
| `"gene"` |  |
| `"max_result"` |  |
| `"ok"` |  |
| `"organism"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"term"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/sequence_search`

#### SequencingReadbackVerify

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"min_supporting_read"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"read"` |  |
| `"reference"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/sequencing_readback_verify`

#### SessionCreate

| Field | Description |
| --- | --- |
| `"entry"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/session_create`

#### SessionGet

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"name"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"session_id"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/session_get`

#### SessionRun

| Field | Description |
| --- | --- |
| `"arg"` |  |
| `"from_session"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"session_id"` |  |
| `"tool"` |  |
| `"write_back"` |  |

Operations: Create.

API path: `/session_run`

#### SessionSet

| Field | Description |
| --- | --- |
| `"entry"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"session_id"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/session_set`

#### SirnaDesign

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"min_reynold"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sh_rna_loop"` |  |
| `"target"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/sirna_design`

#### SiteDirectedMutagenesi

| Field | Description |
| --- | --- |
| `"arm_tm_target"` |  |
| `"dntp_mm"` |  |
| `"edit_kind"` |  |
| `"frame_start"` |  |
| `"gate"` |  |
| `"mg_mm"` |  |
| `"na_mm"` |  |
| `"new_base"` |  |
| `"ok"` |  |
| `"oligo_nm"` |  |
| `"organism"` |  |
| `"position"` |  |
| `"provenance"` |  |
| `"residue"` |  |
| `"result"` |  |
| `"style"` |  |
| `"target_aa"` |  |
| `"template"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/site_directed_mutagenesis`

#### Translate

| Field | Description |
| --- | --- |
| `"frame"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"to_stop"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/translate`

#### VariantAnnotate

| Field | Description |
| --- | --- |
| `"assembly"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |
| `"variant"` |  |

Operations: Create.

API path: `/variant_annotate`

#### VariantComparator

| Field | Description |
| --- | --- |
| `"coding"` |  |
| `"frame_start"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"query"` |  |
| `"reference"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/variant_comparator`

#### VerifyAssembly

| Field | Description |
| --- | --- |
| `"arm_tm_target"` |  |
| `"circular"` |  |
| `"claimed_construct"` |  |
| `"coding"` |  |
| `"enzyme"` |  |
| `"enzyme3"` |  |
| `"enzyme5"` |  |
| `"fragment"` |  |
| `"fragment_pcr"` |  |
| `"frame_start"` |  |
| `"gate"` |  |
| `"insert"` |  |
| `"insert_pcr"` |  |
| `"method"` |  |
| `"name"` |  |
| `"ok"` |  |
| `"overlap_len"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |
| `"vector"` |  |
| `"vector_pcr"` |  |

Operations: Create.

API path: `/verify_assembly`

#### VerifyConstruct

| Field | Description |
| --- | --- |
| `"claimed_construct"` |  |
| `"expected_frame_start"` |  |
| `"gate"` |  |
| `"insert_forward_primer"` |  |
| `"insert_reverse_primer"` |  |
| `"insert_template"` |  |
| `"max_primer_mismatch"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"template_circular"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/verify_construct`

#### VirtualGel

| Field | Description |
| --- | --- |
| `"circular"` |  |
| `"enzyme"` |  |
| `"gate"` |  |
| `"ladder"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/virtual_gel`

#### VolcanoPlotData

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"row"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/volcano_plot_data`

#### WebSearch

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"max_result"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"query"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/web_search`



## Entities


### AlphafoldLookup

Create an instance: `alphafoldLookup := client.AlphafoldLookup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.AlphafoldLookup(nil).Create(map[string]any{
    "accession": "example_accession",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### AsoDesign

Create an instance: `asoDesign := client.AsoDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `length` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `target` | `string` |  |
| `tool` | `string` |  |
| `wing` | `int` |  |

#### Example: Create

```go
result, err := client.AsoDesign(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### BaseEditingDesign

Create an instance: `baseEditingDesign := client.BaseEditingDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `editor` | `string` |  |
| `frame_start` | `int` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `target` | `string` |  |
| `target_position` | `int` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.BaseEditingDesign(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Batch

Create an instance: `batch := client.Batch(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `map[string]any` |  |
| `input` | `string` |  |
| `ok` | `any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Load

```go
batch, err := client.Batch(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(batch) // the loaded record
```

#### Example: Create

```go
result, err := client.Batch(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### BatchWorkflow

Create an instance: `batchWorkflow := client.BatchWorkflow(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `input` | `string` |  |
| `ok` | `any` |  |
| `result` | `map[string]any` |  |
| `step` | `[]any` |  |

#### Example: Load

```go
batchWorkflow, err := client.BatchWorkflow(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(batchWorkflow) // the loaded record
```

#### Example: Create

```go
result, err := client.BatchWorkflow(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "result": map[string]any{},
    "step": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### CharacterizeSequence

Create an instance: `characterizeSequence := client.CharacterizeSequence(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `end_primer_length` | `int` |  |
| `gate` | `any` |  |
| `max_orf` | `int` |  |
| `min_orf_aa` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.CharacterizeSequence(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### CloningSimulate

Create an instance: `cloningSimulate := client.CloningSimulate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_tm_target` | `float64` |  |
| `circular` | `bool` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragment` | `[]any` |  |
| `gate` | `any` |  |
| `insert` | `string` |  |
| `method` | `string` |  |
| `name` | `[]any` |  |
| `ok` | `any` |  |
| `overlap_len` | `int` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |
| `vector` | `string` |  |

#### Example: Create

```go
result, err := client.CloningSimulate(nil).Create(map[string]any{
    "method": "example_method",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### CodonAdaptationIndex

Create an instance: `codonAdaptationIndex := client.CodonAdaptationIndex(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `frame_start` | `int` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `provenance` | `map[string]any` |  |
| `rare_threshold` | `float64` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.CodonAdaptationIndex(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### CodonOptimize

Create an instance: `codonOptimize := client.CodonOptimize(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `protein` | `string` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.CodonOptimize(nil).Create(map[string]any{
    "ok": "example_ok",
    "protein": "example_protein",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ConstructAutofix

Create an instance: `constructAutofix := client.ConstructAutofix(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoid_enzyme` | `[]any` |  |
| `cryptic_orf_min_aa` | `int` |  |
| `frame_start` | `int` |  |
| `gate` | `any` |  |
| `gc_high` | `float64` |  |
| `gc_low` | `float64` |  |
| `gc_window` | `int` |  |
| `homopolymer_min` | `int` |  |
| `max_pass` | `int` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ConstructAutofix(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ConstructQc

Create an instance: `constructQc := client.ConstructQc(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avoid_enzyme` | `[]any` |  |
| `cryptic_orf_min_aa` | `int` |  |
| `frame_start` | `int` |  |
| `gate` | `any` |  |
| `gc_high` | `float64` |  |
| `gc_low` | `float64` |  |
| `gc_window` | `int` |  |
| `homopolymer_min` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ConstructQc(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### CrisprGrnaDesign

Create an instance: `crisprGrnaDesign := client.CrisprGrnaDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `min_score` | `float64` |  |
| `nuclease` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `search_reverse_strand` | `bool` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.CrisprGrnaDesign(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### CrisprHdrDonor

Create an instance: `crisprHdrDonor := client.CrisprHdrDonor(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_length` | `int` |  |
| `block_pam` | `bool` |  |
| `design_genotyping_primer` | `bool` |  |
| `edit_end` | `int` |  |
| `edit_start` | `int` |  |
| `frame_start` | `int` |  |
| `gate` | `any` |  |
| `guide_end` | `int` |  |
| `guide_start` | `int` |  |
| `guide_strand` | `string` |  |
| `nuclease` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `replacement` | `string` |  |
| `result` | `map[string]any` |  |
| `target_sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.CrisprHdrDonor(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "replacement": "example_replacement",
    "result": map[string]any{},
    "target_sequence": "example_target_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### CrisprOfftargetCheck

Create an instance: `crisprOfftargetCheck := client.CrisprOfftargetCheck(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `max_mismatch` | `int` |  |
| `nuclease` | `string` |  |
| `ok` | `any` |  |
| `protospacer` | `string` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.CrisprOfftargetCheck(nil).Create(map[string]any{
    "ok": "example_ok",
    "protospacer": "example_protospacer",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### CrossDimer

Create an instance: `crossDimer := client.CrossDimer(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence_a` | `string` |  |
| `sequence_b` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.CrossDimer(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence_a": "example_sequence_a",
    "sequence_b": "example_sequence_b",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### DnaMolarity

Create an instance: `dnaMolarity := client.DnaMolarity(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `length` | `int` |  |
| `mass_ng` | `float64` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |
| `volume_ul` | `float64` |  |

#### Example: Create

```go
result, err := client.DnaMolarity(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### DoubleDigest

Create an instance: `doubleDigest := client.DoubleDigest(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enzyme_a` | `string` |  |
| `enzyme_b` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.DoubleDigest(nil).Create(map[string]any{
    "enzyme_a": "example_enzyme_a",
    "enzyme_b": "example_enzyme_b",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ExportEchoPicklist

Create an instance: `exportEchoPicklist := client.ExportEchoPicklist(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `reaction` | `[]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ExportEchoPicklist(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reaction": []any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ExportOpentronsProtocol

Create an instance: `exportOpentronsProtocol := client.ExportOpentronsProtocol(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `protocol_name` | `string` |  |
| `provenance` | `map[string]any` |  |
| `reaction` | `[]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ExportOpentronsProtocol(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reaction": []any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ExportPlateLayout

Create an instance: `exportPlateLayout := client.ExportPlateLayout(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `reaction` | `[]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ExportPlateLayout(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reaction": []any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ExpressionHeatmapCluster

Create an instance: `expressionHeatmapCluster := client.ExpressionHeatmapCluster(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cluster_col` | `bool` |  |
| `cluster_row` | `bool` |  |
| `distance_metric` | `string` |  |
| `gate` | `any` |  |
| `gene` | `[]any` |  |
| `linkage` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sample` | `[]any` |  |
| `tool` | `string` |  |
| `value` | `[]any` |  |
| `z_score_row` | `bool` |  |

#### Example: Create

```go
result, err := client.ExpressionHeatmapCluster(nil).Create(map[string]any{
    "gene": []any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sample": []any{},
    "tool": "example_tool",
    "value": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### FastqQcReport

Create an instance: `fastqQcReport := client.FastqQcReport(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `input` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `quality_offset` | `int` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.FastqQcReport(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### FastqTrim

Create an instance: `fastqTrim := client.FastqTrim(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `input` | `string` |  |
| `min_length` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `quality_offset` | `int` |  |
| `quality_threshold` | `int` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.FastqTrim(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### FindOrf

Create an instance: `findOrf := client.FindOrf(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `min_aa_length` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `require_stop` | `bool` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.FindOrf(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### FormatSequence

Create an instance: `formatSequence := client.FormatSequence(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `case_mode` | `string` |  |
| `convert` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `reverse` | `bool` |  |
| `sequence` | `string` |  |
| `strip_non_letter` | `bool` |  |
| `tool` | `string` |  |
| `width` | `int` |  |

#### Example: Create

```go
result, err := client.FormatSequence(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### FunctionalEnrichment

Create an instance: `functionalEnrichment := client.FunctionalEnrichment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `background` | `[]any` |  |
| `collection` | `[]any` |  |
| `gate` | `any` |  |
| `gene` | `[]any` |  |
| `max_term_size` | `int` |  |
| `min_term_size` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.FunctionalEnrichment(nil).Create(map[string]any{
    "gene": []any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### GcContent

Create an instance: `gcContent := client.GcContent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.GcContent(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### GeneDossier

Create an instance: `geneDossier := client.GeneDossier(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `gene` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.GeneDossier(nil).Create(map[string]any{
    "gene": "example_gene",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### GeneExpression

Create an instance: `geneExpression := client.GeneExpression(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `gene` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.GeneExpression(nil).Create(map[string]any{
    "gene": "example_gene",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### GeneModel

Create an instance: `geneModel := client.GeneModel(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `gene` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.GeneModel(nil).Create(map[string]any{
    "gene": "example_gene",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### GoldenGateFidelity

Create an instance: `goldenGateFidelity := client.GoldenGateFidelity(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `compare_to_named_set` | `string` |  |
| `dataset` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `overhang` | `[]any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `risk_threshold` | `float64` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.GoldenGateFidelity(nil).Create(map[string]any{
    "ok": "example_ok",
    "overhang": []any{},
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### HgvsConvert

Create an instance: `hgvsConvert := client.HgvsConvert(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |
| `variant` | `string` |  |

#### Example: Create

```go
result, err := client.HgvsConvert(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
    "variant": "example_variant",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### IdMapPoll

Create an instance: `idMapPoll := client.IdMapPoll(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `job_id` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.IdMapPoll(nil).Create(map[string]any{
    "job_id": "example_job_id",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### IdMapSubmit

Create an instance: `idMapSubmit := client.IdMapSubmit(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from` | `string` |  |
| `gate` | `any` |  |
| `ids` | `[]any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tax_id` | `string` |  |
| `to` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.IdMapSubmit(nil).Create(map[string]any{
    "from": "example_from",
    "ids": []any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "to": "example_to",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### InSilicoPcr

Create an instance: `inSilicoPcr := client.InSilicoPcr(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `forward_primer` | `string` |  |
| `gate` | `any` |  |
| `max_mismatch` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `reverse_primer` | `string` |  |
| `template` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.InSilicoPcr(nil).Create(map[string]any{
    "forward_primer": "example_forward_primer",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "reverse_primer": "example_reverse_primer",
    "template": "example_template",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### KaspPrimerDesign

Create an instance: `kaspPrimerDesign := client.KaspPrimerDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `add_secondary_mismatch` | `bool` |  |
| `allele_a` | `string` |  |
| `allele_b` | `string` |  |
| `gate` | `any` |  |
| `max_amplicon` | `int` |  |
| `min_amplicon` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `snp_position` | `int` |  |
| `target` | `string` |  |
| `target_core_tm` | `float64` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.KaspPrimerDesign(nil).Create(map[string]any{
    "allele_a": "example_allele_a",
    "allele_b": "example_allele_b",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "snp_position": 1,
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ListTool

Create an instance: `listTool := client.ListTool(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Example: Load

```go
listTool, err := client.ListTool(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(listTool) // the loaded record
```


### MeltingTemperature

Create an instance: `meltingTemperature := client.MeltingTemperature(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dntp_mm` | `float64` |  |
| `gate` | `any` |  |
| `mg_mm` | `float64` |  |
| `na_mm` | `float64` |  |
| `ok` | `any` |  |
| `oligo_nm` | `float64` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `target_tm` | `float64` |  |
| `tm_tolerance` | `float64` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.MeltingTemperature(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### MotifFinder

Create an instance: `motifFinder := client.MotifFinder(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `max_mismatch` | `int` |  |
| `motif` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `search_reverse_strand` | `bool` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.MotifFinder(nil).Create(map[string]any{
    "motif": "example_motif",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### MultipleSequenceAlignment

Create an instance: `multipleSequenceAlignment := client.MultipleSequenceAlignment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `input` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.MultipleSequenceAlignment(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### OligoAnalysi

Create an instance: `oligoAnalysi := client.OligoAnalysi(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dntp_mm` | `float64` |  |
| `gate` | `any` |  |
| `mg_mm` | `float64` |  |
| `na_mm` | `float64` |  |
| `ok` | `any` |  |
| `oligo_nm` | `float64` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.OligoAnalysi(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### OrthologMap

Create an instance: `orthologMap := client.OrthologMap(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `source_species` | `string` |  |
| `symbol` | `[]any` |  |
| `target_species` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |

#### Example: Create

```go
result, err := client.OrthologMap(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "symbol": []any{},
    "target_species": "example_target_species",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PairwiseAlignment

Create an instance: `pairwiseAlignment := client.PairwiseAlignment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gap` | `float64` |  |
| `gate` | `any` |  |
| `match` | `float64` |  |
| `mismatch` | `float64` |  |
| `mode` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `seq_a` | `string` |  |
| `seq_b` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.PairwiseAlignment(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "seq_a": "example_seq_a",
    "seq_b": "example_seq_b",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ParseGenbank

Create an instance: `parseGenbank := client.ParseGenbank(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `text` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ParseGenbank(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "text": "example_text",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ParseSangerTrace

Create an instance: `parseSangerTrace := client.ParseSangerTrace(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file_base64` | `string` |  |
| `file_name` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ParseSangerTrace(nil).Create(map[string]any{
    "file_base64": "example_file_base64",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PlasmidAnnotate

Create an instance: `plasmidAnnotate := client.PlasmidAnnotate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.PlasmidAnnotate(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PlasmidDeepAnnotate

Create an instance: `plasmidDeepAnnotate := client.PlasmidDeepAnnotate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.PlasmidDeepAnnotate(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PlasmidFullReport

Create an instance: `plasmidFullReport := client.PlasmidFullReport(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `top_n` | `int` |  |

#### Example: Create

```go
result, err := client.PlasmidFullReport(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PlasmidIdentify

Create an instance: `plasmidIdentify := client.PlasmidIdentify(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `top_n` | `int` |  |

#### Example: Create

```go
result, err := client.PlasmidIdentify(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PrimeEditingDesign

Create an instance: `primeEditingDesign := client.PrimeEditingDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `edit_end` | `int` |  |
| `edit_start` | `int` |  |
| `frame_start` | `int` |  |
| `gate` | `any` |  |
| `inserted_seq` | `string` |  |
| `ok` | `any` |  |
| `pbs_length` | `int` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `rtt_homology` | `int` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.PrimeEditingDesign(nil).Create(map[string]any{
    "edit_end": 1,
    "edit_start": 1,
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PrimeEditingTwinDesign

Create an instance: `primeEditingTwinDesign := client.PrimeEditingTwinDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `new_sequence` | `string` |  |
| `ok` | `any` |  |
| `overlap_length` | `int` |  |
| `pbs_length` | `int` |  |
| `provenance` | `map[string]any` |  |
| `replace_end` | `int` |  |
| `replace_start` | `int` |  |
| `result` | `map[string]any` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.PrimeEditingTwinDesign(nil).Create(map[string]any{
    "new_sequence": "example_new_sequence",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "replace_end": 1,
    "replace_start": 1,
    "result": map[string]any{},
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PrimerDesign

Create an instance: `primerDesign := client.PrimerDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amplicon_max` | `int` |  |
| `amplicon_min` | `int` |  |
| `dntp_mm` | `float64` |  |
| `gate` | `any` |  |
| `gc_max` | `float64` |  |
| `gc_min` | `float64` |  |
| `len_max` | `int` |  |
| `len_min` | `int` |  |
| `len_opt` | `int` |  |
| `max_return` | `int` |  |
| `mg_mm` | `float64` |  |
| `na_mm` | `float64` |  |
| `ok` | `any` |  |
| `oligo_nm` | `float64` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `target_end` | `int` |  |
| `target_start` | `int` |  |
| `template` | `string` |  |
| `tm_max` | `float64` |  |
| `tm_max_diff` | `float64` |  |
| `tm_min` | `float64` |  |
| `tm_opt` | `float64` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.PrimerDesign(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "template": "example_template",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PrimerSpecificity

Create an instance: `primerSpecificity := client.PrimerSpecificity(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `forward_primer` | `string` |  |
| `gate` | `any` |  |
| `max_mismatch` | `int` |  |
| `max_product_length` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `reverse_primer` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.PrimerSpecificity(nil).Create(map[string]any{
    "forward_primer": "example_forward_primer",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "reverse_primer": "example_reverse_primer",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ProteaseDigestion

Create an instance: `proteaseDigestion := client.ProteaseDigestion(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `max_mass` | `float64` |  |
| `max_peptide` | `int` |  |
| `min_mass` | `float64` |  |
| `missed_cleavage` | `int` |  |
| `ok` | `any` |  |
| `protease` | `string` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ProteaseDigestion(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ProteinAnnotatePoll

Create an instance: `proteinAnnotatePoll := client.ProteinAnnotatePoll(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `job_id` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ProteinAnnotatePoll(nil).Create(map[string]any{
    "job_id": "example_job_id",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ProteinAnnotateSubmit

Create an instance: `proteinAnnotateSubmit := client.ProteinAnnotateSubmit(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `appl` | `string` |  |
| `gate` | `any` |  |
| `goterm` | `bool` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ProteinAnnotateSubmit(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ProteinHydrophobicity

Create an instance: `proteinHydrophobicity := client.ProteinHydrophobicity(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `scale` | `string` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `window` | `int` |  |

#### Example: Create

```go
result, err := client.ProteinHydrophobicity(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ProteinProperty

Create an instance: `proteinProperty := client.ProteinProperty(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `charge_step` | `float64` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ProteinProperty(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### RandomSequence

Create an instance: `randomSequence := client.RandomSequence(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `gc_content` | `float64` |  |
| `kind` | `string` |  |
| `length` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.RandomSequence(nil).Create(map[string]any{
    "length": 1,
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### RestrictionSite

Create an instance: `restrictionSite := client.RestrictionSite(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `enzyme` | `[]any` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.RestrictionSite(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ReverseComplement

Create an instance: `reverseComplement := client.ReverseComplement(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |

#### Example: Create

```go
result, err := client.ReverseComplement(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ReverseTranslate

Create an instance: `reverseTranslate := client.ReverseTranslate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `mode` | `string` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `protein` | `string` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ReverseTranslate(nil).Create(map[string]any{
    "ok": "example_ok",
    "protein": "example_protein",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### RnaFold

Create an instance: `rnaFold := client.RnaFold(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.RnaFold(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SangerVsReference

Create an instance: `sangerVsReference := client.SangerVsReference(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `file_base64` | `string` |  |
| `file_name` | `string` |  |
| `gate` | `any` |  |
| `min_coverage` | `float64` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `read` | `string` |  |
| `reference` | `string` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SangerVsReference(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reference": "example_reference",
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SavePermalink

Create an instance: `savePermalink := client.SavePermalink(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `map[string]any` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SavePermalink(nil).Create(map[string]any{
    "arg": map[string]any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SeqfileStat

Create an instance: `seqfileStat := client.SeqfileStat(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `input` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `quality_offset` | `int` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SeqfileStat(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SequenceFetch

Create an instance: `sequenceFetch := client.SequenceFetch(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accession` | `string` |  |
| `db` | `string` |  |
| `format` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SequenceFetch(nil).Create(map[string]any{
    "accession": "example_accession",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SequenceFormatConvert

Create an instance: `sequenceFormatConvert := client.SequenceFormatConvert(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `from` | `string` |  |
| `gate` | `any` |  |
| `input` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `to` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SequenceFormatConvert(nil).Create(map[string]any{
    "input": "example_input",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SequenceReport

Create an instance: `sequenceReport := client.SequenceReport(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `end_primer_length` | `int` |  |
| `gate` | `any` |  |
| `max_orf` | `int` |  |
| `min_orf_aa` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SequenceReport(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SequenceSearch

Create an instance: `sequenceSearch := client.SequenceSearch(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `db` | `string` |  |
| `gate` | `any` |  |
| `gene` | `string` |  |
| `max_result` | `int` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `term` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SequenceSearch(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SequencingReadbackVerify

Create an instance: `sequencingReadbackVerify := client.SequencingReadbackVerify(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `min_supporting_read` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `read` | `string` |  |
| `reference` | `string` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SequencingReadbackVerify(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "read": "example_read",
    "reference": "example_reference",
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SessionCreate

Create an instance: `sessionCreate := client.SessionCreate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entry` | `map[string]any` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SessionCreate(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SessionGet

Create an instance: `sessionGet := client.SessionGet(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `name` | `[]any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SessionGet(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "session_id": "example_session_id",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SessionRun

Create an instance: `sessionRun := client.SessionRun(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arg` | `map[string]any` |  |
| `from_session` | `map[string]any` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |
| `write_back` | `map[string]any` |  |

#### Example: Create

```go
result, err := client.SessionRun(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "session_id": "example_session_id",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SessionSet

Create an instance: `sessionSet := client.SessionSet(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `entry` | `map[string]any` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `session_id` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SessionSet(nil).Create(map[string]any{
    "entry": map[string]any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "session_id": "example_session_id",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SirnaDesign

Create an instance: `sirnaDesign := client.SirnaDesign(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `min_reynold` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sh_rna_loop` | `string` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SirnaDesign(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "target": "example_target",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SiteDirectedMutagenesi

Create an instance: `siteDirectedMutagenesi := client.SiteDirectedMutagenesi(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_tm_target` | `float64` |  |
| `dntp_mm` | `float64` |  |
| `edit_kind` | `string` |  |
| `frame_start` | `int` |  |
| `gate` | `any` |  |
| `mg_mm` | `float64` |  |
| `na_mm` | `float64` |  |
| `new_base` | `string` |  |
| `ok` | `any` |  |
| `oligo_nm` | `float64` |  |
| `organism` | `string` |  |
| `position` | `int` |  |
| `provenance` | `map[string]any` |  |
| `residue` | `int` |  |
| `result` | `map[string]any` |  |
| `style` | `string` |  |
| `target_aa` | `string` |  |
| `template` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SiteDirectedMutagenesi(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "template": "example_template",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Translate

Create an instance: `translate := client.Translate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `frame` | `int` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `to_stop` | `bool` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.Translate(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### VariantAnnotate

Create an instance: `variantAnnotate := client.VariantAnnotate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `assembly` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |
| `variant` | `string` |  |

#### Example: Create

```go
result, err := client.VariantAnnotate(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
    "variant": "example_variant",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### VariantComparator

Create an instance: `variantComparator := client.VariantComparator(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `coding` | `bool` |  |
| `frame_start` | `int` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `query` | `string` |  |
| `reference` | `string` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.VariantComparator(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "query": "example_query",
    "reference": "example_reference",
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### VerifyAssembly

Create an instance: `verifyAssembly := client.VerifyAssembly(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `arm_tm_target` | `float64` |  |
| `circular` | `bool` |  |
| `claimed_construct` | `string` |  |
| `coding` | `bool` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragment` | `[]any` |  |
| `fragment_pcr` | `[]any` |  |
| `frame_start` | `int` |  |
| `gate` | `any` |  |
| `insert` | `string` |  |
| `insert_pcr` | `map[string]any` |  |
| `method` | `string` |  |
| `name` | `[]any` |  |
| `ok` | `any` |  |
| `overlap_len` | `int` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |
| `vector` | `string` |  |
| `vector_pcr` | `map[string]any` |  |

#### Example: Create

```go
result, err := client.VerifyAssembly(nil).Create(map[string]any{
    "claimed_construct": "example_claimed_construct",
    "method": "example_method",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### VerifyConstruct

Create an instance: `verifyConstruct := client.VerifyConstruct(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `claimed_construct` | `string` |  |
| `expected_frame_start` | `int` |  |
| `gate` | `any` |  |
| `insert_forward_primer` | `string` |  |
| `insert_reverse_primer` | `string` |  |
| `insert_template` | `string` |  |
| `max_primer_mismatch` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `template_circular` | `bool` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.VerifyConstruct(nil).Create(map[string]any{
    "claimed_construct": "example_claimed_construct",
    "insert_forward_primer": "example_insert_forward_primer",
    "insert_reverse_primer": "example_insert_reverse_primer",
    "insert_template": "example_insert_template",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### VirtualGel

Create an instance: `virtualGel := client.VirtualGel(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `circular` | `bool` |  |
| `enzyme` | `[]any` |  |
| `gate` | `any` |  |
| `ladder` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.VirtualGel(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequence": "example_sequence",
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### VolcanoPlotData

Create an instance: `volcanoPlotData := client.VolcanoPlotData(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `row` | `[]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.VolcanoPlotData(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "row": []any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### WebSearch

Create an instance: `webSearch := client.WebSearch(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `gate` | `any` |  |
| `max_result` | `float64` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `query` | `string` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.WebSearch(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "query": "example_query",
    "result": map[string]any{},
    "tool": "example_tool",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/seqbench-mcp-sdk/go/
├── seqbench-mcp.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/seqbench-mcp-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
batch := client.Batch(nil)
batch.Load(nil, nil)

// batch.Data() now returns the batch data from the last load
// batch.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
