package renderer

import (
	"log"
	"os"
	temp "text/template"
)

func Render(template string, variables interface{}) {
	type Inventory struct {
		Material string
		Count    uint
	}

	tmpl, err := temp.New("output").Parse(template)

	if err != nil {
		log.Fatalf("cannot create template: %v", err)
	}

	err = tmpl.Execute(os.Stdout, variables)

	if err != nil {
		log.Fatalf("cannot render template: %v", err)
	}
}
