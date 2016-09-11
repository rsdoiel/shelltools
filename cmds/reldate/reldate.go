/**
 * Generates a date in YYYY-MM-DD format based on a relative time
 * description (e.g. -1 week, +3 years)
 *
 * @author R. S. Doiel, <rsdoiel@gmail.com>
 * copyright (c) 2014 all rights reserved.
 * Released under the Simplified BSD License
 * See: http://opensource.org/licenses/bsd-license.php
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	// Local package
	"github.com/rsdoiel/shelltools/reldate"
)

var (
	showHelp    bool
	showVersion bool
	showLicense bool

	endOfMonthFor bool
	relativeTo    string
	relativeT     time.Time
)

var usage = func(exit_code int, appName, version, msg string) {
	var fh = os.Stderr
	if exit_code == 0 {
		fh = os.Stdout
	}
	fmt.Fprintf(fh, `
USAGE %s [TIME_INCREMENT TIME_UNIT|WEEKDAY_NAME]
    

EXAMPLES
 
+ Two days from today: %s 2 days
+ Three weeks ago: %s -- -3 weeks
+ Three weeks from 2014-01-01: %s --from=2014-01-01 3 weeks
+ Three days before 2014-01-01: %s --from=2014-01-01 -- -3 days
+ The Friday of this week: %s Friday
+ The Monday in week containing 2015-02-06: %s --from=2015-02-06 Monday

Time increments are a positive or negative integer. Time unit can be
either day(s), week(s), month(s), or year(s). Weekday names are
case insentive (e.g. Monday and monday). They can be abbreviated
to the first three letters of the name, e.g. Sunday can be Sun, Monday
can be Mon, Tuesday can be Tue, Wednesday can be Wed, Thursday can
be Thu, Friday can be Fri or Saturday can be Sat.

OPTIONS

`, msg, appName, appName, appName, appName, appName, appName)

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(fh, "\t-%s\t(defaults to %s) %s\n", f.Name, f.Value, f.Usage)
	})

	fmt.Fprintf(fh, `
%s

copyright (c) 2014 all rights reserved.
Released under the Simplified BSD License
See: http://opensource.org/licenses/bsd-license.php

`, version)
	os.Exit(exit_code)
}

var license = func(exit_code int, appName, version string) {
	var fh = os.Stderr
	if exit_code == 0 {
		fh = os.Stdout
	}
	fmt.Fprintf(fh, `
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
	os.Exit(exit_code)
}

func init() {
	const (
		relativeToUsage = "Date the relative time is calculated from."
		helpUsage       = "Display this help document."
		endOfMonthUsage = "Display the end of month day. E.g. 2012-02-29"
	)

	flag.BoolVar(&showHelp, "help", false, helpUsage)
	flag.BoolVar(&showHelp, "h", false, helpUsage)
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")

	flag.StringVar(&relativeTo, "from", relativeTo, relativeToUsage)
	flag.StringVar(&relativeTo, "f", relativeTo, relativeToUsage)
	flag.BoolVar(&endOfMonthFor, "end-of-month", endOfMonthFor, endOfMonthUsage)
	flag.BoolVar(&endOfMonthFor, "e", endOfMonthFor, endOfMonthUsage)
}

func assertOk(e error, failMsg string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, " %s\n %s\n", failMsg, e)
		os.Exit(1)
	}
}

func main() {
	var (
		err        error
		unitString string
	)
	appName := path.Base(os.Args[0])

	flag.Parse()
	if showHelp == true {
		usage(0, appName, reldate.Version, "")
	}
	if showLicense == true {
		license(0, appName, reldate.Version)
	}
	if showVersion == true {
		fmt.Fprintf(os.Stdout, "%s version %s\n", appName, reldate.Version)
		os.Exit(0)
	}

	argc := flag.NArg()
	argv := flag.Args()

	if argc < 1 && endOfMonthFor == false {
		usage(1, appName, reldate.Version, "Missing time increment and units (e.g. +2 days) or weekday name (e.g. Monday, Mon).\n")
	} else if argc > 2 {
		usage(1, appName, reldate.Version, "Too many command line arguments.\n")
	}

	relativeT = time.Now()
	if relativeTo != "" {
		relativeT, err = time.Parse(reldate.YYYYMMDD, relativeTo)
		assertOk(err, "Cannot parse the from date.\n")
	}

	if endOfMonthFor == true {
		fmt.Println(reldate.EndOfMonth(relativeT))
		os.Exit(0)
	}

	timeInc := 0
	if argc == 2 {
		unitString = strings.ToLower(argv[1])
		timeInc, err = strconv.Atoi(argv[0])
		assertOk(err, "Time increment should be a positive or negative integer.\n")
	} else {
		// We may have a weekday string
		unitString = strings.ToLower(argv[0])
	}
	t, err := reldate.RelativeTime(relativeT, timeInc, unitString)
	assertOk(err, "Did not understand command.")
	fmt.Println(t.Format(reldate.YYYYMMDD))
}
