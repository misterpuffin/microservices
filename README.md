# Microservices

This repository was used as a sample application for my sharing about Microservices architecture during CVWO.
It showcases how an application can be containerized into several individual microservices which can communicate between one another
through REST api. In this application, we have a simple frontend which communicates with a broker service.
The broker service can then rely on the services of:
- authentication microservice
- logger microservice
- mailer microservice

There are a few exercises with the instructions written in `project/todo.md`. After completing the first exercise, you
should checkout to the other branches to see the solution and try out the newer instructions.

### Task 1: Add mail-service to docker compose file
### Task 2: Implement gRPC in our app


## Navigating the application

Building the application
> cd project && make up_build

Running the application
> cd project && make up
