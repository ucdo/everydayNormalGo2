package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	appx := app.New()
	window := appx.NewWindow("Hello") // 这个是窗口的title
	window.Resize(fyne.NewSize(300, 300))
	clock := widget.NewLabel("")
	window.SetContent(clock)
	//format := time.Now().Format("2006-01-02 15:04:05")
	//clock.SetText(format)
	updateTime(clock)

	window.SetContent(clock)
	// appx.Run()后面的代码不会执行。如果是windows。  appx.Quit()会执行
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	window2 := appx.NewWindow("larger")
	window2.Resize(fyne.NewSize(100, 500))
	//window2.SetContent(widget.NewButton("open new", func() {
	//	window3 := appx.NewWindow("windows3")
	//	window3.SetContent(widget.NewLabel("third"))
	//	window3.Show()
	//}))

	btn := widget.NewButton("open new", func() {
		window3 := appx.NewWindow("windows3")
		window3.SetContent(widget.NewLabel("third"))
		window3.Show()
	})
	btn.Resize(fyne.NewSize(5, 5))

	window2.SetContent(btn)

	window.Show()
	window2.Show()
	appx.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exit!")
}

func updateTime(clock *widget.Label) {
	format := time.Now().Format("2006-01-02 15:04:05")
	clock.SetText(format)
}
