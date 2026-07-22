FROM golang:1.27rc2-bookworm

WORKDIR /app

COPY . .

#RUN go get -d -v ./...

RUN go build -o api

EXPOSE 5432

CMD ["./api"]
