package importer

import (
	"github.com/r3boot/ircinfo/storage"
	"regexp"
)

//const RE_JOIN_DATA string = ".*-!- (.*) [~?(.*)@(.*)] has joined"
const RE_JOIN_DATA string = ".*-!- (.*) \\[~?(.*)@(.*)\\] has joined"

var RE_JOIN *regexp.Regexp

// 14:05 -!- r3boot [~r3boot@shell.r3blog.nl] has joined #bit.org

func CompileRegexps() {
	var err error

	if RE_JOIN, err = regexp.Compile(RE_JOIN_DATA); err != nil {
		Log.Fatal("Invalid regexp: " + RE_JOIN_DATA)
	}
}

func MatchRegexp(line string) {
	result := RE_JOIN.FindAllStringSubmatch(line, -1)
	if len(result) > 0 {
		nick := result[0][1]
		user := result[0][2]
		host := result[0][3]
		mask := nick + "!" + user + "@" + host

		storage.AddMask(mask)
		storage.AddMaskToNickname(nick, mask)
		storage.AddMaskToUsername(user, mask)
		storage.AddMaskToHostname(user, mask)
	}
}
