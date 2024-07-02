package output

func IsItAllNL(result string) bool {
	for _, char := range result {
		if char != '\n' {
			return false
		}
	}
	return true
}
