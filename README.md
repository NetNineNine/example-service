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
