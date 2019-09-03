package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
    // write your code in Go 1.4
    if (A==B && K!=1) {
        if (A%K==0) {
            return 1
        } else {
            return 0
        }
    } else if (K==1) {
        return B-A+1
    }
    
    count := ((B-A)/K)+1
    lbA := (A&1)==1
    lbB := (B&1)==1
    lbK := (K&1)==1
    
    if (lbA && lbB && !lbK) {
        count--
    }
    
    return count
}