package test

import "fmt"

func test() {

	i := 1
	var j = 0
	//flag := true
	var d float32 = 2.3333
	var c int = int(d)
	fmt.Println("test")
	fmt.Println(i + j + c)
	num := 11
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	q := 2
	switch q {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("none")
	}

	for i := 0; i < 4; i++ {
		fmt.Print(i)
	}

	k := 2
	for k < 3 {
		fmt.Print("1")
		k++
	}

	ap := &k
	//var bp *int
	fmt.Println(*ap)
	fmt.Println(&k)

	z := 12
	var zp = &z
	fmt.Println(*zp)
	print(ap)

	zero := 9
	inc(&zero)
	fmt.Println(zero)
}

func print(i *int) {
	fmt.Println(*i)
}
func inc(i *int) {
	*i++
}
