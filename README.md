# Linux kernel Namespaces

[![GoDoc](https://godoc.org/github.com/TheDiveO/lxkns?status.svg)](http://godoc.org/github.com/TheDiveO/lxkns)
[![GitHub](https://img.shields.io/github/license/thediveo/lxkns)](https://img.shields.io/github/license/thediveo/lxkns)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/lxkns)](https://goreportcard.com/report/github.com/thediveo/lxkns)

`lxkns` is a Golang package for discovering Linux kernel namespaces. In
contrast to most well-known CLI tools, such as `lsns`, this package detects
namespaces even in places of a running Linux system other tools do not
consider. In particular:

1. from the procfs filesystem in `/proc/[PID]/ns/*` -- as `lsns` and other tools do.
2. bind-mounted namespaces, via `/proc/[PID]/mountinfo`.
3. file descriptor-referenced namespaces, via `/proc/[PID]/fd/*`.
4. intermediate hierarchical user and PID namespaces, via `NS_GET_PARENT`
   ([man 2 ioctl_ns](http://man7.org/linux/man-pages/man2/ioctl_ns.2.html)).

But `lxkns` is more than "just" a Golang package. It also features CLI tools
build on top of `lxkns`:

- `lsuns`: shows _all_ user namespaces in your Linux host, in a neat
  hierarchy. Moreover, it can also show the non-user namespaces "owned" by
  user namespaces. This ownership information is important with respect to
  capabilities and processes switching namespaces using `setns()` ([man 2
  setns](http://man7.org/linux/man-pages/man2/setns.2.html)).

- `lspns`: shows _all_ PID namespaces in your Linux host, in a neat hierarchy.

## Package Usage

The following example code runs a full namespace discovery using
`Discover(FullDiscovery)` and then prints all namespaces found, sorted by
their type, then by their ID.

```go
package main

import (
    "fmt"
    "github.com/thediveo/lxkns"
)

func main() {
    result := lxkns.Discover(lxkns.FullDiscovery)
    for nsidx := lxkns.MountNS; nsidx < lxkns.NamespaceTypesCount; nsidx++ {
        for _, ns := range result.SortedNamespaces(nsidx) {
            fmt.Println(ns.String())
        }
    }
}
```

## Copyright and License

`lxkns` is Copyright 2020 Harald Albrecht, and licensed under the Apache
License, Version 2.0.