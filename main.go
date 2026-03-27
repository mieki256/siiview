// -*- mode: go; Encoding: utf-8; coding: utf-8 -*-
// Last updated: <2026/03/27 20:42:29 +0900>
//
// 標準入力から受け取った画像をウインドウ表示する。
// fyneを使ってウインドウ表示をしている。
// png/bmp/jpeg/gif/tiff に対応。
// Google Gemini に作成してもらった。
//
// Usage: gmic sp lena o -.png | siiview.exe
// ESC / Q key to exit
//
// Windows11 x64 25H2 + Go 1.25.7 64bit

package main

import (
	"image"
	"log"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	stat, _ := os.Stdin.Stat()

	// Check if data is being piped to stdin
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		log.Println("[siiview] No image data piped to stdin. Usage: cat image.png | siiview")
		return
	}

	// Try to decode image from stdin
	img, format, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatalf("[siiview] Failed to decode image: %v", err)
	}
	log.Printf("[siiview] Loaded image format: %s", format)

	// Create window by fyne
	myApp := app.New()
	myWindow := myApp.NewWindow("siiview - Stdin Image Viewer : ESC key to exit")

	// ESC / Q key to exit
	myWindow.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == fyne.KeyEscape || k.Name == fyne.KeyQ {
			myApp.Quit()
		}
	})

	imgSize := img.Bounds().Size()
	w := float32(imgSize.X)
	h := float32(imgSize.Y)

	imageWidget := canvas.NewImageFromImage(img)
	imageWidget.FillMode = canvas.ImageFillContain
	// imageWidget.FillMode = canvas.ImageFillOriginal

	myWindow.SetContent(imageWidget)
	myWindow.Resize(fyne.NewSize(w, h))
	// myWindow.Resize(fyne.NewSize(800, 600))

	myWindow.ShowAndRun()
}
