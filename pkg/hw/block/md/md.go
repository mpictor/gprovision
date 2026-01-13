// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

// Package md does various things with mdadm.
package md

import (
	"os/exec"

	"github.com/mpictor/gprovision/pkg/log"
)

func AssembleScan() bool {
	_, success := log.Cmd(exec.Command("mdadm", "--assemble", "--scan"))
	return success
}
