package main

import (
	"context"
	"errors"
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
func (app *App) FetchVideos(req VideoRequest, q string, chq string) (*[]Video, error) {
	videos, err := getVideosByQuery(app.DB, q, chq, &req)
	if err != nil {
		return nil, err
	}
	return videos, nil

}
func getVideosByQuery(db *gorm.DB, q string, chq string, videoRequest *VideoRequest) (*[]Video, error) {
	var videos []Video
	query := strings.ToLower(q)
	query = strings.ReplaceAll(query, " ", "%")
	searchQuery := fmt.Sprintf("%%%s%%", query)

	baseScope := db.Where("title_lower LIKE ? OR desc_lower LIKE ?", searchQuery, searchQuery)

	if len(chq) > 0 {
		chq = strings.ToLower(chq)
		chq = fmt.Sprintf("%%%s%%", chq)
		baseScope = baseScope.Where("channel_code like ?", chq)
	}

	err := baseScope.
		Order("start ASC").
		Limit(videoRequest.PerPage).
		Offset((videoRequest.Page - 1) * videoRequest.PerPage).
		Find(&videos).Error
	return &videos, err
}
func (a *App) SetSetting(key string, value string) {
	var val SettingsApp
	err := a.DB.Where("key = ?", key).First(&val).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		var set SettingsApp
		set.Value = value
		set.Key = key
		a.DB.Create(&set)
		return
	} else if err != nil {
		panic(err)
	}
	a.DB.Model(&SettingsApp{}).Where("key = ?", key).Update("value", value)
}
func (a *App) GetSetting(key string) string {
	var val SettingsApp
	a.DB.Where("key = ?", key).First(&val)
	return val.Value
}
