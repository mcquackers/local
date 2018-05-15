FROM golang:1.9

ENV WORKDIR $GOPATH/src/github.com/mcquackers/local

WORKDIR $WORKDIR
COPY . $WORKDIR

RUN apt-get update && apt-get install -y git
RUN go get -u github.com/golang/dep/cmd/dep && dep ensure

CMD ["echo", "y u do dis"]
