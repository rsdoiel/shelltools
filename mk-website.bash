#!/bin/bash

PROJECT="shelltools"

function checkApp() {
    APP_NAME=$(which $1)
    if [ "$APP_NAME" = "" ] && [ ! -f "./bin/$1" ]; then
        echo "Missing $APP_NAME"
        exit 1
    fi
}

function softwareCheck() {
    for APP_NAME in $@; do
        checkApp $APP_NAME
    done
}

function MakePage () {
    nav="$1"
    content="$2"
    html="$3"
    # Always use the latest compiled mkpage
    APP=$(which mkpage)
    if [ -f ./bin/mkpage ]; then
        APP="./bin/mkpage"
    fi

    echo "Rendering $html"
    $APP \
	"title=text:$PROJECT -- a small collection of file and shell utilities" \
        "nav=$nav" \
        "content=$content" \
	    "sitebuilt=text:Updated $(date)" \
        "copyright=copyright.md" \
        page.tmpl > $html
}

echo "Checking necessary software is installed"
softwareCheck mkpage
echo "Generating website index.html"
MakePage nav.md README.md index.html
echo "Generating install.html"
MakePage nav.md INSTALL.md install.html
echo "Generating finddir.html"
MakePage nav.md finddir.md finddir.html
echo "Generating findfile.html"
MakePage nav.md findfile.md findfile.html
echo "Generating pathparts.html"
MakePage nav.md pathparts.md pathparts.html
echo "Generating mergepath.html"
MakePage nav.md mergepath.md mergepath.html
echo "Generating reldate.html"
MakePage nav.md reldate.md reldate.html
echo "Generating range.html"
MakePage nav.md range.md range.html
echo "Generating license.html"
MakePage nav.md "markdown:$(cat LICENSE)" license.html

