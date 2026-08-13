package core

import (
	"fmt"
	"strings"

	vs "github.com/voxgig-sdk/seqbench-mcp-sdk/go/utility/struct"
)

type SeqbenchMcpSDK struct {
	Mode     string
	options  map[string]any
	utility  *Utility
	Features []Feature
	rootctx  *Context
}

func NewSeqbenchMcpSDK(options map[string]any) *SeqbenchMcpSDK {
	sdk := &SeqbenchMcpSDK{
		Mode:     "live",
		Features: []Feature{},
	}

	sdk.utility = NewUtility()

	config := MakeConfig()

	sdk.rootctx = sdk.utility.MakeContext(map[string]any{
		"client":  sdk,
		"utility": sdk.utility,
		"config":  config,
		"options": options,
		"shared":  map[string]any{},
	}, nil)

	sdk.options = sdk.utility.MakeOptions(sdk.rootctx)

	if vs.GetPath([]any{"feature", "test", "active"}, sdk.options) == true {
		sdk.Mode = "test"
	}

	sdk.rootctx.Options = sdk.options

	// Add features in the resolved order (MakeOptions puts an explicit array
	// order first, else defaults to test-first). Ordering matters: the `test`
	// feature installs the base mock transport and the transport features
	// (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
	// must be added before them to sit at the base of the chain.
	featureOpts := ToMapAny(vs.GetProp(sdk.options, "feature"))
	if featureOpts != nil {
		if fo, ok := vs.GetPath([]any{"__derived__", "featureorder"}, sdk.options).([]any); ok {
			for _, n := range fo {
				fname, _ := n.(string)
				fopts := ToMapAny(featureOpts[fname])
				if fopts != nil {
					if active, ok := fopts["active"]; ok {
						if ab, ok := active.(bool); ok && ab {
							sdk.utility.FeatureAdd(sdk.rootctx, makeFeature(fname))
						}
					}
				}
			}
		}
	}

	// Add extension features.
	if extend := vs.GetProp(sdk.options, "extend"); extend != nil {
		if extList, ok := extend.([]any); ok {
			for _, f := range extList {
				if feat, ok := f.(Feature); ok {
					sdk.utility.FeatureAdd(sdk.rootctx, feat)
				}
			}
		}
	}

	// Initialize features.
	for _, f := range sdk.Features {
		sdk.utility.FeatureInit(sdk.rootctx, f)
	}

	sdk.utility.FeatureHook(sdk.rootctx, "PostConstruct")

	return sdk
}

func (sdk *SeqbenchMcpSDK) OptionsMap() map[string]any {
	out := vs.Clone(sdk.options)
	if om, ok := out.(map[string]any); ok {
		return om
	}
	return map[string]any{}
}

func (sdk *SeqbenchMcpSDK) GetUtility() *Utility {
	return CopyUtility(sdk.utility)
}

func (sdk *SeqbenchMcpSDK) GetRootCtx() *Context {
	return sdk.rootctx
}

func (sdk *SeqbenchMcpSDK) Prepare(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "prepare",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	options := sdk.options

	path, _ := vs.GetProp(fetchargs, "path").(string)
	method, _ := vs.GetProp(fetchargs, "method").(string)
	if method == "" {
		method = "GET"
	}

	params := ToMapAny(vs.GetProp(fetchargs, "params"))
	if params == nil {
		params = map[string]any{}
	}
	query := ToMapAny(vs.GetProp(fetchargs, "query"))
	if query == nil {
		query = map[string]any{}
	}

	headers := utility.PrepareHeaders(ctx)

	base, _ := vs.GetProp(options, "base").(string)
	prefix, _ := vs.GetProp(options, "prefix").(string)
	suffix, _ := vs.GetProp(options, "suffix").(string)

	ctx.Spec = NewSpec(map[string]any{
		"base":    base,
		"prefix":  prefix,
		"suffix":  suffix,
		"path":    path,
		"method":  method,
		"params":  params,
		"query":   query,
		"headers": headers,
		"body":    vs.GetProp(fetchargs, "body"),
		"step":    "start",
	})

	// Merge user-provided headers.
	if uh := vs.GetProp(fetchargs, "headers"); uh != nil {
		if uhm, ok := uh.(map[string]any); ok {
			for k, v := range uhm {
				ctx.Spec.Headers[k] = v
			}
		}
	}

	_, err := utility.PrepareAuth(ctx)
	if err != nil {
		return nil, err
	}

	return utility.MakeFetchDef(ctx)
}

// Raw endpoint access is operator-controllable, like every entity op.
// Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
// either one reaches the same endpoint.
func (sdk *SeqbenchMcpSDK) Direct(fetchargs map[string]any) (map[string]any, error) {
	if !sdk.opAllowed("direct") {
		return sdk.opDenied("direct"), nil
	}

	return sdk.rawRequest(fetchargs)
}

