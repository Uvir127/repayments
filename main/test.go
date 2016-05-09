package main
import (
	"testing"
)

func TestCalcDailyRepayment (t *testing.T){
	var tests = []struct {
		initialAmt, rate, period, interest, loanRepayment, dailyRepayment float64
	}{
		{2000, 9, 12, 5.92, 2005.92, 167.16},
		{5000, 12, 20, 5032.88, 251.64},
		//{2000.0, 9.0, 12.0, 6.0},
		//{2000.0, 9.0, 12.0, 6.0},
	}
	for _, c := range tests {
		calcInterest := calcDailyInterest(c.initialAmt, c.rate, c.period)
		if calcInterest != c.interest{
			t.Fatal("Expected", c.interest, " got", calcInterest)
		}

		calcLoanRepayment, calcDailyRepayment := calcDailyLoanRepayment(c.initialAmt, c.rate, c.period)
		if calcLoanRepayment != c.loanRepayment{
			t.Fatal("Expected", c.loanRepayment, " got", calcLoanRepayment)
		}
		if calcDailyRepayment != c.dailyRepayment{
			t.Fatal("Expected", c.dailyRepayment, " got", calcDailyRepayment)
		}
	}
}

//func TestDaily (t *testing.T) {
//	var tests = []struct {
//		s, want string
//	}{
//		{"Backward", "drawkcaB"},
//		{"Hello, World", "dlroW ,olleH"},
//		{"", ""},
//	}
//	for _, c := range tests {
//		got := Reverse(c.s)
//		if got != c.want {
//			t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
//		}
//	}
//}

////Testing URl - Not required
//func Test (t *testing.T){
//	//router := mux.NewRouter().StrictSlash(true)
//	s := httptest.NewServer("/calcMonthly/")
//	defer s.Close()
//
//	resp, err := http.Get(s.URL)
//	if err != nil {
//		t.Error(err)
//	}
//	if body, err := ioutil.ReadAll(resp.Body);err != nil {
//		t.Error(err)
//	} else if string(body) != ""{
//		t.Error("Expected", "Something", "Got", body)
//	}
//}

////Example Code
//type App struct{
//	Message string
//}
//
//func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request){
//	w.Write([]byte(a.Message))
//}
//
//func fakeApp(msg string) *httptest.Server{
//	app:= &App{Message:msg}
//	return httptest.NewServer(app)
//}
//
//func getBody(t *testing.T, s *httptest.Server, path string) string {
//	resp, err := http.Get(s.URL + path)
//	if err != nil {
//		t.Error(err)
//	}
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		t.Error(err)
//	}
//	return string(body)
//}
