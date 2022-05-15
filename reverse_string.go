package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	rune := []rune(input)
	for i := len(rune) - 1; i >= 0; i-- {
		output += string(rune[i])
	}
	return output
}
