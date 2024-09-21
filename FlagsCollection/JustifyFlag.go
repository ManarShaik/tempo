package FlagsCollection

import (
	"fmt"
	"strings"
)

const (
	AlignLeft    = "left"
	AlignRight   = "right"
	AlignCenter  = "center"
	AlignJustify = "justify"
)

// JustifyFlag maps the provided alignment flag string to a valid alignment option
func JustifyFlag(alignFlag string) string {

	// Define a map of valid alignment options
	validAlignments := map[string]bool{
		AlignLeft:    true,
		AlignRight:   true,
		AlignCenter:  true,
		AlignJustify: true,
	}
// Convert the input alignment flag to lowercase for case-insensitive matching
	alignKey := strings.ToLower(alignFlag)

	// Check if the provided alignment exists in the valid alignments map
	if _, exists := validAlignments[alignKey]; !exists {
		// If the alignment is invalid, print a warning and default to left alignment
		fmt.Println("Invalid alignment specified. Defaulting to left.")
		return AlignLeft
	}

	return alignKey
}
