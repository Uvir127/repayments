package main
import (
	"testing"
)

func TestCalcDaily (t *testing.T){
	var tests = []struct {
		initialAmt, rate, period, interest, loanRepayment, installment float64
	}{
		{2000, 7, 5, 1.92, 2001.92, 400.38},
		{5000, 12, 20, 32.88, 5032.88, 251.64},
		{4600, 5, 15, 9.45, 4609.45, 307.30},
		{6438, 3.2, 31, 17.50, 6455.50, 208.24},
	}
	for _, c := range tests {
		calcInterest, calcLoanRepayment, calcDailyRepayment := calcDaily(c.initialAmt, c.rate, c.period)
		if calcInterest != c.interest{
			t.Fatal("Expected", c.interest, " got", calcInterest)
		}
		if calcLoanRepayment != c.loanRepayment{
			t.Fatal("Expected", c.loanRepayment, " got", calcLoanRepayment)
		}
		if calcDailyRepayment != c.installment{
			t.Fatal("Expected", c.installment, " got", calcDailyRepayment)
		}
	}
}

func TestCalcMonthly (t *testing.T){
	var tests = []struct {
		initialAmt, rate, period, interest, loanRepayment, installment float64
	}{
		{2000, 6, 60, 7200, 9200, 153.33},
		{15000, 8, 12, 14400, 29400, 2450},
		{8500, 12.5, 24, 25500, 34000, 1416.67},
		{8924, 7.4, 18, 11886.77, 20810.77, 1156.15},
	}
	for _, c := range tests {
		calcInterest, calcLoanRepayment, calcMonthRepayment := calcMonthly(c.initialAmt, c.rate, c.period)
		if calcInterest != c.interest{
			t.Fatal("Expected", c.interest, " got", calcInterest)
		}
		if calcLoanRepayment != c.loanRepayment{
			t.Fatal("Expected", c.loanRepayment, " got", calcLoanRepayment)
		}
		if calcMonthRepayment != c.installment{
			t.Fatal("Expected", c.installment, " got", calcMonthRepayment)
		}
	}
}
