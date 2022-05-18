package reverse_string

func ReverseString(input string) string {
	output := []rune(input)
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}

	return string(output)
}
