// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux && !netbsd && !openbsd && arm64
// +build !linux,!netbsd,!openbsd,arm64

package cpu

func doinit() {}
