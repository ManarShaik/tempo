package Methods

import (
	"os"
	"fmt"
	"os/exec"
    "strconv"
)

// Define Winsize structure
type Winsize struct {
	Row, Col, Xpixel, Ypixel uint16
}

// IoctlGetWinsize uses the ioctl system call to get the terminal window size
func IoctlGetWinsize() int {

    // Create the command
	cmd := exec.Command("tput", "cols")

	cmd.Stdin=os.Stdin

    out,_:=cmd.Output()

	// Print the number of columns
	result,_:=strconv.Atoi(string(out[:len(out)-1]))
     fmt.Println(result)
	return result
}
