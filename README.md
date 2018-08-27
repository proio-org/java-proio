# proio for Java
[![Build Status](https://travis-ci.org/proio-org/java-proio.svg?branch=master)](https://travis-ci.org/proio-org/java-proio)
[![codecov](https://codecov.io/gh/proio-org/java-proio/branch/master/graph/badge.svg)](https://codecov.io/gh/proio-org/java-proio)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9d03afa4af904c65a288774a9d8b4fcf)](https://www.codacy.com/app/decibelcooper/java-proio?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=proio-org/java-proio&amp;utm_campaign=Badge_Grade)

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
java --illegal-access=deny -cp target/proio-*-jar-with-dependencies.jar proio.Ls ../samples/muons-withmeta.proio | less
```
