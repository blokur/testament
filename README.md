# Testament

For all your testing needs.

1. [Introduction](#introduction)
2. [Development](#development)
   - [Prerequisite](#prerequisite)
   - [Running Tests](#running-tests)
   - [Make Examples](#make-examples)
   - [Changelog](#changelog)
   - [Mocks](#mocks)
3. [References](#references)


## Introduction

These are common tools we usually use in our tests.

+---------------+------------------------------------------------------+
| AssertInError | Check if an error is found in a deeply nested errors |
+---------------+------------------------------------------------------+

## Development

### Prerequisite

This project supports Go > `1.14`. To run targets from the `Makefile` you need
to install GNU make.

In order to install dependencies:

```bash
make dependencies
```

This also installs [reflex][reflex] to help with development process.

### Running Tests

To watch for file changes and run unittest:

```bash
make unittest_watch
# or to run them with race flag:
make unittest_race_watch
```

There is also a `integration_test` target for running integration tests.

### Make Examples

```bash
make                           # shows help on targets.
make unittest
make unittest run=TestMyTest   # runs a specific test with regexp.
make unittest dir=./path/...   # runs tests in a package.
make unittest dir=./path/... run=TestSomethingElse
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

### Mocks

To generate mocks run:

```bash
make mocks
```

## References

[reflex]: https://github.com/cespare/reflex
