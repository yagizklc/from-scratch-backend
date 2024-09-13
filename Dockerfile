FROM golang:1.23-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
RUN go build -o /app/bin/main /app/app/cmd/.
EXPOSE 8080
CMD ["./bin/main"]
