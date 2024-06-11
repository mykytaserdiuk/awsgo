FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /app/main ./cmd

FROM alpine:3.13
WORKDIR /app
RUN apk add --no-cache bash
COPY --from=builder /app/main /app/main
COPY scripts/wait-for-it.sh wait-for-it.sh

EXPOSE 1232
CMD ["./main"]