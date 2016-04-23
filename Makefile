#
# Simple Makefile
#

build: cmds/findfile/findfile.go
	go build -o bin/findfile cmds/findfile/findfile.go 

clean: bin/findfile
	rm bin/findfile

install: cmds/findfile/findfile.go
	go install cmds/findfile/findfile.go

