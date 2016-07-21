package storage

func HostnameHasMask(hostname string, mask string) bool {
	if _, ok := HostList[hostname]; !ok {
		// hostname does not exist
		return false
	}

	for _, item := range HostList[hostname] {
		if item == mask {
			return true
		}
	}
	return false
}

func AddMaskToHostname(hostname string, mask string) bool {
	if HostnameHasMask(hostname, mask) {
		return false
	}

	HostList[hostname] = append(HostList[hostname], mask)
	return true
}
