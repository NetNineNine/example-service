# 2. Use gitlint instead of commitlint

Date: 2024-10-19

## Status

Accepted

## Context

The example file that Wynand van Wyk shared used a pre-built `commitlint` in a repo we can't access. Looking into the
docs on the official commitlint website, they want you to install it using `npm`, even for your CI system. This seems
cumbersome. There are other Docker images available, which could work for us, but we'd need to look into their security.

Also, since this is an example repo, I want an example of an ADR that could be superseded by a later ADR.

## Decision

I've decided to go with `gitlint` instead of `commitlint` in the interim. This at least gives us something working. More
info can be found [here](https://github.com/jorisroovers/gitlint?tab=readme-ov-file)

## Consequences

We have Conventional Commit linting available, but nearly as configurable as `commitlint` allows.
