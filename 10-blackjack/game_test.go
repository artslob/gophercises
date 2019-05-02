package main

import "testing"

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
