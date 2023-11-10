package orm

import (
	"fmt"
	"gorm.io/gorm"
	"gormhasmany/utils"
	"math/rand"
)

type Owner struct {
	gorm.Model
	Name string
	Pets []Pet
}

type Pet struct {
	gorm.Model
	Name    string
	OwnerID uint
}

func CreateRandomOwner() *Owner {
	ownerName := utils.GetRandomName()
	petsNum := rand.Intn(3)
	newOwner := &Owner{
		Name: ownerName,
		Pets: make([]Pet, petsNum),
	}
	for i := 0; i < petsNum; i++ {
		newOwner.Pets[i] = Pet{Name: utils.GetRandomName()}
	}
	return newOwner
}

func PrintOwner(owner *Owner) {
	fmt.Printf("id: %d\nOwner name: %s\nPets: ", owner.Model.ID, owner.Name)
	for i, pet := range owner.Pets {
		fmt.Print(pet.Name)
		if i < len(owner.Pets)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
	fmt.Println()
}
