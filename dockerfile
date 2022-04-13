FROM golang:1.18 as builder
WORKDIR /build
COPY . .
RUN go build -o app server/main.go &&\
  go install github.com/cespare/reflex@latest

FROM alpine:latest
WORKDIR /run
ENV GIN_MODE=release
COPY --from=builder /build/app ./
EXPOSE 8080
ENTRYPOINT ["./app"]