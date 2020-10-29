package utils

func CheckErr(error error) bool {
	if error != nil {
		panic(error)
		return true
	}
	return false
}
