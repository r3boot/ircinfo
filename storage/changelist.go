package storage

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const C_FNAME string = "./data/changelist.csv"

func SaveChangeList() {
	var fd *os.File
	var err error

	fd, err = os.OpenFile(C_FNAME, (os.O_CREATE | os.O_WRONLY), 0644)
	if err != nil {
		Log.Fatal("Failed to create " + C_FNAME + ": " + err.Error())
	}
	defer fd.Close()

	fd.WriteString("# ircinfo-changelist: oldnick,newnick,...\n")
	for oldnick, newnicks := range ChangeList {
		fd.WriteString(oldnick)
		for _, newnick := range newnicks {
			fd.WriteString("," + newnick)
		}
		fd.WriteString("\n")
	}

}

func LoadChangeList() {
	var fs os.FileInfo
	var fd *os.File
	var scanner *bufio.Scanner
	var err error

	if fs, err = os.Stat(C_FNAME); err != nil {
		Log.Warning(C_FNAME + " does not exist, not loading ChangeList")
		return
	}

	if fs.IsDir() {
		Log.Warning(C_FNAME + " is a directory, not loading ChangeList")
		return
	}

	if fd, err = os.Open(C_FNAME); err != nil {
		Log.Fatal("Failed to open " + C_FNAME + ": " + err.Error())
	}
	defer fd.Close()

	ChangeList = make(map[string][]string, L_MAX)
	scanner = bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0:1] == "#" {
			continue
		}

		items := strings.Split(line, ",")
		ChangeList[items[0]] = items[1:]
	}

	Log.Debug("Loaded " + strconv.Itoa(len(ChangeList)) + " items into the ChangeList")
}

func HasNickChange(oldnick string, newnick string) bool {
	if _, ok := ChangeList[oldnick]; !ok {
		// nickname does not exist
		return false
	}

	for _, item := range ChangeList[oldnick] {
		if item == newnick {
			return true
		}
	}
	return false
}

func AddNickToChange(oldnick string, newnick string) bool {
	if HasNickChange(oldnick, newnick) {
		return false
	}

	ChangeList[oldnick] = append(ChangeList[oldnick], newnick)
	return true
}
