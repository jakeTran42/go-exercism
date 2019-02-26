package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// SafeCounter : SafeCounter is a struct that holds a
// FreqMap and a mutex, which is used to only allow one
// goroutine access at a time.
type SafeCounter struct {
	histoMap FreqMap
	mux      sync.Mutex
}

// Increment : Increment is a SafeCounter method that
// increases the count of a rune in SafeCounter's histoMap.
// Inspired from https://tour.golang.org/concurrency/9
func (counter *SafeCounter) Increment(char rune) {
	counter.mux.Lock()
	counter.histoMap[char]++
	counter.mux.Unlock()
}

// ConcurrentFrequency : ConcurrentFrequency is a function
// that accepts a splice of texts, counts the number of each
// rune in the texts concurrently, and returns the
// histogram of rune counts
func ConcurrentFrequency(texts []string) FreqMap {
	count := SafeCounter{histoMap: make(map[rune]int)}

	for _, text := range texts {
		for _, char := range text {
			go count.Increment(char)
		}
	}

	return count.histoMap
}
