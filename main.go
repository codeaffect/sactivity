package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	fmt.Print("Port:" + port)

	if port == "" {
		fmt.Print("$PORT must be set\n")
		port = "8081"
		fmt.Print("Port is set to:" + port)
	}

	setRoutesRequest()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setRoutesRequest() {
	http.HandleFunc("/", apiHome)
	http.HandleFunc("/about", apiAbout)
}

func apiHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home is here!")
}
func apiAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about ME!")
}

func getEnv() string {
	return os.Getenv("APP_ENV")
}
