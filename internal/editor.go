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
    code, key := e.readKeyPress()
    if code > -1 {
        switch code {
        case CTRL_Q:
            return true
        case ARROW_UP, ARROW_DOWN, ARROW_RIGHT, ARROW_LEFT:
            e.moveCursor(code)
            break
        case ENTER:
            e.buf.NewLine(e.cursorX, e.cursorY)
            e.cursorX ++
            e.cursorY = 0
            break
        case BACKSPACE:
            e.buf.Backspace(e.cursorX, e.cursorY)
            e.cursorY --
            if e.cursorY == 0 && e.cursorX > 0 {
                e.cursorX --
            }
            break
        }
    } else {
        e.buf.InsertRune(e.cursorX, e.cursorY, key)
        e.cursorY ++
    }

    e.fixCursorOutbound()

    return false
}

func (e *Editor) moveCursor(keyType int) {
    switch keyType {
    case ARROW_UP:
        e.cursorX--
    case ARROW_DOWN:
        e.cursorX++
    case ARROW_RIGHT:
        e.cursorY++
    case ARROW_LEFT:
        e.cursorY--
    }
}

func (e *Editor) fixCursorOutbound() {
    if e.cursorX < 0 {
        e.cursorX = 0
    }
    if e.cursorX > e.buf.lineNum - 1 {
        e.cursorX = e.buf.lineNum - 1
    }
    if e.cursorY < 0 {
        e.cursorY = 0
    }
    colNum := e.buf.GetCurrColNum(e.cursorX)
    if e.cursorY > colNum {
        e.cursorY = colNum
    }
}