package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	result := 0
	value := -1000000007
	size := 0
	for i, _ := range A {
		if size == 0 {
			size++
			value = A[i]
		} else {
			if A[i] != value {
				size--
			} else {
				size++
			}
		}
	}
	candidate := -1000000003
	if size > 0 {
		candidate = value
	}
	count := 0
	for i, _ := range A {
		if A[i] == candidate {
			count++
		}
	}
	leader := -1000000002
	if (len(A) - count) < count {
		leader = candidate
	}

	leaderCount := 0
	for i, _ := range A {
		if A[i] == leader {
			leaderCount++
		}
		leaderInRightPart := count - leaderCount
		if (leaderCount > (i+1)/2) && (leaderInRightPart > (len(A)-i-1)/2) {
			result++
		}
	}
	return result
}
