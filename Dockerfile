# Build container
FROM golang:latest AS builder

ENV GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

# Runtime Container
FROM alpine
RUN apk update && apk --no-cache add ca-certificates

COPY --from=builder /build ./

ENTRYPOINT ["./main"]