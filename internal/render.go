package ste

import (
    "io"
    "os"
    "fmt"
    // "bufio"
)

type Render struct {
}

func (r *Render) Clear() {
    io.WriteString(os.Stdout, "\x1b[2J")
    io.WriteString(os.Stdout, "\x1b[H")
}

func (r *Render) DrawScreen(
    buf Buffer,
    cursorX int,
    cursorY int,
    rowOffset int,
    colOffset int,
) {
    // bwriter := bufio.NewWriter(os.Stdout)
    bufString := "\x1b[25l"
    r.Clear()
    bufString += r.drawHeader()
    bufString += r.drawBuffer(buf, rowOffset, colOffset)
    fmt.Print(bufString)
    bufString = ""
    bufString += r.drawCursor(cursorX, cursorY, rowOffset, colOffset)
    bufString += "\x1b[25l"
    
    fmt.Print(bufString)
}

func (r *Render) drawHeader() string {
    return "STE\n\n"
}

func (r *Render) drawBuffer(
    buf Buffer,
    rowOffset int,
    colOffset int,
) string {
    res := ""
    for i, row := range buf.lines{
        if i < rowOffset {
            continue
        }
        res += string(row.txt[colOffset:])
        res += "\n"
    }
    return res
}

func (r *Render) drawCursor(
    cursorX int,
    cursorY int,
    rowOffset int,
    colOffset int,
) string {
    return fmt.Sprintf(
        "\x1b[%d;%dH", 
        cursorX - rowOffset + 3, 
        cursorY - colOffset + 1)
}
