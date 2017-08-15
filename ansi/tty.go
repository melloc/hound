// +build !windows

package ansi

import (
	"golang.org/x/sys/unix"
)

// Issue a ioctl syscall to try to read a termios for the descriptor. If
// we are unable to read one, this is not a tty.
func isTTY(fd uintptr) bool {
	_, err := unix.IoctlGetTermios(int(fd), ioctlReadTermios)
	return err == nil
}