// Is this raw-access op permitted by the SDK's allow.op option?
func (sdk *SeqbenchMcpSDK) opAllowed(op string) bool {
	allowOp, _ := vs.GetPath([]any{"allow", "op"}, sdk.options).(string)
	return strings.Contains(allowOp, op)
}

func (sdk *SeqbenchMcpSDK) opDenied(op string) map[string]any {
	allowOp, _ := vs.GetPath([]any{"allow", "op"}, sdk.options).(string)
	return map[string]any{
		"ok": false,
		"err": fmt.Errorf("SeqbenchMcpSDK: %s: operation not allowed by"+
			" SDK option allow.op value: \"%s\"", op, allowOp),
	}
}

// Ungated request path shared by Direct and Graphql, each of which checks
// its own allow.op token first. Unexported, rather than a flag on fetchargs:
// a caller-supplied marker would let anyone opt straight back out of the
// gate by passing it.
func (sdk *SeqbenchMcpSDK) rawRequest(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	fetchdef, err := sdk.Prepare(fetchargs)
	if err != nil {
		return map[string]any{"ok": false, "err": err}, nil
	}

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "direct",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	url, _ := fetchdef["url"].(string)
	fetched, fetchErr := utility.Fetcher(ctx, url, fetchdef)

	if fetchErr != nil {
		return map[string]any{"ok": false, "err": fetchErr}, nil
	}

	if fetched == nil {
		return map[string]any{
			"ok":  false,
			"err": ctx.MakeError("direct_no_response", "response: undefined"),
		}, nil
	}

	if fm, ok := fetched.(map[string]any); ok {
		status := ToInt(vs.GetProp(fm, "status"))
		headers := vs.GetProp(fm, "headers")

		// No-body responses (204, 304) and explicit zero content-length
		// must skip JSON parsing — calling json() on an empty body errors.
		var contentLength string
		if hm, ok := headers.(map[string]any); ok {
			if cl, ok := hm["content-length"]; ok {
				contentLength = fmt.Sprintf("%v", cl)
			}
		}
		noBody := status == 204 || status == 304 || contentLength == "0"

		var jsonData any
		if !noBody {
			if jf := vs.GetProp(fm, "json"); jf != nil {
				if f, ok := jf.(func() any); ok {
					// f() returns nil on parse error in our fetcher.
					jsonData = f()
				}
			}
		}

		return map[string]any{
			"ok":      status >= 200 && status < 300,
			"status":  status,
			"headers": headers,
			"data":    jsonData,
		}, nil
	}

	return map[string]any{"ok": false, "err": ctx.MakeError("direct_invalid", "invalid response type")}, nil
}

// Raw GraphQL access: the pressure valve that makes the generated surface's
// deliberate omissions (per-call selection sets, typed filter builders,
// batching, subscriptions) livable — the whole schema stays reachable.
//
// Thin wrapper over the same prepare/fetch path Direct uses, with the one
// thing raw Direct cannot do for GraphQL: a GraphQL failure rides HTTP 200
// as a top-level `errors` array, so status alone would report a failed query
// as ok.
//
// NOTE: like Direct, this bypasses the feature pipeline — no retry,
// ratelimit or paging features apply.
func (sdk *SeqbenchMcpSDK) Graphql(
	query string, variables map[string]any, ctrl map[string]any,
) (map[string]any, error) {
	if !sdk.opAllowed("graphql") {
		return sdk.opDenied("graphql"), nil
	}

	if variables == nil {
		variables = map[string]any{}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	res, err := sdk.rawRequest(map[string]any{
		"method":  "POST",
		"headers": map[string]any{"content-type": "application/json"},
		"body":    map[string]any{"query": query, "variables": variables},
		"ctrl":    ctrl,
	})

	if err != nil {
		return res, err
	}

	// Errors are read BEFORE any status check: a GraphQL parse or validation
	// failure comes back as HTTP 400 carrying the standard { errors: [...] }
	// body, and the raw path represents a non-2xx as ok:false with no err —
	// so returning early on status would discard the server's own
	// diagnostics, which are the only useful part of that response.
	errors, _ := vs.GetPath([]any{"data", "errors"}, res).([]any)

	if 0 < len(errors) {
		msg, _ := vs.GetProp(errors[0], "message").(string)
		if msg == "" {
			msg = "graphql error"
		}
		res["ok"] = false
		res["err"] = fmt.Errorf("SeqbenchMcpSDK: graphql: %s", msg)
		res["graphql"] = errors
	}

	return res, nil
}


// AlphafoldLookup returns a AlphafoldLookup entity bound to this client.
// Idiomatic usage: client.AlphafoldLookup(nil).List(nil, nil) or
// client.AlphafoldLookup(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) AlphafoldLookup(data map[string]any) SeqbenchMcpEntity {
	return NewAlphafoldLookupEntityFunc(sdk, data)
}


