package utils

// Contains returns bool if search exists within the string slice
func Contains(search string, list []string) bool {
	for index := range list {
		if list[index] == search {
			return true
		}
	}
	return false
}

// ContainsByte returns bool if search exists within the byte slice
func ContainsByte(search byte, list []byte) bool {
	for index := range list {
		if list[index] == search {
			return true
		}
	}
	return false
}
