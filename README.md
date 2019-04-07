# go-cmnt-eol-lint

Lint your EOL of Go comments

## About

- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments#comment-sentences) said:

> Comments should begin with the name of the thing being described and end in a period:

- This command check whether your Go comments have a period or not.

## Usage

```console
$ go-cmnt-eol-lint File_1.go [File_n.go...]
```



## How to build

### Require

- Go 1.11+
    - Using `GO111MODULE=on` in Makefile


### Build

```console
$ make
```


### Test

``` console
$ make test
```


### Install

``` console
$ make install
```

- Install to `$GOPATH/bin`
