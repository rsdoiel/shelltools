#
# Simple Makefile
#

build:
	go build -o bin/findfile cmds/findfile/findfile.go 
	go build -o bin/finddir	 cmds/finddir/finddir.go 
	go build -o bin/pathparts cmds/pathparts/pathparts.go 
	go build -o bin/mergepath	 cmds/mergepath/mergepath.go 
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
	if [ -f fsutils-binary-release.zip ]; then rm -f fsutils-binary-release.zip; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/findfile/findfile.go
	env GOBIN=$(HOME)/bin go install cmds/finddir/finddir.go

release:
	./mk-release.bash
