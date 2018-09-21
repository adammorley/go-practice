package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type wordCount map[string]int

// scan a file, eg a book, and find the count of all the words in the file
// scanning by rune to skip punctuation.  this depends on the underlying
// buffer cache implementation in the filesystem of the operating system
// so it may be required to tune the page algorithm
func (w wordCount) scanFile(file *os.File, logger *log.Logger) {
	var digitCount, invalidCount, symbolCount int
	var sb strings.Builder
	rr := bufio.NewReader(file)
	for {
		if r, _, err := rr.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if unicode.IsLetter(r) {
				sb.WriteRune(r)
			} else if unicode.IsDigit(r) {
				digitCount++
			} else if unicode.IsSymbol(r) {
				symbolCount++
			} else if unicode.IsPunct(r) || unicode.IsSpace(r) {
				if sb.Len() > 0 {
					w[sb.String()]++
					sb.Reset()
				}
			} else {
				invalidCount++
			}
		}
	}
	logger.Print("found ", strconv.Itoa(digitCount), " digits, ", strconv.Itoa(symbolCount), " symbols and ", strconv.Itoa(invalidCount), " invalid runes")
}

type pair struct {
	word  string
	count int
}

// return the top ten words with length > n in a wordCount
func (w wordCount) top10(n int) []pair {
	var words []pair = make([]pair, 0, len(w)) // pre-allocate space
	for k, v := range w {
		if len(k) > n {
			words = append(words, pair{word: k, count: v})
		}
	}
	sort.Slice(words, func(i, j int) bool { return words[i].count > words[j].count })
	return words[0:10]
}

// to parallelize, add scanning and top10() calculation equal to roughly the number of cores, modulo the disk
// bandwidth.  for example, in the context of a modern ssd, the number of parallel io streams available is non
// obvious and requires examination; it also rarely matches the number of cores in most system designs
// additionally, the books may be in a object store, which means building a pipeline for fetching books
// and processing them and storing metadata is another consideration
// parallelization is likely best done using goroutines
func main() {
	var logger *log.Logger = log.New(os.Stderr, "error: ", log.Ltime)
	filename := flag.String("filename", "", "the filename to process")
	wordLength := flag.Int("length", 5, "the word length to calculate the top 10 words for")
	flag.Parse()
	if _, err := os.Stat(*filename); err != nil {
		logger.Fatal("could not find file, please specify as the filename option")
	}
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal("could not open book")
	}
	var wc wordCount = map[string]int{} // defined as a type to reduce copying
	wc.scanFile(f, logger)
	fmt.Println("top ten words:")
	c := wc.top10(*wordLength)
	for _, p := range c {
		fmt.Println(p.word, p.count)
	}
}
