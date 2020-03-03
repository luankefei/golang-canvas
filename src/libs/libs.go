package libs

import (
	"io"
	"net/http"
	"os"
)

// Setup to init font, queue & aliyun
func Setup() {
}

// DownloadFile from url and copy to local file
func DownloadFile(filename string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
