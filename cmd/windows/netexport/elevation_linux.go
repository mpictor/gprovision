// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

package main

import (
	"os"
	"strings"
)

func checkElevation() {
	if !strings.HasSuffix(os.Args[0], ".test") {
		panic("windows only")
	}
}
