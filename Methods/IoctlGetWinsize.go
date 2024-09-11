package Methods

import (
    "fmt"
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

    // Print the dimensions of the terminal
    fmt.Printf("Row: %d\n", ws.Row)
    fmt.Printf("Col: %d\n", ws.Col)
    fmt.Printf("Xpixel: %d\n", ws.Xpixel)
    fmt.Printf("Ypixel: %d\n", ws.Ypixel)

    return nil, int(ws.Col)
}
