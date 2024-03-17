FROM golang:alpine AS build

RUN apk add git

RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go build -o /tmp/otlp_backend main.go

FROM alpine:edge

COPY --from=build /tmp/otlp_backend /sbin/otlp_backend

CMD /sbin/otlp_backend
