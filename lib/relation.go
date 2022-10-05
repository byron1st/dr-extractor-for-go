package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
)

type DepedencyRelationFromSCToEL struct {
	Source         string `json:"source"`
	SourceMetadata string `json:"sourceMetadata"`
	Target         string `json:"target"`
	TargetMetadata string `json:"targetMetadata"`

	prog               *ssa.Program
	edge               *callgraph.Edge
	sourceCodePkgNames []string
}

func NewDependencyRelation(prog *ssa.Program, edge *callgraph.Edge, sourceCodePkgNames []string) DepedencyRelationFromSCToEL {
	return DepedencyRelationFromSCToEL{
		prog:               prog,
		edge:               edge,
		sourceCodePkgNames: sourceCodePkgNames,
	}
}

func (dr DepedencyRelationFromSCToEL) CheckIfDRFromSCToEL() bool {
	sourceModule := dr.edge.Caller.Func.Pkg.Pkg.Path()
	targetModule := dr.edge.Callee.Func.Pkg.Pkg.Path()

	return dr.isSourceCodePkg(sourceModule) && !dr.isSourceCodePkg(targetModule)
}

func (dr DepedencyRelationFromSCToEL) isSourceCodePkg(pkg string) bool {
	for _, sourceCodePkgName := range dr.sourceCodePkgNames {
		if strings.Contains(pkg, sourceCodePkgName) {
			return true
		}
	}

	return false
}

func (dr DepedencyRelationFromSCToEL) Print() error {
	sourceModule := dr.edge.Caller.Func.Pkg.Pkg.Path()
	targetModule := dr.edge.Callee.Func.Pkg.Pkg.Path()
	pos := dr.prog.Fset.Position(dr.edge.Pos())
	line := pos.Line
	filename := pos.Filename

	dr.Source = fmt.Sprintf("%s.%s", sourceModule, dr.edge.Caller.Func.Name())
	dr.SourceMetadata = fmt.Sprintf("%s:%d", filename, line)
	dr.Target = fmt.Sprintf("%s.%s", targetModule, dr.edge.Callee.Func.Name())
	dr.TargetMetadata = ""

	relationBytes, err := json.Marshal(dr)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintln(os.Stdout, string(relationBytes)); err != nil {
		return err
	}

	return nil
}

// func removeCharacters(str string, chars ...string) string {
// 	removed := str
// 	for _, c := range chars {
// 		removed = strings.Replace(removed, c, "", -1)
// 	}

// 	return removed
// }