// AsoDesign returns a AsoDesign entity bound to this client.
// Idiomatic usage: client.AsoDesign(nil).List(nil, nil) or
// client.AsoDesign(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) AsoDesign(data map[string]any) SeqbenchMcpEntity {
	return NewAsoDesignEntityFunc(sdk, data)
}


// BaseEditingDesign returns a BaseEditingDesign entity bound to this client.
// Idiomatic usage: client.BaseEditingDesign(nil).List(nil, nil) or
// client.BaseEditingDesign(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) BaseEditingDesign(data map[string]any) SeqbenchMcpEntity {
	return NewBaseEditingDesignEntityFunc(sdk, data)
}


// Batch returns a Batch entity bound to this client.
// Idiomatic usage: client.Batch(nil).List(nil, nil) or
// client.Batch(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) Batch(data map[string]any) SeqbenchMcpEntity {
	return NewBatchEntityFunc(sdk, data)
}


// BatchWorkflow returns a BatchWorkflow entity bound to this client.
// Idiomatic usage: client.BatchWorkflow(nil).List(nil, nil) or
// client.BatchWorkflow(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) BatchWorkflow(data map[string]any) SeqbenchMcpEntity {
	return NewBatchWorkflowEntityFunc(sdk, data)
}


// CharacterizeSequence returns a CharacterizeSequence entity bound to this client.
// Idiomatic usage: client.CharacterizeSequence(nil).List(nil, nil) or
// client.CharacterizeSequence(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) CharacterizeSequence(data map[string]any) SeqbenchMcpEntity {
	return NewCharacterizeSequenceEntityFunc(sdk, data)
}


// CloningSimulate returns a CloningSimulate entity bound to this client.
// Idiomatic usage: client.CloningSimulate(nil).List(nil, nil) or
// client.CloningSimulate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) CloningSimulate(data map[string]any) SeqbenchMcpEntity {
	return NewCloningSimulateEntityFunc(sdk, data)
}


// CodonAdaptationIndex returns a CodonAdaptationIndex entity bound to this client.
// Idiomatic usage: client.CodonAdaptationIndex(nil).List(nil, nil) or
// client.CodonAdaptationIndex(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) CodonAdaptationIndex(data map[string]any) SeqbenchMcpEntity {
	return NewCodonAdaptationIndexEntityFunc(sdk, data)
}


// CodonOptimize returns a CodonOptimize entity bound to this client.
// Idiomatic usage: client.CodonOptimize(nil).List(nil, nil) or
// client.CodonOptimize(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) CodonOptimize(data map[string]any) SeqbenchMcpEntity {
	return NewCodonOptimizeEntityFunc(sdk, data)
}


// ConstructAutofix returns a ConstructAutofix entity bound to this client.
// Idiomatic usage: client.ConstructAutofix(nil).List(nil, nil) or
// client.ConstructAutofix(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ConstructAutofix(data map[string]any) SeqbenchMcpEntity {
	return NewConstructAutofixEntityFunc(sdk, data)
}


// ConstructQc returns a ConstructQc entity bound to this client.
// Idiomatic usage: client.ConstructQc(nil).List(nil, nil) or
// client.ConstructQc(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ConstructQc(data map[string]any) SeqbenchMcpEntity {
	return NewConstructQcEntityFunc(sdk, data)
}


// CrisprGrnaDesign returns a CrisprGrnaDesign entity bound to this client.
// Idiomatic usage: client.CrisprGrnaDesign(nil).List(nil, nil) or
// client.CrisprGrnaDesign(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) CrisprGrnaDesign(data map[string]any) SeqbenchMcpEntity {
	return NewCrisprGrnaDesignEntityFunc(sdk, data)
}


// CrisprHdrDonor returns a CrisprHdrDonor entity bound to this client.
// Idiomatic usage: client.CrisprHdrDonor(nil).List(nil, nil) or
// client.CrisprHdrDonor(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) CrisprHdrDonor(data map[string]any) SeqbenchMcpEntity {
	return NewCrisprHdrDonorEntityFunc(sdk, data)
}


// CrisprOfftargetCheck returns a CrisprOfftargetCheck entity bound to this client.
// Idiomatic usage: client.CrisprOfftargetCheck(nil).List(nil, nil) or
// client.CrisprOfftargetCheck(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) CrisprOfftargetCheck(data map[string]any) SeqbenchMcpEntity {
	return NewCrisprOfftargetCheckEntityFunc(sdk, data)
}


// CrossDimer returns a CrossDimer entity bound to this client.
// Idiomatic usage: client.CrossDimer(nil).List(nil, nil) or
// client.CrossDimer(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) CrossDimer(data map[string]any) SeqbenchMcpEntity {
	return NewCrossDimerEntityFunc(sdk, data)
}


