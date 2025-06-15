package main

import (
    "context"
    "embed"
    "github.com/wailsapp/wails/v2"
    "github.com/wailsapp/wails/v2/pkg/options"
    "github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type App struct {
    ctx context.Context
}

func NewApp() *App {
    return &App{}
}

func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
}

func main() {
    app := NewApp()

    err := wails.Run(&options.App{
        Title:  "My Plex Server",
        Width:  1280,
        Height: 720,
        URL:    "https://DEINE-PLEX-URL.com", // ← HIER ÄNDERN
        AssetServer: &assetserver.Options{
            Assets: assets,
        },
        OnStartup: app.startup,
    })

    if err != nil {
        println("Error:", err.Error())
    }
}
