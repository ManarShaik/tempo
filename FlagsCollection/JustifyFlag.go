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


func JustifyFlag(alignFlag string) string {
	validAlignments := map[string]bool{
		AlignLeft:    true,
		AlignRight:   true,
		AlignCenter:  true,
		AlignJustify: true,
	}

	alignKey := strings.ToLower(alignFlag)
	if _, exists := validAlignments[alignKey]; !exists {
		fmt.Println("Invalid alignment specified. Defaulting to left.")
		return AlignLeft
	}

	
	return alignKey
}
