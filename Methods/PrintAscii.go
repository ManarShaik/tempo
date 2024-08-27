package Methods

import "fmt"

const (
	Reset = "\033[0m"
)

func PrintAscii(argLettersDivided, fileLines []string, subString string, color *string) {
	desiredColor := color
	tempColor := Reset
	var start int = 0
	count := 0
	for i := 0; i < len(argLettersDivided); i++ {
		if argLettersDivided[i] == "\\n" {
			fmt.Println()
			continue
		}
		for j := 0; j < 8; j++ {
			for k := 0; k < len(argLettersDivided[i]); k++ {
				if (argLettersDivided[i])[k] == subString[0] && i+len(subString)-1 <= len(argLettersDivided[i])-1 {
					tempColor = *desiredColor
					count = len(subString)
				}
				start = ((int((argLettersDivided[i])[k]) - 32) * 9) + (j + 1)
				if (argLettersDivided[i])[k] == ' ' {
					fmt.Print((fileLines[start]))
					continue
				}
				fmt.Print(tempColor + fileLines[start])
				start = 0
			}
			fmt.Print("\n")
			if count==1{
				tempColor=Reset
			}
			count--
		}
	}
}
