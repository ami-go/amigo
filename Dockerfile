FROM golang
ADD . /go/src/amigo/
WORKDIR /go/src/amigo/
EXPOSE 8080
CMD go run main.go