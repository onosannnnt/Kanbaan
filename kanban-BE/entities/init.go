package entities

import "gorm.io/gorm"

func InitEntities(db *gorm.DB) {
	if err := 	db.AutoMigrate(
		&User{},
		&Board{},
		&Column{},
		&Task{},
		&Notification{},
	); err != nil {
		panic(err)
	}
}
