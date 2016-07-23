package importer

import (
	"github.com/r3boot/ircinfo/storage"
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
		ReadAndParse(logfile)
	}

	ml_size := len(storage.MaskList)
	nl_size := len(storage.NickList)
	ul_size := len(storage.UserList)
	hl_size := len(storage.HostList)
	cl_size := len(storage.ChangeList)

	Log.Debug("Found " + strconv.Itoa(ml_size) + " masks")
	Log.Debug("Found " + strconv.Itoa(nl_size) + " nicknames")
	Log.Debug("Found " + strconv.Itoa(ul_size) + " users")
	Log.Debug("Found " + strconv.Itoa(hl_size) + " hosts")
	Log.Debug("Found " + strconv.Itoa(cl_size) + " nickchanges")
}
