package ste

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Render struct {
	viewX, viewY int
}

func (r *Render) Clear() {
	_, err := io.WriteString(os.Stdout, "\x1b[2J")
	_, err = io.WriteString(os.Stdout, "\x1b[H")
	if err != nil {
	    panic(err)
    }
}

func (r *Render) DrawScreen(
	buf Buffer,
	cursorX int,
	cursorY int,
	rowOffset int,
	colOffset int,
) {
	b := bytes.Buffer{}
	b.WriteString("\x1b[?25l")
	b.WriteString("\x1b[H")
	b.WriteString(r.drawBuffer(buf))
	b.WriteString(r.drawCursor(cursorX, cursorY, rowOffset, colOffset))
	b.WriteString("\x1b[?25h")

	_, err := b.WriteTo(os.Stdout)
	if err != nil {
	    panic(err)
    }
}

func (r *Render) drawBuffer(
	buf Buffer,
) string {
	res := ""
	for _, r := range buf.txt {
		if r == '\n' {
			res += "\x1b[K"
			res += "\r\n"
			continue
		}
		res += string(r)
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
		cursorX-rowOffset+1,
		cursorY-colOffset+1)
}