// DnaMolarity returns a DnaMolarity entity bound to this client.
// Idiomatic usage: client.DnaMolarity(nil).List(nil, nil) or
// client.DnaMolarity(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) DnaMolarity(data map[string]any) SeqbenchMcpEntity {
	return NewDnaMolarityEntityFunc(sdk, data)
}


// DoubleDigest returns a DoubleDigest entity bound to this client.
// Idiomatic usage: client.DoubleDigest(nil).List(nil, nil) or
// client.DoubleDigest(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) DoubleDigest(data map[string]any) SeqbenchMcpEntity {
	return NewDoubleDigestEntityFunc(sdk, data)
}


// ExportEchoPicklist returns a ExportEchoPicklist entity bound to this client.
// Idiomatic usage: client.ExportEchoPicklist(nil).List(nil, nil) or
// client.ExportEchoPicklist(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ExportEchoPicklist(data map[string]any) SeqbenchMcpEntity {
	return NewExportEchoPicklistEntityFunc(sdk, data)
}


// ExportOpentronsProtocol returns a ExportOpentronsProtocol entity bound to this client.
// Idiomatic usage: client.ExportOpentronsProtocol(nil).List(nil, nil) or
// client.ExportOpentronsProtocol(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ExportOpentronsProtocol(data map[string]any) SeqbenchMcpEntity {
	return NewExportOpentronsProtocolEntityFunc(sdk, data)
}


// ExportPlateLayout returns a ExportPlateLayout entity bound to this client.
// Idiomatic usage: client.ExportPlateLayout(nil).List(nil, nil) or
// client.ExportPlateLayout(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ExportPlateLayout(data map[string]any) SeqbenchMcpEntity {
	return NewExportPlateLayoutEntityFunc(sdk, data)
}


// ExpressionHeatmapCluster returns a ExpressionHeatmapCluster entity bound to this client.
// Idiomatic usage: client.ExpressionHeatmapCluster(nil).List(nil, nil) or
// client.ExpressionHeatmapCluster(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ExpressionHeatmapCluster(data map[string]any) SeqbenchMcpEntity {
	return NewExpressionHeatmapClusterEntityFunc(sdk, data)
}


// FastqQcReport returns a FastqQcReport entity bound to this client.
// Idiomatic usage: client.FastqQcReport(nil).List(nil, nil) or
// client.FastqQcReport(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) FastqQcReport(data map[string]any) SeqbenchMcpEntity {
	return NewFastqQcReportEntityFunc(sdk, data)
}


// FastqTrim returns a FastqTrim entity bound to this client.
// Idiomatic usage: client.FastqTrim(nil).List(nil, nil) or
// client.FastqTrim(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) FastqTrim(data map[string]any) SeqbenchMcpEntity {
	return NewFastqTrimEntityFunc(sdk, data)
}


// FindOrf returns a FindOrf entity bound to this client.
// Idiomatic usage: client.FindOrf(nil).List(nil, nil) or
// client.FindOrf(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) FindOrf(data map[string]any) SeqbenchMcpEntity {
	return NewFindOrfEntityFunc(sdk, data)
}


// FormatSequence returns a FormatSequence entity bound to this client.
// Idiomatic usage: client.FormatSequence(nil).List(nil, nil) or
// client.FormatSequence(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) FormatSequence(data map[string]any) SeqbenchMcpEntity {
	return NewFormatSequenceEntityFunc(sdk, data)
}


// FunctionalEnrichment returns a FunctionalEnrichment entity bound to this client.
// Idiomatic usage: client.FunctionalEnrichment(nil).List(nil, nil) or
// client.FunctionalEnrichment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) FunctionalEnrichment(data map[string]any) SeqbenchMcpEntity {
	return NewFunctionalEnrichmentEntityFunc(sdk, data)
}


// GcContent returns a GcContent entity bound to this client.
// Idiomatic usage: client.GcContent(nil).List(nil, nil) or
// client.GcContent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) GcContent(data map[string]any) SeqbenchMcpEntity {
	return NewGcContentEntityFunc(sdk, data)
}


// GeneDossier returns a GeneDossier entity bound to this client.
// Idiomatic usage: client.GeneDossier(nil).List(nil, nil) or
// client.GeneDossier(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) GeneDossier(data map[string]any) SeqbenchMcpEntity {
	return NewGeneDossierEntityFunc(sdk, data)
}


// GeneExpression returns a GeneExpression entity bound to this client.
// Idiomatic usage: client.GeneExpression(nil).List(nil, nil) or
// client.GeneExpression(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) GeneExpression(data map[string]any) SeqbenchMcpEntity {
	return NewGeneExpressionEntityFunc(sdk, data)
}


// GeneModel returns a GeneModel entity bound to this client.
// Idiomatic usage: client.GeneModel(nil).List(nil, nil) or
// client.GeneModel(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) GeneModel(data map[string]any) SeqbenchMcpEntity {
	return NewGeneModelEntityFunc(sdk, data)
}


