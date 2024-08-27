package main

import (
	"asciiArtColor/FlagsCollection"
	"asciiArtColor/Methods"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// var output = flag.String("output", "terminal", "where to print?")
	// var align = flag.String("align", "Reset", "printing color")
	color := FlagsCollection.ColorFlag()
	args := flag.Args()
	Methods.ValidateArg(args)
    subString , str , banner := Methods.SubstringAndBanner(args)
	if subString==""{
		return
	}
fmt.Println(subString +" "+ str +" "+ banner)
	File, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fileLines := strings.Split(string(File), "\n")
	resultStr := strings.ReplaceAll(str, "\\n", "\n")
	argLettersDivided := Methods.SplitWithNewlines(resultStr)

	Methods.PrintAscii(argLettersDivided, fileLines, subString, &color)

}
