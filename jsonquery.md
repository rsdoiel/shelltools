
# jsonquery

_jsonquery_ was inspired by Stephen Dolan's [jq](https://github.com/stedolan/jq) and 
Simeji's [jid](https://github.com/simeji/jid). It leverages the latter's package for
proccessing the JSON query path organized as a tool similar to _jq_.

## USAGE: 

`jsonquery [OPTIONS]`

## SYSNOPSIS

jsonquery provides for both interactive exloration of JSON structures like jid 
and command line scripting of data extraction like jq.

## OPTIONS

```
	-e	apply expression to input
	-h	display help
	-i	input filename
	-l	display license
	-m	display output in monochrome
	-o	output filename
	-v	display version
```

## EXAMPLE

If myblob.json contained

```json
    {"name": "Doe, Jane", "email":"jane.doe@example.org"}
```

Getting just the name could be done with

```shell
    jsonquery -i myblob.json -e .name
```

