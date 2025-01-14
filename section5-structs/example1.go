package main

import "fmt"

type user struct {
	email string
}

type admin struct {
	user
	password string
}

type adminExtended admin

func (u *user) toString() string {
	return fmt.Sprintf("user{email: %s}", u.email)
}

func (a *adminExtended) toString() string {
	return fmt.Sprintf("admin{%s, password: %s}", a.user.toString(), a.password)
}

func main() {
	u := user{email: "user@email.com"}
	fmt.Println(u.toString())
	a := admin{user{email: "admin@mail.com"}, "qwerty"}
	fmt.Println(a.toString())
	aExt := adminExtended{user{email: "admin@mail.com"}, "qwerty"}
	fmt.Println(aExt.toString())
}
