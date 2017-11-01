package main

import (
	"fmt"
	"github.com/gidsi/generate"
	"github.com/gidsi/generate/jsonschema"
	validator "github.com/gidsi/go-spaceapi-validator"
	"io"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	for version, schema := range validator.SpaceApiSchemas {
		// dirty fix because of a bug in the generator
		// https://github.com/a-h/generate/issues/12
		oldType := `"type": [
            "boolean",
            "null"
          ]`
		newType := `"type": "boolean"`

		// dirty fix for missing schema attribute
		cleanSchema := strings.Replace(
			strings.Replace(
				schema,
				oldType,
				newType,
				1,
			),
			"{",
			`{ "$schema": "http://json-schema.org/schema#",`,
			1,
		)

		parsedSchema, err := jsonschema.Parse(cleanSchema)
		if err != nil {
			panic("couldnt parse schema!")
		}

		generator := generate.New(parsedSchema)
		structs, _ := generator.CreateStructs()
		os.Mkdir("v"+version, 0755)
		w, err := os.Create("./v" + version + "/spec.go")

		if err == nil {
			output(w, structs)
		}
	}

	// structs should be formatted nicely and gofmt doesn't provide a lib. Not going to implement it by myself
	cmd := exec.Command("gofmt", "-w", "./")
	cmd.Run()
}

// everything below is shamelessly copied from
// https://github.com/a-h/generate/blob/master/cmd/schema-generate/main.go
func getOrderedFieldNames(m map[string]generate.Field) []string {
	keys := make([]string, len(m))
	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}
	sort.Strings(keys)
	return keys
}

func getOrderedStructNames(m map[string]generate.Struct) []string {
	keys := make([]string, len(m))
	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}
	sort.Strings(keys)
	return keys
}

func output(w io.Writer, structs map[string]generate.Struct) {
	fmt.Fprintf(w, "package spaceapiStruct\n")

	for _, k := range getOrderedStructNames(structs) {
		s := structs[k]

		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "// %s %s\n", s.Name, s.Description)
		fmt.Fprintf(w, "type %s struct {\n", s.Name)

		for _, fieldKey := range getOrderedFieldNames(s.Fields) {
			f := s.Fields[fieldKey]

			// Only apply omitempty if the field is not required.
			omitempty := ",omitempty"
			if f.Required {
				omitempty = ""
			}

			fmt.Fprintf(w, "  %s %s `json:\"%s%s\"`\n", f.Name, f.Type, f.JSONName, omitempty)
		}

		fmt.Fprintln(w, "}")
	}
}
