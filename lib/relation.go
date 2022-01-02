package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type RelationByTarget struct {
	Language       string `json:"language"`
	TargetModule   string `json:"targetModule"`
	TargetFunc     string `json:"targetFunc,omitempty"`
	SourceModule   string `json:"sourceModule"`
	SourceLocation string `json:"sourceLocation,omitempty"`
}

func (r RelationByTarget) Print() error {
	relationBytes, err := json.Marshal(r)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintln(os.Stdout, string(relationBytes)); err != nil {
		return err
	}

	return nil
}

func removeCharacters(str string, chars ...string) string {
	removed := str
	for _, c := range chars {
		removed = strings.Replace(removed, c, "", -1)
	}

	return removed
}
