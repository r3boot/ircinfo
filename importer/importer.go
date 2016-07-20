package importer

import (
	"github.com/r3boot/rlib/logger"
)

// Instance of the logger, in the 'import' namespace
var Log logger.Log

// Initialization routine for this library.
func Setup(l logger.Log) {
	Log = l

	SetupManager()
}
