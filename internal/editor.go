package ste

import (
    "fmt"
    "os"
    "io"
    "bufio"
)

type Editor struct {
    term TermConfig
    buf Buffer
    cursorPos cursorPos
    reader *bufio.Reader
}

type cursorPos struct {
    x int
    y int
}

func (e *Editor) Init() {
    e.term = TermConfig{}
    e.buf = Buffer{}
    e.cursorPos = cursorPos{x: 0, y: 0}
    e.reader = bufio.NewReader(os.Stdin)
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
    keyAscii, key := e.readKeyPress()
    switch keyAscii {
    case 17:
        return true
    case '\r':
        fmt.Print("\r\n")
        e.cursorPos.x = 0
        e.cursorPos.y ++
    default:
        fmt.Print(key)
        e.cursorPos.x ++
    }
    return false
}

func (e *Editor) readKeyPress() (int, string){
    var buf [1]byte
    _, _ = e.reader.Read(buf[:])
    return int(buf[0]), string(buf[0])
}
