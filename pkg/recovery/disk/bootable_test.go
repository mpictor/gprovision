// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

package disk

import (
	"strings"
	"testing"
)

// func finalizeGrubConf(root_uuid, extra_opts string) []byte
func TestFinalizeGrubConf(t *testing.T) {
	out := finalizeGrubConf("uuid1234", "opts")
	t.Logf("%s\n", string(out))
	if !strings.Contains(string(out), "uuid1234") {
		t.Errorf("missing uuid")
	}
	if !strings.Contains(string(out), "opts") {
		t.Errorf("missing opts")
	}
}

// ensures templates function
// func (bd *bootData) processTemplate(in []byte, name string) []byte
func TestDefaultTemplates(t *testing.T) {
	for _, td := range []struct {
		name string
		in   []byte
	}{
		{name: "hdd", in: hddboot},
		{name: "fb", in: rfm},
		{name: "mnu", in: menutmpl},
	} {
		t.Run(td.name, func(t *testing.T) {
			bd := &bootData{}
			out, err := bd.processTemplate(td.in, td.name)
			if err != nil {
				t.Error(err)
				t.Log(string(out))
			}
		})
	}
}
