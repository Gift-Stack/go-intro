package stringutil

func ReverseString(value string) string {
	var reversedString string

	for i := len(value) - 1; i >= 0; i-- {
		reversedString += string(value[i])
	}

	return reversedString
}
