// Copyright (C) 2015-2020 the Gprovision Authors. All Rights Reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// SPDX-License-Identifier: MIT
//

//prepare for erase integ test - create filesystem on recovery, write pattern on main volume
package main

import (
	"io"
	"os"
	"syscall"

	"github.com/mpictor/gprovision/pkg/appliance"
	"github.com/mpictor/gprovision/pkg/log"
	"github.com/mpictor/gprovision/pkg/recovery/disk"
)

func prepare(plat *appliance.Variant, disks []*disk.Disk) {
	tgt := disks[0]
	err := tgt.Partition(plat)
	if err != nil {
		log.Fatalf("error %s partitioning", err)
	}
	err = writePatterns(tgt)
	if err != nil {
		log.Fatalf("error %s writing patterns", err)
	} else {
		log.Log("patterns written successfully. exiting.")
	}
}

//FIXME seqRunner
func writePatterns(d *disk.Disk) error {
	//write directly to block device
	blk, err := os.OpenFile(d.Device(), os.O_WRONLY|syscall.O_DIRECT, 0)
	if err != nil {
		return err
	}
	defer blk.Close()
	siz := d.SizeBytes()
	offsets := offsetSeq(siz)
	//determine count/offsets
	for _, o := range offsets {
		err = writePat(blk, o)
		if err != nil {
			return err
		}
	}
	return nil
}
func writePat(f *os.File, offs int64) error {
	_, err := f.Seek(offs, io.SeekStart)
	if err != nil {
		return err
	}
	_, err = f.Write(pat1m(offs))
	return err
}
