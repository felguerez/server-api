# server-api

Building a monolithic API backend with Golang. I'm using [gofiber/fiber](https://github.com/gofiber/fiber) as a web
framework, which looks a lot like [Express](https://github.com/expressjs/express).

## Getting started

Time to set up `web-service`!

### Prerequisites

#### Environment
Before you get the project running you'll need to configure your environment. The app
uses [dotenv](https://github.com/joho/godotenv) to autoload values
defined in [.env](/.env.example). You can also pass these values as environment variables when running the app if you
prefer.

#### Golang

This project uses Go `v1.20.3`. You can find installation instructions at [go.dev](https://go.dev/doc/install) or
install with homebrew:

```shell
brew install go
```

#### AWS/DynamoDB

You should have an AWS account with access keys (AWS_SECRET_ACCESS_KEY, AWS_SECRET_KEY_ID) stored in [.env](/.env.example).
You'll need AWS for storage using DynamoDB. 

After creating a DynamoDB table, add the table name to [.env](.env.example).

#### Spotify Web API

This project also uses the [Spotify Web API](https://developer.spotify.com/documentation/web-api). You'll need to create
an app at [Spotify for Developers](https://developer.spotify.com/dashboard). Add your app's client ID and client secret
to the [.env](/.env.example) file.

#### Gmail App Password

This project sends email via SMTP authenticated through Gmail App Passwords. You can send emails through your own Gmail
account by setting up an [App Password](https://support.google.com/mail/answer/185833?hl=en) which is available if you
have 2FA enabled on your Google account. After generating an App Password in your Google Account under the Security
settings, add it to [.env](/.env) as `SMTP_PASSWORD`, along with your Gmail address as `SMTP_USERNAME`.

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
5. Create an app at [developer.spotify.com](https://developer.spotify.com/dashboard) and copy the client id and client
   secret to `.env` as `SPOTIFY_CLIENT_ID` and `SPOTIFY_CLIENT_SECRET` respectively.
6. Build the project:

```shell
$ go build
```

7. Run 🏃‍♂️ (note the module name is `web-service`):

```shell
$ go run web-service 
```

Pro-tip: You can run the project with [air](https://github.com/cosmtrek/air/) for nodemon-style live reloading:

```shell
$ go install github.com/cosmtrek/air@latest
$ echo -e "\nalias air=$(go env GOPATH)/bin/air"
$ air # runs and reloads the app using .air.toml config
```

8. Congrats! 🍾 You've got an app running on http://localhost:3000. You can also view the API docs
   at http://localhost:3000/swagger.

## Deployment

This project deploys continuously via [Google Cloud Run](https://cloud.google.com/run):

* [🔗 root index](https://server-api-c4m6jglxsq-uc.a.run.app/)
* [🔗 `/api/spotify/recently-played`](https://server-api-c4m6jglxsq-uc.a.run.app/api/spotify/recently-played)
* [🔗 `/api/spotify/currently-playing`](https://server-api-c4m6jglxsq-uc.a.run.app/api/spotify/currently-playing)
* [🔗 `/api/spotify/top/artists`](https://server-api-c4m6jglxsq-uc.a.run.app/api/spotify/top/artists)
* [🔗 `/api/spotify/top/tracks`](https://server-api-c4m6jglxsq-uc.a.run.app/api/spotify/top/tracks)

### TODO

* Register a domain
* Deploy a client