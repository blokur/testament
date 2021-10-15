# Testament

For all your testing needs.

[![Build Status](https://travis-ci.com/blokur/testament.svg?token=TM5LRGpEAwKms8UULFDi&branch=master)](https://travis-ci.com/blokur/testament)

![Testament](https://media.giphy.com/media/RScBAu4Y11IME/giphy.gif)

1. [Introduction](#introduction)
2. [Development](#development)
   - [Prerequisite](#prerequisite)
   - [Running Tests](#running-tests)
   - [Make Examples](#make-examples)
   - [Changelog](#changelog)

## Introduction

These are common tools we usually use in our tests.

```
+----------------+------------------------------------------------------+
| AssertInError | Check if an error is found in a deeply nested errors  |
| RandomString  | Generates a randomly generated string by given length |
+----------------+------------------------------------------------------+
```

## Development

### Prerequisite

This project supports Go >= `1.17`. To run targets from the `Makefile` you need
to install GNU make.

In order to install dependencies:

```bash
make dependencies
```

This also installs [reflex][reflex] to help with development process.

### Running Tests

To watch for file changes and run unit_test:

```bash
make unit_test_watch
# or to run them with race flag:
make unit_test_race_watch
```

There is also a `integration_test` target for running integration tests.

### Make Examples

```bash
make                           # shows help on targets.
make unit_test
make unit_test run=TestMyTest   # runs a specific test with regexp.
make unit_test dir=./path/...   # runs tests in a package.
make unit_test dir=./path/... run=TestSomethingElse
```

Please see the Makefile for more targets.

### Changelog

You need to update the changelogs file before each release. In order to update
the changelogs file run the following:

```bash
make changelog
```

When you are ready to make a commitment and tag the next release, use this
target and pass in the next tag:

```bash
make changelog_release tag=v1.0.1
```

[reflex]: https://github.com/cespare/reflex
