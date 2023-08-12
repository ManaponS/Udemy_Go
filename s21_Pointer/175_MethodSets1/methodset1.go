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

func main() {
	d1 := dog{
		name: "Han",
	}
	d1.walk() //Han
	d1.run()  //Bolt

	d2 := &dog{
		name: "Drake",
	}
	d2.walk() //Drake
	d2.run()  //Bolt

}
