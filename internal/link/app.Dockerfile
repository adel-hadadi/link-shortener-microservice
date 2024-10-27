FROM alpine:latest

WORKDIR /app

COPY linkShortenerApp .

COPY .env .

CMD ["./linkShortenerApp"]