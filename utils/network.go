package utils

// IsNetError checks the status code to see if it is a network error
func IsNetError(resCode int) bool {
	switch resCode {
	case 400:
		return true
	case 401:
		return true
	case 404:
		return true
	default:
		return false
	}
}
