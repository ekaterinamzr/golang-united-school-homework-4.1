package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}

	output = string(runes)

	return output
}
