package main

import (
	"embed"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	// app := NewApp()

	setupDb()
	loader := &DbEpg{db: db}
	loader.LoadAndParse()

	// // Create application with options
	// err := wails.Run(&options.App{
	// 	Title:  "iptv",
	// 	Width:  1024,
	// 	Height: 768,
	// 	AssetServer: &assetserver.Options{
	// 		Assets: assets,
	// 	},
	// 	BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
	// 	OnStartup:        app.startup,
	// 	Bind: []interface{}{
	// 		app,
	// 	},
	// })

	// if err != nil {
	// 	println("Error:", err.Error())
	// }
}
