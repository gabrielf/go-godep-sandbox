go-godep-sandbox
===============

Minimal project using godep vendoring (including dependencies in your project and rewrite imports to these dependencies).

In this case the only dependency is `martini` which in turn depends on `inject`.

Install this project locally by running `go get github.com/gabrielf/go-godep-sandbox` then navigate to the source `cd $GOPATH/src/github.com/gabrielf/go-godep-sandbox` and run with `go run main.go`.
