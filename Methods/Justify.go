package Methods

// Justify calculates the space allocation for a given word based on alignment and width
func Justify(align string, width int, word string, fileLines []string) (int, int) {

	wordWidth := 0
	// Calculate the width of the word based on the character codes
	for i := 0; i < len(word); i++ {
		// Calculate the index for fileLines based on ASCII value of the character
		start := ((int((word)[i]) - 32) * 9) + (1)
		wordWidth += len(fileLines[start]) // Accumulate the width of each character
	}
	// Calculate the remaining space to be filled with spaces
	spacesRemain := width - wordWidth

	// Determine the allocation of spaces based on the alignment
	switch align {
	case "left":
		return 0, spacesRemain
	case "right":
		return spacesRemain, 0
	case "center":
		return spacesRemain / 2, spacesRemain / 2
	default:
		return 0, 0
	}
}
