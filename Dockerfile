FROM golang:1.21

WORKDIR /app

COPY . /app

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o main .

EXPOSE 8080

ENV REDIS_ADDR="redis:6379"
ENV REDIS_PASSWORD=""
ENV BASE_URL="http://localhost:8080"

CMD ["./main"]
