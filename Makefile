#
# Simple Makefile
#

build: cmds/findfile/findfile.go
	go build -o bin/findfile cmds/findfile/findfile.go 
	go build -o bin/finddir	 cmds/finddir/finddir.go 

clean: bin/findfile
	rm bin/findfile
	rm bin/finddir

install: cmds/findfile/findfile.go
	go install cmds/findfile/findfile.go
	go install cmds/finddir/finddir.go

