# Testament

For all your testing needs.

[![Build Status](https://travis-ci.com/blokur/testament.svg?token=TM5LRGpEAwKms8UULFDi&branch=master)](https://travis-ci.com/blokur/testament)

![Testament](https://media.giphy.com/media/RScBAu4Y11IME/giphy.gif)

1. [Introduction](#introduction)
2. [Development](#development)
   - [Prerequisite](#prerequisite)
   - [Running Tests](#running-tests)
   - [Make Examples](#make-examples)

## Introduction

These are common tools we usually use in our tests.

| Function           | Description                                                            |
| :----------------- | :--------------------------------------------------------------------- |
| AssertInError      | Check if an error is found in a deeply nested errors                   |
| RandomString       | Generates a randomly generated string by given length                  |
| GetFreeOpenPort    | Returns a port that is already claimed.                                |
| GetFreePort        | Returns a random open port.                                            |
| RandomString       | Returns a randomly generates string with the length of count.          |
| RandomLowerString  | Returns a randomly generates lower-cased string with a maximum length. |
| StringSlice        | Return a string slice with the provided length.                        |
| RandomStringSlice  | Return a random string slice with maximum length.                      |
| IntSlice           | Returns a slice of int elements with the provided length.              |
| RandIntSlice       | Returns a slice of int elements with random length of less than n.     |
| Int32Slice         |                                                                        |
| RandInt32Slice     |                                                                        |
| Int64Slice         |                                                                        |
| RandInt64Slice     |                                                                        |
| IntSliceComparer   | Is a go-cmp comparer that doesn't care if the slices are not sorted.   |
| Int32SliceComparer |                                                                        |
| Int64SliceComparer |                                                                        |
| RandomEmailAddress | Returns a valid random email address.                                  |

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

[reflex]: https://github.com/cespare/reflex
