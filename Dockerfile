# Build the application
FROM golang:1.23-alpine as builder
RUN apk add --no-cache make

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build


# Create a minimal image to run the application
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/bin/main .

EXPOSE 8080
CMD ["./main"]
