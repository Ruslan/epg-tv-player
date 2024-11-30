package main

import (
	"context"
	"fmt"

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
