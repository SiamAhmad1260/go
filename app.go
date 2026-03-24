package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 1. Create app (always first)
	myApp := app.New()

	// 2. Create window
	myWindow := myApp.NewWindow("My First App")

	// 3. Create widgets (UI elements)
	hello := widget.NewLabel("Hello World!")

	// 4. Add widget to window
	myWindow.SetContent(hello)

	// 5. Show window and run app
	myWindow.ShowAndRun()
}
