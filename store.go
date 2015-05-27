package main

import (
	"strings"
	"sync"
	"unicode"
)

type Store struct {
	sync.Mutex
	Words   sMap
	Letters sMap
}

func (s *Store) Add(words []string) {
	s.Lock()
	for _, word := range words {
		if len(word) >= 2 {

			var w []rune

			for _, runeValue := range word {
				// TODO: use unicode.IsOneOf()
				if unicode.IsLetter(runeValue) || unicode.IsNumber(runeValue) {
					w = append(w, runeValue)
					s.Letters.All = append(s.Letters.All, strings.ToLower(string(runeValue)))
				}
			}

			if len(w) != 0 {
				s.Words.All = append(s.Words.All, strings.ToLower(string(w)))
			}
		}
	}
	s.MakeSorted()
	s.Unlock()
}

func (s *Store) MakeSorted() {
	// TODO: Impl concurrency sort
	s.Words.MakeSorted()
	s.Letters.MakeSorted()
}
