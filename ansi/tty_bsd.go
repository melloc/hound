// +build darwin freebsd openbsd netbsd

package ansi

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA
