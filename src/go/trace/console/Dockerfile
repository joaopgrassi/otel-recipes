FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o go-recipe
ENTRYPOINT [ "./go-recipe" ]
