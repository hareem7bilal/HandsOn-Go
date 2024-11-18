package main

import (
	"fmt"

	// import the proverbs package
	"github.com/jboursiquot/go-proverbs"
)

// getRandomProverb returns a random proverb from the proverbs package
func getRandomProverb() *proverbs.Proverb {
	return proverbs.Random()
}

func main() {
	// print the result of calling your getRandomProverb function
	fmt.Printf("%+v\n", getRandomProverb())
}
