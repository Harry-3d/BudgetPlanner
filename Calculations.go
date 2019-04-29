package main

import "fmt"

func minusTax(before, tax float64) float64 {
	return before * (1 - tax)
}

// divider takes a dollar amount and devides it between 3 percentage amounts.
func divider(after float64, percentage1, percentage2, percentage3 float64) divided {
	var divs divided
	//after := *afterPtr
	divs.required = after * percentage1
	divs.savings = after * percentage2
	divs.spending = after * percentage3

	return divs
}

type divided struct {
	required float64
	savings  float64
	spending float64
}

func payExpenses(m map[string]float64, toPay *float64, overDrawFunds *float64) {
	var overDraw float64
	for key, val := range m {
		*toPay = *toPay - val
		fmt.Printf("%-20s cost $%-12.2f remaining:%.2f\n", key, val, *toPay)
		if *toPay < 0 {
			overDraw += *toPay
			//fmt.Println("overdraw:", overDraw)
		}
	}

	// If the expenses exceed the divided.required funds, pull funds from divided.savings
	if overDraw < 0 {
		fmt.Printf("Exceeded funds, using balance from lower division($%.2f). ", *overDrawFunds)
		*overDrawFunds += overDraw
		//if *overDrawFunds > 0 {
		*toPay = 0
		//}
		fmt.Printf("After correction lower division contains $%.2f\n", *overDrawFunds)

	}
}
