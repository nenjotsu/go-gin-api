# FROM golang:alpine

# ADD . /go/src/github.com/gin-gonic/gin
# ADD . /go/src/github.com/itsjamie/gin-cors
# ADD . /go/src/golang.org/x/crypto/bcrypt
# ADD . /go/src/gopkg.in/appleboy/gin-jwt.v2
# ADD . /go/src/gopkg.in/mgo.v2
# ADD . /go/src/gopkg.in/mgo.v2/bson

# ADD . /go/src/github.com/marcidblue-sales-api
# RUN go install github.com/marcidblue-sales-api
# CMD ["/go/bin/marcidblue-sales-api"]
# EXPOSE 7324

FROM golang:alpine 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
EXPOSE 7324
CMD ["/app/main"]
