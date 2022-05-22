FROM golang:1.16-alpine3.15 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download \
    && GOOS=linux GOARCH=amd64 CGO_ENABLED=0
ADD . .

RUN go build -o ./cmd/binary ./cmd/main.go

FROM alpine:3.15
ENV TZ=Asia/Jakarta \
    MYSQL_USER="root" \
    MYSQL_PASSWORD="root" \
    MYSQL_HOST="localhost" \
    MYSQL_PORT="3306" \
    MYSQL_DBNAME="cakes" \
    MYSQL_MAX_IDLE_POOL="10" \
    MYSQL_MAX_IDLE_TIME="10" \

RUN apk add --no-cache --upgrade \
    bash \
    tzdata\
    && cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN mkdir /app
COPY --from=builder /app/cmd/binary /app
WORKDIR /app
RUN chmod +x binary
EXPOSE 8000
CMD ["./binary"]
