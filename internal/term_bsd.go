// +build darwin freebsd netbsd openbsd solaris dragonfly

package ste

import "syscall"

const TIOGETATT = syscall.TIOCGETA
const TIOSETATT = syscall.TIOCSETA
