// Package letter provides functions to calculate the Frequency of letters in strings.
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

// ConcurrentFrequency concurrently calls `Frequency` on multiple strings and
// Combine the results.
func ConcurrentFrequency(xs []string) FreqMap {
	return concurrentFrequencyWithSyncMaps(xs)
	// return concurrentFrequencyWithRWMutex(xs)
}

func concurrentFrequencyWithSyncMaps(xs []string) FreqMap {
	var (
		sm sync.Map
		wg sync.WaitGroup
	)
	fm := FreqMap{}
	for i, s := range xs {
		wg.Add(1)
		go func(k int, v string) {
			sm.Store(k, Frequency(v))
			wg.Done()
		}(i, s)
	}
	wg.Wait()
	sm.Range(func(key interface{}, value interface{}) bool {
		for r, f := range value.(FreqMap) {
			fm[r] += f
		}
		return true
	})
	return fm
}

func concurrentFrequencyWithRWMutex(xs []string) FreqMap {
	var (
		mx sync.RWMutex
		wg sync.WaitGroup
	)
	xfm := []FreqMap{}
	fm := FreqMap{}
	for _, s := range xs {
		wg.Add(1)
		go func(v string) {
			mx.Lock()
			xfm = append(xfm, Frequency(v))
			mx.Unlock()
			wg.Done()
		}(s)
	}
	wg.Wait()
	for _, m := range xfm {
		for r, f := range m {
			fm[r] += f
		}
	}
	return fm
}
