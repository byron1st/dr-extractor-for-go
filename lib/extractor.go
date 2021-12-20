package lib

import (
	"fmt"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/static"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa/ssautil"
)

func ExtractCallgraph(pkgName string) error {
	cfg := &packages.Config{
		Mode:  packages.LoadAllSyntax,
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
		pos := prog.Fset.Position(edge.Pos())
		caller := edge.Caller.String()
		filename := pos.Filename
		line := pos.Line
		callee := edge.Callee.String()

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
