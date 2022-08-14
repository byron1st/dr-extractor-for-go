package lib

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/static"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa/ssautil"
)

func ExtractCallgraph(pkgName string, baseName string) error {
	cfg := &packages.Config{
		Mode: packages.NeedDeps |
			packages.NeedSyntax |
			packages.NeedTypesInfo |
			packages.NeedTypes |
			packages.NeedTypesSizes |
			packages.NeedImports |
			packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles,
		Tests: false,
	}
	initial, err := packages.Load(cfg, pkgName)
	if err != nil {
		return err
	}
	if packages.PrintErrors(initial) > 0 {
		return fmt.Errorf("packages contain errors")
	}

	prog, _ := ssautil.AllPackages(initial, 0)
	prog.Build()

	cg := static.CallGraph(prog)
	cg.DeleteSyntheticNodes()

	if err := callgraph.GraphVisitEdges(cg, func(edge *callgraph.Edge) error {
		sourceModule := edge.Caller.Func.Pkg.Pkg.Path()
		targetModule := edge.Callee.Func.Pkg.Pkg.Path()
		targetFunc := edge.Callee.Func.Name()
		pos := prog.Fset.Position(edge.Pos())
		line := pos.Line
		filename := pos.Filename

		if !strings.Contains(sourceModule, pkgName) ||
			strings.Contains(targetModule, pkgName) ||
			(baseName != "" && strings.Contains(targetModule, baseName)) ||
			filename == "" {
			return nil
		}

		if !strings.Contains(sourceModule, pkgName) ||
			strings.Contains(targetModule, pkgName) ||
			(baseName != "" && strings.Contains(targetModule, baseName)) ||
			filename == "" {
			return nil
		}

		relation := NewRelation(sourceModule, filename, line, targetModule, targetFunc)
		if err := relation.Print(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
