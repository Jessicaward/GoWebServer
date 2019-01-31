package controller

import (
	"GoWebServer/data"
)

func DisplayTeam(arg int) string {
	teams := data.GetAllTeams()
	team := teams[arg]
	return "<h1>" + team.Name + "</h1><div>" + team.CarColour + "</div>"
}
