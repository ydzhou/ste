package term

import (
    "os"
    "syscall"
    "unsafe"
    "errors"
)

type Term struct {
    prevTerm *syscall.Termios
}

func (t *Term) Raw() (error)  {
    term, err := t.getAttr(os.Stdin.Fd())
    if err != nil {
        return err
    }
    
    t.prevTerm = term
    t.setRaw(term)
    t.setAtt(os.Stdin.Fd(), term)

    return nil
}

func (t *Term) Reset() {
    _ = t.setAtt(os.Stdin.Fd(), t.prevTerm)
}

func (t *Term) setRaw(term *syscall.Termios) {
    // This attempts to replicate the behaviour documented for cfmakeraw in
    // the termios(3) manpage.
    term.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON
    // newState.Oflag &^= syscall.OPOST
    term.Lflag &^= syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
    term.Cflag &^= syscall.CSIZE | syscall.PARENB
    term.Cflag |= syscall.CS8

    term.Cc[syscall.VMIN] = 1
    term.Cc[syscall.VTIME] = 0
}
func (t *Term) getAttr(fd uintptr) (*syscall.Termios, error) {
    var term syscall.Termios
    _, _, err := syscall.Syscall6(
        syscall.SYS_IOCTL,
        fd,
        TIOGETATT,
        uintptr(unsafe.Pointer(&term)),
        0,0,0)
    if err != 0 {
        return nil, errors.New("failed to get term attributes")
    }
    return &term, nil
}

func (t *Term) setAtt(fd uintptr, term *syscall.Termios) (error) {
    _, _, err := syscall.Syscall6(
        syscall.SYS_IOCTL,
        fd,
        TIOSETATT,
        uintptr(unsafe.Pointer(term)),
        0,0,0)
    if err != 0 {
        return errors.New("faled to set term attributes")
    }
    return nil
}
