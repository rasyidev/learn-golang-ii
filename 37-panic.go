package main

func bagi(a, b float64) float64 {
	if b == 0 {
		panic("Eits, ga boleh dibagi nol yaa!")
	}

	return a / b
}

func main() {
	bagi(1, 0)
}
