FROM golang:1.17

COPY ./server /app
WORKDIR /app

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/jinzhu/inflection
RUN go get -u github.com/mattn/go-sqlite3
RUN go get -u github.com/jinzhu/gorm

CMD ["go","run","server.go"]

