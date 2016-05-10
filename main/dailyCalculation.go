package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"strconv"
)

func daily(w http.ResponseWriter, r *http.Request) {
	values := mux.Vars(r)

	//Display
	fmt.Fprintln(w, "Calculation of Daily Repayment")
	fmt.Fprintln(w, "==============================")
	fmt.Fprintln(w, "Initial Amount: R", values["initialAmt"])
	fmt.Fprintln(w, "Period: ", values["period"])

	//Variable initialisation and error checking
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

	//Calls required methods to calculate loan total and installments
	interest, totalLoanRepayment, dailyInstallment := calcDaily(initialAmt, rate, period)
	fmt.Fprintf(w, "Interest On Loan: R%.2f \n", interest)
	fmt.Fprintf(w, "Daily Installments: R%.2f \n", dailyInstallment)
	fmt.Fprintf(w, "Total Loan Repayment: R%.2f", totalLoanRepayment)
}

//Calculations for Daily loan repayment
func calcDaily(initialAmt float64, rate float64, period float64) (float64, float64, float64){
	interest := initialAmt * (rate / 365.00 / 100.00) * period
	totalLoanRepayment := initialAmt + interest
	dailyInstallment := totalLoanRepayment /period
	return Round(interest, 2), Round(totalLoanRepayment, 2), Round(dailyInstallment, 2)
}

