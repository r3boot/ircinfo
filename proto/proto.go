package proto

import (
	"github.com/r3boot/rlib/logger"
)

var Log logger.Log

func Setup(l logger.Log) {
	Log = l
}
