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
//A =
//n = number of times interest compounded per year
//t = number of years the money is invested or borrowed

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

	interest := 0.00
	interest = initialAmt * (rate / 365.00 / 100.00) * period

	loanRepayment := initialAmt + interest
	fmt.Fprintf(w, "Loan Repayment: %.2f ", loanRepayment)
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

	interest := 0.00
	interest = initialAmt * (rate / 100.00) * period

	loanRepayment := initialAmt + interest
	fmt.Fprintf(w, "Loan Repayment: %.2f ", loanRepayment)
}





