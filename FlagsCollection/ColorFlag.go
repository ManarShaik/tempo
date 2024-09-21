package FlagsCollection

import (
	"fmt"
	"strings"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)
// ColorFlag maps the provided color flag string to its corresponding ANSI color code
func ColorFlag(colorFlag string) string {

	// Create a map of string representations to ANSI color codes
	colorMap := map[string]string{
		"red":     Red,
		"green":   Green,
		"yellow":  Yellow,
		"blue":    Blue,
		"magenta": Magenta,
		"cyan":    Cyan,
		"white":   White,
		"reset":   Reset,
	}

	// Convert the input color flag to lowercase for case-insensitive matching
	colorKey := strings.ToLower(colorFlag)

	// Look up the color in the map
	color, exists := colorMap[colorKey]
	if !exists {
		// If the color does not exist, print a warning and default to Reset
		fmt.Println("Invalid color specified. Defaulting to Reset.")
		color = Reset
	}
	return color
}
