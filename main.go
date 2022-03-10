package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const PORT string = ":8080"

func mul(w http.ResponseWriter, r *http.Request) {
	op1, err := strconv.Atoi(mux.Vars(r)["op1"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	op2, err := strconv.Atoi(mux.Vars(r)["op2"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	product := op1 * op2

	fmt.Fprintf(w, "%v", product)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/mul/{op1}/{op2}", mul)

	log.Printf("Listening on port %v", PORT)
	http.ListenAndServe(PORT, r)
}
