FROM golang:1.22-alpine

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o app .

EXPOSE 7860

CMD ["./app"]
