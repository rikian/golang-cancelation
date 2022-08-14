package migrations

import (
	"go/cancel/app/entity"
	"log"

	"gorm.io/gorm"
)

type table struct {
	tableName interface{}
}

func registerTable() []table {
	return []table{
		{tableName: entity.Product{}},
	}
}

func RunMigration(db *gorm.DB) error {
	for _, table := range registerTable() {
		err := db.Migrator().CreateTable(table.tableName)

		if err != nil {
			log.Println("Migrations Failed...")
			log.Print(err)
			break
		}
	}

	log.Println("Migrations successfully ...")
	return nil
}
