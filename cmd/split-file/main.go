package splitfilecmd

import (
	"flag"
	"fmt"
	"github.com/kevinroleke/splitfile"
)

func main() {
	chunks := flag.Int("chunks", 0, "Amount of chunks to split file into.")
    lines := flag.Int("lines", 0, "Amount of lines per chunk.")
    input := flag.String("input", "", "Input file.")
	prefix := flag.String("prefix", "", "Output prefix.")

    flag.Parse()

	if *prefix == "" {
        prefix = *input + "."
    }

	if *lines == 0 && *chunks == 0 {
		fmt.Error("Must supply either --chunks or --lines")
		return
	}

	if *input == "" {
		fmt.Error("Must supply an input file")
		return
	}

	f, err := os.Open(input)
	if err != nil {
		fmt.Error("Input file does not exist, or otherwise is not able to be accessed.")
		return
	}

	if *lines > 0 {
		lc, err = splitfile.LineCounter(f)
		if err != nil {
			fmt.Error("Error with input file")
		}

		chunks = lc/lines
	}

	err = splitfile.SplitFile(chunks, f, prefix)
	if err != nil {
		fmt.Error(err)
	}
}