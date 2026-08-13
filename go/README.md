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
| `"frameStart"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"target"` |  |
| `"targetPosition"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/base_editing_design`

#### Batch

| Field | Description |
| --- | --- |
| `"args"` |  |
| `"capped"` |  |
| `"columns"` |  |
| `"count"` |  |
| `"errors"` |  |
| `"input"` |  |
| `"limit"` |  |
| `"provenance"` |  |
| `"rows"` |  |
| `"tool"` |  |

Operations: Create, Load.

API path: `/batch`

#### BatchWorkflow

| Field | Description |
| --- | --- |
| `"capped"` |  |
| `"columns"` |  |
| `"count"` |  |
| `"errors"` |  |
| `"input"` |  |
| `"limit"` |  |
| `"provenance"` |  |
| `"rows"` |  |
| `"steps"` |  |

Operations: Create, Load.

API path: `/workflow`

#### CharacterizeSequence

| Field | Description |
| --- | --- |
| `"endPrimerLength"` |  |
| `"gate"` |  |
| `"maxOrfs"` |  |
| `"minOrfAa"` |  |
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
| `"armTmTarget"` |  |
| `"circular"` |  |
| `"enzyme"` |  |
| `"enzyme3"` |  |
| `"enzyme5"` |  |
| `"fragments"` |  |
| `"gate"` |  |
| `"insert"` |  |
| `"method"` |  |
| `"names"` |  |
| `"ok"` |  |
| `"overlapLen"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |
| `"vector"` |  |

Operations: Create.

API path: `/cloning_simulate`

#### CodonAdaptationIndex

| Field | Description |
| --- | --- |
| `"frameStart"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"organism"` |  |
| `"provenance"` |  |
| `"rareThreshold"` |  |
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
| `"avoidEnzymes"` |  |
| `"crypticOrfMinAa"` |  |
| `"frameStart"` |  |
| `"gate"` |  |
| `"gcHigh"` |  |
| `"gcLow"` |  |
| `"gcWindow"` |  |
| `"homopolymerMin"` |  |
| `"maxPasses"` |  |
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
| `"avoidEnzymes"` |  |
| `"crypticOrfMinAa"` |  |
| `"frameStart"` |  |
| `"gate"` |  |
| `"gcHigh"` |  |
| `"gcLow"` |  |
| `"gcWindow"` |  |
| `"homopolymerMin"` |  |
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
| `"minScore"` |  |
| `"nuclease"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"searchReverseStrand"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/crispr_grna_design`

#### CrisprHdrDonor

| Field | Description |
| --- | --- |
| `"armLength"` |  |
| `"blockPam"` |  |
| `"designGenotypingPrimers"` |  |
| `"editEnd"` |  |
| `"editStart"` |  |
| `"frameStart"` |  |
| `"gate"` |  |
| `"guideEnd"` |  |
| `"guideStart"` |  |
| `"guideStrand"` |  |
| `"nuclease"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"replacement"` |  |
| `"result"` |  |
| `"targetSequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/crispr_hdr_donor`

#### CrisprOfftargetCheck

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"maxMismatches"` |  |
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
| `"sequenceA"` |  |
| `"sequenceB"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/cross_dimer`

#### DnaMolarity

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"length"` |  |
| `"massNg"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |
| `"type"` |  |
| `"volumeUl"` |  |

Operations: Create.

API path: `/dna_molarity`

#### DoubleDigest

| Field | Description |
| --- | --- |
| `"enzymeA"` |  |
| `"enzymeB"` |  |
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
| `"reactions"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/export_echo_picklist`

#### ExportOpentronsProtocol

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"ok"` |  |
| `"protocolName"` |  |
| `"provenance"` |  |
| `"reactions"` |  |
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
| `"reactions"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/export_plate_layout`

#### ExpressionHeatmapCluster

| Field | Description |
| --- | --- |
| `"clusterCols"` |  |
| `"clusterRows"` |  |
| `"distanceMetric"` |  |
| `"gate"` |  |
| `"genes"` |  |
| `"linkage"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"samples"` |  |
| `"tool"` |  |
| `"values"` |  |
| `"zScoreRows"` |  |

Operations: Create.

API path: `/expression_heatmap_cluster`

#### FastqQcReport

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"input"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"qualityOffset"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/fastq_qc_report`

#### FastqTrim

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"input"` |  |
| `"minLength"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"qualityOffset"` |  |
| `"qualityThreshold"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/fastq_trim`

