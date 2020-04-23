# Template API

Template API is used as template project by flow cli.

### Prerequisites

- Install [go 1.14.x](https://golang.org)
- Install [Docker](https://docker.com/)
- Install [VS Code](https://code.visualstudio.com/)

## Getting started

### Clone the repository

```
git clone git@github.com:go-flow/template-api.git <project_name>
```

### Install dependecies

```
cd <project_name>
go get ./...
```

### Build and run the project

Run application from source

```
cd <project_name>
go run  cmd/api/main.go
```

build docker image

```
docker build -t template-api .
```

Here’s the breakdown of this command: docker build is the command that tells Docker to build an image. -t template-api sets the tag name of the Docker image to template-api, which we can reference later. Please note that in order to have a valid tag name, it must be in lowercase and has no spaces (use snake-case naming). Finally, . tells Docker where to look for the Dockerfile that is needed for the build ( . means here).

run docker image

```
docker run  -p 80:5000 --env-file ./.env  --name=template-api-container template-api
```

command above will run docker container from a image. -p 80:5000 maps port 80 of our machine to port 5000 of the container.
--env-file ./.env tels to docker to read a .env file of environment variables. --name=template-api-container gives our container the name template-api-container.

Build and run with docker-compose

```
cd <project_name>
docker-compose up
```

this command will run all services in docker and build development builds for every service. If all builds are successful services will be available on following locations

- REST API: `http://localhost:80`

## Environment variables

| Variable                       | Required | Default Value | Description                                    |
| -------------------------------| -------- | ------------- | ---------------------------------------------- |
| ENV                            | NO       | production    | indicates in which environment app is running  |
| LOG_LEVEL                      | NO       | error         | logging level                                  |
| DB_DEV_CONNECTION              | YES      | -             | Database connection string for DEV environment |
| DB_TEST_CONNECTION             | YES      | -             | Database connection for TEST environment       |
| DB_PRODUCTION_CONNECTION       | YES      | -             | Database connection for PRODUCTION environment |



## Project structure

The project ships with a directory structure like:

```
|
|---- .github                           # git related templates and script needed for the project
|---- business                          # project business package
|     |---- user_business.go            # user business logic
|     |---- init.go                     # business initialization
|---- cmd                               # project commands holder
|     |---- api                         # api  cli command package
|           |---- main.go               # api cli entry point
|---- config                            # configuration package
|     |---- config.go                   # configuration logic
|---- controllers                       # project controllers package
|     |---- base_controller.go          # base controller extended by all other controllers, contains per controller middlewares & helper methods
|     |---- users_controller.go         # users controller handles values HTTP request & Response
|     |---- index_controller.go         # index controller
|     |---- init.go                     # controllers initialization
|---- docs                              # project related documentation
|     |---- docs.go                     # automatically generated swagger documentation code
|     |---- swagger.json                # automatically generated swagger documentation code
|     |---- swagger.yaml                # automatically generated swagger documentation code
|---- kube                              # Kubernetes related configuration files
|---- middlewares                       # middlewares package
|---- models                            # models package
|     |---- paginated_model.go          # model describing paginated model response
|     |---- response_error.go           # model returned in case of an error
|     |---- response.go                 # model returned in case of a successful operation
|     |---- user.go                     # model describing user record
|---- pkg                               # project libraries package
|     |---- paging                      # package used for handling paging, sorting & filtering
|     |---- swagger                     # package used generating swagger documentation
|     |---- cors                        # package used for handling CORS
|---- repositories                      # repositories package
|     |---- init.go                     # repositories initialization
|     |---- user_repository.go          # user repository handling database communication for user record
|---- services                          # services package
|     |---- init.go                     # services initialization
|     |---- user_service.go             # user service defines set of available operations aroud Users model
|---- CODE_OF_CONDUCT.md                # Contributor Covenant Code of Conduct
|---- CONTRIBUTING.md                   # Contribution guidelines
|---- docker-compose.yml                # docker-compose configuration file
|---- Dockerfile                        # Dockerfile for Docker image
|---- e2e.go                            # e2e tests entry file
|---- go.mod                            # go module definition file
|---- go.sum                            # go module dependecies checksum
|---- Makefile                          # makefile scripts and shortcuts for build, test, run
|---- README.md                         # this file :)
```

## Project dependecies

Dependencies are managed through `go.mod`.

| Package                             | Description                               |
| ----------------------------------- | ----------------------------------------- |
| cloud.google.com/go             | Google Cloud product. libraries.                          |
| github.com/alecthomas/template      | Go’s text/template package with newline elision.                   |
| github.com/go-flow/flow             | Go web framework.                          |
| github.com/go-playground/validator/v10 | Package validator implements value validations for structs and individual fields based on tags.                  |
| github.com/pkg/errors    | Package errors provides simple error handling primitives.               |
| github.com/swaggo/files                | Generate swagger ui embedded files.  |
| github.com/swaggo/swag                | Swag converts Go annotations to Swagger Documentation 2.0.  |
| golang.org/x/net                | Supplementary Go networking libraries.  |
| google.golang.org/genproto            | Go generated proto packages.  |


## Built With

- [go](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
- [go flow](https://github.com/go-flow/flow/) - High Performance minimalist web framework for gophers


## Contributing

Please read [CONTRIBUTING.md](https://github.com/go-flow/template-api/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.
