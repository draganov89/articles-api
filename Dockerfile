
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

# currently no go.sum file is needed
#COPY go.sum ./   

# this is currently not needed
#RUN go mod download

COPY . .

RUN go build -o /articles-api ./cmd/cli/main.go

EXPOSE 8092

CMD [ "/articles-api" ]