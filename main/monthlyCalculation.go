package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"strconv"
)

func monthly(w http.ResponseWriter, r *http.Request){
	values := mux.Vars(r)

	//Display
	fmt.Fprintln(w, "Calculation of Monthly Repayment")
	fmt.Fprintln(w, "================================")
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
	interest, totalLoanRepayment, monthlyInstallment := calcMonthly(initialAmt, rate, period)
	fmt.Fprintf(w, "Interest On Loan: R%.2f \n", interest)
	fmt.Fprintf(w, "Monthly Installment: R%.2f \n", monthlyInstallment)
	fmt.Fprintf(w, "Loan Repayment: R%.2f", totalLoanRepayment)
}

//Calculations for Monthly loan repayment
func calcMonthly(initialAmt float64, rate float64, period float64) (float64, float64, float64){
	interest := initialAmt * (rate / 100.00) * period
	totalLoanRepayment := initialAmt + interest
	monthlyInstallment := totalLoanRepayment /period
	return Round(interest, 2), Round(totalLoanRepayment, 2), Round(monthlyInstallment,2)
}
