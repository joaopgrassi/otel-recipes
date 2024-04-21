# OTel recipes contributing guide

First, welcome to OTel recipes, we are glad you are here!

This document will help you to get started in contributing to OTel recipes.

## Code of Conduct

Please make sure to read and observe our [Code of Conduct](./CODE_OF_CONDUCT.md).

## Required Tools

Working with the project sources requires the following basic tools:

1. [git](https://git-scm.com/)
2. [go](https://golang.org/) (version 1.21 and up)
3. [docker](https://www.docker.com/)

Depending on the programming language there may be other requirements, such as .NET SDK
JDK, Node.js and etc.

## Project architecture

OTel recipes can be split into two parts: The collection of sample applications (recipes) and the website.

### Recipes

The recipes are organized by programming language and telemetry signal.
An example of directory structure for a recipe is:

```shell
src/
└── csharp/
    └── traces/
        └── console/
            ├── [app files]
            ├── Dockerfile
            ├── docker-compose.yaml
            ├── collector-config.yaml
            ├── recipefile.json
            └── test/
                ├── go.mod
                ├── go.sum
                └── traces_test.go
```

Each recipe MUST:

- Be containerized and define a `Dockerfile`
- Define a `recipefile.json` (See below for more info)
- Be configured to export OTLP data to an OpenTelemetry collector
  - Can be via `gRPC` or `HTTP`
  - Thus the need to have a `collector-config.yaml` and a collector container inside the compose file
- Be a container inside the `docker-compose.yaml` file
  - Also declare any dependency it uses. E.g., database, messaging system etc.
- Be testable. Each app must achieve a goal (e.g. record a span) and this goal MUST be tested e2e. See [Testing a recipe](#testing-a-recipe) below

#### The `recipefile.json` file

The `recipefile.json` is a JSON schema file that contains metadata about the recipe.
It is required that each sample defines its recipe file. The recipe file is based on the schema file
found in the root of the repository: [otel-recipes-schema.json](./otel-recipes-schema.json).

The recipe file is the core of OTel recipes. The website contains all recipe files together,
essentially being the "database" the website uses to show up the recipes.

Here is an overview of the most important fields in the recipe file:

- `id`: Follow the convention of `{language}.{shortnameofapp}.{signal}`. For example for a C# console app that exports traces, `csharp.console.metrics`
- `languageId`: Stick with the schema allowed values
- `signal`: Stick with the schema allowed values
- `displayName`: The name of the app that will be shown in the website.
  Ideally the name should be small but still give a good idea of what the recipe has to offer
- `sourceRoot`: The root of the recipe app on GitHub. E.g. https://github.com/joaopgrassi/otel-recipes/tree/main/src/csharp/metrics/console
- `steps`: The list of steps users must follow to get OTel configured in the application
  - `displayName`: The name of each step that will be shown in the website
  - `order`: The order of the step
  - `source`: The `raw` GitHub link to the code file related to the step. E.g, https://raw.githubusercontent.com/joaopgrassi/otel-recipes/main/src/csharp/metrics/console/App.cs
- `dependencies`: The OpenTelemetry-related packages the recipe needs
  - `id`: The exact package name. E.g., `@opentelemetry/api`, `OpenTelemetry.Exporter.OpenTelemetryProtocol`
  - `version`: The version of the package. E.g., `1.7.0`

During a PR, several checks are performed against recipe files, such as unique id and schema validations

You can find an example of a recipe file in [example-recipefile.json](./src/csharp/traces/console/recipefile.json).

#### Testing a recipe

Each recipe app MUST be e2e tested. The goal is that apps are constrained to small objectives
that then can be verified e2e. This is all done automatically by the GitHub workflow `Build and Test`.
You only need to define the below and the rest is all magic ✨.

In a nutshell, the way it works is:

1. A recipe define its container file `Dockerfile`
2. The `id` on the recipe file MUST be used as the `service.name` and as for the [name of the meter](https://opentelemetry.io/docs/specs/otel/metrics/api/#get-a-meter)
3. A recipe defines a `docker-compose.yaml` file, which at minimum contain a container for the app, the collector and the OTLP back-end
4. The app produces *some* telemetry, for example a span named `HelloWorld` with an attribute `foo=bar`.
5. The span is exported to the collector, which then exports to the telemetry back-end [OTLP back-end](./internal/otlp_backend/README.md)
6. The recipe defines a `go` test, which uses the test framework from [Test utils](./internal/common/testutils/README.md)
7. The test is responsible for creating the expected data, then querying the OTLP back-end for the actual data and doing the assertions

For an example, see the [C# console app test](./src/csharp/traces/console/test/).

## Adding a new recipe app

Before starting working on a new recipe app, please open a feature request via a new issue. The app will be discussed
and once everything is clear, development can start. This is to avoid double work and to speed up reviews later on.

### Pull requests

#### Title guidelines

The title for your pull-request should contain a clear statement of the kind of app you are introducing and what signal it is focused on. For instance:

Add ASP.NET API example for traces

#### Description guidelines

The description should go into more details of what the goal of the recipe app is. For example if you want to show how to demonstrate
how to configure OpenTelemetry in an application that records attributes you can add:

> An ASP.NET API configured to export OpenTelemetry spans to a collector via gRPC. The app records a span when the endpoint `helloworld` is called.

#### Automated tests

When a PR is open, several CI checks are executed:

- `Build and Test`: Will build the recipe app and execute the go e2e tests. See [Testing a recipe](#testing-a-recipe)
- `Validate unique recipe ids`: Validates if the `id` in the `recipefile.json` is unique across all other recipes
- `Check recipe file`: Validates the `recipefile.json` against the JSON schema.
- `Validate recipe file dependencies`: Validates if the `dependencies` declared in the `recipefile.json` matches with the ones in the app.
  - This helps ensure that they are in sync. Updating a package and forgetting to update the recipefile will cause a failure

## Found a bug? Something not working?

If you found a bug or something is not correct/working as it should, please open an issue using the existing templates.
