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
We might see some possbile errors in terms of cross container netwoking. In such cases, we need to indivisually build the services. So execute the following commands
```sh
$ docker network create mbenz_default 
$ cd mbenz_poc
$ chmod +x build.sh && ./build.sh
$ cd ../mbenz_planning
$ chmod +x build.sh && ./build.sh
```
