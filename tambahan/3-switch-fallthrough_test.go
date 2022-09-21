package tambahan

import (
	"fmt"
	"testing"
)

func TestSwitchFallthrough(t *testing.T) {
	FinalScore := []int{23, 57, 54, 62, 80}

	for _, score := range FinalScore {
		fmt.Printf("%#v\t: ", score)
		switch {
		case score >= 75:
			fmt.Println("Ntaps")

		case score >= 60:
			fmt.Println("Mayan")

		case score < 60:
			fmt.Println("Duh bro, belajar lagi bro!")
			fallthrough

		case score < 50:
			fmt.Println("Belajar geh bro!")
		}
	}
}

/*
=== RUN   TestSwitchFallthrough
23      : Duh bro, belajar lagi bro!
Belajar geh bro!
57      : Duh bro, belajar lagi bro!
Belajar geh bro!
54      : Duh bro, belajar lagi bro!
Belajar geh bro!
62      : Mayan
80      : Ntaps
--- PASS: TestSwitchFallthrough (0.00s)
PASS
ok      learn-go-ii/tambahan    0.443s
*/
