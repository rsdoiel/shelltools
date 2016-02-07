//
// findfile.go - a simple directory tree walker that looks for files
// by name, basename or extension. Basically a unix "find" light to
// demonstrate walking the file system
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright (c) 2016, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//  list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//  this list of conditions and the following disclaimer in the documentation
//  and/or other materials provided with the distribution.
//
// * Neither the name of findfile nor the names of its
//  contributors may be used to endorse or promote products derived from
//  this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const version = "0.0.0"

var (
	help                bool
	showVersion         bool
	findDirectoriesOnly bool
	findFilesOnly       bool
	findPrefix          bool
	findSuffix          bool
	stopOnErrors        bool
	outputFullPath      bool
)

func display(docroot, p string) {
	var s string
	if outputFullPath == true {
		s, _ = filepath.Abs(p)
	} else {
		s, _ = filepath.Rel(docroot, p)
	}
	fmt.Printf("%s\n", s)
}

func walkPath(docroot string, target string) error {
	return filepath.Walk(docroot, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			if stopOnErrors == true {
				return fmt.Errorf("Can't read %s, %s", p, err)
			}
			return nil
		}
		if findDirectoriesOnly == true && info.IsDir() == false {
			return nil
		}
		if findFilesOnly == true && info.Mode().IsRegular() == false {
			return nil
		}
		s := filepath.Base(p)
		switch {
		case findPrefix == true && strings.HasPrefix(s, target) == true:
			display(docroot, p)
		case findSuffix == true && strings.HasSuffix(s, target) == true:
			display(docroot, p)
		case strings.Compare(s, target) == 0:
			display(docroot, p)
		}
		return nil
	})
}

func init() {
	flag.BoolVar(&help, "h", false, "display this help message")
	flag.BoolVar(&showVersion, "v", false, "display version message")
	flag.BoolVar(&findDirectoriesOnly, "d", false, "find directories only")
	flag.BoolVar(&findFilesOnly, "f", false, "find files only")
	flag.BoolVar(&stopOnErrors, "e", false, "Stop walk on file system errors (e.g. permissions)")
	flag.BoolVar(&findPrefix, "p", false, "find file(s) based on basename prefix")
	flag.BoolVar(&findSuffix, "s", false, "find file(s) based on basename suffix")
	flag.BoolVar(&outputFullPath, "F", false, "list full path for files found")
}

func main() {
	target := ""
	flag.Parse()
	args := flag.Args()

	if showVersion == true {
		fmt.Printf("Version %s\n", version)
		os.Exit(0)
	}

	if help == true || len(args) == 0 {
		fmt.Println("USAGE findfile [OPTIONS] TARGET_FILENAME [DIRECTORIES_TO_SEARCH]")
		flag.PrintDefaults()
		fmt.Printf("Version %s\n", version)
		if len(args) == 0 {
			os.Exit(1)
		}
		os.Exit(0)
	}

	if len(args) == 1 {
		err := walkPath(".", args[0])
		if err != nil {
			log.Fatal(err)
		}
	}

	for i, dir := range args {
		if i == 0 {
			target = dir
		} else {
			err := walkPath(dir, target)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
