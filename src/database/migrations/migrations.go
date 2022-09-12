package migrations

import (
	"go-bitly/src/database"
	"go-bitly/src/database/entities"
)

func RunMigrations() {
	err := database.Db.AutoMigrate(&entities.ShortURLEntity{})
	if err != nil {
		panic(err)
	}
}
