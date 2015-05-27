package main

import "sort"

type sMap struct {
	Counter map[string]int
	All     []string
	Sorted  []string
}

func (sm *sMap) Len() int {
	return len(sm.Counter)
}

func (sm *sMap) Less(i, j int) bool {
	return sm.Counter[sm.Sorted[i]] > sm.Counter[sm.Sorted[j]]
}

func (sm *sMap) Swap(i, j int) {
	sm.Sorted[i], sm.Sorted[j] = sm.Sorted[j], sm.Sorted[i]
}

func (sm *sMap) MakeSorted() {
	sm.Counter = make(map[string]int)

	for _, v := range sm.All {
		sm.Counter[v]++
	}
	sm.Sorted = make([]string, len(sm.Counter))

	var i int
	for key, _ := range sm.Counter {
		sm.Sorted[i] = key
		i++
	}
	sort.Sort(sm)
}
