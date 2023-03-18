package test

import "fmt"

func testType() {
	var a animal
	a = cat{"mao", "miao"}
	fmt.Println(a.description())
	var g = dog{"gou", "wang"}
	fmt.Println(g.description())
}

type animal interface {
	description() string
}

type cat struct {
	Type  string
	Sound string
}

type dog struct {
	Pinyin string
	sound  string
}

func (c cat) description() string {
	return fmt.Sprintf("sound: %v" + c.Sound)
}

func (c dog) description() string {
	return fmt.Sprintf("sound: %v" + c.sound)
}
