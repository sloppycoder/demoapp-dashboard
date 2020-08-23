FROM golang:1.14 as builder

WORKDIR /build
ADD . /build

RUN go build -o bin/server cmd/server.go

FROM gcr.io/distroless/base

COPY --from=builder /build/bin/server /server

EXPOSE 8090

CMD ["/server"]
