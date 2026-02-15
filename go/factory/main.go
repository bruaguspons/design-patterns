package main

import "fmt"

func main() {
	admin, _ := CreateUser("admin", "John", 30)
	guest, _ := CreateUser("guest", "Jane", 25)
	regular, _ := CreateUser("regular", "Bob", 40)

	fmt.Printf("%s (%d) - %s\n", admin.GetName(), admin.GetAge(), admin.GetRole())
	fmt.Printf("%s (%d) - %s\n", guest.GetName(), guest.GetAge(), guest.GetRole())
	fmt.Printf("%s (%d) - %s\n", regular.GetName(), regular.GetAge(), regular.GetRole())
}
