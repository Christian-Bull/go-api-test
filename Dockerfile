# gets latest golang image
FROM golang:alpine AS builder

# needed for go modules
ENV GO111MODULE=on
# ENV CGO_ENABLED=0 # not sure what this does yet

# # sets env vars for host
ARG TARGETOS
ARG TARGETARCH

RUN mkdir /bin/app
ADD cmd /bin/app

WORKDIR /bin/app

# gets go modules
RUN go mod tidy -v
RUN go mod download

# builds app
RUN GOARCH=$TARGETARCH GOOS=$TARGETOS go build -o ./go-api/.
EXPOSE 5050

# run it
ENTRYPOINT ["/bin/app/go-api"]