package Methods

import (
    "syscall"
    "unsafe"
)

// Define Winsize structure
type Winsize struct {
    Row, Col, Xpixel, Ypixel uint16
}

// Define constants
const (
    TIOCGWINSZ = 0x5400 // ioctl request code for getting window size
    SYS_IOCTL  = 0x5410 // ioctl system call number (This value may need to be adjusted based on your OS)
)

// IoctlGetWinsize performs an ioctl system call to get terminal window size
func IoctlGetWinsize(fd int, req uintptr, ws *Winsize) (int, int) {
    ret, _, errno := syscall.Syscall6(
        uintptr(SYS_IOCTL),                        // syscall number for ioctl
        uintptr(fd),                      // file descriptor
        req,                              // request code
        uintptr(unsafe.Pointer(ws)),     // pointer to the Winsize structure
        uintptr(0), uintptr(0), uintptr(0),uintptr(0),                           // unused arguments for syscall.Syscall6
    )
    if errno != 0 {
        return -1, 0
    }
    return int(ret), int(ws.Col) // return width
}