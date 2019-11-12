// Package promhistogram contains an analyzer which detects calls to
// prometheus.NewHistogram with a HistogramOpts with the Buckets field not set.
//
// Usually when you create a histogram you don't want the default buckets and while writing
// code it's easy to forget.
package promhistogram

import "golang.org/x/tools/go/analysis"

var Analyzer = analysis.Analyzer{
	Name: "promhistogram",
	Doc:  "check that the Buckets field is set when creating a Prometheus histogram",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	panic("not implemented")
}
