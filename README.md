# MBenz Hackathon

## Introduction

This hackathon intends to be a hiring event. We will give a problem statement based on real world situations to the aspiring candidates. They will have to develop a well-tested and scalable solution by following the best development and coding practices.

## Hackathon Link
https://www.techgig.com/hackathon/mercedes-benz-hiring-challenge

## How to run
Use the `Makefile` to run it
```sh
$ make all
```

# Possible Errors
We might see some possbile errors in terms of cross container netwoking. In such cases, we need to run the `routing` micro-service locally. To do that, just build and run the `poc` microservice using docker by
```sh
$ docker build -t mbenz-poc:latest -f Dockerfile .
$ docker run --rm --name mbenzPoc -p 9000:9000 mbenz-poc
```

and then just move to the local dir of `mbenz_planning` and do

```
$ go mod download
$ go build -ldflags="-s -w" -o apiserver .
$ ./apiserver
```
This would run the project locally and we can run the apis