package Methods

func Justify(align string, width int, word string, fileLines []string)(int,int){
wordWidth := 0
for i:=0; i<len(word); i++{
		start := ((int((word)[i]) - 32) * 9) + (1)
		wordWidth += len(fileLines[start])
}
spacesRemain := width - wordWidth

switch align {
case "left":
	return 0 , spacesRemain
case "right":
	return spacesRemain , 0
case "center":
	return spacesRemain/2 , spacesRemain/2
default:
	return 0,0
}
}