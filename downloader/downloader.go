package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Downloader struct {
	httpClient *http.Client
}

func NewDownloader() *Downloader {
	return &Downloader{&http.Client{}}
}

func (d *Downloader) Download(url string, filePath string, fileName string) error {
	resp, err := d.httpClient.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath.Join(filePath, fileName))

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
