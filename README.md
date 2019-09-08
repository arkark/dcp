dcp
===

An alias of `docker container cp` and useful completions.

```sh
$ dcp <TAB><TAB>
-            .gitignore   README.md    completion/  debug/       fuga:        go.sum       internal/
.git/        Makefile     bar:         dcp          foo:         go.mod       hoge:        main.go
$ dcp bar:<TAB><TAB>
bin/    dev/    etc/    home/   lib/    media/  mnt/    opt/    proc/   root/   run/    sbin/   srv/    sys/    tmp/    usr/    var/
$ dcp bar:etc/<TAB><TAB>
etc/alpine-release   etc/group            etc/issue            etc/motd             etc/passwd           etc/resolv.conf      etc/ssl/
etc/apk/             etc/hostname         etc/logrotate.d/     etc/mtab             etc/periodic/        etc/securetty        etc/sysctl.conf
etc/conf.d/          etc/hosts            etc/modprobe.d/      etc/network/         etc/profile          etc/services         etc/sysctl.d/
etc/crontabs/        etc/init.d/          etc/modules          etc/opt/             etc/profile.d/       etc/shadow           etc/udhcpd.conf
etc/fstab            etc/inittab          etc/modules-load.d/  etc/os-release       etc/protocols        etc/shells
$ dcp bar:etc/alpine-release .
$ cat alpine-release
3.10.2
```

## Prerequisites

- Docker commands are executable without sudo.
- Bash completion is installed.

## Installation

### Binary

```sh
$ go get -u github.com/ArkArk/dcp
```

### Completion

#### Linux

```sh
$ sudo curl -L https://raw.githubusercontent.com/ArkArk/dcp/master/completion/dcp -o /etc/bash_completion.d/dcp
```

#### Mac

```sh
$ sudo curl -L https://raw.githubusercontent.com/ArkArk/dcp/master/completion/dcp -o /usr/local/etc/bash_completion.d/dcp
```

## Supported shells

- [x] bash
- [ ] zsh
- [ ] fish
