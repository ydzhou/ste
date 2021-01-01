package ste

import ("fmt")

type Buffer struct {
    lines []line
    dirty bool
    id int
}

type line struct {
    txt [] rune
}

func (b *Buffer) New() {
    b.lines = []line{line{}}
}

func (b *Buffer) NewLine(x, y int) {
    if x > len(b.lines) - 1 || y > len(b.lines[x].txt) {
        panic(fmt.Errorf("failed to create new line at (%d,%d)", x, y))
    }

    line := line{}
    if y <= len(b.lines[x].txt) {
        line.txt = make([]rune, len(b.lines[x].txt[:y]))
        copy(line.txt, b.lines[x].txt[:y])
        nextLineTxt := make([]rune, len(b.lines[x].txt[y:]))
        copy(nextLineTxt, b.lines[x].txt[y:])
        b.lines[x].txt = nextLineTxt
    }

    b.lines = append(b.lines, line)
    copy(b.lines[x+1:], b.lines[x:])
    b.lines[x] = line

    return
}

func (b *Buffer) Insert(x, y int, data rune) {
    // Append a new line if cursor is under the last line
    if x == len(b.lines) && y == 0 {
        b.lines = append(b.lines, line{})
    }

    if x > len(b.lines) - 1 || y > len(b.lines[x].txt) {
        panic(fmt.Errorf("failed to insert [%s] at (%d,%d)", string(data), x, y))
    }

    b.lines[x].txt = append(b.lines[x].txt, data)
    if y == len(b.lines[x].txt) - 1 {
        return
    } 

    copy(b.lines[x].txt[y+1:], b.lines[x].txt[y:])
    b.lines[x].txt[y] = data
}
