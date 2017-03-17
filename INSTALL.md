
# Installing plugin-to-SONA

## Prerequisite: Dep

[dep](https://github.com/golang/dep) is an up-and-coming dependency management tool,
which has the goal of being accepted as part of the standard go toolchain. It
is currently pre-alpha. However, it comes the closest to working easily out of
the box.

```sh
$ go get github.com/golang/dep
$ go install github.com/golang/dep/cmd/dep

$ dep init
$ dep ensure k8s.io/client-go@^2.0.0
```

This will set up a `vendor` directory in your current directory.

## Build
```sh
$ make build
```

## Install

### Network configuration
Create a conf file to "/etc/cni/net.d"
```sh
$ more /etc/cni/net.d/sona-cni.conf
TBD

```

### Binary
Copy the executable to "/opt/cni/bin"
```sh
# Make sure build process is done successfully.
$ sudo copy sona-cni /opt/cni/bin
$ sudo chown root /opt/cni/bin/sona-cni

```
