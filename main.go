package main

import (
	"go.rischmann.fr/goanalyze/promhistogram"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	analyzers := []*analysis.Analyzer{
		promhistogram.Analyzer,
	}

	multichecker.Main(analyzers...)
}
