package ste

import (
	"bytes"
	"fmt"
	"os"
)

type Display struct {
	viewX, viewY     int
	offsetX, offsetY int
	txt              [][]rune
}

func (d *Display) Clear() {
	b := bytes.Buffer{}

	for i := 0; i < d.viewX; i++ {
		if i != 0 {
			b.WriteString("\r\n")
		}
		b.WriteString("\x1b[K")
	}

	b.WriteTo(os.Stdout)
}

func (d *Display) DrawScreen(
	buf Buffer,
	cursorX int,
	cursorY int,
) {
	d.scroll(cursorX, cursorY)
	d.convertBuffer(buf)
	b := &bytes.Buffer{}
	b.WriteString("\x1b[?25l")
	b.WriteString("\x1b[H")
	d.drawBuffer(b)
	b.WriteString(d.drawCursor(cursorX, cursorY))
	b.WriteString("\x1b[?25h")

	_, err := b.WriteTo(os.Stdout)
	if err != nil {
		panic(err)
	}
}

// Convert buffer to displayable texts
func (d *Display) convertBuffer(buf Buffer) {
	d.txt = [][]rune{}

	startIdx := d.getTxtIdxByRow(d.offsetX, buf.txt)
	endIdx := d.getTxtIdxByRow(d.offsetX+d.viewX, buf.txt) + 1

	line := []rune{}
	for i := startIdx; i < endIdx; i++ {
		if buf.txt[i] == '\n' {
			d.txt = append(d.txt, line)
			line = []rune{}
			continue
		}
		if buf.txt[i] == '\t' {
			for j := 0; j < TAB_SIZE; j++ {
				line = append(line, ' ')
			}
			continue
		}
		line = append(line, buf.txt[i])
	}
}

func (d *Display) getTxtIdxByRow(x int, txt []rune) int {
	idx := 0
	for idx, _ = range txt {
		if x == 0 {
			break
		}
		if txt[idx] == '\n' {
			x--
		}
	}
	return idx
}

func (d *Display) drawBuffer(
	b *bytes.Buffer,
) {
	for i, l := range d.txt {
		if i != 0 {
			b.WriteString("\r\n")
		}
		for j := d.offsetY ; j < len(l) && j < d.offsetY + d.viewY; j++ {
			b.WriteRune(l[j])
		}
		b.WriteString("\x1b[K")
	}
	for i := len(d.txt); i < d.viewX; i++ {
		if i != 0 {
			b.WriteString("\r\n")
		}
		b.WriteRune('~')
		b.WriteString("\x1b[K")
	}
	b.WriteRune('~')
	b.WriteString("\x1b[K")
}

func (d *Display) drawCursor(
	cursorX int,
	cursorY int,
) string {
	return fmt.Sprintf(
		"\x1b[%d;%dH",
		cursorX-d.offsetX + 1,
		cursorY-d.offsetY + 1)
}

func (d *Display) scroll(cursorX int, cursorY int) {
	if cursorX >= d.offsetX+d.viewX {
		d.offsetX = cursorX - d.viewX + 1
	} else if cursorX < d.offsetX {
		d.offsetX = cursorX
	}
	if cursorY >= d.offsetY+d.viewY {
		d.offsetY = cursorY - d.viewY + 1
	} else if cursorY < d.offsetY {
		d.offsetY = cursorY
	}
}
