
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

# currently no go.sum file is needed
COPY go.sum ./   

# this is currently not needed
RUN go mod download

COPY . .

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]