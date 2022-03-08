FROM golang:1.17 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
    go build --ldflags "-s -w" -a -o ./output/gocrypt ./gocrypt.go

FROM alpine:3.14
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
ENV HOME /home/appuser/
COPY --from=builder /app/output/gocrypt /app/gocrypt

ENTRYPOINT ["/app/gocrypt"]
