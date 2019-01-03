package main

import (
	"fmt"
	"github.com/faiss/faiss_go"
)

func main() {
	x()
}

func x() {
	v := new(faiss_go.Faissindexflatl2)
	fmt.Println(faiss_go.FaissIndexflatl2New(&v))
	x := (*faiss_go.Faissindex)(v)
	faiss_go.FaissIndexIsTrained(x)
	s := [10]float32{1, 2, 3}
	fmt.Println(faiss_go.FaissIndexAdd(x, 10, s[:]))
	fmt.Println(faiss_go.FaissIndexNtotal(x))
}
