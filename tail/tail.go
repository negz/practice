package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// Seek back to the start of the first line we want to output.
func seekBack(f io.ReadSeeker, lines int) error {
	for newlines, offset := 0, -1; newlines < lines; offset-- {
		// Seek back offset from end of file.
		newPos, err := f.Seek(int64(offset), io.SeekEnd)
		if err != nil {
			return err
		}

		// We're at the start of the file. Don't seek back further.
		if newPos == 0 {
			return nil
		}

		b := make([]byte, 1)
		if _, err := f.Read(b); err != nil {
			return err
		}

		if b[0] == byte('\n') {
			newlines++
		}
	}
	return nil
}

func getLines(name string, lines int) ([][]byte, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err = seekBack(f, lines); err != nil {
		return nil, err
	}

	tail, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return bytes.Split(tail, []byte("\n")), nil
}

func outputLines(lines [][]byte) error {
	for _, line := range lines {
		if _, err := os.Stdout.Write(line); err != nil {
			return err
		}
		if _, err := os.Stdout.Write([]byte("\n")); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var (
		app   = kingpin.New(filepath.Base(os.Args[0]), "Tail a file!")
		path  = app.Arg("path", "Path to the file to tail.").ExistingFile()
		lines = app.Flag("lines", "Number of lines to tail.").Short('n').Int()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	l, err := getLines(*path, *lines)
	kingpin.FatalIfError(err, "cannot get last %v lines", *lines)
	kingpin.FatalIfError(outputLines(l), "cannot output last %v lines", *lines)
}
