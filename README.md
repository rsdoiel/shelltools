# findfile

A simple implementation of a file find in go. It demonstrates walking a file path and making simple choices based on file type.


## USAGE 

	findfile [OPTIONS] TARGET_FILENAME [DIRECTORIES_TO_SEARCH]

Version 0.0.0

## OPTIONS

 flags | description
-------|---------------------------------------
 -d    | find directories only
 -f    | find files only
 -F    | list full path for files found
-h     | display this help message
-p     | find file(s) based on basename prefix
-s     | find file(s) based on basename suffix
-v     | display version message


