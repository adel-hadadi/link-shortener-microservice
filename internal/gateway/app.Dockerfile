FROM alpine:latest

WORKDIR /app

COPY gatewayApp .

COPY .env .

CMD ["./gatewayApp"]