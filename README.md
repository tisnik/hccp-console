# hccp-console
PoC: HCC Proxy Console

<!-- vim-markdown-toc GFM -->

* [Description](#description)
* [Building and running](#building-and-running)
* [Make targets](#make-targets)

<!-- vim-markdown-toc -->

## Description

Controller and web UI for HCC Proxy

Please note that this is just PoC at this moment

## Building and running

1. `make build`
1. `./hccp-console`

## Make targets

```
Usage: make <OPTIONS> ... <TARGETS>

Available targets are:

clean                Run go clean
build                Build binary containing service executable
fmt                  Run go fmt -w for all sources
lint                 Run golint
vet                  Run go vet. Report likely mistakes in source code
cyclo                Run gocyclo
ineffassign          Run ineffassign checker
errcheck             Run errcheck
goconst              Run goconst checker
help                 Show this help screen
```
