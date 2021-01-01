package ste

func (e *Editor) readKeyPress() (int, rune, bool){
    special := false
    var b [3]byte
    _, _ = e.reader.Read(b[:])
    r := rune(b[0])
    switch int(b[0]) {
    case 17: 
        return CTRL_Q, r, true
    case 13:
        return ENTER, r, true
    }
    if int(b[0]) == 27 {
        special = true
        switch int(b[2]) {
        case 65: return ARROW_UP, r, special
        case 66: return ARROW_DOWN, r, special
        case 67: return ARROW_RIGHT, r, special
        case 68: return ARROW_LEFT, r, special
        } 
    }
    return -1, r, false
}
