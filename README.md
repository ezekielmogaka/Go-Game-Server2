### Your Own Games Store

Your task will be to use the included game data JSON from the Games Store to create both a web server container and CLI tool to interface with it.  Your web service will be expected to run locally as a Docker container for review. You are also asked to produce a valid Kubernetes specification that can be applied to run this web service on a Kubernetes cluster.

A placeholder `Dockerfile` and `Makefile` with a few comments have been created for you to get you started.

#### Web Service

Please produce a simple unauthenticated web service that serves the data found in `games.json`.  This service should:

- Run in a container
- Have tests
- Listen on [localhost:8080](http://localhost:8080)
- Handle requests to `/` by serving a list of games with the following information:
  - `title`
  - `description`
  - `id`
- Handle requests to `/game?id=${ID}` with the following information for the specific game ID passed:
  - `title`
  - `description`
  - `id`
  - `currentPrice`
  - `sellerName`
  - `developerName`
  - `publisherName`
  - `thumbnailURL`
- `make build` should produce this container image
- `make test` should run tests for this program
- `make run` should run this web server (it does not have to be backgrounded)
- `make kubernetes` should produce a working Kubernetes specification for deploying this application with multiple replicas to Kubernetes
