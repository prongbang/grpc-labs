# 02 Protocol Buffer Basics: Go
[https://developers.google.com/protocol-buffers/docs/proto3#default](https://developers.google.com/protocol-buffers/docs/proto3#default)

## Defining your protocol format

The `.proto` file starts with a package declaration

```go
syntax = "proto3";
package tutorial;
```

Next, you have your message definitions.
```go
message Person {
    string name = 1;
    int32 id = 2;   // Unique ID number for this person
    string email = 3;

    enum PhoneType {
        MOBILE = 0;
        HOME = 1;
        WORK = 2;
    }

    message PhoneNumber {
        string number 1;
        PhoneType type = 2;
    }

    repeated PhoneNumber phones = 4;
}

// Our address book file is just one of these.
message AddressBook {
    repeated Person people = 1;
}
```
The `" = 1"`, `" = 2"` markers on each element identify the unique "tag" that field uses in the binary encoding.

## Compiling your protocol buffers
```
$ protoc -I=02-proto --go_out=02-proto 02-proto/addressbook.proto
```
This generates `addressbook.pb.go` in your specified destination directory.

## Create an instance of Person
```go
import (
    pb "grpc-labs/tutorial"
)
p := pb.Person{
        Id:    1,
        Name:  "prongbang",
        Email: "prongbang@endtry.com",
        Phones: []*pb.Person_PhoneNumber{
                {
                    Number: "099-999-8888", Type: pb.Person_HOME
                },
        },
}
```

## Writing a Message
```go
book := &pb.AddressBook{}
// ...

// Write the new address book back to disk.
out, err := proto.Marshal(book)
if err != nil {
    log.Fatalln("Failed to encode address book:", err)
}
if err := ioutil.WriteFile(fname, out, 0644); err != nil {
    log.Fatalln("Failed to write address book:", err)
}
```

## Reading a Message
```go
// Read the existing address book.
in, err := ioutil.ReadFile(fname)
if err != nil {
    log.Fatalln("Error reading file:", err)
}
book := &pb.AddressBook{}
if err := proto.Unmarshal(in, book); err != nil {
    log.Fatalln("Failed to parse address book:", err)
}
```