package Methods

import (
	"fmt"
	"strings"
)

// ValidateArg checks if the provided arguments meet specific validation criteria
func ValidateArg(args []string) bool {

	if len(args) < 1 || len(args) > 3 {
		return false // Return false if the argument count is invalid
	} else {

		// Check if the last argument is one of the valid banner names
		if args[len(args)-1] != "standard" && args[len(args)-1] != "shadow" && args[len(args)-1] != "thinkertoy" {

			// Handle cases where the last argument is not a valid banner
			if len(args) == 2 && !strings.HasSuffix(args[1], ".txt") && !strings.Contains(args[1], args[0]) {
				// For two arguments, check if the second doesn't end with ".txt" and doesn't contain the first
				fmt.Println("subString dosen't exist")
				return false
			}
			if len(args) == 3 && strings.HasSuffix(args[2], ".txt") && !strings.Contains(args[1], args[0]) {
				// For three arguments, check if the third ends with ".txt" and the second doesn't contain the first
				fmt.Println("subString dosen't exist")
				return false
			}
		}
	}

	// Check if any argument is an empty string
	for i := 0; i < len(args); i++ {
		if len(args[i]) == 0 {
			return false
		}
	}

	// Return true if all checks pass
	return true
}
