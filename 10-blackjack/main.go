package main

import (
	"bufio"
	"fmt"
	"github.com/artslob/gophercises/09-deck/deck"
	"log"
	"os"
	"strings"
)

func stringifyCard(card deck.Card) string {
	return fmt.Sprintf("{%s}", card)
}

func stringifyHand(h hand) string {
	var array []string
	for _, card := range h {
		array = append(array, stringifyCard(card))
	}
	return strings.Join(array, ", ")
}

const blackjack = 21

func maxCloseToBlackjack(normal, soft int) int {
	if normal == soft {
		return normal
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	if normal <= blackjack && soft <= blackjack {
		return max(normal, soft)
	} else if normal <= blackjack {
		return normal
	} else if soft <= blackjack {
		return soft
	}
	return min(normal, soft)
}

func main() {
	fmt.Println("Starting blackjack game.")
	player, dealer := hand{}, hand{}
	d := deck.New()
	d.Shuffle()
	for i := 0; i < 2; i++ {
		player.draw(d.GetTopCard())
		dealer.draw(d.GetTopCard())
	}
	fmt.Printf("Dealer have: %s and *hidden* card.\n", stringifyCard(dealer[0]))

	scanner := bufio.NewScanner(os.Stdin)
F:
	for {
		fmt.Printf("\nYou have: %s.\n", stringifyHand(player))
		fmt.Printf("Score: %s\n", player.scoreString())
		fmt.Print("Will you Stand or Hit? (h) or (s): ")
		if !scanner.Scan() && scanner.Err() != nil {
			log.Fatal("got error while reading input")
		}
		input := strings.TrimSpace(scanner.Text())
		switch input {
		case "h":
			top := d.GetTopCard()
			player.draw(top)
			fmt.Printf("You got %s\n", stringifyCard(top))
			normal, soft := player.getScores()
			fmt.Printf("Score: %s\n", player.stringifyScores(normal, soft))
			if soft == blackjack || normal == blackjack {
				fmt.Println("You won!")
				return
			}
			if soft > blackjack && normal > blackjack {
				fmt.Println("Sorry, you lost.")
				return
			}
		case "s":
			break F
		default:
			fmt.Println("Enter 'h' or 's'.")
		}
	}

	fmt.Printf("\nYour`s score: %s\n", player.scoreString())

	dNormal, dSoft := dealer.getScores()
	fmt.Printf("Dealers turn. His hand: %s, score: %s\n", stringifyHand(dealer), dealer.stringifyScores(dNormal, dSoft))
	if dNormal <= 16 || dSoft == 17 {
		top := d.GetTopCard()
		dealer.draw(top)
		fmt.Printf("Dealer draws and gets the card: %s, score: %s\n", top, dealer.scoreString())
	}

	dealerScore := maxCloseToBlackjack(dealer.getScores())
	playerScore := maxCloseToBlackjack(player.getScores())

	fmt.Println(whoWon(playerScore, dealerScore))
}

func whoWon(playerScore, dealerScore int) string {
	playerWon := "You won!"
	dealerWon := "Dealer won."

	if playerScore == dealerScore {
		return "Draw."
	}
	if playerScore > blackjack && dealerScore > blackjack {
		if playerScore < dealerScore {
			return playerWon
		}
		return dealerWon
	}
	if playerScore <= blackjack && dealerScore <= blackjack {
		if playerScore > dealerScore {
			return playerWon
		}
		return dealerWon
	}
	if playerScore <= blackjack {
		return playerWon
	}
	return dealerWon
}
