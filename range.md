
[![Go Report Card](http://goreportcard.com/badge/rsdoiel/range)](http://goreportcard.com/report/rsdoiel/range)
[![License](https://img.shields.io/badge/License-BSD%202--Clause-blue.svg)](https://opensource.org/licenses/BSD-2-Clause)


# range

Range is a simple utility for shell scripts that emits a list of integers starting with the first command line argument and ending with the last integer command line argument.

If the first argument is greater than the last then it counts down otherwise it counts up.


## USAGE

```
    range [OPTIONS] STARTING_INTEGER ENDING_INTEGER [INCREMENT_INTEGER]
```

## EXAMPLES

+ Count from one through five: range 1 5
+ Count from negative two to six: range -- -2 6
+ Count even number from two to ten: range --increment=2 2 10
+ Count down from ten to one: range 10 1
+ Pick a value from a range: range -r 0 10


## OPTIONS

+ *-e*, *--end* The ending integer (e.g. 10)
+ *-h*, *--help* Display this help document.
+ *-i*, *--increment* The non-zero integer value to increment by (e.g. 1 or -1)
+ *-s*, *--start* The starting integer (e.g. 1)
+ *-r*, *--random* Pick a random element from a range 

## Installation

_range_ can be installed with the *go get* command.

```
    go get github.com/rsdoiel/range/...
```



## LICENSE

copyright (c) 2014 All rights reserved.
Released under the [Simplified BSD License](http://opensource.org/licenses/bsd-license.php)
