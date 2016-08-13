
[![Go Report Card](http://goreportcard.com/badge/rsdoiel/fsutils)](http://goreportcard.com/report/rsdoiel/fsutils)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)


# fsutils

Various utilities for simplifying work on the command line. It includes
a simple implementation of a file finder in go. It demonstrates walking a file path, making choices based on file type, and works as an example of using Go's flag package.


## USAGE 

	findfile [OPTIONS] TARGET_FILENAME [DIRECTORIES_TO_SEARCH]
	finddir [OPTIONS] TARGET_FILENAME [DIRECTORIES_TO_SEARCH]

Use the "-help" option for a full list of options.


## Installation

_fsutils_ is go get-able.

```
    go get github.com/rsdoiel/fsutils/...
```

Or grab the pre-compiled binaries at http://github.com/rsdoiel/fsutils/releases/latest
