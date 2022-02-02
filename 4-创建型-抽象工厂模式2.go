package main

import "fmt"

//手机制作接口
type Phone interface {
	Call() //打电话
}

//ipad制作接口
type Ipad interface {
	Play() //打游戏
}

//工厂(制作手机，ipad)
type Factory interface {
	CreatePhone() Phone
	CreateIpad() Ipad
}

//小米公司
type Xiaomi struct{}

//华为公司
type Huawei struct{}

//小米手机模型
type XiaomiPhone struct{}

//小米iPadd模型
type XiaomiIpad struct{}

//华为手机模型
type HuaweiPhone struct{}
type HuaweiIpad struct{}

//小米手机打电话功能
func (x *XiaomiPhone) Call() {
	fmt.Println("小米，iphone打电话")
}

//小米Ipad玩游戏功能
func (x *XiaomiIpad) Play() {
	fmt.Println("小米，ipad玩游戏")
}

//华为手机打电话功能
func (h *HuaweiPhone) Call() {
	fmt.Println("华为，iphone打电话")
}

//华为ipad玩游戏功能
func (h *HuaweiIpad) Play() {
	fmt.Println("华为，ipad玩游戏")
}

//小米工厂实现
func (x *Xiaomi) CreatePhone() Phone {
	return &XiaomiPhone{}
}

func (x *Xiaomi) CreateIpad() Ipad {
	return &XiaomiIpad{}
}

//华为工厂实现
func (h *Huawei) CreatePhone() Phone {
	return &HuaweiPhone{}
}

func (h *Huawei) CreateIpad() Ipad {
	return &HuaweiIpad{}
}

func main() {
	xm := Xiaomi{}
	xm.CreatePhone().Call()
	xm.CreateIpad().Play()

	hw := Huawei{}
	hw.CreatePhone().Call()
	hw.CreateIpad().Play()
}
