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

	//v := new(calculatedValues)
	//v := calcDaily(initialAmt, rate, period)
	//interest, totalLoanRepayment, installment := calcDaily(initialAmt, rate, period)
	interest := calcDailyInterest(initialAmt, rate, period)
	totalLoanRepayment := calcDailyLoanRepayment(initialAmt, interest)
	installment := calcDailyInstallment(totalLoanRepayment, period)

	fmt.Fprintf(w, "Interest On Loan: R%.2f \n", interest)
	fmt.Fprintf(w, "Daily Installments: R%.2f \n", installment)
	fmt.Fprintf(w, "Total Loan Repayment: R%.2f", totalLoanRepayment)
}

func calcDailyInterest(initialAmt float64, rate float64, period float64) float64 {
	return Round(initialAmt * (rate / 365.00 / 100.00) * period, 2)
}

func calcDailyLoanRepayment(initialAmt float64, interest float64) float64{
	return Round(initialAmt + interest,2)
}

func calcDailyInstallment(totalLoanRepayment float64, period float64) float64 {
	return Round(totalLoanRepayment /period,2 )
}
//Calculations for Daily loan repayment
//func calcDaily(initialAmt float64, rate float64, period float64) (v *calculatedValues){
//	interest := initialAmt * (rate / 365.00 / 100.00) * period
//	totalLoanRepayment := initialAmt + interest
//	dailyInstallment := totalLoanRepayment /period
//	v.interest = Round(interest, 2)
//	v.totalLoanRepayment = Round(totalLoanRepayment, 2)
//	v.installment = Round(dailyInstallment, 2)
//	return v
//}

//func calcDaily(initialAmt float64, rate float64, period float64) (float64, float64, float64){
//	interest := initialAmt * (rate / 365.00 / 100.00) * period
//	totalLoanRepayment := initialAmt + interest
//	installment := totalLoanRepayment /period
//	interest = Round(interest, 2)
//	totalLoanRepayment = Round(totalLoanRepayment, 2)
//	installment = Round(installment, 2)
//	return interest, totalLoanRepayment, installment
//}
