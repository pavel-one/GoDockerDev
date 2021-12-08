FROM golang:alpine

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

RUN go get github.com/joho/godotenv
RUN go mod download

RUN apk add nodejs
RUN apk add nano
RUN apk add --update npm
RUN npm install -g nodemon