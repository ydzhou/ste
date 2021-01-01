// +build darwin freebsd netbsd openbsd solaris dragonfly

package term

import "syscall"

const TIOGETATT = syscall.TIOCGETA
const TIOSETATT = syscall.TIOCSETA
