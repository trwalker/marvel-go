FROM golang:latest

ENV GOPATH=/app/go

COPY . ${GOPATH}/src/github.com/trwalker/marvel-go/

COPY . ${GOPATH}/src/github.com/trwalker/marvel-go/

WORKDIR ${GOPATH}/src/github.com/trwalker/marvel-go/

RUN ["go", "get", "github.com/gorilla/mux"]
RUN ["go", "get", "github.com/gorilla/handlers"]

RUN ["go", "install"]

ENTRYPOINT ["/app/go/bin/marvel-go"]

EXPOSE 9000/tcp