package Methods

func SplitWithNewlines(str string) []string {
	word := ""
	var argLettersDivided []string // Slice to hold the split words and newline markers
	count := 0                     // to track the number of newlines

	// Loop through each character in the input string
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			// If a word has been accumulated, append it to the result
			if word != "" {
				argLettersDivided = append(argLettersDivided, word)
				word = "" // Reset word for the next accumulation
			}
			// Append a newline marker if this isn't the first newline encountered
			if count != 0 {
				argLettersDivided = append(argLettersDivided, "\\n")
			}
			count++ // Increment the newline counter

		} else {
			// Accumulate characters into the word
			word += string(str[i])
		}
	}
	// If there's a word left at the end, append it to the result
	if word != "" {
		argLettersDivided = append(argLettersDivided, word)
	}
	return argLettersDivided
}
