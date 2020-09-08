package main

import (
	"unicode"
	"fmt"
)


type Action interface{
	HandlerAction(c string)
}

type Handler struct {
	Action
	next *Handler
}

func (h *Handler) SetNext(next *Handler) {
	h.next = next
}

func (h *Handler) HandlerAction(c string){
	// this round
	if h.Action != nil {
		h.Action.HandlerAction(c)
	}
	// next part
	if h.next != nil {
		h.next.HandlerAction(c)
	}
}

type SymbolHandler struct{}

func (sh *SymbolHandler) HandlerAction(c string){
	fmt.Println("Is Symbol")
}

type CharacterHandler struct{}

func onlyLetter(c string) bool{
	for _, r := range c {
        if !unicode.IsLetter(r) {
            return false
        }
	}
	return true
}

func (ch *CharacterHandler) HandlerAction(c string){
	if onlyLetter(c){
		fmt.Println("Is Character")
	}
}

type DigitHandler struct{}

func onlyDigit(c string) bool{
	for _, r := range c {
        if !unicode.IsDigit(r) {
            return false
        }
	}
	return true
}

func (dh *DigitHandler) HandlerAction(c string){
	if onlyDigit(c){
        fmt.Println("Is Number")
	}
}
