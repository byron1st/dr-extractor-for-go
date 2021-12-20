package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Relation struct {
	Source   string `json:"source"`
	Target   string `json:"target"`
	Kind     string `json:"kind,omitempty"`
	Location string `json:"location,omitempty"`
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
