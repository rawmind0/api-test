api-test
============

This program runs an api named service in configured port. Used to show and/or test microservice concept

## Build

```
make build

# Build multiplatform binaries
CROSS=1 make build
```

## Usage

```
NAME:
   api-test - api-test [OPTIONS]

USAGE:
   api-test [global options] command [command options] [arguments...]

VERSION:
   dev

AUTHOR:
   Raul Sanchez <rawmind@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d    Debug logging
   --name value   service name (default: "dev")
   --port value   service port (default: "8080")
   --help, -h     show help
   --version, -v  print the version
```
