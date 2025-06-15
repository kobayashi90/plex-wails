package main

import "github.com/webview/webview"

func main() {
    debug := true
    w := webview.New(debug)
    defer w.Destroy()

    w.SetTitle("Discord")
    w.SetSize(1280, 800, webview.HintNone)
    w.Navigate("https://discord.com/app")
    w.Run()
}
