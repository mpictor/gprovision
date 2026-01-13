// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

// Package netexport reads network config data from windows. The package
// ../pkg/systemd/networkd can then be used to write config files compatible
// with systemd-networkd.
//
// Requires Powershell, SaveRestore.ps1 (Intel). Some data is retrieved from
// the registry, while other data comes from the output of SaveRestore.ps1 or
// raw powershell commands.
package netexport
