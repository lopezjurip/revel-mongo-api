# Mounting Go Revel API + MongoDB on Docker

[![Build Status][ci-image]][ci-url]

## See step by step

[https://blog.lopezjuri.com/2015/12/29/mounting-go-revel-api--mongodb-on-docker](https://blog.lopezjuri.com/2015/12/29/mounting-go-revel-api--mongodb-on-docker/).

## Getting Started

### Install

```sh
# Dependencies
go get github.com/revel/cmd/revel
go get github.com/tools/godep

# Project
go get github.com/mrpatiwi/revel-mongo-api
cd $GOPATH/src/github.com/mrpatiwi/revel-mongo-api

# Install dependencies with Godep
godep go install ./app
```

### Start the web server

```sh
revel run github.com/mrpatiwi/revel-mongo-api
```

### Go to http://localhost:9000/ and you'll see

"It works"

[ci-image]: https://travis-ci.org/mrpatiwi/revel-mongo-api.svg
[ci-url]: https://travis-ci.org/mrpatiwi/revel-mongo-api
