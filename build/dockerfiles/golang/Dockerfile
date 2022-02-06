# ------------------------------------------------------------------------------
# Base image
# ------------------------------------------------------------------------------
FROM golang:latest AS golang
WORKDIR $GOPATH/src
ARG UID=1000
ARG GID=1000
ENV UID=${UID}
ENV GID=${GID}
RUN getent passwd $UID || (groupadd --gid $GID user && useradd --uid $UID --gid user --shell /bin/bash --create-home user)
USER $UID
COPY ./ ./
WORKDIR $GOPATH/src

# ------------------------------------------------------------------------------
# Test image
# ------------------------------------------------------------------------------
FROM golang AS golang_test

RUN apk update && apk upgrade && apk add --no-cache git
RUN go mod vendor
RUN go mod tidy
RUN go mod download
ENV GO111MODULE="on" \
    GOOS=linux
CMD ["go", "test"]

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------
FROM golang AS golang_test
RUN go build -mod=vendor -o .\build\bin\main .\cmd\bin\main.go && go build -mod=vendor -o .\build\bin\migration .\cmd\bin\main.go
RUN chmod +x .\build\bin\main && chmod +x .\build\bin\migration
CMD ["./main"]

# ------------------------------------------------------------------------------
# Production image
# ------------------------------------------------------------------------------
FROM alpine:3.7 as golang_img
COPY --from=golang /go/build/bin/ ./
EXPOSE 8080
ENTRYPOINT ["./main"]