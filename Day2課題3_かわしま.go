package main

import "fmt"

type Turtle struct {
	name string
	x    float64
	y    float64
	a    float64
}

func (t *Turtle) addition(a float64) {
	t.a += a
}

func (t *Turtle) change_name(name string) {
	t.name = name
}

func (t *Turtle) message(mes string) {
	fmt.Println(t.name + ":" + mes)
}

func main() {

	var t1 Turtle = Turtle{"師匠", 1000, 5, 180.5}
	var t2 Turtle = Turtle{"弟子", 10, 250, 270.3}
	fmt.Println(t1)
	fmt.Println(t2)

	t1.addition(10.5)
	t1.change_name("大師匠")
	fmt.Println(t1)
	t1.message("こんばんは")
}
