package tambahan

import (
	"fmt"
	"testing"
)

type Assignment struct {
	FinalScore float64
	Presence   int
}

func TestSwitchDenganOperator(t *testing.T) {
	assignments := []Assignment{
		{FinalScore: 32.4, Presence: 89},
		{FinalScore: 93.2, Presence: 38},
		{FinalScore: 54.2, Presence: 95},
		{FinalScore: 32.23, Presence: 73},
	}

	for _, assignment := range assignments {
		// cetak tipe data dan value dari variabel assignment
		fmt.Printf("%#v\t: ", assignment)
		switch {
		case assignment.FinalScore >= 60 && assignment.Presence >= 60:
			fmt.Println("Anda lulus dan rajin masuk kelas")

		case assignment.FinalScore >= 60 && assignment.Presence < 60:
			fmt.Println("Anda lulus, tapi jarang hadir")

		case assignment.FinalScore < 60 && assignment.Presence >= 60:
			fmt.Println("Anda tidak lulus, tapi rajin masuk kelas")

		case assignment.FinalScore < 60 && assignment.Presence < 60:
			fmt.Println("Anda tidak lulus dan jarang masuk kelas")

		default:
			fmt.Println("Unique case nih")
		}
	}
}

/*
=== RUN   TestSwitchDenganOperator
tambahan.Assignment{FinalScore:32.4, Presence:89}       : Anda tidak lulus, tapi rajin masuk kelas
tambahan.Assignment{FinalScore:93.2, Presence:38}       : Anda lulus, tapi jarang hadir
tambahan.Assignment{FinalScore:54.2, Presence:95}       : Anda tidak lulus, tapi rajin masuk kelas
tambahan.Assignment{FinalScore:32.23, Presence:73}      : Anda tidak lulus, tapi rajin masuk kelas
--- PASS: TestSwitchDenganOperator (0.00s)
PASS
ok      learn-go-ii/tambahan    0.657s
*/
