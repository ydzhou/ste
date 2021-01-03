package ste

type Buffer struct {
    txt [] rune
    dirty bool
    id int
    lineNum int
}

func (b *Buffer) New() {
    b.txt = []rune{'\n'}
    b.lineNum = 1
}

func (b *Buffer) NewLine(x, y int) {
    b.InsertRune(x, y, '\n')
    b.lineNum ++
}

func (b *Buffer) InsertRune(x, y int, data rune) {
    b.Insert(x, y, []rune{data})
}

func (b *Buffer) Insert(x, y int, data []rune) {
    idx := b.getIdx(x, y)
    if idx == len(b.txt) {
        b.txt = append(b.txt, data...)
        b.txt = append(b.txt, '\n')
        b.lineNum ++
        return
    }
    txt := make([] rune, idx)
    copy(txt, b.txt[:idx])
    txt = append(txt, data...)
    txt = append(txt, b.txt[idx:]...)
    b.txt = txt
    b.dirty = true
}

func (b *Buffer) Backspace(x, y int) {
    idx := b.getIdx(x, y)
    if idx == 0 {
        return
    }
    if b.txt[idx - 1] == '\n' {
        b.lineNum --
    }
    txt := make([] rune, idx - 1)
    copy(txt, b.txt[:idx - 1])
    txt = append(txt, b.txt[idx:]...)
    b.txt = txt
    b.dirty = true
}

func (b *Buffer) getIdx(x, y int) int {
    idx := 0
    for idx, _ = range b.txt {
        if x == 0 && y == 0 {
            break
        }
        if x > 0 && b.txt[idx] == rune('\n') {
            x--
            continue
        }
        if x == 0 {
            y --
        }
        if b.txt[idx] == rune('\n') {
            break
        }
    }
    return idx
}

func (b *Buffer) GetCurrColNum(x int) int {
    idx := 0
    for i, _ := range b.txt {
        if x == 0 {
            if b.txt[i] == '\n' {
                return idx
            }
            idx ++
        }
        if b.txt[i] == '\n' {
            x --
        }
    }
    return idx
}