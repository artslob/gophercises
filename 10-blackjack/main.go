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
	player, dealer := hand.Hand{}, hand.Hand{}
	d := deck.New()
	d.Shuffle()
	for i := 0; i < 2; i++ {
		player.Draw(d.GetTopCard())
		dealer.Draw(d.GetTopCard())
	}
	fmt.Printf("Dealer have: %s and *hidden* card.\n", stringifyCard((*dealer.Cards)[0]))

	scanner := bufio.NewScanner(os.Stdin)
F:
	for {
		fmt.Printf("\nYou have: %s.\n", stringifyHand(player))
		fmt.Printf("Score: %s\n", player.ScoreString())
		fmt.Print("Will you Stand or Hit? (h) or (s): ")
		if !scanner.Scan() && scanner.Err() != nil {
			log.Fatal("got error while reading input")
		}
		input := strings.TrimSpace(scanner.Text())
		switch input {
		case "h":
			top := d.GetTopCard()
			player.Draw(top)
			fmt.Printf("You got %s\n", stringifyCard(top))
			fmt.Printf("Score: %s\n", player.ScoreString())
			normal, soft := player.GetScores()
			if soft == hand.Blackjack || normal == hand.Blackjack {
				fmt.Println("You won!")
				return
			}
			if soft > hand.Blackjack && normal > hand.Blackjack {
				fmt.Println("Sorry, you lost.")
				return
			}
		case "s":
			break F
		default:
			fmt.Println("Enter 'h' or 's'.")
		}
	}

	fmt.Printf("\nYour`s score: %s\n", player.ScoreString())

	fmt.Printf("Dealers turn. His hand: %s, score: %s\n", stringifyHand(dealer), dealer.ScoreString())
	dNormal, dSoft := dealer.GetScores()
	if dNormal <= 16 || dSoft == 17 {
		top := d.GetTopCard()
		dealer.Draw(top)
		fmt.Printf("Dealer draws and gets the card: %s, score: %s\n", top, dealer.ScoreString())
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
