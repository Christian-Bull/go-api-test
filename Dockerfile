# gets latest golang image
FROM golang:alpine AS builder

# needed for go modules
ENV GO111MODULE=on
# ENV CGO_ENABLED=0 # not sure what this does yet

# # sets env vars for host
ARG TARGETOS
ARG TARGETARCH

RUN mkdir /bin/app
ADD . /bin/app

WORKDIR /bin/app

# gets go modules
RUN go mod tidy -v
RUN go mod download

# builds app
RUN GOARCH=$TARGETARCH GOOS=$TARGETOS go build -o /go/bin/go-api-test .

COPY /go/bin/go-api-test /bin/go-api-test

# run it
ENTRYPOINT ["/bin/go-api-test"]