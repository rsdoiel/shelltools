
# jsonquery

_jsonquery_ was inspired by Stephen Dolan's [jq](https://github.com/stedolan/jq) and 
Simeji's [jid](https://github.com/simeji/jid). It leverages the latter's package for
proccessing the JSON query path organized as a tool similar to _jq_.

## USAGE

```
    jsonquery [OPTIONS] [EXPRESSION] [INPUT_FILENAME] [OUTPUT_FILENAME]
```

## SYSNOPSIS

jsonquery provides for both interactive exploration of JSON structures like jid 
and command line scripting flexibility for data extraction like jq.

+ EXPRESSION can be an empty stirng or dot notation for an object's path
+ INPUT_FILENAME is the filename to read or a dash "-" if you want to explicity read from stdin
	+ if not provided then jsonquery reads from stdin
+ OUTPUT_FILENAME is the filename to write or a dash "-" if you want to explicity write to stdout
	+ if not provided then jsonquery write to stdout

## OPTIONS

```
	-h	display help
	-i	run interactively
	-l	display license
	-m	display output in monochrome
	-v	display version
```

## EXAMPLE

If myblob.json contained

```json
    {"name": "Doe, Jane", "email":"jane.doe@example.org"}
```

Getting just the name could be done with

```shell
    jsonquery .name myblob.json
```

This would yeild

```
    "Doe, Jane"
```

jsonquery v0.0.14-alpha-1
