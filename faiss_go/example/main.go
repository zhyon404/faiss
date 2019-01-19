package main

import (
	"fmt"
	"github.com/facebookresearch/faiss/faiss_go"
	"math/rand"
)

func main() {
	Searching()
}

func add(x *faiss_go.Faissindex, nb int, xb []float32) {
	fmt.Println(faiss_go.FaissIndexIsTrained(x))
	fmt.Println(faiss_go.FaissIndexAdd(x, nb, xb))
	fmt.Println(faiss_go.FaissIndexNtotal(x))
}

func Searching() {
	d, nb, xb, nq, xq := data()
	i := index(d)
	_, _ = nq, xq
	add(i, nb, xb)
	searching(i, nq, xq, 4)
}

func searching(i *faiss_go.Faissindex, n int, x []float32, k int) {
	I := make([]int, k*n)
	D := make([]float32, k*n)
	fmt.Println(faiss_go.FaissIndexSearch(i, n, x, k, D, I))
}

func index(d int) *faiss_go.Faissindex {
	v := new(faiss_go.Faissindexflatl2)
	fmt.Println(faiss_go.FaissIndexflatl2NewWith(&v, d))
	x := (*faiss_go.Faissindex)(v)
	return x
}

func data() (int, int, []float32, int, []float32) {
	d := 64      // dimension
	nb := 100000 // database size
	nq := 10000  // nb of queries
	xb := make([]float32, d*nb)
	xq := make([]float32, d*nq)
	for i := 0; i < nb; i++ {
		for j := 0; j < d; j++ {
			v := rand.Float32()
			xb[d*i+j] = v - float32(4)
		}
		xb[d*i] += float32(i) / 1000.
	}
	for i := 0; i < nq; i++ {
		for j := 0; j < d; j++ {
			xq[d*i+j] = rand.Float32()
		}
		xq[d*i] += float32(i) / 1000.
	}
	return d, nb, xb, nq, xq
}
