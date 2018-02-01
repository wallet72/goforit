# install go...
FROM golang:latest

# make the work dir and fill it with code goodness
RUN mkdir /helloworld
ADD . /helloworld

# move the work dir, build the app and run it
WORKDIR /helloworld
RUN go build -o helloworld .
CMD ["/helloworld/helloworld"]
