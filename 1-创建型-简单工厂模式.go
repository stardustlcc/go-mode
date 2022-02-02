package main

import "fmt"

/*
	go 没有构造函数一说，一般会定义一个NewXXX函数来初始化相关函数
	NewXXX函数返回接口时就是简单工厂模式，golang 一般推荐的做法就是简单工厂模式
	使用场景：
		代码中出现switch的时候就可以考虑使用简单工厂模式

	简单的工厂模式：利用参数来创建不同的产品
*/
type API interface {
	Say(name string) string
}

type userApi struct{}

type adminApi struct{}

func NewApi(t int) API {
	switch t {
	case 1:
		return &userApi{}
	case 2:
		return &adminApi{}
	}
	return nil
}

func (*userApi) Say(name string) string {
	return fmt.Sprintf("hai user %s", name)
}

func (*adminApi) Say(name string) string {
	return fmt.Sprintf("hai admin %s", name)
}

func main() {
	fmt.Println(NewApi(1).Say("mm"))
	fmt.Println(NewApi(2).Say("gg"))
}
