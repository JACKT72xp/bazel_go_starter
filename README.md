# Bazel Go Gin Starter

This repository provides a template for a Go web application using the Bazel build system, the Gin web framework, and integrates live reloading with Reflex.

## Description

The project is structured to use Bazel, a powerful build tool that allows for efficient and reproducible builds. The Go web application is built using the Gin framework, known for its performance and efficiency. To enhance the development experience, this setup includes Reflex, which is used for live reloading of the application during development, allowing changes to be tested quickly without manual recompilation.

## Prerequisites

Before you start, ensure you have the following installed:

- Go programming language
- Bazel build tool
- Reflex for live reloading

You can install Reflex using the following command:

```sh
go get github.com/cespare/reflex
```

### Building the Project
To build the project using Bazel, navigate to the project's root directory and run:

```sh
bazel build //:main
```

This command compiles the application and generates an executable in the bazel-bin/ directory.

### Running the Application
To run the compiled application, use the following Bazel command:

```sh
bazel run //:main
```
This will start the Go web server on port 8080.

## Using Reflex for Live Reloading
To use Reflex for live reloading during development, you can run the following command in the root of the repository:
```sh
reflex -g 'main.go' -s -- sh -c 'bazel run //:main'
```

This command monitors changes to main.go and automatically restarts the server using Bazel whenever changes are detected, enabling a faster development cycle.

- -g 'main.go': Watches for changes in main.go.
- -s: Stops the current process before starting a new one.
- -- sh -c 'bazel run //:main': Specifies the command to run on file change.
