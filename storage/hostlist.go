package storage

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const H_FNAME string = "./data/hostlist.csv"

func SaveHostList() {
	var fd *os.File
	var err error

	fd, err = os.OpenFile(H_FNAME, (os.O_CREATE | os.O_WRONLY), 0644)
	if err != nil {
		Log.Fatal("Failed to create " + H_FNAME + ": " + err.Error())
	}
	defer fd.Close()

	fd.WriteString("# ircinfo-hostlist: hostname,mask,...\n")
	for hostname, masks := range HostList {
		fd.WriteString(hostname)
		for _, mask := range masks {
			fd.WriteString("," + mask)
		}
		fd.WriteString("\n")
	}
}

func LoadHostList() {
	var fs os.FileInfo
	var fd *os.File
	var scanner *bufio.Scanner
	var err error

	if fs, err = os.Stat(H_FNAME); err != nil {
		Log.Warning(H_FNAME + " does not exist, not loading HostList")
		return
	}

	if fs.IsDir() {
		Log.Warning(H_FNAME + " is a directory, not loading HostList")
		return
	}

	if fd, err = os.Open(H_FNAME); err != nil {
		Log.Fatal("Failed to open " + H_FNAME + ": " + err.Error())
	}
	defer fd.Close()

	HostList = make(map[string][]string, L_MAX)
	scanner = bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0:1] == "#" {
			continue
		}

		items := strings.Split(line, ",")
		HostList[items[0]] = items[1:]
	}

	Log.Debug("Loaded " + strconv.Itoa(len(HostList)) + " items into the HostList")
}

func HostnameHasMask(hostname string, mask string) bool {
	if _, ok := HostList[hostname]; !ok {
		// hostname does not exist
		return false
	}

	for _, item := range HostList[hostname] {
		if item == mask {
			return true
		}
	}
	return false
}

func AddMaskToHostname(hostname string, mask string) bool {
	if HostnameHasMask(hostname, mask) {
		return false
	}

	HostList[hostname] = append(HostList[hostname], mask)
	return true
}
