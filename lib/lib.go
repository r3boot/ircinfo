package lib

import (
	"github.com/r3boot/rlib/logger"
)

// Instance of the logger, in the 'lib' namespace
var Log logger.Log

// Initialization routine for this library.
func Setup(l logger.Log) {
	Log = l
}
