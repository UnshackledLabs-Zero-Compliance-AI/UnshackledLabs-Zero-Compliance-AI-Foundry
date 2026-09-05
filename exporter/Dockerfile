FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init exporter && go get github.com/prometheus/client_golang/prometheus && go build -o exporter .

FROM scratch
COPY --from=builder /app/exporter /exporter
EXPOSE 9090
ENTRYPOINT ["/exporter"]