//go:build windows

package reuseable

import (
    "golang.org/x/sys/windows"
    "syscall"
)

func rawControl(rawConn syscall.RawConn) error {
    var err error
    // See syscall.RawConn.Control
    rawConn.Control(func(fd uintptr) {
        err = windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1)
        if err != nil {
            return
        }
    })
    return err
}

// See net.ListenConfig and net.Dialer's Control attribute
func control(network string, address string, rawConn syscall.RawConn) error {
    // See syscall.RawConn.Control
    if err := rawControl(rawConn); err != nil {
        return err
    }
    return nil
}
