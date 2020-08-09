package migration

import (
	"github.com/ybkuroki/go-webapp-project-template/config"
	"github.com/ybkuroki/go-webapp-project-template/model"
	"github.com/ybkuroki/go-webapp-project-template/repository"
)

// CreateDatabase creates the tables used in this application.
func CreateDatabase(config *config.Config) {
	if config.Database.Migration {
		db := repository.GetDB()

		db.DropTableIfExists(&model.Account{})
		db.DropTableIfExists(&model.Authority{})

		db.AutoMigrate(&model.Account{})
		db.AutoMigrate(&model.Authority{})
	}
}
