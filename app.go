package main

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	DB  *gorm.DB
}

// NewApp creates a new App application struct
func NewApp(db *gorm.DB) *App {
	return &App{DB: db}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type ChannelResponse struct {
	Logo  string `json:"logo"`
	Title string `json:"title"`
	Group string `json:"group"`
	ID    uint   `json:"id"`
	TvgID string `json:"tvg_id"`
}
type VideoRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

func (app *App) FetchChannels() (map[string]interface{}, error) {
	var channels []Channel
	var channelsResponse []ChannelResponse
	var videoCount int64

	// Fetch channels from the database
	if err := app.DB.Order("position ASC").Find(&channels).Error; err != nil {
		return nil, err
	}

	// Map to ChannelResponse
	for _, channel := range channels {
		channelsResponse = append(channelsResponse, ChannelResponse{
			Logo:  channel.TvgLogo,
			Title: channel.Title,
			Group: channel.GroupTitle,
			ID:    channel.ID,
			TvgID: channel.TvgID,
		})
	}

	// Count videos
	if err := app.DB.Model(&Video{}).Count(&videoCount).Error; err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"channels":    channelsResponse,
		"totalVideos": videoCount,
	}, nil
}
func (app *App) FetchVideos(req VideoRequest, q string) (*[]Video, error) {
	videos, err := getVideosByQuery(app.DB, q, &req)
	if err != nil {
		return nil, err
	}
	return videos, nil

}
func getVideosByQuery(db *gorm.DB, q string, videoRequest *VideoRequest) (*[]Video, error) {
	var videos []Video
	query := strings.ToLower(q)
	searchQuery := fmt.Sprintf("%%%s%%", query)
	err := db.Where("lower(title_lower) LIKE ? OR lower(desc_lower) LIKE ?", searchQuery, searchQuery).
		Order("start ASC").
		Limit(videoRequest.PerPage).
		Offset(videoRequest.Page * videoRequest.PerPage).
		Find(&videos).Error
	return &videos, err
}
func (a *App) SetSetting(key string, value string) {
	var val SettingsApp
	a.DB.Where("key = ?", key).First(&val)
	if val.Value == "" {
		var set SettingsApp
		set.Value = value
		set.Key = key
		a.DB.Create(&set)
	} else {
		a.DB.Model(&SettingsApp{}).Where("key = ?", key).Update("value", value)
	}
}
func (a *App) GetSetting(key string) string {
	var val SettingsApp
	a.DB.Where("key = ?", key).First(&val)
	return val.Value
}
