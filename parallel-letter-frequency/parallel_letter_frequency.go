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
	//return concurrentFrequencyWithSyncMaps(xs)
	return concurrentFrequencyWithMutex(xs)
}

func concurrentFrequencyWithSyncMaps(xs []string) FreqMap {
	var (
		sm sync.Map
		wg sync.WaitGroup
	)
	fm := FreqMap{}
	for _, s := range xs {
		wg.Add(1)
		go func(s string) {
			for r, f := range Frequency(s) {
				storedF, loaded := sm.LoadOrStore(r, f)
				if loaded {
					sm.Store(r, f+storedF.(int))
				}
			}
			wg.Done()
		}(s)
	}
	wg.Wait()
	sm.Range(func(key interface{}, value interface{}) bool {
		fm[key.(rune)] = value.(int)
		return true
	})
	return fm
}

func concurrentFrequencyWithMutex(xs []string) FreqMap {
	var (
		mx sync.Mutex
		wg sync.WaitGroup
	)
	fm := FreqMap{}
	for _, s := range xs {
		wg.Add(1)
		go func(s string) {
			for r, f := range Frequency(s) {
				mx.Lock()
				fm[r] += f
				mx.Unlock()
			}
			wg.Done()
		}(s)
	}
	wg.Wait()
	return fm
}
