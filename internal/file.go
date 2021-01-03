package ste

import (
	"bufio"
	"os"
)

func (e *Editor) Open(name string) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	e.buf.txt = []rune{}
	e.buf.lineNum = 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		e.buf.txt = append(e.buf.txt, []rune(sc.Text())...)
		e.buf.txt = append(e.buf.txt, '\n')
		e.buf.lineNum++
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}
}

func (e *Editor) Save(name string) {}
