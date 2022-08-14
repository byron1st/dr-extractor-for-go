package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Relation struct {
	SourceModule Module `json:"sourceModule"`
	TargetModule Module `json:"targetModule"`
}

type Module struct {
	ID       string `json:"id"`
	Location string `json:"location"`
}

func NewRelation(sourceModule string, filename string, line int, targetModule string, targetFunc string) Relation {
	return Relation{
		SourceModule: Module{
			ID:       removeCharacters(sourceModule, "(", ")", "*"),
			Location: fmt.Sprintf("%s:%d", filename, line),
		},
		TargetModule: Module{
			ID: fmt.Sprintf("%s.%s", removeCharacters(targetModule, "(", ")", "*"), removeCharacters(targetFunc, "(", ")", "*")),
		},
	}
}

func (r Relation) Print() error {
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
