package day22

import (
	"fmt"
)

type Game struct {
	Deck1 *Deck
	Deck2 *Deck
	PlayedConfigs map[string]struct{}
	SubgamesEnabled bool
	IsSubgame bool
}

type Deck struct {
	TopCard *Card
	BottomCard *Card
	CardCount int
}


type Card struct {
	Above *Card
	Below *Card
	Value int
}

func newDeck(cards []int) *Deck {
	deck := &Deck{
		TopCard: &Card{
			Above: nil,
			Below: nil,
			Value: cards[0],
		},
		CardCount: len(cards),
	}
	deck.BottomCard = deck.TopCard
	for _, cardValue := range cards[1:] {
		card := &Card{
			Above: nil,
			Below: nil,
			Value: cardValue,
		}
		deck.BottomCard.Below = card
		card.Above = deck.BottomCard
		deck.BottomCard = card
	}
	return deck
}


func NewGame(player1 []int, player2 []int, subgamesEnabled, isSubgame bool) *Game {
	g := &Game{
		Deck1: newDeck(player1),
		Deck2: newDeck(player2),
		PlayedConfigs: map[string]struct{}{},
		SubgamesEnabled: subgamesEnabled,
		IsSubgame: isSubgame,

	}
	return g
}

func (d *Deck) removeTopCard() *Card {
	res := d.TopCard
	if res == d.BottomCard {
		d.BottomCard = nil
	}
	d.TopCard = d.TopCard.Below
	if d.TopCard != nil {
		d.TopCard.Above = nil
	}
	res.Above = nil
	res.Below = nil
	d.CardCount--
	return res
}

func (d *Deck) print() {
	card := d.TopCard
	for ; card != nil; {
		fmt.Printf("%v <- %d -> %v \n", card.Above, card.Value, card.Below)
		card = card.Below
	}
}

func (d *Deck) getCardValueSlice(limit int) []int {
	res := []int{}
	for card := d.TopCard; card != nil; {
		res = append(res, card.Value)
		card = card.Below
	}
	return res[:limit]
}

func (d *Deck) addCardsToBottom(cards ...*Card) {
	for _, card := range cards {
		if d.BottomCard != nil {
			d.BottomCard.Below = card
			card.Above = d.BottomCard
			d.BottomCard = card
			card.Below = nil
		} else {
			d.TopCard = card
			d.BottomCard = card
			card.Below = nil
			card.Above = nil
		}
		d.CardCount++
	}
}


func (g *Game) PlayRound() (bool, int) {
	var winner int
	config := g.getCardsConfig()
	if g.SubgamesEnabled {
		if _, ok := g.PlayedConfigs[config]; ok {
			fmt.Println("Player 1 has won the game by infinite recursion protection rule")
			return true, 1
		} else {
			g.PlayedConfigs[config] = struct{}{}
		}
	}
	if g.Deck1.TopCard == nil || g.Deck2.TopCard == nil {
		if g.Deck1.TopCard == nil {
			fmt.Println("Player 2 has won the game")
			winner = 2
		}
		if g.Deck2.TopCard == nil {
			fmt.Println("Player 1 has won the game")
			winner = 1
		}
		return true, winner
	}

	card1, card2 := g.Deck1.removeTopCard(), g.Deck2.removeTopCard()

	if g.SubgamesEnabled && g.Deck1.CardCount >= card1.Value && g.Deck2.CardCount >= card2.Value {
		subgame := NewGame(g.Deck1.getCardValueSlice(card1.Value), g.Deck2.getCardValueSlice(card2.Value), true, true)
		var roundWinner int
		var gg bool
		for ;gg != true; {
			gg, roundWinner = subgame.PlayRound()
		}
		if roundWinner == 1 {
			g.Deck1.addCardsToBottom(card1, card2)
		} else {
			g.Deck2.addCardsToBottom(card2, card1)
		}
	} else {
		if card1.Value > card2.Value {
			g.Deck1.addCardsToBottom(card1, card2)
		} else {
			g.Deck2.addCardsToBottom(card2, card1)
		}
	}

	return false, winner
}

func (d *Deck) getCardsConfig() string {
	res := ""
	for card := d.TopCard; card != nil; {
		res += fmt.Sprintf("_%d_", card.Value)
		card = card.Below
	}
	return res
}

func (g *Game) getCardsConfig() string {
	res := fmt.Sprintf("%s__%s", g.Deck1.getCardsConfig(), g.Deck2.getCardsConfig())
	return res
}


func (g *Game) getScore(player int) int {
	var score int
	var card *Card
	var deck *Deck
	if player == 1 {
		deck = g.Deck1
	}
	if player == 2 {
		deck = g.Deck2
	}
	card = deck.BottomCard
	for i := 1; card != nil; i++ {
		score += card.Value * i
		if card == deck.TopCard {
			break
		}
		card = card.Above
	}
	return score
}