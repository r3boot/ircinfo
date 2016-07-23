package main

import (
	"flag"
	"github.com/r3boot/ircinfo/importer"
	"github.com/r3boot/ircinfo/lib"
	"github.com/r3boot/ircinfo/proto"
	"github.com/r3boot/ircinfo/storage"
	"github.com/r3boot/rlib/logger"
	"os"
)

// Information about this application
const APP_NAME string = "ircinfo"
const APP_VERS string = "0.1"
const APP_DESC string = "An tool which indexes your irc logs, and gets a picture of who-is-who in channels with nickchangers"

// Default options for commandline parameters
const D_DEBUG bool = false

// Commandline parameters
var debug = flag.Bool("D", D_DEBUG, "Enable debug output")

// Application related variables
const OP_IMPORT string = "import"
const OP_SEARCH string = "search"

var Log logger.Log

/*
 * This function is called once during program startup. It is used to
 * perform various tasks needed before this application can start running
 */
func init() {
	var err error

	// Parse command-line flags
	flag.Parse()

	// Initialize logging framework
	Log.UseDebug = *debug
	Log.UseVerbose = *debug
	Log.UseTimestamp = true
	Log.Debug("Logging initialized")

	// Parse non-flag commandline parameters
	if len(flag.Args()) == 0 {
		Log.Fatal("Nothing to do!")
	}

	switch flag.Arg(0) {
	case OP_IMPORT:
		{
			if len(flag.Args()) == 1 {
				Log.Fatal("Need a directory to import")
			}

			if _, err = os.Stat(flag.Arg(1)); err != nil {
				Log.Fatal("Unable to import directory")
			}
		}
	case OP_SEARCH:
		{
			Log.Debug("search")
		}
	default:
		{
			Log.Fatal("Unknown operation")
		}
	}

	// Initialize various submodules
	lib.Setup(Log)
	proto.Setup(Log)
	importer.Setup(Log)
	storage.Setup(Log)

}

/*
 * This function is called when the program starts running, after init is called
 * and it performs the actual work this program is supposed to do
 */
func main() {
	switch flag.Arg(0) {
	case OP_IMPORT:
		{
			directory := flag.Arg(1)
			importer.Manager.Import(directory)
			storage.SaveLists()
		}
	case OP_SEARCH:
		{
			storage.LoadLists()
		}
	default:
		{
			Log.Warning("Function not implemented")
		}
	}
}
