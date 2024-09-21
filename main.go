package main

import (
	"asciiArtColor/FlagsCollection"
	"asciiArtColor/Methods"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Define flags for color and align
var (
	alignFlag = flag.String("align", "AlignLeft", "Text alignment options: left, right, center, justify")
	colorFlag = flag.String("color", "Reset", "printing in colors")
)

func main() {
	preArgs := os.Args[1:]

	// Check for unrecognized flags
	for i := 0; i < len(preArgs); i++ {
		if (strings.HasPrefix(preArgs[i], "--") || strings.HasPrefix(preArgs[i], "-")) && strings.Contains(preArgs[i], "=") && (!strings.Contains(preArgs[i], "color") && !strings.Contains(preArgs[i], "align")) {
			fmt.Println("Unrecognized flag")
			return // Exit if an unrecognized flag is found
		}
	}

	// Parse the command line flags
	flag.Parse()

	// Get justified alignment and color from flag values
	align := FlagsCollection.JustifyFlag(*alignFlag)
	color := FlagsCollection.ColorFlag(*colorFlag)

	// Retrieve remaining command line arguments after flags
	args := flag.Args()

	// Validate the command line arguments
	if !Methods.ValidateArg(args) {
		fmt.Println("Unvalid argument")
		return
	}
	// Get the terminal width
	width := Methods.IoctlGetWinsize()

	// Process the arguments to extract substring and banner file name
	subString, str, banner := Methods.SubstringAndBanner(args)
	if subString == "" {
		return
	}

	// Read the specified banner file
	File, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the file content into lines for ASCII art processing
	fileLines := strings.Split(string(File), "\n")
	// Replace escape sequences for new lines in the string
	resultStr := strings.ReplaceAll(str, "\\n", "\n")
	// Split the result string into separate lines
	argLettersDivided := Methods.SplitWithNewlines(resultStr)
	// Print the ASCII art using the specified parameters
	Methods.PrintAscii(argLettersDivided, fileLines, subString, &color, width, align)

}
