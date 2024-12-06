package main

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}
	err = db.AutoMigrate(&Video{})
	db.AutoMigrate(&SettingsApp{})
	if err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}
	videos := []Video{
		{ID: 1, TitleLower: "test video one", DescLower: "description one", Start: time.Time{}},
		{ID: 2, TitleLower: "test video two", DescLower: "description two", Start: time.Time{}},
		{ID: 3, TitleLower: "another video", DescLower: "another description", Start: time.Time{}},
	}
	if err := db.Create(&videos).Error; err != nil {
		t.Fatalf("Failed to seed test data: %v", err)
	}
	return db
}
func TestFetchVideos(t *testing.T) {
	db := setupTestDB(t)
	app := NewApp(db)

	tests := []struct {
		name     string
		req      VideoRequest
		query    string
		expected int
	}{
		{"Query Matches Title", VideoRequest{Page: 0, PerPage: 2}, "test", 2},
		{"Query Matches Description", VideoRequest{Page: 0, PerPage: 2}, "description", 2},
		{"Query Matches Nothing", VideoRequest{Page: 0, PerPage: 2}, "no-match", 0},
		{"Pagination Test1", VideoRequest{Page: 0, PerPage: 1}, "test", 1},
		{"Pagination Test2", VideoRequest{Page: 1, PerPage: 1}, "test", 1},
		{"Pagination Test3", VideoRequest{Page: 2, PerPage: 1}, "test", 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			videos, err := app.FetchVideos(tc.req, tc.query)
			assert.NoError(t, err)
			assert.Len(t, *videos, tc.expected)
		})
	}
}
func TestSetGetSetting(t *testing.T) {
	db := setupTestDB(t)
	app := NewApp(db)
	tests := []struct {
		name     string
		key      string
		val      string
		expected string
	}{
		{"set and get setting", "theme", "black", "black"},
		{"update and get setting", "theme", "white", "white"},
	}
	for _, tc := range tests {
		app.SetSetting(tc.key, tc.val)
		assert.Equal(t, tc.expected, app.GetSetting(tc.key))
	}
}
