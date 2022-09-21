package tambahan

import (
	"fmt"
	"testing"
)

func TestForLoopCondition1(t *testing.T) {
	i := 0

	for i < 3 {
		fmt.Println("Perulangan ke-", i+1)
		i++
	}
}

/*
=== RUN   TestForLoopCondition1
Perulangan ke- 1
Perulangan ke- 2
Perulangan ke- 3
--- PASS: TestForLoopCondition1 (0.00s)
PASS
ok      learn-go-ii/tambahan    0.374s
*/

func TestForLoopCondition2(t *testing.T) {
	i := 0

	for {
		fmt.Println("Perulangan ke-", i+1)

		// increment
		i++

		// stop condition
		if i == 3 {
			break
		}
	}
}

/*
=== RUN   TestForLoopCondition2
Perulangan ke- 1
Perulangan ke- 2
Perulangan ke- 3
--- PASS: TestForLoopCondition2 (0.00s)
PASS
ok      learn-go-ii/tambahan    0.493s
*/
