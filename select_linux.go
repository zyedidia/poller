// Copyright (c) 2014, Nick Patavalis (npat@efault.net).
// All rights reserved.
// Use of this source code is governed by a BSD-style license that can
// be found in the LICENSE.txt file.

// +build linux,noepoll

package poller

import (
	"syscall"
	"time"
)

func uxSelect(nfd int, r, w, e *fdSet, tmo time.Duration) (n int, err error) {
	tv := NsecToTimeval(tmo.Nanoseconds())
	return syscall.Select(nfd,
		(*syscall.FdSet)(r),
		(*syscall.FdSet)(w),
		(*syscall.FdSet)(e),
		&tv)
}
