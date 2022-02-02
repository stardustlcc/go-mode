package main

import (
	"fmt"
	"sync"
)

//单例模式类
type Singletion struct{}

var singletion *Singletion
var once sync.Once

//获取单例模式对象
func GetInstance() *Singletion {
	//用加锁原子操作（代码包sync/atomic）来保证函数 f 只执行一次。
	once.Do(func() {
		singletion = &Singletion{}
		fmt.Println("instance...")
	})
	return singletion
}

func main() {
	var s *Singletion

	s = GetInstance()
	s = GetInstance()
	fmt.Println(s)

}
