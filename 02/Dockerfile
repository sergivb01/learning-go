FROM golang:1.7.3 as builder

WORKDIR /go/src/github.com/sergivb01/go-exercices/02
COPY * ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/sergivb01/go-exercices/02/app .
CMD ["./app"]

EXPOSE 8080
