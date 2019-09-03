package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    sort.Ints(A)
    answer := 1
    for i := 0; i < len(A); i++ {
        if(A[i] == answer){
            answer++
        }     
    }
    return answer
}