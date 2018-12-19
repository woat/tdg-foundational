package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Welcome to Hangman!\n")
	exec()
	os.Exit(0)
}

func exec() {
	word := newWord("apple")
	// word := readWord("path/to/file.txt")
	dispWord := word.display()
	guessedSet := make(set, 26)
	guesses := 8

	for guesses != 0 {
		if dispWord == word.origin {
			break
		}

		fmt.Printf("The word now looks like this: %s\n", dispWord)

		if guesses > 1 {
			fmt.Printf("You have %d guesses left.\n", guesses)
		} else {
			fmt.Printf("You have only one guess left.\n")
		}

		r := bufio.NewReader(os.Stdin)
		fmt.Print("Enter letter: ")
		input, _ := r.ReadString('\n')
		letter := input[:1]

		if _, ok := guessedSet[letter]; ok {
			fmt.Printf("You already guessed %s, try again.\n", letter)
			continue
		} else {
			guessedSet[letter] = true
		}

		correct := word.guess(letter)
		if correct {
			dispWord = word.display()
			fmt.Print("That guess is correct.\n")
		} else {
			fmt.Printf("There are no %s's in the word.\n", letter)
			guesses--
		}
	}

	if guesses > 0 {
		fmt.Printf("You guessed the word: %s\n", word.origin)
		fmt.Print("You win!\n")
	} else {
		fmt.Printf("You're completely hung.\nThe word was: %s.\nYou lose.\n", word.origin)
	}
}

type word struct {
	origin string
	len    int
	set    set
}

type set map[string]bool

func (w *word) guess(letter string) bool {
	_, ok := w.set[letter]
	if ok {
		delete(w.set, letter)
	}
	return ok
}

// I'm going to have say no for part 2 graphics
func (w word) display() string {
	dw := w.origin
	for k, _ := range w.set {
		dw = strings.Replace(dw, k, "-", -1)
	}
	return string(dw)
}

func newWord(w string) *word {
	m := make(set, len(w))
	for _, r := range w {
		s := string(r)
		m[s] = true
	}

	return &word{w, len(w), m}
}

/*
Read file function for part 3
func readWord(fp string) string {
	f, _ := os.Open(fp)
	defer f.Close()

	var lns []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lns = append(lns, scanner.Text())
	}

	return lns[rand.Intn(len(lns))]
}
*/
