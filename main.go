package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/createIntent", handleCreatePaymentIntent)

	read, ss := http.Get("http://localhost")

	fmt.Println("__________")
	fmt.Println(read, ss)
	var err error = http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handleRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("reqest from ", request.RemoteAddr)
	response := []byte("Server root requested")
	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("the req is correct ")
}

var req struct {
	ProductId string `json:"product_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("reqest from ", request.RemoteAddr)

	response := []byte("Server is up and running")
	writer.Write(response)

}
