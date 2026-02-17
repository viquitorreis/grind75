package main

func main() {

}

type Pair struct {
	timestamp int
	val       string
}

type TimeMap struct {
	store map[string][]Pair
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]Pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	// this.store =
	if _, exists := this.store[key]; exists {
		this.store[key] = append(this.store[key], Pair{
			timestamp: timestamp,
			val:       value,
		})
	} else {
		this.store[key] = []Pair{{
			timestamp: timestamp,
			val:       value,
		}}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	val, exists := this.store[key]
	if !exists {
		return ""
	}

	for i := len(val) - 1; i >= 0; i-- {
		if val[i].timestamp <= timestamp {
			return val[i].val
		}
	}

	return ""
}
