package tambahan

import (
	"fmt"
	"testing"
)

func TestBreakingNestedLoop(t *testing.T) {
OuterLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 3 && j == 3 {
				break OuterLoop
			}

			fmt.Printf("%v ", j)
		}
		fmt.Println("")
	}
}

/*
=== RUN   TestBreakingNestedLoop
0 1 2 3 4 5 6 7 8 9
0 1 2 3 4 5 6 7 8 9
0 1 2 3 4 5 6 7 8 9
0 1 2 --- PASS: TestBreakingNestedLoop (0.00s)
PASS
ok      learn-go-ii/tambahan    0.357s
*/
