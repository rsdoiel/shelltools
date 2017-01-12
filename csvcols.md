
# csvcols

## USAGE

csvcols [OPTIONS] ARGS_AS_COLS

## SYNOPSIS

csvcols converts a set of command line args into columns output in CSV format.

## OPTIONS


+ -d, -delimiter	set delimiter for conversion
+ -h, -help     	display help
+ -l, -license      display license
+ -v, -version      display version

## EXAMPLES

Simple usage of building a CSV file one row at a time.

```shell
    csvcols one two three > 3col.csv
    csvcols 1 2 3 >> 3col.csv
    cat 3col.csv
```

Example parsing a pipe delimited string into a CSV line

```shell
    csvcols -d "|" "one|two|three" > 3col.csv
    csvcols -delimiter "|" "1|2|3" >> 3col.csv
    cat 3col.csv
```

