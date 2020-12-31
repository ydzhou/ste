package ste

import (
    "os"
    "bufio"
)

type Editor struct {
    term TermConfig
    buf Buffer
    reader *bufio.Reader
    render Render
    cursorX, cursorY int
    rowOffset, colOffset int
}

func (e *Editor) Init() {
    e.term = TermConfig{}
    e.buf = Buffer{}
    e.reader = bufio.NewReader(os.Stdin)
    e.render = Render{}
    e.cursorX = 0
    e.cursorY = 0
    e.buf.NewLine(0)
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
    keyAscii, key := e.readKeyPress()
    switch keyAscii {
    case CTRL_Q:
        return true
    case ARROW_UP, ARROW_DOWN, ARROW_RIGHT, ARROW_LEFT:
        e.moveCursor(keyAscii)
        break
    case '\r':
        e.buf.NewLine(e.cursorX)
        e.cursorX ++
        break
    default:
        e.buf.Insert(e.cursorX, e.cursorY, key)
        e.cursorY ++
    }
    return false
}

func (e *Editor) readKeyPress() (int, rune){
    r, _, _ := e.reader.ReadRune()
    if r == 27 {
        r, _, _ = e.reader.ReadRune()
        r, _, _ = e.reader.ReadRune()
    }
    switch int(r) {
        case 17: return CTRL_Q, r 
    }
    switch int(r) {
        case 65: return ARROW_UP, r
        case 66: return ARROW_DOWN, r
        case 67: return ARROW_RIGHT, r
        case 68: return ARROW_LEFT, r
    } 
    return -1, r
}

func (e *Editor) moveCursor(keyType int) {
    switch keyType {
    case ARROW_UP:
        if e.cursorX > 0 {
            e.cursorX --
        }
    case ARROW_DOWN:
        // if e.cursorX < len(e.buf.lines) - 1 {
            e.cursorX ++
        // }
    case ARROW_RIGHT:
        // if len(e.buf.lines) > 0 && e.cursorY < len(e.buf.lines[e.cursorX].txt) - 1 {
            e.cursorY ++
        // }
    case ARROW_LEFT:
        if e.cursorY > 0 {
            e.cursorY -- 
        }
    }
}
