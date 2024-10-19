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

They are most easily managed using [adr-tools](https://github.com/npryce/adr-tools/blob/master/INSTALL.md). On Linux and MacOS 
the simplest is to use Homebrew

```shell
brew install adr-tools
```

## ADRs

- [Record architecture decisions](doc/adr/0001-record-architecture-decisions.md)
- [dsaa](doc/adr/0002-use-gitlint-instead-of-commitlint.md)