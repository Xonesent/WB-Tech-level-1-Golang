/*

№ 1
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

*/

package main

import "fmt"

type Human struct {
	Name       string
	Age        int
	Profession string
}

func (h *Human) Speak() {
	fmt.Printf("%s говорит: Привет!\n", h.Name)
}

type Action struct {
	Human
	ActionName string
}

func (a *Action) Act() {
	fmt.Printf("%s %s, ее профессия %s\n", a.Name, a.ActionName, a.Profession)
}

func main() {
	person := Action{
		Human: Human{
			Name:       "Анна",
			Age:        17,
			Profession: "БизнесВуман",
		},
		ActionName: "работает",
	}

	person.Speak()
	person.Act()
}
