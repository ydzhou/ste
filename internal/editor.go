package ste

import (
    "os"
    "bufio"
    "github.com/ydzhou/ste/internal/term"
)

type Editor struct {
    term term.Term
    buf Buffer
    reader *bufio.Reader
    render Render
    cursorX, cursorY int
    rowOffset, colOffset int
}

func (e *Editor) Init() {
    e.term = term.Term{}
    e.buf = Buffer{}
    e.reader = bufio.NewReader(os.Stdin)
    e.render = Render{}
    e.cursorX = 0
    e.cursorY = 0
    e.buf.New()
}

func (e *Editor) Start() {
    _ = e.term.Raw()

    e.render.Clear()

    defer e.term.Reset()

    for {
        e.render.DrawScreen(e.buf, e.cursorX, e.cursorY, e.rowOffset, e.colOffset)
        if e.process() {
            break
        }
    }

    e.render.Clear()
}

func (e *Editor) process() bool {
    keyAscii, key, special := e.readKeyPress()
    if special {
    switch keyAscii {
    case CTRL_Q:
        return true
    case ARROW_UP, ARROW_DOWN, ARROW_RIGHT, ARROW_LEFT:
        e.moveCursor(keyAscii)
        break
    case ENTER:
        e.buf.NewLine(e.cursorX, e.cursorY)
        e.cursorX ++
        e.cursorY = 0
        break
    }
    } else {
        e.buf.Insert(e.cursorX, e.cursorY, key)
        e.cursorY ++
    }
    return false
}

func (e *Editor) moveCursor(keyType int) {
    switch keyType {
    case ARROW_UP:
        if e.cursorX > 0 {
            e.cursorX --
        }
    case ARROW_DOWN:
        if e.cursorX < len(e.buf.lines) {
            e.cursorX ++
        }
    case ARROW_RIGHT:
        if len(e.buf.lines) > 0 && e.cursorY < len(e.buf.lines[e.cursorX].txt) {
            e.cursorY ++
        }
    case ARROW_LEFT:
        if e.cursorY > 0 {
            e.cursorY -- 
        }
    }
}
