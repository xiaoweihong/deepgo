package main

type User2 struct {
	Name string
}

func (u User2) test() {

}

func main() {
	abc := User2{"www"}
	abc.test()
}
