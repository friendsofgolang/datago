FROM golang:latest

ADD . /go/src/datago

RUN cd /go/src/datago \
    && curl https://glide.sh/get | sh \
    && glide install

RUN go install datago

ENTRYPOINT /go/bin/datago

EXPOSE 8080

