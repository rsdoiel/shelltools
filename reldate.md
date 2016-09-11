
[![Go Report Card](http://goreportcard.com/badge/rsdoiel/reldate)](http://goreportcard.com/report/rsdoiel/reldate)
[![License](https://img.shields.io/badge/License-BSD%202--Clause-blue.svg)](https://opensource.org/licenses/BSD-2-Clause)

# reldate

A small command line utility which returns the relative date in YYYY-MM-DD format. This is helpful
when scripting various time relationships.

## Example

If today was 2014-08-03 and you wanted the date three days in the past try--

```
    reldate 3 days
```

The output would be 

```
    2014-08-06
```

## Time units

Supported time units are

+ day(s)
+ week(s)
+ year(s)

## Specifying a date to calucate from

_reldate_ handles dates in the YYYY-MM-DD format (e.g. March 1, 2014 would be 2014-03-01).  By default _reldate_ uses today as
the date to calculate relative time from.  If you use the *--from* option you can it will calculate the relative date from that 
specific date. 

```
    reldate --from=2014-08-03 3 days
```

Will yield

```
    2014-08-06
```

## Negative increments

Command line arguments traditionally start with a dash which we also use to denote a nagative number. To tell the command line
process that to not treat negative numbers as an "option" preceed your time increment and time unit with a double dash.

```
   reldate --from=2014-08-03 -- -3 days 
```

Will yield

```
    2014-07-31
```

## Relative week days

You can calculate a date from a weekday name (e.g. Saturday, Monday, Tuesday) knowning a day (e.g. 2015-02-10 or the current date of the
week) occuring in a week.  A common case would be wanting to figure out the Monday date of a week containing 2015-02-10. The week is
presumed to start on Sunday (i.e. 0) and finish with Saturday (e.g. 6).

```
    reldate --from=2015-02-10 Monday
```

will yeild

```
    2015-02-09
```

As that is the Monday of the week containing 2015-02-10. Weekday names case insensitive and can be the first three letters
of the English names or full English names (e.g. Monday, monday, Mon, mon).

## Installation

_reldate_ can be installed with the *go get* command.

```
    go get github.com/rsdoiel/reldate/...
```


## License

copyright (c) 2014 All rights reserved.
Released under the [Simplified BSD License](http://opensource.org/licenses/bsd-license.php)
