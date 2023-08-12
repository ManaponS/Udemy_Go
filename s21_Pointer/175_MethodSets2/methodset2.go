package main

import "fmt"

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Printf("%v is walking \n", d.name)
}

func (d *dog) run() {
	d.name = "Bolt"
	fmt.Printf("%v is running \n", d.name)
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{
		name: "Han",
	}
	d1.walk() //Han
	d1.run()  //Bolt
	// cannot use d1 (variable of type dog) as youngin value in argument to youngRun: dog does not implement youngin (method run has pointer receiver)
	// youngRun(d1)

	d2 := &dog{
		name: "Drake",
	}
	d2.run()  //Bolt
	d2.walk() //Bolt
	youngRun(d2)

}
