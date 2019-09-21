# go-keychain

[![Actions Status](https://github.com/jmuldoon/go-keychain/workflows/test/badge.svg)](https://github.com/jmuldoon/go-keychain/actions)

Personal use, password export for applications. Quick and dirty implemenation for password setting and retrieval. Plus quick use case.

<!-- TOC -->

- [go-keychain](#go-keychain)
  - [Example cli arguments](#example-cli-arguments)
    - [Example useage](#example-useage)
  - [.bashrc / .bash_profile](#bashrc--bash_profile)
  - [Makefile use](#makefile-use)
  - [Docker-Compose](#docker-compose)

<!-- /TOC -->

## Example cli arguments

Arguments   | Values        | Description
---         | ---           | ---
account     | string        | user account (required)
group       | string        | user access group
data        | string        | password/encrypted part (required if writing)
generate    | integer       | integer length of the password string to be generated
label       | string        | name (required)
service     | string        | where (required)
read        | bool          | defaults false. If you want to read must be specified

### Example useage

``` bash
go-keychain -account username -service testservice -label testlabel -data tests123
go-keychain -account username -service testservicepath2 -label testlabel -generate 32
```

## .bashrc / .bash_profile

Please add the following line to your bash profile so that when loaded you can automatically export
the default base password you'd like.

``` bash
export DEFAULT_KEY_NAME=$(go-keychain -account username -service testservice -label testlabel -read)
```

## Makefile use

The following `make` commands executed from the top level directory will enable the building of the binary. Pre-requisite is to have go installed. This was built using `go version go1.9 darwin/amd64`, but should be cross compilable with `gox` if needed.

``` bash
make all
```

Please see for a complete listing the below command.

``` bash
make help
```

## Docker-Compose

The following command will allow one to pull and bring up a container that will allow them to work directly with the binary

``` bash
docker-compose run go /bin/sh
```
