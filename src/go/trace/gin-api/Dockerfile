FROM golang:alpine AS base

WORKDIR /app
COPY . .

FROM base AS gin-server
RUN go install ./app.go

CMD ["/go/bin/app"]
