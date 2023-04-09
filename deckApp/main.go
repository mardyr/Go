package main


func main() {
	filename := "deck"
	cards := newDeckFromFile(filename)
	hand,_ := deal(cards,3)
	cards.shuffle();
	cards.print()
}
