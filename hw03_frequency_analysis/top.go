package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(t string) []string {
	wordList := strings.Fields(t)
	m := make(map[string]int)

	for _, word := range wordList {
		m[word]++
	}

	type kv struct {
		word  string
		count int
	}

	wordFreq := make([]kv, 0, len(m))

	for word, count := range m {
		wordFreq = append(wordFreq, kv{word, count})
	}

	sort.Slice(wordFreq, func(i, j int) bool {
		if wordFreq[i].count == wordFreq[j].count {
			return wordFreq[i].word < wordFreq[j].word
		}
		return wordFreq[i].count > wordFreq[j].count
	})

	topWords := make([]string, 0, 10)
	for i := 0; i < len(wordFreq) && i < 10; i++ {
		if i >= len(wordFreq) {
			break
		}
		topWords = append(topWords, wordFreq[i].word)
	}

	return topWords
}
