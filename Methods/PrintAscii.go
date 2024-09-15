package Methods

import (
	"fmt"
	"strings"
)

const (
	Reset = "\033[0m"
)

func PrintAscii(argLettersDivided, fileLines []string, subString string, color *string, width int, align string) {
	tempColor := Reset
	var start int = 0
	count := 0
	for i := 0; i < len(argLettersDivided); i++ {
		if argLettersDivided[i] == "\\n" {
			fmt.Println()
			continue
		}
		leftSpacess , rightSpace := Justify(align, width, argLettersDivided[i], fileLines)
		for j := 0; j < 8; j++ {
			fmt.Print(strings.Repeat(" ", leftSpacess))
			for k := 0; k < len(argLettersDivided[i]); k++ {
				
				if (argLettersDivided[i])[k] == subString[0] && i+len(subString)-1 <= len(argLettersDivided[i])-1 {
					if ValidateSubString(subString, (argLettersDivided[i])[k:k+len(subString)]){
					    tempColor = *color
					    count = len(subString)
					}
				
				}
				start = ((int((argLettersDivided[i])[k]) - 32) * 9) + (j + 1)
				if (argLettersDivided[i])[k] == ' ' {
					fmt.Print((fileLines[start]))
					continue
				}
				fmt.Print(tempColor + fileLines[start])
				start = 0
			
			if count==1{
				tempColor=Reset
			}
			count--

			}
			fmt.Print(strings.Repeat(" ", rightSpace))
			fmt.Print("\n")
			
		}
	}
}
