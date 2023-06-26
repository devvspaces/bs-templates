# Todo Application

This is a simple todo application built with Angular and Golang.

## Tech stack

- [Golang](https://golang.org/) - Backend API
- [Angular](https://angular.io/) - Frontend framework
- [Jaegar Tracing](https://www.jaegertracing.io/) - Distributed tracing

## Databases

You can use either mongodb or redis as your database

- [MongoDB](https://www.mongodb.com/) - NoSQL database
- [Redis](https://redis.io/) - In-memory data structure store

## Application

API - [README.md](./app/README.md).

Web - [README.md](./app/web/README.md).

## Bunnyshell Templates

This project bunnyshell templates are provided in the [.bunnyshell/templates](./.bunnyshell/templates) directory.

Each template has a README.md file with more information about the template.

### Templates

- [Golang + MongoDB + Jaegertracing (Staging/Remote Development)](./.bunnyshell/templates/golang-mongo-jaegartracing-dev/README.md)
    This template is used for development. It will start a mongodb and jaegertracing container and will start the api and web in development mode. It also contains information and configurations that will help you setup your remote development environment. You would be able to work on your local machine and run the application in the remote environment in Bunnyshell.

- [Golang + MongoDB + Jaegertracing (Production)](./.bunnyshell/templates/golang-mongo-jaegartracing-prod/README.md)
    This template is used for production. It will start a mongodb and jaegertracing container and will start the api and web in production mode.

- [Golang + Redis + Jaegertracing (Staging/Remote Development)](./.bunnyshell/templates/golang-redis-jaegartracing-dev/README.md)
    This template is used for development. It will start a redis and jaegertracing container and will start the api and web in development mode. It also contains information and configurations that will help you setup your remote development environment. You would be able to work on your local machine and run the application in the remote environment in Bunnyshell.

- [Golang + Redis + Jaegertracing (Production)](./.bunnyshell/templates/golang-redis-jaegartracing-prod/README.md)
    This template is used for production. It will start a redis and jaegertracing container and will start the api and web in production mode.

## Credits

Thanks to [mecitsemerci](https://github.com/mecitsemerci) for creating the [Todo Application](https://github.com/mecitsemerci/go-todo-app) that was used as a base for this project.
