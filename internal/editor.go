package dt

import (
    "fmt"
    "os"
    "io"
)

type Editor struct {
    term TermConfig
    buf Buffer
    cursorPos cursorPos
}

type cursorPos struct {
    x int
    y int
}

func (e *Editor) Init() {
    e.term = TermConfig{}
    e.buf = Buffer{}
    e.cursorPos = cursorPos{x: 0, y: 0}
}

func (e *Editor) Start() {
    _ = e.term.Raw()

    io.WriteString(os.Stdout, "\x1b[2J")
    io.WriteString(os.Stdout, "\x1b[H")

    defer e.term.Reset()

    for {
        if e.process() {
            break
        }
    }

    io.WriteString(os.Stdout, "\x1b[2J")
    io.WriteString(os.Stdout, "\x1b[H")
}

func (e *Editor) process() bool {
    keyAscii := e.readKeyPress()
    switch keyAscii {
    case 'q':
        return true
    case '\r':
        fmt.Print("\r\n")
        e.cursorPos.x = 0
        e.cursorPos.y ++
    default:
        fmt.Println(keyAscii)
        e.cursorPos.x ++
    }
    return false
}

func (e *Editor) readKeyPress() int {
    var buf []byte
    _, _ = os.Stdin.Read(buf[:])
    return int(buf[0])
}
