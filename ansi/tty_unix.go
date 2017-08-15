// +build linux solaris

package ansi

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TCGETS
