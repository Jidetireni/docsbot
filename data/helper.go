package data

import (
	"io"
	"os"
)

func loadPDFFile(path string) (io.ReaderAt, int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}

	info, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, 0, err
	}

	return file, info.Size(), nil
}
