/**
 * mergepath.go - merge the path variable to avoid duplicates
 *
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2014 All rights reserved.
 * @license released under the Simplified BSD License
 * see: http://opensource.org/licenses/bsd-license.php
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	envPath     string
	dir         string
	showHelp    bool
	appendPath  = true
	prependPath = false
	clipPath    = false
)

var usage = func(exit_code int, msg string) {
	var fh = os.Stderr

	if exit_code == 0 {
		fh = os.Stdout
	}
	fmt.Fprintf(fh, `%s
 USAGE %s [OPTIONS] PATH_TO_ADD PATH_TO_MODIFY

 EXAMPLES

    append work directory to existing path: %s . $PATH
    prepend working directory to existing path: %s -P . $PATH

 OPTIONS

`, msg, os.Args[0], os.Args[0], os.Args[0])
	flag.VisitAll(func(f *flag.Flag) {
		if len(f.Name) > 1 {
			fmt.Fprintf(fh, "    -%s, -%s\t%s\n", f.Name[0:1], f.Name, f.Usage)
		}
	})
	fmt.Fprintf(fh, `

 Copyright (c) 2014 All rights reserved.
 Released under the Simplified BSD License
 See: http://opensource.org/licenses/bsd-license.php
`)
	os.Exit(exit_code)
}

func init() {
	const (
		pathUsage    = "The path you want to merge with."
		dirUsage     = "The directory you want to add to the path."
		appendUsage  = "Append the directory to the path removing any duplication"
		prependUsage = "Prepend the directory to the path removing any duplication"
		clipUsage    = "Remove a directory from the path"
		helpUsage    = "This help document."
	)
	flag.BoolVar(&showHelp, "h", false, helpUsage)
	flag.BoolVar(&showHelp, "help", false, helpUsage)

	envPath = "$PATH"

	flag.StringVar(&envPath, "e", envPath, pathUsage)
	flag.StringVar(&envPath, "envpath", envPath, pathUsage)
	flag.StringVar(&dir, "d", dir, dirUsage)
	flag.StringVar(&dir, "directory", dir, dirUsage)

	flag.BoolVar(&appendPath, "a", appendPath, appendUsage)
	flag.BoolVar(&appendPath, "append", appendPath, appendUsage)
	flag.BoolVar(&prependPath, "p", prependPath, prependUsage)
	flag.BoolVar(&prependPath, "prepend", prependPath, prependUsage)
	flag.BoolVar(&clipPath, "c", clipPath, clipUsage)
	flag.BoolVar(&clipPath, "clip", clipPath, clipUsage)
}

func clip(envPath string, dir string) string {
	oParts := []string{}
	iParts := strings.Split(envPath, ":")
	for _, v := range iParts {
		if v != dir {
			oParts = append(oParts, v)
		}
	}
	return strings.Join(oParts, ":")
}

func main() {
	flag.Parse()

	if showHelp == true {
		usage(0, "")
	}

	if flag.NArg() > 0 {
		dir = flag.Arg(0)
		if flag.NArg() == 2 {
			envPath = flag.Arg(1)
		}
	}

	if envPath == "$PATH" {
		envPath = os.Getenv("PATH")
	}
	if dir == "" {
		usage(1, "Missing directory to add to path")
	}
	if clipPath == true {
		fmt.Printf("%s", clip(envPath, dir))
		os.Exit(0)
	}
	if prependPath == true {
		appendPath = false
	}
	if strings.Contains(envPath, dir) {
		envPath = clip(envPath, dir)
	}
	if appendPath == true {
		fmt.Printf("%s:%s", envPath, dir)
	} else {
		fmt.Printf("%s:%s", dir, envPath)
	}
}
