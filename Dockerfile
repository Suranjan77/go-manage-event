# syntax=docker/dockerfile:1

# Build
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /go-manage-event

#Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-manage-event /go-manage-event

EXPOSE 8080

USER noroot:nonroot

ENTRYPOINT ["/go-manage-event"]