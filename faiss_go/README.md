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
```$xslt
 go run main.go 
```

### Regenerate go bind

download c-for-go

```$xslt
go get github.com/xlab/c-for-go

```

under faiss

```$xslt
c-for-go faiss.yml
```
