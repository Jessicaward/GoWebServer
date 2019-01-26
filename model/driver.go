package model

import (
	"fmt"
	"strconv"
)

type Driver struct {
	ID     int
	Name   string
	Number int
	TeamID int
}

func PrintDriver(d Driver) {
	fmt.Println(strconv.Itoa(d.Number) + ": " + d.Name)
}
