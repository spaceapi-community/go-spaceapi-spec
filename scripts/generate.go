package main

import (
	"github.com/a-h/generate"
	"github.com/spaceapi-community/go-spaceapi-validator"
	"net/url"
	"os"
	"os/exec"
)

func main() {
	for version, schema := range spaceapivalidator.SpaceAPISchemas {
		schemaUrl, _ := url.Parse("https://example.com/")

		parsedSchema, err := generate.Parse(schema, schemaUrl)
		if err != nil {
			panic("couldnt parse schema!")
		}

		generator := generate.New(parsedSchema)
		err = generator.CreateTypes()
		if err != nil {
			panic("couldnt create types")
		}

		w, err := os.Create("./v" + version + "/spec.go")
		if err != nil {
			panic("couldnt write spec file")
		}
		generate.Output(w, generator, "spaceapistruct")
	}

	// structs should be formatted nicely and gofmt doesn't provide a lib. Not going to implement it here
	cmd := exec.Command("gofmt", "-w", "./")
	err := cmd.Run()
	if err != nil {
		panic("couldnt fmt file")
	}
}
