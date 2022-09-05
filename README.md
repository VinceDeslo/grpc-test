# Messing around with gRPC

This repo is a playground to mess around with gRPC in Go.

### Objective
Build out a gRPC service to retrieve users from an in memory database seeded by JSON.

### Requirements
- Go
- gRPC

### Top level Make commands:
- `make gen` : Generates the protobuf definitions for the service and messages.
- `make run` : Starts the gRPC server on `localhost:9090`.
- `make curl-users` : Gets all the data from the in memory database with `gpcurl`.

### Structure
```
grpc-test
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── proto
│   └── ** Protobuf definitions **
├── scripts
│   └── ** cURL scripts **
└── server
    ├── db
    │   └── ** JSON data **
    ├── main.go
    ├── models
    │   └── ** Object DTOs **
    ├── services
    │   └── ** Dependency injection sources for services **
    ├── user
    │   ├── ** Protobuf generated files and in memory database **
    └── utils
        └── ** JSON and I/O helpers **
```
