package main

import (
	"net/http"
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
