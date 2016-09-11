
# mergepath

A small utility to merge PATH lists across different platforms (e.g. Linux, OS X).

## USAGE 

```shell
    mergepath [OPTIONS] PATH_TO_ADD PATH_TO_MODIFY
```

## EXAMPLES

append work directory to existing path: 

```shell
    mergepath . $PATH
```

prepend working directory to existing path: 

```shell
    mergepath -P . $PATH
```

## OPTIONS

```
    -a, -append	Append the directory to the path removing any duplication
    -c, -clip	Remove a directory from the path
    -d, -directory	The directory you want to add to the path.
    -e, -envpath	The path you want to merge with.
    -h, -help	This help document.
    -p, -prepend	Prepend the directory to the path removing any duplication
```
