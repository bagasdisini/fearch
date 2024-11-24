package main

import (
	"fearch/document"
	"fearch/index"
	"fearch/search"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	indexPath := "index.gob"
	docPath := "doc.xml"

	term := "african cat"

	idx := make(index.Index)
	if _, err := os.Stat(indexPath); err == nil {
		start := time.Now()
		idx, err = index.LoadIndex(indexPath)
		if err != nil {
			log.Fatalf("Failed to load index: %v", err)
		}
		fmt.Printf("Loading index took %v seconds\n", time.Since(start).Seconds())
	} else {
		start := time.Now()
		docs, err := document.LoadDocuments(docPath)
		if err != nil {
			log.Fatalf("Failed to load documents: %v", err)
		}
		fmt.Printf("Loading documents took %v seconds\n", time.Since(start).Seconds())

		start = time.Now()
		idx.Add(docs)
		fmt.Printf("Indexing took %v seconds\n", time.Since(start).Seconds())

		if err := index.SaveIndex(indexPath, idx); err != nil {
			log.Fatalf("Failed to save index: %v", err)
		}
	}

	start := time.Now()
	results := search.FullText(idx, term)
	fmt.Printf("Full text search took %v\n", time.Since(start).String())

	fmt.Println(len(results))

	// TODO : Cari cara ambil text dari document berdasarkan ID

	// Get the memory usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
