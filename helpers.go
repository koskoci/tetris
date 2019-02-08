package main

func index(s shape, p [2]int) int {
	for i, pixel := range s {
		if pixel == p {
			return i
		}
	}
	return -1
}

func includes(x []int, n int) bool {
	for _, m := range x {
		if m == n {
			return true
		}
	}
	return false
}

func hasBit(n uint32, pos uint) bool {
	val := n & (1 << pos)
	return val > 0
}
