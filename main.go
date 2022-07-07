package main

import "fmt"

type Dog struct {
	wweight int
	ccount int
}
type Cat struct {
	wweight int
	ccount int
}
type Cow struct {
	wweight int
	ccount int
}

func(co Cow) count()int{return co.ccount}
func(c Cat) count()int{return c.ccount}
func(d Dog) count()int{return d.ccount}

func(co Cow) weight()int{return co.wweight}
func(c Cat) weight()int{return c.wweight}
func(d Dog) weight()int{return d.wweight}

func (Cow) String() string{return "Cow:"}
func (Cat) String() string{return "Cat:"}
func (Dog) String() string{return "Dog:"}


func (co Cow) food() int {
	return (co.wweight * 25)
}

func (c Cat) food() int {
	return (c.wweight * 7)
}

func (d Dog) food() int {
	x := (d.wweight / 5)
	return (x * 10)
}

type consumption interface {
	food() int
	String() string
	weight() int
	count() int
}

func about(x consumption, totaly *int) {
	fmt.Printf("Animal name %v with weight %vkg must eat %vkg food \n",x.String(),x.weight(),x.food())
	fmt.Printf("Total for %v = %vkg \n",x.String(),(x.count() * x.food()))
	*totaly=(*totaly+(x.count() * x.food()))
	fmt.Printf("Minimum food in farm: %vkg\n", *totaly)
}

func main() {
	var mustEat consumption
    var totaly int
	
	mustEat = Cat {wweight: 5, ccount: 3}
	about(mustEat, &totaly)
	
	mustEat = Dog {wweight: 15, ccount: 2}
	about(mustEat, &totaly)
    
    mustEat = Cow {wweight: 115, ccount: 25}
    about(mustEat, &totaly)
}
