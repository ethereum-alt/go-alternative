// Copyright 2017 The go-alternative Authors
// This file is part of the go-alternative library.
//
// The go-alternative library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-alternative library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-alternative library. If not, see <http://www.gnu.org/licenses/>.

// +build linux darwin freebsd

package fuse

import (
	"bazil.org/fuse/fs"
)

var (
	_ fs.Node = (*SwarmDir)(nil)
)

type SwarmRoot struct {
	root *SwarmDir
}

func (filesystem *SwarmRoot) Root() (fs.Node, error) {
	return filesystem.root, nil
}
