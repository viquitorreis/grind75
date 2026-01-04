package main

func main() {
	s := "jjih"
	k := 4
	println(reversePrefix(s, k))
}

func reversePrefix(s string, k int) string {
	// str := make([]byte)
	str := []byte{}
	slow, fast := 0, k

	for range k {
		str = append(str, 'a')
	}

	println("str is: ", string(str))

	for fast >= slow {
		last := s[fast-1]
		first := s[slow]
		temp := first
		// fmt.Printf("fast: %d - slow: %d - last: %s - first: %s - temp: %s - str: %s\n",
		// 	fast, slow, string(last), string(first), string(temp), string(str))
		str[slow] = last
		str[fast-1] = temp
		fast--
		slow++

	}

	for i := k; i < len(s); i++ {
		str = append(str, byte(s[i]))
	}

	return string(str)
}
