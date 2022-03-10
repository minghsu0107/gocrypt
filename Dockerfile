FROM golang:1.17 as builder
WORKDIR /app
COPY . .
ARG VERSION
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build --ldflags "-X github.com/minghsu0107/gocrypt/cmd.Version=`echo $VERSION` -s -w" -a -o ./output/gocrypt ./gocrypt.go

FROM alpine:3.14
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
ENV HOME /home/appuser/
COPY --from=builder /app/output/gocrypt /app/gocrypt

ENTRYPOINT ["/app/gocrypt"]
