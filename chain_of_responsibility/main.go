package main

import "fmt"

func main(){

	h1 := &Handler{Action: &SymbolHandler{},}
	h2 := &Handler{Action: &CharacterHandler{},}
	h3 := &Handler{Action: &DigitHandler{},}
    

	h1.SetNext(h2)
	h2.SetNext(h3)

	var a Action = h1
	
	fmt.Println("---- 1234 ----")
	a.HandlerAction("1234")
	fmt.Println("---- victor ----")
	a.HandlerAction("victor")

}