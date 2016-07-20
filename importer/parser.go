package importer

import (
	"os"
)

/*
 * This function parses a filename using regular expressions, and extracts
 * information which can be used as data objects
 */
func Parse(filename string) {
	var fs os.FileInfo
	var fd *os.File
	var size int
	var data []byte
	var err error

	Log.Debug("Parsing " + filename)

	if fs, err = os.Stat(filename); err != nil {
		Log.Warning("Unable to stat() logfile '" + filename + "': " + err.Error())
		return
	}

	// Sanity check
	if fs.IsDir() {
		Log.Warning("'" + filename + "' is a directory")
		return
	}

	data = make([]byte, fs.Size())

	if fd, err = os.Open(filename); err != nil {
		Log.Warning("Cannot open '" + filename + "': " + err.Error())
		return
	}

	if size, err = fd.Read(data); err != nil {
		Log.Warning("Failed to read '" + filename + "': " + err.Error())
		return
	}

	if int64(size) != fs.Size() {
		Log.Fatal("Failed to read '" + filename + "': size mismatch")
		return
	}
}
