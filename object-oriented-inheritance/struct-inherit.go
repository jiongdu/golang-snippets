package main

import "fmt"

type Pet struct {
	name string
}

// 继承自其他结构体的struct类型可以直接访问父类结构体的字段和方法
type Dog struct {
	Pet   // 匿名嵌入结构体实例 or 匿名嵌入结构体实例指针
	Breed string
}

func (p *Pet) Speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

func (p *Pet) Name() string {
	return p.name
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and i am a %v", d.Pet.Speak(), d.Breed)
}

func main() {
	d := Dog{
		Pet:   Pet{name: "spot"},
		Breed: "pointer",
	}
	fmt.Println(d.Name())
	fmt.Println(d.Speak())
}
