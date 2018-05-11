# 01 Setup

## Install Golang 
[https://golang.org/dl/](https://golang.org/dl/)

## SettingGOPATH: 
[https://github.com/golang/go/wiki/SettingGOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

## Check 
```
$ go version
```

## Install gRPC
```
$ go get -u google.golang.org/grpc
```

## Install Protocol Buffers v3
```
$ go get -u github.com/golang/protobuf/protoc-gen-go
```
It must be in your $PATH for the protocol compiler, protoc, to find it.
```
$ export PATH=$PATH:$GOPATH/bin
```

## Download `protoc`
[https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases)

- protoc-3.5.1-osx-x86_64.zip
- unzip
- set path