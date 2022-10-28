FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY . .
RUN go build -o main ./cmd/web

ENV PORT 4000
EXPOSE $PORT

CMD ["./main"]



