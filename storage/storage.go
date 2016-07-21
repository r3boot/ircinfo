package storage

import (
	"github.com/r3boot/rlib/logger"
)

const L_MAX int64 = 65536

var Log logger.Log

var MaskList map[string]string
var NickList map[string][]string
var UserList map[string][]string
var HostList map[string][]string
var ChangeList map[string][]string

func Setup(l logger.Log) {
	Log = l

	MaskList = make(map[string]string, L_MAX)
	NickList = make(map[string][]string, L_MAX)
	UserList = make(map[string][]string, L_MAX)
	HostList = make(map[string][]string, L_MAX)
	ChangeList = make(map[string][]string, L_MAX)
}
