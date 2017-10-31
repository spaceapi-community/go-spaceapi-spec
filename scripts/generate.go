package main

import (
	"github.com/a-h/generate"
	"github.com/a-h/generate/jsonschema"
	validator "github.com/gidsi/go-spaceapi-validator"
	"fmt"
	"strings"
	"os"
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
		output(version, structs)
	}
}

func output(filename string, structs map[string]generate.Struct) {
	os.Mkdir("v" + filename, 0755)
	w, err := os.Create("./v" + filename + "/spec.go")

	if err == nil {
		fmt.Fprintf(w, "package spaceapi\n")

		for _, s := range structs {

			fmt.Fprintln(w, "")
			fmt.Fprintf(w, "type %s struct {\n", s.Name)

			for _, f := range s.Fields {
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
}