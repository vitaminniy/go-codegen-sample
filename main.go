package main

import (
	"flag"
	"log"

	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	log.SetFlags(0)

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("no filename provided")
	}

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	doc, err := loader.LoadFromFile(args[0])
	if err != nil {
		log.Fatalf("could not parse [%s]: %v", args[0], err)
	}

	for p, pi := range doc.Paths {
		if pi.Get == nil {
			log.Printf("path [%s] has no 'GET' operation", p)
			continue
		}

		log.Println("parameters:")
		for _, param := range pi.Get.Parameters {
			if param == nil || param.Value == nil {
				continue
			}

			log.Printf("\t- ref: %s; name: %s", param.Ref, param.Value.Name)
		}
	}
}
