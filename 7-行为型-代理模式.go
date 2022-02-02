package main

import "fmt"

/*
	https://www.cnblogs.com/loveshes/p/12734990.html
	代理模式，为其它对象提供一种代理以控制对这个对象的访问。
	比如小明向小红写情书，但是又不好意思直接给小红，于是就让小军替自己给小红。
*/

type Girl struct {
	Name string
}

//发送消息的接口
type SendMsger interface {
	SenMsg(*Girl)
}

//真实对象
type TrueLove struct {
	Name string
}

//真实对象的发送动作
func (p *TrueLove) SendMsg(g Girl) {
	fmt.Printf("[%s] 这是[%s]给你的情书\n", g.Name, p.Name)
}

//代理对象
type Proxy struct {
	Name string
	TrueLove
}

//代理对象对真实对象进行包装
func (p *Proxy) Wrap(t TrueLove) {
	p.TrueLove = t
}

//代理对象的发送动作
func (p *Proxy) SendMsg(g Girl) {
	p.TrueLove.SendMsg(g)
	fmt.Printf("-------来自[%v]\n", p.Name)
}

func main() {
	girl := Girl{"小红"}
	trueLove := TrueLove{"小明"}
	trueLove.SendMsg(girl)
	fmt.Println()

	proxy := Proxy{Name: "工具人小刘"}
	proxy.Wrap(trueLove)
	proxy.SendMsg(girl)

}
