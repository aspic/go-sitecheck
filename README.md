# Simple site link checking tool

This tool aims to check referred links for a page (and potential slow
loading sub pages).

## Build

Ensure that go is installed, and that the checked out project is present
in your GOPATH.

    // build binary
    $ go get && go build

## Usage

    $ ./go-sitecheck -url=<some url> -threshold=<threshold in ms> -depth=<some-depth>
