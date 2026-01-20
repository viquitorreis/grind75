package main

func main() {
}

type TimeMap struct {
	store map[string][]Pair // chave -> lista de (timestamp, value)
}

type Pair struct {
	timestamp int
	value     string
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]Pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], Pair{
		timestamp: timestamp,
		value:     value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	var idx int
	val, ok := this.store[key]
	if ok {
		idx = min(val[len(val)-1].timestamp, timestamp)
	} else {
		return ""
	}

	for i := idx; i >= 0; i-- {
		if val[i].timestamp <= timestamp {
			return val[i].value
		}
	}

	return ""
}
