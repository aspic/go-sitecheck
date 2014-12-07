# Simple site link checking tool

This tool aims to check referred links for a page.

## To build

Ensure that the project is checked out and available in your GOPATH.

    // build binary
    go get && go build

    // run
    ./go-sitecheck -url <some url> -threshold <threshold in ms>
