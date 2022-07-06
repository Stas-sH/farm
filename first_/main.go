package main

import "fmt"

type Dog struct {
	weight int
	count int
}
type Cat struct {
	weight int
	count int
}
type Cow struct {
	weight int
	count int
}

func(co Cow) Count()int{return co.count}
func(c Cat) Count()int{return c.count}
func(d Dog) Count()int{return d.count}

func(co Cow) Weight()int{return co.weight}
func(c Cat) Weight()int{return c.weight}
func(d Dog) Weight()int{return d.weight}

func (Cow) Name() string{return "Cow:"}
func (Cat) Name() string{return "Cat:"}
func (Dog) Name() string{return "Dog:"}


func (co Cow) Food() int {
	return (co.weight * 25)
}

func (c Cat) Food() int {
	return (c.weight * 7)
}

func (d Dog) Food() int {
	x := (d.weight / 5)
	return (x * 10)
}

type consumption interface {
	Food() int
	Name() string
	Weight() int
	Count() int
}

func about(x consumption, totaly *int) {
	fmt.Printf("Animal name %v with weight %vkg must eat %vkg food \n",x.Name(),x.Weight(),x.Food())
	fmt.Printf("Total for %v = %vkg \n",x.Name(),(x.Count() * x.Food()))
	*totaly=(*totaly+(x.Count() * x.Food()))
	fmt.Printf("Minimum food in farm: %vkg", *totaly)
}

func main() {
	var mustEat consumption
    var totaly int
	
	mustEat = Cat {weight: 5, count: 3}
	about(mustEat, &totaly)
	
	mustEat = Dog {weight: 15, count: 2}
	about(mustEat, &totaly)
    
    mustEat = Cow {weight: 115, count: 25}
    about(mustEat, &totaly)
}
