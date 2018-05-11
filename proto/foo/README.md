# Generate
```
$ protoc -I. proto/foo/foo.proto --go_out=plugins=grpc:.
```
Or
```
$ protoc -I=proto/foo --go_out=proto/foo/foo.proto
```

Other
```
$ protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```