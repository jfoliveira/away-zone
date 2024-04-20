# simple-hcs

A very simple automated health check system

# Implementation Plan, TO DO:

## Implement a web application

Not clear on the challenge description if needed or I misundertood.
Implement a simple HTTP server app, just in case.

- [ ] Expose endpoints `GET /health` and `GET /status`
- [ ] Return random sample data to simulate `database connectivity status`
- [ ] Log response errors
  - [ ] On error, trigger alert notification

## Write a client script app for health check

- [ ] query a single health check endpoint, using HTTP protocol
- [ ] get the list of endpoints from a `.env` file or env vars
- [ ] get the check interval from external configuration file
- [ ] Collect request/response metrics
- [ ] make it async to query two endpoints simultaneously in random order

## Contaneirize the health check system

- [ ] Build docker image
- [ ] Use docker compose (?) if needed and time allows me to get there

## Provision infrastructure for application

- [ ] IaC-fy infra

## Package the solution with helm

- [ ] Build a helm chart
- [ ] Deploy

## Prepare presentation, simple slides

- [ ] Plan presentation steps. Prepare a few slides, if needed
