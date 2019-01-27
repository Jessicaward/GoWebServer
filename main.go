package main

import (
	"GoWebServer/data"
	"GoWebServer/model"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The best team is %s!", r.URL.Path[1:])

	if r.URL.Path[1:] == "test" {
		testData()
	}
}

func driverHandler(w http.ResponseWriter, r *http.Request) {
	drivers := data.GetAllDrivers()

	//ID = every character after :8080/driver/.
	//err = error stemming from type conversion
	id, err := strconv.Atoi(r.URL.Path[len("/driver/"):])

	//Type conversion error checking
	if err != nil {
		log.Fatal(err)
	}

	//Gets driver by id (rly inefficient atm, I know)
	driver := drivers[id]

	//sends basic HTML as HTTP response text (uses string interpolation to insert driver name and number)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", driver.Name, driver.Number)
}

func main() {
	//Routes
	//todo: move these to a routing packge
	http.HandleFunc("/driver/", driverHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//NOTE: this is temporary, used for getting driver/team IDs etc
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
