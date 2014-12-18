# Simple (and rude) link checking tool

This tool aims to check referred links for a page (and potential slow
loading sub pages).

## Build

Ensure that you have go installed, the project is checked out and available in your GOPATH.

    // build binary
    $ go get && go build

## Usage

    $ ./go-sitecheck -url=<some url> -threshold=<threshold in ms> -depth=<some depth>

### URL mapping

By providing a map parameter this utility can translate between domains
and/or paths. This can be useful if the tool is used for scraping
external resources but testing internal resources.

Example usage:
    
    $ ./go-sitecheck -map=nrk.no:vg.no

This will replace the text fragment in all discovered links that
contains "nrk.no" with "vg.no". 

### Ignore domains

A comma separated list of domains/sub-urls to ignore can be passed to
the tool by specifying the ignore parameter.

Example usage:

    $ ./go-sitecheck -ignore=tv.nrk.no,radio.nrk.no
