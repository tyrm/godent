package util

// ContainsString will return true if a string is found in a given group of strings.
func ContainsString(stack []string, needle string) bool {
	for _, s := range stack {
		if s == needle {
			return true
		}
	}

	return false
}
