# server-api

Building an API backend with Golang. I'm using Fiber as a web framework that looks a lot like Express from Node.

## Getting started

Hello future collaborator, welcome to `web-service`!

### Prerequisites

Before you get the project running you'll need to set up:

#### Golang
This project uses Go `v1.20.3`. You can find  installation instructions at [go.dev](https://go.dev/doc/install) or install with homebrew:

```shell
brew install go
```

#### AWS/DynamoDB
This project assumes you have an AWS account already configured with access keys in  `~/.aws/credentials`. You'll need this to create and use DynamoDB for storage, as well to deploy on AWS Lambda.

#### Spotify Web API
This project also uses the [Spotify Web API](https://developer.spotify.com/documentation/web-api). You'll need to create an app at [Spotify for Developers](https://developer.spotify.com/dashboard).

## Installation and setup (development)

1. Clone the repo:
```shell
$ git clone git@github.com:felguerez/server-api.git
```
2. Sync dependencies:
```shell
$ cd server-api/
$ go get .
```
3. Rename `.env.example` to `.env`:
```shell
$ mv .env.example .env
```
4. Create a DynamoDB table and copy the name to `.env` as `TABLE_NAME`
5. Create an app at [developer.spotify.com](https://developer.spotify.com/dashboard) and copy the client id and client secret to `.env` as `SPOTIFY_CLIENT_ID` and `SPOTIFY_CLIENT_SECRET` respectively. 
6. Build the project:
```shell
$ go build
```
7. Run 🏃‍♂️ (note the module name is `web-service`):
```shell
$ go run web-service 
```
Pro tip: You can run the project with [air](https://github.com/cosmtrek/air/) for nodemon-style live reloading:
```shell
$ go install github.com/cosmtrek/air@latest
$ echo -e "\nalias air=$(go env GOPATH)/bin/air"
$ air # runs and reloads the app using .air.toml config
```

8. Congrats! 🍾 You've got an app running on http://localhost:3000. You can also view the API docs at http://localhost:3000/swagger.

## Deployment
This project deploys continuously via [Google Cloud Run](https://cloud.google.com/run):

* [🔗 root index](https://server-api-c4m6jglxsq-uc.a.run.app/)
* [🔗 `/api/spotify/recently-played`](https://server-api-c4m6jglxsq-uc.a.run.app/api/spotify/recently-played)
* [🔗 `/api/spotify/currently-playing`](https://server-api-c4m6jglxsq-uc.a.run.app/api/spotify/currently-playing)

* **TODO: register a domain**