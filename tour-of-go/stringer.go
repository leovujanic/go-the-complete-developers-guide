package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Car struct {
	speed int
	brand string
}

func (p Car) String() string {
	return fmt.Sprintf("VW je auto, ovo su kola %s koja se vuku %d na sat", p.brand, p.speed)
}

func main() {
	alfa := Car{brand: "Alfa", speed: 100}
	mercedes := Car{brand: "Mercedes", speed: 160}

	fmt.Println(alfa)
	fmt.Println(mercedes)
}
