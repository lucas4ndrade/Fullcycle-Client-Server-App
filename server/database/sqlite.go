package database

import (
	"context"
	"time"

	"github.com/go-fullcicle/server/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db = &gorm.DB{}

func StartDatabase() {
	dial := sqlite.Open(config.GetConfig().DBPath)
	openedDB, err := gorm.Open(dial)
	if err != nil {
		panic(err)
	}

	db = openedDB

	err = db.AutoMigrate(&Bid{})
	if err != nil {
		panic("failed to migrate database")
	}
}

func SaveBid(val string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	b := Bid{Value: val}

	return db.WithContext(ctx).Create(&b).Error
}
