package main

import "testing"

func TestWhoWon(t *testing.T) {
	const playerWon = "You won!"
	const dealerWon = "Dealer won."
	const draw = "Draw."

	type args struct {
		playerScore int
		dealerScore int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "zeros",
			args: args{
				playerScore: 0,
				dealerScore: 0,
			},
			want: draw,
		},
		{
			name: "equals to less than 21",
			args: args{
				playerScore: 19,
				dealerScore: 19,
			},
			want: draw,
		},
		{
			name: "equals to 21",
			args: args{
				playerScore: 21,
				dealerScore: 21,
			},
			want: draw,
		},
		{
			name: "equals to 22",
			args: args{
				playerScore: 22,
				dealerScore: 22,
			},
			want: draw,
		},
		{
			name: "both more than 21",
			args: args{
				playerScore: 22,
				dealerScore: 23,
			},
			want: playerWon,
		},
		{
			name: "player 21 and dealer more than 21",
			args: args{
				playerScore: 21,
				dealerScore: 25,
			},
			want: playerWon,
		},
		{
			name: "dealer 21 and player more than 21",
			args: args{
				playerScore: 22,
				dealerScore: 21,
			},
			want: dealerWon,
		},
		{
			name: "both less than 21",
			args: args{
				playerScore: 19,
				dealerScore: 17,
			},
			want: playerWon,
		},
		{
			name: "player less than 21 and dealer more than 21",
			args: args{
				playerScore: 17,
				dealerScore: 22,
			},
			want: playerWon,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := whoWon(tt.args.playerScore, tt.args.dealerScore); got != tt.want {
				t.Errorf("whoWon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxCloseToBlackjack(t *testing.T) {
	type args struct {
		normal int
		soft   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zeroes",
			args: args{
				normal: 0,
				soft:   0,
			},
			want: 0,
		},
		{
			name: "blackjack",
			args: args{
				normal: 21,
				soft:   21,
			},
			want: 21,
		},
		{
			name: "21 and more than 21",
			args: args{
				normal: 21,
				soft:   22,
			},
			want: 21,
		},
		{
			name: "two more than 21",
			args: args{
				normal: 24,
				soft:   22,
			},
			want: 22,
		},
		{
			name: "two less than 21",
			args: args{
				normal: 19,
				soft:   16,
			},
			want: 19,
		},
		{
			name: "21 and less than 21",
			args: args{
				normal: 21,
				soft:   20,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCloseToBlackjack(tt.args.normal, tt.args.soft); got != tt.want {
				t.Errorf("maxCloseToBlackjack() = %v, want %v", got, tt.want)
			}
		})
	}
}
