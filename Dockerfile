FROM golang:1.20.5 AS builder

WORKDIR /work

# Download dependencies first to optimize Docker caching
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build The binary
COPY . .
RUN make build

# Begin final image
FROM alpine:latest

RUN apk --no-cache add gcompat ca-certificates nano

COPY --from=builder /work/kafctl /usr/local/bin/

WORKDIR /work
ENTRYPOINT [ "sh" ]
