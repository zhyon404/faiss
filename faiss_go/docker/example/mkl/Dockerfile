FROM golang

COPY --from=robbyjo/ubuntu-mkl:16.04-2018.1 /opt/ /opt/

ADD . /go/src/github.com/facebookresearch/faiss

WORKDIR /go/src/github.com/facebookresearch/faiss

RUN cp /opt/intel/compilers_and_libraries/linux/mkl/lib/intel64/*.so /lib

ENV MKLROOT /opt/intel/compilers_and_libraries/linux/mkl/

ENV CGO_LDFLAGS -L/opt/intel/compilers_and_libraries/linux/mkl/lib/intel64 -lmkl_intel_lp64 -lmkl_gnu_thread -lmkl_core -lgomp -lpthread -lm -ldl -fPIC -fopenmp -m64 -Wno-sign-compare -g -O3 -Wall -Wextra -msse4 -mpopcnt

RUN ./configure

RUN make

RUN cd c_api&&make

RUN cp  c_api/libfaiss_c.so /lib/

RUN cd faiss_go&&go build example/main.go

CMD ./faiss_go/main