// GoldenGateFidelity returns a GoldenGateFidelity entity bound to this client.
// Idiomatic usage: client.GoldenGateFidelity(nil).List(nil, nil) or
// client.GoldenGateFidelity(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) GoldenGateFidelity(data map[string]any) SeqbenchMcpEntity {
	return NewGoldenGateFidelityEntityFunc(sdk, data)
}


// HgvsConvert returns a HgvsConvert entity bound to this client.
// Idiomatic usage: client.HgvsConvert(nil).List(nil, nil) or
// client.HgvsConvert(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) HgvsConvert(data map[string]any) SeqbenchMcpEntity {
	return NewHgvsConvertEntityFunc(sdk, data)
}


// IdMapPoll returns a IdMapPoll entity bound to this client.
// Idiomatic usage: client.IdMapPoll(nil).List(nil, nil) or
// client.IdMapPoll(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) IdMapPoll(data map[string]any) SeqbenchMcpEntity {
	return NewIdMapPollEntityFunc(sdk, data)
}


// IdMapSubmit returns a IdMapSubmit entity bound to this client.
// Idiomatic usage: client.IdMapSubmit(nil).List(nil, nil) or
// client.IdMapSubmit(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) IdMapSubmit(data map[string]any) SeqbenchMcpEntity {
	return NewIdMapSubmitEntityFunc(sdk, data)
}


// InSilicoPcr returns a InSilicoPcr entity bound to this client.
// Idiomatic usage: client.InSilicoPcr(nil).List(nil, nil) or
// client.InSilicoPcr(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) InSilicoPcr(data map[string]any) SeqbenchMcpEntity {
	return NewInSilicoPcrEntityFunc(sdk, data)
}


// KaspPrimerDesign returns a KaspPrimerDesign entity bound to this client.
// Idiomatic usage: client.KaspPrimerDesign(nil).List(nil, nil) or
// client.KaspPrimerDesign(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) KaspPrimerDesign(data map[string]any) SeqbenchMcpEntity {
	return NewKaspPrimerDesignEntityFunc(sdk, data)
}


// ListTool returns a ListTool entity bound to this client.
// Idiomatic usage: client.ListTool(nil).List(nil, nil) or
// client.ListTool(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ListTool(data map[string]any) SeqbenchMcpEntity {
	return NewListToolEntityFunc(sdk, data)
}


// MeltingTemperature returns a MeltingTemperature entity bound to this client.
// Idiomatic usage: client.MeltingTemperature(nil).List(nil, nil) or
// client.MeltingTemperature(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) MeltingTemperature(data map[string]any) SeqbenchMcpEntity {
	return NewMeltingTemperatureEntityFunc(sdk, data)
}


// MotifFinder returns a MotifFinder entity bound to this client.
// Idiomatic usage: client.MotifFinder(nil).List(nil, nil) or
// client.MotifFinder(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) MotifFinder(data map[string]any) SeqbenchMcpEntity {
	return NewMotifFinderEntityFunc(sdk, data)
}


// MultipleSequenceAlignment returns a MultipleSequenceAlignment entity bound to this client.
// Idiomatic usage: client.MultipleSequenceAlignment(nil).List(nil, nil) or
// client.MultipleSequenceAlignment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) MultipleSequenceAlignment(data map[string]any) SeqbenchMcpEntity {
	return NewMultipleSequenceAlignmentEntityFunc(sdk, data)
}


// OligoAnalysi returns a OligoAnalysi entity bound to this client.
// Idiomatic usage: client.OligoAnalysi(nil).List(nil, nil) or
// client.OligoAnalysi(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) OligoAnalysi(data map[string]any) SeqbenchMcpEntity {
	return NewOligoAnalysiEntityFunc(sdk, data)
}


// OrthologMap returns a OrthologMap entity bound to this client.
// Idiomatic usage: client.OrthologMap(nil).List(nil, nil) or
// client.OrthologMap(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) OrthologMap(data map[string]any) SeqbenchMcpEntity {
	return NewOrthologMapEntityFunc(sdk, data)
}


// PairwiseAlignment returns a PairwiseAlignment entity bound to this client.
// Idiomatic usage: client.PairwiseAlignment(nil).List(nil, nil) or
// client.PairwiseAlignment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) PairwiseAlignment(data map[string]any) SeqbenchMcpEntity {
	return NewPairwiseAlignmentEntityFunc(sdk, data)
}


// ParseGenbank returns a ParseGenbank entity bound to this client.
// Idiomatic usage: client.ParseGenbank(nil).List(nil, nil) or
// client.ParseGenbank(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ParseGenbank(data map[string]any) SeqbenchMcpEntity {
	return NewParseGenbankEntityFunc(sdk, data)
}


