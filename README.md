# Proctor

<p align="center"><img src="doc/proctor-logo.png" width="360"></p>
<p align="center">
  <a href="https://travis-ci.org/gojektech/proctor"><img src="https://travis-ci.org/gojektech/proctor.svg?branch=master" alt="Build Status"></img></a>
  <a href="https://goreportcard.com/report/github.com/gojektech/proctor"><img src="https://goreportcard.com/badge/github.com/gojektech/proctor"></img></a>
  <a href="https://golangci.com"><img src="https://golangci.com/badges/github.com/gojektech/proctor.svg"></img></a>
</p>

## Description

Proctor is a developer friendly automation orchestrator. It helps everyone use automation and contribute to it

### proctor CLI

Proctor CLI is a command line tool to interact with [proctord](https://github.com/gojektech/proctor/blob/master/proctord).
Users can use it to run procs.

### Dev environment setup

* Install and setup golang
* Install glide
* Clone the repository: `go get github.com/gojektech/proctor`
* Install dependencies using glide: `glide install`
* [Configure proctor CLI](#proctor-cli-configuration)
* Running `go install github.com/gojektech/proctor` will place the CLI binary in your `$GOPATH/bin` directory
* Run `proctor version` to check installation

### Running tests

* [Setup dev environment](#dev-environment-setup)
* `cd proctord`. Refer README to setup test environment of proctord
* After setting up test env for proctord, `cd ..`
* Configure environment variables `source .env.test`
* Run tests: `go test -race -cover $(glide novendor)`

#### Proctor CLI configuration

* Make a directory `.proctor` inside your home directory
* Create a file proctor.yaml inside above directory
* Put the following content in the above file

``` sh
PROCTOR_HOST: [hostname where proctord is running]
```
