package Methods

import (
	"os"
	"os/exec"
	"strconv"
)

// IoctlGetWinsize retrieves the number of columns in the terminal window
func IoctlGetWinsize() int {

// Create the command to get the number of columns using 'tput'
	cmd := exec.Command("tput", "cols")
// Set the standard input to the terminal's stdin
	cmd.Stdin = os.Stdin
// Execute the command and capture the output
	out, _ := cmd.Output()

// Convert the output bytes to an integer, trimming the trailing newline
	result, _ := strconv.Atoi(string(out[:len(out)-1]))
	return result
}

// // Define Winsize structure
// type Winsize struct {
// 	Row, Col, Xpixel, Ypixel uint16
// }