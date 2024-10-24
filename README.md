# example-service
Canonical example golang-service for Net Nine Nine

## Setup

We use Pre-Commit hooks to ensure things are done correctly. You will need to set this up locally, otherwise the CI
system will bring down it's wrath upon you.

You'll need `pip`, which should come standard on most *nix based systems (linux, MacOS) since `python` is there by
default. Otherwise, find installation instructions [here](https://pip.pypa.io/en/stable/installation/).

```shell
pip install pre-commit
pre-commit install --hook-type commit-msg
```
We make use of Architecture Decision Records (as [described by Michael Nygard](http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions))
to track any significant decisions made during the development of this service.

They are most easily managed using [adr-tools](https://github.com/npryce/adr-tools/blob/master/INSTALL.md). On Linux and MacOS the simplest is to use Homebrew.

```shell
brew install adr-tools
```

While we are busy, it would be useful to install `golangci-lint` locally as well, since pre-commit will use this.

```shell
brew install golangci-lint
```

## Build

While we don't have any scripts or Dockerfiles, your simplest option is

```shell
go mod tidy
go test ./...
go build -v ./...
```

If you want an executable, you can do

```shell
go build -o dist/example-service main.go
```

## Run

To run the executable from the last step

```shell
chmod +x dist/example-service
./dist/example-service
```

If you want a JIT compiled version

```shell
go run main.go
```

And you test that it works

```shell
curl http://localhost:8080/
```
which should respond with `Hello, World!`

## ADRs

- [Record architecture decisions](doc/adr/0001-record-architecture-decisions.md)
- [Semantic Releases using go-semantic-release](doc/adr/0002-semantic-releases-using-go-semantic-release.md)
- [Use gitlint instead of commitlint](doc/adr/0003-use-gitlint-instead-of-commitlint)
