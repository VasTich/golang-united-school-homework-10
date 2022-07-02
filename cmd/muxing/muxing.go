package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

/*
| GET    | `/name/{PARAM}`                       | body: `Hello, PARAM!`         |
| GET    | `/bad`                                | Status: `500`                 |
| POST   | `/data` + Body `PARAM`                | body: `I got message:\nPARAM` |
| POST   | `/headers`+ Headers{"a":"2", "b":"3"} | Header `"a+b": "5"`           |
*/
func handleGetParam(w http.ResponseWriter, r *http.Request) {
	param, ok := mux.Vars(r)["PARAM"]
	if !ok {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello, %s!", param)))
}

func handleGetBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func handlePostData(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("I got message:\n%s", d)))
}

func handlePostHeaders(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")

	log.Println(a)
	log.Println(b)

	if len(a) == 0 || len(b) == 0 {
		return
	}

	aval, err := strconv.Atoi(a)
	if err != nil {
		return
	}

	bval, err := strconv.Atoi(b)
	if err != nil {
		return
	}

	w.Header().Set("a+b", strconv.Itoa(aval+bval))
	w.WriteHeader(http.StatusOK)
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", handleGetParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleGetBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handlePostData).Methods(http.MethodPost)
	router.HandleFunc("/headers", handlePostHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
