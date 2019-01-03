FROM golang

RUN apt-get update&&apt-get install -y libblas-dev liblapack-dev

ADD . /go/src/github.com/faiss

WORKDIR /go/src/github.com/faiss

RUN ./configure&&make

RUN cd c_api&&make

RUN cd faiss_go&&go build example/main.go

CMD ./faiss_go/main