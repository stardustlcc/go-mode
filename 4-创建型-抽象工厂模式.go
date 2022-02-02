package main

import "fmt"

type IProductA interface {
	runA()
}
type IProductB interface {
	runB()
}

//一个抽象工厂
type Ifactoty interface {
	createProductA() IProductA
	createProductB() IProductB
}

type ProductA1 struct{}
type ProductA2 struct{}

func (a *ProductA1) runA() {
	fmt.Println("run productA1 A")
}

func (a *ProductA2) runA() {
	fmt.Println("run productA2 A")
}

type ProductB1 struct{}
type ProductB2 struct{}

func (b *ProductB1) runB() {
	fmt.Println("run ProductB1 B")
}
func (b *ProductB2) runB() {
	fmt.Println("run ProductB2 B")
}

type Factory1 struct{}
type Factory2 struct{}

func (f *Factory1) createProductA() IProductA {
	return new(ProductA1)
}
func (f *Factory1) createProductB() IProductB {
	return new(ProductB1)
}

func (f *Factory2) createProductA() IProductA {
	return new(ProductA2)
}
func (f *Factory2) createProductB() IProductB {
	return new(ProductB2)
}

/*
   抽象工厂模式：将工厂抽象化，可以实例化多个厂家，每个厂家各自建造自己的产品
*/
func main() {
	var ifa Ifactoty
	var p1 IProductA
	var p2 IProductB

	ifa = new(Factory1)
	p1 = ifa.createProductA()
	p2 = ifa.createProductB()
	p1.runA()
	p2.runB()

	ifa = new(Factory2)
	p1 = ifa.createProductA()
	p2 = ifa.createProductB()
	p1.runA()
	p2.runB()
}
