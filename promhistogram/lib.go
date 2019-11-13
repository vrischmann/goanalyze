// Package promhistogram contains an analyzer which detects calls to
// prometheus.NewHistogram with a HistogramOpts with the Buckets field not set.
//
// Usually when you create a histogram you don't want the default buckets and while writing
// code it's easy to forget.
package promhistogram

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "promhistogram",
	Doc:      "check that the Buckets field is set when creating a Prometheus histogram",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

var (
// promPackageIdent = &ast.Ident{Name: "prometheus"}
)

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	types := []ast.Node{
		(*ast.CompositeLit)(nil),
	}

	inspect.Preorder(types, func(n ast.Node) {
		lit := n.(*ast.CompositeLit)

		// check if our literal is exactly prometheus.HistogramOpts

		sel, ok := lit.Type.(*ast.SelectorExpr)
		if !ok {
			return
		}
		exprIdent, ok := sel.X.(*ast.Ident)
		if !ok {
			return
		}

		if exprIdent.Name != "prometheus" || sel.Sel.Name != "HistogramOpts" {
			return
		}

		// it is, now check the different composite elements

		// Prometheus requires only the Name field according to the documentation:
		// https://godoc.org/github.com/prometheus/client_golang/prometheus#HistogramOpts
		//
		// In our case we also assume that the Buckets field must be set.

		var (
			hasName    bool
			hasBuckets bool
		)

		for _, el := range lit.Elts {
			kve := el.(*ast.KeyValueExpr)
			if identIs(kve.Key, "Name") {
				hasName = true
			}
			if identIs(kve.Key, "Buckets") {
				hasBuckets = true
			}
		}

		if !hasName {
			pass.ReportRangef(lit, "literal HistogramOpts doesn't have the Name field defined")
		}
		if !hasBuckets {
			pass.ReportRangef(lit, "literal HistogramOpts doesn't have the Buckets field defined")
		}
	})

	return nil, nil
}

func identIs(ident ast.Expr, name string) bool {
	v, ok := ident.(*ast.Ident)
	if !ok {
		return false
	}
	return v.Name == name
}
