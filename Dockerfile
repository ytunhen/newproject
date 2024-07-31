FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go build -o message-processor .

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/message-processor .
CMD ["./message-processor"]
