// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

package server

import (
	"github.com/mpictor/gprovision/pkg/oss/pblog/pb"
)

type Persister interface {
	//StoreLog actually appends. All other store operations overwrite.
	StoreLog(id string, les *pb.LogEvents) error
	RetrieveLog(id string) (m pb.LogEvents, err error)
	StoreMacs(id string, m pb.MACs) error
	RetrieveMacs(id string) (m pb.MACs, err error)
	StoreIpmiMacs(id string, m pb.MACs) error
	RetrieveIpmiMacs(id string) (m pb.MACs, err error)
	StorePass(id string, p *pb.Credentials) error
	RetrievePass(id string) (p *pb.Credentials, err error)
	Ids() []string
	Close() error
}

type Backender interface {
	OpenDB(path string) Persister
}
