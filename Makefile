#
# Simple Makefile
#
PROJECT = shelltools

build:
	go build -o bin/findfile cmds/findfile/findfile.go 
	go build -o bin/finddir cmds/finddir/finddir.go 
	go build -o bin/pathparts cmds/pathparts/pathparts.go 
	go build -o bin/mergepath cmds/mergepath/mergepath.go 
	go build -o bin/reldate cmds/reldate/reldate.go 
	go build -o bin/range cmds/range/range.go 
	./mk-website.bash

website:
	./mk-website.bash

save:
	./mk-website.bash
	git commit -am "quick save"
	git push origin master

publish:
	./mk-website.bash
	./publish.bash

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -f $(PROJECT)-binary-release.zip ]; then rm -f $(PROJECT)-binary-release.zip; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/findfile/findfile.go
	env GOBIN=$(HOME)/bin go install cmds/finddir/finddir.go
	env GOBIN=$(HOME)/bin go install cmds/pathparts/pathparts.go
	env GOBIN=$(HOME)/bin go install cmds/mergepath/mergepath.go
	env GOBIN=$(HOME)/bin go install cmds/reldate/reldate.go
	env GOBIN=$(HOME)/bin go install cmds/range/range.go

release:
	./mk-release.bash
