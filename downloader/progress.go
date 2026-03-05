package downloader

import (
	"fmt"
	"io"
)

const BytesInMegabyte = 1_048_576

type ProgressBar struct {
	writer      io.Writer
	fullSize    int64
	currentSize int64
}

func NewProgressBar(writer io.Writer, fullSize int64) *ProgressBar {
	return &ProgressBar{writer, fullSize, 0}
}

func (pb *ProgressBar) Write(p []byte) (n int, err error) {
	noOfBytes, err := pb.writer.Write(p)
	if err != nil {
		return 0, err
	}

	pb.currentSize += int64(noOfBytes)

	displayHumanReadableProgress(pb.currentSize, pb.fullSize)

	return noOfBytes, nil
}

func displayHumanReadableProgress(cur int64, full int64) {
	base := 1024
	baseReadable := "KB"

	if full > BytesInMegabyte {
		base = BytesInMegabyte
		baseReadable = "MB"
	}

	curReadable := cur / int64(base)
	fullReadable := full / int64(base)

	percentage := float64(cur) / float64(full) * 100

	fmt.Printf("\r %d %s/%d %s %.0f%%", curReadable, baseReadable, fullReadable, baseReadable, percentage)
}
