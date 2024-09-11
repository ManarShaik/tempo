package main

import (
	"asciiArtColor/FlagsCollection"
	"asciiArtColor/Methods"
	"flag"
	"fmt"
	"os"
	"strings"
    "syscall"
)
var(
	alignFlag = flag.String("align", "AlignLeft", "Text alignment options: left, right, center, justify")
    colorFlag = flag.String("color", "Reset", "printing in colors")

)
func main() {
    flag.Parse()

	align := FlagsCollection.JustifyFlag(*alignFlag)
	color := FlagsCollection.ColorFlag(*colorFlag)
	args := flag.Args()
	if !Methods.ValidateArg(args){
		fmt.Println("unvalide argument")
		return
	}
    var ws Methods.Winsize
    _, width := Methods.IoctlGetWinsize(int(os.Stdin.Fd()), syscall.TIOCGWINSZ, &ws)

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
