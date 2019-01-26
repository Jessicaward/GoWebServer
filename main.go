package main

import (
	"GoWebServer/data"
	"GoWebServer/model"
)

func main() {
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
}
