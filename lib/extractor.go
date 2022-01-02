package lib

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/static"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa/ssautil"
)

func ExtractCallgraph(pkgName string) error {
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
		caller := edge.Caller.Func.Pkg.Pkg.Path()
		if !strings.Contains(caller, pkgName) {
			return nil
		}

		pos := prog.Fset.Position(edge.Pos())
		filename := pos.Filename
		if filename == "" {
			return nil
		}

		line := pos.Line
		callee := fmt.Sprintf("%s.%s", edge.Callee.Func.Pkg.Pkg.Path(), edge.Callee.Func.Name())

		relation := Relation{
			Source:   removeCharacters(caller, "(", ")", "*"),
			Target:   removeCharacters(callee, "(", ")", "*"),
			Location: fmt.Sprintf("%s:%d", filename, line),
		}

		if err := relation.Print(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
