package storage

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const U_FNAME string = "./data/userlist.csv"

func SaveUserList() {
	var fd *os.File
	var err error

	fd, err = os.OpenFile(U_FNAME, (os.O_CREATE | os.O_WRONLY), 0644)
	if err != nil {
		Log.Fatal("Failed to create " + U_FNAME + ": " + err.Error())
	}
	defer fd.Close()

	fd.WriteString("# ircinfo-userlist: user,mask,...\n")
	for user, masks := range UserList {
		fd.WriteString(user)
		for _, mask := range masks {
			fd.WriteString("," + mask)
		}
		fd.WriteString("\n")
	}
}

func LoadUserList() {
	var fs os.FileInfo
	var fd *os.File
	var scanner *bufio.Scanner
	var err error

	if fs, err = os.Stat(U_FNAME); err != nil {
		Log.Warning(U_FNAME + " does not exist, not loading UserList")
		return
	}

	if fs.IsDir() {
		Log.Warning(U_FNAME + " is a directory, not loading UserList")
		return
	}

	if fd, err = os.Open(U_FNAME); err != nil {
		Log.Fatal("Failed to open " + U_FNAME + ": " + err.Error())
	}
	defer fd.Close()

	UserList = make(map[string][]string, L_MAX)
	scanner = bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0:1] == "#" {
			continue
		}

		items := strings.Split(line, ",")
		UserList[items[0]] = items[1:]
	}

	Log.Debug("Loaded " + strconv.Itoa(len(UserList)) + " items into the UserList")
}

func UsernameHasMask(username string, mask string) bool {
	if _, ok := UserList[username]; !ok {
		// username does not exist
		return false
	}

	for _, item := range UserList[username] {
		if item == mask {
			return true
		}
	}
	return false
}

func AddMaskToUsername(username string, mask string) bool {
	if UsernameHasMask(username, mask) {
		return false
	}

	UserList[username] = append(UserList[username], mask)
	return true
}
