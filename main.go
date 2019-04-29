package main

// TODO:
// convert yearly expense into monthly feature daily, weekly, fornightly, biannually, yearly

var dived divided

func main() {
	expense.noEsc = make(map[string]float64)
	expense.felxi = make(map[string]float64)
	expense.goals = make(map[string]float64)

	InitSvr()
}

var expense Expenses

type Expenses struct {
	noEsc map[string]float64 // Expenses you have to pay like rent, phone bill
	felxi map[string]float64 // Expenses that are flexible like fuel, entertainment.
	goals map[string]float64 // Saving or debt reduction goals
}
