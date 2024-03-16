package main

import (
	"fmt"
	"io"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	"encore.app/store"
)

// Define the models to generate migrations for.
var models = []any{
	&store.Organization{},
	&store.Config{},
	&store.Reviewer{},
	&store.Message{},
	&store.Pullrequest{},
}

func main() {
	stmts, err := gormschema.New("postgres").Load(models...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
