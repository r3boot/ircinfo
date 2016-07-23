package importer

import (
	"github.com/r3boot/ircinfo/storage"
	"regexp"
)

const RE_JOIN_DATA string = ".*-!- (.*) \\[~?(.*)@(.*)\\] has joined"
const RE_NICKCHANGE_DATA string = ".*-!- (.*) is now known as (.*)"

var RE_JOIN *regexp.Regexp
var RE_NICKCHANGE *regexp.Regexp

// 14:05 -!- r3boot [~r3boot@shell.r3blog.nl] has joined #bit.org

func CompileRegexps() {
	var err error

	if RE_JOIN, err = regexp.Compile(RE_JOIN_DATA); err != nil {
		Log.Fatal("Invalid regexp: " + RE_JOIN_DATA)
	}

	if RE_NICKCHANGE, err = regexp.Compile(RE_NICKCHANGE_DATA); err != nil {
		Log.Fatal("Invalid regexp: " + RE_NICKCHANGE_DATA)
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
		return
	}

	result = RE_NICKCHANGE.FindAllStringSubmatch(line, -1)
	if len(result) > 0 {
		oldnick := result[0][1]
		newnick := result[0][2]

		storage.AddNickToChange(oldnick, newnick)
	}
}
