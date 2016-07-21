package storage

func HasMask(mask string) (result bool) {
	_, ok := MaskList[mask]
	return ok
}

func AddMask(mask string) bool {
	MaskList[mask] = mask
	return true
}