#### FindOrf

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"minAaLength"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"requireStop"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/find_orfs`

#### FormatSequence

| Field | Description |
| --- | --- |
| `"caseMode"` |  |
| `"convert"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"reverse"` |  |
| `"sequence"` |  |
| `"stripNonLetters"` |  |
| `"tool"` |  |
| `"width"` |  |

Operations: Create.

API path: `/format_sequence`

#### FunctionalEnrichment

| Field | Description |
| --- | --- |
| `"background"` |  |
| `"collections"` |  |
| `"gate"` |  |
| `"genes"` |  |
| `"maxTermSize"` |  |
| `"minTermSize"` |  |
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
| `"compareToNamedSet"` |  |
| `"dataset"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"overhangs"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"riskThreshold"` |  |
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
| `"jobId"` |  |
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
| `"taxId"` |  |
| `"to"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/id_map_submit`

#### InSilicoPcr

| Field | Description |
| --- | --- |
| `"circular"` |  |
| `"forwardPrimer"` |  |
| `"gate"` |  |
| `"maxMismatches"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"reversePrimer"` |  |
| `"template"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/in_silico_pcr`

#### KaspPrimerDesign

| Field | Description |
| --- | --- |
| `"addSecondaryMismatch"` |  |
| `"alleleA"` |  |
| `"alleleB"` |  |
| `"gate"` |  |
| `"maxAmplicon"` |  |
| `"minAmplicon"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"snpPosition"` |  |
| `"target"` |  |
| `"targetCoreTm"` |  |
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
| `"dntpMM"` |  |
| `"gate"` |  |
| `"mgMM"` |  |
| `"naMM"` |  |
| `"ok"` |  |
| `"oligoNM"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sequence"` |  |
| `"targetTm"` |  |
| `"tmTolerance"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/melting_temperature`

#### MotifFinder

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"maxMismatches"` |  |
| `"motif"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"searchReverseStrand"` |  |
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
| `"dntpMM"` |  |
| `"gate"` |  |
| `"mgMM"` |  |
| `"naMM"` |  |
| `"ok"` |  |
| `"oligoNM"` |  |
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
| `"sourceSpecies"` |  |
| `"symbols"` |  |
| `"targetSpecies"` |  |
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
| `"seqA"` |  |
| `"seqB"` |  |
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
| `"fileBase64"` |  |
| `"fileName"` |  |
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
| `"topN"` |  |

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
| `"topN"` |  |

Operations: Create.

API path: `/plasmid_identify`

#### PrimeEditingDesign

| Field | Description |
| --- | --- |
| `"editEnd"` |  |
| `"editStart"` |  |
| `"frameStart"` |  |
| `"gate"` |  |
| `"insertedSeq"` |  |
| `"ok"` |  |
| `"pbsLength"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"rttHomology"` |  |
| `"target"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/prime_editing_design`

#### PrimeEditingTwinDesign

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"newSequence"` |  |
| `"ok"` |  |
| `"overlapLength"` |  |
| `"pbsLength"` |  |
| `"provenance"` |  |
| `"replaceEnd"` |  |
| `"replaceStart"` |  |
| `"result"` |  |
| `"target"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/prime_editing_twin_design`

#### PrimerDesign

| Field | Description |
| --- | --- |
| `"ampliconMax"` |  |
| `"ampliconMin"` |  |
| `"dntpMM"` |  |
| `"gate"` |  |
| `"gcMax"` |  |
| `"gcMin"` |  |
| `"lenMax"` |  |
| `"lenMin"` |  |
| `"lenOpt"` |  |
| `"maxReturn"` |  |
| `"mgMM"` |  |
| `"naMM"` |  |
| `"ok"` |  |
| `"oligoNM"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"targetEnd"` |  |
| `"targetStart"` |  |
| `"template"` |  |
| `"tmMax"` |  |
| `"tmMaxDiff"` |  |
| `"tmMin"` |  |
| `"tmOpt"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/primer_design`

#### PrimerSpecificity

| Field | Description |
| --- | --- |
| `"forwardPrimer"` |  |
| `"gate"` |  |
| `"maxMismatches"` |  |
| `"maxProductLength"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"reversePrimer"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/primer_specificity`

