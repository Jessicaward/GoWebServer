package main

import (
	"GoWebServer/data"
	"GoWebServer/model"
)

func main() {
	drivers := data.GetAllDrivers()
	ferrari := new(model.Team)

	ferrari.ID = 1
	ferrari.Name = "Ferrari"
	ferrari.CarColour = "Scarlett red"
	model.PrintTeam(ferrari)

	i := 0
	for i <= len(drivers) {
		model.PrintDriver(drivers[i])
	}
}
