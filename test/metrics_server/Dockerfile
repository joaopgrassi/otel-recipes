FROM golang:alpine AS build

RUN apk add git

RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go build -o /tmp/metrics_server main.go

FROM alpine:edge

COPY --from=build /tmp/metrics_server /sbin/metrics_server

CMD /sbin/metrics_server
