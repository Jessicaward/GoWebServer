package data

import (
	"GoWebServer/model"
)

//NOTE: This is a temporary data source, it's so horrible atm oh god what is happening
func GetAllDrivers() []model.Driver {
	Drivers := make([]model.Driver, 0, 20)

	Drivers = append(Drivers, model.Driver{1, "Lewis Hamilton", 44, 1})
	Drivers = append(Drivers, model.Driver{2, "Valtteri Bottas", 77, 1})
	Drivers = append(Drivers, model.Driver{3, "Sebastian Vettel", 5, 2})
	Drivers = append(Drivers, model.Driver{4, "Charles Leclerc", 16, 2})
	Drivers = append(Drivers, model.Driver{5, "Max Verstappen", 33, 3})
	Drivers = append(Drivers, model.Driver{6, "Pierre Gasly", 10, 3})
	Drivers = append(Drivers, model.Driver{7, "Sergio Perez", 11, 4})
	Drivers = append(Drivers, model.Driver{8, "Lance Stroll", 18, 4})
	Drivers = append(Drivers, model.Driver{9, "George Russell", 63, 5})
	Drivers = append(Drivers, model.Driver{10, "Robert Kubica", 88, 5})
	Drivers = append(Drivers, model.Driver{11, "Nico Hulkenberg", 27, 6})
	Drivers = append(Drivers, model.Driver{12, "Daniel Ricciardo", 3, 6})
	Drivers = append(Drivers, model.Driver{13, "Romain Grosjean", 8, 7})
	Drivers = append(Drivers, model.Driver{14, "Kevin Magnussen", 20, 7})
	Drivers = append(Drivers, model.Driver{15, "Daniil Kvyat", 26, 8})
	Drivers = append(Drivers, model.Driver{16, "Alexander Albon", 23, 8})
	Drivers = append(Drivers, model.Driver{17, "Kimi Raikkonen", 7, 9})
	Drivers = append(Drivers, model.Driver{18, "Antonio Giovanazzi", 99, 8})
	Drivers = append(Drivers, model.Driver{19, "Lando Norris", 99, 10})
	Drivers = append(Drivers, model.Driver{20, "Carlos Sainz", 55, 10})

	return Drivers
}

//NOTE: This is a temporary data souce
func GetAllTeams() []model.Team {
	Teams := make([]model.Team, 0, 10)

	Teams = append(Teams, model.Team{1, "Mercedes", "Silver"})
	Teams = append(Teams, model.Team{2, "Ferrari", "Red"})
	Teams = append(Teams, model.Team{3, "Red Bull", "Navy"})
	Teams = append(Teams, model.Team{4, "Force India", "Pink"})
	Teams = append(Teams, model.Team{5, "Williams", "White"})
	Teams = append(Teams, model.Team{6, "Renault", "Yellow"})
	Teams = append(Teams, model.Team{7, "Haas", "White"})
	Teams = append(Teams, model.Team{8, "Toro Rosso", "Blue"})
	Teams = append(Teams, model.Team{9, "Sauber", "White"})
	Teams = append(Teams, model.Team{10, "McLaren", "Orange"})

	return Teams
}
