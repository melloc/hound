// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package index

import (
  "golang.org/x/sys/unix"
  "log"
  "os"
)

func mmapFile(f *os.File) mmapData {
  st, err := f.Stat()
  if err != nil {
    log.Fatal(err)
  }
  size := st.Size()
  if int64(int(size+4095)) != size+4095 {
    log.Fatalf("%s: too large for mmap", f.Name())
  }
  n := int(size)
  if n == 0 {
    return mmapData{f, nil, nil}
  }
  data, err := unix.Mmap(int(f.Fd()), 0, (n+4095)&^4095, unix.PROT_READ, unix.MAP_SHARED)
  if err != nil {
    log.Fatalf("mmap %s: %v", f.Name(), err)
  }
  return mmapData{f, data[:n], data}
}

func unmmapFile(m *mmapData) error {
  if err := unix.Munmap(m.o); err != nil {
    return err
  }

  return m.f.Close()
}

func unmmap(d []byte) error {
  return unix.Munmap(d)
}
