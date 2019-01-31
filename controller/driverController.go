package controller

import (
	"GoWebServer/data"
	"strconv"
)

func DisplayDriver(arg int) string {
	drivers := data.GetAllDrivers()

	//Gets driver by arg (rly inefficient atm, I know)
	driver := drivers[arg]

	//sends basic HTML as HTTP response text (uses string interpolation to insert driver name and number)
	return "<h1>" + driver.Name + "</h1><div>" + strconv.Itoa(driver.Number) + "</div>"
}
