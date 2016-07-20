package importer

import (
	"io/ioutil"
	"os"
)

/*
 * Recursive function which crawls a directory and gathers filenames of logs
 * which we need to index
 */
func Crawl(directory string) (result []string) {
	var files []os.FileInfo
	var err error

	Log.Debug("Importing all logs from " + directory)
	if files, err = ioutil.ReadDir(directory); err != nil {
		Log.Fatal("Crawl", err.Error())
	}

	for _, file := range files {
		if file.IsDir() {
			result = append(result, Crawl(directory+"/"+file.Name())...)
		} else {
			result = append(result, directory+"/"+file.Name())
		}
	}

	return
}
