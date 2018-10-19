# Golang Database Service (GDS)

A simple golang service to relay access to a database.

## Overview

GDS is a golang webserver designed to act as a thin layer between other services and a database. The service implements two network interfaces;

- a RESTful HTTP1 JSON interface, and
- a gRPC HTTP2 interface

## Connecting

### SSL/TLS

> gRPC has SSL/TLS integration and promotes the use of SSL/TLS to authenticate the server, and to encrypt all the data exchanged between the client and the server.

GDS uses SSL/TLS for all network requests. This may seem unnecessary for an internally facing service, however I chose to do this for a few main reasons:

- The purpose of GDS as a boilerplate is to implement best practices and allow the developer to make choices suited to their particular situation, and
- gRPC and HTTP on the same socket using gRPC's built in `ServeHTTP` method requires HTTP2 which requires the connection to arrive over SSL/TLS

Generate self signed certificates:

    openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt

## Tools

### [Reflex](https://github.com/cespare/reflex)

GDS uses Reflex to auto-reload the webserver after changes to source are detected. 

## Managing Dependencies

### Application Dependencies

Application dependencies are managed using [Dep](https://github.com/golang/dep).

> `dep` was the "official experiment." The Go toolchain, as of 1.11, has (experimentally) adopted an approach that sharply diverges from dep. As a result, we are continuing development of dep, but gearing work primarily towards the development of an alternative prototype for versioning behavior in the toolchain.

I am investigating (and fully intend to migrate to) go 1.11 modules, however, at this moment, it is a low priority.

#### Usage

```
Dep is a tool for managing dependencies for Go projects

Usage: "dep [command]"

Commands:

  init     Set up a new Go project, or migrate an existing one
  status   Report the status of the project's dependencies
  ensure   Ensure a dependency is safely vendored in the project
  prune    Pruning is now performed automatically by dep ensure.
  version  Show the dep version information

Examples:
  dep init                               set up a new project
  dep ensure                             install the project's dependencies
  dep ensure -update                     update the locked versions of all dependencies
  dep ensure -add github.com/pkg/errors  add a dependency to the project

Use "dep help [command]" for more information about a command.
```

Add dependencies:

    $ dep ensure -add github.com/pkg/errors github.com/foo/bar

Dep will also search the imports in your project to ensure your `Gopkg.lock` is in sync with your imports is in sync with `/vendor`

### Tool Dependencies

Tool dependencies are managed using [Retool](https://github.com/twitchtv/retool).

> `retool` helps manage the versions of tools that you use with your repository. These are executables that are a crucial part of your development environment, but aren't imported by any of your code, so they don't get scooped up by glide or godep (or any other vendoring tool).

#### Usage

```
usage: retool (add | remove | upgrade | sync | do | build | help)

use retool with a subcommand:

add will add a tool
remove will remove a tool
upgrade will upgrade a tool
sync will synchronize your _tools with tools.json, downloading if necessary
build will compile all the tools in _tools
do will run stuff using your installed tools

help [command] will describe a command in more detail
version will print the installed version of retool
```

Add a tool:

    retool add github.com/jteeuwen/go-bindata/go-bindata origin/master
               ^ tool                                    ^ commit

Prepend `retool do` to use tools installed by retool:

    retool do go-bindata -pkg testdata -o ./testdata/testdata.go ./testdata/data.json

## Testing

    $ docker-compose run grpc_cli ls app:8080 -l --channel_creds_type=ssl --ssl_target=localhost

    