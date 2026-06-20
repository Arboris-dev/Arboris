FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY go_server/ ./go_server/
COPY generated/go/ ./generated/go/

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/go_server ./go_server/main.go

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/bin/go_server .

EXPOSE 8080

CMD ["./go_server"]