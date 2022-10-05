package lib

import (
	"fmt"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/static"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa/ssautil"
)

func ExtractCallgraph(mainPkgName string, sourceCodePkgNames []string) error {
	pkgs, err := loadPkgs(mainPkgName)
	if err != nil {
		return err
	}

	prog, _ := ssautil.AllPackages(pkgs, 0)
	prog.Build()

	cg := static.CallGraph(prog)
	cg.DeleteSyntheticNodes()

	if err := callgraph.GraphVisitEdges(cg, func(edge *callgraph.Edge) error {
		dr := NewDependencyRelation(prog, edge, sourceCodePkgNames)
		if !dr.CheckIfDRFromSCToEL() {
			return nil
		}

		if err := dr.Print(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

var cfg = &packages.Config{
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

func loadPkgs(mainPkgName string) ([]*packages.Package, error) {
	pkgs, err := packages.Load(cfg, mainPkgName)
	if err != nil {
		return nil, err
	} else if packages.PrintErrors(pkgs) > 0 {
		return nil, fmt.Errorf("packages contain errors")
	}

	return pkgs, nil
}
