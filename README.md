# SplitFile

Split large files into parts. 

## Module Usage

```go
package main

import (
	"fmt"
	"github.com/kevinroleke/splitfile"
)

func main() {
	chunks := 16
	prefix := "test.txt_split."

	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	splitfile.SplitFile(chunks, f, prefix)
}
```

## CLI Usage

`./splitfile --input test.txt --chunks 16 --prefix abc`
`./splitfile --input test.txt --lines 100`
