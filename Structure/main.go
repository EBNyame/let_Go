package main

import "fmt"

type Team struct {
	Name     string
	Location string
	Trophies int
	Owner
}

type Owner struct {
	Name     string
	NetWorth int
}

func main() {
	manUtd := &Team{
		Name:     "Manchester United",
		Location: "Manchester, england",
		Trophies: 44444,

		Owner: Owner{
			Name:     "Exodus",
			NetWorth: 55433345673678,
		},
	}

	updateTrophies(manUtd)

	fmt.Println(manUtd)

	realMadrid := newTeam("Real Madrid", "Madrid, Spain", 3332)
	fmt.Println(realMadrid)
}

func updateTrophies(x *Team) {
	x.Trophies += 1
}

// Constructors

func newTeam(name string, location string, trophies int) Team {
	return Team{
		Name:     name,
		Location: location,
		Trophies: trophies,
	}
}
