
# urlparse

# Overview

A commandline utility that returns a delimited list of URL parts suitable
for use with other Unix utilities like _cut_.

## USAGE 

    urlparse [OPTIONS] URL_TO_PARSE

Display the parsed URL as delimited fields on one line.

## EXAMPLES


Get protocol. Returns "http".
 
     urlparse --protocol http://example.com/my/page.html


Get host or domain name.  Returns "example.com".
 
     urlparse --host http://example.com/my/page.html


Get path. Returns "/my/page.html".
 
     urlparse --path http://example.com/my/page.html


Get basename. Returns "page.html".
 
     urlparse --basename http://example.com/my/page.html


Get extension. Returns ".html".
 
     urlparse --extension http://example.com/my/page.html


## OPTIONS

Without options urlparse returns protocol, host and path fields 
separated by a tab.

+ -b, -basename	Display the base filename at the end of the path.
+ -D, -delimiter Set the output delimited for parsed display. (defaults to tab)
+ -d, --directory Display all but the last element of the path
+ -e, -extension Display the filename extension (e.g. .html).
+ -H, -host	Display the host (domain name) in URL.
+ -p, -path	Display the path after the hostname.
+ -P, -protocol	Display the protocol of URL (defaults to http)

+ -h, -help	Display this help document.
+ -l, -license Display license information.
+ -v, -version Display version information.

