package db

import "medico/data"

func Migrate(repository Repository) {
	if err := repository.DropTableIfExists(data.CommonUserDB{}); err != nil {
		return
	}

	if err := repository.AutoMigrate(data.CommonUserDB{}); err != nil {
		return
	}
}
