package model

import (
	"fmt"
	"strconv"
)

type Team struct {
	ID        int
	Name      string
	CarColour string
}

func PrintTeam(t *Team) {
	fmt.Println(strconv.Itoa(t.ID) + ": " + t.Name + " is " + t.CarColour)
}
