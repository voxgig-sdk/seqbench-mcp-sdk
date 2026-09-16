# SeqBench API

The SeqBench REST API exposes every SeqBench tool as a stateless JSON endpoint: primer design and melting-temperature (Tm) calculation, in-silico PCR, cloning simulation (Gibson Assembly, Golden Gate Assembly, restriction digest), CRISPR guide RNA (gRNA) design, ORF finding, restriction-site mapping, codon optimization, sequence alignment and more. Every tool is callable at `POST /&#123;tool&#125;` with a JSON body matching that tool&#39;s input schema. `GET /` lists all tools and their schemas. `POST /batch` runs one tool over a whole multi-FASTA file, and `POST /workflow` chains several tools into a pipeline over every record. Responses share a common envelope: `&#123; ok, tool, result, gate, provenance &#125;`. The optional `gate` reports typed pass/fail QC checks plus an honest `notChecked` list of what it does not verify. No authentication or API key is required. CORS is open. The same tools are also available over the Model Context Protocol (MCP) at `https://seqbench.com/api/mcp` for AI agents.

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 85 entities and 87 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### AlphafoldLookup

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `accession`: UniProt accession, for example
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### AsoDesign

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `length`: Total gapmer length (nt).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `target`: Nucleotide sequence (raw or FASTA; IUPAC accepted).
- `tool`: The tool slug that ran.

### BaseEditingDesign

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `editor`: Base editor: be3/be4max (CBE, C→T) or abe7.10/abe8e (ABE, A→G).
- `frameStart`: Optional 1-based CDS reading-frame start, to classify each edit&#39;s amino-acid consequence.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `target`: Nucleotide sequence (raw or FASTA; IUPAC accepted).

### Batch

Results: Batch completed (individual rows may carry per-record errors).; Batchable tool list and limits.

SDK operations: `create`, `load`.

Key fields to recognise:

- `args`: Shared tool arguments applied to every record.
- `capped`: True if input exceeded the record limit.
- `input`: Multi-FASTA text or one sequence per line (max ~2,000,000 chars).
- `limit`: Maximum records per call (500).
- `tool`: Tool slug, or an array of slugs (one per step) for a workflow.

### BatchWorkflow

Results: Workflow completed (individual steps may carry per-record errors).; Pipeline-capable tool list and limits.

SDK operations: `create`, `load`.

Key fields to recognise:

- `columns`: Flattened &quot;&lt;step&gt;·&lt;tool&gt;·&lt;key&gt;&quot; column headers.
- `input`: Multi-FASTA text or one sequence per line.
- `limit`: Maximum records per call (200).

### CharacterizeSequence

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `endPrimerLength`: Length of the naive end primers taken from each end.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `maxOrfs`: Maximum number of ORFs to return, longest first.
- `minOrfAa`: Minimum ORF length in amino acids (nucleotide input only).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.

