package main

import (
	"fmt"
	"gormhasmany/orm"
)

func main() {
	orm := orm.New()

	if err := orm.Open(); err != nil {
		fmt.Println("opening error: ", err)
		return
	}
	fmt.Println("orm opened")

	if err := orm.Migrate(); err != nil {
		fmt.Println("migration error: ", err)
		return
	}
	fmt.Println("automigration done")
	fmt.Println()

	for i := 0; i < 3; i++ {
		orm.CreateRandom()
	}

	if err := orm.GetAndPrintAllOwners(); err != nil {
		fmt.Println("printall error: ", err)
		return
	}
}
