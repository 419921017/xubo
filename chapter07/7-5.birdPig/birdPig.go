package main

import "fmt"

func main() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	for name, obj := range animals {
		f, isFlyer := obj.(Flyer)
		w, isWaker := obj.(Walker)
		fmt.Printf("name: %s is Flyer: %v is Walker: %v\n", name, isFlyer, isWaker)
		if isFlyer {
			f.Fly()
		}
		if isWaker {
			w.Walk()
		}
	}

	p1 := new(pig)
	var a Walker = p1
	p2 := a.(*pig)
	fmt.Printf("p1=%p, p2=%p", p1, p2)
}

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct{}

func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

type pig struct{}

func (p *pig) Walk() {
	fmt.Println("pig: walk")
}
