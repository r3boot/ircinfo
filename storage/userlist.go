package storage

func UsernameHasMask(username string, mask string) bool {
	if _, ok := UserList[username]; !ok {
		// username does not exist
		return false
	}

	for _, item := range UserList[username] {
		if item == mask {
			return true
		}
	}
	return false
}

func AddMaskToUsername(username string, mask string) bool {
	if UsernameHasMask(username, mask) {
		return false
	}

	UserList[username] = append(UserList[username], mask)
	return true
}
