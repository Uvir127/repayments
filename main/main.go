package main

import (
	"net/http"
	"fmt"
	"strconv"
	"log"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/calcMonthly/{initialAmt}&{rate}&{period}", calcMonthly)
	router.HandleFunc("/calcDaily/{initialAmt}&{rate}&{period}", calcDaily)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func calcDaily(w http.ResponseWriter, r *http.Request) {
	values := mux.Vars(r)

	fmt.Fprintln(w, "Daily Repayment")
	fmt.Fprintln(w, "Initial Amount: ", values["initialAmt"])
	fmt.Fprintln(w, "Period: ", values["period"])

	initialAmt, err := strconv.ParseFloat(values["initialAmt"], 64)
	if err != nil {
		fmt.Fprintln(w, "Error Converting Initial Amount")
	}

	period, err  := strconv.ParseFloat(values["period"], 64)
	if err != nil {
		fmt.Fprintln(w, "Error Converting Period")
	}

	rate, err := strconv.ParseFloat(values["rate"], 64)
	if err != nil{
		fmt.Fprintln(w, "Error Converting Rate")
	}

	loanRepayment := initialAmt + calcDailyInterest(initialAmt, rate, period)
	dailyRepayment := loanRepayment/period
	fmt.Fprintf(w, "Daily Repayment: %.2f \n", dailyRepayment)
	fmt.Fprintf(w, "Total Loan Repayment: %.2f ", loanRepayment)
}

func calcDailyInterest(initialAmt float64, rate float64, period float64) float64 {

	interest := 0.00
	interest = initialAmt * (rate / 365.00 / 100.00) * period
	return interest
}

func calcMonthly(w http.ResponseWriter, r *http.Request){
	values := mux.Vars(r)

	fmt.Fprintln(w, "Monthly Repayment")
	fmt.Fprintln(w, "Initial Amount: ", values["initialAmt"])
	fmt.Fprintln(w, "Period: ", values["period"])

	initialAmt, err := strconv.ParseFloat(values["initialAmt"], 64)
	if err != nil {
		fmt.Fprintln(w, "Error Converting Initial Amount")
	}

	period, err  := strconv.ParseFloat(values["period"], 64)
	if err != nil {
		fmt.Fprintln(w, "Error Converting Period")
	}

	rate, err := strconv.ParseFloat(values["rate"], 64)
	if err != nil{
		fmt.Fprintln(w, "Error Converting Rate")
	}

	loanRepayment := initialAmt + calcMonthlyInterest(initialAmt, rate, period)
	monthlyRepayment := loanRepayment/period
	fmt.Fprintf(w, "Daily Repayment: %.2f ln", monthlyRepayment)
	fmt.Fprintf(w, "Loan Repayment: %.2f ", loanRepayment)
}

func calcMonthlyInterest(initialAmt float64, rate float64, period float64) float64 {
	interest := 0.00
	interest = initialAmt * (rate / 100.00) * period
	return interest
}