### CloningSimulate

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `armTmTarget`: Target annealing Tm (°C) for primer arms.
- `circular`: Produce a circular product.
- `enzyme`: Type IIS enzyme for Golden Gate (for example
- `enzyme3`: 3′ enzyme (restriction method).
- `enzyme5`: 5′ enzyme (restriction method).

### CodonAdaptationIndex

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `frameStart`: 1-based position to start reading codons.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `rareThreshold`: Relative adaptiveness (w) below this flags a codon as rare.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Coding sequence (DNA/RNA; should start in-frame at ATG).

### CodonOptimize

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `protein`: Protein sequence (one-letter codes).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### ConstructAutofix

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `avoidEnzymes`: Enzyme names whose internal sites should be removed (for example
- `crypticOrfMinAa`: Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged.
- `frameStart`: 1-based nucleotide where the reading frame begins.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `maxPasses`: Repeat full passes until clean or no further progress.

### ConstructQc

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `avoidEnzymes`: Enzyme names whose internal sites should be flagged as errors.
- `crypticOrfMinAa`: Minimum peptide length (aa) for a hidden alternate-frame ORF to be flagged.
- `frameStart`: 1-based nucleotide where the reading frame begins.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `gcHigh`: GC% above this flags a GC-rich window.

### CrisprGrnaDesign

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `minScore`: Only return guides with a heuristic score at least this high (0–100).
- `nuclease`: Nuclease id.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `searchReverseStrand`: Also scan the reverse strand for guides.

### CrisprHdrDonor

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `armLength`: Homology arm length (bp) on each side.
- `blockPam`: When a SpCas9-family guide is supplied and the edit does not already disrupt its PAM, fold a PAM-blocking mutation (silent when a CDS frame is given) into the donor so the edited allele can&#39;t be re-cut.
- `designGenotypingPrimers`: Also design a primer pair (on the original targetSequence) whose product spans the edit site.
- `editEnd`: 1-based inclusive end of the region being replaced; editEnd = editStart-1 denotes a pure insertion with nothing removed.
- `editStart`: 1-based start of the region being replaced.

### CrisprOfftargetCheck

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `maxMismatches`: Mismatches tolerated between the protospacer and a candidate genomic site.
- `nuclease`: Nuclease id, determines the PAM pattern/side required at each candidate site.
- `protospacer`: The guide&#39;s protospacer sequence, 5&#39;→3&#39; (no PAM).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.

### CrossDimer

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequenceA`: First oligo (5&#39;→3&#39;).
- `sequenceB`: Second oligo (5&#39;→3&#39;).
- `tool`: The tool slug that ran.

### DnaMolarity

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `length`: Length in bp (dsDNA) or nt (ssDNA/ssRNA).
- `massNg`: Mass in nanograms.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Optional sequence, overrides length and gives an exact molar mass from base composition.

### DoubleDigest

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `enzymeA`: First enzyme name (for example
- `enzymeB`: Second enzyme name (for example
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### ExportEchoPicklist

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `reactions`: One entry per PCR reaction, up to 96 (a single 96-well plate).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### ExportOpentronsProtocol

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `protocolName`: Optional protocol name (used in the script&#39;s metadata).
- `reactions`: One entry per PCR reaction, up to 96 (a single 96-well plate).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### ExportPlateLayout

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `reactions`: One entry per PCR reaction, up to 96 (a single 96-well plate).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### ExpressionHeatmapCluster

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `clusterCols`: Cluster (reorder) samples.
- `clusterRows`: Cluster (reorder) genes.
- `distanceMetric`: correlation = 1 - Pearson r (the standard expression-heatmap default); euclidean = straight-line distance.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `genes`: Row (gene) labels.

### FastqQcReport

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `input`: FASTQ text: records of an &#39;@id&#39; header, sequence, &#39;+&#39; separator and quality line (four lines each).
- `qualityOffset`: FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### FastqTrim

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `input`: FASTQ text: records of an &#39;@id&#39; header, sequence, &#39;+&#39; separator and quality line (four lines each).
- `minLength`: Reads shorter than this after trimming are dropped.
- `qualityOffset`: FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3-1.7).
- `qualityThreshold`: 3&#39; quality-trim threshold (Phred score).

### FindOrf

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `minAaLength`: Minimum protein length (aa) to report.
- `requireStop`: Only report ORFs terminated by a stop codon.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Nucleotide sequence (raw or FASTA; IUPAC accepted).

### FormatSequence

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `convert`: DNA→RNA (T→U) or RNA→DNA (U→T).
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `reverse`: Reverse the sequence (no complement).
- `sequence`: Nucleotide sequence (raw or FASTA; IUPAC accepted).

### FunctionalEnrichment

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `background`: Custom background/universe gene symbols.
- `collections`: Which term collections to test.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `genes`: Query gene symbols (human, for example
- `maxTermSize`: Skip terms/pathways with more than this many background genes (matches clusterProfiler&#39;s default).

### GcContent

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Nucleotide sequence (raw or FASTA; IUPAC accepted).
- `tool`: The tool slug that ran.

### GeneDossier

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `gene`: A human gene symbol (&quot;TP53&quot;) or Ensembl gene ID (&quot;ENSG00000141510&quot;).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### GeneExpression

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `gene`: A human gene symbol (&quot;TP53&quot;) or Ensembl gene ID (&quot;ENSG00000141510&quot;).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### GeneModel

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `gene`: A human gene symbol (&quot;TP53&quot;) or Ensembl gene ID (&quot;ENSG00000141510&quot;).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### GoldenGateFidelity

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `compareToNamedSet`: Also score this published reference set (see namedSetsAvailable in the output) alongside your candidate set, for comparison.
- `dataset`: Which real ligation dataset to score against, generic T4 ligase, or an enzyme-specific one-pot dataset if that matches your actual digestion enzyme.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `overhangs`: The candidate 4-base overhangs for one assembly (for example
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.

### HgvsConvert

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.
- `variant`: A full HGVS &quot;c.&quot; variant description: &quot;&lt;accession or gene symbol&gt;:c.&lt;edit&gt;&quot;, for example

### IdMapPoll

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### IdMapSubmit

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `from`: Source id type: &quot;Gene_Name&quot;, &quot;Ensembl&quot;, &quot;GeneID&quot;, &quot;RefSeq_Protein&quot;, or &quot;UniProtKB_AC-ID&quot;.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `ids`: The ids to map, up to 1000 (for example
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `taxId`: NCBI taxonomy id to disambiguate a gene symbol (only used when from=&quot;Gene_Name&quot;).

### InSilicoPcr

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `circular`: Treat the template as circular (plasmid).
- `forwardPrimer`: Primer 1, 5&#39;→3&#39;.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `maxMismatches`: Mismatches tolerated per primer.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.

### KaspPrimerDesign

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `addSecondaryMismatch`: Engineer the internal ARMS destabilising mismatch near the 3&#39; end.
- `alleleA`: First allele (single base), gets the FAM tail.
- `alleleB`: Second allele (single base), gets the HEX tail.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `maxAmplicon`: Maximum amplicon length for the common reverse primer.

### ListTool

Results: The API index.

SDK operations: `load`.

### MeltingTemperature

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `dntpMM`: Total [dNTP] (mM), chelates Mg2+.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `mgMM`: Divalent cation [Mg2+] (mM).
- `naMM`: Monovalent cation [Na+]/[K+] (mM).
- `oligoNM`: Total strand concentration (nM).

### MotifFinder

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `maxMismatches`: Maximum allowed mismatches per match.
- `motif`: Query motif; IUPAC ambiguity codes (R Y S W K M B D H V N) allowed.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `searchReverseStrand`: Also search the reverse strand.

### MultipleSequenceAlignment

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `input`: Two or more sequences in multi-FASTA format (&gt;name / sequence).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### OligoAnalysi

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `dntpMM`: Total [dNTP] (mM), chelates Mg2+.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `mgMM`: Divalent cation [Mg2+] (mM).
- `naMM`: Monovalent cation [Na+]/[K+] (mM).
- `oligoNM`: Total strand concentration (nM).

### OrthologMap

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sourceSpecies`: Ensembl species slug the symbols belong to (for example
- `symbols`: Gene symbols to look up, up to 50 (for example
- `targetSpecies`: Ensembl species slug to find homologs in (for example

### PairwiseAlignment

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gap`: Linear gap penalty (per gap position).
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `match`: Match score.
- `mismatch`: Mismatch penalty.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.

### ParseGenbank

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `text`: A GenBank flat file (LOCUS … FEATURES … ORIGIN … //).
- `tool`: The tool slug that ran.

### ParseSangerTrace

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `fileBase64`: The binary ABIF (.ab1 / .abi) trace file, base64-encoded.
- `fileName`: Optional original file name (echoed back).
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### PlasmidAnnotate

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Nucleotide sequence (raw or FASTA; IUPAC accepted).
- `tool`: The tool slug that ran.

### PlasmidDeepAnnotate

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `circular`: Treat the sequence as a circular plasmid (vs.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Nucleotide sequence (raw or FASTA; IUPAC accepted).
- `tool`: The tool slug that ran.

### PlasmidFullReport

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `circular`: Treat the query as a circular molecule (most plasmids are).
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Nucleotide sequence (raw or FASTA; IUPAC accepted).
- `tool`: The tool slug that ran.

### PlasmidIdentify

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `circular`: Treat the query as a circular molecule (most plasmids are).
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Nucleotide sequence (raw or FASTA; IUPAC accepted).
- `tool`: The tool slug that ran.

### PrimeEditingDesign

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `editEnd`: 1-based inclusive end of the region being changed.
- `editStart`: 1-based inclusive start of the region being changed.
- `frameStart`: Optional 1-based CDS reading-frame start, used only to annotate whether a PAM-blocking mutation would be silent.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `insertedSeq`: Replacement bases (forward strand).

### PrimeEditingTwinDesign

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `newSequence`: New sequence (forward strand) to install in place of [replaceStart, replaceEnd].
- `overlapLength`: Length (bp) of the shared overlap built into both pegRNAs&#39; 3&#39; flaps where they meet and anneal.
- `pbsLength`: Optional preferred PBS length to highlight; a full 8-17 nt sweep is always returned.
- `replaceEnd`: 1-based inclusive end of the region being replaced/deleted.

### PrimerDesign

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `dntpMM`: Total [dNTP] (mM), chelates Mg2+.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `maxReturn`: Number of best pairs to return.
- `mgMM`: Divalent cation [Mg2+] (mM).
- `naMM`: Monovalent cation [Na+]/[K+] (mM).

### PrimerSpecificity

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `forwardPrimer`: Forward primer, 5&#39;→3&#39;.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `maxMismatches`: Mismatches tolerated per primer against a reference genome.
- `maxProductLength`: Ignore candidate off-target products longer than this (bp), a search-window cap, not a biological claim.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.

### ProteaseDigestion

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `maxMass`: Optional upper bound on neutral monoisotopic mass (Da).
- `maxPeptides`: Cap on the number of returned peptides.
- `minMass`: Optional lower bound on neutral monoisotopic mass (Da).
- `missedCleavages`: Allowed missed internal cleavages (0–2).

### ProteinAnnotatePoll

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### ProteinAnnotateSubmit

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `appl`: Restrict to one member database (for example
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `goterms`: Include GO-term cross-references.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Protein sequence, one-letter code (FASTA header, if any, is stripped).

### ProteinHydrophobicity

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `scale`: Amino-acid scale.
- `sequence`: Protein sequence (one-letter amino-acid codes; non-AA characters ignored).
- `tool`: The tool slug that ran.

### ProteinProperty

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `chargeStep`: pH step for the net-charge titration curve (0–14).
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Protein sequence (one-letter amino-acid codes; non-AA characters ignored).
- `tool`: The tool slug that ran.

### RandomSequence

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `gcContent`: Target GC percentage 0..100 (dna/rna only); omit for uniform.
- `length`: Number of residues to generate.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### RestrictionSite

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `enzymes`: Enzyme names to scan; omit to scan all curated enzymes.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Nucleotide sequence (raw or FASTA; IUPAC accepted).
- `tool`: The tool slug that ran.

### ReverseComplement

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Nucleotide sequence (raw or FASTA; IUPAC accepted).
- `tool`: The tool slug that ran.

### ReverseTranslate

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `organism`: Codon-usage host (ignored in degenerate mode).
- `protein`: Protein sequence (one-letter codes; * for stop).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### RnaFold

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Nucleotide sequence (raw or FASTA; IUPAC accepted).
- `tool`: The tool slug that ran.

### SangerVsReference

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `fileBase64`: The binary ABIF (.ab1 / .abi) trace file, base64-encoded.
- `fileName`: Optional original file name (echoed back).
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `minCoverage`: Fraction of the reference the read must span before a PASS is meaningful; below this the verdict is &#39;ambiguous_low_coverage&#39; regardless of identity.
- `read`: Sanger read as FASTA or raw text (alternative to uploading an ABIF trace).

### SavePermalink

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `args`: Arguments for that tool, exactly as you would pass to it directly.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### SeqfileStat

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `input`: FASTA or FASTQ text (raw sequence is treated as single-record FASTA).
- `qualityOffset`: FASTQ Phred ASCII offset (33 = Sanger/Illumina 1.8+, 64 = Illumina 1.3–1.7).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### SequenceFetch

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `accession`: GenBank/RefSeq accession (for example
- `db`: Database to query; auto-detects from the accession format.
- `format`: Output format (GenBank is only available for NCBI accessions, UniProt and Ensembl are FASTA-only).
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.

### SequenceFormatConvert

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `from`: Input format; &#39;auto&#39; sniffs it from the first meaningful line.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `input`: A FASTA or GenBank record to convert.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `to`: Output format.

### SequenceReport

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `endPrimerLength`: Length of the naive end primers taken from each end.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `maxOrfs`: Maximum number of ORFs to return, longest first.
- `minOrfAa`: Minimum ORF length in amino acids.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.

### SequenceSearch

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `gene`: Gene symbol/name, for example
- `maxResults`: Up to 20.
- `organism`: Organism name, for example
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.

### SequencingReadbackVerify

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `minSupportingReads`: Minimum number of reads agreeing on a variant position for it to count as a consensus (candidate real) variant rather than single-read noise.
- `reads`: Raw reads in FASTA or FASTQ format (auto-detected).
- `reference`: The claimed/expected reference sequence.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.

### SessionCreate

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `entries`: Initial named entries, for example
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### SessionGet

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `names`: Only return these entries; omit to return all of them.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### SessionRun

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `args`: Additional literal arguments, merged with the ones resolved from the session.
- `fromSession`: Map of &#123; toolArgName: sessionEntryName &#125;, resolves each named tool argument from the session before running.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### SessionSet

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `entries`: Named entries to add/overwrite, for example
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### SirnaDesign

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `minReynolds`: Minimum Reynolds score (0–8) to keep; falls back to best-ranked if none qualify.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `shRnaLoop`: Loop sequence used when assembling the shRNA cassette.
- `target`: Nucleotide sequence (raw or FASTA; IUPAC accepted).

### SiteDirectedMutagenesi

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `armTmTarget`: Target Tm (°C) for each template-binding arm.
- `dntpMM`: Total [dNTP] (mM), chelates Mg2+.
- `editKind`: Edit at the nucleotide or amino-acid level.
- `frameStart`: 1-based position of the first base of codon 1 (editKind=&#39;aa&#39;).
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).

### Translate

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `sequence`: Nucleotide sequence (raw or FASTA; IUPAC accepted).
- `toStop`: Stop at the first stop codon.
- `tool`: The tool slug that ran.

### VariantAnnotate

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `assembly`: Genome build for rsID/chrom-pos-ref-alt/genomic-HGVS lookups (MyVariant.info&#39;s native default is hg19).
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.
- `variant`: An rsID (&quot;rs1042522&quot;), chrom:pos:ref:alt (&quot;17:7676154:G:C&quot;, single-base substitutions only), genomic HGVS (&quot;chr17:g.7676154G&gt;C&quot; or &quot;17:g.7676154G&gt;C&quot;), or transcript HGVS c.

### VariantComparator

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `coding`: Treat as a coding sequence and report amino-acid effects.
- `frameStart`: 1-based reading-frame start (used when coding is true).
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `query`: Query / variant sequence (raw or FASTA).
- `reference`: Reference / wild-type sequence (raw or FASTA).

### VerifyAssembly

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `armTmTarget`: Target annealing Tm (°C) for primer arms.
- `circular`: Treat the product/claimed construct as circular (most plasmids are).
- `claimedConstruct`: The sequence you claim you ended up with.
- `coding`: Report amino-acid effects of any mismatch, assuming claimedConstruct is (or contains) a coding sequence.
- `enzyme`: Type IIS enzyme for Golden Gate.

### VerifyConstruct

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `claimedConstruct`: The final sequence claimed to have been built.
- `expectedFrameStart`: 1-based position in claimedConstruct where the intended reading frame begins.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `insertForwardPrimer`: Forward primer used to amplify the insert, 5&#39;→3&#39;.
- `insertReversePrimer`: Reverse primer used to amplify the insert, 5&#39;→3&#39;.

### VirtualGel

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `circular`: Treat the sequence as circular (plasmid).
- `enzymes`: Enzyme names to digest with.
- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `ladder`: DNA ladder to plot alongside the sample lane.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.

### VolcanoPlotData

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `rows`: Differential expression rows, one per gene.
- `tool`: The tool slug that ran.

### WebSearch

Results: Tool ran successfully.

SDK operations: `create`.

Key fields to recognise:

- `gate`: Typed QC verdict, or null when the tool defines no gate or the call lacked gating inputs (for example no targetTm).
- `max_results`: Maximum number of results to return (default 5, max 10).
- `query`: The search query.
- `result`: Tool-specific output object. Its shape depends on the tool, call `GET /&#123;tool&#125;` or see the docs for each tool&#39;s fields.
- `tool`: The tool slug that ran.

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| AlphafoldLookup | `create` | `POST /alphafold_lookup` | Not required |
| AsoDesign | `create` | `POST /aso_design` | Not required |
| BaseEditingDesign | `create` | `POST /base_editing_design` | Not required |
| Batch | `create` | `POST /batch` | Not required |
| Batch | `load` | `GET /batch` | Not required |
| BatchWorkflow | `create` | `POST /workflow` | Not required |
| BatchWorkflow | `load` | `GET /workflow` | Not required |
| CharacterizeSequence | `create` | `POST /characterize_sequence` | Not required |
| CloningSimulate | `create` | `POST /cloning_simulate` | Not required |
| CodonAdaptationIndex | `create` | `POST /codon_adaptation_index` | Not required |
| CodonOptimize | `create` | `POST /codon_optimize` | Not required |
| ConstructAutofix | `create` | `POST /construct_autofix` | Not required |
| ConstructQc | `create` | `POST /construct_qc` | Not required |
| CrisprGrnaDesign | `create` | `POST /crispr_grna_design` | Not required |
| CrisprHdrDonor | `create` | `POST /crispr_hdr_donor` | Not required |
| CrisprOfftargetCheck | `create` | `POST /crispr_offtarget_check` | Not required |
| CrossDimer | `create` | `POST /cross_dimer` | Not required |
| DnaMolarity | `create` | `POST /dna_molarity` | Not required |
| DoubleDigest | `create` | `POST /double_digest` | Not required |
| ExportEchoPicklist | `create` | `POST /export_echo_picklist` | Not required |
| ExportOpentronsProtocol | `create` | `POST /export_opentrons_protocol` | Not required |
| ExportPlateLayout | `create` | `POST /export_plate_layout` | Not required |
| ExpressionHeatmapCluster | `create` | `POST /expression_heatmap_cluster` | Not required |
| FastqQcReport | `create` | `POST /fastq_qc_report` | Not required |
| FastqTrim | `create` | `POST /fastq_trim` | Not required |
| FindOrf | `create` | `POST /find_orfs` | Not required |
| FormatSequence | `create` | `POST /format_sequence` | Not required |
| FunctionalEnrichment | `create` | `POST /functional_enrichment` | Not required |
| GcContent | `create` | `POST /gc_content` | Not required |
| GeneDossier | `create` | `POST /gene_dossier` | Not required |
| GeneExpression | `create` | `POST /gene_expression` | Not required |
| GeneModel | `create` | `POST /gene_model` | Not required |
| GoldenGateFidelity | `create` | `POST /golden_gate_fidelity` | Not required |
| HgvsConvert | `create` | `POST /hgvs_convert` | Not required |
| IdMapPoll | `create` | `POST /id_map_poll` | Not required |
| IdMapSubmit | `create` | `POST /id_map_submit` | Not required |
| InSilicoPcr | `create` | `POST /in_silico_pcr` | Not required |
| KaspPrimerDesign | `create` | `POST /kasp_primer_design` | Not required |
| ListTool | `load` | `GET /` | Not required |
| MeltingTemperature | `create` | `POST /melting_temperature` | Not required |
| MotifFinder | `create` | `POST /motif_finder` | Not required |
| MultipleSequenceAlignment | `create` | `POST /multiple_sequence_alignment` | Not required |
| OligoAnalysi | `create` | `POST /oligo_analysis` | Not required |
| OrthologMap | `create` | `POST /ortholog_map` | Not required |
| PairwiseAlignment | `create` | `POST /pairwise_alignment` | Not required |
| ParseGenbank | `create` | `POST /parse_genbank` | Not required |
| ParseSangerTrace | `create` | `POST /parse_sanger_trace` | Not required |
| PlasmidAnnotate | `create` | `POST /plasmid_annotate` | Not required |
| PlasmidDeepAnnotate | `create` | `POST /plasmid_deep_annotate` | Not required |
| PlasmidFullReport | `create` | `POST /plasmid_full_report` | Not required |
| PlasmidIdentify | `create` | `POST /plasmid_identify` | Not required |
| PrimeEditingDesign | `create` | `POST /prime_editing_design` | Not required |
| PrimeEditingTwinDesign | `create` | `POST /prime_editing_twin_design` | Not required |
| PrimerDesign | `create` | `POST /primer_design` | Not required |
| PrimerSpecificity | `create` | `POST /primer_specificity` | Not required |
| ProteaseDigestion | `create` | `POST /protease_digestion` | Not required |
| ProteinAnnotatePoll | `create` | `POST /protein_annotate_poll` | Not required |
| ProteinAnnotateSubmit | `create` | `POST /protein_annotate_submit` | Not required |
| ProteinHydrophobicity | `create` | `POST /protein_hydrophobicity` | Not required |
| ProteinProperty | `create` | `POST /protein_properties` | Not required |
| RandomSequence | `create` | `POST /random_sequence` | Not required |
| RestrictionSite | `create` | `POST /restriction_sites` | Not required |
| ReverseComplement | `create` | `POST /reverse_complement` | Not required |
| ReverseTranslate | `create` | `POST /reverse_translate` | Not required |
| RnaFold | `create` | `POST /rna_fold` | Not required |
| SangerVsReference | `create` | `POST /sanger_vs_reference` | Not required |
| SavePermalink | `create` | `POST /save_permalink` | Not required |
| SeqfileStat | `create` | `POST /seqfile_stats` | Not required |
| SequenceFetch | `create` | `POST /sequence_fetch` | Not required |
| SequenceFormatConvert | `create` | `POST /sequence_format_convert` | Not required |
| SequenceReport | `create` | `POST /sequence_report` | Not required |
| SequenceSearch | `create` | `POST /sequence_search` | Not required |
| SequencingReadbackVerify | `create` | `POST /sequencing_readback_verify` | Not required |
| SessionCreate | `create` | `POST /session_create` | Not required |
| SessionGet | `create` | `POST /session_get` | Not required |
| SessionRun | `create` | `POST /session_run` | Not required |
| SessionSet | `create` | `POST /session_set` | Not required |
| SirnaDesign | `create` | `POST /sirna_design` | Not required |
| SiteDirectedMutagenesi | `create` | `POST /site_directed_mutagenesis` | Not required |
| Translate | `create` | `POST /translate` | Not required |
| VariantAnnotate | `create` | `POST /variant_annotate` | Not required |
| VariantComparator | `create` | `POST /variant_comparator` | Not required |
| VerifyAssembly | `create` | `POST /verify_assembly` | Not required |
| VerifyConstruct | `create` | `POST /verify_construct` | Not required |
| VirtualGel | `create` | `POST /virtual_gel` | Not required |
| VolcanoPlotData | `create` | `POST /volcano_plot_data` | Not required |
| WebSearch | `create` | `POST /web_search` | Not required |

## Connect to the API

- Production: `https://seqbench.com/api/v1`

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

A read request without required parameters or authentication is `GET /batch`. For example:

```sh
curl --fail-with-body --silent --show-error 'https://seqbench.com/api/v1/batch'
```

Inspect the response using the Batch reference. This checks the public route; authenticated operations need their own credentials and request data.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| Golang | `go/` | Build from source |
| Lua | `lua/` | Build from source |
| PHP | `php/` | Build from source |
| Python | `py/` | Build from source |
| Ruby | `rb/` | Build from source |
| TypeScript | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### Go CLI

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### Go MCP server

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `seqbench-mcp_list`: List records for an entity. No active entity supports this operation.
- `seqbench-mcp_load`: Load one record for an entity. Supported entities: `batch`, `batch__workflow`, `list_tool`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- `ratelimit`: Client-side rate limiting via a token bucket
- `retry`: Automatic retry of transient failures with exponential backoff
- `test`: In-memory mock transport for testing without a live server
- `timeout`: Per-request timeout with transport abort

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the first-call guide for the setup sequence.
- Read the authentication guide before using protected routes.
- Use the API reference for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

