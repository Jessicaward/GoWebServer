package route

import (
	"GoWebServer/controller"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/driver", driverHandler)
	http.HandleFunc("/team/", teamHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The best team is %s!", r.URL.Path[1:])

	if r.URL.Path[1:] == "test" {
		//todo: display test data?
	}
}

func driverHandler(w http.ResponseWriter, r *http.Request) {
	arg := r.URL.Path[len("/driver/")]

	if string(arg) == "add" {
		//todo: add new driver
	} else {
		controller.DisplayDriver(int(arg))
	}
}

func teamHandler(w http.ResponseWriter, r *http.Request) {
	arg := r.URL.Path[len("/driver/")]

	if string(arg) == "add" {
		//todo: add new team
	} else {
		controller.DisplayTeam(int(arg))
	}
}
