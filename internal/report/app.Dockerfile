FROM alpine:latest

WORKDIR /app

COPY reportApp .

RUN mkdir "storage"

COPY .env .

CMD ["./reportApp"]
