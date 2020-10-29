# Server

## API

The API is specified with [OpenAPI 2.0](https://swagger.io/specification/v2/)
in swagger.yml. JetBrains IDEs have excellent support for [OpenAPI
specifications](https://plugins.jetbrains.com/plugin/14394-openapi-specifications).
For VS Code users, you can consider the [Swagger
Viewer](https://marketplace.visualstudio.com/items?itemName=Arjun.swagger-viewer)
and [OpenAPI (Swagger)
Editor](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi)
extensions.

On the server side, we use
[go-swagger](https://github.com/go-swagger/go-swagger) to automatically
generate some boilerplate for the server. **Every time swagger.yml changes, one
should run `make generate` to ensure the boilerplate is up to date.**

You can view a generated API documentation by running `make start` and visiting
http://localhost:9000/docs.

## Installing Go

Go ≥1.15 is required.

**macOS:** (needs [Homebrew](https://brew.sh/))
```sh
$ brew install go
$ go version
go version go1.15.3 darwin/amd64
```

**Arch Linux and derivatives:**
```sh
$ sudo pacman -S go
$ go version
go version go1.15.3 linux/amd64
```

**Ubuntu**: ([documentation](https://github.com/golang/go/wiki/Ubuntu))
```sh
$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt update
$ sudo apt install golang-go
$ go version
go version go1.15.3 linux/amd64
```

**Others** (including **Windows**): follow official [guide](https://golang.org/doc/install)
```sh
$ go version
go version go1.15.3 windows/amd64
```

## Common operations

**To start the server:**
```sh
make start
```

**To reformat code:**
```sh
make format
```

**To regenerate boilerplate code:**
```sh
make generate
```

## Code organization

Things we can edit:

- `impl` contains endpoint implementations.
- `restapi/configure_down_to_meet.go` contains server configurations. Every
  time we add a new endpoint, we need to add a line to `configureAPI` that sets
  the handler appropriately.

Things that are automatically generated, and should not be edited:

- `cmd/down-to-meet-server` is the generated source code for the server binary.
- `models` contains the generated data structures.
- `restapi` (except `configure_down_to_meet.go`) contains generated server
  code.
