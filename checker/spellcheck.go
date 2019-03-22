package checker

type SpellChecker struct {
	words map[string]int32
}

func(s *SpellChecker) Build() {
	s.words = make(map[string]int32)
}

func (s *SpellChecker) Train(word string) {
	s.words[word] = s.words[word] + 1
}

func (s *SpellChecker) Correct(word string) string {
	// Check if we already have the word
	m := s.words[word]
	if m != 0 {
		return word
	}

	// Collection of possible matches
	//possibleMatches := make(map[string]int32)

	edits := WordEdits(word)

	for _, e := range edits {
		m := s.words[e]

		if m != 0 {
			return e
		}
	}


	return ""
}


// Returns a slice of strings which are similar to the
// supplied word and indicate a "misspelling"
func WordEdits(word string) []string {
	closeWords := []string{}

	// Perhaps they mistyped a letter try mutating one
	for i := 0; i<len(word); i++ {
		for letter := 'a'; letter <= 'z'; letter++ {
			closeWords = append(closeWords, word[:i] + string(letter) + word[i+1:])
		}
	}

	// Perhaps they missed a letter try adding one in
	for i := 0; i<len(word) + 1; i++ {
		for letter := 'a'; letter <= 'z'; letter++ {
			closeWords = append(closeWords, word[:i] + string(letter) + word[i:])
		}
	}

	// Perhaps they added an extra letter try removing it
	for i := 0; i<len(word); i++ {
		closeWords = append(closeWords, word[:i] + word[i+1:])
	}

	return closeWords
}


func NewSpellChecker() *SpellChecker {
	spellChecker := SpellChecker{}
	spellChecker.words = make(map[string]int32)
	return &spellChecker
}