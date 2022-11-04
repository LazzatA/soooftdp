package main

import "obs/user"

func main() {
	sub1 := &user.Subscriber{Name: "Sub"}
	sub2 := &user.Subscriber{Name: "Sub2"}
	sub3 := &user.Subscriber{Name: "Sub3"}
	subN := &user.Subscriber{Name: "Sub4"}

	chanel := user.Publisher{
		Name:  "Publisher channel",
		Users: map[string]user.Users{},
	}
	chanel.Subscribe(sub1)
	chanel.Subscribe(sub2)
	chanel.Subscribe(sub3)
	chanel.Subscribe(subN)
	println("Unsubscribe Sub-3")
	chanel.UnSubscribe(sub3)
	chanel.Notify
}
