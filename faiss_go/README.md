## BUILD

### Step 1: Compiling the C++ Faiss

[link](https://github.com/facebookresearch/faiss/blob/master/INSTALL.md#step-1-compiling-the-c-faiss)

### Step 2: Compiling the C api

under faiss/c_api
```$bash
make
```

### TEST
under faiss/faiss_go/example
```
 go run main.go 
```

see Docker Example for complete guide

### Regenerate go bind

download c-for-go

```
go get github.com/xlab/c-for-go

```

under faiss

```
c-for-go faiss.yml
```

### Docker Example

under faiss

```
docker build -t faiss -f faiss_go/docker/example/openblas/Dockerfile .
```

```
docker run -it faiss
```