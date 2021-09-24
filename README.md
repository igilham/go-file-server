# Go File Server

This is a simple example of a basic HTTP file server in Go.

## Build

```shell
go build fileserve.go
```

## Usage

By default, the application will serve the current working directory on port 8080. The directory and port can be overridden with arguments on the command line.

```shell
./fileserve [-d path/to/directory] [-p port]
```
