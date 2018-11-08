# PlayNet Libs

[![Go Report Card](https://goreportcard.com/badge/github.com/playnet-public/libs)](https://goreportcard.com/report/github.com/playnet-public/libs)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Build Status](https://travis-ci.org/playnet-public/libs.svg?branch=master)](https://travis-ci.org/playnet-public/libs)
[![Join Discord at https://discord.gg/dWZkR6R](https://img.shields.io/badge/style-join-green.svg?style=flat&label=Discord)](https://discord.gg/dWZkR6R)

The repository containing various shared libs for the entire playnet project.

## Libs

### Logging

Our logging setup using `go.uber.org/zap`.
Sentry and Jaeger are being added for production environments.

```go
l := log.New(
    "name",
    "sentryDSN",
    false,
)
defer l.Close()
```

Afterwards the logger can be used just like a default zap.Logger.
When the log level is Error or worse, a sentry message is being sent containing all string and int tags.
If you provide a zap.Error tag, the related stacktrace will also be attached.

Additionally there is a tracer(opentracing/jaeger) available in the logger which should be closed before exiting main.

### Problems

`problems` is a small library which implements the RFC7807 error response format standard for e.g. HTTP API's.

The `problems` lib provides a struct called `Problem` and an interface called `ProblemInfo`.
`Problem` implements the `error` interface, so you can simply return the problem as an error in your application.

```go
func returnError() error {
    return problems.New()
}
```

You can marshal the `Problem` struct to JSON:

```go
func makeJSON(problem *Problem) ([]byte, error) {
    return json.Marshal(&problem)
}
```

If you want to define a HTTP problem, you can do it like this:

```go
var problemNotFound = problems.New().SetTitle("NotFound").SetDetail("The requested url was not found.").SetStatus(404).SetType("https://example.com/problem/description")
```

For custom fields for more detailed problem information, you can implement `Problem` in your custom struct:

```go
type customInfo struct {
    AdditionalField1 bool `json:"additional_field_1"`
}

type customProblem struct {
    *problems.Problem
    *customInfo
}
```

## Contributions

Pull Requests and Issue Reports are welcome.
If you are interested in contributing, feel free to [get in touch](https://discord.gg/WbrXWJB)