#### ProteaseDigestion

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"maxMass"` |  |
| `"maxPeptides"` |  |
| `"minMass"` |  |
| `"missedCleavages"` |  |
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
| `"jobId"` |  |
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
| `"goterms"` |  |
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
| `"chargeStep"` |  |
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
| `"gcContent"` |  |
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
| `"enzymes"` |  |
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
| `"fileBase64"` |  |
| `"fileName"` |  |
| `"gate"` |  |
| `"minCoverage"` |  |
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
| `"args"` |  |
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
| `"qualityOffset"` |  |
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
| `"endPrimerLength"` |  |
| `"gate"` |  |
| `"maxOrfs"` |  |
| `"minOrfAa"` |  |
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
| `"maxResults"` |  |
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
| `"minSupportingReads"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"reads"` |  |
| `"reference"` |  |
| `"result"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/sequencing_readback_verify`

#### SessionCreate

| Field | Description |
| --- | --- |
| `"entries"` |  |
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
| `"names"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sessionId"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/session_get`

#### SessionRun

| Field | Description |
| --- | --- |
| `"args"` |  |
| `"fromSession"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sessionId"` |  |
| `"tool"` |  |
| `"writeBack"` |  |

Operations: Create.

API path: `/session_run`

#### SessionSet

| Field | Description |
| --- | --- |
| `"entries"` |  |
| `"gate"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"sessionId"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/session_set`

