FROM golang:latest

RUN mkdir /helloworld
ADD . /helloworld
WORKDIR /helloworld
RUN go build -o helloworld .
CMD ["/helloworld/helloworld"]
