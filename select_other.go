// Copyright (c) 2014, Nick Patavalis (npat@efault.net).
// All rights reserved.
// Use of this source code is governed by a BSD-style license that can
// be found in the LICENSE.txt file.

// +build linux,noepoll freebsd netbsd openbsd darwin dragonfly solaris

package poller

import (
	"syscall"
	"unsafe"
)

type fdSet syscall.FdSet

const nfdbits = int(unsafe.Sizeof(int64(0)) * 8)

func FDSet(fd int, set *fdSet) {
	set.Bits[fd/nfdbits] |= (1 << (uint)(fd%nfdbits))
}

func FDClr(fd int, set *fdSet) {
	set.Bits[fd/nfdbits] &^= (1 << (uint)(fd%nfdbits))
}

func FDIsSet(fd int, set *fdSet) bool {
	if set.Bits[fd/nfdbits]&(1<<(uint)(fd%nfdbits)) != 0 {
		return true
	} else {
		return false
	}
}

func FDZero(set *fdSet) {
	for i := range set.Bits {
		set.Bits[i] = 0
	}
}

func (fs *fdSet) Clr(fds ...int) {
	for _, fd := range fds {
		FDClr(fd, fs)
	}
}

func (fs *fdSet) Set(fds ...int) {
	for _, fd := range fds {
		FDSet(fd, fs)
	}
}

func (fs *fdSet) Zero() {
	FDZero(fs)
}

func (fs *fdSet) IsSet(fd int) bool {
	return FDIsSet(fd, fs)
}
