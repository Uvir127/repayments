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
	interest := calcMonthlyInterest(initialAmt, rate, period)
	totalLoanRepayment := calcMonthlyLoanRepayment(initialAmt, interest)
	installment := calcMonthlyInstallment(totalLoanRepayment, period)

	fmt.Fprintf(w, "Interest On Loan: R%.2f \n", interest)
	fmt.Fprintf(w, "Monthly Installment: R%.2f \n", installment)
	fmt.Fprintf(w, "Loan Repayment: R%.2f", totalLoanRepayment)
}

func calcMonthlyInterest(initialAmt float64, rate float64, period float64) float64 {
	return Round(initialAmt * (rate / 100.00) * period, 2)
}

func calcMonthlyLoanRepayment(initialAmt float64, interest float64) float64{
	return Round(initialAmt + interest,2)
}

func calcMonthlyInstallment(totalLoanRepayment float64, period float64) float64 {
	return Round(totalLoanRepayment /period,2 )
}