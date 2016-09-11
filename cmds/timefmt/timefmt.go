// datefmt formats a date based on the formatting options available with
// Golang's Time.Format
//
// Copyright (c) 2016, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of datefmt nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
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
	"os"
	"path"
	"strings"
	"time"

	// My package providing additional common time formats
	"github.com/rsdoiel/shelltools/timefmt"
)

const version = timefmt.Version

var (
	showHelp    bool
	showVersion bool
	showLicense bool

	useUTC       bool
	inputFormat  = time.RFC3339
	outputFormat = time.RFC3339
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showLicense, "l", false, "display license")

	flag.BoolVar(&useUTC, "utc", false, "timestamps in UTC")
	flag.StringVar(&inputFormat, "input", inputFormat, "Set format for input")
	flag.StringVar(&outputFormat, "output", outputFormat, "Set format for output")
}

func applyConstants(s string) string {
	switch strings.ToLower(s) {
	case "ansic":
		s = time.ANSIC
	case "unixdate":
		s = time.UnixDate
	case "rubydate":
		s = time.RubyDate
	case "rfc822":
		s = time.RFC822
	case "rfc822z":
		s = time.RFC822Z
	case "rfc850":
		s = time.RFC850
	case "rfc1123":
		s = time.RFC1123
	case "rfc1123z":
		s = time.RFC1123Z
	case "rfc3339":
		s = time.RFC3339
	case "RFC3339Nano":
		s = time.RFC3339Nano
	case "kitchen":
		s = time.Kitchen
	case "stamp":
		s = time.Stamp
	case "stampmilli":
		s = time.StampMilli
	case "stampmicro":
		s = time.StampMicro
	case "stampnano":
		s = time.StampNano
	case "mysql":
		s = timefmt.MySQL
	}
	return s
}

func main() {
	appname := path.Base(os.Args[0])
	now := time.Now()
	flag.Parse()

	if showHelp == true {
		fmt.Printf(` USAGE: %s [OPTIONS] [INPUT_DATE]

 %s formats the current date or INPUT_DATE based on the output format
 provided in options. The default input and  output format is RFC3339. 
 Formats are specified based on Golang's time package including the
 common constants (e.g. RFC822, RFC1123). 

 For details see https://golang.org/pkg/time/#Time.Format.

 An additional time layouts provided by %s
 
 + mysql
 	+ "2006-01-02 15:04:05 -0700" 

 in Golang's time layout.

 EXAMPLE

     %s -output RFC822 %q

         %s

     %s -input mysql -output RFC822 %q 

	     %s

 OPTIONS

`, appname, appname, appname,
			appname, now.Format(time.RFC3339), now.Format(time.RFC822),
			appname, now.Format(timefmt.MySQL), now.Format(time.RFC822))

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("    -%s %s\n", f.Name, f.Usage)
		})

		fmt.Printf("\n\n Version %s\n", version)
		os.Exit(0)
	}

	if showVersion == true {
		fmt.Printf("\n Version %s\n", version)
		os.Exit(0)
	}

	if showLicense == true {
		fmt.Printf(`
 %s
 
 Copyright (c) 2016, R. S. Doiel
 All rights reserved.
 
 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are met:
 
 * Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.
 
 * Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.
 
 * Neither the name of datefmt nor the names of its
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

`, appname)
	}

	args := flag.Args()

	var (
		inputDate time.Time
		err       error
	)

	// Handle constants for formatting
	inputFormat = applyConstants(inputFormat)
	outputFormat = applyConstants(outputFormat)

	if len(args) > 0 {
		for i, dt := range args {
			inputDate, err = time.Parse(inputFormat, dt)
			if err != nil {
				fmt.Fprintf(os.Stderr, "can't read %s, %s\n", dt, err)
				os.Exit(1)
			}
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%s", inputDate.Format(outputFormat))
		}
		os.Exit(0)
	}
	inputDate = time.Now()
	fmt.Printf("%s", inputDate.Format(outputFormat))
}
