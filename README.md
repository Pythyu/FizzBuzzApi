# FizzBuzzApi REST API

This is a Go REST API with two endpoints `/v1/fizzbuzz` and `/v1/stats`
that aim to compute customized FizzBuzz sequence. It is fully ready for 
production and equipped of tooling to allow an easier maintenance by
other developers.

More details on the repository design [here](wiki/Design.md)

This API is publicly available on [http://51.44.217.74:8989/v1/{endpoint}](http://51.44.217.74:8989/swagger/index.html)

## Main Features
 - REST API
 - Input Validation
 - Rate limiting
 - Statistics endpoint
 - OpenAPI documentation
 - Docker support
 - DockerHub image deployment
 - CI/CD
 - Units tests
 - Error translation

More details on these features [here](wiki/Features.md)

## Installation

If you want to simply run it without building it, you can find each commit's docker image on this [dockerhub registry](https://hub.docker.com/repository/docker/marethyuu/public_storage/general).

### Requirements
 - Go 1.26.4
 - Docker (optional)

### Clone
```bash
git clone https://github.com/Pythyu/FizzBuzzApi.git
```

### Run with docker
```bash
docker compose build
```
Once built, you can run the image locally with
```bash
docker compose up
```

### Run locally without docker
```bash
go build -o ./bin/api ./cmd/api && ./bin/api
```

## Configuration
The server configuration can be modified with runtimes environment variables

| Variable             | Default | Description                                                                                                                                                                      |
|----------------------|---------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| SERVER_PORT          | 8989    | Port the Go REST API binary is hosting on.<br/>_Note it is meant when running outside of docker.<br/>Since compose config redirect the default(`8989`) to `EXPOSED_SERVER_PORT`_ |
| SERVER_TIMEOUT_READ  | 3s      | How long the API is willing to wait for a client to read a request.                                                                                                              |
| SERVER_TIMEOUT_WRITE | 5s      | How long the server is allowed to response back.                                                                                                                                 |
| SERVER_TIMEOUT_IDLE  | 5s      | The maximum amount of time to wait for the next request.                                                                                                                         |
| SERVER_DEBUG         | false   | Enable/Disable server debug's log (Currently Unused)                                                                                                                             |

We also have two build-time environment variables, only used for docker image

| Variable            | Default      | Description                                                                                                                                                                                                                    |
|---------------------|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| EXPOSED_SERVER_PORT | 8080         | Used to redirect the docker default exposed port of 8989 to any port we want to expose our API on.<br/>_On the local dev environment it's set as 8989(see `.env`) to avoid conflict with existing 8080 server like nginx_      |
| IMAGE_TAG           | **Required** | Tag used to build our docker image. Set as `local_dev_test` in our dev environment.<br/>It is set as `{branch_name}_{full_sha}_{pipeline_id}` for main commits. The latest one is also tagged to `latest_dev` for easier acces |


## Api Documentation

You can access the OpenAPI endpoint on this [public URL](http://51.44.217.74:8989/swagger/index.html)

### GET /v1/fizzbuzz

Returns a fully customizable FizzBuzz sequence

| Parameter       | Type    | Description                     |
|-----------------|---------|---------------------------------|
| first_multiple  | integer | First multiple for fizz string  |
| second_multiple | integer | Second multiple for buzz string |
| limit_integer   | integer | Limit number of the sequence    |
| fizzString      | string  | Fizz replacement                |
| buzzString      | string  | Buzz replacement                |

### GET /v1/stats

Returns the parameters corresponding to the most frequently requested FizzBuzz computation and the number of times it has been called.

No parameters.