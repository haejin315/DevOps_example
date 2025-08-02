# FROM golang:1.24.5-alpine

# WORKDIR /app
# COPY . .

# RUN go build -o main .

# EXPOSE 8080
# CMD ["./main"]

# 1단계: 빌드용 환경
FROM golang:1.24.5-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o main .

# 2단계: 실행용 최소 환경
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
