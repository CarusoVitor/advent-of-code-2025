package filehandling

import (
	"bufio"
	"io"
	"os"
)

func OpenFile(path string) (io.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ExtractSliceNewLine(rd io.Reader) []string {
	scanner := bufio.NewScanner(rd)
	lines := make([]string, 0, 512)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ExtractSliceSep(rd io.Reader, sep byte, removeSep bool) ([]string, error) {
	reader := bufio.NewReader(rd)
	itens := make([]string, 0, 512)

	var err error = nil
	var item string
	for err != io.EOF {
		item, err = reader.ReadString(sep)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if removeSep && err != io.EOF {
			item = item[:len(item)-1]
		}
		itens = append(itens, item)
	}

	return itens, nil
}
