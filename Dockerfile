FROM golang:1.18-alpine as builder
RUN apk --no-cache add ca-certificates git build-base
WORKDIR /app

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
# CGO_ENABLED=0
RUN go build -v ./cmd/calc-bot

# Create final image
# FROM alpine
# WORKDIR /
# COPY --from=builder /calc-bot . err
EXPOSE 8080
CMD ["./calc-bot"]