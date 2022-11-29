FROM golang:1.19.3-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY /app/go.mod .
COPY /app/go.sum .

RUN go mod download

COPY /app .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/tg-sender .

# Start fresh from a smaller image
FROM alpine:3.17

COPY --from=build_base /tmp/app/out/tg-sender /app/tg-sender

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/tg-sender"]
