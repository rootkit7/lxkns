// Definitions of data types and constants related to Linux kernel namespaces.

// Copyright 2020 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build linux

package species

import (
	"strconv"

	"golang.org/x/sys/unix"
)

// NamespaceType mirrors the data type used in the Linux kernel for the
// namespace type constants. These constants are actually part of the clone()
// syscall options parameter.
type NamespaceType uint64

// The 7 type of Linux namespaces defined at this time (sic!). Well, the 8th
// namespace for time is already ticking along...
//
// These constants are used with several of the namespace-related functions,
// such as clone() in particular, but also setns(), unshare(), and the
// NS_GET_NSTYPE ioctl(). The origin for their definitions is:
// https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/sched.h
//
// Oh, forgo golint with its "helicopter parents" attitude patronizing us about
// how names of Linux kernel definitions have to look like. Go for something
// grown up, such as golangci-lint, and many more.
const (
	CLONE_NEWNS     = NamespaceType(unix.CLONE_NEWNS)
	CLONE_NEWCGROUP = NamespaceType(unix.CLONE_NEWCGROUP)
	CLONE_NEWUTS    = NamespaceType(unix.CLONE_NEWUTS)
	CLONE_NEWIPC    = NamespaceType(unix.CLONE_NEWIPC)
	CLONE_NEWUSER   = NamespaceType(unix.CLONE_NEWUSER)
	CLONE_NEWPID    = NamespaceType(unix.CLONE_NEWPID)
	CLONE_NEWNET    = NamespaceType(unix.CLONE_NEWNET)
)

// NaNS identifies an invalid namespace type.
const NaNS NamespaceType = 0

// Name returns the type name string (such as "mnt", "net", ...) of a
// namespace type value.
func (nstype NamespaceType) Name() string {
	name := typeNames[nstype]
	return name
}

// String returns the Linux kernel namespace constant name for a given
// namespace type value.
func (nstype NamespaceType) String() string {
	switch nstype {
	case NaNS:
		return "NaNS"
	case CLONE_NEWNS:
		return "CLONE_NEWNS"
	case CLONE_NEWCGROUP:
		return "CLONE_NEWCGROUP"
	case CLONE_NEWUTS:
		return "CLONE_NEWUTS"
	case CLONE_NEWIPC:
		return "CLONE_NEWIPC"
	case CLONE_NEWUSER:
		return "CLONE_NEWUSER"
	case CLONE_NEWPID:
		return "CLONE_NEWPID"
	case CLONE_NEWNET:
		return "CLONE_NEWNET"
	default:
		return "NamespaceType(" + strconv.FormatInt(int64(nstype), 10) + ")"
	}
}

// Maps Linux namespace constants to their "short" type names, as used in the
// proc filesystem.
var typeNames = map[NamespaceType]string{
	CLONE_NEWNS:     "mnt",
	CLONE_NEWCGROUP: "cgroup",
	CLONE_NEWUTS:    "uts",
	CLONE_NEWIPC:    "ipc",
	CLONE_NEWUSER:   "user",
	CLONE_NEWPID:    "pid",
	CLONE_NEWNET:    "net",
}

// NameToType returns the namespace type value (constant CLONE_NEWNS, ...)
// corresponding to the specified namespace type name (such as "mnt", "net",
// et cetera).
func NameToType(name string) NamespaceType {
	t := nameTypes[name]
	return t
}

// Maps Linux namespace type names (as used in the proc filesystem) to their
// Linux kernel constants.
var nameTypes = map[string]NamespaceType{
	"mnt":    CLONE_NEWNS,
	"cgroup": CLONE_NEWCGROUP,
	"uts":    CLONE_NEWUTS,
	"ipc":    CLONE_NEWIPC,
	"user":   CLONE_NEWUSER,
	"pid":    CLONE_NEWPID,
	"net":    CLONE_NEWNET,
}
