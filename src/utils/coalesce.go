package utils

func Coalesce(str, defaultValue string) string {
	if str == "" {
		return defaultValue
	}

	return str
}
