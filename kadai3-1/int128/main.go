package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gopherdojo/dojo2/kadai3-1/int128/game"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	g := game.New(game.Things, 30*time.Second)
	fmt.Fprintf(os.Stderr, "Are you ready? Type words in 30 seconds!\n--\n")

	ctx := context.Background()
	score, err := g.Start(ctx)
	switch {
	case err != nil:
		log.Fatal(err)
	case score.GiveUp:
		fmt.Fprintf(os.Stderr, "\nGive up? You got %d point(s).\n", score.CorrectWords)
	default:
		fmt.Fprintf(os.Stderr, "\nTime up! You got %d point(s).\n", score.CorrectWords)
	}
}
