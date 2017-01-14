//
// jsonpath is a command line tool for filter JSON data from standard in or specified files.
// It was inspired by [jq](https://github.com/stedolan/jq) and [jid](https://github.com/simeji/jid).
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	// 3rd Party packages
	//	"github.com/simeji/jid"
	"github.com/rsdoiel/jid"

	// My Packages
	"github.com/rsdoiel/cli"
	"github.com/rsdoiel/shelltools"
)

var (
	usage = `USAGE: %s [OPTIONS] [EXPRESSION] [INPUT_FILENAME] [OUTPUT_FILENAME]`

	description = `
SYSNOPSIS

%s provides for both interactive exploration of JSON structures like jid 
and command line scripting flexibility for data extraction into delimited
columns. This is helpful in flattening content extracted from JSON blobs.
The default delimiter for each value extracted is a comma. This can be
overridden with an option.

+ EXPRESSION can be an empty stirng or dot notation for an object's path
+ INPUT_FILENAME is the filename to read or a dash "-" if you want to 
  explicity read from stdin
	+ if not provided then %s reads from stdin
+ OUTPUT_FILENAME is the filename to write or a dash "-" if you want to 
  explicity write to stdout
	+ if not provided then %s write to stdout
`

	examples = `
EXAMPLES

If myblob.json contained

{"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}

Getting just the name could be done with

    %s .name myblob.json

This would yeild

    "Doe, Jane"

Flipping .name and .age into pipe delimited columns is as 
easy as listing each field in the expression inside a 
space delimited string.

    %s -d\|  ".name .age" myblob.json

This would yeild

    "Doe, Jane"|42
`

	// Basic Options
	showHelp    bool
	showLicense bool
	showVersion bool

	// Application Specific Options
	monochrome     bool
	runInteractive bool
	delimiter      = ","
	expression     string
	inputFName     string
	outputFName    string
)

func init() {
	// Basic Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")

	// Application Specific Options
	flag.BoolVar(&monochrome, "m", false, "display output in monochrome")
	flag.BoolVar(&runInteractive, "i", false, "run interactively")
	flag.StringVar(&delimiter, "d", delimiter, "set the delimiter for multi-field output")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, "JSONQUERY", fmt.Sprintf(shelltools.LicenseText, appName, shelltools.Version), shelltools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName, appName, appName)
	cfg.OptionsText = "OPTIONS\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName, appName)

	//NOTE: Need to handle JSONQUERY_MONOCHROME setting
	monochrome = cfg.MergeEnvBool("monochrome", monochrome)

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

	if runInteractive == true {
		expression = "."
	}

	var (
		in  *os.File
		out *os.File
		err error
	)

	// Handle ordered args.
	for _, arg := range args {
		switch {
		case len(expression) == 0:
			if len(arg) == 0 {
				arg = "."
			}
			expression = arg
		case len(inputFName) == 0:
			if strings.Compare(arg, "-") != 0 {
				inputFName = arg
			}
		case len(outputFName) == 0:
			if strings.Compare(arg, "-") != 0 {
				outputFName = arg
			}
		}
	}

	// FIXME: figure out if this is an interactive session or content read in from a file.
	if len(inputFName) == 0 || strings.Compare(inputFName, "-") == 0 {
		in = os.Stdin
	} else {
		in, err = os.Open(inputFName)
		if err != nil {
			log.Fatalln(err)
		}
		defer in.Close()
	}
	if len(outputFName) == 0 || strings.Compare(outputFName, "-") == 0 {
		out = os.Stdout
	} else {
		out, err = os.Create(outputFName)
		if err != nil {
			log.Fatalln(err)
		}
		defer out.Close()
	}

	// Make sure we are ready to run the engine, else display help with error level
	expression = strings.TrimSpace(expression)
	if len(expression) == 0 {
		fmt.Println(cfg.Usage())
		fmt.Println("Missing expression")
		os.Exit(1)
	}

	// Configure the jid engine
	engineAttributes := &jid.EngineAttribute{
		DefaultQuery: expression,
		Monochrome:   monochrome,
	}

	// Run the jid engine appropriately
	engine, err := jid.NewEngine(in, engineAttributes)
	if err != nil {
		log.Fatalln(err)
	}
	var result jid.EngineResultInterface

	if runInteractive == true {
		result = engine.Run()
		if err := result.GetError(); err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		//FIXME: should honor query mode
		fmt.Fprintf(out, "%s", result.GetContent())
	} else {
		qrys := strings.Split(expression, " ")
		for i, qry := range qrys {
			result = engine.EvalString(qry)
			if err := result.GetError(); err != nil {
				log.Fatalln(err)
				os.Exit(1)
			}
			if i > 0 {
				fmt.Fprintf(out, "%s", delimiter)
			}
			fmt.Fprintf(out, "%s", result.GetContent())
		}
	}
}
