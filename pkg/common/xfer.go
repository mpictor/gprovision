// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

package common

import (
	"os"
)

type Verifyer interface {
	Verify() error
}

type FileTransferer interface {
	Get() error
	GetWithRetry() error
	GetIntermediate() string
	UseIntermediateDir(dir string)
	Finalize() (err error)
}

type TransferableVerifiableFile interface {
	FileTransferer
	Verifyer
	Basename() string
	Mode(m os.FileMode)
}
