FROM golang:1.21.0-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go


FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
COPY Makefile .
RUN echo "APP.NAME=phone-book" > .env
RUN echo "APP.PORT=8080" >> .env
RUN echo "APP.LOG_LEVEL=DEBUG" >> .env
RUN echo "APP.ENVIRONMENT=PROD" >> .env
RUN echo "DATABASE.DSN=postgres://postgres:password@postgres:5432/phone-book?sslmode=disable" >> .env

EXPOSE 8080
CMD ["/app/main"]