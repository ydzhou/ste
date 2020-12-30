package dt

import (
    "os"
    "github.com/google/goterm/term"
)

type TermConfig struct {
    config term.Termios
}

func (t *TermConfig) Raw() (error)  {
    var err error
    t.config, err = term.Attr(os.Stdin)
    if err != nil {
        return err
    }
    config := t.config
    config.Raw()
    config.Set(os.Stdin)
    return nil
}

func (t *TermConfig) Reset() {
    t.config.Set(os.Stdin)
}
