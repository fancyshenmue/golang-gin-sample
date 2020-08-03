package validation

func CheckSringInSlice(s string, stringSlice []string) bool {
	for _, v := range stringSlice {
		if v == s {
			return true
		}
	}

	return false
}
