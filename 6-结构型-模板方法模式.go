package main

import "fmt"

//动物接口
type Animal interface {
	Eat(food string) bool
	Sleep() bool
}

//抽象动物
type animal struct {
	name string
	food []string
	max  int
	full bool
}

func (a *animal) eat(food string, eat, full, refuse func() bool) bool {
	if !a.full {
		a.food = append(a.food, food)
		if len(a.food) < a.max {
			return eat() //留给具体的动物去实现
		} else {
			a.full = true
			return full() //留给具体的动物去实现
		}
	} else {
		return refuse() //留给具体的动物去实现
	}
}

func (a *animal) sleep(full, notFull func() bool) bool {
	if a.full {
		return full() //留给具体的动物去实现
	} else {
		return notFull() //留给具体的动物去实现
	}
}

//具体的动物
type Dog struct {
	*animal
}
type Pig struct {
	*animal
}

func NewDog(name string, max int) *Dog {
	p := &animal{
		name: name,
		max:  max,
	}
	return &Dog{p}
}

func NewPig(name string) *Pig {
	p := &animal{
		name: name,
	}
	return &Pig{p}
}

func (d *Dog) Eat(food string) bool {
	eat := func() bool {
		fmt.Printf("[%s]吃了[%v], 还要继续吃\n", d.name, food)
		return true
	}

	full := func() bool {
		fmt.Printf("[%v]吃了[%v]，已经吃饱了~\n", d.name, food)
		return true
	}

	refuse := func() bool {
		fmt.Printf("[%v]吃饱了，拒绝吃东西~\n", d.name)
		return false
	}
	return d.animal.eat(food, eat, full, refuse)
}

func (d *Dog) Sleep() bool {
	full := func() bool {
		fmt.Printf("[%v]吃饱了，去睡觉zzz\n", d.name)
		return true
	}
	notFull := func() bool {
		fmt.Printf("[%v]没吃饱，拒绝睡觉\n", d.name)
		return false
	}
	return d.animal.sleep(full, notFull)
}

func (p *Pig) Eat(food string) bool {
	eat := func() bool {
		fmt.Printf("[%v]吃了[%v]，还要继续吃~\n", p.name, food)
		return true
	}
	return p.animal.eat(food, eat, eat, eat)
}

func (p *Pig) Sleep() bool {
	full := func() bool {
		fmt.Printf("[%v]去睡觉zzz\n", p.name)
		return true
	}
	return p.animal.sleep(full, full)
}

func main() {
	var animal Animal
	animal = NewPig("小猪")
	fmt.Println(animal.Eat("饲料"))
	fmt.Println(animal.Sleep())

	animal = NewDog("小狗", 2)
	fmt.Println(animal.Eat("骨头"))
	fmt.Println(animal.Sleep())
	fmt.Println(animal.Eat("肉"))
	fmt.Println(animal.Eat("狗粮"))
	fmt.Println(animal.Sleep())
}
