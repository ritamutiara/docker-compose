FROM golang:1.18-alpine

RUN apk add --no-cache curl

WORKDIR /app

COPY main.go .

ENV APP_PORT=8080
ENV MODE=local

EXPOSE 8080

CMD ["go", "run", "main.go"]
