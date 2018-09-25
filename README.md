# proio for Java
[![Build Status](https://travis-ci.org/proio-org/java-proio.svg?branch=master)](https://travis-ci.org/proio-org/java-proio)
[![codecov](https://codecov.io/gh/proio-org/java-proio/branch/master/graph/badge.svg)](https://codecov.io/gh/proio-org/java-proio)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/3540d7b51d034acc8bd47ffac45d32fd)](https://www.codacy.com/app/proio-org/java-proio?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=proio-org/java-proio&amp;utm_campaign=Badge_Grade)

Please see the [main proio repository](https://github.com/proio-org/proio) for general information on proio.

## Installation
At this time, proio is not yet available in the maven central repository.
However, maven can be used to easily build a jar file.

### Requirements
* Maven
* Protobuf compiler (`protoc`)

### Building the code
```shell
git submodule init
git submodule update
mvn install
```

## Running the "Ls" tool
This is a tool that serves as an example for a browser tool.  This one is
simple and only dumps text to the terminal.
```shell
java --illegal-access=deny -cp target/proio-*-jar-with-dependencies.jar proio.Ls sample.proio | less
```
