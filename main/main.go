package main

import (
	"net/http"
	"fmt"
	"strconv"
	"log"
	"github.com/gorilla/mux"
	"math"
)

func main() {
	//Checks URL to see what functions need to be called
	//Stores required details
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/calcMonthly/{initialAmt}&{rate}&{period}", monthly)
	router.HandleFunc("/calcDaily/{initialAmt}&{rate}&{period}", daily)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func daily(w http.ResponseWriter, r *http.Request) {
	values := mux.Vars(r)

	//Display
	fmt.Fprintln(w, "Calculation of Daily Repayment")
	fmt.Fprintln(w, "==============================")
	fmt.Fprintln(w, "Initial Amount: ", values["initialAmt"])
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


func monthly(w http.ResponseWriter, r *http.Request){
	values := mux.Vars(r)

	//Display
	fmt.Fprintln(w, "Calculation of Monthly Repayment")
	fmt.Fprintln(w, "================================")
	fmt.Fprintln(w, "Initial Amount: ", values["initialAmt"])
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

//Rounds Float
func Round(float float64, decimals int) (rValue float64) {
	var round float64
	pow := math.Pow(10, float64(decimals))
	digits := pow * float
	if (digits - 0.5) < math.Floor(digits){
		round = math.Floor(digits)
	} else {
		round = math.Ceil(digits)
	}
	rValue = round/pow
	return
}

