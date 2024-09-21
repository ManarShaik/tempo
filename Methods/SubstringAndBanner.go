package Methods

// SubstringAndBanner processes input arguments to determine the substring and banner file name
func SubstringAndBanner(args []string) (string, string, string) {

	// If only one argument is provided, return it as both the substring and the default banner
	if len(args) == 1 {
		return args[0], args[0], "standard.txt"
	}

	// Define a list of available banners
	allBanners := []string{"standard", "shadow", "thinkertoy"}

	// If two arguments are provided, check the second argument for a valid banner
	if len(args) == 2 {

		// Check if the last argument matches any known banner
		if args[len(args)-1] == "standard" || args[len(args)-1] == "shadow" || args[len(args)-1] == "thinkertoy" {
			return args[0], args[0], args[len(args)-1] + ".txt"
		}

		// Check against the predefined banner list
		for i := 0; i < len(allBanners); i++ {
			if allBanners[i] == args[len(args)-1] {
				return args[0], args[0], allBanners[i] + ".txt"
			}
		}

		// If the first argument is longer than the second, return empty strings
		if len(args[0]) > len(args[1]) {
			return "", "", ""
		}

		// Return the first argument as the substring and the second as a fallback with the default banner
		return args[0], args[1], "standard.txt"
	}

	// If three arguments are provided, ensure the first is not longer than the second
	if len(args[0]) > len(args[1]) {
		return "", "", ""
	}
	return args[0], args[1], args[2] + ".txt"
}
