package main

import "fmt"

type carro struct{
	Id int
}

func (c carro) EmaiorDeDez()bool {
	return c.Id >=10
}

func main(){
	car := carro{}
	car.Id = 100
	if car.EmaiorDeDez() {
		fmt.Println("eh maior")
	}
}