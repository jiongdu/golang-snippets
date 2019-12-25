package main

import (
	"fmt"
)

type Pet struct {
	name string
}

type Dog struct {
	Pet
	Breed string
}

func (p *Pet) Play() {
	fmt.Println(p.Speak())
}

func (p *Pet) Speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

func (p *Pet) Name() string {
	return p.name
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and I am a %v", d.Pet.Speak(), d.Breed)
}

func main() {
	d := Dog{Pet: Pet{name: "spot"}, Breed: "pointer"}
	fmt.Println(d.Name())
	fmt.Println(d.Speak())
	d.Play() //不能实现多态的效果
}
