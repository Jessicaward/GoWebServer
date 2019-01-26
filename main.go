package main

import (
	"GoWebServer/data"
	"GoWebServer/model"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The best team is %s!", r.URL.Path[1:])

	if r.URL.Path[1:] == "test" {
		testData()
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testData() {
	//Terminal padding
	fmt.Println()

	drivers := data.GetAllDrivers()
	teams := data.GetAllTeams()

	i := 0
	for i <= (len(drivers) - 1) {
		model.PrintDriver(drivers[i])
		i++
	}

	i = 0
	for i <= (len(teams) - 1) {
		model.PrintTeam(teams[i])
		i++
	}

	//Terminal padding
	fmt.Println()
}
