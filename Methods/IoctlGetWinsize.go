package Methods

import (
    "syscall"
    "unsafe"
)

// Define Winsize structure
type Winsize struct {
    Row, Col, Xpixel, Ypixel uint16
}

// IoctlGetWinsize uses the ioctl system call to get the terminal window size
func IoctlGetWinsize(fd int, req uintptr, ws *Winsize) (int, int) {
    ret, _, errno := syscall.Syscall6(
        syscall.SYS_IOCTL,
        uintptr(fd),
        req,
        uintptr(unsafe.Pointer(ws)),
        0, // extra arguments are set to 0
        0,
        0,
    )
    if errno != 0 {
        return -1, 0
    }
    return int(ret), int(ws.Col) // return width
}
