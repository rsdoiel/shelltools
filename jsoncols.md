
# jsoncols

_jsoncols_ was inspired by Stephen Dolan's [jq](https://github.com/stedolan/jq) and 
Simeji's [jid](https://github.com/simeji/jid). 

## USAGE

```
    jsoncols [OPTIONS] [EXPRESSION] [INPUT_FILENAME] [OUTPUT_FILENAME]
```

## SYSNOPSIS

_jsoncols_ provides for both interactive exploration of JSON structures like jid 
and command line scripting flexibility for data extraction into delimited
columns. This is helpful in flattening content extracted from JSON blobs.
The default delimiter for each value extracted is a comma. This can be
overridden with an option.

+ EXPRESSION can be an empty stirng or dot notation for an object's path
+ INPUT_FILENAME is the filename to read or a dash "-" if you want to 
  explicity read from stdin
	+ if not provided then _jsoncols_ reads from stdin
+ OUTPUT_FILENAME is the filename to write or a dash "-" if you want to 
  explicity write to stdout
	+ if not provided then _jsoncols_ write to stdout

## OPTIONS

```
	-d	set the delimiter for multi-field output
	-h	display help
	-i	run interactively
	-l	display license
	-m	display output in monochrome
	-v	display version
```

## EXAMPLES

If myblob.json contained

```json
    {"name": "Doe, Jane", "email":"jane.doe@example.org", "age": 42}
```

Getting just the name could be done with

```shell
    jsoncols .name myblob.json
```

This would yeild

```text
    "Doe, Jane"
```

Flipping .name and .age into pipe delimited columns is as 
easy as listing each field in the expression inside a 
space delimited string.

```shell
    jsoncols -d\|  ".name .age" myblob.json
```

This would yeild

```text
    "Doe, Jane"|42
```

