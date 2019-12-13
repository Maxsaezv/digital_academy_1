package main

import "fmt"

type Pila struct {
	Pila []int
}

func (p *Pila) push(num int) {
	p.Pila = append(p.Pila, num)
}

func (p *Pila) pop() ([]int, int) {
	head := len(p.Pila) - 1

	element := p.Pila[head]

	p.Pila = p.Pila[:head]
	return p.Pila, element
}

func main() {

	var x Pila
	for index := 0; index < 10; index++ {
		x.push(index)
	}
	fmt.Println(x)
	fmt.Println(x.pop())

}
