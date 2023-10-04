FROM golang:1.21

WORKDIR /usr/src/app
EXPOSE 8080


COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

CMD ["app"]
















