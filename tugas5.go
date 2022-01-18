package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n1, n2, n3 pelajar
	n1.nama = "Ahmad"
	n1.umur = rand.Intn(100)
	n2.nama = "Abi"
	n2.umur = rand.Intn(100)
	n3.nama = "Zaka"
	n3.umur = rand.Intn(100)
	n1.namasaya()
	n2.namasaya()
	n3.namasaya()

}

type pelajar struct {
	nama string
	umur int
}

func (n pelajar) namasaya() {
	fmt.Println(n.nama, n.umur)
}
