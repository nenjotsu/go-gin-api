FROM golang:alpine 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
EXPOSE 7324
CMD ["/app/main"]
