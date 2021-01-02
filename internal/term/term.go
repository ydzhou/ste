package term

import (
    "os"
    "syscall"
    "unsafe"
    "errors"
)

type Term struct {
    prevTerm *termios
}

type termios struct {
    Iflag  uint64
    Oflag  uint64
    Cflag  uint64
    Lflag  uint64
    Cc     [20]byte
    Ispeed uint64
    Ospeed uint64
}

func (t *Term) Raw() (error)  {
    term, err := t.getAttr(os.Stdin.Fd())
    if err != nil {
        return err
    }
    t.prevTerm = &termios{}
    *t.prevTerm = *term
    t.setRaw(term)
    t.setAtt(os.Stdin.Fd(), term)

    return nil
}

func (t *Term) Reset() {
    _ = t.setAtt(os.Stdin.Fd(), t.prevTerm)
}

func (t *Term) setRaw(term *termios) {
    term.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON
    term.Oflag &^= syscall.OPOST
    term.Lflag &^= syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
    term.Cflag &^= syscall.CSIZE | syscall.PARENB
    term.Cflag |= syscall.CS8
    term.Cc[syscall.VMIN] = 1
    term.Cc[syscall.VTIME] = 0
}
func (t *Term) getAttr(fd uintptr) (*termios, error) {
    var term termios
    _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, TIOGETATT, uintptr(unsafe.Pointer(&term)))
    if err != 0 {
        return nil, errors.New("failed to get term attributes")
    }
    return &term, nil
}

func (t *Term) setAtt(fd uintptr, term *termios) (error) {
    _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, TIOSETATT, uintptr(unsafe.Pointer(term)))
    if err != 0 {
        return errors.New("faled to set term attributes")
    }
    return nil
}
