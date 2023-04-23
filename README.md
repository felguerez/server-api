# server-api

Building an API backend with Golang, I'm using Fiber as a web framework that looks a lot like Express from Node.

## Getting started

### Install Golang

This project uses Go v1.20.3. You can find  installation instructions at [go.dev](https://go.dev/doc/install) or install with homebrew:

```shell
brew install go
```

### Installation

1. Clone the repo
```shell
$ git clone git@github.com:felguerez/server-api.git
```
2. Sync dependencies
```shell
$ cd server-api/
$ go get .
```
3. Build the project
```shell
$ go build
```
4. Run 🏃‍♂️ (note the module name is `web-service`)
```shell
$ go run web-service 
```
You can run the project with [air](https://github.com/cosmtrek/air/) for nodemon-style live reloading:
```shell
$ go install github.com/cosmtrek/air@latest
$ echo -e "\nalias air=$(go env GOPATH)/bin/air"
$ air # runs and reloads the app using .air.toml config
```

5. Congrats! 🍾 You've got an app running on http://localhost:3000

### Running the client
TODO: instructions for building and running the client