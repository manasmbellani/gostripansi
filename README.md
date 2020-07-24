Strip ANSI
==========

This Go package removes ANSI escape codes from strings provided in the input.

This work is based on the Stripping logic written by `acarl005` in repo:
`github.com/acarl005/gostripansi`

## Install

```
$ go get -u github.com/manasmbellani/gostripansi
```

## Usage

To strip coloured output from /tmp/test.txt file
```
$ cat /tmp/test.txt | gostripansi
```