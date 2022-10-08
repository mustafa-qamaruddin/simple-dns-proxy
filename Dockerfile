FROM golang:1.14.2-alpine3.11 as builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/croc-hunter/

ENV GOPATH /go

COPY croc-hunter.go /go/src/croc-hunter/
COPY static/ static/

RUN go get -d -v

RUN go build -o /go/bin/croc-hunter

FROM alpine AS final

WORKDIR /app

COPY static/ static/

COPY --from=builder /go/bin/croc-hunter /app/croc-hunter

RUN apk add libcap && setcap 'cap_net_bind_service=+ep' /app/croc-hunter

USER 1000

EXPOSE ${PORT}

CMD ["/app/croc-hunter"]
