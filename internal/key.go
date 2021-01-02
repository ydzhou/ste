package ste

func (e *Editor) readKeyPress() (int, rune){
    r, n, err := e.reader.ReadRune()
    if err != nil {
        panic(err)
    }

    if n == 0 && err == nil {
        return 0, 0
    }

    // Escape (special) key is pressed.
    if r == 27 {
        // Escape key only
        if e.reader.Buffered() == 0 {
            return 0, 27
        }

        for i := 0; i < 4; i++ {
            r, _, err = e.reader.ReadRune()
            r, _, err = e.reader.ReadRune()
            if err != nil {
                return 0, 27
            }
            switch r {
            case 65: return ARROW_UP, r
            case 66: return ARROW_DOWN, r
            case 67: return ARROW_RIGHT, r
            case 68: return ARROW_LEFT, r
            }
        }
    }

    switch r {
    case 127:
        return BACKSPACE, r
    case 17:
        return CTRL_Q, r
    case 13:
        return ENTER, r
    }

    return -1, r
}

