package reader

import (
	"fmt"
	"io"
	"os"
)

var maxFileSizeMb int64 = 20

func GetReader(filePath string) (r io.Reader, err error) {
	reportFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	fi, err := reportFile.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fi.Size()
	if fileSize == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	if fileSize > maxFileSizeMb*1024*1024 {
		return nil, fmt.Errorf("file is too large(%d bytes), limit: %dMb", fileSize, maxFileSizeMb)
	}
	return reportFile, nil
}
