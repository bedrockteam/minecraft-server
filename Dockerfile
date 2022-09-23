FROM golang:1.19-alpine as builder
WORKDIR /opt/build/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o minecraft-server .


FROM scratch
WORKDIR /opt/app/
COPY --from=builder /opt/build/minecraft-server .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./minecraft-server"]
