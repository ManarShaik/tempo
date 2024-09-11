package Methods

func SubstringAndBanner(args []string) (string, string, string) {
	if len(args) == 1 {
		return args[0], args[0], "standard.txt"
	}
	allBanners := []string{"standard", "shadow", "thinkertoy"}
	if len(args) == 2 {
		if args[len(args)-1]=="standard"||args[len(args)-1]=="shadow"||args[len(args)-1]=="thinkertoy"{
			return args[0], args[0], args[len(args)-1]+".txt"
		}
		for i := 0; i < len(allBanners); i++ {
			if allBanners[i] == args[len(args)-1] {
				return args[0], args[0], allBanners[i]+".txt"
			}
		}
		if len(args[0]) > len(args[1]) {
			return "", "", ""
		}
		return args[0], args[1], "standard.txt"
	}
	if len(args[0]) > len(args[1]) {
		return "", "", ""
	}
	return args[0], args[1], args[2]+".txt"
}
