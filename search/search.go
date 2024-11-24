package search

import (
	"fearch/document"
	"fearch/index"
	"fearch/token"
	"regexp"
	"strings"
)

func Contains(docs []document.Document, term string) []document.Document {
	var r []document.Document
	for _, doc := range docs {
		if strings.Contains(doc.Text, term) {
			r = append(r, doc)
		}
	}
	return r
}

func Regex(docs []document.Document, term string) []document.Document {
	re := regexp.MustCompile(`(?i)\b` + term + `\b`)
	var r []document.Document
	for _, doc := range docs {
		if re.MatchString(doc.Text) {
			r = append(r, doc)
		}
	}
	return r
}

func FullText(idx index.Index, text string) []int {
	var r []int
	for _, t := range token.Analyze(text) {
		if ids, ok := idx[t]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			return nil
		}
	}
	return r
}

func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}
