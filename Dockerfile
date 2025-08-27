FROM golang:1.21 as builder

WORKDIR /fania-bot
COPY . .
RUN go mod tidy

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

# Start a new image for production without build dependencies
FROM alpine
# RUN apk add --no-cache ca-certificates

# Copy the app binary from the builder to the production image
COPY --from=builder /fania-bot /fania-bot

# Run the app when the vm starts
CMD ["./fania-bot/app"]