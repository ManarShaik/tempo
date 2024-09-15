package Methods

func ValidateSubString(subString, str string)bool{
for i:=0; i<len(str); i++{
if subString[i]!=str[i]{
	return false
}
}
return true
}