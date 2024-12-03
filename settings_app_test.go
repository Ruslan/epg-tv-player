package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupSettingsTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	if err := db.AutoMigrate(&SettingsApp{}); err != nil {
		panic(err)
	}
	return db
}

func TestAddDefaultSettings(t *testing.T) {
	db := setupSettingsTestDB()
	AddDefaultSettings(db, "key1", "key2", "key3")
	var settings []SettingsApp
	result := db.Find(&settings)
	assert.NoError(t, result.Error)
	assert.Len(t, settings, 3)
	for _, setting := range settings {
		assert.Equal(t, "default", setting.Value)
	}
}

func TestAddNewSettings(t *testing.T) {
	db := setupSettingsTestDB()
	AddDefaultSettings(db, "key1", "key2")
	AddNewSettings("key1", "updated_value", db)
	var updatedSetting SettingsApp
	result := db.Where("key = ?", "key1").First(&updatedSetting)
	assert.NoError(t, result.Error)
	assert.Equal(t, "updated_value", updatedSetting.Value)
	var defaultSetting SettingsApp
	result = db.Where("key = ?", "key2").First(&defaultSetting)
	assert.NoError(t, result.Error)
	assert.Equal(t, "default", defaultSetting.Value)
}

func TestAddNewSettingsForNonExistingKey(t *testing.T) {
	db := setupSettingsTestDB()
	AddNewSettings("non_existing_key", "new_value", db)
	var nonExistingSetting SettingsApp
	result := db.Where("key = ?", "non_existing_key").First(&nonExistingSetting)
	assert.Error(t, result.Error) // Должна быть ошибка записи не найдено
	assert.Equal(t, int64(0), result.RowsAffected)
}
