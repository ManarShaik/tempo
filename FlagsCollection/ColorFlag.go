package FlagsCollection

import (
	"flag"
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

func ColorFlag() string{
	colorFlag := flag.String("color", "Reset", "printing in colors")
	flag.Parse()
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
	colorKey := strings.ToLower(*colorFlag)
	color, exists := colorMap[colorKey]
	if !exists {
		fmt.Println("Invalid color specified. Defaulting to Reset.")
		color = Reset
	}
	return color
}