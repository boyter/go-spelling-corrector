package main

import (
	"fmt"
	"github.com/boyter/go-spelling-corrector/checker"
)

func main() {
	spellchecker := checker.NewSpellChecker()
	spellchecker.Train("spelling")

	fmt.Println(spellchecker.Correct("spuslling"))

	fmt.Println(checker.WordEdits("spelling"))
}

