# marvel-go

## Install Dependencies
```
> go get -t ./...
```

## Adding Key Config
Add a new file `/config/key/key_config.go`
Get your public/private key from https://developer.marvel.com/account
```
package key

var PublicKey string = ""
var PrivateKey string = ""
```

## Running Locally
```
> go install
> marvel-go
```

## Running with Docker Locally
Pre-requisite: Install Docker
```
// builds the image, copies the source, and builds the application
> docker-compose build
```
```
// start the container, http://[docker_machine_ip]:9000/v1/characters
> docker-compose up -d
```
```
// stops and tears down the container
> docker-compose down
```

## Running Unit Tests
```
> goconvey
```

## Debugging
```
> go get github.com/mailgun/godebug
// set breakpoints in any file, see https://github.com/mailgun/godebug
> godebug run -instrument=github.com/trwalker/marvel-go/auth,github.com/trwalker/marvel-go/characters,github.com/trwalker/marvelo/config,github.com/trwalker/marvel-go/controllers,github.com/trwalker/marvel-go/rest main.go
```
