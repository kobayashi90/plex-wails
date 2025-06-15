package main

import (
    "context"
    "embed"
    "github.com/wailsapp/wails/v2"
    "github.com/wailsapp/wails/v2/pkg/options"
    "github.com/wailsapp/wails/v2/pkg/options/assetserver"
    "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

type App struct {
    ctx context.Context
}

func NewApp() *App {
    return &App{}
}

// Deine bestehende startup Funktion (umbenennen zu OnStartup)
func (a *App) OnStartup(ctx context.Context) {
    a.ctx = ctx
}

// NEUE Fullscreen-Funktionen hinzufügen:
func (a *App) SetFullscreen(fullscreen bool) {
    if fullscreen {
        runtime.WindowFullscreen(a.ctx)
    } else {
        runtime.WindowUnfullscreen(a.ctx)
    }
}

func (a *App) IsFullscreen() bool {
    return runtime.WindowIsFullscreen(a.ctx)
}

func main() {
    app := NewApp()

    err := wails.Run(&options.App{
        Title:  "Plex Desktop",
        Width:  1280,
        Height: 720,
        AssetServer: &assetserver.Options{
            Assets: assets,
        },
        OnStartup: app.OnStartup, // Hier geändert von app.startup zu app.OnStartup
    })

    if err != nil {
        println("Error:", err.Error())
    }
}
