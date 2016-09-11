
# finddir

Lists directories recursively. Can be constrained by prefix, suffix or
basename contents (e.g. contains).

## USAGE

```
    finddir [OPTIONS] [TARGET_FILENAME] [DIRECTORIES_TO_SEARCH]
```

Finds directories based on matching prefix, suffix or contained text in base filename.

```
    -c, -contains	find file(s) based on basename containing text
    -d, -depth	    limit depth of directories walked
    -e, -error-stop	Stop walk on file system errors (e.g. permissions)
    -f, -full-path	list full path for files found
    -h, -help	    display this help message
    -l, -license	display license information
    -m, -mod-time	display file modification time before the path
    -p, -prefix	    find file(s) based on basename prefix
    -s, -suffix	    find file(s) based on basename suffix
    -v, -version	display version message
```

