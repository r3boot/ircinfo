package storage

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const M_FNAME string = "./data/masklist.csv"

func SaveMaskList() {
	var fd *os.File
	var err error

	fd, err = os.OpenFile(M_FNAME, (os.O_CREATE | os.O_WRONLY), 0644)
	if err != nil {
		Log.Fatal("Failed to create " + M_FNAME + ": " + err.Error())
	}
	defer fd.Close()

	fd.WriteString("# ircinfo-masklist: mask\n")
	for _, mask := range MaskList {
		fd.WriteString(mask + "\n")
	}
}

func LoadMaskList() {
	var fs os.FileInfo
	var fd *os.File
	var scanner *bufio.Scanner
	var err error

	if fs, err = os.Stat(M_FNAME); err != nil {
		Log.Warning(M_FNAME + " does not exist, not loading MaskList")
		return
	}

	if fs.IsDir() {
		Log.Warning(M_FNAME + " is a directory, not loading MaskList")
		return
	}

	if fd, err = os.Open(M_FNAME); err != nil {
		Log.Fatal("Failed to open " + M_FNAME + ": " + err.Error())
	}
	defer fd.Close()

	MaskList = make(map[string]string, L_MAX)
	scanner = bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0:1] == "#" {
			continue
		}

		items := strings.Split(line, ",")
		MaskList[items[0]] = items[0]
	}

	Log.Debug("Loaded " + strconv.Itoa(len(MaskList)) + " items into the MaskList")
}

func HasMask(mask string) (result bool) {
	_, ok := MaskList[mask]
	return ok
}

func AddMask(mask string) bool {
	MaskList[mask] = mask
	return true
}
