package splitfile

import (
	"os"
	"bufio"
	"io"
	"bytes"
	"strconv"
)

func Split(chunks int, f *os.File, prefix string) error {
	lineCount, err := LineCounter(f)
	if err != nil {
		return err
	}

	f.Seek(0, io.SeekStart)

	var (
		chunkWriters []*bufio.Writer
		chunkFiles []*os.File
		chunkInd int = 0
		i int = 0
		sc *bufio.Scanner = bufio.NewScanner(f)
		chunkSize int = lineCount / chunks
		remainder int = lineCount % chunks
	)

	if remainder > 0 {
		chunks++
	}

	for ind := 0; ind < chunks; ind++ {
		file, err := os.OpenFile(prefix + strconv.Itoa(ind), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		datawriter := bufio.NewWriter(file)

		chunkFiles = append(chunkFiles, file)
		chunkWriters = append(chunkWriters, datawriter)
	}

	for sc.Scan() {
        i++
        chunkWriters[chunkInd].WriteString(sc.Text() + "\n")

		if i % 100 == 0 {
			chunkWriters[chunkInd].Flush()
		}

		if i >= chunkSize {
			chunkWriters[chunkInd].Flush()
			chunkFiles[chunkInd].Close()

			chunkInd++
			i = 0
			
			if chunkInd >= chunks {
				return nil
			}
		}
    }

	if remainder > 0 {
		chunkWriters[chunkInd].Flush()
		chunkFiles[chunkInd].Close()
	}

	return nil
}

// https://stackoverflow.com/a/52153000
func LineCounter(r io.Reader) (int, error) {
    var count int
    const lineBreak = '\n'

    buf := make([]byte, bufio.MaxScanTokenSize)

    for {
        bufferSize, err := r.Read(buf)
        if err != nil && err != io.EOF {
            return 0, err
        }

        var buffPosition int
        for {
            i := bytes.IndexByte(buf[buffPosition:], lineBreak)
            if i == -1 || bufferSize == buffPosition {
                break
            }
            buffPosition += i + 1
            count++
        }
        if err == io.EOF {
            break
        }
    }

    return count, nil
}