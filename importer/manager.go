package importer

import (
	"strconv"
)

type ManagerStruct struct {
}

var Manager *ManagerStruct

const D_LOGFILES_MAX int = 65535

func SetupManager() {
	Manager = new(ManagerStruct)
}

func (m *ManagerStruct) Import(directory string) {
	var all_logs []string

	// Gather a list of all logfiles we need to parse
	all_logs = Crawl(directory)

	Log.Debug("Found " + strconv.Itoa(len(all_logs)) + " logfiles to parse")

	// Import all found logs
	for _, logfile := range all_logs {
		Parse(logfile)
	}
}
