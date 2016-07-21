package storage

func NicknameHasMask(nickname string, mask string) bool {
	if _, ok := NickList[nickname]; !ok {
		// nickname does not exist
		return false
	}

	for _, item := range NickList[nickname] {
		if item == mask {
			return true
		}
	}
	return false
}

func AddMaskToNickname(nickname string, mask string) bool {
	if NicknameHasMask(nickname, mask) {
		return false
	}

	NickList[nickname] = append(NickList[nickname], mask)
	return true
}
