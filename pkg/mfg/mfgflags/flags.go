// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

// Package mfgflags handles flags used to alter mfg behavior for testing.
package mfgflags

import (
	"os"
	"strings"

	"github.com/mpictor/gprovision/pkg/common/strs"
	"github.com/mpictor/gprovision/pkg/log"
)

var (
	Verbose         bool
	inin_mfg_test   string
	BehaviorAltered bool
)

const (
	VerboseLog        = "Verbose-logging"
	SkipNet           = "Skip-networking"
	ExternalJson      = "Allow-external-json"
	StopAfterValidate = "Stop-after-validate"
	NoRecov           = "No-create-recov"
	NoWrite           = "No-write-files"
	NoMfg             = "No-manufacture"
	NoWipe            = "No-wipe-disks"
	RawDmi            = "Dmi-raw-output"
	NoBiosPw          = "No-bios-pw"
	NoIpmiPw          = "No-ipmi-pw"
)

func init() {
	inin_mfg_test = os.Getenv(strs.MfgTestEnv())
	Verbose = Flag(VerboseLog)
	if len(inin_mfg_test) != 0 && inin_mfg_test != VerboseLog {
		log.Logf("WARNING - env var %s has altered app's behavior. Not safe for production.", strs.MfgTestEnv())
		BehaviorAltered = true
	}
}
func Flag(str string) (b bool) {
	return strings.Contains(inin_mfg_test, str)
}
