package index

import (
	"encoding/gob"
	"fearch/document"
	"fearch/token"
	"os"
)

type Index map[string][]int

func (idx Index) Add(docs []document.Document) {
	for _, doc := range docs {
		for _, t := range token.Analyze(doc.Text) {
			ids := idx[t]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[t] = append(ids, doc.ID)
		}
	}
}

func SaveIndex(path string, idx Index) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(idx); err != nil {
		return err
	}
	return nil
}

func LoadIndex(path string) (Index, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := gob.NewDecoder(f)
	var idx Index
	if err := dec.Decode(&idx); err != nil {
		return nil, err
	}
	return idx, nil
}
