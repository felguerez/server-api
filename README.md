# server-api

Building an API backend with Golang. I'm using Fiber as a web framework that looks a lot like Node's Express.

## How to set up
1. Install Golang 1.20.3^ (instructions at [go.dev](https://go.dev/doc/install))
2. `git clone git@github.com:felguerez/server-api.git`
3. `cd server-api/`
3. `go get .`
4. `go build`
5. `go run`
4. You've got an app running on http://localhost:3000

## Svelte embedded client

This project uses a SvelteKit Embed module based on the [gofiber/recipes example](https://github.com/gofiber/recipes/tree/master/sveltekit-embed). The Svelte client is found in the top-level `client` directory.

### Running the client
TODO: instructions for building and running the client