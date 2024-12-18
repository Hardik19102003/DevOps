package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	fmt.Println("Listening on localhost:4242...")
	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	response := []byte("Server is up and running!")

	fmt.Println("Health endpoint called!")

	_, err := writer.Write(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Health endpoint called!")
}

func returnError(password string) error {
	var secretPassword string = "supersecretpassword"
	if password == secretPassword {
		return nil
	} else {
		return errors.New("invalid password")
	}

}

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Request was correct and method was POST")
}
