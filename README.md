# colortee

Colortee is a simple and naive program imitating UNIX `tee`. It outputs to tty preserving raw input, and to file with ANSI escape codes about colors filtered.

It supports appending to existing files (`-append`), and force outputting raw contents to files (`-raw`).

## Build

`go build`

## Examples

```shell
$ ls --color=always | ./colortee /tmp/test && hexdump -C /tmp/test
colortee
colortee.go
go.mod
README.md
00000000  63 6f 6c 6f 72 74 65 65  0a 63 6f 6c 6f 72 74 65  |colortee.colorte|
00000010  65 2e 67 6f 0a 67 6f 2e  6d 6f 64 0a 52 45 41 44  |e.go.go.mod.READ|
00000020  4d 45 2e 6d 64 0a                                 |ME.md.|
00000026
$ ls --color=always | ./colortee -raw /tmp/test && hexdump -C /tmp/test
colortee
colortee.go
go.mod
README.md
00000000  1b 5b 30 6d 1b 5b 30 31  3b 33 32 6d 63 6f 6c 6f  |.[0m.[01;32mcolo|
00000010  72 74 65 65 1b 5b 30 6d  0a 63 6f 6c 6f 72 74 65  |rtee.[0m.colorte|
00000020  65 2e 67 6f 0a 67 6f 2e  6d 6f 64 0a 52 45 41 44  |e.go.go.mod.READ|
00000030  4d 45 2e 6d 64 0a                                 |ME.md.|
00000036
$ ls --color=always | ./colortee -append /tmp/test && hexdump -C /tmp/test
colortee
colortee.go
go.mod
README.md
00000000  1b 5b 30 6d 1b 5b 30 31  3b 33 32 6d 63 6f 6c 6f  |.[0m.[01;32mcolo|
00000010  72 74 65 65 1b 5b 30 6d  0a 63 6f 6c 6f 72 74 65  |rtee.[0m.colorte|
00000020  65 2e 67 6f 0a 67 6f 2e  6d 6f 64 0a 52 45 41 44  |e.go.go.mod.READ|
00000030  4d 45 2e 6d 64 0a 63 6f  6c 6f 72 74 65 65 0a 63  |ME.md.colortee.c|
00000040  6f 6c 6f 72 74 65 65 2e  67 6f 0a 67 6f 2e 6d 6f  |olortee.go.go.mo|
00000050  64 0a 52 45 41 44 4d 45  2e 6d 64 0a              |d.README.md.|
0000005c
$ ls --color=auto | ./colortee -raw /tmp/test && hexdump -C /tmp/test
colortee
colortee.go
go.mod
README.md
00000000  63 6f 6c 6f 72 74 65 65  0a 63 6f 6c 6f 72 74 65  |colortee.colorte|
00000010  65 2e 67 6f 0a 67 6f 2e  6d 6f 64 0a 52 45 41 44  |e.go.go.mod.READ|
00000020  4d 45 2e 6d 64 0a                                 |ME.md.|
00000026
$ unbuffer ls --color=auto | ./colortee -raw /tmp/test && hexdump -C /tmp/test
colortee  colortee.go  go.mod  README.md
00000000  1b 5b 30 6d 1b 5b 30 31  3b 33 32 6d 63 6f 6c 6f  |.[0m.[01;32mcolo|
00000010  72 74 65 65 1b 5b 30 6d  20 20 63 6f 6c 6f 72 74  |rtee.[0m  colort|
00000020  65 65 2e 67 6f 20 20 67  6f 2e 6d 6f 64 20 20 52  |ee.go  go.mod  R|
00000030  45 41 44 4d 45 2e 6d 64  0a                       |EADME.md.|
00000039
```

As you can see, you can use `unbuffer` from package `expect` to trick programs to forcely output colors even they don't provide an option. You can install `expect` with your system package manager. For example, you can run

```shell
sudo apt install expect
```

to install `expect` in Debian/Ubuntu.

## Reference

[Wikipedia - Tee (Command)](https://en.wikipedia.org/wiki/Tee_(command)#Unix_and_Unix-like_2:~:text=ls%20%2D%2Dcolor%3Dalways%20%7C%20tee%20%3E(sed%20%22s%2F%5Cx1b%5B%5Em%5D*m%2F%2Fg%22%20%3E%20ls.txt))