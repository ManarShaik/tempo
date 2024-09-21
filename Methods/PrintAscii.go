package Methods

import (
	"fmt"
	"strings"
)

const (
	Reset = "\033[0m"
)

// PrintAscii prints ASCII art for the provided letters with specified formatting
func PrintAscii(argLettersDivided, fileLines []string, subString string, color *string, width int, align string) {
	tempColor := Reset
	var start int = 0 // Variable to hold the index for accessing fileLines
	count := 0        // Counter for managing substring highlighting

	// Loop through each letter in the divided argument
	for i := 0; i < len(argLettersDivided); i++ {

		// Check for newline escape sequence and print a newline if found
		if argLettersDivided[i] == "\\n" {
			fmt.Println()
			continue
		}

		// Get spaces based on alignment
		leftSpacess, rightSpace := Justify(align, width, argLettersDivided[i], fileLines)

		// Loop to print each line of the ASCII art (assuming 8 lines)
		for j := 0; j < 8; j++ {

			// Print left leading spaces
			fmt.Print(strings.Repeat(" ", leftSpacess))

			for k := 0; k < len(argLettersDivided[i]); k++ {

				// Check if the current character matches the first character of the substring
				if (argLettersDivided[i])[k] == subString[0] && i+len(subString)-1 <= len(argLettersDivided[i])-1 {

					// Validate if the substring matches
					if ValidateSubString(subString, (argLettersDivided[i])[k:k+len(subString)]) {
						tempColor = *color     // Set the color for the substring
						count = len(subString) // Set count for substring length
					}

				}

				// Calculate the starting index in fileLines based on the character's ASCII value
				start = ((int((argLettersDivided[i])[k]) - 32) * 9) + (j + 1)

				// Print the corresponding ASCII art line
				// Print space
				if (argLettersDivided[i])[k] == ' ' {
					fmt.Print((fileLines[start]))
					continue
				}
				// Print the character in the designated color
				fmt.Print(tempColor + fileLines[start])
				start = 0 // Reset start for next character

				// Reset color after printing the substring
				if count == 1 {
					tempColor = Reset
				}
				count-- // starting count from zero

			}
			// Print right trailing spaces
			fmt.Print(strings.Repeat(" ", rightSpace))
			fmt.Print("\n") // Move to the next line

		}
	}
}
