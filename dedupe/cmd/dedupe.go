package main

import (
	"os"
	"path/filepath"

	"github.com/negz/practice/dedupe"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		app  = kingpin.New(filepath.Base(os.Args[0]), "Deduplicates file").DefaultEnvars()
		root = app.Arg("root", "root of tree in which to dedupe").ExistingDir()
	)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	kingpin.FatalIfError(dedupe.Dedupe(*root), "cannot deduplicate")
}
