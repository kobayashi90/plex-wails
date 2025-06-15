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
    isFullscreen bool
}

func NewApp() *App {
    return &App{
        isFullscreen: false,
    }
}

func (a *App) OnStartup(ctx context.Context) {
    a.ctx = ctx
}

func (a *App) SetFullscreen(fullscreen bool) {
    if fullscreen {
        runtime.WindowFullscreen(a.ctx)
        a.isFullscreen = true
    } else {
        runtime.WindowUnfullscreen(a.ctx)
        a.isFullscreen = false
    }
}

func (a *App) IsFullscreen() bool {
    // Use internal state as runtime.WindowIsFullscreen might not be reliable
    return a.isFullscreen
}

func (a *App) ToggleFullscreen() {
    if a.isFullscreen {
        a.SetFullscreen(false)
    } else {
        a.SetFullscreen(true)
    }
}

// Minimize window
func (a *App) MinimizeWindow() {
    runtime.WindowMinimise(a.ctx)
}

// Close application
func (a *App) CloseApplication() {
    runtime.Quit(a.ctx)
}

// Restore window from minimized state
func (a *App) RestoreWindow() {
    runtime.WindowUnminimise(a.ctx)
}

// Hide window
func (a *App) HideWindow() {
    runtime.WindowHide(a.ctx)
}

// Show window
func (a *App) ShowWindow() {
    runtime.WindowShow(a.ctx)
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
        OnStartup: app.OnStartup,
        // Ensure the window can go fullscreen
        DisableResize: false,
        Fullscreen:    true,
    })

    if err != nil {
        println("Error:", err.Error())
    }
}