#### SirnaDesign

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"minReynolds"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"shRnaLoop"` |  |
| `"target"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/sirna_design`

#### SiteDirectedMutagenesi

| Field | Description |
| --- | --- |
| `"armTmTarget"` |  |
| `"dntpMM"` |  |
| `"editKind"` |  |
| `"frameStart"` |  |
| `"gate"` |  |
| `"mgMM"` |  |
| `"naMM"` |  |
| `"newBase"` |  |
| `"ok"` |  |
| `"oligoNM"` |  |
| `"organism"` |  |
| `"position"` |  |
| `"provenance"` |  |
| `"residue"` |  |
| `"result"` |  |
| `"style"` |  |
| `"targetAa"` |  |
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
| `"toStop"` |  |
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
| `"frameStart"` |  |
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
| `"armTmTarget"` |  |
| `"circular"` |  |
| `"claimedConstruct"` |  |
| `"coding"` |  |
| `"enzyme"` |  |
| `"enzyme3"` |  |
| `"enzyme5"` |  |
| `"fragmentPcrs"` |  |
| `"fragments"` |  |
| `"frameStart"` |  |
| `"gate"` |  |
| `"insert"` |  |
| `"insertPcr"` |  |
| `"method"` |  |
| `"names"` |  |
| `"ok"` |  |
| `"overlapLen"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"tool"` |  |
| `"vector"` |  |
| `"vectorPcr"` |  |

Operations: Create.

API path: `/verify_assembly`

#### VerifyConstruct

| Field | Description |
| --- | --- |
| `"claimedConstruct"` |  |
| `"expectedFrameStart"` |  |
| `"gate"` |  |
| `"insertForwardPrimer"` |  |
| `"insertReversePrimer"` |  |
| `"insertTemplate"` |  |
| `"maxPrimerMismatches"` |  |
| `"ok"` |  |
| `"provenance"` |  |
| `"result"` |  |
| `"templateCircular"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/verify_construct`

#### VirtualGel

| Field | Description |
| --- | --- |
| `"circular"` |  |
| `"enzymes"` |  |
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
| `"rows"` |  |
| `"tool"` |  |

Operations: Create.

API path: `/volcano_plot_data`

#### WebSearch

| Field | Description |
| --- | --- |
| `"gate"` |  |
| `"max_results"` |  |
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
| `frameStart` | `int` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `target` | `string` |  |
| `targetPosition` | `int` |  |
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
| `args` | `map[string]any` |  |
| `capped` | `bool` |  |
| `columns` | `[]any` |  |
| `count` | `int` |  |
| `errors` | `int` |  |
| `input` | `string` |  |
| `limit` | `int` |  |
| `provenance` | `map[string]any` |  |
| `rows` | `[]any` |  |
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
    "capped": true,
    "columns": []any{},
    "count": 1,
    "errors": 1,
    "input": "example_input",
    "limit": 1,
    "provenance": map[string]any{},
    "rows": []any{},
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
| `capped` | `bool` |  |
| `columns` | `[]any` |  |
| `count` | `int` |  |
| `errors` | `int` |  |
| `input` | `string` |  |
| `limit` | `int` |  |
| `provenance` | `map[string]any` |  |
| `rows` | `[]any` |  |
| `steps` | `[]any` |  |

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
    "capped": true,
    "columns": []any{},
    "count": 1,
    "errors": 1,
    "input": "example_input",
    "limit": 1,
    "provenance": map[string]any{},
    "rows": []any{},
    "steps": []any{},
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
| `endPrimerLength` | `int` |  |
| `gate` | `any` |  |
| `maxOrfs` | `int` |  |
| `minOrfAa` | `int` |  |
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
| `armTmTarget` | `float64` |  |
| `circular` | `bool` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragments` | `[]any` |  |
| `gate` | `any` |  |
| `insert` | `string` |  |
| `method` | `string` |  |
| `names` | `[]any` |  |
| `ok` | `any` |  |
| `overlapLen` | `int` |  |
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
| `frameStart` | `int` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `organism` | `string` |  |
| `provenance` | `map[string]any` |  |
| `rareThreshold` | `float64` |  |
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
| `avoidEnzymes` | `[]any` |  |
| `crypticOrfMinAa` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `any` |  |
| `gcHigh` | `float64` |  |
| `gcLow` | `float64` |  |
| `gcWindow` | `int` |  |
| `homopolymerMin` | `int` |  |
| `maxPasses` | `int` |  |
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
| `avoidEnzymes` | `[]any` |  |
| `crypticOrfMinAa` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `any` |  |
| `gcHigh` | `float64` |  |
| `gcLow` | `float64` |  |
| `gcWindow` | `int` |  |
| `homopolymerMin` | `int` |  |
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
| `minScore` | `float64` |  |
| `nuclease` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `searchReverseStrand` | `bool` |  |
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
| `armLength` | `int` |  |
| `blockPam` | `bool` |  |
| `designGenotypingPrimers` | `bool` |  |
| `editEnd` | `int` |  |
| `editStart` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `any` |  |
| `guideEnd` | `int` |  |
| `guideStart` | `int` |  |
| `guideStrand` | `string` |  |
| `nuclease` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `replacement` | `string` |  |
| `result` | `map[string]any` |  |
| `targetSequence` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.CrisprHdrDonor(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "replacement": "example_replacement",
    "result": map[string]any{},
    "targetSequence": "example_targetSequence",
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
| `maxMismatches` | `int` |  |
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
| `sequenceA` | `string` |  |
| `sequenceB` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.CrossDimer(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sequenceA": "example_sequenceA",
    "sequenceB": "example_sequenceB",
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
| `massNg` | `float64` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |
| `volumeUl` | `float64` |  |

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
| `enzymeA` | `string` |  |
| `enzymeB` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.DoubleDigest(nil).Create(map[string]any{
    "enzymeA": "example_enzymeA",
    "enzymeB": "example_enzymeB",
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
| `reactions` | `[]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ExportEchoPicklist(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reactions": []any{},
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
| `protocolName` | `string` |  |
| `provenance` | `map[string]any` |  |
| `reactions` | `[]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ExportOpentronsProtocol(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reactions": []any{},
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
| `reactions` | `[]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ExportPlateLayout(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reactions": []any{},
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
| `clusterCols` | `bool` |  |
| `clusterRows` | `bool` |  |
| `distanceMetric` | `string` |  |
| `gate` | `any` |  |
| `genes` | `[]any` |  |
| `linkage` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `samples` | `[]any` |  |
| `tool` | `string` |  |
| `values` | `[]any` |  |
| `zScoreRows` | `bool` |  |

#### Example: Create

```go
result, err := client.ExpressionHeatmapCluster(nil).Create(map[string]any{
    "genes": []any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "samples": []any{},
    "tool": "example_tool",
    "values": []any{},
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
| `qualityOffset` | `int` |  |
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
| `minLength` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `qualityOffset` | `int` |  |
| `qualityThreshold` | `int` |  |
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
| `minAaLength` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `requireStop` | `bool` |  |
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
| `caseMode` | `string` |  |
| `convert` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `reverse` | `bool` |  |
| `sequence` | `string` |  |
| `stripNonLetters` | `bool` |  |
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
| `collections` | `[]any` |  |
| `gate` | `any` |  |
| `genes` | `[]any` |  |
| `maxTermSize` | `int` |  |
| `minTermSize` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.FunctionalEnrichment(nil).Create(map[string]any{
    "genes": []any{},
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
| `compareToNamedSet` | `string` |  |
| `dataset` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `overhangs` | `[]any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `riskThreshold` | `float64` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.GoldenGateFidelity(nil).Create(map[string]any{
    "ok": "example_ok",
    "overhangs": []any{},
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
| `jobId` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.IdMapPoll(nil).Create(map[string]any{
    "jobId": "example_jobId",
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
| `taxId` | `string` |  |
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
| `forwardPrimer` | `string` |  |
| `gate` | `any` |  |
| `maxMismatches` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `reversePrimer` | `string` |  |
| `template` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.InSilicoPcr(nil).Create(map[string]any{
    "forwardPrimer": "example_forwardPrimer",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "reversePrimer": "example_reversePrimer",
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
| `addSecondaryMismatch` | `bool` |  |
| `alleleA` | `string` |  |
| `alleleB` | `string` |  |
| `gate` | `any` |  |
| `maxAmplicon` | `int` |  |
| `minAmplicon` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `snpPosition` | `int` |  |
| `target` | `string` |  |
| `targetCoreTm` | `float64` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.KaspPrimerDesign(nil).Create(map[string]any{
    "alleleA": "example_alleleA",
    "alleleB": "example_alleleB",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "snpPosition": 1,
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
| `dntpMM` | `float64` |  |
| `gate` | `any` |  |
| `mgMM` | `float64` |  |
| `naMM` | `float64` |  |
| `ok` | `any` |  |
| `oligoNM` | `float64` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sequence` | `string` |  |
| `targetTm` | `float64` |  |
| `tmTolerance` | `float64` |  |
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
| `maxMismatches` | `int` |  |
| `motif` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `searchReverseStrand` | `bool` |  |
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
| `dntpMM` | `float64` |  |
| `gate` | `any` |  |
| `mgMM` | `float64` |  |
| `naMM` | `float64` |  |
| `ok` | `any` |  |
| `oligoNM` | `float64` |  |
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
| `sourceSpecies` | `string` |  |
| `symbols` | `[]any` |  |
| `targetSpecies` | `string` |  |
| `tool` | `string` |  |
| `type` | `string` |  |

#### Example: Create

```go
result, err := client.OrthologMap(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "symbols": []any{},
    "targetSpecies": "example_targetSpecies",
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
| `seqA` | `string` |  |
| `seqB` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.PairwiseAlignment(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "seqA": "example_seqA",
    "seqB": "example_seqB",
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
| `fileBase64` | `string` |  |
| `fileName` | `string` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ParseSangerTrace(nil).Create(map[string]any{
    "fileBase64": "example_fileBase64",
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
| `topN` | `int` |  |

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
| `topN` | `int` |  |

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
| `editEnd` | `int` |  |
| `editStart` | `int` |  |
| `frameStart` | `int` |  |
| `gate` | `any` |  |
| `insertedSeq` | `string` |  |
| `ok` | `any` |  |
| `pbsLength` | `int` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `rttHomology` | `int` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.PrimeEditingDesign(nil).Create(map[string]any{
    "editEnd": 1,
    "editStart": 1,
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
| `newSequence` | `string` |  |
| `ok` | `any` |  |
| `overlapLength` | `int` |  |
| `pbsLength` | `int` |  |
| `provenance` | `map[string]any` |  |
| `replaceEnd` | `int` |  |
| `replaceStart` | `int` |  |
| `result` | `map[string]any` |  |
| `target` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.PrimeEditingTwinDesign(nil).Create(map[string]any{
    "newSequence": "example_newSequence",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "replaceEnd": 1,
    "replaceStart": 1,
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
| `ampliconMax` | `int` |  |
| `ampliconMin` | `int` |  |
| `dntpMM` | `float64` |  |
| `gate` | `any` |  |
| `gcMax` | `float64` |  |
| `gcMin` | `float64` |  |
| `lenMax` | `int` |  |
| `lenMin` | `int` |  |
| `lenOpt` | `int` |  |
| `maxReturn` | `int` |  |
| `mgMM` | `float64` |  |
| `naMM` | `float64` |  |
| `ok` | `any` |  |
| `oligoNM` | `float64` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `targetEnd` | `int` |  |
| `targetStart` | `int` |  |
| `template` | `string` |  |
| `tmMax` | `float64` |  |
| `tmMaxDiff` | `float64` |  |
| `tmMin` | `float64` |  |
| `tmOpt` | `float64` |  |
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
| `forwardPrimer` | `string` |  |
| `gate` | `any` |  |
| `maxMismatches` | `int` |  |
| `maxProductLength` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `reversePrimer` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.PrimerSpecificity(nil).Create(map[string]any{
    "forwardPrimer": "example_forwardPrimer",
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "reversePrimer": "example_reversePrimer",
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
| `maxMass` | `float64` |  |
| `maxPeptides` | `int` |  |
| `minMass` | `float64` |  |
| `missedCleavages` | `int` |  |
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
| `jobId` | `string` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.ProteinAnnotatePoll(nil).Create(map[string]any{
    "jobId": "example_jobId",
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
| `goterms` | `bool` |  |
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
| `chargeStep` | `float64` |  |
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
| `gcContent` | `float64` |  |
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
| `enzymes` | `[]any` |  |
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
| `fileBase64` | `string` |  |
| `fileName` | `string` |  |
| `gate` | `any` |  |
| `minCoverage` | `float64` |  |
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
| `args` | `map[string]any` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SavePermalink(nil).Create(map[string]any{
    "args": map[string]any{},
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
| `qualityOffset` | `int` |  |
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
| `endPrimerLength` | `int` |  |
| `gate` | `any` |  |
| `maxOrfs` | `int` |  |
| `minOrfAa` | `int` |  |
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
| `maxResults` | `int` |  |
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
| `minSupportingReads` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `reads` | `string` |  |
| `reference` | `string` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SequencingReadbackVerify(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "reads": "example_reads",
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
| `entries` | `map[string]any` |  |
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
| `names` | `[]any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sessionId` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SessionGet(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sessionId": "example_sessionId",
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
| `args` | `map[string]any` |  |
| `fromSession` | `map[string]any` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sessionId` | `string` |  |
| `tool` | `string` |  |
| `writeBack` | `map[string]any` |  |

#### Example: Create

```go
result, err := client.SessionRun(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sessionId": "example_sessionId",
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
| `entries` | `map[string]any` |  |
| `gate` | `any` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `sessionId` | `string` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.SessionSet(nil).Create(map[string]any{
    "entries": map[string]any{},
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "sessionId": "example_sessionId",
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
| `minReynolds` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `shRnaLoop` | `string` |  |
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
| `armTmTarget` | `float64` |  |
| `dntpMM` | `float64` |  |
| `editKind` | `string` |  |
| `frameStart` | `int` |  |
| `gate` | `any` |  |
| `mgMM` | `float64` |  |
| `naMM` | `float64` |  |
| `newBase` | `string` |  |
| `ok` | `any` |  |
| `oligoNM` | `float64` |  |
| `organism` | `string` |  |
| `position` | `int` |  |
| `provenance` | `map[string]any` |  |
| `residue` | `int` |  |
| `result` | `map[string]any` |  |
| `style` | `string` |  |
| `targetAa` | `string` |  |
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
| `toStop` | `bool` |  |
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
| `frameStart` | `int` |  |
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
| `armTmTarget` | `float64` |  |
| `circular` | `bool` |  |
| `claimedConstruct` | `string` |  |
| `coding` | `bool` |  |
| `enzyme` | `string` |  |
| `enzyme3` | `string` |  |
| `enzyme5` | `string` |  |
| `fragmentPcrs` | `[]any` |  |
| `fragments` | `[]any` |  |
| `frameStart` | `int` |  |
| `gate` | `any` |  |
| `insert` | `string` |  |
| `insertPcr` | `map[string]any` |  |
| `method` | `string` |  |
| `names` | `[]any` |  |
| `ok` | `any` |  |
| `overlapLen` | `int` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `tool` | `string` |  |
| `vector` | `string` |  |
| `vectorPcr` | `map[string]any` |  |

#### Example: Create

```go
result, err := client.VerifyAssembly(nil).Create(map[string]any{
    "claimedConstruct": "example_claimedConstruct",
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
| `claimedConstruct` | `string` |  |
| `expectedFrameStart` | `int` |  |
| `gate` | `any` |  |
| `insertForwardPrimer` | `string` |  |
| `insertReversePrimer` | `string` |  |
| `insertTemplate` | `string` |  |
| `maxPrimerMismatches` | `int` |  |
| `ok` | `any` |  |
| `provenance` | `map[string]any` |  |
| `result` | `map[string]any` |  |
| `templateCircular` | `bool` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.VerifyConstruct(nil).Create(map[string]any{
    "claimedConstruct": "example_claimedConstruct",
    "insertForwardPrimer": "example_insertForwardPrimer",
    "insertReversePrimer": "example_insertReversePrimer",
    "insertTemplate": "example_insertTemplate",
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
| `enzymes` | `[]any` |  |
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
| `rows` | `[]any` |  |
| `tool` | `string` |  |

#### Example: Create

```go
result, err := client.VolcanoPlotData(nil).Create(map[string]any{
    "ok": "example_ok",
    "provenance": map[string]any{},
    "result": map[string]any{},
    "rows": []any{},
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
| `max_results` | `float64` |  |
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
