#
# Simple Makefile
#
PROJECT = shelltools

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\"  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

build: bin/csvcols bin/findfile bin/finddir bin/jsoncols bin/pathparts bin/mergepath bin/reldate bin/range bin/timefmt bin/urlparse

bin/csvcols: shelltools.go cmds/csvcols/csvcols.go
	go build -o bin/csvcols cmds/csvcols/csvcols.go

bin/findfile: shelltools.go cmds/findfile/findfile.go
	go build -o bin/findfile cmds/findfile/findfile.go 

bin/finddir: shelltools.go cmds/finddir/finddir.go
	go build -o bin/finddir cmds/finddir/finddir.go 

bin/jsoncols: shelltools.go cmds/jsoncols/jsoncols.go
	go build -o bin/jsoncols cmds/jsoncols/jsoncols.go

bin/pathparts: shelltools.go cmds/pathparts/pathparts.go
	go build -o bin/pathparts cmds/pathparts/pathparts.go 

bin/mergepath: shelltools.go cmds/mergepath/mergepath.go
	go build -o bin/mergepath cmds/mergepath/mergepath.go 

bin/reldate: shelltools.go cmds/reldate/reldate.go
	go build -o bin/reldate cmds/reldate/reldate.go 

bin/range: shelltools.go cmds/range/range.go
	go build -o bin/range cmds/range/range.go 

bin/timefmt: shelltools.go cmds/timefmt/timefmt.go
	go build -o bin/timefmt cmds/timefmt/timefmt.go 

bin/urlparse: shelltools.go cmds/urlparse/urlparse.go
	go build -o bin/urlparse cmds/urlparse/urlparse.go 

website:
	./mk-website.bash

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

refresh:
	git fetch origin
	git pull origin $(BRANCH)

publish:
	./mk-website.bash
	./publish.bash

clean: 
	if [ -d bin ]; then /bin/rm -fR bin; fi
	if [ -d dist ]; then /bin/rm -fR dist; fi
	if [ -f $(PROJECT)-$(VERSION)-release.zip ]; then rm -f $(PROJECT)-$(VERSION)-release.zip; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/csvcols/csvcols.go
	env GOBIN=$(HOME)/bin go install cmds/findfile/findfile.go
	env GOBIN=$(HOME)/bin go install cmds/finddir/finddir.go
	env GOBIN=$(HOME)/bin go install cmds/jsoncols/jsoncols.go
	env GOBIN=$(HOME)/bin go install cmds/pathparts/pathparts.go
	env GOBIN=$(HOME)/bin go install cmds/mergepath/mergepath.go
	env GOBIN=$(HOME)/bin go install cmds/reldate/reldate.go
	env GOBIN=$(HOME)/bin go install cmds/range/range.go
	env GOBIN=$(HOME)/bin go install cmds/timefmt/timefmt.go
	env GOBIN=$(HOME)/bin go install cmds/urlparse/urlparse.go

dist/linux-amd64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/csvcols cmds/csvcols/csvcols.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/findfile cmds/findfile/findfile.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/finddir cmds/finddir/finddir.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/jsoncols cmds/jsoncols/jsoncols.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/pathparts cmds/pathparts/pathparts.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/mergepath cmds/mergepath/mergepath.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/reldate cmds/reldate/reldate.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/range cmds/range/range.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/timefmt cmds/timefmt/timefmt.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/urlparse cmds/urlparse/urlparse.go

dist/macosx-amd64:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/csvcols cmds/csvcols/csvcols.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/findfile cmds/findfile/findfile.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/finddir cmds/finddir/finddir.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/jsoncols cmds/jsoncols/jsoncols.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/pathparts cmds/pathparts/pathparts.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/mergepath cmds/mergepath/mergepath.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/reldate cmds/reldate/reldate.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/range cmds/range/range.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/timefmt cmds/timefmt/timefmt.go
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/macosx-amd64/urlparse cmds/urlparse/urlparse.go

dist/windows-amd64:
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/csvcols.exe cmds/csvcols/csvcols.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/findfile.exe cmds/findfile/findfile.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/finddir.exe cmds/finddir/finddir.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/jsoncols.exe cmds/jsoncols/jsoncols.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/pathparts.exe cmds/pathparts/pathparts.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/mergepath.exe cmds/mergepath/mergepath.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/reldate.exe cmds/reldate/reldate.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/range.exe cmds/range/range.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/timefmt.exe cmds/timefmt/timefmt.go
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/urlparse.exe cmds/urlparse/urlparse.go

dist/raspbian-arm7:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/csvcols cmds/csvcols/csvcols.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/findfile cmds/findfile/findfile.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/finddir cmds/finddir/finddir.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/jsoncols cmds/jsoncols/jsoncols.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/pathparts cmds/pathparts/pathparts.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/mergepath cmds/mergepath/mergepath.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/reldate cmds/reldate/reldate.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/range cmds/range/range.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/timefmt cmds/timefmt/timefmt.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspbian-arm7/urlparse cmds/urlparse/urlparse.go

dist/raspbian-arm6:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/csvcols cmds/csvcols/csvcols.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/findfile cmds/findfile/findfile.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/finddir cmds/finddir/finddir.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/jsoncols cmds/jsoncols/jsoncols.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/pathparts cmds/pathparts/pathparts.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/mergepath cmds/mergepath/mergepath.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/reldate cmds/reldate/reldate.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/range cmds/range/range.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/timefmt cmds/timefmt/timefmt.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspbian-arm6/urlparse cmds/urlparse/urlparse.go

dist:
	mkdir -p dist

dist/README.md: README.md
	cp -v README.md dist/

dist/LICENSE: LICENSE
	cp -v LICENSE dist/

dist/INSTALL.md: INSTALL.md
	cp -v INSTALL.md dist/

dist/demo:
	cp -vR demo dist/

targets: dist/linux-amd64 dist/macosx-amd64 dist/windows-amd64 dist/raspbian-arm7 dist/raspbian-arm6

docs: dist dist/README.md dist/LICENSE dist/INSTALL.md dist/demo 

zip: $(PROJECT)-$(VERSION)-release.zip 
	zip -r $(PROJECT)-$(VERSION)-release.zip dist/

release: targets docs zip
