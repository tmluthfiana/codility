package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// write your code in Go 1.4
	max0s := 0
	cur0s := -1
	for i := uint8(0); i < 31; i++ {
		if N&(1<<i) > 0 { // Bit is 1
			if cur0s > max0s {
				max0s = cur0s
			}
			cur0s = 0
		} else if cur0s != -1 {
			cur0s++
		}
	}
	return max0s
}

/*func Solution(N int) int {
	n := int64(N)
	longest := 0
	binary := strconv.FormatInt(n, 2)
	listGaps := strings.Split(binary, "1")
	for index, element := range listGaps {
		if binary[len(binary)-1:] == "0" && index == len(listGaps)-1 {
			break
		}
		if len(element) > longest {
			longest = len(element)
		}
	}
	return longest
}*/
