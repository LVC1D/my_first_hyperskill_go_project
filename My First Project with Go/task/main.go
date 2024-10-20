package main

import "fmt"

func main() {
	// Write your code here
	/*
		fmt.Println("Prices:")
		items := []string{
			"Bubblegum",
			"Toffee",
			"Ice cream",
			"Milk chocolate",
			"Doughnut",
			"Pancake",
		}
		prices := []float32{
			2.0,
			0.2,
			5.0,
			4.0,
			2.5,
			3.2,
		}
	*/
	totalIncome := map[string]int{
		"Bubblegum":      202,
		"Toffee":         118,
		"Ice cream":      2250,
		"Milk chocolate": 1680,
		"Doughnut":       1075,
		"Pancake":        80,
	}

	var income, staffExp, otherExp int

	fmt.Println("Earned amount:")
	for key, value := range totalIncome {
		fmt.Printf("%s: $%d\n", key, value)
	}

	for _, value := range totalIncome {
		income += value
	}

	fmt.Printf("\nIncome: $%d\n", income)
	fmt.Println("Staff expenses:")
	fmt.Scan(&staffExp)

	fmt.Println("Other expenses:")
	fmt.Scan(&otherExp)

	fmt.Printf("\nNet income: $%d\n", income-staffExp-otherExp)

}
