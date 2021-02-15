package model

// Clues is a struct of persons, weapon, places
type Clues struct {
	Persons []string `json:"persons"`
	Weapons []string `json:"weapons"`
	Places  []string `json:"places"`
}
