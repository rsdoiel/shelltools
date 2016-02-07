//
// findfile.go - a simple directory tree walker that looks for files
// by name, basename or extension. Basically a unix "find" light to
// demonstrate walking the file system
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
	findBasename        bool
	findExtension       bool
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
	err := filepath.Walk(docroot, func(p string, info os.FileInfo, err error) error {
		if findDirectoriesOnly == true && info.IsDir() == false {
			return nil
		}
		if findFilesOnly == true && info.Mode().IsRegular() == false {
			return nil
		}
		if err != nil {
			return fmt.Errorf("Can't read %s, %s", p, err)
		}
		s := filepath.Base(p)
		switch {
		case findBasename == true:
			if strings.HasPrefix(s, target) == true {
				display(docroot, p)
			}
		case findExtension == true:
			if strings.HasSuffix(s, target) == true {
				display(docroot, p)
			}
		case strings.Compare(s, target) == 0:
			display(docroot, p)
		}
		return nil
	})
	return err
}

func init() {
	flag.BoolVar(&help, "h", false, "display this help message")
	flag.BoolVar(&showVersion, "v", false, "display version message")
	flag.BoolVar(&findDirectoriesOnly, "d", false, "find directories only")
	flag.BoolVar(&findFilesOnly, "f", false, "find files only")
	flag.BoolVar(&findBasename, "p", false, "find file(s) based on basename prefix")
	flag.BoolVar(&findExtension, "s", false, "find file(s) based on basename suffix")
	flag.BoolVar(&outputFullPath, "full", false, "list full path for files found")
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
