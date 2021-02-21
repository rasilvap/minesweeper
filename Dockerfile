FROM golang:latest AS builder
RUN apt-get update 
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /Users/Usuario/all-repos/minesweeper-API
COPY go.mod .
RUN go mod tidy
COPY . .
RUN go build

FROM scratch
COPY --from=builder /Users/Usuario/all-repos/minesweeper-API .
ENTRYPOINT [ "./minesweeper-service" ]
