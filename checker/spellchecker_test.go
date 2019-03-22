package checker

import (
	"testing"
)

func TestCheckerSingleWord(t *testing.T) {
	spellChecker := NewSpellChecker()
	spellChecker.Train("test")

	suggested := spellChecker.Correct("test")

	if suggested != "test" {
		t.Errorf("Should be test got '%s'", suggested)
	}
}

func TestCheckerTwoWords(t *testing.T) {
	spellChecker := NewSpellChecker()
	spellChecker.Train("test")
	spellChecker.Train("best")

	suggested := spellChecker.Correct("best")

	if suggested != "best" {
		t.Errorf("Should be best got '%s'", suggested)
	}
}

func TestCheckerCorrect(t *testing.T) {
	spellChecker := NewSpellChecker()
	spellChecker.Train("test")

	suggested := spellChecker.Correct("tost")

	if suggested != "test" {
		t.Errorf("Should be test got '%s'", suggested)
	}
}

func TestCheckerCorrectLonger(t *testing.T) {
	spellChecker := NewSpellChecker()
	spellChecker.Train("best")

	suggested := spellChecker.Correct("bost")

	if suggested != "best" {
		t.Errorf("Should be best got '%s'", suggested)
	}
}