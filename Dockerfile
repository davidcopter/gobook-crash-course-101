# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -a -installsuffix cgo -o gobook-build-file ./cmd/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/gobook-build-file /bin/app

EXPOSE 1323

USER nonroot:nonroot

ENTRYPOINT ["app"]