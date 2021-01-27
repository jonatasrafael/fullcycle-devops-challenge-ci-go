FROM golang

COPY . .

RUN go build ./src/main.go

ENTRYPOINT [ "./main"]
