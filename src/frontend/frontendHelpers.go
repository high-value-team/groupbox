package frontend

func IsExist(path string) bool {
	if _, ok := _escData[path]; ok {
		return true
	}
	return false
}
