FROM golang:alpine as builder

COPY src/main.go .

RUN go build -ldflags="-s -w" main.go

FROM scratch

COPY --from=builder /go/main .

ENTRYPOINT [ "./main" ]
