package main

import "fmt"


type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

// 定义另一个notifier
type admin struct {

}
func (a admin) notify() {
	fmt.Println("notify admin")
}

//定义user结构体
type user struct {
	//定义user的域
	name string
	email string
}

//值接收
func (u user) notify() {
	fmt.Println("notify from user")
}
func (u user) printName() {
	fmt.Println(u.name)
}
//指针接收
func (u *user) changeName(newName string) {
	u.name = newName
}


func main() {
	u := user{
		name :"xnwuig",
		email : "xnwuig@netease.com",
	}
	u.printName()
	u.changeName("xiaowucheng")
	u.printName()
	u.notify()
	sendNotification(u)
	a := admin{}
	sendNotification(a)
}