// ParseSangerTrace returns a ParseSangerTrace entity bound to this client.
// Idiomatic usage: client.ParseSangerTrace(nil).List(nil, nil) or
// client.ParseSangerTrace(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ParseSangerTrace(data map[string]any) SeqbenchMcpEntity {
	return NewParseSangerTraceEntityFunc(sdk, data)
}


// PlasmidAnnotate returns a PlasmidAnnotate entity bound to this client.
// Idiomatic usage: client.PlasmidAnnotate(nil).List(nil, nil) or
// client.PlasmidAnnotate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) PlasmidAnnotate(data map[string]any) SeqbenchMcpEntity {
	return NewPlasmidAnnotateEntityFunc(sdk, data)
}


// PlasmidDeepAnnotate returns a PlasmidDeepAnnotate entity bound to this client.
// Idiomatic usage: client.PlasmidDeepAnnotate(nil).List(nil, nil) or
// client.PlasmidDeepAnnotate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) PlasmidDeepAnnotate(data map[string]any) SeqbenchMcpEntity {
	return NewPlasmidDeepAnnotateEntityFunc(sdk, data)
}


// PlasmidFullReport returns a PlasmidFullReport entity bound to this client.
// Idiomatic usage: client.PlasmidFullReport(nil).List(nil, nil) or
// client.PlasmidFullReport(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) PlasmidFullReport(data map[string]any) SeqbenchMcpEntity {
	return NewPlasmidFullReportEntityFunc(sdk, data)
}


// PlasmidIdentify returns a PlasmidIdentify entity bound to this client.
// Idiomatic usage: client.PlasmidIdentify(nil).List(nil, nil) or
// client.PlasmidIdentify(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) PlasmidIdentify(data map[string]any) SeqbenchMcpEntity {
	return NewPlasmidIdentifyEntityFunc(sdk, data)
}


// PrimeEditingDesign returns a PrimeEditingDesign entity bound to this client.
// Idiomatic usage: client.PrimeEditingDesign(nil).List(nil, nil) or
// client.PrimeEditingDesign(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) PrimeEditingDesign(data map[string]any) SeqbenchMcpEntity {
	return NewPrimeEditingDesignEntityFunc(sdk, data)
}


// PrimeEditingTwinDesign returns a PrimeEditingTwinDesign entity bound to this client.
// Idiomatic usage: client.PrimeEditingTwinDesign(nil).List(nil, nil) or
// client.PrimeEditingTwinDesign(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) PrimeEditingTwinDesign(data map[string]any) SeqbenchMcpEntity {
	return NewPrimeEditingTwinDesignEntityFunc(sdk, data)
}


// PrimerDesign returns a PrimerDesign entity bound to this client.
// Idiomatic usage: client.PrimerDesign(nil).List(nil, nil) or
// client.PrimerDesign(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) PrimerDesign(data map[string]any) SeqbenchMcpEntity {
	return NewPrimerDesignEntityFunc(sdk, data)
}


// PrimerSpecificity returns a PrimerSpecificity entity bound to this client.
// Idiomatic usage: client.PrimerSpecificity(nil).List(nil, nil) or
// client.PrimerSpecificity(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) PrimerSpecificity(data map[string]any) SeqbenchMcpEntity {
	return NewPrimerSpecificityEntityFunc(sdk, data)
}


// ProteaseDigestion returns a ProteaseDigestion entity bound to this client.
// Idiomatic usage: client.ProteaseDigestion(nil).List(nil, nil) or
// client.ProteaseDigestion(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ProteaseDigestion(data map[string]any) SeqbenchMcpEntity {
	return NewProteaseDigestionEntityFunc(sdk, data)
}


// ProteinAnnotatePoll returns a ProteinAnnotatePoll entity bound to this client.
// Idiomatic usage: client.ProteinAnnotatePoll(nil).List(nil, nil) or
// client.ProteinAnnotatePoll(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ProteinAnnotatePoll(data map[string]any) SeqbenchMcpEntity {
	return NewProteinAnnotatePollEntityFunc(sdk, data)
}


// ProteinAnnotateSubmit returns a ProteinAnnotateSubmit entity bound to this client.
// Idiomatic usage: client.ProteinAnnotateSubmit(nil).List(nil, nil) or
// client.ProteinAnnotateSubmit(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ProteinAnnotateSubmit(data map[string]any) SeqbenchMcpEntity {
	return NewProteinAnnotateSubmitEntityFunc(sdk, data)
}


// ProteinHydrophobicity returns a ProteinHydrophobicity entity bound to this client.
// Idiomatic usage: client.ProteinHydrophobicity(nil).List(nil, nil) or
// client.ProteinHydrophobicity(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ProteinHydrophobicity(data map[string]any) SeqbenchMcpEntity {
	return NewProteinHydrophobicityEntityFunc(sdk, data)
}


