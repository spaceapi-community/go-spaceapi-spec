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
		_ = generator.CreateTypes()

		w, err := os.Create("./v" + version + "/spec.go")
		generate.Output(w, generator, "spaceapiStruct")
	}

	// structs should be formatted nicely and gofmt doesn't provide a lib. Not going to implement it here
	cmd := exec.Command("gofmt", "-w", "./")
	cmd.Run()
}
