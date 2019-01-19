FROM golang

RUN apt-get update&&apt-get install -y libopenblas-dev

ADD . /go/src/github.com/facebookresearch/faiss

WORKDIR /go/src/github.com/facebookresearch/faiss

ENV CGO_LDFLAGS -L/usr/local/lib -lstdc++ -pthread -lcrypt -lm -fopenmp -lblas -llapack

RUN ./configure&&make

RUN cd c_api&&make

RUN cp  c_api/libfaiss_c.so /lib/

RUN cd faiss_go&&go build example/main.go

CMD ./faiss_go/main