// ProteinProperty returns a ProteinProperty entity bound to this client.
// Idiomatic usage: client.ProteinProperty(nil).List(nil, nil) or
// client.ProteinProperty(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ProteinProperty(data map[string]any) SeqbenchMcpEntity {
	return NewProteinPropertyEntityFunc(sdk, data)
}


// RandomSequence returns a RandomSequence entity bound to this client.
// Idiomatic usage: client.RandomSequence(nil).List(nil, nil) or
// client.RandomSequence(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) RandomSequence(data map[string]any) SeqbenchMcpEntity {
	return NewRandomSequenceEntityFunc(sdk, data)
}


// RestrictionSite returns a RestrictionSite entity bound to this client.
// Idiomatic usage: client.RestrictionSite(nil).List(nil, nil) or
// client.RestrictionSite(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) RestrictionSite(data map[string]any) SeqbenchMcpEntity {
	return NewRestrictionSiteEntityFunc(sdk, data)
}


// ReverseComplement returns a ReverseComplement entity bound to this client.
// Idiomatic usage: client.ReverseComplement(nil).List(nil, nil) or
// client.ReverseComplement(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ReverseComplement(data map[string]any) SeqbenchMcpEntity {
	return NewReverseComplementEntityFunc(sdk, data)
}


// ReverseTranslate returns a ReverseTranslate entity bound to this client.
// Idiomatic usage: client.ReverseTranslate(nil).List(nil, nil) or
// client.ReverseTranslate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) ReverseTranslate(data map[string]any) SeqbenchMcpEntity {
	return NewReverseTranslateEntityFunc(sdk, data)
}


// RnaFold returns a RnaFold entity bound to this client.
// Idiomatic usage: client.RnaFold(nil).List(nil, nil) or
// client.RnaFold(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) RnaFold(data map[string]any) SeqbenchMcpEntity {
	return NewRnaFoldEntityFunc(sdk, data)
}


// SangerVsReference returns a SangerVsReference entity bound to this client.
// Idiomatic usage: client.SangerVsReference(nil).List(nil, nil) or
// client.SangerVsReference(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SangerVsReference(data map[string]any) SeqbenchMcpEntity {
	return NewSangerVsReferenceEntityFunc(sdk, data)
}


// SavePermalink returns a SavePermalink entity bound to this client.
// Idiomatic usage: client.SavePermalink(nil).List(nil, nil) or
// client.SavePermalink(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SavePermalink(data map[string]any) SeqbenchMcpEntity {
	return NewSavePermalinkEntityFunc(sdk, data)
}


// SeqfileStat returns a SeqfileStat entity bound to this client.
// Idiomatic usage: client.SeqfileStat(nil).List(nil, nil) or
// client.SeqfileStat(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SeqfileStat(data map[string]any) SeqbenchMcpEntity {
	return NewSeqfileStatEntityFunc(sdk, data)
}


// SequenceFetch returns a SequenceFetch entity bound to this client.
// Idiomatic usage: client.SequenceFetch(nil).List(nil, nil) or
// client.SequenceFetch(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SequenceFetch(data map[string]any) SeqbenchMcpEntity {
	return NewSequenceFetchEntityFunc(sdk, data)
}


// SequenceFormatConvert returns a SequenceFormatConvert entity bound to this client.
// Idiomatic usage: client.SequenceFormatConvert(nil).List(nil, nil) or
// client.SequenceFormatConvert(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SequenceFormatConvert(data map[string]any) SeqbenchMcpEntity {
	return NewSequenceFormatConvertEntityFunc(sdk, data)
}


// SequenceReport returns a SequenceReport entity bound to this client.
// Idiomatic usage: client.SequenceReport(nil).List(nil, nil) or
// client.SequenceReport(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SequenceReport(data map[string]any) SeqbenchMcpEntity {
	return NewSequenceReportEntityFunc(sdk, data)
}


// SequenceSearch returns a SequenceSearch entity bound to this client.
// Idiomatic usage: client.SequenceSearch(nil).List(nil, nil) or
// client.SequenceSearch(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SequenceSearch(data map[string]any) SeqbenchMcpEntity {
	return NewSequenceSearchEntityFunc(sdk, data)
}


// SequencingReadbackVerify returns a SequencingReadbackVerify entity bound to this client.
// Idiomatic usage: client.SequencingReadbackVerify(nil).List(nil, nil) or
// client.SequencingReadbackVerify(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SequencingReadbackVerify(data map[string]any) SeqbenchMcpEntity {
	return NewSequencingReadbackVerifyEntityFunc(sdk, data)
}


// SessionCreate returns a SessionCreate entity bound to this client.
// Idiomatic usage: client.SessionCreate(nil).List(nil, nil) or
// client.SessionCreate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SessionCreate(data map[string]any) SeqbenchMcpEntity {
	return NewSessionCreateEntityFunc(sdk, data)
}


