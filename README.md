# Scope


## Background
The purpose of this tool is to check if and IP address is in a CIDR block. Useful for red teams to check if the ip they are looking at is in scope.  


## Usage
Releases includes a target for linux. Windows and Mac users will have to build from source.

```
Usage:
        scope [ip] [CIDR]
        scope [ip] -f [CIDR File]
        scope --ip [ip] -f [CIDR File]
Options:
  -f, --file string   CIDR File Path
      --ip string     IP file path
```

## Build
Install the latest version of Golang.

run 

```
go build
```

to build the binary

or 

```
go run .
```

to run without building



