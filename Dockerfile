FROM golang:1.19-alpine as builder
WORKDIR /opt/build/
COPY . .
RUN mkdir /opt/app/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o /opt/app/minecraft-server .

WORKDIR /opt/app/
CMD ["./minecraft-server"]
