package utils

func TruncString(s string, maxLen int) string {
	if maxLen > len(s) {
		maxLen = len(s) - 1
	}
	return s[0:maxLen]
}
