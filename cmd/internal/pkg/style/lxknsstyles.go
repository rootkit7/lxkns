// Defines the rendering styles used by lxkns tools when rendering specific
// elements, such as different namespace styles, process names, PIDs, et
// cetera.

// Copyright 2020 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package style

// The set of styles for styling types of Linux-kernel namespaces differently,
// as well as some more elements, such as process names, user names, et
// cetera. The styles are meant to be directly referenced (used) by other
// packages importing our cmd/internal/style package
var (
	MntStyle    Style // styles mnt: namespaces
	CgroupStyle Style // styles cgroup: namespaces
	UTSStyle    Style // styles uts: namespaces
	IPCStyle    Style // styles ipc: namespaces
	UserStyle   Style // styles utc: namespaces
	PIDStyle    Style // styles pid: namespaces
	NetStyle    Style // styles net: namespaces

	OwnerStyle   Style // styles owner username and UID
	ProcessStyle Style // styles process names
	UnknownStyle Style // styles undetermined elements, such as unknown PIDs.
)

// Styles maps style configuration top-level element names to their
// corresponding Style objects for storing and using specific style information.
var Styles = map[string]*Style{
	"mnt":    &MntStyle,
	"cgroup": &CgroupStyle,
	"uts":    &UTSStyle,
	"ipc":    &IPCStyle,
	"user":   &UserStyle,
	"pid":    &PIDStyle,
	"net":    &NetStyle,

	"owner":   &OwnerStyle,
	"process": &ProcessStyle,
	"unknown": &UnknownStyle,
}
