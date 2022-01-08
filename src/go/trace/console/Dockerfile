FROM golang:1.17-alpine3.15
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o go-recipe
ENTRYPOINT [ "./go-recipe" ]
