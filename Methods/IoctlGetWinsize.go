package Methods

import (
    "syscall"
    "unsafe"
)


type Winsize struct {
    Row    uint16
    Col    uint16
    Xpixel uint16
    Ypixel uint16
}


func IoctlGetWinsize(fd int, req int, ws *Winsize) (error, int) {
    // Perform the syscall to get the terminal size
    _, _, err := syscall.Syscall(
        syscall.SYS_IOCTL,
        uintptr(fd),
        uintptr(req),
        uintptr(unsafe.Pointer(ws)),
    )
    if err != 0 {
        return err, 0
    }


    return nil, int(ws.Col)
}
