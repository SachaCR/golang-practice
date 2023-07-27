# Golang & Monorepo experimentations

Please do not take this repository as a reflection of what I would do for a production service. It's just a playground where I experiment ideas.

This repository is only dedicated to learn Golang and also experiment with a monorepo configuration. This is why there an empty frontend app in the repo.

So far I tried to build a small web service that exposes an API using Gin Gonic framework. I've no real idea if I'm following the Golan best practices or not. For me the objective is to be able to build the same kind of web service I can build with Typescript today. With the same level of quality, testing, decoupling, clean architecture, code coverage, etc... 

So far I've lean and experimented the following things with Go:

- Isolate my package internal implementation by using `internal` packages.
- Build customer errors 
- Create enums with iota
- Use interfaces and struct to build objects and encapsulate their states
- Use the viper package to manage my app configuration
- Implement tests in Golang
- Use the Golang native tests runner
- Add a more friendly test reporter
- Use Taskfile to manage my app scripts
- Use Auth0 on the front end to authenticate users

I was currently starting playing with authorization but this is not finished yet.

Next step to experiment with:
- Extract Authorization logic properly into a different package
- Apply authorization logic with middlewares
- Use Auth0 to propagate roles into the app
- Check roles to authorize actions
- Interact with a Postgres database with Gorm
- Find a use case to play with Go concurency concepts like: goroutine, channels, etc...

# Run tests:

`backend/api-server/$ task test`

# Run coverage:
`backend/api-server/$ task coverage`

# Run Service in local 

`backend/api-server/$ task local`

# Run Service 

It expect the GO_ENV environement variable to be set to a proper environment value.

`backend/api-server/$ task run`