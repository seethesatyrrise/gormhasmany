package orm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type orm struct {
	db *gorm.DB
}

func New() *orm {
	return &orm{}
}

func (o *orm) Open() error {
	cfg := getDefaultConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Etc/UTC",
		cfg.host, cfg.user, cfg.password, cfg.dbname, cfg.port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	o.db = db
	return nil
}

func (o *orm) Migrate() error {
	return o.db.AutoMigrate(&Owner{}, &Pet{})
}

func (o *orm) CreateRandom() {
	o.db.Create(CreateRandomOwner())
}

func (o *orm) GetAndPrintAllOwners() error {
	var owners []Owner
	err := o.db.Model(&Owner{}).Preload("Pets").Find(&owners).Error
	if err != nil {
		return err
	}
	for _, owner := range owners {
		PrintOwner(&owner)
	}
	return nil
}
