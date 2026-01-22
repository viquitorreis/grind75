package main

import "fmt"

func main() {
	ts := Constructor()
	ts.Set("foo", "bar", 1)
	ts.Get("foo", 1)
	ts.Get("foo", 3)
	ts.Set("foo", "bar2", 4)
	fmt.Println(ts.Get("foo", 4))
	fmt.Println(ts.Get("foo", 5))
}

type Pair struct {
	value     string
	timestamp int
}

type TimeMap struct {
	keyVal map[string][]Pair
}

func Constructor() TimeMap {
	return TimeMap{
		keyVal: make(map[string][]Pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.keyVal[key]; ok {
		this.keyVal[key] = append(this.keyVal[key], Pair{value: value, timestamp: timestamp})
	} else {
		this.keyVal[key] = []Pair{
			{value: value, timestamp: timestamp},
		}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	val, ok := this.keyVal[key]
	if !ok || len(val) == 0 {
		return ""
	}

	left, right := 0, len(val)-1
	result := ""

	for left <= right {
		mid := left + (right-left)/2
		if val[mid].timestamp <= timestamp {
			result = val[mid].value // candidato possível
			left = mid + 1          // tenta achar timestamp maior ainda
		} else {
			right = mid - 1
		}
	}

	return result

}
