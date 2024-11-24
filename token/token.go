package token

import (
	"github.com/bbalet/stopwords"
	snowballeng "github.com/kljensen/snowball/english"
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func stopWordFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		wrd := stopwords.CleanString(token, "en", false)
		if wrd != "" {
			r = append(r, wrd)
		}
	}
	return r
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}

func Analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopWordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}
