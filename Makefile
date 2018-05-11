build:
    protoc -I. proto/foo/foo.proto --go_out=plugins=grpc:.
