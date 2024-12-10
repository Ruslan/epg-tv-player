package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SettingsApp struct {
	id    int    `gorm:"primaryKey"`
	Key   string `gorm:"type:text"`
	Value string `gorm:"type:text"`
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("settings.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&SettingsApp{}); err != nil {
		panic(err)
	}
	return db
}
func AddDefaultSettings(db *gorm.DB, keys ...string) {
	for _, key := range keys {
		db.Create(&SettingsApp{Key: key, Value: "default"})
	}
}
func SetSettings(key string, value string, db *gorm.DB) {
	db.Model(SettingsApp{}).Where("key = ?", key).Update("value", value)
}
func GetSettings(key string, db *gorm.DB) string {
	var set SettingsApp
	db.Where("key = ?", key).First(&set)
	return set.Value
}
