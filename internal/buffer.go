package ste

type Buffer struct {
    lines []line
    dirty bool
    id int
}

type line struct {
    txt [] rune
}

func (b *Buffer) NewLine(x int) {
    b.lines = append(b.lines, line{})
    if x == len(b.lines) - 2 {
        return
    }
    copy(b.lines[x+1:], b.lines[x:])
    b.lines[x] = line{}
}

func (b *Buffer) Insert(x, y int, data rune) {
    b.lines[x].txt = append(b.lines[x].txt, data)
    if y == len(b.lines[x].txt) - 1 {
        return
    } 
    copy(b.lines[x].txt[y+1:], b.lines[x].txt[y:])
    b.lines[x].txt[y] = data
}
