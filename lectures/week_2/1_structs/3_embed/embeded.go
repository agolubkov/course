package main

import "fmt"

type Person struct {
	Name string
	inn  string
}

type Stuff struct {
	inn int
}

type SecretAgent struct {
	Person
	Stuff
	LicenseToKill bool
}

func (p Person) GetName() string {
	return p.Name
}

func (s SecretAgent) GetName() string {
	return "CLASSIFIED"
}

func (s SecretAgent) GetInn() string {
	return fmt.Sprintf("%03d", s.Stuff.inn)
}

func (s SecretAgent) String() string {
	return "agent: " + s.GetInn() + " name: " + s.GetName()
}

func main() {
	sa := SecretAgent{Person: Person{"James", "12312321321"}, Stuff: Stuff{inn: 7}, LicenseToKill: true}

	fmt.Printf("info: %T %+v\n", sa, sa)
}
