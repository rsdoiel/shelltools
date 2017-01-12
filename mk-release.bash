#!/bin/bash
#
# Make releases for Linux/amd64, Linux/ARM7 (Raspberry Pi), Windows, and Mac OX X (darwin)
#
PROJECT=shelltools

PROG_LIST="csvcols findfile finddir mergepath pathparts range reldate timefmt urlparse"

VERSION=$(grep -m1 'Version = ' $PROJECT.go | cut -d\"  -f 2)

RELEASE_NAME=$PROJECT-$VERSION

echo "Preparing $RELEASE_NAME"
for PROGNAME in $PROG_LIST; do
    echo "Cross compiling $PROGNAME"
    env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
    env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspberrypi-arm6/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
    env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
    env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/$PROGNAME.exe cmds/$PROGNAME/$PROGNAME.go
    env CGO_ENABLED=0 GOOS=darwin	GOARCH=amd64 go build -o dist/macosx-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
done
echo "Assembling dist/"
for FNAME in README.md LICENSE INSTALL.md; do
    cp -v $FNAME dist/
done
echo "Zipping $RELEASE_NAME"
zip -r $RELEASE_NAME-release.zip dist/*

