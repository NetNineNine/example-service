# 3. Semantic Releases using go-semantic-release

Date: 2024-10-19

## Status

Accepted

## Context

We want to use [Semantic Release](https://github.com/semantic-release/semantic-release?tab=readme-ov-file) to auto-enforce [Semantic Versioning](http://semver.org/) as a best practice.

The `semantic-release` project was originally designed for Node.js packages. While it's marketed as being language-agnostic, it has a strong bias towards the Node.js ecosystem, so when we first tried it, it failed because of a lack of `packages.json`, which a Go project obviously doesn't have.

## Decision

For the time being, we are opting to use [go-semantic-release/semantic-release](https://github.com/go-semantic-release/semantic-release), which is a Go implementation, where the project uses itself for versioning, providing good examples to base our CI off of.

## Consequences

In theory it should be the same, except with better examples since we can reference their own Go project. Lacking a `package.json` could be an issue, and not having an `.releaserc.yaml` file means we need different options for enforcing Conventional Commits.