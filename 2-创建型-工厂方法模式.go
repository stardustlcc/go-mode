package main

import "fmt"

/*
   工厂模式：一个工厂，有多个车间（子工厂），通过实例化车间来创建产品对象
*/
//动物
type animal struct {
	name   string
	master string
}

func (a *animal) SetInfo(name, master string) {
	a.name = name
	a.master = master
}

//这里有三种小动物
type dog struct {
	animal
}
type cat struct {
	animal
}
type pig struct {
	animal
}

//动物都有Speck()方法
type Speaker interface {
	Speak()
	SetInfo(name, master string)
}

func (d *dog) Speak() {
	fmt.Printf("我叫%v, 我的主人是%v, 汪汪汪\n", d.name, d.master)
}
func (c *cat) Speak() {
	fmt.Printf("我叫%v, 我的主人是%v, 喵喵喵\n", c.name, c.master)
}
func (p *pig) Speak() {
	fmt.Printf("我叫%v, 我的主人是%v, 猪猪猪\n", p.name, p.master)
}

//抽象工厂接口
type AnimalFactory interface {
	GetAnimal() Speaker
}

//生产具体动物的具体工厂
type DogFactory struct{}
type CatFactory struct{}
type PigFactory struct{}

//狗工厂
func (d *DogFactory) GetAnimal() Speaker {
	return &dog{}
}

//猫工厂
func (c *CatFactory) GetAnimal() Speaker {
	return &cat{}
}

//猪工厂
func (p *PigFactory) GetAnimal() Speaker {
	return &pig{}
}

func main() {
	//抽象工厂
	var af AnimalFactory

	af = &DogFactory{}
	dog := af.GetAnimal()
	dog.SetInfo("哈士奇", "cici")
	dog.Speak()

	af = &CatFactory{}
	cat := af.GetAnimal()
	cat.SetInfo("咕噜", "mm")
	cat.Speak()

	af = &PigFactory{}
	pig := af.GetAnimal()
	pig.SetInfo("猪猪", "zuzu")
	pig.Speak()
}
