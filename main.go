package main

import (
	"asciiArtColor/FlagsCollection"
	"asciiArtColor/Methods"
	"flag"
	"fmt"
	"os"
	"strings"
)
var(
	alignFlag = flag.String("align", "AlignLeft", "Text alignment options: left, right, center, justify")
    colorFlag = flag.String("color", "Reset", "printing in colors")

)
func main() {
	preArgs := os.Args[1:]
	for i:=0; i<len(preArgs); i++{
        if (strings.HasPrefix(preArgs[i], "--") || strings.HasPrefix(preArgs[i], "-"))&& strings.Contains(preArgs[i], "=")&&(!strings.Contains(preArgs[i], "color")&&!strings.Contains(preArgs[i], "align")){
            fmt.Println("Unrecognized flag")
            return
        }
    }
    flag.Parse()
	align := FlagsCollection.JustifyFlag(*alignFlag)
	color := FlagsCollection.ColorFlag(*colorFlag)
	args := flag.Args()

	if !Methods.ValidateArg(args){
		fmt.Println("Unvalid argument")
		return
	}
    
    width := Methods.IoctlGetWinsize()

    subString , str , banner := Methods.SubstringAndBanner(args)
	if subString==""{
		return
	}
	File, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fileLines := strings.Split(string(File), "\n")
	resultStr := strings.ReplaceAll(str, "\\n", "\n")
	argLettersDivided := Methods.SplitWithNewlines(resultStr)

	Methods.PrintAscii(argLettersDivided, fileLines, subString, &color, width, align)
	

}
