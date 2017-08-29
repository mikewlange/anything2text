# anything2text

A Go library about to be Micro-service to convert PDF, DOC, DOCX, XML, HTML, RTF, ODT, Pages documents and images (Tesseract) to plain text.

## Installation

To fetch and build the code:

    $ go get github.com/anything2text/anything2text/...

This will also build the command line tool `docd` into `$GOPATH/bin` (assumed to be in your `PATH` already).

## Dependencies
tidy, wv, popplerutils, unrtf, https://github.com/JalfResi/justext

Example install of dependencies (not all systems):

    $ sudo apt-get install poppler-utils wv unrtf tidy
    $ go get github.com/JalfResi/justext

Add image support to the `docconv` library you first need to install and build https://github.com/otiai10/gosseract.  Now you can add `-tags ocr` to any `go` command when building/fetching `docconv` to include support for processing images:

    $ go get -tags ocr github.com/mikewlange/anything2text/...
