package main

import "fmt"

type notifer interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func sendNotification(n notifer) {
	n.notify()
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	lisa := admin{"lisa", "lisa@email.com"}
	sendNotification(&lisa)
}
