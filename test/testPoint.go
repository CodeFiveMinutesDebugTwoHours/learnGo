package test

import "fmt"

func testPoint() {
	p := person{name: "zxf", age: 11}
	fmt.Println(p)

	setAge(&p, 22)
	fmt.Println(p)
	p.setAge(33)
	fmt.Println(p)
}

type person struct {
	name string
	age  int
}

func setAge(p *person, age int) {
	p.age = age
}

/*
*
可以改值
*/
func (p *person) setAgeWithPoint(age int) {
	p.age = age
}

/*
*
不能改值
*/
func (p person) setAge(age int) {
	p.age = age
}
