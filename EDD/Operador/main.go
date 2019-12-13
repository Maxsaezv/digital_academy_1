package main

import (
	"fmt"
)

//Product producto
type Product struct {
	Name  string
	Price float32
}

func add(m map[int32]*Product, p Product) int {
	m[0] = &p
	return 1

}

func main() {
	list := make(map[int32]*Product)

	x := Product{Name: "ols", Price: 1.0}

	add(list, x)

	fmt.Printf("%+v\n", *list[0])

}
