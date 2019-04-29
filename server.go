package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type SvrOut struct {
	Remaining float64
}

func InitSvr() {
	// Live pages
	http.HandleFunc("/", calc)
	// Serve static files
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func calc(w http.ResponseWriter, r *http.Request) {
	log.Printf("Func: calc. Method:%s", r.Method)

	if r.Method == "GET" {
		t, err := template.ParseFiles("web/calculation.html")
		if err != nil {
			log.Println(err)
		} else {
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		afterTax := r.FormValue("after-tax")
		staticName := r.Form["static-name"]
		staticCost := r.Form["static-cost"]
		variableName := r.Form["variable-name"]
		variableCost := r.Form["variable-cost"]
		goalName := r.Form["goal-name"]
		goalCost := r.Form["goal-cost"]

		// something and TODO sanitize
		cash, err := strconv.ParseFloat(afterTax, 64)
		if err != nil {
			log.Printf("ERROR: Parsing float")
		}
		division := divider(cash, .60, .30, .10)
		fmt.Printf("division: %+v\n", division)

		// Populate data structures
		// TODO: sanitize user input
		loader(staticName, staticCost, expense.noEsc)
		loader(variableName, variableCost, expense.felxi)
		loader(goalName, goalCost, expense.goals)

		payExpenses(expense.noEsc, &division.required, &division.savings)
		payExpenses(expense.felxi, &division.required, &division.savings)
		payExpenses(expense.goals, &division.savings, &division.spending)

		var output SvrOut
		output.Remaining = division.required + division.savings + division.spending

		t, err := template.ParseFiles("web/calculation.html")
		if err != nil {
			log.Println(err)
		} else {
			t.Execute(w, output)
		}

	}
}

func loader(key []string, payload []string, target map[string]float64) {
	for i := 0; i < len(payload); i++ {

		// Protect []key from out of bounds.
		// Name nameless keys.
		var name string
		if len(key[i]) == 0 {
			name = "Unnamed"
		} else {
			name = key[i]
		}

		var err error
		target[name], err = strconv.ParseFloat(payload[i], 64)
		if err != nil {
			log.Println(err)
		}
	}
}
