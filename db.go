package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Channel struct {
	ID             uint
	Title          string
	Position       int
	Url            string
	GroupTitle     string
	TvgID          string
	TvgLogo        string
	TvgRec         string
	catchup_source string
	catchup_days   int
}

type Video struct {
	ID          uint
	Title       string
	TitleLower  string
	Desc        string
	DescLower   string
	Start       time.Time `gorm:"index"`
	Stop        time.Time `gorm:"index"`
	ChannelId   uint
	Channel     Channel
	ChannelCode string `gorm:"index"`
}
type SettingsApp struct {
	id    int    `gorm:"primaryKey"`
	Key   string `gorm:"type:text"`
	Value string `gorm:"type:text"`
}

var db *gorm.DB

func setupDb() {
	// Connect to SQLite database
	var err error
	db, err = gorm.Open(sqlite.Open("tv.db"), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate
	db.AutoMigrate(&Channel{})
	db.AutoMigrate(&Video{})
	db.AutoMigrate(&SettingsApp{})
	var channelNames []string
	result := db.Model(&Channel{}).Pluck("title", &channelNames)

	// Check for errors
	if result.Error != nil {
		fmt.Println("Error occurred:", result.Error)
		return
	}

	var videoCount int64
	db.Model(&Video{}).Count(&videoCount)

	// Check for errors
	if result.Error != nil {
		fmt.Println("Error occurred:", result.Error)
		return
	}

	fmt.Println("Channels in DB:", len(channelNames), "Videos in DB:", videoCount)
}
