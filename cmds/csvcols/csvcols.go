//
// csvcols - is a command line that takes each argument in order and outputs a line in CSV format.
// It can also take a delimiter and line of text splitting it into a CSV formatted set of columns.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright (c) 2017, R. S. Doiel
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
// * Neither the name of findfile nor the names of its
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
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	// My packages
	"github.com/rsdoiel/cli"
	"github.com/rsdoiel/shelltools"
)

var (
	usage = `USAGE: %s [OPTIONS] ARGS_AS_COLS`

	description = `
SYNOPSIS

%s converts a set of command line args into columns output in CSV formator filters CSV columns from stdin.
`

	examples = `
EXAMPLES

Simple usage of building a CSV file one row at a time.

    %s one two three > 3col.csv
    %s 1 2 3 >> 3col.csv
    cat 3col.csv

Example parsing a pipe delimited string into a CSV line

    %s -d "|" "one|two|three" > 3col.csv
    %s -delimiter "|" "1|2|3" >> 3col.csv
    cat 3col.csv

Filter a 5 columns CSV file saving columns 2 and 4
in 2col.csv.

    cat 5col.csv | %s -f 2 4 > 2col.csv
`

	// Basic Options
	showHelp    bool
	showLicense bool
	showVersion bool

	// App Options
	delimiter     string
	filterColumns bool
)

func selectedColumns(record []string, columnNos []int) []string {
	result := []string{}
	l := len(record)
	for _, col := range columnNos {
		if col >= 0 && col < l {
			result = append(result, record[col])
		} else {
			// If we don't find the column, story an empty string
			result = append(result, "")
		}
	}
	return result
}

func printCSVColumns(columnNos []int) {
	r := csv.NewReader(os.Stdin)
	w := csv.NewWriter(os.Stdout)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read stdin for CSV data, %s", err)
		os.Exit(1)
	}
	for _, record := range records {
		row := selectedColumns(record, columnNos)
		if err := w.Write(row); err != nil {
			fmt.Fprintf(os.Stderr, "error writing record to csv: %s\n", err)
			os.Exit(1)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func init() {
	// Basic Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")

	// App Options
	flag.StringVar(&delimiter, "d", "", "set delimiter for conversion")
	flag.StringVar(&delimiter, "delimiter", "", "set delimiter for conversion")
	flag.BoolVar(&filterColumns, "f", false, "filter CSV input for columns requested")
	flag.BoolVar(&filterColumns, "filter-columns", false, "filter CSV input for columns requested")

}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, "csvcols", fmt.Sprintf(shelltools.LicenseText, appName, shelltools.Version), shelltools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionsText = "OPTIONS\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName, appName, appName)

	if showHelp == true {
		fmt.Println(cfg.Usage())
		os.Exit(0)
	}

	if showLicense == true {
		fmt.Println(cfg.License())
		os.Exit(0)
	}

	if showVersion == true {
		fmt.Println(cfg.Version())
		os.Exit(0)
	}

	if filterColumns == true {
		columnNos := []int{}
		for _, arg := range args {
			i, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Expected a column number (0 - n), %q, %s", arg, err)
				os.Exit(1)
			}
			columnNos = append(columnNos, i)
		}
		printCSVColumns(columnNos)
	}

	if len(delimiter) > 0 && len(args) == 1 {
		args = strings.Split(args[0], delimiter)
	}

	// Clean up fields removing outer quotes if necessary
	fields := []string{}
	for _, val := range args {
		if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
			val = strings.TrimPrefix(strings.TrimSuffix(val, "\""), "\"")
		}
		fields = append(fields, strings.TrimSpace(val))
	}

	out := csv.NewWriter(os.Stdout)
	if err := out.Write(fields); err != nil {
		log.Fatalf("error wrint args as csv, %s", err)
	}
	out.Flush()
	if err := out.Error(); err != nil {
		log.Fatal(err)
	}

}
