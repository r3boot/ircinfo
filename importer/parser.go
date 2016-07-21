package importer

import (
	"bufio"
	"os"
)

/*
 * This function parses a filename using regular expressions, and extracts
 * information which can be used as data objects
 */
func ReadAndParse(filename string) {
	var fd *os.File
	var err error

	Log.Debug("Parsing " + filename)

	if fd, err = os.Open(filename); err != nil {
		Log.Warning("Cannot open '" + filename + "': " + err.Error())
		return
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		MatchRegexp(scanner.Text())
	}

	return
}
