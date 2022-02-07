# ------------------------------------------------------------------------------
# Base image
# ------------------------------------------------------------------------------
FROM golang:1.17.6-bullseye AS golang
RUN apt-get update && apt-get install git
ARG UID=1000
ARG GID=1000
ENV UID=${UID}
ENV GID=${GID}
RUN getent passwd $UID || (groupadd --gid $GID user && useradd --uid $UID --gid user --shell /bin/bash --create-home user)
USER $UID
WORKDIR $GOPATH
COPY . $GOPATH/src
ENV GO111MODULE="on" \
    GOOS=linux
#RUN cd $GOPATH/src && go mod vendor
#RUN cd $GOPATH/src && go mod tidy
#RUN cd $GOPATH/src && go mod download

# ------------------------------------------------------------------------------
# Test image
# ------------------------------------------------------------------------------
FROM golang AS golang_test
WORKDIR $GOPATH/src
ENTRYPOINT go go test ./...

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------
FROM golang AS golang_dev
ENTRYPOINT go run .\cmd\reminder\main.go

# ------------------------------------------------------------------------------
# Build image
# ------------------------------------------------------------------------------
FROM golang_dev AS golang_build
RUN go build -mod=vendor -o ./build/bin/main ./cmd/bin/main.go && go build -mod=vendor -o ./build/bin/migration ./cmd/bin/main.go
RUN chmod +x ./build/bin/main && chmod +x ./build/bin/migration


# ------------------------------------------------------------------------------
# Production image
# ------------------------------------------------------------------------------
FROM alpine:3.7 as golang_prod
COPY --from=golang_build /go/build/bin/ ./
EXPOSE ${APP_PORT}
CMD ["./main"]