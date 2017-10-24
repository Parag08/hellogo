# notes

## Add license

## Add test_files

## travis CI (or build passing flag)

add .travis.yml

For go add following content to it


```
language: go
script: go test
notifications:
  email: false
```

go to travis CI [Travis CI](https://travis-ci.org/)

add project

try building project if it builds hurray else go back and debug.

add following line to README.md

```
[![Build Status](https://secure.travis-ci.org/Parag08/hellogo.png?branch=master)](http://travis-ci.org/Parag08/hellogo)
```

## Coverage (coveralls)

go to [Coveralls](https://coveralls.io/)

add project

get token

add .coveralls.yml

```
repo_token: oN4oo3LVlbiOTX8Gltl5cZqVmklqXHwxv
```

run following commands

```
go get golang.org/x/tools/cmd/cover
go get github.com/mattn/goveralls
go test -v -covermode=count -coverprofile=coverage.out
$GOPATH/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken oN4oo3LVlbiOTX8Gltl5cZqVmklqXHwxv
```


add following line to README.md

```
[![Coverage Status](https://coveralls.io/repos/github/Parag08/hellogo/badge.svg?branch=master)](https://coveralls.io/github/Parag08/hellogo?branch=master) 
```

## CII Best Practices

go to [CII best practices](https://bestpractices.coreinfrastructure.org/)

add project

add following line to README.md


```
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/1343/badge)](https://bestpractices.coreinfrastructure.org/projects/1343)
```

add project to
