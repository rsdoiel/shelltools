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

	// 3rd Party packages
	"github.com/simeji/jid"

	// My Packages
	"github.com/rsdoiel/cli"
	"github.com/rsdoiel/shelltools"
)

var (
	usage = `USAGE: %s [OPTIONS] [EXPRESSION] [INPUT_FILENAME]`

	description = `
SYSNOPSIS

%s provides for both interactive exploration of JSON structures like jid 
and command line scripting flexibility for data extraction like jq.
`

	examples = `
EXAMPLE

If myblob.json contained

    {"name": "Doe, Jane", "email":"jane.doe@example.org"}

Getting just the name could be done with

    %s -i myblob.json -e .name
`

	// Basic Options
	showHelp    bool
	showLicense bool
	showVersion bool

	// Application Specific Options
	monochrome  bool
	expression  string
	inputFName  string
	outputFName string
)

func init() {
	// Basic Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")

	// Application Specific Options
	flag.BoolVar(&monochrome, "m", false, "display output in monochrome")
	flag.StringVar(&expression, "e", "", "apply expression to input")
	flag.StringVar(&inputFName, "i", "", "input filename")
	flag.StringVar(&inputFName, "o", "", "output filename")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	// Configuration and command line interation
	cfg := cli.New(appName, "JSONQUERY", fmt.Sprintf(shelltools.LicenseText, appName, shelltools.Version), shelltools.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionsText = "OPTIONS\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName)

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

	var (
		in  *os.File
		out *os.File
		err error
	)

	// Handle ordered args.
	if len(args) > 0 {
		expression = args[0]
	}
	if len(args) > 1 {
		inputFName = args[1]
	}
	if len(args) > 2 {
		outputFName = args[2]
	}

	// FIXME: figure out if this is an interactive session or content read in from a file.
	if len(inputFName) == 0 {
		in = os.Stdin
	} else {
		in, err = os.Open(inputFName)
		if err != nil {
			log.Fatalln(err)
		}
		defer in.Close()
	}
	if len(outputFName) == 0 {
		out = os.Stdout
	} else {
		out, err = os.Open(outputFName)
		if err != nil {
			log.Fatalln(err)
		}
		defer out.Close()
	}
	// FIXME: Configure the jid engine
	engineAttributes := &jid.EngineAttribute{
		DefaultQuery: expression,
		Monochrome:   monochrome,
	}

	// FIXME: Run the jid engine appropriately
	engine, err := jid.NewEngine(in, engineAttributes)
	if err != nil {
		log.Fatalln(err)
	}
	result := engine.Run()
	if err := result.GetError(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	fmt.Printf("%s", result.GetContent())
}
