package main

import (
	"golang.org/x/tools/go/analysis"
	"go.rischmann.fr/goanalyze/promhistogram"
)

func main() {
	analyzers := []*analysis.Analyzer{
		promhistogram.Analyzer,
	}

	multichecker.Main(analyzers)
}
