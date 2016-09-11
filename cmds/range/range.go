/**
 * range.go - emit a list of integers separated by spaces starting from
 * first command line parameter to last command line parameter.
 *
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2014 all rights reserved.
 * Released under the Simplified BSD License
 * See: http://opensource.org/licenses/bsd-license.php
 */
package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

const version = "v1.0.1"

var (
	showVersion bool
	showHelp    bool
	showLicense bool

	start         int
	end           int
	increment     int
	randomElement bool
)

var usage = func(exit_code int, msg string) {
	var (
		fh      = os.Stderr
		appname = os.Args[0]
	)

	if exit_code == 0 {
		fh = os.Stdout
	}
	fmt.Fprintf(fh, `%s
 USAGE %s STARTING_INTEGER ENDING_INTEGER [INCREMENT_INTEGER]

 EXAMPLES

 Count from one through five: 
     %s 1 5
 Count from negative two to six: 
     %s -- -2 6
 Count even numbers from two to ten: 
     %s --increment=2 2 10
 Count down from ten to one: 
     %s 10 1
 Pick a random number in range one and ten:
     %s -r 1 10
 Pick a random even number in range two to twelve:
     %s 12 --random --increment=2 2 12

 OPTIONS

`, msg, appname, appname, appname, appname, appname, appname, appname)

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(fh, "\t-%s\t(defaults to %s) %s\n", f.Name, f.DefValue, f.Usage)
	})

	fmt.Fprintf(fh, `
 Version %s

 copyright (c) 2014 all rights reserved.
 Released under the Simplified BSD License
 See: http://opensource.org/licenses/bsd-license.php

`, version)
	os.Exit(exit_code)
}

func init() {
	const (
		helpUsage    = "Display this help document."
		versionUsage = "Display version"
		licenseUsage = "Display license"
		startUsage   = "The starting integer."
		endUsage     = "The ending integer."
		incUsage     = "The non-zero integer increment value."
	)

	flag.IntVar(&start, "start", 0, startUsage)
	flag.IntVar(&start, "s", 0, startUsage)
	flag.IntVar(&end, "end", 0, endUsage)
	flag.IntVar(&end, "e", 0, endUsage)
	flag.IntVar(&increment, "increment", 1, incUsage)
	flag.IntVar(&increment, "i", 1, incUsage)
	flag.BoolVar(&randomElement, "r", false, "Pick a range value from range")
	flag.BoolVar(&randomElement, "random", false, "Pick a range value from range")

	flag.BoolVar(&showHelp, "help", showHelp, helpUsage)
	flag.BoolVar(&showHelp, "h", showHelp, helpUsage)
	flag.BoolVar(&showVersion, "v", showVersion, versionUsage)
	flag.BoolVar(&showVersion, "version", showVersion, versionUsage)
	flag.BoolVar(&showLicense, "l", showLicense, licenseUsage)
	flag.BoolVar(&showLicense, "license", showLicense, licenseUsage)
}

func assertOk(e error, failMsg string) {
	if e != nil {
		usage(1, fmt.Sprintf(" %s\n %s\n", failMsg, e))
	}
}

func inRange(i, start, end int) bool {
	if start <= end && i <= end {
		return true
	}
	if start >= end && i >= end {
		return true
	}
	return false
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	if showHelp == true {
		usage(0, "")
	}
	if showVersion == true {
		fmt.Printf("Version %s\n", version)
		os.Exit(0)
	}
	if showLicense == true {
		fmt.Printf(`
%s %s

Copyright (c) 2014-2016, R. S. Doiel
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of findfile nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`, appName, version)
		os.Exit(0)
	}

	argc := flag.NArg()
	argv := flag.Args()

	if argc < 2 {
		usage(1, "Must include start and end integers.")
	} else if argc > 3 {
		usage(1, "Too many command line arguments.")
	}

	start, err := strconv.Atoi(argv[0])
	assertOk(err, "Start value must be an integer.")
	end, err := strconv.Atoi(argv[1])
	assertOk(err, "End value must be an integer.")
	if argc == 3 {
		increment, err = strconv.Atoi(argv[2])
	} else if increment == 0 {
		err = errors.New("increment was zero")
	}
	assertOk(err, "Increment must be a non-zero integer.")

	if start == end {
		fmt.Printf("%d", start)
		os.Exit(0)
	}

	// Normalize to a positive value.
	if start <= end && increment < 0 {
		increment = increment * -1
	}
	if start > end && increment > 0 {
		increment = increment * -1
	}

	// If randonElement than generate range and pick the ith random element from range
	var (
		ithArray []int
		ith      = 0
	)

	// Now count up or down as appropriate.
	for i := start; inRange(i, start, end) == true; i = i + increment {
		if randomElement == true {
			ithArray = append(ithArray, i)
		} else {
			if i == start {
				fmt.Printf("%d", i)
			} else {
				fmt.Printf(" %d", i)
			}
		}
	}
	// if randomElement we should an array we can pick the elements from
	if randomElement == true {
		rand.Seed(time.Now().Unix())
		ith = rand.Intn(len(ithArray))
		fmt.Printf("%d", ithArray[ith])
	}
}
