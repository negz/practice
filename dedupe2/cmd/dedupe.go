package main

import (
	"os"
	"path/filepath"

	"github.com/negz/practice/dedupe2"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		app  = kingpin.New(filepath.Base(os.Args[0]), "Deduplicate files!").DefaultEnvars()
		root = app.Arg("root", "Root of path in which to deduplicate").ExistingDir()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))
	dedupe2.Dedupe(*root)
}
