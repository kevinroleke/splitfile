package main

import (
	"flag"
	"log"
	"os"
	"io"
	"github.com/kevinroleke/splitfile"
)

func main() {
	chunks := flag.Int("chunks", 0, "Amount of chunks to split file into.")
    lines := flag.Int("lines", 0, "Amount of lines per chunk.")
    input := flag.String("input", "", "Input file.")
	prefix := flag.String("prefix", "", "Output prefix.")

    flag.Parse()

	var (
		f_prefix string
		f_chunks int
	)

	if *prefix == "" {
        f_prefix = *input + "."
    } else {
		f_prefix = *prefix
	}

	if *lines == 0 && *chunks == 0 {
		log.Fatal("Must supply either --chunks or --lines")
		return
	}

	if *input == "" {
		log.Fatal("Must supply an input file")
		return
	}

	f, err := os.Open(*input)
	if err != nil {
		log.Fatal("Input file does not exist, or otherwise is not able to be accessed.")
		return
	}

	if *lines > 0 {
		lc, err := splitfile.LineCounter(f)
		if err != nil {
			log.Fatal("Error with input file")
		}

		f.Seek(0, io.SeekStart)

		f_chunks = lc / *lines
	} else {
		f_chunks = *chunks
	}

	err = splitfile.Split(f_chunks, f, f_prefix)
	if err != nil {
		log.Fatal(err)
	}
}
