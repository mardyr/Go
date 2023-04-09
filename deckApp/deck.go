package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	"os"
)

type deck [] string 

func newDeck() deck{
	cards := deck{}

	cardSuits := []string{"Spades","Hearts","Diamonds","Clubs"}
	cardValues := [] string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","Oueen","King"}

	for _,suit := range cardSuits{
		for _,value := range cardValues{
			cards = append(cards,value +" of " + suit)
		}
	}
	return cards
}

func (d deck) print(){	
	for i,card := range d{
		fmt.Println(i,card)
	}
}

func deal(d deck,handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}

func (d deck) toString() string{
	return strings.Join([]string(d),"\n")
}

func (d deck) saveToFile(path string) error{
	return ioutil.WriteFile(path,[]byte(d.toString()),0777)
}

func newDeckFromFile(path string) deck{
	bs, err := ioutil.ReadFile(path)
	if err != nil{
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs),"\n"))
}

func (d deck) shuffle() {
	shuffleSeededSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(shuffleSeededSource)

	for i := range d {
		newPosition := r.Intn(len(d)-1)
		//swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}