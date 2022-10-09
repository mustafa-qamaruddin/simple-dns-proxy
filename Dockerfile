FROM golang:1.19.1-alpine as builder

RUN apk update && apk add git

WORKDIR $GOPATH/src/simple-dns-proxy/

ENV GOPATH /go

COPY . /go/src/simple-dns-proxy/

RUN go get -d -v

RUN go build -o /go/bin/simple-dns-proxy

# use one image to produce some artifacts, which are then placed into another,
# much smaller image, containing only the parts necessary for running the
# artifacts that we’d built
FROM alpine AS final

WORKDIR /app

COPY --from=builder /go/bin/simple-dns-proxy /app/simple-dns-proxy

RUN apk add libcap && setcap 'cap_net_bind_service=+ep' /app/simple-dns-proxy

# first user after root: 1000
USER root

EXPOSE 53

CMD ["/app/simple-dns-proxy"]
