package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB(host, username, password, database, port string) (string, bool) {
	var err error

	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})

	if err != nil {
		return err.Error(), true
	}

	return "Postgres server listening on port " + port + "...\n", false
}
