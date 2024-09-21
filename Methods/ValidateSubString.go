package Methods

// ValidateSubString checks if the beginning of 'str' matches the 'subString'
func ValidateSubString(subString, str string) bool {

	// Loop through each character of 'subString'
	for i := 0; i < len(str); i++ {

		// Check if the characters of 'subString' match those of 'str' at the same position
		if subString[i] != str[i] {
			return false // Return false if any character does not match
		}
	}
	return true
}
