package main

import (
	"bufio"
	"fmt"
	"github.com/artslob/gophercises/09-deck/deck"
	"github.com/artslob/gophercises/10-blackjack/hand"
	"log"
	"os"
	"strings"
)

func stringifyCard(card deck.Card) string {
	return fmt.Sprintf("{%s}", card)
}

func stringifyHand(h hand.Hand) string {
	var array []string
	for _, card := range *h.Cards {
		array = append(array, stringifyCard(card))
	}
	return strings.Join(array, ", ")
}

func main() {
	fmt.Println("Starting blackjack game.")
	d := deck.New()
	d.Shuffle()
	player := hand.New(d.GetTopCard(), d.GetTopCard())
	dealer := hand.New(d.GetTopCard(), d.GetTopCard())
	fmt.Printf("Dealer have: %s and *hidden* card.\n", stringifyCard(dealer.TopCard()))

	scanner := bufio.NewScanner(os.Stdin)
F:
	for {
		fmt.Printf("\nYou have: %s.\n", stringifyHand(player))
		fmt.Printf("Score: %s\n", player.ScoreString())
		if player.IsBlackjack() {
			fmt.Println(playerWon)
			return
		}
		if player.MoreThanBlackjack() {
			fmt.Println("Sorry, you lost.")
			return
		}

		fmt.Print("Will you Stand or Hit? (h) or (s): ")
		if !scanner.Scan() && scanner.Err() != nil {
			log.Fatal("got error while reading input")
		}
		input := strings.TrimSpace(scanner.Text())
		switch input {
		case "h":
			player.Draw(d.GetTopCard())
			fmt.Printf("You got %s\n", stringifyCard(player.TopCard()))
		case "s":
			break F
		default:
			fmt.Println("Enter 'h' or 's'.")
		}
	}

	fmt.Printf("\nYour`s score: %s\n", player.ScoreString())

	fmt.Printf("Dealers turn. His hand: %s, score: %s\n", stringifyHand(dealer), dealer.ScoreString())
	if dealer.Score() <= 16 || dealer.SoftScore() == 17 {
		dealer.Draw(d.GetTopCard())
		fmt.Printf("Dealer draws and gets the card: %s, score: %s\n", dealer.TopCard(), dealer.ScoreString())
	}

	fmt.Println(whoWon(player.BestScore(), dealer.BestScore()))
}

type WinnerPerson string

const (
	playerWon WinnerPerson = "You won!"
	dealerWon              = "Dealer won."
	draw                   = "Draw."
)

func whoWon(playerScore, dealerScore int) WinnerPerson {
	if playerScore == dealerScore {
		return draw
	}
	if playerScore > hand.Blackjack && dealerScore > hand.Blackjack {
		if playerScore < dealerScore {
			return playerWon
		}
		return dealerWon
	}
	if playerScore <= hand.Blackjack && dealerScore <= hand.Blackjack {
		if playerScore > dealerScore {
			return playerWon
		}
		return dealerWon
	}
	if playerScore <= hand.Blackjack {
		return playerWon
	}
	return dealerWon
}
