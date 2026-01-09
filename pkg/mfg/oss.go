// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

package mfg

import (
	"github.com/mpictor/gprovision/pkg/oss/frd"
	"github.com/mpictor/gprovision/pkg/oss/pblog"
	"github.com/mpictor/gprovision/pkg/oss/stash"
)

func init() {
	pblog.UseRLoggerSetup()
	pblog.UseRKeeper()
	stash.UseImpl()
	frd.UseImpl()
}
