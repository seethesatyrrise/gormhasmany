package utils

import "math/rand"

var names = []string{"Rodney", "Schaefer", "Brennen", "Mooney", "Delilah", "Delgado", "Alison",
	"Curry", "Ruben", "Valdivia", "Ashlin", "Seward", "Jarret", "Braun",
	"Edgardo", "Wayne", "Jameson", "Heaton", "Guadalupe", "Scruggs"}

func GetRandomName() string {
	return names[rand.Intn(20)]
}
