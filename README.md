# AwayZone

A very reliable application - that's `Always On` (pun intended) - relying on an automated health check system to notify the team when something is not right.

## Dependencies

- [Docker](https://www.docker.com/products/docker-desktop/)
- [make](https://www.gnu.org/software/make/)
  - `make` is pre-installed in most Linux systems.
  - In `macOS` it is included as part of the `Xcode` command line utils. It can be installed with the following command:
  ```
  xcode-select --install
  ```

## Usage

### Build and init web application and health-checker

```shell
make run-dev
```

### Stop running containers

```shell
make stop-dev
```

### Miscellaneous

To list all available commands just run:

```shell
make
```

## Implementation Plan, check list:

### Implement an HTTP server/web application

- [x] Implement a simple HTTP server app
- [ ] Expose endpoints `GET /health` and `GET /status`
- [ ] Return random sample data to simulate `database connectivity status`
- [x] Log response errors
- [ ] On error, trigger alert notification

### Write a client script app for health check

- [x] query a single health check endpoint, using HTTP protocol
- [ ] get the list of endpoints from a `.env` file or env vars
- [ ] get the check interval from external configuration file
- [ ] Collect request/response metrics
- [x] make it async to query two endpoints simultaneously in random order
- [ ] Run in intervals, instead of just once

### Contaneirize the health check system

- [x] Build docker image
- [x] Create a docker compose file

### Package the solution with helm

- [x] Build a helm chart
- [x] Organize helm values files per environment

### Deploy via IaC to an orchestrated environment:

- [x] Create a pulumi program
- [x] Use pulumi to deploy helm charts to Kubernetes

### Provision infrastructure for application

- [ ] Provision managed Kubernetes cluster, using terraform

### Documentation

- [ ] Improve README
- [ ] Document service's endpoints using OpenAPI spec

### Prepare presentation, simple slides

- [ ] Plan presentation steps. Prepare a few slides, if needed
- [ ] Perhaps list and sort acceptance criterias, plus bonus point, and go through sections that cover group of related criteria? Might be a way to get the presentation organized and move fast, giving room for questions

### README

- [ ] Improve README, branding, logo, be fun, ... AWAA ZONE, from TOW-AWAY ZONE, parking not allowed, not parking, always running, always on, ...
- [ ] Rename repo to use new name
