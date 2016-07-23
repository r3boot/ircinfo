package storage

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const N_FNAME string = "./data/nicklist.csv"

func SaveNickList() {
	var fd *os.File
	var err error

	fd, err = os.OpenFile(N_FNAME, (os.O_CREATE | os.O_WRONLY), 0644)
	if err != nil {
		Log.Fatal("Failed to create " + N_FNAME + ": " + err.Error())
	}
	defer fd.Close()

	fd.WriteString("# ircinfo-nicklist: nickname,mask,...\n")
	for nick, masks := range NickList {
		fd.WriteString(nick)
		for _, mask := range masks {
			fd.WriteString("," + mask)
		}
		fd.WriteString("\n")
	}
}

func LoadNickList() {
	var fs os.FileInfo
	var fd *os.File
	var scanner *bufio.Scanner
	var err error

	if fs, err = os.Stat(N_FNAME); err != nil {
		Log.Warning(N_FNAME + " does not exist, not loading NickList")
		return
	}

	if fs.IsDir() {
		Log.Warning(N_FNAME + " is a directory, not loading NickList")
		return
	}

	if fd, err = os.Open(N_FNAME); err != nil {
		Log.Fatal("Failed to open " + N_FNAME + ": " + err.Error())
	}
	defer fd.Close()

	NickList = make(map[string][]string, L_MAX)
	scanner = bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0:1] == "#" {
			continue
		}

		items := strings.Split(line, ",")
		NickList[items[0]] = items[1:]
	}

	Log.Debug("Loaded " + strconv.Itoa(len(NickList)) + " items into the NickList")
}

func NicknameHasMask(nickname string, mask string) bool {
	if _, ok := NickList[nickname]; !ok {
		// nickname does not exist
		return false
	}

	for _, item := range NickList[nickname] {
		if item == mask {
			return true
		}
	}
	return false
}

func AddMaskToNickname(nickname string, mask string) bool {
	if NicknameHasMask(nickname, mask) {
		return false
	}

	NickList[nickname] = append(NickList[nickname], mask)
	return true
}
