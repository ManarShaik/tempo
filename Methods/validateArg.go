package Methods

func ValidateArg(args []string)bool{
	if len(args)<1|| len(args)>3{
		return false
	}
	for i := 0; i < len(args); i++ {
		if len(args[i]) == 0{
			return false
		}
	}
	return true
}