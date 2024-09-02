package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

var ticker int

func main() {
	ticker = 0
	drawRectangle()
}

func tidyUp() {
	fmt.Println("Exit!")
}

func window2() {
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

func updateTime(clock *widget.Label) {
	format := time.Now().Format("2006-01-02 15:04:05")
	clock.SetText(format)
}

func addInput() {
	a := app.New()
	w := a.NewWindow("Hello Person")

	w.SetContent(container.NewVBox(makeUI()))
	w.ShowAndRun()
}

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello motherfucker")
	in := widget.NewEntry()

	in.OnChanged = func(s string) {
		out.SetText("Hello " + s + "!")
	}

	return out, in
}

func drawRectangle() {
	appx := app.New()
	myWindow := appx.NewWindow("draw")
	myCanvas := myWindow.Canvas()

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	rect := canvas.NewRectangle(blue)
	myCanvas.SetContent(rect)

	for i := 0; i < 10; i++ {
		go func(myCanvas fyne.Canvas) {
			time.Sleep(1 * time.Second)
			//green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
			//rect.FillColor = green
			if ticker%2 == 0 {
				setContext(myCanvas)
			} else {
				setContentToCircle(myCanvas)
			}
			ticker++
			rect.Refresh()
		}(myCanvas)
		time.Sleep(1 * time.Second)
	}

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}

func setContext(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(red)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}
