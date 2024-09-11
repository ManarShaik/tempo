package Methods
import (
"strings"
"fmt"
)
func ValidateArg(args []string)bool{
	if len(args)<1|| len(args)>3{
		return false
	}else{
		if len(args)==2&&!strings.HasSuffix(args[1],".txt")&&!strings.Contains(args[1], args[0]){
			fmt.Println("subString dosen't exist")
return false
		}
		if len(args)==3&&strings.HasSuffix(args[2],".txt")&&!strings.Contains(args[1], args[0]){
			fmt.Println("subString dosen't exist")
			return false
					}
	}
	for i := 0; i < len(args); i++ {
		if len(args[i]) == 0{
			return false
		}
	}
	return true
}