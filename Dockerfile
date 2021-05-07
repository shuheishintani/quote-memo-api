FROM golang:1.16
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get -u github.com/cosmtrek/air