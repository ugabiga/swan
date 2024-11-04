package main

import (
	"fmt"
	"io"
	"log"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
)

var models = []any{}

func main() {
	stmts, err := gormschema.New("sqlite").Load(
		models...,
	)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		if err != nil {
			log.Printf("failed to write to stderr: %v", err)
			return
		}
		os.Exit(1)
	}

	_, err = io.WriteString(os.Stdout, stmts)
	if err != nil {
		log.Printf("failed to write to stdout: %v", err)
		return
	}
}
