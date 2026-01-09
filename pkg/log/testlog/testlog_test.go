// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

package testlog

import (
	"fmt"
	"testing"

	"github.com/mpictor/gprovision/pkg/log"
)

func Example() {
	//hack since examples have no args
	t := &testing.T{}
	tlog := NewTestLog(t, true, false)
	tlog.FatalIsNotErr = true
	log.Msgf("doing something important")
	log.Logf("technical details...")
	//test something that would normally terminate execution
	bad := true
	if bad {
		log.Fatalf("some severe error")
	}
	tlog.Freeze()
	//ensure that log.Fatalf() was called
	if tlog.FatalCount != 1 {
		t.Errorf("FatalCount must be 1")
	}
	fmt.Println(tlog.Buf.String())
	//output: MSG:doing something important
	//LOG:technical details...
	//>>FATAL()<< some severe error
}
