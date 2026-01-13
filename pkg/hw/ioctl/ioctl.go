// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

package ioctl

import (
	"syscall"
	"unsafe"
)

// IMPORTANT
// An ioctl() request has encoded in it whether the argument is an in
// parameter or out parameter, and the size of the argument argp in
// bytes.

type FDer interface {
	Fd() uintptr
}

func Ioctl1(fd uintptr, cmd int) (res uint64, err error) {
	ptr := uintptr(unsafe.Pointer(&res))
	err = ioctl(fd, uintptr(cmd), ptr)
	return res, err
}

func ioctl(fd, cmd, ptr uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, ptr)
	if err == 0 {
		return nil
	}
	return err
}
