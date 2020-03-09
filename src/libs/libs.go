package libs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
)

// Setup to init font, queue & aliyun
func Setup() {
	initDir()
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

// changes the current working directory to the named directory
// @see https://golang.org/pkg/os/#Chdir
// @see https://brandur.org/fragments/testing-go-project-root
func initDir() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../main")
	fmt.Println("initDir", dir)
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
