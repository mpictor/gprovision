// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

package recovery

import (
	"github.com/mpictor/gprovision/pkg/common/rkeep"
	"github.com/mpictor/gprovision/pkg/hw/power"
	"github.com/mpictor/gprovision/pkg/log"
)

var RecFatal = log.FailAction{
	MsgPfx: "ERROR: ",
	Pre: func(f string, va ...interface{}) {
		if log.GetPrefix() == "test" {
			panic("Fatalf called from 'go test'")
		}
		rkeep.ReportFailure(log.GetPrefix() + " failed, rebooting...")
	},
	Terminator: func() {
		power.Reboot(false)
	},
}
