# 실행용 Dockerfile
FROM devops_example/go-builder:1.24.5 AS builder

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
