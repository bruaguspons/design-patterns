package main

import "errors"

type User interface {
	GetName() string
	GetAge() int
	GetRole() string
}

type AdminUser struct {
	Name string
	Age  int
}

func (a AdminUser) GetName() string {
	return a.Name
}

func (a AdminUser) GetAge() int {
	return a.Age
}

func (a AdminUser) GetRole() string {
	return "admin"
}

type GuestUser struct {
	Name string
	Age  int
}

func (g GuestUser) GetName() string {
	return g.Name
}

func (g GuestUser) GetAge() int {
	return g.Age
}

func (g GuestUser) GetRole() string {
	return "guest"
}

type RegularUser struct {
	Name string
	Age  int
}

func (r RegularUser) GetName() string {
	return r.Name
}

func (r RegularUser) GetAge() int {
	return r.Age
}

func (r RegularUser) GetRole() string {
	return "regular"
}

func CreateUser(userType string, name string, age int) (User, error) {
	switch userType {
	case "admin":
		return AdminUser{Name: name, Age: age}, nil
	case "guest":
		return GuestUser{Name: name, Age: age}, nil
	case "regular":
		return RegularUser{Name: name, Age: age}, nil
	default:
		return nil, errors.New("invalid user type")
	}
}