// SessionGet returns a SessionGet entity bound to this client.
// Idiomatic usage: client.SessionGet(nil).List(nil, nil) or
// client.SessionGet(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SessionGet(data map[string]any) SeqbenchMcpEntity {
	return NewSessionGetEntityFunc(sdk, data)
}


// SessionRun returns a SessionRun entity bound to this client.
// Idiomatic usage: client.SessionRun(nil).List(nil, nil) or
// client.SessionRun(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SessionRun(data map[string]any) SeqbenchMcpEntity {
	return NewSessionRunEntityFunc(sdk, data)
}


// SessionSet returns a SessionSet entity bound to this client.
// Idiomatic usage: client.SessionSet(nil).List(nil, nil) or
// client.SessionSet(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SessionSet(data map[string]any) SeqbenchMcpEntity {
	return NewSessionSetEntityFunc(sdk, data)
}


// SirnaDesign returns a SirnaDesign entity bound to this client.
// Idiomatic usage: client.SirnaDesign(nil).List(nil, nil) or
// client.SirnaDesign(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SirnaDesign(data map[string]any) SeqbenchMcpEntity {
	return NewSirnaDesignEntityFunc(sdk, data)
}


// SiteDirectedMutagenesi returns a SiteDirectedMutagenesi entity bound to this client.
// Idiomatic usage: client.SiteDirectedMutagenesi(nil).List(nil, nil) or
// client.SiteDirectedMutagenesi(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) SiteDirectedMutagenesi(data map[string]any) SeqbenchMcpEntity {
	return NewSiteDirectedMutagenesiEntityFunc(sdk, data)
}


// Translate returns a Translate entity bound to this client.
// Idiomatic usage: client.Translate(nil).List(nil, nil) or
// client.Translate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) Translate(data map[string]any) SeqbenchMcpEntity {
	return NewTranslateEntityFunc(sdk, data)
}


// VariantAnnotate returns a VariantAnnotate entity bound to this client.
// Idiomatic usage: client.VariantAnnotate(nil).List(nil, nil) or
// client.VariantAnnotate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) VariantAnnotate(data map[string]any) SeqbenchMcpEntity {
	return NewVariantAnnotateEntityFunc(sdk, data)
}


// VariantComparator returns a VariantComparator entity bound to this client.
// Idiomatic usage: client.VariantComparator(nil).List(nil, nil) or
// client.VariantComparator(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) VariantComparator(data map[string]any) SeqbenchMcpEntity {
	return NewVariantComparatorEntityFunc(sdk, data)
}


// VerifyAssembly returns a VerifyAssembly entity bound to this client.
// Idiomatic usage: client.VerifyAssembly(nil).List(nil, nil) or
// client.VerifyAssembly(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) VerifyAssembly(data map[string]any) SeqbenchMcpEntity {
	return NewVerifyAssemblyEntityFunc(sdk, data)
}


// VerifyConstruct returns a VerifyConstruct entity bound to this client.
// Idiomatic usage: client.VerifyConstruct(nil).List(nil, nil) or
// client.VerifyConstruct(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) VerifyConstruct(data map[string]any) SeqbenchMcpEntity {
	return NewVerifyConstructEntityFunc(sdk, data)
}


// VirtualGel returns a VirtualGel entity bound to this client.
// Idiomatic usage: client.VirtualGel(nil).List(nil, nil) or
// client.VirtualGel(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) VirtualGel(data map[string]any) SeqbenchMcpEntity {
	return NewVirtualGelEntityFunc(sdk, data)
}


// VolcanoPlotData returns a VolcanoPlotData entity bound to this client.
// Idiomatic usage: client.VolcanoPlotData(nil).List(nil, nil) or
// client.VolcanoPlotData(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) VolcanoPlotData(data map[string]any) SeqbenchMcpEntity {
	return NewVolcanoPlotDataEntityFunc(sdk, data)
}


// WebSearch returns a WebSearch entity bound to this client.
// Idiomatic usage: client.WebSearch(nil).List(nil, nil) or
// client.WebSearch(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *SeqbenchMcpSDK) WebSearch(data map[string]any) SeqbenchMcpEntity {
	return NewWebSearchEntityFunc(sdk, data)
}



func TestSDK(testopts map[string]any, sdkopts map[string]any) *SeqbenchMcpSDK {
	if sdkopts == nil {
		sdkopts = map[string]any{}
	}
	sdkopts = vs.Clone(sdkopts).(map[string]any)

	if testopts == nil {
		testopts = map[string]any{}
	}
	testopts = vs.Clone(testopts).(map[string]any)
	testopts["active"] = true

	vs.SetPath(sdkopts, []any{"feature", "test"}, testopts)

	sdk := NewSeqbenchMcpSDK(sdkopts)
	sdk.Mode = "test"

	return sdk
}
