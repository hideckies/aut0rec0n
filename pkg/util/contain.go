package util

// Check if the string array contains the string
func StrArrContains(sarr []string, s string) bool {
	for _, v := range sarr {
		if v == s {
			return true
		}
	}
	return false
}
